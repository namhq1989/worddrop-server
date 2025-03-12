package queue

var TypeNames = struct {
	UpdateMerchantUserStatistic             string
	UpdateStoreUserStatistic                string
	DeleteUserActiveShoppingSessionsCaching string

	UpdateUserCardsStat           string
	CardEnrolledSuccessfully      string
	UnsetDefaultOtherBankAccounts string

	VisaDeleteCard string
}{
	UpdateMerchantUserStatistic:             "merchant.updateMerchantUserStatistic",
	UpdateStoreUserStatistic:                "merchant.updateStoreUserStatistic",
	DeleteUserActiveShoppingSessionsCaching: "merchant.deleteUserActiveShoppingSessionsCaching",

	UpdateUserCardsStat:           "user.updateUserCardsStat",
	CardEnrolledSuccessfully:      "user.cardEnrolledSuccessfully",
	UnsetDefaultOtherBankAccounts: "user.unsetDefaultOtherBankAccounts",

	VisaDeleteCard: "visa.deleteCard",
}
