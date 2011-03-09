package gl

// #include <stdlib.h>
// #define GL_GLEXT_PROTOTYPES
//
// void* cmalloc(int s){return malloc(s);};
//
// #ifdef __APPLE__
// # include <OpenGL/gl.h>
// # include <OpenGL/glext.h>
// #else
// # include <GL/gl.h>
// # include <GL/glext.h>
// #endif
import "C"

// Renderbuffer Objects

type Renderbuffer Object

// void glGenRenderbuffers(GLsizei n, GLuint *renderbuffers)
func GenRenderbuffer() Renderbuffer {
	var b C.GLuint
	C.glGenRenderbuffers(1, &b)
	return Renderbuffer(b)
}

// Fill slice with new renderbuffers
func GenRenderbuffers(bufs []Renderbuffer) {
	C.glGenRenderbuffers(C.GLsizei(len(bufs)), (*C.GLuint)(&bufs[0]))
}

// void glBindRenderbuffer(GLenum target, GLuint renderbuffer);
func (rb Renderbuffer) Bind(target GLenum) {
	C.glBindRenderbuffer(C.GLenum(target), C.GLuint(rb))
}

// void glDeleteRenderbuffers(GLsizei n, GLuint* renderbuffers);
func (rb Renderbuffer) Delete() {
	C.glDeleteRenderbuffers(1, (*C.GLuint)(&rb))
}

func DeleteRenderbuffers(bufs []Renderbuffer) {
	C.glDeleteRenderbuffers(C.GLsizei(len(bufs)), (*C.GLuint)(&bufs[0]))
}

// void glGetRenderbufferParameteriv(GLenum target, GLenum pname, GLint* params);
//func GetRenderbufferParameteriv (target, pname GLenum, params []int) {
//  C.glGetRenderbufferParameteriv (C.GLenum(target), C.GLenum(pname), (*C.GLint)(&params[0]))
//}

// void glRenderbufferStorage(GLenum target, GLenum internalformat, GLsizei width, GLsizei height);
func RenderbufferStorage(target, internalformat GLenum, width int, height int) {
	C.glRenderbufferStorage(C.GLenum(target), C.GLenum(internalformat), C.GLsizei(width), C.GLsizei(height))
}

// void glRenderbufferStorageMultisample(GLenum target, GLsizei samples, GLenum internalformat, GLsizei width, GLsizei height);
func RenderbufferStorageMultisample(target GLenum, samples int, internalformat GLenum, width, height int) {
	C.glRenderbufferStorageMultisample(C.GLenum(target), C.GLsizei(samples), C.GLenum(internalformat), C.GLsizei(width), C.GLsizei(height))
}

// Framebuffer Objects
// TODO: implement GLsync stuff
type Framebuffer Object

// void glBindFramebuffer(GLenum target, GLuint framebuffer);
func (fb Framebuffer) Bind(target GLenum) {
	C.glBindFramebuffer(C.GLenum(target), C.GLuint(fb))
}

// void glBlitFramebuffer(GLint srcX0, GLint srcY0, GLint srcX1, GLint srcY1, GLint dstX0, GLint dstY0, GLint dstX1, GLint dstY1, GLbitfield mask, GLenum filter);
func BlitFramebuffer(srcX0, srcY0, srcX1, srcY1, dstX0, dstY0, dstX1, dstY1 int, mask GLbitfield, filter GLenum) {
	C.glBlitFramebuffer(C.GLint(srcX0), C.GLint(srcY0), C.GLint(srcX1), C.GLint(srcY1), C.GLint(dstX0), C.GLint(dstY0), C.GLint(dstX1), C.GLint(dstY1), C.GLbitfield(mask), C.GLenum(filter))
}

// GLenum glCheckFramebufferStatus(GLenum target);
func CheckFramebufferStatus(target GLenum) GLenum {
	return (GLenum)(C.glCheckFramebufferStatus(C.GLenum(target)))
}

// void glDeleteFramebuffers(GLsizei n, GLuint* framebuffers);
func (fb Framebuffer) DeleteFramebuffer() {
	C.glDeleteFramebuffers(1, (*C.GLuint)(&fb))
}

func DeleteFramebuffers(bufs []Framebuffer) {
	C.glDeleteFramebuffers(C.GLsizei(len(bufs)), (*C.GLuint)(&bufs[0]))
}

// GLsync glFramebufferRenderbuffer(GLenum target, GLenum attachment, GLenum renderbuffertarget, GLuint renderbuffer);
func (rb Renderbuffer) FramebufferRenderbuffer(target, attachment, renderbuffertarget GLenum) /* GLsync */ {
	// TODO: sync stuff.  return (GLsync)(C.glFramebufferRenderbuffer (C.GLenum(target), C.GLenum(attachment), C.GLenum(renderbuffertarget), C.GLuint(rb)))
	C.glFramebufferRenderbuffer(C.GLenum(target), C.GLenum(attachment), C.GLenum(renderbuffertarget), C.GLuint(rb))
}

// void glFramebufferTexture1D(GLenum target, GLenum attachment, GLenum textarget, GLuint texture, GLint level);
func FramebufferTexture1D(target, attachment, textarget GLenum, texture Texture, level int) {
	C.glFramebufferTexture1D(C.GLenum(target), C.GLenum(attachment), C.GLenum(textarget), C.GLuint(texture), C.GLint(level))
}

// void glFramebufferTexture2D(GLenum target, GLenum attachment, GLenum textarget, GLuint texture, GLint level);
func FramebufferTexture2D(target, attachment, textarget GLenum, texture Texture, level int) {
	C.glFramebufferTexture2D(C.GLenum(target), C.GLenum(attachment), C.GLenum(textarget), C.GLuint(texture), C.GLint(level))
}

// void glFramebufferTexture3D(GLenum target, GLenum attachment, GLenum textarget, GLuint texture, GLint level, GLint layer);
func FramebufferTexture3D(target, attachment, textarget GLenum, texture Texture, level int, layer int) {
	C.glFramebufferTexture3D(C.GLenum(target), C.GLenum(attachment), C.GLenum(textarget), C.GLuint(texture), C.GLint(level), C.GLint(layer))
}

// void glFramebufferTextureLayer(GLenum target, GLenum attachment, GLuint texture, GLint level, GLint layer);
func FramebufferTextureLayer(target, attachment GLenum, texture Texture, level, layer int) {
	C.glFramebufferTextureLayer(C.GLenum(target), C.GLenum(attachment), C.GLuint(texture), C.GLint(level), C.GLint(layer))
}

// void glGenFramebuffers(GLsizei n, GLuint* ids);
func GenFramebuffer() Framebuffer {
	var b C.GLuint
	C.glGenRenderbuffers(1, &b)
	return Framebuffer(b)
}

func GenFramebuffers(bufs []Framebuffer) {
	C.glGenFramebuffers(C.GLsizei(len(bufs)), (*C.GLuint)(&bufs[0]))
}

// void glGetFramebufferAttachmentParameter(GLenum target, GLenum attachment, GLenum pname, GLint* params);
//func GetFramebufferAttachmentParameter (target, attachment, pname GLenum, params []int) {
//  C.glGetFramebufferAttachmentParameter (C.GLenum(target), C.GLenum(attachment), C.GLenum(pname), (*C.GLint)(&params[0]))
//}
