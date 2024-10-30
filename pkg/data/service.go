package data

func (s *UdpService) ToOutput() UdpServiceOutput {
	return UdpServiceOutput{
		Slug:        s.Slug,
		Name:        s.Name,
		NameShort:   s.NameShort,
		Description: s.Description,
		Tags:        s.Tags,
		References:  s.References,
	}
}

func (p *UdpProbe) ToOutput() UdpProbeOutput {
	return UdpProbeOutput{
		Slug: p.Slug,
		Name: p.Name,
	}
}
