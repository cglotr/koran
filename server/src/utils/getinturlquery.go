package utils

import (
	"errors"
	"strconv"
)

// GetIntURLQuery .
func GetIntURLQuery(values map[string][]string, key string) (int, error) {
	value, ok := values[key]
	if !ok {
		return 0, errors.New("getIntURLQuery: key isn't in map")
	}
	if len(value) < 1 {
		return 0, errors.New("getIntURLQuery: empty value")
	}
	parsed, err := strconv.Atoi(value[0])
	if err != nil {
		return 0, err
	}
	return parsed, nil
}
