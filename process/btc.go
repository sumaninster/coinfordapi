package process

import (
	"coinford_api/models"
	"coinford_api/configs"
	"github.com/btcsuite/btcrpcclient"
	//"github.com/btcsuite/btcd/btcjson"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil"
	"fmt"
	"time"
	"log"
)

type BTCProcess struct {
	client 		*btcrpcclient.Client
	currencyid 	int64
}

func (btcp *BTCProcess) Process() {
	var err error
	btcp.client, err = btcp.Connect()
	btcp.currencyid = models.GetCurrencyId("BTC")
	if err == nil {
		for {
			fmt.Println("BTC Received: ", time.Now(), "\n------------------------------------------")
			btcp.receive()
			fmt.Println("BTC Sent: ", time.Now(), "\n------------------------------------------")
			btcp.send()
			fmt.Println("BTC Process End: ", time.Now(), "\n------------------------------------------")
			time.Sleep(3 * time.Second)
		}
	}
}

func (btcp *BTCProcess) GetNewAddress() (string, error) {
	var err error
	btcp.client, err = btcp.Connect()
	if err == nil {
		address, err := btcp.client.GetNewAddress(configs.BTC_IN_ACCOUNT)
		if err == nil {
			return address.String(), nil
		}
		return *configs.NullString, err
	}
	return *configs.NullString, err
}

func (btcp *BTCProcess) receive() error {
	var add []btcutil.Address
	add, _ = btcp.client.GetAddressesByAccount(configs.BTC_IN_ACCOUNT)
	for _, address := range add {
		amount,_ := btcp.client.GetReceivedByAddressMinConf(address, configs.BTC_MIN_CONF)
		fmt.Println(address,": ", amount.ToBTC())
		if amount.ToBTC() > 0 {
			walletCrypto := models.WalletCrypto{Address: address.String()}
			err := walletCrypto.Read("address")
			if err == nil {
				wallet := models.Wallet{Id: walletCrypto.WalletId}
				err := wallet.Read("id")
				if err == nil {
					wallet.Amount += amount.ToBTC()
					wallet.UpdatedAt = time.Now()
					err := wallet.Update()
					if err == nil {
						btcp.client.SetAccount(address, configs.BTC_VAULT_ACCOUNT)
						btcp.send_money(configs.BTC_VAULT_ACCOUNT, configs.BTC_OUT_ADDRESS, amount)
						btcp.client.SetAccount(address, configs.BTC_IN_ACCOUNT)
					} else {
						return err
					}
				} else {
					return err
				}
				btcp.client.SetAccount(address, configs.BTC_IN_ACCOUNT)
			} else {
				amount,_ := btcp.client.GetReceivedByAddress(address)
				fmt.Println(address,": ", amount.ToBTC())
				btcp.client.SetAccount(address, configs.BTC_OUT_ACCOUNT)
			}
		}
	}
	return nil
}

func (btcp *BTCProcess) send() {
	var err error
	btcp.client, err = btcp.Connect()
	num, transfers, err := models.TransferProcess(btcp.currencyid)
	fmt.Println(err)
	if err == nil && num > 0 {
		for _, transfer := range transfers {
			amount, _ := btcutil.NewAmount(transfer.Amount)
			btcp.send_money(configs.BTC_OUT_ACCOUNT, transfer.CryptoAddress(), amount)
			models.DebitWallet(transfer.Id, transfer.Amount, btcp.currencyid)
		}
	}
}

func (btcp *BTCProcess) send_money(account string, address string, amount btcutil.Amount) {
	var err error
	btcp.client, err = btcp.Connect()
	toadd, _ := btcutil.DecodeAddress(address, &chaincfg.MainNetParams)
	txSha, err := btcp.client.SendFromMinConf(account, toadd, amount, configs.BTC_MIN_CONF)
	if err != nil {
		log.Fatalf("error SendBitcoin: %v", err)
	}
	log.Printf("SendBitcoin completed! tx sha is: %s", txSha.String())
}

func (btcp *BTCProcess) Connect() (*btcrpcclient.Client, error) {
	wpassphrase := "suman123"
	var wtimeout int64
	wtimeout = 60 //seconds
	client, err := btcrpcclient.New(&btcrpcclient.ConnConfig{
		HTTPPostMode: true,
		DisableTLS:   true,
		Host:         "127.0.0.1:18332",
		User:         "bitcoinrpc",
		Pass:         "7003dd926de77d98d64d9c8152a3ec68",
	}, nil)
	if err == nil {
		client.WalletPassphrase(wpassphrase, wtimeout)
	}
	return client, err
}