package reflectx_test

import (
	"fmt"
	"reflect"

	"github.com/go4x/goal/reflectx"
)

// ExampleIsNil demonstrates how to check if a value is nil
func ExampleIsNil() {
	var ptr *int
	fmt.Println("Pointer is nil:", reflectx.IsNil(ptr))
	fmt.Println("Int is nil:", reflectx.IsNil(42))
	fmt.Println("Nil interface is nil:", reflectx.IsNil(nil))
	// Output:
	// Pointer is nil: true
	// Int is nil: false
	// Nil interface is nil: true
}

// ExampleIsZero demonstrates how to check if a value is zero
func ExampleIsZero() {
	fmt.Println("Zero int:", reflectx.IsZero(0))
	fmt.Println("Non-zero int:", reflectx.IsZero(42))
	fmt.Println("Empty string:", reflectx.IsZero(""))
	fmt.Println("Non-empty string:", reflectx.IsZero("hello"))
	// Output:
	// Zero int: true
	// Non-zero int: false
	// Empty string: true
	// Non-empty string: false
}

// ExampleIsPointer demonstrates various type checking functions
func ExampleIsPointer() {
	// Check if value is a pointer
	ptr := &[]int{1, 2, 3}
	fmt.Println("Is pointer:", reflectx.IsPointer(ptr))
	fmt.Println("Is slice:", reflectx.IsSlice(*ptr))

	// Check if value is a struct
	type Person struct {
		Name string
		Age  int
	}
	p := Person{Name: "Alice", Age: 30}
	fmt.Println("Is struct:", reflectx.IsStruct(p))
	fmt.Println("Is string:", reflectx.IsString(p.Name))
	// Output:
	// Is pointer: true
	// Is slice: true
	// Is struct: true
	// Is string: true
}

// ExampleGetTypeName demonstrates how to get type information
func ExampleGetTypeName() {
	type User struct {
		ID   int
		Name string
	}
	user := User{ID: 1, Name: "Alice"}

	fmt.Println("Type name:", reflectx.GetTypeName(user))
	fmt.Println("Kind:", reflectx.GetKind(user))
	fmt.Println("Size:", reflectx.GetSize(user))
	// Output:
	// Type name: reflectx_test.User
	// Kind: struct
	// Size: 24
}

// ExampleConvert demonstrates type conversion
func ExampleConvert() {
	// Convert int to float64
	result, err := reflectx.Convert(42, reflect.TypeOf(0.0))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Converted value: %v (type: %T)\n", result, result)
	// Output:
	// Converted value: 42 (type: float64)
}

// ExampleGetFieldNames demonstrates struct field operations
func ExampleGetFieldNames() {
	type Person struct {
		Name    string `json:"name"`
		Age     int    `json:"age"`
		Email   string `json:"email"`
		private int    // unexported field
	}

	p := Person{Name: "Alice", Age: 30, Email: "alice@example.com"}

	// Get field names
	names := reflectx.GetFieldNames(p)
	fmt.Println("Field names:", names)

	// Get field tags
	tags := reflectx.GetFieldTags(p)
	// Remove empty tags for cleaner output
	cleanTags := make(map[string]string)
	for k, v := range tags {
		if v != "" {
			cleanTags[k] = v
		}
	}
	fmt.Println("Field tags:", cleanTags)

	// Get field value
	value, err := reflectx.GetFieldValue(p, "Name")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Name field value:", value)
	// Output:
	// Field names: [Name Age Email private]
	// Field tags: map[Age:json:"age" Email:json:"email" Name:json:"name"]
	// Name field value: Alice
}

// ExampleGetFieldValue demonstrates getting and setting field values
func ExampleGetFieldValue() {
	type Config struct {
		Host string
		Port int
	}

	config := Config{Host: "localhost", Port: 8080}

	// Get field value
	host, err := reflectx.GetFieldValue(config, "Host")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Host:", host)

	// Set field value (requires pointer)
	err = reflectx.SetFieldValue(&config, "Port", 9090)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("New port:", config.Port)
	// Output:
	// Host: localhost
	// New port: 9090
}

// ExampleGetFieldInfo demonstrates getting detailed field information
func ExampleGetFieldInfo() {
	type Product struct {
		ID    int     `json:"id" db:"product_id"`
		Name  string  `json:"name" db:"product_name"`
		Price float64 `json:"price" db:"price"`
	}

	p := Product{ID: 1, Name: "Laptop", Price: 999.99}

	// Get field info
	info, err := reflectx.GetFieldInfo(p, "Name")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Field name:", info["name"])
	fmt.Println("Field type:", info["type"])
	fmt.Println("Field tag:", info["tag"])
	fmt.Println("Is exported:", true) // Name field is exported
	// Output:
	// Field name: Name
	// Field type: string
	// Field tag: json:"name" db:"product_name"
	// Is exported: true
}

// ExampleGetValue demonstrates value operations
func ExampleGetValue() {
	// Get value and type
	value := []int{1, 2, 3, 4, 5}
	fmt.Println("Value:", reflectx.GetValue(value))
	fmt.Println("Type:", reflectx.GetType(value))
	length, _ := reflectx.GetLen(value)
	fmt.Println("Length:", length)
	capacity, _ := reflectx.GetCap(value)
	fmt.Println("Capacity:", capacity)
	// Output:
	// Value: [1 2 3 4 5]
	// Type: []int
	// Length: 5
	// Capacity: 5
}

// ExampleGetIndex demonstrates index operations
func ExampleGetIndex() {
	// Get element by index
	slice := []string{"apple", "banana", "cherry"}
	element, err := reflectx.GetIndex(slice, 1)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Element at index 1:", element)

	// Set element by index
	err = reflectx.SetIndex(slice, 1, "orange")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("After setting:", slice)
	// Output:
	// Element at index 1: banana
	// After setting: [apple orange cherry]
}

// ExampleGetMapValue demonstrates map operations
func ExampleGetMapValue() {
	// Create a map
	data := map[string]int{
		"apple":  5,
		"banana": 3,
		"cherry": 8,
	}

	// Get map value
	value, err := reflectx.GetMapValue(data, "banana")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Banana count:", value)

	// Set map value
	err = reflectx.SetMapValue(data, "orange", 7)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Orange count:", data["orange"])
	// Output:
	// Banana count: 3
	// Orange count: 7
}

// ExampleCallMethod demonstrates method calling
func ExampleCallMethod() {
	calc := ExampleCalculator{Value: 10}

	// Call method
	result, err := reflectx.CallMethod(calc, "Add", 5)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Result:", result[0])
	// Output:
	// Result: 15
}

// ExampleCalculator is used for demonstration
type ExampleCalculator struct {
	Value int
}

func (c ExampleCalculator) Add(x int) int {
	return c.Value + x
}

// ExampleHasMethod demonstrates method checking
func ExampleHasMethod() {
	service := ExampleService{}

	// Check if method exists
	hasProcess := reflectx.HasMethod(service, "Process")
	hasNonExistent := reflectx.HasMethod(service, "NonExistent")

	fmt.Println("Has Process method:", hasProcess)
	fmt.Println("Has NonExistent method:", hasNonExistent)
	// Output:
	// Has Process method: true
	// Has NonExistent method: false
}

// ExampleService is used for demonstration
type ExampleService struct{}

func (s ExampleService) Process() string {
	return "processed"
}

// ExampleGetMethodNames demonstrates getting method information
func ExampleGetMethodNames() {
	user := ExampleUser{Name: "Alice"}

	// Get method names
	names := reflectx.GetMethodNames(user)
	fmt.Println("Method names:", names)

	// Get method count
	count := reflectx.GetMethodCount(user)
	fmt.Println("Method count:", count)
	// Output:
	// Method names: [GetName SetName]
	// Method count: 2
}

// ExampleUser is used for demonstration
type ExampleUser struct {
	Name string
}

func (u ExampleUser) GetName() string {
	return u.Name
}

func (u ExampleUser) SetName(name string) {
	u.Name = name
}

// ExampleImplements demonstrates interface checking
func ExampleImplements() {
	book := ExampleBook{Title: "Go Programming"}

	// Check if implements interface
	implements := reflectx.Implements(book, (*ExampleReader)(nil))
	fmt.Println("Implements Reader:", implements)

	// Check assignability
	assignable := reflectx.AssignableTo(book, (*ExampleReader)(nil))
	fmt.Println("Assignable to Reader:", assignable)
	// Output:
	// Implements Reader: true
	// Assignable to Reader: false
}

// ExampleReader interface for demonstration
type ExampleReader interface {
	Read() string
}

// ExampleBook is used for demonstration
type ExampleBook struct {
	Title string
}

func (b ExampleBook) Read() string {
	return b.Title
}

// ExampleGetInterfaceMethods demonstrates interface method operations
func ExampleGetInterfaceMethods() {
	// Get interface methods
	methods := reflectx.GetInterfaceMethods((*ExampleWriter)(nil))
	fmt.Println("Interface methods:", methods)

	// Get method count
	count := reflectx.GetInterfaceMethodCount((*ExampleWriter)(nil))
	fmt.Println("Method count:", count)
	// Output:
	// Interface methods: [Close Write]
	// Method count: 2
}

// ExampleWriter interface for demonstration
type ExampleWriter interface {
	Write(data string) error
	Close() error
}

// ExampleIsGeneric demonstrates generic type checking
func ExampleIsGeneric() {
	// Check if type is generic
	var slice ExampleGenericSlice[int]

	fmt.Println("Is generic:", reflectx.IsGeneric(slice))
	fmt.Println("Is generic type:", reflectx.IsGenericType(reflect.TypeOf(slice)))
	// Output:
	// Is generic: false
	// Is generic type: false
}

// ExampleGenericSlice is used for demonstration
type ExampleGenericSlice[T any] []T

// ExampleMethods demonstrates the original Methods function
func ExampleMethods() {
	service := &ExampleMethodsService{}

	// Get methods
	methods := reflectx.Methods(service)
	fmt.Println("Methods:", methods)
	// Output:
	// Methods: [Process Validate]
}

// ExampleMethodsService is used for demonstration
type ExampleMethodsService struct{}

func (s ExampleMethodsService) Process() string {
	return "processed"
}

func (s ExampleMethodsService) Validate() bool {
	return true
}

// ExamplePrintMethodSet demonstrates the original PrintMethodSet function
func ExamplePrintMethodSet() {
	handler := &ExampleHandler{}

	// Print method set
	reflectx.PrintMethodSet(handler)
	// Output:
	// reflectx_test.ExampleHandler's method set:
	//   - Handle
}

// ExampleHandler is used for demonstration
type ExampleHandler struct{}

func (h ExampleHandler) Handle() string {
	return "handled"
}
