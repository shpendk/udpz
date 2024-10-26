package scan

import (
	"net"
	"time"
	"udpz/pkg/data"

	"github.com/rs/zerolog"
)

type UdpProbeScanner struct {
	HostConcurrency uint
	PortConcurrency uint
	ProbeCount      uint
	Retransmissions uint
	ReadTimeout     time.Duration
	Logger          zerolog.Logger

	resultsLive chan PortResult
	results     []PortResult
	resultsMap  map[string]map[uint16][]PortResult

	scanAllAddresses bool
}

type Target struct {
	Type   string `yaml:"type" json:"type"`
	Target string `yaml:"source" json:"source"`
}

type Host struct {
	Type   string `yaml:"type" json:"type"`
	Host   string `yaml:"host" json:"host"`
	Target Target `yaml:"target" json:"target"`

	ip net.IP
}

type PortResult struct {
	Host      Host            `yaml:"host" json:"host"`
	Port      uint16          `yaml:"port" json:"port"`
	Transport string          `yaml:"transport" json:"transport"`
	Probe     data.UdpProbe   `yaml:"probe" json:"probe"`
	Response  string          `yaml:"response" json:"response"`
	Service   data.UdpService `yaml:"service" json:"service"`
}

type HostResult struct {
	Host  Host
	Ports PortResult
	Start time.Time
	End   time.Time
}
