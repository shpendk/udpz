package data

/*
[RESOURCES]
- https://github.com/Orange-Cyberdefense/awesome-industrial-protocols?tab=readme-ov-file
*/

/*
# TODO
[ ] Find and remove unnecessary payloads
[ ] 631: CUPS/IPP
[ ] 7000: Citrix NetScaler NSIP -> https://community.citrix.com/tech-zone/build/tech-papers/citrix-communication-ports
[ ] 3003: Citrix NetScaler Appliance -> https://community.citrix.com/tech-zone/build/tech-papers/citrix-communication-ports
[ ] 162: Citrix NetScaler ADM Appliance
[ ] 5351: Port Control Protocol -> https://datatracker.ietf.org/doc/html/rfc6887
[X] 20000: DNP3 (ICS)
[ ] 4000: ROC PLus
[ ] 1089-1091: Foundation Fieldbus HSE
[ ] 55000-55003: FL-net
[X] 9600: OMRON
[X] https://www.shadowserver.org/what-we-do/network-reporting/accessible-stun-service-report/
[ ] 502: ModbusUDP https://www.speedguide.net/port.php?port=502
*/

var (
	NUM_PAYLOADS int = 81 // # of payloads that will be sent

	UDP_PROBE_ARD = UdpProbe{
		Service: Service{
			Slug:        "ard",
			Name:        "ARD",
			Long:        "Apple Remote Desktop (ARD)",
			Description: `Apple Remote Desktop (ARD) allows users to remotely control or monitor other computers`,
			References: []string{
				"https://www.speedguide.net/port.php?port=3283",
				"https://wikipedia.org/wiki/Apple_Remote_Desktop",
			},
		},
		Ports: []UdpPort{
			3283,
		},
		Payloads: [][]byte{
			{0, 20, 0, 1, 3},
		},
	}
	UDP_PROBE_BACNET = UdpProbe{ // Validated
		Service: Service{
			Slug:        "bacnet",
			Name:        "BACNet",
			Long:        "Building Automation & Control Networks (BACNet)",
			Description: `Building Automation & Control Networks (BACNet) controls communication of building automation and control systems for applications such as heating, ventilating, air-conditioning control (HVAC), lighting control, access control, and fire detection systems and their associated equipment.`,
			References: []string{
				"https://www.speedguide.net/port.php?port=47808",
				"https://wikipedia.org/wiki/BACnet",
			},
		},
		Ports: []UdpPort{
			47808,
		},
		Payloads: [][]byte{
			{129, 10, 0, 17, 1, 4, 0, 5, 214, 12, 12, 2, 63, 255, 255, 25, 75, 76},
			{129, 10, 0, 37, 1, 4, 2, 5, 1, 14, 12, 2, 0, 0, 0, 30, 9, 12, 9, 28, 9, 44, 9, 56, 9, 57, 9, 58, 9, 70, 9, 77, 9, 120, 9, 121, 31},
		},
	}
	UDP_PROBE_CHARGEN = UdpProbe{
		Service: Service{
			Slug:        "chargen",
			Name:        "CharGen",
			Long:        "Character Generator Protocol",
			Description: `Character Generator Protocol (CharGen) generates and replies with a packet containing arbitrary characters. Should be disabled if there is no specific need for it, source for potential attacks.`,
			References: []string{
				"https://www.speedguide.net/port.php?port=19",
				"https://wikipedia.org/wiki/Character_Generator_Protocol",
				"https://datatracker.ietf.org/doc/html/rfc864",
			},
		},
		Ports: []UdpPort{
			19,
		},
		Payloads: [][]byte{
			{1},
		},
	}
	/*
		UDP_PROBE_WINFRAME = UdpProbe{
			Service: Service{
				Slug: "winframe",
				Name: "Citrix WinFrame",
				Long: "Citrix WinFrame Server",
				References: []string{
					"https://www.speedguide.net/port.php?port=1604",
				},
			},
			Ports: []UdpPort{
				1604,
			}
			Payloads: [][]byte{
				{30, 0, 1, 48, 2, 253, 168, 227, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			},
		}
	*/
	UDP_PROBE_COAP = UdpProbe{
		Service: Service{
			Slug:        "coap",
			Name:        "COAP",
			Description: `The Constrained Application Protocol (CoAP) is a specialized web transfer protocol for use with constrained nodes and constrained (low-power,lossy) networks often comprised of 8-bit microcontrollers with limited memory and ROM.`,
			References: []string{
				"https://www.speedguide.net/port.php?port=5683",
				"https://wikipedia.org/wiki/Constrained_Application_Protocol",
				"https://datatracker.ietf.org/doc/html/rfc7252",
			},
		},
		Ports: []UdpPort{
			5683,
			5684, // alt
		},
		Payloads: [][]byte{
			{64, 1, 125, 112, 187, 46, 119, 101, 108, 108, 45, 107, 110, 111, 119, 110, 4, 99, 111, 114, 101},
		},
	}
	UDP_PROBE_IBM_DB2 = UdpProbe{
		Service: Service{
			Slug:        "ibm-db2",
			Name:        "IBM-DB2",
			Description: `Family of data management products including database servers, developed by IBM.`,
			References: []string{
				"https://www.speedguide.net/port.php?port=523",
				"https://wikipedia.org/wiki/IBM_Db2",
			},
		},
		Ports: []UdpPort{
			523,
		},
		Payloads: [][]byte{
			{68, 66, 50, 71, 69, 84, 65, 68, 68, 82, 0, 83, 81, 76, 48, 57, 48, 49, 48, 0},
			{68, 66, 50, 71, 69, 84, 65, 68, 68, 82, 0, 83, 81, 76, 48, 53, 48, 48, 48, 0},
		},
	}
	/*
		UDP_PROBE_DIGI_ADDP = UdpProbe{
			Service: Service{
				Slug: "digi-addp",
				Name: "Digi ADDP",
				Long: "Digi Advanced Device Discovery Protocol (ADDP)",
				References: []string{
					"https://www.speedguide.net/port.php?port=2362",
					"https://github.com/christophgysin/addp",
					"https://www.digi.com/resources/documentation/digidocs/90001537/references/r_advanced_device_discovery_prot.htm",
				},
			},
			Ports: []UdpPort{
				2362,
			}
			Payloads: [][]byte{
				{68, 73, 71, 73, 0, 1, 0, 6, 255, 255, 255, 255, 255, 255},
				{68, 86, 75, 84, 0, 1, 0, 6, 255, 255, 255, 255, 255, 255},
				{68, 71, 68, 80, 0, 1, 0, 6, 255, 255, 255, 255, 255, 255},
			},
		}
	*/
	UDP_PROBE_DNP3 = UdpProbe{
		Service: Service{
			Slug:        "dnp3",
			Name:        "DNP3",
			Description: ``, // TODO
			References: []string{
				"https://www.speedguide.net/port.php?port=20000",
				"https://en.wikipedia.org/wiki/DNP3",
			},
		},
		Ports: []UdpPort{
			20000,
		},
		Payloads: [][]byte{
			{5, 100, 5, 201, 0, 0, 0, 0, 54, 76, 5, 100, 5, 201, 1, 0, 0, 0, 222, 142, 5, 100, 5, 201, 2, 0, 0, 0, 159, 132, 5, 100, 5, 201, 3, 0, 0, 0, 119, 70, 5, 100, 5, 201, 4, 0, 0, 0, 29, 144, 5, 100, 5, 201, 5, 0, 0, 0, 245, 82, 5, 100, 5, 201, 6, 0, 0, 0, 180, 88, 5, 100, 5, 201, 7, 0, 0, 0, 92, 154, 5, 100, 5, 201, 8, 0, 0, 0, 25, 185, 5, 100, 5, 201, 9, 0, 0, 0, 241, 123, 5, 100, 5, 201, 10, 0, 0, 0, 176, 113, 5, 100, 5, 201, 11, 0, 0, 0, 88, 179, 5, 100, 5, 201, 12, 0, 0, 0, 50, 101, 5, 100, 5, 201, 13, 0, 0, 0, 218, 167, 5, 100, 5, 201, 14, 0, 0, 0, 155, 173, 5, 100, 5, 201, 15, 0, 0, 0, 115, 111, 5, 100, 5, 201, 16, 0, 0, 0, 17, 235, 5, 100, 5, 201, 17, 0, 0, 0, 249, 41, 5, 100, 5, 201, 18, 0, 0, 0, 184, 35, 5, 100, 5, 201, 19, 0, 0, 0, 80, 225, 5, 100, 5, 201, 20, 0, 0, 0, 58, 55, 5, 100, 5, 201, 21, 0, 0, 0, 210, 245, 5, 100, 5, 201, 22, 0, 0, 0, 147, 255, 5, 100, 5, 201, 23, 0, 0, 0, 123, 61, 5, 100, 5, 201, 24, 0, 0, 0, 62, 30, 5, 100, 5, 201, 25, 0, 0, 0, 214, 220, 5, 100, 5, 201, 26, 0, 0, 0, 151, 214, 5, 100, 5, 201, 27, 0, 0, 0, 127, 20, 5, 100, 5, 201, 28, 0, 0, 0, 21, 194, 5, 100, 5, 201, 29, 0, 0, 0, 253, 0, 5, 100, 5, 201, 30, 0, 0, 0, 188, 10, 5, 100, 5, 201, 31, 0, 0, 0, 84, 200, 5, 100, 5, 201, 32, 0, 0, 0, 1, 79, 5, 100, 5, 201, 33, 0, 0, 0, 233, 141, 5, 100, 5, 201, 34, 0, 0, 0, 168, 135, 5, 100, 5, 201, 35, 0, 0, 0, 64, 69, 5, 100, 5, 201, 36, 0, 0, 0, 42, 147, 5, 100, 5, 201, 37, 0, 0, 0, 194, 81, 5, 100, 5, 201, 38, 0, 0, 0, 131, 91, 5, 100, 5, 201, 39, 0, 0, 0, 107, 153, 5, 100, 5, 201, 40, 0, 0, 0, 46, 186, 5, 100, 5, 201, 41, 0, 0, 0, 198, 120, 5, 100, 5, 201, 42, 0, 0, 0, 135, 114, 5, 100, 5, 201, 43, 0, 0, 0, 111, 176, 5, 100, 5, 201, 44, 0, 0, 0, 5, 102, 5, 100, 5, 201, 45, 0, 0, 0, 237, 164, 5, 100, 5, 201, 46, 0, 0, 0, 172, 174, 5, 100, 5, 201, 47, 0, 0, 0, 68, 108, 5, 100, 5, 201, 48, 0, 0, 0, 38, 232, 5, 100, 5, 201, 49, 0, 0, 0, 206, 42, 5, 100, 5, 201, 50, 0, 0, 0, 143, 32, 5, 100, 5, 201, 51, 0, 0, 0, 103, 226, 5, 100, 5, 201, 52, 0, 0, 0, 13, 52, 5, 100, 5, 201, 53, 0, 0, 0, 229, 246, 5, 100, 5, 201, 54, 0, 0, 0, 164, 252, 5, 100, 5, 201, 55, 0, 0, 0, 76, 62, 5, 100, 5, 201, 56, 0, 0, 0, 9, 29, 5, 100, 5, 201, 57, 0, 0, 0, 225, 223, 5, 100, 5, 201, 58, 0, 0, 0, 160, 213, 5, 100, 5, 201, 59, 0, 0, 0, 72, 23, 5, 100, 5, 201, 60, 0, 0, 0, 34, 193, 5, 96, 92, 147, 208, 0, 0, 12, 160, 48, 86, 64, 92, 147, 224, 0, 0, 8, 176, 144, 86, 64, 92, 147, 240, 0, 0, 6, 60, 176, 86, 64, 92, 148, 0, 0, 0, 5, 132, 160, 86, 64, 92, 148, 16, 0, 0, 11, 8, 128, 86, 64, 92, 148, 32, 0, 0, 15, 24, 32, 86, 64, 92, 148, 48, 0, 0, 1, 148, 0, 86, 64, 92, 148, 64, 0, 0, 7, 57, 96, 86, 64, 92, 148, 80, 0, 0, 9, 181, 64, 86, 64, 92, 148, 96, 0, 0, 13, 165, 224, 86, 64, 92, 148, 112, 0, 0, 3, 41, 192, 86, 64, 92, 148, 128, 0, 0, 7, 123, 240, 86, 64, 92, 148, 144, 0, 0, 9, 247, 208, 86, 64, 92, 148, 160, 0, 0, 13, 231, 112},
		},
	}
	UDP_PROBE_DNS = UdpProbe{
		Service: Service{
			Slug:        "dns",
			Name:        "DNS",
			Description: `Domain Name System (DNS) is a hierarchical and distributed name service that provides a naming system for computers, services, and other resources in the Internet or other Internet Protocol (IP) networks`,
			References: []string{
				"https://www.speedguide.net/port.php?port=53",
				"https://wikipedia.org/wiki/Domain_Name_System",
			},
		},
		Ports: []UdpPort{
			53,
		},
		Payloads: [][]byte{
			// SOA
			{252, 142, 1, 32, 0, 1, 0, 0, 0, 0, 0, 1, 0, 0, 2, 0, 1, 0, 0, 41, 4, 208, 0, 0, 0, 0, 0, 12, 0, 10, 0, 8, 61, 35, 192, 10, 249, 210, 251, 58},
			// version.bind
			{158, 64, 1, 32, 0, 1, 0, 0, 0, 0, 0, 1, 7, 118, 101, 114, 115, 105, 111, 110, 4, 98, 105, 110, 100, 0, 0, 16, 0, 3, 0, 0, 41, 4, 208, 0, 0, 0, 0, 0, 12, 0, 10, 0, 8, 208, 122, 48, 228, 216, 89, 46, 160},
			// A "localhost"
			{0, 0, 1, 32, 0, 1, 0, 0, 0, 0, 0, 1, 9, 108, 111, 99, 97, 108, 104, 111, 115, 116, 0, 0, 1, 0, 1, 0, 0, 41, 4, 208, 0, 0, 0, 0, 0, 12, 0, 10, 0, 8, 151, 227, 150, 141, 120, 208, 243, 106},
		},
	}
	UDP_PROBE_IPMI = UdpProbe{
		Service: Service{
			Slug:        "ipmi",
			Name:        "IPMI",
			Description: `Intelligent Platform Management Interface (IPMI) is a set of computer interface specifications for an autonomous computer subsystem that provides management and monitoring capabilities independent of the host system's CPU, firmware (BIOS or UEFI) and operating system`,
			References: []string{
				"https://www.speedguide.net/port.php?port=623",
				"https://wikipedia.org/wiki/Intelligent_Platform_Management_Interface",
				"https://wiki.wireshark.org/IPMI",
			},
		},
		Ports: []UdpPort{
			623,
		},
		Payloads: [][]byte{
			{6, 0, 255, 7, 0, 0, 0, 0, 0, 0, 0, 0, 0, 9, 32, 24, 200, 129, 0, 56, 142, 4, 181},
		},
	}
	UDP_PROBE_CLDAP = UdpProbe{
		Service: Service{
			Slug:        "cldap",
			Name:        "CLDAP",
			Long:        "Connectionless Lightweight Directory Access Protocol (CLDAP)",
			Description: `Connectionless Lightweight Directory Access Protocol (CLDAP) is the connectionless variant of LDAP often used to query the RootDSE of a Microsoft Windows domain controller`,
			References: []string{
				"https://wikipedia.org/wiki/Lightweight_Directory_Access_Protocol",
				"https://www.speedguide.net/port.php?port=389",
				"https://wiki.wireshark.org/MS-CLDAP",
			},
		},
		Ports: []UdpPort{
			389,
		},
		Payloads: [][]byte{
			// This can be unstable - same payload 3x to improve detection probability
			{48, 132, 0, 0, 0, 45, 2, 1, 1, 99, 132, 0, 0, 0, 36, 4, 0, 10, 1, 0, 10, 1, 0, 2, 1, 0, 2, 1, 0, 1, 1, 0, 135, 11, 111, 98, 106, 101, 99, 116, 99, 108, 97, 115, 115, 48, 132, 0, 0, 0, 0, 0, 10},
			{48, 132, 0, 0, 0, 45, 2, 1, 1, 99, 132, 0, 0, 0, 36, 4, 0, 10, 1, 0, 10, 1, 0, 2, 1, 0, 2, 1, 0, 1, 1, 0, 135, 11, 111, 98, 106, 101, 99, 116, 99, 108, 97, 115, 115, 48, 132, 0, 0, 0, 0, 0, 10},
			{48, 132, 0, 0, 0, 45, 2, 1, 1, 99, 132, 0, 0, 0, 36, 4, 0, 10, 1, 0, 10, 1, 0, 2, 1, 0, 2, 1, 0, 1, 1, 0, 135, 11, 111, 98, 106, 101, 99, 116, 99, 108, 97, 115, 115, 48, 132, 0, 0, 0, 0, 0, 10},
		},
	}
	UDP_PROBE_MDNS = UdpProbe{
		Service: Service{
			Slug:        "mdns",
			Name:        "mDNS",
			Long:        "Multicast Domain Name System (mDNS)",
			Description: `Multicast Domain Name System (mDNS) resolves hostnames to IP addresses within small networks that do not include a local name server. It is a zero-configuration service, using essentially the same programming interfaces, packet formats and operating semantics as unicast Domain Name System (DNS)`,
			References: []string{
				"https://www.speedguide.net/port.php?port=5353",
				"https://wikipedia.org/wiki/Multicast_DNS",
			},
		},
		Ports: []UdpPort{
			5353,
		},
		Payloads: [][]byte{
			{27, 108, 1, 32, 0, 1, 0, 0, 0, 0, 0, 1, 1, 49, 1, 48, 1, 48, 3, 49, 50, 55, 7, 105, 110, 45, 97, 100, 100, 114, 4, 97, 114, 112, 97, 0, 0, 12, 0, 1, 0, 0, 41, 4, 208, 0, 0, 0, 0, 0, 12, 0, 10, 0, 8, 127, 1, 250, 112, 14, 12, 142, 176},
			//{0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 9, 95, 115, 101, 114, 118, 105, 99, 101, 115, 7, 95, 100, 110, 115, 45, 115, 100, 4, 95, 117, 100, 112, 5, 108, 111, 99, 97, 108, 0, 0, 12, 0, 1},
		},
	}
	UDP_PROBE_MEMCACHE = UdpProbe{
		Service: Service{
			Slug:        "memcache",
			Name:        "Memcache",
			Description: `Memcache is a general-purpose distributed memory-caching system often used to speed up dynamic database-driven websites by caching data and objects in RAM to reduce the number of times an external data source must be read`,
			References: []string{
				"https://www.speedguide.net/port.php?port=11211",
				"https://wikipedia.org/wiki/Memcached",
				"https://www.wireshark.org/docs/dfref/m/memcache.html",
			},
		},
		Ports: []UdpPort{
			11211,
		},
		Payloads: [][]byte{
			{90, 77, 0, 0, 0, 1, 0, 0, 115, 116, 97, 116, 115, 32, 105, 116, 101, 109, 115, 13, 10},
		},
	}
	UDP_PROBE_MELSEC_Q = UdpProbe{
		Service: Service{
			Slug:        "melsec-q",
			Name:        "MELSEC-Q",
			Long:        "Mitsubishi MELSEC-Q",
			Description: `Mitsubishi Electric MELSEC-Q Series PLCs use a proprietary network protocol for communication. The devices are used by equipment and manufacturing facilities to provide high-speed,  large volume data processing and machine control.`,
			References: []string{
				"https://dl.mitsubishielectric.com/dl/fa/document/manual/plc/sh080008/sh080008ab.pdf",
				"https://github.com/xl7dev/ICSecurity/blob/master/icse-nse/melsecq-discover-udp.nse",
			},
		},
		Ports: []UdpPort{
			5006,
		},
		Payloads: [][]byte{
			{87, 0, 0, 0, 0, 17, 17, 7, 0, 0, 255, 255, 3, 0, 0, 254, 3, 0, 0, 20, 0, 28, 8, 10, 8, 0, 0, 0, 0, 0, 0, 0, 4, 1, 1, 1, 0, 0, 0, 0, 1},
		},
	}
	UDP_PROBE_MOXA_NPORT = UdpProbe{
		Service: Service{
			Slug:        "moxa-nport",
			Name:        "Moxa NPort",
			Description: ``, // TODO
			References: []string{
				"https://github.com/xl7dev/ICSecurity/blob/45693d87b4cc2818d0ddf4a3e8d110eb41ffeec1/icse-nse/moxa-enum.nse",
			},
		},
		Ports: []UdpPort{
			4800,
		},
		Payloads: [][]byte{
			{1, 0, 0, 8, 0, 0, 0, 0},
		},
	}
	UDP_PROBE_MSSQL = UdpProbe{
		Service: Service{
			Slug:        "mssql",
			Name:        "MSSQL",
			Long:        "Microsoft SQL Server",
			Description: `Microsoft SQL Server (Structured Query Language) is a proprietary relational database management system developed by Microsoft. As a database server, it is a software product with the primary function of storing and retrieving data`,
			References: []string{
				"https://www.speedguide.net/port.php?port=1434",
				"https://wikipedia.org/wiki/Microsoft_SQL_Server",
			},
		},
		Ports: []UdpPort{
			1434,
		},
		Payloads: [][]byte{
			{2},
		},
	}
	UDP_PROBE_NAT_PMP = UdpProbe{
		Service: Service{
			Slug:        "nat-pmp",
			Name:        "NAT-PMP",
			Long:        "Network Address Translation Port Mapping Protocol",
			Description: `NAT Port Mapping Protocol (NAT-PMP) is a network protocol for automatically establishing network address translation (NAT) settings and port forwarding configurations`,
			References: []string{
				"https://www.speedguide.net/port.php?port=5351",
				"https://wikipedia.org/wiki/NAT_Port_Mapping_Protocol",
			},
		},
		Ports: []UdpPort{
			5351,
		},
		Payloads: [][]byte{
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 227, 179, 228, 131},
			{0, 0, 0, 0},
		},
	}
	UDP_PROBE_NETBIOS = UdpProbe{
		Service: Service{
			Slug:        "netbios",
			Name:        "NetBIOS",
			Long:        "Network Basic Input/Output System",
			Description: `Network Basic Input/Output System (NetBIOS) is a protocol often associated with Microsoft Windows that allows communication of data for files and printers over a network connection.`,
			References: []string{
				"https://www.speedguide.net/port.php?port=137",
				"https://wikipedia.org/wiki/NetBIOS",
			},
		},
		Ports: []UdpPort{
			137, // 138, 139 are also used (not for discovery though)
		},
		Payloads: [][]byte{
			{229, 216, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 32, 67, 75, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 0, 0, 33, 0, 1},
		},
	}
	UDP_PROBE_NETIS = UdpProbe{
		Service: Service{
			Slug:        "netis",
			Name:        "Netis Backdoor",
			Long:        "Netis Router Backdoor",
			Description: ``, // TODO
			References: []string{
				"https://www.shadowserver.org/what-we-do/network-reporting/netcore-netis-router-vulnerability-scan-report/",
			},
		},
		Ports: []UdpPort{
			53413,
		},
		Payloads: [][]byte{
			{10, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 158, 33, 189, 173},
		},
	}
	UDP_PROBE_NTP = UdpProbe{
		Service: Service{
			Slug:        "ntp",
			Name:        "NTP",
			Description: `Network Time Protocol (NTP) is a networking protocol for clock synchronization between computer systems over packet-switched, variable-latency data networks.`,
			References: []string{
				"https://www.speedguide.net/port.php?port=123",
				"https://wikipedia.org/wiki/Network_Time_Protocol",
			},
		},
		Ports: []UdpPort{
			123,
		},
		Payloads: [][]byte{
			{227, 0, 4, 250, 0, 1, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 197, 79, 35, 75, 113, 177, 82, 243},
			{23, 0, 3, 42, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
	}
	UDP_PROBE_FINS = UdpProbe{
		Service: Service{
			Slug:        "fins",
			Name:        "FINS",
			Description: `Factory Interface Network Service (FINS) is a network protocol used by Omron PLCs, over different physical networks like Ethernet, Controller Link, DeviceNet and RS-232C`,
			References: []string{
				"https://svn.nmap.org/nmap/scripts/omron-info.nse",
			},
		},
		Ports: []UdpPort{
			9600,
		},
		Payloads: [][]byte{
			{128, 0, 2, 0, 0, 0, 0, 99, 0, 239, 5, 1, 0},
			{70, 73, 78, 83, 0, 0, 0, 12, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
	}
	UDP_PROBE_OPENVPN = UdpProbe{ // TODO: make sure this is working - didn't work on HTB machine - Corporate
		Service: Service{
			Slug:        "openvpn",
			Name:        "OpenVPN",
			Long:        "OpenVPN (Virtual Private Networking)",
			Description: ``, // TODO
			References: []string{
				"https://www.speedguide.net/port.php?port=1194",
				"https://wikipedia.org/wiki/OpenVPN",
			},
		},
		Ports: []UdpPort{
			1194,
		},
		Payloads: [][]byte{
			{56, 18, 18, 18, 18, 18, 18, 18, 18, 0, 0, 0, 0, 0, 56, 177, 38, 222},
		},
	}
	UDP_PROBE_PCA = UdpProbe{
		Service: Service{
			Slug:        "pca",
			Name:        "pcAnywhere",
			Long:        "Symantec pcAnywhere",
			Description: ``, // TODO
			References: []string{
				"https://www.speedguide.net/port.php?port=5632",
				"https://wikipedia.org/wiki/PcAnywhere",
			},
		},
		Ports: []UdpPort{
			5632,
		},
		Payloads: [][]byte{
			{78, 81},
			{83, 84},
			{78, 81, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 65, 244, 191, 166},
		},
	}
	UDP_PROBE_PORTMAP = UdpProbe{
		Service: Service{
			Slug:        "portmap",
			Name:        "Portmap/RPC",
			Long:        "Sun Remote Procedure Call (RPC)",
			Description: ``, // TODO
			References: []string{
				"https://www.speedguide.net/port.php?port=111",
				"https://wikipedia.org/wiki/Sun_RPC",
			},
		},
		Ports: []UdpPort{
			111,
		},
		Payloads: [][]byte{
			{26, 169, 255, 225, 0, 0, 0, 0, 0, 0, 0, 2, 0, 1, 134, 160, 0, 0, 0, 2, 0, 0, 0, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
	}
	UDP_PROBE_QOTD = UdpProbe{
		Service: Service{
			Slug:        "qotd",
			Name:        "QOTD",
			Long:        "Quote of the Day (QOTD)",
			Description: ``, // TODO
			References: []string{
				"https://www.speedguide.net/port.php?port=17",
				"https://wikipedia.org/wiki/QOTD",
			},
		},
		Ports: []UdpPort{
			17,
		},
		Payloads: [][]byte{
			{13, 10},
		},
	}
	UDP_PROBE_RDP = UdpProbe{
		Service: Service{
			Slug: "rdp",
			Name: "RDP",
			Long: "Remote Desktop Protocol (RDP)",
			References: []string{
				"https://www.speedguide.net/port.php?port=3389",
				"https://wikipedia.org/wiki/Remote_Desktop_Protocol",
			},
		},
		Ports: []UdpPort{
			3389,
		},
		Payloads: [][]byte{
			{0, 0, 0, 0, 0, 0, 0, 255, 0, 0, 0, 0, 0, 0, 0, 0, 84, 6},
		},
	}
	UDP_PROBE_RIP = UdpProbe{
		Service: Service{
			Slug:        "rip",
			Name:        "RIP",
			Long:        "Routing Information Protocol (RIP)",
			Description: ``, // TODO
			References: []string{
				"https://www.speedguide.net/port.php?port=520",
				"https://wikipedia.org/wiki/Routing_Information_Protocol",
			},
		},
		Ports: []UdpPort{
			520,
		},
		Payloads: [][]byte{
			{1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 16, 0},
		},
	}
	UDP_PROBE_SENTINEL = UdpProbe{
		Service: Service{
			Slug:        "sentinel",
			Name:        "Sentinel LM",
			Long:        "Sentinel LM (License Manager)",
			Description: ``, // TODO
			References: []string{
				"https://www.speedguide.net/port.php?port=5093",
				"https://docs.sentinel.thalesgroup.com/ldk/LDKdocs/SPNL/LDK_SLnP_Guide/Distributing/License_Manager/020-License_Manager_Type.htm",
			},
		},
		Ports: []UdpPort{
			5093,
		},
		Payloads: [][]byte{
			{122, 0, 0, 0, 0, 0},
		},
	}
	UDP_PROBE_SIP = UdpProbe{
		Service: Service{
			Slug:        "sip",
			Name:        "SIP",
			Long:        "Session Initiation Protocol (SIP)",
			Description: `Session Initiation Protocol (SIP) is a signaling protocol used for initiating, maintaining, and terminating communication sessions that include voice, video and messaging applications`,
			References: []string{
				"https://www.speedguide.net/port.php?port=5060",
				"https://wikipedia.org/wiki/Session_Initiation_Protocol",
			},
		},
		Ports: []UdpPort{
			5060,
		},
		Payloads: [][]byte{
			{79, 80, 84, 73, 79, 78, 83},
		},
	}
	UDP_PROBE_SNMP = UdpProbe{
		Service: Service{
			Slug:        "snmp",
			Name:        "SNMP",
			Long:        "Simple Network Management Protocol (SNMP)",
			Description: `Simple Network Management Protocol (SNMP) is an Internet Standard protocol for collecting and organizing information about managed devices on IP networks and for modifying that information to change device behavior. Devices that typically support SNMP include cable modems, routers, switches, servers, workstations, printers, and more. SNMP exposes management data in the form of variables on the managed systems organized in a management information base (MIB), which describes the system status and configuration.`, // TODO: shorten
			References: []string{
				"https://www.speedguide.net/port.php?port=161",
				"https://wikipedia.org/wiki/Simple_Network_Management_Protocol",
			},
		},
		Ports: []UdpPort{
			161,
			10161, // alt
			162,   // alt
			10162, // alt
		},
		Payloads: [][]byte{
			{48, 41, 2, 1, 0, 4, 6, 112, 117, 98, 108, 105, 99, 160, 28, 2, 4, 86, 90, 220, 93, 2, 1, 0, 2, 1, 0, 48, 14, 48, 12, 6, 8, 43, 6, 1, 2, 1, 1, 1, 0, 5, 0},
			{48, 38, 2, 1, 1, 4, 6, 112, 117, 98, 108, 105, 99, 161, 25, 2, 4, 220, 99, 194, 154, 2, 1, 0, 2, 1, 0, 48, 11, 48, 9, 6, 5, 43, 6, 1, 2, 1, 5, 0},
			{48, 58, 2, 1, 3, 48, 15, 2, 2, 74, 105, 2, 3, 0, 255, 227, 4, 1, 4, 2, 1, 3, 4, 16, 48, 14, 4, 0, 2, 1, 0, 2, 1, 0, 4, 0, 4, 0, 4, 0, 48, 18, 4, 0, 4, 0, 160, 12, 2, 2, 55, 240, 2, 1, 0, 2, 1, 0, 48, 0},
		},
	}
	UDP_PROBE_SSDP = UdpProbe{
		Service: Service{
			Slug:        "ssdp",
			Name:        "SSDP",
			Long:        "Simple Service Discovery Protocol (SSDP)",
			Description: `Simple Service Discovery Protocol (SSDP) is a network protocol based on the Internet protocol suite for advertisement and discovery of network services and presence information. It accomplishes this without assistance of server-based configuration mechanisms, such as Dynamic Host Configuration Protocol (DHCP) or Domain Name System (DNS), without special static configuration of a network host`,
			References: []string{
				"https://www.speedguide.net/port.php?port=1900",
				"https://www.speedguide.net/port.php?port=5000",
				"https://wikipedia.org/wiki/Simple_Service_Discovery_Protocol",
			},
		},
		Ports: []UdpPort{
			1900,
			5000, // alt
		},
		Payloads: [][]byte{
			{77, 45, 83, 69, 65, 82, 67, 72, 32, 42, 32, 72, 84, 84, 80, 47, 49, 46, 49, 13, 10, 72, 79, 83, 84, 58, 50, 51, 57, 46, 50, 53, 53, 46, 50, 53, 53, 46, 50, 53, 48, 58, 49, 57, 48, 48, 13, 10, 83, 84, 58, 115, 115, 100, 112, 58, 97, 108, 108, 13, 10, 77, 65, 78, 58, 34, 115, 115, 100, 112, 58, 100, 105, 115, 99, 111, 118, 101, 114, 34, 13, 10, 13, 10},
		},
	}
	UDP_PROBE_STUN = UdpProbe{
		Service: Service{
			Slug:        "stun",
			Name:        "STUN",
			Description: `Session Traversal Utilities for NAT (STUN) is a standardized set of methods, including a network protocol, for traversal of network address translator (NAT) gateways in applications of real-time voice, video, messaging, and other interactive communications`,
			References: []string{
				"https://www.speedguide.net/port.php?port=3478",
				"https://www.speedguide.net/port.php?port=19302",
				"https://www.shadowserver.org/what-we-do/network-reporting/accessible-stun-service-report/",
				"https://www.rfc-editor.org/rfc/rfc5389",
			},
		},
		Ports: []UdpPort{
			3478,
			3470,  // alt
			19302, // alt
		},
		Payloads: [][]byte{
			{0, 1, 0, 0, 33, 18, 164, 66, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
	}
	UDP_PROBE_TFTP = UdpProbe{
		Service: Service{
			Slug:        "tftp",
			Name:        "TFTP",
			Long:        "Trivial File Transfer Protocol (TFTP)",
			Description: `Trivial File Transfer Protocol (TFTP) is a simple lockstep File Transfer Protocol which allows a client to get a file from or put a file onto a remote host. One of its primary uses is in the early stages of nodes booting from a local area network. TFTP has been used for this application because it is very simple to implement`,
			References: []string{
				"https://www.speedguide.net/port.php?port=69",
				"wikipedia.org/wiki/Trivial_File_Transfer_Protocol",
			},
		},
		Ports: []UdpPort{
			69,
			6969, // alt
		},
		Payloads: [][]byte{
			{0, 1, 47, 97, 0, 110, 101, 116, 97, 115, 99, 105, 105, 0},
		},
	}
	UDP_PROBE_UBIQUITI = UdpProbe{
		Service: Service{
			Slug: "ubiquiti",
			Name: "Ubiquiti AirControl",
			Long: "Ubiquiti Networks AirControl Management Discovery Protocol",
			References: []string{
				"https://www.speedguide.net/port.php?port=10001",
			},
		},
		Ports: []UdpPort{
			10001,
		},
		Payloads: [][]byte{
			{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 31, 127, 163, 102},
			{1, 0, 0, 0},
			{2, 8, 0, 0},
		},
	}
	UDP_PROBE_UPNP = UdpProbe{
		Service: Service{
			Slug: "upnp",
			Name: "UPnP",
			Long: "Universal Plug and Play (UPnP)",
			References: []string{
				"https://www.speedguide.net/port.php?port=1900",
				"https://www.speedguide.net/port.php?port=5000",
				"https://wikipedia.org/wiki/Universal_Plug_and_Play",
			},
		},
		Ports: []UdpPort{
			1900,
			5000,
		},
		Payloads: [][]byte{
			{77, 45, 83, 69, 65, 82, 67, 72, 32, 42, 32, 72, 84, 84, 80, 47, 49, 46, 49, 13, 10, 72, 111, 115, 116, 58, 50, 51, 57, 46, 50, 53, 53, 46, 50, 53, 53, 46, 50, 53, 48, 58, 49, 57, 48, 48, 13, 10, 83, 84, 58, 117, 112, 110, 112, 58, 114, 111, 111, 116, 100, 101, 118, 105, 99, 101, 13, 10, 77, 97, 110, 58, 34, 115, 115, 100, 112, 58, 100, 105, 115, 99, 111, 118, 101, 114, 34, 13, 10, 77, 88, 58, 51, 13, 10, 13, 10, 13, 10},
		},
	}
	UDP_PROBE_STEAM_HLTV = UdpProbe{
		Service: Service{
			Slug:        "steam-hltv",
			Name:        "Steam HLTV",
			Long:        "Steam (Valve gaming platform) Half-Life TV",
			Description: ``, // TODO
			References: []string{
				"https://www.speedguide.net/port.php?port=27015",
				"https://help.steampowered.com/en/faqs/view/558D-FD60-531D-98BC",
			},
		},
		Ports: []UdpPort{
			27015,
		},
		Payloads: [][]byte{
			{255, 255, 255, 255, 84, 83, 111, 117, 114, 99, 101, 32, 69, 110, 103, 105, 110, 101, 32, 81, 117, 101, 114, 121, 0},
		},
	}
	UDP_PROBE_WDBRPC = UdpProbe{
		Service: Service{
			Slug:        "wdbrpc",
			Name:        "WDBRPC",
			Long:        "vxWorks WDB remote debugging ONCRPC",
			Description: ``, // TODO
			References: []string{
				"https://www.speedguide.net/port.php?port=17185",
			},
		},
		Ports: []UdpPort{
			17185,
		},
		Payloads: [][]byte{
			{26, 9, 250, 186, 0, 0, 0, 0, 0, 0, 0, 2, 85, 85, 85, 85, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 255, 255, 85, 18, 0, 0, 0, 60, 0, 0, 0, 1, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0},
		},
	}
	UDP_PROBE_WSD = UdpProbe{
		Service: Service{
			Slug:        "wsd",
			Name:        "WSD",
			Long:        "Web Services Discovery (WSD)",
			Description: ``, // TODO
			References: []string{
				"https://www.speedguide.net/port.php?port=3702",
				"https://wikipedia.org/wiki/Web_Services_Discovery",
			},
		},
		Ports: []UdpPort{
			3702,
		},
		Payloads: [][]byte{
			{60, 58, 47, 62, 10},
			{60, 63, 120, 109, 108, 32, 118, 101, 114, 115, 105, 111, 110, 61, 34, 49, 46, 48, 34, 32, 101, 110, 99, 111, 100, 105, 110, 103, 61, 34, 117, 116, 102, 45, 56, 34, 63, 62, 10, 60, 115, 111, 97, 112, 58, 69, 110, 118, 101, 108, 111, 112, 101, 32, 120, 109, 108, 110, 115, 58, 115, 111, 97, 112, 61, 34, 104, 116, 116, 112, 58, 47, 47, 119, 119, 119, 46, 119, 51, 46, 111, 114, 103, 47, 50, 48, 48, 51, 47, 48, 53, 47, 115, 111, 97, 112, 45, 101, 110, 118, 101, 108, 111, 112, 101, 34, 32, 120, 109, 108, 110, 115, 58, 119, 115, 97, 61, 34, 104, 116, 116, 112, 58, 47, 47, 115, 99, 104, 101, 109, 97, 115, 46, 120, 109, 108, 115, 111, 97, 112, 46, 111, 114, 103, 47, 119, 115, 47, 50, 48, 48, 52, 47, 48, 56, 47, 97, 100, 100, 114, 101, 115, 115, 105, 110, 103, 34, 32, 120, 109, 108, 110, 115, 58, 119, 115, 100, 61, 34, 104, 116, 116, 112, 58, 47, 47, 115, 99, 104, 101, 109, 97, 115, 46, 120, 109, 108, 115, 111, 97, 112, 46, 111, 114, 103, 47, 119, 115, 47, 50, 48, 48, 53, 47, 48, 52, 47, 100, 105, 115, 99, 111, 118, 101, 114, 121, 34, 32, 120, 109, 108, 110, 115, 58, 119, 115, 100, 112, 61, 34, 104, 116, 116, 112, 58, 47, 47, 115, 99, 104, 101, 109, 97, 115, 46, 120, 109, 108, 115, 111, 97, 112, 46, 111, 114, 103, 47, 119, 115, 47, 50, 48, 48, 54, 47, 48, 50, 47, 100, 101, 118, 112, 114, 111, 102, 34, 62, 10, 60, 115, 111, 97, 112, 58, 72, 101, 97, 100, 101, 114, 62, 60, 119, 115, 97, 58, 84, 111, 62, 117, 114, 110, 58, 115, 99, 104, 101, 109, 97, 115, 45, 120, 109, 108, 115, 111, 97, 112, 45, 111, 114, 103, 58, 119, 115, 58, 50, 48, 48, 53, 58, 48, 52, 58, 100, 105, 115, 99, 111, 118, 101, 114, 121, 60, 47, 119, 115, 97, 58, 84, 111, 62, 60, 119, 115, 97, 58, 65, 99, 116, 105, 111, 110, 62, 104, 116, 116, 112, 58, 47, 47, 115, 99, 104, 101, 109, 97, 115, 46, 120, 109, 108, 115, 111, 97, 112, 46, 111, 114, 103, 47, 119, 115, 47, 50, 48, 48, 53, 47, 48, 52, 47, 100, 105, 115, 99, 111, 118, 101, 114, 121, 47, 80, 114, 111, 98, 101, 60, 47, 119, 115, 97, 58, 65, 99, 116, 105, 111, 110, 62, 60, 119, 115, 97, 58, 77, 101, 115, 115, 97, 103, 101, 73, 68, 62, 117, 114, 110, 58, 117, 117, 105, 100, 58, 99, 101, 48, 52, 100, 97, 100, 48, 45, 53, 100, 50, 99, 45, 52, 48, 50, 54, 45, 57, 49, 52, 54, 45, 49, 97, 97, 98, 102, 99, 49, 101, 52, 49, 49, 49, 60, 47, 119, 115, 97, 58, 77, 101, 115, 115, 97, 103, 101, 73, 68, 62, 60, 47, 115, 111, 97, 112, 58, 72, 101, 97, 100, 101, 114, 62, 60, 115, 111, 97, 112, 58, 66, 111, 100, 121, 62, 60, 119, 115, 100, 58, 80, 114, 111, 98, 101, 62, 60, 119, 115, 100, 58, 84, 121, 112, 101, 115, 62, 119, 115, 100, 112, 58, 68, 101, 118, 105, 99, 101, 60, 47, 119, 115, 100, 58, 84, 121, 112, 101, 115, 62, 60, 47, 119, 115, 100, 58, 80, 114, 111, 98, 101, 62, 60, 47, 115, 111, 97, 112, 58, 66, 111, 100, 121, 62, 60, 47, 115, 111, 97, 112, 58, 69, 110, 118, 101, 108, 111, 112, 101, 62, 10},
		},
	}
	UDP_PROBE_XDMCP = UdpProbe{
		Service: Service{
			Slug:        "xdmcp",
			Name:        "XDMCP",
			Long:        "X Display Manager Control Protocol (XDMCP)",
			Description: ``, // TODO
			References: []string{
				"https://www.speedguide.net/port.php?port=177",
				"https://wikipedia.org/wiki/X_display_manager",
			},
		},
		Ports: []UdpPort{
			177,
		},
		Payloads: [][]byte{
			{0, 1, 0, 2, 0, 1, 0},
		},
	}
	UDP_PROBE_KERBEROS = UdpProbe{
		Service: Service{
			Slug:        "kerberos",
			Name:        "Kerberos",
			Long:        "Kerberos Key Distribution Center (KDC) Server.",
			Description: `Kerberos is a network authentication protocol that uses tickets to allow devices to prove their identity to others in a secure manner.`,
			References: []string{
				"https://www.speedguide.net/port.php?port=88",
				"https://wikipedia.org/wiki/Kerberos_(protocol)",
			},
		},
		Ports: []UdpPort{
			88,
		},
		Payloads: [][]byte{
			{106, 122, 48, 120, 161, 3, 2, 1, 5, 162, 3, 2, 1, 10, 164, 108, 48, 106, 160, 7, 3, 5, 0, 64, 0, 0, 0, 161, 17, 48, 15, 160, 3, 2, 1, 1, 161, 8, 48, 6, 27, 4, 110, 109, 97, 112, 162, 6, 27, 4, 116, 101, 115, 116, 163, 25, 48, 23, 160, 3, 2, 1, 2, 161, 16, 48, 14, 27, 6, 107, 114, 98, 116, 103, 116, 27, 4, 116, 101, 115, 116, 165, 17, 24, 15, 50, 48, 50, 50, 49, 49, 49, 51, 50, 49, 52, 53, 48, 50, 90, 167, 6, 2, 4, 9, 74, 118, 129, 168, 14, 48, 12, 2, 1, 18, 2, 1, 17, 2, 1, 16, 2, 1, 23},
		},
	}
	UDP_PROBE_LANTRONIX_DISCOVER = UdpProbe{
		Service: Service{
			Slug:        "lantronix-discover",
			Name:        "Lantronix Discovery",
			Description: `Lantronix Discovery allows applications on the network to discover Lantronix serial-to-ethernet devices`,
			References: []string{
				"https://www.speedguide.net/port.php?port=30718",
				"http://wiki.lantronix.com/wiki/Lantronix_Discovery_Protocol",
				"https://www.shodan.io/search?query=product%3A%22Lantronix+X90%22",
			},
		},
		Ports: []UdpPort{
			30718,
		},
		Payloads: [][]byte{
			{0, 0, 0, 246},
		},
	}
	UDP_PROBE_IKE = UdpProbe{
		Service: Service{
			Slug:        "ike",
			Name:        "IKE",
			Long:        "Internet Key Exchange (IKE)",
			Description: ``, // TODO
			References: []string{
				"https://www.speedguide.net/port.php?port=500",
				"https://www.speedguide.net/port.php?port=4500",
				"https://wikipedia.org/wiki/Internet_Key_Exchange",
			},
		},
		Ports: []UdpPort{
			500,
			4500, // alt
		},
		Payloads: [][]byte{
			{91, 94, 100, 192, 62, 153, 181, 17, 0, 0, 0, 0, 0, 0, 0, 0, 1, 16, 2, 0, 0, 0, 0, 0, 0, 0, 1, 80, 0, 0, 1, 52, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 1, 40, 1, 1, 0, 8, 3, 0, 0, 36, 1, 1},
		},
	}
	UDP_PROBE_RADIUS = UdpProbe{
		Service: Service{
			Slug:        "radius",
			Name:        "RADIUS",
			Long:        "Remote Authentication Dial-In User Service (RADIUS)",
			Description: ``, // TODO
			References: []string{
				"https://www.speedguide.net/port.php?port=1645",
				"https://wikipedia.org/wiki/RADIUS",
				"https://datatracker.ietf.org/doc/html/rfc2865",
				"https://datatracker.ietf.org/doc/html/rfc2866",
			},
		},
		Ports: []UdpPort{
			1812,
			1645, // alt
			1813, // alt
		},
		Payloads: [][]byte{
			{1, 103, 0, 87, 64, 182, 100, 219, 245, 214, 129, 178, 173, 189, 23, 105, 81, 81, 24, 200, 1, 7, 115, 116, 101, 118, 101, 2, 18, 219, 198, 196, 183, 88, 190, 20, 240, 5, 179, 135, 124, 158, 47, 182, 1, 4, 6, 192, 168, 0, 28, 5, 6, 0, 0, 0, 123, 80, 18, 95, 15, 134, 71, 232, 200, 155, 216, 129, 54, 66, 104, 252, 208, 69, 50, 79, 12, 2, 102, 0, 10, 1, 115, 116, 101, 118, 101},
		},
	}
	UDP_PROBE_DTLS = UdpProbe{
		Service: Service{
			Slug:        "dtls",
			Name:        "DTLS",
			Long:        "Datagram Transport Layer Security (DTLS)",
			Description: ``, // TODO
			References: []string{
				"https://wikipedia.org/wiki/Datagram_Transport_Layer_Security",
			},
		},
		Ports: []UdpPort{
			4433,
			443,
			5061,
		},
		Payloads: [][]byte{
			// Client Hello
			{22, 254, 255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 192, 1, 0, 0, 180, 0, 0, 0, 0, 0, 0, 0, 180, 254, 253, 248, 44, 180, 106, 34, 111, 3, 224, 108, 208, 21, 167, 243, 221, 105, 38, 76, 57, 33, 178, 146, 212, 68, 8, 190, 161, 212, 128, 59, 247, 131, 245, 0, 0, 0, 56, 192, 44, 192, 48, 0, 159, 204, 169, 204, 168, 204, 170, 192, 43, 192, 47, 0, 158, 192, 36, 192, 40, 0, 107, 192, 35, 192, 39, 0, 103, 192, 10, 192, 20, 0, 57, 192, 9, 192, 19, 0, 51, 0, 157, 0, 156, 0, 61, 0, 60, 0, 53, 0, 47, 0, 255, 1, 0, 0, 82, 0, 11, 0, 4, 3, 0, 1, 2, 0, 10, 0, 12, 0, 10, 0, 29, 0, 23, 0, 30, 0, 25, 0, 24, 0, 35, 0, 0, 0, 22, 0, 0, 0, 23, 0, 0, 0, 13, 0, 42, 0, 40, 4, 3, 5, 3, 6, 3, 8, 7, 8, 8, 8, 9, 8, 10, 8, 11, 8, 4, 8, 5, 8, 6, 4, 1, 5, 1, 6, 1, 3, 3, 3, 1, 3, 2, 4, 2, 5, 2, 6, 2},
			// Application Data
			//{23, 254, 253, 0, 1, 0, 0, 0, 0, 0, 1, 0, 28, 167, 91, 224, 171, 30, 64, 156, 76, 119, 148, 10, 115, 252, 132, 41, 190, 13, 89, 31, 157, 102, 138, 175, 118, 237, 242, 247, 73}
		},
	}
	UDP_PROBE_ENIP = UdpProbe{
		Service: Service{
			Slug:        "enip",
			Name:        "EtherNet/IP",
			Long:        "Ethernet/IP (Industrial Protocol)",
			Description: `Ethernet/IP is an industrial network protocol that adapts the Common Industrial Protocol (CIP) to standard Ethernet. EtherNet/IP is one of the leading industrial protocols in the United States and is widely used in a range of industries including factory, hybrid and process.`,
			References: []string{
				"https://www.speedguide.net/port.php?port=44818",
				"https://www.speedguide.net/port.php?port=2222",
				"https://wikipedia.org/wiki/EtherNet/IP",
			},
		},
		Ports: []UdpPort{
			44818, // also TCP
			2222,  // alt
		},
		Payloads: [][]byte{
			// ENIP List Identity request
			{99, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
	}
	UDP_PROBE_ETHER_SIO = UdpProbe{
		Service: Service{
			Slug:        "ether-s-io",
			Name:        "Ether-S-I/O",
			Description: ``, // TODO
		},
	}
	UDP_PROBE_BITTORRENT = UdpProbe{
		Service: Service{
			Slug:        "bittorrent",
			Name:        "BitTorrent DHT",
			Long:        "BitTorrent Distributed Hash Table (DHT)",
			Description: ``, // TODO
			References: []string{
				"https://www.speedguide.net/port.php?port=6881",
				"http://www.bittorrent.org/beps/bep_0005.html",
			},
		},
		Ports: []UdpPort{
			6881,
		},
		Payloads: [][]byte{
			{100, 49, 58, 97, 100, 50, 58, 105, 100, 50, 48, 58, 97, 98, 99, 100, 101, 102, 103, 104, 105, 106, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 101, 49, 58, 113, 52, 58, 112, 105, 110, 103, 49, 58, 116, 50, 58, 97, 97, 49, 58, 121, 49, 58, 113, 101}, // DHT ping <~ http://www.bittorrent.org/beps/bep_0005.html
		},
	}
	UDP_PROBE_MSRPC = UdpProbe{
		Service: Service{
			Slug:        "msrpc",
			Name:        "MSRPC",
			Long:        "Microsoft Windows Remote Procedure Call (MSRPC)",
			Description: ``, // TODO
			References: []string{
				"https://www.speedguide.net/port.php?port=135",
				"https://wikipedia.org/wiki/Microsoft_RPC",
			},
		},
		Ports: []UdpPort{
			135,
		},
		Payloads: [][]byte{
			// String binding: ncadg_ip_udp
			{5, 0, 11, 3, 16, 0, 0, 0, 72, 0, 0, 0, 1, 0, 0, 0, 184, 16, 184, 16, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 1, 0, 196, 254, 252, 153, 96, 82, 27, 16, 187, 203, 0, 170, 0, 33, 52, 122, 0, 0, 0, 0, 4, 93, 136, 138, 235, 28, 201, 17, 159, 232, 8, 0, 43, 16, 72, 96, 2, 0, 0, 0},
		},
	}
	UDP_PROBE_L2TP = UdpProbe{
		Service: Service{
			Slug: "l2tp",
			Name: "L2TP",
			Long: "Layer 2 Tunneling Protocol (L2TP)",
			References: []string{
				"https://www.speedguide.net/port.php?port=1701",
				"https://wikipedia.org/wiki/Layer_2_Tunneling_Protocol",
			},
		},
		Ports: []UdpPort{
			1702,
		},
		Payloads: [][]byte{
			{200, 2, 0, 107, 0, 0, 0, 0, 0, 0, 0, 0, 128, 8, 0, 0, 0, 0, 0, 1, 128, 8, 0, 0, 0, 2, 1, 0, 128, 10, 0, 0, 0, 3, 0, 0, 0, 3, 128, 10, 0, 0, 0, 4, 0, 0, 0, 0, 128, 8, 0, 0, 0, 6, 6, 144, 128, 16, 0, 0, 0, 7, 49, 48, 57, 46, 54, 46, 49, 46, 55, 50, 128, 19, 0, 0, 0, 8, 120, 101, 108, 101, 114, 97, 110, 99, 101, 46, 99, 111, 109, 128, 8, 0, 0, 0, 9, 74, 50, 128, 8, 0, 0, 0, 10, 0, 4},
		},
	}
	UDP_PROBE_NFS = UdpProbe{
		Service: Service{
			Slug:        "nfs",
			Name:        "NFS",
			Long:        "Network File System (NFS)",
			Description: ``, // TODO
			References: []string{
				"https://www.speedguide.net/port.php?port=2049",
				"https://wikipedia.org/wiki/Network_File_System",
			},
		},
		Ports: []UdpPort{
			2049,
		},
		Payloads: [][]byte{
			{101, 202, 114, 40, 0, 0, 0, 0, 0, 0, 0, 2, 0, 1, 134, 163, 0, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
	}
	UDP_PROBE_SLP = UdpProbe{
		Service: Service{
			Slug: "slp",
			Name: "SLP",
			Long: "Service Location Protocol (SLP)",
			References: []string{
				"https://www.speedguide.net/port.php?port=427",
				"https://wikipedia.org/wiki/Service_Location_Protocol",
			},
		},
		Ports: []UdpPort{
			427,
		},
		Payloads: [][]byte{
			{2, 1, 0, 0, 54, 32, 0, 0, 0, 0, 0, 1, 0, 2, 101, 110, 0, 0, 0, 21, 115, 101, 114, 118, 105, 99, 101, 58, 115, 101, 114, 118, 105, 99, 101, 45, 97, 103, 101, 110, 116, 0, 7, 100, 101, 102, 97, 117, 108, 116, 0, 0, 0, 0},
		},
	}
	UDP_PROBE_HART_IP = UdpProbe{
		Service: Service{
			Slug:        "hart-ip",
			Name:        "HART-IP",
			Long:        "Highway Addressable Remote Transducer Industrial Protocol",
			Description: "The HART Communications Protocol (Highway Addressable Remote Transducer Protocol) is an early implementation of Fieldbus, a digital industrial automation protocol. Its most notable advantage is that it can communicate over legacy wiring.",
			References: []string{
				"https://www.speedguide.net/port.php?port=5094",
				"https://wikipedia.org/wiki/Highway_Addressable_Remote_Transducer_Protocol",
				"https://www.shodan.io/explore/category/industrial-control-systems",
				"https://wiki.wireshark.org/HART-IP",
			},
		},
		Ports: []UdpPort{
			5094,
		},
		Payloads: [][]byte{
			{1, 0, 0, 0, 0, 2, 0, 13, 1, 0, 0, 117, 48},
		},
	}
	UDP_PROBE_HID_DISCOVERY = UdpProbe{
		Service: Service{
			Slug:        "hid-discovery",
			Name:        "HID Discovery",
			Description: ``,
			References: []string{
				"https://www.exploit-db.com/exploits/44992",
				"https://www.pcworld.com/article/420378/flaw-in-popular-door-controllers-allow-hackers-to-easily-unlock-secure-doors.html",
				"https://www.shodan.io/search?query=%22HID+VertX%22+port%3A4070",
			},
		},
		Ports: []UdpPort{
			4070,
		},
		Payloads: [][]byte{
			{100, 105, 115, 99, 111, 118, 101, 114, 59, 48, 49, 51, 59},
		},
	}
	UDP_PROBE_PCWORX = UdpProbe{
		Service: Service{
			Slug:        "pcworx",
			Name:        "PCWorx",
			Description: ``,
			References: []string{
				"https://www.speedguide.net/port.php?port=1962",
				"https://sergiusechel.medium.com/misconfiguration-in-ilc-gsm-gprs-devices-leaves-over-1-200-ics-devices-vulnerable-to-attacks-over-82c2d4a91561",
			},
		},
		Ports: []UdpPort{
			1962,
		},
		Payloads: [][]byte{
			{1, 1, 0, 26, 0, 0, 0, 0, 120, 128, 0, 3, 0, 12, 73, 66, 69, 84, 72, 48, 49, 78, 48, 95, 77, 0},
		},
	}
	/*
		UDP_PROBE_EPL = UdpProbe{ // [TODO]
			Service: Service{
				Slug:        "epl",
				Name:        "EPL",
				Long:        "Ethernet Powerlink",
				Description: "Real-time protocol for standard Ethernet. It is an open protocol managed by the Ethernet POWERLINK Standardization Group (EPSG). It was introduced by Austrian automation company B&R in 2001",
				References: []string{
					"https://wiki.wireshark.org/SampleCaptures#ethernet-powerlink-v1",
					"https://wiki.wireshark.org/SampleCaptures#ethernet-powerlink-v2",
				},
			},
		}

			UDP_PROBE_SYNCHROPHASOR = UdpProbe{
				Service: Service{
					Slug: ""
				},
				Payloads: [][]byte{
					{170, 65, 0, 18, 0, 60, 72, 153, 144, 154, 0, 144, 46, 18, 0, 5, 22, 138}, // [TODO] test this
				},
			}
	*/
)

var (
	PROBES_ALL = map[uint16][]UdpProbe{
		17: {UDP_PROBE_QOTD},
		19: {UDP_PROBE_CHARGEN},
		53: {UDP_PROBE_DNS},
		69: {UDP_PROBE_TFTP},
		//80:  {UDP_PROBE_DTLS},
		88:  {UDP_PROBE_KERBEROS},
		111: {UDP_PROBE_PORTMAP},
		123: {UDP_PROBE_NTP},
		135: {UDP_PROBE_MSRPC},
		137: {UDP_PROBE_NETBIOS},
		//138: {UDP_PROBE_NETBIOS}, // check this
		//139: {UDP_PROBE_NETBIOS}, // check this
		161: {UDP_PROBE_SNMP},
		162: {UDP_PROBE_SNMP},
		177: {UDP_PROBE_XDMCP},
		389: {UDP_PROBE_CLDAP},
		427: {UDP_PROBE_SLP},
		443: {UDP_PROBE_DTLS},
		500: {UDP_PROBE_IKE},
		520: {UDP_PROBE_RIP},
		523: {UDP_PROBE_IBM_DB2},
		623: {UDP_PROBE_IPMI},
		//853:  {UDP_PROBE_DTLS},
		1194: {UDP_PROBE_OPENVPN},
		1434: {UDP_PROBE_MSSQL},
		//1604:  {UDP_PROBE_WINFRAME},
		1645: {UDP_PROBE_RADIUS},
		1701: {UDP_PROBE_L2TP},
		1812: {UDP_PROBE_RADIUS},
		1813: {UDP_PROBE_RADIUS},
		1900: {UDP_PROBE_SSDP, UDP_PROBE_UPNP},
		1962: {UDP_PROBE_PCWORX},
		2049: {UDP_PROBE_NFS},
		2222: {UDP_PROBE_ENIP},
		//2362:  {UDP_PROBE_DIGI_ADDP},
		3283: {UDP_PROBE_ARD},
		3389: {UDP_PROBE_RDP},
		//3391: {UDP_PROBE_DTLS},
		3470: {UDP_PROBE_STUN},
		3478: {UDP_PROBE_STUN},
		3702: {UDP_PROBE_WSD},
		4433: {UDP_PROBE_DTLS},
		4500: {UDP_PROBE_IKE},
		//4740: {UDP_PROBE_DTLS},
		5000: {UDP_PROBE_SSDP, UDP_PROBE_UPNP},
		5060: {UDP_PROBE_SIP},
		5061: {UDP_PROBE_DTLS},
		5093: {UDP_PROBE_SENTINEL},
		5094: {UDP_PROBE_HART_IP},
		//5349: {UDP_PROBE_DTLS},
		5351: {UDP_PROBE_NAT_PMP},
		5353: {UDP_PROBE_MDNS},
		5632: {UDP_PROBE_PCA},
		5683: {UDP_PROBE_COAP},
		/*5684:  {UDP_PROBE_DTLS},
		5868:  {UDP_PROBE_DTLS},
		6514:  {UDP_PROBE_DTLS},
		6636:  {UDP_PROBE_DTLS},*/
		6881: {UDP_PROBE_BITTORRENT},
		// 8232:  {UDP_PROBE_DTLS},
		10001: {UDP_PROBE_UBIQUITI},
		10161: {UDP_PROBE_SNMP},
		10162: {UDP_PROBE_SNMP},
		11211: {UDP_PROBE_MEMCACHE},
		/*12346: {UDP_PROBE_DTLS},
		12446: {UDP_PROBE_DTLS},
		12546: {UDP_PROBE_DTLS},
		12646: {UDP_PROBE_DTLS},
		12746: {UDP_PROBE_DTLS},
		12846: {UDP_PROBE_DTLS},
		12946: {UDP_PROBE_DTLS},
		13046: {UDP_PROBE_DTLS},*/
		17185: {UDP_PROBE_WDBRPC},
		19302: {UDP_PROBE_STUN},
		20000: {UDP_PROBE_DNP3},
		27015: {UDP_PROBE_STEAM_HLTV},
		30718: {UDP_PROBE_LANTRONIX_DISCOVER},
		44818: {UDP_PROBE_ENIP},
		47808: {UDP_PROBE_BACNET},
		53413: {UDP_PROBE_NETIS},
	}
	PROBES_ICS = map[uint16][]UdpProbe{
		1962:  {UDP_PROBE_PCWORX},
		2222:  {UDP_PROBE_ENIP},
		4070:  {UDP_PROBE_HID_DISCOVERY},
		4800:  {UDP_PROBE_MOXA_NPORT},
		5006:  {UDP_PROBE_MELSEC_Q},
		5094:  {UDP_PROBE_HART_IP},
		9600:  {UDP_PROBE_FINS},
		17185: {UDP_PROBE_WDBRPC},
		20000: {UDP_PROBE_DNP3},
		30718: {UDP_PROBE_LANTRONIX_DISCOVER},
		44818: {UDP_PROBE_ENIP},
		47808: {UDP_PROBE_BACNET},
	}
)
