package Utils_test

import (
	"github.com/2840422688/yang-lib/utils"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestCheckParamsEmpty(t *testing.T) {
	// 正常情况：参数列表为空
	assert.False(t, Utils.CheckParamsEmpty())

	// 正常情况：参数列表不包含空值
	assert.False(t, Utils.CheckParamsEmpty("foo", 42))

	// 异常情况：参数列表包含空值
	assert.True(t, Utils.CheckParamsEmpty("", "bar", 0))
}

func TestAutoConvert(t *testing.T) {
	// 正常情况：转换为字符串
	//strs, err := Utils.AutoConvert(reflect.String, "foo", 42)
	//assert.NoError(t, err)
	//assert.Equal(t, []interface{}{"foo", "42"}, strs)

	// 正常情况：转换为整数
	ints, err := Utils.AutoConvert(reflect.String, 123, 456)
	assert.NoError(t, err)
	assert.Equal(t, []interface{}{"123", "456"}, ints)

	// 正常情况：转换为浮点数
	floats, err := Utils.AutoConvert(reflect.Float64, "3.14", 2)
	assert.NoError(t, err)
	assert.Equal(t, []interface{}{3.14, 2.0}, floats)

	// 异常情况：无法识别的类型
	_, err = Utils.AutoConvert(reflect.Ptr, "pointer", "invalid")
	assert.Error(t, err)

	// 异常情况：无法转换类型
	_, err = Utils.AutoConvert(reflect.Int, "abc", "def")
	assert.Error(t, err)
}

func TestGetInterfaceType(t *testing.T) {
	// 正常情况：获取参数类型列表
	types := Utils.GetInterfaceType("foo", 42, 3.14)
	assert.Equal(t, []string{"string", "int", "float64"}, types)

	// 正常情况：空参数列表
	types = Utils.GetInterfaceType()
	assert.Empty(t, types)
}
