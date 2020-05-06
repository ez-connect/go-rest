package core

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"reflect"
)

func TypeOf(i interface{}) reflect.Type {
	return reflect.TypeOf(i)
}

func NewInstance(t reflect.Type) interface{} {
	v := reflect.New(t).Elem()
	return v.Interface()
}

// From struct to map[string]interface{}
// using json.Marshal() and json json.Unmarshal()
func Struct2Map(v interface{}) (map[string]interface{}, error) {
	var res map[string]interface{}
	data, err := json.Marshal(&v)
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(data, &res)
	return res, err
}

// From map to struct
// using json.Marshal() and json json.Unmarshal()
func Map2Struct(v map[string]interface{}, i interface{}) error {
	data, err := json.Marshal(&v)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, &i)
}

// Parse a buffer to an interface
func Buffer2Struct(buf *bytes.Buffer, i interface{}) error {
	body, err := ioutil.ReadAll(buf)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, i)
}
