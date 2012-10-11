package gl

// #cgo darwin LDFLAGS: -framework OpenGL -lGLEW
// #cgo windows LDFLAGS: -lglew32 -lopengl32
// #cgo linux LDFLAGS: -lGLEW -lGL
//
// #include <stdlib.h>
//
// #include <GL/glew.h>
//
// #undef GLEW_GET_FUN
// #define GLEW_GET_FUN(x) (*x)
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

// Object

type Object C.GLuint

func (object Object) IsBuffer() bool { return C.glIsBuffer(C.GLuint(object)) != 0 }

func (object Object) IsProgram() bool { return C.glIsProgram(C.GLuint(object)) != 0 }

func (object Object) IsShader() bool { return C.glIsShader(C.GLuint(object)) != 0 }

func (object Object) IsTexture() bool { return C.glIsTexture(C.GLuint(object)) != 0 }

func (object Object) IsTransformFeedback() bool { return C.glIsTransformFeedback(C.GLuint(object)) != 0 }

func (object Object) IsVertexArray() bool { return C.glIsVertexArray(C.GLuint(object)) != 0 }

// Shader

type Shader Object

func CreateShader(type_ GLenum) Shader { return Shader(C.glCreateShader(C.GLenum(type_))) }

func (shader Shader) Delete() { C.glDeleteShader(C.GLuint(shader)) }

func (shader Shader) GetInfoLog() string {
	var length C.GLint
	C.glGetShaderiv(C.GLuint(shader), C.GLenum(INFO_LOG_LENGTH), &length)
	// length is buffer size including null character

	if length > 1 {
		log := C.malloc(C.size_t(length))
		defer C.free(log)
		C.glGetShaderInfoLog(C.GLuint(shader), C.GLsizei(length), nil, (*C.GLchar)(log))
		return C.GoString((*C.char)(log))
	}
	return ""
}

func (shader Shader) GetSource() string {
	var len C.GLint
	C.glGetShaderiv(C.GLuint(shader), C.GLenum(SHADER_SOURCE_LENGTH), &len)

	log := C.malloc(C.size_t(len + 1))
	C.glGetShaderSource(C.GLuint(shader), C.GLsizei(len), nil, (*C.GLchar)(log))

	defer C.free(log)

	return C.GoString((*C.char)(log))
}

func (shader Shader) Source(source string) {

	csource := glString(source)
	defer freeString(csource)

	var one C.GLint = C.GLint(len(source))

	C.glShaderSource(C.GLuint(shader), 1, &csource, &one)
}

func (shader Shader) Compile() { C.glCompileShader(C.GLuint(shader)) }

func (shader Shader) Get(param GLenum) int {
	var rv C.GLint

	C.glGetShaderiv(C.GLuint(shader), C.GLenum(param), &rv)
	return int(rv)
}

// Program

type Program Object

func CreateProgram() Program { return Program(C.glCreateProgram()) }

func (program Program) Delete() { C.glDeleteProgram(C.GLuint(program)) }

func (program Program) AttachShader(shader Shader) {
	C.glAttachShader(C.GLuint(program), C.GLuint(shader))
}

func (program Program) GetAttachedShaders() []Object {
	var len C.GLint
	C.glGetProgramiv(C.GLuint(program), C.GLenum(ACTIVE_UNIFORM_MAX_LENGTH), &len)

	objects := make([]Object, len)
	C.glGetAttachedShaders(C.GLuint(program), C.GLsizei(len), nil, *((**C.GLuint)(unsafe.Pointer(&objects))))
	return objects
}

func (program Program) DetachShader(shader Shader) {
	C.glDetachShader(C.GLuint(program), C.GLuint(shader))
}

func (program Program) TransformFeedbackVaryings(names []string, buffer_mode GLenum) {
	if len(names) == 0 {
		C.glTransformFeedbackVaryings(C.GLuint(program), 0, (**C.GLchar)(nil), C.GLenum(buffer_mode))
	} else {
		gl_names := make([]*C.GLchar, len(names))

		for i := range names {
			gl_names[i] = glString(names[i])
		}

		C.glTransformFeedbackVaryings(C.GLuint(program), C.GLsizei(len(gl_names)), &gl_names[0], C.GLenum(buffer_mode))

		for _, s := range gl_names {
			freeString(s)
		}
	}
}

func (program Program) Link() { C.glLinkProgram(C.GLuint(program)) }

func (program Program) Validate() { C.glValidateProgram(C.GLuint(program)) }

func (program Program) Use() { C.glUseProgram(C.GLuint(program)) }

func ProgramUnuse() { C.glUseProgram(C.GLuint(0)) }

func (program Program) GetInfoLog() string {
	var length C.GLint
	C.glGetProgramiv(C.GLuint(program), C.GLenum(INFO_LOG_LENGTH), &length)
	// length is buffer size including null character

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

func (program Program) GetUniformiv(location UniformLocation, values []int) {
	if len(values) == 0 {
		panic("Invalid values length")
	}
	// FIXME(jimt): This should really yield only one return value instead of using a slice.
	// http://www.opengl.org/sdk/docs/man/xhtml/glGetUniform.xml
	C.glGetUniformiv(C.GLuint(program), C.GLint(location), (*C.GLint)(unsafe.Pointer(&(values[0]))))
}

func (program Program) GetUniformfv(location UniformLocation, values []float32) {
	if len(values) == 0 {
		panic("Invalid values length")
	}
	// FIXME(jimt): This should really yield only one return value instead of using a slice.
	// http://www.opengl.org/sdk/docs/man/xhtml/glGetUniform.xml
	C.glGetUniformfv(C.GLuint(program), C.GLint(location), (*C.GLfloat)(unsafe.Pointer(&(values[0]))))
}

func (program Program) GetUniformLocation(name string) UniformLocation {

	cname := glString(name)
	defer freeString(cname)

	return UniformLocation(C.glGetUniformLocation(C.GLuint(program), cname))
}

func (program Program) GetAttribLocation(name string) AttribLocation {

	cname := glString(name)
	defer freeString(cname)

	return AttribLocation(C.glGetAttribLocation(C.GLuint(program), cname))
}

func (program Program) BindAttribLocation(index AttribLocation, name string) {

	cname := glString(name)
	defer freeString(cname)

	C.glBindAttribLocation(C.GLuint(program), C.GLuint(index), cname)

}

// Texture

type Texture Object

// Create single texture object
func GenTexture() Texture {
	var b C.GLuint
	C.glGenTextures(1, &b)
	return Texture(b)
}

// Fill slice with new textures
func GenTextures(textures []Texture) {
	if len(textures) == 0 {
		panic("Invalid texture slice length")
	}
	C.glGenTextures(C.GLsizei(len(textures)), (*C.GLuint)(&textures[0]))
}

// Delete texture object
func (texture Texture) Delete() {
	b := C.GLuint(texture)
	C.glDeleteTextures(1, &b)
}

// Delete all textures in slice
func DeleteTextures(textures []Texture) {
	if len(textures) == 0 {
		panic("Invalid texture slice length")
	}
	C.glDeleteTextures(C.GLsizei(len(textures)), (*C.GLuint)(&textures[0]))
}

// Bind this texture as target
func (texture Texture) Bind(target GLenum) {
	C.glBindTexture(C.GLenum(target), C.GLuint(texture))
}

// Unbind this texture
func (texture Texture) Unbind(target GLenum) {
	C.glBindTexture(C.GLenum(target), 0)
}

//void glTexImage1D (GLenum target, int level, int internalformat, int width, int border, GLenum format, GLenum type, const GLvoid *pixels)
func TexImage1D(target GLenum, level int, internalformat int, width int, border int, format, typ GLenum, pixels interface{}) {
	C.glTexImage1D(C.GLenum(target), C.GLint(level), C.GLint(internalformat),
		C.GLsizei(width), C.GLint(border), C.GLenum(format), C.GLenum(typ),
		ptr(pixels))
}

//void glTexImage2D (GLenum target, int level, int internalformat, int width, int height, int border, GLenum format, GLenum type, const GLvoid *pixels)
func TexImage2D(target GLenum, level int, internalformat int, width int, height int, border int, format, typ GLenum, pixels interface{}) {
	C.glTexImage2D(C.GLenum(target), C.GLint(level), C.GLint(internalformat),
		C.GLsizei(width), C.GLsizei(height), C.GLint(border), C.GLenum(format),
		C.GLenum(typ), ptr(pixels))
}

//void glTexImage3D (GLenum target, int level, int internalformat, int width, int height, int depth, int border, GLenum format, GLenum type, const GLvoid *pixels)
func TexImage3D(target GLenum, level int, internalformat int, width, height, depth int, border int, format, typ GLenum, pixels interface{}) {
	C.glTexImage3D(C.GLenum(target), C.GLint(level), C.GLint(internalformat),
		C.GLsizei(width), C.GLsizei(height), C.GLsizei(depth), C.GLint(border),
		C.GLenum(format), C.GLenum(typ), ptr(pixels))
}

//void glTexBuffer (GLenum target, GLenum internalformat, GLuint buffer)
func TexBuffer(target, internalformat GLenum, buffer Buffer) {
	C.glTexBuffer(C.GLenum(target), C.GLenum(internalformat), C.GLuint(buffer))
}

//void glPixelMapfv (GLenum map, int mapsize, const float *values)
func PixelMapfv(map_ GLenum, mapsize int, values *float32) {
	C.glPixelMapfv(C.GLenum(map_), C.GLsizei(mapsize), (*C.GLfloat)(values))
}

//void glPixelMapuiv (GLenum map, int mapsize, const uint *values)
func PixelMapuiv(map_ GLenum, mapsize int, values *uint32) {
	C.glPixelMapuiv(C.GLenum(map_), C.GLsizei(mapsize), (*C.GLuint)(values))
}

//void glPixelMapusv (GLenum map, int mapsize, const uint16 *values)
func PixelMapusv(map_ GLenum, mapsize int, values *uint16) {
	C.glPixelMapusv(C.GLenum(map_), C.GLsizei(mapsize), (*C.GLushort)(values))
}

//void glTexSubImage1D (GLenum target, int level, int xoffset, int width, GLenum format, GLenum type, const GLvoid *pixels)
func TexSubImage1D(target GLenum, level int, xoffset int, width int, format, typ GLenum, pixels interface{}) {
	C.glTexSubImage1D(C.GLenum(target), C.GLint(level), C.GLint(xoffset),
		C.GLsizei(width), C.GLenum(format), C.GLenum(typ), ptr(pixels))
}

//void glTexSubImage2D (GLenum target, int level, int xoffset, int yoffset, int width, int height, GLenum format, GLenum type, const GLvoid *pixels)
func TexSubImage2D(target GLenum, level int, xoffset int, yoffset int, width int, height int, format, typ GLenum, pixels interface{}) {
	C.glTexSubImage2D(C.GLenum(target), C.GLint(level), C.GLint(xoffset),
		C.GLint(yoffset), C.GLsizei(width), C.GLsizei(height), C.GLenum(format),
		C.GLenum(typ), ptr(pixels))
}

//void glTexImage3D (GLenum target, int level, int xoffset, int yoffset, int zoffset, int width, int height, int depth, GLenum format, GLenum type, const GLvoid *pixels)
func TexSubImage3D(target GLenum, level int, xoffset, yoffset, zoffset, width, height, depth int, format, typ GLenum, pixels interface{}) {
	C.glTexSubImage3D(C.GLenum(target), C.GLint(level),
		C.GLint(xoffset), C.GLint(yoffset), C.GLint(zoffset),
		C.GLsizei(width), C.GLsizei(height), C.GLsizei(depth),
		C.GLenum(format), C.GLenum(typ), ptr(pixels))
}

//void glCopyTexImage1D (GLenum target, int level, GLenum internalFormat, int x, int y, int width, int border)
func CopyTexImage1D(target GLenum, level int, internalFormat GLenum, x int, y int, width int, border int) {
	C.glCopyTexImage1D(C.GLenum(target), C.GLint(level), C.GLenum(internalFormat), C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLint(border))
}

//void glCopyTexImage2D (GLenum target, int level, GLenum internalFormat, int x, int y, int width, int height, int border)
func CopyTexImage2D(target GLenum, level int, internalFormat GLenum, x int, y int, width int, height int, border int) {
	C.glCopyTexImage2D(C.GLenum(target), C.GLint(level), C.GLenum(internalFormat), C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height), C.GLint(border))
}

//void glCopyTexSubImage1D (GLenum target, int level, int xoffset, int x, int y, int width)
func CopyTexSubImage1D(target GLenum, level int, xoffset int, x int, y int, width int) {
	C.glCopyTexSubImage1D(C.GLenum(target), C.GLint(level), C.GLint(xoffset), C.GLint(x), C.GLint(y), C.GLsizei(width))
}

//void glCopyTexSubImage2D (GLenum target, int level, int xoffset, int yoffset, int x, int y, int width, int height)
func CopyTexSubImage2D(target GLenum, level int, xoffset int, yoffset int, x int, y int, width int, height int) {
	C.glCopyTexSubImage2D(C.GLenum(target), C.GLint(level), C.GLint(xoffset), C.GLint(yoffset), C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height))
}

// TODO 3D textures

//void glTexEnvf (GLenum target, GLenum pname, float32 param)
func TexEnvf(target GLenum, pname GLenum, param float32) {
	C.glTexEnvf(C.GLenum(target), C.GLenum(pname), C.GLfloat(param))
}

//void glTexEnvfv (GLenum target, GLenum pname, const float *params)
func TexEnvfv(target GLenum, pname GLenum, params []float32) {
	if len(params) == 0 {
		panic("Invalid params slice length")
	}
	C.glTexEnvfv(C.GLenum(target), C.GLenum(pname), (*C.GLfloat)(&params[0]))
}

//void glTexEnvi (GLenum target, GLenum pname, int param)
func TexEnvi(target GLenum, pname GLenum, param int) {
	C.glTexEnvi(C.GLenum(target), C.GLenum(pname), C.GLint(param))
}

//void glTexEnviv (GLenum target, GLenum pname, const int *params)
func TexEnviv(target GLenum, pname GLenum, params []int32) {
	if len(params) == 0 {
		panic("Invalid params slice length")
	}
	C.glTexEnviv(C.GLenum(target), C.GLenum(pname), (*C.GLint)(&params[0]))
}

//void glTexGend (GLenum coord, GLenum pname, float64 param)
func TexGend(coord GLenum, pname GLenum, param float64) {
	C.glTexGend(C.GLenum(coord), C.GLenum(pname), C.GLdouble(param))
}

//void glTexGendv (GLenum coord, GLenum pname, const float64 *params)
func TexGendv(coord GLenum, pname GLenum, params []float64) {
	if len(params) == 0 {
		panic("Invalid params slice length")
	}
	C.glTexGendv(C.GLenum(coord), C.GLenum(pname), (*C.GLdouble)(&params[0]))
}

//void glTexGenf (GLenum coord, GLenum pname, float32 param)
func TexGenf(coord GLenum, pname GLenum, param float32) {
	C.glTexGenf(C.GLenum(coord), C.GLenum(pname), C.GLfloat(param))
}

//void glTexGenfv (GLenum coord, GLenum pname, const float *params)
func TexGenfv(coord GLenum, pname GLenum, params []float32) {
	if len(params) == 0 {
		panic("Invalid params slice length")
	}
	C.glTexGenfv(C.GLenum(coord), C.GLenum(pname), (*C.GLfloat)(&params[0]))
}

//void glTexGeni (GLenum coord, GLenum pname, int param)
func TexGeni(coord GLenum, pname GLenum, param int) {
	C.glTexGeni(C.GLenum(coord), C.GLenum(pname), C.GLint(param))
}

//void glTexGeniv (GLenum coord, GLenum pname, const int *params)
func TexGeniv(coord GLenum, pname GLenum, params []int32) {
	if len(params) == 0 {
		panic("Invalid params slice length")
	}
	C.glTexGeniv(C.GLenum(coord), C.GLenum(pname), (*C.GLint)(&params[0]))
}

//void glTexParameterf (GLenum target, GLenum pname, float32 param)
func TexParameterf(target GLenum, pname GLenum, param float32) {
	C.glTexParameterf(C.GLenum(target), C.GLenum(pname), C.GLfloat(param))
}

//void glTexParameterfv (GLenum target, GLenum pname, const float *params)
func TexParameterfv(target GLenum, pname GLenum, params []float32) {
	if len(params) == 0 {
		panic("Invalid params slice length")
	}
	C.glTexParameterfv(C.GLenum(target), C.GLenum(pname), (*C.GLfloat)(&params[0]))
}

//void glTexParameteri (GLenum target, GLenum pname, int param)
func TexParameteri(target GLenum, pname GLenum, param int) {
	C.glTexParameteri(C.GLenum(target), C.GLenum(pname), C.GLint(param))
}

//void glTexParameteriv (GLenum target, GLenum pname, const int *params)
func TexParameteriv(target GLenum, pname GLenum, params []int32) {
	if len(params) == 0 {
		panic("Invalid params slice length")
	}
	C.glTexParameteriv(C.GLenum(target), C.GLenum(pname), (*C.GLint)(&params[0]))
}

//void glPrioritizeTextures (GLsizei n, const uint *textures, const GLclampf *priorities)
func PrioritizeTextures(n int, textures *uint32, priorities *GLclampf) {
	C.glPrioritizeTextures(C.GLsizei(n), (*C.GLuint)(textures), (*C.GLclampf)(priorities))
}

//void glGetTexEnvfv (GLenum target, GLenum pname, float *params)
func GetTexEnvfv(target GLenum, pname GLenum, params []float32) {
	if len(params) == 0 {
		panic("Invalid params slice length")
	}
	C.glGetTexEnvfv(C.GLenum(target), C.GLenum(pname), (*C.GLfloat)(&params[0]))
}

//void glGetTexEnviv (GLenum target, GLenum pname, int *params)
func GetTexEnviv(target GLenum, pname GLenum, params []int32) {
	if len(params) == 0 {
		panic("Invalid params slice length")
	}
	C.glGetTexEnviv(C.GLenum(target), C.GLenum(pname), (*C.GLint)(&params[0]))
}

//void glGetTexGendv (GLenum coord, GLenum pname, float64 *params)
func GetTexGendv(coord GLenum, pname GLenum, params []float64) {
	if len(params) == 0 {
		panic("Invalid params slice length")
	}
	C.glGetTexGendv(C.GLenum(coord), C.GLenum(pname), (*C.GLdouble)(&params[0]))
}

//void glGetTexGenfv (GLenum coord, GLenum pname, float *params)
func GetTexGenfv(coord GLenum, pname GLenum, params []float32) {
	if len(params) == 0 {
		panic("Invalid params slice length")
	}
	C.glGetTexGenfv(C.GLenum(coord), C.GLenum(pname), (*C.GLfloat)(&params[0]))
}

//void glGetTexGeniv (GLenum coord, GLenum pname, int *params)
func GetTexGeniv(coord GLenum, pname GLenum, params []int32) {
	if len(params) == 0 {
		panic("Invalid params slice length")
	}
	C.glGetTexGeniv(C.GLenum(coord), C.GLenum(pname), (*C.GLint)(&params[0]))
}

//void glGetTexImage (GLenum target, int level, GLenum format, GLenum type, GLvoid *pixels)
func GetTexImage(target GLenum, level int, format, typ GLenum, pixels interface{}) {
	C.glGetTexImage(C.GLenum(target), C.GLint(level), C.GLenum(format),
		C.GLenum(typ), ptr(pixels))
}

//void glGetTexLevelParameterfv (GLenum target, int level, GLenum pname, float *params)
func GetTexLevelParameterfv(target GLenum, level int, pname GLenum, params []float32) {
	if len(params) == 0 {
		panic("Invalid params slice length")
	}
	C.glGetTexLevelParameterfv(C.GLenum(target), C.GLint(level), C.GLenum(pname), (*C.GLfloat)(&params[0]))
}

//void glGetTexLevelParameteriv (GLenum target, int level, GLenum pname, int *params)
func GetTexLevelParameteriv(target GLenum, level int, pname GLenum, params []int32) {
	if len(params) == 0 {
		panic("Invalid params slice length")
	}
	C.glGetTexLevelParameteriv(C.GLenum(target), C.GLint(level), C.GLenum(pname), (*C.GLint)(&params[0]))
}

//void glGetTexParameterfv (GLenum target, GLenum pname, float *params)
func GetTexParameterfv(target GLenum, pname GLenum, params []float32) {
	if len(params) == 0 {
		panic("Invalid params slice length")
	}
	C.glGetTexParameterfv(C.GLenum(target), C.GLenum(pname), (*C.GLfloat)(&params[0]))
}

//void glGetTexParameteriv (GLenum target, GLenum pname, int *params)
func GetTexParameteriv(target GLenum, pname GLenum, params []int32) {
	if len(params) == 0 {
		panic("Invalid params slice length")
	}
	C.glGetTexParameteriv(C.GLenum(target), C.GLenum(pname), (*C.GLint)(&params[0]))
}

// Buffer Objects

type Buffer Object

// Create single buffer object
func GenBuffer() Buffer {
	var b C.GLuint
	C.glGenBuffers(1, &b)
	return Buffer(b)
}

// Fill slice with new buffers
func GenBuffers(buffers []Buffer) {
	if len(buffers) == 0 {
		panic("Invalid buffer slice length")
	}
	C.glGenBuffers(C.GLsizei(len(buffers)), (*C.GLuint)(&buffers[0]))
}

// Delete buffer object
func (buffer Buffer) Delete() {
	b := C.GLuint(buffer)
	C.glDeleteBuffers(1, &b)
}

// Delete all textures in slice
func DeleteBuffers(buffers []Buffer) {
	if len(buffers) == 0 {
		panic("Invalid buffer slice length")
	}
	C.glDeleteBuffers(C.GLsizei(len(buffers)), (*C.GLuint)(&buffers[0]))
}

// Remove buffer binding
func BufferUnbind(target GLenum) {
	C.glBindBuffer(C.GLenum(target), C.GLuint(0))
}

// Bind this buffer as target
func (buffer Buffer) Bind(target GLenum) {
	C.glBindBuffer(C.GLenum(target), C.GLuint(buffer))
}

// Bind this buffer as index of target
func (buffer Buffer) BindBufferBase(target GLenum, index uint) {
	C.glBindBufferBase(C.GLenum(target), C.GLuint(index), C.GLuint(buffer))
}

// Bind this buffer range as index of target
func (buffer Buffer) BindBufferRange(target GLenum, index uint, offset int, size uint) {
	C.glBindBufferRange(C.GLenum(target), C.GLuint(index), C.GLuint(buffer), C.GLintptr(offset), C.GLsizeiptr(size))
}

// Creates and initializes a buffer object's data store
func BufferData(target GLenum, size int, data interface{}, usage GLenum) {
	C.glBufferData(C.GLenum(target), C.GLsizeiptr(size), ptr(data), C.GLenum(usage))
}

//  Update a subset of a buffer object's data store
func BufferSubData(target GLenum, offset int, size int, data interface{}) {
	C.glBufferSubData(C.GLenum(target), C.GLintptr(offset), C.GLsizeiptr(size),
		ptr(data))
}

// Returns a subset of a buffer object's data store
func GetBufferSubData(target GLenum, offset int, size int, data interface{}) {
	C.glGetBufferSubData(C.GLenum(target), C.GLintptr(offset),
		C.GLsizeiptr(size), ptr(data))
}

//  Map a buffer object's data store
func MapBuffer(target GLenum, access GLenum) {
	C.glMapBuffer(C.GLenum(target), C.GLenum(access))
}

//  Unmap a buffer object's data store
func UnmapBuffer(target GLenum) bool {
	return goBool(C.glUnmapBuffer(C.GLenum(target)))
}

// Return buffer pointer
func glGetBufferPointerv(target GLenum, pname GLenum, params []unsafe.Pointer) {
	if len(params) == 0 {
		panic("Invalid param slice length")
	}
	C.glGetBufferPointerv(C.GLenum(target), C.GLenum(pname), &params[0])
}

// Return parameters of a buffer object
func GetBufferParameteriv(target GLenum, pname GLenum, params []int32) {
	if len(params) == 0 {
		panic("Invalid param slice length")
	}
	C.glGetBufferParameteriv(C.GLenum(target), C.GLenum(pname), (*C.GLint)(&params[0]))
}

// Transform Feedback Objects

type TransformFeedback Object

// Create a single transform feedback object
func GenTransformFeedback() TransformFeedback {
	var t C.GLuint
	C.glGenTransformFeedbacks(1, &t)
	return TransformFeedback(t)
}

// Fill slice with new transform feedbacks
func GenTransformFeedbacks(feedbacks []TransformFeedback) {
	if len(feedbacks) == 0 {
		panic("Invalid feedback slice length")
	}
	// FIXME(jimt): glGenBuffers is correct?
	// If so, this is a duplicate of GenBuffers().
	C.glGenBuffers(C.GLsizei(len(feedbacks)), (*C.GLuint)(&feedbacks[0]))
}

// Delete a transform feedback object
func (feedback TransformFeedback) Delete() {
	C.glDeleteTransformFeedbacks(1, (*C.GLuint)(&feedback))
}

// Draw the results of the last Begin/End cycle from this transform feedback using primitive type 'mode'
func (feedback TransformFeedback) Draw(mode GLenum) {
	C.glDrawTransformFeedback(C.GLenum(mode), C.GLuint(feedback))
}

// Delete all transform feedbacks in a slice
func DeleteTransformFeedbacks(feedbacks []TransformFeedback) {
	if len(feedbacks) == 0 {
		panic("Invalid feedback slice length")
	}
	C.glDeleteTransformFeedbacks(C.GLsizei(len(feedbacks)), (*C.GLuint)(&feedbacks[0]))
}

// Bind this transform feedback as target
func (feedback TransformFeedback) Bind(target GLenum) {
	C.glBindTransformFeedback(C.GLenum(target), C.GLuint(feedback))
}

// Begin transform feedback with primitive type 'mode'
func BeginTransformFeedback(mode GLenum) {
	C.glBeginTransformFeedback(C.GLenum(mode))
}

// Pause transform feedback
func PauseTransformFeedback() {
	C.glPauseTransformFeedback()
}

// End transform feedback
func EndTransformFeedback() {
	C.glEndTransformFeedback()
}

// AttribLocation

type AttribLocation int

func (indx AttribLocation) Attrib1f(x float32) {
	C.glVertexAttrib1f(C.GLuint(indx), C.GLfloat(x))
}

func (indx AttribLocation) Attrib1fv(values []float32) {
	if len(values) == 0 {
		panic("Invalid value slice length")
	}
	C.glVertexAttrib1fv(C.GLuint(indx), (*C.GLfloat)(unsafe.Pointer(&values[0])))
}

func (indx AttribLocation) Attrib2f(x float32, y float32) {
	C.glVertexAttrib2f(C.GLuint(indx), C.GLfloat(x), C.GLfloat(y))
}

func (indx AttribLocation) Attrib2fv(values []float32) {
	if len(values) == 0 {
		panic("Invalid value slice length")
	}
	C.glVertexAttrib2fv(C.GLuint(indx), (*C.GLfloat)(unsafe.Pointer(&values[0])))
}

func (indx AttribLocation) Attrib3f(x float32, y float32, z float32) {
	C.glVertexAttrib3f(C.GLuint(indx), C.GLfloat(x), C.GLfloat(y), C.GLfloat(z))
}

func (indx AttribLocation) Attrib3fv(values []float32) {
	if len(values) == 0 {
		panic("Invalid value slice length")
	}
	C.glVertexAttrib3fv(C.GLuint(indx), (*C.GLfloat)(unsafe.Pointer(&values[0])))
}

func (indx AttribLocation) Attrib4f(x float32, y float32, z float32, w float32) {
	C.glVertexAttrib4f(C.GLuint(indx), C.GLfloat(x), C.GLfloat(y), C.GLfloat(z), C.GLfloat(w))
}

func (indx AttribLocation) Attrib4fv(values []float32) {
	if len(values) == 0 {
		panic("Invalid value slice length")
	}
	C.glVertexAttrib4fv(C.GLuint(indx), (*C.GLfloat)(unsafe.Pointer(&values[0])))
}

func (indx AttribLocation) AttribPointer(size uint, typ GLenum, normalized bool, stride int, pointer interface{}) {
	C.glVertexAttribPointer(C.GLuint(indx), C.GLint(size), C.GLenum(typ),
		glBool(normalized), C.GLsizei(stride), ptr(pointer))
}

func (indx AttribLocation) EnableArray() {
	C.glEnableVertexAttribArray(C.GLuint(indx))
}

func (indx AttribLocation) DisableArray() {
	C.glDisableVertexAttribArray(C.GLuint(indx))
}

// Vertex Arrays
type VertexArray Object

func GenVertexArray() VertexArray {
	var a C.GLuint
	C.glGenVertexArrays(1, &a)
	return VertexArray(a)
}

func GenVertexArrays(arrays []VertexArray) {
	if len(arrays) == 0 {
		panic("Invalid array slice length")
	}
	C.glGenVertexArrays(C.GLsizei(len(arrays)), (*C.GLuint)(&arrays[0]))
}

func (array VertexArray) Delete() {
	C.glDeleteVertexArrays(1, (*C.GLuint)(&array))
}

func DeleteVertexArrays(arrays []VertexArray) {
	if len(arrays) == 0 {
		panic("Invalid array slice length")
	}
	C.glDeleteVertexArrays(C.GLsizei(len(arrays)), (*C.GLuint)(&arrays[0]))
}

func (array VertexArray) Bind() {
	C.glBindVertexArray(C.GLuint(array))
}

// UniformLocation
//TODO

type UniformLocation int

func (location UniformLocation) Uniform1f(x float32) {
	C.glUniform1f(C.GLint(location), C.GLfloat(x))
}

func (location UniformLocation) Uniform2f(x float32, y float32) {
	C.glUniform2f(C.GLint(location), C.GLfloat(x), C.GLfloat(y))
}

func (location UniformLocation) Uniform3f(x float32, y float32, z float32) {
	C.glUniform3f(C.GLint(location), C.GLfloat(x), C.GLfloat(y), C.GLfloat(z))
}

func (location UniformLocation) Uniform1fv(v []float32) {
	C.glUniform1fv(C.GLint(location), C.GLsizei(len(v)),
		(*C.GLfloat)(ptr(v)))
}

func (location UniformLocation) Uniform1i(x int) {
	C.glUniform1i(C.GLint(location), C.GLint(x))
}

func (location UniformLocation) Uniform1iv(v []int) {
	panic("unimplemented")
	//	C.glUniform1iv(C.GLint(location), (*C.int)(&v[0]));
}

func (location UniformLocation) Uniform2fv(v []float32) {
	panic("unimplemented")
	//	C.glUniform2fv(C.GLint(location), (*C.float)(&v[0]));
}

func (location UniformLocation) Uniform2i(x int, y int) {
	C.glUniform2i(C.GLint(location), C.GLint(x), C.GLint(y))
}

func (location UniformLocation) Uniform2iv(v []int32) {
	panic("unimplemented")
	//	C.glUniform2iv(C.GLint(location), (*C.int)(&v[0]));
}

func (location UniformLocation) Uniform3fv(v []float32) {
	panic("unimplemented")
	//	C.glUniform3fv(C.GLint(location), (*C.float)(&v[0]));
}

func (location UniformLocation) Uniform3i(x int, y int, z int) {
	C.glUniform3i(C.GLint(location), C.GLint(x), C.GLint(y), C.GLint(z))
}

func (location UniformLocation) Uniform3iv(v []int32) {
	panic("unimplemented")
	//	C.glUniform3iv(C.GLint(location), (*C.int)(&v[0]));
}

func (location UniformLocation) Uniform4f(x float32, y float32, z float32, w float32) {
	C.glUniform4f(C.GLint(location), C.GLfloat(x), C.GLfloat(y), C.GLfloat(z), C.GLfloat(w))
}

func (location UniformLocation) Uniform4fv(v []float32) {
	panic("unimplemented")
	//	C.glUniform4fv(C.GLint(location), (*C.float)(&v[0]));
}

func (location UniformLocation) Uniform4i(x int, y int, z int, w int) {
	C.glUniform4i(C.GLint(location), C.GLint(x), C.GLint(y), C.GLint(z), C.GLint(w))
}

func (location UniformLocation) Uniform4iv(v []int32) {
	panic("unimplemented")
	//	C.glUniform4iv(C.GLint(location), (*C.int)(&v[0]));
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

func ActiveTexture(texture GLenum) { C.glActiveTexture(C.GLenum(texture)) }

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

//void glAccum (GLenum op, float32 value)
func Accum(op GLenum, value float32) {
	C.glAccum(C.GLenum(op), C.GLfloat(value))
}

//void glAlphaFunc (GLenum func, GLclampf ref)
func AlphaFunc(func_ GLenum, ref GLclampf) {
	C.glAlphaFunc(C.GLenum(func_), C.GLclampf(ref))
}

//bool glAreTexturesResident (GLsizei n, const uint *textures, bool *residences)
func AreTexturesResident(n int, textures *uint, residences *bool) bool {
	//TODO
	panic("unimlemented")
	//return C.glAreTexturesResident(C.GLsizei(n), (*C.GLuint)(unsafe.Pointer(textures)), (*C.GLboolean)(unsafe.Pointer(residences))) != 0

}

//void glArrayElement (int i)
func ArrayElement(i int) {
	C.glArrayElement(C.GLint(i))
}

//void glBegin (GLenum mode)
func Begin(mode GLenum) {
	C.glBegin(C.GLenum(mode))
}

//void glBindTexture (GLenum target, uint texture)
func BindTexture(target GLenum, texture uint) {
	C.glBindTexture(C.GLenum(target), C.GLuint(texture))
}

//void glBitmap (GLsizei width, int height, float32 xorig, float32 yorig, float32 xmove, float32 ymove, const uint8 *bitmap)
func Bitmap(width int, height int, xorig float32, yorig float32, xmove float32, ymove float32, bitmap *uint8) {
	C.glBitmap(C.GLsizei(width), C.GLsizei(height), C.GLfloat(xorig), C.GLfloat(yorig), C.GLfloat(xmove), C.GLfloat(ymove), (*C.GLubyte)(bitmap))
}

//void glBlendFunc (GLenum sfactor, GLenum dfactor)
func BlendFunc(sfactor GLenum, dfactor GLenum) {
	C.glBlendFunc(C.GLenum(sfactor), C.GLenum(dfactor))
}

//void glCallList (uint list)
func CallList(list uint) {
	C.glCallList(C.GLuint(list))
}

//void glCallLists (GLsizei n, GLenum type, const GLvoid *lists)
func CallLists(n int, typ GLenum, lists interface{}) {
	C.glCallLists(C.GLsizei(n), C.GLenum(typ), ptr(lists))
}

//void glClear (GLbitfield mask)
func Clear(mask GLbitfield) {
	C.glClear(C.GLbitfield(mask))
}

//void glClearAccum (float32 red, float32 green, float32 blue, float32 alpha)
func ClearAccum(red float32, green float32, blue float32, alpha float32) {
	C.glClearAccum(C.GLfloat(red), C.GLfloat(green), C.GLfloat(blue), C.GLfloat(alpha))
}

//void glClearColor (GLclampf red, GLclampf green, GLclampf blue, GLclampf alpha)
func ClearColor(red GLclampf, green GLclampf, blue GLclampf, alpha GLclampf) {
	C.glClearColor(C.GLclampf(red), C.GLclampf(green), C.GLclampf(blue), C.GLclampf(alpha))
}

//void glClearDepth (GLclampd depth)
func ClearDepth(depth GLclampd) {
	C.glClearDepth(C.GLclampd(depth))
}

//void glClearIndex (float32 c)
func ClearIndex(c float32) {
	C.glClearIndex(C.GLfloat(c))
}

//void glClearStencil (int s)
func ClearStencil(s int) {
	C.glClearStencil(C.GLint(s))
}

//void glClipPlane (GLenum plane, const float64 *equation)
func ClipPlane(plane GLenum, equation *float64) {
	C.glClipPlane(C.GLenum(plane), (*C.GLdouble)(equation))
}

//void glColor3b (int8 red, int8 green, int8 blue)
func Color3b(red int8, green int8, blue int8) {
	C.glColor3b(C.GLbyte(red), C.GLbyte(green), C.GLbyte(blue))
}

//void glColor3bv (const int8 *v)
func Color3bv(v []int8) {
	if len(v) != 3 {
		panic("Invalid color length")
	}
	C.glColor3bv((*C.GLbyte)(&v[0]))
}

//void glColor3d (float64 red, float64 green, float64 blue)
func Color3d(red float64, green float64, blue float64) {
	C.glColor3d(C.GLdouble(red), C.GLdouble(green), C.GLdouble(blue))
}

//void glColor3dv (const float64 *v)
func Color3dv(v []float64) {
	if len(v) != 3 {
		panic("Invalid color length")
	}
	C.glColor3dv((*C.GLdouble)(&v[0]))
}

//void glColor3f (float32 red, float32 green, float32 blue)
func Color3f(red float32, green float32, blue float32) {
	C.glColor3f(C.GLfloat(red), C.GLfloat(green), C.GLfloat(blue))
}

//void glColor3fv (const float *v)
func Color3fv(v []float32) {
	if len(v) != 3 {
		panic("Invalid color length")
	}
	C.glColor3fv((*C.GLfloat)(&v[0]))
}

//void glColor3i (int red, int green, int blue)
func Color3i(red int, green int, blue int) {
	C.glColor3i(C.GLint(red), C.GLint(green), C.GLint(blue))
}

//void glColor3iv (const int *v)
func Color3iv(v []int32) {
	if len(v) != 3 {
		panic("Invalid color length")
	}
	C.glColor3iv((*C.GLint)(&v[0]))
}

//void glColor3s (int16 red, int16 green, int16 blue)
func Color3s(red int16, green int16, blue int16) {
	C.glColor3s(C.GLshort(red), C.GLshort(green), C.GLshort(blue))
}

//void glColor3sv (const int16 *v)
func Color3sv(v []int16) {
	C.glColor3sv((*C.GLshort)(&v[0]))
}

//void glColor3ub (uint8 red, uint8 green, uint8 blue)
func Color3ub(red uint8, green uint8, blue uint8) {
	C.glColor3ub(C.GLubyte(red), C.GLubyte(green), C.GLubyte(blue))
}

//void glColor3ubv (const uint8 *v)
func Color3ubv(v []uint8) {
	if len(v) != 3 {
		panic("Invalid color length")
	}
	C.glColor3ubv((*C.GLubyte)(&v[0]))
}

//void glColor3ui (uint red, uint green, uint blue)
func Color3ui(red uint, green uint, blue uint) {
	C.glColor3ui(C.GLuint(red), C.GLuint(green), C.GLuint(blue))
}

//void glColor3uiv (const uint *v)
func Color3uiv(v []uint32) {
	if len(v) != 3 {
		panic("Invalid color length")
	}
	C.glColor3uiv((*C.GLuint)(&v[0]))
}

//void glColor3us (uint16 red, uint16 green, uint16 blue)
func Color3us(red uint16, green uint16, blue uint16) {
	C.glColor3us(C.GLushort(red), C.GLushort(green), C.GLushort(blue))
}

//void glColor3usv (const uint16 *v)
func Color3usv(v []uint16) {
	if len(v) != 3 {
		panic("Invalid color length")
	}
	C.glColor3usv((*C.GLushort)(&v[0]))
}

//void glColor4b (int8 red, int8 green, int8 blue, int8 alpha)
func Color4b(red int8, green int8, blue int8, alpha int8) {
	C.glColor4b(C.GLbyte(red), C.GLbyte(green), C.GLbyte(blue), C.GLbyte(alpha))
}

//void glColor4bv (const int8 *v)
func Color4bv(v []int8) {
	if len(v) != 4 {
		panic("Invalid color length")
	}
	C.glColor4bv((*C.GLbyte)(&v[0]))
}

//void glColor4d (float64 red, float64 green, float64 blue, float64 alpha)
func Color4d(red float64, green float64, blue float64, alpha float64) {
	C.glColor4d(C.GLdouble(red), C.GLdouble(green), C.GLdouble(blue), C.GLdouble(alpha))
}

//void glColor4dv (const float64 *v)
func Color4dv(v []float64) {
	if len(v) != 4 {
		panic("Invalid color length")
	}
	C.glColor4dv((*C.GLdouble)(&v[0]))
}

//void glColor4f (float32 red, float32 green, float32 blue, float32 alpha)
func Color4f(red float32, green float32, blue float32, alpha float32) {
	C.glColor4f(C.GLfloat(red), C.GLfloat(green), C.GLfloat(blue), C.GLfloat(alpha))
}

//void glColor4fv (const float *v)
func Color4fv(v []float32) {
	if len(v) != 4 {
		panic("Invalid color length")
	}
	C.glColor4fv((*C.GLfloat)(&v[0]))
}

//void glColor4i (int red, int green, int blue, int alpha)
func Color4i(red int, green int, blue int, alpha int) {
	C.glColor4i(C.GLint(red), C.GLint(green), C.GLint(blue), C.GLint(alpha))
}

//void glColor4iv (const int *v)
func Color4iv(v []int32) {
	if len(v) != 4 {
		panic("Invalid color length")
	}
	C.glColor4iv((*C.GLint)(&v[0]))
}

//void glColor4s (int16 red, int16 green, int16 blue, int16 alpha)
func Color4s(red int16, green int16, blue int16, alpha int16) {
	C.glColor4s(C.GLshort(red), C.GLshort(green), C.GLshort(blue), C.GLshort(alpha))
}

//void glColor4sv (const int16 *v)
func Color4sv(v []int16) {
	if len(v) != 4 {
		panic("Invalid color length")
	}
	C.glColor4sv((*C.GLshort)(&v[0]))
}

//void glColor4ub (uint8 red, uint8 green, uint8 blue, uint8 alpha)
func Color4ub(red uint8, green uint8, blue uint8, alpha uint8) {
	C.glColor4ub(C.GLubyte(red), C.GLubyte(green), C.GLubyte(blue), C.GLubyte(alpha))
}

//void glColor4ubv (const uint8 *v)
func Color4ubv(v []uint8) {
	if len(v) != 4 {
		panic("Invalid color length")
	}
	C.glColor4ubv((*C.GLubyte)(&v[0]))
}

//void glColor4ui (uint red, uint green, uint blue, uint alpha)
func Color4ui(red uint, green uint, blue uint, alpha uint) {
	C.glColor4ui(C.GLuint(red), C.GLuint(green), C.GLuint(blue), C.GLuint(alpha))
}

//void glColor4uiv (const uint *v)
func Color4uiv(v []uint32) {
	if len(v) != 4 {
		panic("Invalid color length")
	}
	C.glColor4uiv((*C.GLuint)(&v[0]))
}

//void glColor4us (uint16 red, uint16 green, uint16 blue, uint16 alpha)
func Color4us(red uint16, green uint16, blue uint16, alpha uint16) {
	C.glColor4us(C.GLushort(red), C.GLushort(green), C.GLushort(blue), C.GLushort(alpha))
}

//void glColor4usv (const uint16 *v)
func Color4usv(v []uint16) {
	if len(v) != 4 {
		panic("Invalid color length")
	}
	C.glColor4usv((*C.GLushort)(&v[0]))
}

//void glColorMask (bool red, bool green, bool blue, bool alpha)
func ColorMask(red bool, green bool, blue bool, alpha bool) {
	C.glColorMask(glBool(red), glBool(green), glBool(blue), glBool(alpha))
}

//void glColorMaterial (GLenum face, GLenum mode)
func ColorMaterial(face GLenum, mode GLenum) {
	C.glColorMaterial(C.GLenum(face), C.GLenum(mode))
}

//void glColorPointer (int size, GLenum type, int stride, const GLvoid *pointer)
func ColorPointer(size int, typ GLenum, stride int, pointer interface{}) {
	C.glColorPointer(C.GLint(size), C.GLenum(typ), C.GLsizei(stride),
		ptr(pointer))
}

//void glCopyPixels (int x, int y, int width, int height, GLenum type)
func CopyPixels(x int, y int, width int, height int, type_ GLenum) {
	C.glCopyPixels(C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height), C.GLenum(type_))
}

//void glCullFace (GLenum mode)
func CullFace(mode GLenum) {
	C.glCullFace(C.GLenum(mode))
}

//void glDeleteLists (uint list, int range)
func DeleteLists(list uint, range_ int) {
	C.glDeleteLists(C.GLuint(list), C.GLsizei(range_))
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

//void glDisableClientState (GLenum array)
func DisableClientState(array GLenum) {
	C.glDisableClientState(C.GLenum(array))
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

//void glDrawPixels (GLsizei width, int height, GLenum format, GLenum type, const GLvoid *pixels)
func DrawPixels(width int, height int, format, typ GLenum, pixels interface{}) {
	C.glDrawPixels(C.GLsizei(width), C.GLsizei(height), C.GLenum(format),
		C.GLenum(typ), ptr(pixels))
}

//void glEdgeFlag (bool flag)
func EdgeFlag(flag bool) {
	C.glEdgeFlag(glBool(flag))
}

//void glEdgeFlagPointer (GLsizei stride, const GLvoid *pointer)
func EdgeFlagPointer(stride int, pointer unsafe.Pointer) {
	C.glEdgeFlagPointer(C.GLsizei(stride), pointer)
}

//void glEdgeFlagv (const bool *flag)
func EdgeFlagv(flag *bool) {
	//TODO
	panic("unimplemented")
	//C.glEdgeFlagv((*C.GLboolean)(flag))
}

//void glEnable (GLenum cap)
func Enable(cap GLenum) {
	C.glEnable(C.GLenum(cap))
}

//void glEnableClientState (GLenum array)
func EnableClientState(array GLenum) {
	C.glEnableClientState(C.GLenum(array))
}

//void glEnd (void)
func End() {
	C.glEnd()
}

//void glEndList (void)
func EndList() {
	C.glEndList()
}

//void glEvalCoord1d (float64 u)
func EvalCoord1d(u float64) {
	C.glEvalCoord1d(C.GLdouble(u))
}

//void glEvalCoord1dv (const float64 *u)
func EvalCoord1dv(u *float64) {
	C.glEvalCoord1dv((*C.GLdouble)(u))
}

//void glEvalCoord1f (float32 u)
func EvalCoord1f(u float32) {
	C.glEvalCoord1f(C.GLfloat(u))
}

//void glEvalCoord1fv (const float *u)
func EvalCoord1fv(u []float32) {
	if len(u) != 1 {
		panic("Invalid coordinate length")
	}
	C.glEvalCoord1fv((*C.GLfloat)(&u[0]))
}

//void glEvalCoord2d (float64 u, float64 v)
func EvalCoord2d(u float64, v float64) {
	C.glEvalCoord2d(C.GLdouble(u), C.GLdouble(v))
}

//void glEvalCoord2dv (const float64 *u)
func EvalCoord2dv(u *float64) {
	C.glEvalCoord2dv((*C.GLdouble)(u))
}

//void glEvalCoord2f (float32 u, float32 v)
func EvalCoord2f(u float32, v float32) {
	C.glEvalCoord2f(C.GLfloat(u), C.GLfloat(v))
}

//void glEvalCoord2fv (const float *u)
func EvalCoord2fv(u []float32) {
	if len(u) != 2 {
		panic("Invalid coordinate length")
	}
	C.glEvalCoord2fv((*C.GLfloat)(&u[0]))
}

//void glEvalMesh1 (GLenum mode, int i1, int i2)
func EvalMesh1(mode GLenum, i1 int, i2 int) {
	C.glEvalMesh1(C.GLenum(mode), C.GLint(i1), C.GLint(i2))
}

//void glEvalMesh2 (GLenum mode, int i1, int i2, int j1, int j2)
func EvalMesh2(mode GLenum, i1 int, i2 int, j1 int, j2 int) {
	C.glEvalMesh2(C.GLenum(mode), C.GLint(i1), C.GLint(i2), C.GLint(j1), C.GLint(j2))
}

//void glEvalPoint1 (int i)
func EvalPoint1(i int) {
	C.glEvalPoint1(C.GLint(i))
}

//void glEvalPoint2 (int i, int j)
func EvalPoint2(i int, j int) {
	C.glEvalPoint2(C.GLint(i), C.GLint(j))
}

//void glFeedbackBuffer (GLsizei size, GLenum type, float32 *buffer)
func FeedbackBuffer(size int, type_ GLenum, buffer *float32) {
	C.glFeedbackBuffer(C.GLsizei(size), C.GLenum(type_), (*C.GLfloat)(buffer))
}

//void glFinish (void)
func Finish() {
	C.glFinish()
}

//void glFlush (void)
func Flush() {
	C.glFlush()
}

//void glFogf (GLenum pname, float32 param)
func Fogf(pname GLenum, param float32) {
	C.glFogf(C.GLenum(pname), C.GLfloat(param))
}

//void glFogfv (GLenum pname, const float *params)
func Fogfv(pname GLenum, params []float32) {
	if len(params) == 0 {
		panic("Invalid params length")
	}
	C.glFogfv(C.GLenum(pname), (*C.GLfloat)(&params[0]))
}

//void glFogi (GLenum pname, int param)
func Fogi(pname GLenum, param int) {
	C.glFogi(C.GLenum(pname), C.GLint(param))
}

//void glFogiv (GLenum pname, const int *params)
func Fogiv(pname GLenum, params []int32) {
	if len(params) == 0 {
		panic("Invalid params length")
	}
	C.glFogiv(C.GLenum(pname), (*C.GLint)(&params[0]))
}

//void glFrontFace (GLenum mode)
func FrontFace(mode GLenum) {
	C.glFrontFace(C.GLenum(mode))
}

//void glFrustum (float64 left, float64 right, float64 bottom, float64 top, float64 zNear, float64 zFar)
func Frustum(left float64, right float64, bottom float64, top float64, zNear float64, zFar float64) {
	C.glFrustum(C.GLdouble(left), C.GLdouble(right), C.GLdouble(bottom), C.GLdouble(top), C.GLdouble(zNear), C.GLdouble(zFar))
}

//uint glGenLists (GLsizei range)
func GenLists(range_ int) uint {
	return uint(C.glGenLists(C.GLsizei(range_)))
}

//void glGetBooleanv (GLenum pname, bool *params)
func GetBooleanv(pname GLenum, params []bool) {
	//TODO
	panic("unimplemented")
	//C.glGetBooleanv(C.GLenum(pname), (*C.GLboolean)(&params[0]))
}

//void glGetClipPlane (GLenum plane, float64 *equation)
func GetClipPlane(plane GLenum, equation *float64) {
	C.glGetClipPlane(C.GLenum(plane), (*C.GLdouble)(equation))
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

//void glGetLightfv (GLenum light, GLenum pname, float *params)
func GetLightfv(light GLenum, pname GLenum, params []float32) {
	if len(params) == 0 {
		panic("Invalid params length")
	}
	C.glGetLightfv(C.GLenum(light), C.GLenum(pname), (*C.GLfloat)(&params[0]))
}

//void glGetLightiv (GLenum light, GLenum pname, int *params)
func GetLightiv(light GLenum, pname GLenum, params []int32) {
	if len(params) == 0 {
		panic("Invalid params length")
	}
	C.glGetLightiv(C.GLenum(light), C.GLenum(pname), (*C.GLint)(&params[0]))
}

//void glGetMapdv (GLenum target, GLenum query, float64 *v)
func GetMapdv(target GLenum, query GLenum, v []float64) {
	if len(v) == 0 {
		panic("Invalid slice length")
	}
	C.glGetMapdv(C.GLenum(target), C.GLenum(query), (*C.GLdouble)(&v[0]))
}

//void glGetMapfv (GLenum target, GLenum query, float *v)
func GetMapfv(target GLenum, query GLenum, v []float32) {
	if len(v) == 0 {
		panic("Invalid slice length")
	}
	C.glGetMapfv(C.GLenum(target), C.GLenum(query), (*C.GLfloat)(&v[0]))
}

//void glGetMapiv (GLenum target, GLenum query, int *v)
func GetMapiv(target GLenum, query GLenum, v []int32) {
	if len(v) == 0 {
		panic("Invalid slice length")
	}
	C.glGetMapiv(C.GLenum(target), C.GLenum(query), (*C.GLint)(&v[0]))
}

//void glGetMaterialfv (GLenum face, GLenum pname, float *params)
func GetMaterialfv(face GLenum, pname GLenum, params []float32) {
	if len(params) == 0 {
		panic("Invalid params length")
	}
	C.glGetMaterialfv(C.GLenum(face), C.GLenum(pname), (*C.GLfloat)(&params[0]))
}

//void glGetMaterialiv (GLenum face, GLenum pname, int *params)
func GetMaterialiv(face GLenum, pname GLenum, params []int32) {
	if len(params) == 0 {
		panic("Invalid params length")
	}
	C.glGetMaterialiv(C.GLenum(face), C.GLenum(pname), (*C.GLint)(&params[0]))
}

//void glGetPixelMapfv (GLenum map, float *values)
func GetPixelMapfv(map_ GLenum, values []float32) {
	if len(values) == 0 {
		panic("Invalid values length")
	}
	C.glGetPixelMapfv(C.GLenum(map_), (*C.GLfloat)(&values[0]))
}

//void glGetPixelMapuiv (GLenum map, uint *values)
func GetPixelMapuiv(map_ GLenum, values *uint32) {
	C.glGetPixelMapuiv(C.GLenum(map_), (*C.GLuint)(values))
}

//void glGetPixelMapusv (GLenum map, uint16 *values)
func GetPixelMapusv(map_ GLenum, values *uint16) {
	C.glGetPixelMapusv(C.GLenum(map_), (*C.GLushort)(values))
}

//void glGetPointerv (GLenum pname, GLvoid* *params)
func GetPointerv(pname GLenum, params []unsafe.Pointer) {
	if len(params) == 0 {
		panic("Invalid params length")
	}
	C.glGetPointerv(C.GLenum(pname), &params[0])
}

//void glGetPolygonStipple (uint8 *mask)
func GetPolygonStipple(mask *uint8) {
	C.glGetPolygonStipple((*C.GLubyte)(mask))
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

//void glIndexMask (uint mask)
func IndexMask(mask uint) {
	C.glIndexMask(C.GLuint(mask))
}

//void glIndexPointer (GLenum type, int stride, const GLvoid *pointer)
func IndexPointer(typ GLenum, stride int, pointer interface{}) {
	C.glIndexPointer(C.GLenum(typ), C.GLsizei(stride), ptr(pointer))
}

//void glIndexd (float64 c)
func Indexd(c float64) {
	C.glIndexd(C.GLdouble(c))
}

//void glIndexdv (const float64 *c)
func Indexdv(c *float64) {
	C.glIndexdv((*C.GLdouble)(c))
}

//void glIndexf (float32 c)
func Indexf(c float32) {
	C.glIndexf(C.GLfloat(c))
}

//void glIndexfv (const float32 *c)
func Indexfv(c *float32) {
	C.glIndexfv((*C.GLfloat)(c))
}

//void glIndexi (int c)
func Indexi(c int) {
	C.glIndexi(C.GLint(c))
}

//void glIndexiv (const int *c)
func Indexiv(c *int32) {
	C.glIndexiv((*C.GLint)(c))
}

//void glIndexs (int16 c)
func Indexs(c int16) {
	C.glIndexs(C.GLshort(c))
}

//void glIndexsv (const int16 *c)
func Indexsv(c *int16) {
	C.glIndexsv((*C.GLshort)(c))
}

//void glIndexub (uint8 c)
func Indexub(c uint8) {
	C.glIndexub(C.GLubyte(c))
}

//void glIndexubv (const uint8 *c)
func Indexubv(c *uint8) {
	C.glIndexubv((*C.GLubyte)(c))
}

//void glInitNames (void)
func InitNames() {
	C.glInitNames()
}

//void glInterleavedArrays (GLenum format, int stride, const GLvoid *pointer)
func InterleavedArrays(format GLenum, stride int, pointer unsafe.Pointer) {
	C.glInterleavedArrays(C.GLenum(format), C.GLsizei(stride), pointer)
}

//bool glIsEnabled (GLenum cap)
func IsEnabled(cap GLenum) bool {
	return goBool(C.glIsEnabled(C.GLenum(cap)))
}

//bool glIsList (uint list)
func IsList(list uint) bool {
	return goBool(C.glIsList(C.GLuint(list)))
}

//void glLightModelf (GLenum pname, float32 param)
func LightModelf(pname GLenum, param float32) {
	C.glLightModelf(C.GLenum(pname), C.GLfloat(param))
}

//void glLightModelfv (GLenum pname, const float *params)
func LightModelfv(pname GLenum, params []float32) {
	if len(params) == 0 {
		panic("Invalid params length")
	}
	C.glLightModelfv(C.GLenum(pname), (*C.GLfloat)(&params[0]))
}

//void glLightModeli (GLenum pname, int param)
func LightModeli(pname GLenum, param int) {
	C.glLightModeli(C.GLenum(pname), C.GLint(param))
}

//void glLightModeliv (GLenum pname, const int *params)
func LightModeliv(pname GLenum, params []int32) {
	if len(params) == 0 {
		panic("Invalid params length")
	}
	C.glLightModeliv(C.GLenum(pname), (*C.GLint)(&params[0]))
}

//void glLightf (GLenum light, GLenum pname, float32 param)
func Lightf(light GLenum, pname GLenum, param float32) {
	C.glLightf(C.GLenum(light), C.GLenum(pname), C.GLfloat(param))
}

//void glLightfv (GLenum light, GLenum pname, const float *params)
func Lightfv(light GLenum, pname GLenum, params []float32) {
	if len(params) == 0 {
		panic("Invalid params length")
	}
	C.glLightfv(C.GLenum(light), C.GLenum(pname), (*C.GLfloat)(&params[0]))
}

//void glLighti (GLenum light, GLenum pname, int param)
func Lighti(light GLenum, pname GLenum, param int) {
	C.glLighti(C.GLenum(light), C.GLenum(pname), C.GLint(param))
}

//void glLightiv (GLenum light, GLenum pname, const int *params)
func Lightiv(light GLenum, pname GLenum, params []int32) {
	if len(params) == 0 {
		panic("Invalid params length")
	}
	C.glLightiv(C.GLenum(light), C.GLenum(pname), (*C.GLint)(&params[0]))
}

//void glLineStipple (int factor, uint16 pattern)
func LineStipple(factor int, pattern uint16) {
	C.glLineStipple(C.GLint(factor), C.GLushort(pattern))
}

//void glLineWidth (float32 width)
func LineWidth(width float32) {
	C.glLineWidth(C.GLfloat(width))
}

//void glListBase (uint base)
func ListBase(base uint) {
	C.glListBase(C.GLuint(base))
}

//void glLoadIdentity (void)
func LoadIdentity() {
	C.glLoadIdentity()
}

//void glLoadMatrixd (const float64 *m)
func LoadMatrixd(m *float64) {
	C.glLoadMatrixd((*C.GLdouble)(m))
}

//void glLoadMatrixf (const float32 *m)
func LoadMatrixf(m *float32) {
	C.glLoadMatrixf((*C.GLfloat)(m))
}

//void glLoadName (uint name)
func LoadName(name uint) {
	C.glLoadName(C.GLuint(name))
}

//void glLogicOp (GLenum opcode)
func LogicOp(opcode GLenum) {
	C.glLogicOp(C.GLenum(opcode))
}

//void glMap1d (GLenum target, float64 u1, float64 u2, int stride, int order, const float64 *points)
func Map1d(target GLenum, u1 float64, u2 float64, stride int, order int, points *float64) {
	C.glMap1d(C.GLenum(target), C.GLdouble(u1), C.GLdouble(u2), C.GLint(stride), C.GLint(order), (*C.GLdouble)(points))
}

//void glMap1f (GLenum target, float32 u1, float32 u2, int stride, int order, const float32 *points)
func Map1f(target GLenum, u1 float32, u2 float32, stride int, order int, points *float32) {
	C.glMap1f(C.GLenum(target), C.GLfloat(u1), C.GLfloat(u2), C.GLint(stride), C.GLint(order), (*C.GLfloat)(points))
}

//void glMap2d (GLenum target, float64 u1, float64 u2, int ustride, int uorder, float64 v1, float64 v2, int vstride, int vorder, const float64 *points)
func Map2d(target GLenum, u1 float64, u2 float64, ustride int, uorder int, v1 float64, v2 float64, vstride int, vorder int, points *float64) {
	C.glMap2d(C.GLenum(target), C.GLdouble(u1), C.GLdouble(u2), C.GLint(ustride), C.GLint(uorder), C.GLdouble(v1), C.GLdouble(v2), C.GLint(vstride), C.GLint(vorder), (*C.GLdouble)(points))
}

//void glMap2f (GLenum target, float32 u1, float32 u2, int ustride, int uorder, float32 v1, float32 v2, int vstride, int vorder, const float32 *points)
func Map2f(target GLenum, u1 float32, u2 float32, ustride int, uorder int, v1 float32, v2 float32, vstride int, vorder int, points *float32) {
	C.glMap2f(C.GLenum(target), C.GLfloat(u1), C.GLfloat(u2), C.GLint(ustride), C.GLint(uorder), C.GLfloat(v1), C.GLfloat(v2), C.GLint(vstride), C.GLint(vorder), (*C.GLfloat)(points))
}

//void glMapGrid1d (int un, float64 u1, float64 u2)
func MapGrid1d(un int, u1 float64, u2 float64) {
	C.glMapGrid1d(C.GLint(un), C.GLdouble(u1), C.GLdouble(u2))
}

//void glMapGrid1f (int un, float32 u1, float32 u2)
func MapGrid1f(un int, u1 float32, u2 float32) {
	C.glMapGrid1f(C.GLint(un), C.GLfloat(u1), C.GLfloat(u2))
}

//void glMapGrid2d (int un, float64 u1, float64 u2, int vn, float64 v1, float64 v2)
func MapGrid2d(un int, u1 float64, u2 float64, vn int, v1 float64, v2 float64) {
	C.glMapGrid2d(C.GLint(un), C.GLdouble(u1), C.GLdouble(u2), C.GLint(vn), C.GLdouble(v1), C.GLdouble(v2))
}

//void glMapGrid2f (int un, float32 u1, float32 u2, int vn, float32 v1, float32 v2)
func MapGrid2f(un int, u1 float32, u2 float32, vn int, v1 float32, v2 float32) {
	C.glMapGrid2f(C.GLint(un), C.GLfloat(u1), C.GLfloat(u2), C.GLint(vn), C.GLfloat(v1), C.GLfloat(v2))
}

//void glMaterialf (GLenum face, GLenum pname, float32 param)
func Materialf(face GLenum, pname GLenum, param float32) {
	C.glMaterialf(C.GLenum(face), C.GLenum(pname), C.GLfloat(param))
}

//void glMaterialfv (GLenum face, GLenum pname, const float *params)
func Materialfv(face GLenum, pname GLenum, params []float32) {
	if len(params) == 0 {
		panic("Invalid params length")
	}
	C.glMaterialfv(C.GLenum(face), C.GLenum(pname), (*C.GLfloat)(&params[0]))
}

//void glMateriali (GLenum face, GLenum pname, int param)
func Materiali(face GLenum, pname GLenum, param int) {
	C.glMateriali(C.GLenum(face), C.GLenum(pname), C.GLint(param))
}

//void glMaterialiv (GLenum face, GLenum pname, const int *params)
func Materialiv(face GLenum, pname GLenum, params []int32) {
	if len(params) == 0 {
		panic("Invalid params length")
	}
	C.glMaterialiv(C.GLenum(face), C.GLenum(pname), (*C.GLint)(&params[0]))
}

//void glMatrixMode (GLenum mode)
func MatrixMode(mode GLenum) {
	C.glMatrixMode(C.GLenum(mode))
}

//void glMultMatrixd (const float64 *m)
func MultMatrixd(m []float64) {
	if len(m) != 16 {
		panic("Invalid matrix size")
	}
	C.glMultMatrixd((*C.GLdouble)(unsafe.Pointer(&m[0])))
}

//void glMultMatrixf (const float32 *m)
func MultMatrixf(m []float32) {
	if len(m) != 16 {
		panic("Invalid matrix size")
	}
	C.glMultMatrixf((*C.GLfloat)(unsafe.Pointer(&m[0])))
}

//void glNewList (uint list, GLenum mode)
func NewList(list uint, mode GLenum) {
	C.glNewList(C.GLuint(list), C.GLenum(mode))
}

//void glNormal3b (int8 nx, int8 ny, int8 nz)
func Normal3b(nx int8, ny int8, nz int8) {
	C.glNormal3b(C.GLbyte(nx), C.GLbyte(ny), C.GLbyte(nz))
}

//void glNormal3bv (const int8 *v)
func Normal3bv(v []int8) {
	C.glNormal3bv((*C.GLbyte)(&v[0]))
}

//void glNormal3d (float64 nx, float64 ny, float64 nz)
func Normal3d(nx float64, ny float64, nz float64) {
	C.glNormal3d(C.GLdouble(nx), C.GLdouble(ny), C.GLdouble(nz))
}

//void glNormal3dv (const float64 *v)
func Normal3dv(v []float64) {
	if len(v) != 3 {
		panic("Invalid normal size")
	}
	C.glNormal3dv((*C.GLdouble)(&v[0]))
}

//void glNormal3f (float32 nx, float32 ny, float32 nz)
func Normal3f(nx float32, ny float32, nz float32) {
	C.glNormal3f(C.GLfloat(nx), C.GLfloat(ny), C.GLfloat(nz))
}

//void glNormal3fv (const float *v)
func Normal3fv(v []float32) {
	if len(v) != 3 {
		panic("Invalid normal size")
	}
	C.glNormal3fv((*C.GLfloat)(&v[0]))
}

//void glNormal3i (int nx, int ny, int nz)
func Normal3i(nx int, ny int, nz int) {
	C.glNormal3i(C.GLint(nx), C.GLint(ny), C.GLint(nz))
}

//void glNormal3iv (const int *v)
func Normal3iv(v []int32) {
	if len(v) != 3 {
		panic("Invalid normal size")
	}
	C.glNormal3iv((*C.GLint)(&v[0]))
}

//void glNormal3s (int16 nx, int16 ny, int16 nz)
func Normal3s(nx int16, ny int16, nz int16) {
	C.glNormal3s(C.GLshort(nx), C.GLshort(ny), C.GLshort(nz))
}

//void glNormal3sv (const int16 *v)
func Normal3sv(v []int16) {
	if len(v) != 3 {
		panic("Invalid normal size")
	}
	C.glNormal3sv((*C.GLshort)(&v[0]))
}

//void glNormalPointer (GLenum type, int stride, const GLvoid *pointer)
func NormalPointer(typ GLenum, stride int, pointer interface{}) {
	C.glNormalPointer(C.GLenum(typ), C.GLsizei(stride), ptr(pointer))
}

//void glOrtho (float64 left, float64 right, float64 bottom, float64 top, float64 zNear, float64 zFar)
func Ortho(left float64, right float64, bottom float64, top float64, zNear float64, zFar float64) {
	C.glOrtho(C.GLdouble(left), C.GLdouble(right), C.GLdouble(bottom), C.GLdouble(top), C.GLdouble(zNear), C.GLdouble(zFar))
}

//void glPassThrough (float32 token)
func PassThrough(token float32) {
	C.glPassThrough(C.GLfloat(token))
}

//void glPixelStoref (GLenum pname, float param)
func PixelStoref(pname GLenum, param float32) {
	C.glPixelStoref(C.GLenum(pname), C.GLfloat(param))
}

//void glPixelStorei (GLenum pname, int param)
func PixelStorei(pname GLenum, param int) {
	C.glPixelStorei(C.GLenum(pname), C.GLint(param))
}

//void glPixelTransferf (GLenum pname, float32 param)
func PixelTransferf(pname GLenum, param float32) {
	C.glPixelTransferf(C.GLenum(pname), C.GLfloat(param))
}

//void glPixelTransferi (GLenum pname, int param)
func PixelTransferi(pname GLenum, param int) {
	C.glPixelTransferi(C.GLenum(pname), C.GLint(param))
}

//void glPixelZoom (float32 xfactor, float32 yfactor)
func PixelZoom(xfactor float32, yfactor float32) {
	C.glPixelZoom(C.GLfloat(xfactor), C.GLfloat(yfactor))
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

//void glPolygonStipple (const uint8 *mask)
func PolygonStipple(mask *uint8) {
	C.glPolygonStipple((*C.GLubyte)(mask))
}

//void glPopAttrib (void)
func PopAttrib() {
	C.glPopAttrib()
}

//void glPopClientAttrib (void)
func PopClientAttrib() {
	C.glPopClientAttrib()
}

//void glPopMatrix (void)
func PopMatrix() {
	C.glPopMatrix()
}

//void glPopName (void)
func PopName() {
	C.glPopName()
}

//void glPushAttrib (GLbitfield mask)
func PushAttrib(mask GLbitfield) {
	C.glPushAttrib(C.GLbitfield(mask))
}

//void glPushClientAttrib (GLbitfield mask)
func PushClientAttrib(mask GLbitfield) {
	C.glPushClientAttrib(C.GLbitfield(mask))
}

//void glPushMatrix (void)
func PushMatrix() {
	C.glPushMatrix()
}

//void glPushName (uint name)
func PushName(name uint) {
	C.glPushName(C.GLuint(name))
}

//void glRasterPos2d (float64 x, float64 y)
func RasterPos2d(x float64, y float64) {
	C.glRasterPos2d(C.GLdouble(x), C.GLdouble(y))
}

//void glRasterPos2dv (const float64 *v)
func RasterPos2dv(v []float64) {
	if len(v) != 2 {
		panic("Invalid point size")
	}
	C.glRasterPos2dv((*C.GLdouble)(&v[0]))
}

//void glRasterPos2f (float32 x, float32 y)
func RasterPos2f(x float32, y float32) {
	C.glRasterPos2f(C.GLfloat(x), C.GLfloat(y))
}

//void glRasterPos2fv (const float *v)
func RasterPos2fv(v []float32) {
	if len(v) != 2 {
		panic("Invalid point size")
	}
	C.glRasterPos2fv((*C.GLfloat)(&v[0]))
}

//void glRasterPos2i (int x, int y)
func RasterPos2i(x int, y int) {
	C.glRasterPos2i(C.GLint(x), C.GLint(y))
}

//void glRasterPos2iv (const int *v)
func RasterPos2iv(v []int32) {
	if len(v) != 2 {
		panic("Invalid point size")
	}
	C.glRasterPos2iv((*C.GLint)(&v[0]))
}

//void glRasterPos2s (int16 x, int16 y)
func RasterPos2s(x int16, y int16) {
	C.glRasterPos2s(C.GLshort(x), C.GLshort(y))
}

//void glRasterPos2sv (const int16 *v)
func RasterPos2sv(v []int16) {
	if len(v) != 2 {
		panic("Invalid point size")
	}
	C.glRasterPos2sv((*C.GLshort)(&v[0]))
}

//void glRasterPos3d (float64 x, float64 y, float64 z)
func RasterPos3d(x float64, y float64, z float64) {
	C.glRasterPos3d(C.GLdouble(x), C.GLdouble(y), C.GLdouble(z))
}

//void glRasterPos3dv (const float64 *v)
func RasterPos3dv(v []float64) {
	if len(v) != 3 {
		panic("Invalid point size")
	}
	C.glRasterPos3dv((*C.GLdouble)(&v[0]))
}

//void glRasterPos3f (float32 x, float32 y, float32 z)
func RasterPos3f(x float32, y float32, z float32) {
	C.glRasterPos3f(C.GLfloat(x), C.GLfloat(y), C.GLfloat(z))
}

//void glRasterPos3fv (const float *v)
func RasterPos3fv(v []float32) {
	if len(v) != 3 {
		panic("Invalid point size")
	}
	C.glRasterPos3fv((*C.GLfloat)(&v[0]))
}

//void glRasterPos3i (int x, int y, int z)
func RasterPos3i(x int, y int, z int) {
	C.glRasterPos3i(C.GLint(x), C.GLint(y), C.GLint(z))
}

//void glRasterPos3iv (const int *v)
func RasterPos3iv(v []int32) {
	if len(v) != 3 {
		panic("Invalid point size")
	}
	C.glRasterPos3iv((*C.GLint)(&v[0]))
}

//void glRasterPos3s (int16 x, int16 y, int16 z)
func RasterPos3s(x int16, y int16, z int16) {
	C.glRasterPos3s(C.GLshort(x), C.GLshort(y), C.GLshort(z))
}

//void glRasterPos3sv (const int16 *v)
func RasterPos3sv(v []int16) {
	if len(v) != 3 {
		panic("Invalid point size")
	}
	C.glRasterPos3sv((*C.GLshort)(&v[0]))
}

//void glRasterPos4d (float64 x, float64 y, float64 z, float64 w)
func RasterPos4d(x float64, y float64, z float64, w float64) {
	C.glRasterPos4d(C.GLdouble(x), C.GLdouble(y), C.GLdouble(z), C.GLdouble(w))
}

//void glRasterPos4dv (const float64 *v)
func RasterPos4dv(v []float64) {
	if len(v) != 3 {
		panic("Invalid point size")
	}
	C.glRasterPos4dv((*C.GLdouble)(&v[0]))
}

//void glRasterPos4f (float32 x, float32 y, float32 z, float32 w)
func RasterPos4f(x float32, y float32, z float32, w float32) {
	C.glRasterPos4f(C.GLfloat(x), C.GLfloat(y), C.GLfloat(z), C.GLfloat(w))
}

//void glRasterPos4fv (const float *v)
func RasterPos4fv(v []float32) {
	if len(v) != 4 {
		panic("Invalid point size")
	}
	C.glRasterPos4fv((*C.GLfloat)(&v[0]))
}

//void glRasterPos4i (int x, int y, int z, int w)
func RasterPos4i(x int, y int, z int, w int) {
	C.glRasterPos4i(C.GLint(x), C.GLint(y), C.GLint(z), C.GLint(w))
}

//void glRasterPos4iv (const int *v)
func RasterPos4iv(v []int32) {
	if len(v) != 4 {
		panic("Invalid point size")
	}
	C.glRasterPos4iv((*C.GLint)(&v[0]))
}

//void glRasterPos4s (int16 x, int16 y, int16 z, int16 w)
func RasterPos4s(x int16, y int16, z int16, w int16) {
	C.glRasterPos4s(C.GLshort(x), C.GLshort(y), C.GLshort(z), C.GLshort(w))
}

//void glRasterPos4sv (const int16 *v)
func RasterPos4sv(v []int16) {
	if len(v) != 4 {
		panic("Invalid point size")
	}
	C.glRasterPos4sv((*C.GLshort)(&v[0]))
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

//void glRectd (float64 x1, float64 y1, float64 x2, float64 y2)
func Rectd(x1 float64, y1 float64, x2 float64, y2 float64) {
	C.glRectd(C.GLdouble(x1), C.GLdouble(y1), C.GLdouble(x2), C.GLdouble(y2))
}

//void glRectdv (const float64 *v1, const float64 *v2)
func Rectdv(v1 *float64, v2 *float64) {
	C.glRectdv((*C.GLdouble)(v1), (*C.GLdouble)(v2))
}

//void glRectf (float32 x1, float32 y1, float32 x2, float32 y2)
func Rectf(x1 float32, y1 float32, x2 float32, y2 float32) {
	C.glRectf(C.GLfloat(x1), C.GLfloat(y1), C.GLfloat(x2), C.GLfloat(y2))
}

//void glRectfv (const float *v1, const float *v2)
func Rectfv(v1 []float32, v2 []float32) {
	if len(v1) != 2 || len(v2) != 2 {
		panic("Invalid vertex size")
	}
	C.glRectfv((*C.GLfloat)(&v1[0]), (*C.GLfloat)(&v2[0]))
}

//void glRecti (int x1, int y1, int x2, int y2)
func Recti(x1 int, y1 int, x2 int, y2 int) {
	C.glRecti(C.GLint(x1), C.GLint(y1), C.GLint(x2), C.GLint(y2))
}

//void glRectiv (const int *v1, const int *v2)
func Rectiv(v1 *int32, v2 *int32) {
	C.glRectiv((*C.GLint)(v1), (*C.GLint)(v2))
}

//void glRects (int16 x1, int16 y1, int16 x2, int16 y2)
func Rects(x1 int16, y1 int16, x2 int16, y2 int16) {
	C.glRects(C.GLshort(x1), C.GLshort(y1), C.GLshort(x2), C.GLshort(y2))
}

//void glRectsv (const int16 *v1, const int16 *v2)
func Rectsv(v1 *int16, v2 *int16) {
	C.glRectsv((*C.GLshort)(v1), (*C.GLshort)(v2))
}

//int glRenderMode (GLenum mode)
func RenderMode(mode GLenum) int {
	return int(C.glRenderMode(C.GLenum(mode)))
}

//void glRotated (float64 angle, float64 x, float64 y, float64 z)
func Rotated(angle float64, x float64, y float64, z float64) {
	C.glRotated(C.GLdouble(angle), C.GLdouble(x), C.GLdouble(y), C.GLdouble(z))
}

//void glRotatef (float32 angle, float32 x, float32 y, float32 z)
func Rotatef(angle float32, x float32, y float32, z float32) {
	C.glRotatef(C.GLfloat(angle), C.GLfloat(x), C.GLfloat(y), C.GLfloat(z))
}

//void glScaled (float64 x, float64 y, float64 z)
func Scaled(x float64, y float64, z float64) {
	C.glScaled(C.GLdouble(x), C.GLdouble(y), C.GLdouble(z))
}

//void glScalef (float32 x, float32 y, float32 z)
func Scalef(x float32, y float32, z float32) {
	C.glScalef(C.GLfloat(x), C.GLfloat(y), C.GLfloat(z))
}

//void glScissor (int x, int y, int width, int height)
func Scissor(x int, y int, width int, height int) {
	C.glScissor(C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height))
}

//void glSelectBuffer (GLsizei size, uint *buffer)
func SelectBuffer(buffer []uint32) {
	if len(buffer) < 1 {
		panic("Invalid buffer size")
	}
	C.glSelectBuffer(C.GLsizei(len(buffer)), (*C.GLuint)(&buffer[0]))
}

//void glShadeModel (GLenum mode)
func ShadeModel(mode GLenum) {
	C.glShadeModel(C.GLenum(mode))
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

//void glTexCoord1d (float64 s)
func TexCoord1d(s float64) {
	C.glTexCoord1d(C.GLdouble(s))
}

//void glTexCoord1dv (const float64 *v)
func TexCoord1dv(v []float64) {
	if len(v) != 1 {
		panic("Invalid vector size")
	}
	C.glTexCoord1dv((*C.GLdouble)(&v[0]))
}

//void glTexCoord1f (float32 s)
func TexCoord1f(s float32) {
	C.glTexCoord1f(C.GLfloat(s))
}

//void glTexCoord1fv (const float *v)
func TexCoord1fv(v []float32) {
	if len(v) != 1 {
		panic("Invalid vector size")
	}
	C.glTexCoord1fv((*C.GLfloat)(&v[0]))
}

//void glTexCoord1i (int s)
func TexCoord1i(s int) {
	C.glTexCoord1i(C.GLint(s))
}

//void glTexCoord1iv (const int *v)
func TexCoord1iv(v []int32) {
	if len(v) != 1 {
		panic("Invalid vector size")
	}
	C.glTexCoord1iv((*C.GLint)(&v[0]))
}

//void glTexCoord1s (int16 s)
func TexCoord1s(s int16) {
	C.glTexCoord1s(C.GLshort(s))
}

//void glTexCoord1sv (const int16 *v)
func TexCoord1sv(v []int16) {
	if len(v) != 1 {
		panic("Invalid vector size")
	}
	C.glTexCoord1sv((*C.GLshort)(&v[0]))
}

//void glTexCoord2d (float64 s, float64 t)
func TexCoord2d(s float64, t float64) {
	C.glTexCoord2d(C.GLdouble(s), C.GLdouble(t))
}

//void glTexCoord2dv (const float64 *v)
func TexCoord2dv(v []float64) {
	if len(v) != 2 {
		panic("Invalid vector size")
	}
	C.glTexCoord2dv((*C.GLdouble)(&v[0]))
}

//void glTexCoord2f (float32 s, float32 t)
func TexCoord2f(s float32, t float32) {
	C.glTexCoord2f(C.GLfloat(s), C.GLfloat(t))
}

//void glTexCoord2fv (const float *v)
func TexCoord2fv(v []float32) {
	if len(v) != 2 {
		panic("Invalid vector size")
	}
	C.glTexCoord2fv((*C.GLfloat)(&v[0]))
}

//void glTexCoord2i (int s, int t)
func TexCoord2i(s int, t int) {
	C.glTexCoord2i(C.GLint(s), C.GLint(t))
}

//void glTexCoord2iv (const int *v)
func TexCoord2iv(v []int32) {
	if len(v) != 2 {
		panic("Invalid vector size")
	}
	C.glTexCoord2iv((*C.GLint)(&v[0]))
}

//void glTexCoord2s (int16 s, int16 t)
func TexCoord2s(s int16, t int16) {
	C.glTexCoord2s(C.GLshort(s), C.GLshort(t))
}

//void glTexCoord2sv (const int16 *v)
func TexCoord2sv(v []int16) {
	if len(v) != 2 {
		panic("Invalid vector size")
	}
	C.glTexCoord2sv((*C.GLshort)(&v[0]))
}

//void glTexCoord3d (float64 s, float64 t, float64 r)
func TexCoord3d(s float64, t float64, r float64) {
	C.glTexCoord3d(C.GLdouble(s), C.GLdouble(t), C.GLdouble(r))
}

//void glTexCoord3dv (const float64 *v)
func TexCoord3dv(v []float64) {
	if len(v) != 3 {
		panic("Invalid vector size")
	}
	C.glTexCoord3dv((*C.GLdouble)(&v[0]))
}

//void glTexCoord3f (float32 s, float32 t, float32 r)
func TexCoord3f(s float32, t float32, r float32) {
	C.glTexCoord3f(C.GLfloat(s), C.GLfloat(t), C.GLfloat(r))
}

//void glTexCoord3fv (const float *v)
func TexCoord3fv(v []float32) {
	if len(v) != 3 {
		panic("Invalid vector size")
	}
	C.glTexCoord3fv((*C.GLfloat)(&v[0]))
}

//void glTexCoord3i (int s, int t, int r)
func TexCoord3i(s int, t int, r int) {
	C.glTexCoord3i(C.GLint(s), C.GLint(t), C.GLint(r))
}

//void glTexCoord3iv (const int *v)
func TexCoord3iv(v []int32) {
	if len(v) != 3 {
		panic("Invalid vector size")
	}
	C.glTexCoord3iv((*C.GLint)(&v[0]))
}

//void glTexCoord3s (int16 s, int16 t, int16 r)
func TexCoord3s(s int16, t int16, r int16) {
	C.glTexCoord3s(C.GLshort(s), C.GLshort(t), C.GLshort(r))
}

//void glTexCoord3sv (const int16 *v)
func TexCoord3sv(v []int16) {
	if len(v) != 3 {
		panic("Invalid vector size")
	}
	C.glTexCoord3sv((*C.GLshort)(&v[0]))
}

//void glTexCoord4d (float64 s, float64 t, float64 r, float64 q)
func TexCoord4d(s float64, t float64, r float64, q float64) {
	C.glTexCoord4d(C.GLdouble(s), C.GLdouble(t), C.GLdouble(r), C.GLdouble(q))
}

//void glTexCoord4dv (const float64 *v)
func TexCoord4dv(v []float64) {
	if len(v) != 4 {
		panic("Invalid vector size")
	}
	C.glTexCoord4dv((*C.GLdouble)(&v[0]))
}

//void glTexCoord4f (float32 s, float32 t, float32 r, float32 q)
func TexCoord4f(s float32, t float32, r float32, q float32) {
	C.glTexCoord4f(C.GLfloat(s), C.GLfloat(t), C.GLfloat(r), C.GLfloat(q))
}

//void glTexCoord4fv (const float *v)
func TexCoord4fv(v []float32) {
	if len(v) != 4 {
		panic("Invalid vector size")
	}
	C.glTexCoord4fv((*C.GLfloat)(&v[0]))
}

//void glTexCoord4i (int s, int t, int r, int q)
func TexCoord4i(s int, t int, r int, q int) {
	C.glTexCoord4i(C.GLint(s), C.GLint(t), C.GLint(r), C.GLint(q))
}

//void glTexCoord4iv (const int *v)
func TexCoord4iv(v []int32) {
	if len(v) != 4 {
		panic("Invalid vector size")
	}
	C.glTexCoord4iv((*C.GLint)(&v[0]))
}

//void glTexCoord4s (int16 s, int16 t, int16 r, int16 q)
func TexCoord4s(s int16, t int16, r int16, q int16) {
	C.glTexCoord4s(C.GLshort(s), C.GLshort(t), C.GLshort(r), C.GLshort(q))
}

//void glTexCoord4sv (const int16 *v)
func TexCoord4sv(v []int16) {
	if len(v) != 4 {
		panic("Invalid vector size")
	}
	C.glTexCoord4sv((*C.GLshort)(&v[0]))
}

//void glTexCoordPointer (int size, GLenum type, int stride, const GLvoid *pointer)
func TexCoordPointer(size int, typ GLenum, stride int, pointer interface{}) {
	C.glTexCoordPointer(C.GLint(size), C.GLenum(typ), C.GLsizei(stride),
		ptr(pointer))
}

//void glTranslated (float64 x, float64 y, float64 z)
func Translated(x float64, y float64, z float64) {
	C.glTranslated(C.GLdouble(x), C.GLdouble(y), C.GLdouble(z))
}

//void glTranslatef (float32 x, float32 y, float32 z)
func Translatef(x float32, y float32, z float32) {
	C.glTranslatef(C.GLfloat(x), C.GLfloat(y), C.GLfloat(z))
}

//void glVertex2d (float64 x, float64 y)
func Vertex2d(x float64, y float64) {
	C.glVertex2d(C.GLdouble(x), C.GLdouble(y))
}

//void glVertex2dv (const float64 *v)
func Vertex2dv(v []float64) {
	if len(v) != 2 {
		panic("Invalid vertex coord size")
	}
	C.glVertex2dv((*C.GLdouble)(&v[0]))
}

//void glVertex2f (float32 x, float32 y)
func Vertex2f(x float32, y float32) {
	C.glVertex2f(C.GLfloat(x), C.GLfloat(y))
}

//void glVertex2fv (const float *v)
func Vertex2fv(v []float32) {
	if len(v) != 2 {
		panic("Invalid vertex coord size")
	}
	C.glVertex2fv((*C.GLfloat)(&v[0]))
}

//void glVertex2i (int x, int y)
func Vertex2i(x int, y int) {
	C.glVertex2i(C.GLint(x), C.GLint(y))
}

//void glVertex2iv (const int *v)
func Vertex2iv(v []int32) {
	if len(v) != 2 {
		panic("Invalid vertex coord size")
	}
	C.glVertex2iv((*C.GLint)(&v[0]))
}

//void glVertex2s (int16 x, int16 y)
func Vertex2s(x int16, y int16) {
	C.glVertex2s(C.GLshort(x), C.GLshort(y))
}

//void glVertex2sv (const int16 *v)
func Vertex2sv(v []int16) {
	if len(v) != 2 {
		panic("Invalid vertex coord size")
	}
	C.glVertex2sv((*C.GLshort)(&v[0]))
}

//void glVertex3d (float64 x, float64 y, float64 z)
func Vertex3d(x float64, y float64, z float64) {
	C.glVertex3d(C.GLdouble(x), C.GLdouble(y), C.GLdouble(z))
}

//void glVertex3dv (const float64 *v)
func Vertex3dv(v []float64) {
	if len(v) != 3 {
		panic("Invalid vertex coord size")
	}
	C.glVertex3dv((*C.GLdouble)(&v[0]))
}

//void glVertex3f (float32 x, float32 y, float32 z)
func Vertex3f(x float32, y float32, z float32) {
	C.glVertex3f(C.GLfloat(x), C.GLfloat(y), C.GLfloat(z))
}

//void glVertex3fv (const float *v)
func Vertex3fv(v []float32) {
	if len(v) != 3 {
		panic("Invalid vertex coord size")
	}
	C.glVertex3fv((*C.GLfloat)(&v[0]))
}

//void glVertex3i (int x, int y, int z)
func Vertex3i(x int, y int, z int) {
	C.glVertex3i(C.GLint(x), C.GLint(y), C.GLint(z))
}

//void glVertex3iv (const int *v)
func Vertex3iv(v []int32) {
	if len(v) != 3 {
		panic("Invalid vertex coord size")
	}
	C.glVertex3iv((*C.GLint)(&v[0]))
}

//void glVertex3s (int16 x, int16 y, int16 z)
func Vertex3s(x int16, y int16, z int16) {
	C.glVertex3s(C.GLshort(x), C.GLshort(y), C.GLshort(z))
}

//void glVertex3sv (const int16 *v)
func Vertex3sv(v []int16) {
	if len(v) != 3 {
		panic("Invalid vertex coord size")
	}
	C.glVertex3sv((*C.GLshort)(&v[0]))
}

//void glVertex4d (float64 x, float64 y, float64 z, float64 w)
func Vertex4d(x float64, y float64, z float64, w float64) {
	C.glVertex4d(C.GLdouble(x), C.GLdouble(y), C.GLdouble(z), C.GLdouble(w))
}

//void glVertex4dv (const float64 *v)
func Vertex4dv(v []float64) {
	if len(v) != 4 {
		panic("Invalid vertex coord size")
	}
	C.glVertex4dv((*C.GLdouble)(&v[0]))
}

//void glVertex4f (float32 x, float32 y, float32 z, float32 w)
func Vertex4f(x float32, y float32, z float32, w float32) {
	C.glVertex4f(C.GLfloat(x), C.GLfloat(y), C.GLfloat(z), C.GLfloat(w))
}

//void glVertex4fv (const float *v)
func Vertex4fv(v []float32) {
	if len(v) != 4 {
		panic("Invalid vertex coord size")
	}
	C.glVertex4fv((*C.GLfloat)(&v[0]))
}

//void glVertex4i (int x, int y, int z, int w)
func Vertex4i(x int, y int, z int, w int) {
	C.glVertex4i(C.GLint(x), C.GLint(y), C.GLint(z), C.GLint(w))
}

//void glVertex4iv (const int *v)
func Vertex4iv(v []int32) {
	if len(v) != 4 {
		panic("Invalid vertex coord size")
	}
	C.glVertex4iv((*C.GLint)(&v[0]))
}

//void glVertex4s (int16 x, int16 y, int16 z, int16 w)
func Vertex4s(x int16, y int16, z int16, w int16) {
	C.glVertex4s(C.GLshort(x), C.GLshort(y), C.GLshort(z), C.GLshort(w))
}

//void glVertex4sv (const int16 *v)
func Vertex4sv(v []int16) {
	if len(v) != 4 {
		panic("Invalid vertex coord size")
	}
	C.glVertex4sv((*C.GLshort)(&v[0]))
}

//void glVertexPointer (int size, GLenum type, int stride, const GLvoid *pointer)
func VertexPointer(size int, typ GLenum, stride int, pointer interface{}) {
	C.glVertexPointer(C.GLint(size), C.GLenum(typ), C.GLsizei(stride), ptr(pointer))
}

//void glViewport (int x, int y, int width, int height)
func Viewport(x int, y int, width int, height int) {
	C.glViewport(C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height))
}

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
	if len(bufs) == 0 {
		panic("Invalid buffer size")
	}
	C.glGenRenderbuffers(C.GLsizei(len(bufs)), (*C.GLuint)(&bufs[0]))
}

// void glBindRenderbuffer(GLenum target, GLuint renderbuffer);
func (rb Renderbuffer) Bind() {
	C.glBindRenderbuffer(C.GLenum(RENDERBUFFER), C.GLuint(rb))
}

// Unbind this texture
func (rb Renderbuffer) Unbind() {
	C.glBindRenderbuffer(C.GLenum(RENDERBUFFER), 0)
}

// void glDeleteRenderbuffers(GLsizei n, GLuint* renderbuffers);
func (rb Renderbuffer) Delete() {
	C.glDeleteRenderbuffers(1, (*C.GLuint)(&rb))
}

func DeleteRenderbuffers(bufs []Renderbuffer) {
	if len(bufs) == 0 {
		panic("Invalid buffer size")
	}
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
func (fb Framebuffer) Bind() {
	C.glBindFramebuffer(C.GLenum(FRAMEBUFFER), C.GLuint(fb))
}

func (fb Framebuffer) Unbind() {
	C.glBindFramebuffer(C.GLenum(FRAMEBUFFER), 0)
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
func (fb Framebuffer) Delete() {
	C.glDeleteFramebuffers(1, (*C.GLuint)(&fb))
}

func DeleteFramebuffers(bufs []Framebuffer) {
	if len(bufs) == 0 {
		panic("Invalid buffer size")
	}
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
	C.glGenFramebuffers(1, &b)
	return Framebuffer(b)
}

func GenFramebuffers(bufs []Framebuffer) {
	if len(bufs) == 0 {
		panic("Invalid buffer size")
	}
	C.glGenFramebuffers(C.GLsizei(len(bufs)), (*C.GLuint)(&bufs[0]))
}

// void glGetFramebufferAttachmentParameter(GLenum target, GLenum attachment, GLenum pname, GLint* params);
//func GetFramebufferAttachmentParameter (target, attachment, pname GLenum, params []int) {
//  C.glGetFramebufferAttachmentParameter (C.GLenum(target), C.GLenum(attachment), C.GLenum(pname), (*C.GLint)(&params[0]))
//}

func Init() GLenum {
	return GLenum(C.glewInit())
}
