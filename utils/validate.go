package Utils

//utils基础校验规则
import (
	"errors"
	"regexp"
)

func IsChinese(param string) bool {
	result, err := regexp.MatchString(`^[\p{Han}]+$`, param)
	if err != nil {
		return result
	}
	return result
}

func IsEnglish(param string) bool {
	result, err := regexp.MatchString(`^[a-zA-Z]+$`, param)
	if err != nil {
		return result
	}
	return result
}

func IsPhoneNum(param string) bool {
	result, err := regexp.MatchString(`^(1[3|4|5|8][0-9]\d{4,8})$`, param)
	if err != nil {
		return result
	}
	return result
}

func IsNum(param string) bool {
	result, err := regexp.MatchString(`^[0-9]+$`, param)
	if err != nil {
		return result
	}
	return result
}

func IsIdCard(param string) (bool, error) {
	switch len(param) {
	case 18:
		result, err := regexp.MatchString(`(^\\d{18}$)|(^\\d{17}(\\d|X|x)$)`, param)
		if err != nil {
			return result, err
		}
		break
	case 15:
		result, err := regexp.MatchString(`(^\\d{15}$)`, param)
		if err != nil {
			return result, err
		}
		break
	default:
		return false, errors.New("身份证格式不正确")
	}
	return true, nil
}
