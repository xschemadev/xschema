package compliance

import (
	_ "embed"
	"encoding/json"
	"strings"
	"sync"
)

//go:embed adapter-limitations.json
var adapterLimitationsData []byte

// AdapterLimitation defines unsupported patterns for a specific adapter
type AdapterLimitation struct {
	UnsupportedPatterns []string `json:"unsupported_patterns"`
	Reason              string   `json:"reason"`
}

// AdapterLimitations maps adapter names to their limitations
type AdapterLimitations map[string]AdapterLimitation

// cached data (loaded once)
var (
	loadLimitationsOnce sync.Once
	adapterLimitations  AdapterLimitations
)

// loadAdapterLimitations parses the embedded JSON once
func loadAdapterLimitations() AdapterLimitations {
	loadLimitationsOnce.Do(func() {
		adapterLimitations = make(AdapterLimitations)
		json.Unmarshal(adapterLimitationsData, &adapterLimitations)
	})
	return adapterLimitations
}

// GetAdapterLimitation returns the limitation config for a specific adapter, if any
func GetAdapterLimitation(adapterName string) *AdapterLimitation {
	limitations := loadAdapterLimitations()
	if lim, ok := limitations[adapterName]; ok {
		return &lim
	}
	return nil
}

// MatchesUnsupportedPattern checks if a test group description matches any unsupported pattern for the adapter
func MatchesUnsupportedPattern(adapterName, groupDescription string) (bool, string) {
	lim := GetAdapterLimitation(adapterName)
	if lim == nil {
		return false, ""
	}

	for _, pattern := range lim.UnsupportedPatterns {
		if strings.Contains(groupDescription, pattern) {
			return true, lim.Reason
		}
	}

	return false, ""
}
