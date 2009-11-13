// godefs -g sdl sdl.c

// MACHINE GENERATED - DO NOT EDIT.

package sdl

// Constants
const (
	INIT_AUDIO = 0x10;
	INIT_VIDEO = 0x20;
	INIT_CDROM = 0x100;
	INIT_JOYSTICK = 0x200;
	INIT_NOPARACHUTE = 0x100000;
	INIT_EVENTTHREAD = 0x1000000;
	INIT_EVERYTHING = 0xffff;
	SWSURFACE = 0;
	HWSURFACE = 0x1;
	ASYNCBLIT = 0x4;
	ANYFORMAT = 0x10000000;
	HWPALETTE = 0x20000000;
	DOUBLEBUF = 0x40000000;
	FULLSCREEN = 0x80000000;
	OPENGL = 0x2;
	OPENGLBLIT = 0xa;
	RESIZABLE = 0x10;
	NOFRAME = 0x20;
	HWACCEL = 0x100;
	SRCCOLORKEY = 0x1000;
	RLEACCELOK = 0x2000;
	RLEACCEL = 0x4000;
	SRCALPHA = 0x10000;
	PREALLOC = 0x1000000;
	YV12_OVERLAY = 0x32315659;
	IYUV_OVERLAY = 0x56555949;
	YUY2_OVERLAY = 0x32595559;
	UYVY_OVERLAY = 0x59565955;
	YVYU_OVERLAY = 0x55595659;
	LOGPAL = 0x1;
	PHYSPAL = 0x2;
	NOEVENT = 0;
	ACTIVEEVENT = 0x1;
	KEYDOWN = 0x2;
	KEYUP = 0x3;
	MOUSEMOTION = 0x4;
	MOUSEBUTTONDOWN = 0x5;
	MOUSEBUTTONUP = 0x6;
	JOYAXISMOTION = 0x7;
	JOYBALLMOTION = 0x8;
	JOYHATMOTION = 0x9;
	JOYBUTTONDOWN = 0xa;
	JOYBUTTONUP = 0xb;
	QUIT = 0xc;
	SYSWMEVENT = 0xd;
	EVENT_RESERVEDA = 0xe;
	EVENT_RESERVEDB = 0xf;
	VIDEORESIZE = 0x10;
	VIDEOEXPOSE = 0x11;
	EVENT_RESERVED2 = 0x12;
	EVENT_RESERVED3 = 0x13;
	EVENT_RESERVED4 = 0x14;
	EVENT_RESERVED5 = 0x15;
	EVENT_RESERVED6 = 0x16;
	EVENT_RESERVED7 = 0x17;
	USEREVENT = 0x18;
	NUMEVENTS = 0x20;
	ACTIVEEVENTMASK = 0x2;
	KEYDOWNMASK = 0x4;
	KEYUPMASK = 0x8;
	KEYEVENTMASK = 0xc;
	MOUSEMOTIONMASK = 0x10;
	MOUSEBUTTONDOWNMASK = 0x20;
	MOUSEBUTTONUPMASK = 0x40;
	MOUSEEVENTMASK = 0x70;
	JOYAXISMOTIONMASK = 0x80;
	JOYBALLMOTIONMASK = 0x100;
	JOYHATMOTIONMASK = 0x200;
	JOYBUTTONDOWNMASK = 0x400;
	JOYBUTTONUPMASK = 0x800;
	JOYEVENTMASK = 0xf80;
	VIDEORESIZEMASK = 0x10000;
	VIDEOEXPOSEMASK = 0x20000;
	QUITMASK = 0x1000;
	SYSWMEVENTMASK = 0x2000;
)

// Types

type Surface struct {
	Flags uint32;
	Format *PixelFormat;
	W int32;
	H int32;
	Pitch uint16;
	Pad0 [2]byte;
	Pixels *byte;
	Offset int32;
	Hwdata *[0]byte /* sprivate_hwdata */;
	Clip_rect Rect;
	Unused1 uint32;
	Locked uint32;
	Map *[0]byte /* sSDL_BlitMap */;
	Format_version uint32;
	Refcount int32;
}

type PixelFormat struct {
	Palette *Palette;
	BitsPerPixel uint8;
	BytesPerPixel uint8;
	Rloss uint8;
	Gloss uint8;
	Bloss uint8;
	Aloss uint8;
	Rshift uint8;
	Gshift uint8;
	Bshift uint8;
	Ashift uint8;
	Pad0 [2]byte;
	Rmask uint32;
	Gmask uint32;
	Bmask uint32;
	Amask uint32;
	Colorkey uint32;
	Alpha uint8;
	Pad1 [3]byte;
}

type Rect struct {
	X int16;
	Y int16;
	W uint16;
	H uint16;
}

type Color struct {
	R uint8;
	G uint8;
	B uint8;
	Unused uint8;
}

type Palette struct {
	Ncolors int32;
	Colors *Color;
}

type VideoInfo struct {
	Pad0 [2]byte;
	UnusedBits3 uint16;
	Video_mem uint32;
	Vfmt *PixelFormat;
	Current_w int32;
	Current_h int32;
}

type Overlay struct {
	Format uint32;
	W int32;
	H int32;
	Planes int32;
	Pitches *uint16;
	Pixels **uint8;
	Hwfuncs *[0]byte /* sprivate_yuvhwfuncs */;
	Hwdata *[0]byte /* sprivate_yuvhwdata */;
	Pad0 [4]byte;
}

type ActiveEvent struct {
	Type uint8;
	Gain uint8;
	State uint8;
}

type KeyboardEvent struct {
	Type uint8;
	Which uint8;
	State uint8;
	Pad0 [1]byte;
	Keysym Keysym;
}

type MouseMotionEvent struct {
	Type uint8;
	Which uint8;
	State uint8;
	Pad0 [1]byte;
	X uint16;
	Y uint16;
	Xrel int16;
	Yrel int16;
}

type MouseButtonEvent struct {
	Type uint8;
	Which uint8;
	Button uint8;
	State uint8;
	X uint16;
	Y uint16;
}

type JoyAxisEvent struct {
	Type uint8;
	Which uint8;
	Axis uint8;
	Pad0 [1]byte;
	Value int16;
}

type JoyBallEvent struct {
	Type uint8;
	Which uint8;
	Ball uint8;
	Pad0 [1]byte;
	Xrel int16;
	Yrel int16;
}

type JoyHatEvent struct {
	Type uint8;
	Which uint8;
	Hat uint8;
	Value uint8;
}

type JoyButtonEvent struct {
	Type uint8;
	Which uint8;
	Button uint8;
	State uint8;
}

type ResizeEvent struct {
	Type uint8;
	Pad0 [3]byte;
	W int32;
	H int32;
}

type ExposeEvent struct {
	Type uint8;
}

type QuitEvent struct {
	Type uint8;
}

type UserEvent struct {
	Type uint8;
	Pad0 [3]byte;
	Code int32;
	Data1 *byte;
	Data2 *byte;
}

type SysWMmsg struct {
}

type SysWMEvent struct {
	Type uint8;
	Pad0 [3]byte;
	Msg *SysWMmsg;
}

type Event struct {
	Type uint8;
	Pad0 [19]byte;
}

type Keysym struct {
	Scancode uint8;
	Pad0 [3]byte;
	Sym byte;
	Mod byte;
	Unicode uint16;
	Pad1 [2]byte;
}
