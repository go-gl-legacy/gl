// Copyright 2012 The go-gl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gl

// #include "gl.h"
import "C"

const (
	ALREADY_SIGNALED           = C.GL_ALREADY_SIGNALED
	CONDITION_SATISFIED        = C.GL_CONDITION_SATISFIED
	MAX_SERVER_WAIT_TIMEOUT    = C.GL_MAX_SERVER_WAIT_TIMEOUT
	OBJECT_TYPE                = C.GL_OBJECT_TYPE
	SIGNALED                   = C.GL_SIGNALED
	SYNC_CONDITION             = C.GL_SYNC_CONDITION
	SYNC_FENCE                 = C.GL_SYNC_FENCE
	SYNC_FLAGS                 = C.GL_SYNC_FLAGS
	SYNC_FLUSH_COMMANDS_BIT    = C.GL_SYNC_FLUSH_COMMANDS_BIT
	SYNC_GPU_COMMANDS_COMPLETE = C.GL_SYNC_GPU_COMMANDS_COMPLETE
	SYNC_STATUS                = C.GL_SYNC_STATUS
	TIMEOUT_EXPIRED            = C.GL_TIMEOUT_EXPIRED
	TIMEOUT_IGNORED            = C.GL_TIMEOUT_IGNORED
	UNSIGNALED                 = C.GL_UNSIGNALED
	WAIT_FAILED                = C.GL_WAIT_FAILED
)

var ARB_sync = false

func init() {
	extensions["GL_ARB_sync"] = &ARB_sync
}

func ClientWaitSync(glSync GLsync, flags uint32, timeout uint64) {
	C.glClientWaitSync(C.GLsync(glSync), C.GLbitfield(flags), C.GLuint64(timeout))
}

func DeleteSync(glSync GLsync) {
	C.glDeleteSync(C.GLsync(glSync))
}

func FenceSync(condition GLenum, flags uint32) {
	C.glFenceSync(C.GLenum(condition), C.GLbitfield(flags))
}

func GetInteger64v(pname GLenum, params []int64) {
	C.glGetInteger64v(C.GLenum(pname), (*C.GLint64)(&params[0]))
}

func GetSynciv(glSync GLsync, pname GLenum, values []int32) []int32 {
	var length int32
	C.glGetSynciv(C.GLsync(glSync), C.GLenum(pname), C.GLsizei(len(values)), (*C.GLsizei)(&length), (*C.GLint)(&values[0]))
	return values[:length]
}

func IsSync(glSync GLsync) {
	C.glIsSync(C.GLsync(glSync))
}

func WaitSync(glSync GLsync, flags uint32, timeout uint64) {
	C.glWaitSync(C.GLsync(glSync), C.GLbitfield(flags), C.GLuint64(timeout))
}
