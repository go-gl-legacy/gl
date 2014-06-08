// Copyright 2012 The go-gl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gl

// #include "gl.h"
//
// void goDebugCB(GLenum source, GLenum type, GLuint id, GLenum severity, GLsizei length, char *message);
//
// static inline void debugProcCB(GLenum source, GLenum type, GLuint id,
//     GLenum severity, GLsizei length, const GLchar* message, void* userParam)
// {
//     goDebugCB(source, type, id, severity, length, (char *)message);
// }
//
// static inline void glDebugMessageCB(void)
// {
//     glDebugMessageCallbackARB(debugProcCB, NULL);
// }
//
// /*
//  * Depending on glew version 'message' could be 'GLchar*' or 'char*' which
//  * causes problems in more strict Go type system. Let's work it around in C
//  */
// static inline void __glDebugMessageInsert(GLenum source, GLenum type, GLuint id,
//     GLenum severity, GLsizei length, const char *message)
// {
//     glDebugMessageInsertARB(source, type, id, severity, length, message);
// }
//
// static inline void GetDebugMessageLog(GLuint count, GLsizei bufSize, GLenum *sources,
//     GLenum *types, GLuint *ids, GLenum *severities, GLsizei *lengths, char *messageLog)
// {
//     glGetDebugMessageLogARB(count, bufSize, sources, types, ids, severities, lengths, messageLog);
// }
import "C"
import "unsafe"

const (
	DEBUG_CALLBACK_FUNCTION          = C.GL_DEBUG_CALLBACK_FUNCTION_ARB
	DEBUG_CALLBACK_USER_PARAM        = C.GL_DEBUG_CALLBACK_USER_PARAM_ARB
	DEBUG_LOGGED_MESSAGES            = C.GL_DEBUG_LOGGED_MESSAGES_ARB
	DEBUG_NEXT_LOGGED_MESSAGE_LENGTH = C.GL_DEBUG_NEXT_LOGGED_MESSAGE_LENGTH_ARB
	DEBUG_OUTPUT_SYNCHRONOUS         = C.GL_DEBUG_OUTPUT_SYNCHRONOUS_ARB
	DEBUG_SEVERITY_HIGH              = C.GL_DEBUG_SEVERITY_HIGH_ARB
	DEBUG_SEVERITY_LOW               = C.GL_DEBUG_SEVERITY_LOW_ARB
	DEBUG_SEVERITY_MEDIUM            = C.GL_DEBUG_SEVERITY_MEDIUM_ARB
	DEBUG_SOURCE_API                 = C.GL_DEBUG_SOURCE_API_ARB
	DEBUG_SOURCE_APPLICATION         = C.GL_DEBUG_SOURCE_APPLICATION_ARB
	DEBUG_SOURCE_OTHER               = C.GL_DEBUG_SOURCE_OTHER_ARB
	DEBUG_SOURCE_SHADER_COMPILER     = C.GL_DEBUG_SOURCE_SHADER_COMPILER_ARB
	DEBUG_SOURCE_THIRD_PARTY         = C.GL_DEBUG_SOURCE_THIRD_PARTY_ARB
	DEBUG_SOURCE_WINDOW_SYSTEM       = C.GL_DEBUG_SOURCE_WINDOW_SYSTEM_ARB
	DEBUG_TYPE_DEPRECATED_BEHAVIOR   = C.GL_DEBUG_TYPE_DEPRECATED_BEHAVIOR_ARB
	DEBUG_TYPE_ERROR                 = C.GL_DEBUG_TYPE_ERROR_ARB
	DEBUG_TYPE_OTHER                 = C.GL_DEBUG_TYPE_OTHER_ARB
	DEBUG_TYPE_PERFORMANCE           = C.GL_DEBUG_TYPE_PERFORMANCE_ARB
	DEBUG_TYPE_PORTABILITY           = C.GL_DEBUG_TYPE_PORTABILITY_ARB
	DEBUG_TYPE_UNDEFINED_BEHAVIOR    = C.GL_DEBUG_TYPE_UNDEFINED_BEHAVIOR_ARB
	MAX_DEBUG_LOGGED_MESSAGES        = C.GL_MAX_DEBUG_LOGGED_MESSAGES_ARB
	MAX_DEBUG_MESSAGE_LENGTH         = C.GL_MAX_DEBUG_MESSAGE_LENGTH_ARB
)

var ARB_debug_output = false

func init() {
	extensions["GL_ARB_debug_output"] = &ARB_debug_output
}

type debugProc func(source GLenum, typ GLenum, id uint, severity GLenum, message string)

var debugCB debugProc

//export goDebugCB
func goDebugCB(source GLenum, typ GLenum, id GLuint, severity GLenum, length GLsizei, message *C.char) {
	debugCB(source, typ, uint(id), severity, C.GoStringN(message, C.int(length)))
}

func DebugMessageCallback(cbfunc debugProc) {
	if cbfunc == nil {
		C.glDebugMessageCallbackARB(nil, nil)
	} else {
		debugCB = cbfunc
		C.glDebugMessageCB()
	}
}

func DebugMessageControl(source GLenum, typ GLenum, severity GLenum, ids []uint, enabled bool) {
	C.glDebugMessageControlARB(C.GLenum(source), C.GLenum(typ), C.GLenum(severity),
		C.GLsizei(len(ids)), (*C.GLuint)(unsafe.Pointer(&ids)), glBool(enabled))
}

func DebugMessageInsert(source GLenum, typ GLenum, id uint, severity GLenum, message string) {
	C.__glDebugMessageInsert(C.GLenum(source), C.GLenum(typ), C.GLuint(id), C.GLenum(severity),
		C.GLsizei(len(message)), C.CString(message))
}

func GetNextDebugMessage() (msg string, source GLenum, typ GLenum, id uint, severity GLenum) {
	length := []int32{0}
	GetIntegerv(DEBUG_NEXT_LOGGED_MESSAGE_LENGTH, length)
	if length[0] < 1 {
		msg = ""
		return
	}

	buf := C.malloc(C.size_t(length[0]))
	defer C.free(buf)

	_id := []GLuint{0}
	C.GetDebugMessageLog(C.GLuint(1), C.GLsizei(length[0]),
		(*C.GLenum)(&source), (*C.GLenum)(&typ), (*C.GLuint)(&_id[0]),
		(*C.GLenum)(&severity), nil, (*C.char)(buf))

	id = uint(_id[0])
	msg = C.GoString((*C.char)(buf))
	return
}
