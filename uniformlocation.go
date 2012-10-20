// Copyright 2012 The go-gl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gl

// #cgo darwin LDFLAGS: -framework OpenGL -lGLEW
// #cgo windows LDFLAGS: -lglew32 -lopengl32
// #cgo linux LDFLAGS: -lGLEW -lGL
//
// #include <stdlib.h>
//
// #include <GL/glew.h>
//
// #undef GLEW_GET_FUN
// #define GLEW_GET_FUN(x) (*x)
import "C"

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

func (location UniformLocation) Uniform1fv(v []float32) {
	C.glUniform1fv(C.GLint(location), C.GLsizei(len(v)),
		(*C.GLfloat)(ptr(v)))
}

func (location UniformLocation) Uniform1i(x int) {
	C.glUniform1i(C.GLint(location), C.GLint(x))
}

func (location UniformLocation) Uniform1iv(v []int32) {
	if len(v) > 0 {
		C.glUniform1iv(C.GLint(location), C.GLsizei(len(v)), (*C.GLint)(&v[0]))
	}
}

func (location UniformLocation) Uniform2fv(v []float32) {
	if len(v) > 0 {
		C.glUniform2fv(C.GLint(location), C.GLsizei(len(v)), (*C.GLfloat)(&v[0]))
	}
}

func (location UniformLocation) Uniform2i(x int, y int) {
	C.glUniform2i(C.GLint(location), C.GLint(x), C.GLint(y))
}

func (location UniformLocation) Uniform2iv(v []int32) {
	if len(v) > 0 {
		C.glUniform2iv(C.GLint(location), C.GLsizei(len(v)), (*C.GLint)(&v[0]))
	}
}

func (location UniformLocation) Uniform3fv(v []float32) {
	if len(v) > 0 {
		C.glUniform3fv(C.GLint(location), C.GLsizei(len(v)), (*C.GLfloat)(&v[0]))
	}
}

func (location UniformLocation) Uniform3i(x int, y int, z int) {
	C.glUniform3i(C.GLint(location), C.GLint(x), C.GLint(y), C.GLint(z))
}

func (location UniformLocation) Uniform3iv(v []int32) {
	if len(v) > 0 {
		C.glUniform3iv(C.GLint(location), C.GLsizei(len(v)), (*C.GLint)(&v[0]))
	}
}

func (location UniformLocation) Uniform4f(x float32, y float32, z float32, w float32) {
	C.glUniform4f(C.GLint(location), C.GLfloat(x), C.GLfloat(y), C.GLfloat(z), C.GLfloat(w))
}

func (location UniformLocation) Uniform4fv(v []float32) {
	if len(v) > 0 {
		C.glUniform4fv(C.GLint(location), C.GLsizei(len(v)), (*C.GLfloat)(&v[0]))
	}
}

func (location UniformLocation) Uniform4i(x int, y int, z int, w int) {
	C.glUniform4i(C.GLint(location), C.GLint(x), C.GLint(y), C.GLint(z), C.GLint(w))
}

func (location UniformLocation) Uniform4iv(v []int32) {
	if len(v) > 0 {
		C.glUniform4iv(C.GLint(location), C.GLsizei(len(v)), (*C.GLint)(&v[0]))
	}
}
