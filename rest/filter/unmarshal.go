package filter

import (
	"bytes"
	"encoding/json"
	"reflect"

	"github.com/ez-connect/go-rest/core"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	keywords = []string{
		"$and",
		"$or",
		"$eq",
		"$gt",
		"$lt",
	}
)

func validateMap(v map[string]interface{}, rv reflect.Value) bool {
	for key, value := range v {
		field := rv.FieldByName(key)
		if !field.IsValid() && core.IndexOf(keywords, key) < 0 {
			return false
		}

		switch value := value.(type) {
		case map[string]interface{}:
			if !validateMap(value, rv) {
				return false
			}
		case []interface{}:
			if !validateArray(value, rv) {
				return false
			}
		}

	}
	return true
}

func validateArray(v []interface{}, rv reflect.Value) bool {
	for value := range v {
		switch v[value].(type) {
		case map[string]interface{}:
			if !validateMap(v[value].(map[string]interface{}), rv) {
				return false
			}
		case []interface{}:
			if !validateArray(v[value].([]interface{}), rv) {
				return false
			}
		}
	}
	return true
}

func Unmarshal(query string, v interface{}) map[string]interface{} {
	buf := bytes.NewBufferString(query)
	if !json.Valid(buf.Bytes()) {
		return nil
	}
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return nil
	}

	res := bson.M{}
	err := json.Unmarshal(buf.Bytes(), &res)
	if err != nil {
		return nil
	}

	// validate query
	s := rv.Elem()
	if validateMap(res, s) {
		return res
	} else {
		return nil
	}
}
