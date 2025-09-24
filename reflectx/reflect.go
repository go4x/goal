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

// Type checking and conversion utilities

// IsNil checks if a value is nil
func IsNil(v interface{}) bool {
	if v == nil {
		return true
	}
	rv := reflect.ValueOf(v)
	switch rv.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice:
		return rv.IsNil()
	default:
		return false
	}
}

// IsZero checks if a value is the zero value of its type
func IsZero(v interface{}) bool {
	rv := reflect.ValueOf(v)
	return rv.IsZero()
}

// IsPointer checks if a value is a pointer
func IsPointer(v interface{}) bool {
	rv := reflect.ValueOf(v)
	return rv.Kind() == reflect.Ptr
}

// IsSlice checks if a value is a slice
func IsSlice(v interface{}) bool {
	rv := reflect.ValueOf(v)
	return rv.Kind() == reflect.Slice
}

// IsMap checks if a value is a map
func IsMap(v interface{}) bool {
	rv := reflect.ValueOf(v)
	return rv.Kind() == reflect.Map
}

// IsStruct checks if a value is a struct
func IsStruct(v interface{}) bool {
	rv := reflect.ValueOf(v)
	return rv.Kind() == reflect.Struct
}

// IsInterface checks if a value is an interface
func IsInterface(v interface{}) bool {
	rv := reflect.ValueOf(v)
	return rv.Kind() == reflect.Interface
}

// IsFunc checks if a value is a function
func IsFunc(v interface{}) bool {
	rv := reflect.ValueOf(v)
	return rv.Kind() == reflect.Func
}

// IsChannel checks if a value is a channel
func IsChannel(v interface{}) bool {
	rv := reflect.ValueOf(v)
	return rv.Kind() == reflect.Chan
}

// IsArray checks if a value is an array
func IsArray(v interface{}) bool {
	rv := reflect.ValueOf(v)
	return rv.Kind() == reflect.Array
}

// IsString checks if a value is a string
func IsString(v interface{}) bool {
	rv := reflect.ValueOf(v)
	return rv.Kind() == reflect.String
}

// IsInt checks if a value is an integer type
func IsInt(v interface{}) bool {
	rv := reflect.ValueOf(v)
	switch rv.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return true
	default:
		return false
	}
}

// IsUint checks if a value is an unsigned integer type
func IsUint(v interface{}) bool {
	rv := reflect.ValueOf(v)
	switch rv.Kind() {
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return true
	default:
		return false
	}
}

// IsFloat checks if a value is a float type
func IsFloat(v interface{}) bool {
	rv := reflect.ValueOf(v)
	switch rv.Kind() {
	case reflect.Float32, reflect.Float64:
		return true
	default:
		return false
	}
}

// IsBool checks if a value is a boolean
func IsBool(v interface{}) bool {
	rv := reflect.ValueOf(v)
	return rv.Kind() == reflect.Bool
}

// IsComplex checks if a value is a complex type
func IsComplex(v interface{}) bool {
	rv := reflect.ValueOf(v)
	switch rv.Kind() {
	case reflect.Complex64, reflect.Complex128:
		return true
	default:
		return false
	}
}

// GetTypeName returns the type name of a value
func GetTypeName(v interface{}) string {
	if v == nil {
		return "nil"
	}
	rt := reflect.TypeOf(v)
	return rt.String()
}

// GetKind returns the kind of a value
func GetKind(v interface{}) reflect.Kind {
	if v == nil {
		return reflect.Invalid
	}
	rv := reflect.ValueOf(v)
	return rv.Kind()
}

// GetPackagePath returns the package path of a type
func GetPackagePath(v interface{}) string {
	if v == nil {
		return ""
	}
	rt := reflect.TypeOf(v)
	return rt.PkgPath()
}

// GetName returns the name of a type
func GetName(v interface{}) string {
	if v == nil {
		return ""
	}
	rt := reflect.TypeOf(v)
	return rt.Name()
}

// GetSize returns the size of a type in bytes
func GetSize(v interface{}) uintptr {
	if v == nil {
		return 0
	}
	rt := reflect.TypeOf(v)
	return rt.Size()
}

// GetAlign returns the alignment of a type
func GetAlign(v interface{}) uintptr {
	if v == nil {
		return 0
	}
	rt := reflect.TypeOf(v)
	return uintptr(rt.Align())
}

// GetFieldAlign returns the field alignment of a type
func GetFieldAlign(v interface{}) uintptr {
	if v == nil {
		return 0
	}
	rt := reflect.TypeOf(v)
	return uintptr(rt.FieldAlign())
}

// IsComparable checks if a type is comparable
func IsComparable(v interface{}) bool {
	if v == nil {
		return false
	}
	rt := reflect.TypeOf(v)
	return rt.Comparable()
}

// IsAssignable checks if a type is assignable to another type
func IsAssignable(src, dst interface{}) bool {
	if src == nil || dst == nil {
		return false
	}
	srcType := reflect.TypeOf(src)
	dstType := reflect.TypeOf(dst)
	return srcType.AssignableTo(dstType)
}

// IsConvertible checks if a type is convertible to another type
func IsConvertible(src, dst interface{}) bool {
	if src == nil || dst == nil {
		return false
	}
	srcType := reflect.TypeOf(src)
	dstType := reflect.TypeOf(dst)
	return srcType.ConvertibleTo(dstType)
}

// Convert converts a value to another type
func Convert(src interface{}, dstType reflect.Type) (interface{}, error) {
	if src == nil {
		return nil, fmt.Errorf("source value is nil")
	}
	if dstType == nil {
		return nil, fmt.Errorf("destination type is nil")
	}

	srcValue := reflect.ValueOf(src)
	if !srcValue.IsValid() {
		return nil, fmt.Errorf("source value is invalid")
	}

	if !srcValue.Type().ConvertibleTo(dstType) {
		return nil, fmt.Errorf("cannot convert %s to %s", srcValue.Type(), dstType)
	}

	return srcValue.Convert(dstType).Interface(), nil
}

// ConvertTo converts a value to a specific type
func ConvertTo(src interface{}, dst interface{}) (interface{}, error) {
	if src == nil {
		return nil, fmt.Errorf("source value is nil")
	}
	if dst == nil {
		return nil, fmt.Errorf("destination type is nil")
	}

	dstType := reflect.TypeOf(dst)
	return Convert(src, dstType)
}

// Struct field utilities

// GetFieldNames returns all field names of a struct
func GetFieldNames(v interface{}) []string {
	if v == nil {
		return nil
	}

	rt := reflect.TypeOf(v)

	// Handle pointer to struct
	if rt.Kind() == reflect.Ptr {
		rt = rt.Elem()
	}

	if rt.Kind() != reflect.Struct {
		return nil
	}

	var names []string
	for i := 0; i < rt.NumField(); i++ {
		field := rt.Field(i)
		names = append(names, field.Name)
	}
	return names
}

// GetFieldTags returns all field tags of a struct
func GetFieldTags(v interface{}) map[string]string {
	if v == nil {
		return nil
	}

	rt := reflect.TypeOf(v)

	// Handle pointer to struct
	if rt.Kind() == reflect.Ptr {
		rt = rt.Elem()
	}

	if rt.Kind() != reflect.Struct {
		return nil
	}

	tags := make(map[string]string)
	for i := 0; i < rt.NumField(); i++ {
		field := rt.Field(i)
		tags[field.Name] = string(field.Tag)
	}
	return tags
}

// GetFieldValue returns the value of a field by name
func GetFieldValue(v interface{}, fieldName string) (interface{}, error) {
	if v == nil {
		return nil, fmt.Errorf("value is nil")
	}

	rv := reflect.ValueOf(v)
	rt := reflect.TypeOf(v)

	// Handle pointer to struct
	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
		rt = rt.Elem()
	}

	if rt.Kind() != reflect.Struct {
		return nil, fmt.Errorf("value is not a struct")
	}

	_, found := rt.FieldByName(fieldName)
	if !found {
		return nil, fmt.Errorf("field %s not found", fieldName)
	}

	fieldValue := rv.FieldByName(fieldName)
	if !fieldValue.IsValid() {
		return nil, fmt.Errorf("field %s is invalid", fieldName)
	}

	return fieldValue.Interface(), nil
}

// SetFieldValue sets the value of a field by name
func SetFieldValue(v interface{}, fieldName string, value interface{}) error {
	if v == nil {
		return fmt.Errorf("value is nil")
	}

	rv := reflect.ValueOf(v)
	rt := reflect.TypeOf(v)

	// Handle pointer to struct
	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
		rt = rt.Elem()
	}

	if rt.Kind() != reflect.Struct {
		return fmt.Errorf("value is not a struct")
	}

	_, found := rt.FieldByName(fieldName)
	if !found {
		return fmt.Errorf("field %s not found", fieldName)
	}

	fieldValue := rv.FieldByName(fieldName)
	if !fieldValue.IsValid() {
		return fmt.Errorf("field %s is invalid", fieldName)
	}

	if !fieldValue.CanSet() {
		return fmt.Errorf("field %s cannot be set", fieldName)
	}

	valueType := reflect.TypeOf(value)
	if !valueType.AssignableTo(fieldValue.Type()) {
		return fmt.Errorf("cannot assign %s to field %s of type %s", valueType, fieldName, fieldValue.Type())
	}

	fieldValue.Set(reflect.ValueOf(value))
	return nil
}

// GetFieldType returns the type of a field by name
func GetFieldType(v interface{}, fieldName string) (reflect.Type, error) {
	if v == nil {
		return nil, fmt.Errorf("value is nil")
	}

	rt := reflect.TypeOf(v)

	// Handle pointer to struct
	if rt.Kind() == reflect.Ptr {
		rt = rt.Elem()
	}

	if rt.Kind() != reflect.Struct {
		return nil, fmt.Errorf("value is not a struct")
	}

	field, found := rt.FieldByName(fieldName)
	if !found {
		return nil, fmt.Errorf("field %s not found", fieldName)
	}

	return field.Type, nil
}

// GetFieldTag returns the tag of a field by name
func GetFieldTag(v interface{}, fieldName, tagKey string) (string, error) {
	if v == nil {
		return "", fmt.Errorf("value is nil")
	}

	rt := reflect.TypeOf(v)

	// Handle pointer to struct
	if rt.Kind() == reflect.Ptr {
		rt = rt.Elem()
	}

	if rt.Kind() != reflect.Struct {
		return "", fmt.Errorf("value is not a struct")
	}

	field, found := rt.FieldByName(fieldName)
	if !found {
		return "", fmt.Errorf("field %s not found", fieldName)
	}

	return field.Tag.Get(tagKey), nil
}

// HasField checks if a struct has a field with the given name
func HasField(v interface{}, fieldName string) bool {
	if v == nil {
		return false
	}

	rt := reflect.TypeOf(v)

	// Handle pointer to struct
	if rt.Kind() == reflect.Ptr {
		rt = rt.Elem()
	}

	if rt.Kind() != reflect.Struct {
		return false
	}

	_, found := rt.FieldByName(fieldName)
	return found
}

// GetFieldCount returns the number of fields in a struct
func GetFieldCount(v interface{}) int {
	if v == nil {
		return 0
	}

	rt := reflect.TypeOf(v)

	// Handle pointer to struct
	if rt.Kind() == reflect.Ptr {
		rt = rt.Elem()
	}

	if rt.Kind() != reflect.Struct {
		return 0
	}

	return rt.NumField()
}

// GetFieldInfo returns detailed information about a field
func GetFieldInfo(v interface{}, fieldName string) (map[string]interface{}, error) {
	if v == nil {
		return nil, fmt.Errorf("value is nil")
	}

	rt := reflect.TypeOf(v)

	// Handle pointer to struct
	if rt.Kind() == reflect.Ptr {
		rt = rt.Elem()
	}

	if rt.Kind() != reflect.Struct {
		return nil, fmt.Errorf("value is not a struct")
	}

	field, found := rt.FieldByName(fieldName)
	if !found {
		return nil, fmt.Errorf("field %s not found", fieldName)
	}

	info := map[string]interface{}{
		"name":      field.Name,
		"type":      field.Type.String(),
		"tag":       string(field.Tag),
		"index":     field.Index,
		"anonymous": field.Anonymous,
		"pkgPath":   field.PkgPath,
	}

	return info, nil
}

// GetAnonymousFields returns all anonymous fields of a struct
func GetAnonymousFields(v interface{}) []string {
	if v == nil {
		return nil
	}

	rt := reflect.TypeOf(v)

	// Handle pointer to struct
	if rt.Kind() == reflect.Ptr {
		rt = rt.Elem()
	}

	if rt.Kind() != reflect.Struct {
		return nil
	}

	var fields []string
	for i := 0; i < rt.NumField(); i++ {
		field := rt.Field(i)
		if field.Anonymous {
			fields = append(fields, field.Name)
		}
	}
	return fields
}

// GetExportedFields returns all exported fields of a struct
func GetExportedFields(v interface{}) []string {
	if v == nil {
		return nil
	}

	rt := reflect.TypeOf(v)

	// Handle pointer to struct
	if rt.Kind() == reflect.Ptr {
		rt = rt.Elem()
	}

	if rt.Kind() != reflect.Struct {
		return nil
	}

	var fields []string
	for i := 0; i < rt.NumField(); i++ {
		field := rt.Field(i)
		if field.PkgPath == "" { // Exported field
			fields = append(fields, field.Name)
		}
	}
	return fields
}

// GetUnexportedFields returns all unexported fields of a struct
func GetUnexportedFields(v interface{}) []string {
	if v == nil {
		return nil
	}

	rt := reflect.TypeOf(v)

	// Handle pointer to struct
	if rt.Kind() == reflect.Ptr {
		rt = rt.Elem()
	}

	if rt.Kind() != reflect.Struct {
		return nil
	}

	var fields []string
	for i := 0; i < rt.NumField(); i++ {
		field := rt.Field(i)
		if field.PkgPath != "" { // Unexported field
			fields = append(fields, field.Name)
		}
	}
	return fields
}

// Value operation and conversion utilities

// GetValue returns the reflect.Value of a value
func GetValue(v interface{}) reflect.Value {
	return reflect.ValueOf(v)
}

// GetType returns the reflect.Type of a value
func GetType(v interface{}) reflect.Type {
	return reflect.TypeOf(v)
}

// GetElem returns the element type of a pointer, slice, array, map, or channel
func GetElem(v interface{}) (reflect.Type, error) {
	if v == nil {
		return nil, fmt.Errorf("value is nil")
	}

	rt := reflect.TypeOf(v)
	switch rt.Kind() {
	case reflect.Ptr, reflect.Slice, reflect.Array, reflect.Map, reflect.Chan:
		return rt.Elem(), nil
	default:
		return nil, fmt.Errorf("type %s does not have an element type", rt)
	}
}

// GetKey returns the key type of a map
func GetKey(v interface{}) (reflect.Type, error) {
	if v == nil {
		return nil, fmt.Errorf("value is nil")
	}

	rt := reflect.TypeOf(v)
	if rt.Kind() != reflect.Map {
		return nil, fmt.Errorf("type %s is not a map", rt)
	}

	return rt.Key(), nil
}

// GetLen returns the length of a slice, array, map, string, or channel
func GetLen(v interface{}) (int, error) {
	if v == nil {
		return 0, fmt.Errorf("value is nil")
	}

	rv := reflect.ValueOf(v)
	switch rv.Kind() {
	case reflect.Slice, reflect.Array, reflect.Map, reflect.String, reflect.Chan:
		return rv.Len(), nil
	default:
		return 0, fmt.Errorf("type %s does not have a length", rv.Type())
	}
}

// GetCap returns the capacity of a slice, array, or channel
func GetCap(v interface{}) (int, error) {
	if v == nil {
		return 0, fmt.Errorf("value is nil")
	}

	rv := reflect.ValueOf(v)
	switch rv.Kind() {
	case reflect.Slice, reflect.Array, reflect.Chan:
		return rv.Cap(), nil
	default:
		return 0, fmt.Errorf("type %s does not have a capacity", rv.Type())
	}
}

// GetIndex returns the element at index i of a slice, array, or string
func GetIndex(v interface{}, i int) (interface{}, error) {
	if v == nil {
		return nil, fmt.Errorf("value is nil")
	}

	rv := reflect.ValueOf(v)
	switch rv.Kind() {
	case reflect.Slice, reflect.Array, reflect.String:
		if i < 0 || i >= rv.Len() {
			return nil, fmt.Errorf("index %d out of range", i)
		}
		return rv.Index(i).Interface(), nil
	default:
		return nil, fmt.Errorf("type %s does not support indexing", rv.Type())
	}
}

// SetIndex sets the element at index i of a slice or array
func SetIndex(v interface{}, i int, value interface{}) error {
	if v == nil {
		return fmt.Errorf("value is nil")
	}

	rv := reflect.ValueOf(v)
	switch rv.Kind() {
	case reflect.Slice, reflect.Array:
		if i < 0 || i >= rv.Len() {
			return fmt.Errorf("index %d out of range", i)
		}
		if !rv.Index(i).CanSet() {
			return fmt.Errorf("element at index %d cannot be set", i)
		}
		valueType := reflect.TypeOf(value)
		if !valueType.AssignableTo(rv.Index(i).Type()) {
			return fmt.Errorf("cannot assign %s to element of type %s", valueType, rv.Index(i).Type())
		}
		rv.Index(i).Set(reflect.ValueOf(value))
		return nil
	default:
		return fmt.Errorf("type %s does not support setting by index", rv.Type())
	}
}

// GetMapValue returns the value for a key in a map
func GetMapValue(v interface{}, key interface{}) (interface{}, error) {
	if v == nil {
		return nil, fmt.Errorf("value is nil")
	}

	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Map {
		return nil, fmt.Errorf("type %s is not a map", rv.Type())
	}

	keyValue := reflect.ValueOf(key)
	if !keyValue.Type().AssignableTo(rv.Type().Key()) {
		return nil, fmt.Errorf("key type %s is not assignable to map key type %s", keyValue.Type(), rv.Type().Key())
	}

	mapValue := rv.MapIndex(keyValue)
	if !mapValue.IsValid() {
		return nil, fmt.Errorf("key not found in map")
	}

	return mapValue.Interface(), nil
}

// SetMapValue sets the value for a key in a map
func SetMapValue(v interface{}, key, value interface{}) error {
	if v == nil {
		return fmt.Errorf("value is nil")
	}

	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Map {
		return fmt.Errorf("type %s is not a map", rv.Type())
	}

	keyValue := reflect.ValueOf(key)
	valueValue := reflect.ValueOf(value)

	if !keyValue.Type().AssignableTo(rv.Type().Key()) {
		return fmt.Errorf("key type %s is not assignable to map key type %s", keyValue.Type(), rv.Type().Key())
	}

	if !valueValue.Type().AssignableTo(rv.Type().Elem()) {
		return fmt.Errorf("value type %s is not assignable to map value type %s", valueValue.Type(), rv.Type().Elem())
	}

	rv.SetMapIndex(keyValue, valueValue)
	return nil
}

// GetMapKeys returns all keys in a map
func GetMapKeys(v interface{}) ([]interface{}, error) {
	if v == nil {
		return nil, fmt.Errorf("value is nil")
	}

	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Map {
		return nil, fmt.Errorf("type %s is not a map", rv.Type())
	}

	var keys []interface{}
	for _, key := range rv.MapKeys() {
		keys = append(keys, key.Interface())
	}

	return keys, nil
}

// GetMapValues returns all values in a map
func GetMapValues(v interface{}) ([]interface{}, error) {
	if v == nil {
		return nil, fmt.Errorf("value is nil")
	}

	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Map {
		return nil, fmt.Errorf("type %s is not a map", rv.Type())
	}

	var values []interface{}
	for _, key := range rv.MapKeys() {
		values = append(values, rv.MapIndex(key).Interface())
	}

	return values, nil
}

// GetMapEntries returns all key-value pairs in a map
func GetMapEntries(v interface{}) ([]map[string]interface{}, error) {
	if v == nil {
		return nil, fmt.Errorf("value is nil")
	}

	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Map {
		return nil, fmt.Errorf("type %s is not a map", rv.Type())
	}

	var entries []map[string]interface{}
	for _, key := range rv.MapKeys() {
		entry := map[string]interface{}{
			"key":   key.Interface(),
			"value": rv.MapIndex(key).Interface(),
		}
		entries = append(entries, entry)
	}

	return entries, nil
}

// CallMethod calls a method by name on a value
func CallMethod(v interface{}, methodName string, args ...interface{}) ([]interface{}, error) {
	if v == nil {
		return nil, fmt.Errorf("value is nil")
	}

	rv := reflect.ValueOf(v)
	rt := reflect.TypeOf(v)

	// Handle pointer to struct
	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
		rt = rt.Elem()
	}

	_, found := rt.MethodByName(methodName)
	if !found {
		return nil, fmt.Errorf("method %s not found", methodName)
	}

	methodValue := rv.MethodByName(methodName)
	if !methodValue.IsValid() {
		return nil, fmt.Errorf("method %s is invalid", methodName)
	}

	// Convert args to reflect.Values
	var reflectArgs []reflect.Value
	for _, arg := range args {
		reflectArgs = append(reflectArgs, reflect.ValueOf(arg))
	}

	// Call the method
	results := methodValue.Call(reflectArgs)

	// Convert results back to interfaces
	var interfaces []interface{}
	for _, result := range results {
		interfaces = append(interfaces, result.Interface())
	}

	return interfaces, nil
}

// HasMethod checks if a value has a method with the given name
func HasMethod(v interface{}, methodName string) bool {
	if v == nil {
		return false
	}

	rt := reflect.TypeOf(v)

	// Handle pointer to struct
	if rt.Kind() == reflect.Ptr {
		rt = rt.Elem()
	}

	_, found := rt.MethodByName(methodName)
	return found
}

// GetMethodNames returns all method names of a value
func GetMethodNames(v interface{}) []string {
	if v == nil {
		return nil
	}

	rt := reflect.TypeOf(v)

	// Handle pointer to struct
	if rt.Kind() == reflect.Ptr {
		rt = rt.Elem()
	}

	var names []string
	for i := 0; i < rt.NumMethod(); i++ {
		method := rt.Method(i)
		names = append(names, method.Name)
	}

	return names
}

// GetMethodCount returns the number of methods of a value
func GetMethodCount(v interface{}) int {
	if v == nil {
		return 0
	}

	rt := reflect.TypeOf(v)

	// Handle pointer to struct
	if rt.Kind() == reflect.Ptr {
		rt = rt.Elem()
	}

	return rt.NumMethod()
}

// GetMethodInfo returns detailed information about a method
func GetMethodInfo(v interface{}, methodName string) (map[string]interface{}, error) {
	if v == nil {
		return nil, fmt.Errorf("value is nil")
	}

	rt := reflect.TypeOf(v)

	// Handle pointer to struct
	if rt.Kind() == reflect.Ptr {
		rt = rt.Elem()
	}

	method, found := rt.MethodByName(methodName)
	if !found {
		return nil, fmt.Errorf("method %s not found", methodName)
	}

	info := map[string]interface{}{
		"name":    method.Name,
		"type":    method.Type.String(),
		"index":   method.Index,
		"pkgPath": method.PkgPath,
		"numIn":   method.Type.NumIn(),
		"numOut":  method.Type.NumOut(),
	}

	return info, nil
}

// Interface utilities

// Implements checks if a type implements an interface
func Implements(typ interface{}, iface interface{}) bool {
	if typ == nil || iface == nil {
		return false
	}

	typType := reflect.TypeOf(typ)
	ifaceType := reflect.TypeOf(iface)

	// Handle pointer to interface
	if ifaceType.Kind() == reflect.Ptr {
		ifaceType = ifaceType.Elem()
	}

	if ifaceType.Kind() != reflect.Interface {
		return false
	}

	return typType.Implements(ifaceType)
}

// AssignableTo checks if a type is assignable to another type
func AssignableTo(src, dst interface{}) bool {
	if src == nil || dst == nil {
		return false
	}

	srcType := reflect.TypeOf(src)
	dstType := reflect.TypeOf(dst)

	return srcType.AssignableTo(dstType)
}

// ConvertibleTo checks if a type is convertible to another type
func ConvertibleTo(src, dst interface{}) bool {
	if src == nil || dst == nil {
		return false
	}

	srcType := reflect.TypeOf(src)
	dstType := reflect.TypeOf(dst)

	return srcType.ConvertibleTo(dstType)
}

// GetInterfaceMethods returns all methods of an interface
func GetInterfaceMethods(iface interface{}) []string {
	if iface == nil {
		return nil
	}

	ifaceType := reflect.TypeOf(iface)

	// Handle pointer to interface
	if ifaceType.Kind() == reflect.Ptr {
		ifaceType = ifaceType.Elem()
	}

	if ifaceType.Kind() != reflect.Interface {
		return nil
	}

	var methods []string
	for i := 0; i < ifaceType.NumMethod(); i++ {
		method := ifaceType.Method(i)
		methods = append(methods, method.Name)
	}

	return methods
}

// GetInterfaceMethodCount returns the number of methods in an interface
func GetInterfaceMethodCount(iface interface{}) int {
	if iface == nil {
		return 0
	}

	ifaceType := reflect.TypeOf(iface)

	// Handle pointer to interface
	if ifaceType.Kind() == reflect.Ptr {
		ifaceType = ifaceType.Elem()
	}

	if ifaceType.Kind() != reflect.Interface {
		return 0
	}

	return ifaceType.NumMethod()
}

// GetInterfaceMethodInfo returns detailed information about an interface method
func GetInterfaceMethodInfo(iface interface{}, methodName string) (map[string]interface{}, error) {
	if iface == nil {
		return nil, fmt.Errorf("interface is nil")
	}

	ifaceType := reflect.TypeOf(iface)

	// Handle pointer to interface
	if ifaceType.Kind() == reflect.Ptr {
		ifaceType = ifaceType.Elem()
	}

	if ifaceType.Kind() != reflect.Interface {
		return nil, fmt.Errorf("type is not an interface")
	}

	method, found := ifaceType.MethodByName(methodName)
	if !found {
		return nil, fmt.Errorf("method %s not found", methodName)
	}

	info := map[string]interface{}{
		"name":    method.Name,
		"type":    method.Type.String(),
		"index":   method.Index,
		"pkgPath": method.PkgPath,
		"numIn":   method.Type.NumIn(),
		"numOut":  method.Type.NumOut(),
	}

	return info, nil
}

// IsInterfaceType checks if a type is an interface
func IsInterfaceType(typ interface{}) bool {
	if typ == nil {
		return false
	}

	typType := reflect.TypeOf(typ)

	// Handle pointer to interface
	if typType.Kind() == reflect.Ptr {
		typType = typType.Elem()
	}

	return typType.Kind() == reflect.Interface
}

// GetInterfaceName returns the name of an interface
func GetInterfaceName(iface interface{}) string {
	if iface == nil {
		return ""
	}

	ifaceType := reflect.TypeOf(iface)

	// Handle pointer to interface
	if ifaceType.Kind() == reflect.Ptr {
		ifaceType = ifaceType.Elem()
	}

	if ifaceType.Kind() != reflect.Interface {
		return ""
	}

	return ifaceType.Name()
}

// GetInterfacePackage returns the package path of an interface
func GetInterfacePackage(iface interface{}) string {
	if iface == nil {
		return ""
	}

	ifaceType := reflect.TypeOf(iface)

	// Handle pointer to interface
	if ifaceType.Kind() == reflect.Ptr {
		ifaceType = ifaceType.Elem()
	}

	if ifaceType.Kind() != reflect.Interface {
		return ""
	}

	return ifaceType.PkgPath()
}

// GetInterfaceString returns the string representation of an interface
func GetInterfaceString(iface interface{}) string {
	if iface == nil {
		return "nil"
	}

	ifaceType := reflect.TypeOf(iface)

	// Handle pointer to interface
	if ifaceType.Kind() == reflect.Ptr {
		ifaceType = ifaceType.Elem()
	}

	if ifaceType.Kind() != reflect.Interface {
		return ""
	}

	return ifaceType.String()
}

// GetInterfaceMethodsDetailed returns all methods of an interface with detailed information
func GetInterfaceMethodsDetailed(iface interface{}) ([]map[string]interface{}, error) {
	if iface == nil {
		return nil, fmt.Errorf("interface is nil")
	}

	ifaceType := reflect.TypeOf(iface)

	// Handle pointer to interface
	if ifaceType.Kind() == reflect.Ptr {
		ifaceType = ifaceType.Elem()
	}

	if ifaceType.Kind() != reflect.Interface {
		return nil, fmt.Errorf("type is not an interface")
	}

	var methods []map[string]interface{}
	for i := 0; i < ifaceType.NumMethod(); i++ {
		method := ifaceType.Method(i)
		info := map[string]interface{}{
			"name":    method.Name,
			"type":    method.Type.String(),
			"index":   method.Index,
			"pkgPath": method.PkgPath,
			"numIn":   method.Type.NumIn(),
			"numOut":  method.Type.NumOut(),
		}
		methods = append(methods, info)
	}

	return methods, nil
}

// Generic utilities

// IsGeneric checks if a type is generic
func IsGeneric(typ interface{}) bool {
	if typ == nil {
		return false
	}

	typType := reflect.TypeOf(typ)
	return typType.Kind() == reflect.Interface && typType.NumMethod() > 0
}

// GetGenericTypeParameters returns the type parameters of a generic type
func GetGenericTypeParameters(typ interface{}) ([]reflect.Type, error) {
	if typ == nil {
		return nil, fmt.Errorf("type is nil")
	}

	typType := reflect.TypeOf(typ)
	if typType.Kind() != reflect.Interface {
		return nil, fmt.Errorf("type is not an interface")
	}

	// Note: Go's reflection doesn't provide direct access to type parameters
	// This is a limitation of the current reflection API
	// We can only check if the type has methods (indicating it might be generic)
	if typType.NumMethod() == 0 {
		return nil, fmt.Errorf("type has no methods")
	}

	// Return empty slice as we can't access type parameters directly
	return []reflect.Type{}, nil
}

// GetGenericConstraints returns the constraints of a generic type
func GetGenericConstraints(typ interface{}) ([]string, error) {
	if typ == nil {
		return nil, fmt.Errorf("type is nil")
	}

	typType := reflect.TypeOf(typ)
	if typType.Kind() != reflect.Interface {
		return nil, fmt.Errorf("type is not an interface")
	}

	// Note: Go's reflection doesn't provide direct access to type constraints
	// This is a limitation of the current reflection API
	return []string{}, nil
}

// IsGenericType checks if a type is a generic type
func IsGenericType(typ interface{}) bool {
	if typ == nil {
		return false
	}

	typType := reflect.TypeOf(typ)

	// Check if it's an interface with methods (might be generic)
	if typType.Kind() == reflect.Interface {
		return typType.NumMethod() > 0
	}

	return false
}

// GetGenericTypeName returns the name of a generic type
func GetGenericTypeName(typ interface{}) string {
	if typ == nil {
		return ""
	}

	typType := reflect.TypeOf(typ)
	if typType.Kind() != reflect.Interface {
		return ""
	}

	return typType.Name()
}

// GetGenericTypeString returns the string representation of a generic type
func GetGenericTypeString(typ interface{}) string {
	if typ == nil {
		return "nil"
	}

	typType := reflect.TypeOf(typ)
	return typType.String()
}

// GetGenericTypePackage returns the package path of a generic type
func GetGenericTypePackage(typ interface{}) string {
	if typ == nil {
		return ""
	}

	typType := reflect.TypeOf(typ)
	if typType.Kind() != reflect.Interface {
		return ""
	}

	return typType.PkgPath()
}

// GetGenericTypeSize returns the size of a generic type
func GetGenericTypeSize(typ interface{}) uintptr {
	if typ == nil {
		return 0
	}

	typType := reflect.TypeOf(typ)
	return typType.Size()
}

// GetGenericTypeAlign returns the alignment of a generic type
func GetGenericTypeAlign(typ interface{}) uintptr {
	if typ == nil {
		return 0
	}

	typType := reflect.TypeOf(typ)
	return uintptr(typType.Align())
}

// GetGenericTypeFieldAlign returns the field alignment of a generic type
func GetGenericTypeFieldAlign(typ interface{}) uintptr {
	if typ == nil {
		return 0
	}

	typType := reflect.TypeOf(typ)
	return uintptr(typType.FieldAlign())
}

// GetGenericTypeComparable checks if a generic type is comparable
func GetGenericTypeComparable(typ interface{}) bool {
	if typ == nil {
		return false
	}

	typType := reflect.TypeOf(typ)
	return typType.Comparable()
}

// GetGenericTypeAssignableTo checks if a generic type is assignable to another type
func GetGenericTypeAssignableTo(src, dst interface{}) bool {
	if src == nil || dst == nil {
		return false
	}

	srcType := reflect.TypeOf(src)
	dstType := reflect.TypeOf(dst)

	return srcType.AssignableTo(dstType)
}

// GetGenericTypeConvertibleTo checks if a generic type is convertible to another type
func GetGenericTypeConvertibleTo(src, dst interface{}) bool {
	if src == nil || dst == nil {
		return false
	}

	srcType := reflect.TypeOf(src)
	dstType := reflect.TypeOf(dst)

	return srcType.ConvertibleTo(dstType)
}

// GetGenericTypeMethods returns all methods of a generic type
func GetGenericTypeMethods(typ interface{}) []string {
	if typ == nil {
		return nil
	}

	typType := reflect.TypeOf(typ)
	if typType.Kind() != reflect.Interface {
		return nil
	}

	var methods []string
	for i := 0; i < typType.NumMethod(); i++ {
		method := typType.Method(i)
		methods = append(methods, method.Name)
	}

	return methods
}

// GetGenericTypeMethodCount returns the number of methods of a generic type
func GetGenericTypeMethodCount(typ interface{}) int {
	if typ == nil {
		return 0
	}

	typType := reflect.TypeOf(typ)
	if typType.Kind() != reflect.Interface {
		return 0
	}

	return typType.NumMethod()
}

// GetGenericTypeMethodInfo returns detailed information about a method of a generic type
func GetGenericTypeMethodInfo(typ interface{}, methodName string) (map[string]interface{}, error) {
	if typ == nil {
		return nil, fmt.Errorf("type is nil")
	}

	typType := reflect.TypeOf(typ)
	if typType.Kind() != reflect.Interface {
		return nil, fmt.Errorf("type is not an interface")
	}

	method, found := typType.MethodByName(methodName)
	if !found {
		return nil, fmt.Errorf("method %s not found", methodName)
	}

	info := map[string]interface{}{
		"name":    method.Name,
		"type":    method.Type.String(),
		"index":   method.Index,
		"pkgPath": method.PkgPath,
		"numIn":   method.Type.NumIn(),
		"numOut":  method.Type.NumOut(),
	}

	return info, nil
}

// GetGenericTypeMethodsDetailed returns all methods of a generic type with detailed information
func GetGenericTypeMethodsDetailed(typ interface{}) ([]map[string]interface{}, error) {
	if typ == nil {
		return nil, fmt.Errorf("type is nil")
	}

	typType := reflect.TypeOf(typ)
	if typType.Kind() != reflect.Interface {
		return nil, fmt.Errorf("type is not an interface")
	}

	var methods []map[string]interface{}
	for i := 0; i < typType.NumMethod(); i++ {
		method := typType.Method(i)
		info := map[string]interface{}{
			"name":    method.Name,
			"type":    method.Type.String(),
			"index":   method.Index,
			"pkgPath": method.PkgPath,
			"numIn":   method.Type.NumIn(),
			"numOut":  method.Type.NumOut(),
		}
		methods = append(methods, info)
	}

	return methods, nil
}
