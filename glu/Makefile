# Copyright 2009 The Go Authors.  All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

include $(GOROOT)/src/Make.inc

TARG=gl/glu

CGOFILES:=glu.go

ifeq ($(GOOS),darwin)
CGO_LDFLAGS:=-framework OpenGL
else ifeq ($(GOOS),windows)
CGO_LDFLAGS:=-lglu32
else
CGO_LDFLAGS:=-lGLU
endif

GC="${O}g" -I ../gl/_obj

include $(GOROOT)/src/Make.pkg
