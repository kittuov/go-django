package godjango

import (
	"reflect"
	"fmt"
)

var _regModels map[string]reflect.Type

func RegisterModel(model interface{}){
	fmt.Println("%v",reflect.TypeOf(model))
}
