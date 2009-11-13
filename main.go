package main

import "sdl"
//import "math"

import "fmt"
import "gl"

func main() {
	sdl.Init(sdl.INIT_VIDEO);

	var screen = sdl.SetVideoMode(640, 480, 32, sdl.OPENGL);

	var running = true;

	for running
    {

		e := &sdl.Event{};

		for (e.Poll())
        {
            switch(e.Type)
            {
                case sdl.QUIT:
    				running = false;
                    break;
                case sdl.KEYDOWN:
                    running = false;
                    break;
                default:
                    fmt.Printf("Event: %08x\n",e.Type);
            }
		}

        gl.Begin(0);

		screen.Flip();
        sdl.Delay(25);
	}

	sdl.Quit();

}
