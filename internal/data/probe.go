package data

func (p *UdpProbe) ToOutput() UdpProbeOutput {
	return UdpProbeOutput{
		Slug: p.Slug,
		Name: p.Name,
	}
}
