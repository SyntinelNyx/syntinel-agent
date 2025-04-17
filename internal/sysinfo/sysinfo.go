package sysinfo

import (
	"github.com/SyntinelNyx/syntinel-agent/internal/logger"
	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/disk"
	"github.com/shirou/gopsutil/v4/mem"
)

func SysInfo() string {

	cpuPercent, err := cpu.Percent(1, true)
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
		logger.Info("Memory Info: %v", memStat)
	}

	diskStat, err := disk.Usage("/")
	if err != nil {
		logger.Error("Error getting disk info: %v", err)
	} else {
		logger.Info("Disk Info: %v", diskStat)
	
	}
	
	return 
}
