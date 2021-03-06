// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sym

import "github.com/pgavlin/gc/dwarf"

// LoaderSym holds a loader.Sym value. We can't refer to this
// type from the sym package since loader imports sym.
type LoaderSym int

// CompilationUnit is an abstraction used by DWARF to represent a chunk of
// debug-related data. We create a CompilationUnit per Object file in a
// library (so, one for all the Go code, one for each assembly file, etc.).
type CompilationUnit struct {
	Pkg            string        // The package name, eg ("fmt", or "runtime")
	Lib            *Library      // Our library
	Consts         *Symbol       // Package constants DIEs
	PCs            []dwarf.Range // PC ranges, relative to Textp[0]
	DWInfo         *dwarf.DWDie  // CU root DIE
	FuncDIEs       []*Symbol     // Function DIE subtrees
	AbsFnDIEs      []*Symbol     // Abstract function DIE subtrees
	RangeSyms      []*Symbol     // Symbols for debug_range
	Textp          []*Symbol     // Text symbols in this CU
	DWARFFileTable []string      // The file table used to generate the .debug_lines

	Consts2    LoaderSym   // Package constants DIEs (loader)
	FuncDIEs2  []LoaderSym // Function DIE subtrees (loader)
	AbsFnDIEs2 []LoaderSym // Abstract function DIE subtrees (loader)
	RangeSyms2 []LoaderSym // Symbols for debug_range (loader)
	Textp2     []LoaderSym // Text symbols in this CU (loader)
}
