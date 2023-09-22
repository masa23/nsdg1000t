package main

import (
	"log"

	"github.com/k0kubun/pp/v3"
	"github.com/masa23/nsdg1000t"
)

func main() {
	host := "192.168.1.1"
	user := "admin"
	password := ""

	client := nsdg1000t.NewClient(host, user, password)

	if err := client.Login(); err != nil {
		log.Fatalf("failed to login: %v", err)
	}

	if err := client.GetCSRFToken(); err != nil {
		log.Fatalf("failed to get CSRF token: %v", err)
	}

	// StatConnectivity
	statConnectivity, err := client.StatConnectivity()
	if err != nil {
		log.Fatalf("failed to get stat connectivity: %v", err)
	}

	// StatEnable24G
	statEnable24G, err := client.StatEnable24G()
	if err != nil {
		log.Fatalf("failed to get stat enable 24G: %v", err)
	}

	// StatEnable5G
	statEnable5G, err := client.StatEnable5G()
	if err != nil {
		log.Fatalf("failed to get stat enable 5G: %v", err)
	}

	// StatWPS
	statWPS, err := client.StatWPS()
	if err != nil {
		log.Fatalf("failed to get stat WPS: %v", err)
	}

	// SupportLAN
	supportLAN, err := client.SupportLAN()
	if err != nil {
		log.Fatalf("failed to get support LAN: %v", err)
	}

	// SupportEthernetSwitch
	supportEthernetSwitchStatus, err := client.SupportEthernetSwitch()
	if err != nil {
		log.Fatalf("failed to get support ethernet switch status: %v", err)
	}

	// SupportWifi24G
	supportWifi24G, err := client.SupportWifi24G()
	if err != nil {
		log.Fatalf("failed to get support wifi 24G: %v", err)
	}

	// SupportWifi5G
	supportWifi5G, err := client.SupportWifi5G()
	if err != nil {
		log.Fatalf("failed to get support wifi 5G: %v", err)
	}

	// SupportDhcpLease
	supportDhcpLease, err := client.SupportDhcpLease()
	if err != nil {
		log.Fatalf("failed to get support DHCP lease: %v", err)
	}

	// SupportLanStatistics
	supportLanStatistics, err := client.SupportLanStatistics()
	if err != nil {
		log.Fatalf("failed to get support LAN statistics: %v", err)
	}

	// SupportWlanStatistics
	supportWlanStatistics, err := client.SupportWlanStatistics()
	if err != nil {
		log.Fatalf("failed to get support WLAN statistics: %v", err)
	}

	pp.Println("statConnectivity", statConnectivity)
	pp.Println("statEnable24G", statEnable24G)
	pp.Println("statEnable5G", statEnable5G)
	pp.Println("statWPS", statWPS)
	pp.Println("supportLAN", supportLAN)
	pp.Println("supportEthernetSwitchStatus", supportEthernetSwitchStatus)
	pp.Println("supportWifi24G", supportWifi24G)
	pp.Println("supportWifi5G", supportWifi5G)
	pp.Println("supportDhcpLease", supportDhcpLease)
	pp.Println("supportLanStatistics", supportLanStatistics)
	pp.Println("supportWlanStatistics", supportWlanStatistics)
}
