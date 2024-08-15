package scan

import "fmt"

func (sc *UdpProbeScanner) logInfo(fstr string, args ...any) {
	if !sc.NoLog {
		sc.ProgressWriter.Log("[INFO] %s", fmt.Sprintf(fstr, args...))
	}
}

func (sc *UdpProbeScanner) logDebug(fstr string, args ...any) {
	if !sc.NoLog {
		sc.ProgressWriter.Log("[DEBUG] %s", fmt.Sprintf(fstr, args...))
	}
}

func (sc *UdpProbeScanner) logPortResult(pr PortResult) {
	if !sc.NoLog && sc.ProgressWriter.IsRenderInProgress() {
		sc.logInfo("Open port - %s:%s:%d (%s)", pr.Transport, pr.Target, pr.Port, pr.Service.Slug)
	}
}
