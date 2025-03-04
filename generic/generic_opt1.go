package main

import (
	"fmt"
	"reflect"
)

type UserAttributeType string

func convertSlice[T any, U any](input []T, convert func(T) U) []U {
	result := make([]U, len(input))
	for i, v := range input {
		result[i] = convert(v)
	}
	return result
}

func main() {
	strings := []string{"attribute1", "attribute2", "attribute3"}
	userAttributes := convertSlice(strings, func(s string) UserAttributeType {
		return UserAttributeType(s)
	})
	fmt.Println(reflect.TypeOf(userAttributes), userAttributes) // 输出: [attribute1 attribute2 attribute3]
}
