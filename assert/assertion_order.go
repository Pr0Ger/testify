package assert

import (
	"fmt"
	"reflect"
)

const (
	greater = 1
	equal   = 0
	less    = -1
)

func compare(obj1, obj2 interface{}, kind reflect.Kind) (int, bool) {
	switch kind {
	case reflect.Int:
		{
			intobj1 := obj1.(int)
			intobj2 := obj2.(int)
			if intobj1 > intobj2 {
				return greater, true
			}
			if intobj1 == intobj2 {
				return equal, true
			}
			if intobj1 < intobj2 {
				return less, true
			}
		}
	case reflect.Int8:
		{
			int8obj1 := obj1.(int8)
			int8obj2 := obj2.(int8)
			if int8obj1 > int8obj2 {
				return greater, true
			}
			if int8obj1 == int8obj2 {
				return equal, true
			}
			if int8obj1 < int8obj2 {
				return less, true
			}
		}
	case reflect.Int16:
		{
			int16obj1 := obj1.(int16)
			int16obj2 := obj2.(int16)
			if int16obj1 > int16obj2 {
				return greater, true
			}
			if int16obj1 == int16obj2 {
				return equal, true
			}
			if int16obj1 < int16obj2 {
				return less, true
			}
		}
	case reflect.Int32:
		{
			int32obj1 := obj1.(int32)
			int32obj2 := obj2.(int32)
			if int32obj1 > int32obj2 {
				return greater, true
			}
			if int32obj1 == int32obj2 {
				return equal, true
			}
			if int32obj1 < int32obj2 {
				return less, true
			}
		}
	case reflect.Int64:
		{
			int64obj1 := obj1.(int64)
			int64obj2 := obj2.(int64)
			if int64obj1 > int64obj2 {
				return greater, true
			}
			if int64obj1 == int64obj2 {
				return equal, true
			}
			if int64obj1 < int64obj2 {
				return less, true
			}
		}
	case reflect.Uint:
		{
			uintobj1 := obj1.(uint)
			uintobj2 := obj2.(uint)
			if uintobj1 > uintobj2 {
				return greater, true
			}
			if uintobj1 == uintobj2 {
				return equal, true
			}
			if uintobj1 < uintobj2 {
				return less, true
			}
		}
	case reflect.Uint8:
		{
			uint8obj1 := obj1.(uint8)
			uint8obj2 := obj2.(uint8)
			if uint8obj1 > uint8obj2 {
				return greater, true
			}
			if uint8obj1 == uint8obj2 {
				return equal, true
			}
			if uint8obj1 < uint8obj2 {
				return less, true
			}
		}
	case reflect.Uint16:
		{
			uint16obj1 := obj1.(uint16)
			uint16obj2 := obj2.(uint16)
			if uint16obj1 > uint16obj2 {
				return greater, true
			}
			if uint16obj1 == uint16obj2 {
				return equal, true
			}
			if uint16obj1 < uint16obj2 {
				return less, true
			}
		}
	case reflect.Uint32:
		{
			uint32obj1 := obj1.(uint32)
			uint32obj2 := obj2.(uint32)
			if uint32obj1 > uint32obj2 {
				return greater, true
			}
			if uint32obj1 == uint32obj2 {
				return equal, true
			}
			if uint32obj1 < uint32obj2 {
				return less, true
			}
		}
	case reflect.Uint64:
		{
			uint64obj1 := obj1.(uint64)
			uint64obj2 := obj2.(uint64)
			if uint64obj1 > uint64obj2 {
				return greater, true
			}
			if uint64obj1 == uint64obj2 {
				return equal, true
			}
			if uint64obj1 < uint64obj2 {
				return less, true
			}
		}
	case reflect.Float32:
		{
			float32obj1 := obj1.(float32)
			float32obj2 := obj2.(float32)
			if float32obj1 > float32obj2 {
				return greater, true
			}
			if float32obj1 == float32obj2 {
				return equal, true
			}
			if float32obj1 < float32obj2 {
				return less, true
			}
		}
	case reflect.Float64:
		{
			float64obj1 := obj1.(float64)
			float64obj2 := obj2.(float64)
			if float64obj1 > float64obj2 {
				return greater, true
			}
			if float64obj1 == float64obj2 {
				return equal, true
			}
			if float64obj1 < float64obj2 {
				return less, true
			}
		}
	case reflect.String:
		{
			stringobj1 := obj1.(string)
			stringobj2 := obj2.(string)
			if stringobj1 > stringobj2 {
				return greater, true
			}
			if stringobj1 == stringobj2 {
				return equal, true
			}
			if stringobj1 < stringobj2 {
				return less, true
			}
		}
	}

	return equal, false
}

// Greater asserts that the first element is greater than the second
//
//    assert.Greater(t, 2, 1)
//    assert.Greater(t, float64(2), float64(1))
//    assert.Greater(t, "b", "a")
func Greater(t TestingT, e1 interface{}, e2 interface{}, msgAndArgs ...interface{}) bool {
	return compareTwoValues(t, e1, e2, []int{greater}, "\"%s\" is not greater than \"%s\"", msgAndArgs)
}

// GreaterOrEqual asserts that the first element is greater than or equal to the second
//
//    assert.GreaterOrEqual(t, 2, 1)
//    assert.GreaterOrEqual(t, 2, 2)
//    assert.GreaterOrEqual(t, "b", "a")
//    assert.GreaterOrEqual(t, "b", "b")
func GreaterOrEqual(t TestingT, e1 interface{}, e2 interface{}, msgAndArgs ...interface{}) bool {
	return compareTwoValues(t, e1, e2, []int{greater, equal}, "\"%s\" is not greater or equal than \"%s\"", msgAndArgs)
}

// Less asserts that the first element is less than the second
//
//    assert.Less(t, 1, 2)
//    assert.Less(t, float64(1), float64(2))
//    assert.Less(t, "a", "b")
func Less(t TestingT, e1 interface{}, e2 interface{}, msgAndArgs ...interface{}) bool {
	return compareTwoValues(t, e1, e2, []int{less}, "\"%s\" is not less than \"%s\"", msgAndArgs)
}

// LessOrEqual asserts that the first element is less than or equal to the second
//
//    assert.LessOrEqual(t, 1, 2)
//    assert.LessOrEqual(t, 2, 2)
//    assert.LessOrEqual(t, "a", "b")
//    assert.LessOrEqual(t, "b", "b")
func LessOrEqual(t TestingT, e1 interface{}, e2 interface{}, msgAndArgs ...interface{}) bool {
	return compareTwoValues(t, e1, e2, []int{less, equal}, "\"%s\" is not less or equal than \"%s\"", msgAndArgs)
}

func compareTwoValues(t TestingT, e1 interface{}, e2 interface{}, allowedComparesResults []int, failMessage string, msgAndArgs ...interface{}) bool {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}

	e1Kind := reflect.ValueOf(e1).Kind()
	e2Kind := reflect.ValueOf(e2).Kind()
	if e1Kind != e2Kind {
		return Fail(t, "Elements should be the same type", msgAndArgs...)
	}

	compareResult, isComparable := compare(e1, e2, e1Kind)
	if !isComparable {
		return Fail(t, fmt.Sprintf("Can not compare type \"%s\"", reflect.TypeOf(e1)), msgAndArgs...)
	}

	if !containsValue(allowedComparesResults, compareResult) {
		return Fail(t, fmt.Sprintf(failMessage, e1, e2), msgAndArgs...)
	}

	return true
}

func containsValue(values []int, value int) bool {
	for _, v := range values {
		if v == value {
			return true
		}
	}

	return false
}
