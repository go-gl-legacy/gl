# Copyright 2009 The Go Authors.  All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

include $(GOROOT)/src/Make.inc

.PHONY: all gl install examples clean

all: install

gl:
	gomake -C src/gl
	gomake -C src/glu

install: gl
	gomake -C src/gl install
	gomake -C src/glu install

examples:
	gomake -C src/examples

clean:
	gomake -C src/gl clean
	gomake -C src/glu clean
	gomake -C src/examples clean
