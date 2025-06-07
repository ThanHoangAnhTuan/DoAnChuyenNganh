package payment

import (
	"encoding/json"
	"net/url"
	"sort"
	"strings"
)

func SortObject(params map[string]string) map[string]string {
	// Get all keys
	keys := make([]string, 0, len(params))
	for key := range params {
		keys = append(keys, key)
	}

	// Sort keys
	sort.Strings(keys)

	// Create sorted map
	sorted := make(map[string]string)
	for _, key := range keys {
		// URL encode both key and value, replace %20 with +
		encodedKey := url.QueryEscape(key)
		encodedValue := strings.ReplaceAll(url.QueryEscape(params[key]), "%20", "+")
		sorted[encodedKey] = encodedValue
	}

	return sorted
}

func CreateSignData(params map[string]string) string {
	// Get sorted keys
	keys := make([]string, 0, len(params))
	for key := range params {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	// Create query string for signing
	var parts []string
	for _, key := range keys {
		parts = append(parts, key+"="+params[key])
	}

	return strings.Join(parts, "&")
}

func CreateQueryString(params map[string]string) string {
	// Get sorted keys
	keys := make([]string, 0, len(params))
	for key := range params {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	// Create query string
	var parts []string
	for _, key := range keys {
		parts = append(parts, key+"="+params[key])
	}

	return strings.Join(parts, "&")
}

func MakeAPIRequest(url string, data interface{}) (string, error) {
	// This is a simplified version - you should implement proper HTTP client
	// with timeout, retry logic, etc.
	_, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	return "API request logged", nil
}
