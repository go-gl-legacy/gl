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

    vertex_shader := gl.CreateShader(gl.VERTEX_SHADER);
    vertex_shader.Source("aasa");

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

		sdl.GL_SwapBuffers();
		sdl.Delay(25);
	}

	sdl.Quit();

}
