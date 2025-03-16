package apperrors

import "errors"

var Common = struct {
	Success              error
	BadRequest           error
	NotFound             error
	Unauthorized         error
	Forbidden            error
	AlreadyExisted       error
	EmailAlreadyExisted  error
	InvalidID            error
	InvalidName          error
	InvalidEmail         error
	InvalidCategory      error
	InvalidWord          error
	InvalidLevel         error
	InvalidDefinition    error
	InvalidPartOfSpeech  error
	InvalidIPA           error
	InvalidExample       error
	InvalidNewsTitle     error
	InvalidNewsSummary   error
	InvalidNewsSourceURL error
}{
	Success:              errors.New("success"),
	BadRequest:           errors.New("bad_request"),
	NotFound:             errors.New("not_found"),
	Unauthorized:         errors.New("unauthorized"),
	Forbidden:            errors.New("forbidden"),
	AlreadyExisted:       errors.New("already_existed"),
	EmailAlreadyExisted:  errors.New("email_already_existed"),
	InvalidID:            errors.New("invalid_id"),
	InvalidName:          errors.New("invalid_name"),
	InvalidEmail:         errors.New("invalid_email"),
	InvalidCategory:      errors.New("invalid_category"),
	InvalidWord:          errors.New("invalid_word"),
	InvalidLevel:         errors.New("invalid_level"),
	InvalidDefinition:    errors.New("invalid_definition"),
	InvalidPartOfSpeech:  errors.New("invalid_part_of_speech"),
	InvalidIPA:           errors.New("invalid_ipa"),
	InvalidExample:       errors.New("invalid_example"),
	InvalidNewsTitle:     errors.New("invalid_news_title"),
	InvalidNewsSummary:   errors.New("invalid_news_summary"),
	InvalidNewsSourceURL: errors.New("invalid_news_source_url"),
}
