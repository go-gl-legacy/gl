# Copyright 2009 The Go Authors.  All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

include $(GOROOT)/src/Make.inc

.PHONY: all install examples clean

all:
	gomake -C gl10
	gomake -C gl20
	gomake -C glu

install: all
	gomake -C gl10 install
	gomake -C gl20 install
	gomake -C glu install

examples:
	gomake -C examples

clean:
	gomake -C gl10 clean
	gomake -C gl20 clean
	gomake -C glu clean
	gomake -C examples clean
