package scan

import (
	"bufio"
	"fmt"
	"net"
	"os"

	"github.com/rs/zerolog"
)

func (sc *UdpProbeScanner) ResolveTargetLine(targetSource string, hosts chan Host) (ok bool) {

	sc.Logger.Trace().
		Str("type", "call").
		Str("function", "(*UdpProbeScanner).ResolveTargetLine").
		Dict("arguments", zerolog.Dict().
			Str("targetSource", targetSource).
			Interface("hosts", hosts)).
		Msg("(*UdpProbeScanner).ResolveTargetLine(...)")

	var target Target
	var host Host

	target.Target = targetSource

	if ip := net.ParseIP(targetSource); ip != nil {
		// If target can be parsed as IP address

		target.Type = "IP"
		host.Target = target

		// Check if IP address is IPv4
		if ip4 := ip.To4(); ip4 != nil {

			host.Type = "IPv4"
			host.Host = ip4.String()

		} else if ip16 := ip.To16(); ip16 != nil {
			// If IP address is IPv6

			host.Type = "IPv6"
			host.Host = fmt.Sprintf("[%s]", ip16) // net.Dial requires IPv6 addresses to be in square brackets
		}
		sc.Logger.Debug().
			Str("type", target.Type).
			Str("target", target.Target).
			Str("address_type", host.Type).
			IPAddr("ip", ip).
			Msg("Target resolved")

		hosts <- host // Target is resolved to IP address
		ok = true

	} else if ip, ipNet, err := net.ParseCIDR(targetSource); err == nil {
		// If target can be parsed as CIDR

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
			Str("address_type", addrType).
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
				host.Host = fmt.Sprintf("[%s]", ip)
			}
			hosts <- host

			for j := len(ip) - 1; j >= 0; j-- {
				ip[j]++
				if ip[j] != 0 {
					break
				}
			}
		}
		ok = true

	} else if REGEX_HOSTNAME.MatchString(targetSource) {

		if ips, err := net.LookupIP(targetSource); err == nil {

			target.Type = "hostname"

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
			ok = true
		}
	}

	sc.Logger.Trace().
		Str("type", "return").
		Str("function", "(*UdpProbeScanner).ResolveTargetLine").
		Dict("arguments", zerolog.Dict().
			Str("targetSource", targetSource).
			Interface("hosts", hosts)).
		Dict("return", zerolog.Dict().
			Bool("ok", ok)).
		Msg("return <- (*UdpProbeScanner).ResolveTargetLine()")
	return
}

func (sc *UdpProbeScanner) ResolveTarget(targetSource string, hosts chan Host) (err error) {

	sc.Logger.Trace().
		Str("type", "call").
		Str("function", "(*UdpProbeScanner).ResolveTarget").
		Dict("arguments", zerolog.Dict().
			Str("targetSource", targetSource).
			Interface("hosts", hosts)).
		Msg("(*UdpProbeScanner).ResolveTarget(...)")

	if sc.ResolveTargetLine(targetSource, hosts) {

		sc.Logger.Debug().
			Str("target", targetSource).
			Msg("Target resolved")

	} else if _, err := os.Stat(targetSource); err == nil {

		var file *os.File

		if file, err = os.OpenFile(targetSource, os.O_RDONLY, 0); err != nil {
			sc.Logger.Error().
				Err(err).
				Str("path", targetSource).
				Msg("Failed to open target input file")
		}
		defer func() {
			if err = file.Close(); err != nil {
				sc.Logger.Error().
					Err(err).
					Str("path", targetSource).
					Msg("Failed to close target input file")
			}
		}()

		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)

		for scanner.Scan() {
			sc.ResolveTargetLine(scanner.Text(), hosts)
		}
	} else {
		sc.Logger.Error().
			Str("target", targetSource).
			Msg("Could not resolve target")
	}

	sc.Logger.Trace().
		Str("type", "return").
		Str("function", "(*UdpProbeScanner).ResolveTarget").
		Dict("arguments", zerolog.Dict().
			Str("targetSource", targetSource).
			Interface("hosts", hosts)).
		Dict("return", zerolog.Dict().
			AnErr("err", err)).
		Msg("return <- (*UdpProbeScanner).ResolveTarget()")

	return nil
}
