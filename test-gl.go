package main

import "sdl"
//import "math"

import "fmt"
import "gl"

func load_shader(vertex string,fragment string) gl.GLuint
{

    var program=gl.CreateProgram();

    var fs=gl.CreateShader(gl.FRAGMENT_SHADER);
    gl.ShaderSource(fs,[]string{fragment});
    gl.CompileShader(fs);
    gl.AttachShader(program,fs);

    var vs=gl.CreateShader(gl.VERTEX_SHADER);
    gl.ShaderSource(vs,[]string{vertex});
    gl.CompileShader(vs);
    gl.AttachShader(program,vs);

    gl.LinkProgram(program);

    println(gl.GetProgramInfoLog(program));

    return program;

}

func main() {

	sdl.Init(sdl.INIT_VIDEO);

	sdl.SetVideoMode(640, 480, 32, sdl.OPENGL);

	if gl.Init() != 0 {
		panic("glew error")
	}

    println(gl.GetString(gl.RENDERER));

    var program=load_shader("void main(){gl_Position=gl_Vertex;}","void main(){gl_FragColor = vec4(1.0, 0.0, 0.0, 1.0);}");

    gl.UseProgram(program);

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
			default:
				fmt.Printf("Event: %08x\n", e.Type)
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
