package sysinfo

import (
	"time"

	"github.com/SyntinelNyx/syntinel-agent/internal/logger"
	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/disk"
	"github.com/shirou/gopsutil/v4/mem"
)

func SysInfo() string {

	cpuPercent, err := cpu.Percent(2*time.Second, true)
	if err != nil {
		logger.Error("Error getting CPU percentage: %v", err)
	} else {
		for i := range cpuPercent {
			logger.Info("CPU Usage: %.2f%%", cpuPercent[i])
		}
	}

	memStat, err := mem.VirtualMemory()
	if err != nil {
		logger.Error("Error getting memory info: %v", err)
	} else {
		memoryInfo := map[string]interface{}{
			"Total":       memStat.Total,
			"Available":   memStat.Available,
			"Used":        memStat.Used,
			"UsedPercent": memStat.UsedPercent,
		}
		logger.Info("Memory Info: %v", memoryInfo)
	}

	diskStat, err := disk.Usage("/")
	if err != nil {
		logger.Error("Error getting disk info: %v", err)
	} else {
		diskInfo := map[string]interface{}{
			"Total":       diskStat.Total,
			"Free":        diskStat.Free,
			"Used":        diskStat.Used,
			"UsedPercent": diskStat.UsedPercent,
		}
		logger.Info("Disk Info: %v", diskInfo)
	}

	return ""
}
