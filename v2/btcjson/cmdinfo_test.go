// Copyright (c) 2015 Conformal Systems LLC.
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package btcjson_test

import (
	"reflect"
	"testing"

	"github.com/btcsuite/btcjson/v2/btcjson"
)

// TestCmdMethod tests the CmdMethod function to ensure it retuns the expected
// methods and errors.
func TestCmdMethod(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		cmd    interface{}
		method string
		err    error
	}{
		{
			name: "unregistered type",
			cmd:  (*int)(nil),
			err:  btcjson.Error{ErrorCode: btcjson.ErrUnregisteredMethod},
		},
		{
			name:   "nil pointer of registered type",
			cmd:    (*btcjson.GetBlockCmd)(nil),
			method: "getblock",
		},
		{
			name:   "nil instance of registered type",
			cmd:    &btcjson.GetBlockCountCmd{},
			method: "getblockcount",
		},
	}

	t.Logf("Running %d tests", len(tests))
	for i, test := range tests {
		method, err := btcjson.CmdMethod(test.cmd)
		if reflect.TypeOf(err) != reflect.TypeOf(test.err) {
			t.Errorf("Test #%d (%s) wrong error - got %T (%[3]v), "+
				"want %T", i, test.name, err, test.err)
			continue
		}
		if err != nil {
			gotErrorCode := err.(btcjson.Error).ErrorCode
			if gotErrorCode != test.err.(btcjson.Error).ErrorCode {
				t.Errorf("Test #%d (%s) mismatched error code "+
					"- got %v (%v), want %v", i, test.name,
					gotErrorCode, err,
					test.err.(btcjson.Error).ErrorCode)
				continue
			}

			continue
		}

		// Ensure method matches the expected value.
		if method != test.method {
			t.Errorf("Test #%d (%s) mismatched method - got %v, "+
				"want %v", i, test.name, method, test.method)
			continue
		}
	}
}

// TestMethodUsage tests the MethodUsage function ensure it returns the expected
// flags and errors.
func TestMethodUsage(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		method string
		err    error
		flags  btcjson.UsageFlag
	}{
		{
			name:   "unregistered type",
			method: "bogusmethod",
			err:    btcjson.Error{ErrorCode: btcjson.ErrUnregisteredMethod},
		},
		{
			name:   "getblock",
			method: "getblock",
			flags:  0,
		},
		{
			name:   "walletpassphrase",
			method: "walletpassphrase",
			flags:  btcjson.UFWalletOnly,
		},
	}

	t.Logf("Running %d tests", len(tests))
	for i, test := range tests {
		flags, err := btcjson.MethodUsage(test.method)
		if reflect.TypeOf(err) != reflect.TypeOf(test.err) {
			t.Errorf("Test #%d (%s) wrong error - got %T (%[3]v), "+
				"want %T", i, test.name, err, test.err)
			continue
		}
		if err != nil {
			gotErrorCode := err.(btcjson.Error).ErrorCode
			if gotErrorCode != test.err.(btcjson.Error).ErrorCode {
				t.Errorf("Test #%d (%s) mismatched error code "+
					"- got %v (%v), want %v", i, test.name,
					gotErrorCode, err,
					test.err.(btcjson.Error).ErrorCode)
				continue
			}

			continue
		}

		// Ensure flags match the expected value.
		if flags != test.flags {
			t.Errorf("Test #%d (%s) mismatched flags - got %v, "+
				"want %v", i, test.name, flags, test.flags)
			continue
		}
	}
}
