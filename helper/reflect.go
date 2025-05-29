package helper

import (
	"fmt"
	"math/bits"
	"reflect"
	"strconv"
	"time"
)

// kindToBits将数值类型（如int、float）映射到相应的位数大小。
var kindToBits = map[reflect.Kind]int{
	reflect.Int:     strconv.IntSize,
	reflect.Int8:    8,
	reflect.Int16:   16,
	reflect.Int32:   32,
	reflect.Int64:   64,
	reflect.Uint:    bits.UintSize,
	reflect.Uint16:  16,
	reflect.Uint32:  32,
	reflect.Uint64:  64,
	reflect.Float32: 32,
	reflect.Float64: 64,
}

// setReflectValueFromBool将解析后的布尔值设置给指定的变量。
func setReflectValueFromBool(value reflect.Value, stringValue string) error {
	actualValue, err := strconv.ParseBool(stringValue)
	if err == nil {
		value.SetBool(actualValue)
	}

	return err
}

// setReflectValueFromInt 将解析后的int值设置给指定的变量。
func setReflectValueFromInt(value reflect.Value, stringValue string, bitSize int) error {
	actualValue, err := strconv.ParseInt(stringValue, 10, bitSize)
	if err == nil {
		value.SetInt(actualValue)
	}

	return err
}

// setReflectValueFromUint 将解析后的Uint值设置给指定的变量。
func setReflectValueFromUint(value reflect.Value, stringValue string, bitSize int) error {
	actualValue, err := strconv.ParseUint(stringValue, 10, bitSize)
	if err == nil {
		value.SetUint(actualValue)
	}

	return err
}

// setReflectValueFromFloat 将解析后的float值设置给指定的变量。
func setReflectValueFromFloat(value reflect.Value, stringValue string, bitSize int) error {
	actualValue, err := strconv.ParseFloat(stringValue, bitSize)
	if err == nil {
		value.SetFloat(actualValue)
	}

	return err
}

// setReflectValueFromTime 将解析后的时间日期值设置给指定的变量。
func setReflectValueFromTime(value reflect.Value, stringValue, format string) error {
	actualValue, err := time.Parse(format, stringValue)
	if err == nil {
		value.Set(reflect.ValueOf(actualValue))
	}

	return err
}

// setReflectValue 将解析后的值设置给指定的变量。
func setReflectValue(value reflect.Value, stringValue, format string) error {
	// 获取类型
	kind := value.Kind()

	switch kind {
	case reflect.String:
		value.SetString(stringValue)
		return nil

	case reflect.Bool:
		return setReflectValueFromBool(value, stringValue)

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return setReflectValueFromInt(value, stringValue, kindToBits[kind])

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return setReflectValueFromUint(value, stringValue, kindToBits[kind])

	case reflect.Float32, reflect.Float64:
		return setReflectValueFromFloat(value, stringValue, kindToBits[kind])

	case reflect.Struct:
		typeString := value.Type().String()

		switch typeString {
		case "time.Time":
			return setReflectValueFromTime(value, stringValue, format)

		default:
			return fmt.Errorf("unsupported struct type %s", typeString)
		}

	default:
		return fmt.Errorf("unsupported value kind %s", kind)
	}
}

// getReflectValue 返回给定值的字符串表示形式。
func getReflectValue(value reflect.Value, format string) (string, error) {
	kind := value.Kind()

	switch kind {
	case reflect.String:
		return value.String(), nil

	case reflect.Bool:
		return strconv.FormatBool(value.Bool()), nil

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(value.Int(), 10), nil

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return strconv.FormatUint(value.Uint(), 10), nil

	case reflect.Float32, reflect.Float64:
		return strconv.FormatFloat(value.Float(), 'g', -1, kindToBits[kind]), nil

	case reflect.Struct:
		typeString := value.Type().String()

		switch typeString {
		case "time.Time":
			return value.Interface().(time.Time).Format(format), nil

		default:
			return "", fmt.Errorf("unsupported struct type %s", typeString)
		}

	default:
		return "", fmt.Errorf("unsupported value kind %s", kind)
	}
}
