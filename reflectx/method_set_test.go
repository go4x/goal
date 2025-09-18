package reflectx_test

import (
	"fmt"
	"testing"

	"github.com/gophero/goal/reflectx"
	"github.com/stretchr/testify/assert"
)

type Inter interface {
	M1()
	M2()
}

type InterImpl struct {
}

// 非指针接收者
func (r InterImpl) M1() {

}

// 指针接收者
func (r *InterImpl) M2() {

}

func TestMethods(t *testing.T) {
	var ty InterImpl
	var pty *InterImpl
	methods := reflectx.Methods(&ty)
	fmt.Println(methods)
	assert.Equal(t, len(methods), 1)
	assert.Equal(t, methods[0], "M1")
	methods = reflectx.Methods(&pty)
	fmt.Println(methods)
	assert.Equal(t, len(methods), 2)
	assert.Equal(t, methods[0], "M1")
	assert.Equal(t, methods[1], "M2")
	// nil interface, also has 2 methods: M1 and M2
	methods = reflectx.Methods((*Inter)(nil))
	fmt.Println(methods)
	assert.Equal(t, len(methods), 2)
	assert.Equal(t, methods[0], "M1")
	assert.Equal(t, methods[1], "M2")
}

func TestPrintMethodSet(t *testing.T) {
	var ty InterImpl
	var pty *InterImpl
	reflectx.PrintMethodSet(&ty)
	reflectx.PrintMethodSet(&pty)
	// nil interface
	reflectx.PrintMethodSet((*Inter)(nil))
}

type Inter2 interface {
	M3()
}

type GenericInter[T any] interface {
	Inter
	Print(t T)
}

type GenericImpl[T any] struct {
	InterImpl
}

func (r GenericImpl[T]) M3() {
}

func (r *GenericImpl[T]) Print(t T) {
	fmt.Println(t)
}

func TestPrintMethodSet1(t *testing.T) {
	var ty GenericImpl[int]
	var pty *GenericImpl[int]
	reflectx.PrintMethodSet(&ty)
	reflectx.PrintMethodSet(&pty)
	// nil interface
	reflectx.PrintMethodSet((*GenericInter[int])(nil))
}
