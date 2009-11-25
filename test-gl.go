package main

import "sdl"

import "gl"

func main() {

	sdl.Init(sdl.INIT_VIDEO);

	var screen = sdl.SetVideoMode(640, 480, 32, sdl.OPENGL);

	if screen == nil {
		panic("sdl error")
	}

	if gl.Init() != 0 {
		panic("glew error")
	}

	var running = true;

	for running {

		e := &sdl.Event{};

		for e.Poll() {
			switch e.Type {
			case sdl.QUIT:
				running = false;
				break;
			case sdl.KEYDOWN:
				running = false;
				break;
			}
		}

		gl.Begin(gl.TRIANGLES);
		gl.Vertex3f(0, 0, 0);
		gl.Vertex3f(0, 1, 0);
		gl.Vertex3f(1, 1, 0);
		gl.End();

		sdl.GL_SwapBuffers();
		sdl.Delay(25);
	}

	sdl.Quit();

}
