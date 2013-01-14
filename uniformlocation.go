// Copyright 2012 The go-gl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gl

// #include "gl.h"
import "C"
import "unsafe"

// UniformLocation
//TODO

type UniformLocation int

func (location UniformLocation) Uniform1f(x float32) {
	C.glUniform1f(C.GLint(location), C.GLfloat(x))
}

func (location UniformLocation) Uniform2f(x float32, y float32) {
	C.glUniform2f(C.GLint(location), C.GLfloat(x), C.GLfloat(y))
}

func (location UniformLocation) Uniform3f(x float32, y float32, z float32) {
	C.glUniform3f(C.GLint(location), C.GLfloat(x), C.GLfloat(y), C.GLfloat(z))
}

func (location UniformLocation) Uniform1ui(x uint) {
	C.glUniform1ui(C.GLint(location), C.GLuint(x))
}
func (location UniformLocation) Uniform2ui(x uint, y uint) {
	C.glUniform2ui(C.GLint(location), C.GLuint(x), C.GLuint(y))
}

func (location UniformLocation) Uniform3ui(x uint, y uint, z uint) {
	C.glUniform3ui(C.GLint(location), C.GLuint(x), C.GLuint(y), C.GLuint(z))
}

func (location UniformLocation) Uniform4ui(x uint, y uint, z uint, w uint) {
	C.glUniform4ui(C.GLint(location), C.GLuint(x), C.GLuint(y), C.GLuint(z), C.GLuint(w))
}

func (location UniformLocation) Uniform1uiv(v []uint32) {
	if len(v) < 1 {
		panic("Invalid array length - must be at least 1")
	}
	C.glUniform1uiv(C.GLint(location), C.GLsizei(len(v)), (*C.GLuint)((unsafe.Pointer)(&v[0])))
}

func (location UniformLocation) Uniform2uiv(v ...[2]uint32) {
	if len(v) < 1 {
		panic("Invalid array length - must be at least 1")
	}
	C.glUniform2uiv(C.GLint(location), C.GLsizei(len(v)), (*C.GLuint)((unsafe.Pointer)(&v[0])))
}

func (location UniformLocation) Uniform3uiv(v ...[3]uint32) {
	if len(v) < 1 {
		panic("Invalid array length - must be at least 1")
	}
	C.glUniform3uiv(C.GLint(location), C.GLsizei(len(v)), (*C.GLuint)((unsafe.Pointer)(&v[0])))
}

func (location UniformLocation) Uniform4uiv(v ...[4]uint32) {
	if len(v) < 1 {
		panic("Invalid array length - must be at least 1")
	}
	C.glUniform4uiv(C.GLint(location), C.GLsizei(len(v)), (*C.GLuint)((unsafe.Pointer)(&v[0])))
}

func (location UniformLocation) Uniform1fv(v []float32) {
	if len(v) < 1 {
		panic("Invalid array length - must be at least 1")
	}
	C.glUniform1fv(C.GLint(location), C.GLsizei(len(v)), (*C.GLfloat)((unsafe.Pointer)(&v[0])))
}

func (location UniformLocation) Uniform1i(x int) {
	C.glUniform1i(C.GLint(location), C.GLint(x))
}

func (location UniformLocation) Uniform1iv(v []int32) {
	if len(v) < 1 {
		panic("Invalid array length - must be at least 1")
	}
	C.glUniform1iv(C.GLint(location), C.GLsizei(len(v)), (*C.GLint)((unsafe.Pointer)(&v[0])))
}

func (location UniformLocation) Uniform2fv(v ...[2]float32) {
	if len(v) < 1 {
		panic("Invalid array length - must be at least 1")
	}
	C.glUniform2fv(C.GLint(location), C.GLsizei(len(v)), (*C.GLfloat)((unsafe.Pointer)(&v[0])))
}

func (location UniformLocation) Uniform2i(x int, y int) {
	C.glUniform2i(C.GLint(location), C.GLint(x), C.GLint(y))
}

func (location UniformLocation) Uniform2iv(v ...[2]int32) {
	if len(v) < 1 {
		panic("Invalid array length - must be at least 1")
	}
	C.glUniform2iv(C.GLint(location), C.GLsizei(len(v)), (*C.GLint)((unsafe.Pointer)(&v[0])))
}

func (location UniformLocation) Uniform3fv(v ...[3]float32) {
	if len(v) < 1 {
		panic("Invalid array length - must be at least 1")
	}
	C.glUniform3fv(C.GLint(location), C.GLsizei(len(v)), (*C.GLfloat)((unsafe.Pointer)(&v[0])))
}

func (location UniformLocation) Uniform3i(x int, y int, z int) {
	C.glUniform3i(C.GLint(location), C.GLint(x), C.GLint(y), C.GLint(z))
}

func (location UniformLocation) Uniform3iv(v ...[3]int32) {
	if len(v) < 1 {
		panic("Invalid array length - must be at least 1")
	}
	C.glUniform3iv(C.GLint(location), C.GLsizei(len(v)), (*C.GLint)((unsafe.Pointer)(&v[0])))
}

func (location UniformLocation) Uniform4f(x float32, y float32, z float32, w float32) {
	C.glUniform4f(C.GLint(location), C.GLfloat(x), C.GLfloat(y), C.GLfloat(z), C.GLfloat(w))
}

func (location UniformLocation) Uniform4fv(v ...[4]float32) {
	if len(v) < 1 {
		panic("Invalid array length - must be at least 1")
	}
	C.glUniform4fv(C.GLint(location), C.GLsizei(len(v)), (*C.GLfloat)((unsafe.Pointer)(&v[0])))
}

func (location UniformLocation) Uniform4i(x int, y int, z int, w int) {
	C.glUniform4i(C.GLint(location), C.GLint(x), C.GLint(y), C.GLint(z), C.GLint(w))
}

func (location UniformLocation) Uniform4iv(v ...[4]int32) {
	if len(v) < 1 {
		panic("Invalid array length - must be at least 1")
	}
	C.glUniform4iv(C.GLint(location), C.GLsizei(len(v)), (*C.GLint)((unsafe.Pointer)(&v[0])))
}

func (location UniformLocation) UniformMatrix2fv(transpose bool, list ...[4]float32) {
	if len(list) < 1 {
		panic("Invalid array length - must be at least 1")
	}
	C.glUniformMatrix2fv(C.GLint(location), C.GLsizei(len(list)), glBool(transpose), ((*C.GLfloat)((unsafe.Pointer)(&list[0]))))
}

func (location UniformLocation) UniformMatrix2f(transpose bool, matrix *[4]float32) {
	if matrix == nil {
		panic("Matrix is nil")
	}
	C.glUniformMatrix2fv(C.GLint(location), 1, glBool(transpose), ((*C.GLfloat)((unsafe.Pointer)(&matrix[0]))))
}

func (location UniformLocation) UniformMatrix3fv(transpose bool, list ...[9]float32) {
	if len(list) < 1 {
		panic("Invalid array length - must be at least 1")
	}
	C.glUniformMatrix3fv(C.GLint(location), C.GLsizei(len(list)), glBool(transpose), ((*C.GLfloat)((unsafe.Pointer)(&list[0]))))
}

func (location UniformLocation) UniformMatrix3f(transpose bool, matrix *[9]float32) {
	if matrix == nil {
		panic("Matrix is nil")
	}
	C.glUniformMatrix3fv(C.GLint(location), 1, glBool(transpose), ((*C.GLfloat)((unsafe.Pointer)(&matrix[0]))))
}

func (location UniformLocation) UniformMatrix4fv(transpose bool, list ...[16]float32) {
	if len(list) < 1 {
		panic("Invalid array length - must be at least 1")
	}
	C.glUniformMatrix4fv(C.GLint(location), C.GLsizei(len(list)), glBool(transpose), ((*C.GLfloat)((unsafe.Pointer)(&list[0]))))
}

func (location UniformLocation) UniformMatrix4f(transpose bool, matrix *[16]float32) {
	if matrix == nil {
		panic("Matrix is nil")
	}
	C.glUniformMatrix4fv(C.GLint(location), 1, glBool(transpose), ((*C.GLfloat)((unsafe.Pointer)(&matrix[0]))))
}

func (location UniformLocation) UniformMatrix2x3fv(transpose bool, list ...[6]float32) {
	if len(list) < 1 {
		panic("Invalid array length - must be at least 1")
	}
	C.glUniformMatrix2x3fv(C.GLint(location), C.GLsizei(len(list)), glBool(transpose), ((*C.GLfloat)((unsafe.Pointer)(&list[0]))))
}

func (location UniformLocation) UniformMatrix2x3f(transpose bool, matrix *[6]float32) {
	if matrix == nil {
		panic("Matrix is nil")
	}
	C.glUniformMatrix2x3fv(C.GLint(location), 1, glBool(transpose), ((*C.GLfloat)((unsafe.Pointer)(&matrix[0]))))
}

func (location UniformLocation) UniformMatrix3x2fv(transpose bool, list ...[6]float32) {
	if len(list) < 1 {
		panic("Invalid array length - must be at least 1")
	}
	C.glUniformMatrix3x2fv(C.GLint(location), C.GLsizei(len(list)), glBool(transpose), ((*C.GLfloat)((unsafe.Pointer)(&list[0]))))
}

func (location UniformLocation) UniformMatrix3x2f(transpose bool, matrix *[6]float32) {
	if matrix == nil {
		panic("Matrix is nil")
	}
	C.glUniformMatrix3x2fv(C.GLint(location), 1, glBool(transpose), ((*C.GLfloat)((unsafe.Pointer)(&matrix[0]))))
}

func (location UniformLocation) UniformMatrix2x4fv(transpose bool, list ...[8]float32) {
	if len(list) < 1 {
		panic("Invalid array length - must be at least 1")
	}
	C.glUniformMatrix2x4fv(C.GLint(location), C.GLsizei(len(list)), glBool(transpose), ((*C.GLfloat)((unsafe.Pointer)(&list[0]))))
}

func (location UniformLocation) UniformMatrix2x4f(transpose bool, matrix *[8]float32) {
	if matrix == nil {
		panic("Matrix is nil")
	}
	C.glUniformMatrix2x4fv(C.GLint(location), 1, glBool(transpose), ((*C.GLfloat)((unsafe.Pointer)(&matrix[0]))))
}

func (location UniformLocation) UniformMatrix4x2fv(transpose bool, list ...[8]float32) {
	if len(list) < 1 {
		panic("Invalid array length - must be at least 1")
	}
	C.glUniformMatrix4x2fv(C.GLint(location), C.GLsizei(len(list)), glBool(transpose), ((*C.GLfloat)((unsafe.Pointer)(&list[0]))))
}

func (location UniformLocation) UniformMatrix4x2f(transpose bool, matrix *[8]float32) {
	if matrix == nil {
		panic("Matrix is nil")
	}
	C.glUniformMatrix4x2fv(C.GLint(location), 1, glBool(transpose), ((*C.GLfloat)((unsafe.Pointer)(&matrix[0]))))
}

func (location UniformLocation) UniformMatrix3x4fv(transpose bool, list ...[12]float32) {
	if len(list) < 1 {
		panic("Invalid array length - must be at least 1")
	}
	C.glUniformMatrix3x4fv(C.GLint(location), C.GLsizei(len(list)), glBool(transpose), ((*C.GLfloat)((unsafe.Pointer)(&list[0]))))
}

func (location UniformLocation) UniformMatrix3x4f(transpose bool, matrix *[12]float32) {
	if matrix == nil {
		panic("Matrix is nil")
	}
	C.glUniformMatrix3x4fv(C.GLint(location), 1, glBool(transpose), ((*C.GLfloat)((unsafe.Pointer)(&matrix[0]))))
}

func (location UniformLocation) UniformMatrix4x3fv(transpose bool, list ...[12]float32) {
	if len(list) < 1 {
		panic("Invalid array length - must be at least 1")
	}
	C.glUniformMatrix4x3fv(C.GLint(location), C.GLsizei(len(list)), glBool(transpose), ((*C.GLfloat)((unsafe.Pointer)(&list[0]))))
}

func (location UniformLocation) UniformMatrix4x3f(transpose bool, matrix *[12]float32) {
	if matrix == nil {
		panic("Matrix is nil")
	}
	C.glUniformMatrix4x3fv(C.GLint(location), 1, glBool(transpose), ((*C.GLfloat)((unsafe.Pointer)(&matrix[0]))))
}
