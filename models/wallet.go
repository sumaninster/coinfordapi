package models

import (
	"github.com/astaxie/beego/orm"
	"errors"
)

func (w *Wallet) TableName() string {
    return "wallet"
}

func (w *Wallet) Insert() error {
	if _, err := orm.NewOrm().Insert(w); err != nil {
		return err
	}
	return nil
}

func (w *Wallet) Read(fields ...string) error {
	if err := orm.NewOrm().Read(w, fields...); err != nil {
		return err
	}
	return nil
}

func (w *Wallet) ReadOrCreate(field string, fields ...string) (bool, int64, error) {
	return orm.NewOrm().ReadOrCreate(w, field, fields...)
}

func (w *Wallet) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(w, fields...); err != nil {
		return err
	}
	return nil
}

func (w *Wallet) Delete(fields ...string) error {
	if _, err := orm.NewOrm().Delete(w, fields ...); err != nil {
		return err
	}
	return nil
}

func Wallets(walletMasterId int64) (int64, []Wallet, error) {
	var table Wallet
	var wallets []Wallet
	num, err := orm.NewOrm().QueryTable(table).Filter("wallet_master_id", walletMasterId).Filter("deleted_at__isnull", true).OrderBy("-updated_at").All(&wallets)
	return num, wallets, err
}

func MyWallets(user *User, currencyId int64) (int64, []Wallet, error) {
	walletMaster := WalletMaster{UserId: user.Id, CurrencyId: currencyId}
	err := walletMaster.Read("user_id", "currency_id")
	num, wallets, err := Wallets(walletMaster.Id)
	return num, wallets, err
}

func PrimaryWallet(userId int64, currencyId int64) (Wallet, error) {
	walletMaster := WalletMaster{UserId: userId, CurrencyId: currencyId}
	err := walletMaster.Read("user_id", "currency_id")
	var wallet Wallet
	if err == nil {
		wallet = Wallet{WalletMasterId: walletMaster.Id, Primary: true}
		err = wallet.Read("wallet_master_id", "primary")
	}
	return wallet, err
}

func CheckWalletLockAmount(userId int64,
	walletId int64, currencyId int64,
	receiveWalletId int64, receiveCurrencyId int64,
	amount float64) (error, error) {

	wallet := Wallet{Id: walletId}
	errW := wallet.Read("id")

	receiveWallet := Wallet{Id: receiveWalletId}
	errRW := receiveWallet.Read("id")

	if errW == nil && errRW == nil {

		walletMaster := WalletMaster{Id: wallet.WalletMasterId}
		errW = walletMaster.Read("id")

		receiveWalletMaster := WalletMaster{Id: receiveWallet.WalletMasterId}
		errRW = receiveWalletMaster.Read("id")

		if errW == nil && errRW == nil {
			if walletMaster.UserId == userId &&
			receiveWalletMaster.UserId == userId &&
			walletMaster.CurrencyId == currencyId &&
			receiveWalletMaster.CurrencyId == receiveCurrencyId  {
				if (wallet.Amount - wallet.AmountLocked) >= amount {
					wallet.AmountLocked += amount
					wallet.Update()
				} else {
					errW = errors.New("Insufficient balance in primary wallet")
				}
			} else {
				errW = errors.New("Wrong user or currency")
			}
		}
	}
	return errW, errRW
}

func CheckWalletUnlockAmount(userId int64,
	walletId int64, currencyId int64,
	receiveWalletId int64, receiveCurrencyId int64,
	amount float64) (error, error) {

	wallet := Wallet{Id: walletId}
	errW := wallet.Read("id")

	receiveWallet := Wallet{Id: receiveWalletId}
	errRW := receiveWallet.Read("id")

	if errW == nil && errRW == nil {

		walletMaster := WalletMaster{Id: wallet.WalletMasterId}
		errW = walletMaster.Read("id")

		receiveWalletMaster := WalletMaster{Id: receiveWallet.WalletMasterId}
		errRW = receiveWalletMaster.Read("id")

		if errW == nil && errRW == nil {
			if walletMaster.UserId == userId &&
			receiveWalletMaster.UserId == userId &&
			walletMaster.CurrencyId == currencyId &&
			receiveWalletMaster.CurrencyId == receiveCurrencyId  {
				if wallet.AmountLocked >= amount {
					wallet.AmountLocked -= amount
					wallet.Update()
				} else {
					errW = errors.New("Insufficient amount locked in primary wallet")
				}
			} else {
				errW = errors.New("Wrong user or currency")
			}
		}
	}
	return errW, errRW
}

func WalletPrimaryFalse(walletMasterId int64) (int64, []Wallet, error) {
	var table Wallet
	var wallets []Wallet
	num, err := orm.NewOrm().QueryTable(table).Filter("wallet_master_id", walletMasterId).OrderBy("-updated_at").All(&wallets)
	if err == nil && num > 0 {
		for _, v := range wallets {
			var update = false
			if v.Primary == true {
				v.Primary = false
				update = true
			}
			if v.Nickname == "Primary" {
				v.Nickname = "No Name"
				update = true
			}
			if update {
				v.Update()
			}
		}
	}
	return num, wallets, err
}

func CreditWallet(userId int64, amount float64, currencyId int64) error {
	wallet, err := PrimaryWallet(userId, currencyId)
	if err == nil {
		wallet.Amount += amount
		err = wallet.Update()
		return err
	}
	return err
}

func DebitWallet(userId int64, amount float64, currencyId int64) error {
	wallet, err := PrimaryWallet(userId, currencyId)
	if err == nil {
		wallet.Amount -= amount
		wallet.AmountLocked -= amount
		err = wallet.Update()
		return err
	}
	return err
}

func init() {
    orm.RegisterModel(new(Wallet))
}
