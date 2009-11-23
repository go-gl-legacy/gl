# Copyright 2009 The Go Authors.  All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

include $(GOROOT)/src/Make.$(GOARCH)

TARG=sdl

GOFILES:=sdl_defs.go sdldraw.go

CGOFILES:=sdl.go

CGO_LDFLAGS:=`pkg-config --libs sdl` -lSDL_image

CGO_CFLAGS:=`pkg-config --cflags sdl`

CLEANFILES+=sdl

include $(GOROOT)/src/Make.pkg

sdl_defs.go: sdl.c
	godefs -g sdl sdl.c > sdl_defs.go
