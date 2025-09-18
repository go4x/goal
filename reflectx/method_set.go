package reflectx

import (
	"fmt"
	"reflect"
)

func Methods(i interface{}) []string {
	var t = reflect.TypeOf(i)
	var elem = t.Elem()

	var nm = elem.NumMethod()
	var ret []string
	for i := 0; i < nm; i++ {
		var m = elem.Method(i)
		ret = append(ret, m.Name)
	}
	return ret
}

func PrintMethodSet(i interface{}) {
	var t = reflect.TypeOf(i)
	var elem = t.Elem()

	var nm = elem.NumMethod()
	if nm == 0 {
		fmt.Printf("%s's method set is empty\n", elem)
		return
	}
	fmt.Printf("%s's method set:\n", elem)
	for i := 0; i < nm; i++ {
		var m = elem.Method(i)
		fmt.Printf("  - %s\n", m.Name)
	}
	fmt.Println()
}
