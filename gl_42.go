// Copyright 2012 The go-gl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gl

// #include "gl.h"
import "C"

const (
	COMPRESSED_RGB_BPTC_SIGNED_FLOAT   = C.GL_COMPRESSED_RGB_BPTC_SIGNED_FLOAT
	COMPRESSED_RGB_BPTC_UNSIGNED_FLOAT = C.GL_COMPRESSED_RGB_BPTC_UNSIGNED_FLOAT
	COMPRESSED_RGBA_BPTC_UNORM         = C.GL_COMPRESSED_RGBA_BPTC_UNORM
	COMPRESSED_SRGB_ALPHA_BPTC_UNORM   = C.GL_COMPRESSED_SRGB_ALPHA_BPTC_UNORM
)
