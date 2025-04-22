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

func SysInfo() (string, error) {
	sysInfo := make(map[string]any)

	cpuPercent, err := cpu.Percent(2*time.Second, true)
	if err != nil {
		sysInfo["CPU"] = "Error retrieving CPU info"
	} else {
		formattedCPU := make([]string, len(cpuPercent))
		for i, percent := range cpuPercent {
			formattedCPU[i] = fmt.Sprintf("CPU[%d] %.2f", i, percent)
		}
		sysInfo["CPU"] = formattedCPU
		logger.Info("CPU Info: %v", formattedCPU)
	}

	memStat, err := mem.VirtualMemory()
	if err != nil {
		sysInfo["Memory"] = "Error retrieving memory info"
	} else {
		sysInfo["Memory"] = fmt.Sprintf("Memory Info: Total: %v, Available: %v, Used: %v, UsedPercent: %.2f ", memStat.Total, memStat.Available, memStat.Used, memStat.UsedPercent)
		logger.Info("Memory Info: %v", sysInfo["Memory"])
	}

	diskStat, err := disk.Usage("/")
	if err != nil {
		sysInfo["Disk"] = "Error retrieving disk info"
	} else {
		sysInfo["Disk"] = fmt.Sprintf("Disk Info: Total: %v, Free: %v, Used: %v, UsedPercent: %.2f ", diskStat.Total, diskStat.Free, diskStat.Used, diskStat.UsedPercent)
		logger.Info("Disk Info: %v", sysInfo["Disk"])
	}

	jsonData, err := json.Marshal(sysInfo)
	if err != nil {
		return "", fmt.Errorf("Error marshalling system info to JSON: %v", err)
	}

	return string(jsonData), nil
}
