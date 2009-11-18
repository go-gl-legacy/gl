package gl

// #define GL_GLEXT_PROTOTYPES
// #include <GL/gl.h>
// #include <GL/glext.h>
// GLenum glewInit(void);
// void free(void*);
import "C"
import "unsafe"

type GLenum C.GLenum;
type GLboolean C.GLboolean;
type GLbitfield C.GLbitfield;
type GLbyte C.GLbyte;
type GLshort C.GLshort;
type GLint C.GLint;
type GLsizei C.GLsizei;
type GLubyte C.GLubyte;
type GLushort C.GLushort;
type GLuint C.GLuint;
type GLfloat C.GLfloat;
type GLclampf C.GLclampf;
type GLdouble C.GLdouble;
type GLclampd C.GLclampd;
//void glAccum (GLenum op, GLfloat value)
func Accum(op GLenum, value GLfloat) 
{
	C.glAccum(C.GLenum(op), C.GLfloat(value));
}

//void glAlphaFunc (GLenum func, GLclampf ref)
func AlphaFunc(func_ GLenum, ref GLclampf) 
{
	C.glAlphaFunc(C.GLenum(func_), C.GLclampf(ref));
}

//GLboolean glAreTexturesResident (GLsizei n, const GLuint *textures, GLboolean *residences)
func AreTexturesResident(n GLsizei, textures *GLuint, residences *GLboolean) GLboolean
{
	return GLboolean(C.glAreTexturesResident(C.GLsizei(n), (*C.GLuint)(textures), (*C.GLboolean)(residences)));
}

//void glArrayElement (GLint i)
func ArrayElement(i GLint) 
{
	C.glArrayElement(C.GLint(i));
}

//void glBegin (GLenum mode)
func Begin(mode GLenum) 
{
	C.glBegin(C.GLenum(mode));
}

//void glBindTexture (GLenum target, GLuint texture)
func BindTexture(target GLenum, texture GLuint) 
{
	C.glBindTexture(C.GLenum(target), C.GLuint(texture));
}

//void glBitmap (GLsizei width, GLsizei height, GLfloat xorig, GLfloat yorig, GLfloat xmove, GLfloat ymove, const GLubyte *bitmap)
func Bitmap(width GLsizei, height GLsizei, xorig GLfloat, yorig GLfloat, xmove GLfloat, ymove GLfloat, bitmap *GLubyte) 
{
	C.glBitmap(C.GLsizei(width), C.GLsizei(height), C.GLfloat(xorig), C.GLfloat(yorig), C.GLfloat(xmove), C.GLfloat(ymove), (*C.GLubyte)(bitmap));
}

//void glBlendFunc (GLenum sfactor, GLenum dfactor)
func BlendFunc(sfactor GLenum, dfactor GLenum) 
{
	C.glBlendFunc(C.GLenum(sfactor), C.GLenum(dfactor));
}

//void glCallList (GLuint list)
func CallList(list GLuint) 
{
	C.glCallList(C.GLuint(list));
}

//void glCallLists (GLsizei n, GLenum type, const GLvoid *lists)
func CallLists(n GLsizei, type_ GLenum, lists unsafe.Pointer) 
{
	C.glCallLists(C.GLsizei(n), C.GLenum(type_), lists);
}

//void glClear (GLbitfield mask)
func Clear(mask GLbitfield) 
{
	C.glClear(C.GLbitfield(mask));
}

//void glClearAccum (GLfloat red, GLfloat green, GLfloat blue, GLfloat alpha)
func ClearAccum(red GLfloat, green GLfloat, blue GLfloat, alpha GLfloat) 
{
	C.glClearAccum(C.GLfloat(red), C.GLfloat(green), C.GLfloat(blue), C.GLfloat(alpha));
}

//void glClearColor (GLclampf red, GLclampf green, GLclampf blue, GLclampf alpha)
func ClearColor(red GLclampf, green GLclampf, blue GLclampf, alpha GLclampf) 
{
	C.glClearColor(C.GLclampf(red), C.GLclampf(green), C.GLclampf(blue), C.GLclampf(alpha));
}

//void glClearDepth (GLclampd depth)
func ClearDepth(depth GLclampd) 
{
	C.glClearDepth(C.GLclampd(depth));
}

//void glClearIndex (GLfloat c)
func ClearIndex(c GLfloat) 
{
	C.glClearIndex(C.GLfloat(c));
}

//void glClearStencil (GLint s)
func ClearStencil(s GLint) 
{
	C.glClearStencil(C.GLint(s));
}

//void glClipPlane (GLenum plane, const GLdouble *equation)
func ClipPlane(plane GLenum, equation *GLdouble) 
{
	C.glClipPlane(C.GLenum(plane), (*C.GLdouble)(equation));
}

//void glColor3b (GLbyte red, GLbyte green, GLbyte blue)
func Color3b(red GLbyte, green GLbyte, blue GLbyte) 
{
	C.glColor3b(C.GLbyte(red), C.GLbyte(green), C.GLbyte(blue));
}

//void glColor3bv (const GLbyte *v)
func Color3bv(v *GLbyte) 
{
	C.glColor3bv((*C.GLbyte)(v));
}

//void glColor3d (GLdouble red, GLdouble green, GLdouble blue)
func Color3d(red GLdouble, green GLdouble, blue GLdouble) 
{
	C.glColor3d(C.GLdouble(red), C.GLdouble(green), C.GLdouble(blue));
}

//void glColor3dv (const GLdouble *v)
func Color3dv(v *GLdouble) 
{
	C.glColor3dv((*C.GLdouble)(v));
}

//void glColor3f (GLfloat red, GLfloat green, GLfloat blue)
func Color3f(red GLfloat, green GLfloat, blue GLfloat) 
{
	C.glColor3f(C.GLfloat(red), C.GLfloat(green), C.GLfloat(blue));
}

//void glColor3fv (const GLfloat *v)
func Color3fv(v *GLfloat) 
{
	C.glColor3fv((*C.GLfloat)(v));
}

//void glColor3i (GLint red, GLint green, GLint blue)
func Color3i(red GLint, green GLint, blue GLint) 
{
	C.glColor3i(C.GLint(red), C.GLint(green), C.GLint(blue));
}

//void glColor3iv (const GLint *v)
func Color3iv(v *GLint) 
{
	C.glColor3iv((*C.GLint)(v));
}

//void glColor3s (GLshort red, GLshort green, GLshort blue)
func Color3s(red GLshort, green GLshort, blue GLshort) 
{
	C.glColor3s(C.GLshort(red), C.GLshort(green), C.GLshort(blue));
}

//void glColor3sv (const GLshort *v)
func Color3sv(v *GLshort) 
{
	C.glColor3sv((*C.GLshort)(v));
}

//void glColor3ub (GLubyte red, GLubyte green, GLubyte blue)
func Color3ub(red GLubyte, green GLubyte, blue GLubyte) 
{
	C.glColor3ub(C.GLubyte(red), C.GLubyte(green), C.GLubyte(blue));
}

//void glColor3ubv (const GLubyte *v)
func Color3ubv(v *GLubyte) 
{
	C.glColor3ubv((*C.GLubyte)(v));
}

//void glColor3ui (GLuint red, GLuint green, GLuint blue)
func Color3ui(red GLuint, green GLuint, blue GLuint) 
{
	C.glColor3ui(C.GLuint(red), C.GLuint(green), C.GLuint(blue));
}

//void glColor3uiv (const GLuint *v)
func Color3uiv(v *GLuint) 
{
	C.glColor3uiv((*C.GLuint)(v));
}

//void glColor3us (GLushort red, GLushort green, GLushort blue)
func Color3us(red GLushort, green GLushort, blue GLushort) 
{
	C.glColor3us(C.GLushort(red), C.GLushort(green), C.GLushort(blue));
}

//void glColor3usv (const GLushort *v)
func Color3usv(v *GLushort) 
{
	C.glColor3usv((*C.GLushort)(v));
}

//void glColor4b (GLbyte red, GLbyte green, GLbyte blue, GLbyte alpha)
func Color4b(red GLbyte, green GLbyte, blue GLbyte, alpha GLbyte) 
{
	C.glColor4b(C.GLbyte(red), C.GLbyte(green), C.GLbyte(blue), C.GLbyte(alpha));
}

//void glColor4bv (const GLbyte *v)
func Color4bv(v *GLbyte) 
{
	C.glColor4bv((*C.GLbyte)(v));
}

//void glColor4d (GLdouble red, GLdouble green, GLdouble blue, GLdouble alpha)
func Color4d(red GLdouble, green GLdouble, blue GLdouble, alpha GLdouble) 
{
	C.glColor4d(C.GLdouble(red), C.GLdouble(green), C.GLdouble(blue), C.GLdouble(alpha));
}

//void glColor4dv (const GLdouble *v)
func Color4dv(v *GLdouble) 
{
	C.glColor4dv((*C.GLdouble)(v));
}

//void glColor4f (GLfloat red, GLfloat green, GLfloat blue, GLfloat alpha)
func Color4f(red GLfloat, green GLfloat, blue GLfloat, alpha GLfloat) 
{
	C.glColor4f(C.GLfloat(red), C.GLfloat(green), C.GLfloat(blue), C.GLfloat(alpha));
}

//void glColor4fv (const GLfloat *v)
func Color4fv(v *GLfloat) 
{
	C.glColor4fv((*C.GLfloat)(v));
}

//void glColor4i (GLint red, GLint green, GLint blue, GLint alpha)
func Color4i(red GLint, green GLint, blue GLint, alpha GLint) 
{
	C.glColor4i(C.GLint(red), C.GLint(green), C.GLint(blue), C.GLint(alpha));
}

//void glColor4iv (const GLint *v)
func Color4iv(v *GLint) 
{
	C.glColor4iv((*C.GLint)(v));
}

//void glColor4s (GLshort red, GLshort green, GLshort blue, GLshort alpha)
func Color4s(red GLshort, green GLshort, blue GLshort, alpha GLshort) 
{
	C.glColor4s(C.GLshort(red), C.GLshort(green), C.GLshort(blue), C.GLshort(alpha));
}

//void glColor4sv (const GLshort *v)
func Color4sv(v *GLshort) 
{
	C.glColor4sv((*C.GLshort)(v));
}

//void glColor4ub (GLubyte red, GLubyte green, GLubyte blue, GLubyte alpha)
func Color4ub(red GLubyte, green GLubyte, blue GLubyte, alpha GLubyte) 
{
	C.glColor4ub(C.GLubyte(red), C.GLubyte(green), C.GLubyte(blue), C.GLubyte(alpha));
}

//void glColor4ubv (const GLubyte *v)
func Color4ubv(v *GLubyte) 
{
	C.glColor4ubv((*C.GLubyte)(v));
}

//void glColor4ui (GLuint red, GLuint green, GLuint blue, GLuint alpha)
func Color4ui(red GLuint, green GLuint, blue GLuint, alpha GLuint) 
{
	C.glColor4ui(C.GLuint(red), C.GLuint(green), C.GLuint(blue), C.GLuint(alpha));
}

//void glColor4uiv (const GLuint *v)
func Color4uiv(v *GLuint) 
{
	C.glColor4uiv((*C.GLuint)(v));
}

//void glColor4us (GLushort red, GLushort green, GLushort blue, GLushort alpha)
func Color4us(red GLushort, green GLushort, blue GLushort, alpha GLushort) 
{
	C.glColor4us(C.GLushort(red), C.GLushort(green), C.GLushort(blue), C.GLushort(alpha));
}

//void glColor4usv (const GLushort *v)
func Color4usv(v *GLushort) 
{
	C.glColor4usv((*C.GLushort)(v));
}

//void glColorMask (GLboolean red, GLboolean green, GLboolean blue, GLboolean alpha)
func ColorMask(red GLboolean, green GLboolean, blue GLboolean, alpha GLboolean) 
{
	C.glColorMask(C.GLboolean(red), C.GLboolean(green), C.GLboolean(blue), C.GLboolean(alpha));
}

//void glColorMaterial (GLenum face, GLenum mode)
func ColorMaterial(face GLenum, mode GLenum) 
{
	C.glColorMaterial(C.GLenum(face), C.GLenum(mode));
}

//void glColorPointer (GLint size, GLenum type, GLsizei stride, const GLvoid *pointer)
func ColorPointer(size GLint, type_ GLenum, stride GLsizei, pointer unsafe.Pointer) 
{
	C.glColorPointer(C.GLint(size), C.GLenum(type_), C.GLsizei(stride), pointer);
}

//void glCopyPixels (GLint x, GLint y, GLsizei width, GLsizei height, GLenum type)
func CopyPixels(x GLint, y GLint, width GLsizei, height GLsizei, type_ GLenum) 
{
	C.glCopyPixels(C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height), C.GLenum(type_));
}

//void glCopyTexImage1D (GLenum target, GLint level, GLenum internalFormat, GLint x, GLint y, GLsizei width, GLint border)
func CopyTexImage1D(target GLenum, level GLint, internalFormat GLenum, x GLint, y GLint, width GLsizei, border GLint) 
{
	C.glCopyTexImage1D(C.GLenum(target), C.GLint(level), C.GLenum(internalFormat), C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLint(border));
}

//void glCopyTexImage2D (GLenum target, GLint level, GLenum internalFormat, GLint x, GLint y, GLsizei width, GLsizei height, GLint border)
func CopyTexImage2D(target GLenum, level GLint, internalFormat GLenum, x GLint, y GLint, width GLsizei, height GLsizei, border GLint) 
{
	C.glCopyTexImage2D(C.GLenum(target), C.GLint(level), C.GLenum(internalFormat), C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height), C.GLint(border));
}

//void glCopyTexSubImage1D (GLenum target, GLint level, GLint xoffset, GLint x, GLint y, GLsizei width)
func CopyTexSubImage1D(target GLenum, level GLint, xoffset GLint, x GLint, y GLint, width GLsizei) 
{
	C.glCopyTexSubImage1D(C.GLenum(target), C.GLint(level), C.GLint(xoffset), C.GLint(x), C.GLint(y), C.GLsizei(width));
}

//void glCopyTexSubImage2D (GLenum target, GLint level, GLint xoffset, GLint yoffset, GLint x, GLint y, GLsizei width, GLsizei height)
func CopyTexSubImage2D(target GLenum, level GLint, xoffset GLint, yoffset GLint, x GLint, y GLint, width GLsizei, height GLsizei) 
{
	C.glCopyTexSubImage2D(C.GLenum(target), C.GLint(level), C.GLint(xoffset), C.GLint(yoffset), C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height));
}

//void glCullFace (GLenum mode)
func CullFace(mode GLenum) 
{
	C.glCullFace(C.GLenum(mode));
}

//void glDeleteLists (GLuint list, GLsizei range)
func DeleteLists(list GLuint, range_ GLsizei) 
{
	C.glDeleteLists(C.GLuint(list), C.GLsizei(range_));
}

//void glDeleteTextures (GLsizei n, const GLuint *textures)
func DeleteTextures(n GLsizei, textures *GLuint) 
{
	C.glDeleteTextures(C.GLsizei(n), (*C.GLuint)(textures));
}

//void glDepthFunc (GLenum func)
func DepthFunc(func_ GLenum) 
{
	C.glDepthFunc(C.GLenum(func_));
}

//void glDepthMask (GLboolean flag)
func DepthMask(flag GLboolean) 
{
	C.glDepthMask(C.GLboolean(flag));
}

//void glDepthRange (GLclampd zNear, GLclampd zFar)
func DepthRange(zNear GLclampd, zFar GLclampd) 
{
	C.glDepthRange(C.GLclampd(zNear), C.GLclampd(zFar));
}

//void glDisable (GLenum cap)
func Disable(cap GLenum) 
{
	C.glDisable(C.GLenum(cap));
}

//void glDisableClientState (GLenum array)
func DisableClientState(array GLenum) 
{
	C.glDisableClientState(C.GLenum(array));
}

//void glDrawArrays (GLenum mode, GLint first, GLsizei count)
func DrawArrays(mode GLenum, first GLint, count GLsizei) 
{
	C.glDrawArrays(C.GLenum(mode), C.GLint(first), C.GLsizei(count));
}

//void glDrawBuffer (GLenum mode)
func DrawBuffer(mode GLenum) 
{
	C.glDrawBuffer(C.GLenum(mode));
}

//void glDrawElements (GLenum mode, GLsizei count, GLenum type, const GLvoid *indices)
func DrawElements(mode GLenum, count GLsizei, type_ GLenum, indices unsafe.Pointer) 
{
	C.glDrawElements(C.GLenum(mode), C.GLsizei(count), C.GLenum(type_), indices);
}

//void glDrawPixels (GLsizei width, GLsizei height, GLenum format, GLenum type, const GLvoid *pixels)
func DrawPixels(width GLsizei, height GLsizei, format GLenum, type_ GLenum, pixels unsafe.Pointer) 
{
	C.glDrawPixels(C.GLsizei(width), C.GLsizei(height), C.GLenum(format), C.GLenum(type_), pixels);
}

//void glEdgeFlag (GLboolean flag)
func EdgeFlag(flag GLboolean) 
{
	C.glEdgeFlag(C.GLboolean(flag));
}

//void glEdgeFlagPointer (GLsizei stride, const GLvoid *pointer)
func EdgeFlagPointer(stride GLsizei, pointer unsafe.Pointer) 
{
	C.glEdgeFlagPointer(C.GLsizei(stride), pointer);
}

//void glEdgeFlagv (const GLboolean *flag)
func EdgeFlagv(flag *GLboolean) 
{
	C.glEdgeFlagv((*C.GLboolean)(flag));
}

//void glEnable (GLenum cap)
func Enable(cap GLenum) 
{
	C.glEnable(C.GLenum(cap));
}

//void glEnableClientState (GLenum array)
func EnableClientState(array GLenum) 
{
	C.glEnableClientState(C.GLenum(array));
}

//void glEnd (void)
func End() 
{
	C.glEnd();
}

//void glEndList (void)
func EndList() 
{
	C.glEndList();
}

//void glEvalCoord1d (GLdouble u)
func EvalCoord1d(u GLdouble) 
{
	C.glEvalCoord1d(C.GLdouble(u));
}

//void glEvalCoord1dv (const GLdouble *u)
func EvalCoord1dv(u *GLdouble) 
{
	C.glEvalCoord1dv((*C.GLdouble)(u));
}

//void glEvalCoord1f (GLfloat u)
func EvalCoord1f(u GLfloat) 
{
	C.glEvalCoord1f(C.GLfloat(u));
}

//void glEvalCoord1fv (const GLfloat *u)
func EvalCoord1fv(u *GLfloat) 
{
	C.glEvalCoord1fv((*C.GLfloat)(u));
}

//void glEvalCoord2d (GLdouble u, GLdouble v)
func EvalCoord2d(u GLdouble, v GLdouble) 
{
	C.glEvalCoord2d(C.GLdouble(u), C.GLdouble(v));
}

//void glEvalCoord2dv (const GLdouble *u)
func EvalCoord2dv(u *GLdouble) 
{
	C.glEvalCoord2dv((*C.GLdouble)(u));
}

//void glEvalCoord2f (GLfloat u, GLfloat v)
func EvalCoord2f(u GLfloat, v GLfloat) 
{
	C.glEvalCoord2f(C.GLfloat(u), C.GLfloat(v));
}

//void glEvalCoord2fv (const GLfloat *u)
func EvalCoord2fv(u *GLfloat) 
{
	C.glEvalCoord2fv((*C.GLfloat)(u));
}

//void glEvalMesh1 (GLenum mode, GLint i1, GLint i2)
func EvalMesh1(mode GLenum, i1 GLint, i2 GLint) 
{
	C.glEvalMesh1(C.GLenum(mode), C.GLint(i1), C.GLint(i2));
}

//void glEvalMesh2 (GLenum mode, GLint i1, GLint i2, GLint j1, GLint j2)
func EvalMesh2(mode GLenum, i1 GLint, i2 GLint, j1 GLint, j2 GLint) 
{
	C.glEvalMesh2(C.GLenum(mode), C.GLint(i1), C.GLint(i2), C.GLint(j1), C.GLint(j2));
}

//void glEvalPoint1 (GLint i)
func EvalPoint1(i GLint) 
{
	C.glEvalPoint1(C.GLint(i));
}

//void glEvalPoint2 (GLint i, GLint j)
func EvalPoint2(i GLint, j GLint) 
{
	C.glEvalPoint2(C.GLint(i), C.GLint(j));
}

//void glFeedbackBuffer (GLsizei size, GLenum type, GLfloat *buffer)
func FeedbackBuffer(size GLsizei, type_ GLenum, buffer *GLfloat) 
{
	C.glFeedbackBuffer(C.GLsizei(size), C.GLenum(type_), (*C.GLfloat)(buffer));
}

//void glFinish (void)
func Finish() 
{
	C.glFinish();
}

//void glFlush (void)
func Flush() 
{
	C.glFlush();
}

//void glFogf (GLenum pname, GLfloat param)
func Fogf(pname GLenum, param GLfloat) 
{
	C.glFogf(C.GLenum(pname), C.GLfloat(param));
}

//void glFogfv (GLenum pname, const GLfloat *params)
func Fogfv(pname GLenum, params *GLfloat) 
{
	C.glFogfv(C.GLenum(pname), (*C.GLfloat)(params));
}

//void glFogi (GLenum pname, GLint param)
func Fogi(pname GLenum, param GLint) 
{
	C.glFogi(C.GLenum(pname), C.GLint(param));
}

//void glFogiv (GLenum pname, const GLint *params)
func Fogiv(pname GLenum, params *GLint) 
{
	C.glFogiv(C.GLenum(pname), (*C.GLint)(params));
}

//void glFrontFace (GLenum mode)
func FrontFace(mode GLenum) 
{
	C.glFrontFace(C.GLenum(mode));
}

//void glFrustum (GLdouble left, GLdouble right, GLdouble bottom, GLdouble top, GLdouble zNear, GLdouble zFar)
func Frustum(left GLdouble, right GLdouble, bottom GLdouble, top GLdouble, zNear GLdouble, zFar GLdouble) 
{
	C.glFrustum(C.GLdouble(left), C.GLdouble(right), C.GLdouble(bottom), C.GLdouble(top), C.GLdouble(zNear), C.GLdouble(zFar));
}

//GLuint glGenLists (GLsizei range)
func GenLists(range_ GLsizei) GLuint
{
	return GLuint(C.glGenLists(C.GLsizei(range_)));
}

//void glGenTextures (GLsizei n, GLuint *textures)
func GenTextures(n GLsizei, textures *GLuint) 
{
	C.glGenTextures(C.GLsizei(n), (*C.GLuint)(textures));
}

//void glGetBooleanv (GLenum pname, GLboolean *params)
func GetBooleanv(pname GLenum, params *GLboolean) 
{
	C.glGetBooleanv(C.GLenum(pname), (*C.GLboolean)(params));
}

//void glGetClipPlane (GLenum plane, GLdouble *equation)
func GetClipPlane(plane GLenum, equation *GLdouble) 
{
	C.glGetClipPlane(C.GLenum(plane), (*C.GLdouble)(equation));
}

//void glGetDoublev (GLenum pname, GLdouble *params)
func GetDoublev(pname GLenum, params *GLdouble) 
{
	C.glGetDoublev(C.GLenum(pname), (*C.GLdouble)(params));
}

//GLenum glGetError (void)
func GetError() GLenum
{
	return GLenum(C.glGetError());
}

//void glGetFloatv (GLenum pname, GLfloat *params)
func GetFloatv(pname GLenum, params *GLfloat) 
{
	C.glGetFloatv(C.GLenum(pname), (*C.GLfloat)(params));
}

//void glGetIntegerv (GLenum pname, GLint *params)
func GetIntegerv(pname GLenum, params *GLint) 
{
	C.glGetIntegerv(C.GLenum(pname), (*C.GLint)(params));
}

//void glGetLightfv (GLenum light, GLenum pname, GLfloat *params)
func GetLightfv(light GLenum, pname GLenum, params *GLfloat) 
{
	C.glGetLightfv(C.GLenum(light), C.GLenum(pname), (*C.GLfloat)(params));
}

//void glGetLightiv (GLenum light, GLenum pname, GLint *params)
func GetLightiv(light GLenum, pname GLenum, params *GLint) 
{
	C.glGetLightiv(C.GLenum(light), C.GLenum(pname), (*C.GLint)(params));
}

//void glGetMapdv (GLenum target, GLenum query, GLdouble *v)
func GetMapdv(target GLenum, query GLenum, v *GLdouble) 
{
	C.glGetMapdv(C.GLenum(target), C.GLenum(query), (*C.GLdouble)(v));
}

//void glGetMapfv (GLenum target, GLenum query, GLfloat *v)
func GetMapfv(target GLenum, query GLenum, v *GLfloat) 
{
	C.glGetMapfv(C.GLenum(target), C.GLenum(query), (*C.GLfloat)(v));
}

//void glGetMapiv (GLenum target, GLenum query, GLint *v)
func GetMapiv(target GLenum, query GLenum, v *GLint) 
{
	C.glGetMapiv(C.GLenum(target), C.GLenum(query), (*C.GLint)(v));
}

//void glGetMaterialfv (GLenum face, GLenum pname, GLfloat *params)
func GetMaterialfv(face GLenum, pname GLenum, params *GLfloat) 
{
	C.glGetMaterialfv(C.GLenum(face), C.GLenum(pname), (*C.GLfloat)(params));
}

//void glGetMaterialiv (GLenum face, GLenum pname, GLint *params)
func GetMaterialiv(face GLenum, pname GLenum, params *GLint) 
{
	C.glGetMaterialiv(C.GLenum(face), C.GLenum(pname), (*C.GLint)(params));
}

//void glGetPixelMapfv (GLenum map, GLfloat *values)
func GetPixelMapfv(map_ GLenum, values *GLfloat) 
{
	C.glGetPixelMapfv(C.GLenum(map_), (*C.GLfloat)(values));
}

//void glGetPixelMapuiv (GLenum map, GLuint *values)
func GetPixelMapuiv(map_ GLenum, values *GLuint) 
{
	C.glGetPixelMapuiv(C.GLenum(map_), (*C.GLuint)(values));
}

//void glGetPixelMapusv (GLenum map, GLushort *values)
func GetPixelMapusv(map_ GLenum, values *GLushort) 
{
	C.glGetPixelMapusv(C.GLenum(map_), (*C.GLushort)(values));
}

//void glGetPointerv (GLenum pname, GLvoid* *params)
func GetPointerv(pname GLenum, params *unsafe.Pointer) 
{
	C.glGetPointerv(C.GLenum(pname), params);
}

//void glGetPolygonStipple (GLubyte *mask)
func GetPolygonStipple(mask *GLubyte) 
{
	C.glGetPolygonStipple((*C.GLubyte)(mask));
}

//const GLubyte * glGetString (GLenum name)
func GetString(name GLenum) string
{
	return C.GoString((*C.char)(unsafe.Pointer(C.glGetString(C.GLenum(name)))));
}

//void glGetTexEnvfv (GLenum target, GLenum pname, GLfloat *params)
func GetTexEnvfv(target GLenum, pname GLenum, params *GLfloat) 
{
	C.glGetTexEnvfv(C.GLenum(target), C.GLenum(pname), (*C.GLfloat)(params));
}

//void glGetTexEnviv (GLenum target, GLenum pname, GLint *params)
func GetTexEnviv(target GLenum, pname GLenum, params *GLint) 
{
	C.glGetTexEnviv(C.GLenum(target), C.GLenum(pname), (*C.GLint)(params));
}

//void glGetTexGendv (GLenum coord, GLenum pname, GLdouble *params)
func GetTexGendv(coord GLenum, pname GLenum, params *GLdouble) 
{
	C.glGetTexGendv(C.GLenum(coord), C.GLenum(pname), (*C.GLdouble)(params));
}

//void glGetTexGenfv (GLenum coord, GLenum pname, GLfloat *params)
func GetTexGenfv(coord GLenum, pname GLenum, params *GLfloat) 
{
	C.glGetTexGenfv(C.GLenum(coord), C.GLenum(pname), (*C.GLfloat)(params));
}

//void glGetTexGeniv (GLenum coord, GLenum pname, GLint *params)
func GetTexGeniv(coord GLenum, pname GLenum, params *GLint) 
{
	C.glGetTexGeniv(C.GLenum(coord), C.GLenum(pname), (*C.GLint)(params));
}

//void glGetTexImage (GLenum target, GLint level, GLenum format, GLenum type, GLvoid *pixels)
func GetTexImage(target GLenum, level GLint, format GLenum, type_ GLenum, pixels unsafe.Pointer) 
{
	C.glGetTexImage(C.GLenum(target), C.GLint(level), C.GLenum(format), C.GLenum(type_), pixels);
}

//void glGetTexLevelParameterfv (GLenum target, GLint level, GLenum pname, GLfloat *params)
func GetTexLevelParameterfv(target GLenum, level GLint, pname GLenum, params *GLfloat) 
{
	C.glGetTexLevelParameterfv(C.GLenum(target), C.GLint(level), C.GLenum(pname), (*C.GLfloat)(params));
}

//void glGetTexLevelParameteriv (GLenum target, GLint level, GLenum pname, GLint *params)
func GetTexLevelParameteriv(target GLenum, level GLint, pname GLenum, params *GLint) 
{
	C.glGetTexLevelParameteriv(C.GLenum(target), C.GLint(level), C.GLenum(pname), (*C.GLint)(params));
}

//void glGetTexParameterfv (GLenum target, GLenum pname, GLfloat *params)
func GetTexParameterfv(target GLenum, pname GLenum, params *GLfloat) 
{
	C.glGetTexParameterfv(C.GLenum(target), C.GLenum(pname), (*C.GLfloat)(params));
}

//void glGetTexParameteriv (GLenum target, GLenum pname, GLint *params)
func GetTexParameteriv(target GLenum, pname GLenum, params *GLint) 
{
	C.glGetTexParameteriv(C.GLenum(target), C.GLenum(pname), (*C.GLint)(params));
}

//void glHint (GLenum target, GLenum mode)
func Hint(target GLenum, mode GLenum) 
{
	C.glHint(C.GLenum(target), C.GLenum(mode));
}

//void glIndexMask (GLuint mask)
func IndexMask(mask GLuint) 
{
	C.glIndexMask(C.GLuint(mask));
}

//void glIndexPointer (GLenum type, GLsizei stride, const GLvoid *pointer)
func IndexPointer(type_ GLenum, stride GLsizei, pointer unsafe.Pointer) 
{
	C.glIndexPointer(C.GLenum(type_), C.GLsizei(stride), pointer);
}

//void glIndexd (GLdouble c)
func Indexd(c GLdouble) 
{
	C.glIndexd(C.GLdouble(c));
}

//void glIndexdv (const GLdouble *c)
func Indexdv(c *GLdouble) 
{
	C.glIndexdv((*C.GLdouble)(c));
}

//void glIndexf (GLfloat c)
func Indexf(c GLfloat) 
{
	C.glIndexf(C.GLfloat(c));
}

//void glIndexfv (const GLfloat *c)
func Indexfv(c *GLfloat) 
{
	C.glIndexfv((*C.GLfloat)(c));
}

//void glIndexi (GLint c)
func Indexi(c GLint) 
{
	C.glIndexi(C.GLint(c));
}

//void glIndexiv (const GLint *c)
func Indexiv(c *GLint) 
{
	C.glIndexiv((*C.GLint)(c));
}

//void glIndexs (GLshort c)
func Indexs(c GLshort) 
{
	C.glIndexs(C.GLshort(c));
}

//void glIndexsv (const GLshort *c)
func Indexsv(c *GLshort) 
{
	C.glIndexsv((*C.GLshort)(c));
}

//void glIndexub (GLubyte c)
func Indexub(c GLubyte) 
{
	C.glIndexub(C.GLubyte(c));
}

//void glIndexubv (const GLubyte *c)
func Indexubv(c *GLubyte) 
{
	C.glIndexubv((*C.GLubyte)(c));
}

//void glInitNames (void)
func InitNames() 
{
	C.glInitNames();
}

//void glInterleavedArrays (GLenum format, GLsizei stride, const GLvoid *pointer)
func InterleavedArrays(format GLenum, stride GLsizei, pointer unsafe.Pointer) 
{
	C.glInterleavedArrays(C.GLenum(format), C.GLsizei(stride), pointer);
}

//GLboolean glIsEnabled (GLenum cap)
func IsEnabled(cap GLenum) GLboolean
{
	return GLboolean(C.glIsEnabled(C.GLenum(cap)));
}

//GLboolean glIsList (GLuint list)
func IsList(list GLuint) GLboolean
{
	return GLboolean(C.glIsList(C.GLuint(list)));
}

//GLboolean glIsTexture (GLuint texture)
func IsTexture(texture GLuint) GLboolean
{
	return GLboolean(C.glIsTexture(C.GLuint(texture)));
}

//void glLightModelf (GLenum pname, GLfloat param)
func LightModelf(pname GLenum, param GLfloat) 
{
	C.glLightModelf(C.GLenum(pname), C.GLfloat(param));
}

//void glLightModelfv (GLenum pname, const GLfloat *params)
func LightModelfv(pname GLenum, params *GLfloat) 
{
	C.glLightModelfv(C.GLenum(pname), (*C.GLfloat)(params));
}

//void glLightModeli (GLenum pname, GLint param)
func LightModeli(pname GLenum, param GLint) 
{
	C.glLightModeli(C.GLenum(pname), C.GLint(param));
}

//void glLightModeliv (GLenum pname, const GLint *params)
func LightModeliv(pname GLenum, params *GLint) 
{
	C.glLightModeliv(C.GLenum(pname), (*C.GLint)(params));
}

//void glLightf (GLenum light, GLenum pname, GLfloat param)
func Lightf(light GLenum, pname GLenum, param GLfloat) 
{
	C.glLightf(C.GLenum(light), C.GLenum(pname), C.GLfloat(param));
}

//void glLightfv (GLenum light, GLenum pname, const GLfloat *params)
func Lightfv(light GLenum, pname GLenum, params *GLfloat) 
{
	C.glLightfv(C.GLenum(light), C.GLenum(pname), (*C.GLfloat)(params));
}

//void glLighti (GLenum light, GLenum pname, GLint param)
func Lighti(light GLenum, pname GLenum, param GLint) 
{
	C.glLighti(C.GLenum(light), C.GLenum(pname), C.GLint(param));
}

//void glLightiv (GLenum light, GLenum pname, const GLint *params)
func Lightiv(light GLenum, pname GLenum, params *GLint) 
{
	C.glLightiv(C.GLenum(light), C.GLenum(pname), (*C.GLint)(params));
}

//void glLineStipple (GLint factor, GLushort pattern)
func LineStipple(factor GLint, pattern GLushort) 
{
	C.glLineStipple(C.GLint(factor), C.GLushort(pattern));
}

//void glLineWidth (GLfloat width)
func LineWidth(width GLfloat) 
{
	C.glLineWidth(C.GLfloat(width));
}

//void glListBase (GLuint base)
func ListBase(base GLuint) 
{
	C.glListBase(C.GLuint(base));
}

//void glLoadIdentity (void)
func LoadIdentity() 
{
	C.glLoadIdentity();
}

//void glLoadMatrixd (const GLdouble *m)
func LoadMatrixd(m *GLdouble) 
{
	C.glLoadMatrixd((*C.GLdouble)(m));
}

//void glLoadMatrixf (const GLfloat *m)
func LoadMatrixf(m *GLfloat) 
{
	C.glLoadMatrixf((*C.GLfloat)(m));
}

//void glLoadName (GLuint name)
func LoadName(name GLuint) 
{
	C.glLoadName(C.GLuint(name));
}

//void glLogicOp (GLenum opcode)
func LogicOp(opcode GLenum) 
{
	C.glLogicOp(C.GLenum(opcode));
}

//void glMap1d (GLenum target, GLdouble u1, GLdouble u2, GLint stride, GLint order, const GLdouble *points)
func Map1d(target GLenum, u1 GLdouble, u2 GLdouble, stride GLint, order GLint, points *GLdouble) 
{
	C.glMap1d(C.GLenum(target), C.GLdouble(u1), C.GLdouble(u2), C.GLint(stride), C.GLint(order), (*C.GLdouble)(points));
}

//void glMap1f (GLenum target, GLfloat u1, GLfloat u2, GLint stride, GLint order, const GLfloat *points)
func Map1f(target GLenum, u1 GLfloat, u2 GLfloat, stride GLint, order GLint, points *GLfloat) 
{
	C.glMap1f(C.GLenum(target), C.GLfloat(u1), C.GLfloat(u2), C.GLint(stride), C.GLint(order), (*C.GLfloat)(points));
}

//void glMap2d (GLenum target, GLdouble u1, GLdouble u2, GLint ustride, GLint uorder, GLdouble v1, GLdouble v2, GLint vstride, GLint vorder, const GLdouble *points)
func Map2d(target GLenum, u1 GLdouble, u2 GLdouble, ustride GLint, uorder GLint, v1 GLdouble, v2 GLdouble, vstride GLint, vorder GLint, points *GLdouble) 
{
	C.glMap2d(C.GLenum(target), C.GLdouble(u1), C.GLdouble(u2), C.GLint(ustride), C.GLint(uorder), C.GLdouble(v1), C.GLdouble(v2), C.GLint(vstride), C.GLint(vorder), (*C.GLdouble)(points));
}

//void glMap2f (GLenum target, GLfloat u1, GLfloat u2, GLint ustride, GLint uorder, GLfloat v1, GLfloat v2, GLint vstride, GLint vorder, const GLfloat *points)
func Map2f(target GLenum, u1 GLfloat, u2 GLfloat, ustride GLint, uorder GLint, v1 GLfloat, v2 GLfloat, vstride GLint, vorder GLint, points *GLfloat) 
{
	C.glMap2f(C.GLenum(target), C.GLfloat(u1), C.GLfloat(u2), C.GLint(ustride), C.GLint(uorder), C.GLfloat(v1), C.GLfloat(v2), C.GLint(vstride), C.GLint(vorder), (*C.GLfloat)(points));
}

//void glMapGrid1d (GLint un, GLdouble u1, GLdouble u2)
func MapGrid1d(un GLint, u1 GLdouble, u2 GLdouble) 
{
	C.glMapGrid1d(C.GLint(un), C.GLdouble(u1), C.GLdouble(u2));
}

//void glMapGrid1f (GLint un, GLfloat u1, GLfloat u2)
func MapGrid1f(un GLint, u1 GLfloat, u2 GLfloat) 
{
	C.glMapGrid1f(C.GLint(un), C.GLfloat(u1), C.GLfloat(u2));
}

//void glMapGrid2d (GLint un, GLdouble u1, GLdouble u2, GLint vn, GLdouble v1, GLdouble v2)
func MapGrid2d(un GLint, u1 GLdouble, u2 GLdouble, vn GLint, v1 GLdouble, v2 GLdouble) 
{
	C.glMapGrid2d(C.GLint(un), C.GLdouble(u1), C.GLdouble(u2), C.GLint(vn), C.GLdouble(v1), C.GLdouble(v2));
}

//void glMapGrid2f (GLint un, GLfloat u1, GLfloat u2, GLint vn, GLfloat v1, GLfloat v2)
func MapGrid2f(un GLint, u1 GLfloat, u2 GLfloat, vn GLint, v1 GLfloat, v2 GLfloat) 
{
	C.glMapGrid2f(C.GLint(un), C.GLfloat(u1), C.GLfloat(u2), C.GLint(vn), C.GLfloat(v1), C.GLfloat(v2));
}

//void glMaterialf (GLenum face, GLenum pname, GLfloat param)
func Materialf(face GLenum, pname GLenum, param GLfloat) 
{
	C.glMaterialf(C.GLenum(face), C.GLenum(pname), C.GLfloat(param));
}

//void glMaterialfv (GLenum face, GLenum pname, const GLfloat *params)
func Materialfv(face GLenum, pname GLenum, params *GLfloat) 
{
	C.glMaterialfv(C.GLenum(face), C.GLenum(pname), (*C.GLfloat)(params));
}

//void glMateriali (GLenum face, GLenum pname, GLint param)
func Materiali(face GLenum, pname GLenum, param GLint) 
{
	C.glMateriali(C.GLenum(face), C.GLenum(pname), C.GLint(param));
}

//void glMaterialiv (GLenum face, GLenum pname, const GLint *params)
func Materialiv(face GLenum, pname GLenum, params *GLint) 
{
	C.glMaterialiv(C.GLenum(face), C.GLenum(pname), (*C.GLint)(params));
}

//void glMatrixMode (GLenum mode)
func MatrixMode(mode GLenum) 
{
	C.glMatrixMode(C.GLenum(mode));
}

//void glMultMatrixd (const GLdouble *m)
func MultMatrixd(m *GLdouble) 
{
	C.glMultMatrixd((*C.GLdouble)(m));
}

//void glMultMatrixf (const GLfloat *m)
func MultMatrixf(m *GLfloat) 
{
	C.glMultMatrixf((*C.GLfloat)(m));
}

//void glNewList (GLuint list, GLenum mode)
func NewList(list GLuint, mode GLenum) 
{
	C.glNewList(C.GLuint(list), C.GLenum(mode));
}

//void glNormal3b (GLbyte nx, GLbyte ny, GLbyte nz)
func Normal3b(nx GLbyte, ny GLbyte, nz GLbyte) 
{
	C.glNormal3b(C.GLbyte(nx), C.GLbyte(ny), C.GLbyte(nz));
}

//void glNormal3bv (const GLbyte *v)
func Normal3bv(v *GLbyte) 
{
	C.glNormal3bv((*C.GLbyte)(v));
}

//void glNormal3d (GLdouble nx, GLdouble ny, GLdouble nz)
func Normal3d(nx GLdouble, ny GLdouble, nz GLdouble) 
{
	C.glNormal3d(C.GLdouble(nx), C.GLdouble(ny), C.GLdouble(nz));
}

//void glNormal3dv (const GLdouble *v)
func Normal3dv(v *GLdouble) 
{
	C.glNormal3dv((*C.GLdouble)(v));
}

//void glNormal3f (GLfloat nx, GLfloat ny, GLfloat nz)
func Normal3f(nx GLfloat, ny GLfloat, nz GLfloat) 
{
	C.glNormal3f(C.GLfloat(nx), C.GLfloat(ny), C.GLfloat(nz));
}

//void glNormal3fv (const GLfloat *v)
func Normal3fv(v *GLfloat) 
{
	C.glNormal3fv((*C.GLfloat)(v));
}

//void glNormal3i (GLint nx, GLint ny, GLint nz)
func Normal3i(nx GLint, ny GLint, nz GLint) 
{
	C.glNormal3i(C.GLint(nx), C.GLint(ny), C.GLint(nz));
}

//void glNormal3iv (const GLint *v)
func Normal3iv(v *GLint) 
{
	C.glNormal3iv((*C.GLint)(v));
}

//void glNormal3s (GLshort nx, GLshort ny, GLshort nz)
func Normal3s(nx GLshort, ny GLshort, nz GLshort) 
{
	C.glNormal3s(C.GLshort(nx), C.GLshort(ny), C.GLshort(nz));
}

//void glNormal3sv (const GLshort *v)
func Normal3sv(v *GLshort) 
{
	C.glNormal3sv((*C.GLshort)(v));
}

//void glNormalPointer (GLenum type, GLsizei stride, const GLvoid *pointer)
func NormalPointer(type_ GLenum, stride GLsizei, pointer unsafe.Pointer) 
{
	C.glNormalPointer(C.GLenum(type_), C.GLsizei(stride), pointer);
}

//void glOrtho (GLdouble left, GLdouble right, GLdouble bottom, GLdouble top, GLdouble zNear, GLdouble zFar)
func Ortho(left GLdouble, right GLdouble, bottom GLdouble, top GLdouble, zNear GLdouble, zFar GLdouble) 
{
	C.glOrtho(C.GLdouble(left), C.GLdouble(right), C.GLdouble(bottom), C.GLdouble(top), C.GLdouble(zNear), C.GLdouble(zFar));
}

//void glPassThrough (GLfloat token)
func PassThrough(token GLfloat) 
{
	C.glPassThrough(C.GLfloat(token));
}

//void glPixelMapfv (GLenum map, GLsizei mapsize, const GLfloat *values)
func PixelMapfv(map_ GLenum, mapsize GLsizei, values *GLfloat) 
{
	C.glPixelMapfv(C.GLenum(map_), C.GLsizei(mapsize), (*C.GLfloat)(values));
}

//void glPixelMapuiv (GLenum map, GLsizei mapsize, const GLuint *values)
func PixelMapuiv(map_ GLenum, mapsize GLsizei, values *GLuint) 
{
	C.glPixelMapuiv(C.GLenum(map_), C.GLsizei(mapsize), (*C.GLuint)(values));
}

//void glPixelMapusv (GLenum map, GLsizei mapsize, const GLushort *values)
func PixelMapusv(map_ GLenum, mapsize GLsizei, values *GLushort) 
{
	C.glPixelMapusv(C.GLenum(map_), C.GLsizei(mapsize), (*C.GLushort)(values));
}

//void glPixelStoref (GLenum pname, GLfloat param)
func PixelStoref(pname GLenum, param GLfloat) 
{
	C.glPixelStoref(C.GLenum(pname), C.GLfloat(param));
}

//void glPixelStorei (GLenum pname, GLint param)
func PixelStorei(pname GLenum, param GLint) 
{
	C.glPixelStorei(C.GLenum(pname), C.GLint(param));
}

//void glPixelTransferf (GLenum pname, GLfloat param)
func PixelTransferf(pname GLenum, param GLfloat) 
{
	C.glPixelTransferf(C.GLenum(pname), C.GLfloat(param));
}

//void glPixelTransferi (GLenum pname, GLint param)
func PixelTransferi(pname GLenum, param GLint) 
{
	C.glPixelTransferi(C.GLenum(pname), C.GLint(param));
}

//void glPixelZoom (GLfloat xfactor, GLfloat yfactor)
func PixelZoom(xfactor GLfloat, yfactor GLfloat) 
{
	C.glPixelZoom(C.GLfloat(xfactor), C.GLfloat(yfactor));
}

//void glPointSize (GLfloat size)
func PointSize(size GLfloat) 
{
	C.glPointSize(C.GLfloat(size));
}

//void glPolygonMode (GLenum face, GLenum mode)
func PolygonMode(face GLenum, mode GLenum) 
{
	C.glPolygonMode(C.GLenum(face), C.GLenum(mode));
}

//void glPolygonOffset (GLfloat factor, GLfloat units)
func PolygonOffset(factor GLfloat, units GLfloat) 
{
	C.glPolygonOffset(C.GLfloat(factor), C.GLfloat(units));
}

//void glPolygonStipple (const GLubyte *mask)
func PolygonStipple(mask *GLubyte) 
{
	C.glPolygonStipple((*C.GLubyte)(mask));
}

//void glPopAttrib (void)
func PopAttrib() 
{
	C.glPopAttrib();
}

//void glPopClientAttrib (void)
func PopClientAttrib() 
{
	C.glPopClientAttrib();
}

//void glPopMatrix (void)
func PopMatrix() 
{
	C.glPopMatrix();
}

//void glPopName (void)
func PopName() 
{
	C.glPopName();
}

//void glPrioritizeTextures (GLsizei n, const GLuint *textures, const GLclampf *priorities)
func PrioritizeTextures(n GLsizei, textures *GLuint, priorities *GLclampf) 
{
	C.glPrioritizeTextures(C.GLsizei(n), (*C.GLuint)(textures), (*C.GLclampf)(priorities));
}

//void glPushAttrib (GLbitfield mask)
func PushAttrib(mask GLbitfield) 
{
	C.glPushAttrib(C.GLbitfield(mask));
}

//void glPushClientAttrib (GLbitfield mask)
func PushClientAttrib(mask GLbitfield) 
{
	C.glPushClientAttrib(C.GLbitfield(mask));
}

//void glPushMatrix (void)
func PushMatrix() 
{
	C.glPushMatrix();
}

//void glPushName (GLuint name)
func PushName(name GLuint) 
{
	C.glPushName(C.GLuint(name));
}

//void glRasterPos2d (GLdouble x, GLdouble y)
func RasterPos2d(x GLdouble, y GLdouble) 
{
	C.glRasterPos2d(C.GLdouble(x), C.GLdouble(y));
}

//void glRasterPos2dv (const GLdouble *v)
func RasterPos2dv(v *GLdouble) 
{
	C.glRasterPos2dv((*C.GLdouble)(v));
}

//void glRasterPos2f (GLfloat x, GLfloat y)
func RasterPos2f(x GLfloat, y GLfloat) 
{
	C.glRasterPos2f(C.GLfloat(x), C.GLfloat(y));
}

//void glRasterPos2fv (const GLfloat *v)
func RasterPos2fv(v *GLfloat) 
{
	C.glRasterPos2fv((*C.GLfloat)(v));
}

//void glRasterPos2i (GLint x, GLint y)
func RasterPos2i(x GLint, y GLint) 
{
	C.glRasterPos2i(C.GLint(x), C.GLint(y));
}

//void glRasterPos2iv (const GLint *v)
func RasterPos2iv(v *GLint) 
{
	C.glRasterPos2iv((*C.GLint)(v));
}

//void glRasterPos2s (GLshort x, GLshort y)
func RasterPos2s(x GLshort, y GLshort) 
{
	C.glRasterPos2s(C.GLshort(x), C.GLshort(y));
}

//void glRasterPos2sv (const GLshort *v)
func RasterPos2sv(v *GLshort) 
{
	C.glRasterPos2sv((*C.GLshort)(v));
}

//void glRasterPos3d (GLdouble x, GLdouble y, GLdouble z)
func RasterPos3d(x GLdouble, y GLdouble, z GLdouble) 
{
	C.glRasterPos3d(C.GLdouble(x), C.GLdouble(y), C.GLdouble(z));
}

//void glRasterPos3dv (const GLdouble *v)
func RasterPos3dv(v *GLdouble) 
{
	C.glRasterPos3dv((*C.GLdouble)(v));
}

//void glRasterPos3f (GLfloat x, GLfloat y, GLfloat z)
func RasterPos3f(x GLfloat, y GLfloat, z GLfloat) 
{
	C.glRasterPos3f(C.GLfloat(x), C.GLfloat(y), C.GLfloat(z));
}

//void glRasterPos3fv (const GLfloat *v)
func RasterPos3fv(v *GLfloat) 
{
	C.glRasterPos3fv((*C.GLfloat)(v));
}

//void glRasterPos3i (GLint x, GLint y, GLint z)
func RasterPos3i(x GLint, y GLint, z GLint) 
{
	C.glRasterPos3i(C.GLint(x), C.GLint(y), C.GLint(z));
}

//void glRasterPos3iv (const GLint *v)
func RasterPos3iv(v *GLint) 
{
	C.glRasterPos3iv((*C.GLint)(v));
}

//void glRasterPos3s (GLshort x, GLshort y, GLshort z)
func RasterPos3s(x GLshort, y GLshort, z GLshort) 
{
	C.glRasterPos3s(C.GLshort(x), C.GLshort(y), C.GLshort(z));
}

//void glRasterPos3sv (const GLshort *v)
func RasterPos3sv(v *GLshort) 
{
	C.glRasterPos3sv((*C.GLshort)(v));
}

//void glRasterPos4d (GLdouble x, GLdouble y, GLdouble z, GLdouble w)
func RasterPos4d(x GLdouble, y GLdouble, z GLdouble, w GLdouble) 
{
	C.glRasterPos4d(C.GLdouble(x), C.GLdouble(y), C.GLdouble(z), C.GLdouble(w));
}

//void glRasterPos4dv (const GLdouble *v)
func RasterPos4dv(v *GLdouble) 
{
	C.glRasterPos4dv((*C.GLdouble)(v));
}

//void glRasterPos4f (GLfloat x, GLfloat y, GLfloat z, GLfloat w)
func RasterPos4f(x GLfloat, y GLfloat, z GLfloat, w GLfloat) 
{
	C.glRasterPos4f(C.GLfloat(x), C.GLfloat(y), C.GLfloat(z), C.GLfloat(w));
}

//void glRasterPos4fv (const GLfloat *v)
func RasterPos4fv(v *GLfloat) 
{
	C.glRasterPos4fv((*C.GLfloat)(v));
}

//void glRasterPos4i (GLint x, GLint y, GLint z, GLint w)
func RasterPos4i(x GLint, y GLint, z GLint, w GLint) 
{
	C.glRasterPos4i(C.GLint(x), C.GLint(y), C.GLint(z), C.GLint(w));
}

//void glRasterPos4iv (const GLint *v)
func RasterPos4iv(v *GLint) 
{
	C.glRasterPos4iv((*C.GLint)(v));
}

//void glRasterPos4s (GLshort x, GLshort y, GLshort z, GLshort w)
func RasterPos4s(x GLshort, y GLshort, z GLshort, w GLshort) 
{
	C.glRasterPos4s(C.GLshort(x), C.GLshort(y), C.GLshort(z), C.GLshort(w));
}

//void glRasterPos4sv (const GLshort *v)
func RasterPos4sv(v *GLshort) 
{
	C.glRasterPos4sv((*C.GLshort)(v));
}

//void glReadBuffer (GLenum mode)
func ReadBuffer(mode GLenum) 
{
	C.glReadBuffer(C.GLenum(mode));
}

//void glReadPixels (GLint x, GLint y, GLsizei width, GLsizei height, GLenum format, GLenum type, GLvoid *pixels)
func ReadPixels(x GLint, y GLint, width GLsizei, height GLsizei, format GLenum, type_ GLenum, pixels unsafe.Pointer) 
{
	C.glReadPixels(C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height), C.GLenum(format), C.GLenum(type_), pixels);
}

//void glRectd (GLdouble x1, GLdouble y1, GLdouble x2, GLdouble y2)
func Rectd(x1 GLdouble, y1 GLdouble, x2 GLdouble, y2 GLdouble) 
{
	C.glRectd(C.GLdouble(x1), C.GLdouble(y1), C.GLdouble(x2), C.GLdouble(y2));
}

//void glRectdv (const GLdouble *v1, const GLdouble *v2)
func Rectdv(v1 *GLdouble, v2 *GLdouble) 
{
	C.glRectdv((*C.GLdouble)(v1), (*C.GLdouble)(v2));
}

//void glRectf (GLfloat x1, GLfloat y1, GLfloat x2, GLfloat y2)
func Rectf(x1 GLfloat, y1 GLfloat, x2 GLfloat, y2 GLfloat) 
{
	C.glRectf(C.GLfloat(x1), C.GLfloat(y1), C.GLfloat(x2), C.GLfloat(y2));
}

//void glRectfv (const GLfloat *v1, const GLfloat *v2)
func Rectfv(v1 *GLfloat, v2 *GLfloat) 
{
	C.glRectfv((*C.GLfloat)(v1), (*C.GLfloat)(v2));
}

//void glRecti (GLint x1, GLint y1, GLint x2, GLint y2)
func Recti(x1 GLint, y1 GLint, x2 GLint, y2 GLint) 
{
	C.glRecti(C.GLint(x1), C.GLint(y1), C.GLint(x2), C.GLint(y2));
}

//void glRectiv (const GLint *v1, const GLint *v2)
func Rectiv(v1 *GLint, v2 *GLint) 
{
	C.glRectiv((*C.GLint)(v1), (*C.GLint)(v2));
}

//void glRects (GLshort x1, GLshort y1, GLshort x2, GLshort y2)
func Rects(x1 GLshort, y1 GLshort, x2 GLshort, y2 GLshort) 
{
	C.glRects(C.GLshort(x1), C.GLshort(y1), C.GLshort(x2), C.GLshort(y2));
}

//void glRectsv (const GLshort *v1, const GLshort *v2)
func Rectsv(v1 *GLshort, v2 *GLshort) 
{
	C.glRectsv((*C.GLshort)(v1), (*C.GLshort)(v2));
}

//GLint glRenderMode (GLenum mode)
func RenderMode(mode GLenum) GLint
{
	return GLint(C.glRenderMode(C.GLenum(mode)));
}

//void glRotated (GLdouble angle, GLdouble x, GLdouble y, GLdouble z)
func Rotated(angle GLdouble, x GLdouble, y GLdouble, z GLdouble) 
{
	C.glRotated(C.GLdouble(angle), C.GLdouble(x), C.GLdouble(y), C.GLdouble(z));
}

//void glRotatef (GLfloat angle, GLfloat x, GLfloat y, GLfloat z)
func Rotatef(angle GLfloat, x GLfloat, y GLfloat, z GLfloat) 
{
	C.glRotatef(C.GLfloat(angle), C.GLfloat(x), C.GLfloat(y), C.GLfloat(z));
}

//void glScaled (GLdouble x, GLdouble y, GLdouble z)
func Scaled(x GLdouble, y GLdouble, z GLdouble) 
{
	C.glScaled(C.GLdouble(x), C.GLdouble(y), C.GLdouble(z));
}

//void glScalef (GLfloat x, GLfloat y, GLfloat z)
func Scalef(x GLfloat, y GLfloat, z GLfloat) 
{
	C.glScalef(C.GLfloat(x), C.GLfloat(y), C.GLfloat(z));
}

//void glScissor (GLint x, GLint y, GLsizei width, GLsizei height)
func Scissor(x GLint, y GLint, width GLsizei, height GLsizei) 
{
	C.glScissor(C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height));
}

//void glSelectBuffer (GLsizei size, GLuint *buffer)
func SelectBuffer(size GLsizei, buffer *GLuint) 
{
	C.glSelectBuffer(C.GLsizei(size), (*C.GLuint)(buffer));
}

//void glShadeModel (GLenum mode)
func ShadeModel(mode GLenum) 
{
	C.glShadeModel(C.GLenum(mode));
}

//void glStencilFunc (GLenum func, GLint ref, GLuint mask)
func StencilFunc(func_ GLenum, ref GLint, mask GLuint) 
{
	C.glStencilFunc(C.GLenum(func_), C.GLint(ref), C.GLuint(mask));
}

//void glStencilMask (GLuint mask)
func StencilMask(mask GLuint) 
{
	C.glStencilMask(C.GLuint(mask));
}

//void glStencilOp (GLenum fail, GLenum zfail, GLenum zpass)
func StencilOp(fail GLenum, zfail GLenum, zpass GLenum) 
{
	C.glStencilOp(C.GLenum(fail), C.GLenum(zfail), C.GLenum(zpass));
}

//void glTexCoord1d (GLdouble s)
func TexCoord1d(s GLdouble) 
{
	C.glTexCoord1d(C.GLdouble(s));
}

//void glTexCoord1dv (const GLdouble *v)
func TexCoord1dv(v *GLdouble) 
{
	C.glTexCoord1dv((*C.GLdouble)(v));
}

//void glTexCoord1f (GLfloat s)
func TexCoord1f(s GLfloat) 
{
	C.glTexCoord1f(C.GLfloat(s));
}

//void glTexCoord1fv (const GLfloat *v)
func TexCoord1fv(v *GLfloat) 
{
	C.glTexCoord1fv((*C.GLfloat)(v));
}

//void glTexCoord1i (GLint s)
func TexCoord1i(s GLint) 
{
	C.glTexCoord1i(C.GLint(s));
}

//void glTexCoord1iv (const GLint *v)
func TexCoord1iv(v *GLint) 
{
	C.glTexCoord1iv((*C.GLint)(v));
}

//void glTexCoord1s (GLshort s)
func TexCoord1s(s GLshort) 
{
	C.glTexCoord1s(C.GLshort(s));
}

//void glTexCoord1sv (const GLshort *v)
func TexCoord1sv(v *GLshort) 
{
	C.glTexCoord1sv((*C.GLshort)(v));
}

//void glTexCoord2d (GLdouble s, GLdouble t)
func TexCoord2d(s GLdouble, t GLdouble) 
{
	C.glTexCoord2d(C.GLdouble(s), C.GLdouble(t));
}

//void glTexCoord2dv (const GLdouble *v)
func TexCoord2dv(v *GLdouble) 
{
	C.glTexCoord2dv((*C.GLdouble)(v));
}

//void glTexCoord2f (GLfloat s, GLfloat t)
func TexCoord2f(s GLfloat, t GLfloat) 
{
	C.glTexCoord2f(C.GLfloat(s), C.GLfloat(t));
}

//void glTexCoord2fv (const GLfloat *v)
func TexCoord2fv(v *GLfloat) 
{
	C.glTexCoord2fv((*C.GLfloat)(v));
}

//void glTexCoord2i (GLint s, GLint t)
func TexCoord2i(s GLint, t GLint) 
{
	C.glTexCoord2i(C.GLint(s), C.GLint(t));
}

//void glTexCoord2iv (const GLint *v)
func TexCoord2iv(v *GLint) 
{
	C.glTexCoord2iv((*C.GLint)(v));
}

//void glTexCoord2s (GLshort s, GLshort t)
func TexCoord2s(s GLshort, t GLshort) 
{
	C.glTexCoord2s(C.GLshort(s), C.GLshort(t));
}

//void glTexCoord2sv (const GLshort *v)
func TexCoord2sv(v *GLshort) 
{
	C.glTexCoord2sv((*C.GLshort)(v));
}

//void glTexCoord3d (GLdouble s, GLdouble t, GLdouble r)
func TexCoord3d(s GLdouble, t GLdouble, r GLdouble) 
{
	C.glTexCoord3d(C.GLdouble(s), C.GLdouble(t), C.GLdouble(r));
}

//void glTexCoord3dv (const GLdouble *v)
func TexCoord3dv(v *GLdouble) 
{
	C.glTexCoord3dv((*C.GLdouble)(v));
}

//void glTexCoord3f (GLfloat s, GLfloat t, GLfloat r)
func TexCoord3f(s GLfloat, t GLfloat, r GLfloat) 
{
	C.glTexCoord3f(C.GLfloat(s), C.GLfloat(t), C.GLfloat(r));
}

//void glTexCoord3fv (const GLfloat *v)
func TexCoord3fv(v *GLfloat) 
{
	C.glTexCoord3fv((*C.GLfloat)(v));
}

//void glTexCoord3i (GLint s, GLint t, GLint r)
func TexCoord3i(s GLint, t GLint, r GLint) 
{
	C.glTexCoord3i(C.GLint(s), C.GLint(t), C.GLint(r));
}

//void glTexCoord3iv (const GLint *v)
func TexCoord3iv(v *GLint) 
{
	C.glTexCoord3iv((*C.GLint)(v));
}

//void glTexCoord3s (GLshort s, GLshort t, GLshort r)
func TexCoord3s(s GLshort, t GLshort, r GLshort) 
{
	C.glTexCoord3s(C.GLshort(s), C.GLshort(t), C.GLshort(r));
}

//void glTexCoord3sv (const GLshort *v)
func TexCoord3sv(v *GLshort) 
{
	C.glTexCoord3sv((*C.GLshort)(v));
}

//void glTexCoord4d (GLdouble s, GLdouble t, GLdouble r, GLdouble q)
func TexCoord4d(s GLdouble, t GLdouble, r GLdouble, q GLdouble) 
{
	C.glTexCoord4d(C.GLdouble(s), C.GLdouble(t), C.GLdouble(r), C.GLdouble(q));
}

//void glTexCoord4dv (const GLdouble *v)
func TexCoord4dv(v *GLdouble) 
{
	C.glTexCoord4dv((*C.GLdouble)(v));
}

//void glTexCoord4f (GLfloat s, GLfloat t, GLfloat r, GLfloat q)
func TexCoord4f(s GLfloat, t GLfloat, r GLfloat, q GLfloat) 
{
	C.glTexCoord4f(C.GLfloat(s), C.GLfloat(t), C.GLfloat(r), C.GLfloat(q));
}

//void glTexCoord4fv (const GLfloat *v)
func TexCoord4fv(v *GLfloat) 
{
	C.glTexCoord4fv((*C.GLfloat)(v));
}

//void glTexCoord4i (GLint s, GLint t, GLint r, GLint q)
func TexCoord4i(s GLint, t GLint, r GLint, q GLint) 
{
	C.glTexCoord4i(C.GLint(s), C.GLint(t), C.GLint(r), C.GLint(q));
}

//void glTexCoord4iv (const GLint *v)
func TexCoord4iv(v *GLint) 
{
	C.glTexCoord4iv((*C.GLint)(v));
}

//void glTexCoord4s (GLshort s, GLshort t, GLshort r, GLshort q)
func TexCoord4s(s GLshort, t GLshort, r GLshort, q GLshort) 
{
	C.glTexCoord4s(C.GLshort(s), C.GLshort(t), C.GLshort(r), C.GLshort(q));
}

//void glTexCoord4sv (const GLshort *v)
func TexCoord4sv(v *GLshort) 
{
	C.glTexCoord4sv((*C.GLshort)(v));
}

//void glTexCoordPointer (GLint size, GLenum type, GLsizei stride, const GLvoid *pointer)
func TexCoordPointer(size GLint, type_ GLenum, stride GLsizei, pointer unsafe.Pointer) 
{
	C.glTexCoordPointer(C.GLint(size), C.GLenum(type_), C.GLsizei(stride), pointer);
}

//void glTexEnvf (GLenum target, GLenum pname, GLfloat param)
func TexEnvf(target GLenum, pname GLenum, param GLfloat) 
{
	C.glTexEnvf(C.GLenum(target), C.GLenum(pname), C.GLfloat(param));
}

//void glTexEnvfv (GLenum target, GLenum pname, const GLfloat *params)
func TexEnvfv(target GLenum, pname GLenum, params *GLfloat) 
{
	C.glTexEnvfv(C.GLenum(target), C.GLenum(pname), (*C.GLfloat)(params));
}

//void glTexEnvi (GLenum target, GLenum pname, GLint param)
func TexEnvi(target GLenum, pname GLenum, param GLint) 
{
	C.glTexEnvi(C.GLenum(target), C.GLenum(pname), C.GLint(param));
}

//void glTexEnviv (GLenum target, GLenum pname, const GLint *params)
func TexEnviv(target GLenum, pname GLenum, params *GLint) 
{
	C.glTexEnviv(C.GLenum(target), C.GLenum(pname), (*C.GLint)(params));
}

//void glTexGend (GLenum coord, GLenum pname, GLdouble param)
func TexGend(coord GLenum, pname GLenum, param GLdouble) 
{
	C.glTexGend(C.GLenum(coord), C.GLenum(pname), C.GLdouble(param));
}

//void glTexGendv (GLenum coord, GLenum pname, const GLdouble *params)
func TexGendv(coord GLenum, pname GLenum, params *GLdouble) 
{
	C.glTexGendv(C.GLenum(coord), C.GLenum(pname), (*C.GLdouble)(params));
}

//void glTexGenf (GLenum coord, GLenum pname, GLfloat param)
func TexGenf(coord GLenum, pname GLenum, param GLfloat) 
{
	C.glTexGenf(C.GLenum(coord), C.GLenum(pname), C.GLfloat(param));
}

//void glTexGenfv (GLenum coord, GLenum pname, const GLfloat *params)
func TexGenfv(coord GLenum, pname GLenum, params *GLfloat) 
{
	C.glTexGenfv(C.GLenum(coord), C.GLenum(pname), (*C.GLfloat)(params));
}

//void glTexGeni (GLenum coord, GLenum pname, GLint param)
func TexGeni(coord GLenum, pname GLenum, param GLint) 
{
	C.glTexGeni(C.GLenum(coord), C.GLenum(pname), C.GLint(param));
}

//void glTexGeniv (GLenum coord, GLenum pname, const GLint *params)
func TexGeniv(coord GLenum, pname GLenum, params *GLint) 
{
	C.glTexGeniv(C.GLenum(coord), C.GLenum(pname), (*C.GLint)(params));
}

//void glTexImage1D (GLenum target, GLint level, GLint internalformat, GLsizei width, GLint border, GLenum format, GLenum type, const GLvoid *pixels)
func TexImage1D(target GLenum, level GLint, internalformat GLint, width GLsizei, border GLint, format GLenum, type_ GLenum, pixels unsafe.Pointer) 
{
	C.glTexImage1D(C.GLenum(target), C.GLint(level), C.GLint(internalformat), C.GLsizei(width), C.GLint(border), C.GLenum(format), C.GLenum(type_), pixels);
}

//void glTexImage2D (GLenum target, GLint level, GLint internalformat, GLsizei width, GLsizei height, GLint border, GLenum format, GLenum type, const GLvoid *pixels)
func TexImage2D(target GLenum, level GLint, internalformat GLint, width GLsizei, height GLsizei, border GLint, format GLenum, type_ GLenum, pixels unsafe.Pointer) 
{
	C.glTexImage2D(C.GLenum(target), C.GLint(level), C.GLint(internalformat), C.GLsizei(width), C.GLsizei(height), C.GLint(border), C.GLenum(format), C.GLenum(type_), pixels);
}

//void glTexParameterf (GLenum target, GLenum pname, GLfloat param)
func TexParameterf(target GLenum, pname GLenum, param GLfloat) 
{
	C.glTexParameterf(C.GLenum(target), C.GLenum(pname), C.GLfloat(param));
}

//void glTexParameterfv (GLenum target, GLenum pname, const GLfloat *params)
func TexParameterfv(target GLenum, pname GLenum, params *GLfloat) 
{
	C.glTexParameterfv(C.GLenum(target), C.GLenum(pname), (*C.GLfloat)(params));
}

//void glTexParameteri (GLenum target, GLenum pname, GLint param)
func TexParameteri(target GLenum, pname GLenum, param GLint) 
{
	C.glTexParameteri(C.GLenum(target), C.GLenum(pname), C.GLint(param));
}

//void glTexParameteriv (GLenum target, GLenum pname, const GLint *params)
func TexParameteriv(target GLenum, pname GLenum, params *GLint) 
{
	C.glTexParameteriv(C.GLenum(target), C.GLenum(pname), (*C.GLint)(params));
}

//void glTexSubImage1D (GLenum target, GLint level, GLint xoffset, GLsizei width, GLenum format, GLenum type, const GLvoid *pixels)
func TexSubImage1D(target GLenum, level GLint, xoffset GLint, width GLsizei, format GLenum, type_ GLenum, pixels unsafe.Pointer) 
{
	C.glTexSubImage1D(C.GLenum(target), C.GLint(level), C.GLint(xoffset), C.GLsizei(width), C.GLenum(format), C.GLenum(type_), pixels);
}

//void glTexSubImage2D (GLenum target, GLint level, GLint xoffset, GLint yoffset, GLsizei width, GLsizei height, GLenum format, GLenum type, const GLvoid *pixels)
func TexSubImage2D(target GLenum, level GLint, xoffset GLint, yoffset GLint, width GLsizei, height GLsizei, format GLenum, type_ GLenum, pixels unsafe.Pointer) 
{
	C.glTexSubImage2D(C.GLenum(target), C.GLint(level), C.GLint(xoffset), C.GLint(yoffset), C.GLsizei(width), C.GLsizei(height), C.GLenum(format), C.GLenum(type_), pixels);
}

//void glTranslated (GLdouble x, GLdouble y, GLdouble z)
func Translated(x GLdouble, y GLdouble, z GLdouble) 
{
	C.glTranslated(C.GLdouble(x), C.GLdouble(y), C.GLdouble(z));
}

//void glTranslatef (GLfloat x, GLfloat y, GLfloat z)
func Translatef(x GLfloat, y GLfloat, z GLfloat) 
{
	C.glTranslatef(C.GLfloat(x), C.GLfloat(y), C.GLfloat(z));
}

//void glVertex2d (GLdouble x, GLdouble y)
func Vertex2d(x GLdouble, y GLdouble) 
{
	C.glVertex2d(C.GLdouble(x), C.GLdouble(y));
}

//void glVertex2dv (const GLdouble *v)
func Vertex2dv(v *GLdouble) 
{
	C.glVertex2dv((*C.GLdouble)(v));
}

//void glVertex2f (GLfloat x, GLfloat y)
func Vertex2f(x GLfloat, y GLfloat) 
{
	C.glVertex2f(C.GLfloat(x), C.GLfloat(y));
}

//void glVertex2fv (const GLfloat *v)
func Vertex2fv(v *GLfloat) 
{
	C.glVertex2fv((*C.GLfloat)(v));
}

//void glVertex2i (GLint x, GLint y)
func Vertex2i(x GLint, y GLint) 
{
	C.glVertex2i(C.GLint(x), C.GLint(y));
}

//void glVertex2iv (const GLint *v)
func Vertex2iv(v *GLint) 
{
	C.glVertex2iv((*C.GLint)(v));
}

//void glVertex2s (GLshort x, GLshort y)
func Vertex2s(x GLshort, y GLshort) 
{
	C.glVertex2s(C.GLshort(x), C.GLshort(y));
}

//void glVertex2sv (const GLshort *v)
func Vertex2sv(v *GLshort) 
{
	C.glVertex2sv((*C.GLshort)(v));
}

//void glVertex3d (GLdouble x, GLdouble y, GLdouble z)
func Vertex3d(x GLdouble, y GLdouble, z GLdouble) 
{
	C.glVertex3d(C.GLdouble(x), C.GLdouble(y), C.GLdouble(z));
}

//void glVertex3dv (const GLdouble *v)
func Vertex3dv(v *GLdouble) 
{
	C.glVertex3dv((*C.GLdouble)(v));
}

//void glVertex3f (GLfloat x, GLfloat y, GLfloat z)
func Vertex3f(x GLfloat, y GLfloat, z GLfloat) 
{
	C.glVertex3f(C.GLfloat(x), C.GLfloat(y), C.GLfloat(z));
}

//void glVertex3fv (const GLfloat *v)
func Vertex3fv(v *GLfloat) 
{
	C.glVertex3fv((*C.GLfloat)(v));
}

//void glVertex3i (GLint x, GLint y, GLint z)
func Vertex3i(x GLint, y GLint, z GLint) 
{
	C.glVertex3i(C.GLint(x), C.GLint(y), C.GLint(z));
}

//void glVertex3iv (const GLint *v)
func Vertex3iv(v *GLint) 
{
	C.glVertex3iv((*C.GLint)(v));
}

//void glVertex3s (GLshort x, GLshort y, GLshort z)
func Vertex3s(x GLshort, y GLshort, z GLshort) 
{
	C.glVertex3s(C.GLshort(x), C.GLshort(y), C.GLshort(z));
}

//void glVertex3sv (const GLshort *v)
func Vertex3sv(v *GLshort) 
{
	C.glVertex3sv((*C.GLshort)(v));
}

//void glVertex4d (GLdouble x, GLdouble y, GLdouble z, GLdouble w)
func Vertex4d(x GLdouble, y GLdouble, z GLdouble, w GLdouble) 
{
	C.glVertex4d(C.GLdouble(x), C.GLdouble(y), C.GLdouble(z), C.GLdouble(w));
}

//void glVertex4dv (const GLdouble *v)
func Vertex4dv(v *GLdouble) 
{
	C.glVertex4dv((*C.GLdouble)(v));
}

//void glVertex4f (GLfloat x, GLfloat y, GLfloat z, GLfloat w)
func Vertex4f(x GLfloat, y GLfloat, z GLfloat, w GLfloat) 
{
	C.glVertex4f(C.GLfloat(x), C.GLfloat(y), C.GLfloat(z), C.GLfloat(w));
}

//void glVertex4fv (const GLfloat *v)
func Vertex4fv(v *GLfloat) 
{
	C.glVertex4fv((*C.GLfloat)(v));
}

//void glVertex4i (GLint x, GLint y, GLint z, GLint w)
func Vertex4i(x GLint, y GLint, z GLint, w GLint) 
{
	C.glVertex4i(C.GLint(x), C.GLint(y), C.GLint(z), C.GLint(w));
}

//void glVertex4iv (const GLint *v)
func Vertex4iv(v *GLint) 
{
	C.glVertex4iv((*C.GLint)(v));
}

//void glVertex4s (GLshort x, GLshort y, GLshort z, GLshort w)
func Vertex4s(x GLshort, y GLshort, z GLshort, w GLshort) 
{
	C.glVertex4s(C.GLshort(x), C.GLshort(y), C.GLshort(z), C.GLshort(w));
}

//void glVertex4sv (const GLshort *v)
func Vertex4sv(v *GLshort) 
{
	C.glVertex4sv((*C.GLshort)(v));
}

//void glVertexPointer (GLint size, GLenum type, GLsizei stride, const GLvoid *pointer)
func VertexPointer(size GLint, type_ GLenum, stride GLsizei, pointer unsafe.Pointer) 
{
	C.glVertexPointer(C.GLint(size), C.GLenum(type_), C.GLsizei(stride), pointer);
}

//void glViewport (GLint x, GLint y, GLsizei width, GLsizei height)
func Viewport(x GLint, y GLint, width GLsizei, height GLsizei) 
{
	C.glViewport(C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height));
}

type ptrdiff_t C.ptrdiff_t;
type size_t C.size_t;
type wchar_t C.wchar_t;
type GLchar C.GLchar;
type GLintptr C.GLintptr;
type GLsizeiptr C.GLsizeiptr;
type GLintptrARB C.GLintptrARB;
type GLsizeiptrARB C.GLsizeiptrARB;
type GLcharARB C.GLcharARB;
type GLhandleARB C.GLhandleARB;
type GLhalfARB C.GLhalfARB;
type GLhalfNV C.GLhalfNV;
type GLint64EXT C.GLint64EXT;
type GLuint64EXT C.GLuint64EXT;
//extern void glBlendColor (GLclampf, GLclampf, GLclampf, GLclampf)
func BlendColor(arg0 GLclampf, arg1 GLclampf, arg2 GLclampf, arg3 GLclampf) 
{
	C.glBlendColor(C.GLclampf(arg0), C.GLclampf(arg1), C.GLclampf(arg2), C.GLclampf(arg3));
}

//extern void glBlendEquation (GLenum)
func BlendEquation(arg0 GLenum) 
{
	C.glBlendEquation(C.GLenum(arg0));
}

//extern void glDrawRangeElements (GLenum, GLuint, GLuint, GLsizei, GLenum, const GLvoid *)
func DrawRangeElements(arg0 GLenum, arg1 GLuint, arg2 GLuint, arg3 GLsizei, arg4 GLenum, arg5 unsafe.Pointer) 
{
	C.glDrawRangeElements(C.GLenum(arg0), C.GLuint(arg1), C.GLuint(arg2), C.GLsizei(arg3), C.GLenum(arg4), arg5);
}

//extern void glColorTable (GLenum, GLenum, GLsizei, GLenum, GLenum, const GLvoid *)
func ColorTable(arg0 GLenum, arg1 GLenum, arg2 GLsizei, arg3 GLenum, arg4 GLenum, arg5 unsafe.Pointer) 
{
	C.glColorTable(C.GLenum(arg0), C.GLenum(arg1), C.GLsizei(arg2), C.GLenum(arg3), C.GLenum(arg4), arg5);
}

//extern void glColorTableParameterfv (GLenum, GLenum, const GLfloat *)
func ColorTableParameterfv(arg0 GLenum, arg1 GLenum, arg2 *GLfloat) 
{
	C.glColorTableParameterfv(C.GLenum(arg0), C.GLenum(arg1), (*C.GLfloat)(arg2));
}

//extern void glColorTableParameteriv (GLenum, GLenum, const GLint *)
func ColorTableParameteriv(arg0 GLenum, arg1 GLenum, arg2 *GLint) 
{
	C.glColorTableParameteriv(C.GLenum(arg0), C.GLenum(arg1), (*C.GLint)(arg2));
}

//extern void glCopyColorTable (GLenum, GLenum, GLint, GLint, GLsizei)
func CopyColorTable(arg0 GLenum, arg1 GLenum, arg2 GLint, arg3 GLint, arg4 GLsizei) 
{
	C.glCopyColorTable(C.GLenum(arg0), C.GLenum(arg1), C.GLint(arg2), C.GLint(arg3), C.GLsizei(arg4));
}

//extern void glGetColorTable (GLenum, GLenum, GLenum, GLvoid *)
func GetColorTable(arg0 GLenum, arg1 GLenum, arg2 GLenum, arg3 unsafe.Pointer) 
{
	C.glGetColorTable(C.GLenum(arg0), C.GLenum(arg1), C.GLenum(arg2), arg3);
}

//extern void glGetColorTableParameterfv (GLenum, GLenum, GLfloat *)
func GetColorTableParameterfv(arg0 GLenum, arg1 GLenum, arg2 *GLfloat) 
{
	C.glGetColorTableParameterfv(C.GLenum(arg0), C.GLenum(arg1), (*C.GLfloat)(arg2));
}

//extern void glGetColorTableParameteriv (GLenum, GLenum, GLint *)
func GetColorTableParameteriv(arg0 GLenum, arg1 GLenum, arg2 *GLint) 
{
	C.glGetColorTableParameteriv(C.GLenum(arg0), C.GLenum(arg1), (*C.GLint)(arg2));
}

//extern void glColorSubTable (GLenum, GLsizei, GLsizei, GLenum, GLenum, const GLvoid *)
func ColorSubTable(arg0 GLenum, arg1 GLsizei, arg2 GLsizei, arg3 GLenum, arg4 GLenum, arg5 unsafe.Pointer) 
{
	C.glColorSubTable(C.GLenum(arg0), C.GLsizei(arg1), C.GLsizei(arg2), C.GLenum(arg3), C.GLenum(arg4), arg5);
}

//extern void glCopyColorSubTable (GLenum, GLsizei, GLint, GLint, GLsizei)
func CopyColorSubTable(arg0 GLenum, arg1 GLsizei, arg2 GLint, arg3 GLint, arg4 GLsizei) 
{
	C.glCopyColorSubTable(C.GLenum(arg0), C.GLsizei(arg1), C.GLint(arg2), C.GLint(arg3), C.GLsizei(arg4));
}

//extern void glConvolutionFilter1D (GLenum, GLenum, GLsizei, GLenum, GLenum, const GLvoid *)
func ConvolutionFilter1D(arg0 GLenum, arg1 GLenum, arg2 GLsizei, arg3 GLenum, arg4 GLenum, arg5 unsafe.Pointer) 
{
	C.glConvolutionFilter1D(C.GLenum(arg0), C.GLenum(arg1), C.GLsizei(arg2), C.GLenum(arg3), C.GLenum(arg4), arg5);
}

//extern void glConvolutionFilter2D (GLenum, GLenum, GLsizei, GLsizei, GLenum, GLenum, const GLvoid *)
func ConvolutionFilter2D(arg0 GLenum, arg1 GLenum, arg2 GLsizei, arg3 GLsizei, arg4 GLenum, arg5 GLenum, arg6 unsafe.Pointer) 
{
	C.glConvolutionFilter2D(C.GLenum(arg0), C.GLenum(arg1), C.GLsizei(arg2), C.GLsizei(arg3), C.GLenum(arg4), C.GLenum(arg5), arg6);
}

//extern void glConvolutionParameterf (GLenum, GLenum, GLfloat)
func ConvolutionParameterf(arg0 GLenum, arg1 GLenum, arg2 GLfloat) 
{
	C.glConvolutionParameterf(C.GLenum(arg0), C.GLenum(arg1), C.GLfloat(arg2));
}

//extern void glConvolutionParameterfv (GLenum, GLenum, const GLfloat *)
func ConvolutionParameterfv(arg0 GLenum, arg1 GLenum, arg2 *GLfloat) 
{
	C.glConvolutionParameterfv(C.GLenum(arg0), C.GLenum(arg1), (*C.GLfloat)(arg2));
}

//extern void glConvolutionParameteri (GLenum, GLenum, GLint)
func ConvolutionParameteri(arg0 GLenum, arg1 GLenum, arg2 GLint) 
{
	C.glConvolutionParameteri(C.GLenum(arg0), C.GLenum(arg1), C.GLint(arg2));
}

//extern void glConvolutionParameteriv (GLenum, GLenum, const GLint *)
func ConvolutionParameteriv(arg0 GLenum, arg1 GLenum, arg2 *GLint) 
{
	C.glConvolutionParameteriv(C.GLenum(arg0), C.GLenum(arg1), (*C.GLint)(arg2));
}

//extern void glCopyConvolutionFilter1D (GLenum, GLenum, GLint, GLint, GLsizei)
func CopyConvolutionFilter1D(arg0 GLenum, arg1 GLenum, arg2 GLint, arg3 GLint, arg4 GLsizei) 
{
	C.glCopyConvolutionFilter1D(C.GLenum(arg0), C.GLenum(arg1), C.GLint(arg2), C.GLint(arg3), C.GLsizei(arg4));
}

//extern void glCopyConvolutionFilter2D (GLenum, GLenum, GLint, GLint, GLsizei, GLsizei)
func CopyConvolutionFilter2D(arg0 GLenum, arg1 GLenum, arg2 GLint, arg3 GLint, arg4 GLsizei, arg5 GLsizei) 
{
	C.glCopyConvolutionFilter2D(C.GLenum(arg0), C.GLenum(arg1), C.GLint(arg2), C.GLint(arg3), C.GLsizei(arg4), C.GLsizei(arg5));
}

//extern void glGetConvolutionFilter (GLenum, GLenum, GLenum, GLvoid *)
func GetConvolutionFilter(arg0 GLenum, arg1 GLenum, arg2 GLenum, arg3 unsafe.Pointer) 
{
	C.glGetConvolutionFilter(C.GLenum(arg0), C.GLenum(arg1), C.GLenum(arg2), arg3);
}

//extern void glGetConvolutionParameterfv (GLenum, GLenum, GLfloat *)
func GetConvolutionParameterfv(arg0 GLenum, arg1 GLenum, arg2 *GLfloat) 
{
	C.glGetConvolutionParameterfv(C.GLenum(arg0), C.GLenum(arg1), (*C.GLfloat)(arg2));
}

//extern void glGetConvolutionParameteriv (GLenum, GLenum, GLint *)
func GetConvolutionParameteriv(arg0 GLenum, arg1 GLenum, arg2 *GLint) 
{
	C.glGetConvolutionParameteriv(C.GLenum(arg0), C.GLenum(arg1), (*C.GLint)(arg2));
}

//extern void glGetSeparableFilter (GLenum, GLenum, GLenum, GLvoid *, GLvoid *, GLvoid *)
func GetSeparableFilter(arg0 GLenum, arg1 GLenum, arg2 GLenum, arg3 unsafe.Pointer, arg4 unsafe.Pointer, arg5 unsafe.Pointer) 
{
	C.glGetSeparableFilter(C.GLenum(arg0), C.GLenum(arg1), C.GLenum(arg2), arg3, arg4, arg5);
}

//extern void glSeparableFilter2D (GLenum, GLenum, GLsizei, GLsizei, GLenum, GLenum, const GLvoid *, const GLvoid *)
func SeparableFilter2D(arg0 GLenum, arg1 GLenum, arg2 GLsizei, arg3 GLsizei, arg4 GLenum, arg5 GLenum, arg6 unsafe.Pointer, arg7 unsafe.Pointer) 
{
	C.glSeparableFilter2D(C.GLenum(arg0), C.GLenum(arg1), C.GLsizei(arg2), C.GLsizei(arg3), C.GLenum(arg4), C.GLenum(arg5), arg6, arg7);
}

//extern void glGetHistogram (GLenum, GLboolean, GLenum, GLenum, GLvoid *)
func GetHistogram(arg0 GLenum, arg1 GLboolean, arg2 GLenum, arg3 GLenum, arg4 unsafe.Pointer) 
{
	C.glGetHistogram(C.GLenum(arg0), C.GLboolean(arg1), C.GLenum(arg2), C.GLenum(arg3), arg4);
}

//extern void glGetHistogramParameterfv (GLenum, GLenum, GLfloat *)
func GetHistogramParameterfv(arg0 GLenum, arg1 GLenum, arg2 *GLfloat) 
{
	C.glGetHistogramParameterfv(C.GLenum(arg0), C.GLenum(arg1), (*C.GLfloat)(arg2));
}

//extern void glGetHistogramParameteriv (GLenum, GLenum, GLint *)
func GetHistogramParameteriv(arg0 GLenum, arg1 GLenum, arg2 *GLint) 
{
	C.glGetHistogramParameteriv(C.GLenum(arg0), C.GLenum(arg1), (*C.GLint)(arg2));
}

//extern void glGetMinmax (GLenum, GLboolean, GLenum, GLenum, GLvoid *)
func GetMinmax(arg0 GLenum, arg1 GLboolean, arg2 GLenum, arg3 GLenum, arg4 unsafe.Pointer) 
{
	C.glGetMinmax(C.GLenum(arg0), C.GLboolean(arg1), C.GLenum(arg2), C.GLenum(arg3), arg4);
}

//extern void glGetMinmaxParameterfv (GLenum, GLenum, GLfloat *)
func GetMinmaxParameterfv(arg0 GLenum, arg1 GLenum, arg2 *GLfloat) 
{
	C.glGetMinmaxParameterfv(C.GLenum(arg0), C.GLenum(arg1), (*C.GLfloat)(arg2));
}

//extern void glGetMinmaxParameteriv (GLenum, GLenum, GLint *)
func GetMinmaxParameteriv(arg0 GLenum, arg1 GLenum, arg2 *GLint) 
{
	C.glGetMinmaxParameteriv(C.GLenum(arg0), C.GLenum(arg1), (*C.GLint)(arg2));
}

//extern void glHistogram (GLenum, GLsizei, GLenum, GLboolean)
func Histogram(arg0 GLenum, arg1 GLsizei, arg2 GLenum, arg3 GLboolean) 
{
	C.glHistogram(C.GLenum(arg0), C.GLsizei(arg1), C.GLenum(arg2), C.GLboolean(arg3));
}

//extern void glMinmax (GLenum, GLenum, GLboolean)
func Minmax(arg0 GLenum, arg1 GLenum, arg2 GLboolean) 
{
	C.glMinmax(C.GLenum(arg0), C.GLenum(arg1), C.GLboolean(arg2));
}

//extern void glResetHistogram (GLenum)
func ResetHistogram(arg0 GLenum) 
{
	C.glResetHistogram(C.GLenum(arg0));
}

//extern void glResetMinmax (GLenum)
func ResetMinmax(arg0 GLenum) 
{
	C.glResetMinmax(C.GLenum(arg0));
}

//extern void glTexImage3D (GLenum, GLint, GLint, GLsizei, GLsizei, GLsizei, GLint, GLenum, GLenum, const GLvoid *)
func TexImage3D(arg0 GLenum, arg1 GLint, arg2 GLint, arg3 GLsizei, arg4 GLsizei, arg5 GLsizei, arg6 GLint, arg7 GLenum, arg8 GLenum, arg9 unsafe.Pointer) 
{
	C.glTexImage3D(C.GLenum(arg0), C.GLint(arg1), C.GLint(arg2), C.GLsizei(arg3), C.GLsizei(arg4), C.GLsizei(arg5), C.GLint(arg6), C.GLenum(arg7), C.GLenum(arg8), arg9);
}

//extern void glTexSubImage3D (GLenum, GLint, GLint, GLint, GLint, GLsizei, GLsizei, GLsizei, GLenum, GLenum, const GLvoid *)
func TexSubImage3D(arg0 GLenum, arg1 GLint, arg2 GLint, arg3 GLint, arg4 GLint, arg5 GLsizei, arg6 GLsizei, arg7 GLsizei, arg8 GLenum, arg9 GLenum, arg10 unsafe.Pointer) 
{
	C.glTexSubImage3D(C.GLenum(arg0), C.GLint(arg1), C.GLint(arg2), C.GLint(arg3), C.GLint(arg4), C.GLsizei(arg5), C.GLsizei(arg6), C.GLsizei(arg7), C.GLenum(arg8), C.GLenum(arg9), arg10);
}

//extern void glCopyTexSubImage3D (GLenum, GLint, GLint, GLint, GLint, GLint, GLint, GLsizei, GLsizei)
func CopyTexSubImage3D(arg0 GLenum, arg1 GLint, arg2 GLint, arg3 GLint, arg4 GLint, arg5 GLint, arg6 GLint, arg7 GLsizei, arg8 GLsizei) 
{
	C.glCopyTexSubImage3D(C.GLenum(arg0), C.GLint(arg1), C.GLint(arg2), C.GLint(arg3), C.GLint(arg4), C.GLint(arg5), C.GLint(arg6), C.GLsizei(arg7), C.GLsizei(arg8));
}

//extern void glActiveTexture (GLenum)
func ActiveTexture(arg0 GLenum) 
{
	C.glActiveTexture(C.GLenum(arg0));
}

//extern void glClientActiveTexture (GLenum)
func ClientActiveTexture(arg0 GLenum) 
{
	C.glClientActiveTexture(C.GLenum(arg0));
}

//extern void glMultiTexCoord1d (GLenum, GLdouble)
func MultiTexCoord1d(arg0 GLenum, arg1 GLdouble) 
{
	C.glMultiTexCoord1d(C.GLenum(arg0), C.GLdouble(arg1));
}

//extern void glMultiTexCoord1dv (GLenum, const GLdouble *)
func MultiTexCoord1dv(arg0 GLenum, arg1 *GLdouble) 
{
	C.glMultiTexCoord1dv(C.GLenum(arg0), (*C.GLdouble)(arg1));
}

//extern void glMultiTexCoord1f (GLenum, GLfloat)
func MultiTexCoord1f(arg0 GLenum, arg1 GLfloat) 
{
	C.glMultiTexCoord1f(C.GLenum(arg0), C.GLfloat(arg1));
}

//extern void glMultiTexCoord1fv (GLenum, const GLfloat *)
func MultiTexCoord1fv(arg0 GLenum, arg1 *GLfloat) 
{
	C.glMultiTexCoord1fv(C.GLenum(arg0), (*C.GLfloat)(arg1));
}

//extern void glMultiTexCoord1i (GLenum, GLint)
func MultiTexCoord1i(arg0 GLenum, arg1 GLint) 
{
	C.glMultiTexCoord1i(C.GLenum(arg0), C.GLint(arg1));
}

//extern void glMultiTexCoord1iv (GLenum, const GLint *)
func MultiTexCoord1iv(arg0 GLenum, arg1 *GLint) 
{
	C.glMultiTexCoord1iv(C.GLenum(arg0), (*C.GLint)(arg1));
}

//extern void glMultiTexCoord1s (GLenum, GLshort)
func MultiTexCoord1s(arg0 GLenum, arg1 GLshort) 
{
	C.glMultiTexCoord1s(C.GLenum(arg0), C.GLshort(arg1));
}

//extern void glMultiTexCoord1sv (GLenum, const GLshort *)
func MultiTexCoord1sv(arg0 GLenum, arg1 *GLshort) 
{
	C.glMultiTexCoord1sv(C.GLenum(arg0), (*C.GLshort)(arg1));
}

//extern void glMultiTexCoord2d (GLenum, GLdouble, GLdouble)
func MultiTexCoord2d(arg0 GLenum, arg1 GLdouble, arg2 GLdouble) 
{
	C.glMultiTexCoord2d(C.GLenum(arg0), C.GLdouble(arg1), C.GLdouble(arg2));
}

//extern void glMultiTexCoord2dv (GLenum, const GLdouble *)
func MultiTexCoord2dv(arg0 GLenum, arg1 *GLdouble) 
{
	C.glMultiTexCoord2dv(C.GLenum(arg0), (*C.GLdouble)(arg1));
}

//extern void glMultiTexCoord2f (GLenum, GLfloat, GLfloat)
func MultiTexCoord2f(arg0 GLenum, arg1 GLfloat, arg2 GLfloat) 
{
	C.glMultiTexCoord2f(C.GLenum(arg0), C.GLfloat(arg1), C.GLfloat(arg2));
}

//extern void glMultiTexCoord2fv (GLenum, const GLfloat *)
func MultiTexCoord2fv(arg0 GLenum, arg1 *GLfloat) 
{
	C.glMultiTexCoord2fv(C.GLenum(arg0), (*C.GLfloat)(arg1));
}

//extern void glMultiTexCoord2i (GLenum, GLint, GLint)
func MultiTexCoord2i(arg0 GLenum, arg1 GLint, arg2 GLint) 
{
	C.glMultiTexCoord2i(C.GLenum(arg0), C.GLint(arg1), C.GLint(arg2));
}

//extern void glMultiTexCoord2iv (GLenum, const GLint *)
func MultiTexCoord2iv(arg0 GLenum, arg1 *GLint) 
{
	C.glMultiTexCoord2iv(C.GLenum(arg0), (*C.GLint)(arg1));
}

//extern void glMultiTexCoord2s (GLenum, GLshort, GLshort)
func MultiTexCoord2s(arg0 GLenum, arg1 GLshort, arg2 GLshort) 
{
	C.glMultiTexCoord2s(C.GLenum(arg0), C.GLshort(arg1), C.GLshort(arg2));
}

//extern void glMultiTexCoord2sv (GLenum, const GLshort *)
func MultiTexCoord2sv(arg0 GLenum, arg1 *GLshort) 
{
	C.glMultiTexCoord2sv(C.GLenum(arg0), (*C.GLshort)(arg1));
}

//extern void glMultiTexCoord3d (GLenum, GLdouble, GLdouble, GLdouble)
func MultiTexCoord3d(arg0 GLenum, arg1 GLdouble, arg2 GLdouble, arg3 GLdouble) 
{
	C.glMultiTexCoord3d(C.GLenum(arg0), C.GLdouble(arg1), C.GLdouble(arg2), C.GLdouble(arg3));
}

//extern void glMultiTexCoord3dv (GLenum, const GLdouble *)
func MultiTexCoord3dv(arg0 GLenum, arg1 *GLdouble) 
{
	C.glMultiTexCoord3dv(C.GLenum(arg0), (*C.GLdouble)(arg1));
}

//extern void glMultiTexCoord3f (GLenum, GLfloat, GLfloat, GLfloat)
func MultiTexCoord3f(arg0 GLenum, arg1 GLfloat, arg2 GLfloat, arg3 GLfloat) 
{
	C.glMultiTexCoord3f(C.GLenum(arg0), C.GLfloat(arg1), C.GLfloat(arg2), C.GLfloat(arg3));
}

//extern void glMultiTexCoord3fv (GLenum, const GLfloat *)
func MultiTexCoord3fv(arg0 GLenum, arg1 *GLfloat) 
{
	C.glMultiTexCoord3fv(C.GLenum(arg0), (*C.GLfloat)(arg1));
}

//extern void glMultiTexCoord3i (GLenum, GLint, GLint, GLint)
func MultiTexCoord3i(arg0 GLenum, arg1 GLint, arg2 GLint, arg3 GLint) 
{
	C.glMultiTexCoord3i(C.GLenum(arg0), C.GLint(arg1), C.GLint(arg2), C.GLint(arg3));
}

//extern void glMultiTexCoord3iv (GLenum, const GLint *)
func MultiTexCoord3iv(arg0 GLenum, arg1 *GLint) 
{
	C.glMultiTexCoord3iv(C.GLenum(arg0), (*C.GLint)(arg1));
}

//extern void glMultiTexCoord3s (GLenum, GLshort, GLshort, GLshort)
func MultiTexCoord3s(arg0 GLenum, arg1 GLshort, arg2 GLshort, arg3 GLshort) 
{
	C.glMultiTexCoord3s(C.GLenum(arg0), C.GLshort(arg1), C.GLshort(arg2), C.GLshort(arg3));
}

//extern void glMultiTexCoord3sv (GLenum, const GLshort *)
func MultiTexCoord3sv(arg0 GLenum, arg1 *GLshort) 
{
	C.glMultiTexCoord3sv(C.GLenum(arg0), (*C.GLshort)(arg1));
}

//extern void glMultiTexCoord4d (GLenum, GLdouble, GLdouble, GLdouble, GLdouble)
func MultiTexCoord4d(arg0 GLenum, arg1 GLdouble, arg2 GLdouble, arg3 GLdouble, arg4 GLdouble) 
{
	C.glMultiTexCoord4d(C.GLenum(arg0), C.GLdouble(arg1), C.GLdouble(arg2), C.GLdouble(arg3), C.GLdouble(arg4));
}

//extern void glMultiTexCoord4dv (GLenum, const GLdouble *)
func MultiTexCoord4dv(arg0 GLenum, arg1 *GLdouble) 
{
	C.glMultiTexCoord4dv(C.GLenum(arg0), (*C.GLdouble)(arg1));
}

//extern void glMultiTexCoord4f (GLenum, GLfloat, GLfloat, GLfloat, GLfloat)
func MultiTexCoord4f(arg0 GLenum, arg1 GLfloat, arg2 GLfloat, arg3 GLfloat, arg4 GLfloat) 
{
	C.glMultiTexCoord4f(C.GLenum(arg0), C.GLfloat(arg1), C.GLfloat(arg2), C.GLfloat(arg3), C.GLfloat(arg4));
}

//extern void glMultiTexCoord4fv (GLenum, const GLfloat *)
func MultiTexCoord4fv(arg0 GLenum, arg1 *GLfloat) 
{
	C.glMultiTexCoord4fv(C.GLenum(arg0), (*C.GLfloat)(arg1));
}

//extern void glMultiTexCoord4i (GLenum, GLint, GLint, GLint, GLint)
func MultiTexCoord4i(arg0 GLenum, arg1 GLint, arg2 GLint, arg3 GLint, arg4 GLint) 
{
	C.glMultiTexCoord4i(C.GLenum(arg0), C.GLint(arg1), C.GLint(arg2), C.GLint(arg3), C.GLint(arg4));
}

//extern void glMultiTexCoord4iv (GLenum, const GLint *)
func MultiTexCoord4iv(arg0 GLenum, arg1 *GLint) 
{
	C.glMultiTexCoord4iv(C.GLenum(arg0), (*C.GLint)(arg1));
}

//extern void glMultiTexCoord4s (GLenum, GLshort, GLshort, GLshort, GLshort)
func MultiTexCoord4s(arg0 GLenum, arg1 GLshort, arg2 GLshort, arg3 GLshort, arg4 GLshort) 
{
	C.glMultiTexCoord4s(C.GLenum(arg0), C.GLshort(arg1), C.GLshort(arg2), C.GLshort(arg3), C.GLshort(arg4));
}

//extern void glMultiTexCoord4sv (GLenum, const GLshort *)
func MultiTexCoord4sv(arg0 GLenum, arg1 *GLshort) 
{
	C.glMultiTexCoord4sv(C.GLenum(arg0), (*C.GLshort)(arg1));
}

//extern void glLoadTransposeMatrixf (const GLfloat *)
func LoadTransposeMatrixf(arg0 *GLfloat) 
{
	C.glLoadTransposeMatrixf((*C.GLfloat)(arg0));
}

//extern void glLoadTransposeMatrixd (const GLdouble *)
func LoadTransposeMatrixd(arg0 *GLdouble) 
{
	C.glLoadTransposeMatrixd((*C.GLdouble)(arg0));
}

//extern void glMultTransposeMatrixf (const GLfloat *)
func MultTransposeMatrixf(arg0 *GLfloat) 
{
	C.glMultTransposeMatrixf((*C.GLfloat)(arg0));
}

//extern void glMultTransposeMatrixd (const GLdouble *)
func MultTransposeMatrixd(arg0 *GLdouble) 
{
	C.glMultTransposeMatrixd((*C.GLdouble)(arg0));
}

//extern void glSampleCoverage (GLclampf, GLboolean)
func SampleCoverage(arg0 GLclampf, arg1 GLboolean) 
{
	C.glSampleCoverage(C.GLclampf(arg0), C.GLboolean(arg1));
}

//extern void glCompressedTexImage3D (GLenum, GLint, GLenum, GLsizei, GLsizei, GLsizei, GLint, GLsizei, const GLvoid *)
func CompressedTexImage3D(arg0 GLenum, arg1 GLint, arg2 GLenum, arg3 GLsizei, arg4 GLsizei, arg5 GLsizei, arg6 GLint, arg7 GLsizei, arg8 unsafe.Pointer) 
{
	C.glCompressedTexImage3D(C.GLenum(arg0), C.GLint(arg1), C.GLenum(arg2), C.GLsizei(arg3), C.GLsizei(arg4), C.GLsizei(arg5), C.GLint(arg6), C.GLsizei(arg7), arg8);
}

//extern void glCompressedTexImage2D (GLenum, GLint, GLenum, GLsizei, GLsizei, GLint, GLsizei, const GLvoid *)
func CompressedTexImage2D(arg0 GLenum, arg1 GLint, arg2 GLenum, arg3 GLsizei, arg4 GLsizei, arg5 GLint, arg6 GLsizei, arg7 unsafe.Pointer) 
{
	C.glCompressedTexImage2D(C.GLenum(arg0), C.GLint(arg1), C.GLenum(arg2), C.GLsizei(arg3), C.GLsizei(arg4), C.GLint(arg5), C.GLsizei(arg6), arg7);
}

//extern void glCompressedTexImage1D (GLenum, GLint, GLenum, GLsizei, GLint, GLsizei, const GLvoid *)
func CompressedTexImage1D(arg0 GLenum, arg1 GLint, arg2 GLenum, arg3 GLsizei, arg4 GLint, arg5 GLsizei, arg6 unsafe.Pointer) 
{
	C.glCompressedTexImage1D(C.GLenum(arg0), C.GLint(arg1), C.GLenum(arg2), C.GLsizei(arg3), C.GLint(arg4), C.GLsizei(arg5), arg6);
}

//extern void glCompressedTexSubImage3D (GLenum, GLint, GLint, GLint, GLint, GLsizei, GLsizei, GLsizei, GLenum, GLsizei, const GLvoid *)
func CompressedTexSubImage3D(arg0 GLenum, arg1 GLint, arg2 GLint, arg3 GLint, arg4 GLint, arg5 GLsizei, arg6 GLsizei, arg7 GLsizei, arg8 GLenum, arg9 GLsizei, arg10 unsafe.Pointer) 
{
	C.glCompressedTexSubImage3D(C.GLenum(arg0), C.GLint(arg1), C.GLint(arg2), C.GLint(arg3), C.GLint(arg4), C.GLsizei(arg5), C.GLsizei(arg6), C.GLsizei(arg7), C.GLenum(arg8), C.GLsizei(arg9), arg10);
}

//extern void glCompressedTexSubImage2D (GLenum, GLint, GLint, GLint, GLsizei, GLsizei, GLenum, GLsizei, const GLvoid *)
func CompressedTexSubImage2D(arg0 GLenum, arg1 GLint, arg2 GLint, arg3 GLint, arg4 GLsizei, arg5 GLsizei, arg6 GLenum, arg7 GLsizei, arg8 unsafe.Pointer) 
{
	C.glCompressedTexSubImage2D(C.GLenum(arg0), C.GLint(arg1), C.GLint(arg2), C.GLint(arg3), C.GLsizei(arg4), C.GLsizei(arg5), C.GLenum(arg6), C.GLsizei(arg7), arg8);
}

//extern void glCompressedTexSubImage1D (GLenum, GLint, GLint, GLsizei, GLenum, GLsizei, const GLvoid *)
func CompressedTexSubImage1D(arg0 GLenum, arg1 GLint, arg2 GLint, arg3 GLsizei, arg4 GLenum, arg5 GLsizei, arg6 unsafe.Pointer) 
{
	C.glCompressedTexSubImage1D(C.GLenum(arg0), C.GLint(arg1), C.GLint(arg2), C.GLsizei(arg3), C.GLenum(arg4), C.GLsizei(arg5), arg6);
}

//extern void glGetCompressedTexImage (GLenum, GLint, GLvoid *)
func GetCompressedTexImage(arg0 GLenum, arg1 GLint, arg2 unsafe.Pointer) 
{
	C.glGetCompressedTexImage(C.GLenum(arg0), C.GLint(arg1), arg2);
}

//extern void glBlendFuncSeparate (GLenum, GLenum, GLenum, GLenum)
func BlendFuncSeparate(arg0 GLenum, arg1 GLenum, arg2 GLenum, arg3 GLenum) 
{
	C.glBlendFuncSeparate(C.GLenum(arg0), C.GLenum(arg1), C.GLenum(arg2), C.GLenum(arg3));
}

//extern void glFogCoordf (GLfloat)
func FogCoordf(arg0 GLfloat) 
{
	C.glFogCoordf(C.GLfloat(arg0));
}

//extern void glFogCoordfv (const GLfloat *)
func FogCoordfv(arg0 *GLfloat) 
{
	C.glFogCoordfv((*C.GLfloat)(arg0));
}

//extern void glFogCoordd (GLdouble)
func FogCoordd(arg0 GLdouble) 
{
	C.glFogCoordd(C.GLdouble(arg0));
}

//extern void glFogCoorddv (const GLdouble *)
func FogCoorddv(arg0 *GLdouble) 
{
	C.glFogCoorddv((*C.GLdouble)(arg0));
}

//extern void glFogCoordPointer (GLenum, GLsizei, const GLvoid *)
func FogCoordPointer(arg0 GLenum, arg1 GLsizei, arg2 unsafe.Pointer) 
{
	C.glFogCoordPointer(C.GLenum(arg0), C.GLsizei(arg1), arg2);
}

//extern void glMultiDrawArrays (GLenum, GLint *, GLsizei *, GLsizei)
func MultiDrawArrays(arg0 GLenum, arg1 *GLint, arg2 *GLsizei, arg3 GLsizei) 
{
	C.glMultiDrawArrays(C.GLenum(arg0), (*C.GLint)(arg1), (*C.GLsizei)(arg2), C.GLsizei(arg3));
}

//extern void glMultiDrawElements (GLenum, const GLsizei *, GLenum, const GLvoid* *, GLsizei)
func MultiDrawElements(arg0 GLenum, arg1 *GLsizei, arg2 GLenum, arg3 *unsafe.Pointer, arg4 GLsizei) 
{
	C.glMultiDrawElements(C.GLenum(arg0), (*C.GLsizei)(arg1), C.GLenum(arg2), arg3, C.GLsizei(arg4));
}

//extern void glPointParameterf (GLenum, GLfloat)
func PointParameterf(arg0 GLenum, arg1 GLfloat) 
{
	C.glPointParameterf(C.GLenum(arg0), C.GLfloat(arg1));
}

//extern void glPointParameterfv (GLenum, const GLfloat *)
func PointParameterfv(arg0 GLenum, arg1 *GLfloat) 
{
	C.glPointParameterfv(C.GLenum(arg0), (*C.GLfloat)(arg1));
}

//extern void glPointParameteri (GLenum, GLint)
func PointParameteri(arg0 GLenum, arg1 GLint) 
{
	C.glPointParameteri(C.GLenum(arg0), C.GLint(arg1));
}

//extern void glPointParameteriv (GLenum, const GLint *)
func PointParameteriv(arg0 GLenum, arg1 *GLint) 
{
	C.glPointParameteriv(C.GLenum(arg0), (*C.GLint)(arg1));
}

//extern void glSecondaryColor3b (GLbyte, GLbyte, GLbyte)
func SecondaryColor3b(arg0 GLbyte, arg1 GLbyte, arg2 GLbyte) 
{
	C.glSecondaryColor3b(C.GLbyte(arg0), C.GLbyte(arg1), C.GLbyte(arg2));
}

//extern void glSecondaryColor3bv (const GLbyte *)
func SecondaryColor3bv(arg0 *GLbyte) 
{
	C.glSecondaryColor3bv((*C.GLbyte)(arg0));
}

//extern void glSecondaryColor3d (GLdouble, GLdouble, GLdouble)
func SecondaryColor3d(arg0 GLdouble, arg1 GLdouble, arg2 GLdouble) 
{
	C.glSecondaryColor3d(C.GLdouble(arg0), C.GLdouble(arg1), C.GLdouble(arg2));
}

//extern void glSecondaryColor3dv (const GLdouble *)
func SecondaryColor3dv(arg0 *GLdouble) 
{
	C.glSecondaryColor3dv((*C.GLdouble)(arg0));
}

//extern void glSecondaryColor3f (GLfloat, GLfloat, GLfloat)
func SecondaryColor3f(arg0 GLfloat, arg1 GLfloat, arg2 GLfloat) 
{
	C.glSecondaryColor3f(C.GLfloat(arg0), C.GLfloat(arg1), C.GLfloat(arg2));
}

//extern void glSecondaryColor3fv (const GLfloat *)
func SecondaryColor3fv(arg0 *GLfloat) 
{
	C.glSecondaryColor3fv((*C.GLfloat)(arg0));
}

//extern void glSecondaryColor3i (GLint, GLint, GLint)
func SecondaryColor3i(arg0 GLint, arg1 GLint, arg2 GLint) 
{
	C.glSecondaryColor3i(C.GLint(arg0), C.GLint(arg1), C.GLint(arg2));
}

//extern void glSecondaryColor3iv (const GLint *)
func SecondaryColor3iv(arg0 *GLint) 
{
	C.glSecondaryColor3iv((*C.GLint)(arg0));
}

//extern void glSecondaryColor3s (GLshort, GLshort, GLshort)
func SecondaryColor3s(arg0 GLshort, arg1 GLshort, arg2 GLshort) 
{
	C.glSecondaryColor3s(C.GLshort(arg0), C.GLshort(arg1), C.GLshort(arg2));
}

//extern void glSecondaryColor3sv (const GLshort *)
func SecondaryColor3sv(arg0 *GLshort) 
{
	C.glSecondaryColor3sv((*C.GLshort)(arg0));
}

//extern void glSecondaryColor3ub (GLubyte, GLubyte, GLubyte)
func SecondaryColor3ub(arg0 GLubyte, arg1 GLubyte, arg2 GLubyte) 
{
	C.glSecondaryColor3ub(C.GLubyte(arg0), C.GLubyte(arg1), C.GLubyte(arg2));
}

//extern void glSecondaryColor3ubv (const GLubyte *)
func SecondaryColor3ubv(arg0 *GLubyte) 
{
	C.glSecondaryColor3ubv((*C.GLubyte)(arg0));
}

//extern void glSecondaryColor3ui (GLuint, GLuint, GLuint)
func SecondaryColor3ui(arg0 GLuint, arg1 GLuint, arg2 GLuint) 
{
	C.glSecondaryColor3ui(C.GLuint(arg0), C.GLuint(arg1), C.GLuint(arg2));
}

//extern void glSecondaryColor3uiv (const GLuint *)
func SecondaryColor3uiv(arg0 *GLuint) 
{
	C.glSecondaryColor3uiv((*C.GLuint)(arg0));
}

//extern void glSecondaryColor3us (GLushort, GLushort, GLushort)
func SecondaryColor3us(arg0 GLushort, arg1 GLushort, arg2 GLushort) 
{
	C.glSecondaryColor3us(C.GLushort(arg0), C.GLushort(arg1), C.GLushort(arg2));
}

//extern void glSecondaryColor3usv (const GLushort *)
func SecondaryColor3usv(arg0 *GLushort) 
{
	C.glSecondaryColor3usv((*C.GLushort)(arg0));
}

//extern void glSecondaryColorPointer (GLint, GLenum, GLsizei, const GLvoid *)
func SecondaryColorPointer(arg0 GLint, arg1 GLenum, arg2 GLsizei, arg3 unsafe.Pointer) 
{
	C.glSecondaryColorPointer(C.GLint(arg0), C.GLenum(arg1), C.GLsizei(arg2), arg3);
}

//extern void glWindowPos2d (GLdouble, GLdouble)
func WindowPos2d(arg0 GLdouble, arg1 GLdouble) 
{
	C.glWindowPos2d(C.GLdouble(arg0), C.GLdouble(arg1));
}

//extern void glWindowPos2dv (const GLdouble *)
func WindowPos2dv(arg0 *GLdouble) 
{
	C.glWindowPos2dv((*C.GLdouble)(arg0));
}

//extern void glWindowPos2f (GLfloat, GLfloat)
func WindowPos2f(arg0 GLfloat, arg1 GLfloat) 
{
	C.glWindowPos2f(C.GLfloat(arg0), C.GLfloat(arg1));
}

//extern void glWindowPos2fv (const GLfloat *)
func WindowPos2fv(arg0 *GLfloat) 
{
	C.glWindowPos2fv((*C.GLfloat)(arg0));
}

//extern void glWindowPos2i (GLint, GLint)
func WindowPos2i(arg0 GLint, arg1 GLint) 
{
	C.glWindowPos2i(C.GLint(arg0), C.GLint(arg1));
}

//extern void glWindowPos2iv (const GLint *)
func WindowPos2iv(arg0 *GLint) 
{
	C.glWindowPos2iv((*C.GLint)(arg0));
}

//extern void glWindowPos2s (GLshort, GLshort)
func WindowPos2s(arg0 GLshort, arg1 GLshort) 
{
	C.glWindowPos2s(C.GLshort(arg0), C.GLshort(arg1));
}

//extern void glWindowPos2sv (const GLshort *)
func WindowPos2sv(arg0 *GLshort) 
{
	C.glWindowPos2sv((*C.GLshort)(arg0));
}

//extern void glWindowPos3d (GLdouble, GLdouble, GLdouble)
func WindowPos3d(arg0 GLdouble, arg1 GLdouble, arg2 GLdouble) 
{
	C.glWindowPos3d(C.GLdouble(arg0), C.GLdouble(arg1), C.GLdouble(arg2));
}

//extern void glWindowPos3dv (const GLdouble *)
func WindowPos3dv(arg0 *GLdouble) 
{
	C.glWindowPos3dv((*C.GLdouble)(arg0));
}

//extern void glWindowPos3f (GLfloat, GLfloat, GLfloat)
func WindowPos3f(arg0 GLfloat, arg1 GLfloat, arg2 GLfloat) 
{
	C.glWindowPos3f(C.GLfloat(arg0), C.GLfloat(arg1), C.GLfloat(arg2));
}

//extern void glWindowPos3fv (const GLfloat *)
func WindowPos3fv(arg0 *GLfloat) 
{
	C.glWindowPos3fv((*C.GLfloat)(arg0));
}

//extern void glWindowPos3i (GLint, GLint, GLint)
func WindowPos3i(arg0 GLint, arg1 GLint, arg2 GLint) 
{
	C.glWindowPos3i(C.GLint(arg0), C.GLint(arg1), C.GLint(arg2));
}

//extern void glWindowPos3iv (const GLint *)
func WindowPos3iv(arg0 *GLint) 
{
	C.glWindowPos3iv((*C.GLint)(arg0));
}

//extern void glWindowPos3s (GLshort, GLshort, GLshort)
func WindowPos3s(arg0 GLshort, arg1 GLshort, arg2 GLshort) 
{
	C.glWindowPos3s(C.GLshort(arg0), C.GLshort(arg1), C.GLshort(arg2));
}

//extern void glWindowPos3sv (const GLshort *)
func WindowPos3sv(arg0 *GLshort) 
{
	C.glWindowPos3sv((*C.GLshort)(arg0));
}

//extern void glGenQueries (GLsizei, GLuint *)
func GenQueries(arg0 GLsizei, arg1 *GLuint) 
{
	C.glGenQueries(C.GLsizei(arg0), (*C.GLuint)(arg1));
}

//extern void glDeleteQueries (GLsizei, const GLuint *)
func DeleteQueries(arg0 GLsizei, arg1 *GLuint) 
{
	C.glDeleteQueries(C.GLsizei(arg0), (*C.GLuint)(arg1));
}

//extern GLboolean glIsQuery (GLuint)
func IsQuery(arg0 GLuint) GLboolean
{
	return GLboolean(C.glIsQuery(C.GLuint(arg0)));
}

//extern void glBeginQuery (GLenum, GLuint)
func BeginQuery(arg0 GLenum, arg1 GLuint) 
{
	C.glBeginQuery(C.GLenum(arg0), C.GLuint(arg1));
}

//extern void glEndQuery (GLenum)
func EndQuery(arg0 GLenum) 
{
	C.glEndQuery(C.GLenum(arg0));
}

//extern void glGetQueryiv (GLenum, GLenum, GLint *)
func GetQueryiv(arg0 GLenum, arg1 GLenum, arg2 *GLint) 
{
	C.glGetQueryiv(C.GLenum(arg0), C.GLenum(arg1), (*C.GLint)(arg2));
}

//extern void glGetQueryObjectiv (GLuint, GLenum, GLint *)
func GetQueryObjectiv(arg0 GLuint, arg1 GLenum, arg2 *GLint) 
{
	C.glGetQueryObjectiv(C.GLuint(arg0), C.GLenum(arg1), (*C.GLint)(arg2));
}

//extern void glGetQueryObjectuiv (GLuint, GLenum, GLuint *)
func GetQueryObjectuiv(arg0 GLuint, arg1 GLenum, arg2 *GLuint) 
{
	C.glGetQueryObjectuiv(C.GLuint(arg0), C.GLenum(arg1), (*C.GLuint)(arg2));
}

//extern void glBindBuffer (GLenum, GLuint)
func BindBuffer(arg0 GLenum, arg1 GLuint) 
{
	C.glBindBuffer(C.GLenum(arg0), C.GLuint(arg1));
}

//extern void glDeleteBuffers (GLsizei, const GLuint *)
func DeleteBuffers(arg0 GLsizei, arg1 *GLuint) 
{
	C.glDeleteBuffers(C.GLsizei(arg0), (*C.GLuint)(arg1));
}

//extern void glGenBuffers (GLsizei, GLuint *)
func GenBuffers(arg0 GLsizei, arg1 *GLuint) 
{
	C.glGenBuffers(C.GLsizei(arg0), (*C.GLuint)(arg1));
}

//extern GLboolean glIsBuffer (GLuint)
func IsBuffer(arg0 GLuint) GLboolean
{
	return GLboolean(C.glIsBuffer(C.GLuint(arg0)));
}

//extern void glBufferData (GLenum, GLsizeiptr, const GLvoid *, GLenum)
func BufferData(arg0 GLenum, arg1 GLsizeiptr, arg2 unsafe.Pointer, arg3 GLenum) 
{
	C.glBufferData(C.GLenum(arg0), C.GLsizeiptr(arg1), arg2, C.GLenum(arg3));
}

//extern void glBufferSubData (GLenum, GLintptr, GLsizeiptr, const GLvoid *)
func BufferSubData(arg0 GLenum, arg1 GLintptr, arg2 GLsizeiptr, arg3 unsafe.Pointer) 
{
	C.glBufferSubData(C.GLenum(arg0), C.GLintptr(arg1), C.GLsizeiptr(arg2), arg3);
}

//extern void glGetBufferSubData (GLenum, GLintptr, GLsizeiptr, GLvoid *)
func GetBufferSubData(arg0 GLenum, arg1 GLintptr, arg2 GLsizeiptr, arg3 unsafe.Pointer) 
{
	C.glGetBufferSubData(C.GLenum(arg0), C.GLintptr(arg1), C.GLsizeiptr(arg2), arg3);
}

//extern GLvoid* glMapBuffer (GLenum, GLenum)
func MapBuffer(arg0 GLenum, arg1 GLenum) unsafe.Pointer
{
	return unsafe.Pointer(C.glMapBuffer(C.GLenum(arg0), C.GLenum(arg1)));
}

//extern GLboolean glUnmapBuffer (GLenum)
func UnmapBuffer(arg0 GLenum) GLboolean
{
	return GLboolean(C.glUnmapBuffer(C.GLenum(arg0)));
}

//extern void glGetBufferParameteriv (GLenum, GLenum, GLint *)
func GetBufferParameteriv(arg0 GLenum, arg1 GLenum, arg2 *GLint) 
{
	C.glGetBufferParameteriv(C.GLenum(arg0), C.GLenum(arg1), (*C.GLint)(arg2));
}

//extern void glGetBufferPointerv (GLenum, GLenum, GLvoid* *)
func GetBufferPointerv(arg0 GLenum, arg1 GLenum, arg2 *unsafe.Pointer) 
{
	C.glGetBufferPointerv(C.GLenum(arg0), C.GLenum(arg1), arg2);
}

//extern void glBlendEquationSeparate (GLenum, GLenum)
func BlendEquationSeparate(arg0 GLenum, arg1 GLenum) 
{
	C.glBlendEquationSeparate(C.GLenum(arg0), C.GLenum(arg1));
}

//extern void glDrawBuffers (GLsizei, const GLenum *)
func DrawBuffers(arg0 GLsizei, arg1 *GLenum) 
{
	C.glDrawBuffers(C.GLsizei(arg0), (*C.GLenum)(arg1));
}

//extern void glStencilOpSeparate (GLenum, GLenum, GLenum, GLenum)
func StencilOpSeparate(arg0 GLenum, arg1 GLenum, arg2 GLenum, arg3 GLenum) 
{
	C.glStencilOpSeparate(C.GLenum(arg0), C.GLenum(arg1), C.GLenum(arg2), C.GLenum(arg3));
}

//extern void glStencilFuncSeparate (GLenum, GLenum, GLint, GLuint)
func StencilFuncSeparate(arg0 GLenum, arg1 GLenum, arg2 GLint, arg3 GLuint) 
{
	C.glStencilFuncSeparate(C.GLenum(arg0), C.GLenum(arg1), C.GLint(arg2), C.GLuint(arg3));
}

//extern void glStencilMaskSeparate (GLenum, GLuint)
func StencilMaskSeparate(arg0 GLenum, arg1 GLuint) 
{
	C.glStencilMaskSeparate(C.GLenum(arg0), C.GLuint(arg1));
}

//extern void glAttachShader (GLuint, GLuint)
func AttachShader(arg0 GLuint, arg1 GLuint) 
{
	C.glAttachShader(C.GLuint(arg0), C.GLuint(arg1));
}

//extern void glBindAttribLocation (GLuint, GLuint, const GLchar *)
func BindAttribLocation(arg0 GLuint, arg1 GLuint, arg2 string) 
{
	C.glBindAttribLocation(C.GLuint(arg0), C.GLuint(arg1), (*_C_GLchar)((C.CString)(arg2)));
}

//extern void glCompileShader (GLuint)
func CompileShader(arg0 GLuint) 
{
	C.glCompileShader(C.GLuint(arg0));
}

//extern GLuint glCreateProgram (void)
func CreateProgram() GLuint
{
	return GLuint(C.glCreateProgram());
}

//extern GLuint glCreateShader (GLenum)
func CreateShader(arg0 GLenum) GLuint
{
	return GLuint(C.glCreateShader(C.GLenum(arg0)));
}

//extern void glDeleteProgram (GLuint)
func DeleteProgram(arg0 GLuint) 
{
	C.glDeleteProgram(C.GLuint(arg0));
}

//extern void glDeleteShader (GLuint)
func DeleteShader(arg0 GLuint) 
{
	C.glDeleteShader(C.GLuint(arg0));
}

//extern void glDetachShader (GLuint, GLuint)
func DetachShader(arg0 GLuint, arg1 GLuint) 
{
	C.glDetachShader(C.GLuint(arg0), C.GLuint(arg1));
}

//extern void glDisableVertexAttribArray (GLuint)
func DisableVertexAttribArray(arg0 GLuint) 
{
	C.glDisableVertexAttribArray(C.GLuint(arg0));
}

//extern void glEnableVertexAttribArray (GLuint)
func EnableVertexAttribArray(arg0 GLuint) 
{
	C.glEnableVertexAttribArray(C.GLuint(arg0));
}

//extern void glGetActiveAttrib (GLuint, GLuint, GLsizei, GLsizei *, GLint *, GLenum *, GLchar *)
func GetActiveAttrib(arg0 GLuint, arg1 GLuint, arg2 GLsizei, arg3 *GLsizei, arg4 *GLint, arg5 *GLenum, arg6 string) 
{
	C.glGetActiveAttrib(C.GLuint(arg0), C.GLuint(arg1), C.GLsizei(arg2), (*C.GLsizei)(arg3), (*C.GLint)(arg4), (*C.GLenum)(arg5), (*_C_GLchar)((C.CString)(arg6)));
}

//extern void glGetActiveUniform (GLuint, GLuint, GLsizei, GLsizei *, GLint *, GLenum *, GLchar *)
func GetActiveUniform(arg0 GLuint, arg1 GLuint, arg2 GLsizei, arg3 *GLsizei, arg4 *GLint, arg5 *GLenum, arg6 string) 
{
	C.glGetActiveUniform(C.GLuint(arg0), C.GLuint(arg1), C.GLsizei(arg2), (*C.GLsizei)(arg3), (*C.GLint)(arg4), (*C.GLenum)(arg5), (*_C_GLchar)((C.CString)(arg6)));
}

//extern void glGetAttachedShaders (GLuint, GLsizei, GLsizei *, GLuint *)
func GetAttachedShaders(arg0 GLuint, arg1 GLsizei, arg2 *GLsizei, arg3 *GLuint) 
{
	C.glGetAttachedShaders(C.GLuint(arg0), C.GLsizei(arg1), (*C.GLsizei)(arg2), (*C.GLuint)(arg3));
}

//extern GLint glGetAttribLocation (GLuint, const GLchar *)
func GetAttribLocation(arg0 GLuint, arg1 string) GLint
{
	return GLint(C.glGetAttribLocation(C.GLuint(arg0), (*_C_GLchar)((C.CString)(arg1))));
}

//extern void glGetProgramiv (GLuint, GLenum, GLint *)
func GetProgramiv(arg0 GLuint, arg1 GLenum, arg2 *GLint) 
{
	C.glGetProgramiv(C.GLuint(arg0), C.GLenum(arg1), (*C.GLint)(arg2));
}

//extern void glGetProgramInfoLog (GLuint, GLsizei, GLsizei *, GLchar *)
func GetProgramInfoLog(program GLuint) string
{

    buf := new([1024]C.char);
    var size int;
	C.glGetProgramInfoLog(C.GLuint(program), 1024, (*C.GLsizei)(unsafe.Pointer(&size)), (*C.GLchar)(unsafe.Pointer(buf)));
    buf[size]=0;

    return C.GoString((*C.char)(unsafe.Pointer(buf)));

}

//extern void glGetShaderiv (GLuint, GLenum, GLint *)
func GetShaderiv(arg0 GLuint, arg1 GLenum, arg2 *GLint) 
{
	C.glGetShaderiv(C.GLuint(arg0), C.GLenum(arg1), (*C.GLint)(arg2));
}

//extern void glGetShaderInfoLog (GLuint, GLsizei, GLsizei *, GLchar *)
func GetShaderInfoLog(arg0 GLuint, arg1 GLsizei, arg2 *GLsizei, arg3 string) 
{
	C.glGetShaderInfoLog(C.GLuint(arg0), C.GLsizei(arg1), (*C.GLsizei)(arg2), (*_C_GLchar)((C.CString)(arg3)));
}

//extern void glGetShaderSource (GLuint, GLsizei, GLsizei *, GLchar *)
func GetShaderSource(arg0 GLuint, arg1 GLsizei, arg2 *GLsizei, arg3 string) 
{
	C.glGetShaderSource(C.GLuint(arg0), C.GLsizei(arg1), (*C.GLsizei)(arg2), (*_C_GLchar)((C.CString)(arg3)));
}

//extern GLint glGetUniformLocation (GLuint, const GLchar *)
func GetUniformLocation(arg0 GLuint, arg1 string) GLint
{
	return GLint(C.glGetUniformLocation(C.GLuint(arg0), (*_C_GLchar)((C.CString)(arg1))));
}

//extern void glGetUniformfv (GLuint, GLint, GLfloat *)
func GetUniformfv(arg0 GLuint, arg1 GLint, arg2 *GLfloat) 
{
	C.glGetUniformfv(C.GLuint(arg0), C.GLint(arg1), (*C.GLfloat)(arg2));
}

//extern void glGetUniformiv (GLuint, GLint, GLint *)
func GetUniformiv(arg0 GLuint, arg1 GLint, arg2 *GLint) 
{
	C.glGetUniformiv(C.GLuint(arg0), C.GLint(arg1), (*C.GLint)(arg2));
}

//extern void glGetVertexAttribdv (GLuint, GLenum, GLdouble *)
func GetVertexAttribdv(arg0 GLuint, arg1 GLenum, arg2 *GLdouble) 
{
	C.glGetVertexAttribdv(C.GLuint(arg0), C.GLenum(arg1), (*C.GLdouble)(arg2));
}

//extern void glGetVertexAttribfv (GLuint, GLenum, GLfloat *)
func GetVertexAttribfv(arg0 GLuint, arg1 GLenum, arg2 *GLfloat) 
{
	C.glGetVertexAttribfv(C.GLuint(arg0), C.GLenum(arg1), (*C.GLfloat)(arg2));
}

//extern void glGetVertexAttribiv (GLuint, GLenum, GLint *)
func GetVertexAttribiv(arg0 GLuint, arg1 GLenum, arg2 *GLint) 
{
	C.glGetVertexAttribiv(C.GLuint(arg0), C.GLenum(arg1), (*C.GLint)(arg2));
}

//extern void glGetVertexAttribPointerv (GLuint, GLenum, GLvoid* *)
func GetVertexAttribPointerv(arg0 GLuint, arg1 GLenum, arg2 *unsafe.Pointer) 
{
	C.glGetVertexAttribPointerv(C.GLuint(arg0), C.GLenum(arg1), arg2);
}

//extern GLboolean glIsProgram (GLuint)
func IsProgram(arg0 GLuint) GLboolean
{
	return GLboolean(C.glIsProgram(C.GLuint(arg0)));
}

//extern GLboolean glIsShader (GLuint)
func IsShader(arg0 GLuint) GLboolean
{
	return GLboolean(C.glIsShader(C.GLuint(arg0)));
}

//extern void glLinkProgram (GLuint)
func LinkProgram(arg0 GLuint) 
{
	C.glLinkProgram(C.GLuint(arg0));
}

func array(slice unsafe.Pointer) unsafe.Pointer
{
    return *((*unsafe.Pointer)(slice));
}

//extern void glShaderSource (GLuint, GLsizei, const GLchar* *, const GLint *)
func ShaderSource(shader GLuint, sources []string) 
{

    var count=len(sources);

    var strings=make([](*C.GLchar),count);
    var lengths=make([]C.GLint,count);

    for i := 0; i < count; i++
    {
        strings[i]=(*C.GLchar)(C.CString(sources[i]));
        lengths[i]=C.GLint(len(sources[i]));
    }

	C.glShaderSource(
        C.GLuint(shader),
        C.GLsizei(count),
        (**C.GLchar)(array(unsafe.Pointer(&strings))),
        (*C.GLint)(array(unsafe.Pointer(&lengths)))
    );

    for i := 0; i < count; i++
    {
        C.free(unsafe.Pointer(strings[i]));
    }

}

//extern void glUseProgram (GLuint)
func UseProgram(arg0 GLuint) 
{
	C.glUseProgram(C.GLuint(arg0));
}

//extern void glUniform1f (GLint, GLfloat)
func Uniform1f(arg0 GLint, arg1 GLfloat) 
{
	C.glUniform1f(C.GLint(arg0), C.GLfloat(arg1));
}

//extern void glUniform2f (GLint, GLfloat, GLfloat)
func Uniform2f(arg0 GLint, arg1 GLfloat, arg2 GLfloat) 
{
	C.glUniform2f(C.GLint(arg0), C.GLfloat(arg1), C.GLfloat(arg2));
}

//extern void glUniform3f (GLint, GLfloat, GLfloat, GLfloat)
func Uniform3f(arg0 GLint, arg1 GLfloat, arg2 GLfloat, arg3 GLfloat) 
{
	C.glUniform3f(C.GLint(arg0), C.GLfloat(arg1), C.GLfloat(arg2), C.GLfloat(arg3));
}

//extern void glUniform4f (GLint, GLfloat, GLfloat, GLfloat, GLfloat)
func Uniform4f(arg0 GLint, arg1 GLfloat, arg2 GLfloat, arg3 GLfloat, arg4 GLfloat) 
{
	C.glUniform4f(C.GLint(arg0), C.GLfloat(arg1), C.GLfloat(arg2), C.GLfloat(arg3), C.GLfloat(arg4));
}

//extern void glUniform1i (GLint, GLint)
func Uniform1i(arg0 GLint, arg1 GLint) 
{
	C.glUniform1i(C.GLint(arg0), C.GLint(arg1));
}

//extern void glUniform2i (GLint, GLint, GLint)
func Uniform2i(arg0 GLint, arg1 GLint, arg2 GLint) 
{
	C.glUniform2i(C.GLint(arg0), C.GLint(arg1), C.GLint(arg2));
}

//extern void glUniform3i (GLint, GLint, GLint, GLint)
func Uniform3i(arg0 GLint, arg1 GLint, arg2 GLint, arg3 GLint) 
{
	C.glUniform3i(C.GLint(arg0), C.GLint(arg1), C.GLint(arg2), C.GLint(arg3));
}

//extern void glUniform4i (GLint, GLint, GLint, GLint, GLint)
func Uniform4i(arg0 GLint, arg1 GLint, arg2 GLint, arg3 GLint, arg4 GLint) 
{
	C.glUniform4i(C.GLint(arg0), C.GLint(arg1), C.GLint(arg2), C.GLint(arg3), C.GLint(arg4));
}

//extern void glUniform1fv (GLint, GLsizei, const GLfloat *)
func Uniform1fv(arg0 GLint, arg1 GLsizei, arg2 *GLfloat) 
{
	C.glUniform1fv(C.GLint(arg0), C.GLsizei(arg1), (*C.GLfloat)(arg2));
}

//extern void glUniform2fv (GLint, GLsizei, const GLfloat *)
func Uniform2fv(arg0 GLint, arg1 GLsizei, arg2 *GLfloat) 
{
	C.glUniform2fv(C.GLint(arg0), C.GLsizei(arg1), (*C.GLfloat)(arg2));
}

//extern void glUniform3fv (GLint, GLsizei, const GLfloat *)
func Uniform3fv(arg0 GLint, arg1 GLsizei, arg2 *GLfloat) 
{
	C.glUniform3fv(C.GLint(arg0), C.GLsizei(arg1), (*C.GLfloat)(arg2));
}

//extern void glUniform4fv (GLint, GLsizei, const GLfloat *)
func Uniform4fv(arg0 GLint, arg1 GLsizei, arg2 *GLfloat) 
{
	C.glUniform4fv(C.GLint(arg0), C.GLsizei(arg1), (*C.GLfloat)(arg2));
}

//extern void glUniform1iv (GLint, GLsizei, const GLint *)
func Uniform1iv(arg0 GLint, arg1 GLsizei, arg2 *GLint) 
{
	C.glUniform1iv(C.GLint(arg0), C.GLsizei(arg1), (*C.GLint)(arg2));
}

//extern void glUniform2iv (GLint, GLsizei, const GLint *)
func Uniform2iv(arg0 GLint, arg1 GLsizei, arg2 *GLint) 
{
	C.glUniform2iv(C.GLint(arg0), C.GLsizei(arg1), (*C.GLint)(arg2));
}

//extern void glUniform3iv (GLint, GLsizei, const GLint *)
func Uniform3iv(arg0 GLint, arg1 GLsizei, arg2 *GLint) 
{
	C.glUniform3iv(C.GLint(arg0), C.GLsizei(arg1), (*C.GLint)(arg2));
}

//extern void glUniform4iv (GLint, GLsizei, const GLint *)
func Uniform4iv(arg0 GLint, arg1 GLsizei, arg2 *GLint) 
{
	C.glUniform4iv(C.GLint(arg0), C.GLsizei(arg1), (*C.GLint)(arg2));
}

//extern void glUniformMatrix2fv (GLint, GLsizei, GLboolean, const GLfloat *)
func UniformMatrix2fv(arg0 GLint, arg1 GLsizei, arg2 GLboolean, arg3 *GLfloat) 
{
	C.glUniformMatrix2fv(C.GLint(arg0), C.GLsizei(arg1), C.GLboolean(arg2), (*C.GLfloat)(arg3));
}

//extern void glUniformMatrix3fv (GLint, GLsizei, GLboolean, const GLfloat *)
func UniformMatrix3fv(arg0 GLint, arg1 GLsizei, arg2 GLboolean, arg3 *GLfloat) 
{
	C.glUniformMatrix3fv(C.GLint(arg0), C.GLsizei(arg1), C.GLboolean(arg2), (*C.GLfloat)(arg3));
}

//extern void glUniformMatrix4fv (GLint, GLsizei, GLboolean, const GLfloat *)
func UniformMatrix4fv(arg0 GLint, arg1 GLsizei, arg2 GLboolean, arg3 *GLfloat) 
{
	C.glUniformMatrix4fv(C.GLint(arg0), C.GLsizei(arg1), C.GLboolean(arg2), (*C.GLfloat)(arg3));
}

//extern void glValidateProgram (GLuint)
func ValidateProgram(arg0 GLuint) 
{
	C.glValidateProgram(C.GLuint(arg0));
}

//extern void glVertexAttrib1d (GLuint, GLdouble)
func VertexAttrib1d(arg0 GLuint, arg1 GLdouble) 
{
	C.glVertexAttrib1d(C.GLuint(arg0), C.GLdouble(arg1));
}

//extern void glVertexAttrib1dv (GLuint, const GLdouble *)
func VertexAttrib1dv(arg0 GLuint, arg1 *GLdouble) 
{
	C.glVertexAttrib1dv(C.GLuint(arg0), (*C.GLdouble)(arg1));
}

//extern void glVertexAttrib1f (GLuint, GLfloat)
func VertexAttrib1f(arg0 GLuint, arg1 GLfloat) 
{
	C.glVertexAttrib1f(C.GLuint(arg0), C.GLfloat(arg1));
}

//extern void glVertexAttrib1fv (GLuint, const GLfloat *)
func VertexAttrib1fv(arg0 GLuint, arg1 *GLfloat) 
{
	C.glVertexAttrib1fv(C.GLuint(arg0), (*C.GLfloat)(arg1));
}

//extern void glVertexAttrib1s (GLuint, GLshort)
func VertexAttrib1s(arg0 GLuint, arg1 GLshort) 
{
	C.glVertexAttrib1s(C.GLuint(arg0), C.GLshort(arg1));
}

//extern void glVertexAttrib1sv (GLuint, const GLshort *)
func VertexAttrib1sv(arg0 GLuint, arg1 *GLshort) 
{
	C.glVertexAttrib1sv(C.GLuint(arg0), (*C.GLshort)(arg1));
}

//extern void glVertexAttrib2d (GLuint, GLdouble, GLdouble)
func VertexAttrib2d(arg0 GLuint, arg1 GLdouble, arg2 GLdouble) 
{
	C.glVertexAttrib2d(C.GLuint(arg0), C.GLdouble(arg1), C.GLdouble(arg2));
}

//extern void glVertexAttrib2dv (GLuint, const GLdouble *)
func VertexAttrib2dv(arg0 GLuint, arg1 *GLdouble) 
{
	C.glVertexAttrib2dv(C.GLuint(arg0), (*C.GLdouble)(arg1));
}

//extern void glVertexAttrib2f (GLuint, GLfloat, GLfloat)
func VertexAttrib2f(arg0 GLuint, arg1 GLfloat, arg2 GLfloat) 
{
	C.glVertexAttrib2f(C.GLuint(arg0), C.GLfloat(arg1), C.GLfloat(arg2));
}

//extern void glVertexAttrib2fv (GLuint, const GLfloat *)
func VertexAttrib2fv(arg0 GLuint, arg1 *GLfloat) 
{
	C.glVertexAttrib2fv(C.GLuint(arg0), (*C.GLfloat)(arg1));
}

//extern void glVertexAttrib2s (GLuint, GLshort, GLshort)
func VertexAttrib2s(arg0 GLuint, arg1 GLshort, arg2 GLshort) 
{
	C.glVertexAttrib2s(C.GLuint(arg0), C.GLshort(arg1), C.GLshort(arg2));
}

//extern void glVertexAttrib2sv (GLuint, const GLshort *)
func VertexAttrib2sv(arg0 GLuint, arg1 *GLshort) 
{
	C.glVertexAttrib2sv(C.GLuint(arg0), (*C.GLshort)(arg1));
}

//extern void glVertexAttrib3d (GLuint, GLdouble, GLdouble, GLdouble)
func VertexAttrib3d(arg0 GLuint, arg1 GLdouble, arg2 GLdouble, arg3 GLdouble) 
{
	C.glVertexAttrib3d(C.GLuint(arg0), C.GLdouble(arg1), C.GLdouble(arg2), C.GLdouble(arg3));
}

//extern void glVertexAttrib3dv (GLuint, const GLdouble *)
func VertexAttrib3dv(arg0 GLuint, arg1 *GLdouble) 
{
	C.glVertexAttrib3dv(C.GLuint(arg0), (*C.GLdouble)(arg1));
}

//extern void glVertexAttrib3f (GLuint, GLfloat, GLfloat, GLfloat)
func VertexAttrib3f(arg0 GLuint, arg1 GLfloat, arg2 GLfloat, arg3 GLfloat) 
{
	C.glVertexAttrib3f(C.GLuint(arg0), C.GLfloat(arg1), C.GLfloat(arg2), C.GLfloat(arg3));
}

//extern void glVertexAttrib3fv (GLuint, const GLfloat *)
func VertexAttrib3fv(arg0 GLuint, arg1 *GLfloat) 
{
	C.glVertexAttrib3fv(C.GLuint(arg0), (*C.GLfloat)(arg1));
}

//extern void glVertexAttrib3s (GLuint, GLshort, GLshort, GLshort)
func VertexAttrib3s(arg0 GLuint, arg1 GLshort, arg2 GLshort, arg3 GLshort) 
{
	C.glVertexAttrib3s(C.GLuint(arg0), C.GLshort(arg1), C.GLshort(arg2), C.GLshort(arg3));
}

//extern void glVertexAttrib3sv (GLuint, const GLshort *)
func VertexAttrib3sv(arg0 GLuint, arg1 *GLshort) 
{
	C.glVertexAttrib3sv(C.GLuint(arg0), (*C.GLshort)(arg1));
}

//extern void glVertexAttrib4Nbv (GLuint, const GLbyte *)
func VertexAttrib4Nbv(arg0 GLuint, arg1 *GLbyte) 
{
	C.glVertexAttrib4Nbv(C.GLuint(arg0), (*C.GLbyte)(arg1));
}

//extern void glVertexAttrib4Niv (GLuint, const GLint *)
func VertexAttrib4Niv(arg0 GLuint, arg1 *GLint) 
{
	C.glVertexAttrib4Niv(C.GLuint(arg0), (*C.GLint)(arg1));
}

//extern void glVertexAttrib4Nsv (GLuint, const GLshort *)
func VertexAttrib4Nsv(arg0 GLuint, arg1 *GLshort) 
{
	C.glVertexAttrib4Nsv(C.GLuint(arg0), (*C.GLshort)(arg1));
}

//extern void glVertexAttrib4Nub (GLuint, GLubyte, GLubyte, GLubyte, GLubyte)
func VertexAttrib4Nub(arg0 GLuint, arg1 GLubyte, arg2 GLubyte, arg3 GLubyte, arg4 GLubyte) 
{
	C.glVertexAttrib4Nub(C.GLuint(arg0), C.GLubyte(arg1), C.GLubyte(arg2), C.GLubyte(arg3), C.GLubyte(arg4));
}

//extern void glVertexAttrib4Nubv (GLuint, const GLubyte *)
func VertexAttrib4Nubv(arg0 GLuint, arg1 *GLubyte) 
{
	C.glVertexAttrib4Nubv(C.GLuint(arg0), (*C.GLubyte)(arg1));
}

//extern void glVertexAttrib4Nuiv (GLuint, const GLuint *)
func VertexAttrib4Nuiv(arg0 GLuint, arg1 *GLuint) 
{
	C.glVertexAttrib4Nuiv(C.GLuint(arg0), (*C.GLuint)(arg1));
}

//extern void glVertexAttrib4Nusv (GLuint, const GLushort *)
func VertexAttrib4Nusv(arg0 GLuint, arg1 *GLushort) 
{
	C.glVertexAttrib4Nusv(C.GLuint(arg0), (*C.GLushort)(arg1));
}

//extern void glVertexAttrib4bv (GLuint, const GLbyte *)
func VertexAttrib4bv(arg0 GLuint, arg1 *GLbyte) 
{
	C.glVertexAttrib4bv(C.GLuint(arg0), (*C.GLbyte)(arg1));
}

//extern void glVertexAttrib4d (GLuint, GLdouble, GLdouble, GLdouble, GLdouble)
func VertexAttrib4d(arg0 GLuint, arg1 GLdouble, arg2 GLdouble, arg3 GLdouble, arg4 GLdouble) 
{
	C.glVertexAttrib4d(C.GLuint(arg0), C.GLdouble(arg1), C.GLdouble(arg2), C.GLdouble(arg3), C.GLdouble(arg4));
}

//extern void glVertexAttrib4dv (GLuint, const GLdouble *)
func VertexAttrib4dv(arg0 GLuint, arg1 *GLdouble) 
{
	C.glVertexAttrib4dv(C.GLuint(arg0), (*C.GLdouble)(arg1));
}

//extern void glVertexAttrib4f (GLuint, GLfloat, GLfloat, GLfloat, GLfloat)
func VertexAttrib4f(arg0 GLuint, arg1 GLfloat, arg2 GLfloat, arg3 GLfloat, arg4 GLfloat) 
{
	C.glVertexAttrib4f(C.GLuint(arg0), C.GLfloat(arg1), C.GLfloat(arg2), C.GLfloat(arg3), C.GLfloat(arg4));
}

//extern void glVertexAttrib4fv (GLuint, const GLfloat *)
func VertexAttrib4fv(arg0 GLuint, arg1 *GLfloat) 
{
	C.glVertexAttrib4fv(C.GLuint(arg0), (*C.GLfloat)(arg1));
}

//extern void glVertexAttrib4iv (GLuint, const GLint *)
func VertexAttrib4iv(arg0 GLuint, arg1 *GLint) 
{
	C.glVertexAttrib4iv(C.GLuint(arg0), (*C.GLint)(arg1));
}

//extern void glVertexAttrib4s (GLuint, GLshort, GLshort, GLshort, GLshort)
func VertexAttrib4s(arg0 GLuint, arg1 GLshort, arg2 GLshort, arg3 GLshort, arg4 GLshort) 
{
	C.glVertexAttrib4s(C.GLuint(arg0), C.GLshort(arg1), C.GLshort(arg2), C.GLshort(arg3), C.GLshort(arg4));
}

//extern void glVertexAttrib4sv (GLuint, const GLshort *)
func VertexAttrib4sv(arg0 GLuint, arg1 *GLshort) 
{
	C.glVertexAttrib4sv(C.GLuint(arg0), (*C.GLshort)(arg1));
}

//extern void glVertexAttrib4ubv (GLuint, const GLubyte *)
func VertexAttrib4ubv(arg0 GLuint, arg1 *GLubyte) 
{
	C.glVertexAttrib4ubv(C.GLuint(arg0), (*C.GLubyte)(arg1));
}

//extern void glVertexAttrib4uiv (GLuint, const GLuint *)
func VertexAttrib4uiv(arg0 GLuint, arg1 *GLuint) 
{
	C.glVertexAttrib4uiv(C.GLuint(arg0), (*C.GLuint)(arg1));
}

//extern void glVertexAttrib4usv (GLuint, const GLushort *)
func VertexAttrib4usv(arg0 GLuint, arg1 *GLushort) 
{
	C.glVertexAttrib4usv(C.GLuint(arg0), (*C.GLushort)(arg1));
}

//extern void glVertexAttribPointer (GLuint, GLint, GLenum, GLboolean, GLsizei, const GLvoid *)
func VertexAttribPointer(arg0 GLuint, arg1 GLint, arg2 GLenum, arg3 GLboolean, arg4 GLsizei, arg5 unsafe.Pointer) 
{
	C.glVertexAttribPointer(C.GLuint(arg0), C.GLint(arg1), C.GLenum(arg2), C.GLboolean(arg3), C.GLsizei(arg4), arg5);
}

//extern void glUniformMatrix2x3fv (GLint location, GLsizei count, GLboolean transpose, const GLfloat *value)
func UniformMatrix2x3fv(location GLint, count GLsizei, transpose GLboolean, value *GLfloat) 
{
	C.glUniformMatrix2x3fv(C.GLint(location), C.GLsizei(count), C.GLboolean(transpose), (*C.GLfloat)(value));
}

//extern void glUniformMatrix3x2fv (GLint location, GLsizei count, GLboolean transpose, const GLfloat *value)
func UniformMatrix3x2fv(location GLint, count GLsizei, transpose GLboolean, value *GLfloat) 
{
	C.glUniformMatrix3x2fv(C.GLint(location), C.GLsizei(count), C.GLboolean(transpose), (*C.GLfloat)(value));
}

//extern void glUniformMatrix2x4fv (GLint location, GLsizei count, GLboolean transpose, const GLfloat *value)
func UniformMatrix2x4fv(location GLint, count GLsizei, transpose GLboolean, value *GLfloat) 
{
	C.glUniformMatrix2x4fv(C.GLint(location), C.GLsizei(count), C.GLboolean(transpose), (*C.GLfloat)(value));
}

//extern void glUniformMatrix4x2fv (GLint location, GLsizei count, GLboolean transpose, const GLfloat *value)
func UniformMatrix4x2fv(location GLint, count GLsizei, transpose GLboolean, value *GLfloat) 
{
	C.glUniformMatrix4x2fv(C.GLint(location), C.GLsizei(count), C.GLboolean(transpose), (*C.GLfloat)(value));
}

//extern void glUniformMatrix3x4fv (GLint location, GLsizei count, GLboolean transpose, const GLfloat *value)
func UniformMatrix3x4fv(location GLint, count GLsizei, transpose GLboolean, value *GLfloat) 
{
	C.glUniformMatrix3x4fv(C.GLint(location), C.GLsizei(count), C.GLboolean(transpose), (*C.GLfloat)(value));
}

//extern void glUniformMatrix4x3fv (GLint location, GLsizei count, GLboolean transpose, const GLfloat *value)
func UniformMatrix4x3fv(location GLint, count GLsizei, transpose GLboolean, value *GLfloat) 
{
	C.glUniformMatrix4x3fv(C.GLint(location), C.GLsizei(count), C.GLboolean(transpose), (*C.GLfloat)(value));
}

//extern void glColorMaski (GLuint, GLboolean, GLboolean, GLboolean, GLboolean)
func ColorMaski(arg0 GLuint, arg1 GLboolean, arg2 GLboolean, arg3 GLboolean, arg4 GLboolean) 
{
	C.glColorMaski(C.GLuint(arg0), C.GLboolean(arg1), C.GLboolean(arg2), C.GLboolean(arg3), C.GLboolean(arg4));
}

//extern void glGetBooleani_v (GLenum, GLuint, GLboolean *)
func GetBooleani_v(arg0 GLenum, arg1 GLuint, arg2 *GLboolean) 
{
	C.glGetBooleani_v(C.GLenum(arg0), C.GLuint(arg1), (*C.GLboolean)(arg2));
}

//extern void glGetIntegeri_v (GLenum, GLuint, GLint *)
func GetIntegeri_v(arg0 GLenum, arg1 GLuint, arg2 *GLint) 
{
	C.glGetIntegeri_v(C.GLenum(arg0), C.GLuint(arg1), (*C.GLint)(arg2));
}

//extern void glEnablei (GLenum, GLuint)
func Enablei(arg0 GLenum, arg1 GLuint) 
{
	C.glEnablei(C.GLenum(arg0), C.GLuint(arg1));
}

//extern void glDisablei (GLenum, GLuint)
func Disablei(arg0 GLenum, arg1 GLuint) 
{
	C.glDisablei(C.GLenum(arg0), C.GLuint(arg1));
}

//extern GLboolean glIsEnabledi (GLenum, GLuint)
func IsEnabledi(arg0 GLenum, arg1 GLuint) GLboolean
{
	return GLboolean(C.glIsEnabledi(C.GLenum(arg0), C.GLuint(arg1)));
}

//extern void glBeginTransformFeedback (GLenum)
func BeginTransformFeedback(arg0 GLenum) 
{
	C.glBeginTransformFeedback(C.GLenum(arg0));
}

//extern void glEndTransformFeedback (void)
func EndTransformFeedback() 
{
	C.glEndTransformFeedback();
}

//extern void glBindBufferRange (GLenum, GLuint, GLuint, GLintptr, GLsizeiptr)
func BindBufferRange(arg0 GLenum, arg1 GLuint, arg2 GLuint, arg3 GLintptr, arg4 GLsizeiptr) 
{
	C.glBindBufferRange(C.GLenum(arg0), C.GLuint(arg1), C.GLuint(arg2), C.GLintptr(arg3), C.GLsizeiptr(arg4));
}

//extern void glBindBufferBase (GLenum, GLuint, GLuint)
func BindBufferBase(arg0 GLenum, arg1 GLuint, arg2 GLuint) 
{
	C.glBindBufferBase(C.GLenum(arg0), C.GLuint(arg1), C.GLuint(arg2));
}

//extern void glTransformFeedbackVaryings (GLuint, GLsizei, const GLchar **, GLenum)
func TransformFeedbackVaryings(arg0 GLuint, arg1 GLsizei, arg2 **GLchar, arg3 GLenum) 
{
	C.glTransformFeedbackVaryings(C.GLuint(arg0), C.GLsizei(arg1), (**C.GLchar)(arg2), C.GLenum(arg3));
}

//extern void glGetTransformFeedbackVarying (GLuint, GLuint, GLint *)
func GetTransformFeedbackVarying(arg0 GLuint, arg1 GLuint, arg2 *GLint) 
{
	C.glGetTransformFeedbackVarying(C.GLuint(arg0), C.GLuint(arg1), (*C.GLint)(arg2));
}

//extern void glClampColor (GLenum, GLenum)
func ClampColor(arg0 GLenum, arg1 GLenum) 
{
	C.glClampColor(C.GLenum(arg0), C.GLenum(arg1));
}

//extern void glBeginConditionalRender (GLuint, GLenum)
func BeginConditionalRender(arg0 GLuint, arg1 GLenum) 
{
	C.glBeginConditionalRender(C.GLuint(arg0), C.GLenum(arg1));
}

//extern void glEndConditionalRender (void)
func EndConditionalRender() 
{
	C.glEndConditionalRender();
}

//extern void glVertexAttribI1i (GLuint, GLint)
func VertexAttribI1i(arg0 GLuint, arg1 GLint) 
{
	C.glVertexAttribI1i(C.GLuint(arg0), C.GLint(arg1));
}

//extern void glVertexAttribI2i (GLuint, GLint, GLint)
func VertexAttribI2i(arg0 GLuint, arg1 GLint, arg2 GLint) 
{
	C.glVertexAttribI2i(C.GLuint(arg0), C.GLint(arg1), C.GLint(arg2));
}

//extern void glVertexAttribI3i (GLuint, GLint, GLint, GLint)
func VertexAttribI3i(arg0 GLuint, arg1 GLint, arg2 GLint, arg3 GLint) 
{
	C.glVertexAttribI3i(C.GLuint(arg0), C.GLint(arg1), C.GLint(arg2), C.GLint(arg3));
}

//extern void glVertexAttribI4i (GLuint, GLint, GLint, GLint, GLint)
func VertexAttribI4i(arg0 GLuint, arg1 GLint, arg2 GLint, arg3 GLint, arg4 GLint) 
{
	C.glVertexAttribI4i(C.GLuint(arg0), C.GLint(arg1), C.GLint(arg2), C.GLint(arg3), C.GLint(arg4));
}

//extern void glVertexAttribI1ui (GLuint, GLuint)
func VertexAttribI1ui(arg0 GLuint, arg1 GLuint) 
{
	C.glVertexAttribI1ui(C.GLuint(arg0), C.GLuint(arg1));
}

//extern void glVertexAttribI2ui (GLuint, GLuint, GLuint)
func VertexAttribI2ui(arg0 GLuint, arg1 GLuint, arg2 GLuint) 
{
	C.glVertexAttribI2ui(C.GLuint(arg0), C.GLuint(arg1), C.GLuint(arg2));
}

//extern void glVertexAttribI3ui (GLuint, GLuint, GLuint, GLuint)
func VertexAttribI3ui(arg0 GLuint, arg1 GLuint, arg2 GLuint, arg3 GLuint) 
{
	C.glVertexAttribI3ui(C.GLuint(arg0), C.GLuint(arg1), C.GLuint(arg2), C.GLuint(arg3));
}

//extern void glVertexAttribI4ui (GLuint, GLuint, GLuint, GLuint, GLuint)
func VertexAttribI4ui(arg0 GLuint, arg1 GLuint, arg2 GLuint, arg3 GLuint, arg4 GLuint) 
{
	C.glVertexAttribI4ui(C.GLuint(arg0), C.GLuint(arg1), C.GLuint(arg2), C.GLuint(arg3), C.GLuint(arg4));
}

//extern void glVertexAttribI1iv (GLuint, const GLint *)
func VertexAttribI1iv(arg0 GLuint, arg1 *GLint) 
{
	C.glVertexAttribI1iv(C.GLuint(arg0), (*C.GLint)(arg1));
}

//extern void glVertexAttribI2iv (GLuint, const GLint *)
func VertexAttribI2iv(arg0 GLuint, arg1 *GLint) 
{
	C.glVertexAttribI2iv(C.GLuint(arg0), (*C.GLint)(arg1));
}

//extern void glVertexAttribI3iv (GLuint, const GLint *)
func VertexAttribI3iv(arg0 GLuint, arg1 *GLint) 
{
	C.glVertexAttribI3iv(C.GLuint(arg0), (*C.GLint)(arg1));
}

//extern void glVertexAttribI4iv (GLuint, const GLint *)
func VertexAttribI4iv(arg0 GLuint, arg1 *GLint) 
{
	C.glVertexAttribI4iv(C.GLuint(arg0), (*C.GLint)(arg1));
}

//extern void glVertexAttribI1uiv (GLuint, const GLuint *)
func VertexAttribI1uiv(arg0 GLuint, arg1 *GLuint) 
{
	C.glVertexAttribI1uiv(C.GLuint(arg0), (*C.GLuint)(arg1));
}

//extern void glVertexAttribI2uiv (GLuint, const GLuint *)
func VertexAttribI2uiv(arg0 GLuint, arg1 *GLuint) 
{
	C.glVertexAttribI2uiv(C.GLuint(arg0), (*C.GLuint)(arg1));
}

//extern void glVertexAttribI3uiv (GLuint, const GLuint *)
func VertexAttribI3uiv(arg0 GLuint, arg1 *GLuint) 
{
	C.glVertexAttribI3uiv(C.GLuint(arg0), (*C.GLuint)(arg1));
}

//extern void glVertexAttribI4uiv (GLuint, const GLuint *)
func VertexAttribI4uiv(arg0 GLuint, arg1 *GLuint) 
{
	C.glVertexAttribI4uiv(C.GLuint(arg0), (*C.GLuint)(arg1));
}

//extern void glVertexAttribI4bv (GLuint, const GLbyte *)
func VertexAttribI4bv(arg0 GLuint, arg1 *GLbyte) 
{
	C.glVertexAttribI4bv(C.GLuint(arg0), (*C.GLbyte)(arg1));
}

//extern void glVertexAttribI4sv (GLuint, const GLshort *)
func VertexAttribI4sv(arg0 GLuint, arg1 *GLshort) 
{
	C.glVertexAttribI4sv(C.GLuint(arg0), (*C.GLshort)(arg1));
}

//extern void glVertexAttribI4ubv (GLuint, const GLubyte *)
func VertexAttribI4ubv(arg0 GLuint, arg1 *GLubyte) 
{
	C.glVertexAttribI4ubv(C.GLuint(arg0), (*C.GLubyte)(arg1));
}

//extern void glVertexAttribI4usv (GLuint, const GLushort *)
func VertexAttribI4usv(arg0 GLuint, arg1 *GLushort) 
{
	C.glVertexAttribI4usv(C.GLuint(arg0), (*C.GLushort)(arg1));
}

//extern void glVertexAttribIPointer (GLuint, GLint, GLenum, GLsizei, const GLvoid *)
func VertexAttribIPointer(arg0 GLuint, arg1 GLint, arg2 GLenum, arg3 GLsizei, arg4 unsafe.Pointer) 
{
	C.glVertexAttribIPointer(C.GLuint(arg0), C.GLint(arg1), C.GLenum(arg2), C.GLsizei(arg3), arg4);
}

//extern void glGetVertexAttribIiv (GLuint, GLenum, GLint *)
func GetVertexAttribIiv(arg0 GLuint, arg1 GLenum, arg2 *GLint) 
{
	C.glGetVertexAttribIiv(C.GLuint(arg0), C.GLenum(arg1), (*C.GLint)(arg2));
}

//extern void glGetVertexAttribIuiv (GLuint, GLenum, GLuint *)
func GetVertexAttribIuiv(arg0 GLuint, arg1 GLenum, arg2 *GLuint) 
{
	C.glGetVertexAttribIuiv(C.GLuint(arg0), C.GLenum(arg1), (*C.GLuint)(arg2));
}

//extern void glGetUniformuiv (GLuint, GLint, GLuint *)
func GetUniformuiv(arg0 GLuint, arg1 GLint, arg2 *GLuint) 
{
	C.glGetUniformuiv(C.GLuint(arg0), C.GLint(arg1), (*C.GLuint)(arg2));
}

//extern void glBindFragDataLocation (GLuint, GLuint, const GLchar *)
func BindFragDataLocation(arg0 GLuint, arg1 GLuint, arg2 string) 
{
	C.glBindFragDataLocation(C.GLuint(arg0), C.GLuint(arg1), (*_C_GLchar)((C.CString)(arg2)));
}

//extern GLint glGetFragDataLocation (GLuint, const GLchar *)
func GetFragDataLocation(arg0 GLuint, arg1 string) GLint
{
	return GLint(C.glGetFragDataLocation(C.GLuint(arg0), (*_C_GLchar)((C.CString)(arg1))));
}

//extern void glUniform1ui (GLint, GLuint)
func Uniform1ui(arg0 GLint, arg1 GLuint) 
{
	C.glUniform1ui(C.GLint(arg0), C.GLuint(arg1));
}

//extern void glUniform2ui (GLint, GLuint, GLuint)
func Uniform2ui(arg0 GLint, arg1 GLuint, arg2 GLuint) 
{
	C.glUniform2ui(C.GLint(arg0), C.GLuint(arg1), C.GLuint(arg2));
}

//extern void glUniform3ui (GLint, GLuint, GLuint, GLuint)
func Uniform3ui(arg0 GLint, arg1 GLuint, arg2 GLuint, arg3 GLuint) 
{
	C.glUniform3ui(C.GLint(arg0), C.GLuint(arg1), C.GLuint(arg2), C.GLuint(arg3));
}

//extern void glUniform4ui (GLint, GLuint, GLuint, GLuint, GLuint)
func Uniform4ui(arg0 GLint, arg1 GLuint, arg2 GLuint, arg3 GLuint, arg4 GLuint) 
{
	C.glUniform4ui(C.GLint(arg0), C.GLuint(arg1), C.GLuint(arg2), C.GLuint(arg3), C.GLuint(arg4));
}

//extern void glUniform1uiv (GLint, GLsizei, const GLuint *)
func Uniform1uiv(arg0 GLint, arg1 GLsizei, arg2 *GLuint) 
{
	C.glUniform1uiv(C.GLint(arg0), C.GLsizei(arg1), (*C.GLuint)(arg2));
}

//extern void glUniform2uiv (GLint, GLsizei, const GLuint *)
func Uniform2uiv(arg0 GLint, arg1 GLsizei, arg2 *GLuint) 
{
	C.glUniform2uiv(C.GLint(arg0), C.GLsizei(arg1), (*C.GLuint)(arg2));
}

//extern void glUniform3uiv (GLint, GLsizei, const GLuint *)
func Uniform3uiv(arg0 GLint, arg1 GLsizei, arg2 *GLuint) 
{
	C.glUniform3uiv(C.GLint(arg0), C.GLsizei(arg1), (*C.GLuint)(arg2));
}

//extern void glUniform4uiv (GLint, GLsizei, const GLuint *)
func Uniform4uiv(arg0 GLint, arg1 GLsizei, arg2 *GLuint) 
{
	C.glUniform4uiv(C.GLint(arg0), C.GLsizei(arg1), (*C.GLuint)(arg2));
}

//extern void glTexParameterIiv (GLenum, GLenum, const GLint *)
func TexParameterIiv(arg0 GLenum, arg1 GLenum, arg2 *GLint) 
{
	C.glTexParameterIiv(C.GLenum(arg0), C.GLenum(arg1), (*C.GLint)(arg2));
}

//extern void glTexParameterIuiv (GLenum, GLenum, const GLuint *)
func TexParameterIuiv(arg0 GLenum, arg1 GLenum, arg2 *GLuint) 
{
	C.glTexParameterIuiv(C.GLenum(arg0), C.GLenum(arg1), (*C.GLuint)(arg2));
}

//extern void glGetTexParameterIiv (GLenum, GLenum, GLint *)
func GetTexParameterIiv(arg0 GLenum, arg1 GLenum, arg2 *GLint) 
{
	C.glGetTexParameterIiv(C.GLenum(arg0), C.GLenum(arg1), (*C.GLint)(arg2));
}

//extern void glGetTexParameterIuiv (GLenum, GLenum, GLuint *)
func GetTexParameterIuiv(arg0 GLenum, arg1 GLenum, arg2 *GLuint) 
{
	C.glGetTexParameterIuiv(C.GLenum(arg0), C.GLenum(arg1), (*C.GLuint)(arg2));
}

//extern void glClearBufferiv (GLenum, GLint, const GLint *)
func ClearBufferiv(arg0 GLenum, arg1 GLint, arg2 *GLint) 
{
	C.glClearBufferiv(C.GLenum(arg0), C.GLint(arg1), (*C.GLint)(arg2));
}

//extern void glClearBufferuiv (GLenum, GLint, const GLuint *)
func ClearBufferuiv(arg0 GLenum, arg1 GLint, arg2 *GLuint) 
{
	C.glClearBufferuiv(C.GLenum(arg0), C.GLint(arg1), (*C.GLuint)(arg2));
}

//extern void glClearBufferfv (GLenum, GLint, const GLfloat *)
func ClearBufferfv(arg0 GLenum, arg1 GLint, arg2 *GLfloat) 
{
	C.glClearBufferfv(C.GLenum(arg0), C.GLint(arg1), (*C.GLfloat)(arg2));
}

//extern void glClearBufferfi (GLenum, GLint, GLfloat, GLint)
func ClearBufferfi(arg0 GLenum, arg1 GLint, arg2 GLfloat, arg3 GLint) 
{
	C.glClearBufferfi(C.GLenum(arg0), C.GLint(arg1), C.GLfloat(arg2), C.GLint(arg3));
}

//extern const GLubyte * glGetStringi (GLenum, GLuint)
func GetStringi(arg0 GLenum, arg1 GLuint) *GLubyte
{
	return (*GLubyte)(C.glGetStringi(C.GLenum(arg0), C.GLuint(arg1)));
}

//extern void glActiveTextureARB (GLenum)
func ActiveTextureARB(arg0 GLenum) 
{
	C.glActiveTextureARB(C.GLenum(arg0));
}

//extern void glClientActiveTextureARB (GLenum)
func ClientActiveTextureARB(arg0 GLenum) 
{
	C.glClientActiveTextureARB(C.GLenum(arg0));
}

//extern void glMultiTexCoord1dARB (GLenum, GLdouble)
func MultiTexCoord1dARB(arg0 GLenum, arg1 GLdouble) 
{
	C.glMultiTexCoord1dARB(C.GLenum(arg0), C.GLdouble(arg1));
}

//extern void glMultiTexCoord1dvARB (GLenum, const GLdouble *)
func MultiTexCoord1dvARB(arg0 GLenum, arg1 *GLdouble) 
{
	C.glMultiTexCoord1dvARB(C.GLenum(arg0), (*C.GLdouble)(arg1));
}

//extern void glMultiTexCoord1fARB (GLenum, GLfloat)
func MultiTexCoord1fARB(arg0 GLenum, arg1 GLfloat) 
{
	C.glMultiTexCoord1fARB(C.GLenum(arg0), C.GLfloat(arg1));
}

//extern void glMultiTexCoord1fvARB (GLenum, const GLfloat *)
func MultiTexCoord1fvARB(arg0 GLenum, arg1 *GLfloat) 
{
	C.glMultiTexCoord1fvARB(C.GLenum(arg0), (*C.GLfloat)(arg1));
}

//extern void glMultiTexCoord1iARB (GLenum, GLint)
func MultiTexCoord1iARB(arg0 GLenum, arg1 GLint) 
{
	C.glMultiTexCoord1iARB(C.GLenum(arg0), C.GLint(arg1));
}

//extern void glMultiTexCoord1ivARB (GLenum, const GLint *)
func MultiTexCoord1ivARB(arg0 GLenum, arg1 *GLint) 
{
	C.glMultiTexCoord1ivARB(C.GLenum(arg0), (*C.GLint)(arg1));
}

//extern void glMultiTexCoord1sARB (GLenum, GLshort)
func MultiTexCoord1sARB(arg0 GLenum, arg1 GLshort) 
{
	C.glMultiTexCoord1sARB(C.GLenum(arg0), C.GLshort(arg1));
}

//extern void glMultiTexCoord1svARB (GLenum, const GLshort *)
func MultiTexCoord1svARB(arg0 GLenum, arg1 *GLshort) 
{
	C.glMultiTexCoord1svARB(C.GLenum(arg0), (*C.GLshort)(arg1));
}

//extern void glMultiTexCoord2dARB (GLenum, GLdouble, GLdouble)
func MultiTexCoord2dARB(arg0 GLenum, arg1 GLdouble, arg2 GLdouble) 
{
	C.glMultiTexCoord2dARB(C.GLenum(arg0), C.GLdouble(arg1), C.GLdouble(arg2));
}

//extern void glMultiTexCoord2dvARB (GLenum, const GLdouble *)
func MultiTexCoord2dvARB(arg0 GLenum, arg1 *GLdouble) 
{
	C.glMultiTexCoord2dvARB(C.GLenum(arg0), (*C.GLdouble)(arg1));
}

//extern void glMultiTexCoord2fARB (GLenum, GLfloat, GLfloat)
func MultiTexCoord2fARB(arg0 GLenum, arg1 GLfloat, arg2 GLfloat) 
{
	C.glMultiTexCoord2fARB(C.GLenum(arg0), C.GLfloat(arg1), C.GLfloat(arg2));
}

//extern void glMultiTexCoord2fvARB (GLenum, const GLfloat *)
func MultiTexCoord2fvARB(arg0 GLenum, arg1 *GLfloat) 
{
	C.glMultiTexCoord2fvARB(C.GLenum(arg0), (*C.GLfloat)(arg1));
}

//extern void glMultiTexCoord2iARB (GLenum, GLint, GLint)
func MultiTexCoord2iARB(arg0 GLenum, arg1 GLint, arg2 GLint) 
{
	C.glMultiTexCoord2iARB(C.GLenum(arg0), C.GLint(arg1), C.GLint(arg2));
}

//extern void glMultiTexCoord2ivARB (GLenum, const GLint *)
func MultiTexCoord2ivARB(arg0 GLenum, arg1 *GLint) 
{
	C.glMultiTexCoord2ivARB(C.GLenum(arg0), (*C.GLint)(arg1));
}

//extern void glMultiTexCoord2sARB (GLenum, GLshort, GLshort)
func MultiTexCoord2sARB(arg0 GLenum, arg1 GLshort, arg2 GLshort) 
{
	C.glMultiTexCoord2sARB(C.GLenum(arg0), C.GLshort(arg1), C.GLshort(arg2));
}

//extern void glMultiTexCoord2svARB (GLenum, const GLshort *)
func MultiTexCoord2svARB(arg0 GLenum, arg1 *GLshort) 
{
	C.glMultiTexCoord2svARB(C.GLenum(arg0), (*C.GLshort)(arg1));
}

//extern void glMultiTexCoord3dARB (GLenum, GLdouble, GLdouble, GLdouble)
func MultiTexCoord3dARB(arg0 GLenum, arg1 GLdouble, arg2 GLdouble, arg3 GLdouble) 
{
	C.glMultiTexCoord3dARB(C.GLenum(arg0), C.GLdouble(arg1), C.GLdouble(arg2), C.GLdouble(arg3));
}

//extern void glMultiTexCoord3dvARB (GLenum, const GLdouble *)
func MultiTexCoord3dvARB(arg0 GLenum, arg1 *GLdouble) 
{
	C.glMultiTexCoord3dvARB(C.GLenum(arg0), (*C.GLdouble)(arg1));
}

//extern void glMultiTexCoord3fARB (GLenum, GLfloat, GLfloat, GLfloat)
func MultiTexCoord3fARB(arg0 GLenum, arg1 GLfloat, arg2 GLfloat, arg3 GLfloat) 
{
	C.glMultiTexCoord3fARB(C.GLenum(arg0), C.GLfloat(arg1), C.GLfloat(arg2), C.GLfloat(arg3));
}

//extern void glMultiTexCoord3fvARB (GLenum, const GLfloat *)
func MultiTexCoord3fvARB(arg0 GLenum, arg1 *GLfloat) 
{
	C.glMultiTexCoord3fvARB(C.GLenum(arg0), (*C.GLfloat)(arg1));
}

//extern void glMultiTexCoord3iARB (GLenum, GLint, GLint, GLint)
func MultiTexCoord3iARB(arg0 GLenum, arg1 GLint, arg2 GLint, arg3 GLint) 
{
	C.glMultiTexCoord3iARB(C.GLenum(arg0), C.GLint(arg1), C.GLint(arg2), C.GLint(arg3));
}

//extern void glMultiTexCoord3ivARB (GLenum, const GLint *)
func MultiTexCoord3ivARB(arg0 GLenum, arg1 *GLint) 
{
	C.glMultiTexCoord3ivARB(C.GLenum(arg0), (*C.GLint)(arg1));
}

//extern void glMultiTexCoord3sARB (GLenum, GLshort, GLshort, GLshort)
func MultiTexCoord3sARB(arg0 GLenum, arg1 GLshort, arg2 GLshort, arg3 GLshort) 
{
	C.glMultiTexCoord3sARB(C.GLenum(arg0), C.GLshort(arg1), C.GLshort(arg2), C.GLshort(arg3));
}

//extern void glMultiTexCoord3svARB (GLenum, const GLshort *)
func MultiTexCoord3svARB(arg0 GLenum, arg1 *GLshort) 
{
	C.glMultiTexCoord3svARB(C.GLenum(arg0), (*C.GLshort)(arg1));
}

//extern void glMultiTexCoord4dARB (GLenum, GLdouble, GLdouble, GLdouble, GLdouble)
func MultiTexCoord4dARB(arg0 GLenum, arg1 GLdouble, arg2 GLdouble, arg3 GLdouble, arg4 GLdouble) 
{
	C.glMultiTexCoord4dARB(C.GLenum(arg0), C.GLdouble(arg1), C.GLdouble(arg2), C.GLdouble(arg3), C.GLdouble(arg4));
}

//extern void glMultiTexCoord4dvARB (GLenum, const GLdouble *)
func MultiTexCoord4dvARB(arg0 GLenum, arg1 *GLdouble) 
{
	C.glMultiTexCoord4dvARB(C.GLenum(arg0), (*C.GLdouble)(arg1));
}

//extern void glMultiTexCoord4fARB (GLenum, GLfloat, GLfloat, GLfloat, GLfloat)
func MultiTexCoord4fARB(arg0 GLenum, arg1 GLfloat, arg2 GLfloat, arg3 GLfloat, arg4 GLfloat) 
{
	C.glMultiTexCoord4fARB(C.GLenum(arg0), C.GLfloat(arg1), C.GLfloat(arg2), C.GLfloat(arg3), C.GLfloat(arg4));
}

//extern void glMultiTexCoord4fvARB (GLenum, const GLfloat *)
func MultiTexCoord4fvARB(arg0 GLenum, arg1 *GLfloat) 
{
	C.glMultiTexCoord4fvARB(C.GLenum(arg0), (*C.GLfloat)(arg1));
}

//extern void glMultiTexCoord4iARB (GLenum, GLint, GLint, GLint, GLint)
func MultiTexCoord4iARB(arg0 GLenum, arg1 GLint, arg2 GLint, arg3 GLint, arg4 GLint) 
{
	C.glMultiTexCoord4iARB(C.GLenum(arg0), C.GLint(arg1), C.GLint(arg2), C.GLint(arg3), C.GLint(arg4));
}

//extern void glMultiTexCoord4ivARB (GLenum, const GLint *)
func MultiTexCoord4ivARB(arg0 GLenum, arg1 *GLint) 
{
	C.glMultiTexCoord4ivARB(C.GLenum(arg0), (*C.GLint)(arg1));
}

//extern void glMultiTexCoord4sARB (GLenum, GLshort, GLshort, GLshort, GLshort)
func MultiTexCoord4sARB(arg0 GLenum, arg1 GLshort, arg2 GLshort, arg3 GLshort, arg4 GLshort) 
{
	C.glMultiTexCoord4sARB(C.GLenum(arg0), C.GLshort(arg1), C.GLshort(arg2), C.GLshort(arg3), C.GLshort(arg4));
}

//extern void glMultiTexCoord4svARB (GLenum, const GLshort *)
func MultiTexCoord4svARB(arg0 GLenum, arg1 *GLshort) 
{
	C.glMultiTexCoord4svARB(C.GLenum(arg0), (*C.GLshort)(arg1));
}

//extern void glLoadTransposeMatrixfARB (const GLfloat *)
func LoadTransposeMatrixfARB(arg0 *GLfloat) 
{
	C.glLoadTransposeMatrixfARB((*C.GLfloat)(arg0));
}

//extern void glLoadTransposeMatrixdARB (const GLdouble *)
func LoadTransposeMatrixdARB(arg0 *GLdouble) 
{
	C.glLoadTransposeMatrixdARB((*C.GLdouble)(arg0));
}

//extern void glMultTransposeMatrixfARB (const GLfloat *)
func MultTransposeMatrixfARB(arg0 *GLfloat) 
{
	C.glMultTransposeMatrixfARB((*C.GLfloat)(arg0));
}

//extern void glMultTransposeMatrixdARB (const GLdouble *)
func MultTransposeMatrixdARB(arg0 *GLdouble) 
{
	C.glMultTransposeMatrixdARB((*C.GLdouble)(arg0));
}

//extern void glSampleCoverageARB (GLclampf, GLboolean)
func SampleCoverageARB(arg0 GLclampf, arg1 GLboolean) 
{
	C.glSampleCoverageARB(C.GLclampf(arg0), C.GLboolean(arg1));
}

//extern void glCompressedTexImage3DARB (GLenum, GLint, GLenum, GLsizei, GLsizei, GLsizei, GLint, GLsizei, const GLvoid *)
func CompressedTexImage3DARB(arg0 GLenum, arg1 GLint, arg2 GLenum, arg3 GLsizei, arg4 GLsizei, arg5 GLsizei, arg6 GLint, arg7 GLsizei, arg8 unsafe.Pointer) 
{
	C.glCompressedTexImage3DARB(C.GLenum(arg0), C.GLint(arg1), C.GLenum(arg2), C.GLsizei(arg3), C.GLsizei(arg4), C.GLsizei(arg5), C.GLint(arg6), C.GLsizei(arg7), arg8);
}

//extern void glCompressedTexImage2DARB (GLenum, GLint, GLenum, GLsizei, GLsizei, GLint, GLsizei, const GLvoid *)
func CompressedTexImage2DARB(arg0 GLenum, arg1 GLint, arg2 GLenum, arg3 GLsizei, arg4 GLsizei, arg5 GLint, arg6 GLsizei, arg7 unsafe.Pointer) 
{
	C.glCompressedTexImage2DARB(C.GLenum(arg0), C.GLint(arg1), C.GLenum(arg2), C.GLsizei(arg3), C.GLsizei(arg4), C.GLint(arg5), C.GLsizei(arg6), arg7);
}

//extern void glCompressedTexImage1DARB (GLenum, GLint, GLenum, GLsizei, GLint, GLsizei, const GLvoid *)
func CompressedTexImage1DARB(arg0 GLenum, arg1 GLint, arg2 GLenum, arg3 GLsizei, arg4 GLint, arg5 GLsizei, arg6 unsafe.Pointer) 
{
	C.glCompressedTexImage1DARB(C.GLenum(arg0), C.GLint(arg1), C.GLenum(arg2), C.GLsizei(arg3), C.GLint(arg4), C.GLsizei(arg5), arg6);
}

//extern void glCompressedTexSubImage3DARB (GLenum, GLint, GLint, GLint, GLint, GLsizei, GLsizei, GLsizei, GLenum, GLsizei, const GLvoid *)
func CompressedTexSubImage3DARB(arg0 GLenum, arg1 GLint, arg2 GLint, arg3 GLint, arg4 GLint, arg5 GLsizei, arg6 GLsizei, arg7 GLsizei, arg8 GLenum, arg9 GLsizei, arg10 unsafe.Pointer) 
{
	C.glCompressedTexSubImage3DARB(C.GLenum(arg0), C.GLint(arg1), C.GLint(arg2), C.GLint(arg3), C.GLint(arg4), C.GLsizei(arg5), C.GLsizei(arg6), C.GLsizei(arg7), C.GLenum(arg8), C.GLsizei(arg9), arg10);
}

//extern void glCompressedTexSubImage2DARB (GLenum, GLint, GLint, GLint, GLsizei, GLsizei, GLenum, GLsizei, const GLvoid *)
func CompressedTexSubImage2DARB(arg0 GLenum, arg1 GLint, arg2 GLint, arg3 GLint, arg4 GLsizei, arg5 GLsizei, arg6 GLenum, arg7 GLsizei, arg8 unsafe.Pointer) 
{
	C.glCompressedTexSubImage2DARB(C.GLenum(arg0), C.GLint(arg1), C.GLint(arg2), C.GLint(arg3), C.GLsizei(arg4), C.GLsizei(arg5), C.GLenum(arg6), C.GLsizei(arg7), arg8);
}

//extern void glCompressedTexSubImage1DARB (GLenum, GLint, GLint, GLsizei, GLenum, GLsizei, const GLvoid *)
func CompressedTexSubImage1DARB(arg0 GLenum, arg1 GLint, arg2 GLint, arg3 GLsizei, arg4 GLenum, arg5 GLsizei, arg6 unsafe.Pointer) 
{
	C.glCompressedTexSubImage1DARB(C.GLenum(arg0), C.GLint(arg1), C.GLint(arg2), C.GLsizei(arg3), C.GLenum(arg4), C.GLsizei(arg5), arg6);
}

//extern void glGetCompressedTexImageARB (GLenum, GLint, GLvoid *)
func GetCompressedTexImageARB(arg0 GLenum, arg1 GLint, arg2 unsafe.Pointer) 
{
	C.glGetCompressedTexImageARB(C.GLenum(arg0), C.GLint(arg1), arg2);
}

//extern void glPointParameterfARB (GLenum, GLfloat)
func PointParameterfARB(arg0 GLenum, arg1 GLfloat) 
{
	C.glPointParameterfARB(C.GLenum(arg0), C.GLfloat(arg1));
}

//extern void glPointParameterfvARB (GLenum, const GLfloat *)
func PointParameterfvARB(arg0 GLenum, arg1 *GLfloat) 
{
	C.glPointParameterfvARB(C.GLenum(arg0), (*C.GLfloat)(arg1));
}

//extern void glWeightbvARB (GLint, const GLbyte *)
func WeightbvARB(arg0 GLint, arg1 *GLbyte) 
{
	C.glWeightbvARB(C.GLint(arg0), (*C.GLbyte)(arg1));
}

//extern void glWeightsvARB (GLint, const GLshort *)
func WeightsvARB(arg0 GLint, arg1 *GLshort) 
{
	C.glWeightsvARB(C.GLint(arg0), (*C.GLshort)(arg1));
}

//extern void glWeightivARB (GLint, const GLint *)
func WeightivARB(arg0 GLint, arg1 *GLint) 
{
	C.glWeightivARB(C.GLint(arg0), (*C.GLint)(arg1));
}

//extern void glWeightfvARB (GLint, const GLfloat *)
func WeightfvARB(arg0 GLint, arg1 *GLfloat) 
{
	C.glWeightfvARB(C.GLint(arg0), (*C.GLfloat)(arg1));
}

//extern void glWeightdvARB (GLint, const GLdouble *)
func WeightdvARB(arg0 GLint, arg1 *GLdouble) 
{
	C.glWeightdvARB(C.GLint(arg0), (*C.GLdouble)(arg1));
}

//extern void glWeightubvARB (GLint, const GLubyte *)
func WeightubvARB(arg0 GLint, arg1 *GLubyte) 
{
	C.glWeightubvARB(C.GLint(arg0), (*C.GLubyte)(arg1));
}

//extern void glWeightusvARB (GLint, const GLushort *)
func WeightusvARB(arg0 GLint, arg1 *GLushort) 
{
	C.glWeightusvARB(C.GLint(arg0), (*C.GLushort)(arg1));
}

//extern void glWeightuivARB (GLint, const GLuint *)
func WeightuivARB(arg0 GLint, arg1 *GLuint) 
{
	C.glWeightuivARB(C.GLint(arg0), (*C.GLuint)(arg1));
}

//extern void glWeightPointerARB (GLint, GLenum, GLsizei, const GLvoid *)
func WeightPointerARB(arg0 GLint, arg1 GLenum, arg2 GLsizei, arg3 unsafe.Pointer) 
{
	C.glWeightPointerARB(C.GLint(arg0), C.GLenum(arg1), C.GLsizei(arg2), arg3);
}

//extern void glVertexBlendARB (GLint)
func VertexBlendARB(arg0 GLint) 
{
	C.glVertexBlendARB(C.GLint(arg0));
}

//extern void glCurrentPaletteMatrixARB (GLint)
func CurrentPaletteMatrixARB(arg0 GLint) 
{
	C.glCurrentPaletteMatrixARB(C.GLint(arg0));
}

//extern void glMatrixIndexubvARB (GLint, const GLubyte *)
func MatrixIndexubvARB(arg0 GLint, arg1 *GLubyte) 
{
	C.glMatrixIndexubvARB(C.GLint(arg0), (*C.GLubyte)(arg1));
}

//extern void glMatrixIndexusvARB (GLint, const GLushort *)
func MatrixIndexusvARB(arg0 GLint, arg1 *GLushort) 
{
	C.glMatrixIndexusvARB(C.GLint(arg0), (*C.GLushort)(arg1));
}

//extern void glMatrixIndexuivARB (GLint, const GLuint *)
func MatrixIndexuivARB(arg0 GLint, arg1 *GLuint) 
{
	C.glMatrixIndexuivARB(C.GLint(arg0), (*C.GLuint)(arg1));
}

//extern void glMatrixIndexPointerARB (GLint, GLenum, GLsizei, const GLvoid *)
func MatrixIndexPointerARB(arg0 GLint, arg1 GLenum, arg2 GLsizei, arg3 unsafe.Pointer) 
{
	C.glMatrixIndexPointerARB(C.GLint(arg0), C.GLenum(arg1), C.GLsizei(arg2), arg3);
}

//extern void glWindowPos2dARB (GLdouble, GLdouble)
func WindowPos2dARB(arg0 GLdouble, arg1 GLdouble) 
{
	C.glWindowPos2dARB(C.GLdouble(arg0), C.GLdouble(arg1));
}

//extern void glWindowPos2dvARB (const GLdouble *)
func WindowPos2dvARB(arg0 *GLdouble) 
{
	C.glWindowPos2dvARB((*C.GLdouble)(arg0));
}

//extern void glWindowPos2fARB (GLfloat, GLfloat)
func WindowPos2fARB(arg0 GLfloat, arg1 GLfloat) 
{
	C.glWindowPos2fARB(C.GLfloat(arg0), C.GLfloat(arg1));
}

//extern void glWindowPos2fvARB (const GLfloat *)
func WindowPos2fvARB(arg0 *GLfloat) 
{
	C.glWindowPos2fvARB((*C.GLfloat)(arg0));
}

//extern void glWindowPos2iARB (GLint, GLint)
func WindowPos2iARB(arg0 GLint, arg1 GLint) 
{
	C.glWindowPos2iARB(C.GLint(arg0), C.GLint(arg1));
}

//extern void glWindowPos2ivARB (const GLint *)
func WindowPos2ivARB(arg0 *GLint) 
{
	C.glWindowPos2ivARB((*C.GLint)(arg0));
}

//extern void glWindowPos2sARB (GLshort, GLshort)
func WindowPos2sARB(arg0 GLshort, arg1 GLshort) 
{
	C.glWindowPos2sARB(C.GLshort(arg0), C.GLshort(arg1));
}

//extern void glWindowPos2svARB (const GLshort *)
func WindowPos2svARB(arg0 *GLshort) 
{
	C.glWindowPos2svARB((*C.GLshort)(arg0));
}

//extern void glWindowPos3dARB (GLdouble, GLdouble, GLdouble)
func WindowPos3dARB(arg0 GLdouble, arg1 GLdouble, arg2 GLdouble) 
{
	C.glWindowPos3dARB(C.GLdouble(arg0), C.GLdouble(arg1), C.GLdouble(arg2));
}

//extern void glWindowPos3dvARB (const GLdouble *)
func WindowPos3dvARB(arg0 *GLdouble) 
{
	C.glWindowPos3dvARB((*C.GLdouble)(arg0));
}

//extern void glWindowPos3fARB (GLfloat, GLfloat, GLfloat)
func WindowPos3fARB(arg0 GLfloat, arg1 GLfloat, arg2 GLfloat) 
{
	C.glWindowPos3fARB(C.GLfloat(arg0), C.GLfloat(arg1), C.GLfloat(arg2));
}

//extern void glWindowPos3fvARB (const GLfloat *)
func WindowPos3fvARB(arg0 *GLfloat) 
{
	C.glWindowPos3fvARB((*C.GLfloat)(arg0));
}

//extern void glWindowPos3iARB (GLint, GLint, GLint)
func WindowPos3iARB(arg0 GLint, arg1 GLint, arg2 GLint) 
{
	C.glWindowPos3iARB(C.GLint(arg0), C.GLint(arg1), C.GLint(arg2));
}

//extern void glWindowPos3ivARB (const GLint *)
func WindowPos3ivARB(arg0 *GLint) 
{
	C.glWindowPos3ivARB((*C.GLint)(arg0));
}

//extern void glWindowPos3sARB (GLshort, GLshort, GLshort)
func WindowPos3sARB(arg0 GLshort, arg1 GLshort, arg2 GLshort) 
{
	C.glWindowPos3sARB(C.GLshort(arg0), C.GLshort(arg1), C.GLshort(arg2));
}

//extern void glWindowPos3svARB (const GLshort *)
func WindowPos3svARB(arg0 *GLshort) 
{
	C.glWindowPos3svARB((*C.GLshort)(arg0));
}

//extern void glVertexAttrib1dARB (GLuint, GLdouble)
func VertexAttrib1dARB(arg0 GLuint, arg1 GLdouble) 
{
	C.glVertexAttrib1dARB(C.GLuint(arg0), C.GLdouble(arg1));
}

//extern void glVertexAttrib1dvARB (GLuint, const GLdouble *)
func VertexAttrib1dvARB(arg0 GLuint, arg1 *GLdouble) 
{
	C.glVertexAttrib1dvARB(C.GLuint(arg0), (*C.GLdouble)(arg1));
}

//extern void glVertexAttrib1fARB (GLuint, GLfloat)
func VertexAttrib1fARB(arg0 GLuint, arg1 GLfloat) 
{
	C.glVertexAttrib1fARB(C.GLuint(arg0), C.GLfloat(arg1));
}

//extern void glVertexAttrib1fvARB (GLuint, const GLfloat *)
func VertexAttrib1fvARB(arg0 GLuint, arg1 *GLfloat) 
{
	C.glVertexAttrib1fvARB(C.GLuint(arg0), (*C.GLfloat)(arg1));
}

//extern void glVertexAttrib1sARB (GLuint, GLshort)
func VertexAttrib1sARB(arg0 GLuint, arg1 GLshort) 
{
	C.glVertexAttrib1sARB(C.GLuint(arg0), C.GLshort(arg1));
}

//extern void glVertexAttrib1svARB (GLuint, const GLshort *)
func VertexAttrib1svARB(arg0 GLuint, arg1 *GLshort) 
{
	C.glVertexAttrib1svARB(C.GLuint(arg0), (*C.GLshort)(arg1));
}

//extern void glVertexAttrib2dARB (GLuint, GLdouble, GLdouble)
func VertexAttrib2dARB(arg0 GLuint, arg1 GLdouble, arg2 GLdouble) 
{
	C.glVertexAttrib2dARB(C.GLuint(arg0), C.GLdouble(arg1), C.GLdouble(arg2));
}

//extern void glVertexAttrib2dvARB (GLuint, const GLdouble *)
func VertexAttrib2dvARB(arg0 GLuint, arg1 *GLdouble) 
{
	C.glVertexAttrib2dvARB(C.GLuint(arg0), (*C.GLdouble)(arg1));
}

//extern void glVertexAttrib2fARB (GLuint, GLfloat, GLfloat)
func VertexAttrib2fARB(arg0 GLuint, arg1 GLfloat, arg2 GLfloat) 
{
	C.glVertexAttrib2fARB(C.GLuint(arg0), C.GLfloat(arg1), C.GLfloat(arg2));
}

//extern void glVertexAttrib2fvARB (GLuint, const GLfloat *)
func VertexAttrib2fvARB(arg0 GLuint, arg1 *GLfloat) 
{
	C.glVertexAttrib2fvARB(C.GLuint(arg0), (*C.GLfloat)(arg1));
}

//extern void glVertexAttrib2sARB (GLuint, GLshort, GLshort)
func VertexAttrib2sARB(arg0 GLuint, arg1 GLshort, arg2 GLshort) 
{
	C.glVertexAttrib2sARB(C.GLuint(arg0), C.GLshort(arg1), C.GLshort(arg2));
}

//extern void glVertexAttrib2svARB (GLuint, const GLshort *)
func VertexAttrib2svARB(arg0 GLuint, arg1 *GLshort) 
{
	C.glVertexAttrib2svARB(C.GLuint(arg0), (*C.GLshort)(arg1));
}

//extern void glVertexAttrib3dARB (GLuint, GLdouble, GLdouble, GLdouble)
func VertexAttrib3dARB(arg0 GLuint, arg1 GLdouble, arg2 GLdouble, arg3 GLdouble) 
{
	C.glVertexAttrib3dARB(C.GLuint(arg0), C.GLdouble(arg1), C.GLdouble(arg2), C.GLdouble(arg3));
}

//extern void glVertexAttrib3dvARB (GLuint, const GLdouble *)
func VertexAttrib3dvARB(arg0 GLuint, arg1 *GLdouble) 
{
	C.glVertexAttrib3dvARB(C.GLuint(arg0), (*C.GLdouble)(arg1));
}

//extern void glVertexAttrib3fARB (GLuint, GLfloat, GLfloat, GLfloat)
func VertexAttrib3fARB(arg0 GLuint, arg1 GLfloat, arg2 GLfloat, arg3 GLfloat) 
{
	C.glVertexAttrib3fARB(C.GLuint(arg0), C.GLfloat(arg1), C.GLfloat(arg2), C.GLfloat(arg3));
}

//extern void glVertexAttrib3fvARB (GLuint, const GLfloat *)
func VertexAttrib3fvARB(arg0 GLuint, arg1 *GLfloat) 
{
	C.glVertexAttrib3fvARB(C.GLuint(arg0), (*C.GLfloat)(arg1));
}

//extern void glVertexAttrib3sARB (GLuint, GLshort, GLshort, GLshort)
func VertexAttrib3sARB(arg0 GLuint, arg1 GLshort, arg2 GLshort, arg3 GLshort) 
{
	C.glVertexAttrib3sARB(C.GLuint(arg0), C.GLshort(arg1), C.GLshort(arg2), C.GLshort(arg3));
}

//extern void glVertexAttrib3svARB (GLuint, const GLshort *)
func VertexAttrib3svARB(arg0 GLuint, arg1 *GLshort) 
{
	C.glVertexAttrib3svARB(C.GLuint(arg0), (*C.GLshort)(arg1));
}

//extern void glVertexAttrib4NbvARB (GLuint, const GLbyte *)
func VertexAttrib4NbvARB(arg0 GLuint, arg1 *GLbyte) 
{
	C.glVertexAttrib4NbvARB(C.GLuint(arg0), (*C.GLbyte)(arg1));
}

//extern void glVertexAttrib4NivARB (GLuint, const GLint *)
func VertexAttrib4NivARB(arg0 GLuint, arg1 *GLint) 
{
	C.glVertexAttrib4NivARB(C.GLuint(arg0), (*C.GLint)(arg1));
}

//extern void glVertexAttrib4NsvARB (GLuint, const GLshort *)
func VertexAttrib4NsvARB(arg0 GLuint, arg1 *GLshort) 
{
	C.glVertexAttrib4NsvARB(C.GLuint(arg0), (*C.GLshort)(arg1));
}

//extern void glVertexAttrib4NubARB (GLuint, GLubyte, GLubyte, GLubyte, GLubyte)
func VertexAttrib4NubARB(arg0 GLuint, arg1 GLubyte, arg2 GLubyte, arg3 GLubyte, arg4 GLubyte) 
{
	C.glVertexAttrib4NubARB(C.GLuint(arg0), C.GLubyte(arg1), C.GLubyte(arg2), C.GLubyte(arg3), C.GLubyte(arg4));
}

//extern void glVertexAttrib4NubvARB (GLuint, const GLubyte *)
func VertexAttrib4NubvARB(arg0 GLuint, arg1 *GLubyte) 
{
	C.glVertexAttrib4NubvARB(C.GLuint(arg0), (*C.GLubyte)(arg1));
}

//extern void glVertexAttrib4NuivARB (GLuint, const GLuint *)
func VertexAttrib4NuivARB(arg0 GLuint, arg1 *GLuint) 
{
	C.glVertexAttrib4NuivARB(C.GLuint(arg0), (*C.GLuint)(arg1));
}

//extern void glVertexAttrib4NusvARB (GLuint, const GLushort *)
func VertexAttrib4NusvARB(arg0 GLuint, arg1 *GLushort) 
{
	C.glVertexAttrib4NusvARB(C.GLuint(arg0), (*C.GLushort)(arg1));
}

//extern void glVertexAttrib4bvARB (GLuint, const GLbyte *)
func VertexAttrib4bvARB(arg0 GLuint, arg1 *GLbyte) 
{
	C.glVertexAttrib4bvARB(C.GLuint(arg0), (*C.GLbyte)(arg1));
}

//extern void glVertexAttrib4dARB (GLuint, GLdouble, GLdouble, GLdouble, GLdouble)
func VertexAttrib4dARB(arg0 GLuint, arg1 GLdouble, arg2 GLdouble, arg3 GLdouble, arg4 GLdouble) 
{
	C.glVertexAttrib4dARB(C.GLuint(arg0), C.GLdouble(arg1), C.GLdouble(arg2), C.GLdouble(arg3), C.GLdouble(arg4));
}

//extern void glVertexAttrib4dvARB (GLuint, const GLdouble *)
func VertexAttrib4dvARB(arg0 GLuint, arg1 *GLdouble) 
{
	C.glVertexAttrib4dvARB(C.GLuint(arg0), (*C.GLdouble)(arg1));
}

//extern void glVertexAttrib4fARB (GLuint, GLfloat, GLfloat, GLfloat, GLfloat)
func VertexAttrib4fARB(arg0 GLuint, arg1 GLfloat, arg2 GLfloat, arg3 GLfloat, arg4 GLfloat) 
{
	C.glVertexAttrib4fARB(C.GLuint(arg0), C.GLfloat(arg1), C.GLfloat(arg2), C.GLfloat(arg3), C.GLfloat(arg4));
}

//extern void glVertexAttrib4fvARB (GLuint, const GLfloat *)
func VertexAttrib4fvARB(arg0 GLuint, arg1 *GLfloat) 
{
	C.glVertexAttrib4fvARB(C.GLuint(arg0), (*C.GLfloat)(arg1));
}

//extern void glVertexAttrib4ivARB (GLuint, const GLint *)
func VertexAttrib4ivARB(arg0 GLuint, arg1 *GLint) 
{
	C.glVertexAttrib4ivARB(C.GLuint(arg0), (*C.GLint)(arg1));
}

//extern void glVertexAttrib4sARB (GLuint, GLshort, GLshort, GLshort, GLshort)
func VertexAttrib4sARB(arg0 GLuint, arg1 GLshort, arg2 GLshort, arg3 GLshort, arg4 GLshort) 
{
	C.glVertexAttrib4sARB(C.GLuint(arg0), C.GLshort(arg1), C.GLshort(arg2), C.GLshort(arg3), C.GLshort(arg4));
}

//extern void glVertexAttrib4svARB (GLuint, const GLshort *)
func VertexAttrib4svARB(arg0 GLuint, arg1 *GLshort) 
{
	C.glVertexAttrib4svARB(C.GLuint(arg0), (*C.GLshort)(arg1));
}

//extern void glVertexAttrib4ubvARB (GLuint, const GLubyte *)
func VertexAttrib4ubvARB(arg0 GLuint, arg1 *GLubyte) 
{
	C.glVertexAttrib4ubvARB(C.GLuint(arg0), (*C.GLubyte)(arg1));
}

//extern void glVertexAttrib4uivARB (GLuint, const GLuint *)
func VertexAttrib4uivARB(arg0 GLuint, arg1 *GLuint) 
{
	C.glVertexAttrib4uivARB(C.GLuint(arg0), (*C.GLuint)(arg1));
}

//extern void glVertexAttrib4usvARB (GLuint, const GLushort *)
func VertexAttrib4usvARB(arg0 GLuint, arg1 *GLushort) 
{
	C.glVertexAttrib4usvARB(C.GLuint(arg0), (*C.GLushort)(arg1));
}

//extern void glVertexAttribPointerARB (GLuint, GLint, GLenum, GLboolean, GLsizei, const GLvoid *)
func VertexAttribPointerARB(arg0 GLuint, arg1 GLint, arg2 GLenum, arg3 GLboolean, arg4 GLsizei, arg5 unsafe.Pointer) 
{
	C.glVertexAttribPointerARB(C.GLuint(arg0), C.GLint(arg1), C.GLenum(arg2), C.GLboolean(arg3), C.GLsizei(arg4), arg5);
}

//extern void glEnableVertexAttribArrayARB (GLuint)
func EnableVertexAttribArrayARB(arg0 GLuint) 
{
	C.glEnableVertexAttribArrayARB(C.GLuint(arg0));
}

//extern void glDisableVertexAttribArrayARB (GLuint)
func DisableVertexAttribArrayARB(arg0 GLuint) 
{
	C.glDisableVertexAttribArrayARB(C.GLuint(arg0));
}

//extern void glProgramStringARB (GLenum, GLenum, GLsizei, const GLvoid *)
func ProgramStringARB(arg0 GLenum, arg1 GLenum, arg2 GLsizei, arg3 unsafe.Pointer) 
{
	C.glProgramStringARB(C.GLenum(arg0), C.GLenum(arg1), C.GLsizei(arg2), arg3);
}

//extern void glBindProgramARB (GLenum, GLuint)
func BindProgramARB(arg0 GLenum, arg1 GLuint) 
{
	C.glBindProgramARB(C.GLenum(arg0), C.GLuint(arg1));
}

//extern void glDeleteProgramsARB (GLsizei, const GLuint *)
func DeleteProgramsARB(arg0 GLsizei, arg1 *GLuint) 
{
	C.glDeleteProgramsARB(C.GLsizei(arg0), (*C.GLuint)(arg1));
}

//extern void glGenProgramsARB (GLsizei, GLuint *)
func GenProgramsARB(arg0 GLsizei, arg1 *GLuint) 
{
	C.glGenProgramsARB(C.GLsizei(arg0), (*C.GLuint)(arg1));
}

//extern void glProgramEnvParameter4dARB (GLenum, GLuint, GLdouble, GLdouble, GLdouble, GLdouble)
func ProgramEnvParameter4dARB(arg0 GLenum, arg1 GLuint, arg2 GLdouble, arg3 GLdouble, arg4 GLdouble, arg5 GLdouble) 
{
	C.glProgramEnvParameter4dARB(C.GLenum(arg0), C.GLuint(arg1), C.GLdouble(arg2), C.GLdouble(arg3), C.GLdouble(arg4), C.GLdouble(arg5));
}

//extern void glProgramEnvParameter4dvARB (GLenum, GLuint, const GLdouble *)
func ProgramEnvParameter4dvARB(arg0 GLenum, arg1 GLuint, arg2 *GLdouble) 
{
	C.glProgramEnvParameter4dvARB(C.GLenum(arg0), C.GLuint(arg1), (*C.GLdouble)(arg2));
}

//extern void glProgramEnvParameter4fARB (GLenum, GLuint, GLfloat, GLfloat, GLfloat, GLfloat)
func ProgramEnvParameter4fARB(arg0 GLenum, arg1 GLuint, arg2 GLfloat, arg3 GLfloat, arg4 GLfloat, arg5 GLfloat) 
{
	C.glProgramEnvParameter4fARB(C.GLenum(arg0), C.GLuint(arg1), C.GLfloat(arg2), C.GLfloat(arg3), C.GLfloat(arg4), C.GLfloat(arg5));
}

//extern void glProgramEnvParameter4fvARB (GLenum, GLuint, const GLfloat *)
func ProgramEnvParameter4fvARB(arg0 GLenum, arg1 GLuint, arg2 *GLfloat) 
{
	C.glProgramEnvParameter4fvARB(C.GLenum(arg0), C.GLuint(arg1), (*C.GLfloat)(arg2));
}

//extern void glProgramLocalParameter4dARB (GLenum, GLuint, GLdouble, GLdouble, GLdouble, GLdouble)
func ProgramLocalParameter4dARB(arg0 GLenum, arg1 GLuint, arg2 GLdouble, arg3 GLdouble, arg4 GLdouble, arg5 GLdouble) 
{
	C.glProgramLocalParameter4dARB(C.GLenum(arg0), C.GLuint(arg1), C.GLdouble(arg2), C.GLdouble(arg3), C.GLdouble(arg4), C.GLdouble(arg5));
}

//extern void glProgramLocalParameter4dvARB (GLenum, GLuint, const GLdouble *)
func ProgramLocalParameter4dvARB(arg0 GLenum, arg1 GLuint, arg2 *GLdouble) 
{
	C.glProgramLocalParameter4dvARB(C.GLenum(arg0), C.GLuint(arg1), (*C.GLdouble)(arg2));
}

//extern void glProgramLocalParameter4fARB (GLenum, GLuint, GLfloat, GLfloat, GLfloat, GLfloat)
func ProgramLocalParameter4fARB(arg0 GLenum, arg1 GLuint, arg2 GLfloat, arg3 GLfloat, arg4 GLfloat, arg5 GLfloat) 
{
	C.glProgramLocalParameter4fARB(C.GLenum(arg0), C.GLuint(arg1), C.GLfloat(arg2), C.GLfloat(arg3), C.GLfloat(arg4), C.GLfloat(arg5));
}

//extern void glProgramLocalParameter4fvARB (GLenum, GLuint, const GLfloat *)
func ProgramLocalParameter4fvARB(arg0 GLenum, arg1 GLuint, arg2 *GLfloat) 
{
	C.glProgramLocalParameter4fvARB(C.GLenum(arg0), C.GLuint(arg1), (*C.GLfloat)(arg2));
}

//extern void glGetProgramEnvParameterdvARB (GLenum, GLuint, GLdouble *)
func GetProgramEnvParameterdvARB(arg0 GLenum, arg1 GLuint, arg2 *GLdouble) 
{
	C.glGetProgramEnvParameterdvARB(C.GLenum(arg0), C.GLuint(arg1), (*C.GLdouble)(arg2));
}

//extern void glGetProgramEnvParameterfvARB (GLenum, GLuint, GLfloat *)
func GetProgramEnvParameterfvARB(arg0 GLenum, arg1 GLuint, arg2 *GLfloat) 
{
	C.glGetProgramEnvParameterfvARB(C.GLenum(arg0), C.GLuint(arg1), (*C.GLfloat)(arg2));
}

//extern void glGetProgramLocalParameterdvARB (GLenum, GLuint, GLdouble *)
func GetProgramLocalParameterdvARB(arg0 GLenum, arg1 GLuint, arg2 *GLdouble) 
{
	C.glGetProgramLocalParameterdvARB(C.GLenum(arg0), C.GLuint(arg1), (*C.GLdouble)(arg2));
}

//extern void glGetProgramLocalParameterfvARB (GLenum, GLuint, GLfloat *)
func GetProgramLocalParameterfvARB(arg0 GLenum, arg1 GLuint, arg2 *GLfloat) 
{
	C.glGetProgramLocalParameterfvARB(C.GLenum(arg0), C.GLuint(arg1), (*C.GLfloat)(arg2));
}

//extern void glGetProgramivARB (GLenum, GLenum, GLint *)
func GetProgramivARB(arg0 GLenum, arg1 GLenum, arg2 *GLint) 
{
	C.glGetProgramivARB(C.GLenum(arg0), C.GLenum(arg1), (*C.GLint)(arg2));
}

//extern void glGetProgramStringARB (GLenum, GLenum, GLvoid *)
func GetProgramStringARB(arg0 GLenum, arg1 GLenum, arg2 unsafe.Pointer) 
{
	C.glGetProgramStringARB(C.GLenum(arg0), C.GLenum(arg1), arg2);
}

//extern void glGetVertexAttribdvARB (GLuint, GLenum, GLdouble *)
func GetVertexAttribdvARB(arg0 GLuint, arg1 GLenum, arg2 *GLdouble) 
{
	C.glGetVertexAttribdvARB(C.GLuint(arg0), C.GLenum(arg1), (*C.GLdouble)(arg2));
}

//extern void glGetVertexAttribfvARB (GLuint, GLenum, GLfloat *)
func GetVertexAttribfvARB(arg0 GLuint, arg1 GLenum, arg2 *GLfloat) 
{
	C.glGetVertexAttribfvARB(C.GLuint(arg0), C.GLenum(arg1), (*C.GLfloat)(arg2));
}

//extern void glGetVertexAttribivARB (GLuint, GLenum, GLint *)
func GetVertexAttribivARB(arg0 GLuint, arg1 GLenum, arg2 *GLint) 
{
	C.glGetVertexAttribivARB(C.GLuint(arg0), C.GLenum(arg1), (*C.GLint)(arg2));
}

//extern void glGetVertexAttribPointervARB (GLuint, GLenum, GLvoid* *)
func GetVertexAttribPointervARB(arg0 GLuint, arg1 GLenum, arg2 *unsafe.Pointer) 
{
	C.glGetVertexAttribPointervARB(C.GLuint(arg0), C.GLenum(arg1), arg2);
}

//extern GLboolean glIsProgramARB (GLuint)
func IsProgramARB(arg0 GLuint) GLboolean
{
	return GLboolean(C.glIsProgramARB(C.GLuint(arg0)));
}

//extern void glBindBufferARB (GLenum, GLuint)
func BindBufferARB(arg0 GLenum, arg1 GLuint) 
{
	C.glBindBufferARB(C.GLenum(arg0), C.GLuint(arg1));
}

//extern void glDeleteBuffersARB (GLsizei, const GLuint *)
func DeleteBuffersARB(arg0 GLsizei, arg1 *GLuint) 
{
	C.glDeleteBuffersARB(C.GLsizei(arg0), (*C.GLuint)(arg1));
}

//extern void glGenBuffersARB (GLsizei, GLuint *)
func GenBuffersARB(arg0 GLsizei, arg1 *GLuint) 
{
	C.glGenBuffersARB(C.GLsizei(arg0), (*C.GLuint)(arg1));
}

//extern GLboolean glIsBufferARB (GLuint)
func IsBufferARB(arg0 GLuint) GLboolean
{
	return GLboolean(C.glIsBufferARB(C.GLuint(arg0)));
}

//extern void glBufferDataARB (GLenum, GLsizeiptrARB, const GLvoid *, GLenum)
func BufferDataARB(arg0 GLenum, arg1 GLsizeiptrARB, arg2 unsafe.Pointer, arg3 GLenum) 
{
	C.glBufferDataARB(C.GLenum(arg0), C.GLsizeiptrARB(arg1), arg2, C.GLenum(arg3));
}

//extern void glBufferSubDataARB (GLenum, GLintptrARB, GLsizeiptrARB, const GLvoid *)
func BufferSubDataARB(arg0 GLenum, arg1 GLintptrARB, arg2 GLsizeiptrARB, arg3 unsafe.Pointer) 
{
	C.glBufferSubDataARB(C.GLenum(arg0), C.GLintptrARB(arg1), C.GLsizeiptrARB(arg2), arg3);
}

//extern void glGetBufferSubDataARB (GLenum, GLintptrARB, GLsizeiptrARB, GLvoid *)
func GetBufferSubDataARB(arg0 GLenum, arg1 GLintptrARB, arg2 GLsizeiptrARB, arg3 unsafe.Pointer) 
{
	C.glGetBufferSubDataARB(C.GLenum(arg0), C.GLintptrARB(arg1), C.GLsizeiptrARB(arg2), arg3);
}

//extern GLvoid* glMapBufferARB (GLenum, GLenum)
func MapBufferARB(arg0 GLenum, arg1 GLenum) unsafe.Pointer
{
	return unsafe.Pointer(C.glMapBufferARB(C.GLenum(arg0), C.GLenum(arg1)));
}

//extern GLboolean glUnmapBufferARB (GLenum)
func UnmapBufferARB(arg0 GLenum) GLboolean
{
	return GLboolean(C.glUnmapBufferARB(C.GLenum(arg0)));
}

//extern void glGetBufferParameterivARB (GLenum, GLenum, GLint *)
func GetBufferParameterivARB(arg0 GLenum, arg1 GLenum, arg2 *GLint) 
{
	C.glGetBufferParameterivARB(C.GLenum(arg0), C.GLenum(arg1), (*C.GLint)(arg2));
}

//extern void glGetBufferPointervARB (GLenum, GLenum, GLvoid* *)
func GetBufferPointervARB(arg0 GLenum, arg1 GLenum, arg2 *unsafe.Pointer) 
{
	C.glGetBufferPointervARB(C.GLenum(arg0), C.GLenum(arg1), arg2);
}

//extern void glGenQueriesARB (GLsizei, GLuint *)
func GenQueriesARB(arg0 GLsizei, arg1 *GLuint) 
{
	C.glGenQueriesARB(C.GLsizei(arg0), (*C.GLuint)(arg1));
}

//extern void glDeleteQueriesARB (GLsizei, const GLuint *)
func DeleteQueriesARB(arg0 GLsizei, arg1 *GLuint) 
{
	C.glDeleteQueriesARB(C.GLsizei(arg0), (*C.GLuint)(arg1));
}

//extern GLboolean glIsQueryARB (GLuint)
func IsQueryARB(arg0 GLuint) GLboolean
{
	return GLboolean(C.glIsQueryARB(C.GLuint(arg0)));
}

//extern void glBeginQueryARB (GLenum, GLuint)
func BeginQueryARB(arg0 GLenum, arg1 GLuint) 
{
	C.glBeginQueryARB(C.GLenum(arg0), C.GLuint(arg1));
}

//extern void glEndQueryARB (GLenum)
func EndQueryARB(arg0 GLenum) 
{
	C.glEndQueryARB(C.GLenum(arg0));
}

//extern void glGetQueryivARB (GLenum, GLenum, GLint *)
func GetQueryivARB(arg0 GLenum, arg1 GLenum, arg2 *GLint) 
{
	C.glGetQueryivARB(C.GLenum(arg0), C.GLenum(arg1), (*C.GLint)(arg2));
}

//extern void glGetQueryObjectivARB (GLuint, GLenum, GLint *)
func GetQueryObjectivARB(arg0 GLuint, arg1 GLenum, arg2 *GLint) 
{
	C.glGetQueryObjectivARB(C.GLuint(arg0), C.GLenum(arg1), (*C.GLint)(arg2));
}

//extern void glGetQueryObjectuivARB (GLuint, GLenum, GLuint *)
func GetQueryObjectuivARB(arg0 GLuint, arg1 GLenum, arg2 *GLuint) 
{
	C.glGetQueryObjectuivARB(C.GLuint(arg0), C.GLenum(arg1), (*C.GLuint)(arg2));
}

//extern void glDeleteObjectARB (GLhandleARB)
func DeleteObjectARB(arg0 GLhandleARB) 
{
	C.glDeleteObjectARB(C.GLhandleARB(arg0));
}

//extern GLhandleARB glGetHandleARB (GLenum)
func GetHandleARB(arg0 GLenum) GLhandleARB
{
	return GLhandleARB(C.glGetHandleARB(C.GLenum(arg0)));
}

//extern void glDetachObjectARB (GLhandleARB, GLhandleARB)
func DetachObjectARB(arg0 GLhandleARB, arg1 GLhandleARB) 
{
	C.glDetachObjectARB(C.GLhandleARB(arg0), C.GLhandleARB(arg1));
}

//extern GLhandleARB glCreateShaderObjectARB (GLenum)
func CreateShaderObjectARB(arg0 GLenum) GLhandleARB
{
	return GLhandleARB(C.glCreateShaderObjectARB(C.GLenum(arg0)));
}

//extern void glShaderSourceARB (GLhandleARB, GLsizei, const GLcharARB* *, const GLint *)
func ShaderSourceARB(arg0 GLhandleARB, arg1 GLsizei, arg2 **GLcharARB, arg3 *GLint) 
{
	C.glShaderSourceARB(C.GLhandleARB(arg0), C.GLsizei(arg1), (**C.GLcharARB)(arg2), (*C.GLint)(arg3));
}

//extern void glCompileShaderARB (GLhandleARB)
func CompileShaderARB(arg0 GLhandleARB) 
{
	C.glCompileShaderARB(C.GLhandleARB(arg0));
}

//extern GLhandleARB glCreateProgramObjectARB (void)
func CreateProgramObjectARB() GLhandleARB
{
	return GLhandleARB(C.glCreateProgramObjectARB());
}

//extern void glAttachObjectARB (GLhandleARB, GLhandleARB)
func AttachObjectARB(arg0 GLhandleARB, arg1 GLhandleARB) 
{
	C.glAttachObjectARB(C.GLhandleARB(arg0), C.GLhandleARB(arg1));
}

//extern void glLinkProgramARB (GLhandleARB)
func LinkProgramARB(arg0 GLhandleARB) 
{
	C.glLinkProgramARB(C.GLhandleARB(arg0));
}

//extern void glUseProgramObjectARB (GLhandleARB)
func UseProgramObjectARB(arg0 GLhandleARB) 
{
	C.glUseProgramObjectARB(C.GLhandleARB(arg0));
}

//extern void glValidateProgramARB (GLhandleARB)
func ValidateProgramARB(arg0 GLhandleARB) 
{
	C.glValidateProgramARB(C.GLhandleARB(arg0));
}

//extern void glUniform1fARB (GLint, GLfloat)
func Uniform1fARB(arg0 GLint, arg1 GLfloat) 
{
	C.glUniform1fARB(C.GLint(arg0), C.GLfloat(arg1));
}

//extern void glUniform2fARB (GLint, GLfloat, GLfloat)
func Uniform2fARB(arg0 GLint, arg1 GLfloat, arg2 GLfloat) 
{
	C.glUniform2fARB(C.GLint(arg0), C.GLfloat(arg1), C.GLfloat(arg2));
}

//extern void glUniform3fARB (GLint, GLfloat, GLfloat, GLfloat)
func Uniform3fARB(arg0 GLint, arg1 GLfloat, arg2 GLfloat, arg3 GLfloat) 
{
	C.glUniform3fARB(C.GLint(arg0), C.GLfloat(arg1), C.GLfloat(arg2), C.GLfloat(arg3));
}

//extern void glUniform4fARB (GLint, GLfloat, GLfloat, GLfloat, GLfloat)
func Uniform4fARB(arg0 GLint, arg1 GLfloat, arg2 GLfloat, arg3 GLfloat, arg4 GLfloat) 
{
	C.glUniform4fARB(C.GLint(arg0), C.GLfloat(arg1), C.GLfloat(arg2), C.GLfloat(arg3), C.GLfloat(arg4));
}

//extern void glUniform1iARB (GLint, GLint)
func Uniform1iARB(arg0 GLint, arg1 GLint) 
{
	C.glUniform1iARB(C.GLint(arg0), C.GLint(arg1));
}

//extern void glUniform2iARB (GLint, GLint, GLint)
func Uniform2iARB(arg0 GLint, arg1 GLint, arg2 GLint) 
{
	C.glUniform2iARB(C.GLint(arg0), C.GLint(arg1), C.GLint(arg2));
}

//extern void glUniform3iARB (GLint, GLint, GLint, GLint)
func Uniform3iARB(arg0 GLint, arg1 GLint, arg2 GLint, arg3 GLint) 
{
	C.glUniform3iARB(C.GLint(arg0), C.GLint(arg1), C.GLint(arg2), C.GLint(arg3));
}

//extern void glUniform4iARB (GLint, GLint, GLint, GLint, GLint)
func Uniform4iARB(arg0 GLint, arg1 GLint, arg2 GLint, arg3 GLint, arg4 GLint) 
{
	C.glUniform4iARB(C.GLint(arg0), C.GLint(arg1), C.GLint(arg2), C.GLint(arg3), C.GLint(arg4));
}

//extern void glUniform1fvARB (GLint, GLsizei, const GLfloat *)
func Uniform1fvARB(arg0 GLint, arg1 GLsizei, arg2 *GLfloat) 
{
	C.glUniform1fvARB(C.GLint(arg0), C.GLsizei(arg1), (*C.GLfloat)(arg2));
}

//extern void glUniform2fvARB (GLint, GLsizei, const GLfloat *)
func Uniform2fvARB(arg0 GLint, arg1 GLsizei, arg2 *GLfloat) 
{
	C.glUniform2fvARB(C.GLint(arg0), C.GLsizei(arg1), (*C.GLfloat)(arg2));
}

//extern void glUniform3fvARB (GLint, GLsizei, const GLfloat *)
func Uniform3fvARB(arg0 GLint, arg1 GLsizei, arg2 *GLfloat) 
{
	C.glUniform3fvARB(C.GLint(arg0), C.GLsizei(arg1), (*C.GLfloat)(arg2));
}

//extern void glUniform4fvARB (GLint, GLsizei, const GLfloat *)
func Uniform4fvARB(arg0 GLint, arg1 GLsizei, arg2 *GLfloat) 
{
	C.glUniform4fvARB(C.GLint(arg0), C.GLsizei(arg1), (*C.GLfloat)(arg2));
}

//extern void glUniform1ivARB (GLint, GLsizei, const GLint *)
func Uniform1ivARB(arg0 GLint, arg1 GLsizei, arg2 *GLint) 
{
	C.glUniform1ivARB(C.GLint(arg0), C.GLsizei(arg1), (*C.GLint)(arg2));
}

//extern void glUniform2ivARB (GLint, GLsizei, const GLint *)
func Uniform2ivARB(arg0 GLint, arg1 GLsizei, arg2 *GLint) 
{
	C.glUniform2ivARB(C.GLint(arg0), C.GLsizei(arg1), (*C.GLint)(arg2));
}

//extern void glUniform3ivARB (GLint, GLsizei, const GLint *)
func Uniform3ivARB(arg0 GLint, arg1 GLsizei, arg2 *GLint) 
{
	C.glUniform3ivARB(C.GLint(arg0), C.GLsizei(arg1), (*C.GLint)(arg2));
}

//extern void glUniform4ivARB (GLint, GLsizei, const GLint *)
func Uniform4ivARB(arg0 GLint, arg1 GLsizei, arg2 *GLint) 
{
	C.glUniform4ivARB(C.GLint(arg0), C.GLsizei(arg1), (*C.GLint)(arg2));
}

//extern void glUniformMatrix2fvARB (GLint, GLsizei, GLboolean, const GLfloat *)
func UniformMatrix2fvARB(arg0 GLint, arg1 GLsizei, arg2 GLboolean, arg3 *GLfloat) 
{
	C.glUniformMatrix2fvARB(C.GLint(arg0), C.GLsizei(arg1), C.GLboolean(arg2), (*C.GLfloat)(arg3));
}

//extern void glUniformMatrix3fvARB (GLint, GLsizei, GLboolean, const GLfloat *)
func UniformMatrix3fvARB(arg0 GLint, arg1 GLsizei, arg2 GLboolean, arg3 *GLfloat) 
{
	C.glUniformMatrix3fvARB(C.GLint(arg0), C.GLsizei(arg1), C.GLboolean(arg2), (*C.GLfloat)(arg3));
}

//extern void glUniformMatrix4fvARB (GLint, GLsizei, GLboolean, const GLfloat *)
func UniformMatrix4fvARB(arg0 GLint, arg1 GLsizei, arg2 GLboolean, arg3 *GLfloat) 
{
	C.glUniformMatrix4fvARB(C.GLint(arg0), C.GLsizei(arg1), C.GLboolean(arg2), (*C.GLfloat)(arg3));
}

//extern void glGetObjectParameterfvARB (GLhandleARB, GLenum, GLfloat *)
func GetObjectParameterfvARB(arg0 GLhandleARB, arg1 GLenum, arg2 *GLfloat) 
{
	C.glGetObjectParameterfvARB(C.GLhandleARB(arg0), C.GLenum(arg1), (*C.GLfloat)(arg2));
}

//extern void glGetObjectParameterivARB (GLhandleARB, GLenum, GLint *)
func GetObjectParameterivARB(arg0 GLhandleARB, arg1 GLenum, arg2 *GLint) 
{
	C.glGetObjectParameterivARB(C.GLhandleARB(arg0), C.GLenum(arg1), (*C.GLint)(arg2));
}

//extern void glGetInfoLogARB (GLhandleARB, GLsizei, GLsizei *, GLcharARB *)
func GetInfoLogARB(arg0 GLhandleARB, arg1 GLsizei, arg2 *GLsizei, arg3 *GLcharARB) 
{
	C.glGetInfoLogARB(C.GLhandleARB(arg0), C.GLsizei(arg1), (*C.GLsizei)(arg2), (*C.GLcharARB)(arg3));
}

//extern void glGetAttachedObjectsARB (GLhandleARB, GLsizei, GLsizei *, GLhandleARB *)
func GetAttachedObjectsARB(arg0 GLhandleARB, arg1 GLsizei, arg2 *GLsizei, arg3 *GLhandleARB) 
{
	C.glGetAttachedObjectsARB(C.GLhandleARB(arg0), C.GLsizei(arg1), (*C.GLsizei)(arg2), (*C.GLhandleARB)(arg3));
}

//extern GLint glGetUniformLocationARB (GLhandleARB, const GLcharARB *)
func GetUniformLocationARB(arg0 GLhandleARB, arg1 *GLcharARB) GLint
{
	return GLint(C.glGetUniformLocationARB(C.GLhandleARB(arg0), (*C.GLcharARB)(arg1)));
}

//extern void glGetActiveUniformARB (GLhandleARB, GLuint, GLsizei, GLsizei *, GLint *, GLenum *, GLcharARB *)
func GetActiveUniformARB(arg0 GLhandleARB, arg1 GLuint, arg2 GLsizei, arg3 *GLsizei, arg4 *GLint, arg5 *GLenum, arg6 *GLcharARB) 
{
	C.glGetActiveUniformARB(C.GLhandleARB(arg0), C.GLuint(arg1), C.GLsizei(arg2), (*C.GLsizei)(arg3), (*C.GLint)(arg4), (*C.GLenum)(arg5), (*C.GLcharARB)(arg6));
}

//extern void glGetUniformfvARB (GLhandleARB, GLint, GLfloat *)
func GetUniformfvARB(arg0 GLhandleARB, arg1 GLint, arg2 *GLfloat) 
{
	C.glGetUniformfvARB(C.GLhandleARB(arg0), C.GLint(arg1), (*C.GLfloat)(arg2));
}

//extern void glGetUniformivARB (GLhandleARB, GLint, GLint *)
func GetUniformivARB(arg0 GLhandleARB, arg1 GLint, arg2 *GLint) 
{
	C.glGetUniformivARB(C.GLhandleARB(arg0), C.GLint(arg1), (*C.GLint)(arg2));
}

//extern void glGetShaderSourceARB (GLhandleARB, GLsizei, GLsizei *, GLcharARB *)
func GetShaderSourceARB(arg0 GLhandleARB, arg1 GLsizei, arg2 *GLsizei, arg3 *GLcharARB) 
{
	C.glGetShaderSourceARB(C.GLhandleARB(arg0), C.GLsizei(arg1), (*C.GLsizei)(arg2), (*C.GLcharARB)(arg3));
}

//extern void glBindAttribLocationARB (GLhandleARB, GLuint, const GLcharARB *)
func BindAttribLocationARB(arg0 GLhandleARB, arg1 GLuint, arg2 *GLcharARB) 
{
	C.glBindAttribLocationARB(C.GLhandleARB(arg0), C.GLuint(arg1), (*C.GLcharARB)(arg2));
}

//extern void glGetActiveAttribARB (GLhandleARB, GLuint, GLsizei, GLsizei *, GLint *, GLenum *, GLcharARB *)
func GetActiveAttribARB(arg0 GLhandleARB, arg1 GLuint, arg2 GLsizei, arg3 *GLsizei, arg4 *GLint, arg5 *GLenum, arg6 *GLcharARB) 
{
	C.glGetActiveAttribARB(C.GLhandleARB(arg0), C.GLuint(arg1), C.GLsizei(arg2), (*C.GLsizei)(arg3), (*C.GLint)(arg4), (*C.GLenum)(arg5), (*C.GLcharARB)(arg6));
}

//extern GLint glGetAttribLocationARB (GLhandleARB, const GLcharARB *)
func GetAttribLocationARB(arg0 GLhandleARB, arg1 *GLcharARB) GLint
{
	return GLint(C.glGetAttribLocationARB(C.GLhandleARB(arg0), (*C.GLcharARB)(arg1)));
}

//extern void glDrawBuffersARB (GLsizei, const GLenum *)
func DrawBuffersARB(arg0 GLsizei, arg1 *GLenum) 
{
	C.glDrawBuffersARB(C.GLsizei(arg0), (*C.GLenum)(arg1));
}

//extern void glClampColorARB (GLenum, GLenum)
func ClampColorARB(arg0 GLenum, arg1 GLenum) 
{
	C.glClampColorARB(C.GLenum(arg0), C.GLenum(arg1));
}

//extern void glDrawArraysInstancedARB (GLenum, GLint, GLsizei, GLsizei)
func DrawArraysInstancedARB(arg0 GLenum, arg1 GLint, arg2 GLsizei, arg3 GLsizei) 
{
	C.glDrawArraysInstancedARB(C.GLenum(arg0), C.GLint(arg1), C.GLsizei(arg2), C.GLsizei(arg3));
}

//extern void glDrawElementsInstancedARB (GLenum, GLsizei, GLenum, const GLvoid *, GLsizei)
func DrawElementsInstancedARB(arg0 GLenum, arg1 GLsizei, arg2 GLenum, arg3 unsafe.Pointer, arg4 GLsizei) 
{
	C.glDrawElementsInstancedARB(C.GLenum(arg0), C.GLsizei(arg1), C.GLenum(arg2), arg3, C.GLsizei(arg4));
}

//extern GLboolean glIsRenderbuffer (GLuint)
func IsRenderbuffer(arg0 GLuint) GLboolean
{
	return GLboolean(C.glIsRenderbuffer(C.GLuint(arg0)));
}

//extern void glBindRenderbuffer (GLenum, GLuint)
func BindRenderbuffer(arg0 GLenum, arg1 GLuint) 
{
	C.glBindRenderbuffer(C.GLenum(arg0), C.GLuint(arg1));
}

//extern void glDeleteRenderbuffers (GLsizei, const GLuint *)
func DeleteRenderbuffers(arg0 GLsizei, arg1 *GLuint) 
{
	C.glDeleteRenderbuffers(C.GLsizei(arg0), (*C.GLuint)(arg1));
}

//extern void glGenRenderbuffers (GLsizei, GLuint *)
func GenRenderbuffers(arg0 GLsizei, arg1 *GLuint) 
{
	C.glGenRenderbuffers(C.GLsizei(arg0), (*C.GLuint)(arg1));
}

//extern void glRenderbufferStorage (GLenum, GLenum, GLsizei, GLsizei)
func RenderbufferStorage(arg0 GLenum, arg1 GLenum, arg2 GLsizei, arg3 GLsizei) 
{
	C.glRenderbufferStorage(C.GLenum(arg0), C.GLenum(arg1), C.GLsizei(arg2), C.GLsizei(arg3));
}

//extern void glGetRenderbufferParameteriv (GLenum, GLenum, GLint *)
func GetRenderbufferParameteriv(arg0 GLenum, arg1 GLenum, arg2 *GLint) 
{
	C.glGetRenderbufferParameteriv(C.GLenum(arg0), C.GLenum(arg1), (*C.GLint)(arg2));
}

//extern GLboolean glIsFramebuffer (GLuint)
func IsFramebuffer(arg0 GLuint) GLboolean
{
	return GLboolean(C.glIsFramebuffer(C.GLuint(arg0)));
}

//extern void glBindFramebuffer (GLenum, GLuint)
func BindFramebuffer(arg0 GLenum, arg1 GLuint) 
{
	C.glBindFramebuffer(C.GLenum(arg0), C.GLuint(arg1));
}

//extern void glDeleteFramebuffers (GLsizei, const GLuint *)
func DeleteFramebuffers(arg0 GLsizei, arg1 *GLuint) 
{
	C.glDeleteFramebuffers(C.GLsizei(arg0), (*C.GLuint)(arg1));
}

//extern void glGenFramebuffers (GLsizei, GLuint *)
func GenFramebuffers(arg0 GLsizei, arg1 *GLuint) 
{
	C.glGenFramebuffers(C.GLsizei(arg0), (*C.GLuint)(arg1));
}

//extern GLenum glCheckFramebufferStatus (GLenum)
func CheckFramebufferStatus(arg0 GLenum) GLenum
{
	return GLenum(C.glCheckFramebufferStatus(C.GLenum(arg0)));
}

//extern void glFramebufferTexture1D (GLenum, GLenum, GLenum, GLuint, GLint)
func FramebufferTexture1D(arg0 GLenum, arg1 GLenum, arg2 GLenum, arg3 GLuint, arg4 GLint) 
{
	C.glFramebufferTexture1D(C.GLenum(arg0), C.GLenum(arg1), C.GLenum(arg2), C.GLuint(arg3), C.GLint(arg4));
}

//extern void glFramebufferTexture2D (GLenum, GLenum, GLenum, GLuint, GLint)
func FramebufferTexture2D(arg0 GLenum, arg1 GLenum, arg2 GLenum, arg3 GLuint, arg4 GLint) 
{
	C.glFramebufferTexture2D(C.GLenum(arg0), C.GLenum(arg1), C.GLenum(arg2), C.GLuint(arg3), C.GLint(arg4));
}

//extern void glFramebufferTexture3D (GLenum, GLenum, GLenum, GLuint, GLint, GLint)
func FramebufferTexture3D(arg0 GLenum, arg1 GLenum, arg2 GLenum, arg3 GLuint, arg4 GLint, arg5 GLint) 
{
	C.glFramebufferTexture3D(C.GLenum(arg0), C.GLenum(arg1), C.GLenum(arg2), C.GLuint(arg3), C.GLint(arg4), C.GLint(arg5));
}

//extern void glFramebufferRenderbuffer (GLenum, GLenum, GLenum, GLuint)
func FramebufferRenderbuffer(arg0 GLenum, arg1 GLenum, arg2 GLenum, arg3 GLuint) 
{
	C.glFramebufferRenderbuffer(C.GLenum(arg0), C.GLenum(arg1), C.GLenum(arg2), C.GLuint(arg3));
}

//extern void glGetFramebufferAttachmentParameteriv (GLenum, GLenum, GLenum, GLint *)
func GetFramebufferAttachmentParameteriv(arg0 GLenum, arg1 GLenum, arg2 GLenum, arg3 *GLint) 
{
	C.glGetFramebufferAttachmentParameteriv(C.GLenum(arg0), C.GLenum(arg1), C.GLenum(arg2), (*C.GLint)(arg3));
}

//extern void glGenerateMipmap (GLenum)
func GenerateMipmap(arg0 GLenum) 
{
	C.glGenerateMipmap(C.GLenum(arg0));
}

//extern void glBlitFramebuffer (GLint, GLint, GLint, GLint, GLint, GLint, GLint, GLint, GLbitfield, GLenum)
func BlitFramebuffer(arg0 GLint, arg1 GLint, arg2 GLint, arg3 GLint, arg4 GLint, arg5 GLint, arg6 GLint, arg7 GLint, arg8 GLbitfield, arg9 GLenum) 
{
	C.glBlitFramebuffer(C.GLint(arg0), C.GLint(arg1), C.GLint(arg2), C.GLint(arg3), C.GLint(arg4), C.GLint(arg5), C.GLint(arg6), C.GLint(arg7), C.GLbitfield(arg8), C.GLenum(arg9));
}

//extern void glRenderbufferStorageMultisample (GLenum, GLsizei, GLenum, GLsizei, GLsizei)
func RenderbufferStorageMultisample(arg0 GLenum, arg1 GLsizei, arg2 GLenum, arg3 GLsizei, arg4 GLsizei) 
{
	C.glRenderbufferStorageMultisample(C.GLenum(arg0), C.GLsizei(arg1), C.GLenum(arg2), C.GLsizei(arg3), C.GLsizei(arg4));
}

//extern void glFramebufferTextureLayer (GLenum, GLenum, GLuint, GLint, GLint)
func FramebufferTextureLayer(arg0 GLenum, arg1 GLenum, arg2 GLuint, arg3 GLint, arg4 GLint) 
{
	C.glFramebufferTextureLayer(C.GLenum(arg0), C.GLenum(arg1), C.GLuint(arg2), C.GLint(arg3), C.GLint(arg4));
}

//extern void glProgramParameteriARB (GLuint, GLenum, GLint)
func ProgramParameteriARB(arg0 GLuint, arg1 GLenum, arg2 GLint) 
{
	C.glProgramParameteriARB(C.GLuint(arg0), C.GLenum(arg1), C.GLint(arg2));
}

//extern void glFramebufferTextureARB (GLenum, GLenum, GLuint, GLint)
func FramebufferTextureARB(arg0 GLenum, arg1 GLenum, arg2 GLuint, arg3 GLint) 
{
	C.glFramebufferTextureARB(C.GLenum(arg0), C.GLenum(arg1), C.GLuint(arg2), C.GLint(arg3));
}

//extern void glFramebufferTextureLayerARB (GLenum, GLenum, GLuint, GLint, GLint)
func FramebufferTextureLayerARB(arg0 GLenum, arg1 GLenum, arg2 GLuint, arg3 GLint, arg4 GLint) 
{
	C.glFramebufferTextureLayerARB(C.GLenum(arg0), C.GLenum(arg1), C.GLuint(arg2), C.GLint(arg3), C.GLint(arg4));
}

//extern void glFramebufferTextureFaceARB (GLenum, GLenum, GLuint, GLint, GLenum)
func FramebufferTextureFaceARB(arg0 GLenum, arg1 GLenum, arg2 GLuint, arg3 GLint, arg4 GLenum) 
{
	C.glFramebufferTextureFaceARB(C.GLenum(arg0), C.GLenum(arg1), C.GLuint(arg2), C.GLint(arg3), C.GLenum(arg4));
}

//extern void glVertexAttribDivisorARB (GLuint, GLuint)
func VertexAttribDivisorARB(arg0 GLuint, arg1 GLuint) 
{
	C.glVertexAttribDivisorARB(C.GLuint(arg0), C.GLuint(arg1));
}

//extern void glMapBufferRange (GLenum, GLintptr, GLsizeiptr, GLbitfield)
func MapBufferRange(arg0 GLenum, arg1 GLintptr, arg2 GLsizeiptr, arg3 GLbitfield) 
{
	C.glMapBufferRange(C.GLenum(arg0), C.GLintptr(arg1), C.GLsizeiptr(arg2), C.GLbitfield(arg3));
}

//extern void glFlushMappedBufferRange (GLenum, GLintptr, GLsizeiptr)
func FlushMappedBufferRange(arg0 GLenum, arg1 GLintptr, arg2 GLsizeiptr) 
{
	C.glFlushMappedBufferRange(C.GLenum(arg0), C.GLintptr(arg1), C.GLsizeiptr(arg2));
}

//extern void glTexBufferARB (GLenum, GLenum, GLuint)
func TexBufferARB(arg0 GLenum, arg1 GLenum, arg2 GLuint) 
{
	C.glTexBufferARB(C.GLenum(arg0), C.GLenum(arg1), C.GLuint(arg2));
}

//extern void glBindVertexArray (GLuint)
func BindVertexArray(arg0 GLuint) 
{
	C.glBindVertexArray(C.GLuint(arg0));
}

//extern void glDeleteVertexArrays (GLsizei, const GLuint *)
func DeleteVertexArrays(arg0 GLsizei, arg1 *GLuint) 
{
	C.glDeleteVertexArrays(C.GLsizei(arg0), (*C.GLuint)(arg1));
}

//extern void glGenVertexArrays (GLsizei, GLuint *)
func GenVertexArrays(arg0 GLsizei, arg1 *GLuint) 
{
	C.glGenVertexArrays(C.GLsizei(arg0), (*C.GLuint)(arg1));
}

//extern GLboolean glIsVertexArray (GLuint)
func IsVertexArray(arg0 GLuint) GLboolean
{
	return GLboolean(C.glIsVertexArray(C.GLuint(arg0)));
}

//extern void glBlendColorEXT (GLclampf, GLclampf, GLclampf, GLclampf)
func BlendColorEXT(arg0 GLclampf, arg1 GLclampf, arg2 GLclampf, arg3 GLclampf) 
{
	C.glBlendColorEXT(C.GLclampf(arg0), C.GLclampf(arg1), C.GLclampf(arg2), C.GLclampf(arg3));
}

//extern void glPolygonOffsetEXT (GLfloat, GLfloat)
func PolygonOffsetEXT(arg0 GLfloat, arg1 GLfloat) 
{
	C.glPolygonOffsetEXT(C.GLfloat(arg0), C.GLfloat(arg1));
}

//extern void glTexImage3DEXT (GLenum, GLint, GLenum, GLsizei, GLsizei, GLsizei, GLint, GLenum, GLenum, const GLvoid *)
func TexImage3DEXT(arg0 GLenum, arg1 GLint, arg2 GLenum, arg3 GLsizei, arg4 GLsizei, arg5 GLsizei, arg6 GLint, arg7 GLenum, arg8 GLenum, arg9 unsafe.Pointer) 
{
	C.glTexImage3DEXT(C.GLenum(arg0), C.GLint(arg1), C.GLenum(arg2), C.GLsizei(arg3), C.GLsizei(arg4), C.GLsizei(arg5), C.GLint(arg6), C.GLenum(arg7), C.GLenum(arg8), arg9);
}

//extern void glTexSubImage3DEXT (GLenum, GLint, GLint, GLint, GLint, GLsizei, GLsizei, GLsizei, GLenum, GLenum, const GLvoid *)
func TexSubImage3DEXT(arg0 GLenum, arg1 GLint, arg2 GLint, arg3 GLint, arg4 GLint, arg5 GLsizei, arg6 GLsizei, arg7 GLsizei, arg8 GLenum, arg9 GLenum, arg10 unsafe.Pointer) 
{
	C.glTexSubImage3DEXT(C.GLenum(arg0), C.GLint(arg1), C.GLint(arg2), C.GLint(arg3), C.GLint(arg4), C.GLsizei(arg5), C.GLsizei(arg6), C.GLsizei(arg7), C.GLenum(arg8), C.GLenum(arg9), arg10);
}

//extern void glGetTexFilterFuncSGIS (GLenum, GLenum, GLfloat *)
func GetTexFilterFuncSGIS(arg0 GLenum, arg1 GLenum, arg2 *GLfloat) 
{
	C.glGetTexFilterFuncSGIS(C.GLenum(arg0), C.GLenum(arg1), (*C.GLfloat)(arg2));
}

//extern void glTexFilterFuncSGIS (GLenum, GLenum, GLsizei, const GLfloat *)
func TexFilterFuncSGIS(arg0 GLenum, arg1 GLenum, arg2 GLsizei, arg3 *GLfloat) 
{
	C.glTexFilterFuncSGIS(C.GLenum(arg0), C.GLenum(arg1), C.GLsizei(arg2), (*C.GLfloat)(arg3));
}

//extern void glTexSubImage1DEXT (GLenum, GLint, GLint, GLsizei, GLenum, GLenum, const GLvoid *)
func TexSubImage1DEXT(arg0 GLenum, arg1 GLint, arg2 GLint, arg3 GLsizei, arg4 GLenum, arg5 GLenum, arg6 unsafe.Pointer) 
{
	C.glTexSubImage1DEXT(C.GLenum(arg0), C.GLint(arg1), C.GLint(arg2), C.GLsizei(arg3), C.GLenum(arg4), C.GLenum(arg5), arg6);
}

//extern void glTexSubImage2DEXT (GLenum, GLint, GLint, GLint, GLsizei, GLsizei, GLenum, GLenum, const GLvoid *)
func TexSubImage2DEXT(arg0 GLenum, arg1 GLint, arg2 GLint, arg3 GLint, arg4 GLsizei, arg5 GLsizei, arg6 GLenum, arg7 GLenum, arg8 unsafe.Pointer) 
{
	C.glTexSubImage2DEXT(C.GLenum(arg0), C.GLint(arg1), C.GLint(arg2), C.GLint(arg3), C.GLsizei(arg4), C.GLsizei(arg5), C.GLenum(arg6), C.GLenum(arg7), arg8);
}

//extern void glCopyTexImage1DEXT (GLenum, GLint, GLenum, GLint, GLint, GLsizei, GLint)
func CopyTexImage1DEXT(arg0 GLenum, arg1 GLint, arg2 GLenum, arg3 GLint, arg4 GLint, arg5 GLsizei, arg6 GLint) 
{
	C.glCopyTexImage1DEXT(C.GLenum(arg0), C.GLint(arg1), C.GLenum(arg2), C.GLint(arg3), C.GLint(arg4), C.GLsizei(arg5), C.GLint(arg6));
}

//extern void glCopyTexImage2DEXT (GLenum, GLint, GLenum, GLint, GLint, GLsizei, GLsizei, GLint)
func CopyTexImage2DEXT(arg0 GLenum, arg1 GLint, arg2 GLenum, arg3 GLint, arg4 GLint, arg5 GLsizei, arg6 GLsizei, arg7 GLint) 
{
	C.glCopyTexImage2DEXT(C.GLenum(arg0), C.GLint(arg1), C.GLenum(arg2), C.GLint(arg3), C.GLint(arg4), C.GLsizei(arg5), C.GLsizei(arg6), C.GLint(arg7));
}

//extern void glCopyTexSubImage1DEXT (GLenum, GLint, GLint, GLint, GLint, GLsizei)
func CopyTexSubImage1DEXT(arg0 GLenum, arg1 GLint, arg2 GLint, arg3 GLint, arg4 GLint, arg5 GLsizei) 
{
	C.glCopyTexSubImage1DEXT(C.GLenum(arg0), C.GLint(arg1), C.GLint(arg2), C.GLint(arg3), C.GLint(arg4), C.GLsizei(arg5));
}

//extern void glCopyTexSubImage2DEXT (GLenum, GLint, GLint, GLint, GLint, GLint, GLsizei, GLsizei)
func CopyTexSubImage2DEXT(arg0 GLenum, arg1 GLint, arg2 GLint, arg3 GLint, arg4 GLint, arg5 GLint, arg6 GLsizei, arg7 GLsizei) 
{
	C.glCopyTexSubImage2DEXT(C.GLenum(arg0), C.GLint(arg1), C.GLint(arg2), C.GLint(arg3), C.GLint(arg4), C.GLint(arg5), C.GLsizei(arg6), C.GLsizei(arg7));
}

//extern void glCopyTexSubImage3DEXT (GLenum, GLint, GLint, GLint, GLint, GLint, GLint, GLsizei, GLsizei)
func CopyTexSubImage3DEXT(arg0 GLenum, arg1 GLint, arg2 GLint, arg3 GLint, arg4 GLint, arg5 GLint, arg6 GLint, arg7 GLsizei, arg8 GLsizei) 
{
	C.glCopyTexSubImage3DEXT(C.GLenum(arg0), C.GLint(arg1), C.GLint(arg2), C.GLint(arg3), C.GLint(arg4), C.GLint(arg5), C.GLint(arg6), C.GLsizei(arg7), C.GLsizei(arg8));
}

//extern void glGetHistogramEXT (GLenum, GLboolean, GLenum, GLenum, GLvoid *)
func GetHistogramEXT(arg0 GLenum, arg1 GLboolean, arg2 GLenum, arg3 GLenum, arg4 unsafe.Pointer) 
{
	C.glGetHistogramEXT(C.GLenum(arg0), C.GLboolean(arg1), C.GLenum(arg2), C.GLenum(arg3), arg4);
}

//extern void glGetHistogramParameterfvEXT (GLenum, GLenum, GLfloat *)
func GetHistogramParameterfvEXT(arg0 GLenum, arg1 GLenum, arg2 *GLfloat) 
{
	C.glGetHistogramParameterfvEXT(C.GLenum(arg0), C.GLenum(arg1), (*C.GLfloat)(arg2));
}

//extern void glGetHistogramParameterivEXT (GLenum, GLenum, GLint *)
func GetHistogramParameterivEXT(arg0 GLenum, arg1 GLenum, arg2 *GLint) 
{
	C.glGetHistogramParameterivEXT(C.GLenum(arg0), C.GLenum(arg1), (*C.GLint)(arg2));
}

//extern void glGetMinmaxEXT (GLenum, GLboolean, GLenum, GLenum, GLvoid *)
func GetMinmaxEXT(arg0 GLenum, arg1 GLboolean, arg2 GLenum, arg3 GLenum, arg4 unsafe.Pointer) 
{
	C.glGetMinmaxEXT(C.GLenum(arg0), C.GLboolean(arg1), C.GLenum(arg2), C.GLenum(arg3), arg4);
}

//extern void glGetMinmaxParameterfvEXT (GLenum, GLenum, GLfloat *)
func GetMinmaxParameterfvEXT(arg0 GLenum, arg1 GLenum, arg2 *GLfloat) 
{
	C.glGetMinmaxParameterfvEXT(C.GLenum(arg0), C.GLenum(arg1), (*C.GLfloat)(arg2));
}

//extern void glGetMinmaxParameterivEXT (GLenum, GLenum, GLint *)
func GetMinmaxParameterivEXT(arg0 GLenum, arg1 GLenum, arg2 *GLint) 
{
	C.glGetMinmaxParameterivEXT(C.GLenum(arg0), C.GLenum(arg1), (*C.GLint)(arg2));
}

//extern void glHistogramEXT (GLenum, GLsizei, GLenum, GLboolean)
func HistogramEXT(arg0 GLenum, arg1 GLsizei, arg2 GLenum, arg3 GLboolean) 
{
	C.glHistogramEXT(C.GLenum(arg0), C.GLsizei(arg1), C.GLenum(arg2), C.GLboolean(arg3));
}

//extern void glMinmaxEXT (GLenum, GLenum, GLboolean)
func MinmaxEXT(arg0 GLenum, arg1 GLenum, arg2 GLboolean) 
{
	C.glMinmaxEXT(C.GLenum(arg0), C.GLenum(arg1), C.GLboolean(arg2));
}

//extern void glResetHistogramEXT (GLenum)
func ResetHistogramEXT(arg0 GLenum) 
{
	C.glResetHistogramEXT(C.GLenum(arg0));
}

//extern void glResetMinmaxEXT (GLenum)
func ResetMinmaxEXT(arg0 GLenum) 
{
	C.glResetMinmaxEXT(C.GLenum(arg0));
}

//extern void glConvolutionFilter1DEXT (GLenum, GLenum, GLsizei, GLenum, GLenum, const GLvoid *)
func ConvolutionFilter1DEXT(arg0 GLenum, arg1 GLenum, arg2 GLsizei, arg3 GLenum, arg4 GLenum, arg5 unsafe.Pointer) 
{
	C.glConvolutionFilter1DEXT(C.GLenum(arg0), C.GLenum(arg1), C.GLsizei(arg2), C.GLenum(arg3), C.GLenum(arg4), arg5);
}

//extern void glConvolutionFilter2DEXT (GLenum, GLenum, GLsizei, GLsizei, GLenum, GLenum, const GLvoid *)
func ConvolutionFilter2DEXT(arg0 GLenum, arg1 GLenum, arg2 GLsizei, arg3 GLsizei, arg4 GLenum, arg5 GLenum, arg6 unsafe.Pointer) 
{
	C.glConvolutionFilter2DEXT(C.GLenum(arg0), C.GLenum(arg1), C.GLsizei(arg2), C.GLsizei(arg3), C.GLenum(arg4), C.GLenum(arg5), arg6);
}

//extern void glConvolutionParameterfEXT (GLenum, GLenum, GLfloat)
func ConvolutionParameterfEXT(arg0 GLenum, arg1 GLenum, arg2 GLfloat) 
{
	C.glConvolutionParameterfEXT(C.GLenum(arg0), C.GLenum(arg1), C.GLfloat(arg2));
}

//extern void glConvolutionParameterfvEXT (GLenum, GLenum, const GLfloat *)
func ConvolutionParameterfvEXT(arg0 GLenum, arg1 GLenum, arg2 *GLfloat) 
{
	C.glConvolutionParameterfvEXT(C.GLenum(arg0), C.GLenum(arg1), (*C.GLfloat)(arg2));
}

//extern void glConvolutionParameteriEXT (GLenum, GLenum, GLint)
func ConvolutionParameteriEXT(arg0 GLenum, arg1 GLenum, arg2 GLint) 
{
	C.glConvolutionParameteriEXT(C.GLenum(arg0), C.GLenum(arg1), C.GLint(arg2));
}

//extern void glConvolutionParameterivEXT (GLenum, GLenum, const GLint *)
func ConvolutionParameterivEXT(arg0 GLenum, arg1 GLenum, arg2 *GLint) 
{
	C.glConvolutionParameterivEXT(C.GLenum(arg0), C.GLenum(arg1), (*C.GLint)(arg2));
}

//extern void glCopyConvolutionFilter1DEXT (GLenum, GLenum, GLint, GLint, GLsizei)
func CopyConvolutionFilter1DEXT(arg0 GLenum, arg1 GLenum, arg2 GLint, arg3 GLint, arg4 GLsizei) 
{
	C.glCopyConvolutionFilter1DEXT(C.GLenum(arg0), C.GLenum(arg1), C.GLint(arg2), C.GLint(arg3), C.GLsizei(arg4));
}

//extern void glCopyConvolutionFilter2DEXT (GLenum, GLenum, GLint, GLint, GLsizei, GLsizei)
func CopyConvolutionFilter2DEXT(arg0 GLenum, arg1 GLenum, arg2 GLint, arg3 GLint, arg4 GLsizei, arg5 GLsizei) 
{
	C.glCopyConvolutionFilter2DEXT(C.GLenum(arg0), C.GLenum(arg1), C.GLint(arg2), C.GLint(arg3), C.GLsizei(arg4), C.GLsizei(arg5));
}

//extern void glGetConvolutionFilterEXT (GLenum, GLenum, GLenum, GLvoid *)
func GetConvolutionFilterEXT(arg0 GLenum, arg1 GLenum, arg2 GLenum, arg3 unsafe.Pointer) 
{
	C.glGetConvolutionFilterEXT(C.GLenum(arg0), C.GLenum(arg1), C.GLenum(arg2), arg3);
}

//extern void glGetConvolutionParameterfvEXT (GLenum, GLenum, GLfloat *)
func GetConvolutionParameterfvEXT(arg0 GLenum, arg1 GLenum, arg2 *GLfloat) 
{
	C.glGetConvolutionParameterfvEXT(C.GLenum(arg0), C.GLenum(arg1), (*C.GLfloat)(arg2));
}

//extern void glGetConvolutionParameterivEXT (GLenum, GLenum, GLint *)
func GetConvolutionParameterivEXT(arg0 GLenum, arg1 GLenum, arg2 *GLint) 
{
	C.glGetConvolutionParameterivEXT(C.GLenum(arg0), C.GLenum(arg1), (*C.GLint)(arg2));
}

//extern void glGetSeparableFilterEXT (GLenum, GLenum, GLenum, GLvoid *, GLvoid *, GLvoid *)
func GetSeparableFilterEXT(arg0 GLenum, arg1 GLenum, arg2 GLenum, arg3 unsafe.Pointer, arg4 unsafe.Pointer, arg5 unsafe.Pointer) 
{
	C.glGetSeparableFilterEXT(C.GLenum(arg0), C.GLenum(arg1), C.GLenum(arg2), arg3, arg4, arg5);
}

//extern void glSeparableFilter2DEXT (GLenum, GLenum, GLsizei, GLsizei, GLenum, GLenum, const GLvoid *, const GLvoid *)
func SeparableFilter2DEXT(arg0 GLenum, arg1 GLenum, arg2 GLsizei, arg3 GLsizei, arg4 GLenum, arg5 GLenum, arg6 unsafe.Pointer, arg7 unsafe.Pointer) 
{
	C.glSeparableFilter2DEXT(C.GLenum(arg0), C.GLenum(arg1), C.GLsizei(arg2), C.GLsizei(arg3), C.GLenum(arg4), C.GLenum(arg5), arg6, arg7);
}

//extern void glColorTableSGI (GLenum, GLenum, GLsizei, GLenum, GLenum, const GLvoid *)
func ColorTableSGI(arg0 GLenum, arg1 GLenum, arg2 GLsizei, arg3 GLenum, arg4 GLenum, arg5 unsafe.Pointer) 
{
	C.glColorTableSGI(C.GLenum(arg0), C.GLenum(arg1), C.GLsizei(arg2), C.GLenum(arg3), C.GLenum(arg4), arg5);
}

//extern void glColorTableParameterfvSGI (GLenum, GLenum, const GLfloat *)
func ColorTableParameterfvSGI(arg0 GLenum, arg1 GLenum, arg2 *GLfloat) 
{
	C.glColorTableParameterfvSGI(C.GLenum(arg0), C.GLenum(arg1), (*C.GLfloat)(arg2));
}

//extern void glColorTableParameterivSGI (GLenum, GLenum, const GLint *)
func ColorTableParameterivSGI(arg0 GLenum, arg1 GLenum, arg2 *GLint) 
{
	C.glColorTableParameterivSGI(C.GLenum(arg0), C.GLenum(arg1), (*C.GLint)(arg2));
}

//extern void glCopyColorTableSGI (GLenum, GLenum, GLint, GLint, GLsizei)
func CopyColorTableSGI(arg0 GLenum, arg1 GLenum, arg2 GLint, arg3 GLint, arg4 GLsizei) 
{
	C.glCopyColorTableSGI(C.GLenum(arg0), C.GLenum(arg1), C.GLint(arg2), C.GLint(arg3), C.GLsizei(arg4));
}

//extern void glGetColorTableSGI (GLenum, GLenum, GLenum, GLvoid *)
func GetColorTableSGI(arg0 GLenum, arg1 GLenum, arg2 GLenum, arg3 unsafe.Pointer) 
{
	C.glGetColorTableSGI(C.GLenum(arg0), C.GLenum(arg1), C.GLenum(arg2), arg3);
}

//extern void glGetColorTableParameterfvSGI (GLenum, GLenum, GLfloat *)
func GetColorTableParameterfvSGI(arg0 GLenum, arg1 GLenum, arg2 *GLfloat) 
{
	C.glGetColorTableParameterfvSGI(C.GLenum(arg0), C.GLenum(arg1), (*C.GLfloat)(arg2));
}

//extern void glGetColorTableParameterivSGI (GLenum, GLenum, GLint *)
func GetColorTableParameterivSGI(arg0 GLenum, arg1 GLenum, arg2 *GLint) 
{
	C.glGetColorTableParameterivSGI(C.GLenum(arg0), C.GLenum(arg1), (*C.GLint)(arg2));
}

//extern void glPixelTexGenSGIX (GLenum)
func PixelTexGenSGIX(arg0 GLenum) 
{
	C.glPixelTexGenSGIX(C.GLenum(arg0));
}

//extern void glPixelTexGenParameteriSGIS (GLenum, GLint)
func PixelTexGenParameteriSGIS(arg0 GLenum, arg1 GLint) 
{
	C.glPixelTexGenParameteriSGIS(C.GLenum(arg0), C.GLint(arg1));
}

//extern void glPixelTexGenParameterivSGIS (GLenum, const GLint *)
func PixelTexGenParameterivSGIS(arg0 GLenum, arg1 *GLint) 
{
	C.glPixelTexGenParameterivSGIS(C.GLenum(arg0), (*C.GLint)(arg1));
}

//extern void glPixelTexGenParameterfSGIS (GLenum, GLfloat)
func PixelTexGenParameterfSGIS(arg0 GLenum, arg1 GLfloat) 
{
	C.glPixelTexGenParameterfSGIS(C.GLenum(arg0), C.GLfloat(arg1));
}

//extern void glPixelTexGenParameterfvSGIS (GLenum, const GLfloat *)
func PixelTexGenParameterfvSGIS(arg0 GLenum, arg1 *GLfloat) 
{
	C.glPixelTexGenParameterfvSGIS(C.GLenum(arg0), (*C.GLfloat)(arg1));
}

//extern void glGetPixelTexGenParameterivSGIS (GLenum, GLint *)
func GetPixelTexGenParameterivSGIS(arg0 GLenum, arg1 *GLint) 
{
	C.glGetPixelTexGenParameterivSGIS(C.GLenum(arg0), (*C.GLint)(arg1));
}

//extern void glGetPixelTexGenParameterfvSGIS (GLenum, GLfloat *)
func GetPixelTexGenParameterfvSGIS(arg0 GLenum, arg1 *GLfloat) 
{
	C.glGetPixelTexGenParameterfvSGIS(C.GLenum(arg0), (*C.GLfloat)(arg1));
}

//extern void glTexImage4DSGIS (GLenum, GLint, GLenum, GLsizei, GLsizei, GLsizei, GLsizei, GLint, GLenum, GLenum, const GLvoid *)
func TexImage4DSGIS(arg0 GLenum, arg1 GLint, arg2 GLenum, arg3 GLsizei, arg4 GLsizei, arg5 GLsizei, arg6 GLsizei, arg7 GLint, arg8 GLenum, arg9 GLenum, arg10 unsafe.Pointer) 
{
	C.glTexImage4DSGIS(C.GLenum(arg0), C.GLint(arg1), C.GLenum(arg2), C.GLsizei(arg3), C.GLsizei(arg4), C.GLsizei(arg5), C.GLsizei(arg6), C.GLint(arg7), C.GLenum(arg8), C.GLenum(arg9), arg10);
}

//extern void glTexSubImage4DSGIS (GLenum, GLint, GLint, GLint, GLint, GLint, GLsizei, GLsizei, GLsizei, GLsizei, GLenum, GLenum, const GLvoid *)
func TexSubImage4DSGIS(arg0 GLenum, arg1 GLint, arg2 GLint, arg3 GLint, arg4 GLint, arg5 GLint, arg6 GLsizei, arg7 GLsizei, arg8 GLsizei, arg9 GLsizei, arg10 GLenum, arg11 GLenum, arg12 unsafe.Pointer) 
{
	C.glTexSubImage4DSGIS(C.GLenum(arg0), C.GLint(arg1), C.GLint(arg2), C.GLint(arg3), C.GLint(arg4), C.GLint(arg5), C.GLsizei(arg6), C.GLsizei(arg7), C.GLsizei(arg8), C.GLsizei(arg9), C.GLenum(arg10), C.GLenum(arg11), arg12);
}

//extern GLboolean glAreTexturesResidentEXT (GLsizei, const GLuint *, GLboolean *)
func AreTexturesResidentEXT(arg0 GLsizei, arg1 *GLuint, arg2 *GLboolean) GLboolean
{
	return GLboolean(C.glAreTexturesResidentEXT(C.GLsizei(arg0), (*C.GLuint)(arg1), (*C.GLboolean)(arg2)));
}

//extern void glBindTextureEXT (GLenum, GLuint)
func BindTextureEXT(arg0 GLenum, arg1 GLuint) 
{
	C.glBindTextureEXT(C.GLenum(arg0), C.GLuint(arg1));
}

//extern void glDeleteTexturesEXT (GLsizei, const GLuint *)
func DeleteTexturesEXT(arg0 GLsizei, arg1 *GLuint) 
{
	C.glDeleteTexturesEXT(C.GLsizei(arg0), (*C.GLuint)(arg1));
}

//extern void glGenTexturesEXT (GLsizei, GLuint *)
func GenTexturesEXT(arg0 GLsizei, arg1 *GLuint) 
{
	C.glGenTexturesEXT(C.GLsizei(arg0), (*C.GLuint)(arg1));
}

//extern GLboolean glIsTextureEXT (GLuint)
func IsTextureEXT(arg0 GLuint) GLboolean
{
	return GLboolean(C.glIsTextureEXT(C.GLuint(arg0)));
}

//extern void glPrioritizeTexturesEXT (GLsizei, const GLuint *, const GLclampf *)
func PrioritizeTexturesEXT(arg0 GLsizei, arg1 *GLuint, arg2 *GLclampf) 
{
	C.glPrioritizeTexturesEXT(C.GLsizei(arg0), (*C.GLuint)(arg1), (*C.GLclampf)(arg2));
}

//extern void glDetailTexFuncSGIS (GLenum, GLsizei, const GLfloat *)
func DetailTexFuncSGIS(arg0 GLenum, arg1 GLsizei, arg2 *GLfloat) 
{
	C.glDetailTexFuncSGIS(C.GLenum(arg0), C.GLsizei(arg1), (*C.GLfloat)(arg2));
}

//extern void glGetDetailTexFuncSGIS (GLenum, GLfloat *)
func GetDetailTexFuncSGIS(arg0 GLenum, arg1 *GLfloat) 
{
	C.glGetDetailTexFuncSGIS(C.GLenum(arg0), (*C.GLfloat)(arg1));
}

//extern void glSharpenTexFuncSGIS (GLenum, GLsizei, const GLfloat *)
func SharpenTexFuncSGIS(arg0 GLenum, arg1 GLsizei, arg2 *GLfloat) 
{
	C.glSharpenTexFuncSGIS(C.GLenum(arg0), C.GLsizei(arg1), (*C.GLfloat)(arg2));
}

//extern void glGetSharpenTexFuncSGIS (GLenum, GLfloat *)
func GetSharpenTexFuncSGIS(arg0 GLenum, arg1 *GLfloat) 
{
	C.glGetSharpenTexFuncSGIS(C.GLenum(arg0), (*C.GLfloat)(arg1));
}

//extern void glSampleMaskSGIS (GLclampf, GLboolean)
func SampleMaskSGIS(arg0 GLclampf, arg1 GLboolean) 
{
	C.glSampleMaskSGIS(C.GLclampf(arg0), C.GLboolean(arg1));
}

//extern void glSamplePatternSGIS (GLenum)
func SamplePatternSGIS(arg0 GLenum) 
{
	C.glSamplePatternSGIS(C.GLenum(arg0));
}

//extern void glArrayElementEXT (GLint)
func ArrayElementEXT(arg0 GLint) 
{
	C.glArrayElementEXT(C.GLint(arg0));
}

//extern void glColorPointerEXT (GLint, GLenum, GLsizei, GLsizei, const GLvoid *)
func ColorPointerEXT(arg0 GLint, arg1 GLenum, arg2 GLsizei, arg3 GLsizei, arg4 unsafe.Pointer) 
{
	C.glColorPointerEXT(C.GLint(arg0), C.GLenum(arg1), C.GLsizei(arg2), C.GLsizei(arg3), arg4);
}

//extern void glDrawArraysEXT (GLenum, GLint, GLsizei)
func DrawArraysEXT(arg0 GLenum, arg1 GLint, arg2 GLsizei) 
{
	C.glDrawArraysEXT(C.GLenum(arg0), C.GLint(arg1), C.GLsizei(arg2));
}

//extern void glEdgeFlagPointerEXT (GLsizei, GLsizei, const GLboolean *)
func EdgeFlagPointerEXT(arg0 GLsizei, arg1 GLsizei, arg2 *GLboolean) 
{
	C.glEdgeFlagPointerEXT(C.GLsizei(arg0), C.GLsizei(arg1), (*C.GLboolean)(arg2));
}

//extern void glGetPointervEXT (GLenum, GLvoid* *)
func GetPointervEXT(arg0 GLenum, arg1 *unsafe.Pointer) 
{
	C.glGetPointervEXT(C.GLenum(arg0), arg1);
}

//extern void glIndexPointerEXT (GLenum, GLsizei, GLsizei, const GLvoid *)
func IndexPointerEXT(arg0 GLenum, arg1 GLsizei, arg2 GLsizei, arg3 unsafe.Pointer) 
{
	C.glIndexPointerEXT(C.GLenum(arg0), C.GLsizei(arg1), C.GLsizei(arg2), arg3);
}

//extern void glNormalPointerEXT (GLenum, GLsizei, GLsizei, const GLvoid *)
func NormalPointerEXT(arg0 GLenum, arg1 GLsizei, arg2 GLsizei, arg3 unsafe.Pointer) 
{
	C.glNormalPointerEXT(C.GLenum(arg0), C.GLsizei(arg1), C.GLsizei(arg2), arg3);
}

//extern void glTexCoordPointerEXT (GLint, GLenum, GLsizei, GLsizei, const GLvoid *)
func TexCoordPointerEXT(arg0 GLint, arg1 GLenum, arg2 GLsizei, arg3 GLsizei, arg4 unsafe.Pointer) 
{
	C.glTexCoordPointerEXT(C.GLint(arg0), C.GLenum(arg1), C.GLsizei(arg2), C.GLsizei(arg3), arg4);
}

//extern void glVertexPointerEXT (GLint, GLenum, GLsizei, GLsizei, const GLvoid *)
func VertexPointerEXT(arg0 GLint, arg1 GLenum, arg2 GLsizei, arg3 GLsizei, arg4 unsafe.Pointer) 
{
	C.glVertexPointerEXT(C.GLint(arg0), C.GLenum(arg1), C.GLsizei(arg2), C.GLsizei(arg3), arg4);
}

//extern void glBlendEquationEXT (GLenum)
func BlendEquationEXT(arg0 GLenum) 
{
	C.glBlendEquationEXT(C.GLenum(arg0));
}

//extern void glSpriteParameterfSGIX (GLenum, GLfloat)
func SpriteParameterfSGIX(arg0 GLenum, arg1 GLfloat) 
{
	C.glSpriteParameterfSGIX(C.GLenum(arg0), C.GLfloat(arg1));
}

//extern void glSpriteParameterfvSGIX (GLenum, const GLfloat *)
func SpriteParameterfvSGIX(arg0 GLenum, arg1 *GLfloat) 
{
	C.glSpriteParameterfvSGIX(C.GLenum(arg0), (*C.GLfloat)(arg1));
}

//extern void glSpriteParameteriSGIX (GLenum, GLint)
func SpriteParameteriSGIX(arg0 GLenum, arg1 GLint) 
{
	C.glSpriteParameteriSGIX(C.GLenum(arg0), C.GLint(arg1));
}

//extern void glSpriteParameterivSGIX (GLenum, const GLint *)
func SpriteParameterivSGIX(arg0 GLenum, arg1 *GLint) 
{
	C.glSpriteParameterivSGIX(C.GLenum(arg0), (*C.GLint)(arg1));
}

//extern void glPointParameterfEXT (GLenum, GLfloat)
func PointParameterfEXT(arg0 GLenum, arg1 GLfloat) 
{
	C.glPointParameterfEXT(C.GLenum(arg0), C.GLfloat(arg1));
}

//extern void glPointParameterfvEXT (GLenum, const GLfloat *)
func PointParameterfvEXT(arg0 GLenum, arg1 *GLfloat) 
{
	C.glPointParameterfvEXT(C.GLenum(arg0), (*C.GLfloat)(arg1));
}

//extern void glPointParameterfSGIS (GLenum, GLfloat)
func PointParameterfSGIS(arg0 GLenum, arg1 GLfloat) 
{
	C.glPointParameterfSGIS(C.GLenum(arg0), C.GLfloat(arg1));
}

//extern void glPointParameterfvSGIS (GLenum, const GLfloat *)
func PointParameterfvSGIS(arg0 GLenum, arg1 *GLfloat) 
{
	C.glPointParameterfvSGIS(C.GLenum(arg0), (*C.GLfloat)(arg1));
}

//extern GLint glGetInstrumentsSGIX (void)
func GetInstrumentsSGIX() GLint
{
	return GLint(C.glGetInstrumentsSGIX());
}

//extern void glInstrumentsBufferSGIX (GLsizei, GLint *)
func InstrumentsBufferSGIX(arg0 GLsizei, arg1 *GLint) 
{
	C.glInstrumentsBufferSGIX(C.GLsizei(arg0), (*C.GLint)(arg1));
}

//extern GLint glPollInstrumentsSGIX (GLint *)
func PollInstrumentsSGIX(arg0 *GLint) GLint
{
	return GLint(C.glPollInstrumentsSGIX((*C.GLint)(arg0)));
}

//extern void glReadInstrumentsSGIX (GLint)
func ReadInstrumentsSGIX(arg0 GLint) 
{
	C.glReadInstrumentsSGIX(C.GLint(arg0));
}

//extern void glStartInstrumentsSGIX (void)
func StartInstrumentsSGIX() 
{
	C.glStartInstrumentsSGIX();
}

//extern void glStopInstrumentsSGIX (GLint)
func StopInstrumentsSGIX(arg0 GLint) 
{
	C.glStopInstrumentsSGIX(C.GLint(arg0));
}

//extern void glFrameZoomSGIX (GLint)
func FrameZoomSGIX(arg0 GLint) 
{
	C.glFrameZoomSGIX(C.GLint(arg0));
}

//extern void glTagSampleBufferSGIX (void)
func TagSampleBufferSGIX() 
{
	C.glTagSampleBufferSGIX();
}

//extern void glDeformationMap3dSGIX (GLenum, GLdouble, GLdouble, GLint, GLint, GLdouble, GLdouble, GLint, GLint, GLdouble, GLdouble, GLint, GLint, const GLdouble *)
func DeformationMap3dSGIX(arg0 GLenum, arg1 GLdouble, arg2 GLdouble, arg3 GLint, arg4 GLint, arg5 GLdouble, arg6 GLdouble, arg7 GLint, arg8 GLint, arg9 GLdouble, arg10 GLdouble, arg11 GLint, arg12 GLint, arg13 *GLdouble) 
{
	C.glDeformationMap3dSGIX(C.GLenum(arg0), C.GLdouble(arg1), C.GLdouble(arg2), C.GLint(arg3), C.GLint(arg4), C.GLdouble(arg5), C.GLdouble(arg6), C.GLint(arg7), C.GLint(arg8), C.GLdouble(arg9), C.GLdouble(arg10), C.GLint(arg11), C.GLint(arg12), (*C.GLdouble)(arg13));
}

//extern void glDeformationMap3fSGIX (GLenum, GLfloat, GLfloat, GLint, GLint, GLfloat, GLfloat, GLint, GLint, GLfloat, GLfloat, GLint, GLint, const GLfloat *)
func DeformationMap3fSGIX(arg0 GLenum, arg1 GLfloat, arg2 GLfloat, arg3 GLint, arg4 GLint, arg5 GLfloat, arg6 GLfloat, arg7 GLint, arg8 GLint, arg9 GLfloat, arg10 GLfloat, arg11 GLint, arg12 GLint, arg13 *GLfloat) 
{
	C.glDeformationMap3fSGIX(C.GLenum(arg0), C.GLfloat(arg1), C.GLfloat(arg2), C.GLint(arg3), C.GLint(arg4), C.GLfloat(arg5), C.GLfloat(arg6), C.GLint(arg7), C.GLint(arg8), C.GLfloat(arg9), C.GLfloat(arg10), C.GLint(arg11), C.GLint(arg12), (*C.GLfloat)(arg13));
}

//extern void glDeformSGIX (GLbitfield)
func DeformSGIX(arg0 GLbitfield) 
{
	C.glDeformSGIX(C.GLbitfield(arg0));
}

//extern void glLoadIdentityDeformationMapSGIX (GLbitfield)
func LoadIdentityDeformationMapSGIX(arg0 GLbitfield) 
{
	C.glLoadIdentityDeformationMapSGIX(C.GLbitfield(arg0));
}

//extern void glReferencePlaneSGIX (const GLdouble *)
func ReferencePlaneSGIX(arg0 *GLdouble) 
{
	C.glReferencePlaneSGIX((*C.GLdouble)(arg0));
}

//extern void glFlushRasterSGIX (void)
func FlushRasterSGIX() 
{
	C.glFlushRasterSGIX();
}

//extern void glFogFuncSGIS (GLsizei, const GLfloat *)
func FogFuncSGIS(arg0 GLsizei, arg1 *GLfloat) 
{
	C.glFogFuncSGIS(C.GLsizei(arg0), (*C.GLfloat)(arg1));
}

//extern void glGetFogFuncSGIS (GLfloat *)
func GetFogFuncSGIS(arg0 *GLfloat) 
{
	C.glGetFogFuncSGIS((*C.GLfloat)(arg0));
}

//extern void glImageTransformParameteriHP (GLenum, GLenum, GLint)
func ImageTransformParameteriHP(arg0 GLenum, arg1 GLenum, arg2 GLint) 
{
	C.glImageTransformParameteriHP(C.GLenum(arg0), C.GLenum(arg1), C.GLint(arg2));
}

//extern void glImageTransformParameterfHP (GLenum, GLenum, GLfloat)
func ImageTransformParameterfHP(arg0 GLenum, arg1 GLenum, arg2 GLfloat) 
{
	C.glImageTransformParameterfHP(C.GLenum(arg0), C.GLenum(arg1), C.GLfloat(arg2));
}

//extern void glImageTransformParameterivHP (GLenum, GLenum, const GLint *)
func ImageTransformParameterivHP(arg0 GLenum, arg1 GLenum, arg2 *GLint) 
{
	C.glImageTransformParameterivHP(C.GLenum(arg0), C.GLenum(arg1), (*C.GLint)(arg2));
}

//extern void glImageTransformParameterfvHP (GLenum, GLenum, const GLfloat *)
func ImageTransformParameterfvHP(arg0 GLenum, arg1 GLenum, arg2 *GLfloat) 
{
	C.glImageTransformParameterfvHP(C.GLenum(arg0), C.GLenum(arg1), (*C.GLfloat)(arg2));
}

//extern void glGetImageTransformParameterivHP (GLenum, GLenum, GLint *)
func GetImageTransformParameterivHP(arg0 GLenum, arg1 GLenum, arg2 *GLint) 
{
	C.glGetImageTransformParameterivHP(C.GLenum(arg0), C.GLenum(arg1), (*C.GLint)(arg2));
}

//extern void glGetImageTransformParameterfvHP (GLenum, GLenum, GLfloat *)
func GetImageTransformParameterfvHP(arg0 GLenum, arg1 GLenum, arg2 *GLfloat) 
{
	C.glGetImageTransformParameterfvHP(C.GLenum(arg0), C.GLenum(arg1), (*C.GLfloat)(arg2));
}

//extern void glColorSubTableEXT (GLenum, GLsizei, GLsizei, GLenum, GLenum, const GLvoid *)
func ColorSubTableEXT(arg0 GLenum, arg1 GLsizei, arg2 GLsizei, arg3 GLenum, arg4 GLenum, arg5 unsafe.Pointer) 
{
	C.glColorSubTableEXT(C.GLenum(arg0), C.GLsizei(arg1), C.GLsizei(arg2), C.GLenum(arg3), C.GLenum(arg4), arg5);
}

//extern void glCopyColorSubTableEXT (GLenum, GLsizei, GLint, GLint, GLsizei)
func CopyColorSubTableEXT(arg0 GLenum, arg1 GLsizei, arg2 GLint, arg3 GLint, arg4 GLsizei) 
{
	C.glCopyColorSubTableEXT(C.GLenum(arg0), C.GLsizei(arg1), C.GLint(arg2), C.GLint(arg3), C.GLsizei(arg4));
}

//extern void glHintPGI (GLenum, GLint)
func HintPGI(arg0 GLenum, arg1 GLint) 
{
	C.glHintPGI(C.GLenum(arg0), C.GLint(arg1));
}

//extern void glColorTableEXT (GLenum, GLenum, GLsizei, GLenum, GLenum, const GLvoid *)
func ColorTableEXT(arg0 GLenum, arg1 GLenum, arg2 GLsizei, arg3 GLenum, arg4 GLenum, arg5 unsafe.Pointer) 
{
	C.glColorTableEXT(C.GLenum(arg0), C.GLenum(arg1), C.GLsizei(arg2), C.GLenum(arg3), C.GLenum(arg4), arg5);
}

//extern void glGetColorTableEXT (GLenum, GLenum, GLenum, GLvoid *)
func GetColorTableEXT(arg0 GLenum, arg1 GLenum, arg2 GLenum, arg3 unsafe.Pointer) 
{
	C.glGetColorTableEXT(C.GLenum(arg0), C.GLenum(arg1), C.GLenum(arg2), arg3);
}

//extern void glGetColorTableParameterivEXT (GLenum, GLenum, GLint *)
func GetColorTableParameterivEXT(arg0 GLenum, arg1 GLenum, arg2 *GLint) 
{
	C.glGetColorTableParameterivEXT(C.GLenum(arg0), C.GLenum(arg1), (*C.GLint)(arg2));
}

//extern void glGetColorTableParameterfvEXT (GLenum, GLenum, GLfloat *)
func GetColorTableParameterfvEXT(arg0 GLenum, arg1 GLenum, arg2 *GLfloat) 
{
	C.glGetColorTableParameterfvEXT(C.GLenum(arg0), C.GLenum(arg1), (*C.GLfloat)(arg2));
}

//extern void glGetListParameterfvSGIX (GLuint, GLenum, GLfloat *)
func GetListParameterfvSGIX(arg0 GLuint, arg1 GLenum, arg2 *GLfloat) 
{
	C.glGetListParameterfvSGIX(C.GLuint(arg0), C.GLenum(arg1), (*C.GLfloat)(arg2));
}

//extern void glGetListParameterivSGIX (GLuint, GLenum, GLint *)
func GetListParameterivSGIX(arg0 GLuint, arg1 GLenum, arg2 *GLint) 
{
	C.glGetListParameterivSGIX(C.GLuint(arg0), C.GLenum(arg1), (*C.GLint)(arg2));
}

//extern void glListParameterfSGIX (GLuint, GLenum, GLfloat)
func ListParameterfSGIX(arg0 GLuint, arg1 GLenum, arg2 GLfloat) 
{
	C.glListParameterfSGIX(C.GLuint(arg0), C.GLenum(arg1), C.GLfloat(arg2));
}

//extern void glListParameterfvSGIX (GLuint, GLenum, const GLfloat *)
func ListParameterfvSGIX(arg0 GLuint, arg1 GLenum, arg2 *GLfloat) 
{
	C.glListParameterfvSGIX(C.GLuint(arg0), C.GLenum(arg1), (*C.GLfloat)(arg2));
}

//extern void glListParameteriSGIX (GLuint, GLenum, GLint)
func ListParameteriSGIX(arg0 GLuint, arg1 GLenum, arg2 GLint) 
{
	C.glListParameteriSGIX(C.GLuint(arg0), C.GLenum(arg1), C.GLint(arg2));
}

//extern void glListParameterivSGIX (GLuint, GLenum, const GLint *)
func ListParameterivSGIX(arg0 GLuint, arg1 GLenum, arg2 *GLint) 
{
	C.glListParameterivSGIX(C.GLuint(arg0), C.GLenum(arg1), (*C.GLint)(arg2));
}

//extern void glIndexMaterialEXT (GLenum, GLenum)
func IndexMaterialEXT(arg0 GLenum, arg1 GLenum) 
{
	C.glIndexMaterialEXT(C.GLenum(arg0), C.GLenum(arg1));
}

//extern void glIndexFuncEXT (GLenum, GLclampf)
func IndexFuncEXT(arg0 GLenum, arg1 GLclampf) 
{
	C.glIndexFuncEXT(C.GLenum(arg0), C.GLclampf(arg1));
}

//extern void glLockArraysEXT (GLint, GLsizei)
func LockArraysEXT(arg0 GLint, arg1 GLsizei) 
{
	C.glLockArraysEXT(C.GLint(arg0), C.GLsizei(arg1));
}

//extern void glUnlockArraysEXT (void)
func UnlockArraysEXT() 
{
	C.glUnlockArraysEXT();
}

//extern void glCullParameterdvEXT (GLenum, GLdouble *)
func CullParameterdvEXT(arg0 GLenum, arg1 *GLdouble) 
{
	C.glCullParameterdvEXT(C.GLenum(arg0), (*C.GLdouble)(arg1));
}

//extern void glCullParameterfvEXT (GLenum, GLfloat *)
func CullParameterfvEXT(arg0 GLenum, arg1 *GLfloat) 
{
	C.glCullParameterfvEXT(C.GLenum(arg0), (*C.GLfloat)(arg1));
}

//extern void glFragmentColorMaterialSGIX (GLenum, GLenum)
func FragmentColorMaterialSGIX(arg0 GLenum, arg1 GLenum) 
{
	C.glFragmentColorMaterialSGIX(C.GLenum(arg0), C.GLenum(arg1));
}

//extern void glFragmentLightfSGIX (GLenum, GLenum, GLfloat)
func FragmentLightfSGIX(arg0 GLenum, arg1 GLenum, arg2 GLfloat) 
{
	C.glFragmentLightfSGIX(C.GLenum(arg0), C.GLenum(arg1), C.GLfloat(arg2));
}

//extern void glFragmentLightfvSGIX (GLenum, GLenum, const GLfloat *)
func FragmentLightfvSGIX(arg0 GLenum, arg1 GLenum, arg2 *GLfloat) 
{
	C.glFragmentLightfvSGIX(C.GLenum(arg0), C.GLenum(arg1), (*C.GLfloat)(arg2));
}

//extern void glFragmentLightiSGIX (GLenum, GLenum, GLint)
func FragmentLightiSGIX(arg0 GLenum, arg1 GLenum, arg2 GLint) 
{
	C.glFragmentLightiSGIX(C.GLenum(arg0), C.GLenum(arg1), C.GLint(arg2));
}

//extern void glFragmentLightivSGIX (GLenum, GLenum, const GLint *)
func FragmentLightivSGIX(arg0 GLenum, arg1 GLenum, arg2 *GLint) 
{
	C.glFragmentLightivSGIX(C.GLenum(arg0), C.GLenum(arg1), (*C.GLint)(arg2));
}

//extern void glFragmentLightModelfSGIX (GLenum, GLfloat)
func FragmentLightModelfSGIX(arg0 GLenum, arg1 GLfloat) 
{
	C.glFragmentLightModelfSGIX(C.GLenum(arg0), C.GLfloat(arg1));
}

//extern void glFragmentLightModelfvSGIX (GLenum, const GLfloat *)
func FragmentLightModelfvSGIX(arg0 GLenum, arg1 *GLfloat) 
{
	C.glFragmentLightModelfvSGIX(C.GLenum(arg0), (*C.GLfloat)(arg1));
}

//extern void glFragmentLightModeliSGIX (GLenum, GLint)
func FragmentLightModeliSGIX(arg0 GLenum, arg1 GLint) 
{
	C.glFragmentLightModeliSGIX(C.GLenum(arg0), C.GLint(arg1));
}

//extern void glFragmentLightModelivSGIX (GLenum, const GLint *)
func FragmentLightModelivSGIX(arg0 GLenum, arg1 *GLint) 
{
	C.glFragmentLightModelivSGIX(C.GLenum(arg0), (*C.GLint)(arg1));
}

//extern void glFragmentMaterialfSGIX (GLenum, GLenum, GLfloat)
func FragmentMaterialfSGIX(arg0 GLenum, arg1 GLenum, arg2 GLfloat) 
{
	C.glFragmentMaterialfSGIX(C.GLenum(arg0), C.GLenum(arg1), C.GLfloat(arg2));
}

//extern void glFragmentMaterialfvSGIX (GLenum, GLenum, const GLfloat *)
func FragmentMaterialfvSGIX(arg0 GLenum, arg1 GLenum, arg2 *GLfloat) 
{
	C.glFragmentMaterialfvSGIX(C.GLenum(arg0), C.GLenum(arg1), (*C.GLfloat)(arg2));
}

//extern void glFragmentMaterialiSGIX (GLenum, GLenum, GLint)
func FragmentMaterialiSGIX(arg0 GLenum, arg1 GLenum, arg2 GLint) 
{
	C.glFragmentMaterialiSGIX(C.GLenum(arg0), C.GLenum(arg1), C.GLint(arg2));
}

//extern void glFragmentMaterialivSGIX (GLenum, GLenum, const GLint *)
func FragmentMaterialivSGIX(arg0 GLenum, arg1 GLenum, arg2 *GLint) 
{
	C.glFragmentMaterialivSGIX(C.GLenum(arg0), C.GLenum(arg1), (*C.GLint)(arg2));
}

//extern void glGetFragmentLightfvSGIX (GLenum, GLenum, GLfloat *)
func GetFragmentLightfvSGIX(arg0 GLenum, arg1 GLenum, arg2 *GLfloat) 
{
	C.glGetFragmentLightfvSGIX(C.GLenum(arg0), C.GLenum(arg1), (*C.GLfloat)(arg2));
}

//extern void glGetFragmentLightivSGIX (GLenum, GLenum, GLint *)
func GetFragmentLightivSGIX(arg0 GLenum, arg1 GLenum, arg2 *GLint) 
{
	C.glGetFragmentLightivSGIX(C.GLenum(arg0), C.GLenum(arg1), (*C.GLint)(arg2));
}

//extern void glGetFragmentMaterialfvSGIX (GLenum, GLenum, GLfloat *)
func GetFragmentMaterialfvSGIX(arg0 GLenum, arg1 GLenum, arg2 *GLfloat) 
{
	C.glGetFragmentMaterialfvSGIX(C.GLenum(arg0), C.GLenum(arg1), (*C.GLfloat)(arg2));
}

//extern void glGetFragmentMaterialivSGIX (GLenum, GLenum, GLint *)
func GetFragmentMaterialivSGIX(arg0 GLenum, arg1 GLenum, arg2 *GLint) 
{
	C.glGetFragmentMaterialivSGIX(C.GLenum(arg0), C.GLenum(arg1), (*C.GLint)(arg2));
}

//extern void glLightEnviSGIX (GLenum, GLint)
func LightEnviSGIX(arg0 GLenum, arg1 GLint) 
{
	C.glLightEnviSGIX(C.GLenum(arg0), C.GLint(arg1));
}

//extern void glDrawRangeElementsEXT (GLenum, GLuint, GLuint, GLsizei, GLenum, const GLvoid *)
func DrawRangeElementsEXT(arg0 GLenum, arg1 GLuint, arg2 GLuint, arg3 GLsizei, arg4 GLenum, arg5 unsafe.Pointer) 
{
	C.glDrawRangeElementsEXT(C.GLenum(arg0), C.GLuint(arg1), C.GLuint(arg2), C.GLsizei(arg3), C.GLenum(arg4), arg5);
}

//extern void glApplyTextureEXT (GLenum)
func ApplyTextureEXT(arg0 GLenum) 
{
	C.glApplyTextureEXT(C.GLenum(arg0));
}

//extern void glTextureLightEXT (GLenum)
func TextureLightEXT(arg0 GLenum) 
{
	C.glTextureLightEXT(C.GLenum(arg0));
}

//extern void glTextureMaterialEXT (GLenum, GLenum)
func TextureMaterialEXT(arg0 GLenum, arg1 GLenum) 
{
	C.glTextureMaterialEXT(C.GLenum(arg0), C.GLenum(arg1));
}

//extern void glAsyncMarkerSGIX (GLuint)
func AsyncMarkerSGIX(arg0 GLuint) 
{
	C.glAsyncMarkerSGIX(C.GLuint(arg0));
}

//extern GLint glFinishAsyncSGIX (GLuint *)
func FinishAsyncSGIX(arg0 *GLuint) GLint
{
	return GLint(C.glFinishAsyncSGIX((*C.GLuint)(arg0)));
}

//extern GLint glPollAsyncSGIX (GLuint *)
func PollAsyncSGIX(arg0 *GLuint) GLint
{
	return GLint(C.glPollAsyncSGIX((*C.GLuint)(arg0)));
}

//extern GLuint glGenAsyncMarkersSGIX (GLsizei)
func GenAsyncMarkersSGIX(arg0 GLsizei) GLuint
{
	return GLuint(C.glGenAsyncMarkersSGIX(C.GLsizei(arg0)));
}

//extern void glDeleteAsyncMarkersSGIX (GLuint, GLsizei)
func DeleteAsyncMarkersSGIX(arg0 GLuint, arg1 GLsizei) 
{
	C.glDeleteAsyncMarkersSGIX(C.GLuint(arg0), C.GLsizei(arg1));
}

//extern GLboolean glIsAsyncMarkerSGIX (GLuint)
func IsAsyncMarkerSGIX(arg0 GLuint) GLboolean
{
	return GLboolean(C.glIsAsyncMarkerSGIX(C.GLuint(arg0)));
}

//extern void glVertexPointervINTEL (GLint, GLenum, const GLvoid* *)
func VertexPointervINTEL(arg0 GLint, arg1 GLenum, arg2 *unsafe.Pointer) 
{
	C.glVertexPointervINTEL(C.GLint(arg0), C.GLenum(arg1), arg2);
}

//extern void glNormalPointervINTEL (GLenum, const GLvoid* *)
func NormalPointervINTEL(arg0 GLenum, arg1 *unsafe.Pointer) 
{
	C.glNormalPointervINTEL(C.GLenum(arg0), arg1);
}

//extern void glColorPointervINTEL (GLint, GLenum, const GLvoid* *)
func ColorPointervINTEL(arg0 GLint, arg1 GLenum, arg2 *unsafe.Pointer) 
{
	C.glColorPointervINTEL(C.GLint(arg0), C.GLenum(arg1), arg2);
}

//extern void glTexCoordPointervINTEL (GLint, GLenum, const GLvoid* *)
func TexCoordPointervINTEL(arg0 GLint, arg1 GLenum, arg2 *unsafe.Pointer) 
{
	C.glTexCoordPointervINTEL(C.GLint(arg0), C.GLenum(arg1), arg2);
}

//extern void glPixelTransformParameteriEXT (GLenum, GLenum, GLint)
func PixelTransformParameteriEXT(arg0 GLenum, arg1 GLenum, arg2 GLint) 
{
	C.glPixelTransformParameteriEXT(C.GLenum(arg0), C.GLenum(arg1), C.GLint(arg2));
}

//extern void glPixelTransformParameterfEXT (GLenum, GLenum, GLfloat)
func PixelTransformParameterfEXT(arg0 GLenum, arg1 GLenum, arg2 GLfloat) 
{
	C.glPixelTransformParameterfEXT(C.GLenum(arg0), C.GLenum(arg1), C.GLfloat(arg2));
}

//extern void glPixelTransformParameterivEXT (GLenum, GLenum, const GLint *)
func PixelTransformParameterivEXT(arg0 GLenum, arg1 GLenum, arg2 *GLint) 
{
	C.glPixelTransformParameterivEXT(C.GLenum(arg0), C.GLenum(arg1), (*C.GLint)(arg2));
}

//extern void glPixelTransformParameterfvEXT (GLenum, GLenum, const GLfloat *)
func PixelTransformParameterfvEXT(arg0 GLenum, arg1 GLenum, arg2 *GLfloat) 
{
	C.glPixelTransformParameterfvEXT(C.GLenum(arg0), C.GLenum(arg1), (*C.GLfloat)(arg2));
}

//extern void glSecondaryColor3bEXT (GLbyte, GLbyte, GLbyte)
func SecondaryColor3bEXT(arg0 GLbyte, arg1 GLbyte, arg2 GLbyte) 
{
	C.glSecondaryColor3bEXT(C.GLbyte(arg0), C.GLbyte(arg1), C.GLbyte(arg2));
}

//extern void glSecondaryColor3bvEXT (const GLbyte *)
func SecondaryColor3bvEXT(arg0 *GLbyte) 
{
	C.glSecondaryColor3bvEXT((*C.GLbyte)(arg0));
}

//extern void glSecondaryColor3dEXT (GLdouble, GLdouble, GLdouble)
func SecondaryColor3dEXT(arg0 GLdouble, arg1 GLdouble, arg2 GLdouble) 
{
	C.glSecondaryColor3dEXT(C.GLdouble(arg0), C.GLdouble(arg1), C.GLdouble(arg2));
}

//extern void glSecondaryColor3dvEXT (const GLdouble *)
func SecondaryColor3dvEXT(arg0 *GLdouble) 
{
	C.glSecondaryColor3dvEXT((*C.GLdouble)(arg0));
}

//extern void glSecondaryColor3fEXT (GLfloat, GLfloat, GLfloat)
func SecondaryColor3fEXT(arg0 GLfloat, arg1 GLfloat, arg2 GLfloat) 
{
	C.glSecondaryColor3fEXT(C.GLfloat(arg0), C.GLfloat(arg1), C.GLfloat(arg2));
}

//extern void glSecondaryColor3fvEXT (const GLfloat *)
func SecondaryColor3fvEXT(arg0 *GLfloat) 
{
	C.glSecondaryColor3fvEXT((*C.GLfloat)(arg0));
}

//extern void glSecondaryColor3iEXT (GLint, GLint, GLint)
func SecondaryColor3iEXT(arg0 GLint, arg1 GLint, arg2 GLint) 
{
	C.glSecondaryColor3iEXT(C.GLint(arg0), C.GLint(arg1), C.GLint(arg2));
}

//extern void glSecondaryColor3ivEXT (const GLint *)
func SecondaryColor3ivEXT(arg0 *GLint) 
{
	C.glSecondaryColor3ivEXT((*C.GLint)(arg0));
}

//extern void glSecondaryColor3sEXT (GLshort, GLshort, GLshort)
func SecondaryColor3sEXT(arg0 GLshort, arg1 GLshort, arg2 GLshort) 
{
	C.glSecondaryColor3sEXT(C.GLshort(arg0), C.GLshort(arg1), C.GLshort(arg2));
}

//extern void glSecondaryColor3svEXT (const GLshort *)
func SecondaryColor3svEXT(arg0 *GLshort) 
{
	C.glSecondaryColor3svEXT((*C.GLshort)(arg0));
}

//extern void glSecondaryColor3ubEXT (GLubyte, GLubyte, GLubyte)
func SecondaryColor3ubEXT(arg0 GLubyte, arg1 GLubyte, arg2 GLubyte) 
{
	C.glSecondaryColor3ubEXT(C.GLubyte(arg0), C.GLubyte(arg1), C.GLubyte(arg2));
}

//extern void glSecondaryColor3ubvEXT (const GLubyte *)
func SecondaryColor3ubvEXT(arg0 *GLubyte) 
{
	C.glSecondaryColor3ubvEXT((*C.GLubyte)(arg0));
}

//extern void glSecondaryColor3uiEXT (GLuint, GLuint, GLuint)
func SecondaryColor3uiEXT(arg0 GLuint, arg1 GLuint, arg2 GLuint) 
{
	C.glSecondaryColor3uiEXT(C.GLuint(arg0), C.GLuint(arg1), C.GLuint(arg2));
}

//extern void glSecondaryColor3uivEXT (const GLuint *)
func SecondaryColor3uivEXT(arg0 *GLuint) 
{
	C.glSecondaryColor3uivEXT((*C.GLuint)(arg0));
}

//extern void glSecondaryColor3usEXT (GLushort, GLushort, GLushort)
func SecondaryColor3usEXT(arg0 GLushort, arg1 GLushort, arg2 GLushort) 
{
	C.glSecondaryColor3usEXT(C.GLushort(arg0), C.GLushort(arg1), C.GLushort(arg2));
}

//extern void glSecondaryColor3usvEXT (const GLushort *)
func SecondaryColor3usvEXT(arg0 *GLushort) 
{
	C.glSecondaryColor3usvEXT((*C.GLushort)(arg0));
}

//extern void glSecondaryColorPointerEXT (GLint, GLenum, GLsizei, const GLvoid *)
func SecondaryColorPointerEXT(arg0 GLint, arg1 GLenum, arg2 GLsizei, arg3 unsafe.Pointer) 
{
	C.glSecondaryColorPointerEXT(C.GLint(arg0), C.GLenum(arg1), C.GLsizei(arg2), arg3);
}

//extern void glTextureNormalEXT (GLenum)
func TextureNormalEXT(arg0 GLenum) 
{
	C.glTextureNormalEXT(C.GLenum(arg0));
}

//extern void glMultiDrawArraysEXT (GLenum, GLint *, GLsizei *, GLsizei)
func MultiDrawArraysEXT(arg0 GLenum, arg1 *GLint, arg2 *GLsizei, arg3 GLsizei) 
{
	C.glMultiDrawArraysEXT(C.GLenum(arg0), (*C.GLint)(arg1), (*C.GLsizei)(arg2), C.GLsizei(arg3));
}

//extern void glMultiDrawElementsEXT (GLenum, const GLsizei *, GLenum, const GLvoid* *, GLsizei)
func MultiDrawElementsEXT(arg0 GLenum, arg1 *GLsizei, arg2 GLenum, arg3 *unsafe.Pointer, arg4 GLsizei) 
{
	C.glMultiDrawElementsEXT(C.GLenum(arg0), (*C.GLsizei)(arg1), C.GLenum(arg2), arg3, C.GLsizei(arg4));
}

//extern void glFogCoordfEXT (GLfloat)
func FogCoordfEXT(arg0 GLfloat) 
{
	C.glFogCoordfEXT(C.GLfloat(arg0));
}

//extern void glFogCoordfvEXT (const GLfloat *)
func FogCoordfvEXT(arg0 *GLfloat) 
{
	C.glFogCoordfvEXT((*C.GLfloat)(arg0));
}

//extern void glFogCoorddEXT (GLdouble)
func FogCoorddEXT(arg0 GLdouble) 
{
	C.glFogCoorddEXT(C.GLdouble(arg0));
}

//extern void glFogCoorddvEXT (const GLdouble *)
func FogCoorddvEXT(arg0 *GLdouble) 
{
	C.glFogCoorddvEXT((*C.GLdouble)(arg0));
}

//extern void glFogCoordPointerEXT (GLenum, GLsizei, const GLvoid *)
func FogCoordPointerEXT(arg0 GLenum, arg1 GLsizei, arg2 unsafe.Pointer) 
{
	C.glFogCoordPointerEXT(C.GLenum(arg0), C.GLsizei(arg1), arg2);
}

//extern void glTangent3bEXT (GLbyte, GLbyte, GLbyte)
func Tangent3bEXT(arg0 GLbyte, arg1 GLbyte, arg2 GLbyte) 
{
	C.glTangent3bEXT(C.GLbyte(arg0), C.GLbyte(arg1), C.GLbyte(arg2));
}

//extern void glTangent3bvEXT (const GLbyte *)
func Tangent3bvEXT(arg0 *GLbyte) 
{
	C.glTangent3bvEXT((*C.GLbyte)(arg0));
}

//extern void glTangent3dEXT (GLdouble, GLdouble, GLdouble)
func Tangent3dEXT(arg0 GLdouble, arg1 GLdouble, arg2 GLdouble) 
{
	C.glTangent3dEXT(C.GLdouble(arg0), C.GLdouble(arg1), C.GLdouble(arg2));
}

//extern void glTangent3dvEXT (const GLdouble *)
func Tangent3dvEXT(arg0 *GLdouble) 
{
	C.glTangent3dvEXT((*C.GLdouble)(arg0));
}

//extern void glTangent3fEXT (GLfloat, GLfloat, GLfloat)
func Tangent3fEXT(arg0 GLfloat, arg1 GLfloat, arg2 GLfloat) 
{
	C.glTangent3fEXT(C.GLfloat(arg0), C.GLfloat(arg1), C.GLfloat(arg2));
}

//extern void glTangent3fvEXT (const GLfloat *)
func Tangent3fvEXT(arg0 *GLfloat) 
{
	C.glTangent3fvEXT((*C.GLfloat)(arg0));
}

//extern void glTangent3iEXT (GLint, GLint, GLint)
func Tangent3iEXT(arg0 GLint, arg1 GLint, arg2 GLint) 
{
	C.glTangent3iEXT(C.GLint(arg0), C.GLint(arg1), C.GLint(arg2));
}

//extern void glTangent3ivEXT (const GLint *)
func Tangent3ivEXT(arg0 *GLint) 
{
	C.glTangent3ivEXT((*C.GLint)(arg0));
}

//extern void glTangent3sEXT (GLshort, GLshort, GLshort)
func Tangent3sEXT(arg0 GLshort, arg1 GLshort, arg2 GLshort) 
{
	C.glTangent3sEXT(C.GLshort(arg0), C.GLshort(arg1), C.GLshort(arg2));
}

//extern void glTangent3svEXT (const GLshort *)
func Tangent3svEXT(arg0 *GLshort) 
{
	C.glTangent3svEXT((*C.GLshort)(arg0));
}

//extern void glBinormal3bEXT (GLbyte, GLbyte, GLbyte)
func Binormal3bEXT(arg0 GLbyte, arg1 GLbyte, arg2 GLbyte) 
{
	C.glBinormal3bEXT(C.GLbyte(arg0), C.GLbyte(arg1), C.GLbyte(arg2));
}

//extern void glBinormal3bvEXT (const GLbyte *)
func Binormal3bvEXT(arg0 *GLbyte) 
{
	C.glBinormal3bvEXT((*C.GLbyte)(arg0));
}

//extern void glBinormal3dEXT (GLdouble, GLdouble, GLdouble)
func Binormal3dEXT(arg0 GLdouble, arg1 GLdouble, arg2 GLdouble) 
{
	C.glBinormal3dEXT(C.GLdouble(arg0), C.GLdouble(arg1), C.GLdouble(arg2));
}

//extern void glBinormal3dvEXT (const GLdouble *)
func Binormal3dvEXT(arg0 *GLdouble) 
{
	C.glBinormal3dvEXT((*C.GLdouble)(arg0));
}

//extern void glBinormal3fEXT (GLfloat, GLfloat, GLfloat)
func Binormal3fEXT(arg0 GLfloat, arg1 GLfloat, arg2 GLfloat) 
{
	C.glBinormal3fEXT(C.GLfloat(arg0), C.GLfloat(arg1), C.GLfloat(arg2));
}

//extern void glBinormal3fvEXT (const GLfloat *)
func Binormal3fvEXT(arg0 *GLfloat) 
{
	C.glBinormal3fvEXT((*C.GLfloat)(arg0));
}

//extern void glBinormal3iEXT (GLint, GLint, GLint)
func Binormal3iEXT(arg0 GLint, arg1 GLint, arg2 GLint) 
{
	C.glBinormal3iEXT(C.GLint(arg0), C.GLint(arg1), C.GLint(arg2));
}

//extern void glBinormal3ivEXT (const GLint *)
func Binormal3ivEXT(arg0 *GLint) 
{
	C.glBinormal3ivEXT((*C.GLint)(arg0));
}

//extern void glBinormal3sEXT (GLshort, GLshort, GLshort)
func Binormal3sEXT(arg0 GLshort, arg1 GLshort, arg2 GLshort) 
{
	C.glBinormal3sEXT(C.GLshort(arg0), C.GLshort(arg1), C.GLshort(arg2));
}

//extern void glBinormal3svEXT (const GLshort *)
func Binormal3svEXT(arg0 *GLshort) 
{
	C.glBinormal3svEXT((*C.GLshort)(arg0));
}

//extern void glTangentPointerEXT (GLenum, GLsizei, const GLvoid *)
func TangentPointerEXT(arg0 GLenum, arg1 GLsizei, arg2 unsafe.Pointer) 
{
	C.glTangentPointerEXT(C.GLenum(arg0), C.GLsizei(arg1), arg2);
}

//extern void glBinormalPointerEXT (GLenum, GLsizei, const GLvoid *)
func BinormalPointerEXT(arg0 GLenum, arg1 GLsizei, arg2 unsafe.Pointer) 
{
	C.glBinormalPointerEXT(C.GLenum(arg0), C.GLsizei(arg1), arg2);
}

//extern void glFinishTextureSUNX (void)
func FinishTextureSUNX() 
{
	C.glFinishTextureSUNX();
}

//extern void glGlobalAlphaFactorbSUN (GLbyte)
func GlobalAlphaFactorbSUN(arg0 GLbyte) 
{
	C.glGlobalAlphaFactorbSUN(C.GLbyte(arg0));
}

//extern void glGlobalAlphaFactorsSUN (GLshort)
func GlobalAlphaFactorsSUN(arg0 GLshort) 
{
	C.glGlobalAlphaFactorsSUN(C.GLshort(arg0));
}

//extern void glGlobalAlphaFactoriSUN (GLint)
func GlobalAlphaFactoriSUN(arg0 GLint) 
{
	C.glGlobalAlphaFactoriSUN(C.GLint(arg0));
}

//extern void glGlobalAlphaFactorfSUN (GLfloat)
func GlobalAlphaFactorfSUN(arg0 GLfloat) 
{
	C.glGlobalAlphaFactorfSUN(C.GLfloat(arg0));
}

//extern void glGlobalAlphaFactordSUN (GLdouble)
func GlobalAlphaFactordSUN(arg0 GLdouble) 
{
	C.glGlobalAlphaFactordSUN(C.GLdouble(arg0));
}

//extern void glGlobalAlphaFactorubSUN (GLubyte)
func GlobalAlphaFactorubSUN(arg0 GLubyte) 
{
	C.glGlobalAlphaFactorubSUN(C.GLubyte(arg0));
}

//extern void glGlobalAlphaFactorusSUN (GLushort)
func GlobalAlphaFactorusSUN(arg0 GLushort) 
{
	C.glGlobalAlphaFactorusSUN(C.GLushort(arg0));
}

//extern void glGlobalAlphaFactoruiSUN (GLuint)
func GlobalAlphaFactoruiSUN(arg0 GLuint) 
{
	C.glGlobalAlphaFactoruiSUN(C.GLuint(arg0));
}

//extern void glReplacementCodeuiSUN (GLuint)
func ReplacementCodeuiSUN(arg0 GLuint) 
{
	C.glReplacementCodeuiSUN(C.GLuint(arg0));
}

//extern void glReplacementCodeusSUN (GLushort)
func ReplacementCodeusSUN(arg0 GLushort) 
{
	C.glReplacementCodeusSUN(C.GLushort(arg0));
}

//extern void glReplacementCodeubSUN (GLubyte)
func ReplacementCodeubSUN(arg0 GLubyte) 
{
	C.glReplacementCodeubSUN(C.GLubyte(arg0));
}

//extern void glReplacementCodeuivSUN (const GLuint *)
func ReplacementCodeuivSUN(arg0 *GLuint) 
{
	C.glReplacementCodeuivSUN((*C.GLuint)(arg0));
}

//extern void glReplacementCodeusvSUN (const GLushort *)
func ReplacementCodeusvSUN(arg0 *GLushort) 
{
	C.glReplacementCodeusvSUN((*C.GLushort)(arg0));
}

//extern void glReplacementCodeubvSUN (const GLubyte *)
func ReplacementCodeubvSUN(arg0 *GLubyte) 
{
	C.glReplacementCodeubvSUN((*C.GLubyte)(arg0));
}

//extern void glReplacementCodePointerSUN (GLenum, GLsizei, const GLvoid* *)
func ReplacementCodePointerSUN(arg0 GLenum, arg1 GLsizei, arg2 *unsafe.Pointer) 
{
	C.glReplacementCodePointerSUN(C.GLenum(arg0), C.GLsizei(arg1), arg2);
}

//extern void glColor4ubVertex2fSUN (GLubyte, GLubyte, GLubyte, GLubyte, GLfloat, GLfloat)
func Color4ubVertex2fSUN(arg0 GLubyte, arg1 GLubyte, arg2 GLubyte, arg3 GLubyte, arg4 GLfloat, arg5 GLfloat) 
{
	C.glColor4ubVertex2fSUN(C.GLubyte(arg0), C.GLubyte(arg1), C.GLubyte(arg2), C.GLubyte(arg3), C.GLfloat(arg4), C.GLfloat(arg5));
}

//extern void glColor4ubVertex2fvSUN (const GLubyte *, const GLfloat *)
func Color4ubVertex2fvSUN(arg0 *GLubyte, arg1 *GLfloat) 
{
	C.glColor4ubVertex2fvSUN((*C.GLubyte)(arg0), (*C.GLfloat)(arg1));
}

//extern void glColor4ubVertex3fSUN (GLubyte, GLubyte, GLubyte, GLubyte, GLfloat, GLfloat, GLfloat)
func Color4ubVertex3fSUN(arg0 GLubyte, arg1 GLubyte, arg2 GLubyte, arg3 GLubyte, arg4 GLfloat, arg5 GLfloat, arg6 GLfloat) 
{
	C.glColor4ubVertex3fSUN(C.GLubyte(arg0), C.GLubyte(arg1), C.GLubyte(arg2), C.GLubyte(arg3), C.GLfloat(arg4), C.GLfloat(arg5), C.GLfloat(arg6));
}

//extern void glColor4ubVertex3fvSUN (const GLubyte *, const GLfloat *)
func Color4ubVertex3fvSUN(arg0 *GLubyte, arg1 *GLfloat) 
{
	C.glColor4ubVertex3fvSUN((*C.GLubyte)(arg0), (*C.GLfloat)(arg1));
}

//extern void glColor3fVertex3fSUN (GLfloat, GLfloat, GLfloat, GLfloat, GLfloat, GLfloat)
func Color3fVertex3fSUN(arg0 GLfloat, arg1 GLfloat, arg2 GLfloat, arg3 GLfloat, arg4 GLfloat, arg5 GLfloat) 
{
	C.glColor3fVertex3fSUN(C.GLfloat(arg0), C.GLfloat(arg1), C.GLfloat(arg2), C.GLfloat(arg3), C.GLfloat(arg4), C.GLfloat(arg5));
}

//extern void glColor3fVertex3fvSUN (const GLfloat *, const GLfloat *)
func Color3fVertex3fvSUN(arg0 *GLfloat, arg1 *GLfloat) 
{
	C.glColor3fVertex3fvSUN((*C.GLfloat)(arg0), (*C.GLfloat)(arg1));
}

//extern void glNormal3fVertex3fSUN (GLfloat, GLfloat, GLfloat, GLfloat, GLfloat, GLfloat)
func Normal3fVertex3fSUN(arg0 GLfloat, arg1 GLfloat, arg2 GLfloat, arg3 GLfloat, arg4 GLfloat, arg5 GLfloat) 
{
	C.glNormal3fVertex3fSUN(C.GLfloat(arg0), C.GLfloat(arg1), C.GLfloat(arg2), C.GLfloat(arg3), C.GLfloat(arg4), C.GLfloat(arg5));
}

//extern void glNormal3fVertex3fvSUN (const GLfloat *, const GLfloat *)
func Normal3fVertex3fvSUN(arg0 *GLfloat, arg1 *GLfloat) 
{
	C.glNormal3fVertex3fvSUN((*C.GLfloat)(arg0), (*C.GLfloat)(arg1));
}

//extern void glColor4fNormal3fVertex3fSUN (GLfloat, GLfloat, GLfloat, GLfloat, GLfloat, GLfloat, GLfloat, GLfloat, GLfloat, GLfloat)
func Color4fNormal3fVertex3fSUN(arg0 GLfloat, arg1 GLfloat, arg2 GLfloat, arg3 GLfloat, arg4 GLfloat, arg5 GLfloat, arg6 GLfloat, arg7 GLfloat, arg8 GLfloat, arg9 GLfloat) 
{
	C.glColor4fNormal3fVertex3fSUN(C.GLfloat(arg0), C.GLfloat(arg1), C.GLfloat(arg2), C.GLfloat(arg3), C.GLfloat(arg4), C.GLfloat(arg5), C.GLfloat(arg6), C.GLfloat(arg7), C.GLfloat(arg8), C.GLfloat(arg9));
}

//extern void glColor4fNormal3fVertex3fvSUN (const GLfloat *, const GLfloat *, const GLfloat *)
func Color4fNormal3fVertex3fvSUN(arg0 *GLfloat, arg1 *GLfloat, arg2 *GLfloat) 
{
	C.glColor4fNormal3fVertex3fvSUN((*C.GLfloat)(arg0), (*C.GLfloat)(arg1), (*C.GLfloat)(arg2));
}

//extern void glTexCoord2fVertex3fSUN (GLfloat, GLfloat, GLfloat, GLfloat, GLfloat)
func TexCoord2fVertex3fSUN(arg0 GLfloat, arg1 GLfloat, arg2 GLfloat, arg3 GLfloat, arg4 GLfloat) 
{
	C.glTexCoord2fVertex3fSUN(C.GLfloat(arg0), C.GLfloat(arg1), C.GLfloat(arg2), C.GLfloat(arg3), C.GLfloat(arg4));
}

//extern void glTexCoord2fVertex3fvSUN (const GLfloat *, const GLfloat *)
func TexCoord2fVertex3fvSUN(arg0 *GLfloat, arg1 *GLfloat) 
{
	C.glTexCoord2fVertex3fvSUN((*C.GLfloat)(arg0), (*C.GLfloat)(arg1));
}

//extern void glTexCoord4fVertex4fSUN (GLfloat, GLfloat, GLfloat, GLfloat, GLfloat, GLfloat, GLfloat, GLfloat)
func TexCoord4fVertex4fSUN(arg0 GLfloat, arg1 GLfloat, arg2 GLfloat, arg3 GLfloat, arg4 GLfloat, arg5 GLfloat, arg6 GLfloat, arg7 GLfloat) 
{
	C.glTexCoord4fVertex4fSUN(C.GLfloat(arg0), C.GLfloat(arg1), C.GLfloat(arg2), C.GLfloat(arg3), C.GLfloat(arg4), C.GLfloat(arg5), C.GLfloat(arg6), C.GLfloat(arg7));
}

//extern void glTexCoord4fVertex4fvSUN (const GLfloat *, const GLfloat *)
func TexCoord4fVertex4fvSUN(arg0 *GLfloat, arg1 *GLfloat) 
{
	C.glTexCoord4fVertex4fvSUN((*C.GLfloat)(arg0), (*C.GLfloat)(arg1));
}

//extern void glTexCoord2fColor4ubVertex3fSUN (GLfloat, GLfloat, GLubyte, GLubyte, GLubyte, GLubyte, GLfloat, GLfloat, GLfloat)
func TexCoord2fColor4ubVertex3fSUN(arg0 GLfloat, arg1 GLfloat, arg2 GLubyte, arg3 GLubyte, arg4 GLubyte, arg5 GLubyte, arg6 GLfloat, arg7 GLfloat, arg8 GLfloat) 
{
	C.glTexCoord2fColor4ubVertex3fSUN(C.GLfloat(arg0), C.GLfloat(arg1), C.GLubyte(arg2), C.GLubyte(arg3), C.GLubyte(arg4), C.GLubyte(arg5), C.GLfloat(arg6), C.GLfloat(arg7), C.GLfloat(arg8));
}

//extern void glTexCoord2fColor4ubVertex3fvSUN (const GLfloat *, const GLubyte *, const GLfloat *)
func TexCoord2fColor4ubVertex3fvSUN(arg0 *GLfloat, arg1 *GLubyte, arg2 *GLfloat) 
{
	C.glTexCoord2fColor4ubVertex3fvSUN((*C.GLfloat)(arg0), (*C.GLubyte)(arg1), (*C.GLfloat)(arg2));
}

//extern void glTexCoord2fColor3fVertex3fSUN (GLfloat, GLfloat, GLfloat, GLfloat, GLfloat, GLfloat, GLfloat, GLfloat)
func TexCoord2fColor3fVertex3fSUN(arg0 GLfloat, arg1 GLfloat, arg2 GLfloat, arg3 GLfloat, arg4 GLfloat, arg5 GLfloat, arg6 GLfloat, arg7 GLfloat) 
{
	C.glTexCoord2fColor3fVertex3fSUN(C.GLfloat(arg0), C.GLfloat(arg1), C.GLfloat(arg2), C.GLfloat(arg3), C.GLfloat(arg4), C.GLfloat(arg5), C.GLfloat(arg6), C.GLfloat(arg7));
}

//extern void glTexCoord2fColor3fVertex3fvSUN (const GLfloat *, const GLfloat *, const GLfloat *)
func TexCoord2fColor3fVertex3fvSUN(arg0 *GLfloat, arg1 *GLfloat, arg2 *GLfloat) 
{
	C.glTexCoord2fColor3fVertex3fvSUN((*C.GLfloat)(arg0), (*C.GLfloat)(arg1), (*C.GLfloat)(arg2));
}

//extern void glTexCoord2fNormal3fVertex3fSUN (GLfloat, GLfloat, GLfloat, GLfloat, GLfloat, GLfloat, GLfloat, GLfloat)
func TexCoord2fNormal3fVertex3fSUN(arg0 GLfloat, arg1 GLfloat, arg2 GLfloat, arg3 GLfloat, arg4 GLfloat, arg5 GLfloat, arg6 GLfloat, arg7 GLfloat) 
{
	C.glTexCoord2fNormal3fVertex3fSUN(C.GLfloat(arg0), C.GLfloat(arg1), C.GLfloat(arg2), C.GLfloat(arg3), C.GLfloat(arg4), C.GLfloat(arg5), C.GLfloat(arg6), C.GLfloat(arg7));
}

//extern void glTexCoord2fNormal3fVertex3fvSUN (const GLfloat *, const GLfloat *, const GLfloat *)
func TexCoord2fNormal3fVertex3fvSUN(arg0 *GLfloat, arg1 *GLfloat, arg2 *GLfloat) 
{
	C.glTexCoord2fNormal3fVertex3fvSUN((*C.GLfloat)(arg0), (*C.GLfloat)(arg1), (*C.GLfloat)(arg2));
}

//extern void glTexCoord2fColor4fNormal3fVertex3fSUN (GLfloat, GLfloat, GLfloat, GLfloat, GLfloat, GLfloat, GLfloat, GLfloat, GLfloat, GLfloat, GLfloat, GLfloat)
func TexCoord2fColor4fNormal3fVertex3fSUN(arg0 GLfloat, arg1 GLfloat, arg2 GLfloat, arg3 GLfloat, arg4 GLfloat, arg5 GLfloat, arg6 GLfloat, arg7 GLfloat, arg8 GLfloat, arg9 GLfloat, arg10 GLfloat, arg11 GLfloat) 
{
	C.glTexCoord2fColor4fNormal3fVertex3fSUN(C.GLfloat(arg0), C.GLfloat(arg1), C.GLfloat(arg2), C.GLfloat(arg3), C.GLfloat(arg4), C.GLfloat(arg5), C.GLfloat(arg6), C.GLfloat(arg7), C.GLfloat(arg8), C.GLfloat(arg9), C.GLfloat(arg10), C.GLfloat(arg11));
}

//extern void glTexCoord2fColor4fNormal3fVertex3fvSUN (const GLfloat *, const GLfloat *, const GLfloat *, const GLfloat *)
func TexCoord2fColor4fNormal3fVertex3fvSUN(arg0 *GLfloat, arg1 *GLfloat, arg2 *GLfloat, arg3 *GLfloat) 
{
	C.glTexCoord2fColor4fNormal3fVertex3fvSUN((*C.GLfloat)(arg0), (*C.GLfloat)(arg1), (*C.GLfloat)(arg2), (*C.GLfloat)(arg3));
}

//extern void glTexCoord4fColor4fNormal3fVertex4fSUN (GLfloat, GLfloat, GLfloat, GLfloat, GLfloat, GLfloat, GLfloat, GLfloat, GLfloat, GLfloat, GLfloat, GLfloat, GLfloat, GLfloat, GLfloat)
func TexCoord4fColor4fNormal3fVertex4fSUN(arg0 GLfloat, arg1 GLfloat, arg2 GLfloat, arg3 GLfloat, arg4 GLfloat, arg5 GLfloat, arg6 GLfloat, arg7 GLfloat, arg8 GLfloat, arg9 GLfloat, arg10 GLfloat, arg11 GLfloat, arg12 GLfloat, arg13 GLfloat, arg14 GLfloat) 
{
	C.glTexCoord4fColor4fNormal3fVertex4fSUN(C.GLfloat(arg0), C.GLfloat(arg1), C.GLfloat(arg2), C.GLfloat(arg3), C.GLfloat(arg4), C.GLfloat(arg5), C.GLfloat(arg6), C.GLfloat(arg7), C.GLfloat(arg8), C.GLfloat(arg9), C.GLfloat(arg10), C.GLfloat(arg11), C.GLfloat(arg12), C.GLfloat(arg13), C.GLfloat(arg14));
}

//extern void glTexCoord4fColor4fNormal3fVertex4fvSUN (const GLfloat *, const GLfloat *, const GLfloat *, const GLfloat *)
func TexCoord4fColor4fNormal3fVertex4fvSUN(arg0 *GLfloat, arg1 *GLfloat, arg2 *GLfloat, arg3 *GLfloat) 
{
	C.glTexCoord4fColor4fNormal3fVertex4fvSUN((*C.GLfloat)(arg0), (*C.GLfloat)(arg1), (*C.GLfloat)(arg2), (*C.GLfloat)(arg3));
}

//extern void glReplacementCodeuiVertex3fSUN (GLuint, GLfloat, GLfloat, GLfloat)
func ReplacementCodeuiVertex3fSUN(arg0 GLuint, arg1 GLfloat, arg2 GLfloat, arg3 GLfloat) 
{
	C.glReplacementCodeuiVertex3fSUN(C.GLuint(arg0), C.GLfloat(arg1), C.GLfloat(arg2), C.GLfloat(arg3));
}

//extern void glReplacementCodeuiVertex3fvSUN (const GLuint *, const GLfloat *)
func ReplacementCodeuiVertex3fvSUN(arg0 *GLuint, arg1 *GLfloat) 
{
	C.glReplacementCodeuiVertex3fvSUN((*C.GLuint)(arg0), (*C.GLfloat)(arg1));
}

//extern void glReplacementCodeuiColor4ubVertex3fSUN (GLuint, GLubyte, GLubyte, GLubyte, GLubyte, GLfloat, GLfloat, GLfloat)
func ReplacementCodeuiColor4ubVertex3fSUN(arg0 GLuint, arg1 GLubyte, arg2 GLubyte, arg3 GLubyte, arg4 GLubyte, arg5 GLfloat, arg6 GLfloat, arg7 GLfloat) 
{
	C.glReplacementCodeuiColor4ubVertex3fSUN(C.GLuint(arg0), C.GLubyte(arg1), C.GLubyte(arg2), C.GLubyte(arg3), C.GLubyte(arg4), C.GLfloat(arg5), C.GLfloat(arg6), C.GLfloat(arg7));
}

//extern void glReplacementCodeuiColor4ubVertex3fvSUN (const GLuint *, const GLubyte *, const GLfloat *)
func ReplacementCodeuiColor4ubVertex3fvSUN(arg0 *GLuint, arg1 *GLubyte, arg2 *GLfloat) 
{
	C.glReplacementCodeuiColor4ubVertex3fvSUN((*C.GLuint)(arg0), (*C.GLubyte)(arg1), (*C.GLfloat)(arg2));
}

//extern void glReplacementCodeuiColor3fVertex3fSUN (GLuint, GLfloat, GLfloat, GLfloat, GLfloat, GLfloat, GLfloat)
func ReplacementCodeuiColor3fVertex3fSUN(arg0 GLuint, arg1 GLfloat, arg2 GLfloat, arg3 GLfloat, arg4 GLfloat, arg5 GLfloat, arg6 GLfloat) 
{
	C.glReplacementCodeuiColor3fVertex3fSUN(C.GLuint(arg0), C.GLfloat(arg1), C.GLfloat(arg2), C.GLfloat(arg3), C.GLfloat(arg4), C.GLfloat(arg5), C.GLfloat(arg6));
}

//extern void glReplacementCodeuiColor3fVertex3fvSUN (const GLuint *, const GLfloat *, const GLfloat *)
func ReplacementCodeuiColor3fVertex3fvSUN(arg0 *GLuint, arg1 *GLfloat, arg2 *GLfloat) 
{
	C.glReplacementCodeuiColor3fVertex3fvSUN((*C.GLuint)(arg0), (*C.GLfloat)(arg1), (*C.GLfloat)(arg2));
}

//extern void glReplacementCodeuiNormal3fVertex3fSUN (GLuint, GLfloat, GLfloat, GLfloat, GLfloat, GLfloat, GLfloat)
func ReplacementCodeuiNormal3fVertex3fSUN(arg0 GLuint, arg1 GLfloat, arg2 GLfloat, arg3 GLfloat, arg4 GLfloat, arg5 GLfloat, arg6 GLfloat) 
{
	C.glReplacementCodeuiNormal3fVertex3fSUN(C.GLuint(arg0), C.GLfloat(arg1), C.GLfloat(arg2), C.GLfloat(arg3), C.GLfloat(arg4), C.GLfloat(arg5), C.GLfloat(arg6));
}

//extern void glReplacementCodeuiNormal3fVertex3fvSUN (const GLuint *, const GLfloat *, const GLfloat *)
func ReplacementCodeuiNormal3fVertex3fvSUN(arg0 *GLuint, arg1 *GLfloat, arg2 *GLfloat) 
{
	C.glReplacementCodeuiNormal3fVertex3fvSUN((*C.GLuint)(arg0), (*C.GLfloat)(arg1), (*C.GLfloat)(arg2));
}

//extern void glReplacementCodeuiColor4fNormal3fVertex3fSUN (GLuint, GLfloat, GLfloat, GLfloat, GLfloat, GLfloat, GLfloat, GLfloat, GLfloat, GLfloat, GLfloat)
func ReplacementCodeuiColor4fNormal3fVertex3fSUN(arg0 GLuint, arg1 GLfloat, arg2 GLfloat, arg3 GLfloat, arg4 GLfloat, arg5 GLfloat, arg6 GLfloat, arg7 GLfloat, arg8 GLfloat, arg9 GLfloat, arg10 GLfloat) 
{
	C.glReplacementCodeuiColor4fNormal3fVertex3fSUN(C.GLuint(arg0), C.GLfloat(arg1), C.GLfloat(arg2), C.GLfloat(arg3), C.GLfloat(arg4), C.GLfloat(arg5), C.GLfloat(arg6), C.GLfloat(arg7), C.GLfloat(arg8), C.GLfloat(arg9), C.GLfloat(arg10));
}

//extern void glReplacementCodeuiColor4fNormal3fVertex3fvSUN (const GLuint *, const GLfloat *, const GLfloat *, const GLfloat *)
func ReplacementCodeuiColor4fNormal3fVertex3fvSUN(arg0 *GLuint, arg1 *GLfloat, arg2 *GLfloat, arg3 *GLfloat) 
{
	C.glReplacementCodeuiColor4fNormal3fVertex3fvSUN((*C.GLuint)(arg0), (*C.GLfloat)(arg1), (*C.GLfloat)(arg2), (*C.GLfloat)(arg3));
}

//extern void glReplacementCodeuiTexCoord2fVertex3fSUN (GLuint, GLfloat, GLfloat, GLfloat, GLfloat, GLfloat)
func ReplacementCodeuiTexCoord2fVertex3fSUN(arg0 GLuint, arg1 GLfloat, arg2 GLfloat, arg3 GLfloat, arg4 GLfloat, arg5 GLfloat) 
{
	C.glReplacementCodeuiTexCoord2fVertex3fSUN(C.GLuint(arg0), C.GLfloat(arg1), C.GLfloat(arg2), C.GLfloat(arg3), C.GLfloat(arg4), C.GLfloat(arg5));
}

//extern void glReplacementCodeuiTexCoord2fVertex3fvSUN (const GLuint *, const GLfloat *, const GLfloat *)
func ReplacementCodeuiTexCoord2fVertex3fvSUN(arg0 *GLuint, arg1 *GLfloat, arg2 *GLfloat) 
{
	C.glReplacementCodeuiTexCoord2fVertex3fvSUN((*C.GLuint)(arg0), (*C.GLfloat)(arg1), (*C.GLfloat)(arg2));
}

//extern void glReplacementCodeuiTexCoord2fNormal3fVertex3fSUN (GLuint, GLfloat, GLfloat, GLfloat, GLfloat, GLfloat, GLfloat, GLfloat, GLfloat)
func ReplacementCodeuiTexCoord2fNormal3fVertex3fSUN(arg0 GLuint, arg1 GLfloat, arg2 GLfloat, arg3 GLfloat, arg4 GLfloat, arg5 GLfloat, arg6 GLfloat, arg7 GLfloat, arg8 GLfloat) 
{
	C.glReplacementCodeuiTexCoord2fNormal3fVertex3fSUN(C.GLuint(arg0), C.GLfloat(arg1), C.GLfloat(arg2), C.GLfloat(arg3), C.GLfloat(arg4), C.GLfloat(arg5), C.GLfloat(arg6), C.GLfloat(arg7), C.GLfloat(arg8));
}

//extern void glReplacementCodeuiTexCoord2fNormal3fVertex3fvSUN (const GLuint *, const GLfloat *, const GLfloat *, const GLfloat *)
func ReplacementCodeuiTexCoord2fNormal3fVertex3fvSUN(arg0 *GLuint, arg1 *GLfloat, arg2 *GLfloat, arg3 *GLfloat) 
{
	C.glReplacementCodeuiTexCoord2fNormal3fVertex3fvSUN((*C.GLuint)(arg0), (*C.GLfloat)(arg1), (*C.GLfloat)(arg2), (*C.GLfloat)(arg3));
}

//extern void glReplacementCodeuiTexCoord2fColor4fNormal3fVertex3fSUN (GLuint, GLfloat, GLfloat, GLfloat, GLfloat, GLfloat, GLfloat, GLfloat, GLfloat, GLfloat, GLfloat, GLfloat, GLfloat)
func ReplacementCodeuiTexCoord2fColor4fNormal3fVertex3fSUN(arg0 GLuint, arg1 GLfloat, arg2 GLfloat, arg3 GLfloat, arg4 GLfloat, arg5 GLfloat, arg6 GLfloat, arg7 GLfloat, arg8 GLfloat, arg9 GLfloat, arg10 GLfloat, arg11 GLfloat, arg12 GLfloat) 
{
	C.glReplacementCodeuiTexCoord2fColor4fNormal3fVertex3fSUN(C.GLuint(arg0), C.GLfloat(arg1), C.GLfloat(arg2), C.GLfloat(arg3), C.GLfloat(arg4), C.GLfloat(arg5), C.GLfloat(arg6), C.GLfloat(arg7), C.GLfloat(arg8), C.GLfloat(arg9), C.GLfloat(arg10), C.GLfloat(arg11), C.GLfloat(arg12));
}

//extern void glReplacementCodeuiTexCoord2fColor4fNormal3fVertex3fvSUN (const GLuint *, const GLfloat *, const GLfloat *, const GLfloat *, const GLfloat *)
func ReplacementCodeuiTexCoord2fColor4fNormal3fVertex3fvSUN(arg0 *GLuint, arg1 *GLfloat, arg2 *GLfloat, arg3 *GLfloat, arg4 *GLfloat) 
{
	C.glReplacementCodeuiTexCoord2fColor4fNormal3fVertex3fvSUN((*C.GLuint)(arg0), (*C.GLfloat)(arg1), (*C.GLfloat)(arg2), (*C.GLfloat)(arg3), (*C.GLfloat)(arg4));
}

//extern void glBlendFuncSeparateEXT (GLenum, GLenum, GLenum, GLenum)
func BlendFuncSeparateEXT(arg0 GLenum, arg1 GLenum, arg2 GLenum, arg3 GLenum) 
{
	C.glBlendFuncSeparateEXT(C.GLenum(arg0), C.GLenum(arg1), C.GLenum(arg2), C.GLenum(arg3));
}

//extern void glBlendFuncSeparateINGR (GLenum, GLenum, GLenum, GLenum)
func BlendFuncSeparateINGR(arg0 GLenum, arg1 GLenum, arg2 GLenum, arg3 GLenum) 
{
	C.glBlendFuncSeparateINGR(C.GLenum(arg0), C.GLenum(arg1), C.GLenum(arg2), C.GLenum(arg3));
}

//extern void glVertexWeightfEXT (GLfloat)
func VertexWeightfEXT(arg0 GLfloat) 
{
	C.glVertexWeightfEXT(C.GLfloat(arg0));
}

//extern void glVertexWeightfvEXT (const GLfloat *)
func VertexWeightfvEXT(arg0 *GLfloat) 
{
	C.glVertexWeightfvEXT((*C.GLfloat)(arg0));
}

//extern void glVertexWeightPointerEXT (GLsizei, GLenum, GLsizei, const GLvoid *)
func VertexWeightPointerEXT(arg0 GLsizei, arg1 GLenum, arg2 GLsizei, arg3 unsafe.Pointer) 
{
	C.glVertexWeightPointerEXT(C.GLsizei(arg0), C.GLenum(arg1), C.GLsizei(arg2), arg3);
}

//extern void glFlushVertexArrayRangeNV (void)
func FlushVertexArrayRangeNV() 
{
	C.glFlushVertexArrayRangeNV();
}

//extern void glVertexArrayRangeNV (GLsizei, const GLvoid *)
func VertexArrayRangeNV(arg0 GLsizei, arg1 unsafe.Pointer) 
{
	C.glVertexArrayRangeNV(C.GLsizei(arg0), arg1);
}

//extern void glCombinerParameterfvNV (GLenum, const GLfloat *)
func CombinerParameterfvNV(arg0 GLenum, arg1 *GLfloat) 
{
	C.glCombinerParameterfvNV(C.GLenum(arg0), (*C.GLfloat)(arg1));
}

//extern void glCombinerParameterfNV (GLenum, GLfloat)
func CombinerParameterfNV(arg0 GLenum, arg1 GLfloat) 
{
	C.glCombinerParameterfNV(C.GLenum(arg0), C.GLfloat(arg1));
}

//extern void glCombinerParameterivNV (GLenum, const GLint *)
func CombinerParameterivNV(arg0 GLenum, arg1 *GLint) 
{
	C.glCombinerParameterivNV(C.GLenum(arg0), (*C.GLint)(arg1));
}

//extern void glCombinerParameteriNV (GLenum, GLint)
func CombinerParameteriNV(arg0 GLenum, arg1 GLint) 
{
	C.glCombinerParameteriNV(C.GLenum(arg0), C.GLint(arg1));
}

//extern void glCombinerInputNV (GLenum, GLenum, GLenum, GLenum, GLenum, GLenum)
func CombinerInputNV(arg0 GLenum, arg1 GLenum, arg2 GLenum, arg3 GLenum, arg4 GLenum, arg5 GLenum) 
{
	C.glCombinerInputNV(C.GLenum(arg0), C.GLenum(arg1), C.GLenum(arg2), C.GLenum(arg3), C.GLenum(arg4), C.GLenum(arg5));
}

//extern void glCombinerOutputNV (GLenum, GLenum, GLenum, GLenum, GLenum, GLenum, GLenum, GLboolean, GLboolean, GLboolean)
func CombinerOutputNV(arg0 GLenum, arg1 GLenum, arg2 GLenum, arg3 GLenum, arg4 GLenum, arg5 GLenum, arg6 GLenum, arg7 GLboolean, arg8 GLboolean, arg9 GLboolean) 
{
	C.glCombinerOutputNV(C.GLenum(arg0), C.GLenum(arg1), C.GLenum(arg2), C.GLenum(arg3), C.GLenum(arg4), C.GLenum(arg5), C.GLenum(arg6), C.GLboolean(arg7), C.GLboolean(arg8), C.GLboolean(arg9));
}

//extern void glFinalCombinerInputNV (GLenum, GLenum, GLenum, GLenum)
func FinalCombinerInputNV(arg0 GLenum, arg1 GLenum, arg2 GLenum, arg3 GLenum) 
{
	C.glFinalCombinerInputNV(C.GLenum(arg0), C.GLenum(arg1), C.GLenum(arg2), C.GLenum(arg3));
}

//extern void glGetCombinerInputParameterfvNV (GLenum, GLenum, GLenum, GLenum, GLfloat *)
func GetCombinerInputParameterfvNV(arg0 GLenum, arg1 GLenum, arg2 GLenum, arg3 GLenum, arg4 *GLfloat) 
{
	C.glGetCombinerInputParameterfvNV(C.GLenum(arg0), C.GLenum(arg1), C.GLenum(arg2), C.GLenum(arg3), (*C.GLfloat)(arg4));
}

//extern void glGetCombinerInputParameterivNV (GLenum, GLenum, GLenum, GLenum, GLint *)
func GetCombinerInputParameterivNV(arg0 GLenum, arg1 GLenum, arg2 GLenum, arg3 GLenum, arg4 *GLint) 
{
	C.glGetCombinerInputParameterivNV(C.GLenum(arg0), C.GLenum(arg1), C.GLenum(arg2), C.GLenum(arg3), (*C.GLint)(arg4));
}

//extern void glGetCombinerOutputParameterfvNV (GLenum, GLenum, GLenum, GLfloat *)
func GetCombinerOutputParameterfvNV(arg0 GLenum, arg1 GLenum, arg2 GLenum, arg3 *GLfloat) 
{
	C.glGetCombinerOutputParameterfvNV(C.GLenum(arg0), C.GLenum(arg1), C.GLenum(arg2), (*C.GLfloat)(arg3));
}

//extern void glGetCombinerOutputParameterivNV (GLenum, GLenum, GLenum, GLint *)
func GetCombinerOutputParameterivNV(arg0 GLenum, arg1 GLenum, arg2 GLenum, arg3 *GLint) 
{
	C.glGetCombinerOutputParameterivNV(C.GLenum(arg0), C.GLenum(arg1), C.GLenum(arg2), (*C.GLint)(arg3));
}

//extern void glGetFinalCombinerInputParameterfvNV (GLenum, GLenum, GLfloat *)
func GetFinalCombinerInputParameterfvNV(arg0 GLenum, arg1 GLenum, arg2 *GLfloat) 
{
	C.glGetFinalCombinerInputParameterfvNV(C.GLenum(arg0), C.GLenum(arg1), (*C.GLfloat)(arg2));
}

//extern void glGetFinalCombinerInputParameterivNV (GLenum, GLenum, GLint *)
func GetFinalCombinerInputParameterivNV(arg0 GLenum, arg1 GLenum, arg2 *GLint) 
{
	C.glGetFinalCombinerInputParameterivNV(C.GLenum(arg0), C.GLenum(arg1), (*C.GLint)(arg2));
}

//extern void glResizeBuffersMESA (void)
func ResizeBuffersMESA() 
{
	C.glResizeBuffersMESA();
}

//extern void glWindowPos2dMESA (GLdouble, GLdouble)
func WindowPos2dMESA(arg0 GLdouble, arg1 GLdouble) 
{
	C.glWindowPos2dMESA(C.GLdouble(arg0), C.GLdouble(arg1));
}

//extern void glWindowPos2dvMESA (const GLdouble *)
func WindowPos2dvMESA(arg0 *GLdouble) 
{
	C.glWindowPos2dvMESA((*C.GLdouble)(arg0));
}

//extern void glWindowPos2fMESA (GLfloat, GLfloat)
func WindowPos2fMESA(arg0 GLfloat, arg1 GLfloat) 
{
	C.glWindowPos2fMESA(C.GLfloat(arg0), C.GLfloat(arg1));
}

//extern void glWindowPos2fvMESA (const GLfloat *)
func WindowPos2fvMESA(arg0 *GLfloat) 
{
	C.glWindowPos2fvMESA((*C.GLfloat)(arg0));
}

//extern void glWindowPos2iMESA (GLint, GLint)
func WindowPos2iMESA(arg0 GLint, arg1 GLint) 
{
	C.glWindowPos2iMESA(C.GLint(arg0), C.GLint(arg1));
}

//extern void glWindowPos2ivMESA (const GLint *)
func WindowPos2ivMESA(arg0 *GLint) 
{
	C.glWindowPos2ivMESA((*C.GLint)(arg0));
}

//extern void glWindowPos2sMESA (GLshort, GLshort)
func WindowPos2sMESA(arg0 GLshort, arg1 GLshort) 
{
	C.glWindowPos2sMESA(C.GLshort(arg0), C.GLshort(arg1));
}

//extern void glWindowPos2svMESA (const GLshort *)
func WindowPos2svMESA(arg0 *GLshort) 
{
	C.glWindowPos2svMESA((*C.GLshort)(arg0));
}

//extern void glWindowPos3dMESA (GLdouble, GLdouble, GLdouble)
func WindowPos3dMESA(arg0 GLdouble, arg1 GLdouble, arg2 GLdouble) 
{
	C.glWindowPos3dMESA(C.GLdouble(arg0), C.GLdouble(arg1), C.GLdouble(arg2));
}

//extern void glWindowPos3dvMESA (const GLdouble *)
func WindowPos3dvMESA(arg0 *GLdouble) 
{
	C.glWindowPos3dvMESA((*C.GLdouble)(arg0));
}

//extern void glWindowPos3fMESA (GLfloat, GLfloat, GLfloat)
func WindowPos3fMESA(arg0 GLfloat, arg1 GLfloat, arg2 GLfloat) 
{
	C.glWindowPos3fMESA(C.GLfloat(arg0), C.GLfloat(arg1), C.GLfloat(arg2));
}

//extern void glWindowPos3fvMESA (const GLfloat *)
func WindowPos3fvMESA(arg0 *GLfloat) 
{
	C.glWindowPos3fvMESA((*C.GLfloat)(arg0));
}

//extern void glWindowPos3iMESA (GLint, GLint, GLint)
func WindowPos3iMESA(arg0 GLint, arg1 GLint, arg2 GLint) 
{
	C.glWindowPos3iMESA(C.GLint(arg0), C.GLint(arg1), C.GLint(arg2));
}

//extern void glWindowPos3ivMESA (const GLint *)
func WindowPos3ivMESA(arg0 *GLint) 
{
	C.glWindowPos3ivMESA((*C.GLint)(arg0));
}

//extern void glWindowPos3sMESA (GLshort, GLshort, GLshort)
func WindowPos3sMESA(arg0 GLshort, arg1 GLshort, arg2 GLshort) 
{
	C.glWindowPos3sMESA(C.GLshort(arg0), C.GLshort(arg1), C.GLshort(arg2));
}

//extern void glWindowPos3svMESA (const GLshort *)
func WindowPos3svMESA(arg0 *GLshort) 
{
	C.glWindowPos3svMESA((*C.GLshort)(arg0));
}

//extern void glWindowPos4dMESA (GLdouble, GLdouble, GLdouble, GLdouble)
func WindowPos4dMESA(arg0 GLdouble, arg1 GLdouble, arg2 GLdouble, arg3 GLdouble) 
{
	C.glWindowPos4dMESA(C.GLdouble(arg0), C.GLdouble(arg1), C.GLdouble(arg2), C.GLdouble(arg3));
}

//extern void glWindowPos4dvMESA (const GLdouble *)
func WindowPos4dvMESA(arg0 *GLdouble) 
{
	C.glWindowPos4dvMESA((*C.GLdouble)(arg0));
}

//extern void glWindowPos4fMESA (GLfloat, GLfloat, GLfloat, GLfloat)
func WindowPos4fMESA(arg0 GLfloat, arg1 GLfloat, arg2 GLfloat, arg3 GLfloat) 
{
	C.glWindowPos4fMESA(C.GLfloat(arg0), C.GLfloat(arg1), C.GLfloat(arg2), C.GLfloat(arg3));
}

//extern void glWindowPos4fvMESA (const GLfloat *)
func WindowPos4fvMESA(arg0 *GLfloat) 
{
	C.glWindowPos4fvMESA((*C.GLfloat)(arg0));
}

//extern void glWindowPos4iMESA (GLint, GLint, GLint, GLint)
func WindowPos4iMESA(arg0 GLint, arg1 GLint, arg2 GLint, arg3 GLint) 
{
	C.glWindowPos4iMESA(C.GLint(arg0), C.GLint(arg1), C.GLint(arg2), C.GLint(arg3));
}

//extern void glWindowPos4ivMESA (const GLint *)
func WindowPos4ivMESA(arg0 *GLint) 
{
	C.glWindowPos4ivMESA((*C.GLint)(arg0));
}

//extern void glWindowPos4sMESA (GLshort, GLshort, GLshort, GLshort)
func WindowPos4sMESA(arg0 GLshort, arg1 GLshort, arg2 GLshort, arg3 GLshort) 
{
	C.glWindowPos4sMESA(C.GLshort(arg0), C.GLshort(arg1), C.GLshort(arg2), C.GLshort(arg3));
}

//extern void glWindowPos4svMESA (const GLshort *)
func WindowPos4svMESA(arg0 *GLshort) 
{
	C.glWindowPos4svMESA((*C.GLshort)(arg0));
}

//extern void glMultiModeDrawArraysIBM (const GLenum *, const GLint *, const GLsizei *, GLsizei, GLint)
func MultiModeDrawArraysIBM(arg0 *GLenum, arg1 *GLint, arg2 *GLsizei, arg3 GLsizei, arg4 GLint) 
{
	C.glMultiModeDrawArraysIBM((*C.GLenum)(arg0), (*C.GLint)(arg1), (*C.GLsizei)(arg2), C.GLsizei(arg3), C.GLint(arg4));
}

//extern void glMultiModeDrawElementsIBM (const GLenum *, const GLsizei *, GLenum, const GLvoid* const *, GLsizei, GLint)
func MultiModeDrawElementsIBM(arg0 *GLenum, arg1 *GLsizei, arg2 GLenum, arg3 *unsafe.Pointer, arg4 GLsizei, arg5 GLint) 
{
	C.glMultiModeDrawElementsIBM((*C.GLenum)(arg0), (*C.GLsizei)(arg1), C.GLenum(arg2), arg3, C.GLsizei(arg4), C.GLint(arg5));
}

//extern void glColorPointerListIBM (GLint, GLenum, GLint, const GLvoid* *, GLint)
func ColorPointerListIBM(arg0 GLint, arg1 GLenum, arg2 GLint, arg3 *unsafe.Pointer, arg4 GLint) 
{
	C.glColorPointerListIBM(C.GLint(arg0), C.GLenum(arg1), C.GLint(arg2), arg3, C.GLint(arg4));
}

//extern void glSecondaryColorPointerListIBM (GLint, GLenum, GLint, const GLvoid* *, GLint)
func SecondaryColorPointerListIBM(arg0 GLint, arg1 GLenum, arg2 GLint, arg3 *unsafe.Pointer, arg4 GLint) 
{
	C.glSecondaryColorPointerListIBM(C.GLint(arg0), C.GLenum(arg1), C.GLint(arg2), arg3, C.GLint(arg4));
}

//extern void glEdgeFlagPointerListIBM (GLint, const GLboolean* *, GLint)
func EdgeFlagPointerListIBM(arg0 GLint, arg1 **GLboolean, arg2 GLint) 
{
	C.glEdgeFlagPointerListIBM(C.GLint(arg0), (**C.GLboolean)(arg1), C.GLint(arg2));
}

//extern void glFogCoordPointerListIBM (GLenum, GLint, const GLvoid* *, GLint)
func FogCoordPointerListIBM(arg0 GLenum, arg1 GLint, arg2 *unsafe.Pointer, arg3 GLint) 
{
	C.glFogCoordPointerListIBM(C.GLenum(arg0), C.GLint(arg1), arg2, C.GLint(arg3));
}

//extern void glIndexPointerListIBM (GLenum, GLint, const GLvoid* *, GLint)
func IndexPointerListIBM(arg0 GLenum, arg1 GLint, arg2 *unsafe.Pointer, arg3 GLint) 
{
	C.glIndexPointerListIBM(C.GLenum(arg0), C.GLint(arg1), arg2, C.GLint(arg3));
}

//extern void glNormalPointerListIBM (GLenum, GLint, const GLvoid* *, GLint)
func NormalPointerListIBM(arg0 GLenum, arg1 GLint, arg2 *unsafe.Pointer, arg3 GLint) 
{
	C.glNormalPointerListIBM(C.GLenum(arg0), C.GLint(arg1), arg2, C.GLint(arg3));
}

//extern void glTexCoordPointerListIBM (GLint, GLenum, GLint, const GLvoid* *, GLint)
func TexCoordPointerListIBM(arg0 GLint, arg1 GLenum, arg2 GLint, arg3 *unsafe.Pointer, arg4 GLint) 
{
	C.glTexCoordPointerListIBM(C.GLint(arg0), C.GLenum(arg1), C.GLint(arg2), arg3, C.GLint(arg4));
}

//extern void glVertexPointerListIBM (GLint, GLenum, GLint, const GLvoid* *, GLint)
func VertexPointerListIBM(arg0 GLint, arg1 GLenum, arg2 GLint, arg3 *unsafe.Pointer, arg4 GLint) 
{
	C.glVertexPointerListIBM(C.GLint(arg0), C.GLenum(arg1), C.GLint(arg2), arg3, C.GLint(arg4));
}

//extern void glTbufferMask3DFX (GLuint)
func TbufferMask3DFX(arg0 GLuint) 
{
	C.glTbufferMask3DFX(C.GLuint(arg0));
}

//extern void glSampleMaskEXT (GLclampf, GLboolean)
func SampleMaskEXT(arg0 GLclampf, arg1 GLboolean) 
{
	C.glSampleMaskEXT(C.GLclampf(arg0), C.GLboolean(arg1));
}

//extern void glSamplePatternEXT (GLenum)
func SamplePatternEXT(arg0 GLenum) 
{
	C.glSamplePatternEXT(C.GLenum(arg0));
}

//extern void glTextureColorMaskSGIS (GLboolean, GLboolean, GLboolean, GLboolean)
func TextureColorMaskSGIS(arg0 GLboolean, arg1 GLboolean, arg2 GLboolean, arg3 GLboolean) 
{
	C.glTextureColorMaskSGIS(C.GLboolean(arg0), C.GLboolean(arg1), C.GLboolean(arg2), C.GLboolean(arg3));
}

//extern void glIglooInterfaceSGIX (GLenum, const GLvoid *)
func IglooInterfaceSGIX(arg0 GLenum, arg1 unsafe.Pointer) 
{
	C.glIglooInterfaceSGIX(C.GLenum(arg0), arg1);
}

//extern void glDeleteFencesNV (GLsizei, const GLuint *)
func DeleteFencesNV(arg0 GLsizei, arg1 *GLuint) 
{
	C.glDeleteFencesNV(C.GLsizei(arg0), (*C.GLuint)(arg1));
}

//extern void glGenFencesNV (GLsizei, GLuint *)
func GenFencesNV(arg0 GLsizei, arg1 *GLuint) 
{
	C.glGenFencesNV(C.GLsizei(arg0), (*C.GLuint)(arg1));
}

//extern GLboolean glIsFenceNV (GLuint)
func IsFenceNV(arg0 GLuint) GLboolean
{
	return GLboolean(C.glIsFenceNV(C.GLuint(arg0)));
}

//extern GLboolean glTestFenceNV (GLuint)
func TestFenceNV(arg0 GLuint) GLboolean
{
	return GLboolean(C.glTestFenceNV(C.GLuint(arg0)));
}

//extern void glGetFenceivNV (GLuint, GLenum, GLint *)
func GetFenceivNV(arg0 GLuint, arg1 GLenum, arg2 *GLint) 
{
	C.glGetFenceivNV(C.GLuint(arg0), C.GLenum(arg1), (*C.GLint)(arg2));
}

//extern void glFinishFenceNV (GLuint)
func FinishFenceNV(arg0 GLuint) 
{
	C.glFinishFenceNV(C.GLuint(arg0));
}

//extern void glSetFenceNV (GLuint, GLenum)
func SetFenceNV(arg0 GLuint, arg1 GLenum) 
{
	C.glSetFenceNV(C.GLuint(arg0), C.GLenum(arg1));
}

//extern void glMapControlPointsNV (GLenum, GLuint, GLenum, GLsizei, GLsizei, GLint, GLint, GLboolean, const GLvoid *)
func MapControlPointsNV(arg0 GLenum, arg1 GLuint, arg2 GLenum, arg3 GLsizei, arg4 GLsizei, arg5 GLint, arg6 GLint, arg7 GLboolean, arg8 unsafe.Pointer) 
{
	C.glMapControlPointsNV(C.GLenum(arg0), C.GLuint(arg1), C.GLenum(arg2), C.GLsizei(arg3), C.GLsizei(arg4), C.GLint(arg5), C.GLint(arg6), C.GLboolean(arg7), arg8);
}

//extern void glMapParameterivNV (GLenum, GLenum, const GLint *)
func MapParameterivNV(arg0 GLenum, arg1 GLenum, arg2 *GLint) 
{
	C.glMapParameterivNV(C.GLenum(arg0), C.GLenum(arg1), (*C.GLint)(arg2));
}

//extern void glMapParameterfvNV (GLenum, GLenum, const GLfloat *)
func MapParameterfvNV(arg0 GLenum, arg1 GLenum, arg2 *GLfloat) 
{
	C.glMapParameterfvNV(C.GLenum(arg0), C.GLenum(arg1), (*C.GLfloat)(arg2));
}

//extern void glGetMapControlPointsNV (GLenum, GLuint, GLenum, GLsizei, GLsizei, GLboolean, GLvoid *)
func GetMapControlPointsNV(arg0 GLenum, arg1 GLuint, arg2 GLenum, arg3 GLsizei, arg4 GLsizei, arg5 GLboolean, arg6 unsafe.Pointer) 
{
	C.glGetMapControlPointsNV(C.GLenum(arg0), C.GLuint(arg1), C.GLenum(arg2), C.GLsizei(arg3), C.GLsizei(arg4), C.GLboolean(arg5), arg6);
}

//extern void glGetMapParameterivNV (GLenum, GLenum, GLint *)
func GetMapParameterivNV(arg0 GLenum, arg1 GLenum, arg2 *GLint) 
{
	C.glGetMapParameterivNV(C.GLenum(arg0), C.GLenum(arg1), (*C.GLint)(arg2));
}

//extern void glGetMapParameterfvNV (GLenum, GLenum, GLfloat *)
func GetMapParameterfvNV(arg0 GLenum, arg1 GLenum, arg2 *GLfloat) 
{
	C.glGetMapParameterfvNV(C.GLenum(arg0), C.GLenum(arg1), (*C.GLfloat)(arg2));
}

//extern void glGetMapAttribParameterivNV (GLenum, GLuint, GLenum, GLint *)
func GetMapAttribParameterivNV(arg0 GLenum, arg1 GLuint, arg2 GLenum, arg3 *GLint) 
{
	C.glGetMapAttribParameterivNV(C.GLenum(arg0), C.GLuint(arg1), C.GLenum(arg2), (*C.GLint)(arg3));
}

//extern void glGetMapAttribParameterfvNV (GLenum, GLuint, GLenum, GLfloat *)
func GetMapAttribParameterfvNV(arg0 GLenum, arg1 GLuint, arg2 GLenum, arg3 *GLfloat) 
{
	C.glGetMapAttribParameterfvNV(C.GLenum(arg0), C.GLuint(arg1), C.GLenum(arg2), (*C.GLfloat)(arg3));
}

//extern void glEvalMapsNV (GLenum, GLenum)
func EvalMapsNV(arg0 GLenum, arg1 GLenum) 
{
	C.glEvalMapsNV(C.GLenum(arg0), C.GLenum(arg1));
}

//extern void glCombinerStageParameterfvNV (GLenum, GLenum, const GLfloat *)
func CombinerStageParameterfvNV(arg0 GLenum, arg1 GLenum, arg2 *GLfloat) 
{
	C.glCombinerStageParameterfvNV(C.GLenum(arg0), C.GLenum(arg1), (*C.GLfloat)(arg2));
}

//extern void glGetCombinerStageParameterfvNV (GLenum, GLenum, GLfloat *)
func GetCombinerStageParameterfvNV(arg0 GLenum, arg1 GLenum, arg2 *GLfloat) 
{
	C.glGetCombinerStageParameterfvNV(C.GLenum(arg0), C.GLenum(arg1), (*C.GLfloat)(arg2));
}

//extern GLboolean glAreProgramsResidentNV (GLsizei, const GLuint *, GLboolean *)
func AreProgramsResidentNV(arg0 GLsizei, arg1 *GLuint, arg2 *GLboolean) GLboolean
{
	return GLboolean(C.glAreProgramsResidentNV(C.GLsizei(arg0), (*C.GLuint)(arg1), (*C.GLboolean)(arg2)));
}

//extern void glBindProgramNV (GLenum, GLuint)
func BindProgramNV(arg0 GLenum, arg1 GLuint) 
{
	C.glBindProgramNV(C.GLenum(arg0), C.GLuint(arg1));
}

//extern void glDeleteProgramsNV (GLsizei, const GLuint *)
func DeleteProgramsNV(arg0 GLsizei, arg1 *GLuint) 
{
	C.glDeleteProgramsNV(C.GLsizei(arg0), (*C.GLuint)(arg1));
}

//extern void glExecuteProgramNV (GLenum, GLuint, const GLfloat *)
func ExecuteProgramNV(arg0 GLenum, arg1 GLuint, arg2 *GLfloat) 
{
	C.glExecuteProgramNV(C.GLenum(arg0), C.GLuint(arg1), (*C.GLfloat)(arg2));
}

//extern void glGenProgramsNV (GLsizei, GLuint *)
func GenProgramsNV(arg0 GLsizei, arg1 *GLuint) 
{
	C.glGenProgramsNV(C.GLsizei(arg0), (*C.GLuint)(arg1));
}

//extern void glGetProgramParameterdvNV (GLenum, GLuint, GLenum, GLdouble *)
func GetProgramParameterdvNV(arg0 GLenum, arg1 GLuint, arg2 GLenum, arg3 *GLdouble) 
{
	C.glGetProgramParameterdvNV(C.GLenum(arg0), C.GLuint(arg1), C.GLenum(arg2), (*C.GLdouble)(arg3));
}

//extern void glGetProgramParameterfvNV (GLenum, GLuint, GLenum, GLfloat *)
func GetProgramParameterfvNV(arg0 GLenum, arg1 GLuint, arg2 GLenum, arg3 *GLfloat) 
{
	C.glGetProgramParameterfvNV(C.GLenum(arg0), C.GLuint(arg1), C.GLenum(arg2), (*C.GLfloat)(arg3));
}

//extern void glGetProgramivNV (GLuint, GLenum, GLint *)
func GetProgramivNV(arg0 GLuint, arg1 GLenum, arg2 *GLint) 
{
	C.glGetProgramivNV(C.GLuint(arg0), C.GLenum(arg1), (*C.GLint)(arg2));
}

//extern void glGetProgramStringNV (GLuint, GLenum, GLubyte *)
func GetProgramStringNV(arg0 GLuint, arg1 GLenum, arg2 *GLubyte) 
{
	C.glGetProgramStringNV(C.GLuint(arg0), C.GLenum(arg1), (*C.GLubyte)(arg2));
}

//extern void glGetTrackMatrixivNV (GLenum, GLuint, GLenum, GLint *)
func GetTrackMatrixivNV(arg0 GLenum, arg1 GLuint, arg2 GLenum, arg3 *GLint) 
{
	C.glGetTrackMatrixivNV(C.GLenum(arg0), C.GLuint(arg1), C.GLenum(arg2), (*C.GLint)(arg3));
}

//extern void glGetVertexAttribdvNV (GLuint, GLenum, GLdouble *)
func GetVertexAttribdvNV(arg0 GLuint, arg1 GLenum, arg2 *GLdouble) 
{
	C.glGetVertexAttribdvNV(C.GLuint(arg0), C.GLenum(arg1), (*C.GLdouble)(arg2));
}

//extern void glGetVertexAttribfvNV (GLuint, GLenum, GLfloat *)
func GetVertexAttribfvNV(arg0 GLuint, arg1 GLenum, arg2 *GLfloat) 
{
	C.glGetVertexAttribfvNV(C.GLuint(arg0), C.GLenum(arg1), (*C.GLfloat)(arg2));
}

//extern void glGetVertexAttribivNV (GLuint, GLenum, GLint *)
func GetVertexAttribivNV(arg0 GLuint, arg1 GLenum, arg2 *GLint) 
{
	C.glGetVertexAttribivNV(C.GLuint(arg0), C.GLenum(arg1), (*C.GLint)(arg2));
}

//extern void glGetVertexAttribPointervNV (GLuint, GLenum, GLvoid* *)
func GetVertexAttribPointervNV(arg0 GLuint, arg1 GLenum, arg2 *unsafe.Pointer) 
{
	C.glGetVertexAttribPointervNV(C.GLuint(arg0), C.GLenum(arg1), arg2);
}

//extern GLboolean glIsProgramNV (GLuint)
func IsProgramNV(arg0 GLuint) GLboolean
{
	return GLboolean(C.glIsProgramNV(C.GLuint(arg0)));
}

//extern void glLoadProgramNV (GLenum, GLuint, GLsizei, const GLubyte *)
func LoadProgramNV(arg0 GLenum, arg1 GLuint, arg2 GLsizei, arg3 *GLubyte) 
{
	C.glLoadProgramNV(C.GLenum(arg0), C.GLuint(arg1), C.GLsizei(arg2), (*C.GLubyte)(arg3));
}

//extern void glProgramParameter4dNV (GLenum, GLuint, GLdouble, GLdouble, GLdouble, GLdouble)
func ProgramParameter4dNV(arg0 GLenum, arg1 GLuint, arg2 GLdouble, arg3 GLdouble, arg4 GLdouble, arg5 GLdouble) 
{
	C.glProgramParameter4dNV(C.GLenum(arg0), C.GLuint(arg1), C.GLdouble(arg2), C.GLdouble(arg3), C.GLdouble(arg4), C.GLdouble(arg5));
}

//extern void glProgramParameter4dvNV (GLenum, GLuint, const GLdouble *)
func ProgramParameter4dvNV(arg0 GLenum, arg1 GLuint, arg2 *GLdouble) 
{
	C.glProgramParameter4dvNV(C.GLenum(arg0), C.GLuint(arg1), (*C.GLdouble)(arg2));
}

//extern void glProgramParameter4fNV (GLenum, GLuint, GLfloat, GLfloat, GLfloat, GLfloat)
func ProgramParameter4fNV(arg0 GLenum, arg1 GLuint, arg2 GLfloat, arg3 GLfloat, arg4 GLfloat, arg5 GLfloat) 
{
	C.glProgramParameter4fNV(C.GLenum(arg0), C.GLuint(arg1), C.GLfloat(arg2), C.GLfloat(arg3), C.GLfloat(arg4), C.GLfloat(arg5));
}

//extern void glProgramParameter4fvNV (GLenum, GLuint, const GLfloat *)
func ProgramParameter4fvNV(arg0 GLenum, arg1 GLuint, arg2 *GLfloat) 
{
	C.glProgramParameter4fvNV(C.GLenum(arg0), C.GLuint(arg1), (*C.GLfloat)(arg2));
}

//extern void glProgramParameters4dvNV (GLenum, GLuint, GLuint, const GLdouble *)
func ProgramParameters4dvNV(arg0 GLenum, arg1 GLuint, arg2 GLuint, arg3 *GLdouble) 
{
	C.glProgramParameters4dvNV(C.GLenum(arg0), C.GLuint(arg1), C.GLuint(arg2), (*C.GLdouble)(arg3));
}

//extern void glProgramParameters4fvNV (GLenum, GLuint, GLuint, const GLfloat *)
func ProgramParameters4fvNV(arg0 GLenum, arg1 GLuint, arg2 GLuint, arg3 *GLfloat) 
{
	C.glProgramParameters4fvNV(C.GLenum(arg0), C.GLuint(arg1), C.GLuint(arg2), (*C.GLfloat)(arg3));
}

//extern void glRequestResidentProgramsNV (GLsizei, const GLuint *)
func RequestResidentProgramsNV(arg0 GLsizei, arg1 *GLuint) 
{
	C.glRequestResidentProgramsNV(C.GLsizei(arg0), (*C.GLuint)(arg1));
}

//extern void glTrackMatrixNV (GLenum, GLuint, GLenum, GLenum)
func TrackMatrixNV(arg0 GLenum, arg1 GLuint, arg2 GLenum, arg3 GLenum) 
{
	C.glTrackMatrixNV(C.GLenum(arg0), C.GLuint(arg1), C.GLenum(arg2), C.GLenum(arg3));
}

//extern void glVertexAttribPointerNV (GLuint, GLint, GLenum, GLsizei, const GLvoid *)
func VertexAttribPointerNV(arg0 GLuint, arg1 GLint, arg2 GLenum, arg3 GLsizei, arg4 unsafe.Pointer) 
{
	C.glVertexAttribPointerNV(C.GLuint(arg0), C.GLint(arg1), C.GLenum(arg2), C.GLsizei(arg3), arg4);
}

//extern void glVertexAttrib1dNV (GLuint, GLdouble)
func VertexAttrib1dNV(arg0 GLuint, arg1 GLdouble) 
{
	C.glVertexAttrib1dNV(C.GLuint(arg0), C.GLdouble(arg1));
}

//extern void glVertexAttrib1dvNV (GLuint, const GLdouble *)
func VertexAttrib1dvNV(arg0 GLuint, arg1 *GLdouble) 
{
	C.glVertexAttrib1dvNV(C.GLuint(arg0), (*C.GLdouble)(arg1));
}

//extern void glVertexAttrib1fNV (GLuint, GLfloat)
func VertexAttrib1fNV(arg0 GLuint, arg1 GLfloat) 
{
	C.glVertexAttrib1fNV(C.GLuint(arg0), C.GLfloat(arg1));
}

//extern void glVertexAttrib1fvNV (GLuint, const GLfloat *)
func VertexAttrib1fvNV(arg0 GLuint, arg1 *GLfloat) 
{
	C.glVertexAttrib1fvNV(C.GLuint(arg0), (*C.GLfloat)(arg1));
}

//extern void glVertexAttrib1sNV (GLuint, GLshort)
func VertexAttrib1sNV(arg0 GLuint, arg1 GLshort) 
{
	C.glVertexAttrib1sNV(C.GLuint(arg0), C.GLshort(arg1));
}

//extern void glVertexAttrib1svNV (GLuint, const GLshort *)
func VertexAttrib1svNV(arg0 GLuint, arg1 *GLshort) 
{
	C.glVertexAttrib1svNV(C.GLuint(arg0), (*C.GLshort)(arg1));
}

//extern void glVertexAttrib2dNV (GLuint, GLdouble, GLdouble)
func VertexAttrib2dNV(arg0 GLuint, arg1 GLdouble, arg2 GLdouble) 
{
	C.glVertexAttrib2dNV(C.GLuint(arg0), C.GLdouble(arg1), C.GLdouble(arg2));
}

//extern void glVertexAttrib2dvNV (GLuint, const GLdouble *)
func VertexAttrib2dvNV(arg0 GLuint, arg1 *GLdouble) 
{
	C.glVertexAttrib2dvNV(C.GLuint(arg0), (*C.GLdouble)(arg1));
}

//extern void glVertexAttrib2fNV (GLuint, GLfloat, GLfloat)
func VertexAttrib2fNV(arg0 GLuint, arg1 GLfloat, arg2 GLfloat) 
{
	C.glVertexAttrib2fNV(C.GLuint(arg0), C.GLfloat(arg1), C.GLfloat(arg2));
}

//extern void glVertexAttrib2fvNV (GLuint, const GLfloat *)
func VertexAttrib2fvNV(arg0 GLuint, arg1 *GLfloat) 
{
	C.glVertexAttrib2fvNV(C.GLuint(arg0), (*C.GLfloat)(arg1));
}

//extern void glVertexAttrib2sNV (GLuint, GLshort, GLshort)
func VertexAttrib2sNV(arg0 GLuint, arg1 GLshort, arg2 GLshort) 
{
	C.glVertexAttrib2sNV(C.GLuint(arg0), C.GLshort(arg1), C.GLshort(arg2));
}

//extern void glVertexAttrib2svNV (GLuint, const GLshort *)
func VertexAttrib2svNV(arg0 GLuint, arg1 *GLshort) 
{
	C.glVertexAttrib2svNV(C.GLuint(arg0), (*C.GLshort)(arg1));
}

//extern void glVertexAttrib3dNV (GLuint, GLdouble, GLdouble, GLdouble)
func VertexAttrib3dNV(arg0 GLuint, arg1 GLdouble, arg2 GLdouble, arg3 GLdouble) 
{
	C.glVertexAttrib3dNV(C.GLuint(arg0), C.GLdouble(arg1), C.GLdouble(arg2), C.GLdouble(arg3));
}

//extern void glVertexAttrib3dvNV (GLuint, const GLdouble *)
func VertexAttrib3dvNV(arg0 GLuint, arg1 *GLdouble) 
{
	C.glVertexAttrib3dvNV(C.GLuint(arg0), (*C.GLdouble)(arg1));
}

//extern void glVertexAttrib3fNV (GLuint, GLfloat, GLfloat, GLfloat)
func VertexAttrib3fNV(arg0 GLuint, arg1 GLfloat, arg2 GLfloat, arg3 GLfloat) 
{
	C.glVertexAttrib3fNV(C.GLuint(arg0), C.GLfloat(arg1), C.GLfloat(arg2), C.GLfloat(arg3));
}

//extern void glVertexAttrib3fvNV (GLuint, const GLfloat *)
func VertexAttrib3fvNV(arg0 GLuint, arg1 *GLfloat) 
{
	C.glVertexAttrib3fvNV(C.GLuint(arg0), (*C.GLfloat)(arg1));
}

//extern void glVertexAttrib3sNV (GLuint, GLshort, GLshort, GLshort)
func VertexAttrib3sNV(arg0 GLuint, arg1 GLshort, arg2 GLshort, arg3 GLshort) 
{
	C.glVertexAttrib3sNV(C.GLuint(arg0), C.GLshort(arg1), C.GLshort(arg2), C.GLshort(arg3));
}

//extern void glVertexAttrib3svNV (GLuint, const GLshort *)
func VertexAttrib3svNV(arg0 GLuint, arg1 *GLshort) 
{
	C.glVertexAttrib3svNV(C.GLuint(arg0), (*C.GLshort)(arg1));
}

//extern void glVertexAttrib4dNV (GLuint, GLdouble, GLdouble, GLdouble, GLdouble)
func VertexAttrib4dNV(arg0 GLuint, arg1 GLdouble, arg2 GLdouble, arg3 GLdouble, arg4 GLdouble) 
{
	C.glVertexAttrib4dNV(C.GLuint(arg0), C.GLdouble(arg1), C.GLdouble(arg2), C.GLdouble(arg3), C.GLdouble(arg4));
}

//extern void glVertexAttrib4dvNV (GLuint, const GLdouble *)
func VertexAttrib4dvNV(arg0 GLuint, arg1 *GLdouble) 
{
	C.glVertexAttrib4dvNV(C.GLuint(arg0), (*C.GLdouble)(arg1));
}

//extern void glVertexAttrib4fNV (GLuint, GLfloat, GLfloat, GLfloat, GLfloat)
func VertexAttrib4fNV(arg0 GLuint, arg1 GLfloat, arg2 GLfloat, arg3 GLfloat, arg4 GLfloat) 
{
	C.glVertexAttrib4fNV(C.GLuint(arg0), C.GLfloat(arg1), C.GLfloat(arg2), C.GLfloat(arg3), C.GLfloat(arg4));
}

//extern void glVertexAttrib4fvNV (GLuint, const GLfloat *)
func VertexAttrib4fvNV(arg0 GLuint, arg1 *GLfloat) 
{
	C.glVertexAttrib4fvNV(C.GLuint(arg0), (*C.GLfloat)(arg1));
}

//extern void glVertexAttrib4sNV (GLuint, GLshort, GLshort, GLshort, GLshort)
func VertexAttrib4sNV(arg0 GLuint, arg1 GLshort, arg2 GLshort, arg3 GLshort, arg4 GLshort) 
{
	C.glVertexAttrib4sNV(C.GLuint(arg0), C.GLshort(arg1), C.GLshort(arg2), C.GLshort(arg3), C.GLshort(arg4));
}

//extern void glVertexAttrib4svNV (GLuint, const GLshort *)
func VertexAttrib4svNV(arg0 GLuint, arg1 *GLshort) 
{
	C.glVertexAttrib4svNV(C.GLuint(arg0), (*C.GLshort)(arg1));
}

//extern void glVertexAttrib4ubNV (GLuint, GLubyte, GLubyte, GLubyte, GLubyte)
func VertexAttrib4ubNV(arg0 GLuint, arg1 GLubyte, arg2 GLubyte, arg3 GLubyte, arg4 GLubyte) 
{
	C.glVertexAttrib4ubNV(C.GLuint(arg0), C.GLubyte(arg1), C.GLubyte(arg2), C.GLubyte(arg3), C.GLubyte(arg4));
}

//extern void glVertexAttrib4ubvNV (GLuint, const GLubyte *)
func VertexAttrib4ubvNV(arg0 GLuint, arg1 *GLubyte) 
{
	C.glVertexAttrib4ubvNV(C.GLuint(arg0), (*C.GLubyte)(arg1));
}

//extern void glVertexAttribs1dvNV (GLuint, GLsizei, const GLdouble *)
func VertexAttribs1dvNV(arg0 GLuint, arg1 GLsizei, arg2 *GLdouble) 
{
	C.glVertexAttribs1dvNV(C.GLuint(arg0), C.GLsizei(arg1), (*C.GLdouble)(arg2));
}

//extern void glVertexAttribs1fvNV (GLuint, GLsizei, const GLfloat *)
func VertexAttribs1fvNV(arg0 GLuint, arg1 GLsizei, arg2 *GLfloat) 
{
	C.glVertexAttribs1fvNV(C.GLuint(arg0), C.GLsizei(arg1), (*C.GLfloat)(arg2));
}

//extern void glVertexAttribs1svNV (GLuint, GLsizei, const GLshort *)
func VertexAttribs1svNV(arg0 GLuint, arg1 GLsizei, arg2 *GLshort) 
{
	C.glVertexAttribs1svNV(C.GLuint(arg0), C.GLsizei(arg1), (*C.GLshort)(arg2));
}

//extern void glVertexAttribs2dvNV (GLuint, GLsizei, const GLdouble *)
func VertexAttribs2dvNV(arg0 GLuint, arg1 GLsizei, arg2 *GLdouble) 
{
	C.glVertexAttribs2dvNV(C.GLuint(arg0), C.GLsizei(arg1), (*C.GLdouble)(arg2));
}

//extern void glVertexAttribs2fvNV (GLuint, GLsizei, const GLfloat *)
func VertexAttribs2fvNV(arg0 GLuint, arg1 GLsizei, arg2 *GLfloat) 
{
	C.glVertexAttribs2fvNV(C.GLuint(arg0), C.GLsizei(arg1), (*C.GLfloat)(arg2));
}

//extern void glVertexAttribs2svNV (GLuint, GLsizei, const GLshort *)
func VertexAttribs2svNV(arg0 GLuint, arg1 GLsizei, arg2 *GLshort) 
{
	C.glVertexAttribs2svNV(C.GLuint(arg0), C.GLsizei(arg1), (*C.GLshort)(arg2));
}

//extern void glVertexAttribs3dvNV (GLuint, GLsizei, const GLdouble *)
func VertexAttribs3dvNV(arg0 GLuint, arg1 GLsizei, arg2 *GLdouble) 
{
	C.glVertexAttribs3dvNV(C.GLuint(arg0), C.GLsizei(arg1), (*C.GLdouble)(arg2));
}

//extern void glVertexAttribs3fvNV (GLuint, GLsizei, const GLfloat *)
func VertexAttribs3fvNV(arg0 GLuint, arg1 GLsizei, arg2 *GLfloat) 
{
	C.glVertexAttribs3fvNV(C.GLuint(arg0), C.GLsizei(arg1), (*C.GLfloat)(arg2));
}

//extern void glVertexAttribs3svNV (GLuint, GLsizei, const GLshort *)
func VertexAttribs3svNV(arg0 GLuint, arg1 GLsizei, arg2 *GLshort) 
{
	C.glVertexAttribs3svNV(C.GLuint(arg0), C.GLsizei(arg1), (*C.GLshort)(arg2));
}

//extern void glVertexAttribs4dvNV (GLuint, GLsizei, const GLdouble *)
func VertexAttribs4dvNV(arg0 GLuint, arg1 GLsizei, arg2 *GLdouble) 
{
	C.glVertexAttribs4dvNV(C.GLuint(arg0), C.GLsizei(arg1), (*C.GLdouble)(arg2));
}

//extern void glVertexAttribs4fvNV (GLuint, GLsizei, const GLfloat *)
func VertexAttribs4fvNV(arg0 GLuint, arg1 GLsizei, arg2 *GLfloat) 
{
	C.glVertexAttribs4fvNV(C.GLuint(arg0), C.GLsizei(arg1), (*C.GLfloat)(arg2));
}

//extern void glVertexAttribs4svNV (GLuint, GLsizei, const GLshort *)
func VertexAttribs4svNV(arg0 GLuint, arg1 GLsizei, arg2 *GLshort) 
{
	C.glVertexAttribs4svNV(C.GLuint(arg0), C.GLsizei(arg1), (*C.GLshort)(arg2));
}

//extern void glVertexAttribs4ubvNV (GLuint, GLsizei, const GLubyte *)
func VertexAttribs4ubvNV(arg0 GLuint, arg1 GLsizei, arg2 *GLubyte) 
{
	C.glVertexAttribs4ubvNV(C.GLuint(arg0), C.GLsizei(arg1), (*C.GLubyte)(arg2));
}

//extern void glTexBumpParameterivATI (GLenum, const GLint *)
func TexBumpParameterivATI(arg0 GLenum, arg1 *GLint) 
{
	C.glTexBumpParameterivATI(C.GLenum(arg0), (*C.GLint)(arg1));
}

//extern void glTexBumpParameterfvATI (GLenum, const GLfloat *)
func TexBumpParameterfvATI(arg0 GLenum, arg1 *GLfloat) 
{
	C.glTexBumpParameterfvATI(C.GLenum(arg0), (*C.GLfloat)(arg1));
}

//extern void glGetTexBumpParameterivATI (GLenum, GLint *)
func GetTexBumpParameterivATI(arg0 GLenum, arg1 *GLint) 
{
	C.glGetTexBumpParameterivATI(C.GLenum(arg0), (*C.GLint)(arg1));
}

//extern void glGetTexBumpParameterfvATI (GLenum, GLfloat *)
func GetTexBumpParameterfvATI(arg0 GLenum, arg1 *GLfloat) 
{
	C.glGetTexBumpParameterfvATI(C.GLenum(arg0), (*C.GLfloat)(arg1));
}

//extern GLuint glGenFragmentShadersATI (GLuint)
func GenFragmentShadersATI(arg0 GLuint) GLuint
{
	return GLuint(C.glGenFragmentShadersATI(C.GLuint(arg0)));
}

//extern void glBindFragmentShaderATI (GLuint)
func BindFragmentShaderATI(arg0 GLuint) 
{
	C.glBindFragmentShaderATI(C.GLuint(arg0));
}

//extern void glDeleteFragmentShaderATI (GLuint)
func DeleteFragmentShaderATI(arg0 GLuint) 
{
	C.glDeleteFragmentShaderATI(C.GLuint(arg0));
}

//extern void glBeginFragmentShaderATI (void)
func BeginFragmentShaderATI() 
{
	C.glBeginFragmentShaderATI();
}

//extern void glEndFragmentShaderATI (void)
func EndFragmentShaderATI() 
{
	C.glEndFragmentShaderATI();
}

//extern void glPassTexCoordATI (GLuint, GLuint, GLenum)
func PassTexCoordATI(arg0 GLuint, arg1 GLuint, arg2 GLenum) 
{
	C.glPassTexCoordATI(C.GLuint(arg0), C.GLuint(arg1), C.GLenum(arg2));
}

//extern void glSampleMapATI (GLuint, GLuint, GLenum)
func SampleMapATI(arg0 GLuint, arg1 GLuint, arg2 GLenum) 
{
	C.glSampleMapATI(C.GLuint(arg0), C.GLuint(arg1), C.GLenum(arg2));
}

//extern void glColorFragmentOp1ATI (GLenum, GLuint, GLuint, GLuint, GLuint, GLuint, GLuint)
func ColorFragmentOp1ATI(arg0 GLenum, arg1 GLuint, arg2 GLuint, arg3 GLuint, arg4 GLuint, arg5 GLuint, arg6 GLuint) 
{
	C.glColorFragmentOp1ATI(C.GLenum(arg0), C.GLuint(arg1), C.GLuint(arg2), C.GLuint(arg3), C.GLuint(arg4), C.GLuint(arg5), C.GLuint(arg6));
}

//extern void glColorFragmentOp2ATI (GLenum, GLuint, GLuint, GLuint, GLuint, GLuint, GLuint, GLuint, GLuint, GLuint)
func ColorFragmentOp2ATI(arg0 GLenum, arg1 GLuint, arg2 GLuint, arg3 GLuint, arg4 GLuint, arg5 GLuint, arg6 GLuint, arg7 GLuint, arg8 GLuint, arg9 GLuint) 
{
	C.glColorFragmentOp2ATI(C.GLenum(arg0), C.GLuint(arg1), C.GLuint(arg2), C.GLuint(arg3), C.GLuint(arg4), C.GLuint(arg5), C.GLuint(arg6), C.GLuint(arg7), C.GLuint(arg8), C.GLuint(arg9));
}

//extern void glColorFragmentOp3ATI (GLenum, GLuint, GLuint, GLuint, GLuint, GLuint, GLuint, GLuint, GLuint, GLuint, GLuint, GLuint, GLuint)
func ColorFragmentOp3ATI(arg0 GLenum, arg1 GLuint, arg2 GLuint, arg3 GLuint, arg4 GLuint, arg5 GLuint, arg6 GLuint, arg7 GLuint, arg8 GLuint, arg9 GLuint, arg10 GLuint, arg11 GLuint, arg12 GLuint) 
{
	C.glColorFragmentOp3ATI(C.GLenum(arg0), C.GLuint(arg1), C.GLuint(arg2), C.GLuint(arg3), C.GLuint(arg4), C.GLuint(arg5), C.GLuint(arg6), C.GLuint(arg7), C.GLuint(arg8), C.GLuint(arg9), C.GLuint(arg10), C.GLuint(arg11), C.GLuint(arg12));
}

//extern void glAlphaFragmentOp1ATI (GLenum, GLuint, GLuint, GLuint, GLuint, GLuint)
func AlphaFragmentOp1ATI(arg0 GLenum, arg1 GLuint, arg2 GLuint, arg3 GLuint, arg4 GLuint, arg5 GLuint) 
{
	C.glAlphaFragmentOp1ATI(C.GLenum(arg0), C.GLuint(arg1), C.GLuint(arg2), C.GLuint(arg3), C.GLuint(arg4), C.GLuint(arg5));
}

//extern void glAlphaFragmentOp2ATI (GLenum, GLuint, GLuint, GLuint, GLuint, GLuint, GLuint, GLuint, GLuint)
func AlphaFragmentOp2ATI(arg0 GLenum, arg1 GLuint, arg2 GLuint, arg3 GLuint, arg4 GLuint, arg5 GLuint, arg6 GLuint, arg7 GLuint, arg8 GLuint) 
{
	C.glAlphaFragmentOp2ATI(C.GLenum(arg0), C.GLuint(arg1), C.GLuint(arg2), C.GLuint(arg3), C.GLuint(arg4), C.GLuint(arg5), C.GLuint(arg6), C.GLuint(arg7), C.GLuint(arg8));
}

//extern void glAlphaFragmentOp3ATI (GLenum, GLuint, GLuint, GLuint, GLuint, GLuint, GLuint, GLuint, GLuint, GLuint, GLuint, GLuint)
func AlphaFragmentOp3ATI(arg0 GLenum, arg1 GLuint, arg2 GLuint, arg3 GLuint, arg4 GLuint, arg5 GLuint, arg6 GLuint, arg7 GLuint, arg8 GLuint, arg9 GLuint, arg10 GLuint, arg11 GLuint) 
{
	C.glAlphaFragmentOp3ATI(C.GLenum(arg0), C.GLuint(arg1), C.GLuint(arg2), C.GLuint(arg3), C.GLuint(arg4), C.GLuint(arg5), C.GLuint(arg6), C.GLuint(arg7), C.GLuint(arg8), C.GLuint(arg9), C.GLuint(arg10), C.GLuint(arg11));
}

//extern void glSetFragmentShaderConstantATI (GLuint, const GLfloat *)
func SetFragmentShaderConstantATI(arg0 GLuint, arg1 *GLfloat) 
{
	C.glSetFragmentShaderConstantATI(C.GLuint(arg0), (*C.GLfloat)(arg1));
}

//extern void glPNTrianglesiATI (GLenum, GLint)
func PNTrianglesiATI(arg0 GLenum, arg1 GLint) 
{
	C.glPNTrianglesiATI(C.GLenum(arg0), C.GLint(arg1));
}

//extern void glPNTrianglesfATI (GLenum, GLfloat)
func PNTrianglesfATI(arg0 GLenum, arg1 GLfloat) 
{
	C.glPNTrianglesfATI(C.GLenum(arg0), C.GLfloat(arg1));
}

//extern GLuint glNewObjectBufferATI (GLsizei, const GLvoid *, GLenum)
func NewObjectBufferATI(arg0 GLsizei, arg1 unsafe.Pointer, arg2 GLenum) GLuint
{
	return GLuint(C.glNewObjectBufferATI(C.GLsizei(arg0), arg1, C.GLenum(arg2)));
}

//extern GLboolean glIsObjectBufferATI (GLuint)
func IsObjectBufferATI(arg0 GLuint) GLboolean
{
	return GLboolean(C.glIsObjectBufferATI(C.GLuint(arg0)));
}

//extern void glUpdateObjectBufferATI (GLuint, GLuint, GLsizei, const GLvoid *, GLenum)
func UpdateObjectBufferATI(arg0 GLuint, arg1 GLuint, arg2 GLsizei, arg3 unsafe.Pointer, arg4 GLenum) 
{
	C.glUpdateObjectBufferATI(C.GLuint(arg0), C.GLuint(arg1), C.GLsizei(arg2), arg3, C.GLenum(arg4));
}

//extern void glGetObjectBufferfvATI (GLuint, GLenum, GLfloat *)
func GetObjectBufferfvATI(arg0 GLuint, arg1 GLenum, arg2 *GLfloat) 
{
	C.glGetObjectBufferfvATI(C.GLuint(arg0), C.GLenum(arg1), (*C.GLfloat)(arg2));
}

//extern void glGetObjectBufferivATI (GLuint, GLenum, GLint *)
func GetObjectBufferivATI(arg0 GLuint, arg1 GLenum, arg2 *GLint) 
{
	C.glGetObjectBufferivATI(C.GLuint(arg0), C.GLenum(arg1), (*C.GLint)(arg2));
}

//extern void glFreeObjectBufferATI (GLuint)
func FreeObjectBufferATI(arg0 GLuint) 
{
	C.glFreeObjectBufferATI(C.GLuint(arg0));
}

//extern void glArrayObjectATI (GLenum, GLint, GLenum, GLsizei, GLuint, GLuint)
func ArrayObjectATI(arg0 GLenum, arg1 GLint, arg2 GLenum, arg3 GLsizei, arg4 GLuint, arg5 GLuint) 
{
	C.glArrayObjectATI(C.GLenum(arg0), C.GLint(arg1), C.GLenum(arg2), C.GLsizei(arg3), C.GLuint(arg4), C.GLuint(arg5));
}

//extern void glGetArrayObjectfvATI (GLenum, GLenum, GLfloat *)
func GetArrayObjectfvATI(arg0 GLenum, arg1 GLenum, arg2 *GLfloat) 
{
	C.glGetArrayObjectfvATI(C.GLenum(arg0), C.GLenum(arg1), (*C.GLfloat)(arg2));
}

//extern void glGetArrayObjectivATI (GLenum, GLenum, GLint *)
func GetArrayObjectivATI(arg0 GLenum, arg1 GLenum, arg2 *GLint) 
{
	C.glGetArrayObjectivATI(C.GLenum(arg0), C.GLenum(arg1), (*C.GLint)(arg2));
}

//extern void glVariantArrayObjectATI (GLuint, GLenum, GLsizei, GLuint, GLuint)
func VariantArrayObjectATI(arg0 GLuint, arg1 GLenum, arg2 GLsizei, arg3 GLuint, arg4 GLuint) 
{
	C.glVariantArrayObjectATI(C.GLuint(arg0), C.GLenum(arg1), C.GLsizei(arg2), C.GLuint(arg3), C.GLuint(arg4));
}

//extern void glGetVariantArrayObjectfvATI (GLuint, GLenum, GLfloat *)
func GetVariantArrayObjectfvATI(arg0 GLuint, arg1 GLenum, arg2 *GLfloat) 
{
	C.glGetVariantArrayObjectfvATI(C.GLuint(arg0), C.GLenum(arg1), (*C.GLfloat)(arg2));
}

//extern void glGetVariantArrayObjectivATI (GLuint, GLenum, GLint *)
func GetVariantArrayObjectivATI(arg0 GLuint, arg1 GLenum, arg2 *GLint) 
{
	C.glGetVariantArrayObjectivATI(C.GLuint(arg0), C.GLenum(arg1), (*C.GLint)(arg2));
}

//extern void glBeginVertexShaderEXT (void)
func BeginVertexShaderEXT() 
{
	C.glBeginVertexShaderEXT();
}

//extern void glEndVertexShaderEXT (void)
func EndVertexShaderEXT() 
{
	C.glEndVertexShaderEXT();
}

//extern void glBindVertexShaderEXT (GLuint)
func BindVertexShaderEXT(arg0 GLuint) 
{
	C.glBindVertexShaderEXT(C.GLuint(arg0));
}

//extern GLuint glGenVertexShadersEXT (GLuint)
func GenVertexShadersEXT(arg0 GLuint) GLuint
{
	return GLuint(C.glGenVertexShadersEXT(C.GLuint(arg0)));
}

//extern void glDeleteVertexShaderEXT (GLuint)
func DeleteVertexShaderEXT(arg0 GLuint) 
{
	C.glDeleteVertexShaderEXT(C.GLuint(arg0));
}

//extern void glShaderOp1EXT (GLenum, GLuint, GLuint)
func ShaderOp1EXT(arg0 GLenum, arg1 GLuint, arg2 GLuint) 
{
	C.glShaderOp1EXT(C.GLenum(arg0), C.GLuint(arg1), C.GLuint(arg2));
}

//extern void glShaderOp2EXT (GLenum, GLuint, GLuint, GLuint)
func ShaderOp2EXT(arg0 GLenum, arg1 GLuint, arg2 GLuint, arg3 GLuint) 
{
	C.glShaderOp2EXT(C.GLenum(arg0), C.GLuint(arg1), C.GLuint(arg2), C.GLuint(arg3));
}

//extern void glShaderOp3EXT (GLenum, GLuint, GLuint, GLuint, GLuint)
func ShaderOp3EXT(arg0 GLenum, arg1 GLuint, arg2 GLuint, arg3 GLuint, arg4 GLuint) 
{
	C.glShaderOp3EXT(C.GLenum(arg0), C.GLuint(arg1), C.GLuint(arg2), C.GLuint(arg3), C.GLuint(arg4));
}

//extern void glSwizzleEXT (GLuint, GLuint, GLenum, GLenum, GLenum, GLenum)
func SwizzleEXT(arg0 GLuint, arg1 GLuint, arg2 GLenum, arg3 GLenum, arg4 GLenum, arg5 GLenum) 
{
	C.glSwizzleEXT(C.GLuint(arg0), C.GLuint(arg1), C.GLenum(arg2), C.GLenum(arg3), C.GLenum(arg4), C.GLenum(arg5));
}

//extern void glWriteMaskEXT (GLuint, GLuint, GLenum, GLenum, GLenum, GLenum)
func WriteMaskEXT(arg0 GLuint, arg1 GLuint, arg2 GLenum, arg3 GLenum, arg4 GLenum, arg5 GLenum) 
{
	C.glWriteMaskEXT(C.GLuint(arg0), C.GLuint(arg1), C.GLenum(arg2), C.GLenum(arg3), C.GLenum(arg4), C.GLenum(arg5));
}

//extern void glInsertComponentEXT (GLuint, GLuint, GLuint)
func InsertComponentEXT(arg0 GLuint, arg1 GLuint, arg2 GLuint) 
{
	C.glInsertComponentEXT(C.GLuint(arg0), C.GLuint(arg1), C.GLuint(arg2));
}

//extern void glExtractComponentEXT (GLuint, GLuint, GLuint)
func ExtractComponentEXT(arg0 GLuint, arg1 GLuint, arg2 GLuint) 
{
	C.glExtractComponentEXT(C.GLuint(arg0), C.GLuint(arg1), C.GLuint(arg2));
}

//extern GLuint glGenSymbolsEXT (GLenum, GLenum, GLenum, GLuint)
func GenSymbolsEXT(arg0 GLenum, arg1 GLenum, arg2 GLenum, arg3 GLuint) GLuint
{
	return GLuint(C.glGenSymbolsEXT(C.GLenum(arg0), C.GLenum(arg1), C.GLenum(arg2), C.GLuint(arg3)));
}

//extern void glSetInvariantEXT (GLuint, GLenum, const GLvoid *)
func SetInvariantEXT(arg0 GLuint, arg1 GLenum, arg2 unsafe.Pointer) 
{
	C.glSetInvariantEXT(C.GLuint(arg0), C.GLenum(arg1), arg2);
}

//extern void glSetLocalConstantEXT (GLuint, GLenum, const GLvoid *)
func SetLocalConstantEXT(arg0 GLuint, arg1 GLenum, arg2 unsafe.Pointer) 
{
	C.glSetLocalConstantEXT(C.GLuint(arg0), C.GLenum(arg1), arg2);
}

//extern void glVariantbvEXT (GLuint, const GLbyte *)
func VariantbvEXT(arg0 GLuint, arg1 *GLbyte) 
{
	C.glVariantbvEXT(C.GLuint(arg0), (*C.GLbyte)(arg1));
}

//extern void glVariantsvEXT (GLuint, const GLshort *)
func VariantsvEXT(arg0 GLuint, arg1 *GLshort) 
{
	C.glVariantsvEXT(C.GLuint(arg0), (*C.GLshort)(arg1));
}

//extern void glVariantivEXT (GLuint, const GLint *)
func VariantivEXT(arg0 GLuint, arg1 *GLint) 
{
	C.glVariantivEXT(C.GLuint(arg0), (*C.GLint)(arg1));
}

//extern void glVariantfvEXT (GLuint, const GLfloat *)
func VariantfvEXT(arg0 GLuint, arg1 *GLfloat) 
{
	C.glVariantfvEXT(C.GLuint(arg0), (*C.GLfloat)(arg1));
}

//extern void glVariantdvEXT (GLuint, const GLdouble *)
func VariantdvEXT(arg0 GLuint, arg1 *GLdouble) 
{
	C.glVariantdvEXT(C.GLuint(arg0), (*C.GLdouble)(arg1));
}

//extern void glVariantubvEXT (GLuint, const GLubyte *)
func VariantubvEXT(arg0 GLuint, arg1 *GLubyte) 
{
	C.glVariantubvEXT(C.GLuint(arg0), (*C.GLubyte)(arg1));
}

//extern void glVariantusvEXT (GLuint, const GLushort *)
func VariantusvEXT(arg0 GLuint, arg1 *GLushort) 
{
	C.glVariantusvEXT(C.GLuint(arg0), (*C.GLushort)(arg1));
}

//extern void glVariantuivEXT (GLuint, const GLuint *)
func VariantuivEXT(arg0 GLuint, arg1 *GLuint) 
{
	C.glVariantuivEXT(C.GLuint(arg0), (*C.GLuint)(arg1));
}

//extern void glVariantPointerEXT (GLuint, GLenum, GLuint, const GLvoid *)
func VariantPointerEXT(arg0 GLuint, arg1 GLenum, arg2 GLuint, arg3 unsafe.Pointer) 
{
	C.glVariantPointerEXT(C.GLuint(arg0), C.GLenum(arg1), C.GLuint(arg2), arg3);
}

//extern void glEnableVariantClientStateEXT (GLuint)
func EnableVariantClientStateEXT(arg0 GLuint) 
{
	C.glEnableVariantClientStateEXT(C.GLuint(arg0));
}

//extern void glDisableVariantClientStateEXT (GLuint)
func DisableVariantClientStateEXT(arg0 GLuint) 
{
	C.glDisableVariantClientStateEXT(C.GLuint(arg0));
}

//extern GLuint glBindLightParameterEXT (GLenum, GLenum)
func BindLightParameterEXT(arg0 GLenum, arg1 GLenum) GLuint
{
	return GLuint(C.glBindLightParameterEXT(C.GLenum(arg0), C.GLenum(arg1)));
}

//extern GLuint glBindMaterialParameterEXT (GLenum, GLenum)
func BindMaterialParameterEXT(arg0 GLenum, arg1 GLenum) GLuint
{
	return GLuint(C.glBindMaterialParameterEXT(C.GLenum(arg0), C.GLenum(arg1)));
}

//extern GLuint glBindTexGenParameterEXT (GLenum, GLenum, GLenum)
func BindTexGenParameterEXT(arg0 GLenum, arg1 GLenum, arg2 GLenum) GLuint
{
	return GLuint(C.glBindTexGenParameterEXT(C.GLenum(arg0), C.GLenum(arg1), C.GLenum(arg2)));
}

//extern GLuint glBindTextureUnitParameterEXT (GLenum, GLenum)
func BindTextureUnitParameterEXT(arg0 GLenum, arg1 GLenum) GLuint
{
	return GLuint(C.glBindTextureUnitParameterEXT(C.GLenum(arg0), C.GLenum(arg1)));
}

//extern GLuint glBindParameterEXT (GLenum)
func BindParameterEXT(arg0 GLenum) GLuint
{
	return GLuint(C.glBindParameterEXT(C.GLenum(arg0)));
}

//extern GLboolean glIsVariantEnabledEXT (GLuint, GLenum)
func IsVariantEnabledEXT(arg0 GLuint, arg1 GLenum) GLboolean
{
	return GLboolean(C.glIsVariantEnabledEXT(C.GLuint(arg0), C.GLenum(arg1)));
}

//extern void glGetVariantBooleanvEXT (GLuint, GLenum, GLboolean *)
func GetVariantBooleanvEXT(arg0 GLuint, arg1 GLenum, arg2 *GLboolean) 
{
	C.glGetVariantBooleanvEXT(C.GLuint(arg0), C.GLenum(arg1), (*C.GLboolean)(arg2));
}

//extern void glGetVariantIntegervEXT (GLuint, GLenum, GLint *)
func GetVariantIntegervEXT(arg0 GLuint, arg1 GLenum, arg2 *GLint) 
{
	C.glGetVariantIntegervEXT(C.GLuint(arg0), C.GLenum(arg1), (*C.GLint)(arg2));
}

//extern void glGetVariantFloatvEXT (GLuint, GLenum, GLfloat *)
func GetVariantFloatvEXT(arg0 GLuint, arg1 GLenum, arg2 *GLfloat) 
{
	C.glGetVariantFloatvEXT(C.GLuint(arg0), C.GLenum(arg1), (*C.GLfloat)(arg2));
}

//extern void glGetVariantPointervEXT (GLuint, GLenum, GLvoid* *)
func GetVariantPointervEXT(arg0 GLuint, arg1 GLenum, arg2 *unsafe.Pointer) 
{
	C.glGetVariantPointervEXT(C.GLuint(arg0), C.GLenum(arg1), arg2);
}

//extern void glGetInvariantBooleanvEXT (GLuint, GLenum, GLboolean *)
func GetInvariantBooleanvEXT(arg0 GLuint, arg1 GLenum, arg2 *GLboolean) 
{
	C.glGetInvariantBooleanvEXT(C.GLuint(arg0), C.GLenum(arg1), (*C.GLboolean)(arg2));
}

//extern void glGetInvariantIntegervEXT (GLuint, GLenum, GLint *)
func GetInvariantIntegervEXT(arg0 GLuint, arg1 GLenum, arg2 *GLint) 
{
	C.glGetInvariantIntegervEXT(C.GLuint(arg0), C.GLenum(arg1), (*C.GLint)(arg2));
}

//extern void glGetInvariantFloatvEXT (GLuint, GLenum, GLfloat *)
func GetInvariantFloatvEXT(arg0 GLuint, arg1 GLenum, arg2 *GLfloat) 
{
	C.glGetInvariantFloatvEXT(C.GLuint(arg0), C.GLenum(arg1), (*C.GLfloat)(arg2));
}

//extern void glGetLocalConstantBooleanvEXT (GLuint, GLenum, GLboolean *)
func GetLocalConstantBooleanvEXT(arg0 GLuint, arg1 GLenum, arg2 *GLboolean) 
{
	C.glGetLocalConstantBooleanvEXT(C.GLuint(arg0), C.GLenum(arg1), (*C.GLboolean)(arg2));
}

//extern void glGetLocalConstantIntegervEXT (GLuint, GLenum, GLint *)
func GetLocalConstantIntegervEXT(arg0 GLuint, arg1 GLenum, arg2 *GLint) 
{
	C.glGetLocalConstantIntegervEXT(C.GLuint(arg0), C.GLenum(arg1), (*C.GLint)(arg2));
}

//extern void glGetLocalConstantFloatvEXT (GLuint, GLenum, GLfloat *)
func GetLocalConstantFloatvEXT(arg0 GLuint, arg1 GLenum, arg2 *GLfloat) 
{
	C.glGetLocalConstantFloatvEXT(C.GLuint(arg0), C.GLenum(arg1), (*C.GLfloat)(arg2));
}

//extern void glVertexStream1sATI (GLenum, GLshort)
func VertexStream1sATI(arg0 GLenum, arg1 GLshort) 
{
	C.glVertexStream1sATI(C.GLenum(arg0), C.GLshort(arg1));
}

//extern void glVertexStream1svATI (GLenum, const GLshort *)
func VertexStream1svATI(arg0 GLenum, arg1 *GLshort) 
{
	C.glVertexStream1svATI(C.GLenum(arg0), (*C.GLshort)(arg1));
}

//extern void glVertexStream1iATI (GLenum, GLint)
func VertexStream1iATI(arg0 GLenum, arg1 GLint) 
{
	C.glVertexStream1iATI(C.GLenum(arg0), C.GLint(arg1));
}

//extern void glVertexStream1ivATI (GLenum, const GLint *)
func VertexStream1ivATI(arg0 GLenum, arg1 *GLint) 
{
	C.glVertexStream1ivATI(C.GLenum(arg0), (*C.GLint)(arg1));
}

//extern void glVertexStream1fATI (GLenum, GLfloat)
func VertexStream1fATI(arg0 GLenum, arg1 GLfloat) 
{
	C.glVertexStream1fATI(C.GLenum(arg0), C.GLfloat(arg1));
}

//extern void glVertexStream1fvATI (GLenum, const GLfloat *)
func VertexStream1fvATI(arg0 GLenum, arg1 *GLfloat) 
{
	C.glVertexStream1fvATI(C.GLenum(arg0), (*C.GLfloat)(arg1));
}

//extern void glVertexStream1dATI (GLenum, GLdouble)
func VertexStream1dATI(arg0 GLenum, arg1 GLdouble) 
{
	C.glVertexStream1dATI(C.GLenum(arg0), C.GLdouble(arg1));
}

//extern void glVertexStream1dvATI (GLenum, const GLdouble *)
func VertexStream1dvATI(arg0 GLenum, arg1 *GLdouble) 
{
	C.glVertexStream1dvATI(C.GLenum(arg0), (*C.GLdouble)(arg1));
}

//extern void glVertexStream2sATI (GLenum, GLshort, GLshort)
func VertexStream2sATI(arg0 GLenum, arg1 GLshort, arg2 GLshort) 
{
	C.glVertexStream2sATI(C.GLenum(arg0), C.GLshort(arg1), C.GLshort(arg2));
}

//extern void glVertexStream2svATI (GLenum, const GLshort *)
func VertexStream2svATI(arg0 GLenum, arg1 *GLshort) 
{
	C.glVertexStream2svATI(C.GLenum(arg0), (*C.GLshort)(arg1));
}

//extern void glVertexStream2iATI (GLenum, GLint, GLint)
func VertexStream2iATI(arg0 GLenum, arg1 GLint, arg2 GLint) 
{
	C.glVertexStream2iATI(C.GLenum(arg0), C.GLint(arg1), C.GLint(arg2));
}

//extern void glVertexStream2ivATI (GLenum, const GLint *)
func VertexStream2ivATI(arg0 GLenum, arg1 *GLint) 
{
	C.glVertexStream2ivATI(C.GLenum(arg0), (*C.GLint)(arg1));
}

//extern void glVertexStream2fATI (GLenum, GLfloat, GLfloat)
func VertexStream2fATI(arg0 GLenum, arg1 GLfloat, arg2 GLfloat) 
{
	C.glVertexStream2fATI(C.GLenum(arg0), C.GLfloat(arg1), C.GLfloat(arg2));
}

//extern void glVertexStream2fvATI (GLenum, const GLfloat *)
func VertexStream2fvATI(arg0 GLenum, arg1 *GLfloat) 
{
	C.glVertexStream2fvATI(C.GLenum(arg0), (*C.GLfloat)(arg1));
}

//extern void glVertexStream2dATI (GLenum, GLdouble, GLdouble)
func VertexStream2dATI(arg0 GLenum, arg1 GLdouble, arg2 GLdouble) 
{
	C.glVertexStream2dATI(C.GLenum(arg0), C.GLdouble(arg1), C.GLdouble(arg2));
}

//extern void glVertexStream2dvATI (GLenum, const GLdouble *)
func VertexStream2dvATI(arg0 GLenum, arg1 *GLdouble) 
{
	C.glVertexStream2dvATI(C.GLenum(arg0), (*C.GLdouble)(arg1));
}

//extern void glVertexStream3sATI (GLenum, GLshort, GLshort, GLshort)
func VertexStream3sATI(arg0 GLenum, arg1 GLshort, arg2 GLshort, arg3 GLshort) 
{
	C.glVertexStream3sATI(C.GLenum(arg0), C.GLshort(arg1), C.GLshort(arg2), C.GLshort(arg3));
}

//extern void glVertexStream3svATI (GLenum, const GLshort *)
func VertexStream3svATI(arg0 GLenum, arg1 *GLshort) 
{
	C.glVertexStream3svATI(C.GLenum(arg0), (*C.GLshort)(arg1));
}

//extern void glVertexStream3iATI (GLenum, GLint, GLint, GLint)
func VertexStream3iATI(arg0 GLenum, arg1 GLint, arg2 GLint, arg3 GLint) 
{
	C.glVertexStream3iATI(C.GLenum(arg0), C.GLint(arg1), C.GLint(arg2), C.GLint(arg3));
}

//extern void glVertexStream3ivATI (GLenum, const GLint *)
func VertexStream3ivATI(arg0 GLenum, arg1 *GLint) 
{
	C.glVertexStream3ivATI(C.GLenum(arg0), (*C.GLint)(arg1));
}

//extern void glVertexStream3fATI (GLenum, GLfloat, GLfloat, GLfloat)
func VertexStream3fATI(arg0 GLenum, arg1 GLfloat, arg2 GLfloat, arg3 GLfloat) 
{
	C.glVertexStream3fATI(C.GLenum(arg0), C.GLfloat(arg1), C.GLfloat(arg2), C.GLfloat(arg3));
}

//extern void glVertexStream3fvATI (GLenum, const GLfloat *)
func VertexStream3fvATI(arg0 GLenum, arg1 *GLfloat) 
{
	C.glVertexStream3fvATI(C.GLenum(arg0), (*C.GLfloat)(arg1));
}

//extern void glVertexStream3dATI (GLenum, GLdouble, GLdouble, GLdouble)
func VertexStream3dATI(arg0 GLenum, arg1 GLdouble, arg2 GLdouble, arg3 GLdouble) 
{
	C.glVertexStream3dATI(C.GLenum(arg0), C.GLdouble(arg1), C.GLdouble(arg2), C.GLdouble(arg3));
}

//extern void glVertexStream3dvATI (GLenum, const GLdouble *)
func VertexStream3dvATI(arg0 GLenum, arg1 *GLdouble) 
{
	C.glVertexStream3dvATI(C.GLenum(arg0), (*C.GLdouble)(arg1));
}

//extern void glVertexStream4sATI (GLenum, GLshort, GLshort, GLshort, GLshort)
func VertexStream4sATI(arg0 GLenum, arg1 GLshort, arg2 GLshort, arg3 GLshort, arg4 GLshort) 
{
	C.glVertexStream4sATI(C.GLenum(arg0), C.GLshort(arg1), C.GLshort(arg2), C.GLshort(arg3), C.GLshort(arg4));
}

//extern void glVertexStream4svATI (GLenum, const GLshort *)
func VertexStream4svATI(arg0 GLenum, arg1 *GLshort) 
{
	C.glVertexStream4svATI(C.GLenum(arg0), (*C.GLshort)(arg1));
}

//extern void glVertexStream4iATI (GLenum, GLint, GLint, GLint, GLint)
func VertexStream4iATI(arg0 GLenum, arg1 GLint, arg2 GLint, arg3 GLint, arg4 GLint) 
{
	C.glVertexStream4iATI(C.GLenum(arg0), C.GLint(arg1), C.GLint(arg2), C.GLint(arg3), C.GLint(arg4));
}

//extern void glVertexStream4ivATI (GLenum, const GLint *)
func VertexStream4ivATI(arg0 GLenum, arg1 *GLint) 
{
	C.glVertexStream4ivATI(C.GLenum(arg0), (*C.GLint)(arg1));
}

//extern void glVertexStream4fATI (GLenum, GLfloat, GLfloat, GLfloat, GLfloat)
func VertexStream4fATI(arg0 GLenum, arg1 GLfloat, arg2 GLfloat, arg3 GLfloat, arg4 GLfloat) 
{
	C.glVertexStream4fATI(C.GLenum(arg0), C.GLfloat(arg1), C.GLfloat(arg2), C.GLfloat(arg3), C.GLfloat(arg4));
}

//extern void glVertexStream4fvATI (GLenum, const GLfloat *)
func VertexStream4fvATI(arg0 GLenum, arg1 *GLfloat) 
{
	C.glVertexStream4fvATI(C.GLenum(arg0), (*C.GLfloat)(arg1));
}

//extern void glVertexStream4dATI (GLenum, GLdouble, GLdouble, GLdouble, GLdouble)
func VertexStream4dATI(arg0 GLenum, arg1 GLdouble, arg2 GLdouble, arg3 GLdouble, arg4 GLdouble) 
{
	C.glVertexStream4dATI(C.GLenum(arg0), C.GLdouble(arg1), C.GLdouble(arg2), C.GLdouble(arg3), C.GLdouble(arg4));
}

//extern void glVertexStream4dvATI (GLenum, const GLdouble *)
func VertexStream4dvATI(arg0 GLenum, arg1 *GLdouble) 
{
	C.glVertexStream4dvATI(C.GLenum(arg0), (*C.GLdouble)(arg1));
}

//extern void glNormalStream3bATI (GLenum, GLbyte, GLbyte, GLbyte)
func NormalStream3bATI(arg0 GLenum, arg1 GLbyte, arg2 GLbyte, arg3 GLbyte) 
{
	C.glNormalStream3bATI(C.GLenum(arg0), C.GLbyte(arg1), C.GLbyte(arg2), C.GLbyte(arg3));
}

//extern void glNormalStream3bvATI (GLenum, const GLbyte *)
func NormalStream3bvATI(arg0 GLenum, arg1 *GLbyte) 
{
	C.glNormalStream3bvATI(C.GLenum(arg0), (*C.GLbyte)(arg1));
}

//extern void glNormalStream3sATI (GLenum, GLshort, GLshort, GLshort)
func NormalStream3sATI(arg0 GLenum, arg1 GLshort, arg2 GLshort, arg3 GLshort) 
{
	C.glNormalStream3sATI(C.GLenum(arg0), C.GLshort(arg1), C.GLshort(arg2), C.GLshort(arg3));
}

//extern void glNormalStream3svATI (GLenum, const GLshort *)
func NormalStream3svATI(arg0 GLenum, arg1 *GLshort) 
{
	C.glNormalStream3svATI(C.GLenum(arg0), (*C.GLshort)(arg1));
}

//extern void glNormalStream3iATI (GLenum, GLint, GLint, GLint)
func NormalStream3iATI(arg0 GLenum, arg1 GLint, arg2 GLint, arg3 GLint) 
{
	C.glNormalStream3iATI(C.GLenum(arg0), C.GLint(arg1), C.GLint(arg2), C.GLint(arg3));
}

//extern void glNormalStream3ivATI (GLenum, const GLint *)
func NormalStream3ivATI(arg0 GLenum, arg1 *GLint) 
{
	C.glNormalStream3ivATI(C.GLenum(arg0), (*C.GLint)(arg1));
}

//extern void glNormalStream3fATI (GLenum, GLfloat, GLfloat, GLfloat)
func NormalStream3fATI(arg0 GLenum, arg1 GLfloat, arg2 GLfloat, arg3 GLfloat) 
{
	C.glNormalStream3fATI(C.GLenum(arg0), C.GLfloat(arg1), C.GLfloat(arg2), C.GLfloat(arg3));
}

//extern void glNormalStream3fvATI (GLenum, const GLfloat *)
func NormalStream3fvATI(arg0 GLenum, arg1 *GLfloat) 
{
	C.glNormalStream3fvATI(C.GLenum(arg0), (*C.GLfloat)(arg1));
}

//extern void glNormalStream3dATI (GLenum, GLdouble, GLdouble, GLdouble)
func NormalStream3dATI(arg0 GLenum, arg1 GLdouble, arg2 GLdouble, arg3 GLdouble) 
{
	C.glNormalStream3dATI(C.GLenum(arg0), C.GLdouble(arg1), C.GLdouble(arg2), C.GLdouble(arg3));
}

//extern void glNormalStream3dvATI (GLenum, const GLdouble *)
func NormalStream3dvATI(arg0 GLenum, arg1 *GLdouble) 
{
	C.glNormalStream3dvATI(C.GLenum(arg0), (*C.GLdouble)(arg1));
}

//extern void glClientActiveVertexStreamATI (GLenum)
func ClientActiveVertexStreamATI(arg0 GLenum) 
{
	C.glClientActiveVertexStreamATI(C.GLenum(arg0));
}

//extern void glVertexBlendEnviATI (GLenum, GLint)
func VertexBlendEnviATI(arg0 GLenum, arg1 GLint) 
{
	C.glVertexBlendEnviATI(C.GLenum(arg0), C.GLint(arg1));
}

//extern void glVertexBlendEnvfATI (GLenum, GLfloat)
func VertexBlendEnvfATI(arg0 GLenum, arg1 GLfloat) 
{
	C.glVertexBlendEnvfATI(C.GLenum(arg0), C.GLfloat(arg1));
}

//extern void glElementPointerATI (GLenum, const GLvoid *)
func ElementPointerATI(arg0 GLenum, arg1 unsafe.Pointer) 
{
	C.glElementPointerATI(C.GLenum(arg0), arg1);
}

//extern void glDrawElementArrayATI (GLenum, GLsizei)
func DrawElementArrayATI(arg0 GLenum, arg1 GLsizei) 
{
	C.glDrawElementArrayATI(C.GLenum(arg0), C.GLsizei(arg1));
}

//extern void glDrawRangeElementArrayATI (GLenum, GLuint, GLuint, GLsizei)
func DrawRangeElementArrayATI(arg0 GLenum, arg1 GLuint, arg2 GLuint, arg3 GLsizei) 
{
	C.glDrawRangeElementArrayATI(C.GLenum(arg0), C.GLuint(arg1), C.GLuint(arg2), C.GLsizei(arg3));
}

//extern void glDrawMeshArraysSUN (GLenum, GLint, GLsizei, GLsizei)
func DrawMeshArraysSUN(arg0 GLenum, arg1 GLint, arg2 GLsizei, arg3 GLsizei) 
{
	C.glDrawMeshArraysSUN(C.GLenum(arg0), C.GLint(arg1), C.GLsizei(arg2), C.GLsizei(arg3));
}

//extern void glGenOcclusionQueriesNV (GLsizei, GLuint *)
func GenOcclusionQueriesNV(arg0 GLsizei, arg1 *GLuint) 
{
	C.glGenOcclusionQueriesNV(C.GLsizei(arg0), (*C.GLuint)(arg1));
}

//extern void glDeleteOcclusionQueriesNV (GLsizei, const GLuint *)
func DeleteOcclusionQueriesNV(arg0 GLsizei, arg1 *GLuint) 
{
	C.glDeleteOcclusionQueriesNV(C.GLsizei(arg0), (*C.GLuint)(arg1));
}

//extern GLboolean glIsOcclusionQueryNV (GLuint)
func IsOcclusionQueryNV(arg0 GLuint) GLboolean
{
	return GLboolean(C.glIsOcclusionQueryNV(C.GLuint(arg0)));
}

//extern void glBeginOcclusionQueryNV (GLuint)
func BeginOcclusionQueryNV(arg0 GLuint) 
{
	C.glBeginOcclusionQueryNV(C.GLuint(arg0));
}

//extern void glEndOcclusionQueryNV (void)
func EndOcclusionQueryNV() 
{
	C.glEndOcclusionQueryNV();
}

//extern void glGetOcclusionQueryivNV (GLuint, GLenum, GLint *)
func GetOcclusionQueryivNV(arg0 GLuint, arg1 GLenum, arg2 *GLint) 
{
	C.glGetOcclusionQueryivNV(C.GLuint(arg0), C.GLenum(arg1), (*C.GLint)(arg2));
}

//extern void glGetOcclusionQueryuivNV (GLuint, GLenum, GLuint *)
func GetOcclusionQueryuivNV(arg0 GLuint, arg1 GLenum, arg2 *GLuint) 
{
	C.glGetOcclusionQueryuivNV(C.GLuint(arg0), C.GLenum(arg1), (*C.GLuint)(arg2));
}

//extern void glPointParameteriNV (GLenum, GLint)
func PointParameteriNV(arg0 GLenum, arg1 GLint) 
{
	C.glPointParameteriNV(C.GLenum(arg0), C.GLint(arg1));
}

//extern void glPointParameterivNV (GLenum, const GLint *)
func PointParameterivNV(arg0 GLenum, arg1 *GLint) 
{
	C.glPointParameterivNV(C.GLenum(arg0), (*C.GLint)(arg1));
}

//extern void glActiveStencilFaceEXT (GLenum)
func ActiveStencilFaceEXT(arg0 GLenum) 
{
	C.glActiveStencilFaceEXT(C.GLenum(arg0));
}

//extern void glElementPointerAPPLE (GLenum, const GLvoid *)
func ElementPointerAPPLE(arg0 GLenum, arg1 unsafe.Pointer) 
{
	C.glElementPointerAPPLE(C.GLenum(arg0), arg1);
}

//extern void glDrawElementArrayAPPLE (GLenum, GLint, GLsizei)
func DrawElementArrayAPPLE(arg0 GLenum, arg1 GLint, arg2 GLsizei) 
{
	C.glDrawElementArrayAPPLE(C.GLenum(arg0), C.GLint(arg1), C.GLsizei(arg2));
}

//extern void glDrawRangeElementArrayAPPLE (GLenum, GLuint, GLuint, GLint, GLsizei)
func DrawRangeElementArrayAPPLE(arg0 GLenum, arg1 GLuint, arg2 GLuint, arg3 GLint, arg4 GLsizei) 
{
	C.glDrawRangeElementArrayAPPLE(C.GLenum(arg0), C.GLuint(arg1), C.GLuint(arg2), C.GLint(arg3), C.GLsizei(arg4));
}

//extern void glMultiDrawElementArrayAPPLE (GLenum, const GLint *, const GLsizei *, GLsizei)
func MultiDrawElementArrayAPPLE(arg0 GLenum, arg1 *GLint, arg2 *GLsizei, arg3 GLsizei) 
{
	C.glMultiDrawElementArrayAPPLE(C.GLenum(arg0), (*C.GLint)(arg1), (*C.GLsizei)(arg2), C.GLsizei(arg3));
}

//extern void glMultiDrawRangeElementArrayAPPLE (GLenum, GLuint, GLuint, const GLint *, const GLsizei *, GLsizei)
func MultiDrawRangeElementArrayAPPLE(arg0 GLenum, arg1 GLuint, arg2 GLuint, arg3 *GLint, arg4 *GLsizei, arg5 GLsizei) 
{
	C.glMultiDrawRangeElementArrayAPPLE(C.GLenum(arg0), C.GLuint(arg1), C.GLuint(arg2), (*C.GLint)(arg3), (*C.GLsizei)(arg4), C.GLsizei(arg5));
}

//extern void glGenFencesAPPLE (GLsizei, GLuint *)
func GenFencesAPPLE(arg0 GLsizei, arg1 *GLuint) 
{
	C.glGenFencesAPPLE(C.GLsizei(arg0), (*C.GLuint)(arg1));
}

//extern void glDeleteFencesAPPLE (GLsizei, const GLuint *)
func DeleteFencesAPPLE(arg0 GLsizei, arg1 *GLuint) 
{
	C.glDeleteFencesAPPLE(C.GLsizei(arg0), (*C.GLuint)(arg1));
}

//extern void glSetFenceAPPLE (GLuint)
func SetFenceAPPLE(arg0 GLuint) 
{
	C.glSetFenceAPPLE(C.GLuint(arg0));
}

//extern GLboolean glIsFenceAPPLE (GLuint)
func IsFenceAPPLE(arg0 GLuint) GLboolean
{
	return GLboolean(C.glIsFenceAPPLE(C.GLuint(arg0)));
}

//extern GLboolean glTestFenceAPPLE (GLuint)
func TestFenceAPPLE(arg0 GLuint) GLboolean
{
	return GLboolean(C.glTestFenceAPPLE(C.GLuint(arg0)));
}

//extern void glFinishFenceAPPLE (GLuint)
func FinishFenceAPPLE(arg0 GLuint) 
{
	C.glFinishFenceAPPLE(C.GLuint(arg0));
}

//extern GLboolean glTestObjectAPPLE (GLenum, GLuint)
func TestObjectAPPLE(arg0 GLenum, arg1 GLuint) GLboolean
{
	return GLboolean(C.glTestObjectAPPLE(C.GLenum(arg0), C.GLuint(arg1)));
}

//extern void glFinishObjectAPPLE (GLenum, GLint)
func FinishObjectAPPLE(arg0 GLenum, arg1 GLint) 
{
	C.glFinishObjectAPPLE(C.GLenum(arg0), C.GLint(arg1));
}

//extern void glBindVertexArrayAPPLE (GLuint)
func BindVertexArrayAPPLE(arg0 GLuint) 
{
	C.glBindVertexArrayAPPLE(C.GLuint(arg0));
}

//extern void glDeleteVertexArraysAPPLE (GLsizei, const GLuint *)
func DeleteVertexArraysAPPLE(arg0 GLsizei, arg1 *GLuint) 
{
	C.glDeleteVertexArraysAPPLE(C.GLsizei(arg0), (*C.GLuint)(arg1));
}

//extern void glGenVertexArraysAPPLE (GLsizei, const GLuint *)
func GenVertexArraysAPPLE(arg0 GLsizei, arg1 *GLuint) 
{
	C.glGenVertexArraysAPPLE(C.GLsizei(arg0), (*C.GLuint)(arg1));
}

//extern GLboolean glIsVertexArrayAPPLE (GLuint)
func IsVertexArrayAPPLE(arg0 GLuint) GLboolean
{
	return GLboolean(C.glIsVertexArrayAPPLE(C.GLuint(arg0)));
}

//extern void glVertexArrayRangeAPPLE (GLsizei, GLvoid *)
func VertexArrayRangeAPPLE(arg0 GLsizei, arg1 unsafe.Pointer) 
{
	C.glVertexArrayRangeAPPLE(C.GLsizei(arg0), arg1);
}

//extern void glFlushVertexArrayRangeAPPLE (GLsizei, GLvoid *)
func FlushVertexArrayRangeAPPLE(arg0 GLsizei, arg1 unsafe.Pointer) 
{
	C.glFlushVertexArrayRangeAPPLE(C.GLsizei(arg0), arg1);
}

//extern void glVertexArrayParameteriAPPLE (GLenum, GLint)
func VertexArrayParameteriAPPLE(arg0 GLenum, arg1 GLint) 
{
	C.glVertexArrayParameteriAPPLE(C.GLenum(arg0), C.GLint(arg1));
}

//extern void glDrawBuffersATI (GLsizei, const GLenum *)
func DrawBuffersATI(arg0 GLsizei, arg1 *GLenum) 
{
	C.glDrawBuffersATI(C.GLsizei(arg0), (*C.GLenum)(arg1));
}

//extern void glProgramNamedParameter4fNV (GLuint, GLsizei, const GLubyte *, GLfloat, GLfloat, GLfloat, GLfloat)
func ProgramNamedParameter4fNV(arg0 GLuint, arg1 GLsizei, arg2 *GLubyte, arg3 GLfloat, arg4 GLfloat, arg5 GLfloat, arg6 GLfloat) 
{
	C.glProgramNamedParameter4fNV(C.GLuint(arg0), C.GLsizei(arg1), (*C.GLubyte)(arg2), C.GLfloat(arg3), C.GLfloat(arg4), C.GLfloat(arg5), C.GLfloat(arg6));
}

//extern void glProgramNamedParameter4dNV (GLuint, GLsizei, const GLubyte *, GLdouble, GLdouble, GLdouble, GLdouble)
func ProgramNamedParameter4dNV(arg0 GLuint, arg1 GLsizei, arg2 *GLubyte, arg3 GLdouble, arg4 GLdouble, arg5 GLdouble, arg6 GLdouble) 
{
	C.glProgramNamedParameter4dNV(C.GLuint(arg0), C.GLsizei(arg1), (*C.GLubyte)(arg2), C.GLdouble(arg3), C.GLdouble(arg4), C.GLdouble(arg5), C.GLdouble(arg6));
}

//extern void glProgramNamedParameter4fvNV (GLuint, GLsizei, const GLubyte *, const GLfloat *)
func ProgramNamedParameter4fvNV(arg0 GLuint, arg1 GLsizei, arg2 *GLubyte, arg3 *GLfloat) 
{
	C.glProgramNamedParameter4fvNV(C.GLuint(arg0), C.GLsizei(arg1), (*C.GLubyte)(arg2), (*C.GLfloat)(arg3));
}

//extern void glProgramNamedParameter4dvNV (GLuint, GLsizei, const GLubyte *, const GLdouble *)
func ProgramNamedParameter4dvNV(arg0 GLuint, arg1 GLsizei, arg2 *GLubyte, arg3 *GLdouble) 
{
	C.glProgramNamedParameter4dvNV(C.GLuint(arg0), C.GLsizei(arg1), (*C.GLubyte)(arg2), (*C.GLdouble)(arg3));
}

//extern void glGetProgramNamedParameterfvNV (GLuint, GLsizei, const GLubyte *, GLfloat *)
func GetProgramNamedParameterfvNV(arg0 GLuint, arg1 GLsizei, arg2 *GLubyte, arg3 *GLfloat) 
{
	C.glGetProgramNamedParameterfvNV(C.GLuint(arg0), C.GLsizei(arg1), (*C.GLubyte)(arg2), (*C.GLfloat)(arg3));
}

//extern void glGetProgramNamedParameterdvNV (GLuint, GLsizei, const GLubyte *, GLdouble *)
func GetProgramNamedParameterdvNV(arg0 GLuint, arg1 GLsizei, arg2 *GLubyte, arg3 *GLdouble) 
{
	C.glGetProgramNamedParameterdvNV(C.GLuint(arg0), C.GLsizei(arg1), (*C.GLubyte)(arg2), (*C.GLdouble)(arg3));
}

//extern void glVertex2hNV (GLhalfNV, GLhalfNV)
func Vertex2hNV(arg0 GLhalfNV, arg1 GLhalfNV) 
{
	C.glVertex2hNV(C.GLhalfNV(arg0), C.GLhalfNV(arg1));
}

//extern void glVertex2hvNV (const GLhalfNV *)
func Vertex2hvNV(arg0 *GLhalfNV) 
{
	C.glVertex2hvNV((*C.GLhalfNV)(arg0));
}

//extern void glVertex3hNV (GLhalfNV, GLhalfNV, GLhalfNV)
func Vertex3hNV(arg0 GLhalfNV, arg1 GLhalfNV, arg2 GLhalfNV) 
{
	C.glVertex3hNV(C.GLhalfNV(arg0), C.GLhalfNV(arg1), C.GLhalfNV(arg2));
}

//extern void glVertex3hvNV (const GLhalfNV *)
func Vertex3hvNV(arg0 *GLhalfNV) 
{
	C.glVertex3hvNV((*C.GLhalfNV)(arg0));
}

//extern void glVertex4hNV (GLhalfNV, GLhalfNV, GLhalfNV, GLhalfNV)
func Vertex4hNV(arg0 GLhalfNV, arg1 GLhalfNV, arg2 GLhalfNV, arg3 GLhalfNV) 
{
	C.glVertex4hNV(C.GLhalfNV(arg0), C.GLhalfNV(arg1), C.GLhalfNV(arg2), C.GLhalfNV(arg3));
}

//extern void glVertex4hvNV (const GLhalfNV *)
func Vertex4hvNV(arg0 *GLhalfNV) 
{
	C.glVertex4hvNV((*C.GLhalfNV)(arg0));
}

//extern void glNormal3hNV (GLhalfNV, GLhalfNV, GLhalfNV)
func Normal3hNV(arg0 GLhalfNV, arg1 GLhalfNV, arg2 GLhalfNV) 
{
	C.glNormal3hNV(C.GLhalfNV(arg0), C.GLhalfNV(arg1), C.GLhalfNV(arg2));
}

//extern void glNormal3hvNV (const GLhalfNV *)
func Normal3hvNV(arg0 *GLhalfNV) 
{
	C.glNormal3hvNV((*C.GLhalfNV)(arg0));
}

//extern void glColor3hNV (GLhalfNV, GLhalfNV, GLhalfNV)
func Color3hNV(arg0 GLhalfNV, arg1 GLhalfNV, arg2 GLhalfNV) 
{
	C.glColor3hNV(C.GLhalfNV(arg0), C.GLhalfNV(arg1), C.GLhalfNV(arg2));
}

//extern void glColor3hvNV (const GLhalfNV *)
func Color3hvNV(arg0 *GLhalfNV) 
{
	C.glColor3hvNV((*C.GLhalfNV)(arg0));
}

//extern void glColor4hNV (GLhalfNV, GLhalfNV, GLhalfNV, GLhalfNV)
func Color4hNV(arg0 GLhalfNV, arg1 GLhalfNV, arg2 GLhalfNV, arg3 GLhalfNV) 
{
	C.glColor4hNV(C.GLhalfNV(arg0), C.GLhalfNV(arg1), C.GLhalfNV(arg2), C.GLhalfNV(arg3));
}

//extern void glColor4hvNV (const GLhalfNV *)
func Color4hvNV(arg0 *GLhalfNV) 
{
	C.glColor4hvNV((*C.GLhalfNV)(arg0));
}

//extern void glTexCoord1hNV (GLhalfNV)
func TexCoord1hNV(arg0 GLhalfNV) 
{
	C.glTexCoord1hNV(C.GLhalfNV(arg0));
}

//extern void glTexCoord1hvNV (const GLhalfNV *)
func TexCoord1hvNV(arg0 *GLhalfNV) 
{
	C.glTexCoord1hvNV((*C.GLhalfNV)(arg0));
}

//extern void glTexCoord2hNV (GLhalfNV, GLhalfNV)
func TexCoord2hNV(arg0 GLhalfNV, arg1 GLhalfNV) 
{
	C.glTexCoord2hNV(C.GLhalfNV(arg0), C.GLhalfNV(arg1));
}

//extern void glTexCoord2hvNV (const GLhalfNV *)
func TexCoord2hvNV(arg0 *GLhalfNV) 
{
	C.glTexCoord2hvNV((*C.GLhalfNV)(arg0));
}

//extern void glTexCoord3hNV (GLhalfNV, GLhalfNV, GLhalfNV)
func TexCoord3hNV(arg0 GLhalfNV, arg1 GLhalfNV, arg2 GLhalfNV) 
{
	C.glTexCoord3hNV(C.GLhalfNV(arg0), C.GLhalfNV(arg1), C.GLhalfNV(arg2));
}

//extern void glTexCoord3hvNV (const GLhalfNV *)
func TexCoord3hvNV(arg0 *GLhalfNV) 
{
	C.glTexCoord3hvNV((*C.GLhalfNV)(arg0));
}

//extern void glTexCoord4hNV (GLhalfNV, GLhalfNV, GLhalfNV, GLhalfNV)
func TexCoord4hNV(arg0 GLhalfNV, arg1 GLhalfNV, arg2 GLhalfNV, arg3 GLhalfNV) 
{
	C.glTexCoord4hNV(C.GLhalfNV(arg0), C.GLhalfNV(arg1), C.GLhalfNV(arg2), C.GLhalfNV(arg3));
}

//extern void glTexCoord4hvNV (const GLhalfNV *)
func TexCoord4hvNV(arg0 *GLhalfNV) 
{
	C.glTexCoord4hvNV((*C.GLhalfNV)(arg0));
}

//extern void glMultiTexCoord1hNV (GLenum, GLhalfNV)
func MultiTexCoord1hNV(arg0 GLenum, arg1 GLhalfNV) 
{
	C.glMultiTexCoord1hNV(C.GLenum(arg0), C.GLhalfNV(arg1));
}

//extern void glMultiTexCoord1hvNV (GLenum, const GLhalfNV *)
func MultiTexCoord1hvNV(arg0 GLenum, arg1 *GLhalfNV) 
{
	C.glMultiTexCoord1hvNV(C.GLenum(arg0), (*C.GLhalfNV)(arg1));
}

//extern void glMultiTexCoord2hNV (GLenum, GLhalfNV, GLhalfNV)
func MultiTexCoord2hNV(arg0 GLenum, arg1 GLhalfNV, arg2 GLhalfNV) 
{
	C.glMultiTexCoord2hNV(C.GLenum(arg0), C.GLhalfNV(arg1), C.GLhalfNV(arg2));
}

//extern void glMultiTexCoord2hvNV (GLenum, const GLhalfNV *)
func MultiTexCoord2hvNV(arg0 GLenum, arg1 *GLhalfNV) 
{
	C.glMultiTexCoord2hvNV(C.GLenum(arg0), (*C.GLhalfNV)(arg1));
}

//extern void glMultiTexCoord3hNV (GLenum, GLhalfNV, GLhalfNV, GLhalfNV)
func MultiTexCoord3hNV(arg0 GLenum, arg1 GLhalfNV, arg2 GLhalfNV, arg3 GLhalfNV) 
{
	C.glMultiTexCoord3hNV(C.GLenum(arg0), C.GLhalfNV(arg1), C.GLhalfNV(arg2), C.GLhalfNV(arg3));
}

//extern void glMultiTexCoord3hvNV (GLenum, const GLhalfNV *)
func MultiTexCoord3hvNV(arg0 GLenum, arg1 *GLhalfNV) 
{
	C.glMultiTexCoord3hvNV(C.GLenum(arg0), (*C.GLhalfNV)(arg1));
}

//extern void glMultiTexCoord4hNV (GLenum, GLhalfNV, GLhalfNV, GLhalfNV, GLhalfNV)
func MultiTexCoord4hNV(arg0 GLenum, arg1 GLhalfNV, arg2 GLhalfNV, arg3 GLhalfNV, arg4 GLhalfNV) 
{
	C.glMultiTexCoord4hNV(C.GLenum(arg0), C.GLhalfNV(arg1), C.GLhalfNV(arg2), C.GLhalfNV(arg3), C.GLhalfNV(arg4));
}

//extern void glMultiTexCoord4hvNV (GLenum, const GLhalfNV *)
func MultiTexCoord4hvNV(arg0 GLenum, arg1 *GLhalfNV) 
{
	C.glMultiTexCoord4hvNV(C.GLenum(arg0), (*C.GLhalfNV)(arg1));
}

//extern void glFogCoordhNV (GLhalfNV)
func FogCoordhNV(arg0 GLhalfNV) 
{
	C.glFogCoordhNV(C.GLhalfNV(arg0));
}

//extern void glFogCoordhvNV (const GLhalfNV *)
func FogCoordhvNV(arg0 *GLhalfNV) 
{
	C.glFogCoordhvNV((*C.GLhalfNV)(arg0));
}

//extern void glSecondaryColor3hNV (GLhalfNV, GLhalfNV, GLhalfNV)
func SecondaryColor3hNV(arg0 GLhalfNV, arg1 GLhalfNV, arg2 GLhalfNV) 
{
	C.glSecondaryColor3hNV(C.GLhalfNV(arg0), C.GLhalfNV(arg1), C.GLhalfNV(arg2));
}

//extern void glSecondaryColor3hvNV (const GLhalfNV *)
func SecondaryColor3hvNV(arg0 *GLhalfNV) 
{
	C.glSecondaryColor3hvNV((*C.GLhalfNV)(arg0));
}

//extern void glVertexWeighthNV (GLhalfNV)
func VertexWeighthNV(arg0 GLhalfNV) 
{
	C.glVertexWeighthNV(C.GLhalfNV(arg0));
}

//extern void glVertexWeighthvNV (const GLhalfNV *)
func VertexWeighthvNV(arg0 *GLhalfNV) 
{
	C.glVertexWeighthvNV((*C.GLhalfNV)(arg0));
}

//extern void glVertexAttrib1hNV (GLuint, GLhalfNV)
func VertexAttrib1hNV(arg0 GLuint, arg1 GLhalfNV) 
{
	C.glVertexAttrib1hNV(C.GLuint(arg0), C.GLhalfNV(arg1));
}

//extern void glVertexAttrib1hvNV (GLuint, const GLhalfNV *)
func VertexAttrib1hvNV(arg0 GLuint, arg1 *GLhalfNV) 
{
	C.glVertexAttrib1hvNV(C.GLuint(arg0), (*C.GLhalfNV)(arg1));
}

//extern void glVertexAttrib2hNV (GLuint, GLhalfNV, GLhalfNV)
func VertexAttrib2hNV(arg0 GLuint, arg1 GLhalfNV, arg2 GLhalfNV) 
{
	C.glVertexAttrib2hNV(C.GLuint(arg0), C.GLhalfNV(arg1), C.GLhalfNV(arg2));
}

//extern void glVertexAttrib2hvNV (GLuint, const GLhalfNV *)
func VertexAttrib2hvNV(arg0 GLuint, arg1 *GLhalfNV) 
{
	C.glVertexAttrib2hvNV(C.GLuint(arg0), (*C.GLhalfNV)(arg1));
}

//extern void glVertexAttrib3hNV (GLuint, GLhalfNV, GLhalfNV, GLhalfNV)
func VertexAttrib3hNV(arg0 GLuint, arg1 GLhalfNV, arg2 GLhalfNV, arg3 GLhalfNV) 
{
	C.glVertexAttrib3hNV(C.GLuint(arg0), C.GLhalfNV(arg1), C.GLhalfNV(arg2), C.GLhalfNV(arg3));
}

//extern void glVertexAttrib3hvNV (GLuint, const GLhalfNV *)
func VertexAttrib3hvNV(arg0 GLuint, arg1 *GLhalfNV) 
{
	C.glVertexAttrib3hvNV(C.GLuint(arg0), (*C.GLhalfNV)(arg1));
}

//extern void glVertexAttrib4hNV (GLuint, GLhalfNV, GLhalfNV, GLhalfNV, GLhalfNV)
func VertexAttrib4hNV(arg0 GLuint, arg1 GLhalfNV, arg2 GLhalfNV, arg3 GLhalfNV, arg4 GLhalfNV) 
{
	C.glVertexAttrib4hNV(C.GLuint(arg0), C.GLhalfNV(arg1), C.GLhalfNV(arg2), C.GLhalfNV(arg3), C.GLhalfNV(arg4));
}

//extern void glVertexAttrib4hvNV (GLuint, const GLhalfNV *)
func VertexAttrib4hvNV(arg0 GLuint, arg1 *GLhalfNV) 
{
	C.glVertexAttrib4hvNV(C.GLuint(arg0), (*C.GLhalfNV)(arg1));
}

//extern void glVertexAttribs1hvNV (GLuint, GLsizei, const GLhalfNV *)
func VertexAttribs1hvNV(arg0 GLuint, arg1 GLsizei, arg2 *GLhalfNV) 
{
	C.glVertexAttribs1hvNV(C.GLuint(arg0), C.GLsizei(arg1), (*C.GLhalfNV)(arg2));
}

//extern void glVertexAttribs2hvNV (GLuint, GLsizei, const GLhalfNV *)
func VertexAttribs2hvNV(arg0 GLuint, arg1 GLsizei, arg2 *GLhalfNV) 
{
	C.glVertexAttribs2hvNV(C.GLuint(arg0), C.GLsizei(arg1), (*C.GLhalfNV)(arg2));
}

//extern void glVertexAttribs3hvNV (GLuint, GLsizei, const GLhalfNV *)
func VertexAttribs3hvNV(arg0 GLuint, arg1 GLsizei, arg2 *GLhalfNV) 
{
	C.glVertexAttribs3hvNV(C.GLuint(arg0), C.GLsizei(arg1), (*C.GLhalfNV)(arg2));
}

//extern void glVertexAttribs4hvNV (GLuint, GLsizei, const GLhalfNV *)
func VertexAttribs4hvNV(arg0 GLuint, arg1 GLsizei, arg2 *GLhalfNV) 
{
	C.glVertexAttribs4hvNV(C.GLuint(arg0), C.GLsizei(arg1), (*C.GLhalfNV)(arg2));
}

//extern void glPixelDataRangeNV (GLenum, GLsizei, GLvoid *)
func PixelDataRangeNV(arg0 GLenum, arg1 GLsizei, arg2 unsafe.Pointer) 
{
	C.glPixelDataRangeNV(C.GLenum(arg0), C.GLsizei(arg1), arg2);
}

//extern void glFlushPixelDataRangeNV (GLenum)
func FlushPixelDataRangeNV(arg0 GLenum) 
{
	C.glFlushPixelDataRangeNV(C.GLenum(arg0));
}

//extern void glPrimitiveRestartNV (void)
func PrimitiveRestartNV() 
{
	C.glPrimitiveRestartNV();
}

//extern void glPrimitiveRestartIndexNV (GLuint)
func PrimitiveRestartIndexNV(arg0 GLuint) 
{
	C.glPrimitiveRestartIndexNV(C.GLuint(arg0));
}

//extern GLvoid* glMapObjectBufferATI (GLuint)
func MapObjectBufferATI(arg0 GLuint) unsafe.Pointer
{
	return unsafe.Pointer(C.glMapObjectBufferATI(C.GLuint(arg0)));
}

//extern void glUnmapObjectBufferATI (GLuint)
func UnmapObjectBufferATI(arg0 GLuint) 
{
	C.glUnmapObjectBufferATI(C.GLuint(arg0));
}

//extern void glStencilOpSeparateATI (GLenum, GLenum, GLenum, GLenum)
func StencilOpSeparateATI(arg0 GLenum, arg1 GLenum, arg2 GLenum, arg3 GLenum) 
{
	C.glStencilOpSeparateATI(C.GLenum(arg0), C.GLenum(arg1), C.GLenum(arg2), C.GLenum(arg3));
}

//extern void glStencilFuncSeparateATI (GLenum, GLenum, GLint, GLuint)
func StencilFuncSeparateATI(arg0 GLenum, arg1 GLenum, arg2 GLint, arg3 GLuint) 
{
	C.glStencilFuncSeparateATI(C.GLenum(arg0), C.GLenum(arg1), C.GLint(arg2), C.GLuint(arg3));
}

//extern void glVertexAttribArrayObjectATI (GLuint, GLint, GLenum, GLboolean, GLsizei, GLuint, GLuint)
func VertexAttribArrayObjectATI(arg0 GLuint, arg1 GLint, arg2 GLenum, arg3 GLboolean, arg4 GLsizei, arg5 GLuint, arg6 GLuint) 
{
	C.glVertexAttribArrayObjectATI(C.GLuint(arg0), C.GLint(arg1), C.GLenum(arg2), C.GLboolean(arg3), C.GLsizei(arg4), C.GLuint(arg5), C.GLuint(arg6));
}

//extern void glGetVertexAttribArrayObjectfvATI (GLuint, GLenum, GLfloat *)
func GetVertexAttribArrayObjectfvATI(arg0 GLuint, arg1 GLenum, arg2 *GLfloat) 
{
	C.glGetVertexAttribArrayObjectfvATI(C.GLuint(arg0), C.GLenum(arg1), (*C.GLfloat)(arg2));
}

//extern void glGetVertexAttribArrayObjectivATI (GLuint, GLenum, GLint *)
func GetVertexAttribArrayObjectivATI(arg0 GLuint, arg1 GLenum, arg2 *GLint) 
{
	C.glGetVertexAttribArrayObjectivATI(C.GLuint(arg0), C.GLenum(arg1), (*C.GLint)(arg2));
}

//extern void glDepthBoundsEXT (GLclampd, GLclampd)
func DepthBoundsEXT(arg0 GLclampd, arg1 GLclampd) 
{
	C.glDepthBoundsEXT(C.GLclampd(arg0), C.GLclampd(arg1));
}

//extern void glBlendEquationSeparateEXT (GLenum, GLenum)
func BlendEquationSeparateEXT(arg0 GLenum, arg1 GLenum) 
{
	C.glBlendEquationSeparateEXT(C.GLenum(arg0), C.GLenum(arg1));
}

//extern GLboolean glIsRenderbufferEXT (GLuint)
func IsRenderbufferEXT(arg0 GLuint) GLboolean
{
	return GLboolean(C.glIsRenderbufferEXT(C.GLuint(arg0)));
}

//extern void glBindRenderbufferEXT (GLenum, GLuint)
func BindRenderbufferEXT(arg0 GLenum, arg1 GLuint) 
{
	C.glBindRenderbufferEXT(C.GLenum(arg0), C.GLuint(arg1));
}

//extern void glDeleteRenderbuffersEXT (GLsizei, const GLuint *)
func DeleteRenderbuffersEXT(arg0 GLsizei, arg1 *GLuint) 
{
	C.glDeleteRenderbuffersEXT(C.GLsizei(arg0), (*C.GLuint)(arg1));
}

//extern void glGenRenderbuffersEXT (GLsizei, GLuint *)
func GenRenderbuffersEXT(arg0 GLsizei, arg1 *GLuint) 
{
	C.glGenRenderbuffersEXT(C.GLsizei(arg0), (*C.GLuint)(arg1));
}

//extern void glRenderbufferStorageEXT (GLenum, GLenum, GLsizei, GLsizei)
func RenderbufferStorageEXT(arg0 GLenum, arg1 GLenum, arg2 GLsizei, arg3 GLsizei) 
{
	C.glRenderbufferStorageEXT(C.GLenum(arg0), C.GLenum(arg1), C.GLsizei(arg2), C.GLsizei(arg3));
}

//extern void glGetRenderbufferParameterivEXT (GLenum, GLenum, GLint *)
func GetRenderbufferParameterivEXT(arg0 GLenum, arg1 GLenum, arg2 *GLint) 
{
	C.glGetRenderbufferParameterivEXT(C.GLenum(arg0), C.GLenum(arg1), (*C.GLint)(arg2));
}

//extern GLboolean glIsFramebufferEXT (GLuint)
func IsFramebufferEXT(arg0 GLuint) GLboolean
{
	return GLboolean(C.glIsFramebufferEXT(C.GLuint(arg0)));
}

//extern void glBindFramebufferEXT (GLenum, GLuint)
func BindFramebufferEXT(arg0 GLenum, arg1 GLuint) 
{
	C.glBindFramebufferEXT(C.GLenum(arg0), C.GLuint(arg1));
}

//extern void glDeleteFramebuffersEXT (GLsizei, const GLuint *)
func DeleteFramebuffersEXT(arg0 GLsizei, arg1 *GLuint) 
{
	C.glDeleteFramebuffersEXT(C.GLsizei(arg0), (*C.GLuint)(arg1));
}

//extern void glGenFramebuffersEXT (GLsizei, GLuint *)
func GenFramebuffersEXT(arg0 GLsizei, arg1 *GLuint) 
{
	C.glGenFramebuffersEXT(C.GLsizei(arg0), (*C.GLuint)(arg1));
}

//extern GLenum glCheckFramebufferStatusEXT (GLenum)
func CheckFramebufferStatusEXT(arg0 GLenum) GLenum
{
	return GLenum(C.glCheckFramebufferStatusEXT(C.GLenum(arg0)));
}

//extern void glFramebufferTexture1DEXT (GLenum, GLenum, GLenum, GLuint, GLint)
func FramebufferTexture1DEXT(arg0 GLenum, arg1 GLenum, arg2 GLenum, arg3 GLuint, arg4 GLint) 
{
	C.glFramebufferTexture1DEXT(C.GLenum(arg0), C.GLenum(arg1), C.GLenum(arg2), C.GLuint(arg3), C.GLint(arg4));
}

//extern void glFramebufferTexture2DEXT (GLenum, GLenum, GLenum, GLuint, GLint)
func FramebufferTexture2DEXT(arg0 GLenum, arg1 GLenum, arg2 GLenum, arg3 GLuint, arg4 GLint) 
{
	C.glFramebufferTexture2DEXT(C.GLenum(arg0), C.GLenum(arg1), C.GLenum(arg2), C.GLuint(arg3), C.GLint(arg4));
}

//extern void glFramebufferTexture3DEXT (GLenum, GLenum, GLenum, GLuint, GLint, GLint)
func FramebufferTexture3DEXT(arg0 GLenum, arg1 GLenum, arg2 GLenum, arg3 GLuint, arg4 GLint, arg5 GLint) 
{
	C.glFramebufferTexture3DEXT(C.GLenum(arg0), C.GLenum(arg1), C.GLenum(arg2), C.GLuint(arg3), C.GLint(arg4), C.GLint(arg5));
}

//extern void glFramebufferRenderbufferEXT (GLenum, GLenum, GLenum, GLuint)
func FramebufferRenderbufferEXT(arg0 GLenum, arg1 GLenum, arg2 GLenum, arg3 GLuint) 
{
	C.glFramebufferRenderbufferEXT(C.GLenum(arg0), C.GLenum(arg1), C.GLenum(arg2), C.GLuint(arg3));
}

//extern void glGetFramebufferAttachmentParameterivEXT (GLenum, GLenum, GLenum, GLint *)
func GetFramebufferAttachmentParameterivEXT(arg0 GLenum, arg1 GLenum, arg2 GLenum, arg3 *GLint) 
{
	C.glGetFramebufferAttachmentParameterivEXT(C.GLenum(arg0), C.GLenum(arg1), C.GLenum(arg2), (*C.GLint)(arg3));
}

//extern void glGenerateMipmapEXT (GLenum)
func GenerateMipmapEXT(arg0 GLenum) 
{
	C.glGenerateMipmapEXT(C.GLenum(arg0));
}

//extern void glStringMarkerGREMEDY (GLsizei, const GLvoid *)
func StringMarkerGREMEDY(arg0 GLsizei, arg1 unsafe.Pointer) 
{
	C.glStringMarkerGREMEDY(C.GLsizei(arg0), arg1);
}

//extern void glGetQueryObjecti64vEXT (GLuint id, GLenum pname, GLint64EXT *params)
func GetQueryObjecti64vEXT(id GLuint, pname GLenum, params *GLint64EXT) 
{
	C.glGetQueryObjecti64vEXT(C.GLuint(id), C.GLenum(pname), (*C.GLint64EXT)(params));
}

//extern void glGetQueryObjectui64vEXT (GLuint id, GLenum pname, GLuint64EXT *params)
func GetQueryObjectui64vEXT(id GLuint, pname GLenum, params *GLuint64EXT) 
{
	C.glGetQueryObjectui64vEXT(C.GLuint(id), C.GLenum(pname), (*C.GLuint64EXT)(params));
}

//extern void glTexBufferEXT (GLenum target, GLenum internalformat, GLuint buffer)
func TexBufferEXT(target GLenum, internalformat GLenum, buffer GLuint) 
{
	C.glTexBufferEXT(C.GLenum(target), C.GLenum(internalformat), C.GLuint(buffer));
}

//extern void glBeginTransformFeedbackNV (GLenum primitiveMode)
func BeginTransformFeedbackNV(primitiveMode GLenum) 
{
	C.glBeginTransformFeedbackNV(C.GLenum(primitiveMode));
}

//extern void glEndTransformFeedbackNV (void)
func EndTransformFeedbackNV() 
{
	C.glEndTransformFeedbackNV();
}

//extern void glTransformFeedbackAttribsNV (GLuint count, const GLint *attribs, GLenum bufferMode)
func TransformFeedbackAttribsNV(count GLuint, attribs *GLint, bufferMode GLenum) 
{
	C.glTransformFeedbackAttribsNV(C.GLuint(count), (*C.GLint)(attribs), C.GLenum(bufferMode));
}

//extern void glBindBufferRangeNV (GLenum target, GLuint index, GLuint buffer, GLintptr offset, GLsizeiptr size)
func BindBufferRangeNV(target GLenum, index GLuint, buffer GLuint, offset GLintptr, size GLsizeiptr) 
{
	C.glBindBufferRangeNV(C.GLenum(target), C.GLuint(index), C.GLuint(buffer), C.GLintptr(offset), C.GLsizeiptr(size));
}

//extern void glBindBufferOffsetNV (GLenum target, GLuint index, GLuint buffer, GLintptr offset)
func BindBufferOffsetNV(target GLenum, index GLuint, buffer GLuint, offset GLintptr) 
{
	C.glBindBufferOffsetNV(C.GLenum(target), C.GLuint(index), C.GLuint(buffer), C.GLintptr(offset));
}

//extern void glBindBufferBaseNV (GLenum target, GLuint index, GLuint buffer)
func BindBufferBaseNV(target GLenum, index GLuint, buffer GLuint) 
{
	C.glBindBufferBaseNV(C.GLenum(target), C.GLuint(index), C.GLuint(buffer));
}

//extern void glTransformFeedbackVaryingsNV (GLuint program, GLsizei count, const GLint *locations, GLenum bufferMode)
func TransformFeedbackVaryingsNV(program GLuint, count GLsizei, locations *GLint, bufferMode GLenum) 
{
	C.glTransformFeedbackVaryingsNV(C.GLuint(program), C.GLsizei(count), (*C.GLint)(locations), C.GLenum(bufferMode));
}

//extern void glActiveVaryingNV (GLuint program, const GLchar *name)
func ActiveVaryingNV(program GLuint, name string) 
{
	C.glActiveVaryingNV(C.GLuint(program), (*_C_GLchar)((C.CString)(name)));
}

//extern GLint glGetVaryingLocationNV (GLuint program, const GLchar *name)
func GetVaryingLocationNV(program GLuint, name string) GLint
{
	return GLint(C.glGetVaryingLocationNV(C.GLuint(program), (*_C_GLchar)((C.CString)(name))));
}

//extern void glGetActiveVaryingNV (GLuint program, GLuint index, GLsizei bufSize, GLsizei *length, GLsizei *size, GLenum *type, GLchar *name)
func GetActiveVaryingNV(program GLuint, index GLuint, bufSize GLsizei, length *GLsizei, size *GLsizei, type_ *GLenum, name string) 
{
	C.glGetActiveVaryingNV(C.GLuint(program), C.GLuint(index), C.GLsizei(bufSize), (*C.GLsizei)(length), (*C.GLsizei)(size), (*C.GLenum)(type_), (*_C_GLchar)((C.CString)(name)));
}

//extern void glGetTransformFeedbackVaryingNV (GLuint program, GLuint index, GLint *location)
func GetTransformFeedbackVaryingNV(program GLuint, index GLuint, location *GLint) 
{
	C.glGetTransformFeedbackVaryingNV(C.GLuint(program), C.GLuint(index), (*C.GLint)(location));
}

//extern void glDepthRangedNV (GLdouble zNear, GLdouble zFar)
func DepthRangedNV(zNear GLdouble, zFar GLdouble) 
{
	C.glDepthRangedNV(C.GLdouble(zNear), C.GLdouble(zFar));
}

//extern void glClearDepthdNV (GLdouble depth)
func ClearDepthdNV(depth GLdouble) 
{
	C.glClearDepthdNV(C.GLdouble(depth));
}

//extern void glDepthBoundsdNV (GLdouble zmin, GLdouble zmax)
func DepthBoundsdNV(zmin GLdouble, zmax GLdouble) 
{
	C.glDepthBoundsdNV(C.GLdouble(zmin), C.GLdouble(zmax));
}

//extern void glColorMaskIndexedEXT (GLuint index, GLboolean r, GLboolean g, GLboolean b, GLboolean a)
func ColorMaskIndexedEXT(index GLuint, r GLboolean, g GLboolean, b GLboolean, a GLboolean) 
{
	C.glColorMaskIndexedEXT(C.GLuint(index), C.GLboolean(r), C.GLboolean(g), C.GLboolean(b), C.GLboolean(a));
}

//extern void glGetBooleanIndexedvEXT (GLenum target, GLuint index, GLboolean *data)
func GetBooleanIndexedvEXT(target GLenum, index GLuint, data *GLboolean) 
{
	C.glGetBooleanIndexedvEXT(C.GLenum(target), C.GLuint(index), (*C.GLboolean)(data));
}

//extern void glGetIntegerIndexedvEXT (GLenum target, GLuint index, GLint *data)
func GetIntegerIndexedvEXT(target GLenum, index GLuint, data *GLint) 
{
	C.glGetIntegerIndexedvEXT(C.GLenum(target), C.GLuint(index), (*C.GLint)(data));
}

//extern void glEnableIndexedEXT (GLenum target, GLuint index)
func EnableIndexedEXT(target GLenum, index GLuint) 
{
	C.glEnableIndexedEXT(C.GLenum(target), C.GLuint(index));
}

//extern void glDisableIndexedEXT (GLenum target, GLuint index)
func DisableIndexedEXT(target GLenum, index GLuint) 
{
	C.glDisableIndexedEXT(C.GLenum(target), C.GLuint(index));
}

//extern GLboolean glIsEnabledIndexedEXT (GLenum target, GLuint index)
func IsEnabledIndexedEXT(target GLenum, index GLuint) GLboolean
{
	return GLboolean(C.glIsEnabledIndexedEXT(C.GLenum(target), C.GLuint(index)));
}

//extern void glTexParameterIivEXT (GLenum target, GLenum pname, const GLint *params)
func TexParameterIivEXT(target GLenum, pname GLenum, params *GLint) 
{
	C.glTexParameterIivEXT(C.GLenum(target), C.GLenum(pname), (*C.GLint)(params));
}

//extern void glTexParameterIuivEXT (GLenum target, GLenum pname, const GLuint *params)
func TexParameterIuivEXT(target GLenum, pname GLenum, params *GLuint) 
{
	C.glTexParameterIuivEXT(C.GLenum(target), C.GLenum(pname), (*C.GLuint)(params));
}

//extern void glGetTexParameterIivEXT (GLenum target, GLenum pname, GLint *params)
func GetTexParameterIivEXT(target GLenum, pname GLenum, params *GLint) 
{
	C.glGetTexParameterIivEXT(C.GLenum(target), C.GLenum(pname), (*C.GLint)(params));
}

//extern void glGetTexParameterIuivEXT (GLenum target, GLenum pname, GLuint *params)
func GetTexParameterIuivEXT(target GLenum, pname GLenum, params *GLuint) 
{
	C.glGetTexParameterIuivEXT(C.GLenum(target), C.GLenum(pname), (*C.GLuint)(params));
}

//extern void glClearColorIiEXT (GLint red, GLint green, GLint blue, GLint alpha)
func ClearColorIiEXT(red GLint, green GLint, blue GLint, alpha GLint) 
{
	C.glClearColorIiEXT(C.GLint(red), C.GLint(green), C.GLint(blue), C.GLint(alpha));
}

//extern void glClearColorIuiEXT (GLuint red, GLuint green, GLuint blue, GLuint alpha)
func ClearColorIuiEXT(red GLuint, green GLuint, blue GLuint, alpha GLuint) 
{
	C.glClearColorIuiEXT(C.GLuint(red), C.GLuint(green), C.GLuint(blue), C.GLuint(alpha));
}

//extern void glUniformBufferEXT (GLuint program, GLint location, GLuint buffer)
func UniformBufferEXT(program GLuint, location GLint, buffer GLuint) 
{
	C.glUniformBufferEXT(C.GLuint(program), C.GLint(location), C.GLuint(buffer));
}

//extern GLint glGetUniformBufferSizeEXT (GLuint program, GLint location)
func GetUniformBufferSizeEXT(program GLuint, location GLint) GLint
{
	return GLint(C.glGetUniformBufferSizeEXT(C.GLuint(program), C.GLint(location)));
}

//extern GLintptr glGetUniformOffsetEXT (GLuint program, GLint location)
func GetUniformOffsetEXT(program GLuint, location GLint) GLintptr
{
	return GLintptr(C.glGetUniformOffsetEXT(C.GLuint(program), C.GLint(location)));
}

//extern void glGetUniformuivEXT (GLuint program, GLint location, GLuint *params)
func GetUniformuivEXT(program GLuint, location GLint, params *GLuint) 
{
	C.glGetUniformuivEXT(C.GLuint(program), C.GLint(location), (*C.GLuint)(params));
}

//extern void glBindFragDataLocationEXT (GLuint program, GLuint color, const GLchar *name)
func BindFragDataLocationEXT(program GLuint, color GLuint, name string) 
{
	C.glBindFragDataLocationEXT(C.GLuint(program), C.GLuint(color), (*_C_GLchar)((C.CString)(name)));
}

//extern GLint glGetFragDataLocationEXT (GLuint program, const GLchar *name)
func GetFragDataLocationEXT(program GLuint, name string) GLint
{
	return GLint(C.glGetFragDataLocationEXT(C.GLuint(program), (*_C_GLchar)((C.CString)(name))));
}

//extern void glUniform1uiEXT (GLint location, GLuint v0)
func Uniform1uiEXT(location GLint, v0 GLuint) 
{
	C.glUniform1uiEXT(C.GLint(location), C.GLuint(v0));
}

//extern void glUniform2uiEXT (GLint location, GLuint v0, GLuint v1)
func Uniform2uiEXT(location GLint, v0 GLuint, v1 GLuint) 
{
	C.glUniform2uiEXT(C.GLint(location), C.GLuint(v0), C.GLuint(v1));
}

//extern void glUniform3uiEXT (GLint location, GLuint v0, GLuint v1, GLuint v2)
func Uniform3uiEXT(location GLint, v0 GLuint, v1 GLuint, v2 GLuint) 
{
	C.glUniform3uiEXT(C.GLint(location), C.GLuint(v0), C.GLuint(v1), C.GLuint(v2));
}

//extern void glUniform4uiEXT (GLint location, GLuint v0, GLuint v1, GLuint v2, GLuint v3)
func Uniform4uiEXT(location GLint, v0 GLuint, v1 GLuint, v2 GLuint, v3 GLuint) 
{
	C.glUniform4uiEXT(C.GLint(location), C.GLuint(v0), C.GLuint(v1), C.GLuint(v2), C.GLuint(v3));
}

//extern void glUniform1uivEXT (GLint location, GLsizei count, const GLuint *value)
func Uniform1uivEXT(location GLint, count GLsizei, value *GLuint) 
{
	C.glUniform1uivEXT(C.GLint(location), C.GLsizei(count), (*C.GLuint)(value));
}

//extern void glUniform2uivEXT (GLint location, GLsizei count, const GLuint *value)
func Uniform2uivEXT(location GLint, count GLsizei, value *GLuint) 
{
	C.glUniform2uivEXT(C.GLint(location), C.GLsizei(count), (*C.GLuint)(value));
}

//extern void glUniform3uivEXT (GLint location, GLsizei count, const GLuint *value)
func Uniform3uivEXT(location GLint, count GLsizei, value *GLuint) 
{
	C.glUniform3uivEXT(C.GLint(location), C.GLsizei(count), (*C.GLuint)(value));
}

//extern void glUniform4uivEXT (GLint location, GLsizei count, const GLuint *value)
func Uniform4uivEXT(location GLint, count GLsizei, value *GLuint) 
{
	C.glUniform4uivEXT(C.GLint(location), C.GLsizei(count), (*C.GLuint)(value));
}

//extern void glVertexAttribI1iEXT (GLuint index, GLint x)
func VertexAttribI1iEXT(index GLuint, x GLint) 
{
	C.glVertexAttribI1iEXT(C.GLuint(index), C.GLint(x));
}

//extern void glVertexAttribI2iEXT (GLuint index, GLint x, GLint y)
func VertexAttribI2iEXT(index GLuint, x GLint, y GLint) 
{
	C.glVertexAttribI2iEXT(C.GLuint(index), C.GLint(x), C.GLint(y));
}

//extern void glVertexAttribI3iEXT (GLuint index, GLint x, GLint y, GLint z)
func VertexAttribI3iEXT(index GLuint, x GLint, y GLint, z GLint) 
{
	C.glVertexAttribI3iEXT(C.GLuint(index), C.GLint(x), C.GLint(y), C.GLint(z));
}

//extern void glVertexAttribI4iEXT (GLuint index, GLint x, GLint y, GLint z, GLint w)
func VertexAttribI4iEXT(index GLuint, x GLint, y GLint, z GLint, w GLint) 
{
	C.glVertexAttribI4iEXT(C.GLuint(index), C.GLint(x), C.GLint(y), C.GLint(z), C.GLint(w));
}

//extern void glVertexAttribI1uiEXT (GLuint index, GLuint x)
func VertexAttribI1uiEXT(index GLuint, x GLuint) 
{
	C.glVertexAttribI1uiEXT(C.GLuint(index), C.GLuint(x));
}

//extern void glVertexAttribI2uiEXT (GLuint index, GLuint x, GLuint y)
func VertexAttribI2uiEXT(index GLuint, x GLuint, y GLuint) 
{
	C.glVertexAttribI2uiEXT(C.GLuint(index), C.GLuint(x), C.GLuint(y));
}

//extern void glVertexAttribI3uiEXT (GLuint index, GLuint x, GLuint y, GLuint z)
func VertexAttribI3uiEXT(index GLuint, x GLuint, y GLuint, z GLuint) 
{
	C.glVertexAttribI3uiEXT(C.GLuint(index), C.GLuint(x), C.GLuint(y), C.GLuint(z));
}

//extern void glVertexAttribI4uiEXT (GLuint index, GLuint x, GLuint y, GLuint z, GLuint w)
func VertexAttribI4uiEXT(index GLuint, x GLuint, y GLuint, z GLuint, w GLuint) 
{
	C.glVertexAttribI4uiEXT(C.GLuint(index), C.GLuint(x), C.GLuint(y), C.GLuint(z), C.GLuint(w));
}

//extern void glVertexAttribI1ivEXT (GLuint index, const GLint *v)
func VertexAttribI1ivEXT(index GLuint, v *GLint) 
{
	C.glVertexAttribI1ivEXT(C.GLuint(index), (*C.GLint)(v));
}

//extern void glVertexAttribI2ivEXT (GLuint index, const GLint *v)
func VertexAttribI2ivEXT(index GLuint, v *GLint) 
{
	C.glVertexAttribI2ivEXT(C.GLuint(index), (*C.GLint)(v));
}

//extern void glVertexAttribI3ivEXT (GLuint index, const GLint *v)
func VertexAttribI3ivEXT(index GLuint, v *GLint) 
{
	C.glVertexAttribI3ivEXT(C.GLuint(index), (*C.GLint)(v));
}

//extern void glVertexAttribI4ivEXT (GLuint index, const GLint *v)
func VertexAttribI4ivEXT(index GLuint, v *GLint) 
{
	C.glVertexAttribI4ivEXT(C.GLuint(index), (*C.GLint)(v));
}

//extern void glVertexAttribI1uivEXT (GLuint index, const GLuint *v)
func VertexAttribI1uivEXT(index GLuint, v *GLuint) 
{
	C.glVertexAttribI1uivEXT(C.GLuint(index), (*C.GLuint)(v));
}

//extern void glVertexAttribI2uivEXT (GLuint index, const GLuint *v)
func VertexAttribI2uivEXT(index GLuint, v *GLuint) 
{
	C.glVertexAttribI2uivEXT(C.GLuint(index), (*C.GLuint)(v));
}

//extern void glVertexAttribI3uivEXT (GLuint index, const GLuint *v)
func VertexAttribI3uivEXT(index GLuint, v *GLuint) 
{
	C.glVertexAttribI3uivEXT(C.GLuint(index), (*C.GLuint)(v));
}

//extern void glVertexAttribI4uivEXT (GLuint index, const GLuint *v)
func VertexAttribI4uivEXT(index GLuint, v *GLuint) 
{
	C.glVertexAttribI4uivEXT(C.GLuint(index), (*C.GLuint)(v));
}

//extern void glVertexAttribI4bvEXT (GLuint index, const GLbyte *v)
func VertexAttribI4bvEXT(index GLuint, v *GLbyte) 
{
	C.glVertexAttribI4bvEXT(C.GLuint(index), (*C.GLbyte)(v));
}

//extern void glVertexAttribI4svEXT (GLuint index, const GLshort *v)
func VertexAttribI4svEXT(index GLuint, v *GLshort) 
{
	C.glVertexAttribI4svEXT(C.GLuint(index), (*C.GLshort)(v));
}

//extern void glVertexAttribI4ubvEXT (GLuint index, const GLubyte *v)
func VertexAttribI4ubvEXT(index GLuint, v *GLubyte) 
{
	C.glVertexAttribI4ubvEXT(C.GLuint(index), (*C.GLubyte)(v));
}

//extern void glVertexAttribI4usvEXT (GLuint index, const GLushort *v)
func VertexAttribI4usvEXT(index GLuint, v *GLushort) 
{
	C.glVertexAttribI4usvEXT(C.GLuint(index), (*C.GLushort)(v));
}

//extern void glVertexAttribIPointerEXT (GLuint index, GLint size, GLenum type, GLsizei stride, const GLvoid *pointer)
func VertexAttribIPointerEXT(index GLuint, size GLint, type_ GLenum, stride GLsizei, pointer unsafe.Pointer) 
{
	C.glVertexAttribIPointerEXT(C.GLuint(index), C.GLint(size), C.GLenum(type_), C.GLsizei(stride), pointer);
}

//extern void glGetVertexAttribIivEXT (GLuint index, GLenum pname, GLint *params)
func GetVertexAttribIivEXT(index GLuint, pname GLenum, params *GLint) 
{
	C.glGetVertexAttribIivEXT(C.GLuint(index), C.GLenum(pname), (*C.GLint)(params));
}

//extern void glGetVertexAttribIuivEXT (GLuint index, GLenum pname, GLuint *params)
func GetVertexAttribIuivEXT(index GLuint, pname GLenum, params *GLuint) 
{
	C.glGetVertexAttribIuivEXT(C.GLuint(index), C.GLenum(pname), (*C.GLuint)(params));
}

//extern void glProgramParameteriEXT (GLuint program, GLenum pname, GLint value)
func ProgramParameteriEXT(program GLuint, pname GLenum, value GLint) 
{
	C.glProgramParameteriEXT(C.GLuint(program), C.GLenum(pname), C.GLint(value));
}

//extern void glFramebufferTextureEXT (GLenum target, GLenum attachment, GLuint texture, GLint level)
func FramebufferTextureEXT(target GLenum, attachment GLenum, texture GLuint, level GLint) 
{
	C.glFramebufferTextureEXT(C.GLenum(target), C.GLenum(attachment), C.GLuint(texture), C.GLint(level));
}

//extern void glFramebufferTextureLayerEXT (GLenum target, GLenum attachment, GLuint texture, GLint level, GLint layer)
func FramebufferTextureLayerEXT(target GLenum, attachment GLenum, texture GLuint, level GLint, layer GLint) 
{
	C.glFramebufferTextureLayerEXT(C.GLenum(target), C.GLenum(attachment), C.GLuint(texture), C.GLint(level), C.GLint(layer));
}

//extern void glFramebufferTextureFaceEXT (GLenum target, GLenum attachment, GLuint texture, GLint level, GLenum face)
func FramebufferTextureFaceEXT(target GLenum, attachment GLenum, texture GLuint, level GLint, face GLenum) 
{
	C.glFramebufferTextureFaceEXT(C.GLenum(target), C.GLenum(attachment), C.GLuint(texture), C.GLint(level), C.GLenum(face));
}

//extern void glProgramVertexLimitNV (GLenum target, GLint limit)
func ProgramVertexLimitNV(target GLenum, limit GLint) 
{
	C.glProgramVertexLimitNV(C.GLenum(target), C.GLint(limit));
}

//extern void glProgramLocalParameterI4iNV (GLenum target, GLuint index, GLint x, GLint y, GLint z, GLint w)
func ProgramLocalParameterI4iNV(target GLenum, index GLuint, x GLint, y GLint, z GLint, w GLint) 
{
	C.glProgramLocalParameterI4iNV(C.GLenum(target), C.GLuint(index), C.GLint(x), C.GLint(y), C.GLint(z), C.GLint(w));
}

//extern void glProgramLocalParameterI4ivNV (GLenum target, GLuint index, const GLint *params)
func ProgramLocalParameterI4ivNV(target GLenum, index GLuint, params *GLint) 
{
	C.glProgramLocalParameterI4ivNV(C.GLenum(target), C.GLuint(index), (*C.GLint)(params));
}

//extern void glProgramLocalParametersI4ivNV (GLenum target, GLuint index, GLsizei count, const GLint *params)
func ProgramLocalParametersI4ivNV(target GLenum, index GLuint, count GLsizei, params *GLint) 
{
	C.glProgramLocalParametersI4ivNV(C.GLenum(target), C.GLuint(index), C.GLsizei(count), (*C.GLint)(params));
}

//extern void glProgramLocalParameterI4uiNV (GLenum target, GLuint index, GLuint x, GLuint y, GLuint z, GLuint w)
func ProgramLocalParameterI4uiNV(target GLenum, index GLuint, x GLuint, y GLuint, z GLuint, w GLuint) 
{
	C.glProgramLocalParameterI4uiNV(C.GLenum(target), C.GLuint(index), C.GLuint(x), C.GLuint(y), C.GLuint(z), C.GLuint(w));
}

//extern void glProgramLocalParameterI4uivNV (GLenum target, GLuint index, const GLuint *params)
func ProgramLocalParameterI4uivNV(target GLenum, index GLuint, params *GLuint) 
{
	C.glProgramLocalParameterI4uivNV(C.GLenum(target), C.GLuint(index), (*C.GLuint)(params));
}

//extern void glProgramLocalParametersI4uivNV (GLenum target, GLuint index, GLsizei count, const GLuint *params)
func ProgramLocalParametersI4uivNV(target GLenum, index GLuint, count GLsizei, params *GLuint) 
{
	C.glProgramLocalParametersI4uivNV(C.GLenum(target), C.GLuint(index), C.GLsizei(count), (*C.GLuint)(params));
}

//extern void glProgramEnvParameterI4iNV (GLenum target, GLuint index, GLint x, GLint y, GLint z, GLint w)
func ProgramEnvParameterI4iNV(target GLenum, index GLuint, x GLint, y GLint, z GLint, w GLint) 
{
	C.glProgramEnvParameterI4iNV(C.GLenum(target), C.GLuint(index), C.GLint(x), C.GLint(y), C.GLint(z), C.GLint(w));
}

//extern void glProgramEnvParameterI4ivNV (GLenum target, GLuint index, const GLint *params)
func ProgramEnvParameterI4ivNV(target GLenum, index GLuint, params *GLint) 
{
	C.glProgramEnvParameterI4ivNV(C.GLenum(target), C.GLuint(index), (*C.GLint)(params));
}

//extern void glProgramEnvParametersI4ivNV (GLenum target, GLuint index, GLsizei count, const GLint *params)
func ProgramEnvParametersI4ivNV(target GLenum, index GLuint, count GLsizei, params *GLint) 
{
	C.glProgramEnvParametersI4ivNV(C.GLenum(target), C.GLuint(index), C.GLsizei(count), (*C.GLint)(params));
}

//extern void glProgramEnvParameterI4uiNV (GLenum target, GLuint index, GLuint x, GLuint y, GLuint z, GLuint w)
func ProgramEnvParameterI4uiNV(target GLenum, index GLuint, x GLuint, y GLuint, z GLuint, w GLuint) 
{
	C.glProgramEnvParameterI4uiNV(C.GLenum(target), C.GLuint(index), C.GLuint(x), C.GLuint(y), C.GLuint(z), C.GLuint(w));
}

//extern void glProgramEnvParameterI4uivNV (GLenum target, GLuint index, const GLuint *params)
func ProgramEnvParameterI4uivNV(target GLenum, index GLuint, params *GLuint) 
{
	C.glProgramEnvParameterI4uivNV(C.GLenum(target), C.GLuint(index), (*C.GLuint)(params));
}

//extern void glProgramEnvParametersI4uivNV (GLenum target, GLuint index, GLsizei count, const GLuint *params)
func ProgramEnvParametersI4uivNV(target GLenum, index GLuint, count GLsizei, params *GLuint) 
{
	C.glProgramEnvParametersI4uivNV(C.GLenum(target), C.GLuint(index), C.GLsizei(count), (*C.GLuint)(params));
}

//extern void glGetProgramLocalParameterIivNV (GLenum target, GLuint index, GLint *params)
func GetProgramLocalParameterIivNV(target GLenum, index GLuint, params *GLint) 
{
	C.glGetProgramLocalParameterIivNV(C.GLenum(target), C.GLuint(index), (*C.GLint)(params));
}

//extern void glGetProgramLocalParameterIuivNV (GLenum target, GLuint index, GLuint *params)
func GetProgramLocalParameterIuivNV(target GLenum, index GLuint, params *GLuint) 
{
	C.glGetProgramLocalParameterIuivNV(C.GLenum(target), C.GLuint(index), (*C.GLuint)(params));
}

//extern void glGetProgramEnvParameterIivNV (GLenum target, GLuint index, GLint *params)
func GetProgramEnvParameterIivNV(target GLenum, index GLuint, params *GLint) 
{
	C.glGetProgramEnvParameterIivNV(C.GLenum(target), C.GLuint(index), (*C.GLint)(params));
}

//extern void glGetProgramEnvParameterIuivNV (GLenum target, GLuint index, GLuint *params)
func GetProgramEnvParameterIuivNV(target GLenum, index GLuint, params *GLuint) 
{
	C.glGetProgramEnvParameterIuivNV(C.GLenum(target), C.GLuint(index), (*C.GLuint)(params));
}

//extern void glProgramBufferParametersfvNV (GLenum target, GLuint buffer, GLuint index, GLsizei count, const GLfloat *params)
func ProgramBufferParametersfvNV(target GLenum, buffer GLuint, index GLuint, count GLsizei, params *GLfloat) 
{
	C.glProgramBufferParametersfvNV(C.GLenum(target), C.GLuint(buffer), C.GLuint(index), C.GLsizei(count), (*C.GLfloat)(params));
}

//extern void glProgramBufferParametersIivNV (GLenum target, GLuint buffer, GLuint index, GLsizei count, const GLint *params)
func ProgramBufferParametersIivNV(target GLenum, buffer GLuint, index GLuint, count GLsizei, params *GLint) 
{
	C.glProgramBufferParametersIivNV(C.GLenum(target), C.GLuint(buffer), C.GLuint(index), C.GLsizei(count), (*C.GLint)(params));
}

//extern void glProgramBufferParametersIuivNV (GLenum target, GLuint buffer, GLuint index, GLsizei count, const GLuint *params)
func ProgramBufferParametersIuivNV(target GLenum, buffer GLuint, index GLuint, count GLsizei, params *GLuint) 
{
	C.glProgramBufferParametersIuivNV(C.GLenum(target), C.GLuint(buffer), C.GLuint(index), C.GLsizei(count), (*C.GLuint)(params));
}

//extern void glRenderbufferStorageMultisampleEXT (GLenum target, GLsizei samples, GLenum internalformat, GLsizei width, GLsizei height)
func RenderbufferStorageMultisampleEXT(target GLenum, samples GLsizei, internalformat GLenum, width GLsizei, height GLsizei) 
{
	C.glRenderbufferStorageMultisampleEXT(C.GLenum(target), C.GLsizei(samples), C.GLenum(internalformat), C.GLsizei(width), C.GLsizei(height));
}

//extern void glRenderbufferStorageMultisampleCoverageNV (GLenum target, GLsizei coverageSamples, GLsizei colorSamples, GLenum internalformat, GLsizei width, GLsizei height)
func RenderbufferStorageMultisampleCoverageNV(target GLenum, coverageSamples GLsizei, colorSamples GLsizei, internalformat GLenum, width GLsizei, height GLsizei) 
{
	C.glRenderbufferStorageMultisampleCoverageNV(C.GLenum(target), C.GLsizei(coverageSamples), C.GLsizei(colorSamples), C.GLenum(internalformat), C.GLsizei(width), C.GLsizei(height));
}

//extern void glBlitFramebufferEXT (GLint srcX0, GLint srcY0, GLint srcX1, GLint srcY1, GLint dstX0, GLint dstY0, GLint dstX1, GLint dstY1, GLbitfield mask, GLenum filter)
func BlitFramebufferEXT(srcX0 GLint, srcY0 GLint, srcX1 GLint, srcY1 GLint, dstX0 GLint, dstY0 GLint, dstX1 GLint, dstY1 GLint, mask GLbitfield, filter GLenum) 
{
	C.glBlitFramebufferEXT(C.GLint(srcX0), C.GLint(srcY0), C.GLint(srcX1), C.GLint(srcY1), C.GLint(dstX0), C.GLint(dstY0), C.GLint(dstX1), C.GLint(dstY1), C.GLbitfield(mask), C.GLenum(filter));
}

//extern void glDrawArraysInstancedEXT (GLenum mode, GLint start, GLsizei count, GLsizei primcount)
func DrawArraysInstancedEXT(mode GLenum, start GLint, count GLsizei, primcount GLsizei) 
{
	C.glDrawArraysInstancedEXT(C.GLenum(mode), C.GLint(start), C.GLsizei(count), C.GLsizei(primcount));
}

//extern void glDrawElementsInstancedEXT (GLenum mode, GLsizei count, GLenum type, const GLvoid *indices, GLsizei primcount)
func DrawElementsInstancedEXT(mode GLenum, count GLsizei, type_ GLenum, indices unsafe.Pointer, primcount GLsizei) 
{
	C.glDrawElementsInstancedEXT(C.GLenum(mode), C.GLsizei(count), C.GLenum(type_), indices, C.GLsizei(primcount));
}

//extern void glGetVideoivNV (GLuint video_slot, GLenum pname, GLint *params)
func GetVideoivNV(video_slot GLuint, pname GLenum, params *GLint) 
{
	C.glGetVideoivNV(C.GLuint(video_slot), C.GLenum(pname), (*C.GLint)(params));
}

//extern void glGetVideouivNV (GLuint video_slot, GLenum pname, GLuint *params)
func GetVideouivNV(video_slot GLuint, pname GLenum, params *GLuint) 
{
	C.glGetVideouivNV(C.GLuint(video_slot), C.GLenum(pname), (*C.GLuint)(params));
}

//extern void glGetVideoi64vNV (GLuint video_slot, GLenum pname, GLint64EXT *params)
func GetVideoi64vNV(video_slot GLuint, pname GLenum, params *GLint64EXT) 
{
	C.glGetVideoi64vNV(C.GLuint(video_slot), C.GLenum(pname), (*C.GLint64EXT)(params));
}

//extern void glGetVideoui64vNV (GLuint video_slot, GLenum pname, GLuint64EXT *params)
func GetVideoui64vNV(video_slot GLuint, pname GLenum, params *GLuint64EXT) 
{
	C.glGetVideoui64vNV(C.GLuint(video_slot), C.GLenum(pname), (*C.GLuint64EXT)(params));
}

//extern void glBeginConditionalRenderNV (GLuint id, GLenum mode)
func BeginConditionalRenderNV(id GLuint, mode GLenum) 
{
	C.glBeginConditionalRenderNV(C.GLuint(id), C.GLenum(mode));
}

//extern void glEndConditionalRenderNV (void)
func EndConditionalRenderNV() 
{
	C.glEndConditionalRenderNV();
}

//extern void glBeginTransformFeedbackEXT (GLenum)
func BeginTransformFeedbackEXT(arg0 GLenum) 
{
	C.glBeginTransformFeedbackEXT(C.GLenum(arg0));
}

//extern void glEndTransformFeedbackEXT (void)
func EndTransformFeedbackEXT() 
{
	C.glEndTransformFeedbackEXT();
}

//extern void glBindBufferRangeEXT (GLenum, GLuint, GLuint, GLintptr, GLsizeiptr)
func BindBufferRangeEXT(arg0 GLenum, arg1 GLuint, arg2 GLuint, arg3 GLintptr, arg4 GLsizeiptr) 
{
	C.glBindBufferRangeEXT(C.GLenum(arg0), C.GLuint(arg1), C.GLuint(arg2), C.GLintptr(arg3), C.GLsizeiptr(arg4));
}

//extern void glBindBufferOffsetEXT (GLenum, GLuint, GLuint, GLintptr)
func BindBufferOffsetEXT(arg0 GLenum, arg1 GLuint, arg2 GLuint, arg3 GLintptr) 
{
	C.glBindBufferOffsetEXT(C.GLenum(arg0), C.GLuint(arg1), C.GLuint(arg2), C.GLintptr(arg3));
}

//extern void glBindBufferBaseEXT (GLenum, GLuint, GLuint)
func BindBufferBaseEXT(arg0 GLenum, arg1 GLuint, arg2 GLuint) 
{
	C.glBindBufferBaseEXT(C.GLenum(arg0), C.GLuint(arg1), C.GLuint(arg2));
}

//extern void glTransformFeedbackVaryingsEXT (GLuint, GLsizei, const GLchar **, GLenum)
func TransformFeedbackVaryingsEXT(arg0 GLuint, arg1 GLsizei, arg2 **GLchar, arg3 GLenum) 
{
	C.glTransformFeedbackVaryingsEXT(C.GLuint(arg0), C.GLsizei(arg1), (**C.GLchar)(arg2), C.GLenum(arg3));
}

//extern void glGetTransformFeedbackVaryingEXT (GLuint, GLuint, GLint *)
func GetTransformFeedbackVaryingEXT(arg0 GLuint, arg1 GLuint, arg2 *GLint) 
{
	C.glGetTransformFeedbackVaryingEXT(C.GLuint(arg0), C.GLuint(arg1), (*C.GLint)(arg2));
}

//extern void glClientAttribDefaultEXT (GLbitfield)
func ClientAttribDefaultEXT(arg0 GLbitfield) 
{
	C.glClientAttribDefaultEXT(C.GLbitfield(arg0));
}

//extern void glPushClientAttribDefaultEXT (GLbitfield)
func PushClientAttribDefaultEXT(arg0 GLbitfield) 
{
	C.glPushClientAttribDefaultEXT(C.GLbitfield(arg0));
}

//extern void glMatrixLoadfEXT (GLenum, const GLfloat *)
func MatrixLoadfEXT(arg0 GLenum, arg1 *GLfloat) 
{
	C.glMatrixLoadfEXT(C.GLenum(arg0), (*C.GLfloat)(arg1));
}

//extern void glMatrixLoaddEXT (GLenum, const GLdouble *)
func MatrixLoaddEXT(arg0 GLenum, arg1 *GLdouble) 
{
	C.glMatrixLoaddEXT(C.GLenum(arg0), (*C.GLdouble)(arg1));
}

//extern void glMatrixMultfEXT (GLenum, const GLfloat *)
func MatrixMultfEXT(arg0 GLenum, arg1 *GLfloat) 
{
	C.glMatrixMultfEXT(C.GLenum(arg0), (*C.GLfloat)(arg1));
}

//extern void glMatrixMultdEXT (GLenum, const GLdouble *)
func MatrixMultdEXT(arg0 GLenum, arg1 *GLdouble) 
{
	C.glMatrixMultdEXT(C.GLenum(arg0), (*C.GLdouble)(arg1));
}

//extern void glMatrixLoadIdentityEXT (GLenum)
func MatrixLoadIdentityEXT(arg0 GLenum) 
{
	C.glMatrixLoadIdentityEXT(C.GLenum(arg0));
}

//extern void glMatrixRotatefEXT (GLenum, GLfloat, GLfloat, GLfloat, GLfloat)
func MatrixRotatefEXT(arg0 GLenum, arg1 GLfloat, arg2 GLfloat, arg3 GLfloat, arg4 GLfloat) 
{
	C.glMatrixRotatefEXT(C.GLenum(arg0), C.GLfloat(arg1), C.GLfloat(arg2), C.GLfloat(arg3), C.GLfloat(arg4));
}

//extern void glMatrixRotatedEXT (GLenum, GLdouble, GLdouble, GLdouble, GLdouble)
func MatrixRotatedEXT(arg0 GLenum, arg1 GLdouble, arg2 GLdouble, arg3 GLdouble, arg4 GLdouble) 
{
	C.glMatrixRotatedEXT(C.GLenum(arg0), C.GLdouble(arg1), C.GLdouble(arg2), C.GLdouble(arg3), C.GLdouble(arg4));
}

//extern void glMatrixScalefEXT (GLenum, GLfloat, GLfloat, GLfloat)
func MatrixScalefEXT(arg0 GLenum, arg1 GLfloat, arg2 GLfloat, arg3 GLfloat) 
{
	C.glMatrixScalefEXT(C.GLenum(arg0), C.GLfloat(arg1), C.GLfloat(arg2), C.GLfloat(arg3));
}

//extern void glMatrixScaledEXT (GLenum, GLdouble, GLdouble, GLdouble)
func MatrixScaledEXT(arg0 GLenum, arg1 GLdouble, arg2 GLdouble, arg3 GLdouble) 
{
	C.glMatrixScaledEXT(C.GLenum(arg0), C.GLdouble(arg1), C.GLdouble(arg2), C.GLdouble(arg3));
}

//extern void glMatrixTranslatefEXT (GLenum, GLfloat, GLfloat, GLfloat)
func MatrixTranslatefEXT(arg0 GLenum, arg1 GLfloat, arg2 GLfloat, arg3 GLfloat) 
{
	C.glMatrixTranslatefEXT(C.GLenum(arg0), C.GLfloat(arg1), C.GLfloat(arg2), C.GLfloat(arg3));
}

//extern void glMatrixTranslatedEXT (GLenum, GLdouble, GLdouble, GLdouble)
func MatrixTranslatedEXT(arg0 GLenum, arg1 GLdouble, arg2 GLdouble, arg3 GLdouble) 
{
	C.glMatrixTranslatedEXT(C.GLenum(arg0), C.GLdouble(arg1), C.GLdouble(arg2), C.GLdouble(arg3));
}

//extern void glMatrixFrustumEXT (GLenum, GLdouble, GLdouble, GLdouble, GLdouble, GLdouble, GLdouble)
func MatrixFrustumEXT(arg0 GLenum, arg1 GLdouble, arg2 GLdouble, arg3 GLdouble, arg4 GLdouble, arg5 GLdouble, arg6 GLdouble) 
{
	C.glMatrixFrustumEXT(C.GLenum(arg0), C.GLdouble(arg1), C.GLdouble(arg2), C.GLdouble(arg3), C.GLdouble(arg4), C.GLdouble(arg5), C.GLdouble(arg6));
}

//extern void glMatrixOrthoEXT (GLenum, GLdouble, GLdouble, GLdouble, GLdouble, GLdouble, GLdouble)
func MatrixOrthoEXT(arg0 GLenum, arg1 GLdouble, arg2 GLdouble, arg3 GLdouble, arg4 GLdouble, arg5 GLdouble, arg6 GLdouble) 
{
	C.glMatrixOrthoEXT(C.GLenum(arg0), C.GLdouble(arg1), C.GLdouble(arg2), C.GLdouble(arg3), C.GLdouble(arg4), C.GLdouble(arg5), C.GLdouble(arg6));
}

//extern void glMatrixPopEXT (GLenum)
func MatrixPopEXT(arg0 GLenum) 
{
	C.glMatrixPopEXT(C.GLenum(arg0));
}

//extern void glMatrixPushEXT (GLenum)
func MatrixPushEXT(arg0 GLenum) 
{
	C.glMatrixPushEXT(C.GLenum(arg0));
}

//extern void glMatrixLoadTransposefEXT (GLenum, const GLfloat *)
func MatrixLoadTransposefEXT(arg0 GLenum, arg1 *GLfloat) 
{
	C.glMatrixLoadTransposefEXT(C.GLenum(arg0), (*C.GLfloat)(arg1));
}

//extern void glMatrixLoadTransposedEXT (GLenum, const GLdouble *)
func MatrixLoadTransposedEXT(arg0 GLenum, arg1 *GLdouble) 
{
	C.glMatrixLoadTransposedEXT(C.GLenum(arg0), (*C.GLdouble)(arg1));
}

//extern void glMatrixMultTransposefEXT (GLenum, const GLfloat *)
func MatrixMultTransposefEXT(arg0 GLenum, arg1 *GLfloat) 
{
	C.glMatrixMultTransposefEXT(C.GLenum(arg0), (*C.GLfloat)(arg1));
}

//extern void glMatrixMultTransposedEXT (GLenum, const GLdouble *)
func MatrixMultTransposedEXT(arg0 GLenum, arg1 *GLdouble) 
{
	C.glMatrixMultTransposedEXT(C.GLenum(arg0), (*C.GLdouble)(arg1));
}

//extern void glTextureParameterfEXT (GLuint, GLenum, GLenum, GLfloat)
func TextureParameterfEXT(arg0 GLuint, arg1 GLenum, arg2 GLenum, arg3 GLfloat) 
{
	C.glTextureParameterfEXT(C.GLuint(arg0), C.GLenum(arg1), C.GLenum(arg2), C.GLfloat(arg3));
}

//extern void glTextureParameterfvEXT (GLuint, GLenum, GLenum, const GLfloat *)
func TextureParameterfvEXT(arg0 GLuint, arg1 GLenum, arg2 GLenum, arg3 *GLfloat) 
{
	C.glTextureParameterfvEXT(C.GLuint(arg0), C.GLenum(arg1), C.GLenum(arg2), (*C.GLfloat)(arg3));
}

//extern void glTextureParameteriEXT (GLuint, GLenum, GLenum, GLint)
func TextureParameteriEXT(arg0 GLuint, arg1 GLenum, arg2 GLenum, arg3 GLint) 
{
	C.glTextureParameteriEXT(C.GLuint(arg0), C.GLenum(arg1), C.GLenum(arg2), C.GLint(arg3));
}

//extern void glTextureParameterivEXT (GLuint, GLenum, GLenum, const GLint *)
func TextureParameterivEXT(arg0 GLuint, arg1 GLenum, arg2 GLenum, arg3 *GLint) 
{
	C.glTextureParameterivEXT(C.GLuint(arg0), C.GLenum(arg1), C.GLenum(arg2), (*C.GLint)(arg3));
}

//extern void glTextureImage1DEXT (GLuint, GLenum, GLint, GLenum, GLsizei, GLint, GLenum, GLenum, const GLvoid *)
func TextureImage1DEXT(arg0 GLuint, arg1 GLenum, arg2 GLint, arg3 GLenum, arg4 GLsizei, arg5 GLint, arg6 GLenum, arg7 GLenum, arg8 unsafe.Pointer) 
{
	C.glTextureImage1DEXT(C.GLuint(arg0), C.GLenum(arg1), C.GLint(arg2), C.GLenum(arg3), C.GLsizei(arg4), C.GLint(arg5), C.GLenum(arg6), C.GLenum(arg7), arg8);
}

//extern void glTextureImage2DEXT (GLuint, GLenum, GLint, GLenum, GLsizei, GLsizei, GLint, GLenum, GLenum, const GLvoid *)
func TextureImage2DEXT(arg0 GLuint, arg1 GLenum, arg2 GLint, arg3 GLenum, arg4 GLsizei, arg5 GLsizei, arg6 GLint, arg7 GLenum, arg8 GLenum, arg9 unsafe.Pointer) 
{
	C.glTextureImage2DEXT(C.GLuint(arg0), C.GLenum(arg1), C.GLint(arg2), C.GLenum(arg3), C.GLsizei(arg4), C.GLsizei(arg5), C.GLint(arg6), C.GLenum(arg7), C.GLenum(arg8), arg9);
}

//extern void glTextureSubImage1DEXT (GLuint, GLenum, GLint, GLint, GLsizei, GLenum, GLenum, const GLvoid *)
func TextureSubImage1DEXT(arg0 GLuint, arg1 GLenum, arg2 GLint, arg3 GLint, arg4 GLsizei, arg5 GLenum, arg6 GLenum, arg7 unsafe.Pointer) 
{
	C.glTextureSubImage1DEXT(C.GLuint(arg0), C.GLenum(arg1), C.GLint(arg2), C.GLint(arg3), C.GLsizei(arg4), C.GLenum(arg5), C.GLenum(arg6), arg7);
}

//extern void glTextureSubImage2DEXT (GLuint, GLenum, GLint, GLint, GLint, GLsizei, GLsizei, GLenum, GLenum, const GLvoid *)
func TextureSubImage2DEXT(arg0 GLuint, arg1 GLenum, arg2 GLint, arg3 GLint, arg4 GLint, arg5 GLsizei, arg6 GLsizei, arg7 GLenum, arg8 GLenum, arg9 unsafe.Pointer) 
{
	C.glTextureSubImage2DEXT(C.GLuint(arg0), C.GLenum(arg1), C.GLint(arg2), C.GLint(arg3), C.GLint(arg4), C.GLsizei(arg5), C.GLsizei(arg6), C.GLenum(arg7), C.GLenum(arg8), arg9);
}

//extern void glCopyTextureImage1DEXT (GLuint, GLenum, GLint, GLenum, GLint, GLint, GLsizei, GLint)
func CopyTextureImage1DEXT(arg0 GLuint, arg1 GLenum, arg2 GLint, arg3 GLenum, arg4 GLint, arg5 GLint, arg6 GLsizei, arg7 GLint) 
{
	C.glCopyTextureImage1DEXT(C.GLuint(arg0), C.GLenum(arg1), C.GLint(arg2), C.GLenum(arg3), C.GLint(arg4), C.GLint(arg5), C.GLsizei(arg6), C.GLint(arg7));
}

//extern void glCopyTextureImage2DEXT (GLuint, GLenum, GLint, GLenum, GLint, GLint, GLsizei, GLsizei, GLint)
func CopyTextureImage2DEXT(arg0 GLuint, arg1 GLenum, arg2 GLint, arg3 GLenum, arg4 GLint, arg5 GLint, arg6 GLsizei, arg7 GLsizei, arg8 GLint) 
{
	C.glCopyTextureImage2DEXT(C.GLuint(arg0), C.GLenum(arg1), C.GLint(arg2), C.GLenum(arg3), C.GLint(arg4), C.GLint(arg5), C.GLsizei(arg6), C.GLsizei(arg7), C.GLint(arg8));
}

//extern void glCopyTextureSubImage1DEXT (GLuint, GLenum, GLint, GLint, GLint, GLint, GLsizei)
func CopyTextureSubImage1DEXT(arg0 GLuint, arg1 GLenum, arg2 GLint, arg3 GLint, arg4 GLint, arg5 GLint, arg6 GLsizei) 
{
	C.glCopyTextureSubImage1DEXT(C.GLuint(arg0), C.GLenum(arg1), C.GLint(arg2), C.GLint(arg3), C.GLint(arg4), C.GLint(arg5), C.GLsizei(arg6));
}

//extern void glCopyTextureSubImage2DEXT (GLuint, GLenum, GLint, GLint, GLint, GLint, GLint, GLsizei, GLsizei)
func CopyTextureSubImage2DEXT(arg0 GLuint, arg1 GLenum, arg2 GLint, arg3 GLint, arg4 GLint, arg5 GLint, arg6 GLint, arg7 GLsizei, arg8 GLsizei) 
{
	C.glCopyTextureSubImage2DEXT(C.GLuint(arg0), C.GLenum(arg1), C.GLint(arg2), C.GLint(arg3), C.GLint(arg4), C.GLint(arg5), C.GLint(arg6), C.GLsizei(arg7), C.GLsizei(arg8));
}

//extern void glGetTextureImageEXT (GLuint, GLenum, GLint, GLenum, GLenum, GLvoid *)
func GetTextureImageEXT(arg0 GLuint, arg1 GLenum, arg2 GLint, arg3 GLenum, arg4 GLenum, arg5 unsafe.Pointer) 
{
	C.glGetTextureImageEXT(C.GLuint(arg0), C.GLenum(arg1), C.GLint(arg2), C.GLenum(arg3), C.GLenum(arg4), arg5);
}

//extern void glGetTextureParameterfvEXT (GLuint, GLenum, GLenum, GLfloat *)
func GetTextureParameterfvEXT(arg0 GLuint, arg1 GLenum, arg2 GLenum, arg3 *GLfloat) 
{
	C.glGetTextureParameterfvEXT(C.GLuint(arg0), C.GLenum(arg1), C.GLenum(arg2), (*C.GLfloat)(arg3));
}

//extern void glGetTextureParameterivEXT (GLuint, GLenum, GLenum, GLint *)
func GetTextureParameterivEXT(arg0 GLuint, arg1 GLenum, arg2 GLenum, arg3 *GLint) 
{
	C.glGetTextureParameterivEXT(C.GLuint(arg0), C.GLenum(arg1), C.GLenum(arg2), (*C.GLint)(arg3));
}

//extern void glGetTextureLevelParameterfvEXT (GLuint, GLenum, GLint, GLenum, GLfloat *)
func GetTextureLevelParameterfvEXT(arg0 GLuint, arg1 GLenum, arg2 GLint, arg3 GLenum, arg4 *GLfloat) 
{
	C.glGetTextureLevelParameterfvEXT(C.GLuint(arg0), C.GLenum(arg1), C.GLint(arg2), C.GLenum(arg3), (*C.GLfloat)(arg4));
}

//extern void glGetTextureLevelParameterivEXT (GLuint, GLenum, GLint, GLenum, GLint *)
func GetTextureLevelParameterivEXT(arg0 GLuint, arg1 GLenum, arg2 GLint, arg3 GLenum, arg4 *GLint) 
{
	C.glGetTextureLevelParameterivEXT(C.GLuint(arg0), C.GLenum(arg1), C.GLint(arg2), C.GLenum(arg3), (*C.GLint)(arg4));
}

//extern void glTextureImage3DEXT (GLuint, GLenum, GLint, GLenum, GLsizei, GLsizei, GLsizei, GLint, GLenum, GLenum, const GLvoid *)
func TextureImage3DEXT(arg0 GLuint, arg1 GLenum, arg2 GLint, arg3 GLenum, arg4 GLsizei, arg5 GLsizei, arg6 GLsizei, arg7 GLint, arg8 GLenum, arg9 GLenum, arg10 unsafe.Pointer) 
{
	C.glTextureImage3DEXT(C.GLuint(arg0), C.GLenum(arg1), C.GLint(arg2), C.GLenum(arg3), C.GLsizei(arg4), C.GLsizei(arg5), C.GLsizei(arg6), C.GLint(arg7), C.GLenum(arg8), C.GLenum(arg9), arg10);
}

//extern void glTextureSubImage3DEXT (GLuint, GLenum, GLint, GLint, GLint, GLint, GLsizei, GLsizei, GLsizei, GLenum, GLenum, const GLvoid *)
func TextureSubImage3DEXT(arg0 GLuint, arg1 GLenum, arg2 GLint, arg3 GLint, arg4 GLint, arg5 GLint, arg6 GLsizei, arg7 GLsizei, arg8 GLsizei, arg9 GLenum, arg10 GLenum, arg11 unsafe.Pointer) 
{
	C.glTextureSubImage3DEXT(C.GLuint(arg0), C.GLenum(arg1), C.GLint(arg2), C.GLint(arg3), C.GLint(arg4), C.GLint(arg5), C.GLsizei(arg6), C.GLsizei(arg7), C.GLsizei(arg8), C.GLenum(arg9), C.GLenum(arg10), arg11);
}

//extern void glCopyTextureSubImage3DEXT (GLuint, GLenum, GLint, GLint, GLint, GLint, GLint, GLint, GLsizei, GLsizei)
func CopyTextureSubImage3DEXT(arg0 GLuint, arg1 GLenum, arg2 GLint, arg3 GLint, arg4 GLint, arg5 GLint, arg6 GLint, arg7 GLint, arg8 GLsizei, arg9 GLsizei) 
{
	C.glCopyTextureSubImage3DEXT(C.GLuint(arg0), C.GLenum(arg1), C.GLint(arg2), C.GLint(arg3), C.GLint(arg4), C.GLint(arg5), C.GLint(arg6), C.GLint(arg7), C.GLsizei(arg8), C.GLsizei(arg9));
}

//extern void glMultiTexParameterfEXT (GLenum, GLenum, GLenum, GLfloat)
func MultiTexParameterfEXT(arg0 GLenum, arg1 GLenum, arg2 GLenum, arg3 GLfloat) 
{
	C.glMultiTexParameterfEXT(C.GLenum(arg0), C.GLenum(arg1), C.GLenum(arg2), C.GLfloat(arg3));
}

//extern void glMultiTexParameterfvEXT (GLenum, GLenum, GLenum, const GLfloat *)
func MultiTexParameterfvEXT(arg0 GLenum, arg1 GLenum, arg2 GLenum, arg3 *GLfloat) 
{
	C.glMultiTexParameterfvEXT(C.GLenum(arg0), C.GLenum(arg1), C.GLenum(arg2), (*C.GLfloat)(arg3));
}

//extern void glMultiTexParameteriEXT (GLenum, GLenum, GLenum, GLint)
func MultiTexParameteriEXT(arg0 GLenum, arg1 GLenum, arg2 GLenum, arg3 GLint) 
{
	C.glMultiTexParameteriEXT(C.GLenum(arg0), C.GLenum(arg1), C.GLenum(arg2), C.GLint(arg3));
}

//extern void glMultiTexParameterivEXT (GLenum, GLenum, GLenum, const GLint *)
func MultiTexParameterivEXT(arg0 GLenum, arg1 GLenum, arg2 GLenum, arg3 *GLint) 
{
	C.glMultiTexParameterivEXT(C.GLenum(arg0), C.GLenum(arg1), C.GLenum(arg2), (*C.GLint)(arg3));
}

//extern void glMultiTexImage1DEXT (GLenum, GLenum, GLint, GLenum, GLsizei, GLint, GLenum, GLenum, const GLvoid *)
func MultiTexImage1DEXT(arg0 GLenum, arg1 GLenum, arg2 GLint, arg3 GLenum, arg4 GLsizei, arg5 GLint, arg6 GLenum, arg7 GLenum, arg8 unsafe.Pointer) 
{
	C.glMultiTexImage1DEXT(C.GLenum(arg0), C.GLenum(arg1), C.GLint(arg2), C.GLenum(arg3), C.GLsizei(arg4), C.GLint(arg5), C.GLenum(arg6), C.GLenum(arg7), arg8);
}

//extern void glMultiTexImage2DEXT (GLenum, GLenum, GLint, GLenum, GLsizei, GLsizei, GLint, GLenum, GLenum, const GLvoid *)
func MultiTexImage2DEXT(arg0 GLenum, arg1 GLenum, arg2 GLint, arg3 GLenum, arg4 GLsizei, arg5 GLsizei, arg6 GLint, arg7 GLenum, arg8 GLenum, arg9 unsafe.Pointer) 
{
	C.glMultiTexImage2DEXT(C.GLenum(arg0), C.GLenum(arg1), C.GLint(arg2), C.GLenum(arg3), C.GLsizei(arg4), C.GLsizei(arg5), C.GLint(arg6), C.GLenum(arg7), C.GLenum(arg8), arg9);
}

//extern void glMultiTexSubImage1DEXT (GLenum, GLenum, GLint, GLint, GLsizei, GLenum, GLenum, const GLvoid *)
func MultiTexSubImage1DEXT(arg0 GLenum, arg1 GLenum, arg2 GLint, arg3 GLint, arg4 GLsizei, arg5 GLenum, arg6 GLenum, arg7 unsafe.Pointer) 
{
	C.glMultiTexSubImage1DEXT(C.GLenum(arg0), C.GLenum(arg1), C.GLint(arg2), C.GLint(arg3), C.GLsizei(arg4), C.GLenum(arg5), C.GLenum(arg6), arg7);
}

//extern void glMultiTexSubImage2DEXT (GLenum, GLenum, GLint, GLint, GLint, GLsizei, GLsizei, GLenum, GLenum, const GLvoid *)
func MultiTexSubImage2DEXT(arg0 GLenum, arg1 GLenum, arg2 GLint, arg3 GLint, arg4 GLint, arg5 GLsizei, arg6 GLsizei, arg7 GLenum, arg8 GLenum, arg9 unsafe.Pointer) 
{
	C.glMultiTexSubImage2DEXT(C.GLenum(arg0), C.GLenum(arg1), C.GLint(arg2), C.GLint(arg3), C.GLint(arg4), C.GLsizei(arg5), C.GLsizei(arg6), C.GLenum(arg7), C.GLenum(arg8), arg9);
}

//extern void glCopyMultiTexImage1DEXT (GLenum, GLenum, GLint, GLenum, GLint, GLint, GLsizei, GLint)
func CopyMultiTexImage1DEXT(arg0 GLenum, arg1 GLenum, arg2 GLint, arg3 GLenum, arg4 GLint, arg5 GLint, arg6 GLsizei, arg7 GLint) 
{
	C.glCopyMultiTexImage1DEXT(C.GLenum(arg0), C.GLenum(arg1), C.GLint(arg2), C.GLenum(arg3), C.GLint(arg4), C.GLint(arg5), C.GLsizei(arg6), C.GLint(arg7));
}

//extern void glCopyMultiTexImage2DEXT (GLenum, GLenum, GLint, GLenum, GLint, GLint, GLsizei, GLsizei, GLint)
func CopyMultiTexImage2DEXT(arg0 GLenum, arg1 GLenum, arg2 GLint, arg3 GLenum, arg4 GLint, arg5 GLint, arg6 GLsizei, arg7 GLsizei, arg8 GLint) 
{
	C.glCopyMultiTexImage2DEXT(C.GLenum(arg0), C.GLenum(arg1), C.GLint(arg2), C.GLenum(arg3), C.GLint(arg4), C.GLint(arg5), C.GLsizei(arg6), C.GLsizei(arg7), C.GLint(arg8));
}

//extern void glCopyMultiTexSubImage1DEXT (GLenum, GLenum, GLint, GLint, GLint, GLint, GLsizei)
func CopyMultiTexSubImage1DEXT(arg0 GLenum, arg1 GLenum, arg2 GLint, arg3 GLint, arg4 GLint, arg5 GLint, arg6 GLsizei) 
{
	C.glCopyMultiTexSubImage1DEXT(C.GLenum(arg0), C.GLenum(arg1), C.GLint(arg2), C.GLint(arg3), C.GLint(arg4), C.GLint(arg5), C.GLsizei(arg6));
}

//extern void glCopyMultiTexSubImage2DEXT (GLenum, GLenum, GLint, GLint, GLint, GLint, GLint, GLsizei, GLsizei)
func CopyMultiTexSubImage2DEXT(arg0 GLenum, arg1 GLenum, arg2 GLint, arg3 GLint, arg4 GLint, arg5 GLint, arg6 GLint, arg7 GLsizei, arg8 GLsizei) 
{
	C.glCopyMultiTexSubImage2DEXT(C.GLenum(arg0), C.GLenum(arg1), C.GLint(arg2), C.GLint(arg3), C.GLint(arg4), C.GLint(arg5), C.GLint(arg6), C.GLsizei(arg7), C.GLsizei(arg8));
}

//extern void glGetMultiTexImageEXT (GLenum, GLenum, GLint, GLenum, GLenum, GLvoid *)
func GetMultiTexImageEXT(arg0 GLenum, arg1 GLenum, arg2 GLint, arg3 GLenum, arg4 GLenum, arg5 unsafe.Pointer) 
{
	C.glGetMultiTexImageEXT(C.GLenum(arg0), C.GLenum(arg1), C.GLint(arg2), C.GLenum(arg3), C.GLenum(arg4), arg5);
}

//extern void glGetMultiTexParameterfvEXT (GLenum, GLenum, GLenum, GLfloat *)
func GetMultiTexParameterfvEXT(arg0 GLenum, arg1 GLenum, arg2 GLenum, arg3 *GLfloat) 
{
	C.glGetMultiTexParameterfvEXT(C.GLenum(arg0), C.GLenum(arg1), C.GLenum(arg2), (*C.GLfloat)(arg3));
}

//extern void glGetMultiTexParameterivEXT (GLenum, GLenum, GLenum, GLint *)
func GetMultiTexParameterivEXT(arg0 GLenum, arg1 GLenum, arg2 GLenum, arg3 *GLint) 
{
	C.glGetMultiTexParameterivEXT(C.GLenum(arg0), C.GLenum(arg1), C.GLenum(arg2), (*C.GLint)(arg3));
}

//extern void glGetMultiTexLevelParameterfvEXT (GLenum, GLenum, GLint, GLenum, GLfloat *)
func GetMultiTexLevelParameterfvEXT(arg0 GLenum, arg1 GLenum, arg2 GLint, arg3 GLenum, arg4 *GLfloat) 
{
	C.glGetMultiTexLevelParameterfvEXT(C.GLenum(arg0), C.GLenum(arg1), C.GLint(arg2), C.GLenum(arg3), (*C.GLfloat)(arg4));
}

//extern void glGetMultiTexLevelParameterivEXT (GLenum, GLenum, GLint, GLenum, GLint *)
func GetMultiTexLevelParameterivEXT(arg0 GLenum, arg1 GLenum, arg2 GLint, arg3 GLenum, arg4 *GLint) 
{
	C.glGetMultiTexLevelParameterivEXT(C.GLenum(arg0), C.GLenum(arg1), C.GLint(arg2), C.GLenum(arg3), (*C.GLint)(arg4));
}

//extern void glMultiTexImage3DEXT (GLenum, GLenum, GLint, GLenum, GLsizei, GLsizei, GLsizei, GLint, GLenum, GLenum, const GLvoid *)
func MultiTexImage3DEXT(arg0 GLenum, arg1 GLenum, arg2 GLint, arg3 GLenum, arg4 GLsizei, arg5 GLsizei, arg6 GLsizei, arg7 GLint, arg8 GLenum, arg9 GLenum, arg10 unsafe.Pointer) 
{
	C.glMultiTexImage3DEXT(C.GLenum(arg0), C.GLenum(arg1), C.GLint(arg2), C.GLenum(arg3), C.GLsizei(arg4), C.GLsizei(arg5), C.GLsizei(arg6), C.GLint(arg7), C.GLenum(arg8), C.GLenum(arg9), arg10);
}

//extern void glMultiTexSubImage3DEXT (GLenum, GLenum, GLint, GLint, GLint, GLint, GLsizei, GLsizei, GLsizei, GLenum, GLenum, const GLvoid *)
func MultiTexSubImage3DEXT(arg0 GLenum, arg1 GLenum, arg2 GLint, arg3 GLint, arg4 GLint, arg5 GLint, arg6 GLsizei, arg7 GLsizei, arg8 GLsizei, arg9 GLenum, arg10 GLenum, arg11 unsafe.Pointer) 
{
	C.glMultiTexSubImage3DEXT(C.GLenum(arg0), C.GLenum(arg1), C.GLint(arg2), C.GLint(arg3), C.GLint(arg4), C.GLint(arg5), C.GLsizei(arg6), C.GLsizei(arg7), C.GLsizei(arg8), C.GLenum(arg9), C.GLenum(arg10), arg11);
}

//extern void glCopyMultiTexSubImage3DEXT (GLenum, GLenum, GLint, GLint, GLint, GLint, GLint, GLint, GLsizei, GLsizei)
func CopyMultiTexSubImage3DEXT(arg0 GLenum, arg1 GLenum, arg2 GLint, arg3 GLint, arg4 GLint, arg5 GLint, arg6 GLint, arg7 GLint, arg8 GLsizei, arg9 GLsizei) 
{
	C.glCopyMultiTexSubImage3DEXT(C.GLenum(arg0), C.GLenum(arg1), C.GLint(arg2), C.GLint(arg3), C.GLint(arg4), C.GLint(arg5), C.GLint(arg6), C.GLint(arg7), C.GLsizei(arg8), C.GLsizei(arg9));
}

//extern void glBindMultiTextureEXT (GLenum, GLenum, GLuint)
func BindMultiTextureEXT(arg0 GLenum, arg1 GLenum, arg2 GLuint) 
{
	C.glBindMultiTextureEXT(C.GLenum(arg0), C.GLenum(arg1), C.GLuint(arg2));
}

//extern void glEnableClientStateIndexedEXT (GLenum, GLuint)
func EnableClientStateIndexedEXT(arg0 GLenum, arg1 GLuint) 
{
	C.glEnableClientStateIndexedEXT(C.GLenum(arg0), C.GLuint(arg1));
}

//extern void glDisableClientStateIndexedEXT (GLenum, GLuint)
func DisableClientStateIndexedEXT(arg0 GLenum, arg1 GLuint) 
{
	C.glDisableClientStateIndexedEXT(C.GLenum(arg0), C.GLuint(arg1));
}

//extern void glMultiTexCoordPointerEXT (GLenum, GLint, GLenum, GLsizei, const GLvoid *)
func MultiTexCoordPointerEXT(arg0 GLenum, arg1 GLint, arg2 GLenum, arg3 GLsizei, arg4 unsafe.Pointer) 
{
	C.glMultiTexCoordPointerEXT(C.GLenum(arg0), C.GLint(arg1), C.GLenum(arg2), C.GLsizei(arg3), arg4);
}

//extern void glMultiTexEnvfEXT (GLenum, GLenum, GLenum, GLfloat)
func MultiTexEnvfEXT(arg0 GLenum, arg1 GLenum, arg2 GLenum, arg3 GLfloat) 
{
	C.glMultiTexEnvfEXT(C.GLenum(arg0), C.GLenum(arg1), C.GLenum(arg2), C.GLfloat(arg3));
}

//extern void glMultiTexEnvfvEXT (GLenum, GLenum, GLenum, const GLfloat *)
func MultiTexEnvfvEXT(arg0 GLenum, arg1 GLenum, arg2 GLenum, arg3 *GLfloat) 
{
	C.glMultiTexEnvfvEXT(C.GLenum(arg0), C.GLenum(arg1), C.GLenum(arg2), (*C.GLfloat)(arg3));
}

//extern void glMultiTexEnviEXT (GLenum, GLenum, GLenum, GLint)
func MultiTexEnviEXT(arg0 GLenum, arg1 GLenum, arg2 GLenum, arg3 GLint) 
{
	C.glMultiTexEnviEXT(C.GLenum(arg0), C.GLenum(arg1), C.GLenum(arg2), C.GLint(arg3));
}

//extern void glMultiTexEnvivEXT (GLenum, GLenum, GLenum, const GLint *)
func MultiTexEnvivEXT(arg0 GLenum, arg1 GLenum, arg2 GLenum, arg3 *GLint) 
{
	C.glMultiTexEnvivEXT(C.GLenum(arg0), C.GLenum(arg1), C.GLenum(arg2), (*C.GLint)(arg3));
}

//extern void glMultiTexGendEXT (GLenum, GLenum, GLenum, GLdouble)
func MultiTexGendEXT(arg0 GLenum, arg1 GLenum, arg2 GLenum, arg3 GLdouble) 
{
	C.glMultiTexGendEXT(C.GLenum(arg0), C.GLenum(arg1), C.GLenum(arg2), C.GLdouble(arg3));
}

//extern void glMultiTexGendvEXT (GLenum, GLenum, GLenum, const GLdouble *)
func MultiTexGendvEXT(arg0 GLenum, arg1 GLenum, arg2 GLenum, arg3 *GLdouble) 
{
	C.glMultiTexGendvEXT(C.GLenum(arg0), C.GLenum(arg1), C.GLenum(arg2), (*C.GLdouble)(arg3));
}

//extern void glMultiTexGenfEXT (GLenum, GLenum, GLenum, GLfloat)
func MultiTexGenfEXT(arg0 GLenum, arg1 GLenum, arg2 GLenum, arg3 GLfloat) 
{
	C.glMultiTexGenfEXT(C.GLenum(arg0), C.GLenum(arg1), C.GLenum(arg2), C.GLfloat(arg3));
}

//extern void glMultiTexGenfvEXT (GLenum, GLenum, GLenum, const GLfloat *)
func MultiTexGenfvEXT(arg0 GLenum, arg1 GLenum, arg2 GLenum, arg3 *GLfloat) 
{
	C.glMultiTexGenfvEXT(C.GLenum(arg0), C.GLenum(arg1), C.GLenum(arg2), (*C.GLfloat)(arg3));
}

//extern void glMultiTexGeniEXT (GLenum, GLenum, GLenum, GLint)
func MultiTexGeniEXT(arg0 GLenum, arg1 GLenum, arg2 GLenum, arg3 GLint) 
{
	C.glMultiTexGeniEXT(C.GLenum(arg0), C.GLenum(arg1), C.GLenum(arg2), C.GLint(arg3));
}

//extern void glMultiTexGenivEXT (GLenum, GLenum, GLenum, const GLint *)
func MultiTexGenivEXT(arg0 GLenum, arg1 GLenum, arg2 GLenum, arg3 *GLint) 
{
	C.glMultiTexGenivEXT(C.GLenum(arg0), C.GLenum(arg1), C.GLenum(arg2), (*C.GLint)(arg3));
}

//extern void glGetMultiTexEnvfvEXT (GLenum, GLenum, GLenum, GLfloat *)
func GetMultiTexEnvfvEXT(arg0 GLenum, arg1 GLenum, arg2 GLenum, arg3 *GLfloat) 
{
	C.glGetMultiTexEnvfvEXT(C.GLenum(arg0), C.GLenum(arg1), C.GLenum(arg2), (*C.GLfloat)(arg3));
}

//extern void glGetMultiTexEnvivEXT (GLenum, GLenum, GLenum, GLint *)
func GetMultiTexEnvivEXT(arg0 GLenum, arg1 GLenum, arg2 GLenum, arg3 *GLint) 
{
	C.glGetMultiTexEnvivEXT(C.GLenum(arg0), C.GLenum(arg1), C.GLenum(arg2), (*C.GLint)(arg3));
}

//extern void glGetMultiTexGendvEXT (GLenum, GLenum, GLenum, GLdouble *)
func GetMultiTexGendvEXT(arg0 GLenum, arg1 GLenum, arg2 GLenum, arg3 *GLdouble) 
{
	C.glGetMultiTexGendvEXT(C.GLenum(arg0), C.GLenum(arg1), C.GLenum(arg2), (*C.GLdouble)(arg3));
}

//extern void glGetMultiTexGenfvEXT (GLenum, GLenum, GLenum, GLfloat *)
func GetMultiTexGenfvEXT(arg0 GLenum, arg1 GLenum, arg2 GLenum, arg3 *GLfloat) 
{
	C.glGetMultiTexGenfvEXT(C.GLenum(arg0), C.GLenum(arg1), C.GLenum(arg2), (*C.GLfloat)(arg3));
}

//extern void glGetMultiTexGenivEXT (GLenum, GLenum, GLenum, GLint *)
func GetMultiTexGenivEXT(arg0 GLenum, arg1 GLenum, arg2 GLenum, arg3 *GLint) 
{
	C.glGetMultiTexGenivEXT(C.GLenum(arg0), C.GLenum(arg1), C.GLenum(arg2), (*C.GLint)(arg3));
}

//extern void glGetFloatIndexedvEXT (GLenum, GLuint, GLfloat *)
func GetFloatIndexedvEXT(arg0 GLenum, arg1 GLuint, arg2 *GLfloat) 
{
	C.glGetFloatIndexedvEXT(C.GLenum(arg0), C.GLuint(arg1), (*C.GLfloat)(arg2));
}

//extern void glGetDoubleIndexedvEXT (GLenum, GLuint, GLdouble *)
func GetDoubleIndexedvEXT(arg0 GLenum, arg1 GLuint, arg2 *GLdouble) 
{
	C.glGetDoubleIndexedvEXT(C.GLenum(arg0), C.GLuint(arg1), (*C.GLdouble)(arg2));
}

//extern void glGetPointerIndexedvEXT (GLenum, GLuint, GLvoid* *)
func GetPointerIndexedvEXT(arg0 GLenum, arg1 GLuint, arg2 *unsafe.Pointer) 
{
	C.glGetPointerIndexedvEXT(C.GLenum(arg0), C.GLuint(arg1), arg2);
}

//extern void glCompressedTextureImage3DEXT (GLuint, GLenum, GLint, GLenum, GLsizei, GLsizei, GLsizei, GLint, GLsizei, const GLvoid *)
func CompressedTextureImage3DEXT(arg0 GLuint, arg1 GLenum, arg2 GLint, arg3 GLenum, arg4 GLsizei, arg5 GLsizei, arg6 GLsizei, arg7 GLint, arg8 GLsizei, arg9 unsafe.Pointer) 
{
	C.glCompressedTextureImage3DEXT(C.GLuint(arg0), C.GLenum(arg1), C.GLint(arg2), C.GLenum(arg3), C.GLsizei(arg4), C.GLsizei(arg5), C.GLsizei(arg6), C.GLint(arg7), C.GLsizei(arg8), arg9);
}

//extern void glCompressedTextureImage2DEXT (GLuint, GLenum, GLint, GLenum, GLsizei, GLsizei, GLint, GLsizei, const GLvoid *)
func CompressedTextureImage2DEXT(arg0 GLuint, arg1 GLenum, arg2 GLint, arg3 GLenum, arg4 GLsizei, arg5 GLsizei, arg6 GLint, arg7 GLsizei, arg8 unsafe.Pointer) 
{
	C.glCompressedTextureImage2DEXT(C.GLuint(arg0), C.GLenum(arg1), C.GLint(arg2), C.GLenum(arg3), C.GLsizei(arg4), C.GLsizei(arg5), C.GLint(arg6), C.GLsizei(arg7), arg8);
}

//extern void glCompressedTextureImage1DEXT (GLuint, GLenum, GLint, GLenum, GLsizei, GLint, GLsizei, const GLvoid *)
func CompressedTextureImage1DEXT(arg0 GLuint, arg1 GLenum, arg2 GLint, arg3 GLenum, arg4 GLsizei, arg5 GLint, arg6 GLsizei, arg7 unsafe.Pointer) 
{
	C.glCompressedTextureImage1DEXT(C.GLuint(arg0), C.GLenum(arg1), C.GLint(arg2), C.GLenum(arg3), C.GLsizei(arg4), C.GLint(arg5), C.GLsizei(arg6), arg7);
}

//extern void glCompressedTextureSubImage3DEXT (GLuint, GLenum, GLint, GLint, GLint, GLint, GLsizei, GLsizei, GLsizei, GLenum, GLsizei, const GLvoid *)
func CompressedTextureSubImage3DEXT(arg0 GLuint, arg1 GLenum, arg2 GLint, arg3 GLint, arg4 GLint, arg5 GLint, arg6 GLsizei, arg7 GLsizei, arg8 GLsizei, arg9 GLenum, arg10 GLsizei, arg11 unsafe.Pointer) 
{
	C.glCompressedTextureSubImage3DEXT(C.GLuint(arg0), C.GLenum(arg1), C.GLint(arg2), C.GLint(arg3), C.GLint(arg4), C.GLint(arg5), C.GLsizei(arg6), C.GLsizei(arg7), C.GLsizei(arg8), C.GLenum(arg9), C.GLsizei(arg10), arg11);
}

//extern void glCompressedTextureSubImage2DEXT (GLuint, GLenum, GLint, GLint, GLint, GLsizei, GLsizei, GLenum, GLsizei, const GLvoid *)
func CompressedTextureSubImage2DEXT(arg0 GLuint, arg1 GLenum, arg2 GLint, arg3 GLint, arg4 GLint, arg5 GLsizei, arg6 GLsizei, arg7 GLenum, arg8 GLsizei, arg9 unsafe.Pointer) 
{
	C.glCompressedTextureSubImage2DEXT(C.GLuint(arg0), C.GLenum(arg1), C.GLint(arg2), C.GLint(arg3), C.GLint(arg4), C.GLsizei(arg5), C.GLsizei(arg6), C.GLenum(arg7), C.GLsizei(arg8), arg9);
}

//extern void glCompressedTextureSubImage1DEXT (GLuint, GLenum, GLint, GLint, GLsizei, GLenum, GLsizei, const GLvoid *)
func CompressedTextureSubImage1DEXT(arg0 GLuint, arg1 GLenum, arg2 GLint, arg3 GLint, arg4 GLsizei, arg5 GLenum, arg6 GLsizei, arg7 unsafe.Pointer) 
{
	C.glCompressedTextureSubImage1DEXT(C.GLuint(arg0), C.GLenum(arg1), C.GLint(arg2), C.GLint(arg3), C.GLsizei(arg4), C.GLenum(arg5), C.GLsizei(arg6), arg7);
}

//extern void glGetCompressedTextureImageEXT (GLuint, GLenum, GLint, GLvoid *)
func GetCompressedTextureImageEXT(arg0 GLuint, arg1 GLenum, arg2 GLint, arg3 unsafe.Pointer) 
{
	C.glGetCompressedTextureImageEXT(C.GLuint(arg0), C.GLenum(arg1), C.GLint(arg2), arg3);
}

//extern void glCompressedMultiTexImage3DEXT (GLenum, GLenum, GLint, GLenum, GLsizei, GLsizei, GLsizei, GLint, GLsizei, const GLvoid *)
func CompressedMultiTexImage3DEXT(arg0 GLenum, arg1 GLenum, arg2 GLint, arg3 GLenum, arg4 GLsizei, arg5 GLsizei, arg6 GLsizei, arg7 GLint, arg8 GLsizei, arg9 unsafe.Pointer) 
{
	C.glCompressedMultiTexImage3DEXT(C.GLenum(arg0), C.GLenum(arg1), C.GLint(arg2), C.GLenum(arg3), C.GLsizei(arg4), C.GLsizei(arg5), C.GLsizei(arg6), C.GLint(arg7), C.GLsizei(arg8), arg9);
}

//extern void glCompressedMultiTexImage2DEXT (GLenum, GLenum, GLint, GLenum, GLsizei, GLsizei, GLint, GLsizei, const GLvoid *)
func CompressedMultiTexImage2DEXT(arg0 GLenum, arg1 GLenum, arg2 GLint, arg3 GLenum, arg4 GLsizei, arg5 GLsizei, arg6 GLint, arg7 GLsizei, arg8 unsafe.Pointer) 
{
	C.glCompressedMultiTexImage2DEXT(C.GLenum(arg0), C.GLenum(arg1), C.GLint(arg2), C.GLenum(arg3), C.GLsizei(arg4), C.GLsizei(arg5), C.GLint(arg6), C.GLsizei(arg7), arg8);
}

//extern void glCompressedMultiTexImage1DEXT (GLenum, GLenum, GLint, GLenum, GLsizei, GLint, GLsizei, const GLvoid *)
func CompressedMultiTexImage1DEXT(arg0 GLenum, arg1 GLenum, arg2 GLint, arg3 GLenum, arg4 GLsizei, arg5 GLint, arg6 GLsizei, arg7 unsafe.Pointer) 
{
	C.glCompressedMultiTexImage1DEXT(C.GLenum(arg0), C.GLenum(arg1), C.GLint(arg2), C.GLenum(arg3), C.GLsizei(arg4), C.GLint(arg5), C.GLsizei(arg6), arg7);
}

//extern void glCompressedMultiTexSubImage3DEXT (GLenum, GLenum, GLint, GLint, GLint, GLint, GLsizei, GLsizei, GLsizei, GLenum, GLsizei, const GLvoid *)
func CompressedMultiTexSubImage3DEXT(arg0 GLenum, arg1 GLenum, arg2 GLint, arg3 GLint, arg4 GLint, arg5 GLint, arg6 GLsizei, arg7 GLsizei, arg8 GLsizei, arg9 GLenum, arg10 GLsizei, arg11 unsafe.Pointer) 
{
	C.glCompressedMultiTexSubImage3DEXT(C.GLenum(arg0), C.GLenum(arg1), C.GLint(arg2), C.GLint(arg3), C.GLint(arg4), C.GLint(arg5), C.GLsizei(arg6), C.GLsizei(arg7), C.GLsizei(arg8), C.GLenum(arg9), C.GLsizei(arg10), arg11);
}

//extern void glCompressedMultiTexSubImage2DEXT (GLenum, GLenum, GLint, GLint, GLint, GLsizei, GLsizei, GLenum, GLsizei, const GLvoid *)
func CompressedMultiTexSubImage2DEXT(arg0 GLenum, arg1 GLenum, arg2 GLint, arg3 GLint, arg4 GLint, arg5 GLsizei, arg6 GLsizei, arg7 GLenum, arg8 GLsizei, arg9 unsafe.Pointer) 
{
	C.glCompressedMultiTexSubImage2DEXT(C.GLenum(arg0), C.GLenum(arg1), C.GLint(arg2), C.GLint(arg3), C.GLint(arg4), C.GLsizei(arg5), C.GLsizei(arg6), C.GLenum(arg7), C.GLsizei(arg8), arg9);
}

//extern void glCompressedMultiTexSubImage1DEXT (GLenum, GLenum, GLint, GLint, GLsizei, GLenum, GLsizei, const GLvoid *)
func CompressedMultiTexSubImage1DEXT(arg0 GLenum, arg1 GLenum, arg2 GLint, arg3 GLint, arg4 GLsizei, arg5 GLenum, arg6 GLsizei, arg7 unsafe.Pointer) 
{
	C.glCompressedMultiTexSubImage1DEXT(C.GLenum(arg0), C.GLenum(arg1), C.GLint(arg2), C.GLint(arg3), C.GLsizei(arg4), C.GLenum(arg5), C.GLsizei(arg6), arg7);
}

//extern void glGetCompressedMultiTexImageEXT (GLenum, GLenum, GLint, GLvoid *)
func GetCompressedMultiTexImageEXT(arg0 GLenum, arg1 GLenum, arg2 GLint, arg3 unsafe.Pointer) 
{
	C.glGetCompressedMultiTexImageEXT(C.GLenum(arg0), C.GLenum(arg1), C.GLint(arg2), arg3);
}

//extern void glNamedProgramStringEXT (GLuint, GLenum, GLenum, GLsizei, const GLvoid *)
func NamedProgramStringEXT(arg0 GLuint, arg1 GLenum, arg2 GLenum, arg3 GLsizei, arg4 unsafe.Pointer) 
{
	C.glNamedProgramStringEXT(C.GLuint(arg0), C.GLenum(arg1), C.GLenum(arg2), C.GLsizei(arg3), arg4);
}

//extern void glNamedProgramLocalParameter4dEXT (GLuint, GLenum, GLuint, GLdouble, GLdouble, GLdouble, GLdouble)
func NamedProgramLocalParameter4dEXT(arg0 GLuint, arg1 GLenum, arg2 GLuint, arg3 GLdouble, arg4 GLdouble, arg5 GLdouble, arg6 GLdouble) 
{
	C.glNamedProgramLocalParameter4dEXT(C.GLuint(arg0), C.GLenum(arg1), C.GLuint(arg2), C.GLdouble(arg3), C.GLdouble(arg4), C.GLdouble(arg5), C.GLdouble(arg6));
}

//extern void glNamedProgramLocalParameter4dvEXT (GLuint, GLenum, GLuint, const GLdouble *)
func NamedProgramLocalParameter4dvEXT(arg0 GLuint, arg1 GLenum, arg2 GLuint, arg3 *GLdouble) 
{
	C.glNamedProgramLocalParameter4dvEXT(C.GLuint(arg0), C.GLenum(arg1), C.GLuint(arg2), (*C.GLdouble)(arg3));
}

//extern void glNamedProgramLocalParameter4fEXT (GLuint, GLenum, GLuint, GLfloat, GLfloat, GLfloat, GLfloat)
func NamedProgramLocalParameter4fEXT(arg0 GLuint, arg1 GLenum, arg2 GLuint, arg3 GLfloat, arg4 GLfloat, arg5 GLfloat, arg6 GLfloat) 
{
	C.glNamedProgramLocalParameter4fEXT(C.GLuint(arg0), C.GLenum(arg1), C.GLuint(arg2), C.GLfloat(arg3), C.GLfloat(arg4), C.GLfloat(arg5), C.GLfloat(arg6));
}

//extern void glNamedProgramLocalParameter4fvEXT (GLuint, GLenum, GLuint, const GLfloat *)
func NamedProgramLocalParameter4fvEXT(arg0 GLuint, arg1 GLenum, arg2 GLuint, arg3 *GLfloat) 
{
	C.glNamedProgramLocalParameter4fvEXT(C.GLuint(arg0), C.GLenum(arg1), C.GLuint(arg2), (*C.GLfloat)(arg3));
}

//extern void glGetNamedProgramLocalParameterdvEXT (GLuint, GLenum, GLuint, GLdouble *)
func GetNamedProgramLocalParameterdvEXT(arg0 GLuint, arg1 GLenum, arg2 GLuint, arg3 *GLdouble) 
{
	C.glGetNamedProgramLocalParameterdvEXT(C.GLuint(arg0), C.GLenum(arg1), C.GLuint(arg2), (*C.GLdouble)(arg3));
}

//extern void glGetNamedProgramLocalParameterfvEXT (GLuint, GLenum, GLuint, GLfloat *)
func GetNamedProgramLocalParameterfvEXT(arg0 GLuint, arg1 GLenum, arg2 GLuint, arg3 *GLfloat) 
{
	C.glGetNamedProgramLocalParameterfvEXT(C.GLuint(arg0), C.GLenum(arg1), C.GLuint(arg2), (*C.GLfloat)(arg3));
}

//extern void glGetNamedProgramivEXT (GLuint, GLenum, GLenum, GLint *)
func GetNamedProgramivEXT(arg0 GLuint, arg1 GLenum, arg2 GLenum, arg3 *GLint) 
{
	C.glGetNamedProgramivEXT(C.GLuint(arg0), C.GLenum(arg1), C.GLenum(arg2), (*C.GLint)(arg3));
}

//extern void glGetNamedProgramStringEXT (GLuint, GLenum, GLenum, GLvoid *)
func GetNamedProgramStringEXT(arg0 GLuint, arg1 GLenum, arg2 GLenum, arg3 unsafe.Pointer) 
{
	C.glGetNamedProgramStringEXT(C.GLuint(arg0), C.GLenum(arg1), C.GLenum(arg2), arg3);
}

//extern void glNamedProgramLocalParameters4fvEXT (GLuint, GLenum, GLuint, GLsizei, const GLfloat *)
func NamedProgramLocalParameters4fvEXT(arg0 GLuint, arg1 GLenum, arg2 GLuint, arg3 GLsizei, arg4 *GLfloat) 
{
	C.glNamedProgramLocalParameters4fvEXT(C.GLuint(arg0), C.GLenum(arg1), C.GLuint(arg2), C.GLsizei(arg3), (*C.GLfloat)(arg4));
}

//extern void glNamedProgramLocalParameterI4iEXT (GLuint, GLenum, GLuint, GLint, GLint, GLint, GLint)
func NamedProgramLocalParameterI4iEXT(arg0 GLuint, arg1 GLenum, arg2 GLuint, arg3 GLint, arg4 GLint, arg5 GLint, arg6 GLint) 
{
	C.glNamedProgramLocalParameterI4iEXT(C.GLuint(arg0), C.GLenum(arg1), C.GLuint(arg2), C.GLint(arg3), C.GLint(arg4), C.GLint(arg5), C.GLint(arg6));
}

//extern void glNamedProgramLocalParameterI4ivEXT (GLuint, GLenum, GLuint, const GLint *)
func NamedProgramLocalParameterI4ivEXT(arg0 GLuint, arg1 GLenum, arg2 GLuint, arg3 *GLint) 
{
	C.glNamedProgramLocalParameterI4ivEXT(C.GLuint(arg0), C.GLenum(arg1), C.GLuint(arg2), (*C.GLint)(arg3));
}

//extern void glNamedProgramLocalParametersI4ivEXT (GLuint, GLenum, GLuint, GLsizei, const GLint *)
func NamedProgramLocalParametersI4ivEXT(arg0 GLuint, arg1 GLenum, arg2 GLuint, arg3 GLsizei, arg4 *GLint) 
{
	C.glNamedProgramLocalParametersI4ivEXT(C.GLuint(arg0), C.GLenum(arg1), C.GLuint(arg2), C.GLsizei(arg3), (*C.GLint)(arg4));
}

//extern void glNamedProgramLocalParameterI4uiEXT (GLuint, GLenum, GLuint, GLuint, GLuint, GLuint, GLuint)
func NamedProgramLocalParameterI4uiEXT(arg0 GLuint, arg1 GLenum, arg2 GLuint, arg3 GLuint, arg4 GLuint, arg5 GLuint, arg6 GLuint) 
{
	C.glNamedProgramLocalParameterI4uiEXT(C.GLuint(arg0), C.GLenum(arg1), C.GLuint(arg2), C.GLuint(arg3), C.GLuint(arg4), C.GLuint(arg5), C.GLuint(arg6));
}

//extern void glNamedProgramLocalParameterI4uivEXT (GLuint, GLenum, GLuint, const GLuint *)
func NamedProgramLocalParameterI4uivEXT(arg0 GLuint, arg1 GLenum, arg2 GLuint, arg3 *GLuint) 
{
	C.glNamedProgramLocalParameterI4uivEXT(C.GLuint(arg0), C.GLenum(arg1), C.GLuint(arg2), (*C.GLuint)(arg3));
}

//extern void glNamedProgramLocalParametersI4uivEXT (GLuint, GLenum, GLuint, GLsizei, const GLuint *)
func NamedProgramLocalParametersI4uivEXT(arg0 GLuint, arg1 GLenum, arg2 GLuint, arg3 GLsizei, arg4 *GLuint) 
{
	C.glNamedProgramLocalParametersI4uivEXT(C.GLuint(arg0), C.GLenum(arg1), C.GLuint(arg2), C.GLsizei(arg3), (*C.GLuint)(arg4));
}

//extern void glGetNamedProgramLocalParameterIivEXT (GLuint, GLenum, GLuint, GLint *)
func GetNamedProgramLocalParameterIivEXT(arg0 GLuint, arg1 GLenum, arg2 GLuint, arg3 *GLint) 
{
	C.glGetNamedProgramLocalParameterIivEXT(C.GLuint(arg0), C.GLenum(arg1), C.GLuint(arg2), (*C.GLint)(arg3));
}

//extern void glGetNamedProgramLocalParameterIuivEXT (GLuint, GLenum, GLuint, GLuint *)
func GetNamedProgramLocalParameterIuivEXT(arg0 GLuint, arg1 GLenum, arg2 GLuint, arg3 *GLuint) 
{
	C.glGetNamedProgramLocalParameterIuivEXT(C.GLuint(arg0), C.GLenum(arg1), C.GLuint(arg2), (*C.GLuint)(arg3));
}

//extern void glTextureParameterIivEXT (GLuint, GLenum, GLenum, const GLint *)
func TextureParameterIivEXT(arg0 GLuint, arg1 GLenum, arg2 GLenum, arg3 *GLint) 
{
	C.glTextureParameterIivEXT(C.GLuint(arg0), C.GLenum(arg1), C.GLenum(arg2), (*C.GLint)(arg3));
}

//extern void glTextureParameterIuivEXT (GLuint, GLenum, GLenum, const GLuint *)
func TextureParameterIuivEXT(arg0 GLuint, arg1 GLenum, arg2 GLenum, arg3 *GLuint) 
{
	C.glTextureParameterIuivEXT(C.GLuint(arg0), C.GLenum(arg1), C.GLenum(arg2), (*C.GLuint)(arg3));
}

//extern void glGetTextureParameterIivEXT (GLuint, GLenum, GLenum, GLint *)
func GetTextureParameterIivEXT(arg0 GLuint, arg1 GLenum, arg2 GLenum, arg3 *GLint) 
{
	C.glGetTextureParameterIivEXT(C.GLuint(arg0), C.GLenum(arg1), C.GLenum(arg2), (*C.GLint)(arg3));
}

//extern void glGetTextureParameterIuivEXT (GLuint, GLenum, GLenum, GLuint *)
func GetTextureParameterIuivEXT(arg0 GLuint, arg1 GLenum, arg2 GLenum, arg3 *GLuint) 
{
	C.glGetTextureParameterIuivEXT(C.GLuint(arg0), C.GLenum(arg1), C.GLenum(arg2), (*C.GLuint)(arg3));
}

//extern void glMultiTexParameterIivEXT (GLenum, GLenum, GLenum, const GLint *)
func MultiTexParameterIivEXT(arg0 GLenum, arg1 GLenum, arg2 GLenum, arg3 *GLint) 
{
	C.glMultiTexParameterIivEXT(C.GLenum(arg0), C.GLenum(arg1), C.GLenum(arg2), (*C.GLint)(arg3));
}

//extern void glMultiTexParameterIuivEXT (GLenum, GLenum, GLenum, const GLuint *)
func MultiTexParameterIuivEXT(arg0 GLenum, arg1 GLenum, arg2 GLenum, arg3 *GLuint) 
{
	C.glMultiTexParameterIuivEXT(C.GLenum(arg0), C.GLenum(arg1), C.GLenum(arg2), (*C.GLuint)(arg3));
}

//extern void glGetMultiTexParameterIivEXT (GLenum, GLenum, GLenum, GLint *)
func GetMultiTexParameterIivEXT(arg0 GLenum, arg1 GLenum, arg2 GLenum, arg3 *GLint) 
{
	C.glGetMultiTexParameterIivEXT(C.GLenum(arg0), C.GLenum(arg1), C.GLenum(arg2), (*C.GLint)(arg3));
}

//extern void glGetMultiTexParameterIuivEXT (GLenum, GLenum, GLenum, GLuint *)
func GetMultiTexParameterIuivEXT(arg0 GLenum, arg1 GLenum, arg2 GLenum, arg3 *GLuint) 
{
	C.glGetMultiTexParameterIuivEXT(C.GLenum(arg0), C.GLenum(arg1), C.GLenum(arg2), (*C.GLuint)(arg3));
}

//extern void glProgramUniform1fEXT (GLuint, GLint, GLfloat)
func ProgramUniform1fEXT(arg0 GLuint, arg1 GLint, arg2 GLfloat) 
{
	C.glProgramUniform1fEXT(C.GLuint(arg0), C.GLint(arg1), C.GLfloat(arg2));
}

//extern void glProgramUniform2fEXT (GLuint, GLint, GLfloat, GLfloat)
func ProgramUniform2fEXT(arg0 GLuint, arg1 GLint, arg2 GLfloat, arg3 GLfloat) 
{
	C.glProgramUniform2fEXT(C.GLuint(arg0), C.GLint(arg1), C.GLfloat(arg2), C.GLfloat(arg3));
}

//extern void glProgramUniform3fEXT (GLuint, GLint, GLfloat, GLfloat, GLfloat)
func ProgramUniform3fEXT(arg0 GLuint, arg1 GLint, arg2 GLfloat, arg3 GLfloat, arg4 GLfloat) 
{
	C.glProgramUniform3fEXT(C.GLuint(arg0), C.GLint(arg1), C.GLfloat(arg2), C.GLfloat(arg3), C.GLfloat(arg4));
}

//extern void glProgramUniform4fEXT (GLuint, GLint, GLfloat, GLfloat, GLfloat, GLfloat)
func ProgramUniform4fEXT(arg0 GLuint, arg1 GLint, arg2 GLfloat, arg3 GLfloat, arg4 GLfloat, arg5 GLfloat) 
{
	C.glProgramUniform4fEXT(C.GLuint(arg0), C.GLint(arg1), C.GLfloat(arg2), C.GLfloat(arg3), C.GLfloat(arg4), C.GLfloat(arg5));
}

//extern void glProgramUniform1iEXT (GLuint, GLint, GLint)
func ProgramUniform1iEXT(arg0 GLuint, arg1 GLint, arg2 GLint) 
{
	C.glProgramUniform1iEXT(C.GLuint(arg0), C.GLint(arg1), C.GLint(arg2));
}

//extern void glProgramUniform2iEXT (GLuint, GLint, GLint, GLint)
func ProgramUniform2iEXT(arg0 GLuint, arg1 GLint, arg2 GLint, arg3 GLint) 
{
	C.glProgramUniform2iEXT(C.GLuint(arg0), C.GLint(arg1), C.GLint(arg2), C.GLint(arg3));
}

//extern void glProgramUniform3iEXT (GLuint, GLint, GLint, GLint, GLint)
func ProgramUniform3iEXT(arg0 GLuint, arg1 GLint, arg2 GLint, arg3 GLint, arg4 GLint) 
{
	C.glProgramUniform3iEXT(C.GLuint(arg0), C.GLint(arg1), C.GLint(arg2), C.GLint(arg3), C.GLint(arg4));
}

//extern void glProgramUniform4iEXT (GLuint, GLint, GLint, GLint, GLint, GLint)
func ProgramUniform4iEXT(arg0 GLuint, arg1 GLint, arg2 GLint, arg3 GLint, arg4 GLint, arg5 GLint) 
{
	C.glProgramUniform4iEXT(C.GLuint(arg0), C.GLint(arg1), C.GLint(arg2), C.GLint(arg3), C.GLint(arg4), C.GLint(arg5));
}

//extern void glProgramUniform1fvEXT (GLuint, GLint, GLsizei, const GLfloat *)
func ProgramUniform1fvEXT(arg0 GLuint, arg1 GLint, arg2 GLsizei, arg3 *GLfloat) 
{
	C.glProgramUniform1fvEXT(C.GLuint(arg0), C.GLint(arg1), C.GLsizei(arg2), (*C.GLfloat)(arg3));
}

//extern void glProgramUniform2fvEXT (GLuint, GLint, GLsizei, const GLfloat *)
func ProgramUniform2fvEXT(arg0 GLuint, arg1 GLint, arg2 GLsizei, arg3 *GLfloat) 
{
	C.glProgramUniform2fvEXT(C.GLuint(arg0), C.GLint(arg1), C.GLsizei(arg2), (*C.GLfloat)(arg3));
}

//extern void glProgramUniform3fvEXT (GLuint, GLint, GLsizei, const GLfloat *)
func ProgramUniform3fvEXT(arg0 GLuint, arg1 GLint, arg2 GLsizei, arg3 *GLfloat) 
{
	C.glProgramUniform3fvEXT(C.GLuint(arg0), C.GLint(arg1), C.GLsizei(arg2), (*C.GLfloat)(arg3));
}

//extern void glProgramUniform4fvEXT (GLuint, GLint, GLsizei, const GLfloat *)
func ProgramUniform4fvEXT(arg0 GLuint, arg1 GLint, arg2 GLsizei, arg3 *GLfloat) 
{
	C.glProgramUniform4fvEXT(C.GLuint(arg0), C.GLint(arg1), C.GLsizei(arg2), (*C.GLfloat)(arg3));
}

//extern void glProgramUniform1ivEXT (GLuint, GLint, GLsizei, const GLint *)
func ProgramUniform1ivEXT(arg0 GLuint, arg1 GLint, arg2 GLsizei, arg3 *GLint) 
{
	C.glProgramUniform1ivEXT(C.GLuint(arg0), C.GLint(arg1), C.GLsizei(arg2), (*C.GLint)(arg3));
}

//extern void glProgramUniform2ivEXT (GLuint, GLint, GLsizei, const GLint *)
func ProgramUniform2ivEXT(arg0 GLuint, arg1 GLint, arg2 GLsizei, arg3 *GLint) 
{
	C.glProgramUniform2ivEXT(C.GLuint(arg0), C.GLint(arg1), C.GLsizei(arg2), (*C.GLint)(arg3));
}

//extern void glProgramUniform3ivEXT (GLuint, GLint, GLsizei, const GLint *)
func ProgramUniform3ivEXT(arg0 GLuint, arg1 GLint, arg2 GLsizei, arg3 *GLint) 
{
	C.glProgramUniform3ivEXT(C.GLuint(arg0), C.GLint(arg1), C.GLsizei(arg2), (*C.GLint)(arg3));
}

//extern void glProgramUniform4ivEXT (GLuint, GLint, GLsizei, const GLint *)
func ProgramUniform4ivEXT(arg0 GLuint, arg1 GLint, arg2 GLsizei, arg3 *GLint) 
{
	C.glProgramUniform4ivEXT(C.GLuint(arg0), C.GLint(arg1), C.GLsizei(arg2), (*C.GLint)(arg3));
}

//extern void glProgramUniformMatrix2fvEXT (GLuint, GLint, GLsizei, GLboolean, const GLfloat *)
func ProgramUniformMatrix2fvEXT(arg0 GLuint, arg1 GLint, arg2 GLsizei, arg3 GLboolean, arg4 *GLfloat) 
{
	C.glProgramUniformMatrix2fvEXT(C.GLuint(arg0), C.GLint(arg1), C.GLsizei(arg2), C.GLboolean(arg3), (*C.GLfloat)(arg4));
}

//extern void glProgramUniformMatrix3fvEXT (GLuint, GLint, GLsizei, GLboolean, const GLfloat *)
func ProgramUniformMatrix3fvEXT(arg0 GLuint, arg1 GLint, arg2 GLsizei, arg3 GLboolean, arg4 *GLfloat) 
{
	C.glProgramUniformMatrix3fvEXT(C.GLuint(arg0), C.GLint(arg1), C.GLsizei(arg2), C.GLboolean(arg3), (*C.GLfloat)(arg4));
}

//extern void glProgramUniformMatrix4fvEXT (GLuint, GLint, GLsizei, GLboolean, const GLfloat *)
func ProgramUniformMatrix4fvEXT(arg0 GLuint, arg1 GLint, arg2 GLsizei, arg3 GLboolean, arg4 *GLfloat) 
{
	C.glProgramUniformMatrix4fvEXT(C.GLuint(arg0), C.GLint(arg1), C.GLsizei(arg2), C.GLboolean(arg3), (*C.GLfloat)(arg4));
}

//extern void glProgramUniformMatrix2x3fvEXT (GLuint, GLint, GLsizei, GLboolean, const GLfloat *)
func ProgramUniformMatrix2x3fvEXT(arg0 GLuint, arg1 GLint, arg2 GLsizei, arg3 GLboolean, arg4 *GLfloat) 
{
	C.glProgramUniformMatrix2x3fvEXT(C.GLuint(arg0), C.GLint(arg1), C.GLsizei(arg2), C.GLboolean(arg3), (*C.GLfloat)(arg4));
}

//extern void glProgramUniformMatrix3x2fvEXT (GLuint, GLint, GLsizei, GLboolean, const GLfloat *)
func ProgramUniformMatrix3x2fvEXT(arg0 GLuint, arg1 GLint, arg2 GLsizei, arg3 GLboolean, arg4 *GLfloat) 
{
	C.glProgramUniformMatrix3x2fvEXT(C.GLuint(arg0), C.GLint(arg1), C.GLsizei(arg2), C.GLboolean(arg3), (*C.GLfloat)(arg4));
}

//extern void glProgramUniformMatrix2x4fvEXT (GLuint, GLint, GLsizei, GLboolean, const GLfloat *)
func ProgramUniformMatrix2x4fvEXT(arg0 GLuint, arg1 GLint, arg2 GLsizei, arg3 GLboolean, arg4 *GLfloat) 
{
	C.glProgramUniformMatrix2x4fvEXT(C.GLuint(arg0), C.GLint(arg1), C.GLsizei(arg2), C.GLboolean(arg3), (*C.GLfloat)(arg4));
}

//extern void glProgramUniformMatrix4x2fvEXT (GLuint, GLint, GLsizei, GLboolean, const GLfloat *)
func ProgramUniformMatrix4x2fvEXT(arg0 GLuint, arg1 GLint, arg2 GLsizei, arg3 GLboolean, arg4 *GLfloat) 
{
	C.glProgramUniformMatrix4x2fvEXT(C.GLuint(arg0), C.GLint(arg1), C.GLsizei(arg2), C.GLboolean(arg3), (*C.GLfloat)(arg4));
}

//extern void glProgramUniformMatrix3x4fvEXT (GLuint, GLint, GLsizei, GLboolean, const GLfloat *)
func ProgramUniformMatrix3x4fvEXT(arg0 GLuint, arg1 GLint, arg2 GLsizei, arg3 GLboolean, arg4 *GLfloat) 
{
	C.glProgramUniformMatrix3x4fvEXT(C.GLuint(arg0), C.GLint(arg1), C.GLsizei(arg2), C.GLboolean(arg3), (*C.GLfloat)(arg4));
}

//extern void glProgramUniformMatrix4x3fvEXT (GLuint, GLint, GLsizei, GLboolean, const GLfloat *)
func ProgramUniformMatrix4x3fvEXT(arg0 GLuint, arg1 GLint, arg2 GLsizei, arg3 GLboolean, arg4 *GLfloat) 
{
	C.glProgramUniformMatrix4x3fvEXT(C.GLuint(arg0), C.GLint(arg1), C.GLsizei(arg2), C.GLboolean(arg3), (*C.GLfloat)(arg4));
}

//extern void glProgramUniform1uiEXT (GLuint, GLint, GLuint)
func ProgramUniform1uiEXT(arg0 GLuint, arg1 GLint, arg2 GLuint) 
{
	C.glProgramUniform1uiEXT(C.GLuint(arg0), C.GLint(arg1), C.GLuint(arg2));
}

//extern void glProgramUniform2uiEXT (GLuint, GLint, GLuint, GLuint)
func ProgramUniform2uiEXT(arg0 GLuint, arg1 GLint, arg2 GLuint, arg3 GLuint) 
{
	C.glProgramUniform2uiEXT(C.GLuint(arg0), C.GLint(arg1), C.GLuint(arg2), C.GLuint(arg3));
}

//extern void glProgramUniform3uiEXT (GLuint, GLint, GLuint, GLuint, GLuint)
func ProgramUniform3uiEXT(arg0 GLuint, arg1 GLint, arg2 GLuint, arg3 GLuint, arg4 GLuint) 
{
	C.glProgramUniform3uiEXT(C.GLuint(arg0), C.GLint(arg1), C.GLuint(arg2), C.GLuint(arg3), C.GLuint(arg4));
}

//extern void glProgramUniform4uiEXT (GLuint, GLint, GLuint, GLuint, GLuint, GLuint)
func ProgramUniform4uiEXT(arg0 GLuint, arg1 GLint, arg2 GLuint, arg3 GLuint, arg4 GLuint, arg5 GLuint) 
{
	C.glProgramUniform4uiEXT(C.GLuint(arg0), C.GLint(arg1), C.GLuint(arg2), C.GLuint(arg3), C.GLuint(arg4), C.GLuint(arg5));
}

//extern void glProgramUniform1uivEXT (GLuint, GLint, GLsizei, const GLuint *)
func ProgramUniform1uivEXT(arg0 GLuint, arg1 GLint, arg2 GLsizei, arg3 *GLuint) 
{
	C.glProgramUniform1uivEXT(C.GLuint(arg0), C.GLint(arg1), C.GLsizei(arg2), (*C.GLuint)(arg3));
}

//extern void glProgramUniform2uivEXT (GLuint, GLint, GLsizei, const GLuint *)
func ProgramUniform2uivEXT(arg0 GLuint, arg1 GLint, arg2 GLsizei, arg3 *GLuint) 
{
	C.glProgramUniform2uivEXT(C.GLuint(arg0), C.GLint(arg1), C.GLsizei(arg2), (*C.GLuint)(arg3));
}

//extern void glProgramUniform3uivEXT (GLuint, GLint, GLsizei, const GLuint *)
func ProgramUniform3uivEXT(arg0 GLuint, arg1 GLint, arg2 GLsizei, arg3 *GLuint) 
{
	C.glProgramUniform3uivEXT(C.GLuint(arg0), C.GLint(arg1), C.GLsizei(arg2), (*C.GLuint)(arg3));
}

//extern void glProgramUniform4uivEXT (GLuint, GLint, GLsizei, const GLuint *)
func ProgramUniform4uivEXT(arg0 GLuint, arg1 GLint, arg2 GLsizei, arg3 *GLuint) 
{
	C.glProgramUniform4uivEXT(C.GLuint(arg0), C.GLint(arg1), C.GLsizei(arg2), (*C.GLuint)(arg3));
}

//extern void glNamedBufferDataEXT (GLuint, GLsizeiptr, const GLvoid *, GLenum)
func NamedBufferDataEXT(arg0 GLuint, arg1 GLsizeiptr, arg2 unsafe.Pointer, arg3 GLenum) 
{
	C.glNamedBufferDataEXT(C.GLuint(arg0), C.GLsizeiptr(arg1), arg2, C.GLenum(arg3));
}

//extern void glNamedBufferSubDataEXT (GLuint, GLintptr, GLsizeiptr, const GLvoid *)
func NamedBufferSubDataEXT(arg0 GLuint, arg1 GLintptr, arg2 GLsizeiptr, arg3 unsafe.Pointer) 
{
	C.glNamedBufferSubDataEXT(C.GLuint(arg0), C.GLintptr(arg1), C.GLsizeiptr(arg2), arg3);
}

//extern GLvoid* glMapNamedBufferEXT (GLuint, GLenum)
func MapNamedBufferEXT(arg0 GLuint, arg1 GLenum) unsafe.Pointer
{
	return unsafe.Pointer(C.glMapNamedBufferEXT(C.GLuint(arg0), C.GLenum(arg1)));
}

//extern GLboolean glUnmapNamedBufferEXT (GLuint)
func UnmapNamedBufferEXT(arg0 GLuint) GLboolean
{
	return GLboolean(C.glUnmapNamedBufferEXT(C.GLuint(arg0)));
}

//extern void glGetNamedBufferParameterivEXT (GLuint, GLenum, GLint *)
func GetNamedBufferParameterivEXT(arg0 GLuint, arg1 GLenum, arg2 *GLint) 
{
	C.glGetNamedBufferParameterivEXT(C.GLuint(arg0), C.GLenum(arg1), (*C.GLint)(arg2));
}

//extern void glGetNamedBufferPointervEXT (GLuint, GLenum, GLvoid* *)
func GetNamedBufferPointervEXT(arg0 GLuint, arg1 GLenum, arg2 *unsafe.Pointer) 
{
	C.glGetNamedBufferPointervEXT(C.GLuint(arg0), C.GLenum(arg1), arg2);
}

//extern void glGetNamedBufferSubDataEXT (GLuint, GLintptr, GLsizeiptr, GLvoid *)
func GetNamedBufferSubDataEXT(arg0 GLuint, arg1 GLintptr, arg2 GLsizeiptr, arg3 unsafe.Pointer) 
{
	C.glGetNamedBufferSubDataEXT(C.GLuint(arg0), C.GLintptr(arg1), C.GLsizeiptr(arg2), arg3);
}

//extern void glTextureBufferEXT (GLuint, GLenum, GLenum, GLuint)
func TextureBufferEXT(arg0 GLuint, arg1 GLenum, arg2 GLenum, arg3 GLuint) 
{
	C.glTextureBufferEXT(C.GLuint(arg0), C.GLenum(arg1), C.GLenum(arg2), C.GLuint(arg3));
}

//extern void glMultiTexBufferEXT (GLenum, GLenum, GLenum, GLuint)
func MultiTexBufferEXT(arg0 GLenum, arg1 GLenum, arg2 GLenum, arg3 GLuint) 
{
	C.glMultiTexBufferEXT(C.GLenum(arg0), C.GLenum(arg1), C.GLenum(arg2), C.GLuint(arg3));
}

//extern void glNamedRenderbufferStorageEXT (GLuint, GLenum, GLsizei, GLsizei)
func NamedRenderbufferStorageEXT(arg0 GLuint, arg1 GLenum, arg2 GLsizei, arg3 GLsizei) 
{
	C.glNamedRenderbufferStorageEXT(C.GLuint(arg0), C.GLenum(arg1), C.GLsizei(arg2), C.GLsizei(arg3));
}

//extern void glGetNamedRenderbufferParameterivEXT (GLuint, GLenum, GLint *)
func GetNamedRenderbufferParameterivEXT(arg0 GLuint, arg1 GLenum, arg2 *GLint) 
{
	C.glGetNamedRenderbufferParameterivEXT(C.GLuint(arg0), C.GLenum(arg1), (*C.GLint)(arg2));
}

//extern GLenum glCheckNamedFramebufferStatusEXT (GLuint, GLenum)
func CheckNamedFramebufferStatusEXT(arg0 GLuint, arg1 GLenum) GLenum
{
	return GLenum(C.glCheckNamedFramebufferStatusEXT(C.GLuint(arg0), C.GLenum(arg1)));
}

//extern void glNamedFramebufferTexture1DEXT (GLuint, GLenum, GLenum, GLuint, GLint)
func NamedFramebufferTexture1DEXT(arg0 GLuint, arg1 GLenum, arg2 GLenum, arg3 GLuint, arg4 GLint) 
{
	C.glNamedFramebufferTexture1DEXT(C.GLuint(arg0), C.GLenum(arg1), C.GLenum(arg2), C.GLuint(arg3), C.GLint(arg4));
}

//extern void glNamedFramebufferTexture2DEXT (GLuint, GLenum, GLenum, GLuint, GLint)
func NamedFramebufferTexture2DEXT(arg0 GLuint, arg1 GLenum, arg2 GLenum, arg3 GLuint, arg4 GLint) 
{
	C.glNamedFramebufferTexture2DEXT(C.GLuint(arg0), C.GLenum(arg1), C.GLenum(arg2), C.GLuint(arg3), C.GLint(arg4));
}

//extern void glNamedFramebufferTexture3DEXT (GLuint, GLenum, GLenum, GLuint, GLint, GLint)
func NamedFramebufferTexture3DEXT(arg0 GLuint, arg1 GLenum, arg2 GLenum, arg3 GLuint, arg4 GLint, arg5 GLint) 
{
	C.glNamedFramebufferTexture3DEXT(C.GLuint(arg0), C.GLenum(arg1), C.GLenum(arg2), C.GLuint(arg3), C.GLint(arg4), C.GLint(arg5));
}

//extern void glNamedFramebufferRenderbufferEXT (GLuint, GLenum, GLenum, GLuint)
func NamedFramebufferRenderbufferEXT(arg0 GLuint, arg1 GLenum, arg2 GLenum, arg3 GLuint) 
{
	C.glNamedFramebufferRenderbufferEXT(C.GLuint(arg0), C.GLenum(arg1), C.GLenum(arg2), C.GLuint(arg3));
}

//extern void glGetNamedFramebufferAttachmentParameterivEXT (GLuint, GLenum, GLenum, GLint *)
func GetNamedFramebufferAttachmentParameterivEXT(arg0 GLuint, arg1 GLenum, arg2 GLenum, arg3 *GLint) 
{
	C.glGetNamedFramebufferAttachmentParameterivEXT(C.GLuint(arg0), C.GLenum(arg1), C.GLenum(arg2), (*C.GLint)(arg3));
}

//extern void glGenerateTextureMipmapEXT (GLuint, GLenum)
func GenerateTextureMipmapEXT(arg0 GLuint, arg1 GLenum) 
{
	C.glGenerateTextureMipmapEXT(C.GLuint(arg0), C.GLenum(arg1));
}

//extern void glGenerateMultiTexMipmapEXT (GLenum, GLenum)
func GenerateMultiTexMipmapEXT(arg0 GLenum, arg1 GLenum) 
{
	C.glGenerateMultiTexMipmapEXT(C.GLenum(arg0), C.GLenum(arg1));
}

//extern void glFramebufferDrawBufferEXT (GLuint, GLenum)
func FramebufferDrawBufferEXT(arg0 GLuint, arg1 GLenum) 
{
	C.glFramebufferDrawBufferEXT(C.GLuint(arg0), C.GLenum(arg1));
}

//extern void glFramebufferDrawBuffersEXT (GLuint, GLsizei, const GLenum *)
func FramebufferDrawBuffersEXT(arg0 GLuint, arg1 GLsizei, arg2 *GLenum) 
{
	C.glFramebufferDrawBuffersEXT(C.GLuint(arg0), C.GLsizei(arg1), (*C.GLenum)(arg2));
}

//extern void glFramebufferReadBufferEXT (GLuint, GLenum)
func FramebufferReadBufferEXT(arg0 GLuint, arg1 GLenum) 
{
	C.glFramebufferReadBufferEXT(C.GLuint(arg0), C.GLenum(arg1));
}

//extern void glGetFramebufferParameterivEXT (GLuint, GLenum, GLint *)
func GetFramebufferParameterivEXT(arg0 GLuint, arg1 GLenum, arg2 *GLint) 
{
	C.glGetFramebufferParameterivEXT(C.GLuint(arg0), C.GLenum(arg1), (*C.GLint)(arg2));
}

//extern void glNamedRenderbufferStorageMultisampleEXT (GLuint, GLsizei, GLenum, GLsizei, GLsizei)
func NamedRenderbufferStorageMultisampleEXT(arg0 GLuint, arg1 GLsizei, arg2 GLenum, arg3 GLsizei, arg4 GLsizei) 
{
	C.glNamedRenderbufferStorageMultisampleEXT(C.GLuint(arg0), C.GLsizei(arg1), C.GLenum(arg2), C.GLsizei(arg3), C.GLsizei(arg4));
}

//extern void glNamedRenderbufferStorageMultisampleCoverageEXT (GLuint, GLsizei, GLsizei, GLenum, GLsizei, GLsizei)
func NamedRenderbufferStorageMultisampleCoverageEXT(arg0 GLuint, arg1 GLsizei, arg2 GLsizei, arg3 GLenum, arg4 GLsizei, arg5 GLsizei) 
{
	C.glNamedRenderbufferStorageMultisampleCoverageEXT(C.GLuint(arg0), C.GLsizei(arg1), C.GLsizei(arg2), C.GLenum(arg3), C.GLsizei(arg4), C.GLsizei(arg5));
}

//extern void glNamedFramebufferTextureEXT (GLuint, GLenum, GLuint, GLint)
func NamedFramebufferTextureEXT(arg0 GLuint, arg1 GLenum, arg2 GLuint, arg3 GLint) 
{
	C.glNamedFramebufferTextureEXT(C.GLuint(arg0), C.GLenum(arg1), C.GLuint(arg2), C.GLint(arg3));
}

//extern void glNamedFramebufferTextureLayerEXT (GLuint, GLenum, GLuint, GLint, GLint)
func NamedFramebufferTextureLayerEXT(arg0 GLuint, arg1 GLenum, arg2 GLuint, arg3 GLint, arg4 GLint) 
{
	C.glNamedFramebufferTextureLayerEXT(C.GLuint(arg0), C.GLenum(arg1), C.GLuint(arg2), C.GLint(arg3), C.GLint(arg4));
}

//extern void glNamedFramebufferTextureFaceEXT (GLuint, GLenum, GLuint, GLint, GLenum)
func NamedFramebufferTextureFaceEXT(arg0 GLuint, arg1 GLenum, arg2 GLuint, arg3 GLint, arg4 GLenum) 
{
	C.glNamedFramebufferTextureFaceEXT(C.GLuint(arg0), C.GLenum(arg1), C.GLuint(arg2), C.GLint(arg3), C.GLenum(arg4));
}

//extern void glTextureRenderbufferEXT (GLuint, GLenum, GLuint)
func TextureRenderbufferEXT(arg0 GLuint, arg1 GLenum, arg2 GLuint) 
{
	C.glTextureRenderbufferEXT(C.GLuint(arg0), C.GLenum(arg1), C.GLuint(arg2));
}

//extern void glMultiTexRenderbufferEXT (GLenum, GLenum, GLuint)
func MultiTexRenderbufferEXT(arg0 GLenum, arg1 GLenum, arg2 GLuint) 
{
	C.glMultiTexRenderbufferEXT(C.GLenum(arg0), C.GLenum(arg1), C.GLuint(arg2));
}

//extern void glGetMultisamplefvNV (GLenum, GLuint, GLfloat *)
func GetMultisamplefvNV(arg0 GLenum, arg1 GLuint, arg2 *GLfloat) 
{
	C.glGetMultisamplefvNV(C.GLenum(arg0), C.GLuint(arg1), (*C.GLfloat)(arg2));
}

//extern void glSampleMaskIndexedNV (GLuint, GLbitfield)
func SampleMaskIndexedNV(arg0 GLuint, arg1 GLbitfield) 
{
	C.glSampleMaskIndexedNV(C.GLuint(arg0), C.GLbitfield(arg1));
}

//extern void glTexRenderbufferNV (GLenum, GLuint)
func TexRenderbufferNV(arg0 GLenum, arg1 GLuint) 
{
	C.glTexRenderbufferNV(C.GLenum(arg0), C.GLuint(arg1));
}

//extern void glBindTransformFeedbackNV (GLenum, GLuint)
func BindTransformFeedbackNV(arg0 GLenum, arg1 GLuint) 
{
	C.glBindTransformFeedbackNV(C.GLenum(arg0), C.GLuint(arg1));
}

//extern void glDeleteTransformFeedbacksNV (GLsizei, const GLuint *)
func DeleteTransformFeedbacksNV(arg0 GLsizei, arg1 *GLuint) 
{
	C.glDeleteTransformFeedbacksNV(C.GLsizei(arg0), (*C.GLuint)(arg1));
}

//extern void glGenTransformFeedbacksNV (GLsizei, GLuint *)
func GenTransformFeedbacksNV(arg0 GLsizei, arg1 *GLuint) 
{
	C.glGenTransformFeedbacksNV(C.GLsizei(arg0), (*C.GLuint)(arg1));
}

//extern GLboolean glIsTransformFeedbackNV (GLuint)
func IsTransformFeedbackNV(arg0 GLuint) GLboolean
{
	return GLboolean(C.glIsTransformFeedbackNV(C.GLuint(arg0)));
}

//extern void glPauseTransformFeedbackNV (void)
func PauseTransformFeedbackNV() 
{
	C.glPauseTransformFeedbackNV();
}

//extern void glResumeTransformFeedbackNV (void)
func ResumeTransformFeedbackNV() 
{
	C.glResumeTransformFeedbackNV();
}

//extern void glDrawTransformFeedbackNV (GLenum, GLuint)
func DrawTransformFeedbackNV(arg0 GLenum, arg1 GLuint) 
{
	C.glDrawTransformFeedbackNV(C.GLenum(arg0), C.GLuint(arg1));
}

//extern void glBufferAddressRangeNV (GLenum pname, GLuint index, GLuint64EXT address, GLsizeiptr length)
func BufferAddressRangeNV(pname GLenum, index GLuint, address GLuint64EXT, length GLsizeiptr) 
{
	C.glBufferAddressRangeNV(C.GLenum(pname), C.GLuint(index), C.GLuint64EXT(address), C.GLsizeiptr(length));
}

//extern void glVertexFormatNV (GLint size, GLenum type, GLsizei stride)
func VertexFormatNV(size GLint, type_ GLenum, stride GLsizei) 
{
	C.glVertexFormatNV(C.GLint(size), C.GLenum(type_), C.GLsizei(stride));
}

//extern void glNormalFormatNV (GLenum type, GLsizei stride)
func NormalFormatNV(type_ GLenum, stride GLsizei) 
{
	C.glNormalFormatNV(C.GLenum(type_), C.GLsizei(stride));
}

//extern void glColorFormatNV (GLint size, GLenum type, GLsizei stride)
func ColorFormatNV(size GLint, type_ GLenum, stride GLsizei) 
{
	C.glColorFormatNV(C.GLint(size), C.GLenum(type_), C.GLsizei(stride));
}

//extern void glIndexFormatNV (GLenum type, GLsizei stride)
func IndexFormatNV(type_ GLenum, stride GLsizei) 
{
	C.glIndexFormatNV(C.GLenum(type_), C.GLsizei(stride));
}

//extern void glTexCoordFormatNV (GLint size, GLenum type, GLsizei stride)
func TexCoordFormatNV(size GLint, type_ GLenum, stride GLsizei) 
{
	C.glTexCoordFormatNV(C.GLint(size), C.GLenum(type_), C.GLsizei(stride));
}

//extern void glEdgeFlagFormatNV (GLsizei stride)
func EdgeFlagFormatNV(stride GLsizei) 
{
	C.glEdgeFlagFormatNV(C.GLsizei(stride));
}

//extern void glSecondaryColorFormatNV (GLint size, GLenum type, GLsizei stride)
func SecondaryColorFormatNV(size GLint, type_ GLenum, stride GLsizei) 
{
	C.glSecondaryColorFormatNV(C.GLint(size), C.GLenum(type_), C.GLsizei(stride));
}

//extern void glFogCoordFormatNV (GLenum type, GLsizei stride)
func FogCoordFormatNV(type_ GLenum, stride GLsizei) 
{
	C.glFogCoordFormatNV(C.GLenum(type_), C.GLsizei(stride));
}

//extern void glVertexAttribFormatNV (GLuint index, GLint size, GLenum type, GLboolean normalized, GLsizei stride)
func VertexAttribFormatNV(index GLuint, size GLint, type_ GLenum, normalized GLboolean, stride GLsizei) 
{
	C.glVertexAttribFormatNV(C.GLuint(index), C.GLint(size), C.GLenum(type_), C.GLboolean(normalized), C.GLsizei(stride));
}

//extern void glVertexAttribIFormatNV (GLuint index, GLint size, GLenum type, GLsizei stride)
func VertexAttribIFormatNV(index GLuint, size GLint, type_ GLenum, stride GLsizei) 
{
	C.glVertexAttribIFormatNV(C.GLuint(index), C.GLint(size), C.GLenum(type_), C.GLsizei(stride));
}

//extern void glGetIntegerui64i_vNV (GLenum target, GLuint index, GLuint64EXT *data)
func GetIntegerui64i_vNV(target GLenum, index GLuint, data *GLuint64EXT) 
{
	C.glGetIntegerui64i_vNV(C.GLenum(target), C.GLuint(index), (*C.GLuint64EXT)(data));
}

//extern void glGetBufferParameterui64vNV (GLenum target, GLenum pname, GLuint64EXT *params)
func GetBufferParameterui64vNV(target GLenum, pname GLenum, params *GLuint64EXT) 
{
	C.glGetBufferParameterui64vNV(C.GLenum(target), C.GLenum(pname), (*C.GLuint64EXT)(params));
}

//extern void glGetIntegerui64vNV (GLenum target, GLuint64EXT *data)
func GetIntegerui64vNV(target GLenum, data *GLuint64EXT) 
{
	C.glGetIntegerui64vNV(C.GLenum(target), (*C.GLuint64EXT)(data));
}

//extern void glGetNamedBufferParameterui64vNV (GLuint buffer, GLenum pname, GLuint64EXT *params)
func GetNamedBufferParameterui64vNV(buffer GLuint, pname GLenum, params *GLuint64EXT) 
{
	C.glGetNamedBufferParameterui64vNV(C.GLuint(buffer), C.GLenum(pname), (*C.GLuint64EXT)(params));
}

//extern GLboolean glIsBufferResidentNV (GLenum target)
func IsBufferResidentNV(target GLenum) GLboolean
{
	return GLboolean(C.glIsBufferResidentNV(C.GLenum(target)));
}

//extern GLboolean glIsNamedBufferResidentNV (GLuint buffer)
func IsNamedBufferResidentNV(buffer GLuint) GLboolean
{
	return GLboolean(C.glIsNamedBufferResidentNV(C.GLuint(buffer)));
}

//extern void glMakeBufferNonResidentNV (GLenum target)
func MakeBufferNonResidentNV(target GLenum) 
{
	C.glMakeBufferNonResidentNV(C.GLenum(target));
}

//extern void glMakeBufferResidentNV (GLenum target, GLenum access)
func MakeBufferResidentNV(target GLenum, access GLenum) 
{
	C.glMakeBufferResidentNV(C.GLenum(target), C.GLenum(access));
}

//extern void glMakeNamedBufferNonResidentNV (GLuint buffer)
func MakeNamedBufferNonResidentNV(buffer GLuint) 
{
	C.glMakeNamedBufferNonResidentNV(C.GLuint(buffer));
}

//extern void glMakeNamedBufferResidentNV (GLuint buffer, GLenum access)
func MakeNamedBufferResidentNV(buffer GLuint, access GLenum) 
{
	C.glMakeNamedBufferResidentNV(C.GLuint(buffer), C.GLenum(access));
}

//extern void glUniformui64NV (GLint location, GLuint64EXT v0)
func Uniformui64NV(location GLint, v0 GLuint64EXT) 
{
	C.glUniformui64NV(C.GLint(location), C.GLuint64EXT(v0));
}

//extern void glUniformui64vNV (GLint location, GLsizei count, const GLuint64EXT *value)
func Uniformui64vNV(location GLint, count GLsizei, value *GLuint64EXT) 
{
	C.glUniformui64vNV(C.GLint(location), C.GLsizei(count), (*C.GLuint64EXT)(value));
}

//extern void glGetUniformui64vNV (GLuint program, GLint location, GLuint64EXT *params)
func GetUniformui64vNV(program GLuint, location GLint, params *GLuint64EXT) 
{
	C.glGetUniformui64vNV(C.GLuint(program), C.GLint(location), (*C.GLuint64EXT)(params));
}

//extern void glProgramUniformui64NV (GLuint program, GLint location, GLuint64EXT v0)
func ProgramUniformui64NV(program GLuint, location GLint, v0 GLuint64EXT) 
{
	C.glProgramUniformui64NV(C.GLuint(program), C.GLint(location), C.GLuint64EXT(v0));
}

//extern void glProgramUniformui64vNV (GLuint program, GLint location, GLsizei count, const GLuint64EXT *value)
func ProgramUniformui64vNV(program GLuint, location GLint, count GLsizei, value *GLuint64EXT) 
{
	C.glProgramUniformui64vNV(C.GLuint(program), C.GLint(location), C.GLsizei(count), (*C.GLuint64EXT)(value));
}

//GLenum glewInit(void)
func Init() GLenum
{
	return GLenum(C.glewInit());
}

