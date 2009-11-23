package sdl

//implementation of draw.Context and draw.Image interfaces

import "exp/draw"
import "os"
import "image"
import "unsafe"
import "time"

type Context struct {
	screen		*Surface;
	mouse_chan	chan draw.Mouse;
	key_chan	chan int;
	resize_chan	chan bool;
	quit_chan	chan bool;
}

func InitContext(w int, h int) (*Context, os.Error) {

	var this = new(Context);

	Init(INIT_VIDEO);

	this.screen = SetVideoMode(w, h, 32, SWSURFACE);

	this.screen.Lock();

	// TODO \/

	this.resize_chan = make(chan bool, 64);
	this.key_chan = make(chan int, 64);
	this.mouse_chan = make(chan draw.Mouse, 1024);
	this.quit_chan = make(chan bool);

	return this, nil;

}

func (this *Context) Screen() draw.Image	{ return this.screen }

func (this *Context) FlushImage() {

	this.screen.Unlock();
	this.screen.Flip();
	this.screen.Lock();

	e := &Event{};

	for e.Poll() {
		switch e.Type {
		case QUIT:
			this.quit_chan <- true;
			break;
		case KEYDOWN:
			this.key_chan <- int(e.Keyboard().Keysym.Sym);
			break;
		case MOUSEBUTTONDOWN, MOUSEBUTTONUP:
			m := e.MouseButton();
			this.mouse_chan <- draw.Mouse{int(GetMouseState(nil, nil)), draw.Point{int(m.X), int(m.Y)}, time.Nanoseconds()};
			break;
		case MOUSEMOTION:
			m := e.MouseMotion();
			this.mouse_chan <- draw.Mouse{int(GetMouseState(nil, nil)), draw.Point{int(m.X), int(m.Y)}, time.Nanoseconds()};
			break;
		case VIDEORESIZE:
			this.resize_chan <- true;
			break;
		default:
			break
		}
	}

}

func (this *Context) KeyboardChan() <-chan int {
	//TODO conversion
	return this.key_chan
}


func (this *Context) MouseChan() <-chan draw.Mouse {
	return this.mouse_chan
}

func (this *Context) ResizeChan() <-chan bool	{ return this.resize_chan }
func (this *Context) QuitChan() <-chan bool	{ return this.quit_chan }

// surface --> Image

func (surface *Surface) ColorModel() image.ColorModel {
	//TODO
	return nil
}

func (surface *Surface) Width() int	{ return int(surface.W) }

func (surface *Surface) Height() int	{ return int(surface.H) }

func (surface *Surface) Set(x, y int, c image.Color) {
	//TODO endianess, bpp, alpha, etc

	var bpp = int(surface.Format.BytesPerPixel);

	var pixel = uintptr(unsafe.Pointer(surface.Pixels));

	pixel += uintptr(y*int(surface.Pitch) + x*bpp);

	var p = (*image.RGBAColor)(unsafe.Pointer(pixel));

	var r, g, b, a = c.RGBA();

	p.R = uint8(r);
	p.G = uint8(g);
	p.R = uint8(b);
	p.A = uint8(255 - a);

}


func (surface *Surface) At(x, y int) image.Color {

	var bpp = int(surface.Format.BytesPerPixel);

	var pixel = uintptr(unsafe.Pointer(surface.Pixels));

	pixel += uintptr(y*int(surface.Pitch) + x*bpp);

	var color = *((*uint32)(unsafe.Pointer(pixel)));

	var r uint8;
	var g uint8;
	var b uint8;
	var a uint8;

	GetRGBA(color, surface.Format, &r, &g, &b, &a);

	return image.RGBAColor{uint8(r), uint8(g), uint8(b), uint8(a)};

}
