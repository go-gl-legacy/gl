// Copyright 2012 The go-gl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gl

// #include "gl.h"
import "C"

const (
	ALIASED_LINE_WIDTH_RANGE      = C.GL_ALIASED_LINE_WIDTH_RANGE
	ALIASED_POINT_SIZE_RANGE      = C.GL_ALIASED_POINT_SIZE_RANGE
	BGR                           = C.GL_BGR
	BGRA                          = C.GL_BGRA
	CLAMP_TO_EDGE                 = C.GL_CLAMP_TO_EDGE
	LIGHT_MODEL_COLOR_CONTROL     = C.GL_LIGHT_MODEL_COLOR_CONTROL
	MAX_3D_TEXTURE_SIZE           = C.GL_MAX_3D_TEXTURE_SIZE
	MAX_ELEMENTS_INDICES          = C.GL_MAX_ELEMENTS_INDICES
	MAX_ELEMENTS_VERTICES         = C.GL_MAX_ELEMENTS_VERTICES
	PACK_IMAGE_HEIGHT             = C.GL_PACK_IMAGE_HEIGHT
	PACK_SKIP_IMAGES              = C.GL_PACK_SKIP_IMAGES
	PROXY_TEXTURE_3D              = C.GL_PROXY_TEXTURE_3D
	RESCALE_NORMAL                = C.GL_RESCALE_NORMAL
	SEPARATE_SPECULAR_COLOR       = C.GL_SEPARATE_SPECULAR_COLOR
	SINGLE_COLOR                  = C.GL_SINGLE_COLOR
	SMOOTH_LINE_WIDTH_GRANULARITY = C.GL_SMOOTH_LINE_WIDTH_GRANULARITY
	SMOOTH_LINE_WIDTH_RANGE       = C.GL_SMOOTH_LINE_WIDTH_RANGE
	SMOOTH_POINT_SIZE_GRANULARITY = C.GL_SMOOTH_POINT_SIZE_GRANULARITY
	SMOOTH_POINT_SIZE_RANGE       = C.GL_SMOOTH_POINT_SIZE_RANGE
	TEXTURE_3D                    = C.GL_TEXTURE_3D
	TEXTURE_BASE_LEVEL            = C.GL_TEXTURE_BASE_LEVEL
	TEXTURE_BINDING_3D            = C.GL_TEXTURE_BINDING_3D
	TEXTURE_DEPTH                 = C.GL_TEXTURE_DEPTH
	TEXTURE_MAX_LEVEL             = C.GL_TEXTURE_MAX_LEVEL
	TEXTURE_MAX_LOD               = C.GL_TEXTURE_MAX_LOD
	TEXTURE_MIN_LOD               = C.GL_TEXTURE_MIN_LOD
	TEXTURE_WRAP_R                = C.GL_TEXTURE_WRAP_R
	UNPACK_IMAGE_HEIGHT           = C.GL_UNPACK_IMAGE_HEIGHT
	UNPACK_SKIP_IMAGES            = C.GL_UNPACK_SKIP_IMAGES
	UNSIGNED_BYTE_2_3_3_REV       = C.GL_UNSIGNED_BYTE_2_3_3_REV
	UNSIGNED_BYTE_3_3_2           = C.GL_UNSIGNED_BYTE_3_3_2
	UNSIGNED_INT_10_10_10_2       = C.GL_UNSIGNED_INT_10_10_10_2
	UNSIGNED_INT_8_8_8_8          = C.GL_UNSIGNED_INT_8_8_8_8
	UNSIGNED_INT_8_8_8_8_REV      = C.GL_UNSIGNED_INT_8_8_8_8_REV
	UNSIGNED_SHORT_1_5_5_5_REV    = C.GL_UNSIGNED_SHORT_1_5_5_5_REV
	UNSIGNED_SHORT_4_4_4_4        = C.GL_UNSIGNED_SHORT_4_4_4_4
	UNSIGNED_SHORT_4_4_4_4_REV    = C.GL_UNSIGNED_SHORT_4_4_4_4_REV
	UNSIGNED_SHORT_5_5_5_1        = C.GL_UNSIGNED_SHORT_5_5_5_1
	UNSIGNED_SHORT_5_6_5          = C.GL_UNSIGNED_SHORT_5_6_5
	UNSIGNED_SHORT_5_6_5_REV      = C.GL_UNSIGNED_SHORT_5_6_5_REV
)

func CopyTexSubImage3D(target GLenum, level int, xoffset int, yoffset int,
	zoffset int, x int, y int, width int, height int) {
	C.glCopyTexSubImage3D(C.GLenum(target), C.GLint(level), C.GLint(xoffset),
		C.GLint(yoffset), C.GLint(zoffset), C.GLint(x), C.GLint(y),
		C.GLsizei(width), C.GLsizei(height))
}

func DrawRangeElements(mode GLenum, start uint, end uint, count GLsizei,
	typ GLenum, indices interface{}) {
	C.glDrawRangeElements(C.GLenum(mode), C.GLuint(start), C.GLuint(end),
		C.GLsizei(count), C.GLenum(typ), ptr(indices))
}

func TexImage3D(target GLenum, level int, internalFormat int, width int,
	height int, depth int, border int, format GLenum, typ GLenum,
	pixels interface{}) {
	C.glTexImage3D(C.GLenum(target), C.GLint(level), C.GLint(internalFormat),
		C.GLsizei(width), C.GLsizei(height), C.GLsizei(depth), C.GLint(border),
		C.GLenum(format), C.GLenum(typ), ptr(pixels))
}

func TexSubImage3D(target GLenum, level int, xoffset, yoffset, zoffset, width,
	height, depth int, format, typ GLenum, pixels interface{}) {
	C.glTexSubImage3D(C.GLenum(target), C.GLint(level), C.GLint(xoffset),
		C.GLint(yoffset), C.GLint(zoffset), C.GLsizei(width), C.GLsizei(height),
		C.GLsizei(depth),	C.GLenum(format), C.GLenum(typ), ptr(pixels))
}
