package filter

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strconv"
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

func validateMap(v map[string]interface{}, rv reflect.Value, mustBeObjectId bool) (bson.M, error) {
	result := bson.M{}
	for key, value := range v {
		fieldName := strings.Title(key)
		field := rv.FieldByName(fieldName)
		if !field.IsValid() && core.IndexOf(keywords, key) < 0 {
			return nil, fmt.Errorf("Fields/keyword %s is not exist", key)
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

		var err error
		switch value := value.(type) {
		case map[string]interface{}:
			result[key], err = validateMap(value, rv, isObjectId || mustBeObjectId)
			if err != nil {
				return nil, err
			}
		case []interface{}:
			if result[key], err = validateArray(value, rv, isObjectId || mustBeObjectId); err != nil {
				return nil, err
			}
		case string:
			if isObjectId || mustBeObjectId {
				objectId, err := primitive.ObjectIDFromHex(value)
				if err != nil {
					return nil, err
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
				return nil, fmt.Errorf("Field %s must be objectid string", key)
			} else {
				result[key] = value
			}
		}

	}
	return result, nil
}

func validateArray(v []interface{}, rv reflect.Value, mustBeObjectId bool) ([]interface{}, error) {
	result := make([]interface{}, 0)
	for value := range v {
		switch x := v[value].(type) {
		case map[string]interface{}:
			t, err := validateMap(x, rv, mustBeObjectId)
			if err != nil {
				return nil, err
			}
			result = append(result, t)
		case []interface{}:
			t, err := validateArray(x, rv, mustBeObjectId)
			if err != nil {
				return nil, err
			}
			result = append(result, t)
		case string:
			if mustBeObjectId {
				objectId, err := primitive.ObjectIDFromHex(x)
				if err != nil {
					return nil, err
				}
				result = append(result, objectId)
			} else {
				result = append(result, x)
			}
		default:
			if !mustBeObjectId {
				result = append(result, x)
			} else {
				return nil, errors.New("value must be objectid")
			}
		}
	}
	return result, nil
}

func UnmarshalQueryParam(query string, v interface{}) (map[string]interface{}, error) {
	buf := bytes.NewBufferString(query)
	if !json.Valid(buf.Bytes()) {
		return nil, errors.New("Filter must be json")
	}
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return nil, errors.New("Internal server error")
	}

	res := bson.M{}
	err := json.Unmarshal(buf.Bytes(), &res)
	if err != nil {
		return nil, err
	}

	// validate query
	return validateMap(res, rv.Elem(), false)
}

func UnmarshalPathParams(params map[string]string, v interface{}) (map[string]interface{}, error) {
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return nil, errors.New("Internal server error")
	}

	rv = rv.Elem()

	res := bson.M{}
	for key, value := range params {
		fieldName := strings.Title(key)
		field := rv.FieldByName(fieldName)
		if !field.IsValid() {
			return nil, fmt.Errorf("Field %s is not exist", key)
		}

		// get name of field from bson
		st, found := rv.Type().FieldByName(fieldName)
		if found {
			bsonInfo := strings.Split(st.Tag.Get("bson"), ",")
			if len(bsonInfo) > 0 && bsonInfo[0] != "" {
				key = bsonInfo[0]
			}
		}
		// fmt.Println(fieldName)
		value, err := getValue(field.Type(), value)
		if err != nil {
			return nil, err
		}
		res[key] = value
	}

	return res, nil
}

func getValue(fieldType reflect.Type, value string) (interface{}, error) {
	if fieldType == objectIdType {
		objectId, err := primitive.ObjectIDFromHex(value)
		if err != nil {
			return nil, err
		}
		return objectId, nil
	}
	// fmt.Println(fieldType)
	var v interface{}
	var err error
	switch fieldType.Kind() {
	case reflect.Ptr:
		v, err = getValue(fieldType.Elem(), value)
		return &v, err
	// case reflect.Struct:
	// 	if fieldType == objectIdType {
	// 		if err != nil {
	// 			return nil, err
	// 		}
	// 		return objectId, nil
	// 	}
	case reflect.Slice:
		return getValue(fieldType.Elem(), value)
	case reflect.Float32:
		v, err = strconv.ParseFloat(value, 32)
	case reflect.Float64:
		v, err = strconv.ParseFloat(value, 64)
	case reflect.Int32:
		v, err = strconv.ParseInt(value, 10, 32)
	case reflect.Int64:
		v, err = strconv.ParseInt(value, 10, 64)
	case reflect.Bool:
		v, err = strconv.ParseBool(value)
	default:
		v = value
	}

	if err != nil {
		return nil, err
	}
	return v, nil
}
