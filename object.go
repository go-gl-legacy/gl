// Copyright 2012 The go-gl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gl

// #include "gl.h"
import "C"

// Object

func (object Object) IsProgram() bool { return C.glIsProgram(C.GLuint(object)) != 0 }

func (object Object) IsShader() bool { return C.glIsShader(C.GLuint(object)) != 0 }

func (object Object) IsTransformFeedback() bool { return C.glIsTransformFeedback(C.GLuint(object)) != 0 }

func (object Object) IsVertexArray() bool { return C.glIsVertexArray(C.GLuint(object)) != 0 }
