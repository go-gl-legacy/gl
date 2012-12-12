#include <stdlib.h>
#ifdef __APPLE__
#include <OpenGL/gl3.h>
#include <OpenGL/gl3ext.h>
#else
#include <GL/glew.h>
#undef GLEW_GET_FUN
#define GLEW_GET_FUN(x) (*x)
#endif

GLenum goglInit() {
#ifdef __APPLE__
	return (GLenum)0;
#else
	return glewInit();
#endif
}

