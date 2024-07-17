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

package errdefs

import (
	"fmt"
	"strings"

	"github.com/containerd/errdefs/internal/types"
)

type joinError struct {
	errs []error
}

// Join will join the errors together and ensure stack traces
// are appropriately formatted.
func Join(errs ...error) error {
	var e error
	n := 0
	for _, err := range errs {
		if err != nil {
			e = err
			n++
		}
	}

	switch n {
	case 0:
		return nil
	case 1:
		switch e.(type) {
		case *errorValue, *joinError:
			// Don't wrap the types defined by this package
			// as that could interfere with the formatting.
			return e
		}
		return &errorValue{e}
	}

	joined := make([]error, 0, n)
	for _, err := range errs {
		if err != nil {
			joined = append(joined, err)
		}
	}
	return &joinError{errs: joined}
}

func (e *joinError) Error() string {
	var b strings.Builder
	fmt.Fprintf(&b, "%v", e)
	return b.String()
}

func (e *joinError) Format(st fmt.State, verb rune) {
	format := fmt.FormatString(st, verb)
	collapsed := verb == 'v' && st.Flag('+')
	first := true
	for _, err := range e.errs {
		if !collapsed {
			if _, ok := err.(types.CollapsibleError); ok {
				continue
			}
		}
		if !first {
			fmt.Fprintln(st)
		}
		fmt.Fprintf(st, format, err)
		first = false
	}
}

func (e *joinError) Unwrap() []error {
	return e.errs
}
