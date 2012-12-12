// Copyright 2012 The go-gl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gl

// #include "gl.h"
import "C"

// Begin transform feedback with primitive type 'mode'
func BeginTransformFeedback(mode GLenum) {
	C.glBeginTransformFeedback(C.GLenum(mode))
}

// End transform feedback
func EndTransformFeedback() {
	C.glEndTransformFeedback()
}
