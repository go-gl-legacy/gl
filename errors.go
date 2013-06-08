// Copyright 2012 The go-gl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gl

import (
  "errors"
)

var (
  ErrorInputSize = errors.New("Invalid input size")
  ErrorPointerType = errors.New("type must be a pointer, a slice, uintptr or nil")
  ErrorEqualSliceLength = errors.New("Residences slice must be equal in length to textures slice")
)
