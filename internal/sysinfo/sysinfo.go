package sysinfo

import (
	"encoding/json"
	"os/user"

	"github.com/zcalusic/sysinfo"

	"github.com/SyntinelNyx/syntinel-agent/internal/logger"
)

func SysInfo() string {
	current, err := user.Current()
	if err != nil {
		logger.Fatal("error getting current user: %v", err)
	}

	if current.Uid != "0" {
		logger.Fatal("requires superuser privilege")
	}

	var si sysinfo.SysInfo
	si.GetSysInfo()

	data, err := json.MarshalIndent(&si, "", "  ")
	if err != nil {
		logger.Fatal("Error marshaling hardware info to JSON: %v", err)
	}

	return string(data)
}
