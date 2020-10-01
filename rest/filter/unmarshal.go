package filter

import (
	"bytes"
	"encoding/json"
	"reflect"
	"strings"

	"github.com/ez-connect/go-rest/core"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	keywords = []string{
		"$and",
		"$or",
		"$eq",
		"$gt",
		"$lt",
	}

	objectIdType = reflect.TypeOf(primitive.NewObjectID())
)

func validateMap(v map[string]interface{}, rv reflect.Value) bson.M {
	result := bson.M{}
	for key, value := range v {
		field := rv.FieldByName(key)
		if !field.IsValid() && core.IndexOf(keywords, key) < 0 {
			return nil
		}

		isObjectId := false
		if field.IsValid() {
			// check field is objectId
			switch field.Kind() {
			case reflect.Ptr:
				isObjectId = field.Type().Elem() == objectIdType
			case reflect.Struct:
				isObjectId = field.Type() == objectIdType
			}

			if isObjectId && reflect.TypeOf(value).Kind() != reflect.String {
				return nil
			}

			// get name of field from bson
			st, found := rv.Type().FieldByName(key)
			if found {
				bsonInfo := strings.Split(st.Tag.Get("bson"), ",")
				if len(bsonInfo) > 0 {
					key = bsonInfo[0]
				}
			}
		}

		switch value := value.(type) {
		case map[string]interface{}:
			if result[key] = validateMap(value, rv); result[key] == nil {
				return nil
			}
		case []interface{}:
			if result[key] = validateArray(value, rv); result[key] == nil {
				return nil
			}
		case string:
			if isObjectId {
				objectId, err := primitive.ObjectIDFromHex(value)
				if err != nil {
					return nil
				}
				switch field.Kind() {
				case reflect.Ptr:
					result[key] = &objectId
				case reflect.Struct:
					result[key] = objectId
				}
			} else {
				result[key] = value
			}
		default:
			result[key] = value
		}

	}
	return result
}

func validateArray(v []interface{}, rv reflect.Value) []interface{} {
	result := make([]interface{}, 0)
	for value := range v {
		switch x := v[value].(type) {
		case map[string]interface{}:
			t := validateMap(x, rv)
			if t == nil {
				return nil
			}
			result = append(result, t)
		case []interface{}:
			t := validateArray(x, rv)
			if t == nil {
				return nil
			}
			result = append(result, t)
		default:
			result = append(result, x)
		}
	}
	return result
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
	return validateMap(res, s)
}
