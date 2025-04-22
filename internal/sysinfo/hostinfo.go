package sysinfo

import (
	"fmt"

	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/disk"
	"github.com/shirou/gopsutil/v4/host"
	"github.com/shirou/gopsutil/v4/mem"
)

type HostInfo struct {
	Host   host.InfoStat `json:"host"`
	Cpu    CpuStat       `json:"cpu"`
	Memory uint64        `json:"memory"`
	Disk   uint64        `json:"disk"`
}

type CpuStat struct {
	VendorID  string  `json:"vendorId"`
	Cores     int32   `json:"cores"`
	ModelName string  `json:"modelName"`
	Mhz       float64 `json:"mhz"`
	CacheSize int32   `json:"cacheSize"`
}

func System() (*HostInfo, error) {
	hostStat, err := host.Info()
	if err != nil {
		return nil, fmt.Errorf("error getting host info: %v", err)
	}

	cpuStat, err := cpu.Info()
	if err != nil {
		return nil, fmt.Errorf("error getting cpu info: %v", err)
	}

	memStat, err := mem.VirtualMemory()
	if err != nil {
		return nil, fmt.Errorf("error getting memory info: %v", err)
	}

	diskStat, err := disk.Usage("/")
	if err != nil {
		return nil, fmt.Errorf("error getting disk info: %v", err)
	}

	host := HostInfo{
		Host: *hostStat,
		Cpu: CpuStat{
			VendorID:  cpuStat[0].VendorID,
			Cores:     cpuStat[0].Cores,
			ModelName: cpuStat[0].ModelName,
			Mhz:       cpuStat[0].Mhz,
			CacheSize: cpuStat[0].CacheSize,
		},
		Memory: memStat.Total,
		Disk:   diskStat.Total,
	}

	return &host, nil
}
