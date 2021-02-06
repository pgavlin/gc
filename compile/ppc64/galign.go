// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ppc64

import (
	"github.com/pgavlin/gc/compile/gc"
	"github.com/pgavlin/gc/obj/ppc64"
	"github.com/pgavlin/gc/objabi"
)

func Init(arch *gc.Arch) {
	arch.LinkArch = &ppc64.Linkppc64
	if objabi.GOARCH == "ppc64le" {
		arch.LinkArch = &ppc64.Linkppc64le
	}
	arch.REGSP = ppc64.REGSP
	arch.MAXWIDTH = 1 << 60

	arch.ZeroRange = zerorange
	arch.Ginsnop = ginsnop
	arch.Ginsnopdefer = ginsnopdefer

	arch.SSAMarkMoves = ssaMarkMoves
	arch.SSAGenValue = ssaGenValue
	arch.SSAGenBlock = ssaGenBlock
}
