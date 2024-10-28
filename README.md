# UDPz

UDPz is a speedy, portable, cross-platform UDP port scanner written in Go.

## Methodology

UDPz was created to address the need for a fast and efficient tool to scan UDP services across multiple hosts. 
Traditional network scanning tools like `nmap` often provide slower UDP scanning capabilities, which can be a bottleneck for network administrators and security professionals.

UDPz aims to fill this gap by providing a robust solution that can be easily integrated into existing workflows due to his logging capabilities and flexible output options.

## Features

- **Concurrent Scanning**: Utilizes goroutines and channels to perform flexible concurrent scans, significantly speeding up the scanning process.
- **Structured Logging**: Uses `zerolog` for detailed and structured logging, making it easier to analyze scan results.
- **Flexible Target Resolution**: Supports loading IP addresses, CIDR ranges, and hostnames from arguments, or from a file.
- **Customizable Probes**: Allows for the definition of custom probes for different UDP services.
- **Error Handling**: Gracefully handles errors during scanning, ensuring the process continues even if some targets fail.


## Installation

> [!TIP]
> Docker installation is recommended for higher performance.

### Docker

1. Clone the repository:
```sh
git clone https://github.com/FalconOps-Cybersecurity/udpz.git
cd udpz
```

2. Build the Docker image:
```sh
docker build --tag="udpz" --network="host" .
```

3. Run the Docker container:
```sh
docker run --rm --network="host" --name="udpz" --volume="$PWD:/output" -it udpz [flags] [targets ...]
```

### Standard Installation

1. Clone the repository:
```sh
git clone https://github.com/FalconOps-Cybersecurity/udpz.git
cd udpz
```

2. Build the project:
```sh
go build
```

3. Run the binary:
```sh
./udpz [flags] [targets ...]
```

## Usage

```
Usage:
  udpz [flags] [IP|hostname|CIDR|file ...]

Flags:
  -v, --version             version for udpz
  -o, --output string       Save results to file
  -O, --log string          Output log messages to file
  -a, --append              Append results to output file (default true)
  -f, --format string       Output format [text, pretty, csv, tsv, json, yaml, auto] (default "auto")
  -L, --log-format string   Output log format [pretty, json, auto] (default "auto")
  -c, --host-tasks uint     Maximum Number of hosts to scan concurrently (default 10)
  -p, --port-tasks uint     Number of Concurrent scan tasks per host (default 50)
  -r, --retries uint        Number of probe retransmissions per probe (default 2)
  -t, --timeout uint        UDP Probe timeout in milliseconds (default 3000)
  -A, --all                 Scan all resolved addresses instead of just the first (default true)
  -D, --debug               Enable debug logging (Very noisy!)
  -T, --trace               Enable trace logging (Very noisy!)
  -q, --quiet               Disable info logging
  -s, --silent              Disable ALL logging
  -h, --help                help for udpz
```

The target argument(s) can be an IP address, hostname, CIDR, or file(s) containing targets.


### Examples

- Simple UDP scan of a single host with prettyfied output and no informational logging:
```
./udpz -f pretty localhost --quiet

╭───────────┬──────────┬───────┬─────────────┬─────────────────────╮
│ HOST      │ PORT     │ STATE │ SERVICE     │ PROBES              │
├───────────┼──────────┼───────┼─────────────┼─────────────────────┤
│ 127.0.0.1 │ 5353/UDP │ OPEN  │ mDNS        │ mDNS reverse lookup │
│ 127.0.0.1 │ 111/UDP  │ OPEN  │ Portmap/RPC │ Portmap RPC dump    │
╰───────────┴──────────┴───────┴─────────────┴─────────────────────╯
```

- UDP scan with custom number of workers and timeout with debug logging:
```
./udpz -f pretty -p 100 -t 2000 localhost --debug

2:31PM INF cmd/root.go:187 > Starting scanner
2:31PM DBG pkg/scan/scan.go:297 > Calculating unique probe count service_count=49
2:31PM DBG pkg/scan/scan.go:304 > Calculated unique probe count probe_count=93
2:31PM DBG pkg/scan/scan.go:310 > Calculated total probe count total_probes=279
2:31PM DBG pkg/scan/scan.go:316 > Resolving targets target_count=1
2:31PM DBG pkg/scan/scan.go:139 > Resolved target hostname addresses=1 target=localhost
2:31PM DBG pkg/scan/scan.go:389 > Port closed host=127.0.0.1 port=427 target=localhost
2:31PM DBG pkg/scan/scan.go:377 > Skipping closed port host=127.0.0.1 port=427 target=localhost

...
```

- Scan multiple hosts using a CIDR range:
```
./udpz -f pretty 10.10.14.0/24
```


## Supported Services

- Apple Remote Desktop (ARD)
- BitTorrent DHT
- Building Automation & Control Networks (BACNet)
- Character Generator Protocol (CharGen)
- Citrix WinFrame Remote Desktop Server
- Connectionless Lightweight Directory Access Protocol (CLDAP)
- Constrained Application Protocol (CoAP)
- Datagram Transport Layer Security (DTLS)
- Distributed Network Protocol 3 (DNP3)
- Domain Name System (DNS)
- EtherNet/IP
- Factory Interface Network Service (FINS)
- HID Discovery Protocol
- Highway Addressable Remote Transducer Industrial Protocol
- IBM-DB2
- Intelligent Platform Management Interface (IPMI)
- Internet Key Exchange (IKE)
- Kerberos
- Lantronix Discovery
- Layer 2 Tunneling Protocol (L2TP)
- Memcache
- Microsoft Structured Query Language (SQL) Server
- Microsoft Windows Remote Procedure Call (MSRPC)
- Mitsubishi MELSEC-Q
- Moxa NPort
- Multicast Domain Name System (mDNS)
- Network Address Translation Port Mapping Protocol (NAT-PMP)
- Network Basic Input/Output System (NetBIOS)
- Network File System (NFS)
- Network Time Protocol (NTP)
- OpenVPN (Virtual Private Networking)
- PCWorx
- Quote of the Day (QOTD)
- Remote Authentication Dial-In User Service (RADIUS)
- Remote Desktop Protocol (RDP) over UDP
- Remote Procedure Call (RPC)
- Routing Information Protocol next generation (RIPng)
- Routing Information Protocol (RIP)
- Service Location Protocol (SLP)
- Session Initiation Protocol (SIP)
- Session Traversal Utilities for NAT (STUN)
- Simple Network Management Protocol (SNMP) - v1, v2c, v3
- Symantec PCAnywhere
- Trivial File Transfer Protocol (TFTP)
- Ubiquiti Networks AirControl Management Discovery Protocol
- Universal Plug and Play (UPnP)
- VxWorks Wind Debug Agent ONCRPC
- Web Services Discovery (WSD)
- X Display Manager Control Protocol (XDMCP)

## Inspiration / Credits
- [Nmap](https://nmap.org/)
- [UDPx](https://github.com/nullt3r/udpx)