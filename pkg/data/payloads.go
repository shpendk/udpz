package data

var (
	UDP_PAYLOAD_MSSQL_PING = [][]byte{
		{2},
	}
	UDP_PAYLOAD_IPMI_RMCP = [][]byte{
		{6, 0, 255, 7, 0, 0, 0, 0, 0, 0, 0, 0, 0, 9, 32, 24, 200, 129, 0, 56, 142, 4, 181},
	}
)
