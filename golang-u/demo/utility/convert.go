// Package utilities provides ...
package utility

import (
	"errors"
	"reflect"
	"strconv"
)

// because Config contains private structs we can't use reflect.Value
// directly, instead we need to "unsafely" address the variable
// func unsafeValueOf(val reflect.Value) reflect.Value {
// 	uptr := unsafe.Pointer(val.UnsafeAddr())
// 	return reflect.NewAt(val.Type(), uptr).Elem()
// }
//
// func valueCompare(v1 reflect.Value, v2 reflect.Value) int {
// 	switch v1.Type().String() {
// 	case "int", "int16", "int32", "int64":
// 		if v1.Int() > v2.Int() {
// 			return 1
// 		} else if v1.Int() < v2.Int() {
// 			return -1
// 		}
// 		return 0
// 	case "uint", "uint16", "uint32", "uint64":
// 		if v1.Uint() > v2.Uint() {
// 			return 1
// 		} else if v1.Uint() < v2.Uint() {
// 			return -1
// 		}
// 		return 0
// 	case "float32", "float64":
// 		if v1.Float() > v2.Float() {
// 			return 1
// 		} else if v1.Float() < v2.Float() {
// 			return -1
// 		}
// 		return 0
// 	case "time.Duration":
// 		if v1.Interface().(time.Duration) > v2.Interface().(time.Duration) {
// 			return 1
// 		} else if v1.Interface().(time.Duration) < v2.Interface().(time.Duration) {
// 			return -1
// 		}
// 		return 0
// 	}
// 	panic("impossible")
// }
//
// func coerce(v interface{}, typ reflect.Type) (reflect.Value, error) {
// 	var err error
// 	if typ.Kind() == reflect.Ptr {
// 		return reflect.ValueOf(v), nil
// 	}
// 	switch typ.String() {
// 	case "string":
// 		v, err = coerceString(v)
// 	case "int", "int16", "int32", "int64":
// 		v, err = coerceInt64(v)
// 	case "uint", "uint16", "uint32", "uint64":
// 		v, err = coerceUint64(v)
// 	case "float32", "float64":
// 		v, err = coerceFloat64(v)
// 	case "bool":
// 		v, err = coerceBool(v)
// 	case "time.Duration":
// 		v, err = coerceDuration(v)
// 	case "net.Addr":
// 		v, err = coerceAddr(v)
// 	case "nsq.BackoffStrategy":
// 		v, err = coerceBackoffStrategy(v)
// 	default:
// 		v = nil
// 		err = fmt.Errorf("invalid type %s", typ.String())
// 	}
// 	return valueTypeCoerce(v, typ), err
// }
//
// func valueTypeCoerce(v interface{}, typ reflect.Type) reflect.Value {
// 	val := reflect.ValueOf(v)
// 	if reflect.TypeOf(v) == typ {
// 		return val
// 	}
// 	tval := reflect.New(typ).Elem()
// 	switch typ.String() {
// 	case "int", "int16", "int32", "int64":
// 		tval.SetInt(val.Int())
// 	case "uint", "uint16", "uint32", "uint64":
// 		tval.SetUint(val.Uint())
// 	case "float32", "float64":
// 		tval.SetFloat(val.Float())
// 	default:
// 		tval.Set(val)
// 	}
// 	return tval
// }
//
// func coerceString(v interface{}) (string, error) {
// 	switch v := v.(type) {
// 	case string:
// 		return v, nil
// 	case int, int16, int32, int64, uint, uint16, uint32, uint64:
// 		return fmt.Sprintf("%d", v), nil
// 	case float32, float64:
// 		return fmt.Sprintf("%f", v), nil
// 	}
// 	return fmt.Sprintf("%s", v), nil
// }
//
// func coerceDuration(v interface{}) (time.Duration, error) {
// 	switch v := v.(type) {
// 	case string:
// 		return time.ParseDuration(v)
// 	case int, int16, int32, int64:
// 		// treat like ms
// 		return time.Duration(reflect.ValueOf(v).Int()) * time.Millisecond, nil
// 	case uint, uint16, uint32, uint64:
// 		// treat like ms
// 		return time.Duration(reflect.ValueOf(v).Uint()) * time.Millisecond, nil
// 	case time.Duration:
// 		return v, nil
// 	}
// 	return 0, errors.New("invalid value type")
// }
//
// func coerceAddr(v interface{}) (net.Addr, error) {
// 	switch v := v.(type) {
// 	case string:
// 		return net.ResolveTCPAddr("tcp", v)
// 	case net.Addr:
// 		return v, nil
// 	}
// 	return nil, errors.New("invalid value type")
// }
//
// func coerceBackoffStrategy(v interface{}) (BackoffStrategy, error) {
// 	switch v := v.(type) {
// 	case string:
// 		switch v {
// 		case "", "exponential":
// 			return &ExponentialStrategy{}, nil
// 		case "full_jitter":
// 			return &FullJitterStrategy{}, nil
// 		}
// 	case BackoffStrategy:
// 		return v, nil
// 	}
// 	return nil, errors.New("invalid value type")
// }
//
// func coerceBool(v interface{}) (bool, error) {
// 	switch v := v.(type) {
// 	case bool:
// 		return v, nil
// 	case string:
// 		return strconv.ParseBool(v)
// 	case int, int16, int32, int64:
// 		return reflect.ValueOf(v).Int() != 0, nil
// 	case uint, uint16, uint32, uint64:
// 		return reflect.ValueOf(v).Uint() != 0, nil
// 	}
// 	return false, errors.New("invalid value type")
// }

func CoerceInt(v interface{}) (int, error) {
	switch v := v.(type) {
	case string:
		i64, err := strconv.ParseInt(v, 10, 0)
		return int(i64), err
	case int, int16, int32, int64:
		return int(reflect.ValueOf(v).Int()), nil
	case uint, uint16, uint32, uint64:
		return int(reflect.ValueOf(v).Uint()), nil
	}
	return 0, errors.New("invalid value type")
}

func CoerceInt64(v interface{}) (int64, error) {
	switch v := v.(type) {
	case string:
		println("string")
		return strconv.ParseInt(v, 10, 64)
	case int, int16, int32, int64:
		println("int*")
		return reflect.ValueOf(v).Int(), nil
	case uint, uint16, uint32, uint64:
		println("uint*")
		return int64(reflect.ValueOf(v).Uint()), nil
	}
	return 0, errors.New("invalid value type")
}

// func coerceUint64(v interface{}) (uint64, error) {
// 	switch v := v.(type) {
// 	case string:
// 		return strconv.ParseUint(v, 10, 64)
// 	case int, int16, int32, int64:
// 		return uint64(reflect.ValueOf(v).Int()), nil
// 	case uint, uint16, uint32, uint64:
// 		return reflect.ValueOf(v).Uint(), nil
// 	}
// 	return 0, errors.New("invalid value type")
// }

func CoerceFloat64(v interface{}) (float64, error) {
	switch v := v.(type) {
	case string:
		println("string")
		return strconv.ParseFloat(v, 64)
	case int, int16, int32, int64:
		println("int*")
		return float64(reflect.ValueOf(v).Int()), nil
	case uint, uint16, uint32, uint64:
		println("uint*")
		return float64(reflect.ValueOf(v).Uint()), nil
	case float32:
		println("float32")
		return float64(v), nil
	case float64:
		println("float64")
		return v, nil
	}
	return 0, errors.New("invalid value type")
}
