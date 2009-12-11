package gl

// #include <stdlib.h>
// #define GL_GLEXT_PROTOTYPES
// #include <GL/gl.h>
// #include <GL/glext.h>
import "C";
import "unsafe";

type GLenum C.GLenum;
type GLbitfield C.GLbitfield;
type GLint C.GLint;
type GLsizei C.GLsizei;
type GLsizeiptr C.GLsizeiptr;
type GLuint C.GLuint;
type GLfloat C.GLfloat;
type GLclampf C.GLclampf;

type Pointer unsafe.Pointer;

type Object C.GLuint;

// helpers

func glBool(v bool) C.GLboolean
{
    if v
    {
        return 1;
    }

    return 0;

}

func glString(s string) *C.GLchar
{
    return (*C.GLchar)(C.CString(s));
}

func freeString(ptr *C.GLchar)
{
    C.free(unsafe.Pointer(ptr));
}


// Shader

type Shader Object;

func (shader Shader)  Compile()
{
	C.glCompileShader(C.GLuint(shader));
}

// Program

type Program Object;

func (program Program) AttachShader(shader Shader) 
{
	C.glAttachShader(C.GLuint(program), C.GLuint(shader));
}

func (program Program) BindAttribLocation(index GLuint, name string) 
{

    cname := glString(name);
    defer freeString(cname);

	C.glBindAttribLocation(C.GLuint(program), C.GLuint(index), cname);

}

// Buffer

type Buffer Object;

func (buffer Buffer) Bind(target GLenum) 
{
    C.glBindBuffer(C.GLenum(target), C.GLuint(buffer));
}

// FrameBuffer

type Framebuffer Object;

func (framebuffer Framebuffer) Bind(target GLenum) 
{
	C.glBindFramebuffer(C.GLenum(target), C.GLuint(framebuffer));
}

// RenderBuffer

type Renderbuffer Object;

func (renderbuffer Renderbuffer) Bind(target GLenum) 
{
	C.glBindRenderbuffer(C.GLenum(target), C.GLuint(renderbuffer));
}

// Texture

type Texture Object;

func (texture Texture) Bind(target GLenum) 
{
	C.glBindTexture(C.GLenum(target), C.GLuint(texture));
}

// Main

func ActiveTexture(texture GLenum) 
{
	C.glActiveTexture(C.GLenum(texture));
}

func BlendColor(red GLclampf, green GLclampf, blue GLclampf, alpha GLclampf) 
{
//	C.glBlendColor(C.GLclampf(red), C.GLclampf(green), C.GLclampf(blue), C.GLclampf(alpha));
}

func BlendEquation(mode GLenum) 
{
	C.glBlendEquation(C.GLenum(mode));
}

func BlendEquationSeparate(modeRGB GLenum, modeAlpha GLenum) 
{
	C.glBlendEquationSeparate(C.GLenum(modeRGB), C.GLenum(modeAlpha));
}

func BlendFunc(sfactor GLenum, dfactor GLenum) 
{
	C.glBlendFunc(C.GLenum(sfactor), C.GLenum(dfactor));
}

func BufferData(target GLenum, data Pointer, usage GLenum) 
{
//	C.glBufferData(C.GLenum(target), C.Pointer(data), C.GLenum(usage));
}

func BufferSubData(target GLenum, offset GLsizeiptr, data Pointer) 
{
//	C.glBufferSubData(C.GLenum(target), C.GLsizeiptr(offset), C.Pointer(data));
}

func CheckFramebufferStatus(target GLenum) GLenum
{
    return 0;
	//return GLenum(C.glCheckFramebufferStatus(C.GLenum(target)));
}

func Clear(mask GLbitfield) 
{
	C.glClear(C.GLbitfield(mask));
}

func ClearColor(red GLclampf, green GLclampf, blue GLclampf, alpha GLclampf) 
{
	C.glClearColor(C.GLclampf(red), C.GLclampf(green), C.GLclampf(blue), C.GLclampf(alpha));
}

func ClearDepth(depth GLclampf) 
{
	C.glClearDepth(C.GLclampd(depth));
}

func ClearStencil(s GLint) 
{
	C.glClearStencil(C.GLint(s));
}

func ColorMask(red bool, green bool, blue bool, alpha bool) 
{
	C.glColorMask(glBool(red), glBool(green), glBool(blue), glBool(alpha));
}


func CreateBuffer() Buffer
{
    var b C.GLuint;
    C.glGenBuffers(1,&b);
	return Buffer(b);
}

func CreateFramebuffer() Framebuffer
{
    var b C.GLuint;
    C.glGenFramebuffers(1,&b);
	return Framebuffer(b);
}

func CreateProgram() Program
{
	return Program(C.glCreateProgram());
}

func CreateRenderbuffer() Renderbuffer
{
    var b C.GLuint;
    C.glGenRenderbuffers(1,&b);
	return Renderbuffer(b);
}

func CreateShader(type_ GLenum) Shader
{
	return Shader(C.glCreateShader(C.GLenum(type_)));
}

func CreateTexture() Texture
{
    var b C.GLuint;
    C.glGenTextures(1,&b);
	return Texture(b);
}

func CullFace(mode GLenum) 
{
	C.glCullFace(C.GLenum(mode));
}

func (buffer Buffer) Delete()
{
    b := C.GLuint(buffer);
	C.glDeleteBuffers(1,&b);
}

func (framebuffer Framebuffer) Delete()
{
    b := C.GLuint(framebuffer);
	C.glDeleteFramebuffers(1,&b);
}

func (program Program)  Delete()
{
	C.glDeleteProgram(C.GLuint(program));
}

func (renderbuffer Renderbuffer) Delete()
{
    b := C.GLuint(renderbuffer);
	C.glDeleteRenderbuffers(1,&b);
}

func DeleteShader(shader Shader) 
{
	C.glDeleteShader(C.GLuint(shader));
}

func (texture Texture) Delete()
{
    b := C.GLuint(texture);
	C.glDeleteTextures(1,&b);
}

func DepthFunc(func_ GLenum) 
{
	C.glDepthFunc(C.GLenum(func_));
}

func DepthMask(flag bool) 
{
	C.glDepthMask(glBool(flag));
}

func DepthRange(zNear GLclampf, zFar GLclampf) 
{
	C.glDepthRange(C.GLclampd(zNear), C.GLclampd(zFar));
}

func (program Program) DetachShader(shader Shader) 
{
	C.glDetachShader(C.GLuint(program), C.GLuint(shader));
}

func Disable(cap GLenum) 
{
	C.glDisable(C.GLenum(cap));
}

func DisableVertexAttribArray(index GLuint) 
{
	C.glDisableVertexAttribArray(C.GLuint(index));
}

func DrawArrays(mode GLenum, first GLint, count GLsizei) 
{
	C.glDrawArrays(C.GLenum(mode), C.GLint(first), C.GLsizei(count));
}

func DrawElements(mode GLenum, count GLsizei, type_ GLenum, offset uintptr) 
{
	C.glDrawElements(C.GLenum(mode), C.GLsizei(count), C.GLenum(type_), unsafe.Pointer(offset));
}

func Enable(cap GLenum) 
{
	C.glEnable(C.GLenum(cap));
}

func EnableVertexAttribArray(index GLuint) 
{
	C.glEnableVertexAttribArray(C.GLuint(index));
}

func Finish() 
{
	C.glFinish();
}

func Flush() 
{
	C.glFlush();
}

func FrontFace(mode GLenum) 
{
	C.glFrontFace(C.GLenum(mode));
}

func GenerateMipmap(target GLenum) 
{
	C.glGenerateMipmap(C.GLenum(target));
}

func GetAttachedShaders(program GLuint) []Object
{
	return nil;//	return Object[](C.getAttachedShaders(C.GLuint(program)));
}

type AttribLocation GLuint;

func (program Program) GetAttribLocation(name string) AttribLocation
{

    cname := glString(name);
    defer freeString(cname);

	return AttribLocation(C.glGetAttribLocation(C.GLuint(program), cname));
}

func GetError() GLenum
{
	return GLenum(C.glGetError());
}

func GetProgramInfoLog(program Program) string
{
	return "";//	return C.GoString(C.getProgramInfoLog(C.GLuint(program)));
}

func GetShaderInfoLog(shader Shader) string
{
	return "";//	return C.GoString(C.getShaderInfoLog(C.GLuint(shader)));
}

func GetShaderSource(shader Shader) string
{
	return "";//	return C.GoString(C.getShaderSource(C.GLuint(shader)));
}

func GetString(name GLenum) string
{
    s:=(*C.char)(unsafe.Pointer(C.glGetString(C.GLenum(name))));
	return C.GoString(s);
}

type UniformLocation GLint;

func (location UniformLocation) GetUniform(program Program) int
{
    return 0;
	//return int(C.glGetUniform(C.GLuint(program), C.GLuint(location)));
}

func (program Program) GetUniformLocation(name string) UniformLocation
{

    cname := glString(name);
    defer freeString(cname);

	return UniformLocation(C.glGetUniformLocation(C.GLuint(program), cname));
}


func GetVertexAttribOffset(index GLuint, pname GLenum) GLsizeiptr
{
    return 0;
	//return GLsizeiptr(C.getVertexAttribOffset(C.GLuint(index), C.GLenum(pname)));
}

func Hint(target GLenum, mode GLenum) 
{
	C.glHint(C.GLenum(target), C.GLenum(mode));
}

func IsEnabled(cap GLenum) bool
{
	return C.glIsEnabled(C.GLenum(cap))!=0;
}

func (buffer Object) IsBuffer() bool
{
	return C.glIsBuffer(C.GLuint(buffer))!=0;
}


func (framebuffer Object) IsFramebuffer() bool
{
	return C.glIsFramebuffer(C.GLuint(framebuffer))!=0;
}

func (program Object) IsProgram() bool
{
	return C.glIsProgram(C.GLuint(program))!=0;
}

func (renderbuffer Object) IsRenderbuffer() bool
{
	return C.glIsRenderbuffer(C.GLuint(renderbuffer))!=0;
}

func (shader Object) IsShader() bool
{
	return C.glIsShader(C.GLuint(shader))!=0;
}

func (texture Object) IsTexture() bool
{
	return C.glIsTexture(C.GLuint(texture))!=0;
}

func LineWidth(width GLfloat) 
{
	C.glLineWidth(C.GLfloat(width));
}

func (program Program) Link() 
{
	C.glLinkProgram(C.GLuint(program));
}

func PixelStorei(pname GLenum, param GLint) 
{
	C.glPixelStorei(C.GLenum(pname), C.GLint(param));
}

func PolygonOffset(factor GLfloat, units GLfloat) 
{
	C.glPolygonOffset(C.GLfloat(factor), C.GLfloat(units));
}

func SampleCoverage(value GLclampf, invert bool) 
{
	C.glSampleCoverage(C.GLclampf(value), glBool(invert));
}

func Scissor(x GLint, y GLint, width GLsizei, height GLsizei) 
{
	C.glScissor(C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height));
}

func (shader Shader) ShaderSource(source string) 
{

    csource := glString(source);
    defer freeString(csource);

    var one C.GLint=1;

	C.glShaderSource(C.GLuint(shader), 1, &csource,&one);
}

func StencilFunc(func_ GLenum, ref GLint, mask GLuint) 
{
	C.glStencilFunc(C.GLenum(func_), C.GLint(ref), C.GLuint(mask));
}

func StencilFuncSeparate(face GLenum, func_ GLenum, ref GLint, mask GLuint) 
{
	C.glStencilFuncSeparate(C.GLenum(face), C.GLenum(func_), C.GLint(ref), C.GLuint(mask));
}

func StencilMask(mask GLuint) 
{
	C.glStencilMask(C.GLuint(mask));
}

func StencilMaskSeparate(face GLenum, mask GLuint) 
{
	C.glStencilMaskSeparate(C.GLenum(face), C.GLuint(mask));
}

func StencilOp(fail GLenum, zfail GLenum, zpass GLenum) 
{
	C.glStencilOp(C.GLenum(fail), C.GLenum(zfail), C.GLenum(zpass));
}

func StencilOpSeparate(face GLenum, fail GLenum, zfail GLenum, zpass GLenum) 
{
	C.glStencilOpSeparate(C.GLenum(face), C.GLenum(fail), C.GLenum(zfail), C.GLenum(zpass));
}

func TexParameterf(target GLenum, pname GLenum, param GLfloat) 
{
	C.glTexParameterf(C.GLenum(target), C.GLenum(pname), C.GLfloat(param));
}

func TexParameteri(target GLenum, pname GLenum, param GLint) 
{
	C.glTexParameteri(C.GLenum(target), C.GLenum(pname), C.GLint(param));
}

func (location UniformLocation) setFloat(x GLfloat) 
{
	C.glUniform1f(C.GLint(location), C.GLfloat(x));
}

func (location UniformLocation) setFloatv(v []float) 
{
//	C.glUniform1fv(C.GLint(location), (*C.float)(v));
}

func (location UniformLocation) setInt(x GLint) 
{
	C.glUniform1i(C.GLint(location), C.GLint(x));
}

func (location UniformLocation) setIntv(v []int) 
{
//	C.glUniform1iv(C.GLint(location), (*C.int)(v));
}

func (location UniformLocation) Set2Float(x GLfloat, y GLfloat) 
{
	C.glUniform2f(C.GLint(location), C.GLfloat(x), C.GLfloat(y));
}

func Uniform2fv(location UniformLocation, v *float) 
{
//	C.glUniform2fv(C.GLint(location), (*C.float)(v));
}

func Uniform2i(location UniformLocation, x GLint, y GLint) 
{
	C.glUniform2i(C.GLint(location), C.GLint(x), C.GLint(y));
}

func Uniform2iv(location UniformLocation, v *int) 
{
//	C.glUniform2iv(C.GLint(location), (*C.int)(v));
}

func Uniform3f(location UniformLocation, x GLfloat, y GLfloat, z GLfloat) 
{
	C.glUniform3f(C.GLint(location), C.GLfloat(x), C.GLfloat(y), C.GLfloat(z));
}

func Uniform3fv(location UniformLocation, v *float) 
{
//	C.glUniform3fv(C.GLint(location), (*C.float)(v));
}

func Uniform3i(location UniformLocation, x GLint, y GLint, z GLint) 
{
	C.glUniform3i(C.GLint(location), C.GLint(x), C.GLint(y), C.GLint(z));
}

func Uniform3iv(location UniformLocation, v *int) 
{
//	C.glUniform3iv(C.GLint(location), (*C.int)(v));
}

func Uniform4f(location UniformLocation, x GLfloat, y GLfloat, z GLfloat, w GLfloat) 
{
	C.glUniform4f(C.GLint(location), C.GLfloat(x), C.GLfloat(y), C.GLfloat(z), C.GLfloat(w));
}

func Uniform4fv(location UniformLocation, v *float) 
{
//	C.glUniform4fv(C.GLint(location), (*C.float)(v));
}

func Uniform4i(location UniformLocation, x GLint, y GLint, z GLint, w GLint) 
{
	C.glUniform4i(C.GLint(location), C.GLint(x), C.GLint(y), C.GLint(z), C.GLint(w));
}

func Uniform4iv(location UniformLocation, v *int) 
{
//	C.glUniform4iv(C.GLint(location), (*C.int)(v));
}

func (program Program) Use() 
{
	C.glUseProgram(C.GLuint(program));
}

func (program Program) Validate()
{
	C.glValidateProgram(C.GLuint(program));
}

func VertexAttrib1f(indx GLuint, x GLfloat) 
{
	C.glVertexAttrib1f(C.GLuint(indx), C.GLfloat(x));
}

func VertexAttrib1fv(indx GLuint, values *float) 
{
//	C.glVertexAttrib1fv(C.GLuint(indx), (*C.float)(values));
}

func VertexAttrib2f(indx GLuint, x GLfloat, y GLfloat) 
{
	C.glVertexAttrib2f(C.GLuint(indx), C.GLfloat(x), C.GLfloat(y));
}

func VertexAttrib2fv(indx GLuint, values *float) 
{
//	C.glVertexAttrib2fv(C.GLuint(indx), (*C.float)(values));
}

func VertexAttrib3f(indx GLuint, x GLfloat, y GLfloat, z GLfloat) 
{
	C.glVertexAttrib3f(C.GLuint(indx), C.GLfloat(x), C.GLfloat(y), C.GLfloat(z));
}

func VertexAttrib3fv(indx GLuint, values *float) 
{
//	C.glVertexAttrib3fv(C.GLuint(indx), (*C.float)(values));
}

func VertexAttrib4f(indx GLuint, x GLfloat, y GLfloat, z GLfloat, w GLfloat) 
{
	C.glVertexAttrib4f(C.GLuint(indx), C.GLfloat(x), C.GLfloat(y), C.GLfloat(z), C.GLfloat(w));
}

func VertexAttrib4fv(indx GLuint, values *float) 
{
//	C.glVertexAttrib4fv(C.GLuint(indx), (*C.float)(values));
}

func Viewport(x GLint, y GLint, width GLsizei, height GLsizei) 
{
	C.glViewport(C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height));
}

