package sysinfo

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSysInfo(t *testing.T) {
	info := SysInfo()
	assert.NotEmpty(t, info, "SysInfo should return non-empty JSON string")

	var result map[string]interface{}
	err := json.Unmarshal([]byte(info), &result)
	assert.NoError(t, err, "SysInfo should return valid JSON")

	_, ok := result["sysinfo"].(map[string]interface{})
	assert.True(t, ok, "SysInfo JSON should contain 'sysinfo' field")

	_, ok = result["node"].(map[string]interface{})
	assert.True(t, ok, "SysInfo JSON should contain 'node' field")

	_, ok = result["os"].(map[string]interface{})
	assert.True(t, ok, "SysInfo JSON should contain 'os' field")

	_, ok = result["kernel"].(map[string]interface{})
	assert.True(t, ok, "SysInfo JSON should contain 'kernel' field")

	_, ok = result["bios"].(map[string]interface{})
	assert.True(t, ok, "SysInfo JSON should contain 'bios' field")

	_, ok = result["cpu"].(map[string]interface{})
	assert.True(t, ok, "SysInfo JSON should contain 'cpu' field")

	_, ok = result["memory"].(map[string]interface{})
	assert.True(t, ok, "SysInfo JSON should contain 'memory' field")

	_, ok = result["storage"].([]interface{})
	assert.True(t, ok, "SysInfo JSON should contain 'storage' field")
}
