package sysinfo

import (
	"encoding/json"

	"github.com/shirou/gopsutil/v4/host"
	"github.com/shirou/gopsutil/v4/net"
	"github.com/SyntinelNyx/syntinel-agent/internal/logger"
)

func HostInfo() string {
	HostStat, err := host.Info()
	
	if err != nil {
		logger.Fatal("Error getting host info: %v", err)
	}
	data, err := json.MarshalIndent(&HostStat, "", "  ")
	if err != nil {
		logger.Fatal("Error marshaling host info to JSON: %v", err)
	}
	logger.Info("Host info: %s", string(data))
	return string(data)
}

func NetInfo() string {
	NetStat, err := net.Interfaces()
	if err != nil {
		logger.Fatal("Error getting network interfaces: %v", err)
	}
	data, err := json.MarshalIndent(&NetStat, "", "  ")
	if err != nil {
		logger.Fatal("Error marshaling network interfaces to JSON: %v", err)
	}
	logger.Info("Network interfaces: %s", string(data))
	return string(data)
}