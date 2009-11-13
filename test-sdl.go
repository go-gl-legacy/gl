package main

import "sdl"
//import "math"

import "fmt"
import "image/png"
import "os"

func main() {
	sdl.Init(sdl.INIT_VIDEO);

	var screen = sdl.SetVideoMode(640, 480, 32, 0);

	var image = sdl.Load("test.png");

	var running = true;

	var x, y int16;

	for running {

		x++;
		y++;

		e := &sdl.Event{};

		for e.Poll() {
			switch e.Type {
			case sdl.QUIT:
				file, err := os.Open("shoot.png", os.O_CREATE|os.O_WRONLY, 0766);
				println(err);
				png.Encode(file, screen);
				file.Close();
				running = false;
				break;
			case sdl.KEYDOWN:
				println(e.Keyboard().Keysym.Sym)
			case sdl.MOUSEBUTTONDOWN:
				println("Click:", e.MouseButton().X, e.MouseButton().Y);
				x = int16(e.MouseButton().X);
				y = int16(e.MouseButton().Y);
			default:
				fmt.Printf("Event: %08x\n", e.Type)
			}
		}

		screen.FillRect(nil, 0x302019);
		screen.Blit(&sdl.Rect{x, y, 0, 0}, image, nil);
		screen.Flip();
		sdl.Delay(25);
	}

	image.Free();
	screen.Free();

	sdl.Quit();

}
