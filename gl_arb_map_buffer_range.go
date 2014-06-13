// Copyright 2012 The go-gl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gl

// #include "gl.h"
import "C"
import "unsafe"

const (
	MAP_FLUSH_EXPLICIT_BIT = C.GL_MAP_FLUSH_EXPLICIT_BIT
	MAP_INVALIDATE_BUFFER_BIT = C.GL_MAP_INVALIDATE_BUFFER_BIT
	MAP_INVALIDATE_RANGE_BIT = C.GL_MAP_INVALIDATE_RANGE_BIT
	MAP_READ_BIT = C.GL_MAP_READ_BIT
	MAP_UNSYNCHRONIZED_BIT = C.GL_MAP_UNSYNCHRONIZED_BIT
	MAP_WRITE_BIT = C.GL_MAP_WRITE_BIT
)

var ARB_map_buffer_range = false

func init() {
	extensions["GL_ARB_map_buffer_range"] = &ARB_map_buffer_range
}

func FlushMappedBufferRange(target GLenum, offset int, length int) {
	C.glFlushMappedBufferRange(C.GLenum(target), C.GLintptr(offset), C.GLsizeiptr(length))
}

func MapBufferRange(target GLenum, offset int, length int, access uint32) unsafe.Pointer {
	return C.glMapBufferRange(C.GLenum(target), C.GLintptr(offset), C.GLsizeiptr(length), C.GLbitfield(access))
}
