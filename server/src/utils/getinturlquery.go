package utils

import (
	"fmt"
	"strconv"
)

// GetIntURLQuery .
func GetIntURLQuery(values map[string][]string, key string) (int, error) {
	value, ok := values[key]
	if !ok {
		return 0, fmt.Errorf("key %#v is not in %#v", key, values)
	}
	if len(value) < 1 {
		return 0, fmt.Errorf("key %#v is mapped to an empty value", key)
	}
	parsed, err := strconv.Atoi(value[0])
	if err != nil {
		return 0, err
	}
	return parsed, nil
}
