package sysinfo

import (
    "encoding/json"
    "testing"

    "github.com/stretchr/testify/assert"
)

func TestSysInfo(t *testing.T) {
    // Call the SysInfo function
    data, err := SysInfo()
    
    // Check that there was no error
    assert.NoError(t, err, "SysInfo should not return an error")
    
    // Check that the returned data is not empty
    assert.NotEmpty(t, data, "SysInfo should return non-empty data")
    
    // Try to parse the returned JSON data
    var sysInfo map[string]interface{}
    err = json.Unmarshal([]byte(data), &sysInfo)
    assert.NoError(t, err, "SysInfo should return valid JSON")
    
    // Check that the expected keys exist in the returned data
    assert.Contains(t, sysInfo, "CPU", "SysInfo should include CPU information")
    assert.Contains(t, sysInfo, "Memory", "SysInfo should include Memory information")
    assert.Contains(t, sysInfo, "Disk", "SysInfo should include Disk information")
}

func TestSysInfoContents(t *testing.T) {
    // Call the SysInfo function
    data, err := SysInfo()
    assert.NoError(t, err)
    
    var sysInfo map[string]any
    err = json.Unmarshal([]byte(data), &sysInfo)
    assert.NoError(t, err)
    
    // Only perform content checks if the values are not error messages
    if cpuInfo, ok := sysInfo["CPU"].([]any); ok {
        assert.NotEmpty(t, cpuInfo, "CPU information should not be empty")
    }
    
    if memInfo, ok := sysInfo["Memory"].(map[string]any); ok {
        assert.Contains(t, memInfo, "Total", "Memory info should include Total")
        assert.Contains(t, memInfo, "Available", "Memory info should include Available")
        assert.Contains(t, memInfo, "Used", "Memory info should include Used")
        assert.Contains(t, memInfo, "UsedPercent", "Memory info should include UsedPercent")
    }
    
    if diskInfo, ok := sysInfo["Disk"].(map[string]any); ok {
        assert.Contains(t, diskInfo, "Total", "Disk info should include Total")
        assert.Contains(t, diskInfo, "Free", "Disk info should include Free")
        assert.Contains(t, diskInfo, "Used", "Disk info should include Used")
        assert.Contains(t, diskInfo, "UsedPercent", "Disk info should include UsedPercent")
    }
}