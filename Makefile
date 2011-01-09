# Copyright 2009 The Go Authors.  All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

include $(GOROOT)/src/Make.inc

.PHONY: gl10 install gl20 install-gl20 all install-all examples clean

gl10:
	gomake -C gl10
	gomake -C glu-gl10

install: gl10
	gomake -C gl10 install
	gomake -C glu-gl10 install

gl20:
	gomake -C gl20
	gomake -C glu

install-gl20: gl20
	gomake -C gl20 install
	gomake -C glu install

all: gl10 gl20

install-all: install install-gl20

examples:
	gomake -C examples

clean:
	gomake -C gl10 clean
	gomake -C glu-gl10 clean
	gomake -C gl20 clean
	gomake -C glu clean
	gomake -C examples clean
