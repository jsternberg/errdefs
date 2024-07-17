/*
   Copyright The containerd Authors.

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

// Package errors is a shim package as a drop-in replacement for
// github.com/pkg/errors users to migrate to using this package.
//
// It implements most of the functionality from that package
// to ease with the transition. Applications using this package
// can use this and the standard methods simultaneously.
package errors

import (
	"errors"

	"github.com/containerd/errdefs"
	"github.com/containerd/errdefs/stack"
)

func New(message string) error {
	return stack.Errorf("%s", message)
}

func Errorf(format string, args ...any) error {
	return stack.Errorf(format, args...)
}

func Wrap(err error, message string) error {
	err = errdefs.New("%s: %w", message, err)
	return errdefs.Join(err, stack.ErrStack(1))
}

func Wrapf(err error, format string, args ...any) error {
	err = errdefs.New("%w: %w", errdefs.New(format, args...), err)
	return errdefs.Join(err, stack.ErrStack(1))
}

func Is(err, target error) bool {
	return errors.Is(err, target)
}

func As(err error, target any) bool {
	return errors.As(err, target)
}

func Unwrap(err error) error {
	return errors.Unwrap(err)
}

func WithMessage(err error, message string) error {
	return errdefs.New("%s: %w", message, err)
}

func WithMessagef(err error, format string, args ...any) error {
	return errdefs.New("%w: %w", errdefs.New(format, args...), err)
}

func WithStack(err error) error {
	return errdefs.Join(err, stack.ErrStack(1))
}
