// Copyright 2012 The go-gl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gl

// #include "gl.h"
import "C"

const (
	CONTEXT_COMPATIBILITY_PROFILE_BIT    = C.GL_CONTEXT_COMPATIBILITY_PROFILE_BIT
	CONTEXT_CORE_PROFILE_BIT             = C.GL_CONTEXT_CORE_PROFILE_BIT
	CONTEXT_PROFILE_MASK                 = C.GL_CONTEXT_PROFILE_MASK
	FRAMEBUFFER_ATTACHMENT_LAYERED       = C.GL_FRAMEBUFFER_ATTACHMENT_LAYERED
	FRAMEBUFFER_INCOMPLETE_LAYER_TARGETS = C.GL_FRAMEBUFFER_INCOMPLETE_LAYER_TARGETS
	GEOMETRY_INPUT_TYPE                  = C.GL_GEOMETRY_INPUT_TYPE
	GEOMETRY_OUTPUT_TYPE                 = C.GL_GEOMETRY_OUTPUT_TYPE
	GEOMETRY_SHADER                      = C.GL_GEOMETRY_SHADER
	GEOMETRY_VERTICES_OUT                = C.GL_GEOMETRY_VERTICES_OUT
	LINE_STRIP_ADJACENCY                 = C.GL_LINE_STRIP_ADJACENCY
	LINES_ADJACENCY                      = C.GL_LINES_ADJACENCY
	MAX_FRAGMENT_INPUT_COMPONENTS        = C.GL_MAX_FRAGMENT_INPUT_COMPONENTS
	MAX_GEOMETRY_INPUT_COMPONENTS        = C.GL_MAX_GEOMETRY_INPUT_COMPONENTS
	MAX_GEOMETRY_OUTPUT_COMPONENTS       = C.GL_MAX_GEOMETRY_OUTPUT_COMPONENTS
	MAX_GEOMETRY_OUTPUT_VERTICES         = C.GL_MAX_GEOMETRY_OUTPUT_VERTICES
	MAX_GEOMETRY_TEXTURE_IMAGE_UNITS     = C.GL_MAX_GEOMETRY_TEXTURE_IMAGE_UNITS
	MAX_GEOMETRY_TOTAL_OUTPUT_COMPONENTS = C.GL_MAX_GEOMETRY_TOTAL_OUTPUT_COMPONENTS
	MAX_GEOMETRY_UNIFORM_COMPONENTS      = C.GL_MAX_GEOMETRY_UNIFORM_COMPONENTS
	MAX_VERTEX_OUTPUT_COMPONENTS         = C.GL_MAX_VERTEX_OUTPUT_COMPONENTS
	PROGRAM_POINT_SIZE                   = C.GL_PROGRAM_POINT_SIZE
	TRIANGLE_STRIP_ADJACENCY             = C.GL_TRIANGLE_STRIP_ADJACENCY
	TRIANGLES_ADJACENCY                  = C.GL_TRIANGLES_ADJACENCY
)

func FramebufferTexture(target GLenum, attachment GLenum, texture uint, level int) {
	C.glFramebufferTexture(C.GLenum(target), C.GLenum(attachment), C.GLuint(texture), C.GLint(level))
}

func GetBufferParameteri64v(target GLenum, value GLenum, data []int64) {
	C.glGetBufferParameteri64v(C.GLenum(target), C.GLenum(value), (*C.GLint64)(&data[0]))
}

func GetInteger64i_v(pname GLenum, index uint, data []int64) {
	C.glGetInteger64i_v(C.GLenum(pname), C.GLuint(index), (*C.GLint64)(&data[0]))
}
