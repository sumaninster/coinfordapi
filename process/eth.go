package process

import (
	"coinford_api/models"
	"coinford_api/configs"
	"fmt"
	"time"
	"net/http"
	"encoding/json"
	"bytes"
	"io/ioutil"
	"strconv"
	//"errors"
)

type ETHProcess struct {
	currencyid 	int64
}

func (ethp *ETHProcess) Process() {
	ethp.currencyid = models.GetCurrencyId("ETH")
	//for {
		ethp.receive()
		ethp.send()
		fmt.Println("ETH: ", time.Now())
		time.Sleep(3 * time.Second)
	//}
}

func (ethp *ETHProcess) GetNewAddress(passphrase string) (string, error) {
	jsonstr := fmt.Sprintf(`{"jsonrpc":"2.0","method":"personal_newAccount","params":["%s"],"id":1}`, passphrase)
	body, err := ethp.query(jsonstr)
	var rqd ETHNewAccount
	json.Unmarshal(body, &rqd)
	return rqd.Result, err
}

func (ethp *ETHProcess) receive() error {
	jsonstr := `{"jsonrpc":"2.0","method":"eth_accounts","params":[],"id":1}`
	body, _ := ethp.query(jsonstr)
	var rqd ETHListAccount
	json.Unmarshal(body, &rqd)
	accounts := rqd.Result
	for _, address := range accounts {
		amount := ethp.getBalance(address)
		fmt.Println(address,": ", ethp.toETH(amount))
		if amount > 0 {
			walletCrypto := models.WalletCrypto{Address: address}
			err := walletCrypto.Read("address")
			if err == nil {
				walletPassphrase := models.WalletPassphrase{WalletId: walletCrypto.WalletId}
				err := walletPassphrase.Read("wallet_id")
				if err == nil {
					wallet := models.Wallet{Id: walletCrypto.WalletId}
					err := wallet.Read("id")
					if err == nil {
						wallet.Amount += ethp.toETH(amount)
						wallet.UpdatedAt = time.Now()
						err := wallet.Update()
						if err == nil {
							ethp.transfer(address, walletPassphrase.Passphrase, configs.ETH_OUT_ADDRESS, amount)
						} else {
							return err
						}
					} else {
						return err
					}
				} else {
					return err
				}
			} else {
				return err
			}
		}
	}
	return nil
}

func (ethp *ETHProcess) send() {
	num, transfers, err := models.TransferProcess(ethp.currencyid)
	fmt.Println(err)
	if err == nil && num > 0 {
		for _, transfer := range transfers {
			amount := int64(transfer.Amount)*ethp.oneEther()
			ethp.transfer(configs.ETH_OUT_ADDRESS, configs.ETH_OUT_PASSPHRASE, transfer.CryptoAddress(), amount)
			models.DebitWallet(transfer.Id, transfer.Amount, ethp.currencyid)
		}
	}
}

func (ethp *ETHProcess) oneEther() int64 {
	return 1000000000000000000
}

func (ethp *ETHProcess) toETH(amount int64) float64 {
	return float64(amount)/float64(ethp.oneEther())
}

func (ethp *ETHProcess) getBalance(address string) int64 {
	jsonstr := fmt.Sprintf(`{"jsonrpc":"2.0","method":"eth_getBalance","params":["%s", "latest"], "id":1}`, address)
	body, _ := ethp.query(jsonstr)
	var rqd ETHCommonResponse
	json.Unmarshal(body, &rqd)
	balance, _ := strconv.ParseInt(rqd.Result, 0, 64)
	return balance
}

func (ethp *ETHProcess) unlockAccount(address string, passphrase string) {
	jsonstr := fmt.Sprintf(`{"method": "personal_unlockAccount", "params": ["%s", "%s"], "id":1}`, address, passphrase)
	ethp.query(jsonstr)
}

func (ethp *ETHProcess) transfer(fromaddress string, frompassphrase string, toaddress string, amount int64) bool {
	amount_ether := uint64(amount)
	tx := fmt.Sprintf(`{"from": "%s", "to": "%s", "value": "0x%s"}`, fromaddress, toaddress, strconv.FormatUint(amount_ether, 16))
	jsonstr := fmt.Sprintf(`{"jsonrpc":"2.0","method":"personal_sendTransaction","params":[%s, "%s"], "id":1}`, tx, frompassphrase)
	ethp.query(jsonstr)
	return true
}

func (ethp *ETHProcess) query(jsonstr string) ([]byte, error) {
	fmt.Println("Request Body:", string(jsonstr))
	jsonbytes := []byte(jsonstr)
	resp, _ := http.Post(configs.ETHAuthUrl, "application/json", bytes.NewBuffer(jsonbytes))
	body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("Response Body:", string(body))
    return body, nil
}