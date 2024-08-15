package scan

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"net"
	"os"
	"strings"
	"sync"
	"time"
	"udpz/pkg/data"

	"github.com/jedib0t/go-pretty/v6/progress"
)

type Host struct {
	Addresses []net.IP
	Target    string
}

type UdpProbeScanner struct {
	Concurrency     uint
	ProbeCount      int
	Retransmissions int
	ReadTimeout     time.Duration
	ProgressWriter  progress.Writer
	NoLog           bool
	Output          *os.File

	tasks       chan struct{}
	resultsLive chan PortResult
	results     []PortResult
	errors      chan error
}

type PortResult struct {
	Port      uint16        `yaml:"port" json:"port"`
	Transport string        `yaml:"transport" json:"transport"`
	Target    string        `yaml:"host" json:"host"`
	Response  string        `yaml:"response" json:"response"`
	Service   *data.Service `yaml:"service" json:"service"`
}

type ScanTaskRequest struct {
	Target    string
	Transport string
	Port      int
}

func (sc *UdpProbeScanner) handleResult(pr PortResult) {
	if !sc.NoLog {
		sc.logPortResult(pr)
	}
	sc.results = append(sc.results, pr)
}

func (sc *UdpProbeScanner) Length() int {
	return len(sc.results)
}

// TODO: ctx
func (sc *UdpProbeScanner) scanTask(target string, port uint16, payload []byte, service *data.Service) (result PortResult, err error) {

	var conn net.Conn
	var readLen int

	if conn, err = net.Dial("udp", fmt.Sprintf("%s:%d", target, port)); err == nil {
		if err = conn.SetReadDeadline(time.Now().Add(sc.ReadTimeout)); err == nil {

			response := make([]byte, 0x800)
			conn.Write(payload)

			if readLen, err = bufio.NewReader(conn).Read(response); err == nil {
				result = PortResult{
					Port:      port,
					Transport: "udp",
					Target:    target,
					Service:   service,
				}
				if readLen > 0 {
					result.Response = base64.StdEncoding.EncodeToString(response[:readLen])
				}
			}
		}
	}

	return
}

func (sc *UdpProbeScanner) Scan(targets []string) {

	var wg sync.WaitGroup

	sem := make(chan struct{}, sc.Concurrency)
	total := len(targets) * data.NUM_PAYLOADS * (1 + sc.Retransmissions)

	tracker := &progress.Tracker{
		Message:          "Probing targets",
		DeferStart:       false,
		ExpectedDuration: 4 * time.Second,
		Total:            int64(total),
		Units:            progress.UnitsDefault,
	}
	sc.ProgressWriter.AppendTracker(tracker)

	go func() {
		sm := make(map[string]map[uint16]bool)
		for r := range sc.resultsLive {

			if _, ok := sm[r.Target]; !ok {
				sm[r.Target] = make(map[uint16]bool)
			} else if sm[r.Target][r.Port] {
				continue // Ignore duplicates
			}
			sc.handleResult(r)
			sm[r.Target][r.Port] = true
		}
	}()

	for _, target := range targets {
		sc.logInfo("Scanning target: %s", target)

		if addr := net.ParseIP(target); addr != nil && addr.To4() == nil {
			target = fmt.Sprintf("[%s]", target)
		}

		for port, probes := range data.PROBES_ALL {
			var portStatus uint8

			for _, probe := range probes {
				for _, payload := range probe.Payloads {

					sem <- struct{}{}
					wg.Add(1)

					go func(target string, port uint16, payload []byte, service *data.Service, portStatus *uint8, tracker *progress.Tracker) {

						defer func() {
							wg.Done()
							<-sem
						}()

						for i := 0; i <= sc.Retransmissions; i++ {
							if *portStatus != STATE_UNRESPONSIVE {
								tracker.Increment(1)
							}
							if result, err := sc.scanTask(target, port, payload, service); err != nil {
								if strings.Contains(err.Error(), "connection refused") {
									//tracker.Increment(int64(sc.Retransmissions - i + 1))
									*portStatus = STATE_CLOSED
									break
								} else if strings.Contains(err.Error(), "i/o timeout") {
									tracker.Increment(1)
								} else {
									tracker.IncrementWithError(1)
									sc.errors <- err
								}
							} else {
								tracker.Increment(1)
								sc.resultsLive <- result
								*portStatus = STATE_RESPONSIVE
							}
						}
					}(target, port, payload, &probe.Service, &portStatus, tracker)
				}
			}
		}
	}
	wg.Wait()
	for !tracker.IsDone() {
		time.Sleep(5 * time.Millisecond)
	}
}

func NewUdpProbeScanner(concurrency uint, retransmissions uint, readTimeout time.Duration,
	progressWriter progress.Writer) (sc UdpProbeScanner) {

	sc.tasks = make(chan struct{}, concurrency)
	sc.resultsLive = make(chan PortResult)

	sc.Concurrency = concurrency
	sc.ReadTimeout = readTimeout
	sc.ProgressWriter = progressWriter
	sc.Retransmissions = int(retransmissions)
	return
}
