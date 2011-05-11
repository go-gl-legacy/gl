# Copyright 2009 The Go Authors.  All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

include $(GOROOT)/src/Make.inc

TARG=gl

GOFILES:=gl_defs.go

CGOFILES:=gl.go

ifeq ($(GOOS),darwin)
CGO_LDFLAGS:=-framework OpenGL -lGLEW -lGL
else
CGO_LDFLAGS:=-lGLEW -lGL
endif

CLEANFILES+=gl

include $(GOROOT)/src/Make.pkg

DOLLAR:="$"

gl_defs.go: gl_defs.c
	godefs `echo -n $(CGO_CFLAGS) | sed 's/\([^ ^$(DOLLAR)]*\)/-f \1/g'` -g gl gl_defs.c > gl_defs.go
	gofmt -w gl_defs.go
