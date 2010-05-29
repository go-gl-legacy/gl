package glu

// #ifdef __Darwin
// # include <OpenGL/glu.h>
// #else
// # include <GL/glu.h>
// #endif
//
//
import "C"
import "unsafe"
import "gl"

func Build2DMipmaps(target gl.GLenum, internalFormat gl.GLint, width, height gl.GLsizei, format, kind gl.GLenum, data unsafe.Pointer) gl.GLint {
	return gl.GLint(C.gluBuild2DMipmaps(
		C.GLenum(target),
		C.GLint(internalFormat),
		C.GLsizei(width),
		C.GLsizei(height),
		C.GLenum(format),
		C.GLenum(kind),
		data,
	))
}

func Perspective(fovy, aspect, zNear, zFar gl.GLdouble) {
	C.gluPerspective(
		C.GLdouble(fovy),
		C.GLdouble(aspect),
		C.GLdouble(zNear),
		C.GLdouble(zFar),
	)
}
