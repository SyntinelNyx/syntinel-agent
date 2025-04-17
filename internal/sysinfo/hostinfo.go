package sysinfo

import (
	"encoding/json"

	"github.com/SyntinelNyx/syntinel-agent/internal/logger"
	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/disk"
	"github.com/shirou/gopsutil/v4/host"
	"github.com/shirou/gopsutil/v4/mem"
	"github.com/shirou/gopsutil/v4/net"
)

func CpuInfo() string {
	CpuStat, err := cpu.Info()
	if err != nil {
		logger.Fatal("Error getting CPU info: %v", err)
	}
	var individualCpuStats []map[string]interface{}

	for _, stat := range CpuStat {
		// Convert each InfoStat struct to a map for JSON marshaling
		cpuInfo := map[string]interface{}{
			"CPU":       stat.CPU,
			"VendorID":  stat.VendorID,
			"CoreID":    stat.CoreID,
			"ModelName": stat.ModelName,
			"Mhz":       stat.Mhz,
			"CacheSize": stat.CacheSize,
		}
		individualCpuStats = append(individualCpuStats, cpuInfo)
	}

	data, err := json.MarshalIndent(&individualCpuStats, "", "  ")
	if err != nil {
		logger.Fatal("Error marshaling CPU info to JSON: %v", err)
	}
	logger.Info("CPU info: %s", string(data))
	return string(data)

}

func MemInfo() string {
	MemStat, err := mem.VirtualMemory()
	if err != nil {
		logger.Fatal("Error getting memory info: %v", err)
	}

	data, err := json.MarshalIndent(&MemStat.Total, "", "  ")
	if err != nil {
		logger.Fatal("Error marshaling memory info to JSON: %v", err)
	}
	logger.Info("Memory info: %s", string(data))
	return string(data)
}

func DiskInfo() string {
	DiskStat, err := disk.Usage("/")
	if err != nil {
		logger.Fatal("Error getting disk info: %v", err)
	}
	data, err := json.MarshalIndent(&DiskStat.Total, "", "  ")
	if err != nil {
		logger.Fatal("Error marshaling disk info to JSON: %v", err)
	}
	logger.Info("Disk info: %s", string(data))
	return string(data)
}

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

func CombinedInfo() string {
    combinedData := map[string]interface{}{
        "CPU":     json.RawMessage(CpuInfo()),
        "Memory":  json.RawMessage(MemInfo()),
        "Disk":    json.RawMessage(DiskInfo()),
        "Host":    json.RawMessage(HostInfo()),
        "Network": json.RawMessage(NetInfo()),
    }

    data, err := json.MarshalIndent(combinedData, "", "  ")
    if err != nil {
        logger.Fatal("Error marshaling combined info to JSON: %v", err)
    }
    logger.Info("Combined system info: %s", string(data))
    return string(data)
}