// Copyright 2012 The go-gl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gl

// #include "gl.h"
//
// // Workaround for https://github.com/go-gl/gl/issues/104
// void gogl_glGetShaderSource(GLuint shader, GLsizei bufsize, GLsizei* len, GLchar* source) {
//     glGetShaderSource(shader, bufsize, len, source);
// }
//
import "C"

// Shader

type Shader Object

func CreateShader(type_ GLenum) Shader { return Shader(C.glCreateShader(C.GLenum(type_))) }

func (shader Shader) Delete() { C.glDeleteShader(C.GLuint(shader)) }

func (shader Shader) GetInfoLog() string {
	var length C.GLint
	C.glGetShaderiv(C.GLuint(shader), C.GLenum(INFO_LOG_LENGTH), &length)
	// length is buffer size including null character

	if length > 1 {
		log := C.malloc(C.size_t(length))
		defer C.free(log)
		C.glGetShaderInfoLog(C.GLuint(shader), C.GLsizei(length), nil, (*C.GLchar)(log))
		return C.GoString((*C.char)(log))
	}
	return ""
}

func (shader Shader) GetSource() string {
	var length C.GLint
	C.glGetShaderiv(C.GLuint(shader), C.GLenum(SHADER_SOURCE_LENGTH), &length)

	log := C.malloc(C.size_t(len + 1))
	C.gogl_glGetShaderSource(C.GLuint(shader), C.GLsizei(len), nil, (*C.GLchar)(log))

	defer C.free(log)

	if length > 1 {
		log := C.malloc(C.size_t(length + 1))
		defer C.free(log)
		C.glGetShaderSource(C.GLuint(shader), C.GLsizei(length), nil, (*C.GLchar)(log))
		return C.GoString((*C.char)(log))
	}
	return ""
}

func (shader Shader) Source(source string) {

	csource := glString(source)
	defer freeString(csource)

	var one C.GLint = C.GLint(len(source))

	C.glShaderSource(C.GLuint(shader), 1, &csource, &one)
}

func (shader Shader) Compile() { C.glCompileShader(C.GLuint(shader)) }

func (shader Shader) Get(param GLenum) int {
	var rv C.GLint

	C.glGetShaderiv(C.GLuint(shader), C.GLenum(param), &rv)
	return int(rv)
}
