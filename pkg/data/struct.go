package data

type UdpPort uint16

type Service struct {
	Slug        string   `yaml:"slug" json:"slug"`
	Name        string   `yaml:"name" json:"name"`
	Long        string   `yaml:"long" json:"long"`
	Description string   `yaml:"description" json:"description"`
	References  []string `yaml:"references" json:"references"`
}

type UdpProbe struct {
	Payloads [][]byte
	Ports    []UdpPort
	Service  Service
}
