package nsdg1000t

// SupportSystem /api/support/system
type SupportSystem struct {
	BootloaderVersion     string `json:"bootloader_version"`
	BuildTime             string `json:"build_time"`
	CpuUsage              string `json:"cpu_usage"`
	CurrentFirmwareImage  string `json:"current_firmware_image"`
	DataTime              string `json:"data_time"`
	FirmwareVersion       string `json:"firmware_version"`
	FW0Checksum           string `json:"fw0_checksum"`
	FW0ImageSize          string `json:"fw0_image_size"`
	FW0Version            string `json:"fw0_version"`
	FW1Checksum           string `json:"fw1_checksum"`
	FW1ImageSize          string `json:"fw1_image_size"`
	FW1Version            string `json:"fw1_version"`
	HardwareTypeVersion   string `json:"hardware_type_version"`
	MemoryUsage           string `json:"memory_usage"`
	ModelName             string `json:"model_name"`
	OmciSoftwareVersion   string `json:"omci_software_version"`
	RebootCause           string `json:"reboot_cause"`
	SerialNumber          string `json:"serial_number"`
	UptimeSinceLastReboot string `json:"uptime_since_last_reboot"`
	WirelessDriverVersion string `json:"wireless_driver_version"`
}

func (c *Client) SupportSystem() (*SupportSystem, error) {
	url := "http://" + c.host + "/api/support/system&csrf_token=" + c.csrfToken
	var s SupportSystem
	err := c.getEndpointData(url, &s)
	return &s, err
}

// SupportWan /api/support/wan
type SupportWan struct {
	Firewall              string `json:"firewall"`
	Gateway               string `json:"gateway"`
	IpAddress             string `json:"ip_address"`
	Ipv4Interface         string `json:"ipv4_interface"`
	Ipv4RecvdBytes        int64  `json:"ipv4_recvd_bytes"`
	Ipv4RxDrops           int    `json:"ipv4_rx_drops"`
	Ipv4RxErrs            int    `json:"ipv4_rx_errs"`
	Ipv4RxPkts            int    `json:"ipv4_rx_pkts"`
	Ipv4TransBytes        int64  `json:"ipv4_trans_bytes"`
	Ipv4TxDrops           int    `json:"ipv4_tx_drops"`
	Ipv4TxErrs            int    `json:"ipv4_tx_errs"`
	Ipv4TxPkts            int    `json:"ipv4_tx_pkts"`
	Ipv4Type              string `json:"ipv4_type"`
	Ipv6DnsAddress        string `json:"ipv6_dns_address"`
	Ipv6Gateway           string `json:"ipv6_gateway"`
	Ipv6Interface         string `json:"ipv6_interface"`
	Ipv6LinkGlobalAddress string `json:"ipv6_link_global_address"`
	Ipv6LinkLocalAddress  string `json:"ipv6_link_local_address"`
	Ipv6PrefixDelegation  string `json:"ipv6_prefix_delegation"`
	Ipv6RecvdBytes        int64  `json:"ipv6_recvd_bytes"`
	Ipv6RxDrops           int    `json:"ipv6_rx_drops"`
	Ipv6RxErrs            int    `json:"ipv6_rx_errs"`
	Ipv6RxPkts            int    `json:"ipv6_rx_pkts"`
	Ipv6TransBytes        int64  `json:"ipv6_trans_bytes"`
	Ipv6TxDrops           int    `json:"ipv6_tx_drops"`
	Ipv6TxErrs            int    `json:"ipv6_tx_errs"`
	Ipv6TxPkts            int    `json:"ipv6_tx_pkts"`
	Ipv6Type              string `json:"ipv6_type"`
	PrimaryDnsIpAddress   string `json:"primary_dns_ip_address"`
	SecondaryDnsIpAddress string `json:"secondary_dns_ip_address"`
}

func (c *Client) SupportWan() (*SupportWan, error) {
	url := "http://" + c.host + "/api/support/wan&csrf_token=" + c.csrfToken
	var s SupportWan
	err := c.getEndpointData(url, &s)
	return &s, err
}

// SupportLAN /api/support/lan
type SupportLan struct {
	DHCPServer          string `json:"dhcp_server"`
	DHCPv6Server        string `json:"dhcpv6_server"`
	IPAddress           string `json:"ip_address"`
	IPv6Address         string `json:"ipv6_address"`
	MACAddress          string `json:"mac_address"`
	RouterAdvertisement string `json:"router_advertisement"`
}

func (c *Client) SupportLAN() (*SupportLan, error) {
	url := "http://" + c.host + "/api/support/lan&csrf_token=" + c.csrfToken
	var s SupportLan
	err := c.getEndpointData(url, &s)
	return &s, err
}

// SupportEthernetSwitch /api/support/ethernet_switch
type SupportEthernetSwitch struct {
	ID     string `json:"id"`
	Mode   string `json:"mode"`
	Speed  string `json:"speed"`
	Status string `json:"status"`
}

func (c *Client) SupportEthernetSwitch() (*[]SupportEthernetSwitch, error) {
	url := "http://" + c.host + "/api/support/ethernet_switch&csrf_token=" + c.csrfToken
	var s []SupportEthernetSwitch
	err := c.getEndpointData(url, &s)
	return &s, err
}

// /api/support/wifi_24g
type SupportWifi24G struct {
	Status     string `json:"status"`
	Ssid       string `json:"ssid"`
	MACAddress string `json:"mac_address"`
	Security   string `json:"security"`
	Channel    string `json:"channel"`
	Bandwidth  string `json:"bandwidth"`
}

func (c *Client) SupportWifi24G() (*SupportWifi24G, error) {
	url := "http://" + c.host + "/api/support/wifi_24g&csrf_token=" + c.csrfToken
	var s SupportWifi24G
	err := c.getEndpointData(url, &s)
	return &s, err
}

// /api/support/wifi_5g
type SupportWifi5G struct {
	Status     string `json:"status"`
	Ssid       string `json:"ssid"`
	MACAddress string `json:"mac_address"`
	Security   string `json:"security"`
	Channel    string `json:"channel"`
	Bandwidth  string `json:"bandwidth"`
}

func (c *Client) SupportWifi5G() (*SupportWifi5G, error) {
	url := "http://" + c.host + "/api/support/wifi_5g&csrf_token=" + c.csrfToken
	var s SupportWifi5G
	err := c.getEndpointData(url, &s)
	return &s, err
}

// SupportDhcpLease /api/support/dhcp_lease
type SupportDhcpLease struct {
	ID         int    `json:"id"`
	MACAddress string `json:"mac_address"`
	HostName   string `json:"host_name"`
	IPAddress  string `json:"ip_address"`
	Type       string `json:"type"`
	Status     string `json:"status"`
	ExpiresIn  string `json:"expires_in"`
}

func (c *Client) SupportDhcpLease() (*[]SupportDhcpLease, error) {
	url := "http://" + c.host + "/api/support/dhcp_lease&csrf_token=" + c.csrfToken
	var s []SupportDhcpLease
	err := c.getEndpointData(url, &s)
	return &s, err
}

// SupportLanStatistics /api/support/lan_statistics
type SupportLanStatistics struct {
	ID            int    `json:"id"`
	LANInterface  string `json:"lan_interface"`
	LANRecvdBytes int64  `json:"lan_recvd_bytes"`
	LANRxPkts     int    `json:"lan_rx_pkts"`
	LANRxErrs     int    `json:"lan_rx_errs"`
	LANRxDrops    int    `json:"lan_rx_drops"`
	LANTransBytes int64  `json:"lan_trans_bytes"`
	LANTxPkts     int    `json:"lan_tx_pkts"`
	LANTxDrops    int    `json:"lan_tx_drops"`
}

func (c *Client) SupportLanStatistics() (*[]SupportLanStatistics, error) {
	url := "http://" + c.host + "/api/support/lan_statistics&csrf_token=" + c.csrfToken
	var s []SupportLanStatistics
	err := c.getEndpointData(url, &s)
	return &s, err
}

// SupportWlanStatistics /api/support/wlan_statistics
type SupportWlanStatistics struct {
	ID             int    `json:"id"`
	Band           string `json:"band"`
	SSID           string `json:"ssid"`
	WLANRecvdBytes int64  `json:"wlan_recvd_bytes"`
	WLANRxPkts     int    `json:"wlan_rx_pkts"`
	WLANRxErrs     int    `json:"wlan_rx_errs"`
	WLANRxDrops    int    `json:"wlan_rx_drops"`
	WLANTransBytes int64  `json:"wlan_trans_bytes"`
	WLANTxPkts     int    `json:"wlan_tx_pkts"`
	WLANTxDrops    int    `json:"wlan_tx_drops"`
}

func (c *Client) SupportWlanStatistics() (*[]SupportWlanStatistics, error) {
	url := "http://" + c.host + "/api/support/wlan_statistics&csrf_token=" + c.csrfToken
	var s []SupportWlanStatistics
	err := c.getEndpointData(url, &s)
	return &s, err
}
