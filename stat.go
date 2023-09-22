package nsdg1000t

// StatConnectivity /api/stat/connectivity response
type StatConnectivity struct {
	ConnectivityType string `json:"connectivity_type"`
	FirmwareVersion  string `json:"firmware_version"`
	MapEEnable       string `json:"map_e_enable"`
	PonConnectivity  string `json:"pon_connectibity"`
	ReceivedPower    string `json:"received_power"`
	WanAddressIPv4   string `json:"wan_address_ipv4"`
	WanAddressIPv6   string `json:"wan_address_ipv6"`
	WanSSID5G        string `json:"wan_ssid_5g"`
	WanSSID24G       string `json:"wan_ssid_24g"`
}

func (c *Client) StatConnectivity() (*StatConnectivity, error) {
	url := "http://" + c.host + "/api/stat/connectivity&csrf_token=" + c.csrfToken
	var stat StatConnectivity
	err := c.getEndpointData(url, &stat)
	return &stat, err
}

// StatEnable24G /api/stat/24g_enable
type StatEnable24G struct {
	Enabled string `json:"24g_enabled"`
	Ssid    string `json:"24g_ssid"`
}

func (c *Client) StatEnable24G() (*StatEnable24G, error) {
	url := "http://" + c.host + "/api/stat/24g_enable&csrf_token=" + c.csrfToken
	var stat StatEnable24G
	err := c.getEndpointData(url, &stat)
	return &stat, err
}

// StatEnable5G /api/stat/5g_enable
type StatEnable5G struct {
	Enabled string `json:"5g_enabled"`
	Ssid    string `json:"5g_ssid"`
}

func (c *Client) StatEnable5G() (*StatEnable5G, error) {
	url := "http://" + c.host + "/api/stat/5g_enable&csrf_token=" + c.csrfToken
	var stat StatEnable5G
	err := c.getEndpointData(url, &stat)
	return &stat, err
}

// StatWPS /api/stat/wps
type StatWPS struct {
	WpsFunction5G  bool `json:"5g_wps_function"`
	WpsFunction24G bool `json:"24g_wps_function"`
}

func (c *Client) StatWPS() (*StatWPS, error) {
	url := "http://" + c.host + "/api/stat/wps&csrf_token=" + c.csrfToken
	var stat StatWPS
	err := c.getEndpointData(url, &stat)
	return &stat, err
}
