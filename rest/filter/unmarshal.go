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
		"$in",
	}

	objectIdType = reflect.TypeOf(primitive.NewObjectID())
)

func validateMap(v map[string]interface{}, rv reflect.Value, mustBeObjectId bool) bson.M {
	result := bson.M{}
	for key, value := range v {
		fieldName := strings.Title(key)
		field := rv.FieldByName(fieldName)
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

			// get name of field from bson
			st, found := rv.Type().FieldByName(fieldName)
			if found {
				bsonInfo := strings.Split(st.Tag.Get("bson"), ",")
				if len(bsonInfo) > 0 && bsonInfo[0] != "" {
					key = bsonInfo[0]
				}
			}
		}

		switch value := value.(type) {
		case map[string]interface{}:
			if result[key] = validateMap(value, rv, isObjectId || mustBeObjectId); result[key] == nil {
				return nil
			}
		case []interface{}:
			if result[key] = validateArray(value, rv, isObjectId || mustBeObjectId); result[key] == nil {
				return nil
			}
		case string:
			if isObjectId || mustBeObjectId {
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
			if isObjectId || mustBeObjectId {
				return nil
			} else {
				result[key] = value
			}
		}

	}
	return result
}

func validateArray(v []interface{}, rv reflect.Value, mustBeObjectId bool) []interface{} {
	result := make([]interface{}, 0)
	for value := range v {
		switch x := v[value].(type) {
		case map[string]interface{}:
			t := validateMap(x, rv, mustBeObjectId)
			if t == nil {
				return nil
			}
			result = append(result, t)
		case []interface{}:
			t := validateArray(x, rv, mustBeObjectId)
			if t == nil {
				return nil
			}
			result = append(result, t)
		case string:
			if mustBeObjectId {
				objectId, err := primitive.ObjectIDFromHex(x)
				if err != nil {
					return nil
				}
				result = append(result, objectId)
			} else {
				result = append(result, x)
			}
		default:
			if !mustBeObjectId {
				result = append(result, x)
			} else {
				return nil
			}
		}
	}
	return result
}

func UnmarshalQueryParam(query string, v interface{}) map[string]interface{} {
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
	return validateMap(res, rv.Elem(), false)
}

func UnmarshalPathParams(params map[string]string, v interface{}) map[string]interface{} {
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return nil
	}

	rv = rv.Elem()

	res := bson.M{}
	for key, value := range params {
		fieldName := strings.Title(key)
		field := rv.FieldByName(fieldName)
		if !field.IsValid() {
			return nil
		}

		// get name of field from bson
		st, found := rv.Type().FieldByName(fieldName)
		if found {
			bsonInfo := strings.Split(st.Tag.Get("bson"), ",")
			if len(bsonInfo) > 0 && bsonInfo[0] != "" {
				key = bsonInfo[0]
			}
		}

		isObjectId := false

		// check field is objectId
		switch field.Kind() {
		case reflect.Ptr:
			isObjectId = field.Type().Elem() == objectIdType
		case reflect.Struct:
			isObjectId = field.Type() == objectIdType
		}

		if isObjectId {
			objectId, err := primitive.ObjectIDFromHex(value)
			if err != nil {
				return nil
			}
			switch field.Kind() {
			case reflect.Ptr:
				res[key] = &objectId
			case reflect.Struct:
				res[key] = objectId
			}
		} else {
			res[key] = value
		}
	}

	return res
}
