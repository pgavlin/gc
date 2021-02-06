// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"github.com/pgavlin/gc/compile/amd64"
	"github.com/pgavlin/gc/compile/arm"
	"github.com/pgavlin/gc/compile/arm64"
	"github.com/pgavlin/gc/compile/gc"
	"github.com/pgavlin/gc/compile/mips"
	"github.com/pgavlin/gc/compile/mips64"
	"github.com/pgavlin/gc/compile/ppc64"
	"github.com/pgavlin/gc/compile/riscv64"
	"github.com/pgavlin/gc/compile/s390x"
	"github.com/pgavlin/gc/compile/wasm"
	"github.com/pgavlin/gc/compile/x86"
	"github.com/pgavlin/gc/objabi"
	"log"
	"os"
)

var archInits = map[string]func(*gc.Arch){
	"386":      x86.Init,
	"amd64":    amd64.Init,
	"arm":      arm.Init,
	"arm64":    arm64.Init,
	"mips":     mips.Init,
	"mipsle":   mips.Init,
	"mips64":   mips64.Init,
	"mips64le": mips64.Init,
	"ppc64":    ppc64.Init,
	"ppc64le":  ppc64.Init,
	"riscv64":  riscv64.Init,
	"s390x":    s390x.Init,
	"wasm":     wasm.Init,
}

func main() {
	// disable timestamps for reproducible output
	log.SetFlags(0)
	log.SetPrefix("compile: ")

	archInit, ok := archInits[objabi.GOARCH]
	if !ok {
		fmt.Fprintf(os.Stderr, "compile: unknown architecture %q\n", objabi.GOARCH)
		os.Exit(2)
	}

	gc.Main(archInit)
	gc.Exit(0)
}
