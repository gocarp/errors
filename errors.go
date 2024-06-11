// Copyright (c) 2022-2024 The Focela Authors, All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package errors provides rich functionalities to manipulate errors.
//
// For maintainers, please very note that,
// this package is quite a basic package, which SHOULD NOT import extra packages
// except standard packages and internal packages, to avoid cycle imports.
package errors

import "github.com/gocarp/codes"

// IsInterface is the interface for Is feature.
type IsInterface interface {
	Error() string
	Is(target error) bool
}

// EqualInterface is the interface for Equal feature.
type EqualInterface interface {
	Error() string
	Equal(target error) bool
}

// CodeInterface is the interface for Code feature.
type CodeInterface interface {
	Error() string
	Code() codes.Code
}

// StackInterface is the interface for Stack feature.
type StackInterface interface {
	Error() string
	Stack() string
}

// CauseInterface is the interface for Cause feature.
type CauseInterface interface {
	Error() string
	Cause() error
}

// CurrentInterface is the interface for Current feature.
type CurrentInterface interface {
	Error() string
	Current() error
}

// UnwrapInterface is the interface for Unwrap feature.
type UnwrapInterface interface {
	Error() string
	Unwrap() error
}

const (
	commaSeparatorSpace = ", " // commaSeparatorSpace is the comma separator with space.
)
