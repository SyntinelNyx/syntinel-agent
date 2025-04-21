package sysinfo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCpuInfo(t *testing.T) {
	data, err := CpuInfo()
	assert.NoError(t, err, "CpuInfo should not return an error")
	assert.NotEmpty(t, data, "CpuInfo should return non-empty data")
}

func TestMemInfo(t *testing.T) {
	data, err := MemInfo()
	assert.NoError(t, err, "MemInfo should not return an error")
	assert.NotEmpty(t, data, "MemInfo should return non-empty data")
}

func TestDiskInfo(t *testing.T) {
	data, err := DiskInfo()
	assert.NoError(t, err, "DiskInfo should not return an error")
	assert.NotEmpty(t, data, "DiskInfo should return non-empty data")
}

func TestHostInfo(t *testing.T) {
	data, err := HostInfo()
	assert.NoError(t, err, "HostInfo should not return an error")
	assert.NotEmpty(t, data, "HostInfo should return non-empty data")
}

func TestCombinedInfo(t *testing.T) {
	data, err := CombinedInfo()
	assert.NoError(t, err, "CombinedInfo should not return an error")
	assert.NotEmpty(t, data, "CombinedInfo should return non-empty data")
}