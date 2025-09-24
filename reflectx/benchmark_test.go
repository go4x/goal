package reflectx_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/go4x/goal/reflectx"
)

// BenchmarkIsNil benchmarks the IsNil function
func BenchmarkIsNil(b *testing.B) {
	var ptr *int
	var value int = 42

	b.Run("NilPointer", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			reflectx.IsNil(ptr)
		}
	})

	b.Run("NonNilValue", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			reflectx.IsNil(value)
		}
	})
}

// BenchmarkIsZero benchmarks the IsZero function
func BenchmarkIsZero(b *testing.B) {
	zero := 0
	nonZero := 42

	b.Run("ZeroValue", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			reflectx.IsZero(zero)
		}
	})

	b.Run("NonZeroValue", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			reflectx.IsZero(nonZero)
		}
	})
}

// BenchmarkTypeChecks benchmarks type checking functions
func BenchmarkTypeChecks(b *testing.B) {
	ptr := &[]int{1, 2, 3}
	slice := []int{1, 2, 3}

	b.Run("IsPointer", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			reflectx.IsPointer(ptr)
		}
	})

	b.Run("IsSlice", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			reflectx.IsSlice(slice)
		}
	})
}

// BenchmarkGetTypeName benchmarks type information functions
func BenchmarkGetTypeName(b *testing.B) {
	type User struct {
		ID   int
		Name string
	}
	user := User{ID: 1, Name: "Alice"}

	b.Run("GetTypeName", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			reflectx.GetTypeName(user)
		}
	})

	b.Run("GetKind", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			reflectx.GetKind(user)
		}
	})

	b.Run("GetSize", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			reflectx.GetSize(user)
		}
	})
}

// BenchmarkConvert benchmarks type conversion
func BenchmarkConvert(b *testing.B) {
	targetType := reflect.TypeOf(0.0)

	b.Run("Convert", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			reflectx.Convert(42, targetType)
		}
	})
}

// BenchmarkGetFieldNames benchmarks struct field operations
func BenchmarkGetFieldNames(b *testing.B) {
	type Person struct {
		Name    string `json:"name"`
		Age     int    `json:"age"`
		Email   string `json:"email"`
		Address string `json:"address"`
		Phone   string `json:"phone"`
	}

	p := Person{
		Name:    "Alice",
		Age:     30,
		Email:   "alice@example.com",
		Address: "123 Main St",
		Phone:   "555-1234",
	}

	b.Run("GetFieldNames", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			reflectx.GetFieldNames(p)
		}
	})

	b.Run("GetFieldTags", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			reflectx.GetFieldTags(p)
		}
	})

	b.Run("GetFieldValue", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			reflectx.GetFieldValue(p, "Name")
		}
	})
}

// BenchmarkGetFieldInfo benchmarks detailed field information
func BenchmarkGetFieldInfo(b *testing.B) {
	type Product struct {
		ID          int     `json:"id" db:"product_id"`
		Name        string  `json:"name" db:"product_name"`
		Price       float64 `json:"price" db:"price"`
		Description string  `json:"description" db:"description"`
		Category    string  `json:"category" db:"category"`
	}

	p := Product{
		ID:          1,
		Name:        "Laptop",
		Price:       999.99,
		Description: "High-performance laptop",
		Category:    "Electronics",
	}

	b.Run("GetFieldInfo", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			reflectx.GetFieldInfo(p, "Name")
		}
	})
}

// BenchmarkGetValue benchmarks value operations
func BenchmarkGetValue(b *testing.B) {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	b.Run("GetValue", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			reflectx.GetValue(slice)
		}
	})

	b.Run("GetLen", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			reflectx.GetLen(slice)
		}
	})

	b.Run("GetCap", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			reflectx.GetCap(slice)
		}
	})
}

// BenchmarkGetIndex benchmarks index operations
func BenchmarkGetIndex(b *testing.B) {
	slice := []string{"apple", "banana", "cherry", "date", "elderberry"}

	b.Run("GetIndex", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			reflectx.GetIndex(slice, 2)
		}
	})
}

// BenchmarkGetMapValue benchmarks map operations
func BenchmarkGetMapValue(b *testing.B) {
	data := map[string]int{
		"apple":      5,
		"banana":     3,
		"cherry":     8,
		"date":       2,
		"elderberry": 1,
	}

	b.Run("GetMapValue", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			reflectx.GetMapValue(data, "banana")
		}
	})

	b.Run("GetMapKeys", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			reflectx.GetMapKeys(data)
		}
	})

	b.Run("GetMapValues", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			reflectx.GetMapValues(data)
		}
	})
}

// BenchmarkCallMethod benchmarks method calling
func BenchmarkCallMethod(b *testing.B) {
	calc := BenchmarkCalculator{Value: 10}

	b.Run("CallMethod", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			reflectx.CallMethod(calc, "Add", 5)
		}
	})
}

// BenchmarkCalculator is used for benchmarking
type BenchmarkCalculator struct {
	Value int
}

func (c BenchmarkCalculator) Add(x int) int {
	return c.Value + x
}

// BenchmarkHasMethod benchmarks method checking
func BenchmarkHasMethod(b *testing.B) {
	service := BenchmarkService{}

	b.Run("HasMethod", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			reflectx.HasMethod(service, "Process")
		}
	})
}

// BenchmarkService is used for benchmarking
type BenchmarkService struct{}

func (s BenchmarkService) Process() string {
	return "processed"
}

// BenchmarkGetMethodNames benchmarks method information
func BenchmarkGetMethodNames(b *testing.B) {
	user := BenchmarkUser{Name: "Alice"}

	b.Run("GetMethodNames", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			reflectx.GetMethodNames(user)
		}
	})

	b.Run("GetMethodCount", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			reflectx.GetMethodCount(user)
		}
	})
}

// BenchmarkUser is used for benchmarking
type BenchmarkUser struct {
	Name string
}

func (u BenchmarkUser) GetName() string {
	return u.Name
}

func (u BenchmarkUser) SetName(name string) {
	u.Name = name
}

func (u BenchmarkUser) Validate() bool {
	return len(u.Name) > 0
}

// BenchmarkImplements benchmarks interface checking
func BenchmarkImplements(b *testing.B) {
	book := BenchmarkBook{Title: "Go Programming"}

	b.Run("Implements", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			reflectx.Implements(book, (*BenchmarkReader)(nil))
		}
	})

	b.Run("AssignableTo", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			reflectx.AssignableTo(book, (*BenchmarkReader)(nil))
		}
	})

	b.Run("ConvertibleTo", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			reflectx.ConvertibleTo(42, 0.0)
		}
	})
}

// BenchmarkReader interface for benchmarking
type BenchmarkReader interface {
	Read() string
}

// BenchmarkBook is used for benchmarking
type BenchmarkBook struct {
	Title string
}

func (b BenchmarkBook) Read() string {
	return b.Title
}

// BenchmarkGetInterfaceMethods benchmarks interface method operations
func BenchmarkGetInterfaceMethods(b *testing.B) {

	b.Run("GetInterfaceMethods", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			reflectx.GetInterfaceMethods((*BenchmarkWriter)(nil))
		}
	})

	b.Run("GetInterfaceMethodCount", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			reflectx.GetInterfaceMethodCount((*BenchmarkWriter)(nil))
		}
	})
}

// BenchmarkWriter interface for benchmarking
type BenchmarkWriter interface {
	Write(data string) error
	Close() error
	Flush() error
}

// BenchmarkIsGeneric benchmarks generic type checking
func BenchmarkIsGeneric(b *testing.B) {
	var slice BenchmarkGenericSlice[int]

	b.Run("IsGeneric", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			reflectx.IsGeneric(slice)
		}
	})

	b.Run("IsGenericType", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			reflectx.IsGenericType(reflect.TypeOf(slice))
		}
	})
}

// BenchmarkGenericSlice is used for benchmarking
type BenchmarkGenericSlice[T any] []T

// BenchmarkMethods benchmarks the original Methods function
func BenchmarkMethods(b *testing.B) {
	service := &BenchmarkMethodsService{}

	b.Run("Methods", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			reflectx.Methods(service)
		}
	})
}

// BenchmarkMethodsService is used for benchmarking
type BenchmarkMethodsService struct{}

func (s BenchmarkMethodsService) Process() string {
	return "processed"
}

func (s BenchmarkMethodsService) Validate() bool {
	return true
}

func (s BenchmarkMethodsService) Initialize() error {
	return nil
}

// BenchmarkLargeStruct benchmarks operations on large structs
func BenchmarkLargeStruct(b *testing.B) {
	type LargeStruct struct {
		Field1  string
		Field2  int
		Field3  float64
		Field4  bool
		Field5  string
		Field6  int
		Field7  float64
		Field8  bool
		Field9  string
		Field10 int
		Field11 float64
		Field12 bool
		Field13 string
		Field14 int
		Field15 float64
		Field16 bool
		Field17 string
		Field18 int
		Field19 float64
		Field20 bool
	}

	large := LargeStruct{
		Field1: "value1", Field2: 1, Field3: 1.1, Field4: true,
		Field5: "value5", Field6: 6, Field7: 6.6, Field8: false,
		Field9: "value9", Field10: 10, Field11: 10.10, Field12: true,
		Field13: "value13", Field14: 14, Field15: 14.14, Field16: false,
		Field17: "value17", Field18: 18, Field19: 18.18, Field20: true,
	}

	b.Run("GetFieldNames", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			reflectx.GetFieldNames(large)
		}
	})

	b.Run("GetFieldCount", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			reflectx.GetFieldCount(large)
		}
	})
}

// BenchmarkLargeSlice benchmarks operations on large slices
func BenchmarkLargeSlice(b *testing.B) {
	// Create a large slice
	slice := make([]int, 1000)
	for i := range slice {
		slice[i] = i
	}

	b.Run("GetLen", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			reflectx.GetLen(slice)
		}
	})

	b.Run("GetCap", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			reflectx.GetCap(slice)
		}
	})

	b.Run("GetIndex", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			reflectx.GetIndex(slice, 500)
		}
	})
}

// BenchmarkLargeMap benchmarks operations on large maps
func BenchmarkLargeMap(b *testing.B) {
	// Create a large map
	data := make(map[string]int, 1000)
	for i := 0; i < 1000; i++ {
		data[fmt.Sprintf("key%d", i)] = i
	}

	b.Run("GetMapKeys", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			reflectx.GetMapKeys(data)
		}
	})

	b.Run("GetMapValues", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			reflectx.GetMapValues(data)
		}
	})

	b.Run("GetMapValue", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			reflectx.GetMapValue(data, "key500")
		}
	})
}
