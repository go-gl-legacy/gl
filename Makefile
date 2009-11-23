# Copyright 2009 The Go Authors.  All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

include $(GOROOT)/src/Make.$(GOARCH)

all: test-sdl test-gl

sdl:
	make -C sdl install

gl:
	make -C gl install

ttf:
	make -C ttf install

test-sdl: test-sdl.go sdl ttf
	$(GC) test-sdl.go
	$(LD) -o $@ test-sdl.$(O)

test-gl: test-gl.go sdl gl
	$(GC) test-gl.go
	$(LD) -o $@ test-gl.$(O)

clean:
	make -C sdl clean
	make -C gl clean
	make -C ttf clean
	make -C 4s clean
	rm -f -r *.8 *.6 *.o */*.8 */*.6 */*.o */_obj test-sdl test-gl shoot.png
