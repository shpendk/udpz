package scan

import (
	"net"
	"time"
	"udpz/pkg/data"

	"github.com/rs/zerolog"
	"github.com/txthinking/socks5"
)

type UdpProbeScanner struct {
	HostConcurrency  uint
	PortConcurrency  uint
	ProbeCount       uint
	Retransmissions  uint
	scanAllAddresses bool
	ReadTimeout      time.Duration

	Logger   zerolog.Logger
	proxy    *socks5.Client
	useProxy bool

	resultsLive chan PortResult
	results     []PortResult
	resultsMap  map[string]map[uint16][]PortResult
}

type Target struct {
	Type   string `yaml:"type" json:"type"`
	Target string `yaml:"source" json:"source"`
}

type Host struct {
	Type   string `yaml:"type" json:"type"`
	Host   string `yaml:"host" json:"host"`
	Target Target `yaml:"target" json:"target"`
	ip     net.IP
}

type PortResult struct {
	Host      Host                  `yaml:"host" json:"host"`
	Port      uint16                `yaml:"port" json:"port"`
	Transport string                `yaml:"transport" json:"transport"`
	Service   data.UdpServiceOutput `yaml:"service" json:"service"`
	Probe     data.UdpProbeOutput   `yaml:"probe" json:"probe"`
	Response  string                `yaml:"response" json:"response"`
}
