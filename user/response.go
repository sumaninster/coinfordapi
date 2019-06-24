
package user

import (
	"coinford_api/models"
	)

type CommonResponse struct {
	ResponseCode 			int
	ResponseDescription 	string
}

type LoginResponse struct {
	Token 					string
	Name 					string
	ResponseCode 			int
	ResponseDescription 	string
}

type EmailList struct {
	Emails 					[]models.Email
	ResponseCode 			int
	ResponseDescription 	string
}

type WalletData struct {
	Wallet 					models.Wallet
	Crypto 					*models.WalletCrypto
	Fiat 					*models.DataText
}

type WalletMasterData struct {
	WalletMaster 			models.WalletMaster
	Currency 				models.Currency
	WalletData 				[]WalletData
}

type WalletList struct {
	WalletMasterData 		[]WalletMasterData
	ResponseCode 			int
	ResponseDescription 	string
}

type MyOrderDetail struct {
	Order 						models.Order
	OrderType 				string
}

type AllOrderList struct {
	OrderBuys 						[]models.OrderBuy
	OrderSells 						[]models.OrderSell
	Orders 								[]models.Order
	MyOrderBuys						[]models.OrderBuy
	MyOrderSells					[]models.OrderSell
	MyOrders 							[]MyOrderDetail
	CurrencyWallets				[]models.Wallet
	RateCurrencyWallets 	[]models.Wallet
	OrderWallet 					*models.OrderWallet
	ResponseCode 					int
	ResponseDescription 	string
}

type AllTradeHistory struct {
	Low										float64
	High 									float64
	Last 									float64
	Orders 								[]models.Order
	ResponseCode 					int
	ResponseDescription 	string
}

type OrderList struct {
	Orders 								[]models.Order
	ResponseCode 					int
	ResponseDescription 	string
}

type OrderWalletList struct {
	CurrencyWallets				[]models.Wallet
	RateCurrencyWallets 	[]models.Wallet
	OrderWallet 					*models.OrderWallet
	ResponseCode 					int
	ResponseDescription 	string
}

type MyBuyOrderList struct {
	Orders 								[]models.OrderBuy
	CurrencyWallets				[]models.Wallet
	RateCurrencyWallets 	[]models.Wallet
	OrderWallet 					*models.OrderWallet
	ResponseCode 					int
	ResponseDescription 	string
}

type MySellOrderList struct {
	Orders 								[]models.OrderSell
	CurrencyWallets				[]models.Wallet
	RateCurrencyWallets 	[]models.Wallet
	OrderWallet 					*models.OrderWallet
	ResponseCode 					int
	ResponseDescription 	string
}

type BuyOrderList struct {
	Orders 								[]models.OrderBuy
	ResponseCode 					int
	ResponseDescription 	string
}

type SellOrderList struct {
	Orders 								[]models.OrderSell
	ResponseCode 					int
	ResponseDescription 	string
}

type DefaultCurrencyPair struct {
	DefaultCurrencyPair 	models.OrderCurrency
	ResponseCode 					int
	ResponseDescription 	string
}

type UserCountryDetail struct {
	UserCountry 			models.UserCountry
	Country 				models.Country
}

type UserCountryList struct {
	Countries 				[]UserCountryDetail
	ResponseCode 			int
	ResponseDescription 	string
}

type CurrencyList struct {
	Currencies 				[]models.Currency
	ResponseCode 			int
	ResponseDescription 	string
}

type CurrencyListAll struct {
	AllCurrencies 			[]models.Currency
	FiatCurrencies 			[]models.Currency
	CryptoCurrencies 		[]models.Currency
	ResponseCode 			int
	ResponseDescription 	string
}

type CountryList struct {
	Countries 				[]models.Country
	ResponseCode 			int
	ResponseDescription 	string
}

type FieldDetails struct {
	Field 					models.Field
	Category 				[]models.FieldCategory
}

type FieldList struct {
	Fields 					[]FieldDetails
	ResponseCode 			int
	ResponseDescription 	string
}

type FieldListAll struct {
	KYC 					[]FieldDetails
	ADDRESS 				[]FieldDetails
	BANK 					[]FieldDetails
	ResponseCode 			int
	ResponseDescription 	string
}

type DataDetail struct {
    Field 					models.Field
    FieldCategory 			*models.FieldCategory
    DataText 				*models.DataText
    DataCategory 			*models.DataCategory
    DataFile 				*models.DataFile
}

type DataGroup struct {
    Data 					models.Data
    DataDetail      		[]DataDetail
}

type DataList struct {
	DataGroup 				[]DataGroup
	ResponseCode 			int
	ResponseDescription 	string
}

type DataFileId struct {
	Id 						int64
	ResponseCode 			int
	ResponseDescription 	string
}

type TransferWalletDetail struct {
	//Transfer 				models.Transfer
	FromWallet 				*models.Wallet
	ToWallet 				*models.Wallet
}

type TransferPayeeDetail struct {
	//Transfer 				models.Transfer
	FromWallet 				*models.Wallet
	ToPayee 				*models.Payee
}

type TransferBankDetail struct {
	//Transfer 				models.Transfer
	FromWallet 				*models.Wallet
	ToData 					*models.DataText
}

type TransferData struct {
	Transfer 				models.Transfer
	Wallet 					TransferWalletDetail
	Payee 					TransferPayeeDetail
	Bank 					TransferBankDetail
}

type TransferMasterData struct {
	TransferMaster 			models.TransferMaster
	Currency 				models.Currency
	TransferData 			[]TransferData
}

type TransferList struct {
	TransferMasterData 		[]TransferMasterData
	ResponseCode 			int
	ResponseDescription 	string
}

type PayeeData struct {
	Payee 					models.Payee
	Crypto 					*models.PayeeCrypto
}

type PayeeMasterData struct {
	PayeeMaster 			models.PayeeMaster
	PayeeData 				[]PayeeData
}

type PayeeList struct {
	PayeeMasterData 		[]PayeeMasterData
	ResponseCode 			int
	ResponseDescription 	string
}

type Payee struct {
    WalletCrypto 			[]WalletData
    WalletFiat 				[]WalletData
	PayeeCrypto 			[]PayeeData
}

type PayeeListAll struct {
	Payee 					Payee
	ResponseCode 			int
	ResponseDescription 	string
}
