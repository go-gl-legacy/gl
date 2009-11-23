// Copyright 2009 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// this is SDL version of src/pkg/exp/4s/4s.go

package main

import (
	"sdl";
	"log";
	"os";
	"runtime";
)


func main() {

	runtime.LockOSThread();

	args := os.Args;
	p := pieces4;
	if len(args) > 1 && args[1] == "-5" {
		p = pieces5
	}
	dx, dy := 500, 500;
	w, err := sdl.InitContext(dx, dy);
	if err != nil {
		log.Exit(err)
	}

	Play(p, w);
}

func PlaySound(b []uint16) {
	// no audio yet
}

var whoosh = []uint16{
// Insert your favorite sound samples here.
}
