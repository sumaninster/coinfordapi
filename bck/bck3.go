/*
func WalletFiatPrimaryFalse(walletId int64) (int64, []WalletFiat, error) {
	var table WalletFiat
	var walletFiats []WalletFiat
	num, err := orm.NewOrm().QueryTable(table).Filter("wallet_id", walletId).OrderBy("-created_at").All(&walletFiats)
	if err == nil && num > 0 {
		for _, v := range walletFiats {
			v.Primary = false
			v.Update()
		}
	}
	return num, walletFiats, err
}

func PrimaryWalletFiat(walletId int64) (int64, WalletFiat, error) {
	var table WalletFiat
	var walletFiats []WalletFiat
	num, err := orm.NewOrm().QueryTable(table).Filter("wallet_id", walletId).Filter("primary", true).Filter("deleted_at__isnull", true).OrderBy("-created_at").All(&walletFiats)
	return num, walletFiats[0], err
}*/

type TransferWalletDetail struct {
    Transfer 				models.Transfer
    Payee 					models.WalletCrypto
}

type TransferPayeeDetail struct {
    Transfer 				models.Transfer
    Payee 					models.PayeeCrypto
}

type TransferBankDetail struct {
    Transfer 				models.Transfer
    Payee 					models.Data
}

type TransferDetail struct {
	TransferMaster 			*models.TransferMaster
    Wallet 					[]models.
    Payee 					[]TransferPayeeDetail
    Bank 					[]TransferBankDetail
}

				switch transferMaster.TransferType {
				case "WALLET":
					transferCrypto := models.TransferCrypto{TransferId: transfer.Id}
					transferCrypto.Read("transfer_id")
					transferData.Wallet = &transferCrypto
				case "PAYEE":
				case "BANK":
					transferFiat := models.TransferFiat{TransferId: transfer.Id}
					err := transferFiat.Read("transfer_id")
					if err == nil {
						field := models.Field{CountryId: currency.CountryId, FieldType: "BANK", Key: "PRIMARY"}
						err = field.Read("country_id", "field_type", "key")
						if err == nil {
							dataText := models.DataText{DataId: transferFiat.DataId, FieldId: field.Id}
							transferData.Fiat = &dataText
						}
					}
				}

if(rqd.Primary) {
	models.UpdatePrimaryFalse(user.Id, rqd.CountryId, rqd.FieldType)
}
	

func (ft *TransferController) transferWalletDetail(transferId int64) models.Wallet {
	sc := models.TransferSelfCrypto{TransferId: transferId}
	err := sc.Read()
	if err == nil {
		walletAddress := models.WalletCrypto{Id: sc.WalletCryptoId}
		err = walletAddress.Read()
		if err == nil {
			return walletAddress
		}
	}
	return models.WalletCrypto{}
}

func (ft *TransferController) transferPayeeDetail(transferId int64) models.Payee {
	pc := models.TransferPayeeCrypto{TransferId: transferId}
	err := pc.Read()
	if err == nil {
		payeeCrypto := models.PayeeCrypto{Id: pc.PayeeCryptoId}
		err = payeeCrypto.Read()
		if err == nil {
			return payeeCrypto
		}
	}
	return models.PayeeCrypto{}
}

func (ft *TransferController) transferBankDetail(transferId int64) models.DataText {
	sb := models.TransferSelfBank{TransferId: transferId}
	err := sb.Read()
	if err == nil {
		data := models.Data{Id: sb.DataId}
		err = data.Read()
		if err == nil {
			return data
		}
	}
	return models.Data{}
}

/*
func (ft *TransferController) prepareTransferList(userId int64, transferSearch TransferSearch) TransferList {
	var walletCryptos []TransferSelfCryptoDetail
	var payeeCryptos []TransferPayeeCryptoDetail
	var selfBanks []TransferSelfBankDetail

	num, transfers, err := models.UserTransfers(userId, transferSearch.CurrencyId)
	if err == nil && num > 0 {
		for _, transfer := range transfers {
			switch transfer.TransferType {
			case "BANK":
				sb_data := ft.transferSelfBankDetail(transfer.Id)
				selfBank := TransferSelfBankDetail{Transfer: transfer, Payee: sb_data}
				selfBanks = append(selfBanks, selfBank)
			case "WALLET":
				sc_data := ft.transferSelfCryptoDetail(transfer.Id)
				walletCrypto := TransferSelfCryptoDetail{Transfer: transfer, Payee: sc_data}
				walletCryptos = append(walletCryptos, walletCrypto)

				pb_data := ft.transferPayeeCryptoDetail(transfer.Id)
				payeeCrypto := TransferPayeeCryptoDetail{Transfer: transfer, Payee: pb_data}
				payeeCryptos = append(payeeCryptos, payeeCrypto)
			case "PAYEE":
				sb_data := ft.transferSelfBankDetail(transfer.Id)
				selfBank := TransferSelfBankDetail{Transfer: transfer, Payee: sb_data}
				selfBanks = append(selfBanks, selfBank)
			}
		}
	}

	return TransferList{TransferMasterData: transferMasterData, ResponseCode: 200, ResponseDescription: "List of Transfers"}
}*/

/*
func (ft *TransferController) prepareTransferList(userId int64, transferSearch TransferSearch) TransferList {
	var walletCryptos []TransferSelfCryptoDetail
	var payeeCryptos []TransferPayeeCryptoDetail
	var selfBanks []TransferSelfBankDetail

	num, transfers, err := models.UserTransfers(userId, transferSearch.CurrencyId)
	if err == nil && num > 0 {
		for _, transfer := range transfers {
			switch transferSearch.TransferType {
			case "SELF":
				switch transferSearch.CurrencyType {
				case "FIAT":
					sb_data := ft.transferSelfBankDetail(transfer.Id)
					selfBank := TransferSelfBankDetail{Transfer: transfer, Payee: sb_data}
					selfBanks = append(selfBanks, selfBank)
				case "CRYPTO":
					sc_data := ft.transferSelfCryptoDetail(transfer.Id)
					walletCrypto := TransferSelfCryptoDetail{Transfer: transfer, Payee: sc_data}
					walletCryptos = append(walletCryptos, walletCrypto)
				}
			case "PAYEE":
				switch transferSearch.CurrencyType {
				case "CRYPTO":
					pc_data := ft.transferPayeeCryptoDetail(transfer.Id)
					payeeCrypto := TransferPayeeCryptoDetail{Transfer: transfer, Payee: pc_data}
					payeeCryptos = append(payeeCryptos, payeeCrypto)
				}
			case "ALL":
				switch transferSearch.CurrencyType {
				case "FIAT":
					sb_data := ft.transferSelfBankDetail(transfer.Id)
					selfBank := TransferSelfBankDetail{Transfer: transfer, Payee: sb_data}
					selfBanks = append(selfBanks, selfBank)
				case "CRYPTO":
					sc_data := ft.transferSelfCryptoDetail(transfer.Id)
					walletCrypto := TransferSelfCryptoDetail{Transfer: transfer, Payee: sc_data}
					walletCryptos = append(walletCryptos, walletCrypto)

					pb_data := ft.transferPayeeCryptoDetail(transfer.Id)
					payeeCrypto := TransferPayeeCryptoDetail{Transfer: transfer, Payee: pb_data}
					payeeCryptos = append(payeeCryptos, payeeCrypto)
				}
			}
		}
	}

	return TransferList{SelfCrypto: walletCryptos, PayeeCrypto: payeeCryptos, SelfBank: selfBanks, ResponseCode: 200, ResponseDescription: "List of Transfers"}
}*/

type PayeeCryptoAdd struct {
	Address 			string
	CurrencyId 			int64
	Email 				string
	Nickname 			string
	Token 				string
}

type SelfFiatDetail struct {
    Data 					models.Data
    DataText 				*models.DataText
    DataCategory 			*models.DataCategory
    FieldCategory 			*models.FieldCategory
}

type SelfFiat struct {
    SelfFiatDetail 			[]SelfFiatDetail
    Field 					*models.Field
}

    //SelfFiat 				*SelfFiat

		//if err == nil {
			payee := Payee{WalletCrypto: walletCrypto, WalletFiat: walletFiat, PayeeCrypto: payeeCrypto}
			p.Data["json"] = PayeeListAll{Payee: payee, ResponseCode: 200, ResponseDescription: "Payee List"}
		/*} else {
			fmt.Println(err)
			p.Data["json"] = CommonResponse{ResponseCode: 404, ResponseDescription: "Unable to read data"}
		}*/
		
func (p *PayeeController) updateStatus(payeeMasterId int64, rqd *PayeeDetails) {
	var num int64
	var err error
	if rqd.Primary {
		num, _, err = models.PayeePrimaryFalse(payeeMasterId)
	} else {
		num, _, err = models.Payees(payeeMasterId)
	}

	if num == 0 && err == nil {
		if rqd.Nickname == *configs.NullString {
			rqd.Nickname = "Primary"
		}
		rqd.Primary = true
	} else {
		if rqd.Nickname == *configs.NullString {
			rqd.Nickname = "No Name"
		}
	}
}

func (dc *DataController) prepareFieldCategory(field *models.Field, dataId int64) {
	if field.HasCategory {
		fc := models.FieldCategory{Id: data.Id, FieldId: field.Id}
		err = fc.Read("data_id","field_id")
		if err == nil {
			datadetail.FieldCategory = &fc
		}
	}
}

func (ft *TransferController) ListPayee() {
	var rqd TransferSearchAll
	json.Unmarshal(ft.Ctx.Input.RequestBody, &rqd)
	jres, user, isLogin, _ := apiAuthenticate(rqd.Token)
	if isLogin {
	    var walletCrypto 	[]models.WalletData
	    var walletFiat 		[]models.WalletData
	    var payeeCrypto   	[]models.PayeeData
	    var err 			error
		switch rqd.TransferSearch.TransferType {
		case "WALLET":
			walletMaster := models.WalletMaster{UserId: user.Id, CurrencyId: rqd.TransferSearch.CurrencyId}
			walletMaster.Read("user_id", "currency_id")
			wallets := models.Wallets(walletMaster.Id)
			for _, wallet := wallets {
				switch rqd.TransferSearch.CurrencyType {
				case "CRYPTO":
					wc := models.WalletCrypto{WalletId: wallet.Id}
					wc.Read()
					walletCrypto = WalletData{Wallet: wallet, Crypto: wc, Fiat: nil}
				case "FIAT":
					wf := models.WalletFiat{WalletId: wallet.Id}
					wf.Read()
					dt := models.DataText{DataId: wf.DataId}
					dt.Read()
					walletFiat = WalletData{Wallet: wallet, Crypto: nil, Fiat: dt}
				}
			}
		case "PAYEE":
			payeeMaster := models.PayeeMaster{UserId: user.Id, CurrencyId: rqd.TransferSearch.CurrencyId}
			payeeMaster.Read("user_id", "currency_id")
			payees := models.Payees(payeetMaster.Id)
			for _, payee := payees {
				switch rqd.TransferSearch.CurrencyType {
				case "CRYPTO":
					pc := models.PayeeCrypto{PayeeId: payee.Id}
					pc.Read()
					payeeCrypto = PayeeData{Payee: payee, Crypto: pc}
				}
			}
		}
	} else {
		ft.Data["json"] = jres
	}
	ft.ServeJSON()
}


func (ft *TransferController) ListPayee() {
	var rqd TransferSearchAll
	json.Unmarshal(ft.Ctx.Input.RequestBody, &rqd)
	jres, user, isLogin, _ := apiAuthenticate(rqd.Token)
	if isLogin {
	    var walletCrypto 	[]models.WalletCrypto
	    var walletFiat 		[]models.WalletFiat
	    var payeeCrypto   	[]models.PayeeCrypto
	    var selfFiat 		*SelfFiat
	    var err 			error
		switch rqd.TransferSearch.TransferType {
		case "SELF":
			switch rqd.TransferSearch.CurrencyType {
			case "FIAT":
				currency := models.Currency{Id: rqd.TransferSearch.CurrencyId}
				err = currency.Read("id")
				if err == nil {
					field := models.Field{CountryId: currency.CountryId, FieldType: "BANK", Key: "PRIMARY"}
					err = field.Read("country_id", "field_type", "key")
					if err == nil {
						var datas 			[]models.Data
						var selfFiatDetails []SelfFiatDetail
						_, datas, err = models.DataList(user.Id, currency.CountryId, "BANK")
						for _, data := range datas {
						    selfFiatDetail := SelfFiatDetail{Data: data}
							if field.HasInputText {
								dataText := models.DataText{DataId: data.Id, FieldId: field.Id}
								err = dataText.Read("data_id", "field_id")
								selfFiatDetail.DataText = &dataText
							} else if field.HasCategory {
								dataCategory := models.DataCategory{DataId: data.Id, FieldId: field.Id}
								err = dataCategory.Read("data_id", "field_id")
								selfFiatDetail.DataCategory = &dataCategory
								if err == nil {
									fieldCategory := models.FieldCategory{FieldId: field.Id}
									err = fieldCategory.Read("field_id")
									selfFiatDetail.FieldCategory = &fieldCategory
								}
							}
							switch field.DataType {
							case "TEXT":
								dt := models.DataText{DataId: data.Id, FieldId: field.Id}
								err := dt.Read("data_id", "field_id")
								if err == nil {
									datadetail.DataText = &dt
								}
							case "CATEGORY":
					 			dc := models.DataCategory{DataId: data.Id, FieldId: field.Id}
								err := dc.Read("data_id", "field_id")
								if err == nil {
									datadetail.DataCategory = &dc
								}
							case "FILE":
								df := models.DataFile{DataId: data.Id, FieldId: field.Id}
								err := df.Read("data_id", "field_id")
								if err == nil {
									datadetail.DataFile = &df
								}
							}
							selfFiatDetails = append(selfFiatDetails, selfFiatDetail)
						}
						selfFiat1 := SelfFiat{Field: &field, SelfFiatDetail: selfFiatDetails}
						selfFiat = &selfFiat1
					}
				}
				wallet := models.Wallet{WalletMasterId: user.Id}//, CurrencyId: rqd.TransferSearch.CurrencyId}
				err = wallet.Read("user_id", "currency_id")
				if err == nil {
					_, walletFiat, err = models.WalletFiats(wallet.Id)
				}
			case "CRYPTO":
				wallet := models.Wallet{WalletMasterId: user.Id}
				err = wallet.Read("user_id", "currency_id")
				if err == nil {
					_, walletCrypto, err = models.WalletCryptos(wallet.Id)
				}
			}
		case "PAYEE":
			switch rqd.TransferSearch.CurrencyType {
			case "CRYPTO":
				wallet := models.Wallet{WalletMasterId: user.Id}
				err = wallet.Read("user_id", "currency_id")
				if err == nil {
					_, walletCrypto, err = models.WalletCryptos(wallet.Id)
				}
				_, payeeCrypto, err = models.ListPayeeCryptos(user.Id, rqd.TransferSearch.CurrencyId)
			}
		}
		if err == nil {
			payee := Payee{WalletCrypto: walletCrypto, WalletFiat: walletFiat, PayeeCrypto: payeeCrypto, SelfFiat: selfFiat}
			ft.Data["json"] = PayeeList{Payee: payee, ResponseCode: 200, ResponseDescription: "List of Transfers"}
		} else {
			fmt.Println(err)
			ft.Data["json"] = CommonResponse{ResponseCode: 404, ResponseDescription: "Unable to read data"}
		}
	} else {
		ft.Data["json"] = jres
	}
	ft.ServeJSON()
}

_, _, err := wallet.ReadOrCreate("user_id", "currency_id")
if err == nil {
	notDeleted := true
	if wallet.DeletedAt != *configs.NullTime {
		wallet.DeletedAt = *configs.NullTime
		err = wallet.Update()
		notDeleted = false
	}
	if err == nil {
		if notDeleted {
			err = w.addWalletCrypto(&wallet, &rqd)
		}
		if err == nil {
			w.Data["json"] = w.prepareWalletList(user.Id)
		} else {
			w.Data["json"] = CommonResponse{ResponseCode: 404, ResponseDescription: "Error Creating New Wallet 1"}
		}
	} else {
		w.Data["json"] = CommonResponse{ResponseCode: 404, ResponseDescription: "Error Creating New Wallet 2"}
	}
}

		transfer := models.Transfer{, Amount: rqd.Amount}
		err := transfer.Insert()
		if err == nil {
			switch rqd.TransferSearch.TransferType {
			case "SELF":
				switch rqd.TransferSearch.CurrencyType {
				case "FIAT":
					transferSelfBank := models.TransferSelfBank{TransferId: transfer.Id, DataId: rqd.PayeeId}
					err = transferSelfBank.Insert()
				case "CRYPTO":
					transferSelfCrypto := models.TransferSelfCrypto{TransferId: transfer.Id, WalletCryptoId: rqd.PayeeId}
					err = transferSelfCrypto.Insert()
				}
			case "PAYEE":
				switch rqd.TransferSearch.CurrencyType {
				case "CRYPTO":
					transferPayeeCrypto := models.TransferPayeeCrypto{TransferId: transfer.Id, PayeeCryptoId: rqd.PayeeId}
					err = transferPayeeCrypto.Insert()
				}
			}
			if err == nil {
				ft.Data["json"] = ft.prepareTransferList(user.Id, rqd.TransferSearch)
			} else {
				ft.Data["json"] = CommonResponse{ResponseCode: 404, ResponseDescription: "Error Creating New Transfer 1"}
			}
		}