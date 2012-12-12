// Copyright 2012 The go-gl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gl

// #cgo darwin LDFLAGS: -framework OpenGL -lGLEW
// #cgo windows LDFLAGS: -lglew32 -lopengl32
// #cgo linux LDFLAGS: -lGLEW -lGL
// #include "gl.h"
import "C"
import "unsafe"
import "reflect"

type GLenum C.GLenum
type GLbitfield C.GLbitfield
type GLclampf C.GLclampf
type GLclampd C.GLclampd

type Pointer unsafe.Pointer

// those types are left for compatibility reasons
type GLboolean C.GLboolean
type GLbyte C.GLbyte
type GLshort C.GLshort
type GLint C.GLint
type GLsizei C.GLsizei
type GLubyte C.GLubyte
type GLushort C.GLushort
type GLuint C.GLuint
type GLfloat C.GLfloat
type GLdouble C.GLdouble

// helpers

func glBool(v bool) C.GLboolean {
	if v {
		return 1
	}

	return 0
}

func goBool(v C.GLboolean) bool {
	return v != 0
}

func glString(s string) *C.GLchar { return (*C.GLchar)(C.CString(s)) }

func freeString(ptr *C.GLchar) { C.free(unsafe.Pointer(ptr)) }

func ptr(v interface{}) unsafe.Pointer {

	if v == nil {
		return unsafe.Pointer(nil)
	}

	rv := reflect.ValueOf(v)
	var et reflect.Value
	switch rv.Type().Kind() {
	case reflect.Uintptr:
		offset, _ := v.(uintptr)
		return unsafe.Pointer(offset)
	case reflect.Ptr:
		et = rv.Elem()
	case reflect.Slice:
		et = rv.Index(0)
	default:
		panic("type must be a pointer, a slice, uintptr or nil")
	}

	return unsafe.Pointer(et.UnsafeAddr())
}

// Main

func BlendColor(red GLclampf, green GLclampf, blue GLclampf, alpha GLclampf) {
	C.glBlendColor(C.GLclampf(red), C.GLclampf(green), C.GLclampf(blue), C.GLclampf(alpha))
}

func BlendEquation(mode GLenum) { C.glBlendEquation(C.GLenum(mode)) }

func BlendEquationSeparate(modeRGB GLenum, modeAlpha GLenum) {
	C.glBlendEquationSeparate(C.GLenum(modeRGB), C.GLenum(modeAlpha))
}

func BlendFuncSeparate(srcRGB GLenum, dstRGB GLenum, srcAlpha GLenum, dstAlpha GLenum) {
	C.glBlendFuncSeparate(C.GLenum(srcRGB), C.GLenum(dstRGB), C.GLenum(srcAlpha), C.GLenum(dstAlpha))
}

func SampleCoverage(value GLclampf, invert bool) {
	C.glSampleCoverage(C.GLclampf(value), glBool(invert))
}

func StencilFuncSeparate(face GLenum, func_ GLenum, ref int, mask uint) {
	C.glStencilFuncSeparate(C.GLenum(face), C.GLenum(func_), C.GLint(ref), C.GLuint(mask))
}

func StencilMaskSeparate(face GLenum, mask uint) {
	C.glStencilMaskSeparate(C.GLenum(face), C.GLuint(mask))
}

func StencilOpSeparate(face GLenum, fail GLenum, zfail GLenum, zpass GLenum) {
	C.glStencilOpSeparate(C.GLenum(face), C.GLenum(fail), C.GLenum(zfail), C.GLenum(zpass))
}

//void glBlendFunc (GLenum sfactor, GLenum dfactor)
func BlendFunc(sfactor GLenum, dfactor GLenum) {
	C.glBlendFunc(C.GLenum(sfactor), C.GLenum(dfactor))
}

//void glClear (GLbitfield mask)
func Clear(mask GLbitfield) {
	C.glClear(C.GLbitfield(mask))
}

//void glClearColor (GLclampf red, GLclampf green, GLclampf blue, GLclampf alpha)
func ClearColor(red GLclampf, green GLclampf, blue GLclampf, alpha GLclampf) {
	C.glClearColor(C.GLclampf(red), C.GLclampf(green), C.GLclampf(blue), C.GLclampf(alpha))
}

//void glClearDepth (GLclampd depth)
func ClearDepth(depth GLclampd) {
	C.glClearDepth(C.GLclampd(depth))
}

//void glClearStencil (int s)
func ClearStencil(s int) {
	C.glClearStencil(C.GLint(s))
}

//void glCullFace (GLenum mode)
func CullFace(mode GLenum) {
	C.glCullFace(C.GLenum(mode))
}

//void glDepthFunc (GLenum func)
func DepthFunc(func_ GLenum) {
	C.glDepthFunc(C.GLenum(func_))
}

//void glDepthMask (bool flag)
func DepthMask(flag bool) {
	C.glDepthMask(glBool(flag))
}

//void glDepthRange (GLclampd zNear, GLclampd zFar)
func DepthRange(zNear GLclampd, zFar GLclampd) {
	C.glDepthRange(C.GLclampd(zNear), C.GLclampd(zFar))
}

//void glDisable (GLenum cap)
func Disable(cap GLenum) {
	C.glDisable(C.GLenum(cap))
}

//void glDrawArrays (GLenum mode, int first, int count)
func DrawArrays(mode GLenum, first int, count int) {
	C.glDrawArrays(C.GLenum(mode), C.GLint(first), C.GLsizei(count))
}

//void glDrawBuffer (GLenum mode)
func DrawBuffer(mode GLenum) {
	C.glDrawBuffer(C.GLenum(mode))
}

//void glDrawElements (GLenum mode, int count, GLenum type, const GLvoid *indices)
func DrawElements(mode GLenum, count int, typ GLenum, indices interface{}) {
	C.glDrawElements(C.GLenum(mode), C.GLsizei(count), C.GLenum(typ),
		ptr(indices))
}

//void glEnable (GLenum cap)
func Enable(cap GLenum) {
	C.glEnable(C.GLenum(cap))
}

//void glFinish (void)
func Finish() {
	C.glFinish()
}

//void glFlush (void)
func Flush() {
	C.glFlush()
}

//void glFrontFace (GLenum mode)
func FrontFace(mode GLenum) {
	C.glFrontFace(C.GLenum(mode))
}

//void glGetBooleanv (GLenum pname, bool *params)
func GetBooleanv(pname GLenum, params []bool) {
	if len(params) == 0 {
		panic("Invalid params length")
	}
	C.glGetBooleanv(C.GLenum(pname), (*C.GLboolean)(unsafe.Pointer(&params[0])))
}

//void glGetDoublev (GLenum pname, float64 *params)
func GetDoublev(pname GLenum, params []float64) {
	if len(params) == 0 {
		panic("Invalid params length")
	}
	C.glGetDoublev(C.GLenum(pname), (*C.GLdouble)(&params[0]))
}

//GLenum glGetError (void)
func GetError() GLenum {
	return GLenum(C.glGetError())
}

//void glGetFloatv (GLenum pname, float *params)
func GetFloatv(pname GLenum, params []float32) {
	if len(params) == 0 {
		panic("Invalid params length")
	}
	C.glGetFloatv(C.GLenum(pname), (*C.GLfloat)(&params[0]))
}

//void glGetIntegerv (GLenum pname, int *params)
func GetIntegerv(pname GLenum, params []int32) {
	if len(params) == 0 {
		panic("Invalid params length")
	}
	C.glGetIntegerv(C.GLenum(pname), (*C.GLint)(&params[0]))
}

//const uint8 * glGetString (GLenum name)
func GetString(name GLenum) string {
	s := unsafe.Pointer(C.glGetString(C.GLenum(name)))
	return C.GoString((*C.char)(s))
}

//void glHint (GLenum target, GLenum mode)
func Hint(target GLenum, mode GLenum) {
	C.glHint(C.GLenum(target), C.GLenum(mode))
}

//bool glIsEnabled (GLenum cap)
func IsEnabled(cap GLenum) bool {
	return goBool(C.glIsEnabled(C.GLenum(cap)))
}

//void glLineWidth (float32 width)
func LineWidth(width float32) {
	C.glLineWidth(C.GLfloat(width))
}

//void glLogicOp (GLenum opcode)
func LogicOp(opcode GLenum) {
	C.glLogicOp(C.GLenum(opcode))
}

//void glPixelStoref (GLenum pname, float param)
func PixelStoref(pname GLenum, param float32) {
	C.glPixelStoref(C.GLenum(pname), C.GLfloat(param))
}

//void glPixelStorei (GLenum pname, int param)
func PixelStorei(pname GLenum, param int) {
	C.glPixelStorei(C.GLenum(pname), C.GLint(param))
}

//void glPointSize (float32 size)
func PointSize(size float32) {
	C.glPointSize(C.GLfloat(size))
}

//void glPolygonMode (GLenum face, GLenum mode)
func PolygonMode(face GLenum, mode GLenum) {
	C.glPolygonMode(C.GLenum(face), C.GLenum(mode))
}

//void glPolygonOffset (float32 factor, float32 units)
func PolygonOffset(factor float32, units float32) {
	C.glPolygonOffset(C.GLfloat(factor), C.GLfloat(units))
}

//void glReadBuffer (GLenum mode)
func ReadBuffer(mode GLenum) {
	C.glReadBuffer(C.GLenum(mode))
}

//void glReadPixels (int x, int y, int width, int height, GLenum format, GLenum type, GLvoid *pixels)
func ReadPixels(x int, y int, width int, height int, format, typ GLenum, pixels interface{}) {
	C.glReadPixels(C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height),
		C.GLenum(format), C.GLenum(typ), ptr(pixels))
}

//void glScissor (int x, int y, int width, int height)
func Scissor(x int, y int, width int, height int) {
	C.glScissor(C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height))
}

//void glStencilFunc (GLenum func, int ref, uint mask)
func StencilFunc(func_ GLenum, ref int, mask uint) {
	C.glStencilFunc(C.GLenum(func_), C.GLint(ref), C.GLuint(mask))
}

//void glStencilMask (uint mask)
func StencilMask(mask uint) {
	C.glStencilMask(C.GLuint(mask))
}

//void glStencilOp (GLenum fail, GLenum zfail, GLenum zpass)
func StencilOp(fail GLenum, zfail GLenum, zpass GLenum) {
	C.glStencilOp(C.GLenum(fail), C.GLenum(zfail), C.GLenum(zpass))
}

//void glViewport (int x, int y, int width, int height)
func Viewport(x int, y int, width int, height int) {
	C.glViewport(C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height))
}

func Init() GLenum {
	return GLenum(C.goglInit())
}
