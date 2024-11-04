package data

import "fmt"

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

func (s *UdpService) hasTag(tag string) bool {
	if tag == "all" {
		return true
	}
	for _, existingTag := range s.Tags {
		if existingTag == tag {
			return true
		}
	}
	return false
}

func (s *UdpService) getProbes(slug string) ([]UdpProbe, bool) {

	if slug == "all" {
		return s.Probes, true
	}
	for _, probe := range s.Probes {
		if probe.Slug == slug {
			return []UdpProbe{probe}, true
		}
	}
	if slug == s.Slug || slug == fmt.Sprintf("%s:*", s.Slug) || slug == fmt.Sprintf("%s:", s.Slug) {
		return s.Probes, true
	}

	return nil, false
}
