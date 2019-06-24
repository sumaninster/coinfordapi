
func PrimaryWalletLockAmount(userId int64, currencyId int64, amount float64) error {
	wallet, err := PrimaryWallet(userId, currencyId)
	if err == nil {
		if (wallet.Amount - wallet.AmountLocked) >= amount {
			wallet.AmountLocked += amount
			wallet.Update()
		} else {
			errors.New("Insufficient balance in primary wallet")
		}
	} else {
		return err
	}
	return nil
}

func PrimaryWalletUnlockAmount(userId int64, currencyId int64, amount float64) error {
	wallet, err := PrimaryWallet(userId, currencyId)
	if err == nil {
		if wallet.AmountLocked >= amount {
			wallet.AmountLocked -= amount
			wallet.Update()
		} else {
			errors.New("Insufficient amount locked in primary wallet")
		}
	} else {
		return err
	}
	return nil
}

/*
func CheckWalletUnlockAmount(userId int64, currencyId int64, receiveCurrencyId int64, walletId int64, amount float64) (Wallet, error) {
	wallet := Wallet{Id: walletId}
	err := wallet.Read("id")
	if err == nil {
		walletMaster := WalletMaster{Id: wallet.WalletMasterId}
		err = walletMaster.Read("id")
		if err == nil {
			if walletMaster.UserId == userId && walletMaster.CurrencyId == receiveCurrencyId {
				if wallet.AmountLocked >= amount {
					wallet.AmountLocked -= amount
					wallet.Update()
				} else {
					errors.New("Insufficient amount locked in primary wallet")
				}
			} else {
				err = errors.New("Wrong user or currency")
			}
		}
	}
	return wallet, err
}*/
