// Copyright 2012 The go-gl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gl

// #include "gl.h"
import "C"

const (
	BUFFER_ACCESS_FLAGS               = C.GL_BUFFER_ACCESS_FLAGS
	BUFFER_MAP_LENGTH                 = C.GL_BUFFER_MAP_LENGTH
	BUFFER_MAP_OFFSET                 = C.GL_BUFFER_MAP_OFFSET
	INT_SAMPLER_2D_RECT               = C.GL_INT_SAMPLER_2D_RECT
	INT_SAMPLER_BUFFER                = C.GL_INT_SAMPLER_BUFFER
	MAX_RECTANGLE_TEXTURE_SIZE        = C.GL_MAX_RECTANGLE_TEXTURE_SIZE
	MAX_TEXTURE_BUFFER_SIZE           = C.GL_MAX_TEXTURE_BUFFER_SIZE
	PRIMITIVE_RESTART                 = C.GL_PRIMITIVE_RESTART
	PRIMITIVE_RESTART_INDEX           = C.GL_PRIMITIVE_RESTART_INDEX
	PROXY_TEXTURE_RECTANGLE           = C.GL_PROXY_TEXTURE_RECTANGLE
	R16_SNORM                         = C.GL_R16_SNORM
	R8_SNORM                          = C.GL_R8_SNORM
	RED_SNORM                         = C.GL_RED_SNORM
	RG16_SNORM                        = C.GL_RG16_SNORM
	RG8_SNORM                         = C.GL_RG8_SNORM
	RG_SNORM                          = C.GL_RG_SNORM
	RGB16_SNORM                       = C.GL_RGB16_SNORM
	RGB8_SNORM                        = C.GL_RGB8_SNORM
	RGB_SNORM                         = C.GL_RGB_SNORM
	RGBA16_SNORM                      = C.GL_RGBA16_SNORM
	RGBA8_SNORM                       = C.GL_RGBA8_SNORM
	RGBA_SNORM                        = C.GL_RGBA_SNORM
	SAMPLER_2D_RECT                   = C.GL_SAMPLER_2D_RECT
	SAMPLER_2D_RECT_SHADOW            = C.GL_SAMPLER_2D_RECT_SHADOW
	SAMPLER_BUFFER                    = C.GL_SAMPLER_BUFFER
	SIGNED_NORMALIZED                 = C.GL_SIGNED_NORMALIZED
	TEXTURE_BINDING_BUFFER            = C.GL_TEXTURE_BINDING_BUFFER
	TEXTURE_BINDING_RECTANGLE         = C.GL_TEXTURE_BINDING_RECTANGLE
	TEXTURE_BUFFER                    = C.GL_TEXTURE_BUFFER
	TEXTURE_BUFFER_DATA_STORE_BINDING = C.GL_TEXTURE_BUFFER_DATA_STORE_BINDING
	TEXTURE_BUFFER_FORMAT             = C.GL_TEXTURE_BUFFER_FORMAT
	TEXTURE_RECTANGLE                 = C.GL_TEXTURE_RECTANGLE
	UNSIGNED_INT_SAMPLER_2D_RECT      = C.GL_UNSIGNED_INT_SAMPLER_2D_RECT
	UNSIGNED_INT_SAMPLER_BUFFER       = C.GL_UNSIGNED_INT_SAMPLER_BUFFER
)

func DrawArraysInstanced(mode GLenum, first int, count int, primcount int) {
	C.glDrawArraysInstanced(C.GLenum(mode), C.GLint(first), C.GLsizei(count), C.GLsizei(primcount))
}

func DrawElementsInstanced(mode GLenum, count GLsizei, typ GLenum, indices interface{}, primcount GLsizei) {
	C.glDrawElementsInstanced(C.GLenum(mode), C.GLsizei(count), C.GLenum(typ), ptr(indices), C.GLsizei(primcount))
}

func PrimitiveRestartIndex(index uint) {
	C.glPrimitiveRestartIndex(C.GLuint(index))
}

func (buf Buffer) TexBuffer(target GLenum, internalFormat GLenum) {
	C.glTexBuffer(C.GLenum(target), C.GLenum(internalFormat), C.GLuint(buf))
}
