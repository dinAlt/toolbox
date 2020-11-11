package maps

import (
	"errors"
	"reflect"

	"github.com/dinalt/toolbox/tryget"
)

type Map map[string]interface{}

func New(interfacedMap interface{}) (Map, error) {
	if interfacedMap == nil {
		return nil, nil
	}

	if v, ok := interfacedMap.(map[string]interface{}); ok {
		return Map(v), nil
	}

	return nil, errors.New("interfaced value is not a map")
}

func (m Map) ValueOrDefault(key string, defaultValue interface{}) interface{} {
	if m == nil {
		return defaultValue
	}

	v, ok := m[key]
	if !ok {
		return defaultValue
	}

	if defaultValue != interface{}(nil) &&
		reflect.TypeOf(defaultValue) != reflect.TypeOf(v) {
		return defaultValue
	}

	return v
}

func (m Map) Int64OrDefault(key string, defaultValue int64) int64 {
	if m == nil {
		return defaultValue
	}

	interfaced, ok := m[key]
	if !ok {
		return defaultValue
	}

	v, err := tryget.Int64(interfaced)

	if err != nil {
		return defaultValue
	}

	return v
}

func (m Map) Float64OrDefault(key string, defaultValue float64) float64 {
	if m == nil {
		return defaultValue
	}

	interfaced, ok := m[key]
	if !ok {
		return defaultValue
	}

	v, err := tryget.Float64(interfaced)

	if err != nil {
		return defaultValue
	}

	return v
}

func (m Map) StringOrDefault(key string, defaultValue string) string {
	if m == nil {
		return defaultValue
	}

	interfaced, ok := m[key]
	if !ok {
		return defaultValue
	}

	v, err := tryget.String(interfaced)

	if err != nil {
		return defaultValue
	}

	return v
}

func (m Map) BoolOrDefault(key string, defaultValue bool) bool {
	if m == nil {
		return defaultValue
	}

	interfaced, ok := m[key]
	if !ok {
		return defaultValue
	}

	v, err := tryget.Bool(interfaced)

	if err != nil {
		return defaultValue
	}

	return v
}
