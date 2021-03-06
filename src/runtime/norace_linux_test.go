// Copyright 2015 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The file contains tests that can not run under race detector for some reason.
// +build !race

package runtime_test

import (
	"runtime"
	"testing"
	"time"
	"unsafe"
)

var newOSProcDone bool

//go:nosplit
func newOSProcCreated() {
	newOSProcDone = true
}

func TestNewOSProc0(t *testing.T) {
	runtime.NewOSProc0(0x800000, unsafe.Pointer(runtime.FuncPC(newOSProcCreated)))
	check, end := time.Tick(1*time.Second), time.Tick(5*time.Second)
	for {
		select {
		case <-check:
			if newOSProcDone {
				return
			}
		case <-end:
			t.Fatalf("couldn't create new OS process")
		}
	}
}
