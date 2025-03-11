package apperrors

import "errors"

var Platform = struct {
	InvalidPlatformID error
	PlatformNotFound  error
}{
	InvalidPlatformID: errors.New("platform_invalid_id"),
	PlatformNotFound:  errors.New("platform_not_found"),
}
