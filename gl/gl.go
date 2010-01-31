package gl

// #include <stdlib.h>
// #define GL_GLEXT_PROTOTYPES
//
// int glewInit(void);
//
// void* cmalloc(int s){return malloc(s);};
//
// #include <GL/gl.h>
// #include <GL/glext.h>
import "C"
import "unsafe"

type GLenum C.GLenum
type GLbitfield C.GLbitfield
type GLint C.GLint
type GLsizei C.GLsizei
type GLsizeiptr C.GLsizeiptr
type GLuint C.GLuint
type GLfloat C.GLfloat
type GLclampf C.GLclampf

type Pointer unsafe.Pointer

// helpers

func glBool(v bool) C.GLboolean {
	if v {
		return 1
	}

	return 0;

}

func glString(s string) *C.GLchar	{ return (*C.GLchar)(C.CString(s)) }

func freeString(ptr *C.GLchar)	{ C.free(unsafe.Pointer(ptr)) }

// GLEW

func Init() int{
    return int(C.glewInit());
}

// Object

type Object C.GLuint

func (buffer Object) IsBuffer() bool	{ return C.glIsBuffer(C.GLuint(buffer)) != 0 }


func (framebuffer Object) IsFramebuffer() bool {
	return C.glIsFramebuffer(C.GLuint(framebuffer)) != 0
}

func (program Object) IsProgram() bool	{ return C.glIsProgram(C.GLuint(program)) != 0 }

func (renderbuffer Object) IsRenderbuffer() bool {
	return C.glIsRenderbuffer(C.GLuint(renderbuffer)) != 0
}

func (shader Object) IsShader() bool	{ return C.glIsShader(C.GLuint(shader)) != 0 }

func (texture Object) IsTexture() bool	{ return C.glIsTexture(C.GLuint(texture)) != 0 }

// Shader

type Shader Object

func CreateShader(type_ GLenum) Shader	{ return Shader(C.glCreateShader(C.GLenum(type_))) }

func DeleteShader(shader Shader)	{ C.glDeleteShader(C.GLuint(shader)) }


func (shader Shader) GetInfoLog() string {
	var len C.GLint;
	C.glGetShaderiv(C.GLuint(shader), C.GLenum(INFO_LOG_LENGTH), &len);

	log := C.cmalloc(C.int(len + 1));
	C.glGetShaderInfoLog(C.GLuint(shader), C.GLsizei(len), nil, (*C.GLchar)(log));



	defer C.free(log);

	return C.GoString((*C.char)(log));
}

func (shader Shader) GetSource() string {
	var len C.GLint;
	C.glGetShaderiv(C.GLuint(shader), C.GLenum(SHADER_SOURCE_LENGTH), &len);

	log := C.cmalloc(C.int(len + 1));
	C.glGetShaderSource(C.GLuint(shader), C.GLsizei(len), nil, (*C.GLchar)(log));

	defer C.free(log);

	return C.GoString((*C.char)(log));
}

func (shader Shader) Source(source string) {

	csource := glString(source);
	defer freeString(csource);

	var one C.GLint = C.GLint(len(source));

	C.glShaderSource(C.GLuint(shader), 1, &csource, &one);
}


func (shader Shader) Compile()	{ C.glCompileShader(C.GLuint(shader)) }

// Program

type Program Object

func CreateProgram() Program	{ return Program(C.glCreateProgram()) }

func (program Program) Delete()	{ C.glDeleteProgram(C.GLuint(program)) }

func (program Program) AttachShader(shader Shader) {
	C.glAttachShader(C.GLuint(program), C.GLuint(shader))
}


func (program Program) GetAttachedShaders() []Object {
	var len C.GLint;
	C.glGetProgramiv(C.GLuint(program), C.GLenum(ACTIVE_UNIFORM_MAX_LENGTH), &len);

	objects := make([]Object, len);
	C.glGetAttachedShaders(C.GLuint(program), C.GLsizei(len), nil, *((**C.GLuint)(unsafe.Pointer(&objects))));
	return objects;
}

func (program Program) DetachShader(shader Shader) {
	C.glDetachShader(C.GLuint(program), C.GLuint(shader))
}


func (program Program) Link()	{ C.glLinkProgram(C.GLuint(program)) }

func (program Program) Validate()	{ C.glValidateProgram(C.GLuint(program)) }

func (program Program) Use()	{ C.glUseProgram(C.GLuint(program)) }


func (program Program) GetInfoLog() string {

	var len C.GLint;
	C.glGetProgramiv(C.GLuint(program), C.GLenum(INFO_LOG_LENGTH), &len);

	log := C.cmalloc(C.int(len + 1));
	C.glGetProgramInfoLog(C.GLuint(program), C.GLsizei(len), nil, (*C.GLchar)(log));

	defer C.free(log);

	return C.GoString((*C.char)(log));

}

func (program Program) GetUniform(location UniformLocation) int {
	//TODO
	return 0
	//return int(C.glGetUniform(C.GLuint(program), C.GLuint(location)));
}

func (program Program) GetUniformLocation(name string) UniformLocation {

	cname := glString(name);
	defer freeString(cname);

	return UniformLocation(C.glGetUniformLocation(C.GLuint(program), cname));
}

func (program Program) GetAttribLocation(name string) VertexAttrib {

	cname := glString(name);
	defer freeString(cname);

	return VertexAttrib(C.glGetAttribLocation(C.GLuint(program), cname));
}


func (program Program) BindAttribLocation(index GLuint, name string) {

	cname := glString(name);
	defer freeString(cname);

	C.glBindAttribLocation(C.GLuint(program), C.GLuint(index), cname);

}

// Buffer

type Buffer Object

func CreateBuffer() Buffer {
	var b C.GLuint;
	C.glGenBuffers(1, &b);
	return Buffer(b);
}

func (buffer Buffer) Bind(target GLenum)	{ C.glBindBuffer(C.GLenum(target), C.GLuint(buffer)) }

func (buffer Buffer) Delete() {
	b := C.GLuint(buffer);
	C.glDeleteBuffers(1, &b);
}

// FrameBuffer

type Framebuffer Object

func CreateFramebuffer() Framebuffer {
	var b C.GLuint;
	C.glGenFramebuffers(1, &b);
	return Framebuffer(b);
}

func (framebuffer Framebuffer) Bind(target GLenum) {
	C.glBindFramebuffer(C.GLenum(target), C.GLuint(framebuffer))
}

func (framebuffer Framebuffer) Delete() {
	b := C.GLuint(framebuffer);
	C.glDeleteFramebuffers(1, &b);
}


// RenderBuffer

type Renderbuffer Object

func CreateRenderbuffer() Renderbuffer {
	var b C.GLuint;
	C.glGenRenderbuffers(1, &b);
	return Renderbuffer(b);
}

func (renderbuffer Renderbuffer) Bind(target GLenum) {
	C.glBindRenderbuffer(C.GLenum(target), C.GLuint(renderbuffer))
}

func (renderbuffer Renderbuffer) Delete() {
	b := C.GLuint(renderbuffer);
	C.glDeleteRenderbuffers(1, &b);
}

// Texture

type Texture Object

func CreateTexture() Texture {
	var b C.GLuint;
	C.glGenTextures(1, &b);
	return Texture(b);
}

func (texture Texture) Bind(target GLenum) {
	C.glBindTexture(C.GLenum(target), C.GLuint(texture))
}

func (texture Texture) Delete() {
	b := C.GLuint(texture);
	C.glDeleteTextures(1, &b);
}

// VertexAttrib

type VertexAttrib GLuint;

func VertexAttrib1f(indx GLuint, x GLfloat) {
	C.glVertexAttrib1f(C.GLuint(indx), C.GLfloat(x))
}

func VertexAttrib1fv(indx GLuint, values *float) {
	//TODO
	//	C.glVertexAttrib1fv(C.GLuint(indx), (*C.float)(values));
}

func VertexAttrib2f(indx GLuint, x GLfloat, y GLfloat) {
	C.glVertexAttrib2f(C.GLuint(indx), C.GLfloat(x), C.GLfloat(y))
}

func VertexAttrib2fv(indx GLuint, values *float) {
	//TODO
	//	C.glVertexAttrib2fv(C.GLuint(indx), (*C.float)(values));
}

func VertexAttrib3f(indx GLuint, x GLfloat, y GLfloat, z GLfloat) {
	C.glVertexAttrib3f(C.GLuint(indx), C.GLfloat(x), C.GLfloat(y), C.GLfloat(z))
}

func VertexAttrib3fv(indx GLuint, values *float) {
	//TODO
	//	C.glVertexAttrib3fv(C.GLuint(indx), (*C.float)(values));
}

func VertexAttrib4f(indx GLuint, x GLfloat, y GLfloat, z GLfloat, w GLfloat) {
	C.glVertexAttrib4f(C.GLuint(indx), C.GLfloat(x), C.GLfloat(y), C.GLfloat(z), C.GLfloat(w))
}

func VertexAttrib4fv(indx GLuint, values *float) {
	//TODO
	//	C.glVertexAttrib4fv(C.GLuint(indx), (*C.float)(values));
}

func VertexAttribPointer(indx VertexAttrib, size GLuint, type_ GLenum, normalized bool,  stride GLsizei,  pointer unsafe.Pointer)  {
    C.glVertexAttribPointer(C.GLuint(indx), C.GLint(size), C.GLenum(type_), glBool(normalized), C.GLsizei(stride), pointer);
}


// UniformLocation
//TODO


type UniformLocation GLint

func (location UniformLocation) setFloat(x GLfloat) {
	C.glUniform1f(C.GLint(location), C.GLfloat(x))
}

func (location UniformLocation) setFloatv(v []float) {
	//	C.glUniform1fv(C.GLint(location), (*C.float)(v));
}

func (location UniformLocation) setInt(x GLint) {
	C.glUniform1i(C.GLint(location), C.GLint(x))
}

func (location UniformLocation) setIntv(v []int) {
	//	C.glUniform1iv(C.GLint(location), (*C.int)(v));
}

func (location UniformLocation) Set2Float(x GLfloat, y GLfloat) {
	C.glUniform2f(C.GLint(location), C.GLfloat(x), C.GLfloat(y))
}

func Uniform2fv(location UniformLocation, v *float) {
	//	C.glUniform2fv(C.GLint(location), (*C.float)(v));
}

func Uniform2i(location UniformLocation, x GLint, y GLint) {
	C.glUniform2i(C.GLint(location), C.GLint(x), C.GLint(y))
}

func Uniform2iv(location UniformLocation, v *int) {
	//	C.glUniform2iv(C.GLint(location), (*C.int)(v));
}

func Uniform3f(location UniformLocation, x GLfloat, y GLfloat, z GLfloat) {
	C.glUniform3f(C.GLint(location), C.GLfloat(x), C.GLfloat(y), C.GLfloat(z))
}

func Uniform3fv(location UniformLocation, v *float) {
	//	C.glUniform3fv(C.GLint(location), (*C.float)(v));
}

func Uniform3i(location UniformLocation, x GLint, y GLint, z GLint) {
	C.glUniform3i(C.GLint(location), C.GLint(x), C.GLint(y), C.GLint(z))
}

func Uniform3iv(location UniformLocation, v *int) {
	//	C.glUniform3iv(C.GLint(location), (*C.int)(v));
}

func Uniform4f(location UniformLocation, x GLfloat, y GLfloat, z GLfloat, w GLfloat) {
	C.glUniform4f(C.GLint(location), C.GLfloat(x), C.GLfloat(y), C.GLfloat(z), C.GLfloat(w))
}

func Uniform4fv(location UniformLocation, v *float) {
	//	C.glUniform4fv(C.GLint(location), (*C.float)(v));
}

func Uniform4i(location UniformLocation, x GLint, y GLint, z GLint, w GLint) {
	C.glUniform4i(C.GLint(location), C.GLint(x), C.GLint(y), C.GLint(z), C.GLint(w))
}

func Uniform4iv(location UniformLocation, v *int) {
	//	C.glUniform4iv(C.GLint(location), (*C.int)(v));
}

/*
uniformMatrix2fv
uniformMatrix2fv
uniformMatrix3fv
uniformMatrix3fv
uniformMatrix4fv
uniformMatrix4fv
*/

// Main

func ActiveTexture(texture GLenum)	{ C.glActiveTexture(C.GLenum(texture)) }

func BlendColor(red GLclampf, green GLclampf, blue GLclampf, alpha GLclampf) {
	C.glBlendColor(C.GLclampf(red), C.GLclampf(green), C.GLclampf(blue), C.GLclampf(alpha))
}

func BlendEquation(mode GLenum)	{ C.glBlendEquation(C.GLenum(mode)) }

func BlendEquationSeparate(modeRGB GLenum, modeAlpha GLenum) {
	C.glBlendEquationSeparate(C.GLenum(modeRGB), C.GLenum(modeAlpha))
}

func BlendFunc(sfactor GLenum, dfactor GLenum) {
	C.glBlendFunc(C.GLenum(sfactor), C.GLenum(dfactor))
}

func blendFuncSeparate(srcRGB GLenum, dstRGB GLenum, srcAlpha GLenum, dstAlpha GLenum) {
    C.glBlendFuncSeparate(C.GLenum(srcRGB), C.GLenum(dstRGB), C.GLenum(srcAlpha), C.GLenum(dstAlpha))
}

func BufferData(target GLenum, size int, data unsafe.Pointer, usage GLenum) {
	C.glBufferData(C.GLenum(target), C.GLsizeiptr(size), data, C.GLenum(usage));
}

func BufferSubData(target GLenum, offset GLsizeiptr, size int, data unsafe.Pointer) {
	C.glBufferSubData(C.GLenum(target), C.GLintptr(offset), C.GLsizeiptr(size), data);
}

func CheckFramebufferStatus(target GLenum) GLenum {
	return GLenum(C.glCheckFramebufferStatus(C.GLenum(target)))
}

func Clear(mask GLbitfield)	{ C.glClear(C.GLbitfield(mask)) }

func ClearColor(red GLclampf, green GLclampf, blue GLclampf, alpha GLclampf) {
	C.glClearColor(C.GLclampf(red), C.GLclampf(green), C.GLclampf(blue), C.GLclampf(alpha))
}

func ClearDepth(depth GLclampf)	{ C.glClearDepth(C.GLclampd(depth)) }

func ClearStencil(s GLint)	{ C.glClearStencil(C.GLint(s)) }

func ColorMask(red bool, green bool, blue bool, alpha bool) {
	C.glColorMask(glBool(red), glBool(green), glBool(blue), glBool(alpha))
}

func CullFace(mode GLenum)	{ C.glCullFace(C.GLenum(mode)) }

func DepthFunc(func_ GLenum)	{ C.glDepthFunc(C.GLenum(func_)) }

func DepthMask(flag bool)	{ C.glDepthMask(glBool(flag)) }

func DepthRange(zNear GLclampf, zFar GLclampf) {
	C.glDepthRange(C.GLclampd(zNear), C.GLclampd(zFar))
}

func Disable(cap GLenum)	{ C.glDisable(C.GLenum(cap)) }

func DisableVertexAttribArray(index VertexAttrib)	{ C.glDisableVertexAttribArray(C.GLuint(index)) }

func DrawArrays(mode GLenum, first GLint, count int) {
	C.glDrawArrays(C.GLenum(mode), C.GLint(first), C.GLsizei(count))
}

func DrawElements(mode GLenum, count GLsizei, type_ GLenum, offset uintptr) {
	C.glDrawElements(C.GLenum(mode), C.GLsizei(count), C.GLenum(type_), unsafe.Pointer(offset))
}

func Enable(cap GLenum)	{ C.glEnable(C.GLenum(cap)) }

func EnableVertexAttribArray(index VertexAttrib)	{ C.glEnableVertexAttribArray(C.GLuint(index)) }

func Finish()	{ C.glFinish() }

func Flush()	{ C.glFlush() }

func FrontFace(mode GLenum)	{ C.glFrontFace(C.GLenum(mode)) }

func GenerateMipmap(target GLenum)	{ C.glGenerateMipmap(C.GLenum(target)) }

func GetError() GLenum	{ return GLenum(C.glGetError()) }

func GetString(name GLenum) string {
	s := (*C.char)(unsafe.Pointer(C.glGetString(C.GLenum(name))));
	return C.GoString(s);
}

func Hint(target GLenum, mode GLenum)	{ C.glHint(C.GLenum(target), C.GLenum(mode)) }

func IsEnabled(cap GLenum) bool	{ return C.glIsEnabled(C.GLenum(cap)) != 0 }

func LineWidth(width GLfloat)	{ C.glLineWidth(C.GLfloat(width)) }


func PixelStorei(pname GLenum, param GLint) {
	C.glPixelStorei(C.GLenum(pname), C.GLint(param))
}

func PolygonOffset(factor GLfloat, units GLfloat) {
	C.glPolygonOffset(C.GLfloat(factor), C.GLfloat(units))
}

func SampleCoverage(value GLclampf, invert bool) {
	C.glSampleCoverage(C.GLclampf(value), glBool(invert))
}

func Scissor(x GLint, y GLint, width GLsizei, height GLsizei) {
	C.glScissor(C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height))
}

func StencilFunc(func_ GLenum, ref GLint, mask GLuint) {
	C.glStencilFunc(C.GLenum(func_), C.GLint(ref), C.GLuint(mask))
}

func StencilFuncSeparate(face GLenum, func_ GLenum, ref GLint, mask GLuint) {
	C.glStencilFuncSeparate(C.GLenum(face), C.GLenum(func_), C.GLint(ref), C.GLuint(mask))
}

func StencilMask(mask GLuint)	{ C.glStencilMask(C.GLuint(mask)) }

func StencilMaskSeparate(face GLenum, mask GLuint) {
	C.glStencilMaskSeparate(C.GLenum(face), C.GLuint(mask))
}

func StencilOp(fail GLenum, zfail GLenum, zpass GLenum) {
	C.glStencilOp(C.GLenum(fail), C.GLenum(zfail), C.GLenum(zpass))
}

func StencilOpSeparate(face GLenum, fail GLenum, zfail GLenum, zpass GLenum) {
	C.glStencilOpSeparate(C.GLenum(face), C.GLenum(fail), C.GLenum(zfail), C.GLenum(zpass))
}

func TexParameterf(target GLenum, pname GLenum, param GLfloat) {
	C.glTexParameterf(C.GLenum(target), C.GLenum(pname), C.GLfloat(param))
}

func TexParameteri(target GLenum, pname GLenum, param GLint) {
	C.glTexParameteri(C.GLenum(target), C.GLenum(pname), C.GLint(param))
}

func Viewport(x GLint, y GLint, width GLsizei, height GLsizei) {
	C.glViewport(C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height))
}

func CopyTexImage2D(target GLenum, level int, internalformat GLenum, x int, y int, width int, height int, border int) {
	C.glCopyTexImage2D(C.GLenum(target), C.GLint(level), C.GLenum(internalformat), C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height), C.GLint(border));
}

//void copyTexSubImage2D(GLenum target, GLint level, GLint xoffset, GLint yoffset, GLint x, GLint y, GLsizei width, GLsizei height)
func CopyTexSubImage2D(target GLenum, level int, xoffset int, yoffset int, x int, y int, width int, height int) {
	C.glCopyTexSubImage2D(C.GLenum(target), C.GLint(level), C.GLint(xoffset), C.GLint(yoffset), C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height));
}

//void renderbufferStorage(GLenum target, GLenum internalformat, GLsizei width, GLsizei height)
func RenderbufferStorage(target GLenum, internalformat GLenum, width GLsizei, height GLsizei) {
	C.glRenderbufferStorage(C.GLenum(target), C.GLenum(internalformat), C.GLsizei(width), C.GLsizei(height));
}


//void texImage2D(GLenum target, GLint level, GLenum internalformat, GLsizei width, GLsizei height, GLint border, GLenum format, GLenum type, WebGLArray pixels)
//func TexImage2D(target GLenum, level GLint, internalformat GLenum, width GLsizei, height GLsizei, border GLint, format GLenum, type_ GLenum, pixels WebGLArray) {
//	C.glTexImage2D(C.GLenum(target), C.GLint(level), C.GLenum(internalformat), C.GLsizei(width), C.GLsizei(height), C.GLint(border), C.GLenum(format), C.GLenum(type_), C.WebGLArray(pixels));
//}

//void texSubImage2D(GLenum target, GLint level, GLint xoffset, GLint yoffset, GLsizei width, GLsizei height, GLenum format, GLenum type, WebGLArray pixels)
//func TexSubImage2D(target GLenum, level GLint, xoffset GLint, yoffset GLint, width GLsizei, height GLsizei, format GLenum, type_ GLenum, pixels WebGLArray) {
//	C.glTexSubImage2D(C.GLenum(target), C.GLint(level), C.GLint(xoffset), C.GLint(yoffset), C.GLsizei(width), C.GLsizei(height), C.GLenum(format), C.GLenum(type_), C.WebGLArray(pixels));
//}

//WebGLArray readPixels(GLint x, GLint y, GLsizei width, GLsizei height, GLenum format, GLenum type)
//func ReadPixels(x GLint, y GLint, width GLsizei, height GLsizei, format GLenum, type_ GLenum) WebGLArray
//{
//	return WebGLArray(C.glRreadPixels(C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height), C.GLenum(format), C.GLenum(type_)));
//}

/*
func GetActiveUniform(program GLuint, index GLuint) WebGLActiveInfo
{
	return WebGLActiveInfo(C.glGetActiveUniform(C.GLuint(program), C.GLuint(index)));
}

func GetActiveAttrib(program GLuint, index GLuint) WebGLActiveInfo
{
	return WebGLActiveInfo(C.glGetActiveAttrib(C.GLuint(program), C.GLuint(index)));
}
*/

func FramebufferTexture2D(target GLenum, attachment GLenum, textarget GLenum, texture Texture, level GLint) {
	C.glFramebufferTexture2D(C.GLenum(target), C.GLenum(attachment), C.GLenum(textarget), C.GLuint(texture), C.GLint(level));
}

func FramebufferRenderbuffer(target Framebuffer, attachment GLenum, renderbuffertarget GLenum, renderbuffer Renderbuffer) {
	C.glFramebufferRenderbuffer(C.GLenum(target), C.GLenum(attachment), C.GLenum(renderbuffertarget), C.GLuint(renderbuffer));
}


/* TODO
webgl:
    getParameter
    getBufferParameter
    getFramebufferAttachmentParameter
    getProgramParameter
    getRenderbufferParameter
    getShaderParameter
    getTexParameter

    sizeInBytes
    getContextAttributes


gles:
    glClearDepthf
    glCompressedTexImage2D
    glCompressedTexSubImage2D
    glDepthRangef
    glGetBufferParameteriv
    glGetFramebufferAttachmentParameteriv
    glGetRenderbufferParameteriv
    glGetShaderPrecisionFormat
    glGetTexParameter
    glGetVertexAttrib
    glGetVertexAttribPointerv
    glReadPixels
    glReleaseShaderCompiler
    glShaderBinary


*/

// old

func Vertex3f(x GLfloat, y GLfloat, z GLfloat) {
	C.glVertex3f(C.GLfloat(x), C.GLfloat(y), C.GLfloat(z))
}
func Begin(mode GLenum) { C.glBegin(C.GLenum(mode)) }
func End() { C.glEnd() }

func VertexPointer(size GLint, type_ GLenum, stride GLsizei, pointer unsafe.Pointer) {
	C.glVertexPointer(C.GLint(size), C.GLenum(type_), C.GLsizei(stride), pointer)
}

func EnableClientState(array GLenum) { C.glEnableClientState(C.GLenum(array)) }
func DisableClientState(array GLenum) { C.glDisableClientState(C.GLenum(array)) }



