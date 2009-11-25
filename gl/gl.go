package gl

// #include <GL/gl.h>
// GLenum glewInit(void);
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
func GetString(name GLenum) *GLubyte
{
	return (*GLubyte)(C.glGetString(C.GLenum(name)));
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
//GLenum glewInit(void)
func Init() GLenum
{
	return GLenum(C.glewInit());
}

