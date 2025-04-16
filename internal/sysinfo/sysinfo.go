package sysinfo

import (
	// "encoding/json"
	// "os/user"

	// "github.com/zcalusic/sysinfo"

	// "github.com/SyntinelNyx/syntinel-agent/internal/logger"

	"github.com/SyntinelNyx/syntinel-agent/internal/logger"
	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/disk"
	"github.com/shirou/gopsutil/v4/mem"
)

// func SysInfo() string {
// 	current, err := user.Current()
// 	if err != nil {
// 		logger.Fatal("error getting current user: %v", err)
// 	}

// 	if current.Uid != "0" {
// 		logger.Fatal("requires superuser privilege")
// 	}

// 	var si sysinfo.SysInfo
// 	si.GetSysInfo()

// 	data, err := json.MarshalIndent(&si, "", "  ")
// 	if err != nil {
// 		logger.Fatal("Error marshaling hardware info to JSON: %v", err)
// 	}

// 	return string(data)
// }

func SysInfo() string {
	// var si sysinfo.SysInfo
	// si.GetSysInfo()

	cpuStat, err := cpu.Info()
	if err != nil {
		// logger.Fatal("Error getting CPU info: %v", err)
	}

	memStat, err := mem.VirtualMemory()
	if err != nil {
		// logger.Fatal("Error getting memory info: %v", err)
	}

	diskStat, err := disk.Usage("/")
	if err != nil {
		// logger.Fatal("Error getting disk info: %v", err)
	}

	cpuPercent, err := cpu.Percent(0, false)
	if err != nil {
		logger.Error("Error getting CPU percentage: %v", err)
	} else if len(cpuPercent) > 0 {
		logger.Info("CPU Usage: %.2f%%", cpuPercent[0])
	}



	logger.Info("CPU: %s", cpuStat[0].Cores)
	logger.Info("Memory: %v", memStat.Total)
	logger.Info("Disk: %v", diskStat.Used)

	return string(cpuStat[0].ModelName) + " " + string(memStat.Total) + " " + string(diskStat.Used)

}
