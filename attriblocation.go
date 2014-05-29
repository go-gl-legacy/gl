// Copyright 2012 The go-gl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gl

// #include "gl.h"
import "C"

// AttribLocation

func (indx AttribLocation) EnableArray() {
	C.glEnableVertexAttribArray(C.GLuint(indx))
}

func (indx AttribLocation) DisableArray() {
	C.glDisableVertexAttribArray(C.GLuint(indx))
}
