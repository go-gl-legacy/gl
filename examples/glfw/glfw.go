package glfw

// #include <GL/glfw.h>
import "C"

const (
	KEY_SPECIAL = 256
	KEY_ESC     = (KEY_SPECIAL + 1)
	OPENED      = 0x00020001
	WINDOW      = 0x00010001
)

func Init() int {
	return int(C.glfwInit())
}

func OpenWindow(width int, height int, redbits int, greenbits int, bluebits int, alphabits int, depthbits int, stencilbits int, mode int) int {
	return int(C.glfwOpenWindow(C.int(width), C.int(height), C.int(redbits), C.int(greenbits), C.int(bluebits), C.int(alphabits), C.int(depthbits), C.int(stencilbits), C.int(mode)))
}

func SwapBuffers() {
	C.glfwSwapBuffers()
}

func GetWindowParam(param int) int {
	return int(C.glfwGetWindowParam(C.int(param)))
}

func GetKey(param int) int {
	return int(C.glfwGetKey(C.int(param)))
}
