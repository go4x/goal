package ptr

import (
	"testing"
)

func TestTo(t *testing.T) {
	// Test basic types
	intVal := 42
	intPtr := To(intVal)
	if *intPtr != 42 {
		t.Errorf("To(42) = %d, want 42", *intPtr)
	}

	// Test string
	strVal := "hello"
	strPtr := To(strVal)
	if *strPtr != "hello" {
		t.Errorf("To(\"hello\") = %s, want \"hello\"", *strPtr)
	}

	// Test boolean
	boolVal := true
	boolPtr := To(boolVal)
	if *boolPtr != true {
		t.Errorf("To(true) = %t, want true", *boolPtr)
	}

	// Test struct
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
	// Test basic types
	intVal := 42
	intPtr := &intVal
	result := From(intPtr)
	if result != 42 {
		t.Errorf("From(&42) = %d, want 42", result)
	}

	// Test string
	strVal := "world"
	strPtr := &strVal
	resultStr := From(strPtr)
	if resultStr != "world" {
		t.Errorf("From(&\"world\") = %s, want \"world\"", resultStr)
	}

	// Test boolean
	boolVal := false
	boolPtr := &boolVal
	resultBool := From(boolPtr)
	if resultBool != false {
		t.Errorf("From(&false) = %t, want false", resultBool)
	}

	// Test struct
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
	// Test integer slice
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

	// Test string切片
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

	// Test empty slice
	emptySlice := []int{}
	emptyPtrSlice := ToSlice(emptySlice)
	if len(emptyPtrSlice) != 0 {
		t.Errorf("ToSlice(empty) length = %d, want 0", len(emptyPtrSlice))
	}

	// Test struct切片
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
	// Test integer pointer slice
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

	// Test string指针切片
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

	// Test empty pointer slice
	emptyPtrSlice := []*int{}
	emptySlice := FromSlice(emptyPtrSlice)
	if len(emptySlice) != 0 {
		t.Errorf("FromSlice(empty) length = %d, want 0", len(emptySlice))
	}

	// Test struct指针切片
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
	// Test To and From round trip
	original := 12345
	ptr := To(original)
	restored := From(ptr)

	if restored != original {
		t.Errorf("To/From round trip: got %d, want %d", restored, original)
	}
}

func TestToSliceFromSliceRoundTrip(t *testing.T) {
	// Test ToSlice and FromSlice round trip
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

func TestNilptr(t *testing.T) {
	// Test nil pointer behavior
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic when calling From with nil ptr")
		}
	}()

	var nilPtr *int
	From(nilPtr) // 这应该会 panic
}

func TestNilSlice(t *testing.T) {
	// Test nil slice
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

func TestIsNil(t *testing.T) {
	var nilPtr *int
	notNilPtr := To(42)

	if !IsNil(nilPtr) {
		t.Error("IsNil(nil) should return true")
	}

	if IsNil(notNilPtr) {
		t.Error("IsNil(not nil) should return false")
	}
}

func TestIsNotNil(t *testing.T) {
	var nilPtr *int
	notNilPtr := To(42)

	if IsNotNil(nilPtr) {
		t.Error("IsNotNil(nil) should return false")
	}

	if !IsNotNil(notNilPtr) {
		t.Error("IsNotNil(not nil) should return true")
	}
}

func TestValueOr(t *testing.T) {
	var nilPtr *int
	notNilPtr := To(42)

	// Test nil pointer
	result := ValueOr(nilPtr, 100)
	if result != 100 {
		t.Errorf("ValueOr(nil, 100) = %d, want 100", result)
	}

	// Test non-nil pointer
	result = ValueOr(notNilPtr, 100)
	if result != 42 {
		t.Errorf("ValueOr(not nil, 100) = %d, want 42", result)
	}
}

func TestValueOrDefault(t *testing.T) {
	var nilPtr *int
	notNilPtr := To(42)

	// Test nil pointer
	result := ValueOrDefault(nilPtr)
	if result != 0 {
		t.Errorf("ValueOrDefault(nil) = %d, want 0", result)
	}

	// Test non-nil pointer
	result = ValueOrDefault(notNilPtr)
	if result != 42 {
		t.Errorf("ValueOrDefault(not nil) = %d, want 42", result)
	}
}

func TestDeref(t *testing.T) {
	var nilPtr *int
	notNilPtr := To(42)

	// Test nil pointer
	result := Deref(nilPtr)
	if result != 0 {
		t.Errorf("Deref(nil) = %d, want 0", result)
	}

	// Test non-nil pointer
	result = Deref(notNilPtr)
	if result != 42 {
		t.Errorf("Deref(not nil) = %d, want 42", result)
	}
}

func TestDerefOr(t *testing.T) {
	var nilPtr *int
	notNilPtr := To(42)

	// Test nil pointer
	result := DerefOr(nilPtr, 100)
	if result != 100 {
		t.Errorf("DerefOr(nil, 100) = %d, want 100", result)
	}

	// Test non-nil pointer
	result = DerefOr(notNilPtr, 100)
	if result != 42 {
		t.Errorf("DerefOr(not nil, 100) = %d, want 42", result)
	}
}

func TestEqual(t *testing.T) {
	var nilPtr1, nilPtr2 *int
	ptr1 := To(42)
	ptr2 := To(42)
	ptr3 := To(100)

	// Test two nil pointers
	if !Equal(nilPtr1, nilPtr2) {
		t.Error("Equal(nil, nil) should return true")
	}

	// Test one nil one non-nil
	if Equal(nilPtr1, ptr1) {
		t.Error("Equal(nil, not nil) should return false")
	}

	// Test two equal non-nil pointers
	if !Equal(ptr1, ptr2) {
		t.Error("Equal(same values) should return true")
	}

	// Test two different non-nil pointers
	if Equal(ptr1, ptr3) {
		t.Error("Equal(different values) should return false")
	}
}

func TestClone(t *testing.T) {
	original := To(42)
	cloned := Clone(original)

	if cloned == original {
		t.Error("Clone should return a different ptr")
	}

	if *cloned != *original {
		t.Errorf("Clone value = %d, want %d", *cloned, *original)
	}

	// Test nil pointer
	var nilPtr *int
	clonedNil := Clone(nilPtr)
	if clonedNil != nil {
		t.Error("Clone(nil) should return nil")
	}
}

func TestFilter(t *testing.T) {
	slice := []*int{To(1), nil, To(3), nil, To(5)}
	filtered := Filter(slice)

	expected := []*int{To(1), To(3), To(5)}
	if len(filtered) != len(expected) {
		t.Errorf("Filter length = %d, want %d", len(filtered), len(expected))
	}

	for i, ptr := range filtered {
		if *ptr != *expected[i] {
			t.Errorf("Filter[%d] = %d, want %d", i, *ptr, *expected[i])
		}
	}

	// Test nil slice
	var nilSlice []*int
	filteredNil := Filter(nilSlice)
	if filteredNil != nil {
		t.Error("Filter(nil) should return nil")
	}
}

func TestFilterValues(t *testing.T) {
	slice := []*int{To(1), nil, To(3), nil, To(5)}
	filtered := FilterValues(slice)

	expected := []int{1, 3, 5}
	if len(filtered) != len(expected) {
		t.Errorf("FilterValues length = %d, want %d", len(filtered), len(expected))
	}

	for i, val := range filtered {
		if val != expected[i] {
			t.Errorf("FilterValues[%d] = %d, want %d", i, val, expected[i])
		}
	}
}

func TestMap(t *testing.T) {
	slice := []*int{To(1), To(2), To(3)}
	mapped := Map(slice, func(ptr *int) *string {
		if ptr == nil {
			return nil
		}
		return To(string(rune(*ptr + 64)))
	})

	expected := []*string{To("A"), To("B"), To("C")}
	if len(mapped) != len(expected) {
		t.Errorf("Map length = %d, want %d", len(mapped), len(expected))
	}

	for i, ptr := range mapped {
		if *ptr != *expected[i] {
			t.Errorf("Map[%d] = %s, want %s", i, *ptr, *expected[i])
		}
	}
}

func TestMapValues(t *testing.T) {
	slice := []*int{To(1), To(2), To(3)}
	mapped := MapValues(slice, func(val int) string {
		return string(rune(val + 64))
	})

	expected := []string{"A", "B", "C"}
	if len(mapped) != len(expected) {
		t.Errorf("MapValues length = %d, want %d", len(mapped), len(expected))
	}

	for i, val := range mapped {
		if val != expected[i] {
			t.Errorf("MapValues[%d] = %s, want %s", i, val, expected[i])
		}
	}
}

func TestAny(t *testing.T) {
	// Test slice with non-nil pointers
	slice1 := []*int{To(1), nil, To(3)}
	if !Any(slice1) {
		t.Error("Any should return true for slice with non-nil ptrs")
	}

	// Test slice with all nil pointers
	slice2 := []*int{nil, nil, nil}
	if Any(slice2) {
		t.Error("Any should return false for slice with all nil ptrs")
	}

	// Test empty slice
	var emptySlice []*int
	if Any(emptySlice) {
		t.Error("Any should return false for empty slice")
	}

	// Test nil slice
	var nilSlice []*int
	if Any(nilSlice) {
		t.Error("Any should return false for nil slice")
	}
}

func TestAll(t *testing.T) {
	// Test slice with all non-nil pointers
	slice1 := []*int{To(1), To(2), To(3)}
	if !All(slice1) {
		t.Error("All should return true for slice with all non-nil ptrs")
	}

	// Test slice with nil pointers
	slice2 := []*int{To(1), nil, To(3)}
	if All(slice2) {
		t.Error("All should return false for slice with nil ptrs")
	}

	// Test empty slice
	var emptySlice []*int
	if !All(emptySlice) {
		t.Error("All should return true for empty slice")
	}

	// Test nil slice
	var nilSlice []*int
	if !All(nilSlice) {
		t.Error("All should return true for nil slice")
	}
}

func TestCount(t *testing.T) {
	slice := []*int{To(1), nil, To(3), nil, To(5)}
	count := Count(slice)
	if count != 3 {
		t.Errorf("Count = %d, want 3", count)
	}

	// Test empty slice
	var emptySlice []*int
	count = Count(emptySlice)
	if count != 0 {
		t.Errorf("Count(empty) = %d, want 0", count)
	}
}

func TestFirst(t *testing.T) {
	slice := []*int{nil, To(2), To(3)}
	first := First(slice)
	if first == nil || *first != 2 {
		t.Errorf("First = %v, want ptr to 2", first)
	}

	// Test all nil slice
	allNilSlice := []*int{nil, nil, nil}
	first = First(allNilSlice)
	if first != nil {
		t.Error("First(all nil) should return nil")
	}
}

func TestLast(t *testing.T) {
	slice := []*int{To(1), To(2), nil, To(4)}
	last := Last(slice)
	if last == nil || *last != 4 {
		t.Errorf("Last = %v, want ptr to 4", last)
	}

	// Test all nil slice
	allNilSlice := []*int{nil, nil, nil}
	last = Last(allNilSlice)
	if last != nil {
		t.Error("Last(all nil) should return nil")
	}
}

func TestSet(t *testing.T) {
	value := 42
	ptr := &value
	Set(ptr, 100)
	if *ptr != 100 {
		t.Errorf("Set failed: got %d, want 100", *ptr)
	}

	// Test nil pointer
	var nilPtr *int
	Set(nilPtr, 100) // 应该不会 panic
}

func TestZero(t *testing.T) {
	value := 42
	ptr := &value
	Zero(ptr)
	if *ptr != 0 {
		t.Errorf("Zero failed: got %d, want 0", *ptr)
	}

	// Test nil pointer
	var nilPtr *int
	Zero(nilPtr) // 应该不会 panic
}

func TestSwap(t *testing.T) {
	a := To(1)
	b := To(2)
	Swap(a, b)
	if *a != 2 || *b != 1 {
		t.Errorf("Swap failed: got a=%d, b=%d, want a=2, b=1", *a, *b)
	}

	// Test nil pointer
	var nilPtr *int
	Swap(a, nilPtr) // 应该不会 panic
	Swap(nilPtr, b) // 应该不会 panic
}

func TestDeepEqual(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}

	p1 := &Person{Name: "Alice", Age: 30}
	p2 := &Person{Name: "Alice", Age: 30}
	p3 := &Person{Name: "Bob", Age: 25}

	// Test equal structs
	if !DeepEqual(p1, p2) {
		t.Error("DeepEqual should return true for equal structs")
	}

	// Test different structs
	if DeepEqual(p1, p3) {
		t.Error("DeepEqual should return false for different structs")
	}

	// Test nil pointer
	var nilPtr1, nilPtr2 *Person
	if !DeepEqual(nilPtr1, nilPtr2) {
		t.Error("DeepEqual should return true for both nil")
	}

	if DeepEqual(p1, nilPtr1) {
		t.Error("DeepEqual should return false for one nil")
	}
}
