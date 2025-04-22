package sysinfo

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/disk"
	"github.com/shirou/gopsutil/v4/mem"
)

type SysInfo struct {
	CpuUsage  float64 `json:"cpuUsage"`
	MemUsage  Memory  `json:"memoryUsage"`
	DiskUsage Disk    `json:"diskUsage"`
}

type Memory struct {
	Total       uint64  `json:"total"`
	Available   uint64  `json:"available"`
	Used        uint64  `json:"used"`
	UsedPercent float64 `json:"usedPercent"`
}

type Disk struct {
	Total       uint64  `json:"total"`
	Free        uint64  `json:"free"`
	Used        uint64  `json:"used"`
	UsedPercent float64 `json:"usedPercent"`
}

func Monitor() (*SysInfo, error) {
	cpuStat, err := cpu.Percent(2*time.Second, false)
	if err != nil {
		return nil, fmt.Errorf("error measuring cpu usage: %v", err)
	}

	memStat, err := mem.VirtualMemory()
	if err != nil {
		return nil, fmt.Errorf("error measuring memory usage: %v", err)
	}

	diskStat, err := disk.Usage("/")
	if err != nil {
		return nil, fmt.Errorf("error measuring disk usage: %v", err)
	}

	sysInfo := SysInfo{
		CpuUsage: cpuStat[0],
		MemUsage: Memory{
			Total:       memStat.Total,
			Available:   memStat.Available,
			Used:        memStat.Used,
			UsedPercent: memStat.UsedPercent,
		},
		DiskUsage: Disk{
			Total:       diskStat.Total,
			Free:        diskStat.Free,
			Used:        diskStat.Used,
			UsedPercent: diskStat.UsedPercent,
		},
	}

	return &sysInfo, nil
}
