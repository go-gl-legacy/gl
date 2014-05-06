// Copyright 2012 The go-gl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gl

// #include "gl.h"
import "C"

const (
	GL_TEXTURE_SPARSE                         = C.GL_TEXTURE_SPARSE_ARB
	GL_VIRTUAL_PAGE_SIZE_INDEX                = C.GL_VIRTUAL_PAGE_SIZE_INDEX_ARB
	GL_NUM_SPARSE_LEVELS                      = C.GL_NUM_SPARSE_LEVELS_ARB
	GL_NUM_VIRTUAL_PAGE_SIZES                 = C.GL_NUM_VIRTUAL_PAGE_SIZES_ARB
	GL_VIRTUAL_PAGE_SIZE_X                    = C.GL_VIRTUAL_PAGE_SIZE_X_ARB
	GL_VIRTUAL_PAGE_SIZE_Y                    = C.GL_VIRTUAL_PAGE_SIZE_Y_ARB
	GL_VIRTUAL_PAGE_SIZE_Z                    = C.GL_VIRTUAL_PAGE_SIZE_Z_ARB
	GL_MAX_SPARSE_TEXTURE_SIZE                = C.GL_MAX_SPARSE_TEXTURE_SIZE_ARB
	GL_MAX_SPARSE_3D_TEXTURE_SIZE             = C.GL_MAX_SPARSE_3D_TEXTURE_SIZE_ARB
	GL_MAX_SPARSE_ARRAY_TEXTURE_LAYERS        = C.GL_MAX_SPARSE_ARRAY_TEXTURE_LAYERS_ARB
	GL_SPARSE_TEXTURE_FULL_ARRAY_CUBE_MIPMAPS = C.GL_SPARSE_TEXTURE_FULL_ARRAY_CUBE_MIPMAPS_ARB
)

func TexPageCommitment(
	target  GLenum,
	level   int,
	xoffset int,
	yoffset int,
	zoffset int,
	width   uint,
	height  uint,
	depth   uint,
	commit  bool) {
	C.glTexPageCommitmentARB(
		C.GLenum (target),
		C.GLint  (level),
		C.GLint  (xoffset),
		C.GLint  (yoffset),
		C.GLint  (zoffset),
		C.GLsizei(width),
		C.GLsizei(height),
		C.GLsizei(depth),
		glBool   (commit))
}

func TexturePageCommitment(
	texture uint,
	target  GLenum,
	level   int,
	xoffset int,
	yoffset int,
	zoffset int,
	width   uint,
	height  uint,
	depth   uint,
	commit  bool) {
	C.glTexturePageCommitmentEXT(
		C.GLuint (texture),
		C.GLenum (target),
		C.GLint  (level),
		C.GLint  (xoffset),
		C.GLint  (yoffset),
		C.GLint  (zoffset),
		C.GLsizei(width),
		C.GLsizei(height),
		C.GLsizei(depth),
		glBool   (commit))
}
