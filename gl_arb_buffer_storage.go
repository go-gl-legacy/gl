// Copyright 2012 The go-gl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gl

// #include "gl.h"
import "C"
import "unsafe"

const (
	BUFFER_IMMUTABLE_STORAGE         = C.GL_BUFFER_IMMUTABLE_STORAGE
	BUFFER_STORAGE_FLAGS             = C.GL_BUFFER_STORAGE_FLAGS
	CLIENT_MAPPED_BUFFER_BARRIER_BIT = C.GL_CLIENT_MAPPED_BUFFER_BARRIER_BIT
	CLIENT_STORAGE_BIT               = C.GL_CLIENT_STORAGE_BIT
	DYNAMIC_STORAGE_BIT              = C.GL_DYNAMIC_STORAGE_BIT
	MAP_COHERENT_BIT                 = C.GL_MAP_COHERENT_BIT
	MAP_PERSISTENT_BIT               = C.GL_MAP_PERSISTENT_BIT
)

var ARB_buffer_storage = false

func init() {
	extensions["GL_ARB_buffer_storage"] = &ARB_buffer_storage
}

func BufferStorage(target GLenum, size int, data unsafe.Pointer, flags uint32) {
	C.glBufferStorage(C.GLenum(target), C.GLsizeiptr(size), data, C.GLbitfield(flags))
}

func (buffer Buffer) NamedBufferStorage(size int, data unsafe.Pointer, flags uint32) {
	C.glNamedBufferStorageEXT(C.GLuint(buffer), C.GLsizeiptr(size), data, C.GLbitfield(flags))
}
