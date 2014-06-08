// Copyright 2012 The go-gl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gl

// #include "gl.h"
import "C"

const (
	COLOR_ATTACHMENT0                            = C.GL_COLOR_ATTACHMENT0
	COLOR_ATTACHMENT1                            = C.GL_COLOR_ATTACHMENT1
	COLOR_ATTACHMENT10                           = C.GL_COLOR_ATTACHMENT10
	COLOR_ATTACHMENT11                           = C.GL_COLOR_ATTACHMENT11
	COLOR_ATTACHMENT12                           = C.GL_COLOR_ATTACHMENT12
	COLOR_ATTACHMENT13                           = C.GL_COLOR_ATTACHMENT13
	COLOR_ATTACHMENT14                           = C.GL_COLOR_ATTACHMENT14
	COLOR_ATTACHMENT15                           = C.GL_COLOR_ATTACHMENT15
	COLOR_ATTACHMENT2                            = C.GL_COLOR_ATTACHMENT2
	COLOR_ATTACHMENT3                            = C.GL_COLOR_ATTACHMENT3
	COLOR_ATTACHMENT4                            = C.GL_COLOR_ATTACHMENT4
	COLOR_ATTACHMENT5                            = C.GL_COLOR_ATTACHMENT5
	COLOR_ATTACHMENT6                            = C.GL_COLOR_ATTACHMENT6
	COLOR_ATTACHMENT7                            = C.GL_COLOR_ATTACHMENT7
	COLOR_ATTACHMENT8                            = C.GL_COLOR_ATTACHMENT8
	COLOR_ATTACHMENT9                            = C.GL_COLOR_ATTACHMENT9
	DEPTH24_STENCIL8                             = C.GL_DEPTH24_STENCIL8
	DEPTH_ATTACHMENT                             = C.GL_DEPTH_ATTACHMENT
	DEPTH_STENCIL                                = C.GL_DEPTH_STENCIL
	DEPTH_STENCIL_ATTACHMENT                     = C.GL_DEPTH_STENCIL_ATTACHMENT
	DRAW_FRAMEBUFFER                             = C.GL_DRAW_FRAMEBUFFER
	DRAW_FRAMEBUFFER_BINDING                     = C.GL_DRAW_FRAMEBUFFER_BINDING
	FRAMEBUFFER                                  = C.GL_FRAMEBUFFER
	FRAMEBUFFER_ATTACHMENT_ALPHA_SIZE            = C.GL_FRAMEBUFFER_ATTACHMENT_ALPHA_SIZE
	FRAMEBUFFER_ATTACHMENT_BLUE_SIZE             = C.GL_FRAMEBUFFER_ATTACHMENT_BLUE_SIZE
	FRAMEBUFFER_ATTACHMENT_COLOR_ENCODING        = C.GL_FRAMEBUFFER_ATTACHMENT_COLOR_ENCODING
	FRAMEBUFFER_ATTACHMENT_COMPONENT_TYPE        = C.GL_FRAMEBUFFER_ATTACHMENT_COMPONENT_TYPE
	FRAMEBUFFER_ATTACHMENT_DEPTH_SIZE            = C.GL_FRAMEBUFFER_ATTACHMENT_DEPTH_SIZE
	FRAMEBUFFER_ATTACHMENT_GREEN_SIZE            = C.GL_FRAMEBUFFER_ATTACHMENT_GREEN_SIZE
	FRAMEBUFFER_ATTACHMENT_OBJECT_NAME           = C.GL_FRAMEBUFFER_ATTACHMENT_OBJECT_NAME
	FRAMEBUFFER_ATTACHMENT_OBJECT_TYPE           = C.GL_FRAMEBUFFER_ATTACHMENT_OBJECT_TYPE
	FRAMEBUFFER_ATTACHMENT_RED_SIZE              = C.GL_FRAMEBUFFER_ATTACHMENT_RED_SIZE
	FRAMEBUFFER_ATTACHMENT_STENCIL_SIZE          = C.GL_FRAMEBUFFER_ATTACHMENT_STENCIL_SIZE
	FRAMEBUFFER_ATTACHMENT_TEXTURE_CUBE_MAP_FACE = C.GL_FRAMEBUFFER_ATTACHMENT_TEXTURE_CUBE_MAP_FACE
	FRAMEBUFFER_ATTACHMENT_TEXTURE_LAYER         = C.GL_FRAMEBUFFER_ATTACHMENT_TEXTURE_LAYER
	FRAMEBUFFER_ATTACHMENT_TEXTURE_LEVEL         = C.GL_FRAMEBUFFER_ATTACHMENT_TEXTURE_LEVEL
	FRAMEBUFFER_BINDING                          = C.GL_FRAMEBUFFER_BINDING
	FRAMEBUFFER_COMPLETE                         = C.GL_FRAMEBUFFER_COMPLETE
	FRAMEBUFFER_DEFAULT                          = C.GL_FRAMEBUFFER_DEFAULT
	FRAMEBUFFER_INCOMPLETE_ATTACHMENT            = C.GL_FRAMEBUFFER_INCOMPLETE_ATTACHMENT
	FRAMEBUFFER_INCOMPLETE_DRAW_BUFFER           = C.GL_FRAMEBUFFER_INCOMPLETE_DRAW_BUFFER
	FRAMEBUFFER_INCOMPLETE_MISSING_ATTACHMENT    = C.GL_FRAMEBUFFER_INCOMPLETE_MISSING_ATTACHMENT
	FRAMEBUFFER_INCOMPLETE_MULTISAMPLE           = C.GL_FRAMEBUFFER_INCOMPLETE_MULTISAMPLE
	FRAMEBUFFER_INCOMPLETE_READ_BUFFER           = C.GL_FRAMEBUFFER_INCOMPLETE_READ_BUFFER
	FRAMEBUFFER_UNDEFINED                        = C.GL_FRAMEBUFFER_UNDEFINED
	FRAMEBUFFER_UNSUPPORTED                      = C.GL_FRAMEBUFFER_UNSUPPORTED
	INDEX                                        = C.GL_INDEX
	INVALID_FRAMEBUFFER_OPERATION                = C.GL_INVALID_FRAMEBUFFER_OPERATION
	MAX_COLOR_ATTACHMENTS                        = C.GL_MAX_COLOR_ATTACHMENTS
	MAX_RENDERBUFFER_SIZE                        = C.GL_MAX_RENDERBUFFER_SIZE
	MAX_SAMPLES                                  = C.GL_MAX_SAMPLES
	READ_FRAMEBUFFER                             = C.GL_READ_FRAMEBUFFER
	READ_FRAMEBUFFER_BINDING                     = C.GL_READ_FRAMEBUFFER_BINDING
	RENDERBUFFER                                 = C.GL_RENDERBUFFER
	RENDERBUFFER_ALPHA_SIZE                      = C.GL_RENDERBUFFER_ALPHA_SIZE
	RENDERBUFFER_BINDING                         = C.GL_RENDERBUFFER_BINDING
	RENDERBUFFER_BLUE_SIZE                       = C.GL_RENDERBUFFER_BLUE_SIZE
	RENDERBUFFER_DEPTH_SIZE                      = C.GL_RENDERBUFFER_DEPTH_SIZE
	RENDERBUFFER_GREEN_SIZE                      = C.GL_RENDERBUFFER_GREEN_SIZE
	RENDERBUFFER_HEIGHT                          = C.GL_RENDERBUFFER_HEIGHT
	RENDERBUFFER_INTERNAL_FORMAT                 = C.GL_RENDERBUFFER_INTERNAL_FORMAT
	RENDERBUFFER_RED_SIZE                        = C.GL_RENDERBUFFER_RED_SIZE
	RENDERBUFFER_SAMPLES                         = C.GL_RENDERBUFFER_SAMPLES
	RENDERBUFFER_STENCIL_SIZE                    = C.GL_RENDERBUFFER_STENCIL_SIZE
	RENDERBUFFER_WIDTH                           = C.GL_RENDERBUFFER_WIDTH
	STENCIL_ATTACHMENT                           = C.GL_STENCIL_ATTACHMENT
	STENCIL_INDEX1                               = C.GL_STENCIL_INDEX1
	STENCIL_INDEX16                              = C.GL_STENCIL_INDEX16
	STENCIL_INDEX4                               = C.GL_STENCIL_INDEX4
	STENCIL_INDEX8                               = C.GL_STENCIL_INDEX8
	TEXTURE_STENCIL_SIZE                         = C.GL_TEXTURE_STENCIL_SIZE
	UNSIGNED_INT_24_8                            = C.GL_UNSIGNED_INT_24_8
	UNSIGNED_NORMALIZED                          = C.GL_UNSIGNED_NORMALIZED
)

var ARB_framebuffer_object = false

func init() {
	extensions["GL_ARB_framebuffer_object"] = &ARB_framebuffer_object
}

func GenerateMipmap(target GLenum) {
	C.glGenerateMipmap(C.GLenum(target))
}

func (object Object) IsFramebuffer() bool {
	return C.glIsFramebuffer(C.GLuint(object)) == TRUE
}

func (object Object) IsRenderbuffer() bool {
	return C.glIsRenderbuffer(C.GLuint(object)) == TRUE
}

type Framebuffer Object

func GenFramebuffer() Framebuffer {
	var b C.GLuint
	C.glGenFramebuffers(1, &b)
	return Framebuffer(b)
}

func GenFramebuffers(bufs []Framebuffer) {
	if len(bufs) > 0 {
		C.glGenFramebuffers(C.GLsizei(len(bufs)), (*C.GLuint)(&bufs[0]))
	}
}

func (fb Framebuffer) Bind(target GLenum) {
	C.glBindFramebuffer(C.GLenum(target), C.GLuint(fb))
}

func (fb Framebuffer) Unbind(target GLenum) {
	C.glBindFramebuffer(C.GLenum(target), 0)
}

func (fb Framebuffer) Delete() {
	C.glDeleteFramebuffers(1, (*C.GLuint)(&fb))
}

func BlitFramebuffer(srcX0, srcY0, srcX1, srcY1, dstX0, dstY0, dstX1, dstY1 int, mask GLbitfield, filter GLenum) {
	C.glBlitFramebuffer(C.GLint(srcX0), C.GLint(srcY0), C.GLint(srcX1), C.GLint(srcY1), C.GLint(dstX0), C.GLint(dstY0), C.GLint(dstX1), C.GLint(dstY1), C.GLbitfield(mask), C.GLenum(filter))
}

func CheckFramebufferStatus(target GLenum) GLenum {
	return (GLenum)(C.glCheckFramebufferStatus(C.GLenum(target)))
}

func DeleteFramebuffers(bufs []Framebuffer) {
	if len(bufs) > 0 {
		C.glDeleteFramebuffers(C.GLsizei(len(bufs)), (*C.GLuint)(&bufs[0]))
	}
}

func FramebufferTexture1D(target, attachment, textarget GLenum, texture Texture, level int) {
	C.glFramebufferTexture1D(C.GLenum(target), C.GLenum(attachment), C.GLenum(textarget), C.GLuint(texture), C.GLint(level))
}

func FramebufferTexture2D(target, attachment, textarget GLenum, texture Texture, level int) {
	C.glFramebufferTexture2D(C.GLenum(target), C.GLenum(attachment), C.GLenum(textarget), C.GLuint(texture), C.GLint(level))
}

func FramebufferTexture3D(target, attachment, textarget GLenum, texture Texture, level int, layer int) {
	C.glFramebufferTexture3D(C.GLenum(target), C.GLenum(attachment), C.GLenum(textarget), C.GLuint(texture), C.GLint(level), C.GLint(layer))
}

func FramebufferTextureLayer(target, attachment GLenum, texture Texture, level, layer int) {
	C.glFramebufferTextureLayer(C.GLenum(target), C.GLenum(attachment), C.GLuint(texture), C.GLint(level), C.GLint(layer))
}

type Renderbuffer Object

func GenRenderbuffer() Renderbuffer {
	var b C.GLuint
	C.glGenRenderbuffers(1, &b)
	return Renderbuffer(b)
}

func GenRenderbuffers(bufs []Renderbuffer) {
	if len(bufs) > 0 {
		C.glGenRenderbuffers(C.GLsizei(len(bufs)), (*C.GLuint)(&bufs[0]))
	}
}

func (rb Renderbuffer) Bind(target GLenum) {
	C.glBindRenderbuffer(C.GLenum(target), C.GLuint(rb))
}

func (rb Renderbuffer) Delete() {
	C.glDeleteRenderbuffers(1, (*C.GLuint)(&rb))
}

func (rb Renderbuffer) FramebufferRenderbuffer(target GLenum, attachment GLenum, renderbuffertarget GLenum) {
	C.glFramebufferRenderbuffer(C.GLenum(target), C.GLenum(attachment), C.GLenum(renderbuffertarget), C.GLuint(rb))
}

func DeleteRenderbuffers(bufs []Renderbuffer) {
	if len(bufs) > 0 {
		C.glDeleteRenderbuffers(C.GLsizei(len(bufs)), (*C.GLuint)(&bufs[0]))
	}
}

func GetFramebufferAttachmentParameteriv(target GLenum, attachment GLenum, pname GLenum, params []int32) {
	if len(params) == 0 {
		panic("Invalid params size")
	}
	C.glGetFramebufferAttachmentParameteriv(C.GLenum(target), C.GLenum(attachment), C.GLenum(pname), (*C.GLint)(&params[0]))
}

func GetRenderbufferParameteriv(target GLenum, pname GLenum, params []int32) {
	C.glGetRenderbufferParameteriv(C.GLenum(target), C.GLenum(pname), (*C.GLint)(&params[0]))
}

func RenderbufferStorage(target GLenum, internalformat GLenum, width int, height int) {
	C.glRenderbufferStorage(C.GLenum(target), C.GLenum(internalformat), C.GLsizei(width), C.GLsizei(height))
}

func RenderbufferStorageMultisample(target GLenum, samples int, internalformat GLenum, width int, height int) {
	C.glRenderbufferStorageMultisample(C.GLenum(target), C.GLsizei(samples), C.GLenum(internalformat), C.GLsizei(width), C.GLsizei(height))
}
