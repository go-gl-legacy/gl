// Copyright 2012 The go-gl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gl

// #include "gl.h"
import "C"

const (
	COMPRESSED_SLUMINANCE          = C.GL_COMPRESSED_SLUMINANCE
	COMPRESSED_SLUMINANCE_ALPHA    = C.GL_COMPRESSED_SLUMINANCE_ALPHA
	COMPRESSED_SRGB                = C.GL_COMPRESSED_SRGB
	COMPRESSED_SRGB_ALPHA          = C.GL_COMPRESSED_SRGB_ALPHA
	CURRENT_RASTER_SECONDARY_COLOR = C.GL_CURRENT_RASTER_SECONDARY_COLOR
	FLOAT_MAT2x3                   = C.GL_FLOAT_MAT2x3
	FLOAT_MAT2x4                   = C.GL_FLOAT_MAT2x4
	FLOAT_MAT3x2                   = C.GL_FLOAT_MAT3x2
	FLOAT_MAT3x4                   = C.GL_FLOAT_MAT3x4
	FLOAT_MAT4x2                   = C.GL_FLOAT_MAT4x2
	FLOAT_MAT4x3                   = C.GL_FLOAT_MAT4x3
	PIXEL_PACK_BUFFER              = C.GL_PIXEL_PACK_BUFFER
	PIXEL_PACK_BUFFER_BINDING      = C.GL_PIXEL_PACK_BUFFER_BINDING
	PIXEL_UNPACK_BUFFER            = C.GL_PIXEL_UNPACK_BUFFER
	PIXEL_UNPACK_BUFFER_BINDING    = C.GL_PIXEL_UNPACK_BUFFER_BINDING
	SLUMINANCE                     = C.GL_SLUMINANCE
	SLUMINANCE8                    = C.GL_SLUMINANCE8
	SLUMINANCE8_ALPHA8             = C.GL_SLUMINANCE8_ALPHA8
	SLUMINANCE_ALPHA               = C.GL_SLUMINANCE_ALPHA
	SRGB                           = C.GL_SRGB
	SRGB8                          = C.GL_SRGB8
	SRGB8_ALPHA8                   = C.GL_SRGB8_ALPHA8
	SRGB_ALPHA                     = C.GL_SRGB_ALPHA
)
