package data

import (
	"strings"

	"github.com/rs/zerolog"
)

type UdpDataBroker struct {
	Probes   map[string]UdpProbe
	Services map[string]UdpService

	Logger    zerolog.Logger
	populated bool
}

func (u *UdpDataBroker) populate() {
	/*
		This just populates the UdpDataBroker structure with available
		probes & services
	*/

	u.Probes = make(map[string]UdpProbe)

	if !u.populated {
		u.Services = UDP_SERVICES_DEFAULT
		u.Probes = make(map[string]UdpProbe)

		for _, service := range UDP_SERVICES_DEFAULT {

			for _, probe := range service.Probes {
				u.Probes[probe.Slug] = probe
			}
		}
		u.populated = true
	}
}

func (u *UdpDataBroker) Filter(slugs []string, tags []string) (services map[string]UdpService, probeMap map[string]UdpProbe) {
	/*
		From a list of slugs and tags, this creates a map of relevant
		services & a map of relevant probes
	*/

	u.populate()

	services = make(map[string]UdpService)
	probeMap = make(map[string]UdpProbe)

	for serviceSlug, service := range u.Services {

		// Filter by tags
		for _, tag := range tags {

			if service.hasTag(tag) {
				services[serviceSlug] = service

				for _, probe := range service.Probes {
					probeMap[probe.Slug] = probe
				}
			}
		}

		// Filter by slugs
		for _, slug := range slugs {

			// Check if probe is already found
			if _, ok := probeMap[slug]; !ok {

				if strings.HasPrefix(slug, serviceSlug) {

					if probes, ok := service.getProbes(slug); ok {
						// Create new service struct & only add matching probes

						for _, probe := range probes {
							probeMap[probe.Slug] = probe
						}
						var newService UdpService

						if newService, ok = services[serviceSlug]; ok {
							newService.Probes = append(newService.Probes, probes...)
						} else {
							newService = service
							newService.Probes = probes
						}
						services[serviceSlug] = newService
					}
				}
			}
		}
	}
	return
}

func NewUdpDataBroker(logger zerolog.Logger) (broker UdpDataBroker) {
	/*
		Instantiate & populate a new UdpDataBroker structure
	*/

	broker.Logger = logger
	broker.populate()
	return
}
