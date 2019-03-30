package jute

import (
	"encoding/binary"
	"fmt"
	"math"
	"reflect"
)

func Marshal(v interface{}) ([]byte, error) {
	rv := reflect.Indirect(reflect.ValueOf(v))
	bs := make([]byte, size(rv))
	err := marshal(rv, bs)
	//switch reflect.Indirect(rv).Kind() {
	//case reflect.Struct:
	//	for i := 0; i < rv.NumField(); i++ {
	//
	//	}
	//}
	return bs, err
}

func marshal(rv reflect.Value, bs []byte) error {
	switch rv.Kind() {
	case reflect.Uint8:
		bs[0] = rv.Interface().(uint8)
	case reflect.Uint32:
		binary.BigEndian.PutUint32(bs, rv.Interface().(uint32))
	case reflect.Uint64:
		binary.BigEndian.PutUint64(bs, rv.Interface().(uint64))
	case reflect.Float32:
		binary.BigEndian.PutUint32(bs, math.Float32bits(rv.Interface().(float32)))
	case reflect.Float64:
		binary.BigEndian.PutUint64(bs, math.Float64bits(rv.Interface().(float64)))
	case reflect.Bool:
		if rv.Interface().(bool) {
			bs[0] = 1
		} else {
			bs[0] = 0
		}
	case reflect.String:
		binary.BigEndian.PutUint32(bs, uint32(rv.Len()))
		for i := 0; i < rv.Len(); i++ {
			bs[4+i] = rv.Index(i).Interface().(uint8)
		}
	case reflect.Struct:
		var o int
		for i := 0; i < rv.NumField(); i++ {
			err := marshal(rv.Field(i), bs[o:])
			if err != nil {
				return err
			}
			o += size(rv.Field(i))
		}
	case reflect.Slice:
		length := reflect.ValueOf(uint32(rv.Len()))
		err := marshal(length, bs)
		if err != nil {
			return err
		}
		o := size(length)
		for i := 0; i < rv.Len(); i++ {
			err = marshal(rv.Index(i), bs[o:])
			if err != nil {
				return err
			}
			o += size(rv.Index(i))
		}
	default:
		return fmt.Errorf("unknown: %T(%v)", rv.Interface(), rv.Interface())
	}
	return nil
}

func size(rv reflect.Value) int {
	switch rv.Kind() {
	case reflect.Uint8:
		return 1
	case reflect.Uint32:
		return 4
	case reflect.Uint64:
		return 8
	case reflect.Float32:
		return 4
	case reflect.Float64:
		return 8
	case reflect.Bool:
		return 1
	case reflect.String:
		return 4 + rv.Len()
	case reflect.Struct:
		var n int
		for i := 0; i < rv.NumField(); i++ {
			n += size(rv.Field(i))
		}
		return n
	case reflect.Slice:
		n := size(reflect.Zero(rv.Type().Elem()))
		return rv.Len()*n + 4
	default:
		panic(fmt.Sprintf("%s %v", rv.Type(), rv))
	}
}
