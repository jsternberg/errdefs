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

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func ToGRPC(err error) error {
	if err == nil {
		return nil
	}

	if isGRPCError(err) {
		// error has already been mapped to grpc
		return err
	}

	switch {
	case IsInvalidArgument(err):
		return status.Error(codes.InvalidArgument, err.Error())
	case IsNotFound(err):
		return status.Error(codes.NotFound, err.Error())
	case IsAlreadyExists(err):
		return status.Error(codes.AlreadyExists, err.Error())
	case IsFailedPrecondition(err) || IsConflict(err) || IsNotModified(err):
		return status.Error(codes.FailedPrecondition, err.Error())
	case IsUnavailable(err):
		return status.Error(codes.Unavailable, err.Error())
	case IsNotImplemented(err):
		return status.Error(codes.Unimplemented, err.Error())
	case IsCanceled(err):
		return status.Error(codes.Canceled, err.Error())
	case IsDeadlineExceeded(err):
		return status.Error(codes.DeadlineExceeded, err.Error())
	case IsUnauthorized(err):
		return status.Error(codes.Unauthenticated, err.Error())
	case IsPermissionDenied(err):
		return status.Error(codes.PermissionDenied, err.Error())
	case IsInternal(err):
		return status.Error(codes.Internal, err.Error())
	case IsDataLoss(err):
		return status.Error(codes.DataLoss, err.Error())
	case IsAborted(err):
		return status.Error(codes.Aborted, err.Error())
	case IsOutOfRange(err):
		return status.Error(codes.OutOfRange, err.Error())
	case IsResourceExhausted(err):
		return status.Error(codes.ResourceExhausted, err.Error())
	case IsUnknown(err):
		return status.Error(codes.Unknown, err.Error())
	}

	return err
}

// ToGRPCf maps the error to grpc error codes, assembling the formatting string
// and combining it with the target error string.
//
// This is equivalent to grpc.ToGRPC(fmt.Errorf("%s: %w", fmt.Sprintf(format, args...), err))
func ToGRPCf(err error, format string, args ...interface{}) error {
	return ToGRPC(fmt.Errorf("%s: %w", fmt.Sprintf(format, args...), err))
}

func isGRPCError(err error) bool {
	_, ok := status.FromError(err)
	return ok
}
