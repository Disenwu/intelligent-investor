package errors

import errorsx "intelligent-investor/pkg/error"

var (
	ErrParamInvalid = &errorsx.ErrorsX{
		Code:     errorsx.BadRequest,
		Reason:   "ParamInvalid",
		Message:  "Invalid Parameter",
		Metadata: nil,
	}
	ErrPageNotFound = &errorsx.ErrorsX{
		Code:     errorsx.NotFound,
		Reason:   "NotFound.PageNotFound",
		Message:  "Page Not Found",
		Metadata: nil,
	}
	ErrAuthorizationFailed = &errorsx.ErrorsX{
		Code:     errorsx.AuthorizationFailed,
		Reason:   "AuthorizationFailed",
		Message:  "Authorization Failed",
		Metadata: nil,
	}
	ErrForbidden = &errorsx.ErrorsX{
		Code:     errorsx.Forbidden,
		Reason:   "Forbidden",
		Message:  "Forbidden",
		Metadata: nil,
	}
	ErrInternalServerError = &errorsx.ErrorsX{
		Code:     errorsx.InternalServerError,
		Reason:   "InternalServerError",
		Message:  "Internal Server Error",
		Metadata: nil,
	}
)
