# Copyright 2009 The Go Authors.  All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

include $(GOROOT)/src/Make.inc

ver=10

all:
	make -C "gl$(ver)"
	make -C glu

install: all
	make -C "gl$(ver)" install
	make -C glu install

examples:
	make -C examples

clean:
	make -C gl10 clean
	make -C gl20 clean
	make -C glu clean
	make -C examples clean
