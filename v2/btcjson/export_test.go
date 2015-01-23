// Copyright (c) 2014 Conformal Systems LLC.
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package btcjson

import "reflect"

// TstHighestUsageFlagBit makes the internal highestUsageFlagBit parameter
// available to the test package.
var TstHighestUsageFlagBit = highestUsageFlagBit

// TstNumErrorCodes makes the internal numErrorCodes parameter available to the
// test package.
var TstNumErrorCodes = numErrorCodes

// TstAssignField makes the internal assignField function available to the test
// package.
func TstAssignField(paramNum int, fieldName string, dest reflect.Value, src reflect.Value) error {
	return assignField(paramNum, fieldName, dest, src)
}

// TstFieldUsage makes the internal fieldUsage function available to the test
// package.
var TstFieldUsage = fieldUsage
