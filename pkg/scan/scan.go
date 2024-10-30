package scan

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"net"
	"strings"
	"sync"
	"time"
	"udpz/pkg/data"

	"github.com/rs/zerolog"
	"github.com/txthinking/socks5"
)

func (sc *UdpProbeScanner) handleResult(pr PortResult) {

	sc.Logger.Debug().
		Str("target", pr.Host.Target.Target).
		Str("host", fmt.Sprintf("%s:%s", pr.Host.Type, pr.Host.Host)).
		Int("port", int(pr.Port)).
		Str("service", pr.Service.Slug).
		Str("probe", pr.Probe.Slug).
		Str("response", pr.Response).
		Msg("Received response")

	if _, ok := sc.resultsMap[pr.Host.Host]; !ok {
		sc.resultsMap[pr.Host.Host] = make(map[uint16][]PortResult)
	}
	if _, ok := sc.resultsMap[pr.Host.Host][pr.Port]; !ok {
		sc.Logger.Info().
			Str("target", pr.Host.Target.Target).
			Str("host", fmt.Sprintf("%s:%s", pr.Host.Type, pr.Host.Host)).
			Int("port", int(pr.Port)).
			Str("service", pr.Service.Slug).
			Str("probe", pr.Probe.Slug).
			Msg("Discovered UDP service")

		sc.results = append(sc.results, pr)
		sc.resultsMap[pr.Host.Host][pr.Port] = []PortResult{pr}
	} else {
		sc.resultsMap[pr.Host.Host][pr.Port] = append(sc.resultsMap[pr.Host.Host][pr.Port], pr)
	}
}

func (sc *UdpProbeScanner) Length() int {
	return len(sc.results)
}

// TODO: ctx
func (sc *UdpProbeScanner) scanTask(host Host, port uint16, payload []byte) (result PortResult, err error) {

	sc.Logger.Trace().
		Str("type", "call").
		Str("function", "(*UdpProbeScanner).scanTask").
		Dict("arguments", zerolog.Dict().
			Interface("host", host).
			Uint16("port", port).
			Bytes("payload", payload)).
		Msg("(*UdpProbeScanner).scanTask(...)")

	var conn net.Conn
	var readLen int

	for {
		transport := "udp"
		address := fmt.Sprintf("%s:%d", host.Host, port)

		if sc.useProxy {
			sc.Logger.Trace().
				Str("type", "(*socks5.Client).Dial").
				Str("proxy", sc.proxy.Server). //sc.proxy.Server).
				Str("transport", transport).
				Str("address", address).
				Msg("(*socks5.Client).Dial(...)")

			conn, err = sc.proxy.Dial(transport, address)
		} else {
			sc.Logger.Trace().
				Str("type", "net.Dial").
				Str("transport", transport).
				Str("address", address).
				Msg("net.Dial(...)")

			conn, err = net.Dial(transport, address)
		}
		if err == nil {
			if err = conn.SetReadDeadline(time.Now().Add(sc.ReadTimeout)); err == nil {

				response := make([]byte, RESPONSE_MAX_LEN)

				sc.Logger.Trace().
					Str("type", "(net.).write").
					Str("transport", transport).
					Str("address", address).
					Bytes("data", payload).
					Msg("(net.Conn).Write(data)")

				conn.Write(payload)

				sc.Logger.Trace().
					Str("type", "connection.read").
					Str("transport", transport).
					Str("address", address).
					Bytes("data", response).
					Msg("(net.Conn).Read(data)")

				readLen, err = bufio.NewReader(conn).Read(response)

				sc.Logger.Trace().
					Str("type", "connection.close").
					Str("transport", transport).
					Str("address", address).
					Bytes("data", response).
					Msg("(net.Conn).Close()")

				if terr := conn.Close(); terr != nil {
					sc.Logger.Trace().
						Err(terr).
						Str("transport", transport).
						Str("address", address).
						Msg("Failed to close connection")
				}

				if err == nil {

					result = PortResult{
						Port:      port,
						Transport: transport,
						Host:      host,
					}

					if readLen > 0 {
						result.Response = base64.StdEncoding.EncodeToString(response[:readLen])
					}
				}
			}
			break
		} else {
			if strings.Contains(err.Error(), "connect: resource temporarily unavailable") {
				sc.Logger.Debug().
					Err(err).
					Msg("Resource limit reached")
				time.Sleep(10 * time.Millisecond)
				continue
			}
			break
		}
	}

	sc.Logger.Trace().
		Str("type", "return").
		Str("function", "(*UdpProbeScanner).scanTask").
		Dict("arguments", zerolog.Dict().
			Interface("host", host).
			Uint16("port", port).
			Bytes("payload", payload)).
		Dict("return", zerolog.Dict().
			AnErr("err", err).
			Interface("result", result)).
		Msg("return <- (*UdpProbeScanner).scanTask()")
	return
}

func (sc *UdpProbeScanner) Scan(targetSourceList []string) {

	sc.Logger.Trace().
		Str("type", "call").
		Str("function", "(*UdpProbeScanner).Scan").
		Dict("arguments", zerolog.Dict().
			Strs("targetSourceList", targetSourceList)).
		Msg("(*UdpProbeScanner).Scan(...)")

	var hostWg sync.WaitGroup
	//var portWg sync.WaitGroup
	var probeCount uint
	var totalCount uint

	hostSem := make(chan struct{}, sc.HostConcurrency)
	//portSem := make(chan struct{}, sc.PortConcurrency)
	hosts := make(chan Host)

	sc.Logger.Debug().
		Int("service_count", len(data.UDP_SERVICES)).
		Msg("Calculating unique probe count")

	for _, service := range data.UDP_SERVICES {
		probeCount += uint(len(service.Ports) * len(service.Probes))
	}
	sc.Logger.Debug().
		Uint("probe_count", probeCount).
		Msg("Calculated unique probe count")

	totalCount = probeCount * (sc.Retransmissions + 1)

	sc.Logger.Debug().
		Uint("total_probes", totalCount).
		Msg("Calculated total probe count")

	go func(wg *sync.WaitGroup, c chan Host) {

		sc.Logger.Debug().
			Int("target_count", len(targetSourceList)).
			Msg("Resolving targets")

		for _, ts := range targetSourceList {
			sc.ResolveTarget(ts, c)
		}
		wg.Wait()
		close(c)

	}(&hostWg, hosts)

	go func() {
		for r := range sc.resultsLive {
			sc.handleResult(r)
		}
	}()

	for host := range hosts {

		host := host // Shadow variable
		portWg := sync.WaitGroup{}
		portSem := make(chan struct{}, sc.PortConcurrency)

		hostSem <- struct{}{}
		hostWg.Add(1)

		go func() {

			defer func() {
				hostWg.Done()
				<-hostSem
			}()

			for _, service := range data.UDP_SERVICES {
				for _, port := range service.Ports {

					var portStatus uint8 = STATE_UNRESPONSIVE

					for _, probe := range service.Probes {

						probe := probe

						portSem <- struct{}{}
						portWg.Add(1)

						go func(wg *sync.WaitGroup,
							h Host, port uint16, probe data.UdpProbe,
							service *data.UdpService, portStatus *uint8) {

							defer func() {
								wg.Done()
								<-portSem
							}()

							if probeBytes, err := base64.StdEncoding.DecodeString(probe.EncodedData); err == nil {
								for i := 0; i <= int(sc.Retransmissions); i++ {

									if *portStatus == STATE_CLOSED {

										sc.Logger.Debug().
											Str("target", h.Target.Target).
											Str("host", h.Host).
											Uint16("port", port).
											Msg("Skipping closed port")
									}

									if result, err := sc.scanTask(h, port, probeBytes); err != nil {

										if strings.Contains(err.Error(), "connection refused") {
											*portStatus = STATE_CLOSED

											sc.Logger.Debug().
												Str("target", h.Target.Target).
												Str("host", h.Host).
												Uint16("port", port).
												Msg("Port closed")

										} else if strings.Contains(err.Error(), "i/o timeout") {
											sc.Logger.Debug().
												Str("target", h.Target.Target).
												Str("host", h.Host).
												Uint16("port", port).
												Str("probe", probe.Slug).
												Msg("Port unresponsive")

										} else {
											sc.Logger.Error().
												Err(err).
												Str("target", h.Target.Target).
												Str("host", h.Host).
												Uint16("port", port).
												Msg("Error in scan task")
										}
									} else {
										service := data.UDP_SERVICES[probe.Service]
										result.Service = service.ToOutput()
										result.Probe = probe.ToOutput()

										sc.resultsLive <- result
										*portStatus = STATE_RESPONSIVE
									}
								}
							} else {
								sc.Logger.Error().
									Interface("probe", probe).
									Err(err).
									Msg("Failed to decode probe data")
							}
						}(&portWg, host, port, probe, &service, &portStatus)
					}
				}
			}
			portWg.Wait()
		}()
	}

	hostWg.Wait()
}

func NewUdpProbeScanner(logger zerolog.Logger, scanAllAddresses bool,
	hostConcurrency uint, portConcurrency uint, retransmissions uint, readTimeout time.Duration,
	socks5Address string, socks5User string, socks5Password string, socks5Timeout int) (sc UdpProbeScanner, err error) {

	// Assemble a new UdpProbeScanner using supplied arguments

	sc.resultsLive = make(chan PortResult)
	sc.resultsMap = make(map[string]map[uint16][]PortResult)

	sc.HostConcurrency = hostConcurrency
	sc.PortConcurrency = portConcurrency
	sc.Retransmissions = retransmissions
	sc.ReadTimeout = readTimeout

	sc.Logger = logger

	if socks5Address != "" {

		sc.useProxy = true
		sc.Logger.Debug().
			Str("address", socks5Address).
			Int("timeout", socks5Timeout).
			Msg("Using SOCKS5 proxy")

		if socks5User != "" || socks5Password != "" {
			sc.Logger.Debug().
				Str("address", socks5Address).
				Str("user", socks5User).
				Msg("Using SOCKS5 authentication")
		}

		if sc.proxy, err = socks5.NewClient(
			socks5Address, socks5User, socks5Password,
			socks5Timeout, socks5Timeout); err != nil {

			sc.Logger.Error().
				Err(err).
				Str("address", socks5Address).
				Msg("Failed to initialize SOCKS5 proxy dialer")

		} else if _, err := sc.proxy.Dial("udp", "127.0.0.1:1"); err != nil {
			// This shouldn't send any packets through the proxy, justs tests if its available.

			sc.Logger.Fatal().
				Err(err).
				Str("address", socks5Address).
				Msg("SOCKS5 connection failed")
		}

	}

	return
}
