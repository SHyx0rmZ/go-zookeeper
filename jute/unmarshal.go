package jute

import (
	"encoding/binary"
	"fmt"
	"reflect"
)

func Unmarshal(bs []byte, v interface{}) error {
	err := unmarshal(bs, reflect.ValueOf(v))
	return err
}

func unmarshal(bs []byte, rv reflect.Value) error {
	ri := reflect.Indirect(rv)
	switch ri.Kind() {
	case reflect.Uint8:
		rv.SetUint(uint64(bs[0]))
	case reflect.Uint32:
		rv.SetUint(uint64(binary.BigEndian.Uint32(bs)))
	case reflect.Uint64:
		rv.SetUint(uint64(binary.BigEndian.Uint64(bs)))
	case reflect.Struct:
		var o int
		for i := 0; i < ri.NumField(); i++ {
			err := unmarshal(bs[o:], ri.Field(i))
			if err != nil {
				return err
			}
			o += size(ri.Field(i))
		}
	case reflect.Slice:
		length := int(binary.BigEndian.Uint32(bs))
		rv.Set(reflect.MakeSlice(ri.Type(), length, length))
		o := 4
		for i := 0; i < length; i++ {
			err := unmarshal(bs[o:], rv.Index(i))
			if err != nil {
				return err
			}
			o += size(rv.Index(i))
		}
	case reflect.String:
		length := int(binary.BigEndian.Uint32(bs))
		rv.SetString(string(bs[4 : 4+length]))
	default:
		return fmt.Errorf("unknown: %#v", rv.Interface())
	}
	return nil
}
