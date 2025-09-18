package pointer

import (
	"testing"
)

func TestTo(t *testing.T) {
	// 测试基本类型
	intVal := 42
	intPtr := To(intVal)
	if *intPtr != 42 {
		t.Errorf("To(42) = %d, want 42", *intPtr)
	}

	// 测试字符串
	strVal := "hello"
	strPtr := To(strVal)
	if *strPtr != "hello" {
		t.Errorf("To(\"hello\") = %s, want \"hello\"", *strPtr)
	}

	// 测试布尔值
	boolVal := true
	boolPtr := To(boolVal)
	if *boolPtr != true {
		t.Errorf("To(true) = %t, want true", *boolPtr)
	}

	// 测试结构体
	type Person struct {
		Name string
		Age  int
	}
	person := Person{Name: "Alice", Age: 30}
	personPtr := To(person)
	if personPtr.Name != "Alice" || personPtr.Age != 30 {
		t.Errorf("To(Person) = %+v, want {Name: Alice, Age: 30}", *personPtr)
	}
}

func TestFrom(t *testing.T) {
	// 测试基本类型
	intVal := 42
	intPtr := &intVal
	result := From(intPtr)
	if result != 42 {
		t.Errorf("From(&42) = %d, want 42", result)
	}

	// 测试字符串
	strVal := "world"
	strPtr := &strVal
	resultStr := From(strPtr)
	if resultStr != "world" {
		t.Errorf("From(&\"world\") = %s, want \"world\"", resultStr)
	}

	// 测试布尔值
	boolVal := false
	boolPtr := &boolVal
	resultBool := From(boolPtr)
	if resultBool != false {
		t.Errorf("From(&false) = %t, want false", resultBool)
	}

	// 测试结构体
	type Point struct {
		X, Y int
	}
	point := Point{X: 10, Y: 20}
	pointPtr := &point
	resultPoint := From(pointPtr)
	if resultPoint.X != 10 || resultPoint.Y != 20 {
		t.Errorf("From(&Point) = %+v, want {X: 10, Y: 20}", resultPoint)
	}
}

func TestToSlice(t *testing.T) {
	// 测试整数切片
	intSlice := []int{1, 2, 3, 4, 5}
	intPtrSlice := ToSlice(intSlice)

	if len(intPtrSlice) != len(intSlice) {
		t.Errorf("ToSlice length = %d, want %d", len(intPtrSlice), len(intSlice))
	}

	for i, ptr := range intPtrSlice {
		if *ptr != intSlice[i] {
			t.Errorf("ToSlice[%d] = %d, want %d", i, *ptr, intSlice[i])
		}
	}

	// 测试字符串切片
	strSlice := []string{"a", "b", "c"}
	strPtrSlice := ToSlice(strSlice)

	if len(strPtrSlice) != len(strSlice) {
		t.Errorf("ToSlice length = %d, want %d", len(strPtrSlice), len(strSlice))
	}

	for i, ptr := range strPtrSlice {
		if *ptr != strSlice[i] {
			t.Errorf("ToSlice[%d] = %s, want %s", i, *ptr, strSlice[i])
		}
	}

	// 测试空切片
	emptySlice := []int{}
	emptyPtrSlice := ToSlice(emptySlice)
	if len(emptyPtrSlice) != 0 {
		t.Errorf("ToSlice(empty) length = %d, want 0", len(emptyPtrSlice))
	}

	// 测试结构体切片
	type Item struct {
		ID   int
		Name string
	}
	items := []Item{{ID: 1, Name: "item1"}, {ID: 2, Name: "item2"}}
	itemPtrSlice := ToSlice(items)

	if len(itemPtrSlice) != len(items) {
		t.Errorf("ToSlice length = %d, want %d", len(itemPtrSlice), len(items))
	}

	for i, ptr := range itemPtrSlice {
		if ptr.ID != items[i].ID || ptr.Name != items[i].Name {
			t.Errorf("ToSlice[%d] = %+v, want %+v", i, *ptr, items[i])
		}
	}
}

func TestFromSlice(t *testing.T) {
	// 测试整数指针切片
	intPtrSlice := []*int{To(1), To(2), To(3)}
	intSlice := FromSlice(intPtrSlice)

	expected := []int{1, 2, 3}
	if len(intSlice) != len(expected) {
		t.Errorf("FromSlice length = %d, want %d", len(intSlice), len(expected))
	}

	for i, val := range intSlice {
		if val != expected[i] {
			t.Errorf("FromSlice[%d] = %d, want %d", i, val, expected[i])
		}
	}

	// 测试字符串指针切片
	strPtrSlice := []*string{To("x"), To("y"), To("z")}
	strSlice := FromSlice(strPtrSlice)

	expectedStr := []string{"x", "y", "z"}
	if len(strSlice) != len(expectedStr) {
		t.Errorf("FromSlice length = %d, want %d", len(strSlice), len(expectedStr))
	}

	for i, val := range strSlice {
		if val != expectedStr[i] {
			t.Errorf("FromSlice[%d] = %s, want %s", i, val, expectedStr[i])
		}
	}

	// 测试空指针切片
	emptyPtrSlice := []*int{}
	emptySlice := FromSlice(emptyPtrSlice)
	if len(emptySlice) != 0 {
		t.Errorf("FromSlice(empty) length = %d, want 0", len(emptySlice))
	}

	// 测试结构体指针切片
	type Product struct {
		ID    int
		Price float64
	}
	products := []Product{{ID: 1, Price: 10.5}, {ID: 2, Price: 20.0}}
	productPtrSlice := ToSlice(products)
	productSlice := FromSlice(productPtrSlice)

	if len(productSlice) != len(products) {
		t.Errorf("FromSlice length = %d, want %d", len(productSlice), len(products))
	}

	for i, product := range productSlice {
		if product.ID != products[i].ID || product.Price != products[i].Price {
			t.Errorf("FromSlice[%d] = %+v, want %+v", i, product, products[i])
		}
	}
}

func TestToFromRoundTrip(t *testing.T) {
	// 测试 To 和 From 的往返转换
	original := 12345
	ptr := To(original)
	restored := From(ptr)

	if restored != original {
		t.Errorf("To/From round trip: got %d, want %d", restored, original)
	}
}

func TestToSliceFromSliceRoundTrip(t *testing.T) {
	// 测试 ToSlice 和 FromSlice 的往返转换
	original := []int{10, 20, 30, 40, 50}
	ptrSlice := ToSlice(original)
	restored := FromSlice(ptrSlice)

	if len(restored) != len(original) {
		t.Errorf("ToSlice/FromSlice round trip length: got %d, want %d", len(restored), len(original))
	}

	for i, val := range restored {
		if val != original[i] {
			t.Errorf("ToSlice/FromSlice round trip[%d]: got %d, want %d", i, val, original[i])
		}
	}
}

func TestNilPointer(t *testing.T) {
	// 测试 nil 指针的行为
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic when calling From with nil pointer")
		}
	}()

	var nilPtr *int
	From(nilPtr) // 这应该会 panic
}

func TestNilSlice(t *testing.T) {
	// 测试 nil 切片
	var nilSlice []int
	ptrSlice := ToSlice(nilSlice)
	if ptrSlice != nil {
		t.Errorf("ToSlice(nil) = %v, want nil", ptrSlice)
	}

	var nilPtrSlice []*int
	slice := FromSlice(nilPtrSlice)
	if slice != nil {
		t.Errorf("FromSlice(nil) = %v, want nil", slice)
	}
}
