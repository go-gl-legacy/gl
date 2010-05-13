# Copyright 2009 The Go Authors.  All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

include $(GOROOT)/src/Make.$(GOARCH)

all: test-gl libs

libs:
	make -C gl10 install

test-sdl: test-sdl.go libs
	$(GC) test-sdl.go
	$(LD) -o $@ test-sdl.$(O)

test-gl: test-gl.go libs
	$(GC) test-gl.go
	$(LD) -o $@ test-gl.$(O)

clean:
	make -C gl10 clean
	make -C gl20 clean
	rm -f -r *.8 *.6 *.o */*.8 */*.6 */*.o */_obj test-gl
