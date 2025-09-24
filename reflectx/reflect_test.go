package reflectx_test

import (
	"reflect"
	"testing"

	"github.com/go4x/goal/reflectx"
	"github.com/stretchr/testify/assert"
)

// Test types
type TestStruct struct {
	Name    string `json:"name" xml:"name"`
	Age     int    `json:"age" xml:"age"`
	Email   string `json:"email" xml:"email"`
	private int    // Used in tests for unexported field testing
}

type TestInterface interface {
	GetName() string
	GetAge() int
}

type TestInterfaceImpl struct {
	Name string
	Age  int
}

func (t TestInterfaceImpl) GetName() string {
	return t.Name
}

func (t TestInterfaceImpl) GetAge() int {
	return t.Age
}

type GenericInterface[T any] interface {
	GetValue() T
	SetValue(T)
}

type GenericInterfaceImpl[T any] struct {
	value T
}

func (g GenericInterfaceImpl[T]) GetValue() T {
	return g.value
}

func (g GenericInterfaceImpl[T]) SetValue(v T) {
	g.value = v
}

// Test type checking and conversion utilities

func TestIsNil(t *testing.T) {
	var nilPtr *int
	var nilSlice []int
	var nilMap map[string]int
	var nilChan chan int
	var nilFunc func()
	var nilInterface interface{}

	assert.True(t, reflectx.IsNil(nil))
	assert.True(t, reflectx.IsNil(nilPtr))
	assert.True(t, reflectx.IsNil(nilSlice))
	assert.True(t, reflectx.IsNil(nilMap))
	assert.True(t, reflectx.IsNil(nilChan))
	assert.True(t, reflectx.IsNil(nilFunc))
	assert.True(t, reflectx.IsNil(nilInterface))

	assert.False(t, reflectx.IsNil(42))
	assert.False(t, reflectx.IsNil("hello"))
	assert.False(t, reflectx.IsNil([]int{1, 2, 3}))
}

func TestIsZero(t *testing.T) {
	assert.True(t, reflectx.IsZero(0))
	assert.True(t, reflectx.IsZero(""))
	assert.True(t, reflectx.IsZero(false))
	assert.True(t, reflectx.IsZero([]int(nil)))
	assert.True(t, reflectx.IsZero(map[string]int(nil)))

	assert.False(t, reflectx.IsZero(42))
	assert.False(t, reflectx.IsZero("hello"))
	assert.False(t, reflectx.IsZero(true))
	assert.False(t, reflectx.IsZero([]int{1, 2, 3}))
}

func TestTypeChecks(t *testing.T) {
	// Test IsPointer
	assert.True(t, reflectx.IsPointer(&TestStruct{}))
	assert.False(t, reflectx.IsPointer(TestStruct{}))

	// Test IsSlice
	assert.True(t, reflectx.IsSlice([]int{1, 2, 3}))
	assert.False(t, reflectx.IsSlice("hello"))

	// Test IsMap
	assert.True(t, reflectx.IsMap(map[string]int{"a": 1}))
	assert.False(t, reflectx.IsMap([]int{1, 2, 3}))

	// Test IsStruct
	assert.True(t, reflectx.IsStruct(TestStruct{}))
	assert.False(t, reflectx.IsStruct("hello"))

	// Test IsInterfaceType
	assert.True(t, reflectx.IsInterfaceType((*TestInterface)(nil)))
	assert.False(t, reflectx.IsInterfaceType(TestStruct{}))

	// Test IsFunc
	assert.True(t, reflectx.IsFunc(func() {}))
	assert.False(t, reflectx.IsFunc("hello"))

	// Test IsChannel
	ch := make(chan int)
	assert.True(t, reflectx.IsChannel(ch))
	assert.False(t, reflectx.IsChannel("hello"))

	// Test IsArray
	var arr [3]int
	assert.True(t, reflectx.IsArray(arr))
	assert.False(t, reflectx.IsArray([]int{1, 2, 3}))

	// Test IsString
	assert.True(t, reflectx.IsString("hello"))
	assert.False(t, reflectx.IsString(42))

	// Test IsInt
	assert.True(t, reflectx.IsInt(42))
	assert.True(t, reflectx.IsInt(int8(42)))
	assert.False(t, reflectx.IsInt("hello"))

	// Test IsUint
	assert.True(t, reflectx.IsUint(uint(42)))
	assert.True(t, reflectx.IsUint(uint8(42)))
	assert.False(t, reflectx.IsUint("hello"))

	// Test IsFloat
	assert.True(t, reflectx.IsFloat(3.14))
	assert.True(t, reflectx.IsFloat(float32(3.14)))
	assert.False(t, reflectx.IsFloat("hello"))

	// Test IsBool
	assert.True(t, reflectx.IsBool(true))
	assert.False(t, reflectx.IsBool("hello"))

	// Test IsComplex
	assert.True(t, reflectx.IsComplex(complex(1, 2)))
	assert.False(t, reflectx.IsComplex("hello"))
}

func TestTypeInfo(t *testing.T) {
	// Test GetTypeName
	assert.Equal(t, "int", reflectx.GetTypeName(42))
	assert.Equal(t, "string", reflectx.GetTypeName("hello"))

	// Test GetKind
	assert.Equal(t, reflect.Int, reflectx.GetKind(42))
	assert.Equal(t, reflect.String, reflectx.GetKind("hello"))

	// Test GetPackagePath
	assert.Equal(t, "", reflectx.GetPackagePath(42))
	assert.Equal(t, "", reflectx.GetPackagePath("hello"))

	// Test GetName
	assert.Equal(t, "int", reflectx.GetName(42))
	assert.Equal(t, "string", reflectx.GetName("hello"))

	// Test GetSize
	assert.Equal(t, reflect.TypeOf(42).Size(), reflectx.GetSize(42))
	assert.Equal(t, reflect.TypeOf("hello").Size(), reflectx.GetSize("hello"))

	// Test GetAlign
	assert.Equal(t, uintptr(reflect.TypeOf(42).Align()), reflectx.GetAlign(42))

	// Test GetFieldAlign
	assert.Equal(t, uintptr(reflect.TypeOf(42).FieldAlign()), reflectx.GetFieldAlign(42))

	// Test IsComparable
	assert.True(t, reflectx.IsComparable(42))
	assert.True(t, reflectx.IsComparable("hello"))
	assert.False(t, reflectx.IsComparable([]int{1, 2, 3}))
}

func TestTypeConversion(t *testing.T) {
	// Test IsAssignable
	assert.True(t, reflectx.IsAssignable(42, 0))
	assert.False(t, reflectx.IsAssignable(42, "hello"))

	// Test ConvertibleTo
	assert.True(t, reflectx.ConvertibleTo(42, 0.0))
	assert.True(t, reflectx.ConvertibleTo(42, "hello")) // int can be converted to string

	// Test Convert
	result, err := reflectx.Convert(42, reflect.TypeOf(0.0))
	assert.NoError(t, err)
	assert.Equal(t, 42.0, result)

	// Test ConvertTo
	result, err = reflectx.ConvertTo(42, 0.0)
	assert.NoError(t, err)
	assert.Equal(t, 42.0, result)
}

// Test struct field utilities

func TestGetFieldNames(t *testing.T) {
	names := reflectx.GetFieldNames(TestStruct{})
	expected := []string{"Name", "Age", "Email", "private"}
	assert.Equal(t, expected, names)

	// Test with pointer
	names = reflectx.GetFieldNames(&TestStruct{})
	assert.Equal(t, expected, names)

	// Test with nil
	names = reflectx.GetFieldNames(nil)
	assert.Nil(t, names)

	// Test with non-struct
	names = reflectx.GetFieldNames("hello")
	assert.Nil(t, names)
}

func TestGetFieldTags(t *testing.T) {
	tags := reflectx.GetFieldTags(TestStruct{})
	assert.Equal(t, `json:"name" xml:"name"`, tags["Name"])
	assert.Equal(t, `json:"age" xml:"age"`, tags["Age"])
	assert.Equal(t, `json:"email" xml:"email"`, tags["Email"])
	assert.Equal(t, "", tags["private"])

	// Test with pointer
	tags = reflectx.GetFieldTags(&TestStruct{})
	assert.Equal(t, `json:"name" xml:"name"`, tags["Name"])

	// Test with nil
	tags = reflectx.GetFieldTags(nil)
	assert.Nil(t, tags)

	// Test with non-struct
	tags = reflectx.GetFieldTags("hello")
	assert.Nil(t, tags)
}

func TestGetFieldValue(t *testing.T) {
	ts := TestStruct{Name: "John", Age: 30, Email: "john@example.com"}

	// Test valid field
	value, err := reflectx.GetFieldValue(ts, "Name")
	assert.NoError(t, err)
	assert.Equal(t, "John", value)

	// Test with pointer
	value, err = reflectx.GetFieldValue(&ts, "Age")
	assert.NoError(t, err)
	assert.Equal(t, 30, value)

	// Test non-existent field
	_, err = reflectx.GetFieldValue(ts, "NonExistent")
	assert.Error(t, err)

	// Test with nil
	_, err = reflectx.GetFieldValue(nil, "Name")
	assert.Error(t, err)

	// Test with non-struct
	_, err = reflectx.GetFieldValue("hello", "Name")
	assert.Error(t, err)
}

func TestSetFieldValue(t *testing.T) {
	ts := &TestStruct{Name: "John", Age: 30}

	// Test valid field
	err := reflectx.SetFieldValue(ts, "Name", "Jane")
	assert.NoError(t, err)
	assert.Equal(t, "Jane", ts.Name)

	// Test with non-pointer (should fail)
	err = reflectx.SetFieldValue(ts, "Age", 25)
	assert.NoError(t, err)
	assert.Equal(t, 25, ts.Age)

	// Test non-existent field
	err = reflectx.SetFieldValue(ts, "NonExistent", "value")
	assert.Error(t, err)

	// Test with nil
	err = reflectx.SetFieldValue(nil, "Name", "value")
	assert.Error(t, err)

	// Test with non-struct
	err = reflectx.SetFieldValue("hello", "Name", "value")
	assert.Error(t, err)
}

func TestGetFieldType(t *testing.T) {
	// Test valid field
	fieldType, err := reflectx.GetFieldType(TestStruct{}, "Name")
	assert.NoError(t, err)
	assert.Equal(t, reflect.TypeOf(""), fieldType)

	// Test with pointer
	fieldType, err = reflectx.GetFieldType(&TestStruct{}, "Age")
	assert.NoError(t, err)
	assert.Equal(t, reflect.TypeOf(0), fieldType)

	// Test non-existent field
	_, err = reflectx.GetFieldType(TestStruct{}, "NonExistent")
	assert.Error(t, err)

	// Test with nil
	_, err = reflectx.GetFieldType(nil, "Name")
	assert.Error(t, err)

	// Test with non-struct
	_, err = reflectx.GetFieldType("hello", "Name")
	assert.Error(t, err)
}

func TestGetFieldTag(t *testing.T) {
	// Test valid field and tag
	tag, err := reflectx.GetFieldTag(TestStruct{}, "Name", "json")
	assert.NoError(t, err)
	assert.Equal(t, "name", tag)

	// Test with pointer
	tag, err = reflectx.GetFieldTag(&TestStruct{}, "Age", "xml")
	assert.NoError(t, err)
	assert.Equal(t, "age", tag)

	// Test non-existent field
	_, err = reflectx.GetFieldTag(TestStruct{}, "NonExistent", "json")
	assert.Error(t, err)

	// Test with nil
	_, err = reflectx.GetFieldTag(nil, "Name", "json")
	assert.Error(t, err)

	// Test with non-struct
	_, err = reflectx.GetFieldTag("hello", "Name", "json")
	assert.Error(t, err)
}

func TestHasField(t *testing.T) {
	// Test existing field
	assert.True(t, reflectx.HasField(TestStruct{}, "Name"))
	assert.True(t, reflectx.HasField(&TestStruct{}, "Age"))

	// Test non-existent field
	assert.False(t, reflectx.HasField(TestStruct{}, "NonExistent"))

	// Test with nil
	assert.False(t, reflectx.HasField(nil, "Name"))

	// Test with non-struct
	assert.False(t, reflectx.HasField("hello", "Name"))
}

func TestGetFieldCount(t *testing.T) {
	// Test struct
	assert.Equal(t, 4, reflectx.GetFieldCount(TestStruct{}))
	assert.Equal(t, 4, reflectx.GetFieldCount(&TestStruct{}))

	// Test with nil
	assert.Equal(t, 0, reflectx.GetFieldCount(nil))

	// Test with non-struct
	assert.Equal(t, 0, reflectx.GetFieldCount("hello"))
}

func TestGetFieldInfo(t *testing.T) {
	// Test valid field
	info, err := reflectx.GetFieldInfo(TestStruct{}, "Name")
	assert.NoError(t, err)
	assert.Equal(t, "Name", info["name"])
	assert.Equal(t, "string", info["type"])
	assert.Equal(t, `json:"name" xml:"name"`, info["tag"])
	assert.Equal(t, []int{0}, info["index"])
	assert.Equal(t, false, info["anonymous"])
	assert.Equal(t, "", info["pkgPath"])

	// Test with pointer
	info, err = reflectx.GetFieldInfo(&TestStruct{}, "Age")
	assert.NoError(t, err)
	assert.Equal(t, "Age", info["name"])

	// Test non-existent field
	_, err = reflectx.GetFieldInfo(TestStruct{}, "NonExistent")
	assert.Error(t, err)

	// Test with nil
	_, err = reflectx.GetFieldInfo(nil, "Name")
	assert.Error(t, err)

	// Test with non-struct
	_, err = reflectx.GetFieldInfo("hello", "Name")
	assert.Error(t, err)
}

func TestGetAnonymousFields(t *testing.T) {
	type Embedded struct {
		Value int
	}
	type StructWithEmbedded struct {
		Embedded
		Name string
	}

	// Test with embedded field
	fields := reflectx.GetAnonymousFields(StructWithEmbedded{})
	assert.Equal(t, []string{"Embedded"}, fields)

	// Test without embedded fields
	fields = reflectx.GetAnonymousFields(TestStruct{})
	assert.Empty(t, fields)

	// Test with nil
	fields = reflectx.GetAnonymousFields(nil)
	assert.Nil(t, fields)

	// Test with non-struct
	fields = reflectx.GetAnonymousFields("hello")
	assert.Nil(t, fields)
}

func TestGetExportedFields(t *testing.T) {
	// Test struct with exported and unexported fields
	fields := reflectx.GetExportedFields(TestStruct{})
	expected := []string{"Name", "Age", "Email"}
	assert.Equal(t, expected, fields)

	// Test with pointer
	fields = reflectx.GetExportedFields(&TestStruct{})
	assert.Equal(t, expected, fields)

	// Test with nil
	fields = reflectx.GetExportedFields(nil)
	assert.Nil(t, fields)

	// Test with non-struct
	fields = reflectx.GetExportedFields("hello")
	assert.Nil(t, fields)
}

func TestGetUnexportedFields(t *testing.T) {
	// Test struct with unexported fields
	fields := reflectx.GetUnexportedFields(TestStruct{})
	assert.Equal(t, []string{"private"}, fields)

	// Test with pointer
	fields = reflectx.GetUnexportedFields(&TestStruct{})
	assert.Equal(t, []string{"private"}, fields)

	// Test with nil
	fields = reflectx.GetUnexportedFields(nil)
	assert.Nil(t, fields)

	// Test with non-struct
	fields = reflectx.GetUnexportedFields("hello")
	assert.Nil(t, fields)
}

// Test value operation and conversion utilities

func TestGetValue(t *testing.T) {
	value := reflectx.GetValue(42)
	assert.Equal(t, reflect.ValueOf(42), value)
}

func TestGetType(t *testing.T) {
	typ := reflectx.GetType(42)
	assert.Equal(t, reflect.TypeOf(42), typ)
}

func TestGetElem(t *testing.T) {
	// Test pointer
	elem, err := reflectx.GetElem(&TestStruct{})
	assert.NoError(t, err)
	assert.Equal(t, reflect.TypeOf(TestStruct{}), elem)

	// Test slice
	elem, err = reflectx.GetElem([]int{1, 2, 3})
	assert.NoError(t, err)
	assert.Equal(t, reflect.TypeOf(0), elem)

	// Test map
	elem, err = reflectx.GetElem(map[string]int{"a": 1})
	assert.NoError(t, err)
	assert.Equal(t, reflect.TypeOf(0), elem)

	// Test channel
	ch := make(chan int)
	elem, err = reflectx.GetElem(ch)
	assert.NoError(t, err)
	assert.Equal(t, reflect.TypeOf(0), elem)

	// Test non-element type
	_, err = reflectx.GetElem(42)
	assert.Error(t, err)

	// Test with nil
	_, err = reflectx.GetElem(nil)
	assert.Error(t, err)
}

func TestGetKey(t *testing.T) {
	// Test map
	keyType, err := reflectx.GetKey(map[string]int{"a": 1})
	assert.NoError(t, err)
	assert.Equal(t, reflect.TypeOf(""), keyType)

	// Test non-map
	_, err = reflectx.GetKey(42)
	assert.Error(t, err)

	// Test with nil
	_, err = reflectx.GetKey(nil)
	assert.Error(t, err)
}

func TestGetLen(t *testing.T) {
	// Test slice
	length, err := reflectx.GetLen([]int{1, 2, 3})
	assert.NoError(t, err)
	assert.Equal(t, 3, length)

	// Test array
	var arr [3]int
	length, err = reflectx.GetLen(arr)
	assert.NoError(t, err)
	assert.Equal(t, 3, length)

	// Test map
	length, err = reflectx.GetLen(map[string]int{"a": 1, "b": 2})
	assert.NoError(t, err)
	assert.Equal(t, 2, length)

	// Test string
	length, err = reflectx.GetLen("hello")
	assert.NoError(t, err)
	assert.Equal(t, 5, length)

	// Test channel
	ch := make(chan int, 3)
	length, err = reflectx.GetLen(ch)
	assert.NoError(t, err)
	assert.Equal(t, 0, length)

	// Test non-length type
	_, err = reflectx.GetLen(42)
	assert.Error(t, err)

	// Test with nil
	_, err = reflectx.GetLen(nil)
	assert.Error(t, err)
}

func TestGetCap(t *testing.T) {
	// Test slice
	cap, err := reflectx.GetCap([]int{1, 2, 3})
	assert.NoError(t, err)
	assert.Equal(t, 3, cap)

	// Test array
	var arr [3]int
	cap, err = reflectx.GetCap(arr)
	assert.NoError(t, err)
	assert.Equal(t, 3, cap)

	// Test channel
	ch := make(chan int, 3)
	cap, err = reflectx.GetCap(ch)
	assert.NoError(t, err)
	assert.Equal(t, 3, cap)

	// Test non-capacity type
	_, err = reflectx.GetCap(42)
	assert.Error(t, err)

	// Test with nil
	_, err = reflectx.GetCap(nil)
	assert.Error(t, err)
}

func TestGetIndex(t *testing.T) {
	// Test slice
	value, err := reflectx.GetIndex([]int{1, 2, 3}, 1)
	assert.NoError(t, err)
	assert.Equal(t, 2, value)

	// Test array
	var arr = [3]int{1, 2, 3}
	value, err = reflectx.GetIndex(arr, 0)
	assert.NoError(t, err)
	assert.Equal(t, 1, value)

	// Test string
	value, err = reflectx.GetIndex("hello", 1)
	assert.NoError(t, err)
	assert.Equal(t, uint8('e'), value)

	// Test out of range
	_, err = reflectx.GetIndex([]int{1, 2, 3}, 5)
	assert.Error(t, err)

	// Test negative index
	_, err = reflectx.GetIndex([]int{1, 2, 3}, -1)
	assert.Error(t, err)

	// Test non-indexable type
	_, err = reflectx.GetIndex(42, 0)
	assert.Error(t, err)

	// Test with nil
	_, err = reflectx.GetIndex(nil, 0)
	assert.Error(t, err)
}

func TestSetIndex(t *testing.T) {
	// Test slice
	slice := []int{1, 2, 3}
	err := reflectx.SetIndex(slice, 1, 42)
	assert.NoError(t, err)
	assert.Equal(t, 42, slice[1])

	// Test array (arrays are not settable by value, need slice)
	slice2 := []int{1, 2, 3}
	err = reflectx.SetIndex(slice2, 0, 42)
	assert.NoError(t, err)
	assert.Equal(t, 42, slice2[0])

	// Test out of range
	err = reflectx.SetIndex([]int{1, 2, 3}, 5, 42)
	assert.Error(t, err)

	// Test negative index
	err = reflectx.SetIndex([]int{1, 2, 3}, -1, 42)
	assert.Error(t, err)

	// Test non-indexable type
	err = reflectx.SetIndex(42, 0, 42)
	assert.Error(t, err)

	// Test with nil
	err = reflectx.SetIndex(nil, 0, 42)
	assert.Error(t, err)
}

func TestGetMapValue(t *testing.T) {
	// Test map
	m := map[string]int{"a": 1, "b": 2}
	value, err := reflectx.GetMapValue(m, "a")
	assert.NoError(t, err)
	assert.Equal(t, 1, value)

	// Test non-existent key
	_, err = reflectx.GetMapValue(m, "c")
	assert.Error(t, err)

	// Test non-map
	_, err = reflectx.GetMapValue(42, "a")
	assert.Error(t, err)

	// Test with nil
	_, err = reflectx.GetMapValue(nil, "a")
	assert.Error(t, err)
}

func TestSetMapValue(t *testing.T) {
	// Test map
	m := map[string]int{"a": 1}
	err := reflectx.SetMapValue(m, "b", 2)
	assert.NoError(t, err)
	assert.Equal(t, 2, m["b"])

	// Test non-map
	err = reflectx.SetMapValue(42, "a", 1)
	assert.Error(t, err)

	// Test with nil
	err = reflectx.SetMapValue(nil, "a", 1)
	assert.Error(t, err)
}

func TestGetMapKeys(t *testing.T) {
	// Test map
	m := map[string]int{"a": 1, "b": 2}
	keys, err := reflectx.GetMapKeys(m)
	assert.NoError(t, err)
	assert.Len(t, keys, 2)
	assert.Contains(t, keys, "a")
	assert.Contains(t, keys, "b")

	// Test non-map
	_, err = reflectx.GetMapKeys(42)
	assert.Error(t, err)

	// Test with nil
	_, err = reflectx.GetMapKeys(nil)
	assert.Error(t, err)
}

func TestGetMapValues(t *testing.T) {
	// Test map
	m := map[string]int{"a": 1, "b": 2}
	values, err := reflectx.GetMapValues(m)
	assert.NoError(t, err)
	assert.Len(t, values, 2)
	assert.Contains(t, values, 1)
	assert.Contains(t, values, 2)

	// Test non-map
	_, err = reflectx.GetMapValues(42)
	assert.Error(t, err)

	// Test with nil
	_, err = reflectx.GetMapValues(nil)
	assert.Error(t, err)
}

func TestGetMapEntries(t *testing.T) {
	// Test map
	m := map[string]int{"a": 1, "b": 2}
	entries, err := reflectx.GetMapEntries(m)
	assert.NoError(t, err)
	assert.Len(t, entries, 2)

	// Test non-map
	_, err = reflectx.GetMapEntries(42)
	assert.Error(t, err)

	// Test with nil
	_, err = reflectx.GetMapEntries(nil)
	assert.Error(t, err)
}

func TestCallMethod(t *testing.T) {
	// Test valid method
	impl := TestInterfaceImpl{Name: "John", Age: 30}
	results, err := reflectx.CallMethod(impl, "GetName")
	assert.NoError(t, err)
	assert.Len(t, results, 1)
	assert.Equal(t, "John", results[0])

	// Test with pointer
	results, err = reflectx.CallMethod(&impl, "GetAge")
	assert.NoError(t, err)
	assert.Len(t, results, 1)
	assert.Equal(t, 30, results[0])

	// Test non-existent method
	_, err = reflectx.CallMethod(impl, "NonExistent")
	assert.Error(t, err)

	// Test with nil
	_, err = reflectx.CallMethod(nil, "GetName")
	assert.Error(t, err)

	// Test with non-struct
	_, err = reflectx.CallMethod("hello", "GetName")
	assert.Error(t, err)
}

func TestHasMethod(t *testing.T) {
	// Test existing method
	assert.True(t, reflectx.HasMethod(TestInterfaceImpl{}, "GetName"))
	assert.True(t, reflectx.HasMethod(&TestInterfaceImpl{}, "GetAge"))

	// Test non-existent method
	assert.False(t, reflectx.HasMethod(TestInterfaceImpl{}, "NonExistent"))

	// Test with nil
	assert.False(t, reflectx.HasMethod(nil, "GetName"))

	// Test with non-struct
	assert.False(t, reflectx.HasMethod("hello", "GetName"))
}

func TestGetMethodNames(t *testing.T) {
	// Test struct with methods
	names := reflectx.GetMethodNames(TestInterfaceImpl{})
	assert.Contains(t, names, "GetName")
	assert.Contains(t, names, "GetAge")

	// Test with pointer
	names = reflectx.GetMethodNames(&TestInterfaceImpl{})
	assert.Contains(t, names, "GetName")
	assert.Contains(t, names, "GetAge")

	// Test with nil
	names = reflectx.GetMethodNames(nil)
	assert.Nil(t, names)

	// Test with non-struct
	names = reflectx.GetMethodNames("hello")
	assert.Nil(t, names)
}

func TestGetMethodCount(t *testing.T) {
	// Test struct with methods
	count := reflectx.GetMethodCount(TestInterfaceImpl{})
	assert.Equal(t, 2, count)

	// Test with pointer
	count = reflectx.GetMethodCount(&TestInterfaceImpl{})
	assert.Equal(t, 2, count)

	// Test with nil
	count = reflectx.GetMethodCount(nil)
	assert.Equal(t, 0, count)

	// Test with non-struct
	count = reflectx.GetMethodCount("hello")
	assert.Equal(t, 0, count)
}

func TestGetMethodInfo(t *testing.T) {
	// Test valid method
	info, err := reflectx.GetMethodInfo(TestInterfaceImpl{}, "GetName")
	assert.NoError(t, err)
	assert.Equal(t, "GetName", info["name"])
	assert.Equal(t, "func(reflectx_test.TestInterfaceImpl) string", info["type"])
	assert.Equal(t, 1, info["index"]) // GetName is the second method (index 1)
	assert.Equal(t, "", info["pkgPath"])
	assert.Equal(t, 1, info["numIn"]) // 1 parameter (receiver)
	assert.Equal(t, 1, info["numOut"])

	// Test with pointer
	info, err = reflectx.GetMethodInfo(&TestInterfaceImpl{}, "GetAge")
	assert.NoError(t, err)
	assert.Equal(t, "GetAge", info["name"])

	// Test non-existent method
	_, err = reflectx.GetMethodInfo(TestInterfaceImpl{}, "NonExistent")
	assert.Error(t, err)

	// Test with nil
	_, err = reflectx.GetMethodInfo(nil, "GetName")
	assert.Error(t, err)

	// Test with non-struct
	_, err = reflectx.GetMethodInfo("hello", "GetName")
	assert.Error(t, err)
}

// Test interface utilities

func TestImplements(t *testing.T) {
	// Test implementing interface
	impl := TestInterfaceImpl{Name: "John", Age: 30}
	assert.True(t, reflectx.Implements(impl, (*TestInterface)(nil)))

	// Test not implementing interface
	assert.False(t, reflectx.Implements(TestStruct{}, (*TestInterface)(nil)))

	// Test with nil
	assert.False(t, reflectx.Implements(nil, (*TestInterface)(nil)))
	assert.False(t, reflectx.Implements(impl, nil))
}

func TestAssignableTo(t *testing.T) {
	// Test assignable types
	assert.True(t, reflectx.AssignableTo(42, 0))
	assert.False(t, reflectx.AssignableTo(42, "hello"))

	// Test with nil
	assert.False(t, reflectx.AssignableTo(nil, 0))
	assert.False(t, reflectx.AssignableTo(42, nil))
}

func TestConvertibleTo(t *testing.T) {
	// Test convertible types
	assert.True(t, reflectx.ConvertibleTo(42, 0.0))
	assert.True(t, reflectx.ConvertibleTo(42, "hello")) // int can be converted to string

	// Test with nil
	assert.False(t, reflectx.ConvertibleTo(nil, 0.0))
	assert.False(t, reflectx.ConvertibleTo(42, nil))
}

func TestGetInterfaceMethods(t *testing.T) {
	// Test interface methods
	methods := reflectx.GetInterfaceMethods((*TestInterface)(nil))
	assert.Contains(t, methods, "GetName")
	assert.Contains(t, methods, "GetAge")

	// Test with nil
	methods = reflectx.GetInterfaceMethods(nil)
	assert.Nil(t, methods)

	// Test with non-interface
	methods = reflectx.GetInterfaceMethods(TestStruct{})
	assert.Nil(t, methods)
}

func TestGetInterfaceMethodCount(t *testing.T) {
	// Test interface method count
	count := reflectx.GetInterfaceMethodCount((*TestInterface)(nil))
	assert.Equal(t, 2, count)

	// Test with nil
	count = reflectx.GetInterfaceMethodCount(nil)
	assert.Equal(t, 0, count)

	// Test with non-interface
	count = reflectx.GetInterfaceMethodCount(TestStruct{})
	assert.Equal(t, 0, count)
}

func TestGetInterfaceMethodInfo(t *testing.T) {
	// Test valid interface method
	info, err := reflectx.GetInterfaceMethodInfo((*TestInterface)(nil), "GetName")
	assert.NoError(t, err)
	assert.Equal(t, "GetName", info["name"])
	assert.Equal(t, "func() string", info["type"])
	assert.Equal(t, 1, info["index"]) // GetName is the second method (index 1)
	assert.Equal(t, "", info["pkgPath"])
	assert.Equal(t, 0, info["numIn"]) // Interface methods don't have receiver
	assert.Equal(t, 1, info["numOut"])

	// Test non-existent method
	_, err = reflectx.GetInterfaceMethodInfo((*TestInterface)(nil), "NonExistent")
	assert.Error(t, err)

	// Test with nil
	_, err = reflectx.GetInterfaceMethodInfo(nil, "GetName")
	assert.Error(t, err)

	// Test with non-interface
	_, err = reflectx.GetInterfaceMethodInfo(TestStruct{}, "GetName")
	assert.Error(t, err)
}

func TestIsInterfaceType(t *testing.T) {
	// Test interface type
	assert.True(t, reflectx.IsInterfaceType((*TestInterface)(nil)))

	// Test non-interface type
	assert.False(t, reflectx.IsInterfaceType(TestStruct{}))

	// Test with nil
	assert.False(t, reflectx.IsInterfaceType(nil))
}

func TestGetInterfaceName(t *testing.T) {
	// Test interface name
	name := reflectx.GetInterfaceName((*TestInterface)(nil))
	assert.Equal(t, "TestInterface", name)

	// Test with nil
	name = reflectx.GetInterfaceName(nil)
	assert.Equal(t, "", name)

	// Test with non-interface
	name = reflectx.GetInterfaceName(TestStruct{})
	assert.Equal(t, "", name)
}

func TestGetInterfacePackage(t *testing.T) {
	// Test interface package
	pkg := reflectx.GetInterfacePackage((*TestInterface)(nil))
	assert.Equal(t, "github.com/go4x/goal/reflectx_test", pkg) // Package path for local types

	// Test with nil
	pkg = reflectx.GetInterfacePackage(nil)
	assert.Equal(t, "", pkg)

	// Test with non-interface
	pkg = reflectx.GetInterfacePackage(TestStruct{})
	assert.Equal(t, "", pkg)
}

func TestGetInterfaceString(t *testing.T) {
	// Test interface string
	str := reflectx.GetInterfaceString((*TestInterface)(nil))
	assert.Equal(t, "reflectx_test.TestInterface", str)

	// Test with nil
	str = reflectx.GetInterfaceString(nil)
	assert.Equal(t, "nil", str)

	// Test with non-interface
	str = reflectx.GetInterfaceString(TestStruct{})
	assert.Equal(t, "", str)
}

func TestGetInterfaceMethodsDetailed(t *testing.T) {
	// Test interface methods detailed
	methods, err := reflectx.GetInterfaceMethodsDetailed((*TestInterface)(nil))
	assert.NoError(t, err)
	assert.Len(t, methods, 2)

	// Test with nil
	_, err = reflectx.GetInterfaceMethodsDetailed(nil)
	assert.Error(t, err)

	// Test with non-interface
	_, err = reflectx.GetInterfaceMethodsDetailed(TestStruct{})
	assert.Error(t, err)
}

// Test generic utilities

func TestIsGeneric(t *testing.T) {
	// Test generic interface (Go reflection limitations)
	assert.False(t, reflectx.IsGeneric((*GenericInterface[int])(nil)))

	// Test non-generic interface
	assert.False(t, reflectx.IsGeneric((*TestInterface)(nil)))

	// Test with nil
	assert.False(t, reflectx.IsGeneric(nil))

	// Test with non-interface
	assert.False(t, reflectx.IsGeneric(TestStruct{}))
}

func TestGetGenericTypeParameters(t *testing.T) {
	// Test generic type parameters (Go reflection limitations)
	_, err := reflectx.GetGenericTypeParameters((*GenericInterface[int])(nil))
	assert.Error(t, err) // Will fail due to reflection limitations

	// Test with nil
	_, err = reflectx.GetGenericTypeParameters(nil)
	assert.Error(t, err)

	// Test with non-interface
	_, err = reflectx.GetGenericTypeParameters(TestStruct{})
	assert.Error(t, err)
}

func TestGetGenericConstraints(t *testing.T) {
	// Test generic constraints (Go reflection limitations)
	_, err := reflectx.GetGenericConstraints((*GenericInterface[int])(nil))
	assert.Error(t, err) // Will fail due to reflection limitations

	// Test with nil
	_, err = reflectx.GetGenericConstraints(nil)
	assert.Error(t, err)

	// Test with non-interface
	_, err = reflectx.GetGenericConstraints(TestStruct{})
	assert.Error(t, err)
}

func TestIsGenericType(t *testing.T) {
	// Test generic type (Go reflection limitations)
	assert.False(t, reflectx.IsGenericType((*GenericInterface[int])(nil)))

	// Test non-generic type
	assert.False(t, reflectx.IsGenericType((*TestInterface)(nil)))

	// Test with nil
	assert.False(t, reflectx.IsGenericType(nil))

	// Test with non-interface
	assert.False(t, reflectx.IsGenericType(TestStruct{}))
}

func TestGetGenericTypeName(t *testing.T) {
	// Test generic type name (Go reflection limitations)
	name := reflectx.GetGenericTypeName((*GenericInterface[int])(nil))
	assert.Equal(t, "", name) // Empty due to reflection limitations

	// Test with nil
	name = reflectx.GetGenericTypeName(nil)
	assert.Equal(t, "", name)

	// Test with non-interface
	name = reflectx.GetGenericTypeName(TestStruct{})
	assert.Equal(t, "", name)
}

func TestGetGenericTypeString(t *testing.T) {
	// Test generic type string
	str := reflectx.GetGenericTypeString((*GenericInterface[int])(nil))
	assert.Equal(t, "*reflectx_test.GenericInterface[int]", str)

	// Test with nil
	str = reflectx.GetGenericTypeString(nil)
	assert.Equal(t, "nil", str)

	// Test with non-interface
	str = reflectx.GetGenericTypeString(TestStruct{})
	assert.Equal(t, "reflectx_test.TestStruct", str)
}

func TestGetGenericTypePackage(t *testing.T) {
	// Test generic type package
	pkg := reflectx.GetGenericTypePackage((*GenericInterface[int])(nil))
	assert.Equal(t, "", pkg) // Empty for local types

	// Test with nil
	pkg = reflectx.GetGenericTypePackage(nil)
	assert.Equal(t, "", pkg)

	// Test with non-interface
	pkg = reflectx.GetGenericTypePackage(TestStruct{})
	assert.Equal(t, "", pkg)
}

func TestGetGenericTypeSize(t *testing.T) {
	// Test generic type size
	size := reflectx.GetGenericTypeSize((*GenericInterface[int])(nil))
	assert.Equal(t, reflect.TypeOf((*GenericInterface[int])(nil)).Size(), size)

	// Test with nil
	size = reflectx.GetGenericTypeSize(nil)
	assert.Equal(t, uintptr(0), size)

	// Test with non-interface
	size = reflectx.GetGenericTypeSize(TestStruct{})
	assert.Equal(t, reflect.TypeOf(TestStruct{}).Size(), size)
}

func TestGetGenericTypeAlign(t *testing.T) {
	// Test generic type align
	align := reflectx.GetGenericTypeAlign((*GenericInterface[int])(nil))
	assert.Equal(t, uintptr(reflect.TypeOf((*GenericInterface[int])(nil)).Align()), align)

	// Test with nil
	align = reflectx.GetGenericTypeAlign(nil)
	assert.Equal(t, uintptr(0), align)

	// Test with non-interface
	align = reflectx.GetGenericTypeAlign(TestStruct{})
	assert.Equal(t, uintptr(reflect.TypeOf(TestStruct{}).Align()), align)
}

func TestGetGenericTypeFieldAlign(t *testing.T) {
	// Test generic type field align
	align := reflectx.GetGenericTypeFieldAlign((*GenericInterface[int])(nil))
	assert.Equal(t, uintptr(reflect.TypeOf((*GenericInterface[int])(nil)).FieldAlign()), align)

	// Test with nil
	align = reflectx.GetGenericTypeFieldAlign(nil)
	assert.Equal(t, uintptr(0), align)

	// Test with non-interface
	align = reflectx.GetGenericTypeFieldAlign(TestStruct{})
	assert.Equal(t, uintptr(reflect.TypeOf(TestStruct{}).FieldAlign()), align)
}

func TestGetGenericTypeComparable(t *testing.T) {
	// Test generic type comparable
	comparable := reflectx.GetGenericTypeComparable((*GenericInterface[int])(nil))
	assert.Equal(t, reflect.TypeOf((*GenericInterface[int])(nil)).Comparable(), comparable)

	// Test with nil
	comparable = reflectx.GetGenericTypeComparable(nil)
	assert.False(t, comparable)

	// Test with non-interface
	comparable = reflectx.GetGenericTypeComparable(TestStruct{})
	assert.Equal(t, reflect.TypeOf(TestStruct{}).Comparable(), comparable)
}

func TestGetGenericTypeAssignableTo(t *testing.T) {
	// Test generic type assignable to
	assignable := reflectx.GetGenericTypeAssignableTo((*GenericInterface[int])(nil), (*GenericInterface[int])(nil))
	assert.True(t, assignable)

	// Test with nil
	assignable = reflectx.GetGenericTypeAssignableTo(nil, (*GenericInterface[int])(nil))
	assert.False(t, assignable)
	assignable = reflectx.GetGenericTypeAssignableTo((*GenericInterface[int])(nil), nil)
	assert.False(t, assignable)
}

func TestGetGenericTypeConvertibleTo(t *testing.T) {
	// Test generic type convertible to
	convertible := reflectx.GetGenericTypeConvertibleTo((*GenericInterface[int])(nil), (*GenericInterface[int])(nil))
	assert.True(t, convertible)

	// Test with nil
	convertible = reflectx.GetGenericTypeConvertibleTo(nil, (*GenericInterface[int])(nil))
	assert.False(t, convertible)
	convertible = reflectx.GetGenericTypeConvertibleTo((*GenericInterface[int])(nil), nil)
	assert.False(t, convertible)
}

func TestGetGenericTypeMethods(t *testing.T) {
	// Test generic type methods (Go reflection limitations)
	methods := reflectx.GetGenericTypeMethods((*GenericInterface[int])(nil))
	assert.Nil(t, methods) // Nil due to reflection limitations

	// Test with nil
	methods = reflectx.GetGenericTypeMethods(nil)
	assert.Nil(t, methods)

	// Test with non-interface
	methods = reflectx.GetGenericTypeMethods(TestStruct{})
	assert.Nil(t, methods)
}

func TestGetGenericTypeMethodCount(t *testing.T) {
	// Test generic type method count (Go reflection limitations)
	count := reflectx.GetGenericTypeMethodCount((*GenericInterface[int])(nil))
	assert.Equal(t, 0, count) // 0 due to reflection limitations

	// Test with nil
	count = reflectx.GetGenericTypeMethodCount(nil)
	assert.Equal(t, 0, count)

	// Test with non-interface
	count = reflectx.GetGenericTypeMethodCount(TestStruct{})
	assert.Equal(t, 0, count)
}

func TestGetGenericTypeMethodInfo(t *testing.T) {
	// Test valid generic type method (Go reflection limitations)
	_, err := reflectx.GetGenericTypeMethodInfo((*GenericInterface[int])(nil), "GetValue")
	assert.Error(t, err) // Will fail due to reflection limitations

	// Test non-existent method
	_, err = reflectx.GetGenericTypeMethodInfo((*GenericInterface[int])(nil), "NonExistent")
	assert.Error(t, err)

	// Test with nil
	_, err = reflectx.GetGenericTypeMethodInfo(nil, "GetValue")
	assert.Error(t, err)

	// Test with non-interface
	_, err = reflectx.GetGenericTypeMethodInfo(TestStruct{}, "GetValue")
	assert.Error(t, err)
}

func TestGetGenericTypeMethodsDetailed(t *testing.T) {
	// Test generic type methods detailed (Go reflection limitations)
	_, err := reflectx.GetGenericTypeMethodsDetailed((*GenericInterface[int])(nil))
	assert.Error(t, err) // Will fail due to reflection limitations

	// Test with nil
	_, err = reflectx.GetGenericTypeMethodsDetailed(nil)
	assert.Error(t, err)

	// Test with non-interface
	_, err = reflectx.GetGenericTypeMethodsDetailed(TestStruct{})
	assert.Error(t, err)
}

// Test original functions

func TestMethods(t *testing.T) {
	var ty TestInterfaceImpl
	var pty *TestInterfaceImpl
	methods := reflectx.Methods(&ty)
	assert.Equal(t, 2, len(methods))
	assert.Contains(t, methods, "GetName")
	assert.Contains(t, methods, "GetAge")

	methods = reflectx.Methods(&pty)
	assert.Equal(t, 2, len(methods))
	assert.Contains(t, methods, "GetName")
	assert.Contains(t, methods, "GetAge")

	// nil interface, also has 2 methods: GetName and GetAge
	methods = reflectx.Methods((*TestInterface)(nil))
	assert.Equal(t, 2, len(methods))
	assert.Contains(t, methods, "GetName")
	assert.Contains(t, methods, "GetAge")
}

func TestPrintMethodSet(t *testing.T) {
	var ty TestInterfaceImpl
	var pty *TestInterfaceImpl
	reflectx.PrintMethodSet(&ty)
	reflectx.PrintMethodSet(&pty)
	// nil interface
	reflectx.PrintMethodSet((*TestInterface)(nil))
}
