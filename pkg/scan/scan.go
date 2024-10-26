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

func (sc *UdpProbeScanner) ResolveTarget(targetSource string, hosts chan Host) (err error) {

	sc.Logger.Trace().
		Str("type", "call").
		Str("function", "(*UdpProbeScanner).Scan").
		Dict("arguments", zerolog.Dict().
			Str("targetSource", targetSource).
			Interface("hosts", hosts)).
		Msg("(*UdpProbeScanner).Scan()")

	var target Target
	var host Host

	target.Target = targetSource

	if ip := net.ParseIP(targetSource); ip != nil {

		target.Type = "IP"
		host.Target = target

		if ip4 := ip.To4(); ip4 != nil {
			host.Type = "IPv4"
			host.Host = ip4.String()

		} else if ip16 := ip.To16(); ip16 != nil {
			host.Type = "IPv6"
			host.Host = fmt.Sprintf("[%s]", ip16)
		}
		sc.Logger.Debug().
			Str("type", target.Type).
			Str("target", target.Target).
			Str("address_type", host.Type).
			IPAddr("ip", ip).
			Msg("Target resolved")

		hosts <- host

	} else if ip, ipNet, err := net.ParseCIDR(targetSource); err == nil {

		var ipv6 bool

		target.Type = "CIDR"
		addrType := "Unknown"

		if ip4 := ip.To4(); ip4 != nil {
			addrType = "IPv4"
		} else if ip16 := ip.To16(); ip16 != nil {
			addrType = "IPv6"
			ipv6 = true
		}
		sc.Logger.Debug().
			Str("type", target.Type).
			Str("target", target.Target).
			Str("address_type", host.Type).
			IPAddr("ip", ip).
			Str("cidr", ipNet.String()).
			Msg("Target CIDR resolved")

		for ipNet.Contains(ip) {
			host = Host{
				Target: target,
				Type:   addrType,
				Host:   ip.String(),
				ip:     ip,
			}
			if ipv6 {
				host.Host = "[" + host.Host + "]" // Faster than fmt.Sprintf
			}
			hosts <- host

			for j := len(ip) - 1; j >= 0; j-- {
				ip[j]++
				if ip[j] != 0 {
					break
				}
			}
		}

	} else if REGEX_HOSTNAME.MatchString(targetSource) {

		target.Type = "hostname"

		// TODO: check if this handles hostnames correctly
		if ips, err := net.LookupIP(targetSource); err == nil {

			sc.Logger.Debug().
				Str("target", targetSource).
				Int("addresses", len(ips)).
				Msg("Resolved target hostname")

			for _, ip := range ips {

				host = Host{Target: target}

				if ip4 := ip.To4(); ip4 != nil {
					host.Type = "IPv4"
					host.Host = ip4.String()
					host.ip = ip4

				} else if ip16 := ip.To16(); ip16 != nil {
					host.Type = "IPv6"
					host.Host = fmt.Sprintf("[%s]", ip16)
					host.ip = ip16
				}
				hosts <- host

				if !sc.scanAllAddresses {
					break
				}
			}
		} else {
			sc.Logger.Error().
				Err(err).
				Str("target", targetSource).
				Msg("Failed to resolve target hostname")
		}

	} else {
		sc.Logger.Error().
			Err(err).
			Str("target", targetSource).
			Msg("Could not resolve target. Invalid format")
	}

	sc.Logger.Trace().
		Str("type", "return").
		Str("function", "(*UdpProbeScanner).Scan").
		Dict("arguments", zerolog.Dict().
			Str("targetSource", targetSource).
			Interface("hosts", hosts)).
		Dict("return", zerolog.Dict().
			AnErr("err", err)).
		Msg("return <- (*UdpProbeScanner).Scan()")

	return nil
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

		if conn, err = net.Dial(transport, address); err == nil {
			if err = conn.SetReadDeadline(time.Now().Add(sc.ReadTimeout)); err == nil {

				response := make([]byte, 0x400)

				sc.Logger.Trace().
					Str("type", "connection.write").
					Str("transport", transport).
					Str("address", address).
					Bytes("data", payload).
					Msg("(net.Conn).Write(data)")

				conn.Write(payload)

				if readLen, err = bufio.NewReader(conn).Read(response); err == nil {
					result = PortResult{
						Port:      port,
						Transport: transport,
						Host:      host,
					}
					sc.Logger.Trace().
						Str("type", "connection.read").
						Str("transport", transport).
						Str("address", address).
						Bytes("data", response).
						Msg("(net.Conn).Read(data)")

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

	if sc.Logger.GetLevel().String() == zerolog.LevelTraceValue {
		var loggedResult interface{}

		if result.Transport == "" {
			loggedResult = nil
		} else {
			loggedResult = result
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
				Interface("result", loggedResult)).
			Msg("return <- (*UdpProbeScanner).scanTask()")
	}
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
		Uint("total_probe_count", totalCount).
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

			var portStatus uint8 = STATE_UNRESPONSIVE

			for _, service := range data.UDP_SERVICES {
				for _, port := range service.Ports {
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
										result.Service = data.UDP_SERVICES[probe.Service]
										result.Probe = probe
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

func NewUdpProbeScanner(logger zerolog.Logger,
	hostConcurrency uint, portConcurrency uint,
	retransmissions uint, readTimeout time.Duration,
	scanAllAddresses bool) (sc UdpProbeScanner) {

	sc.resultsLive = make(chan PortResult)
	sc.resultsMap = make(map[string]map[uint16][]PortResult)
	sc.HostConcurrency = hostConcurrency
	sc.PortConcurrency = portConcurrency
	sc.ReadTimeout = readTimeout
	sc.Logger = logger
	sc.Retransmissions = retransmissions
	return
}
