package apperrors

import "errors"

var Common = struct {
	Success                  error
	BadRequest               error
	NotFound                 error
	Unauthorized             error
	Forbidden                error
	AlreadyExisted           error
	EmailAlreadyExisted      error
	InvalidID                error
	InvalidName              error
	InvalidEmail             error
	InvalidCode              error
	InvalidStatus            error
	InvalidRole              error
	InvalidProvince          error
	InvalidDistrict          error
	InvalidWard              error
	InvalidPosition          error
	InvalidIdentifier        error
	InvalidBank              error
	InvalidBankAccountName   error
	InvalidBankAccountNumber error
	InvalidArticle           error
}{
	Success:                  errors.New("success"),
	BadRequest:               errors.New("bad_request"),
	NotFound:                 errors.New("not_found"),
	Unauthorized:             errors.New("unauthorized"),
	Forbidden:                errors.New("forbidden"),
	AlreadyExisted:           errors.New("already_existed"),
	EmailAlreadyExisted:      errors.New("email_already_existed"),
	InvalidID:                errors.New("invalid_id"),
	InvalidName:              errors.New("invalid_name"),
	InvalidEmail:             errors.New("invalid_email"),
	InvalidCode:              errors.New("invalid_code"),
	InvalidStatus:            errors.New("invalid_status"),
	InvalidRole:              errors.New("invalid_role"),
	InvalidProvince:          errors.New("invalid_province"),
	InvalidDistrict:          errors.New("invalid_district"),
	InvalidWard:              errors.New("invalid_ward"),
	InvalidPosition:          errors.New("invalid_position"),
	InvalidIdentifier:        errors.New("invalid_identifier"),
	InvalidBank:              errors.New("invalid_bank"),
	InvalidBankAccountName:   errors.New("invalid_bank_account_name"),
	InvalidBankAccountNumber: errors.New("invalid_bank_account_number"),
	InvalidArticle:           errors.New("invalid_article"),
}
