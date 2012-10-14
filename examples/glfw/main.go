// Copyright 2012 The go-gl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
this example uses glfw instead of Go-SDL. To try it type:
	
	gomake
	sudo gomake install
	gomake main
	./main

*/

package main

import "glfw"
import "gl"

func main() {

	ret := glfw.Init()

	print(ret)

	ret = glfw.OpenWindow(300, 300, 0, 0, 0, 0, 0, 0, glfw.WINDOW)

	print(ret)

	if gl.Init() != 0 {
		panic("glew error")
	}

	running := true

	for running {

		gl.Begin(gl.TRIANGLES)
		gl.Vertex3f(0, 0, 0)
		gl.Vertex3f(0, 1, 0)
		gl.Vertex3f(1, 1, 0)
		gl.End()

		glfw.SwapBuffers()

		running = glfw.GetKey(glfw.KEY_ESC) == 0 && glfw.GetWindowParam(glfw.OPENED) != 0

	}

}
