package ptr

import (
	"testing"
)

func BenchmarkTo(b *testing.B) {
	value := 42
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = To(value)
	}
}

func BenchmarkFrom(b *testing.B) {
	value := 42
	ptr := &value
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = From(ptr)
	}
}

func BenchmarkIsNil(b *testing.B) {
	var nilPtr *int
	notNilPtr := To(42)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = IsNil(nilPtr)
		_ = IsNil(notNilPtr)
	}
}

func BenchmarkDeref(b *testing.B) {
	var nilPtr *int
	notNilPtr := To(42)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Deref(nilPtr)
		_ = Deref(notNilPtr)
	}
}

func BenchmarkValueOr(b *testing.B) {
	var nilPtr *int
	notNilPtr := To(42)
	defaultValue := 100
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ValueOr(nilPtr, defaultValue)
		_ = ValueOr(notNilPtr, defaultValue)
	}
}

func BenchmarkEqual(b *testing.B) {
	ptr1 := To(42)
	ptr2 := To(42)
	ptr3 := To(100)
	var nilPtr *int
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Equal(ptr1, ptr2)
		_ = Equal(ptr1, ptr3)
		_ = Equal(ptr1, nilPtr)
	}
}

func BenchmarkClone(b *testing.B) {
	original := To(42)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Clone(original)
	}
}

func BenchmarkToSlice(b *testing.B) {
	slice := make([]int, 1000)
	for i := range slice {
		slice[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ToSlice(slice)
	}
}

func BenchmarkFromSlice(b *testing.B) {
	slice := make([]int, 1000)
	for i := range slice {
		slice[i] = i
	}
	ptrSlice := ToSlice(slice)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = FromSlice(ptrSlice)
	}
}

func BenchmarkFilter(b *testing.B) {
	slice := make([]*int, 1000)
	for i := range slice {
		if i%2 == 0 {
			slice[i] = To(i)
		} else {
			slice[i] = nil
		}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Filter(slice)
	}
}

func BenchmarkFilterValues(b *testing.B) {
	slice := make([]*int, 1000)
	for i := range slice {
		if i%2 == 0 {
			slice[i] = To(i)
		} else {
			slice[i] = nil
		}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = FilterValues(slice)
	}
}

func BenchmarkMapValues(b *testing.B) {
	slice := make([]*int, 1000)
	for i := range slice {
		slice[i] = To(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = MapValues(slice, func(val int) int {
			return val * 2
		})
	}
}

func BenchmarkAny(b *testing.B) {
	slice := make([]*int, 1000)
	for i := range slice {
		if i%2 == 0 {
			slice[i] = To(i)
		} else {
			slice[i] = nil
		}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Any(slice)
	}
}

func BenchmarkAll(b *testing.B) {
	slice := make([]*int, 1000)
	for i := range slice {
		slice[i] = To(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = All(slice)
	}
}

func BenchmarkCount(b *testing.B) {
	slice := make([]*int, 1000)
	for i := range slice {
		if i%2 == 0 {
			slice[i] = To(i)
		} else {
			slice[i] = nil
		}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Count(slice)
	}
}

func BenchmarkFirst(b *testing.B) {
	slice := make([]*int, 1000)
	for i := range slice {
		if i%2 == 0 {
			slice[i] = To(i)
		} else {
			slice[i] = nil
		}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = First(slice)
	}
}

func BenchmarkLast(b *testing.B) {
	slice := make([]*int, 1000)
	for i := range slice {
		if i%2 == 0 {
			slice[i] = To(i)
		} else {
			slice[i] = nil
		}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Last(slice)
	}
}

func BenchmarkDeepEqual(b *testing.B) {
	type Person struct {
		Name string
		Age  int
	}

	p1 := &Person{Name: "Alice", Age: 30}
	p2 := &Person{Name: "Alice", Age: 30}
	p3 := &Person{Name: "Bob", Age: 25}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = DeepEqual(p1, p2)
		_ = DeepEqual(p1, p3)
	}
}
