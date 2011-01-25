# Copyright 2009 The Go Authors.  All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

include $(GOROOT)/src/Make.inc

TARG=gl/glu

CGOFILES:=glu.go

PLATFORM:=$(shell uname -s)

ifeq ($(PLATFORM),Darwin)
CGO_LDFLAGS:=-framework OpenGL
else
CGO_LDFLAGS:=-lGLU
endif
CGO_CFLAGS:=-D__$(PLATFORM)

GC="${O}g" -I ../gl20/_obj

include $(GOROOT)/src/Make.pkg
