// Copyright 2012 The go-gl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gl

// #include "gl.h"
import "C"
import "unsafe"

const (
	NUM_PROGRAM_BINARY_FORMATS      = C.GL_NUM_PROGRAM_BINARY_FORMATS
	PROGRAM_BINARY_FORMATS          = C.GL_PROGRAM_BINARY_FORMATS
	PROGRAM_BINARY_LENGTH           = C.GL_PROGRAM_BINARY_LENGTH
	PROGRAM_BINARY_RETRIEVABLE_HINT = C.GL_PROGRAM_BINARY_RETRIEVABLE_HINT
)

// Binary loads a program object with a binary previously returned from GetBinary.
func (program Program) Binary(format GLenum, binary []byte) {
	C.glProgramBinary(C.GLuint(program), C.GLenum(format), unsafe.Pointer(&binary[0]), C.GLsizei(len(binary)))
}

// Parameteri sets an integer program parameter to the specified value.
func (program Program) Parameteri(pname GLenum, value int) {
	C.glProgramParameteri(C.GLuint(program), C.GLenum(pname), C.GLint(value))
}

// GetBinary retrieves a program's binary, it returns the actual program length
// (which may be different than len(binary)) and its binary encoding format.
func (program Program) GetBinary(binary []byte) (length int, format GLenum) {
	var glformat C.GLenum
	var gllength C.GLsizei

	C.glGetProgramBinary(C.GLuint(program), C.GLsizei(len(binary)), &gllength, &glformat, unsafe.Pointer(&binary[0]))
	length, format = int(gllength), GLenum(glformat)
	return
}
