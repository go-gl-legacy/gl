# Copyright 2009 The Go Authors.  All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

include $(GOROOT)/src/Make.inc

.PHONY: all gl install examples clean

all: install examples

gl:
	gomake -C gl
	gomake -C glu

install: gl
	gomake -C gl install
	gomake -C glu install

examples:
	gomake -C examples

clean:
	gomake -C gl clean
	gomake -C glu clean
	gomake -C examples clean
