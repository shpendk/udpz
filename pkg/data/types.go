package data

type UdpServiceOutput struct {
	Slug        string   `yaml:"slug" json:"slug"`
	Name        string   `yaml:"name" json:"name"`
	NameShort   string   `yaml:"short" json:"short"`
	Description string   `yaml:"description" json:"description"`
	Tags        []string `yaml:"tags" json:"tags"`
	References  []string `yaml:"references" json:"references"`
}

type UdpService struct {
	Slug        string     `yaml:"slug" json:"slug"`
	Name        string     `yaml:"name" json:"name"`
	NameShort   string     `yaml:"short" json:"short"`
	Description string     `yaml:"description" json:"description"`
	Tags        []string   `yaml:"tags" json:"tags"`
	References  []string   `yaml:"references" json:"references"`
	Ports       []uint16   `yaml:"ports" json:"ports"`
	Probes      []UdpProbe `yaml:"probes" json:"probes"`
}

type UdpProbeOutput struct {
	Slug string `yaml:"slug" json:"slug"`
	Name string `yaml:"name" json:"name"`
}

type UdpProbe struct {
	Slug        string `yaml:"slug" json:"slug"`
	Name        string `yaml:"name" json:"name"`
	Service     string `yaml:"service" json:"service"`
	EncodedData string `yaml:"data" json:"data"`
}
