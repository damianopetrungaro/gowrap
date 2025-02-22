package templatestests

// Code generated by gowrap. DO NOT EDIT.
// template: ../templates/grpc_validate
// gowrap: http://github.com/hexdigest/gowrap

//go:generate gowrap gen -p github.com/hexdigest/gowrap/templates_tests -i InterfaceWithValidtableArg -t ../templates/grpc_validate -o interface_with_grpc_validation.go -l ""

import (
	"context"

	grpc_codes "google.golang.org/grpc/codes"
	grpc_status "google.golang.org/grpc/status"
)

// InterfaceWithValidtableArgWithGRPCValidation implements InterfaceWithValidtableArg interface instrumented with GRPC request validation
type InterfaceWithValidtableArgWithGRPCValidation struct {
	InterfaceWithValidtableArg
}

// NewInterfaceWithValidtableArgWithGRPCValidation returns InterfaceWithValidtableArgWithGRPCValidation
func NewInterfaceWithValidtableArgWithGRPCValidation(base InterfaceWithValidtableArg) InterfaceWithValidtableArgWithGRPCValidation {
	return InterfaceWithValidtableArgWithGRPCValidation{
		InterfaceWithValidtableArg: base,
	}
}

// Method implements InterfaceWithValidtableArg
func (_d InterfaceWithValidtableArgWithGRPCValidation) Method(ctx context.Context, r *ValidatableRequest) (err error) {
	if _v, _ok := interface{}(r).(interface{ Validate() error }); _ok {
		if err = _v.Validate(); err != nil {
			err = grpc_status.Error(grpc_codes.InvalidArgument, err.Error())
			return
		}
	}

	return _d.InterfaceWithValidtableArg.Method(ctx, r)
}
