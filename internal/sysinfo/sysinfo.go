package sysinfo

import (
	"encoding/json"
	"log"
	"os/user"


	"github.com/zcalusic/sysinfo"
)

func SysInfo() string {
	current, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	if current.Uid != "0" {
		log.Fatal("Requires superuser privilege")
	}

	var si sysinfo.SysInfo

	si.GetSysInfo()

	// Marshal to JSON
	data, err := json.MarshalIndent(&si, "", "  ")
	if err != nil {
		log.Fatalf("Error marshaling hardware info to JSON: %v", err)
	}

	return string(data)
}

