// Copyright 2012 The go-gl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gl

// #include "gl.h"
import "C"
import "unsafe"

const (
	BLEND_DST_ALPHA               = C.GL_BLEND_DST_ALPHA
	BLEND_DST_RGB                 = C.GL_BLEND_DST_RGB
	BLEND_SRC_ALPHA               = C.GL_BLEND_SRC_ALPHA
	BLEND_SRC_RGB                 = C.GL_BLEND_SRC_RGB
	COLOR_SUM                     = C.GL_COLOR_SUM
	COMPARE_R_TO_TEXTURE          = C.GL_COMPARE_R_TO_TEXTURE
	CURRENT_FOG_COORDINATE        = C.GL_CURRENT_FOG_COORDINATE
	CURRENT_SECONDARY_COLOR       = C.GL_CURRENT_SECONDARY_COLOR
	DECR_WRAP                     = C.GL_DECR_WRAP
	DEPTH_COMPONENT16             = C.GL_DEPTH_COMPONENT16
	DEPTH_COMPONENT24             = C.GL_DEPTH_COMPONENT24
	DEPTH_COMPONENT32             = C.GL_DEPTH_COMPONENT32
	DEPTH_TEXTURE_MODE            = C.GL_DEPTH_TEXTURE_MODE
	FOG_COORDINATE                = C.GL_FOG_COORDINATE
	FOG_COORDINATE_ARRAY          = C.GL_FOG_COORDINATE_ARRAY
	FOG_COORDINATE_ARRAY_POINTER  = C.GL_FOG_COORDINATE_ARRAY_POINTER
	FOG_COORDINATE_ARRAY_STRIDE   = C.GL_FOG_COORDINATE_ARRAY_STRIDE
	FOG_COORDINATE_ARRAY_TYPE     = C.GL_FOG_COORDINATE_ARRAY_TYPE
	FOG_COORDINATE_SOURCE         = C.GL_FOG_COORDINATE_SOURCE
	FRAGMENT_DEPTH                = C.GL_FRAGMENT_DEPTH
	GENERATE_MIPMAP               = C.GL_GENERATE_MIPMAP
	GENERATE_MIPMAP_HINT          = C.GL_GENERATE_MIPMAP_HINT
	INCR_WRAP                     = C.GL_INCR_WRAP
	MAX_TEXTURE_LOD_BIAS          = C.GL_MAX_TEXTURE_LOD_BIAS
	MIRRORED_REPEAT               = C.GL_MIRRORED_REPEAT
	POINT_DISTANCE_ATTENUATION    = C.GL_POINT_DISTANCE_ATTENUATION
	POINT_FADE_THRESHOLD_SIZE     = C.GL_POINT_FADE_THRESHOLD_SIZE
	POINT_SIZE_MAX                = C.GL_POINT_SIZE_MAX
	POINT_SIZE_MIN                = C.GL_POINT_SIZE_MIN
	SECONDARY_COLOR_ARRAY         = C.GL_SECONDARY_COLOR_ARRAY
	SECONDARY_COLOR_ARRAY_POINTER = C.GL_SECONDARY_COLOR_ARRAY_POINTER
	SECONDARY_COLOR_ARRAY_SIZE    = C.GL_SECONDARY_COLOR_ARRAY_SIZE
	SECONDARY_COLOR_ARRAY_STRIDE  = C.GL_SECONDARY_COLOR_ARRAY_STRIDE
	SECONDARY_COLOR_ARRAY_TYPE    = C.GL_SECONDARY_COLOR_ARRAY_TYPE
	TEXTURE_COMPARE_FUNC          = C.GL_TEXTURE_COMPARE_FUNC
	TEXTURE_COMPARE_MODE          = C.GL_TEXTURE_COMPARE_MODE
	TEXTURE_DEPTH_SIZE            = C.GL_TEXTURE_DEPTH_SIZE
	TEXTURE_FILTER_CONTROL        = C.GL_TEXTURE_FILTER_CONTROL
	TEXTURE_LOD_BIAS              = C.GL_TEXTURE_LOD_BIAS
)

func BlendColor(red float32, green float32, blue float32, alpha float32) {
	C.glBlendColor(C.GLclampf(red), C.GLclampf(green), C.GLclampf(blue), C.GLclampf(alpha))
}

func BlendEquation(mode GLenum) {
	C.glBlendEquation(C.GLenum(mode))
}

func BlendFuncSeparate(sfactorRGB GLenum, dfactorRGB GLenum, sfactorAlpha GLenum, dfactorAlpha GLenum) {
	C.glBlendFuncSeparate(C.GLenum(sfactorRGB), C.GLenum(dfactorRGB), C.GLenum(sfactorAlpha), C.GLenum(dfactorAlpha))
}

func FogCoordPointer(typ GLenum, stride int, pointer interface{}) {
	C.glFogCoordPointer(C.GLenum(typ), C.GLsizei(stride), glPointer(pointer))
}

func FogCoordd(coord float64) {
	C.glFogCoordd(C.GLdouble(coord))
}

func FogCoorddv(coord []float64) {
	C.glFogCoorddv((*C.GLdouble)(&coord[0]))
}

func FogCoordf(coord float32) {
	C.glFogCoordf(C.GLfloat(coord))
}

func FogCoordfv(coord []float32) {
	C.glFogCoordfv((*C.GLfloat)(&coord[0]))
}

func MultiDrawArrays(mode GLenum, first []int32, count []int32, drawcount int) {
	C.glMultiDrawArrays(C.GLenum(mode), (*C.GLint)(&first[0]), (*C.GLsizei)(&count[0]), C.GLsizei(drawcount))
}

func MultiDrawElements(mode GLenum, count []int32, typ GLenum, indices []interface{}) {
	indicesPtr := make([]unsafe.Pointer, len(indices))
	for i := range indices {
		indicesPtr[i] = glPointer(indices[i])
	}
	C.glMultiDrawElements(C.GLenum(mode), (*C.GLsizei)(&count[0]), C.GLenum(typ), &indicesPtr[0], C.GLsizei(len(indices)))
}

func PointParameterf(pname GLenum, param float32) {
	C.glPointParameterf(C.GLenum(pname), C.GLfloat(param))
}

func PointParameterfv(pname GLenum, params []float32) {
	C.glPointParameterfv(C.GLenum(pname), (*C.GLfloat)(&params[0]))
}

func PointParameteri(pname GLenum, param int) {
	C.glPointParameteri(C.GLenum(pname), C.GLint(param))
}

func PointParameteriv(pname GLenum, params []int32) {
	C.glPointParameteriv(C.GLenum(pname), (*C.GLint)(&params[0]))
}

func SecondaryColor3b(red int8, green int8, blue int8) {
	C.glSecondaryColor3b(C.GLbyte(red), C.GLbyte(green), C.GLbyte(blue))
}

func SecondaryColor3bv(v []int8) {
	C.glSecondaryColor3bv((*C.GLbyte)(&v[0]))
}

func SecondaryColor3d(red float64, green float64, blue float64) {
	C.glSecondaryColor3d(C.GLdouble(red), C.GLdouble(green), C.GLdouble(blue))
}

func SecondaryColor3dv(v []float64) {
	C.glSecondaryColor3dv((*C.GLdouble)(&v[0]))
}

func SecondaryColor3f(red float32, green float32, blue float32) {
	C.glSecondaryColor3f(C.GLfloat(red), C.GLfloat(green), C.GLfloat(blue))
}

func SecondaryColor3fv(v []float32) {
	C.glSecondaryColor3fv((*C.GLfloat)(&v[0]))
}

func SecondaryColor3i(red int, green int, blue int) {
	C.glSecondaryColor3i(C.GLint(red), C.GLint(green), C.GLint(blue))
}

func SecondaryColor3iv(v []int32) {
	C.glSecondaryColor3iv((*C.GLint)(&v[0]))
}

func SecondaryColor3s(red int16, green int16, blue int16) {
	C.glSecondaryColor3s(C.GLshort(red), C.GLshort(green), C.GLshort(blue))
}

func SecondaryColor3sv(v []int16) {
	C.glSecondaryColor3sv((*C.GLshort)(&v[0]))
}

func SecondaryColor3ub(red uint8, green uint8, blue uint8) {
	C.glSecondaryColor3ub(C.GLubyte(red), C.GLubyte(green), C.GLubyte(blue))
}

func SecondaryColor3ubv(v []uint8) {
	C.glSecondaryColor3ubv((*C.GLubyte)(&v[0]))
}

func SecondaryColor3ui(red uint, green uint, blue uint) {
	C.glSecondaryColor3ui(C.GLuint(red), C.GLuint(green), C.GLuint(blue))
}

func SecondaryColor3uiv(v []uint32) {
	C.glSecondaryColor3uiv((*C.GLuint)(&v[0]))
}

func SecondaryColor3us(red uint16, green uint16, blue uint16) {
	C.glSecondaryColor3us(C.GLushort(red), C.GLushort(green), C.GLushort(blue))
}

func SecondaryColor3usv(v []uint16) {
	C.glSecondaryColor3usv((*C.GLushort)(&v[0]))
}

func SecondaryColorPointer(size int, typ GLenum, stride int, pointer interface{}) {
	C.glSecondaryColorPointer(C.GLint(size), C.GLenum(typ), C.GLsizei(stride), glPointer(pointer))
}

func WindowPos2d(x float64, y float64) {
	C.glWindowPos2d(C.GLdouble(x), C.GLdouble(y))
}

func WindowPos2dv(p []float64) {
	C.glWindowPos2dv((*C.GLdouble)(&p[0]))
}

func WindowPos2f(x float32, y float32) {
	C.glWindowPos2f(C.GLfloat(x), C.GLfloat(y))
}

func WindowPos2fv(p []float32) {
	C.glWindowPos2fv((*C.GLfloat)(&p[0]))
}

func WindowPos2i(x int, y int) {
	C.glWindowPos2i(C.GLint(x), C.GLint(y))
}

func WindowPos2iv(p []int32) {
	C.glWindowPos2iv((*C.GLint)(&p[0]))
}

func WindowPos2s(x int16, y int16) {
	C.glWindowPos2s(C.GLshort(x), C.GLshort(y))
}

func WindowPos2sv(p []int16) {
	C.glWindowPos2sv((*C.GLshort)(&p[0]))
}

func WindowPos3d(x float64, y float64, z float64) {
	C.glWindowPos3d(C.GLdouble(x), C.GLdouble(y), C.GLdouble(z))
}

func WindowPos3dv(p []float64) {
	C.glWindowPos3dv((*C.GLdouble)(&p[0]))
}

func WindowPos3f(x float32, y float32, z float32) {
	C.glWindowPos3f(C.GLfloat(x), C.GLfloat(y), C.GLfloat(z))
}

func WindowPos3fv(p []float32) {
	C.glWindowPos3fv((*C.GLfloat)(&p[0]))
}

func WindowPos3i(x int, y int, z int) {
	C.glWindowPos3i(C.GLint(x), C.GLint(y), C.GLint(z))
}

func WindowPos3iv(p []int32) {
	C.glWindowPos3iv((*C.GLint)(&p[0]))
}

func WindowPos3s(x int16, y int16, z int16) {
	C.glWindowPos3s(C.GLshort(x), C.GLshort(y), C.GLshort(z))
}

func WindowPos3sv(p []int16) {
	C.glWindowPos3sv((*C.GLshort)(&p[0]))
}
