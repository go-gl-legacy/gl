// Copyright 2012 The go-gl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gl

// #include "gl.h"
//
// // Workaround for https://github.com/go-gl/gl/issues/104
// void gogl_glGetShaderSource(GLuint shader, GLsizei bufsize, GLsizei* len, GLchar* source) {
//     glGetShaderSource(shader, bufsize, len, source);
// }
import "C"
import "unsafe"

const (
	ACTIVE_ATTRIBUTE_MAX_LENGTH      = C.GL_ACTIVE_ATTRIBUTE_MAX_LENGTH
	ACTIVE_ATTRIBUTES                = C.GL_ACTIVE_ATTRIBUTES
	ACTIVE_UNIFORM_MAX_LENGTH        = C.GL_ACTIVE_UNIFORM_MAX_LENGTH
	ACTIVE_UNIFORMS                  = C.GL_ACTIVE_UNIFORMS
	ATTACHED_SHADERS                 = C.GL_ATTACHED_SHADERS
	BLEND_EQUATION_ALPHA             = C.GL_BLEND_EQUATION_ALPHA
	BLEND_EQUATION_RGB               = C.GL_BLEND_EQUATION_RGB
	BOOL                             = C.GL_BOOL
	BOOL_VEC2                        = C.GL_BOOL_VEC2
	BOOL_VEC3                        = C.GL_BOOL_VEC3
	BOOL_VEC4                        = C.GL_BOOL_VEC4
	COMPILE_STATUS                   = C.GL_COMPILE_STATUS
	COORD_REPLACE                    = C.GL_COORD_REPLACE
	CURRENT_PROGRAM                  = C.GL_CURRENT_PROGRAM
	CURRENT_VERTEX_ATTRIB            = C.GL_CURRENT_VERTEX_ATTRIB
	DELETE_STATUS                    = C.GL_DELETE_STATUS
	DRAW_BUFFER0                     = C.GL_DRAW_BUFFER0
	DRAW_BUFFER1                     = C.GL_DRAW_BUFFER1
	DRAW_BUFFER10                    = C.GL_DRAW_BUFFER10
	DRAW_BUFFER11                    = C.GL_DRAW_BUFFER11
	DRAW_BUFFER12                    = C.GL_DRAW_BUFFER12
	DRAW_BUFFER13                    = C.GL_DRAW_BUFFER13
	DRAW_BUFFER14                    = C.GL_DRAW_BUFFER14
	DRAW_BUFFER15                    = C.GL_DRAW_BUFFER15
	DRAW_BUFFER2                     = C.GL_DRAW_BUFFER2
	DRAW_BUFFER3                     = C.GL_DRAW_BUFFER3
	DRAW_BUFFER4                     = C.GL_DRAW_BUFFER4
	DRAW_BUFFER5                     = C.GL_DRAW_BUFFER5
	DRAW_BUFFER6                     = C.GL_DRAW_BUFFER6
	DRAW_BUFFER7                     = C.GL_DRAW_BUFFER7
	DRAW_BUFFER8                     = C.GL_DRAW_BUFFER8
	DRAW_BUFFER9                     = C.GL_DRAW_BUFFER9
	FLOAT_MAT2                       = C.GL_FLOAT_MAT2
	FLOAT_MAT3                       = C.GL_FLOAT_MAT3
	FLOAT_MAT4                       = C.GL_FLOAT_MAT4
	FLOAT_VEC2                       = C.GL_FLOAT_VEC2
	FLOAT_VEC3                       = C.GL_FLOAT_VEC3
	FLOAT_VEC4                       = C.GL_FLOAT_VEC4
	FRAGMENT_SHADER                  = C.GL_FRAGMENT_SHADER
	FRAGMENT_SHADER_DERIVATIVE_HINT  = C.GL_FRAGMENT_SHADER_DERIVATIVE_HINT
	INFO_LOG_LENGTH                  = C.GL_INFO_LOG_LENGTH
	INT_VEC2                         = C.GL_INT_VEC2
	INT_VEC3                         = C.GL_INT_VEC3
	INT_VEC4                         = C.GL_INT_VEC4
	LINK_STATUS                      = C.GL_LINK_STATUS
	LOWER_LEFT                       = C.GL_LOWER_LEFT
	MAX_COMBINED_TEXTURE_IMAGE_UNITS = C.GL_MAX_COMBINED_TEXTURE_IMAGE_UNITS
	MAX_DRAW_BUFFERS                 = C.GL_MAX_DRAW_BUFFERS
	MAX_FRAGMENT_UNIFORM_COMPONENTS  = C.GL_MAX_FRAGMENT_UNIFORM_COMPONENTS
	MAX_TEXTURE_COORDS               = C.GL_MAX_TEXTURE_COORDS
	MAX_TEXTURE_IMAGE_UNITS          = C.GL_MAX_TEXTURE_IMAGE_UNITS
	MAX_VARYING_FLOATS               = C.GL_MAX_VARYING_FLOATS
	MAX_VERTEX_ATTRIBS               = C.GL_MAX_VERTEX_ATTRIBS
	MAX_VERTEX_TEXTURE_IMAGE_UNITS   = C.GL_MAX_VERTEX_TEXTURE_IMAGE_UNITS
	MAX_VERTEX_UNIFORM_COMPONENTS    = C.GL_MAX_VERTEX_UNIFORM_COMPONENTS
	POINT_SPRITE                     = C.GL_POINT_SPRITE
	POINT_SPRITE_COORD_ORIGIN        = C.GL_POINT_SPRITE_COORD_ORIGIN
	SAMPLER_1D                       = C.GL_SAMPLER_1D
	SAMPLER_1D_SHADOW                = C.GL_SAMPLER_1D_SHADOW
	SAMPLER_2D                       = C.GL_SAMPLER_2D
	SAMPLER_2D_SHADOW                = C.GL_SAMPLER_2D_SHADOW
	SAMPLER_3D                       = C.GL_SAMPLER_3D
	SAMPLER_CUBE                     = C.GL_SAMPLER_CUBE
	SHADER_SOURCE_LENGTH             = C.GL_SHADER_SOURCE_LENGTH
	SHADER_TYPE                      = C.GL_SHADER_TYPE
	SHADING_LANGUAGE_VERSION         = C.GL_SHADING_LANGUAGE_VERSION
	STENCIL_BACK_FAIL                = C.GL_STENCIL_BACK_FAIL
	STENCIL_BACK_FUNC                = C.GL_STENCIL_BACK_FUNC
	STENCIL_BACK_PASS_DEPTH_FAIL     = C.GL_STENCIL_BACK_PASS_DEPTH_FAIL
	STENCIL_BACK_PASS_DEPTH_PASS     = C.GL_STENCIL_BACK_PASS_DEPTH_PASS
	STENCIL_BACK_REF                 = C.GL_STENCIL_BACK_REF
	STENCIL_BACK_VALUE_MASK          = C.GL_STENCIL_BACK_VALUE_MASK
	STENCIL_BACK_WRITEMASK           = C.GL_STENCIL_BACK_WRITEMASK
	UPPER_LEFT                       = C.GL_UPPER_LEFT
	VALIDATE_STATUS                  = C.GL_VALIDATE_STATUS
	VERTEX_ATTRIB_ARRAY_ENABLED      = C.GL_VERTEX_ATTRIB_ARRAY_ENABLED
	VERTEX_ATTRIB_ARRAY_NORMALIZED   = C.GL_VERTEX_ATTRIB_ARRAY_NORMALIZED
	VERTEX_ATTRIB_ARRAY_POINTER      = C.GL_VERTEX_ATTRIB_ARRAY_POINTER
	VERTEX_ATTRIB_ARRAY_SIZE         = C.GL_VERTEX_ATTRIB_ARRAY_SIZE
	VERTEX_ATTRIB_ARRAY_STRIDE       = C.GL_VERTEX_ATTRIB_ARRAY_STRIDE
	VERTEX_ATTRIB_ARRAY_TYPE         = C.GL_VERTEX_ATTRIB_ARRAY_TYPE
	VERTEX_PROGRAM_POINT_SIZE        = C.GL_VERTEX_PROGRAM_POINT_SIZE
	VERTEX_PROGRAM_TWO_SIDE          = C.GL_VERTEX_PROGRAM_TWO_SIDE
	VERTEX_SHADER                    = C.GL_VERTEX_SHADER
)

func (program Program) AttachShader(shader Shader) {
	C.glAttachShader(C.GLuint(program), C.GLuint(shader))
}

func (program Program) BindAttribLocation(index AttribLocation, name string) {
	cname := glString(name)
	defer freeString(cname)
	C.glBindAttribLocation(C.GLuint(program), C.GLuint(index), cname)
}

func BlendEquationSeparate(modeRGB GLenum, modeAlpha GLenum) {
	C.glBlendEquationSeparate(C.GLenum(modeRGB), C.GLenum(modeAlpha))
}

func (shader Shader) Compile() {
	C.glCompileShader(C.GLuint(shader))
}

func CreateProgram() Program {
	return Program(C.glCreateProgram())
}

func CreateShader(typ GLenum) Shader {
	return Shader(C.glCreateShader(C.GLenum(typ)))
}

func (program Program) Delete() {
	C.glDeleteProgram(C.GLuint(program))
}

func (shader Shader) Delete() {
	C.glDeleteShader(C.GLuint(shader))
}

func (program Program) DetachShader(shader Shader) {
	C.glDetachShader(C.GLuint(program), C.GLuint(shader))
}

func (index AttribLocation) DisableArray() {
	C.glDisableVertexAttribArray(C.GLuint(index))
}

func DrawBuffers(bufs []GLenum) {
	C.glDrawBuffers(C.GLsizei(len(bufs)), (*C.GLenum)(&bufs[0]))
}

func (index AttribLocation) EnableArray() {
	C.glEnableVertexAttribArray(C.GLuint(index))
}

func (program Program) GetActiveAttrib(index uint, length []int32, size []int32, typ []GLenum, name string) {
	cname := glString(name)
	defer freeString(cname)
	C.glGetActiveAttrib(C.GLuint(program), C.GLuint(index), C.GLsizei(len(typ)), (*C.GLsizei)(&length[0]), (*C.GLint)(&size[0]), (*C.GLenum)(&typ[0]), cname)
}

func (program Program) GetActiveUniform(index uint, length []int32, size []int32, typ []GLenum, name string) {
	cname := glString(name)
	defer freeString(cname)
	C.glGetActiveUniform(C.GLuint(program), C.GLuint(index), C.GLsizei(len(typ)), (*C.GLsizei)(&length[0]), (*C.GLint)(&size[0]), (*C.GLenum)(&typ[0]), cname)
}

func (program Program) GetAttachedShaders(count []int32, shaders []Shader) {
	C.glGetAttachedShaders(C.GLuint(program), C.GLsizei(len(shaders)), (*C.GLsizei)(&count[0]), (*C.GLuint)(&shaders[0]))
}

func (program Program) GetAttribLocation(name string) AttribLocation {
	cname := glString(name)
	defer freeString(cname)
	return AttribLocation(C.glGetAttribLocation(C.GLuint(program), cname))
}

func (program Program) GetInfoLog() string {
	var length C.GLint
	C.glGetProgramiv(C.GLuint(program), C.GLenum(INFO_LOG_LENGTH), &length)

	if length > 1 {
		log := C.malloc(C.size_t(length))
		defer C.free(log)
		C.glGetProgramInfoLog(C.GLuint(program), C.GLsizei(length), nil, (*C.GLchar)(log))
		return C.GoString((*C.char)(log))
	}
	return ""
}

func (program Program) Get(param GLenum) int {
	var rv C.GLint
	C.glGetProgramiv(C.GLuint(program), C.GLenum(param), &rv)
	return int(rv)
}

func (shader Shader) GetInfoLog() string {
	var length C.GLint
	C.glGetShaderiv(C.GLuint(shader), C.GLenum(INFO_LOG_LENGTH), &length)

	if length > 1 {
		log := C.malloc(C.size_t(length))
		defer C.free(log)
		C.glGetShaderInfoLog(C.GLuint(shader), C.GLsizei(length), nil, (*C.GLchar)(log))
		return C.GoString((*C.char)(log))
	}
	return ""
}

func (shader Shader) GetSource() string {
	var length C.GLint
	C.glGetShaderiv(C.GLuint(shader), C.GLenum(SHADER_SOURCE_LENGTH), &length)

	log := C.malloc(C.size_t(length + 1))
	C.gogl_glGetShaderSource(C.GLuint(shader), C.GLsizei(length), nil, (*C.GLchar)(log))
	defer C.free(log)

	if length > 1 {
		log := C.malloc(C.size_t(length + 1))
		defer C.free(log)
		C.gogl_glGetShaderSource(C.GLuint(shader), C.GLsizei(length), nil, (*C.GLchar)(log))
		return C.GoString((*C.char)(log))
	}
	return ""
}

func (shader Shader) Get(param GLenum) int {
	var rv C.GLint
	C.glGetShaderiv(C.GLuint(shader), C.GLenum(param), &rv)
	return int(rv)
}

func (program Program) GetUniformLocation(name string) UniformLocation {
	cname := glString(name)
	defer freeString(cname)
	return UniformLocation(C.glGetUniformLocation(C.GLuint(program), cname))
}

func (program Program) GetUniformfv(location UniformLocation, params []float32) {
	C.glGetUniformfv(C.GLuint(program), C.GLint(location), (*C.GLfloat)(&params[0]))
}

func (program Program) GetUniformiv(location UniformLocation, params []int32) {
	C.glGetUniformiv(C.GLuint(program), C.GLint(location), (*C.GLint)(&params[0]))
}

func GetVertexAttribPointerv(index uint, pname GLenum) (ptr unsafe.Pointer) {
	C.glGetVertexAttribPointerv(C.GLuint(index), C.GLenum(pname), &ptr)
	return
}

func GetVertexAttribdv(index uint, pname GLenum, params []float64) {
	C.glGetVertexAttribdv(C.GLuint(index), C.GLenum(pname), (*C.GLdouble)(&params[0]))
}

func GetVertexAttribfv(index uint, pname GLenum, params []float32) {
	C.glGetVertexAttribfv(C.GLuint(index), C.GLenum(pname), (*C.GLfloat)(&params[0]))
}

func GetVertexAttribiv(index uint, pname GLenum, params []int32) {
	C.glGetVertexAttribiv(C.GLuint(index), C.GLenum(pname), (*C.GLint)(&params[0]))
}

func (object Object) IsProgram() bool {
	return goBool(C.glIsProgram(C.GLuint(object)))
}

func (object Object) IsShader() bool {
	return goBool(C.glIsShader(C.GLuint(object)))
}

func (program Program) Link() {
	C.glLinkProgram(C.GLuint(program))
}

func (shader Shader) Source(source ...string) {
	count := C.GLsizei(len(source))
	cstrings := make([]*C.GLchar, count)
	length := make([]C.GLint, count)

	for i, s := range source {
		csource := glString(s)
		cstrings[i] = csource
		length[i] = C.GLint(len(s))
		defer freeString(csource)
	}

	C.glShaderSource(C.GLuint(shader), count, (**_Ctype_GLchar)(unsafe.Pointer(&cstrings[0])), (*_Ctype_GLint)(unsafe.Pointer(&length[0])))
}

func StencilFuncSeparate(frontfunc GLenum, backfunc GLenum, ref int, mask uint32) {
	C.glStencilFuncSeparate(C.GLenum(frontfunc), C.GLenum(backfunc), C.GLint(ref), C.GLuint(mask))
}

func StencilMaskSeparate(face GLenum, mask uint32) {
	C.glStencilMaskSeparate(C.GLenum(face), C.GLuint(mask))
}

func StencilOpSeparate(face GLenum, sfail GLenum, dpfail GLenum, dppass GLenum) {
	C.glStencilOpSeparate(C.GLenum(face), C.GLenum(sfail), C.GLenum(dpfail), C.GLenum(dppass))
}

func (location UniformLocation) Uniform1f(x float32) {
	C.glUniform1f(C.GLint(location), C.GLfloat(x))
}

func (location UniformLocation) Uniform2f(x float32, y float32) {
	C.glUniform2f(C.GLint(location), C.GLfloat(x), C.GLfloat(y))
}

func (location UniformLocation) Uniform3f(x float32, y float32, z float32) {
	C.glUniform3f(C.GLint(location), C.GLfloat(x), C.GLfloat(y), C.GLfloat(z))
}

func (location UniformLocation) Uniform1fv(count int, v []float32) {
	C.glUniform1fv(C.GLint(location), C.GLsizei(count), (*C.GLfloat)(&v[0]))
}

func (location UniformLocation) Uniform1i(x int) {
	C.glUniform1i(C.GLint(location), C.GLint(x))
}

func (location UniformLocation) Uniform1iv(count int, v []int32) {
	C.glUniform1iv(C.GLint(location), C.GLsizei(count), (*C.GLint)(&v[0]))
}

func (location UniformLocation) Uniform2fv(count int, v []float32) {
	C.glUniform2fv(C.GLint(location), C.GLsizei(count), (*C.GLfloat)(&v[0]))
}

func (location UniformLocation) Uniform2i(x int, y int) {
	C.glUniform2i(C.GLint(location), C.GLint(x), C.GLint(y))
}

func (location UniformLocation) Uniform2iv(count int, v []int32) {
	C.glUniform2iv(C.GLint(location), C.GLsizei(count), (*C.GLint)(&v[0]))
}

func (location UniformLocation) Uniform3fv(count int, v []float32) {
	C.glUniform3fv(C.GLint(location), C.GLsizei(count), (*C.GLfloat)(&v[0]))
}

func (location UniformLocation) Uniform3i(x int, y int, z int) {
	C.glUniform3i(C.GLint(location), C.GLint(x), C.GLint(y), C.GLint(z))
}

func (location UniformLocation) Uniform3iv(count int, v []int32) {
	C.glUniform3iv(C.GLint(location), C.GLsizei(count), (*C.GLint)(&v[0]))
}

func (location UniformLocation) Uniform4f(x float32, y float32, z float32, w float32) {
	C.glUniform4f(C.GLint(location), C.GLfloat(x), C.GLfloat(y), C.GLfloat(z), C.GLfloat(w))
}

func (location UniformLocation) Uniform4fv(count int, v []float32) {
	C.glUniform4fv(C.GLint(location), C.GLsizei(count), (*C.GLfloat)(&v[0]))
}

func (location UniformLocation) Uniform4i(x int, y int, z int, w int) {
	C.glUniform4i(C.GLint(location), C.GLint(x), C.GLint(y), C.GLint(z), C.GLint(w))
}

func (location UniformLocation) Uniform4iv(count int, v []int32) {
	C.glUniform4iv(C.GLint(location), C.GLsizei(count), (*C.GLint)(&v[0]))
}

func (location UniformLocation) UniformMatrix2fv(transpose bool, list ...[4]float32) {
	C.glUniformMatrix2fv(C.GLint(location), C.GLsizei(len(list)), glBool(transpose), ((*C.GLfloat)((unsafe.Pointer)(&list[0]))))
}

func (location UniformLocation) UniformMatrix2f(transpose bool, matrix *[4]float32) {
	C.glUniformMatrix2fv(C.GLint(location), 1, glBool(transpose), ((*C.GLfloat)((unsafe.Pointer)(&matrix[0]))))
}

func (location UniformLocation) UniformMatrix3fv(transpose bool, list ...[9]float32) {
	C.glUniformMatrix3fv(C.GLint(location), C.GLsizei(len(list)), glBool(transpose), ((*C.GLfloat)((unsafe.Pointer)(&list[0]))))
}

func (location UniformLocation) UniformMatrix3f(transpose bool, matrix *[9]float32) {
	C.glUniformMatrix3fv(C.GLint(location), 1, glBool(transpose), ((*C.GLfloat)((unsafe.Pointer)(&matrix[0]))))
}

func (location UniformLocation) UniformMatrix4fv(transpose bool, list ...[16]float32) {
	C.glUniformMatrix4fv(C.GLint(location), C.GLsizei(len(list)), glBool(transpose), ((*C.GLfloat)((unsafe.Pointer)(&list[0]))))
}

func (location UniformLocation) UniformMatrix4f(transpose bool, matrix *[16]float32) {
	C.glUniformMatrix4fv(C.GLint(location), 1, glBool(transpose), ((*C.GLfloat)((unsafe.Pointer)(&matrix[0]))))
}

func (program Program) Use() {
	C.glUseProgram(C.GLuint(program))
}

func (program Program) Unuse() {
	C.glUseProgram(C.GLuint(0))
}

func (program Program) Validate() {
	C.glValidateProgram(C.GLuint(program))
}

func (index AttribLocation) Attrib1d(x float64) {
	C.glVertexAttrib1d(C.GLuint(index), C.GLdouble(x))
}

func (index AttribLocation) Attrib1dv(v []float64) {
	C.glVertexAttrib1dv(C.GLuint(index), (*C.GLdouble)(&v[0]))
}

func (index AttribLocation) Attrib1f(x float32) {
	C.glVertexAttrib1f(C.GLuint(index), C.GLfloat(x))
}

func (index AttribLocation) Attrib1fv(v []float32) {
	C.glVertexAttrib1fv(C.GLuint(index), (*C.GLfloat)(&v[0]))
}

func (index AttribLocation) Attrib1s(x int16) {
	C.glVertexAttrib1s(C.GLuint(index), C.GLshort(x))
}

func (index AttribLocation) Attrib1sv(v []int16) {
	C.glVertexAttrib1sv(C.GLuint(index), (*C.GLshort)(&v[0]))
}

func (index AttribLocation) Attrib2d(x float64, y float64) {
	C.glVertexAttrib2d(C.GLuint(index), C.GLdouble(x), C.GLdouble(y))
}

func (index AttribLocation) Attrib2dv(v []float64) {
	C.glVertexAttrib2dv(C.GLuint(index), (*C.GLdouble)(&v[0]))
}

func (index AttribLocation) Attrib2f(x float32, y float32) {
	C.glVertexAttrib2f(C.GLuint(index), C.GLfloat(x), C.GLfloat(y))
}

func (index AttribLocation) Attrib2fv(v []float32) {
	C.glVertexAttrib2fv(C.GLuint(index), (*C.GLfloat)(&v[0]))
}

func (index AttribLocation) Attrib2s(x int16, y int16) {
	C.glVertexAttrib2s(C.GLuint(index), C.GLshort(x), C.GLshort(y))
}

func (index AttribLocation) Attrib2sv(v []int16) {
	C.glVertexAttrib2sv(C.GLuint(index), (*C.GLshort)(&v[0]))
}

func (index AttribLocation) Attrib3d(x float64, y float64, z float64) {
	C.glVertexAttrib3d(C.GLuint(index), C.GLdouble(x), C.GLdouble(y), C.GLdouble(z))
}

func (index AttribLocation) Attrib3dv(v []float64) {
	C.glVertexAttrib3dv(C.GLuint(index), (*C.GLdouble)(&v[0]))
}

func (index AttribLocation) Attrib3f(x float32, y float32, z float32) {
	C.glVertexAttrib3f(C.GLuint(index), C.GLfloat(x), C.GLfloat(y), C.GLfloat(z))
}

func (index AttribLocation) Attrib3fv(v []float32) {
	C.glVertexAttrib3fv(C.GLuint(index), (*C.GLfloat)(&v[0]))
}

func (index AttribLocation) Attrib3s(x int16, y int16, z int16) {
	C.glVertexAttrib3s(C.GLuint(index), C.GLshort(x), C.GLshort(y), C.GLshort(z))
}

func (index AttribLocation) Attrib3sv(v []int16) {
	C.glVertexAttrib3sv(C.GLuint(index), (*C.GLshort)(&v[0]))
}

func (index AttribLocation) Attrib4Nbv(v []int8) {
	C.glVertexAttrib4Nbv(C.GLuint(index), (*C.GLbyte)(&v[0]))
}

func (index AttribLocation) Attrib4Niv(v []int32) {
	C.glVertexAttrib4Niv(C.GLuint(index), (*C.GLint)(&v[0]))
}

func (index AttribLocation) Attrib4Nsv(v []int16) {
	C.glVertexAttrib4Nsv(C.GLuint(index), (*C.GLshort)(&v[0]))
}

func (index AttribLocation) Attrib4Nub(x uint8, y uint8, z uint8, w uint8) {
	C.glVertexAttrib4Nub(C.GLuint(index), C.GLubyte(x), C.GLubyte(y), C.GLubyte(z), C.GLubyte(w))
}

func (index AttribLocation) Attrib4Nubv(v []uint8) {
	C.glVertexAttrib4Nubv(C.GLuint(index), (*C.GLubyte)(&v[0]))
}

func (index AttribLocation) Attrib4Nuiv(v []uint32) {
	C.glVertexAttrib4Nuiv(C.GLuint(index), (*C.GLuint)(&v[0]))
}

func (index AttribLocation) Attrib4Nusv(v []uint16) {
	C.glVertexAttrib4Nusv(C.GLuint(index), (*C.GLushort)(&v[0]))
}

func (index AttribLocation) Attrib4bv(v []int8) {
	C.glVertexAttrib4bv(C.GLuint(index), (*C.GLbyte)(&v[0]))
}

func (index AttribLocation) Attrib4d(x float64, y float64, z float64, w float64) {
	C.glVertexAttrib4d(C.GLuint(index), C.GLdouble(x), C.GLdouble(y), C.GLdouble(z), C.GLdouble(w))
}

func (index AttribLocation) Attrib4dv(v []float64) {
	C.glVertexAttrib4dv(C.GLuint(index), (*C.GLdouble)(&v[0]))
}

func (index AttribLocation) Attrib4f(x float32, y float32, z float32, w float32) {
	C.glVertexAttrib4f(C.GLuint(index), C.GLfloat(x), C.GLfloat(y), C.GLfloat(z), C.GLfloat(w))
}

func (index AttribLocation) Attrib4fv(v []float32) {
	C.glVertexAttrib4fv(C.GLuint(index), (*C.GLfloat)(&v[0]))
}

func (index AttribLocation) Attrib4iv(v []int32) {
	C.glVertexAttrib4iv(C.GLuint(index), (*C.GLint)(&v[0]))
}

func (index AttribLocation) Attrib4s(x int16, y int16, z int16, w int16) {
	C.glVertexAttrib4s(C.GLuint(index), C.GLshort(x), C.GLshort(y), C.GLshort(z), C.GLshort(w))
}

func (index AttribLocation) Attrib4sv(v []int16) {
	C.glVertexAttrib4sv(C.GLuint(index), (*C.GLshort)(&v[0]))
}

func (index AttribLocation) Attrib4ubv(v []uint8) {
	C.glVertexAttrib4ubv(C.GLuint(index), (*C.GLubyte)(&v[0]))
}

func (index AttribLocation) Attrib4uiv(v []uint32) {
	C.glVertexAttrib4uiv(C.GLuint(index), (*C.GLuint)(&v[0]))
}

func (index AttribLocation) Attrib4usv(v []uint16) {
	C.glVertexAttrib4usv(C.GLuint(index), (*C.GLushort)(&v[0]))
}

func (index AttribLocation) AttribPointer(size int, typ GLenum, normalized bool, stride int, pointer interface{}) {
	C.glVertexAttribPointer(C.GLuint(index), C.GLint(size), C.GLenum(typ), glBool(normalized), C.GLsizei(stride), glPointer(pointer))
}
