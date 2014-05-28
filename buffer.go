// Copyright 2012 The go-gl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gl

// #include "gl.h"
import "C"
import "unsafe"

// Buffer Objects

type Buffer Object

// Create single buffer object
func GenBuffer() Buffer {
	var b C.GLuint
	C.glGenBuffers(1, &b)
	return Buffer(b)
}

// Fill slice with new buffers
func GenBuffers(buffers []Buffer) {
	if len(buffers) > 0 {
		C.glGenBuffers(C.GLsizei(len(buffers)), (*C.GLuint)(&buffers[0]))
	}
}

// Delete buffer object
func (buffer Buffer) Delete() {
	b := C.GLuint(buffer)
	C.glDeleteBuffers(1, &b)
}

// Delete all buffers in slice
func DeleteBuffers(buffers []Buffer) {
	if len(buffers) > 0 {
		C.glDeleteBuffers(C.GLsizei(len(buffers)), (*C.GLuint)(&buffers[0]))
	}
}

// Bind this buffer as target
func (buffer Buffer) Bind(target GLenum) {
	C.glBindBuffer(C.GLenum(target), C.GLuint(buffer))
}

// Remove buffer binding
func (buffer Buffer) Unbind(target GLenum) {
	C.glBindBuffer(C.GLenum(target), C.GLuint(0))
}

// Bind this buffer as index of target
func (buffer Buffer) BindBufferBase(target GLenum, index uint) {
	C.glBindBufferBase(C.GLenum(target), C.GLuint(index), C.GLuint(buffer))
}

// Bind this buffer range as index of target
func (buffer Buffer) BindBufferRange(target GLenum, index uint, offset int, size uint) {
	C.glBindBufferRange(C.GLenum(target), C.GLuint(index), C.GLuint(buffer), C.GLintptr(offset), C.GLsizeiptr(size))
}

// Creates and initializes a buffer object's data store
func BufferData(target GLenum, size int, data interface{}, usage GLenum) {
	C.glBufferData(C.GLenum(target), C.GLsizeiptr(size), glPointer(data), C.GLenum(usage))
}

//  Update a subset of a buffer object's data store
func BufferSubData(target GLenum, offset int, size int, data interface{}) {
	C.glBufferSubData(C.GLenum(target), C.GLintptr(offset), C.GLsizeiptr(size),
		glPointer(data))
}

// Returns a subset of a buffer object's data store
func GetBufferSubData(target GLenum, offset int, size int, data interface{}) {
	C.glGetBufferSubData(C.GLenum(target), C.GLintptr(offset),
		C.GLsizeiptr(size), glPointer(data))
}

//  Map a buffer object's data store
func MapBuffer(target GLenum, access GLenum) unsafe.Pointer {
	return unsafe.Pointer(C.glMapBuffer(C.GLenum(target), C.GLenum(access)))
}

//  Unmap a buffer object's data store
func UnmapBuffer(target GLenum) bool {
	return goBool(C.glUnmapBuffer(C.GLenum(target)))
}

// Return buffer pointer
func GetBufferPointerv(target GLenum, pname GLenum) unsafe.Pointer {
	var ptr unsafe.Pointer
	C.glGetBufferPointerv(C.GLenum(target), C.GLenum(pname), &ptr)
	return ptr
}

// Return parameters of a buffer object
func GetBufferParameteriv(target GLenum, pname GLenum) int32 {
	var param C.GLint
	C.glGetBufferParameteriv(C.GLenum(target), C.GLenum(pname), &param)
	return int32(param)
}
