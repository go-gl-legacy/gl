// Copyright 2012 The go-gl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gl

// #include "gl.h"
import "C"
import "unsafe"

const (
	ARRAY_BUFFER                         = C.GL_ARRAY_BUFFER
	ARRAY_BUFFER_BINDING                 = C.GL_ARRAY_BUFFER_BINDING
	BUFFER_ACCESS                        = C.GL_BUFFER_ACCESS
	BUFFER_MAP_POINTER                   = C.GL_BUFFER_MAP_POINTER
	BUFFER_MAPPED                        = C.GL_BUFFER_MAPPED
	BUFFER_SIZE                          = C.GL_BUFFER_SIZE
	BUFFER_USAGE                         = C.GL_BUFFER_USAGE
	COLOR_ARRAY_BUFFER_BINDING           = C.GL_COLOR_ARRAY_BUFFER_BINDING
	CURRENT_FOG_COORD                    = C.GL_CURRENT_FOG_COORD
	CURRENT_QUERY                        = C.GL_CURRENT_QUERY
	DYNAMIC_COPY                         = C.GL_DYNAMIC_COPY
	DYNAMIC_DRAW                         = C.GL_DYNAMIC_DRAW
	DYNAMIC_READ                         = C.GL_DYNAMIC_READ
	EDGE_FLAG_ARRAY_BUFFER_BINDING       = C.GL_EDGE_FLAG_ARRAY_BUFFER_BINDING
	ELEMENT_ARRAY_BUFFER                 = C.GL_ELEMENT_ARRAY_BUFFER
	ELEMENT_ARRAY_BUFFER_BINDING         = C.GL_ELEMENT_ARRAY_BUFFER_BINDING
	FOG_COORD                            = C.GL_FOG_COORD
	FOG_COORD_ARRAY                      = C.GL_FOG_COORD_ARRAY
	FOG_COORD_ARRAY_BUFFER_BINDING       = C.GL_FOG_COORD_ARRAY_BUFFER_BINDING
	FOG_COORD_ARRAY_POINTER              = C.GL_FOG_COORD_ARRAY_POINTER
	FOG_COORD_ARRAY_STRIDE               = C.GL_FOG_COORD_ARRAY_STRIDE
	FOG_COORD_ARRAY_TYPE                 = C.GL_FOG_COORD_ARRAY_TYPE
	FOG_COORD_SRC                        = C.GL_FOG_COORD_SRC
	FOG_COORDINATE_ARRAY_BUFFER_BINDING  = C.GL_FOG_COORDINATE_ARRAY_BUFFER_BINDING
	INDEX_ARRAY_BUFFER_BINDING           = C.GL_INDEX_ARRAY_BUFFER_BINDING
	NORMAL_ARRAY_BUFFER_BINDING          = C.GL_NORMAL_ARRAY_BUFFER_BINDING
	QUERY_COUNTER_BITS                   = C.GL_QUERY_COUNTER_BITS
	QUERY_RESULT                         = C.GL_QUERY_RESULT
	QUERY_RESULT_AVAILABLE               = C.GL_QUERY_RESULT_AVAILABLE
	READ_ONLY                            = C.GL_READ_ONLY
	READ_WRITE                           = C.GL_READ_WRITE
	SAMPLES_PASSED                       = C.GL_SAMPLES_PASSED
	SECONDARY_COLOR_ARRAY_BUFFER_BINDING = C.GL_SECONDARY_COLOR_ARRAY_BUFFER_BINDING
	SRC0_ALPHA                           = C.GL_SRC0_ALPHA
	SRC0_RGB                             = C.GL_SRC0_RGB
	SRC1_ALPHA                           = C.GL_SRC1_ALPHA
	SRC1_RGB                             = C.GL_SRC1_RGB
	SRC2_ALPHA                           = C.GL_SRC2_ALPHA
	SRC2_RGB                             = C.GL_SRC2_RGB
	STATIC_COPY                          = C.GL_STATIC_COPY
	STATIC_DRAW                          = C.GL_STATIC_DRAW
	STATIC_READ                          = C.GL_STATIC_READ
	STREAM_COPY                          = C.GL_STREAM_COPY
	STREAM_DRAW                          = C.GL_STREAM_DRAW
	STREAM_READ                          = C.GL_STREAM_READ
	TEXTURE_COORD_ARRAY_BUFFER_BINDING   = C.GL_TEXTURE_COORD_ARRAY_BUFFER_BINDING
	VERTEX_ARRAY_BUFFER_BINDING          = C.GL_VERTEX_ARRAY_BUFFER_BINDING
	VERTEX_ATTRIB_ARRAY_BUFFER_BINDING   = C.GL_VERTEX_ATTRIB_ARRAY_BUFFER_BINDING
	WEIGHT_ARRAY_BUFFER_BINDING          = C.GL_WEIGHT_ARRAY_BUFFER_BINDING
	WRITE_ONLY                           = C.GL_WRITE_ONLY
)

func (query Query) Begin(target GLenum) {
	C.glBeginQuery(C.GLenum(target), C.GLuint(query))
}

func (buffer Buffer) Bind(target GLenum) {
	C.glBindBuffer(C.GLenum(target), C.GLuint(buffer))
}

func (_ Buffer) Unbind(target GLenum) {
	C.glBindBuffer(C.GLenum(target), C.GLuint(0))
}

func BufferData(target GLenum, size int, data interface{}, usage GLenum) {
	C.glBufferData(C.GLenum(target), C.GLsizeiptr(size), glPointer(data), C.GLenum(usage))
}

func BufferSubData(target GLenum, offset int, size uint, data interface{}) {
	C.glBufferSubData(C.GLenum(target), C.GLintptr(offset), C.GLsizeiptr(size), glPointer(data))
}

func DeleteBuffers(buffers []Buffer) {
	C.glDeleteBuffers(C.GLsizei(len(buffers)), (*C.GLuint)(&buffers[0]))
}

func (buffer Buffer) Delete() {
	C.glDeleteBuffers(C.GLsizei(1), (*C.GLuint)(&buffer))
}

func DeleteQueries(ids []Query) {
	C.glDeleteQueries(C.GLsizei(len(ids)), (*C.GLuint)(&ids[0]))
}

func (query Query) Delete() {
	C.glDeleteQueries(C.GLsizei(1), (*C.GLuint)(&query))
}

func EndQuery(target GLenum) {
	C.glEndQuery(C.GLenum(target))
}

func GenBuffers(buffers []Buffer) {
	C.glGenBuffers(C.GLsizei(len(buffers)), (*C.GLuint)(&buffers[0]))
}

func GenBuffer() (buffer Buffer) {
	C.glGenBuffers(C.GLsizei(1), (*C.GLuint)(&buffer))
	return
}

func GenQueries(ids []Query) {
	C.glGenQueries(C.GLsizei(len(ids)), (*C.GLuint)(&ids[0]))
}

func GenQuery() (query Query) {
	C.glGenQueries(C.GLsizei(1), (*C.GLuint)(&query))
	return
}

func GetBufferParameteriv(target GLenum, pname GLenum, params []int32) {
	C.glGetBufferParameteriv(C.GLenum(target), C.GLenum(pname), (*C.GLint)(&params[0]))
}

func GetBufferPointerv(target GLenum, pname GLenum, params []unsafe.Pointer) {
	C.glGetBufferPointerv(C.GLenum(target), C.GLenum(pname), &params[0])
}

func GetBufferSubData(target GLenum, offset int, size int, data interface{}) {
	C.glGetBufferSubData(C.GLenum(target), C.GLintptr(offset), C.GLsizeiptr(size), glPointer(data))
}

func GetQueryObjectiv(id uint, pname GLenum, params []int32) {
	C.glGetQueryObjectiv(C.GLuint(id), C.GLenum(pname), (*C.GLint)(&params[0]))
}

func GetQueryObjectuiv(id uint, pname GLenum, params []uint32) {
	C.glGetQueryObjectuiv(C.GLuint(id), C.GLenum(pname), (*C.GLuint)(&params[0]))
}

func GetQueryiv(target GLenum, pname GLenum, params []int32) {
	C.glGetQueryiv(C.GLenum(target), C.GLenum(pname), (*C.GLint)(&params[0]))
}

func (object Object) IsBuffer() bool {
	return goBool(C.glIsBuffer(C.GLuint(object)))
}

func (object Object) IsQuery() bool {
	return goBool(C.glIsQuery(C.GLuint(object)))
}

func MapBuffer(target GLenum, access GLenum) unsafe.Pointer {
	return C.glMapBuffer(C.GLenum(target), C.GLenum(access))
}

func UnmapBuffer(target GLenum) bool {
	return goBool(C.glUnmapBuffer(C.GLenum(target)))
}
