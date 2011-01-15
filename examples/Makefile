# Copyright 2009 The Go Authors.  All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

include $(GOROOT)/src/Make.inc

all: draw gomandel gogears

draw: draw.go
	$(GC) draw.go
	$(LD) -o $@ draw.$(O)

gomandel: gomandel.go
	$(GC) gomandel.go
	$(LD) -o $@ gomandel.$(O)

gogears: gogears.go
	$(GC) gogears.go
	$(LD) -o $@ gogears.$(O)

clean:
	rm -f -r *.8 *.6 *.o */*.8 */*.6 */*.o */_obj test-gl
