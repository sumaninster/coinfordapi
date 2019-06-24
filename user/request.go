package user

type UserAdd struct {
	Name 			string
	Username		string
	Password		string
	Email 			string
	Token			string
}

type UserLogin struct {
	Username 		string
	Password		string
	Token			string
}

type UserToken struct {
	Token			string
}

type UserUsername struct {
	Username		string
	Token			string
}

type UserChangePassword struct {
	CurrentPassword	string
	NewPassword	 	string
	Token			string
}

type UserChangeUsername struct {
	CurrentPassword	string
	NewUsername	 	string
	Token			string
}

type UserChangeName struct {
	CurrentPassword	string
	NewName		 	string
	Token			string
}

type UserChangeEmail struct {
	Id				int64
	NewEmail	 	string
	Token			string
}

type UserCountryIdAdd struct {
	CountryId			int64
	Token				string
	RequestCountries 	RequestCountries
}

type UserCountryId struct {
	Id					int64
	Token				string
	RequestCountries 	RequestCountries
}

type WalletDetails struct {
	CurrencyId		int64
	Nickname 		string
	Primary 		bool
	Token			string
}

type WalletMasterId struct {
	Id	 			int64
	Token			string
}

type PayeeDetails struct {
	CurrencyId		int64
	Name 			string
	Email 			string
	Nickname 		string
	Address 		string
	Token			string
}

type PayeeMasterId struct {
	Id	 			int64
	Token			string
}

type RequestOrderList struct {
	CurrencyId			int64
	RateCurrencyId 	int64
	Token						string
}

type OrderWalletUpdate struct {
	CurrencyId						int64
	RateCurrencyId 				int64
	CurrencyWalletId 			int64
	RateCurrencyWalletId 	int64
	Token									string
}

type OrderBuy struct {
	CurrencyId			int64
	Amount 					float64
	Rate						float64
	RateCurrencyId 	int64
	CurrencyWalletId 			int64
	RateCurrencyWalletId 	int64
	Token						string
}

type OrderSell struct {
	CurrencyId						int64
	Amount 								float64
	Rate									float64
	RateCurrencyId 				int64
	CurrencyWalletId 			int64
	RateCurrencyWalletId 	int64
	Token									string
}

type OrderId struct {
	OrderId 	 				int64
	CurrencyId				int64
	RateCurrencyId 		int64
	Token							string
}

type TradeHistory struct {
	Duration 	 				string
	CurrencyId				int64
	RateCurrencyId 		int64
	Token							string
}

type TransferSearch struct {
	CurrencyId 			int64
	CurrencyType 		string
	TransferType 		string
}

type TransferAdd struct {
	FromWalletId 		int64
	ToWalletId 			int64
	ToPayeeId 			int64
	ToDataId			int64
	Amount 				float64
	TransferSearch 		TransferSearch
	Token 				string
}

type TransferId struct {
	Id 					int64
	TransferSearch 		TransferSearch
	Token 				string
}

type TransferSearchAll struct {
	TransferSearch 		TransferSearch
	Token 				string
}

type CurrencyDetails struct {
	CurrencyType 	string
	Token			string
}

type CurrencyId struct {
	CurrencyId 	 	int64
	Token			string
}

type RequestCountries struct {
	OnlyMine			string
	Eligible 			string
}

type RequestCountryList struct {
	Token				string
	RequestCountries 	RequestCountries
}

type CountryDetails struct {
	Name				string
	IsoCode 			string
	DialCode			string
	Code 				string
	Token				string
	RequestCountries 	RequestCountries
}

type CountryId struct {
	Id 	 				int64
	Token				string
	RequestCountries 	RequestCountries
}

type FieldType struct {
	Token				string
	RequestField 		RequestField
}

type RequestField struct {
	FieldType 	 		string
	CountryId 			int64
}

type DataText struct {
	FieldId 	 		int64
	InputText 			string
}

type DataCategory struct {
	FieldId 	 		int64
	FieldCategoryId 	int64
}

type DataFile struct {
	FieldId 	 		int64
	Name 				string
	Extension 			string
}

type DataUpload struct {
	DataText 	 		[]DataText
	DataCategory 		[]DataCategory
	DataFile 			[]DataFile
	CountryId 			int64
	FieldType 			string
	Nickname 			string
	Primary 			bool
	Active 				bool
	Token				string
}

type DataId struct {
	Id 	 				int64
	Token				string
}
