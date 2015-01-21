// Copyright (c) 2015 Conformal Systems LLC.
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

// NOTE: This file is intended to house the RPC commands that are supported by
// a wallet server with btcwallet extensions.

package btcjson

// CreateNewAccountCmd defines the createnewaccount JSON-RPC command.
type CreateNewAccountCmd struct {
	Account string
}

// NewCreateNewAccountCmd returns a new instance which can be used to issue a
// createnewaccount JSON-RPC command.
func NewCreateNewAccountCmd(account string) *CreateNewAccountCmd {
	return &CreateNewAccountCmd{
		Account: account,
	}
}

// DumpWalletCmd defines the dumpwallet JSON-RPC command.
type DumpWalletCmd struct {
	Filename string
}

// NewDumpWalletCmd returns a new instance which can be used to issue a
// dumpwallet JSON-RPC command.
func NewDumpWalletCmd(filename string) *DumpWalletCmd {
	return &DumpWalletCmd{
		Filename: filename,
	}
}

// ImportWalletCmd defines the importwallet JSON-RPC command.
type ImportWalletCmd struct {
	Filename string
}

// NewImportWalletCmd returns a new instance which can be used to issue a
// importwallet JSON-RPC command.
func NewImportWalletCmd(filename string) *ImportWalletCmd {
	return &ImportWalletCmd{
		Filename: filename,
	}
}

// RenameAccountCmd defines the renameaccount JSON-RPC command.
type RenameAccountCmd struct {
	OldAccount string
	NewAccount string
}

// NewRenameAccountCmd returns a new instance which can be used to issue a
// renameaccount JSON-RPC command.
func NewRenameAccountCmd(oldAccount, newAccount string) *RenameAccountCmd {
	return &RenameAccountCmd{
		OldAccount: oldAccount,
		NewAccount: newAccount,
	}
}

func init() {
	// No special flags for commands in this file.
	flags := UsageFlag(0)

	MustRegisterCmd("createnewaccount", (*CreateNewAccountCmd)(nil), flags)
	MustRegisterCmd("dumpwallet", (*DumpWalletCmd)(nil), flags)
	MustRegisterCmd("importwallet", (*ImportWalletCmd)(nil), flags)
	MustRegisterCmd("renameaccount", (*RenameAccountCmd)(nil), flags)
}
