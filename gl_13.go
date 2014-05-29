// Copyright 2012 The go-gl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gl

// #include "gl.h"
import "C"

const (
	ACTIVE_TEXTURE                 = C.GL_ACTIVE_TEXTURE
	ADD_SIGNED                     = C.GL_ADD_SIGNED
	CLAMP_TO_BORDER                = C.GL_CLAMP_TO_BORDER
	CLIENT_ACTIVE_TEXTURE          = C.GL_CLIENT_ACTIVE_TEXTURE
	COMBINE                        = C.GL_COMBINE
	COMBINE_ALPHA                  = C.GL_COMBINE_ALPHA
	COMBINE_RGB                    = C.GL_COMBINE_RGB
	COMPRESSED_ALPHA               = C.GL_COMPRESSED_ALPHA
	COMPRESSED_INTENSITY           = C.GL_COMPRESSED_INTENSITY
	COMPRESSED_LUMINANCE           = C.GL_COMPRESSED_LUMINANCE
	COMPRESSED_LUMINANCE_ALPHA     = C.GL_COMPRESSED_LUMINANCE_ALPHA
	COMPRESSED_RGB                 = C.GL_COMPRESSED_RGB
	COMPRESSED_RGBA                = C.GL_COMPRESSED_RGBA
	COMPRESSED_TEXTURE_FORMATS     = C.GL_COMPRESSED_TEXTURE_FORMATS
	CONSTANT                       = C.GL_CONSTANT
	DOT3_RGB                       = C.GL_DOT3_RGB
	DOT3_RGBA                      = C.GL_DOT3_RGBA
	INTERPOLATE                    = C.GL_INTERPOLATE
	MAX_CUBE_MAP_TEXTURE_SIZE      = C.GL_MAX_CUBE_MAP_TEXTURE_SIZE
	MAX_TEXTURE_UNITS              = C.GL_MAX_TEXTURE_UNITS
	MULTISAMPLE                    = C.GL_MULTISAMPLE
	MULTISAMPLE_BIT                = C.GL_MULTISAMPLE_BIT
	NORMAL_MAP                     = C.GL_NORMAL_MAP
	NUM_COMPRESSED_TEXTURE_FORMATS = C.GL_NUM_COMPRESSED_TEXTURE_FORMATS
	OPERAND0_ALPHA                 = C.GL_OPERAND0_ALPHA
	OPERAND0_RGB                   = C.GL_OPERAND0_RGB
	OPERAND1_ALPHA                 = C.GL_OPERAND1_ALPHA
	OPERAND1_RGB                   = C.GL_OPERAND1_RGB
	OPERAND2_ALPHA                 = C.GL_OPERAND2_ALPHA
	OPERAND2_RGB                   = C.GL_OPERAND2_RGB
	PREVIOUS                       = C.GL_PREVIOUS
	PRIMARY_COLOR                  = C.GL_PRIMARY_COLOR
	PROXY_TEXTURE_CUBE_MAP         = C.GL_PROXY_TEXTURE_CUBE_MAP
	REFLECTION_MAP                 = C.GL_REFLECTION_MAP
	RGB_SCALE                      = C.GL_RGB_SCALE
	SAMPLE_ALPHA_TO_COVERAGE       = C.GL_SAMPLE_ALPHA_TO_COVERAGE
	SAMPLE_ALPHA_TO_ONE            = C.GL_SAMPLE_ALPHA_TO_ONE
	SAMPLE_BUFFERS                 = C.GL_SAMPLE_BUFFERS
	SAMPLE_COVERAGE                = C.GL_SAMPLE_COVERAGE
	SAMPLE_COVERAGE_INVERT         = C.GL_SAMPLE_COVERAGE_INVERT
	SAMPLE_COVERAGE_VALUE          = C.GL_SAMPLE_COVERAGE_VALUE
	SAMPLES                        = C.GL_SAMPLES
	SOURCE0_ALPHA                  = C.GL_SOURCE0_ALPHA
	SOURCE0_RGB                    = C.GL_SOURCE0_RGB
	SOURCE1_ALPHA                  = C.GL_SOURCE1_ALPHA
	SOURCE1_RGB                    = C.GL_SOURCE1_RGB
	SOURCE2_ALPHA                  = C.GL_SOURCE2_ALPHA
	SOURCE2_RGB                    = C.GL_SOURCE2_RGB
	SUBTRACT                       = C.GL_SUBTRACT
	TEXTURE0                       = C.GL_TEXTURE0
	TEXTURE1                       = C.GL_TEXTURE1
	TEXTURE10                      = C.GL_TEXTURE10
	TEXTURE11                      = C.GL_TEXTURE11
	TEXTURE12                      = C.GL_TEXTURE12
	TEXTURE13                      = C.GL_TEXTURE13
	TEXTURE14                      = C.GL_TEXTURE14
	TEXTURE15                      = C.GL_TEXTURE15
	TEXTURE16                      = C.GL_TEXTURE16
	TEXTURE17                      = C.GL_TEXTURE17
	TEXTURE18                      = C.GL_TEXTURE18
	TEXTURE19                      = C.GL_TEXTURE19
	TEXTURE2                       = C.GL_TEXTURE2
	TEXTURE20                      = C.GL_TEXTURE20
	TEXTURE21                      = C.GL_TEXTURE21
	TEXTURE22                      = C.GL_TEXTURE22
	TEXTURE23                      = C.GL_TEXTURE23
	TEXTURE24                      = C.GL_TEXTURE24
	TEXTURE25                      = C.GL_TEXTURE25
	TEXTURE26                      = C.GL_TEXTURE26
	TEXTURE27                      = C.GL_TEXTURE27
	TEXTURE28                      = C.GL_TEXTURE28
	TEXTURE29                      = C.GL_TEXTURE29
	TEXTURE3                       = C.GL_TEXTURE3
	TEXTURE30                      = C.GL_TEXTURE30
	TEXTURE31                      = C.GL_TEXTURE31
	TEXTURE4                       = C.GL_TEXTURE4
	TEXTURE5                       = C.GL_TEXTURE5
	TEXTURE6                       = C.GL_TEXTURE6
	TEXTURE7                       = C.GL_TEXTURE7
	TEXTURE8                       = C.GL_TEXTURE8
	TEXTURE9                       = C.GL_TEXTURE9
	TEXTURE_BINDING_CUBE_MAP       = C.GL_TEXTURE_BINDING_CUBE_MAP
	TEXTURE_COMPRESSED             = C.GL_TEXTURE_COMPRESSED
	TEXTURE_COMPRESSED_IMAGE_SIZE  = C.GL_TEXTURE_COMPRESSED_IMAGE_SIZE
	TEXTURE_COMPRESSION_HINT       = C.GL_TEXTURE_COMPRESSION_HINT
	TEXTURE_CUBE_MAP               = C.GL_TEXTURE_CUBE_MAP
	TEXTURE_CUBE_MAP_NEGATIVE_X    = C.GL_TEXTURE_CUBE_MAP_NEGATIVE_X
	TEXTURE_CUBE_MAP_NEGATIVE_Y    = C.GL_TEXTURE_CUBE_MAP_NEGATIVE_Y
	TEXTURE_CUBE_MAP_NEGATIVE_Z    = C.GL_TEXTURE_CUBE_MAP_NEGATIVE_Z
	TEXTURE_CUBE_MAP_POSITIVE_X    = C.GL_TEXTURE_CUBE_MAP_POSITIVE_X
	TEXTURE_CUBE_MAP_POSITIVE_Y    = C.GL_TEXTURE_CUBE_MAP_POSITIVE_Y
	TEXTURE_CUBE_MAP_POSITIVE_Z    = C.GL_TEXTURE_CUBE_MAP_POSITIVE_Z
	TRANSPOSE_COLOR_MATRIX         = C.GL_TRANSPOSE_COLOR_MATRIX
	TRANSPOSE_MODELVIEW_MATRIX     = C.GL_TRANSPOSE_MODELVIEW_MATRIX
	TRANSPOSE_PROJECTION_MATRIX    = C.GL_TRANSPOSE_PROJECTION_MATRIX
	TRANSPOSE_TEXTURE_MATRIX       = C.GL_TRANSPOSE_TEXTURE_MATRIX
)

func ActiveTexture(texture GLenum) {
	C.glActiveTexture(C.GLenum(texture))
}

func ClientActiveTexture(texture GLenum) {
	C.glClientActiveTexture(C.GLenum(texture))
}

func CompressedTexImage1D(target GLenum, level int, internalformat GLenum, width int, border int, imageSize int, data interface{}) {
	C.glCompressedTexImage1D(C.GLenum(target), C.GLint(level), C.GLenum(internalformat), C.GLsizei(width), C.GLint(border), C.GLsizei(imageSize), glPointer(data))
}

func CompressedTexImage2D(target GLenum, level int, internalformat GLenum, width int, height int, border int, imageSize int, data interface{}) {
	C.glCompressedTexImage2D(C.GLenum(target), C.GLint(level), C.GLenum(internalformat), C.GLsizei(width), C.GLsizei(height), C.GLint(border), C.GLsizei(imageSize), glPointer(data))
}

func CompressedTexImage3D(target GLenum, level int, internalformat GLenum, width int, height int, depth int, border int, imageSize int, data interface{}) {
	C.glCompressedTexImage3D(C.GLenum(target), C.GLint(level), C.GLenum(internalformat), C.GLsizei(width), C.GLsizei(height), C.GLsizei(depth), C.GLint(border), C.GLsizei(imageSize), glPointer(data))
}

func CompressedTexSubImage1D(target GLenum, level int, xoffset int, width int, format GLenum, imageSize int, data interface{}) {
	C.glCompressedTexSubImage1D(C.GLenum(target), C.GLint(level), C.GLint(xoffset), C.GLsizei(width), C.GLenum(format), C.GLsizei(imageSize), glPointer(data))
}

func CompressedTexSubImage2D(target GLenum, level int, xoffset int, yoffset int, width int, height int, format GLenum, imageSize int, data interface{}) {
	C.glCompressedTexSubImage2D(C.GLenum(target), C.GLint(level), C.GLint(xoffset), C.GLint(yoffset), C.GLsizei(width), C.GLsizei(height), C.GLenum(format), C.GLsizei(imageSize), glPointer(data))
}

func CompressedTexSubImage3D(target GLenum, level int, xoffset int, yoffset int, zoffset int, width int, height int, depth int, format GLenum, imageSize int, data interface{}) {
	C.glCompressedTexSubImage3D(C.GLenum(target), C.GLint(level), C.GLint(xoffset), C.GLint(yoffset), C.GLint(zoffset), C.GLsizei(width), C.GLsizei(height), C.GLsizei(depth), C.GLenum(format), C.GLsizei(imageSize), glPointer(data))
}

func GetCompressedTexImage(target GLenum, lod int, img interface{}) {
	C.glGetCompressedTexImage(C.GLenum(target), C.GLint(lod), glPointer(img))
}

func LoadTransposeMatrixd(m []float64) {
	if len(m) != 16 {
		panic("Matrix must contain 16 elements")
	}
	C.glLoadTransposeMatrixd((*C.GLdouble)(&m[0]))
}

func LoadTransposeMatrixf(m []float32) {
	if len(m) != 16 {
		panic("Matrix must contain 16 elements")
	}
	C.glLoadTransposeMatrixf((*C.GLfloat)(&m[0]))
}

func MultTransposeMatrixd(m []float64) {
	if len(m) != 16 {
		panic("Matrix must contain 16 elements")
	}
	C.glMultTransposeMatrixd((*C.GLdouble)(&m[0]))
}

func MultTransposeMatrixf(m []float32) {
	if len(m) != 16 {
		panic("Matrix must contain 16 elements")
	}
	C.glMultTransposeMatrixf((*C.GLfloat)(&m[0]))
}

func MultiTexCoord1d(target GLenum, s float64) {
	C.glMultiTexCoord1d(C.GLenum(target), C.GLdouble(s))
}

func MultiTexCoord1dv(target GLenum, v []float64) {
	C.glMultiTexCoord1dv(C.GLenum(target), (*C.GLdouble)(&v[0]))
}

func MultiTexCoord1f(target GLenum, s float32) {
	C.glMultiTexCoord1f(C.GLenum(target), C.GLfloat(s))
}

func MultiTexCoord1fv(target GLenum, v []float32) {
	C.glMultiTexCoord1fv(C.GLenum(target), (*C.GLfloat)(&v[0]))
}

func MultiTexCoord1i(target GLenum, s int) {
	C.glMultiTexCoord1i(C.GLenum(target), C.GLint(s))
}

func MultiTexCoord1iv(target GLenum, v []int32) {
	C.glMultiTexCoord1iv(C.GLenum(target), (*C.GLint)(&v[0]))
}

func MultiTexCoord1s(target GLenum, s int16) {
	C.glMultiTexCoord1s(C.GLenum(target), C.GLshort(s))
}

func MultiTexCoord1sv(target GLenum, v []int16) {
	C.glMultiTexCoord1sv(C.GLenum(target), (*C.GLshort)(&v[0]))
}

func MultiTexCoord2d(target GLenum, s float64, t float64) {
	C.glMultiTexCoord2d(C.GLenum(target), C.GLdouble(s), C.GLdouble(t))
}

func MultiTexCoord2dv(target GLenum, v []float64) {
	C.glMultiTexCoord2dv(C.GLenum(target), (*C.GLdouble)(&v[0]))
}

func MultiTexCoord2f(target GLenum, s float32, t float32) {
	C.glMultiTexCoord2f(C.GLenum(target), C.GLfloat(s), C.GLfloat(t))
}

func MultiTexCoord2fv(target GLenum, v []float32) {
	C.glMultiTexCoord2fv(C.GLenum(target), (*C.GLfloat)(&v[0]))
}

func MultiTexCoord2i(target GLenum, s int, t int) {
	C.glMultiTexCoord2i(C.GLenum(target), C.GLint(s), C.GLint(t))
}

func MultiTexCoord2iv(target GLenum, v []int32) {
	C.glMultiTexCoord2iv(C.GLenum(target), (*C.GLint)(&v[0]))
}

func MultiTexCoord2s(target GLenum, s int16, t int16) {
	C.glMultiTexCoord2s(C.GLenum(target), C.GLshort(s), C.GLshort(t))
}

func MultiTexCoord2sv(target GLenum, v []int16) {
	C.glMultiTexCoord2sv(C.GLenum(target), (*C.GLshort)(&v[0]))
}

func MultiTexCoord3d(target GLenum, s float64, t float64, r float64) {
	C.glMultiTexCoord3d(C.GLenum(target), C.GLdouble(s), C.GLdouble(t), C.GLdouble(r))
}

func MultiTexCoord3dv(target GLenum, v []float64) {
	C.glMultiTexCoord3dv(C.GLenum(target), (*C.GLdouble)(&v[0]))
}

func MultiTexCoord3f(target GLenum, s float32, t float32, r float32) {
	C.glMultiTexCoord3f(C.GLenum(target), C.GLfloat(s), C.GLfloat(t), C.GLfloat(r))
}

func MultiTexCoord3fv(target GLenum, v []float32) {
	C.glMultiTexCoord3fv(C.GLenum(target), (*C.GLfloat)(&v[0]))
}

func MultiTexCoord3i(target GLenum, s int, t int, r int) {
	C.glMultiTexCoord3i(C.GLenum(target), C.GLint(s), C.GLint(t), C.GLint(r))
}

func MultiTexCoord3iv(target GLenum, v []int32) {
	C.glMultiTexCoord3iv(C.GLenum(target), (*C.GLint)(&v[0]))
}

func MultiTexCoord3s(target GLenum, s int16, t int16, r int16) {
	C.glMultiTexCoord3s(C.GLenum(target), C.GLshort(s), C.GLshort(t), C.GLshort(r))
}

func MultiTexCoord3sv(target GLenum, v []int16) {
	C.glMultiTexCoord3sv(C.GLenum(target), (*C.GLshort)(&v[0]))
}

func MultiTexCoord4d(target GLenum, s float64, t float64, r float64, q float64) {
	C.glMultiTexCoord4d(C.GLenum(target), C.GLdouble(s), C.GLdouble(t), C.GLdouble(r), C.GLdouble(q))
}

func MultiTexCoord4dv(target GLenum, v []float64) {
	C.glMultiTexCoord4dv(C.GLenum(target), (*C.GLdouble)(&v[0]))
}

func MultiTexCoord4f(target GLenum, s float32, t float32, r float32, q float32) {
	C.glMultiTexCoord4f(C.GLenum(target), C.GLfloat(s), C.GLfloat(t), C.GLfloat(r), C.GLfloat(q))
}

func MultiTexCoord4fv(target GLenum, v []float32) {
	C.glMultiTexCoord4fv(C.GLenum(target), (*C.GLfloat)(&v[0]))
}

func MultiTexCoord4i(target GLenum, s int, t int, r int, q int) {
	C.glMultiTexCoord4i(C.GLenum(target), C.GLint(s), C.GLint(t), C.GLint(r), C.GLint(q))
}

func MultiTexCoord4iv(target GLenum, v []int32) {
	C.glMultiTexCoord4iv(C.GLenum(target), (*C.GLint)(&v[0]))
}

func MultiTexCoord4s(target GLenum, s int16, t int16, r int16, q int16) {
	C.glMultiTexCoord4s(C.GLenum(target), C.GLshort(s), C.GLshort(t), C.GLshort(r), C.GLshort(q))
}

func MultiTexCoord4sv(target GLenum, v []int16) {
	C.glMultiTexCoord4sv(C.GLenum(target), (*C.GLshort)(&v[0]))
}

func SampleCoverage(value float32, invert bool) {
	C.glSampleCoverage(C.GLclampf(value), glBool(invert))
}
