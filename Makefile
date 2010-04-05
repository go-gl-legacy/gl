# Copyright 2009 The Go Authors.  All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

include $(GOROOT)/src/Make.$(GOARCH)

all: test-gl libs

libs:
	make -C gl install

test-gl: test-gl.go suzanne.go libs
	$(GC) test-gl.go suzanne.go
	$(LD) -o $@ test-gl.$(O)

clean:
	make -C gl clean
	rm -f -r *.8 *.6 *.o */*.8 */*.6 */*.o */_obj test-gl
