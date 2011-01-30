package glu

// #ifdef __APPLE__
// # include <OpenGL/glu.h>
// #else
// # include <GL/glu.h>
// #endif
//
//
import "C"
import "gl"

func Build2DMipmaps(target gl.GLenum, internalFormat int, width, height int, format gl.GLenum, data interface{}) int {
	t, p := gl.GetGLenumType(data)
	return int(C.gluBuild2DMipmaps(
		C.GLenum(target),
		C.GLint(internalFormat),
		C.GLsizei(width),
		C.GLsizei(height),
		C.GLenum(format),
		C.GLenum(t),
		p,
	))
}

func Perspective(fovy, aspect, zNear, zFar float64) {
	C.gluPerspective(
		C.GLdouble(fovy),
		C.GLdouble(aspect),
		C.GLdouble(zNear),
		C.GLdouble(zFar),
	)
}

func LookAt(eyeX, eyeY, eyeZ, centerX, centerY, centerZ, upX, upY, upZ float64) {
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
