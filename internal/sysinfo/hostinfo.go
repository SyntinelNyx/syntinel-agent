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
	cpuInfo, err := CpuInfo()
	if err != nil {
		combinedData["CPU"] = err.Error()
	}
	combinedData["CPU"] = json.RawMessage(cpuInfo)

	// Collect memory info
	memInfo, err := MemInfo()
	if err != nil {
		combinedData["Memory"] = err.Error()
	}
	combinedData["Memory"] = json.RawMessage(memInfo)

	// Collect disk info
	diskInfo, err := DiskInfo()
	if err != nil {
		combinedData["Disk"] = err.Error()
	}
	combinedData["Disk"] = json.RawMessage(diskInfo)

	// Collect host info
	hostInfo, err := HostInfo()
	if err != nil {
		combinedData["Host"] = err.Error()
	}
	combinedData["Host"] = json.RawMessage(hostInfo)

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
