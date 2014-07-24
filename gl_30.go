// Copyright 2012 The go-gl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gl

// #include "gl.h"
import "C"
import "unsafe"

const (
	ALPHA_INTEGER                                 = C.GL_ALPHA_INTEGER
	BGR_INTEGER                                   = C.GL_BGR_INTEGER
	BGRA_INTEGER                                  = C.GL_BGRA_INTEGER
	BLUE_INTEGER                                  = C.GL_BLUE_INTEGER
	CLAMP_FRAGMENT_COLOR                          = C.GL_CLAMP_FRAGMENT_COLOR
	CLAMP_READ_COLOR                              = C.GL_CLAMP_READ_COLOR
	CLAMP_VERTEX_COLOR                            = C.GL_CLAMP_VERTEX_COLOR
	CLIP_DISTANCE0                                = C.GL_CLIP_DISTANCE0
	CLIP_DISTANCE1                                = C.GL_CLIP_DISTANCE1
	CLIP_DISTANCE2                                = C.GL_CLIP_DISTANCE2
	CLIP_DISTANCE3                                = C.GL_CLIP_DISTANCE3
	CLIP_DISTANCE4                                = C.GL_CLIP_DISTANCE4
	CLIP_DISTANCE5                                = C.GL_CLIP_DISTANCE5
	COMPARE_REF_TO_TEXTURE                        = C.GL_COMPARE_REF_TO_TEXTURE
	CONTEXT_FLAG_FORWARD_COMPATIBLE_BIT           = C.GL_CONTEXT_FLAG_FORWARD_COMPATIBLE_BIT
	CONTEXT_FLAGS                                 = C.GL_CONTEXT_FLAGS
	DEPTH_BUFFER                                  = C.GL_DEPTH_BUFFER
	FIXED_ONLY                                    = C.GL_FIXED_ONLY
	GREEN_INTEGER                                 = C.GL_GREEN_INTEGER
	INT_SAMPLER_1D                                = C.GL_INT_SAMPLER_1D
	INT_SAMPLER_1D_ARRAY                          = C.GL_INT_SAMPLER_1D_ARRAY
	INT_SAMPLER_2D                                = C.GL_INT_SAMPLER_2D
	INT_SAMPLER_2D_ARRAY                          = C.GL_INT_SAMPLER_2D_ARRAY
	INT_SAMPLER_3D                                = C.GL_INT_SAMPLER_3D
	INT_SAMPLER_CUBE                              = C.GL_INT_SAMPLER_CUBE
	INTERLEAVED_ATTRIBS                           = C.GL_INTERLEAVED_ATTRIBS
	MAJOR_VERSION                                 = C.GL_MAJOR_VERSION
	MAX_ARRAY_TEXTURE_LAYERS                      = C.GL_MAX_ARRAY_TEXTURE_LAYERS
	MAX_CLIP_DISTANCES                            = C.GL_MAX_CLIP_DISTANCES
	MAX_PROGRAM_TEXEL_OFFSET                      = C.GL_MAX_PROGRAM_TEXEL_OFFSET
	MAX_TRANSFORM_FEEDBACK_INTERLEAVED_COMPONENTS = C.GL_MAX_TRANSFORM_FEEDBACK_INTERLEAVED_COMPONENTS
	MAX_TRANSFORM_FEEDBACK_SEPARATE_ATTRIBS       = C.GL_MAX_TRANSFORM_FEEDBACK_SEPARATE_ATTRIBS
	MAX_TRANSFORM_FEEDBACK_SEPARATE_COMPONENTS    = C.GL_MAX_TRANSFORM_FEEDBACK_SEPARATE_COMPONENTS
	MAX_VARYING_COMPONENTS                        = C.GL_MAX_VARYING_COMPONENTS
	MIN_PROGRAM_TEXEL_OFFSET                      = C.GL_MIN_PROGRAM_TEXEL_OFFSET
	MINOR_VERSION                                 = C.GL_MINOR_VERSION
	NUM_EXTENSIONS                                = C.GL_NUM_EXTENSIONS
	PRIMITIVES_GENERATED                          = C.GL_PRIMITIVES_GENERATED
	PROXY_TEXTURE_1D_ARRAY                        = C.GL_PROXY_TEXTURE_1D_ARRAY
	PROXY_TEXTURE_2D_ARRAY                        = C.GL_PROXY_TEXTURE_2D_ARRAY
	QUERY_BY_REGION_NO_WAIT                       = C.GL_QUERY_BY_REGION_NO_WAIT
	QUERY_BY_REGION_WAIT                          = C.GL_QUERY_BY_REGION_WAIT
	QUERY_NO_WAIT                                 = C.GL_QUERY_NO_WAIT
	QUERY_WAIT                                    = C.GL_QUERY_WAIT
	R11F_G11F_B10F                                = C.GL_R11F_G11F_B10F
	RASTERIZER_DISCARD                            = C.GL_RASTERIZER_DISCARD
	RED_INTEGER                                   = C.GL_RED_INTEGER
	RGB16F                                        = C.GL_RGB16F
	RGB16I                                        = C.GL_RGB16I
	RGB16UI                                       = C.GL_RGB16UI
	RGB32F                                        = C.GL_RGB32F
	RGB32I                                        = C.GL_RGB32I
	RGB32UI                                       = C.GL_RGB32UI
	RGB8I                                         = C.GL_RGB8I
	RGB8UI                                        = C.GL_RGB8UI
	RGB9_E5                                       = C.GL_RGB9_E5
	RGB_INTEGER                                   = C.GL_RGB_INTEGER
	RGBA16F                                       = C.GL_RGBA16F
	RGBA16I                                       = C.GL_RGBA16I
	RGBA16UI                                      = C.GL_RGBA16UI
	RGBA32F                                       = C.GL_RGBA32F
	RGBA32I                                       = C.GL_RGBA32I
	RGBA32UI                                      = C.GL_RGBA32UI
	RGBA8I                                        = C.GL_RGBA8I
	RGBA8UI                                       = C.GL_RGBA8UI
	RGBA_INTEGER                                  = C.GL_RGBA_INTEGER
	SAMPLER_1D_ARRAY                              = C.GL_SAMPLER_1D_ARRAY
	SAMPLER_1D_ARRAY_SHADOW                       = C.GL_SAMPLER_1D_ARRAY_SHADOW
	SAMPLER_2D_ARRAY                              = C.GL_SAMPLER_2D_ARRAY
	SAMPLER_2D_ARRAY_SHADOW                       = C.GL_SAMPLER_2D_ARRAY_SHADOW
	SAMPLER_CUBE_SHADOW                           = C.GL_SAMPLER_CUBE_SHADOW
	SEPARATE_ATTRIBS                              = C.GL_SEPARATE_ATTRIBS
	STENCIL_BUFFER                                = C.GL_STENCIL_BUFFER
	TEXTURE_1D_ARRAY                              = C.GL_TEXTURE_1D_ARRAY
	TEXTURE_2D_ARRAY                              = C.GL_TEXTURE_2D_ARRAY
	TEXTURE_ALPHA_TYPE                            = C.GL_TEXTURE_ALPHA_TYPE
	TEXTURE_BINDING_1D_ARRAY                      = C.GL_TEXTURE_BINDING_1D_ARRAY
	TEXTURE_BINDING_2D_ARRAY                      = C.GL_TEXTURE_BINDING_2D_ARRAY
	TEXTURE_BLUE_TYPE                             = C.GL_TEXTURE_BLUE_TYPE
	TEXTURE_DEPTH_TYPE                            = C.GL_TEXTURE_DEPTH_TYPE
	TEXTURE_GREEN_TYPE                            = C.GL_TEXTURE_GREEN_TYPE
	TEXTURE_INTENSITY_TYPE                        = C.GL_TEXTURE_INTENSITY_TYPE
	TEXTURE_LUMINANCE_TYPE                        = C.GL_TEXTURE_LUMINANCE_TYPE
	TEXTURE_RED_TYPE                              = C.GL_TEXTURE_RED_TYPE
	TEXTURE_SHARED_SIZE                           = C.GL_TEXTURE_SHARED_SIZE
	TRANSFORM_FEEDBACK_BUFFER                     = C.GL_TRANSFORM_FEEDBACK_BUFFER
	TRANSFORM_FEEDBACK_BUFFER_BINDING             = C.GL_TRANSFORM_FEEDBACK_BUFFER_BINDING
	TRANSFORM_FEEDBACK_BUFFER_MODE                = C.GL_TRANSFORM_FEEDBACK_BUFFER_MODE
	TRANSFORM_FEEDBACK_BUFFER_SIZE                = C.GL_TRANSFORM_FEEDBACK_BUFFER_SIZE
	TRANSFORM_FEEDBACK_BUFFER_START               = C.GL_TRANSFORM_FEEDBACK_BUFFER_START
	TRANSFORM_FEEDBACK_PRIMITIVES_WRITTEN         = C.GL_TRANSFORM_FEEDBACK_PRIMITIVES_WRITTEN
	TRANSFORM_FEEDBACK_VARYING_MAX_LENGTH         = C.GL_TRANSFORM_FEEDBACK_VARYING_MAX_LENGTH
	TRANSFORM_FEEDBACK_VARYINGS                   = C.GL_TRANSFORM_FEEDBACK_VARYINGS
	UNSIGNED_INT_10F_11F_11F_REV                  = C.GL_UNSIGNED_INT_10F_11F_11F_REV
	UNSIGNED_INT_5_9_9_9_REV                      = C.GL_UNSIGNED_INT_5_9_9_9_REV
	UNSIGNED_INT_SAMPLER_1D                       = C.GL_UNSIGNED_INT_SAMPLER_1D
	UNSIGNED_INT_SAMPLER_1D_ARRAY                 = C.GL_UNSIGNED_INT_SAMPLER_1D_ARRAY
	UNSIGNED_INT_SAMPLER_2D                       = C.GL_UNSIGNED_INT_SAMPLER_2D
	UNSIGNED_INT_SAMPLER_2D_ARRAY                 = C.GL_UNSIGNED_INT_SAMPLER_2D_ARRAY
	UNSIGNED_INT_SAMPLER_3D                       = C.GL_UNSIGNED_INT_SAMPLER_3D
	UNSIGNED_INT_SAMPLER_CUBE                     = C.GL_UNSIGNED_INT_SAMPLER_CUBE
	UNSIGNED_INT_VEC2                             = C.GL_UNSIGNED_INT_VEC2
	UNSIGNED_INT_VEC3                             = C.GL_UNSIGNED_INT_VEC3
	UNSIGNED_INT_VEC4                             = C.GL_UNSIGNED_INT_VEC4
	VERTEX_ATTRIB_ARRAY_INTEGER                   = C.GL_VERTEX_ATTRIB_ARRAY_INTEGER
)

func (query Query) BeginConditionalRender(mode GLenum) {
	C.glBeginConditionalRender(C.GLuint(query), C.GLenum(mode))
}

func BeginTransformFeedback(primitiveMode GLenum) {
	C.glBeginTransformFeedback(C.GLenum(primitiveMode))
}

func (program Program) BindFragDataLocation(colorNumber uint, name string) {
	cname := glString(name)
	defer freeString(cname)
	C.glBindFragDataLocation(C.GLuint(program), C.GLuint(colorNumber), cname)
}

func ClampColor(target GLenum, clamp GLenum) {
	C.glClampColor(C.GLenum(target), C.GLenum(clamp))
}

func ClearBufferfi(buffer GLenum, drawBuffer int, depth float32, stencil int) {
	C.glClearBufferfi(C.GLenum(buffer), C.GLint(drawBuffer), C.GLfloat(depth), C.GLint(stencil))
}

func ClearBufferfv(buffer GLenum, drawBuffer int, value []float32) {
	C.glClearBufferfv(C.GLenum(buffer), C.GLint(drawBuffer), (*C.GLfloat)(&value[0]))
}

func ClearBufferiv(buffer GLenum, drawBuffer int, value []int32) {
	C.glClearBufferiv(C.GLenum(buffer), C.GLint(drawBuffer), (*C.GLint)(&value[0]))
}

func ClearBufferuiv(buffer GLenum, drawBuffer int, value []uint32) {
	C.glClearBufferuiv(C.GLenum(buffer), C.GLint(drawBuffer), (*C.GLuint)(&value[0]))
}

func ColorMaski(buf uint, red bool, green bool, blue bool, alpha bool) {
	C.glColorMaski(C.GLuint(buf), glBool(red), glBool(green), glBool(blue), glBool(alpha))
}

func Disablei(cap GLenum, index uint) {
	C.glDisablei(C.GLenum(cap), C.GLuint(index))
}

func Enablei(cap GLenum, index uint) {
	C.glEnablei(C.GLenum(cap), C.GLuint(index))
}

func (_ Query) EndConditionalRender() {
	C.glEndConditionalRender()
}

func EndTransformFeedback() {
	C.glEndTransformFeedback()
}

func GetBooleani_v(pname GLenum, index uint, data []bool) {
	C.glGetBooleani_v(C.GLenum(pname), C.GLuint(index), (*C.GLboolean)(unsafe.Pointer(&data[0])))
}

// TODO type?
func (program Program) GetFragDataLocation(name string) Object {
	cname := glString(name)
	defer freeString(cname)
	return Object(C.glGetFragDataLocation(C.GLuint(program), cname))
}

func GetStringi(name GLenum, index uint) string {
	s := unsafe.Pointer(C.glGetStringi(C.GLenum(name), C.GLuint(index)))
	return C.GoString((*C.char)(s))
}

func GetTexParameterIiv(target GLenum, pname GLenum, params []int32) {
	C.glGetTexParameterIiv(C.GLenum(target), C.GLenum(pname), (*C.GLint)(&params[0]))
}

func GetTexParameterIuiv(target GLenum, pname GLenum, params []uint32) {
	C.glGetTexParameterIuiv(C.GLenum(target), C.GLenum(pname), (*C.GLuint)(&params[0]))
}

func (program Program) GetTransformFeedbackVarying(index uint) (int, GLenum, string) {
	var typ C.GLenum
	var recvLen C.GLsizei
	var varyingLength C.GLsizei
	length := program.Get(TRANSFORM_FEEDBACK_VARYING_MAX_LENGTH)
	buf := make([]C.GLchar, length)
	C.glGetTransformFeedbackVarying(C.GLuint(program), C.GLuint(index), C.GLsizei(length), &recvLen, &varyingLength, (*C.GLenum)(&typ), &buf[0])
	return int(varyingLength), GLenum(typ), C.GoString((*C.char)(&buf[0]))
}

func (program Program) GetUniformuiv(location UniformLocation, params []uint32) {
	C.glGetUniformuiv(C.GLuint(program), C.GLint(location), (*C.GLuint)(&params[0]))
}

func GetVertexAttribIiv(index uint, pname GLenum, params []int32) {
	C.glGetVertexAttribIiv(C.GLuint(index), C.GLenum(pname), (*C.GLint)(&params[0]))
}

func GetVertexAttribIuiv(index uint, pname GLenum, params []uint32) {
	C.glGetVertexAttribIuiv(C.GLuint(index), C.GLenum(pname), (*C.GLuint)(&params[0]))
}

func IsEnabledi(cap GLenum, index uint) bool {
	return goBool(C.glIsEnabledi(C.GLenum(cap), C.GLuint(index)))
}

func TexParameterIiv(target GLenum, pname GLenum, params []int32) {
	C.glTexParameterIiv(C.GLenum(target), C.GLenum(pname), (*C.GLint)(&params[0]))
}

func TexParameterIuiv(target GLenum, pname GLenum, params []uint32) {
	C.glTexParameterIuiv(C.GLenum(target), C.GLenum(pname), (*C.GLuint)(&params[0]))
}

func (program Program) TransformFeedbackVaryings(varyings []string, bufferMode GLenum) {
	cvaryings := make([]*C.GLchar, len(varyings))
	for i := range varyings {
		cvaryings[i] = glString(varyings[i])
		defer freeString(cvaryings[i])
	}
	C.glTransformFeedbackVaryings(C.GLuint(program), C.GLsizei(len(varyings)), (**C.GLchar)(&cvaryings[0]), C.GLenum(bufferMode))
}

func (location UniformLocation) Uniform1ui(v0 uint) {
	C.glUniform1ui(C.GLint(location), C.GLuint(v0))
}

func (location UniformLocation) Uniform1uiv(values []uint32) {
	C.glUniform1uiv(C.GLint(location), C.GLsizei(len(values)), (*C.GLuint)(&values[0]))
}

func (location UniformLocation) Uniform2ui(v0 uint, v1 uint) {
	C.glUniform2ui(C.GLint(location), C.GLuint(v0), C.GLuint(v1))
}

func (location UniformLocation) Uniform2uiv(values []uint32) {
	C.glUniform2uiv(C.GLint(location), C.GLsizei(len(values)), (*C.GLuint)(&values[0]))
}

func (location UniformLocation) Uniform3ui(v0 uint, v1 uint, v2 uint) {
	C.glUniform3ui(C.GLint(location), C.GLuint(v0), C.GLuint(v1), C.GLuint(v2))
}

func (location UniformLocation) Uniform3uiv(values []uint32) {
	C.glUniform3uiv(C.GLint(location), C.GLsizei(len(values)), (*C.GLuint)(&values[0]))
}

func (location UniformLocation) Uniform4ui(v0 uint, v1 uint, v2 uint, v3 uint) {
	C.glUniform4ui(C.GLint(location), C.GLuint(v0), C.GLuint(v1), C.GLuint(v2), C.GLuint(v3))
}

func (location UniformLocation) Uniform4uiv(values []uint32) {
	C.glUniform4uiv(C.GLint(location), C.GLsizei(len(values)), (*C.GLuint)(&values[0]))
}

func (location AttribLocation) AttribI1i(v0 int) {
	C.glVertexAttribI1i(C.GLuint(location), C.GLint(v0))
}

func (location AttribLocation) AttribI1iv(v []int32) {
	C.glVertexAttribI1iv(C.GLuint(location), (*C.GLint)(&v[0]))
}

func (location AttribLocation) AttribI1ui(v0 uint) {
	C.glVertexAttribI1ui(C.GLuint(location), C.GLuint(v0))
}

func (location AttribLocation) AttribI1uiv(v []uint32) {
	C.glVertexAttribI1uiv(C.GLuint(location), (*C.GLuint)(&v[0]))
}

func (location AttribLocation) AttribI2i(v0 int, v1 int) {
	C.glVertexAttribI2i(C.GLuint(location), C.GLint(v0), C.GLint(v1))
}

func (location AttribLocation) AttribI2iv(v []int32) {
	C.glVertexAttribI2iv(C.GLuint(location), (*C.GLint)(&v[0]))
}

func (location AttribLocation) AttribI2ui(v0 uint, v1 uint) {
	C.glVertexAttribI2ui(C.GLuint(location), C.GLuint(v0), C.GLuint(v1))
}

func (location AttribLocation) AttribI2uiv(v []uint32) {
	C.glVertexAttribI2uiv(C.GLuint(location), (*C.GLuint)(&v[0]))
}

func (location AttribLocation) AttribI3i(v0 int, v1 int, v2 int) {
	C.glVertexAttribI3i(C.GLuint(location), C.GLint(v0), C.GLint(v1), C.GLint(v2))
}

func (location AttribLocation) AttribI3iv(v []int32) {
	C.glVertexAttribI3iv(C.GLuint(location), (*C.GLint)(&v[0]))
}

func (location AttribLocation) AttribI3ui(v0 uint, v1 uint, v2 uint) {
	C.glVertexAttribI3ui(C.GLuint(location), C.GLuint(v0), C.GLuint(v1), C.GLuint(v2))
}

func (location AttribLocation) AttribI3uiv(v []uint32) {
	C.glVertexAttribI3uiv(C.GLuint(location), (*C.GLuint)(&v[0]))
}

func (location AttribLocation) AttribI4bv(v []int8) {
	C.glVertexAttribI4bv(C.GLuint(location), (*C.GLbyte)(&v[0]))
}

func (location AttribLocation) AttribI4i(v0 int, v1 int, v2 int, v3 int) {
	C.glVertexAttribI4i(C.GLuint(location), C.GLint(v0), C.GLint(v1), C.GLint(v2), C.GLint(v3))
}

func (location AttribLocation) AttribI4iv(v []int32) {
	C.glVertexAttribI4iv(C.GLuint(location), (*C.GLint)(&v[0]))
}

func (location AttribLocation) AttribI4sv(v []int16) {
	C.glVertexAttribI4sv(C.GLuint(location), (*C.GLshort)(&v[0]))
}

func (location AttribLocation) AttribI4ubv(v []uint8) {
	C.glVertexAttribI4ubv(C.GLuint(location), (*C.GLubyte)(&v[0]))
}

func (location AttribLocation) AttribI4ui(v0 uint, v1 uint, v2 uint, v3 uint) {
	C.glVertexAttribI4ui(C.GLuint(location), C.GLuint(v0), C.GLuint(v1), C.GLuint(v2), C.GLuint(v3))
}

func (location AttribLocation) AttribI4uiv(v []uint32) {
	C.glVertexAttribI4uiv(C.GLuint(location), (*C.GLuint)(&v[0]))
}

func (location AttribLocation) AttribI4usv(v []uint16) {
	C.glVertexAttribI4usv(C.GLuint(location), (*C.GLushort)(&v[0]))
}

func (location AttribLocation) AttribIPointer(size int, typ GLenum, stride int, pointer interface{}) {
	C.glVertexAttribIPointer(C.GLuint(location), C.GLint(size), C.GLenum(typ), C.GLsizei(stride), glPointer(pointer))
}
