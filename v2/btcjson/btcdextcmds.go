// Copyright (c) 2014 Conformal Systems LLC.
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package btcjson

// DebugLevelCmd defines the debuglevel JSON-RPC command.  This command is not a
// standard Bitcoin command.  It is an extension for btcd.
type DebugLevelCmd struct {
	LevelSpec string
}

// NewDebugLevelCmd returns a new DebugLevelCmd which can be used to issue a
// debuglevel JSON-RPC command.  This command is not a standard Bitcoin command.
// It is an extension for btcd.
func NewDebugLevelCmd(levelSpec string) *DebugLevelCmd {
	return &DebugLevelCmd{
		LevelSpec: levelSpec,
	}
}

func init() {
	MustRegisterCmd("debuglevel", (*DebugLevelCmd)(nil))
}
