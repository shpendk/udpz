package scan

import (
	"bufio"
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
	//var host Host

	target.Target = targetSource

	if ip := net.ParseIP(targetSource); ip != nil {
		// If target can be parsed as IP address

		sc.Logger.Debug().
			Str("target", target.String()).
			IPAddr("ip", ip).
			Msg("Target resolved")

		hosts <- Host{
			Target: target,
			Host:   ip,
		} // Target is resolved to IP address
		ok = true

	} else if ip, ipNet, err := net.ParseCIDR(targetSource); err == nil {
		// If target can be parsed as CIDR

		target.Type = "CIDR"

		sc.Logger.Debug().
			Str("target", target.String()).
			IPAddr("ip", ip).
			Str("cidr", ipNet.String()).
			Msg("Target CIDR resolved")

		for ipNet.Contains(ip) {
			hosts <- Host{
				Target: target,
				Host:   ip,
			}

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

				hosts <- Host{
					Target: target,
					Host:   ip,
				}

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
