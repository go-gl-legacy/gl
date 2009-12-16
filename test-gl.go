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
    vertex_shader.Source("void main(void){gl_Position = ftransform();}");
    vertex_shader.Compile();
    print(vertex_shader.GetInfoLog());

    fragment_shader := gl.CreateShader(gl.FRAGMENT_SHADER);
    fragment_shader.Source("void main(void){gl_FragColor = vec4(1.0, 0.0, 0.0, 1.0);}");
    fragment_shader.Compile();
    print(fragment_shader.GetInfoLog());

    program := gl.CreateProgram();
    program.AttachShader(vertex_shader);
    program.AttachShader(fragment_shader);
    program.Link();
    print(program.GetInfoLog());
    program.Use();

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
