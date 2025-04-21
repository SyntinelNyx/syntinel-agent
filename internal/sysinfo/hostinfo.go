package sysinfo

import (
	"encoding/json"
	"fmt"

	"github.com/SyntinelNyx/syntinel-agent/internal/logger"
	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/disk"
	"github.com/shirou/gopsutil/v4/host"
	"github.com/shirou/gopsutil/v4/mem"
)

func CpuInfo() (string, error) {
	CpuStat, err := cpu.Info()
	if err != nil {
		return "", fmt.Errorf("Error getting CPU info: %v", err)
	}
	
    // Extract general CPU info from the first CPU stat
    generalCpuInfo := map[string]any{
        "VendorID":  CpuStat[0].VendorID,
        "ModelName": CpuStat[0].ModelName,
        "Mhz":       CpuStat[0].Mhz,
        "CacheSize": CpuStat[0].CacheSize,
    }

    generalCpuInfo["Cores"] = len(CpuStat)

	data, err := json.MarshalIndent(&generalCpuInfo, "", "  ")
	if err != nil {
		return "", fmt.Errorf("Error marshaling CPU info to JSON: %v", err)
	}
	logger.Info("CPU info: %s", string(data))
	return string(data), nil

}

func MemInfo() (string, error) {
	MemStat, err := mem.VirtualMemory()
	if err != nil {
		return "", fmt.Errorf("Error getting memory info: %v", err)
	}

	data, err := json.MarshalIndent(&MemStat.Total, "", "  ")
	if err != nil {
		return "", fmt.Errorf("Error marshaling memory info to JSON: %v", err)
	}
	logger.Info("Memory info: %s", string(data))
	return string(data), nil
}

func DiskInfo() (string, error) {
	DiskStat, err := disk.Usage("/")
	if err != nil {
		return "", fmt.Errorf("Error getting disk info: %v", err)
	}
	data, err := json.MarshalIndent(&DiskStat.Total, "", "  ")
	if err != nil {
		return "", fmt.Errorf("Error marshaling disk info to JSON: %v", err)
	}
	logger.Info("Disk info: %s", string(data))
	return string(data), nil
}

func HostInfo() (string, error) {
	HostStat, err := host.Info()

	if err != nil {
		return "", fmt.Errorf("Error getting host info: %v", err)
	}
	data, err := json.MarshalIndent(&HostStat, "", "  ")
	if err != nil {
		return "", fmt.Errorf("Error marshaling host info to JSON: %v", err)
	}
	logger.Info("Host info: %s", string(data))
	return string(data), nil
}

func CombinedInfo() (string, error) {
    combinedData := make(map[string]any)
    
    // Collect CPU info
    if cpuInfo, err := CpuInfo(); err == nil {
        combinedData["CPU"] = json.RawMessage(cpuInfo)
    } else {
        logger.Warn("Skipping CPU info: %v", err)
    }

    // Collect memory info
    if memInfo, err := MemInfo(); err == nil {
        combinedData["Memory"] = json.RawMessage(memInfo)
    } else {
        logger.Warn("Skipping memory info: %v", err)
    }

    // Collect disk info
    if diskInfo, err := DiskInfo(); err == nil {
        combinedData["Disk"] = json.RawMessage(diskInfo)
    } else {
        logger.Warn("Skipping disk info: %v", err)
    }

    // Collect host info
    if hostInfo, err := HostInfo(); err == nil {
        combinedData["Host"] = json.RawMessage(hostInfo)
    } else {
        logger.Warn("Skipping host info: %v", err)
    }

    // Return error only if no data was collected
    if len(combinedData) == 0 {
        return "", fmt.Errorf("Failed to collect any system information")
    }

    data, err := json.MarshalIndent(combinedData, "", "  ")
    if err != nil {
        return "", fmt.Errorf("Error marshaling combined info to JSON: %v", err)
    }
    
    logger.Info("Combined system info collected (%d components)", len(combinedData))
    return string(data), nil
}
