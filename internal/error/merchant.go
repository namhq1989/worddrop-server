package apperrors

import "errors"

var Merchant = struct {
	InvalidMerchantID error
	MerchantNotFound  error
	SettingNotFound   error

	InvalidStoreID error
	StoreNotFound  error

	InvalidCollectionID error
	CollectionNotFound  error

	InvalidTransactionID error
	TransactionNotFound  error
}{
	InvalidMerchantID: errors.New("merchant_invalid_id"),
	MerchantNotFound:  errors.New("merchant_not_found"),
	SettingNotFound:   errors.New("merchant_setting_not_found"),

	InvalidStoreID: errors.New("merchant_store_invalid_id"),
	StoreNotFound:  errors.New("merchant_store_not_found"),

	InvalidCollectionID: errors.New("merchant_collection_invalid_id"),
	CollectionNotFound:  errors.New("merchant_collection_not_found"),

	InvalidTransactionID: errors.New("merchant_transaction_invalid_id"),
	TransactionNotFound:  errors.New("merchant_transaction_not_found"),
}
