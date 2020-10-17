package filter

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ValueType int

const (
	Value ValueType = iota
	Array
	Map
	Bool
	Regex
	String
	Any
)

var (
	keywords = map[string][]ValueType{
		// keywork: {ValueType, (Array)Element ValueType}
		"$and": {Array, Map},
		"$or":  {Array, Map},
		"$not": {Map},
		"$nor": {Array, Map},

		"$eq":  {Value},
		"$ne":  {Value},
		"$gt":  {Value},
		"$gte": {Value},
		"$lt":  {Value},
		"$lte": {Value},
		"$in":  {Array, Value},
		"$nin": {Array, Value},

		"$exist": {Value}, // Bool
	}

	objectIdType = reflect.TypeOf(primitive.NewObjectID())
)

func (me ValueType) String() string {
	return [...]string{"Value", "Array", "Map", "Bool", "Regex", "String", "Any"}[me]
}

func getField(rv reflect.Value, key string) (bool, string, reflect.Type) {
	keys := strings.Split(key, ".")
	field := reflect.Value{}
	fieldType := rv.Type()
	newKey := []string{}
	for i, k := range keys {
		fieldName := strings.Title(k)
		field = rv.FieldByName(fieldName)
		if !field.IsValid() {
			return false, key, nil
		}
		fieldType = field.Type()
		switch field.Kind() {
		case reflect.Ptr:
			if fieldType.Elem() == objectIdType && i < len(keys)-1 {
				return false, key, nil
			}
			fieldType = fieldType.Elem()
		case reflect.Struct:
			if fieldType == objectIdType && i < len(keys)-1 {
				return false, key, nil
			}
		case reflect.Slice:
			fieldType = fieldType.Elem()
			switch fieldType.Kind() {
			case reflect.Ptr:
				if fieldType.Elem() == objectIdType && i < len(keys)-1 {
					return false, key, nil
				}
				fieldType = fieldType.Elem()
			case reflect.Struct:
				if fieldType == objectIdType && i < len(keys)-1 {
					return false, key, nil
				}
			}
		}

		// get name of field from bson
		st, found := rv.Type().FieldByName(fieldName)
		if found {
			bsonInfo := strings.Split(st.Tag.Get("bson"), ",")
			if len(bsonInfo) > 0 && bsonInfo[0] != "" {
				newKey = append(newKey, bsonInfo[0])
			} else {
				newKey = append(newKey, fieldName)
			}
		}

		// update rv
		if i < len(keys)-1 {
			rv = reflect.New(fieldType).Elem()
			if rv.Kind() != reflect.Struct {
				return false, key, nil
			}
		}
	}
	if field.IsValid() {
		return field.IsValid(), strings.Join(newKey, "."), fieldType
	}
	return field.IsValid(), key, fieldType
}

func validateMap(v map[string]interface{}, rv reflect.Value, mustBeObjectId bool) (bson.M, error) {
	result := bson.M{}
	for key, value := range v {
		isObjectId := mustBeObjectId
		valueType := Any
		if keywords[key] == nil {
			isValid := false
			var fieldType reflect.Type
			isValid, key, fieldType = getField(rv, key)
			if !isValid {
				return nil, fmt.Errorf("\"%s\" is not exist", key)
			}
			isObjectId = (isObjectId || (fieldType == objectIdType))
		} else {
			valueType = keywords[key][0]
		}

		var err error
		switch value := value.(type) {
		case map[string]interface{}:
			if valueType != Map && valueType != Any {
				return nil, fmt.Errorf("Value of \"%s\" is must be %s", key, valueType.String())
			}
			result[key], err = validateMap(value, rv, isObjectId)
			if err != nil {
				return nil, fmt.Errorf("\"%s\": %s", key, err.Error())
			}
		case []interface{}:
			if valueType != Array && valueType != Any {
				return nil, fmt.Errorf("Value of \"%s\" is must be %s", key, valueType.String())
			}
			if result[key], err = validateArray(value, rv, isObjectId, keywords[key][1]); err != nil {
				return nil, fmt.Errorf("\"%s\": %s", key, err.Error())
			}
		case string:
			if valueType != Value && valueType != Any {
				return nil, fmt.Errorf("Value of \"%s\" is must be %s", key, valueType.String())
			}
			if isObjectId {
				objectId, err := primitive.ObjectIDFromHex(value)
				if err != nil {
					return nil, fmt.Errorf("\"%s\": %s", key, err.Error())
				}
				result[key] = objectId
			} else {
				result[key] = value
			}
		default:
			if valueType != Value && valueType != Any {
				return nil, fmt.Errorf("Value of \"%s\" is must be %s", key, valueType.String())
			}
			if isObjectId {
				return nil, fmt.Errorf("\"%s\" must be objectid string", key)
			} else {
				result[key] = value
			}
		}

	}
	return result, nil
}

func validateArray(v []interface{}, rv reflect.Value, mustBeObjectId bool, elementType ValueType) ([]interface{}, error) {
	result := make([]interface{}, 0)
	for value := range v {
		switch x := v[value].(type) {
		case map[string]interface{}:
			if elementType != Map && elementType != Any {
				return nil, fmt.Errorf("element of array must be %s", elementType.String())
			}
			t, err := validateMap(x, rv, mustBeObjectId)
			if err != nil {
				return nil, err
			}
			result = append(result, t)
		case []interface{}:
			if elementType != Array && elementType != Any {
				return nil, fmt.Errorf("element of array must be %s", elementType.String())
			}
			t, err := validateArray(x, rv, mustBeObjectId, Any)
			if err != nil {
				return nil, err
			}
			result = append(result, t)
		case string:
			if elementType != Value && elementType != Any {
				return nil, fmt.Errorf("element of array must be %s", elementType.String())
			}
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
			if elementType != Value && elementType != Any {
				return nil, fmt.Errorf("element of array must be %s", elementType.String())
			}
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
		isValid := false
		var fieldType reflect.Type
		isValid, key, fieldType = getField(rv, key)
		if !isValid {
			return nil, fmt.Errorf("Fields/keyword %s is not exist", key)
		}

		// fmt.Println(fieldType)
		value, err := getValue(fieldType, value)
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
		return getValue(fieldType.Elem(), value)
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
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32:
		v, err = strconv.ParseInt(value, 10, 32)
	case reflect.Int64:
		v, err = strconv.ParseInt(value, 10, 64)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32:
		v, err = strconv.ParseUint(value, 10, 32)
	case reflect.Uint64:
		v, err = strconv.ParseUint(value, 10, 64)
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
