// Copyright 2012 The go-gl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gl

import (
	"testing"
	"unsafe"
)

func TestPtr(t *testing.T) {
	// test nil
	if p, q := unsafe.Pointer(nil), ptr(nil); p != q {
		t.Fatalf("expected %#v, got %#v\n", p, q)
	}

	// test nil interface
	var r interface{}
	if p, q := unsafe.Pointer(nil), ptr(r); p != q {
		t.Fatalf("expected %#v, got %#v\n", p, q)
	}

	// test nil pointer
	var s *int
	if p, q := unsafe.Pointer(nil), ptr(s); p != q {
		t.Fatalf("expected %#v, got %#v\n", p, q)
	}

	// test uinptr
	for _, n := range []uintptr{0, 1, 2, 42} {
		if p, q := unsafe.Pointer(n), ptr(n); p != q {
			t.Fatalf("expected %#v, got %#v\n", p, q)
		}
	}
}
