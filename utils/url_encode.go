package utils

import (
	"fmt"
	"net/url"
	"strings"
)

func EncodeMap(params map[string]interface{}) string {
	var encodedParams []string

	for key, value := range params {
		encodedParams = append(encodedParams, EncodeValue(key, value))
	}

	return strings.Join(encodedParams, "&")
}

func EncodeValue(key string, value interface{}) string {
	switch v := value.(type) {
	case string:
		return url.QueryEscape(key) + "=" + url.QueryEscape(v)
	case int:
		return url.QueryEscape(key) + "=" + url.QueryEscape(fmt.Sprintf("%d", v))
	case []string:
		var encodedValues []string
		for _, item := range v {
			encodedValues = append(encodedValues, url.QueryEscape(item))
			encodedValues = append(encodedValues, url.QueryEscape(item))
		}
		return url.QueryEscape(key) + "=" + strings.Join(encodedValues, "&"+url.QueryEscape(key)+"=")
	case map[string]string:
		var encodedValues []string
		for nestedKey, nestedValue := range v {
			encodedValues = append(encodedValues, url.QueryEscape(key)+"["+url.QueryEscape(nestedKey)+"]="+url.QueryEscape(nestedValue))
		}
		return strings.Join(encodedValues, "&")
	default:
		return ""
	}
}
