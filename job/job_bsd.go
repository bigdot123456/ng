// Copyright 2015 The Neugram Authors. All rights reserved.
// See the LICENSE file for rights to use this source code.

// +build darwin

package job

import "syscall"

const (
	termiosGet = syscall.TIOCGETA
	termiosSet = syscall.TIOCSETA
)
