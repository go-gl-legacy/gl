// Copyright 2012 The go-gl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gl

// #include "gl.h"
import "C"
import "unsafe"

const (
	PROGRAM_MATRIX             = C.GL_PROGRAM_MATRIX_EXT
	PROGRAM_MATRIX_STACK_DEPTH = C.GL_PROGRAM_MATRIX_STACK_DEPTH_EXT
	TRANSPOSE_PROGRAM_MATRIX   = C.GL_TRANSPOSE_PROGRAM_MATRIX_EXT
)

func BindMultiTexture(texunit GLenum, target GLenum, texture uint) {
	C.glBindMultiTextureEXT(C.GLenum(texunit), C.GLenum(target), C.GLuint(texture))
}

func (framebuffer Framebuffer) CheckNamedFramebufferStatus(target GLenum) GLenum {
	return GLenum(C.glCheckNamedFramebufferStatusEXT(C.GLuint(framebuffer), C.GLenum(target)))
}

func ClientAttribDefault(mask uint32) {
	C.glClientAttribDefaultEXT(C.GLbitfield(mask))
}

func CompressedMultiTexImage1D(texunit GLenum, target GLenum, level int, internalformat GLenum, width int, border int, imageSize int, data interface{}) {
	C.glCompressedMultiTexImage1DEXT(C.GLenum(texunit), C.GLenum(target), C.GLint(level), C.GLenum(internalformat), C.GLsizei(width), C.GLint(border), C.GLsizei(imageSize), glPointer(data))
}

func CompressedMultiTexImage2D(texunit GLenum, target GLenum, level int, internalformat GLenum, width int, height int, border int, imageSize int, data interface{}) {
	C.glCompressedMultiTexImage2DEXT(C.GLenum(texunit), C.GLenum(target), C.GLint(level), C.GLenum(internalformat), C.GLsizei(width), C.GLsizei(height), C.GLint(border), C.GLsizei(imageSize), glPointer(data))
}

func CompressedMultiTexImage3D(texunit GLenum, target GLenum, level int, internalformat GLenum, width int, height int, depth int, border int, imageSize int, data interface{}) {
	C.glCompressedMultiTexImage3DEXT(C.GLenum(texunit), C.GLenum(target), C.GLint(level), C.GLenum(internalformat), C.GLsizei(width), C.GLsizei(height), C.GLsizei(depth), C.GLint(border), C.GLsizei(imageSize), glPointer(data))
}

func CompressedMultiTexSubImage1D(texunit GLenum, target GLenum, level int, xoffset int, width int, format GLenum, imageSize int, data interface{}) {
	C.glCompressedMultiTexSubImage1DEXT(C.GLenum(texunit), C.GLenum(target), C.GLint(level), C.GLint(xoffset), C.GLsizei(width), C.GLenum(format), C.GLsizei(imageSize), glPointer(data))
}

func CompressedMultiTexSubImage2D(texunit GLenum, target GLenum, level int, xoffset int, yoffset int, width int, height int, format GLenum, imageSize int, data interface{}) {
	C.glCompressedMultiTexSubImage2DEXT(C.GLenum(texunit), C.GLenum(target), C.GLint(level), C.GLint(xoffset), C.GLint(yoffset), C.GLsizei(width), C.GLsizei(height), C.GLenum(format), C.GLsizei(imageSize), glPointer(data))
}

func CompressedMultiTexSubImage3D(texunit GLenum, target GLenum, level int, xoffset int, yoffset int, zoffset int, width int, height int, depth int, format GLenum, imageSize int, data interface{}) {
	C.glCompressedMultiTexSubImage3DEXT(C.GLenum(texunit), C.GLenum(target), C.GLint(level), C.GLint(xoffset), C.GLint(yoffset), C.GLint(zoffset), C.GLsizei(width), C.GLsizei(height), C.GLsizei(depth), C.GLenum(format), C.GLsizei(imageSize), glPointer(data))
}

func (texture Texture) CompressedTextureImage1D(target GLenum, level int, internalformat GLenum, width int, border int, imageSize int, data interface{}) {
	C.glCompressedTextureImage1DEXT(C.GLuint(texture), C.GLenum(target), C.GLint(level), C.GLenum(internalformat), C.GLsizei(width), C.GLint(border), C.GLsizei(imageSize), glPointer(data))
}

func (texture Texture) CompressedTextureImage2D(target GLenum, level int, internalformat GLenum, width int, height int, border int, imageSize int, data interface{}) {
	C.glCompressedTextureImage2DEXT(C.GLuint(texture), C.GLenum(target), C.GLint(level), C.GLenum(internalformat), C.GLsizei(width), C.GLsizei(height), C.GLint(border), C.GLsizei(imageSize), glPointer(data))
}

func (texture Texture) CompressedTextureImage3D(target GLenum, level int, internalformat GLenum, width int, height int, depth int, border int, imageSize int, data interface{}) {
	C.glCompressedTextureImage3DEXT(C.GLuint(texture), C.GLenum(target), C.GLint(level), C.GLenum(internalformat), C.GLsizei(width), C.GLsizei(height), C.GLsizei(depth), C.GLint(border), C.GLsizei(imageSize), glPointer(data))
}

func (texture Texture) CompressedTextureSubImage1D(target GLenum, level int, xoffset int, width int, format GLenum, imageSize int, data interface{}) {
	C.glCompressedTextureSubImage1DEXT(C.GLuint(texture), C.GLenum(target), C.GLint(level), C.GLint(xoffset), C.GLsizei(width), C.GLenum(format), C.GLsizei(imageSize), glPointer(data))
}

func (texture Texture) CompressedTextureSubImage2D(target GLenum, level int, xoffset int, yoffset int, width int, height int, format GLenum, imageSize int, data interface{}) {
	C.glCompressedTextureSubImage2DEXT(C.GLuint(texture), C.GLenum(target), C.GLint(level), C.GLint(xoffset), C.GLint(yoffset), C.GLsizei(width), C.GLsizei(height), C.GLenum(format), C.GLsizei(imageSize), glPointer(data))
}

func (texture Texture) CompressedTextureSubImage3D(target GLenum, level int, xoffset int, yoffset int, zoffset int, width int, height int, depth int, format GLenum, imageSize int, data interface{}) {
	C.glCompressedTextureSubImage3DEXT(C.GLuint(texture), C.GLenum(target), C.GLint(level), C.GLint(xoffset), C.GLint(yoffset), C.GLint(zoffset), C.GLsizei(width), C.GLsizei(height), C.GLsizei(depth), C.GLenum(format), C.GLsizei(imageSize), glPointer(data))
}

func CopyMultiTexImage1D(texunit GLenum, target GLenum, level int, internalformat GLenum, x int, y int, width int, border int) {
	C.glCopyMultiTexImage1DEXT(C.GLenum(texunit), C.GLenum(target), C.GLint(level), C.GLenum(internalformat), C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLint(border))
}

func CopyMultiTexImage2D(texunit GLenum, target GLenum, level int, internalformat GLenum, x int, y int, width int, height int, border int) {
	C.glCopyMultiTexImage2DEXT(C.GLenum(texunit), C.GLenum(target), C.GLint(level), C.GLenum(internalformat), C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height), C.GLint(border))
}

func CopyMultiTexSubImage1D(texunit GLenum, target GLenum, level int, xoffset int, x int, y int, width int) {
	C.glCopyMultiTexSubImage1DEXT(C.GLenum(texunit), C.GLenum(target), C.GLint(level), C.GLint(xoffset), C.GLint(x), C.GLint(y), C.GLsizei(width))
}

func CopyMultiTexSubImage2D(texunit GLenum, target GLenum, level int, xoffset int, yoffset int, x int, y int, width int, height int) {
	C.glCopyMultiTexSubImage2DEXT(C.GLenum(texunit), C.GLenum(target), C.GLint(level), C.GLint(xoffset), C.GLint(yoffset), C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height))
}

func CopyMultiTexSubImage3D(texunit GLenum, target GLenum, level int, xoffset int, yoffset int, zoffset int, x int, y int, width int, height int) {
	C.glCopyMultiTexSubImage3DEXT(C.GLenum(texunit), C.GLenum(target), C.GLint(level), C.GLint(xoffset), C.GLint(yoffset), C.GLint(zoffset), C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height))
}

func (texture Texture) CopyTextureImage1D(target GLenum, level int, internalformat GLenum, x int, y int, width int, border int) {
	C.glCopyTextureImage1DEXT(C.GLuint(texture), C.GLenum(target), C.GLint(level), C.GLenum(internalformat), C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLint(border))
}

func (texture Texture) CopyTextureImage2D(target GLenum, level int, internalformat GLenum, x int, y int, width int, height int, border int) {
	C.glCopyTextureImage2DEXT(C.GLuint(texture), C.GLenum(target), C.GLint(level), C.GLenum(internalformat), C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height), C.GLint(border))
}

func (texture Texture) CopyTextureSubImage1D(target GLenum, level int, xoffset int, x int, y int, width int) {
	C.glCopyTextureSubImage1DEXT(C.GLuint(texture), C.GLenum(target), C.GLint(level), C.GLint(xoffset), C.GLint(x), C.GLint(y), C.GLsizei(width))
}

func (texture Texture) CopyTextureSubImage2D(target GLenum, level int, xoffset int, yoffset int, x int, y int, width int, height int) {
	C.glCopyTextureSubImage2DEXT(C.GLuint(texture), C.GLenum(target), C.GLint(level), C.GLint(xoffset), C.GLint(yoffset), C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height))
}

func (texture Texture) CopyTextureSubImage3D(target GLenum, level int, xoffset int, yoffset int, zoffset int, x int, y int, width int, height int) {
	C.glCopyTextureSubImage3DEXT(C.GLuint(texture), C.GLenum(target), C.GLint(level), C.GLint(xoffset), C.GLint(yoffset), C.GLint(zoffset), C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height))
}

func DisableClientStateIndexed(array GLenum, index uint) {
	C.glDisableClientStateIndexedEXT(C.GLenum(array), C.GLuint(index))
}

func DisableClientStatei(array GLenum, index uint) {
	C.glDisableClientStateiEXT(C.GLenum(array), C.GLuint(index))
}

func (vaobj VertexArray) DisableVertexArrayAttrib(index uint) {
	C.glDisableVertexArrayAttribEXT(C.GLuint(vaobj), C.GLuint(index))
}

func (vaobj VertexArray) DisableVertexArray(array GLenum) {
	C.glDisableVertexArrayEXT(C.GLuint(vaobj), C.GLenum(array))
}

func EnableClientStateIndexed(array GLenum, index uint) {
	C.glEnableClientStateIndexedEXT(C.GLenum(array), C.GLuint(index))
}

func EnableClientStatei(array GLenum, index uint) {
	C.glEnableClientStateiEXT(C.GLenum(array), C.GLuint(index))
}

func (vaobj VertexArray) EnableVertexArrayAttrib(index uint) {
	C.glEnableVertexArrayAttribEXT(C.GLuint(vaobj), C.GLuint(index))
}

func (vaobj VertexArray) EnableVertexArray(array GLenum) {
	C.glEnableVertexArrayEXT(C.GLuint(vaobj), C.GLenum(array))
}

func (buffer Buffer) FlushMappedNamedBufferRange(offset uintptr, length uintptr) {
	C.glFlushMappedNamedBufferRangeEXT(C.GLuint(buffer), C.GLintptr(offset), C.GLsizeiptr(length))
}

func (framebuffer Framebuffer) FramebufferDrawBuffer(mode GLenum) {
	C.glFramebufferDrawBufferEXT(C.GLuint(framebuffer), C.GLenum(mode))
}

func (framebuffer Framebuffer) FramebufferDrawBuffers(n int, bufs []GLenum) {
	C.glFramebufferDrawBuffersEXT(C.GLuint(framebuffer), C.GLsizei(n), (*C.GLenum)(&bufs[0]))
}

func (framebuffer Framebuffer) FramebufferReadBuffer(mode GLenum) {
	C.glFramebufferReadBufferEXT(C.GLuint(framebuffer), C.GLenum(mode))
}

func GenerateMultiTexMipmap(texunit GLenum, target GLenum) {
	C.glGenerateMultiTexMipmapEXT(C.GLenum(texunit), C.GLenum(target))
}

func (texture Texture) GenerateTextureMipmap(target GLenum) {
	C.glGenerateTextureMipmapEXT(C.GLuint(texture), C.GLenum(target))
}

func GetCompressedMultiTexImage(texunit GLenum, target GLenum, level int, img interface{}) {
	C.glGetCompressedMultiTexImageEXT(C.GLenum(texunit), C.GLenum(target), C.GLint(level), glPointer(img))
}

func (texture Texture) GetCompressedTextureImage(target GLenum, level int, img interface{}) {
	C.glGetCompressedTextureImageEXT(C.GLuint(texture), C.GLenum(target), C.GLint(level), glPointer(img))
}

func GetDoubleIndexedv(target GLenum, index uint, params []float64) {
	C.glGetDoubleIndexedvEXT(C.GLenum(target), C.GLuint(index), (*C.GLdouble)(&params[0]))
}

func GetDoublei_v(pname GLenum, index uint, params []float64) {
	C.glGetDoublei_vEXT(C.GLenum(pname), C.GLuint(index), (*C.GLdouble)(&params[0]))
}

func GetFloatIndexedv(target GLenum, index uint, params []float32) {
	C.glGetFloatIndexedvEXT(C.GLenum(target), C.GLuint(index), (*C.GLfloat)(&params[0]))
}

func GetFloati_v(pname GLenum, index uint, params []float32) {
	C.glGetFloati_vEXT(C.GLenum(pname), C.GLuint(index), (*C.GLfloat)(&params[0]))
}

func (framebuffer Framebuffer) GetFramebufferParameteriv(pname GLenum, params []int32) {
	C.glGetFramebufferParameterivEXT(C.GLuint(framebuffer), C.GLenum(pname), (*C.GLint)(&params[0]))
}

func GetMultiTexEnvfv(texunit GLenum, target GLenum, pname GLenum, params []float32) {
	C.glGetMultiTexEnvfvEXT(C.GLenum(texunit), C.GLenum(target), C.GLenum(pname), (*C.GLfloat)(&params[0]))
}

func GetMultiTexEnviv(texunit GLenum, target GLenum, pname GLenum, params []int32) {
	C.glGetMultiTexEnvivEXT(C.GLenum(texunit), C.GLenum(target), C.GLenum(pname), (*C.GLint)(&params[0]))
}

func GetMultiTexGendv(texunit GLenum, coord GLenum, pname GLenum, params []float64) {
	C.glGetMultiTexGendvEXT(C.GLenum(texunit), C.GLenum(coord), C.GLenum(pname), (*C.GLdouble)(&params[0]))
}

func GetMultiTexGenfv(texunit GLenum, coord GLenum, pname GLenum, params []float32) {
	C.glGetMultiTexGenfvEXT(C.GLenum(texunit), C.GLenum(coord), C.GLenum(pname), (*C.GLfloat)(&params[0]))
}

func GetMultiTexGeniv(texunit GLenum, coord GLenum, pname GLenum, params []int32) {
	C.glGetMultiTexGenivEXT(C.GLenum(texunit), C.GLenum(coord), C.GLenum(pname), (*C.GLint)(&params[0]))
}

func GetMultiTexImage(texunit GLenum, target GLenum, level int, format GLenum, typ GLenum, pixels interface{}) {
	C.glGetMultiTexImageEXT(C.GLenum(texunit), C.GLenum(target), C.GLint(level), C.GLenum(format), C.GLenum(typ), glPointer(pixels))
}

func GetMultiTexLevelParameterfv(texunit GLenum, target GLenum, level int, pname GLenum, params []float32) {
	C.glGetMultiTexLevelParameterfvEXT(C.GLenum(texunit), C.GLenum(target), C.GLint(level), C.GLenum(pname), (*C.GLfloat)(&params[0]))
}

func GetMultiTexLevelParameteriv(texunit GLenum, target GLenum, level int, pname GLenum, params []int32) {
	C.glGetMultiTexLevelParameterivEXT(C.GLenum(texunit), C.GLenum(target), C.GLint(level), C.GLenum(pname), (*C.GLint)(&params[0]))
}

func GetMultiTexParameterIiv(texunit GLenum, target GLenum, pname GLenum, params []int32) {
	C.glGetMultiTexParameterIivEXT(C.GLenum(texunit), C.GLenum(target), C.GLenum(pname), (*C.GLint)(&params[0]))
}

func GetMultiTexParameterIuiv(texunit GLenum, target GLenum, pname GLenum, params []uint32) {
	C.glGetMultiTexParameterIuivEXT(C.GLenum(texunit), C.GLenum(target), C.GLenum(pname), (*C.GLuint)(&params[0]))
}

func GetMultiTexParameterfv(texunit GLenum, target GLenum, pname GLenum, params []float32) {
	C.glGetMultiTexParameterfvEXT(C.GLenum(texunit), C.GLenum(target), C.GLenum(pname), (*C.GLfloat)(&params[0]))
}

func GetMultiTexParameteriv(texunit GLenum, target GLenum, pname GLenum, params []int32) {
	C.glGetMultiTexParameterivEXT(C.GLenum(texunit), C.GLenum(target), C.GLenum(pname), (*C.GLint)(&params[0]))
}

func (buffer Buffer) GetNamedBufferParameteriv(pname GLenum, params []int32) {
	C.glGetNamedBufferParameterivEXT(C.GLuint(buffer), C.GLenum(pname), (*C.GLint)(&params[0]))
}

func (buffer Buffer) GetNamedBufferPointerv(pname GLenum) (ptr unsafe.Pointer) {
	C.glGetNamedBufferPointervEXT(C.GLuint(buffer), C.GLenum(pname), &ptr)
	return
}

func (buffer Buffer) GetNamedBufferSubData(offset uintptr, size uintptr, data unsafe.Pointer) {
	C.glGetNamedBufferSubDataEXT(C.GLuint(buffer), C.GLintptr(offset), C.GLsizeiptr(size), data)
}

func (framebuffer Framebuffer) GetNamedFramebufferAttachmentParameteriv(attachment GLenum, pname GLenum, params []int32) {
	C.glGetNamedFramebufferAttachmentParameterivEXT(C.GLuint(framebuffer), C.GLenum(attachment), C.GLenum(pname), (*C.GLint)(&params[0]))
}

func (program Program) GetNamedProgramLocalParameterIiv(target GLenum, index uint, params []int32) {
	C.glGetNamedProgramLocalParameterIivEXT(C.GLuint(program), C.GLenum(target), C.GLuint(index), (*C.GLint)(&params[0]))
}

func (program Program) GetNamedProgramLocalParameterIuiv(target GLenum, index uint, params []uint32) {
	C.glGetNamedProgramLocalParameterIuivEXT(C.GLuint(program), C.GLenum(target), C.GLuint(index), (*C.GLuint)(&params[0]))
}

func (program Program) GetNamedProgramLocalParameterdv(target GLenum, index uint, params []float64) {
	C.glGetNamedProgramLocalParameterdvEXT(C.GLuint(program), C.GLenum(target), C.GLuint(index), (*C.GLdouble)(&params[0]))
}

func (program Program) GetNamedProgramLocalParameterfv(target GLenum, index uint, params []float32) {
	C.glGetNamedProgramLocalParameterfvEXT(C.GLuint(program), C.GLenum(target), C.GLuint(index), (*C.GLfloat)(&params[0]))
}

// TODO
// func (program Program) GetNamedProgramString(target GLenum, pname GLenum) string {
// 	C.glGetNamedProgramStringEXT(C.GLuint(program), C.GLenum(target), C.GLenum(pname), C.GLvoid(*string))
// }

func (program Program) GetNamedProgramiv(target GLenum, pname GLenum, params []int32) {
	C.glGetNamedProgramivEXT(C.GLuint(program), C.GLenum(target), C.GLenum(pname), (*C.GLint)(&params[0]))
}

func (renderbuffer Renderbuffer) GetNamedRenderbufferParameteriv(pname GLenum, params []int32) {
	C.glGetNamedRenderbufferParameterivEXT(C.GLuint(renderbuffer), C.GLenum(pname), (*C.GLint)(&params[0]))
}

// Slice?
func GetPointerIndexedv(target GLenum, index uint) (ptr unsafe.Pointer) {
	C.glGetPointerIndexedvEXT(C.GLenum(target), C.GLuint(index), &ptr)
	return
}

// Slice?
func GetPointeri_v(pname GLenum, index uint) (ptr unsafe.Pointer) {
	C.glGetPointeri_vEXT(C.GLenum(pname), C.GLuint(index), &ptr)
	return
}

func (texture Texture) GetTextureImage(target GLenum, level int, format GLenum, typ GLenum, pixels interface{}) {
	C.glGetTextureImageEXT(C.GLuint(texture), C.GLenum(target), C.GLint(level), C.GLenum(format), C.GLenum(typ), glPointer(pixels))
}

func (texture Texture) GetTextureLevelParameterfv(target GLenum, level int, pname GLenum, params []float32) {
	C.glGetTextureLevelParameterfvEXT(C.GLuint(texture), C.GLenum(target), C.GLint(level), C.GLenum(pname), (*C.GLfloat)(&params[0]))
}

func (texture Texture) GetTextureLevelParameteriv(target GLenum, level int, pname GLenum, params []int32) {
	C.glGetTextureLevelParameterivEXT(C.GLuint(texture), C.GLenum(target), C.GLint(level), C.GLenum(pname), (*C.GLint)(&params[0]))
}

func (texture Texture) GetTextureParameterIiv(target GLenum, pname GLenum, params []int32) {
	C.glGetTextureParameterIivEXT(C.GLuint(texture), C.GLenum(target), C.GLenum(pname), (*C.GLint)(&params[0]))
}

func (texture Texture) GetTextureParameterIuiv(target GLenum, pname GLenum, params []uint32) {
	C.glGetTextureParameterIuivEXT(C.GLuint(texture), C.GLenum(target), C.GLenum(pname), (*C.GLuint)(&params[0]))
}

func (texture Texture) GetTextureParameterfv(target GLenum, pname GLenum, params []float32) {
	C.glGetTextureParameterfvEXT(C.GLuint(texture), C.GLenum(target), C.GLenum(pname), (*C.GLfloat)(&params[0]))
}

func (texture Texture) GetTextureParameteriv(target GLenum, pname GLenum, params []int32) {
	C.glGetTextureParameterivEXT(C.GLuint(texture), C.GLenum(target), C.GLenum(pname), (*C.GLint)(&params[0]))
}

func (vaobj VertexArray) GetVertexArrayIntegeri_v(index uint, pname GLenum, param []int32) {
	C.glGetVertexArrayIntegeri_vEXT(C.GLuint(vaobj), C.GLuint(index), C.GLenum(pname), (*C.GLint)(&param[0]))
}

func (vaobj VertexArray) GetVertexArrayIntegerv(pname GLenum, param []int32) {
	C.glGetVertexArrayIntegervEXT(C.GLuint(vaobj), C.GLenum(pname), (*C.GLint)(&param[0]))
}

// Slice?
func (vaobj VertexArray) GetVertexArrayPointeri_v(index uint, pname GLenum) (ptr unsafe.Pointer) {
	C.glGetVertexArrayPointeri_vEXT(C.GLuint(vaobj), C.GLuint(index), C.GLenum(pname), &ptr)
	return
}

// Slice?
func (vaobj VertexArray) GetVertexArrayPointerv(pname GLenum) (ptr unsafe.Pointer) {
	C.glGetVertexArrayPointervEXT(C.GLuint(vaobj), C.GLenum(pname), &ptr)
	return
}

// Slice?
func (buffer Buffer) MapNamedBuffer(access GLenum) unsafe.Pointer {
	return C.glMapNamedBufferEXT(C.GLuint(buffer), C.GLenum(access))
}

// Slice?
func (buffer Buffer) MapNamedBufferRange(offset uintptr, length uintptr, access uint32) unsafe.Pointer {
	return C.glMapNamedBufferRangeEXT(C.GLuint(buffer), C.GLintptr(offset), C.GLsizeiptr(length), C.GLbitfield(access))
}

func MatrixFrustum(matrixMode GLenum, l float64, r float64, b float64, t float64, n float64, f float64) {
	C.glMatrixFrustumEXT(C.GLenum(matrixMode), C.GLdouble(l), C.GLdouble(r), C.GLdouble(b), C.GLdouble(t), C.GLdouble(n), C.GLdouble(f))
}

func MatrixLoadIdentity(matrixMode GLenum) {
	C.glMatrixLoadIdentityEXT(C.GLenum(matrixMode))
}

func MatrixLoadTransposed(matrixMode GLenum, m *[16]float64) {
	C.glMatrixLoadTransposedEXT(C.GLenum(matrixMode), (*C.GLdouble)(&m[0]))
}

func MatrixLoadTransposef(matrixMode GLenum, m *[16]float32) {
	C.glMatrixLoadTransposefEXT(C.GLenum(matrixMode), (*C.GLfloat)(&m[0]))
}

func MatrixLoadd(matrixMode GLenum, m *[16]float64) {
	C.glMatrixLoaddEXT(C.GLenum(matrixMode), (*C.GLdouble)(&m[0]))
}

func MatrixLoadf(matrixMode GLenum, m *[16]float32) {
	C.glMatrixLoadfEXT(C.GLenum(matrixMode), (*C.GLfloat)(&m[0]))
}

func MatrixMultTransposed(matrixMode GLenum, m *[16]float64) {
	C.glMatrixMultTransposedEXT(C.GLenum(matrixMode), (*C.GLdouble)(&m[0]))
}

func MatrixMultTransposef(matrixMode GLenum, m *[16]float32) {
	C.glMatrixMultTransposefEXT(C.GLenum(matrixMode), (*C.GLfloat)(&m[0]))
}

func MatrixMultd(matrixMode GLenum, m *[16]float64) {
	C.glMatrixMultdEXT(C.GLenum(matrixMode), (*C.GLdouble)(&m[0]))
}

func MatrixMultf(matrixMode GLenum, m *[16]float32) {
	C.glMatrixMultfEXT(C.GLenum(matrixMode), (*C.GLfloat)(&m[0]))
}

func MatrixOrtho(matrixMode GLenum, l float64, r float64, b float64, t float64, n float64, f float64) {
	C.glMatrixOrthoEXT(C.GLenum(matrixMode), C.GLdouble(l), C.GLdouble(r), C.GLdouble(b), C.GLdouble(t), C.GLdouble(n), C.GLdouble(f))
}

func MatrixPop(matrixMode GLenum) {
	C.glMatrixPopEXT(C.GLenum(matrixMode))
}

func MatrixPush(matrixMode GLenum) {
	C.glMatrixPushEXT(C.GLenum(matrixMode))
}

func MatrixRotated(matrixMode GLenum, angle float64, x float64, y float64, z float64) {
	C.glMatrixRotatedEXT(C.GLenum(matrixMode), C.GLdouble(angle), C.GLdouble(x), C.GLdouble(y), C.GLdouble(z))
}

func MatrixRotatef(matrixMode GLenum, angle float32, x float32, y float32, z float32) {
	C.glMatrixRotatefEXT(C.GLenum(matrixMode), C.GLfloat(angle), C.GLfloat(x), C.GLfloat(y), C.GLfloat(z))
}

func MatrixScaled(matrixMode GLenum, x float64, y float64, z float64) {
	C.glMatrixScaledEXT(C.GLenum(matrixMode), C.GLdouble(x), C.GLdouble(y), C.GLdouble(z))
}

func MatrixScalef(matrixMode GLenum, x float32, y float32, z float32) {
	C.glMatrixScalefEXT(C.GLenum(matrixMode), C.GLfloat(x), C.GLfloat(y), C.GLfloat(z))
}

func MatrixTranslated(matrixMode GLenum, x float64, y float64, z float64) {
	C.glMatrixTranslatedEXT(C.GLenum(matrixMode), C.GLdouble(x), C.GLdouble(y), C.GLdouble(z))
}

func MatrixTranslatef(matrixMode GLenum, x float32, y float32, z float32) {
	C.glMatrixTranslatefEXT(C.GLenum(matrixMode), C.GLfloat(x), C.GLfloat(y), C.GLfloat(z))
}

func MultiTexBuffer(texunit GLenum, target GLenum, internalformat GLenum, buffer Buffer) {
	C.glMultiTexBufferEXT(C.GLenum(texunit), C.GLenum(target), C.GLenum(internalformat), C.GLuint(buffer))
}

func MultiTexCoordPointer(texunit GLenum, size int, typ GLenum, stride int, pointer interface{}) {
	C.glMultiTexCoordPointerEXT(C.GLenum(texunit), C.GLint(size), C.GLenum(typ), C.GLsizei(stride), glPointer(pointer))
}

func MultiTexEnvf(texunit GLenum, target GLenum, pname GLenum, param float32) {
	C.glMultiTexEnvfEXT(C.GLenum(texunit), C.GLenum(target), C.GLenum(pname), C.GLfloat(param))
}

func MultiTexEnvfv(texunit GLenum, target GLenum, pname GLenum, params []float32) {
	C.glMultiTexEnvfvEXT(C.GLenum(texunit), C.GLenum(target), C.GLenum(pname), (*C.GLfloat)(&params[0]))
}

func MultiTexEnvi(texunit GLenum, target GLenum, pname GLenum, param int) {
	C.glMultiTexEnviEXT(C.GLenum(texunit), C.GLenum(target), C.GLenum(pname), C.GLint(param))
}

func MultiTexEnviv(texunit GLenum, target GLenum, pname GLenum, params []int32) {
	C.glMultiTexEnvivEXT(C.GLenum(texunit), C.GLenum(target), C.GLenum(pname), (*C.GLint)(&params[0]))
}

func MultiTexGend(texunit GLenum, coord GLenum, pname GLenum, param float64) {
	C.glMultiTexGendEXT(C.GLenum(texunit), C.GLenum(coord), C.GLenum(pname), C.GLdouble(param))
}

func MultiTexGendv(texunit GLenum, coord GLenum, pname GLenum, params []float64) {
	C.glMultiTexGendvEXT(C.GLenum(texunit), C.GLenum(coord), C.GLenum(pname), (*C.GLdouble)(&params[0]))
}

func MultiTexGenf(texunit GLenum, coord GLenum, pname GLenum, param float32) {
	C.glMultiTexGenfEXT(C.GLenum(texunit), C.GLenum(coord), C.GLenum(pname), C.GLfloat(param))
}

func MultiTexGenfv(texunit GLenum, coord GLenum, pname GLenum, params []float32) {
	C.glMultiTexGenfvEXT(C.GLenum(texunit), C.GLenum(coord), C.GLenum(pname), (*C.GLfloat)(&params[0]))
}

func MultiTexGeni(texunit GLenum, coord GLenum, pname GLenum, param int) {
	C.glMultiTexGeniEXT(C.GLenum(texunit), C.GLenum(coord), C.GLenum(pname), C.GLint(param))
}

func MultiTexGeniv(texunit GLenum, coord GLenum, pname GLenum, params []int32) {
	C.glMultiTexGenivEXT(C.GLenum(texunit), C.GLenum(coord), C.GLenum(pname), (*C.GLint)(&params[0]))
}

func MultiTexImage1D(texunit GLenum, target GLenum, level int, internalformat int, width int, border int, format GLenum, typ GLenum, pixels interface{}) {
	C.glMultiTexImage1DEXT(C.GLenum(texunit), C.GLenum(target), C.GLint(level), C.GLint(internalformat), C.GLsizei(width), C.GLint(border), C.GLenum(format), C.GLenum(typ), glPointer(pixels))
}

func MultiTexImage2D(texunit GLenum, target GLenum, level int, internalformat int, width int, height int, border int, format GLenum, typ GLenum, pixels interface{}) {
	C.glMultiTexImage2DEXT(C.GLenum(texunit), C.GLenum(target), C.GLint(level), C.GLint(internalformat), C.GLsizei(width), C.GLsizei(height), C.GLint(border), C.GLenum(format), C.GLenum(typ), glPointer(pixels))
}

func MultiTexImage3D(texunit GLenum, target GLenum, level int, internalformat int, width int, height int, depth int, border int, format GLenum, typ GLenum, pixels interface{}) {
	C.glMultiTexImage3DEXT(C.GLenum(texunit), C.GLenum(target), C.GLint(level), C.GLint(internalformat), C.GLsizei(width), C.GLsizei(height), C.GLsizei(depth), C.GLint(border), C.GLenum(format), C.GLenum(typ), glPointer(pixels))
}

func MultiTexParameterIiv(texunit GLenum, target GLenum, pname GLenum, params []int32) {
	C.glMultiTexParameterIivEXT(C.GLenum(texunit), C.GLenum(target), C.GLenum(pname), (*C.GLint)(&params[0]))
}

func MultiTexParameterIuiv(texunit GLenum, target GLenum, pname GLenum, params []uint32) {
	C.glMultiTexParameterIuivEXT(C.GLenum(texunit), C.GLenum(target), C.GLenum(pname), (*C.GLuint)(&params[0]))
}

func MultiTexParameterf(texunit GLenum, target GLenum, pname GLenum, param float32) {
	C.glMultiTexParameterfEXT(C.GLenum(texunit), C.GLenum(target), C.GLenum(pname), C.GLfloat(param))
}

func MultiTexParameterfv(texunit GLenum, target GLenum, pname GLenum, param []float32) {
	C.glMultiTexParameterfvEXT(C.GLenum(texunit), C.GLenum(target), C.GLenum(pname), (*C.GLfloat)(&param[0]))
}

func MultiTexParameteri(texunit GLenum, target GLenum, pname GLenum, param int) {
	C.glMultiTexParameteriEXT(C.GLenum(texunit), C.GLenum(target), C.GLenum(pname), C.GLint(param))
}

func MultiTexParameteriv(texunit GLenum, target GLenum, pname GLenum, param []int32) {
	C.glMultiTexParameterivEXT(C.GLenum(texunit), C.GLenum(target), C.GLenum(pname), (*C.GLint)(&param[0]))
}

func MultiTexRenderbuffer(texunit GLenum, target GLenum, renderbuffer Renderbuffer) {
	C.glMultiTexRenderbufferEXT(C.GLenum(texunit), C.GLenum(target), C.GLuint(renderbuffer))
}

func MultiTexSubImage1D(texunit GLenum, target GLenum, level int, xoffset int, width int, format GLenum, typ GLenum, pixels interface{}) {
	C.glMultiTexSubImage1DEXT(C.GLenum(texunit), C.GLenum(target), C.GLint(level), C.GLint(xoffset), C.GLsizei(width), C.GLenum(format), C.GLenum(typ), glPointer(pixels))
}

func MultiTexSubImage2D(texunit GLenum, target GLenum, level int, xoffset int, yoffset int, width int, height int, format GLenum, typ GLenum, pixels interface{}) {
	C.glMultiTexSubImage2DEXT(C.GLenum(texunit), C.GLenum(target), C.GLint(level), C.GLint(xoffset), C.GLint(yoffset), C.GLsizei(width), C.GLsizei(height), C.GLenum(format), C.GLenum(typ), glPointer(pixels))
}

func MultiTexSubImage3D(texunit GLenum, target GLenum, level int, xoffset int, yoffset int, zoffset int, width int, height int, depth int, format GLenum, typ GLenum, pixels interface{}) {
	C.glMultiTexSubImage3DEXT(C.GLenum(texunit), C.GLenum(target), C.GLint(level), C.GLint(xoffset), C.GLint(yoffset), C.GLint(zoffset), C.GLsizei(width), C.GLsizei(height), C.GLsizei(depth), C.GLenum(format), C.GLenum(typ), glPointer(pixels))
}

func (buffer Buffer) NamedBufferData(size uintptr, data interface{}, usage GLenum) {
	C.glNamedBufferDataEXT(C.GLuint(buffer), C.GLsizeiptr(size), glPointer(data), C.GLenum(usage))
}

func (buffer Buffer) NamedBufferSubData(offset uintptr, size uintptr, data interface{}) {
	C.glNamedBufferSubDataEXT(C.GLuint(buffer), C.GLintptr(offset), C.GLsizeiptr(size), glPointer(data))
}

func NamedCopyBufferSubData(readBuffer Buffer, writeBuffer Buffer, readOffset uintptr, writeOffset uintptr, size uintptr) {
	C.glNamedCopyBufferSubDataEXT(C.GLuint(readBuffer), C.GLuint(writeBuffer), C.GLintptr(readOffset), C.GLintptr(writeOffset), C.GLsizeiptr(size))
}

func (framebuffer Framebuffer) NamedFramebufferRenderbuffer(attachment GLenum, renderbuffertarget GLenum, renderbuffer Renderbuffer) {
	C.glNamedFramebufferRenderbufferEXT(C.GLuint(framebuffer), C.GLenum(attachment), C.GLenum(renderbuffertarget), C.GLuint(renderbuffer))
}

func (framebuffer Framebuffer) NamedFramebufferTexture1D(attachment GLenum, textarget GLenum, texture Texture, level int) {
	C.glNamedFramebufferTexture1DEXT(C.GLuint(framebuffer), C.GLenum(attachment), C.GLenum(textarget), C.GLuint(texture), C.GLint(level))
}

func (framebuffer Framebuffer) NamedFramebufferTexture2D(attachment GLenum, textarget GLenum, texture Texture, level int) {
	C.glNamedFramebufferTexture2DEXT(C.GLuint(framebuffer), C.GLenum(attachment), C.GLenum(textarget), C.GLuint(texture), C.GLint(level))
}

func (framebuffer Framebuffer) NamedFramebufferTexture3D(attachment GLenum, textarget GLenum, texture Texture, level int, zoffset int) {
	C.glNamedFramebufferTexture3DEXT(C.GLuint(framebuffer), C.GLenum(attachment), C.GLenum(textarget), C.GLuint(texture), C.GLint(level), C.GLint(zoffset))
}

func (framebuffer Framebuffer) NamedFramebufferTexture(attachment GLenum, texture Texture, level int) {
	C.glNamedFramebufferTextureEXT(C.GLuint(framebuffer), C.GLenum(attachment), C.GLuint(texture), C.GLint(level))
}

func (framebuffer Framebuffer) NamedFramebufferTextureFace(attachment GLenum, texture Texture, level int, face GLenum) {
	C.glNamedFramebufferTextureFaceEXT(C.GLuint(framebuffer), C.GLenum(attachment), C.GLuint(texture), C.GLint(level), C.GLenum(face))
}

func (framebuffer Framebuffer) NamedFramebufferTextureLayer(attachment GLenum, texture Texture, level int, layer int) {
	C.glNamedFramebufferTextureLayerEXT(C.GLuint(framebuffer), C.GLenum(attachment), C.GLuint(texture), C.GLint(level), C.GLint(layer))
}

func (program Program) NamedProgramLocalParameter4d(target GLenum, index uint, x float64, y float64, z float64, w float64) {
	C.glNamedProgramLocalParameter4dEXT(C.GLuint(program), C.GLenum(target), C.GLuint(index), C.GLdouble(x), C.GLdouble(y), C.GLdouble(z), C.GLdouble(w))
}

func (program Program) NamedProgramLocalParameter4dv(target GLenum, index uint, params []float64) {
	C.glNamedProgramLocalParameter4dvEXT(C.GLuint(program), C.GLenum(target), C.GLuint(index), (*C.GLdouble)(&params[0]))
}

func (program Program) NamedProgramLocalParameter4f(target GLenum, index uint, x float32, y float32, z float32, w float32) {
	C.glNamedProgramLocalParameter4fEXT(C.GLuint(program), C.GLenum(target), C.GLuint(index), C.GLfloat(x), C.GLfloat(y), C.GLfloat(z), C.GLfloat(w))
}

func (program Program) NamedProgramLocalParameter4fv(target GLenum, index uint, params []float32) {
	C.glNamedProgramLocalParameter4fvEXT(C.GLuint(program), C.GLenum(target), C.GLuint(index), (*C.GLfloat)(&params[0]))
}

func (program Program) NamedProgramLocalParameterI4i(target GLenum, index uint, x int, y int, z int, w int) {
	C.glNamedProgramLocalParameterI4iEXT(C.GLuint(program), C.GLenum(target), C.GLuint(index), C.GLint(x), C.GLint(y), C.GLint(z), C.GLint(w))
}

func (program Program) NamedProgramLocalParameterI4iv(target GLenum, index uint, params []int32) {
	C.glNamedProgramLocalParameterI4ivEXT(C.GLuint(program), C.GLenum(target), C.GLuint(index), (*C.GLint)(&params[0]))
}

func (program Program) NamedProgramLocalParameterI4ui(target GLenum, index uint, x uint, y uint, z uint, w uint) {
	C.glNamedProgramLocalParameterI4uiEXT(C.GLuint(program), C.GLenum(target), C.GLuint(index), C.GLuint(x), C.GLuint(y), C.GLuint(z), C.GLuint(w))
}

func (program Program) NamedProgramLocalParameterI4uiv(target GLenum, index uint, params []uint32) {
	C.glNamedProgramLocalParameterI4uivEXT(C.GLuint(program), C.GLenum(target), C.GLuint(index), (*C.GLuint)(&params[0]))
}

func (program Program) NamedProgramLocalParameters4fv(target GLenum, index uint, count int, params []float32) {
	C.glNamedProgramLocalParameters4fvEXT(C.GLuint(program), C.GLenum(target), C.GLuint(index), C.GLsizei(count), (*C.GLfloat)(&params[0]))
}

func (program Program) NamedProgramLocalParametersI4iv(target GLenum, index uint, count int, params []int32) {
	C.glNamedProgramLocalParametersI4ivEXT(C.GLuint(program), C.GLenum(target), C.GLuint(index), C.GLsizei(count), (*C.GLint)(&params[0]))
}

func (program Program) NamedProgramLocalParametersI4uiv(target GLenum, index uint, count int, params []uint32) {
	C.glNamedProgramLocalParametersI4uivEXT(C.GLuint(program), C.GLenum(target), C.GLuint(index), C.GLsizei(count), (*C.GLuint)(&params[0]))
}

// TODO
// func (program Program) NamedProgramString(target GLenum, format GLenum, len int, GLvoid const *string) {
// 	C.glNamedProgramStringEXT(C.GLuint(program), C.GLenum(target), C.GLenum(format), C.GLsizei(len), C.const(GLvoid) *string)
// }

func (renderbuffer Renderbuffer) NamedRenderbufferStorage(internalformat GLenum, width int, height int) {
	C.glNamedRenderbufferStorageEXT(C.GLuint(renderbuffer), C.GLenum(internalformat), C.GLsizei(width), C.GLsizei(height))
}

func (renderbuffer Renderbuffer) NamedRenderbufferStorageMultisampleCoverage(coverageSamples int, colorSamples int, internalformat GLenum, width int, height int) {
	C.glNamedRenderbufferStorageMultisampleCoverageEXT(C.GLuint(renderbuffer), C.GLsizei(coverageSamples), C.GLsizei(colorSamples), C.GLenum(internalformat), C.GLsizei(width), C.GLsizei(height))
}

func (renderbuffer Renderbuffer) NamedRenderbufferStorageMultisample(samples int, internalformat GLenum, width int, height int) {
	C.glNamedRenderbufferStorageMultisampleEXT(C.GLuint(renderbuffer), C.GLsizei(samples), C.GLenum(internalformat), C.GLsizei(width), C.GLsizei(height))
}

func (program Program) ProgramUniform1f(location UniformLocation, v0 float32) {
	C.glProgramUniform1fEXT(C.GLuint(program), C.GLint(location), C.GLfloat(v0))
}

func (program Program) ProgramUniform1fv(location UniformLocation, count int, value []float32) {
	C.glProgramUniform1fvEXT(C.GLuint(program), C.GLint(location), C.GLsizei(count), (*C.GLfloat)(&value[0]))
}

func (program Program) ProgramUniform1i(location UniformLocation, v0 int) {
	C.glProgramUniform1iEXT(C.GLuint(program), C.GLint(location), C.GLint(v0))
}

func (program Program) ProgramUniform1iv(location UniformLocation, count int, value []int32) {
	C.glProgramUniform1ivEXT(C.GLuint(program), C.GLint(location), C.GLsizei(count), (*C.GLint)(&value[0]))
}

func (program Program) ProgramUniform1ui(location UniformLocation, v0 uint) {
	C.glProgramUniform1uiEXT(C.GLuint(program), C.GLint(location), C.GLuint(v0))
}

func (program Program) ProgramUniform1uiv(location UniformLocation, count int, value []uint32) {
	C.glProgramUniform1uivEXT(C.GLuint(program), C.GLint(location), C.GLsizei(count), (*C.GLuint)(&value[0]))
}

func (program Program) ProgramUniform2f(location UniformLocation, v0 float32, v1 float32) {
	C.glProgramUniform2fEXT(C.GLuint(program), C.GLint(location), C.GLfloat(v0), C.GLfloat(v1))
}

func (program Program) ProgramUniform2fv(location UniformLocation, count int, value []float32) {
	C.glProgramUniform2fvEXT(C.GLuint(program), C.GLint(location), C.GLsizei(count), (*C.GLfloat)(&value[0]))
}

func (program Program) ProgramUniform2i(location UniformLocation, v0 int, v1 int) {
	C.glProgramUniform2iEXT(C.GLuint(program), C.GLint(location), C.GLint(v0), C.GLint(v1))
}

func (program Program) ProgramUniform2iv(location UniformLocation, count int, value []int32) {
	C.glProgramUniform2ivEXT(C.GLuint(program), C.GLint(location), C.GLsizei(count), (*C.GLint)(&value[0]))
}

func (program Program) ProgramUniform2ui(location UniformLocation, v0 uint, v1 uint) {
	C.glProgramUniform2uiEXT(C.GLuint(program), C.GLint(location), C.GLuint(v0), C.GLuint(v1))
}

func (program Program) ProgramUniform2uiv(location UniformLocation, count int, value []uint32) {
	C.glProgramUniform2uivEXT(C.GLuint(program), C.GLint(location), C.GLsizei(count), (*C.GLuint)(&value[0]))
}

func (program Program) ProgramUniform3f(location UniformLocation, v0 float32, v1 float32, v2 float32) {
	C.glProgramUniform3fEXT(C.GLuint(program), C.GLint(location), C.GLfloat(v0), C.GLfloat(v1), C.GLfloat(v2))
}

func (program Program) ProgramUniform3fv(location UniformLocation, count int, value []float32) {
	C.glProgramUniform3fvEXT(C.GLuint(program), C.GLint(location), C.GLsizei(count), (*C.GLfloat)(&value[0]))
}

func (program Program) ProgramUniform3i(location UniformLocation, v0 int, v1 int, v2 int) {
	C.glProgramUniform3iEXT(C.GLuint(program), C.GLint(location), C.GLint(v0), C.GLint(v1), C.GLint(v2))
}

func (program Program) ProgramUniform3iv(location UniformLocation, count int, value []int32) {
	C.glProgramUniform3ivEXT(C.GLuint(program), C.GLint(location), C.GLsizei(count), (*C.GLint)(&value[0]))
}

func (program Program) ProgramUniform3ui(location UniformLocation, v0 uint, v1 uint, v2 uint) {
	C.glProgramUniform3uiEXT(C.GLuint(program), C.GLint(location), C.GLuint(v0), C.GLuint(v1), C.GLuint(v2))
}

func (program Program) ProgramUniform3uiv(location UniformLocation, count int, value []uint32) {
	C.glProgramUniform3uivEXT(C.GLuint(program), C.GLint(location), C.GLsizei(count), (*C.GLuint)(&value[0]))
}

func (program Program) ProgramUniform4f(location UniformLocation, v0 float32, v1 float32, v2 float32, v3 float32) {
	C.glProgramUniform4fEXT(C.GLuint(program), C.GLint(location), C.GLfloat(v0), C.GLfloat(v1), C.GLfloat(v2), C.GLfloat(v3))
}

func (program Program) ProgramUniform4fv(location UniformLocation, count int, value []float32) {
	C.glProgramUniform4fvEXT(C.GLuint(program), C.GLint(location), C.GLsizei(count), (*C.GLfloat)(&value[0]))
}

func (program Program) ProgramUniform4i(location UniformLocation, v0 int, v1 int, v2 int, v3 int) {
	C.glProgramUniform4iEXT(C.GLuint(program), C.GLint(location), C.GLint(v0), C.GLint(v1), C.GLint(v2), C.GLint(v3))
}

func (program Program) ProgramUniform4iv(location UniformLocation, count int, value []int32) {
	C.glProgramUniform4ivEXT(C.GLuint(program), C.GLint(location), C.GLsizei(count), (*C.GLint)(&value[0]))
}

func (program Program) ProgramUniform4ui(location UniformLocation, v0 uint, v1 uint, v2 uint, v3 uint) {
	C.glProgramUniform4uiEXT(C.GLuint(program), C.GLint(location), C.GLuint(v0), C.GLuint(v1), C.GLuint(v2), C.GLuint(v3))
}

func (program Program) ProgramUniform4uiv(location UniformLocation, count int, value []uint32) {
	C.glProgramUniform4uivEXT(C.GLuint(program), C.GLint(location), C.GLsizei(count), (*C.GLuint)(&value[0]))
}

func (program Program) ProgramUniformMatrix2fv(location UniformLocation, count int, transpose bool, value []float32) {
	C.glProgramUniformMatrix2fvEXT(C.GLuint(program), C.GLint(location), C.GLsizei(count), glBool(transpose), (*C.GLfloat)(&value[0]))
}

func (program Program) ProgramUniformMatrix2x3fv(location UniformLocation, count int, transpose bool, value []float32) {
	C.glProgramUniformMatrix2x3fvEXT(C.GLuint(program), C.GLint(location), C.GLsizei(count), glBool(transpose), (*C.GLfloat)(&value[0]))
}

func (program Program) ProgramUniformMatrix2x4fv(location UniformLocation, count int, transpose bool, value []float32) {
	C.glProgramUniformMatrix2x4fvEXT(C.GLuint(program), C.GLint(location), C.GLsizei(count), glBool(transpose), (*C.GLfloat)(&value[0]))
}

func (program Program) ProgramUniformMatrix3fv(location UniformLocation, count int, transpose bool, value []float32) {
	C.glProgramUniformMatrix3fvEXT(C.GLuint(program), C.GLint(location), C.GLsizei(count), glBool(transpose), (*C.GLfloat)(&value[0]))
}

func (program Program) ProgramUniformMatrix3x2fv(location UniformLocation, count int, transpose bool, value []float32) {
	C.glProgramUniformMatrix3x2fvEXT(C.GLuint(program), C.GLint(location), C.GLsizei(count), glBool(transpose), (*C.GLfloat)(&value[0]))
}

func (program Program) ProgramUniformMatrix3x4fv(location UniformLocation, count int, transpose bool, value []float32) {
	C.glProgramUniformMatrix3x4fvEXT(C.GLuint(program), C.GLint(location), C.GLsizei(count), glBool(transpose), (*C.GLfloat)(&value[0]))
}

func (program Program) ProgramUniformMatrix4fv(location UniformLocation, count int, transpose bool, value []float32) {
	C.glProgramUniformMatrix4fvEXT(C.GLuint(program), C.GLint(location), C.GLsizei(count), glBool(transpose), (*C.GLfloat)(&value[0]))
}

func (program Program) ProgramUniformMatrix4x2fv(location UniformLocation, count int, transpose bool, value []float32) {
	C.glProgramUniformMatrix4x2fvEXT(C.GLuint(program), C.GLint(location), C.GLsizei(count), glBool(transpose), (*C.GLfloat)(&value[0]))
}

func (program Program) ProgramUniformMatrix4x3fv(location UniformLocation, count int, transpose bool, value []float32) {
	C.glProgramUniformMatrix4x3fvEXT(C.GLuint(program), C.GLint(location), C.GLsizei(count), glBool(transpose), (*C.GLfloat)(&value[0]))
}

func PushClientAttribDefault(mask uint32) {
	C.glPushClientAttribDefaultEXT(C.GLbitfield(mask))
}

func (texture Texture) TextureBuffer(target GLenum, internalformat GLenum, buffer Buffer) {
	C.glTextureBufferEXT(C.GLuint(texture), C.GLenum(target), C.GLenum(internalformat), C.GLuint(buffer))
}

func (texture Texture) TextureImage1D(target GLenum, level int, internalformat int, width int, border int, format GLenum, typ GLenum, pixels interface{}) {
	C.glTextureImage1DEXT(C.GLuint(texture), C.GLenum(target), C.GLint(level), C.GLint(internalformat), C.GLsizei(width), C.GLint(border), C.GLenum(format), C.GLenum(typ), glPointer(pixels))
}

func (texture Texture) TextureImage2D(target GLenum, level int, internalformat int, width int, height int, border int, format GLenum, typ GLenum, pixels interface{}) {
	C.glTextureImage2DEXT(C.GLuint(texture), C.GLenum(target), C.GLint(level), C.GLint(internalformat), C.GLsizei(width), C.GLsizei(height), C.GLint(border), C.GLenum(format), C.GLenum(typ), glPointer(pixels))
}

func (texture Texture) TextureImage3D(target GLenum, level int, internalformat int, width int, height int, depth int, border int, format GLenum, typ GLenum, pixels interface{}) {
	C.glTextureImage3DEXT(C.GLuint(texture), C.GLenum(target), C.GLint(level), C.GLint(internalformat), C.GLsizei(width), C.GLsizei(height), C.GLsizei(depth), C.GLint(border), C.GLenum(format), C.GLenum(typ), glPointer(pixels))
}

func (texture Texture) TextureParameterIiv(target GLenum, pname GLenum, params []int32) {
	C.glTextureParameterIivEXT(C.GLuint(texture), C.GLenum(target), C.GLenum(pname), (*C.GLint)(&params[0]))
}

func (texture Texture) TextureParameterIuiv(target GLenum, pname GLenum, params []uint32) {
	C.glTextureParameterIuivEXT(C.GLuint(texture), C.GLenum(target), C.GLenum(pname), (*C.GLuint)(&params[0]))
}

func (texture Texture) TextureParameterf(target GLenum, pname GLenum, param float32) {
	C.glTextureParameterfEXT(C.GLuint(texture), C.GLenum(target), C.GLenum(pname), C.GLfloat(param))
}

func (texture Texture) TextureParameterfv(target GLenum, pname GLenum, param []float32) {
	C.glTextureParameterfvEXT(C.GLuint(texture), C.GLenum(target), C.GLenum(pname), (*C.GLfloat)(&param[0]))
}

func (texture Texture) TextureParameteri(target GLenum, pname GLenum, param int) {
	C.glTextureParameteriEXT(C.GLuint(texture), C.GLenum(target), C.GLenum(pname), C.GLint(param))
}

func (texture Texture) TextureParameteriv(target GLenum, pname GLenum, param []int32) {
	C.glTextureParameterivEXT(C.GLuint(texture), C.GLenum(target), C.GLenum(pname), (*C.GLint)(&param[0]))
}

func (texture Texture) TextureRenderbuffer(target GLenum, renderbuffer Renderbuffer) {
	C.glTextureRenderbufferEXT(C.GLuint(texture), C.GLenum(target), C.GLuint(renderbuffer))
}

func (texture Texture) TextureSubImage1D(target GLenum, level int, xoffset int, width int, format GLenum, typ GLenum, pixels interface{}) {
	C.glTextureSubImage1DEXT(C.GLuint(texture), C.GLenum(target), C.GLint(level), C.GLint(xoffset), C.GLsizei(width), C.GLenum(format), C.GLenum(typ), glPointer(pixels))
}

func (texture Texture) TextureSubImage2D(target GLenum, level int, xoffset int, yoffset int, width int, height int, format GLenum, typ GLenum, pixels interface{}) {
	C.glTextureSubImage2DEXT(C.GLuint(texture), C.GLenum(target), C.GLint(level), C.GLint(xoffset), C.GLint(yoffset), C.GLsizei(width), C.GLsizei(height), C.GLenum(format), C.GLenum(typ), glPointer(pixels))
}

func (texture Texture) TextureSubImage3D(target GLenum, level int, xoffset int, yoffset int, zoffset int, width int, height int, depth int, format GLenum, typ GLenum, pixels interface{}) {
	C.glTextureSubImage3DEXT(C.GLuint(texture), C.GLenum(target), C.GLint(level), C.GLint(xoffset), C.GLint(yoffset), C.GLint(zoffset), C.GLsizei(width), C.GLsizei(height), C.GLsizei(depth), C.GLenum(format), C.GLenum(typ), glPointer(pixels))
}

func (buffer Buffer) UnmapNamedBuffer() bool {
	return goBool(C.glUnmapNamedBufferEXT(C.GLuint(buffer)))
}

func (vaobj VertexArray) VertexArrayColorOffset(buffer Buffer, size int, typ GLenum, stride int, offset uintptr) {
	C.glVertexArrayColorOffsetEXT(C.GLuint(vaobj), C.GLuint(buffer), C.GLint(size), C.GLenum(typ), C.GLsizei(stride), C.GLintptr(offset))
}

func (vaobj VertexArray) VertexArrayEdgeFlagOffset(buffer Buffer, stride int, offset uintptr) {
	C.glVertexArrayEdgeFlagOffsetEXT(C.GLuint(vaobj), C.GLuint(buffer), C.GLsizei(stride), C.GLintptr(offset))
}

func (vaobj VertexArray) VertexArrayFogCoordOffset(buffer Buffer, typ GLenum, stride int, offset uintptr) {
	C.glVertexArrayFogCoordOffsetEXT(C.GLuint(vaobj), C.GLuint(buffer), C.GLenum(typ), C.GLsizei(stride), C.GLintptr(offset))
}

func (vaobj VertexArray) VertexArrayIndexOffset(buffer Buffer, typ GLenum, stride int, offset uintptr) {
	C.glVertexArrayIndexOffsetEXT(C.GLuint(vaobj), C.GLuint(buffer), C.GLenum(typ), C.GLsizei(stride), C.GLintptr(offset))
}

func (vaobj VertexArray) VertexArrayMultiTexCoordOffset(buffer Buffer, texunit GLenum, size int, typ GLenum, stride int, offset uintptr) {
	C.glVertexArrayMultiTexCoordOffsetEXT(C.GLuint(vaobj), C.GLuint(buffer), C.GLenum(texunit), C.GLint(size), C.GLenum(typ), C.GLsizei(stride), C.GLintptr(offset))
}

func (vaobj VertexArray) VertexArrayNormalOffset(buffer Buffer, typ GLenum, stride int, offset uintptr) {
	C.glVertexArrayNormalOffsetEXT(C.GLuint(vaobj), C.GLuint(buffer), C.GLenum(typ), C.GLsizei(stride), C.GLintptr(offset))
}

func (vaobj VertexArray) VertexArraySecondaryColorOffset(buffer Buffer, size int, typ GLenum, stride int, offset uintptr) {
	C.glVertexArraySecondaryColorOffsetEXT(C.GLuint(vaobj), C.GLuint(buffer), C.GLint(size), C.GLenum(typ), C.GLsizei(stride), C.GLintptr(offset))
}

func (vaobj VertexArray) VertexArrayTexCoordOffset(buffer Buffer, size int, typ GLenum, stride int, offset uintptr) {
	C.glVertexArrayTexCoordOffsetEXT(C.GLuint(vaobj), C.GLuint(buffer), C.GLint(size), C.GLenum(typ), C.GLsizei(stride), C.GLintptr(offset))
}

func (vaobj VertexArray) VertexArrayVertexAttribIOffset(buffer Buffer, index uint, size int, typ GLenum, stride int, offset uintptr) {
	C.glVertexArrayVertexAttribIOffsetEXT(C.GLuint(vaobj), C.GLuint(buffer), C.GLuint(index), C.GLint(size), C.GLenum(typ), C.GLsizei(stride), C.GLintptr(offset))
}

func (vaobj VertexArray) VertexArrayVertexAttribOffset(buffer Buffer, index uint, size int, typ GLenum, normalized bool, stride int, offset uintptr) {
	C.glVertexArrayVertexAttribOffsetEXT(C.GLuint(vaobj), C.GLuint(buffer), C.GLuint(index), C.GLint(size), C.GLenum(typ), glBool(normalized), C.GLsizei(stride), C.GLintptr(offset))
}

func (vaobj VertexArray) VertexArrayVertexOffset(buffer Buffer, size int, typ GLenum, stride int, offset uintptr) {
	C.glVertexArrayVertexOffsetEXT(C.GLuint(vaobj), C.GLuint(buffer), C.GLint(size), C.GLenum(typ), C.GLsizei(stride), C.GLintptr(offset))
}
