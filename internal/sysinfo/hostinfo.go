package sysinfo

import (
	"encoding/json"

	"github.com/SyntinelNyx/syntinel-agent/internal/logger"
	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/disk"
	"github.com/shirou/gopsutil/v4/host"
	"github.com/shirou/gopsutil/v4/mem"
)

func CpuInfo() string {
	CpuStat, err := cpu.Info()
	if err != nil {
		logger.Fatal("Error getting CPU info: %v", err)
	}
	
    // Extract general CPU info from the first CPU stat
    generalCpuInfo := map[string]interface{}{
        "VendorID":  CpuStat[0].VendorID,
        "ModelName": CpuStat[0].ModelName,
        "Mhz":       CpuStat[0].Mhz,
        "CacheSize": CpuStat[0].CacheSize,
    }

    generalCpuInfo["Cores"] = len(CpuStat)

	data, err := json.MarshalIndent(&generalCpuInfo, "", "  ")
	if err != nil {
		logger.Error("Error marshaling CPU info to JSON: %v", err)
	}
	logger.Info("CPU info: %s", string(data))
	return string(data)

}

func MemInfo() string {
	MemStat, err := mem.VirtualMemory()
	if err != nil {
		logger.Error("Error getting memory info: %v", err)
	}

	data, err := json.MarshalIndent(&MemStat.Total, "", "  ")
	if err != nil {
		logger.Error("Error marshaling memory info to JSON: %v", err)
	}
	logger.Info("Memory info: %s", string(data))
	return string(data)
}

func DiskInfo() string {
	DiskStat, err := disk.Usage("/")
	if err != nil {
		logger.Error("Error getting disk info: %v", err)
	}
	data, err := json.MarshalIndent(&DiskStat.Total, "", "  ")
	if err != nil {
		logger.Error("Error marshaling disk info to JSON: %v", err)
	}
	logger.Info("Disk info: %s", string(data))
	return string(data)
}

func HostInfo() string {
	HostStat, err := host.Info()

	if err != nil {
		logger.Error("Error getting host info: %v", err)
	}
	data, err := json.MarshalIndent(&HostStat, "", "  ")
	if err != nil {
		logger.Error("Error marshaling host info to JSON: %v", err)
	}
	logger.Info("Host info: %s", string(data))
	return string(data)
}

func CombinedInfo() string {
	combinedData := map[string]interface{}{
		"CPU":    json.RawMessage(CpuInfo()),
		"Memory": json.RawMessage(MemInfo()),
		"Disk":   json.RawMessage(DiskInfo()),
		"Host":   json.RawMessage(HostInfo()),
	}

	data, err := json.MarshalIndent(combinedData, "", "  ")
	if err != nil {
		logger.Error("Error marshaling combined info to JSON: %v", err)
	}
	logger.Info("Combined system info: %s", string(data))
	return string(data)
}
