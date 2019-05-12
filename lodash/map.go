package lodash

import (
	"reflect"
)

// Map lodash helper
func Map(source, selector interface{}) interface{} {
	var tempRV reflect.Value
	// Each(source, selector, func(resRV, valueRV, _ reflect.Value) bool {
	// 	if !tempRV.IsValid() {
	// 		tempRT := reflect.SliceOf(resRV.Type())
	// 		tempRV = reflect.MakeSlice(tempRT, 0, 0)
	// 	}

	// 	tempRV = reflect.Append(tempRV, resRV)
	// 	return false
	// })
	if tempRV.IsValid() {
		return tempRV.Interface()
	}

	return nil
}

// MapBy is 从source中取出所有property
func MapBy(source interface{}, property string) interface{} {
	// getPropertyRV := PropertyRV(property)
	// return Map(source, func(value, _ interface{}) facade {
	// 	return facade{
	// 		getPropertyRV(value),
	// 	}
	// })

	return nil
}
