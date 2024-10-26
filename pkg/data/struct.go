package data

type UdpService struct {
	Slug        string `yaml:"slug" json:"slug"`
	Name        string `yaml:"name" json:"name"`
	NameShort   string `yaml:"short" json:"short"`
	Description string `yaml:"description" json:"description"`

	Ports      []uint16   `yaml:"ports" json:"ports"`
	Probes     []UdpProbe `yaml:"probes" json:"probes"`
	Tags       []string   `yaml:"tags" json:"tags"`
	References []string   `yaml:"references" json:"references"`

	//Vendor      string `yaml:"vendor" json:"vendor"`
}

type UdpProbe struct {
	Slug        string `yaml:"slug" json:"slug"`
	Name        string `yaml:"name" json:"name"`
	Service     string `yaml:"service" json:"service"`
	EncodedData string `yaml:"data" json:"data"`
}
