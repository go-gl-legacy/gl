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

func (location UniformLocation) Uniform1fv(count int, v []float32) {
	if len(v) < 1 {
		panic(ErrorInputSize)
	}
	C.glUniform1fv(C.GLint(location), C.GLsizei(count), (*C.GLfloat)(&v[0]))
}

func (location UniformLocation) Uniform1i(x int) {
	C.glUniform1i(C.GLint(location), C.GLint(x))
}

func (location UniformLocation) Uniform1iv(count int, v []int32) {
	if len(v) < 1 {
		panic(ErrorInputSize)
	}
	C.glUniform1iv(C.GLint(location), C.GLsizei(count), (*C.GLint)(&v[0]))
}

func (location UniformLocation) Uniform2fv(count int, v []float32) {
	if len(v) < 1 {
		panic(ErrorInputSize)
	}
	C.glUniform2fv(C.GLint(location), C.GLsizei(count), (*C.GLfloat)(&v[0]))
}

func (location UniformLocation) Uniform2i(x int, y int) {
	C.glUniform2i(C.GLint(location), C.GLint(x), C.GLint(y))
}

func (location UniformLocation) Uniform2iv(count int, v []int32) {
	if len(v) < 1 {
		panic(ErrorInputSize)
	}
	C.glUniform2iv(C.GLint(location), C.GLsizei(count), (*C.GLint)(&v[0]))
}

func (location UniformLocation) Uniform3fv(count int, v []float32) {
	if len(v) < 1 {
		panic(ErrorInputSize)
	}
	C.glUniform3fv(C.GLint(location), C.GLsizei(count), (*C.GLfloat)(&v[0]))
}

func (location UniformLocation) Uniform3i(x int, y int, z int) {
	C.glUniform3i(C.GLint(location), C.GLint(x), C.GLint(y), C.GLint(z))
}

func (location UniformLocation) Uniform3iv(count int, v []int32) {
	if len(v) < 1 {
		panic(ErrorInputSize)
	}
	C.glUniform3iv(C.GLint(location), C.GLsizei(count), (*C.GLint)(&v[0]))
}

func (location UniformLocation) Uniform4f(x float32, y float32, z float32, w float32) {
	C.glUniform4f(C.GLint(location), C.GLfloat(x), C.GLfloat(y), C.GLfloat(z), C.GLfloat(w))
}

func (location UniformLocation) Uniform4fv(count int, v []float32) {
	if len(v) < 1 {
		panic(ErrorInputSize)
	}
	C.glUniform4fv(C.GLint(location), C.GLsizei(count), (*C.GLfloat)(&v[0]))
}

func (location UniformLocation) Uniform4i(x int, y int, z int, w int) {
	C.glUniform4i(C.GLint(location), C.GLint(x), C.GLint(y), C.GLint(z), C.GLint(w))
}

func (location UniformLocation) Uniform4iv(count int, v []int32) {
	if len(v) < 1 {
		panic(ErrorInputSize)
	}
	C.glUniform4iv(C.GLint(location), C.GLsizei(count), (*C.GLint)(&v[0]))
}

func (location UniformLocation) UniformMatrix2fv(transpose bool, list ...[]float32) {
	if len(list) < 1 {
		panic(ErrorInputSize)
	}
	// Maybe this isn't the best solution because of added overhead.
	for i, _ := range list {
		if len(list[i]) != 2 {
			panic(ErrorInputSize)
		}
	}
	C.glUniformMatrix2fv(C.GLint(location), C.GLsizei(len(list)), glBool(transpose), ((*C.GLfloat)((unsafe.Pointer)(&list[0]))))
}

func (location UniformLocation) UniformMatrix2f(transpose bool, matrix []float32) {
	if len(matrix) != 2 {
		panic(ErrorInputSize)
	}
	C.glUniformMatrix2fv(C.GLint(location), 1, glBool(transpose), ((*C.GLfloat)((unsafe.Pointer)(&matrix[0]))))
}

func (location UniformLocation) UniformMatrix3fv(transpose bool, list ...[]float32) {
	if len(list) < 1 {
		panic(ErrorInputSize)
	}
	// Maybe this isn't the best solution because of added overhead.
	for i, _ := range list {
		if len(list[i]) != 9 {
			panic(ErrorInputSize)
		}
	}
	C.glUniformMatrix3fv(C.GLint(location), C.GLsizei(len(list)), glBool(transpose), ((*C.GLfloat)((unsafe.Pointer)(&list[0]))))
}

func (location UniformLocation) UniformMatrix3f(transpose bool, matrix []float32) {
	if len(matrix) != 9 {
		panic(ErrorInputSize)
	}
	C.glUniformMatrix3fv(C.GLint(location), 1, glBool(transpose), ((*C.GLfloat)((unsafe.Pointer)(&matrix[0]))))
}

func (location UniformLocation) UniformMatrix4fv(transpose bool, list ...[]float32) {
	if len(list) < 1 {
		panic(ErrorInputSize)
	}
	// Maybe this isn't the best solution because of added overhead.
	for i, _ := range list {
		if len(list[i]) != 16 {
			panic(ErrorInputSize)
		}
	}
	C.glUniformMatrix4fv(C.GLint(location), C.GLsizei(len(list)), glBool(transpose), ((*C.GLfloat)((unsafe.Pointer)(&list[0]))))
}

func (location UniformLocation) UniformMatrix4f(transpose bool, matrix []float32) {
	if len(matrix) != 16 {
		panic(ErrorInputSize)
	}
	C.glUniformMatrix4fv(C.GLint(location), 1, glBool(transpose), ((*C.GLfloat)((unsafe.Pointer)(&matrix[0]))))
}

func (location UniformLocation) UniformMatrix2x3fv(transpose bool, list ...[]float32) {
	if len(list) < 1 {
		panic(ErrorInputSize)
	}
	// Maybe this isn't the best solution because of added overhead.
	for i, _ := range list {
		if len(list[i]) != 6 {
			panic(ErrorInputSize)
		}
	}
	C.glUniformMatrix2x3fv(C.GLint(location), C.GLsizei(len(list)), glBool(transpose), ((*C.GLfloat)((unsafe.Pointer)(&list[0]))))
}

func (location UniformLocation) UniformMatrix2x3f(transpose bool, matrix []float32) {
	if len(matrix) != 6 {
		panic(ErrorInputSize)
	}
	C.glUniformMatrix2x3fv(C.GLint(location), 1, glBool(transpose), ((*C.GLfloat)((unsafe.Pointer)(&matrix[0]))))
}

func (location UniformLocation) UniformMatrix3x2fv(transpose bool, list ...[]float32) {
	if len(list) < 1 {
		panic(ErrorInputSize)
	}
	// Maybe this isn't the best solution because of added overhead.
	for i, _ := range list {
		if len(list[i]) != 6 {
			panic(ErrorInputSize)
		}
	}
	C.glUniformMatrix3x2fv(C.GLint(location), C.GLsizei(len(list)), glBool(transpose), ((*C.GLfloat)((unsafe.Pointer)(&list[0]))))
}

func (location UniformLocation) UniformMatrix3x2f(transpose bool, matrix []float32) {
	if len(matrix) != 6 {
		panic(ErrorInputSize)
	}
	C.glUniformMatrix3x2fv(C.GLint(location), 1, glBool(transpose), ((*C.GLfloat)((unsafe.Pointer)(&matrix[0]))))
}

func (location UniformLocation) UniformMatrix2x4fv(transpose bool, list ...[]float32) {
	if len(list) < 1 {
		panic(ErrorInputSize)
	}
	// Maybe this isn't the best solution because of added overhead.
	for i, _ := range list {
		if len(list[i]) != 6 {
			panic(ErrorInputSize)
		}
	}
	C.glUniformMatrix2x4fv(C.GLint(location), C.GLsizei(len(list)), glBool(transpose), ((*C.GLfloat)((unsafe.Pointer)(&list[0]))))
}

func (location UniformLocation) UniformMatrix2x4f(transpose bool, matrix []float32) {
	if len(matrix) != 8 {
		panic(ErrorInputSize)
	}
	C.glUniformMatrix2x4fv(C.GLint(location), 1, glBool(transpose), ((*C.GLfloat)((unsafe.Pointer)(&matrix[0]))))
}

func (location UniformLocation) UniformMatrix4x2fv(transpose bool, list ...[]float32) {
	if len(list) < 1 {
		panic(ErrorInputSize)
	}
	// Maybe this isn't the best solution because of added overhead.
	for i, _ := range list {
		if len(list[i]) != 8 {
			panic(ErrorInputSize)
		}
	}
	C.glUniformMatrix4x2fv(C.GLint(location), C.GLsizei(len(list)), glBool(transpose), ((*C.GLfloat)((unsafe.Pointer)(&list[0]))))
}

func (location UniformLocation) UniformMatrix4x2f(transpose bool, matrix []float32) {
	if len(matrix) != 8 {
		panic(ErrorInputSize)
	}
	C.glUniformMatrix4x2fv(C.GLint(location), 1, glBool(transpose), ((*C.GLfloat)((unsafe.Pointer)(&matrix[0]))))
}

func (location UniformLocation) UniformMatrix3x4fv(transpose bool, list ...[]float32) {
	if len(list) < 1 {
		panic(ErrorInputSize)
	}
	// Maybe this isn't the best solution because of added overhead.
	for i, _ := range list {
		if len(list[i]) != 12 {
			panic(ErrorInputSize)
		}
	}
	C.glUniformMatrix3x4fv(C.GLint(location), C.GLsizei(len(list)), glBool(transpose), ((*C.GLfloat)((unsafe.Pointer)(&list[0]))))
}

func (location UniformLocation) UniformMatrix3x4f(transpose bool, matrix []float32) {
	if len(matrix) != 12 {
		panic(ErrorInputSize)
	}
	C.glUniformMatrix3x4fv(C.GLint(location), 1, glBool(transpose), ((*C.GLfloat)((unsafe.Pointer)(&matrix[0]))))
}

func (location UniformLocation) UniformMatrix4x3fv(transpose bool, list ...[]float32) {
	if len(list) < 1 {
		panic(ErrorInputSize)
	}
	// Maybe this isn't the best solution because of added overhead.
	for i, _ := range list {
		if len(list[i]) != 12 {
			panic(ErrorInputSize)
		}
	}
	C.glUniformMatrix4x3fv(C.GLint(location), C.GLsizei(len(list)), glBool(transpose), ((*C.GLfloat)((unsafe.Pointer)(&list[0]))))
}

func (location UniformLocation) UniformMatrix4x3f(transpose bool, matrix []float32) {
	if len(matrix) != 12 {
		panic(ErrorInputSize)
	}
	C.glUniformMatrix4x3fv(C.GLint(location), 1, glBool(transpose), ((*C.GLfloat)((unsafe.Pointer)(&matrix[0]))))
}
