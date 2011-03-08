# Copyright 2009 The Go Authors.  All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

include $(GOROOT)/src/Make.inc

TARGETS=draw gomandel gogears

.PHONY: all clean

all: $(TARGETS)

$(TARGETS):
	$(GC) $@.go
	$(LD) -o $@ $@.$(O)

clean:
	rm -f -r *.8 *.6 *.o */*.8 */*.6 */*.o */_obj test-gl $(TARGETS)
