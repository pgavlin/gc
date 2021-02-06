// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package wasm

import (
	"github.com/pgavlin/gc/sys"
	"github.com/pgavlin/gc/link/ld"
)

func Init() (*sys.Arch, ld.Arch) {
	theArch := ld.Arch{
		Funcalign: 16,
		Maxalign:  32,
		Minalign:  1,

		Archinit:      archinit,
		AssignAddress: assignAddress,
		Asmb:          asmb,
		Asmb2:         asmb2,
		Gentext2:      gentext2,
	}

	return sys.ArchWasm, theArch
}

func archinit(ctxt *ld.Link) {
	if *ld.FlagRound == -1 {
		*ld.FlagRound = 4096
	}
	if *ld.FlagTextAddr == -1 {
		*ld.FlagTextAddr = 0
	}
}
