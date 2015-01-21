// Copyright (c) 2015 Conformal Systems LLC.
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package btcjson

import (
	"fmt"
	"reflect"
)

// CmdMethod returns the method for the passed command.  The provided command
// type must be a registered type.  All commands provided by this package are
// registered by default.
func CmdMethod(cmd interface{}) (string, error) {
	// Look up the cmd type and error out if not registered.
	rt := reflect.TypeOf(cmd)
	registerLock.RLock()
	method, ok := concreteTypeToMethod[rt]
	registerLock.RUnlock()
	if !ok {
		str := fmt.Sprintf("%q is not registered", method)
		return "", makeError(ErrUnregisteredMethod, str)
	}

	return method, nil
}

// MethodUsageFlags returns the usage flags for the passed command method.  The
// provided method must be associated with a registered type.  All commands
// provided by this package are registered by default.
func MethodUsageFlags(method string) (UsageFlag, error) {
	// Look up details about the provided method and error out if not
	// registered.
	registerLock.RLock()
	info, ok := methodToInfo[method]
	registerLock.RUnlock()
	if !ok {
		str := fmt.Sprintf("%q is not registered", method)
		return 0, makeError(ErrUnregisteredMethod, str)
	}

	return info.flags, nil
}
