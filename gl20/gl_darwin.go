package gl

// #include <OpenGL/gl.h>
// #include <OpenGL/glext.h>
import "C"
import "unsafe"

//void glTexImage1D (GLenum target, int level, int internalformat, int width, int border, GLenum format, GLenum type, const GLvoid *pixels)
func TexImage1D(target GLenum, level int, internalformat int, width int, border int, format GLenum, type_ GLenum, pixels unsafe.Pointer) {
	C.glTexImage1D(C.GLenum(target), C.GLint(level), C.GLenum(internalformat), C.GLsizei(width), C.GLint(border), C.GLenum(format), C.GLenum(type_), pixels)
}

//void glTexImage2D (GLenum target, int level, int internalformat, int width, int height, int border, GLenum format, GLenum type, const GLvoid *pixels)
func TexImage2D(target GLenum, level int, internalformat int, width int, height int, border int, format GLenum, type_ GLenum, pixels unsafe.Pointer) {
	C.glTexImage2D(C.GLenum(target), C.GLint(level), C.GLenum(internalformat), C.GLsizei(width), C.GLsizei(height), C.GLint(border), C.GLenum(format), C.GLenum(type_), pixels)
}

//void glPixelMapfv (GLenum map, int mapsize, const float *values)
func PixelMapfv(map_ GLenum, mapsize int, values *float32) {
	C.glPixelMapfv(C.GLenum(map_), C.GLint(mapsize), (*C.GLfloat)(values))
}

//void glPixelMapuiv (GLenum map, int mapsize, const uint *values)
func PixelMapuiv(map_ GLenum, mapsize int, values *uint32) {
	C.glPixelMapuiv(C.GLenum(map_), C.GLint(mapsize), (*C.GLuint)(values))
}

//void glPixelMapusv (GLenum map, int mapsize, const uint16 *values)
func PixelMapusv(map_ GLenum, mapsize int, values *uint16) {
	C.glPixelMapusv(C.GLenum(map_), C.GLint(mapsize), (*C.GLushort)(values))
}
