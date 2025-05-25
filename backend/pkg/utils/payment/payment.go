package payment

import (
	"encoding/json"
	"net/url"
	"sort"
	"strings"

	"github.com/thanhoanganhtuan/go-ecommerce-backend-api/internal/vo"
)

func SortObject(obj vo.VNPayParams) vo.VNPayParams {
	sorted := make(vo.VNPayParams)
	keys := make([]string, 0, len(obj))

	// Get all keys
	for key := range obj {
		keys = append(keys, key)
	}

	// Sort keys
	sort.Strings(keys)

	// Create sorted map
	for _, key := range keys {
		sorted[key] = obj[key]
	}

	return sorted
}

func BuildQueryString(params vo.VNPayParams, encode bool) string {
	var parts []string

	keys := make([]string, 0, len(params))
	for key := range params {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	for _, key := range keys {
		value := params[key]
		if encode {
			key = url.QueryEscape(key)
			value = strings.ReplaceAll(url.QueryEscape(value), "%20", "+")
		}
		parts = append(parts, key+"="+value)
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
