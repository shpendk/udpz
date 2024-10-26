package scan

import "regexp"

var (
	REGEX_HOSTNAME = regexp.MustCompile(`^[0-9A-Za-z_.-]{1,253}$`)
)

const (
	CAPTURE_SNAP_LEN   = 262144
	STATE_UNRESPONSIVE = 0
	STATE_RESPONSIVE   = 1
	STATE_CLOSED       = 2

	ADDR_IPV4
	ADDR_IPV6
)
