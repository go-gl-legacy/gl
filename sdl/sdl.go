package sdl

// struct private_hwdata{};
// struct SDL_BlitMap{};
// #define map _map
// #include <SDL/SDL.h>
// #include <SDL/SDL_image.h>
import "C"
import "unsafe"

//type Surface C.SDL_Surface;

type cast unsafe.Pointer

//SDL

func GetError() string	{ return C.GoString(C.SDL_GetError()) }

func Init(flags uint32) int	{ return int(C.SDL_Init(C.Uint32(flags))) }
func Quit()			{ C.SDL_Quit() }

//SDL_video

func SetVideoMode(w int, h int, bpp int, flags uint32) *Surface {
	var screen = C.SDL_SetVideoMode(C.int(w), C.int(h), C.int(bpp), C.Uint32(flags));
	return (*Surface)(cast(screen));
}

func GetVideoSurface() *Surface	{ return (*Surface)(cast(C.SDL_GetVideoSurface())) }

func VideoModeOK(width int, height int, bpp int, flags uint32) int {
	return int(C.SDL_VideoModeOK(C.int(width), C.int(height), C.int(bpp), C.Uint32(flags)))
}

func SDL_UpdateRect(screen *Surface, x int32, y int32, w uint32, h uint32) {
	C.SDL_UpdateRect((*C.SDL_Surface)(cast(screen)), C.Sint32(x), C.Sint32(y), C.Uint32(w), C.Uint32(h))
}

func WM_SetCaption(title string, icon string) {
	ctitle := C.CString(title);
	cicon := C.CString(icon);
	C.SDL_WM_SetCaption(ctitle, cicon);
	C.free(unsafe.Pointer(ctitle));
	C.free(unsafe.Pointer(cicon));
}

func GL_SwapBuffers()	{ C.SDL_GL_SwapBuffers() }

func (screen *Surface) Flip() int	{ return int(C.SDL_Flip((*C.SDL_Surface)(cast(screen)))) }

func (screen *Surface) Free()	{ C.SDL_FreeSurface((*C.SDL_Surface)(cast(screen))) }

func (screen *Surface) Lock() int {
	return int(C.SDL_LockSurface((*C.SDL_Surface)(cast(screen))))
}

func (screen *Surface) Unlock() int	{ return int(C.SDL_Flip((*C.SDL_Surface)(cast(screen)))) }

func (dst *Surface) Blit(dstrect *Rect, src *Surface, srcrect *Rect) int {
	var ret = C.SDL_UpperBlit(
		(*C.SDL_Surface)(cast(src)),
		(*C.SDL_Rect)(cast(srcrect)),
		(*C.SDL_Surface)(cast(dst)),
		(*C.SDL_Rect)(cast(dstrect)));

	return int(ret);
}

func (dst *Surface) FillRect(dstrect *Rect, color uint32) int {
	var ret = C.SDL_FillRect(
		(*C.SDL_Surface)(cast(dst)),
		(*C.SDL_Rect)(cast(dstrect)),
		C.Uint32(color));

	return int(ret);
}

func GetRGBA(color uint32, format *PixelFormat, r *uint8, g *uint8, b *uint8, a *uint8) {
	C.SDL_GetRGBA(C.Uint32(color), (*C.SDL_PixelFormat)(cast(format)), (*C.Uint8)(r), (*C.Uint8)(g), (*C.Uint8)(b), (*C.Uint8)(a))
}

//SDL image

func Load(file string) *Surface {
	cfile := C.CString(file);
	var screen = C.IMG_Load(cfile);
	C.free(unsafe.Pointer(cfile));
	return (*Surface)(cast(screen));
}

// SDL keys

func EnableUNICODE(enable int) int	{ return int(C.SDL_EnableUNICODE(C.int(enable))) }

func EnableKeyRepeat(delay int, interval int) int {
	return int(C.SDL_EnableKeyRepeat(C.int(delay), C.int(interval)))
}

func GetKeyRepeat() (int, int) {

	var delay int;
	var interval int;

	C.SDL_GetKeyRepeat((*C.int)(cast(&delay)), (*C.int)(cast(&interval)));

	return delay, interval;

}

// TODO
// Uint8 * SDL_GetKeyState(int *numkeys)
// func GetKeyState(numkeys *int) *uint8
// {
// 	return (*uint8)(C.SDL_GetKeyState((*C.int)(numkeys)));
// }

type Mod C.int
type Key C.int

// SDLMod SDL_GetModState(void)
func GetModState() Mod	{ return Mod(C.SDL_GetModState()) }

// void SDL_SetModState(SDLMod modstate)
func SetModState(modstate Mod)	{ C.SDL_SetModState(C.SDLMod(modstate)) }

// char * SDL_GetKeyName(SDLKey key)
func GetKeyName(key Key) string	{ return C.GoString(C.SDL_GetKeyName(C.SDLKey(key))) }

//SDL events

func (event *Event) Wait() bool {
	var ret = C.SDL_WaitEvent((*C.SDL_Event)(cast(event)));
	return ret != 0;
}

func (event *Event) Poll() bool {
	var ret = C.SDL_PollEvent((*C.SDL_Event)(cast(event)));
	return ret != 0;
}

func (event *Event) Keyboard() *KeyboardEvent {
	if event.Type == KEYUP || event.Type == KEYDOWN {
		return (*KeyboardEvent)(cast(event))
	}

	return nil;
}

func (event *Event) MouseButton() *MouseButtonEvent {
	if event.Type == MOUSEBUTTONDOWN || event.Type == MOUSEBUTTONUP {
		return (*MouseButtonEvent)(cast(event))
	}

	return nil;
}

func (event *Event) MouseMotion() *MouseMotionEvent {
	if event.Type == MOUSEMOTION {
		return (*MouseMotionEvent)(cast(event))
	}

	return nil;
}

//SDL time

func Delay(ms uint32)	{ C.SDL_Delay(C.Uint32(ms)) }
