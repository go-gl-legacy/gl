# Copyright 2009 The Go Authors.  All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

include $(GOROOT)/src/Make.$(GOARCH)

TARG=gl

GOFILES:=gl_defs.go

CGOFILES:=gl.go

CGO_LDFLAGS:=-lGLEW

CGO_CFLAGS:=

CLEANFILES+=gl

include $(GOROOT)/src/Make.pkg

gl.go: gl.c convert.pl
	./convert.pl gl.c > gl.go

gl_defs.go: gl.c
	godefs -g gl gl.c > gl_defs.go

main: main.go install
	$(GC) main.go
	$(LD) -o $@ main.$(O)

