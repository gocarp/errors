// Copyright (c) 2022-2024 The Focela Authors, All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package errors

import "github.com/gocarp/codes"

// Code returns the error code.
// It returns CodeNil if it has no error code.
func (err *Error) Code() codes.Code {
	if err == nil {
		return codes.CodeNil
	}
	if err.code == codes.CodeNil {
		return Code(err.Unwrap())
	}
	return err.code
}

// SetCode updates the internal code with given code.
func (err *Error) SetCode(code codes.Code) {
	if err == nil {
		return
	}
	err.code = code
}
