// Copyright 2012 The go-gl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gl

// #include "gl.h"
import "C"

const (
	INT_SAMPLER_CUBE_MAP_ARRAY            = C.GL_INT_SAMPLER_CUBE_MAP_ARRAY
	MAX_PROGRAM_TEXTURE_GATHER_COMPONENTS = C.GL_MAX_PROGRAM_TEXTURE_GATHER_COMPONENTS
	MAX_PROGRAM_TEXTURE_GATHER_OFFSET     = C.GL_MAX_PROGRAM_TEXTURE_GATHER_OFFSET
	MIN_PROGRAM_TEXTURE_GATHER_OFFSET     = C.GL_MIN_PROGRAM_TEXTURE_GATHER_OFFSET
	MIN_SAMPLE_SHADING_VALUE              = C.GL_MIN_SAMPLE_SHADING_VALUE
	PROXY_TEXTURE_CUBE_MAP_ARRAY          = C.GL_PROXY_TEXTURE_CUBE_MAP_ARRAY
	SAMPLE_SHADING                        = C.GL_SAMPLE_SHADING
	SAMPLER_CUBE_MAP_ARRAY                = C.GL_SAMPLER_CUBE_MAP_ARRAY
	SAMPLER_CUBE_MAP_ARRAY_SHADOW         = C.GL_SAMPLER_CUBE_MAP_ARRAY_SHADOW
	TEXTURE_BINDING_CUBE_MAP_ARRAY        = C.GL_TEXTURE_BINDING_CUBE_MAP_ARRAY
	TEXTURE_CUBE_MAP_ARRAY                = C.GL_TEXTURE_CUBE_MAP_ARRAY
	UNSIGNED_INT_SAMPLER_CUBE_MAP_ARRAY   = C.GL_UNSIGNED_INT_SAMPLER_CUBE_MAP_ARRAY
)

func BlendEquationSeparatei(buf uint, modeRGB GLenum, modeAlpha GLenum) {
	C.glBlendEquationSeparatei(C.GLuint(buf), C.GLenum(modeRGB), C.GLenum(modeAlpha))
}

func BlendEquationi(buf uint, mode GLenum) {
	C.glBlendEquationi(C.GLuint(buf), C.GLenum(mode))
}

func BlendFuncSeparatei(buf uint, srcRGB GLenum, dstRGB GLenum, srcAlpha GLenum,
	dstAlpha GLenum) {
	C.glBlendFuncSeparatei(C.GLuint(buf), C.GLenum(srcRGB), C.GLenum(dstRGB),
		C.GLenum(srcAlpha), C.GLenum(dstAlpha))
}

func BlendFunci(buf uint, src GLenum, dst GLenum) {
	C.glBlendFunci(C.GLuint(buf), C.GLenum(src), C.GLenum(dst))
}

func MinSampleShading(value float32) {
	C.glMinSampleShading(C.GLclampf(value))
}
