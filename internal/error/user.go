package apperrors

import "errors"

var User = struct {
	InvalidUserID   error
	UserNotFound    error
	InvalidSourceID error

	InvalidCardIssuer  error
	InvalidCard        error
	CardAlreadyExisted error
	CardNotFound       error
	InvalidBankAccount error
}{
	InvalidUserID:   errors.New("user_invalid_id"),
	UserNotFound:    errors.New("user_not_found"),
	InvalidSourceID: errors.New("user_invalid_source_id"),

	InvalidCardIssuer:  errors.New("user_invalid_card_issuer"),
	InvalidCard:        errors.New("user_invalid_card"),
	CardAlreadyExisted: errors.New("user_card_already_existed"),
	CardNotFound:       errors.New("user_card_not_found"),
	InvalidBankAccount: errors.New("user_invalid_bank_account"),
}
