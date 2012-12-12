// Copyright 2012 The go-gl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gl

// #include "gl.h"
import "C"

func ActiveTexture(texture GLenum) { C.glActiveTexture(C.GLenum(texture)) }

// Texture

type Texture Object

// Create single texture object
func GenTexture() Texture {
	var b C.GLuint
	C.glGenTextures(1, &b)
	return Texture(b)
}

// Fill slice with new textures
func GenTextures(textures []Texture) {
	if len(textures) > 0 {
		C.glGenTextures(C.GLsizei(len(textures)), (*C.GLuint)(&textures[0]))
	}
}

// Delete texture object
func (texture Texture) Delete() {
	b := C.GLuint(texture)
	C.glDeleteTextures(1, &b)
}

// Delete all textures in slice
func DeleteTextures(textures []Texture) {
	if len(textures) > 0 {
		C.glDeleteTextures(C.GLsizei(len(textures)), (*C.GLuint)(&textures[0]))
	}
}

// Bind this texture as target
func (texture Texture) Bind(target GLenum) {
	C.glBindTexture(C.GLenum(target), C.GLuint(texture))
}

// Unbind this texture
func (texture Texture) Unbind(target GLenum) {
	C.glBindTexture(C.GLenum(target), 0)
}

//void glTexImage1D (GLenum target, int level, int internalformat, int width, int border, GLenum format, GLenum type, const GLvoid *pixels)
func TexImage1D(target GLenum, level int, internalformat int, width int, border int, format, typ GLenum, pixels interface{}) {
	C.glTexImage1D(C.GLenum(target), C.GLint(level), C.GLint(internalformat),
		C.GLsizei(width), C.GLint(border), C.GLenum(format), C.GLenum(typ),
		ptr(pixels))
}

//void glTexImage2D (GLenum target, int level, int internalformat, int width, int height, int border, GLenum format, GLenum type, const GLvoid *pixels)
func TexImage2D(target GLenum, level int, internalformat int, width int, height int, border int, format, typ GLenum, pixels interface{}) {
	C.glTexImage2D(C.GLenum(target), C.GLint(level), C.GLint(internalformat),
		C.GLsizei(width), C.GLsizei(height), C.GLint(border), C.GLenum(format),
		C.GLenum(typ), ptr(pixels))
}

//void glTexImage3D (GLenum target, int level, int internalformat, int width, int height, int depth, int border, GLenum format, GLenum type, const GLvoid *pixels)
func TexImage3D(target GLenum, level int, internalformat int, width, height, depth int, border int, format, typ GLenum, pixels interface{}) {
	C.glTexImage3D(C.GLenum(target), C.GLint(level), C.GLint(internalformat),
		C.GLsizei(width), C.GLsizei(height), C.GLsizei(depth), C.GLint(border),
		C.GLenum(format), C.GLenum(typ), ptr(pixels))
}

//void glTexBuffer (GLenum target, GLenum internalformat, GLuint buffer)
func TexBuffer(target, internalformat GLenum, buffer Buffer) {
	C.glTexBuffer(C.GLenum(target), C.GLenum(internalformat), C.GLuint(buffer))
}

//void glTexSubImage1D (GLenum target, int level, int xoffset, int width, GLenum format, GLenum type, const GLvoid *pixels)
func TexSubImage1D(target GLenum, level int, xoffset int, width int, format, typ GLenum, pixels interface{}) {
	C.glTexSubImage1D(C.GLenum(target), C.GLint(level), C.GLint(xoffset),
		C.GLsizei(width), C.GLenum(format), C.GLenum(typ), ptr(pixels))
}

//void glTexSubImage2D (GLenum target, int level, int xoffset, int yoffset, int width, int height, GLenum format, GLenum type, const GLvoid *pixels)
func TexSubImage2D(target GLenum, level int, xoffset int, yoffset int, width int, height int, format, typ GLenum, pixels interface{}) {
	C.glTexSubImage2D(C.GLenum(target), C.GLint(level), C.GLint(xoffset),
		C.GLint(yoffset), C.GLsizei(width), C.GLsizei(height), C.GLenum(format),
		C.GLenum(typ), ptr(pixels))
}

//void glTexImage3D (GLenum target, int level, int xoffset, int yoffset, int zoffset, int width, int height, int depth, GLenum format, GLenum type, const GLvoid *pixels)
func TexSubImage3D(target GLenum, level int, xoffset, yoffset, zoffset, width, height, depth int, format, typ GLenum, pixels interface{}) {
	C.glTexSubImage3D(C.GLenum(target), C.GLint(level),
		C.GLint(xoffset), C.GLint(yoffset), C.GLint(zoffset),
		C.GLsizei(width), C.GLsizei(height), C.GLsizei(depth),
		C.GLenum(format), C.GLenum(typ), ptr(pixels))
}

//void glCopyTexImage1D (GLenum target, int level, GLenum internalFormat, int x, int y, int width, int border)
func CopyTexImage1D(target GLenum, level int, internalFormat GLenum, x int, y int, width int, border int) {
	C.glCopyTexImage1D(C.GLenum(target), C.GLint(level), C.GLenum(internalFormat), C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLint(border))
}

//void glCopyTexImage2D (GLenum target, int level, GLenum internalFormat, int x, int y, int width, int height, int border)
func CopyTexImage2D(target GLenum, level int, internalFormat GLenum, x int, y int, width int, height int, border int) {
	C.glCopyTexImage2D(C.GLenum(target), C.GLint(level), C.GLenum(internalFormat), C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height), C.GLint(border))
}

//void glCopyTexSubImage1D (GLenum target, int level, int xoffset, int x, int y, int width)
func CopyTexSubImage1D(target GLenum, level int, xoffset int, x int, y int, width int) {
	C.glCopyTexSubImage1D(C.GLenum(target), C.GLint(level), C.GLint(xoffset), C.GLint(x), C.GLint(y), C.GLsizei(width))
}

//void glCopyTexSubImage2D (GLenum target, int level, int xoffset, int yoffset, int x, int y, int width, int height)
func CopyTexSubImage2D(target GLenum, level int, xoffset int, yoffset int, x int, y int, width int, height int) {
	C.glCopyTexSubImage2D(C.GLenum(target), C.GLint(level), C.GLint(xoffset), C.GLint(yoffset), C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height))
}

// TODO 3D textures

//void glTexParameterf (GLenum target, GLenum pname, float32 param)
func TexParameterf(target GLenum, pname GLenum, param float32) {
	C.glTexParameterf(C.GLenum(target), C.GLenum(pname), C.GLfloat(param))
}

//void glTexParameterfv (GLenum target, GLenum pname, const float *params)
func TexParameterfv(target GLenum, pname GLenum, params []float32) {
	if len(params) == 0 {
		panic("Invalid params slice length")
	}
	C.glTexParameterfv(C.GLenum(target), C.GLenum(pname), (*C.GLfloat)(&params[0]))
}

//void glTexParameteri (GLenum target, GLenum pname, int param)
func TexParameteri(target GLenum, pname GLenum, param int) {
	C.glTexParameteri(C.GLenum(target), C.GLenum(pname), C.GLint(param))
}

//void glTexParameteriv (GLenum target, GLenum pname, const int *params)
func TexParameteriv(target GLenum, pname GLenum, params []int32) {
	if len(params) == 0 {
		panic("Invalid params slice length")
	}
	C.glTexParameteriv(C.GLenum(target), C.GLenum(pname), (*C.GLint)(&params[0]))
}

//void glGetTexImage (GLenum target, int level, GLenum format, GLenum type, GLvoid *pixels)
func GetTexImage(target GLenum, level int, format, typ GLenum, pixels interface{}) {
	C.glGetTexImage(C.GLenum(target), C.GLint(level), C.GLenum(format),
		C.GLenum(typ), ptr(pixels))
}

//void glGetTexLevelParameterfv (GLenum target, int level, GLenum pname, float *params)
func GetTexLevelParameterfv(target GLenum, level int, pname GLenum, params []float32) {
	if len(params) == 0 {
		panic("Invalid params slice length")
	}
	C.glGetTexLevelParameterfv(C.GLenum(target), C.GLint(level), C.GLenum(pname), (*C.GLfloat)(&params[0]))
}

//void glGetTexLevelParameteriv (GLenum target, int level, GLenum pname, int *params)
func GetTexLevelParameteriv(target GLenum, level int, pname GLenum, params []int32) {
	if len(params) == 0 {
		panic("Invalid params slice length")
	}
	C.glGetTexLevelParameteriv(C.GLenum(target), C.GLint(level), C.GLenum(pname), (*C.GLint)(&params[0]))
}

//void glGetTexParameterfv (GLenum target, GLenum pname, float *params)
func GetTexParameterfv(target GLenum, pname GLenum, params []float32) {
	if len(params) == 0 {
		panic("Invalid params slice length")
	}
	C.glGetTexParameterfv(C.GLenum(target), C.GLenum(pname), (*C.GLfloat)(&params[0]))
}

//void glGetTexParameteriv (GLenum target, GLenum pname, int *params)
func GetTexParameteriv(target GLenum, pname GLenum, params []int32) {
	if len(params) == 0 {
		panic("Invalid params slice length")
	}
	C.glGetTexParameteriv(C.GLenum(target), C.GLenum(pname), (*C.GLint)(&params[0]))
}

func GenerateMipmap(target GLenum) {
	C.glGenerateMipmap(C.GLenum(target))
}

