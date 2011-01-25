package gl

// #include <OpenGL/gl.h>
import "C"
import "unsafe"

//void glPixelMapfv (GLenum map, GLsizei mapsize, const GLfloat *values)
func PixelMapfv(map_ GLenum, mapsize GLsizei, values *GLfloat) {
	C.glPixelMapfv(C.GLenum(map_), C.GLint(mapsize), (*C.GLfloat)(values))
}

//void glPixelMapuiv (GLenum map, GLsizei mapsize, const GLuint *values)
func PixelMapuiv(map_ GLenum, mapsize GLsizei, values *GLuint) {
	C.glPixelMapuiv(C.GLenum(map_), C.GLint(mapsize), (*C.GLuint)(values))
}

//void glPixelMapusv (GLenum map, GLsizei mapsize, const GLushort *values)
func PixelMapusv(map_ GLenum, mapsize GLsizei, values *GLushort) {
	C.glPixelMapusv(C.GLenum(map_), C.GLint(mapsize), (*C.GLushort)(values))
}

//void glTexImage1D (GLenum target, GLint level, GLint internalformat, GLsizei width, GLint border, GLenum format, GLenum type, const GLvoid *pixels)
func TexImage1D(target GLenum, level GLint, internalformat GLint, width GLsizei, border GLint, format GLenum, type_ GLenum, pixels unsafe.Pointer) {
	C.glTexImage1D(C.GLenum(target), C.GLint(level), C.GLenum(internalformat), C.GLsizei(width), C.GLint(border), C.GLenum(format), C.GLenum(type_), pixels)
}

//void glTexImage2D (GLenum target, GLint level, GLint internalformat, GLsizei width, GLsizei height, GLint border, GLenum format, GLenum type, const GLvoid *pixels)
func TexImage2D(target GLenum, level GLint, internalformat GLint, width GLsizei, height GLsizei, border GLint, format GLenum, type_ GLenum, pixels unsafe.Pointer) {
	C.glTexImage2D(C.GLenum(target), C.GLint(level), C.GLenum(internalformat), C.GLsizei(width), C.GLsizei(height), C.GLint(border), C.GLenum(format), C.GLenum(type_), pixels)
}