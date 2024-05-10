package Utils

//类型值处理工具类
import (
	"errors"
	"reflect"
	"strconv"
)

func CheckParamsEmpty(params ...interface{}) bool {
	for i := 0; i < len(params); i++ {
		t := reflect.TypeOf(params[i])
		v := reflect.ValueOf(params[i])
		switch t.Kind() {
		case reflect.String:
			if v.String() == "" {
				return true
			}
		case reflect.Float32, reflect.Float64, reflect.Int, reflect.Int32, reflect.Int64, reflect.Int16:
			if v.IsZero() {
				return true
			}
			break
		case reflect.Ptr, reflect.Chan, reflect.Map:
			if v.IsNil() {
				return true
			}
			break
		}
	}
	return false
}

func GetInterfaceType(params ...interface{}) []string {
	var typeSlice = make([]string, len(params))
	for i := 0; i < len(params); i++ {
		t := reflect.TypeOf(params[i])
		typeSlice[i] = t.Kind().String()
	}
	return typeSlice
}

func GetValueOfBaseType(params ...interface{}) ([]interface{}, error) {
	var typeSlice = make([]interface{}, len(params))
	if !CheckParamsEmpty(params) {
		return nil, errors.New("存在空（零）值，请检查")
	}
	for i := 0; i < len(params); i++ {
		v := reflect.ValueOf(params[i])
		t := GetInterfaceType(params)
		switch t[i] {
		case reflect.String.String():
			typeSlice = append(typeSlice, v.String())
			break
		case reflect.Int.String(), reflect.Int32.String(), reflect.Int64.String(), reflect.Int16.String():
			typeSlice = append(typeSlice, v.Int())
			break
		case reflect.Float32.String(), reflect.Float64.String():
			typeSlice = append(typeSlice, v.Float())
			break
		}

	}
	return typeSlice, errors.New("无法识别的类型或非基础数据类型（string,int(64,32...),float(...)）")
}

func AutoConvert(convertType reflect.Kind, params ...interface{}) ([]interface{}, error) {
	var err error
	for i := 0; i < len(params); i++ {
		v := reflect.ValueOf(params[i])
		t := reflect.TypeOf(params[i])
		switch convertType {
		case reflect.String:
			switch t.Kind() {
			case reflect.String:
				params[i] = v.String()
				break
			case reflect.Int:
				params[i] = strconv.Itoa(int(v.Int()))
				break
			case reflect.Float32, reflect.Float64:
				params[i] = strconv.FormatFloat(v.Float(), 'f', 6, 64)
				break
			default:
				return nil, errors.New("值(" + v.String() + ")无法转为" + convertType.String() + "类型")
			}
		case reflect.Int, reflect.Int32, reflect.Int64, reflect.Int16:
			switch t.Kind() {
			case reflect.Int, reflect.Int32, reflect.Int64, reflect.Int16:
				params[i] = v.Int()
			case reflect.String:
				if params[i], err = strconv.Atoi(v.String()); err != nil {
					return nil, err
				}
			case reflect.Float64, reflect.Float32:
				params[i] = int(v.Float())
			default:
				return nil, errors.New("值无法转为" + convertType.String() + "类型")
			}
		case reflect.Float32:
			switch t.Kind() {
			case reflect.Float32, reflect.Float64:
				params[i] = v.Float()
			case reflect.String:
				if params[i], err = strconv.ParseFloat(v.String(), 64); err != nil {
					return nil, err
				}
			case reflect.Int, reflect.Int32, reflect.Int64, reflect.Int16:
				params[i] = float32(v.Float())
			default:
				return nil, errors.New("值无法转为" + convertType.String() + "类型")
			}
			break
		case reflect.Float64:
			switch t.Kind() {
			case reflect.Float32, reflect.Float64:
				params[i] = v.Float()
			case reflect.String:
				if params[i], err = strconv.ParseFloat(v.String(), 64); err != nil {
					return nil, err
				}
			case reflect.Int, reflect.Int32, reflect.Int64, reflect.Int16:
				params[i] = float64(v.Int())
			default:
				return nil, errors.New("值无法转为" + convertType.String() + "类型")
			}
			break
		default:
			return nil, errors.New("没有可以匹配转换的类型")
		}
	}
	return params, nil
}
