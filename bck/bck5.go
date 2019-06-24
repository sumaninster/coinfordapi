
func (ft *TransferController) transferWalletDetail(transferId int64) []TransferWalletDetail{
	var transferWalletDetail TransferWalletDetail
	t := models.TransferWallet{TransferId: transferId}
	err := t.Read("transfer_id")
	if err == nil {
		fromWallet := models.Wallet{Id: t.FromWalletId}
		err = fromWallet.Read("id")
		if err == nil {
			transferWalletDetail.FromWallet = &fromWallet
		}
		toWallet := models.Wallet{Id: t.ToWalletId}
		err = toWallet.Read("id")
		if err == nil {
			transferWalletDetail.ToWallet = &toWallet
		}
	}
	return transferWalletDetail
}

func (ft *TransferController) transferPayeeDetail(transferId int64) []TransferPayeeDetail {
	var transferPayeeDetail TransferPayeeDetail
	t := models.TransferPayee{TransferId: transferId}
	err := t.Read("transfer_id")
	if err == nil {
		fromWallet := models.Wallet{Id: t.FromWalletId}
		err = fromWallet.Read("id")
		if err == nil {
			transferPayeeDetail.FromWallet = &fromWallet
		}
		toPayee := models.Payee{Id: t.ToPayeeId}
		err = toPayee.Read("id")
		if err == nil {
			transferPayeeDetail.ToPayee = &toPayee
		}
	}
	return transferPayeeDetail
}

func (ft *TransferController) transferBankDetail(transferId int64) []TransferBankDetail {
	var transferBankDetail TransferBankDetail
	t := models.TransferBank{TransferId: transferId}
	err := t.Read("transfer_id")
	if err == nil {
		fromWallet := models.Wallet{Id: t.FromWalletId}
		err = fromWallet.Read("id")
		if err == nil {
			transferBankDetail.FromWallet = &fromWallet
		}
		toData := models.DataText{Id: t.ToDataId}
		err = toData.Read("id")
		if err == nil {
			transferBankDetail.ToData = &toData
		}
	}
	return transferBankDetail
}

func (ft *TransferController) prepareTransferList(userId int64) TransferList {
	var transferMasterDatas []TransferMasterData
	_, transferMasters, _ := models.TransferMasters(userId)
	for _, transferMaster := range transferMasters {
		currency := models.Currency{Id: transferMaster.CurrencyId}
		err := currency.Read("Id")
		if err == nil {
			transferMasterData := TransferMasterData{TransferMaster: transferMaster, Currency: currency}
			_, transfers, _ := models.Transfers(transferMaster.Id)
			var transferDatas []TransferData
			transferData := TransferData{Transfer: transfer}

			wallets := ft.transferWalletDetail(transferMaster.Id)
			transferData.Wallet = wallets
			
			payees := ft.transferPayeeDetail(transferMaster.Id)
			transferData.Payee = payees

			bank := ft.transferBankDetail(transferMaster.Id)
			transferData.Bank = banks

			transferDatas = append(transferDatas, transferData)
			transferMasterData.TransferData = transferDatas
			transferMasterDatas = append(transferMasterDatas, transferMasterData)
		}
	}
	return TransferList{TransferMasterData: transferMasterDatas, ResponseCode: 200, ResponseDescription: "List of Transfers"}
}