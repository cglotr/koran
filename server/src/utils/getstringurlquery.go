package utils

import (
	"fmt"
)

// GetStringURLQuery .
func GetStringURLQuery(values map[string][]string, key string) (string, error) {
	value, ok := values[key]
	if !ok {
		return "", fmt.Errorf("key %#v is not in %#v", key, values)
	}
	if len(value) < 1 {
		return "", fmt.Errorf("key %#v is mapped to an empty value", key)
	}
	return value[0], nil
}
