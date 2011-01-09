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
import "gl/gl20"

func Build2DMipmaps(target gl.GLenum, internalFormat int, width, height int, format, kind gl.GLenum, data unsafe.Pointer) int {
	return int(C.gluBuild2DMipmaps(
		C.GLenum(target),
		C.GLint(internalFormat),
		C.GLsizei(width),
		C.GLsizei(height),
		C.GLenum(format),
		C.GLenum(kind),
		data,
	))
}

func Perspective(fovy, aspect, zNear, zFar float) {
	C.gluPerspective(
		C.GLdouble(fovy),
		C.GLdouble(aspect),
		C.GLdouble(zNear),
		C.GLdouble(zFar),
	)
}

func LookAt(eyeX, eyeY, eyeZ, centerX, centerY, centerZ, upX, upY, upZ float) {
	C.gluLookAt(
		C.GLdouble(eyeX),
		C.GLdouble(eyeY),
		C.GLdouble(eyeZ),
		C.GLdouble(centerX),
		C.GLdouble(centerY),
		C.GLdouble(centerZ),
		C.GLdouble(upX),
		C.GLdouble(upY),
		C.GLdouble(upZ),
	)
}
