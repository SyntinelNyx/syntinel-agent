package sysinfo

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/SyntinelNyx/syntinel-agent/internal/logger"
	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/disk"
	"github.com/shirou/gopsutil/v4/mem"
)

func SysInfo() string {
	sysInfo := make(map[string]any)

	cpuPercent, err := cpu.Percent(2*time.Second, true)
	if err != nil {
		sysInfo["CPU"] = "Error retrieving CPU info"
		logger.Error("Error getting CPU percentage: %v", err)
	} else {
		for i := range cpuPercent {
			sysInfo["CPU"] = cpuPercent
			logger.Info("CPU Usage: %.2f%%", cpuPercent[i])
		}
	}

	memStat, err := mem.VirtualMemory()
	if err != nil {
		sysInfo["Memory"] = "Error retrieving memory info"
		logger.Error("Error getting memory info: %v", err)
	} else {
		memoryInfo := map[string]any{
			"Total":       memStat.Total,
			"Available":   memStat.Available,
			"Used":        memStat.Used,
			"UsedPercent": memStat.UsedPercent,
		}
		sysInfo["Memory"] = memoryInfo
		logger.Info("Memory Info: %v", memoryInfo)
	}

	diskStat, err := disk.Usage("/")
	if err != nil {
		sysInfo["Disk"] = "Error retrieving disk info"
		logger.Error("Error getting disk info: %v", err)
	} else {
		diskInfo := map[string]any{
			"Total":       diskStat.Total,
			"Free":        diskStat.Free,
			"Used":        diskStat.Used,
			"UsedPercent": diskStat.UsedPercent,
		}
		sysInfo["Disk"] = diskInfo
		logger.Info("Disk Info: %v", diskInfo)
	}

	jsonData, err := json.Marshal(sysInfo)
	if err != nil {
		logger.Error("Error marshalling system info to JSON: %v", err)
		return fmt.Sprintf("Error marshalling system info to JSON: %v", err)
	}

	return string(jsonData)
}
