// Copyright 2012 The go-gl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gl

// #cgo darwin LDFLAGS: -framework OpenGL -lGLEW
// #cgo windows LDFLAGS: -lglew32 -lopengl32
// #cgo linux LDFLAGS: -lGLEW -lGL
// #cgo freebsd  CFLAGS: -I/usr/local/include
// #cgo freebsd LDFLAGS: -L/usr/local/lib -lglfw
// #include "gl.h"
//
// const char *gogl_glewGetErrorString(GLenum name) {
//     return glewGetErrorString(name);
// }
import "C"
import "reflect"
import "unsafe"
import "fmt"

type GLbitfield C.GLbitfield
type GLboolean C.GLboolean
type GLbyte C.GLbyte
type GLclampd C.GLclampd
type GLclampf C.GLclampf
type GLdouble C.GLdouble
type GLenum C.GLenum
type GLfloat C.GLfloat
type GLint C.GLint
type GLshort C.GLshort
type GLsizei C.GLsizei
type GLsync unsafe.Pointer
type GLubyte C.GLubyte
type GLuint C.GLuint
type GLushort C.GLushort

type AttribLocation C.GLint
type UniformBlockIndex C.GLuint
type UniformLocation C.GLint

type Object C.GLuint
type Buffer Object
type List Object
type Program Object
type Query Object
type Shader Object
type Texture Object
type TransformFeedback Object
type VertexArray Object

var extensions = map[string]*bool{}

func glBool(v bool) C.GLboolean {
	if v {
		return TRUE
	}
	return FALSE
}

func goBool(v C.GLboolean) bool {
	return v == TRUE
}

func glString(s string) *C.GLchar {
	return (*C.GLchar)(C.CString(s))
}

func freeString(ptr *C.GLchar) {
	C.free(unsafe.Pointer(ptr))
}

func glPointer(v interface{}) unsafe.Pointer {
	if v == nil {
		return unsafe.Pointer(nil)
	}

	rv := reflect.ValueOf(v)
	var et reflect.Value
	switch rv.Type().Kind() {
	case reflect.Uintptr:
		offset, _ := v.(uintptr)
		return unsafe.Pointer(offset)
	case reflect.Ptr:
		et = rv.Elem()
	case reflect.Slice:
		et = rv.Index(0)
	default:
		panic("type must be a pointer, a slice, uintptr or nil")
	}

	return unsafe.Pointer(et.UnsafeAddr())
}

func Init() error {
	C.glewExperimental = glBool(true)

	err := C.glewInit()
	if err != C.GLEW_OK {
		errStr := C.GoString(C.gogl_glewGetErrorString(err))
		return fmt.Errorf(errStr)
	}

	for name, val := range extensions {
		*val = IsSupported(name)
	}
	extensions = nil

	return nil
}

func IsSupported(name string) bool {
	return goBool(C.glewIsSupported(C.CString(name)))
}

// Creates a slice from a raw buffer address. The pointer returned points to a slice.
// Convert it like this: *(*[]TYPE)(gl.Slice(buf, len))
// Length is the total amount of elements in the slice.
// WARNING: This function makes use of reflect.SliceHeader which may reduce
// portability of your application. See the reflect.SliceHeader documentation
// for more information.
func Slice(buffer unsafe.Pointer, length int) unsafe.Pointer {
	return unsafe.Pointer(&reflect.SliceHeader{
		Data: uintptr(buffer),
		Len:  length,
		Cap:  length,
	})
}

// Convenience function for Slice converting the raw slice to []int8
func SliceInt8(buffer unsafe.Pointer, length int) []int8 {
	return *(*[]int8)(Slice(buffer, length))
}

// Convenience function for Slice converting the raw slice to []uint8
func SliceUint8(buffer unsafe.Pointer, length int) []uint8 {
	return *(*[]uint8)(Slice(buffer, length))
}

// Convenience function for Slice converting the raw slice to []int32
func SliceInt32(buffer unsafe.Pointer, length int) []int32 {
	return *(*[]int32)(Slice(buffer, length))
}

// Convenience function for Slice converting the raw slice to []uint32
func SliceUint32(buffer unsafe.Pointer, length int) []uint32 {
	return *(*[]uint32)(Slice(buffer, length))
}

// Convenience function for Slice converting the raw slice to []float32
func SliceFloat32(buffer unsafe.Pointer, length int) []float32 {
	return *(*[]float32)(Slice(buffer, length))
}

// Convenience function for Slice converting the raw slice to []float64
func SliceFloat64(buffer unsafe.Pointer, length int) []float64 {
	return *(*[]float64)(Slice(buffer, length))
}
