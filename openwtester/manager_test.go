package openwtester

import (
	"fmt"
	"github.com/blocktree/openwallet/log"
	"github.com/blocktree/openwallet/openw"
	"github.com/blocktree/openwallet/openwallet"
	"path/filepath"
	"testing"
)

var (
	testApp        = "blockethereum-adapter"
	configFilePath = filepath.Join("conf")
)

func testInitWalletManager() *openw.WalletManager {
	log.SetLogFuncCall(true)
	tc := openw.NewConfig()

	tc.ConfigDir = configFilePath
	tc.EnableBlockScan = false
	tc.SupportAssets = []string{
		"BETH",
	}
	return openw.NewWalletManager(tc)
	//tm.Init()
}

func TestWalletManager_CreateWallet(t *testing.T) {
	tm := testInitWalletManager()
	w := &openwallet.Wallet{Alias: "HELLO BETH", IsTrust: true, Password: "12345678"}
	nw, key, err := tm.CreateWallet(testApp, w)
	if err != nil {
		log.Error(err)
		return
	}

	log.Info("wallet:", nw)
	log.Info("key:", key)

}

func TestWalletManager_GetWalletInfo(t *testing.T) {

	tm := testInitWalletManager()

	wallet, err := tm.GetWalletInfo(testApp, "WBup5yKfqFW17qTT3yR4HpdDt8uEj4MUc1")
	if err != nil {
		log.Error("unexpected error:", err)
		return
	}
	log.Info("wallet:", wallet)
}

func TestWalletManager_GetWalletList(t *testing.T) {

	tm := testInitWalletManager()

	list, err := tm.GetWalletList(testApp, 0, 10000000)
	if err != nil {
		log.Error("unexpected error:", err)
		return
	}
	for i, w := range list {
		log.Info("wallet[", i, "] :", w)
	}
	log.Info("wallet count:", len(list))

	tm.CloseDB(testApp)
}

func TestWalletManager_CreateAssetsAccount(t *testing.T) {

	tm := testInitWalletManager()

	walletID := "WBup5yKfqFW17qTT3yR4HpdDt8uEj4MUc1"
	account := &openwallet.AssetsAccount{Alias: "feesSupport", WalletID: walletID, Required: 1, Symbol: "BETH", IsTrust: true}
	account, address, err := tm.CreateAssetsAccount(testApp, walletID, "12345678", account, nil)
	if err != nil {
		log.Error(err)
		return
	}

	log.Info("account:", account)
	log.Info("address:", address)

	tm.CloseDB(testApp)
}

func TestWalletManager_GetAssetsAccountList(t *testing.T) {

	tm := testInitWalletManager()

	walletID := "WBup5yKfqFW17qTT3yR4HpdDt8uEj4MUc1"
	list, err := tm.GetAssetsAccountList(testApp, walletID, 0, 10000000)
	if err != nil {
		log.Error("unexpected error:", err)
		return
	}
	for i, w := range list {
		log.Info("account[", i, "] :", w)
	}
	log.Info("account count:", len(list))

	tm.CloseDB(testApp)

}

func TestWalletManager_CreateAddress(t *testing.T) {

	tm := testInitWalletManager()

	walletID := "WBup5yKfqFW17qTT3yR4HpdDt8uEj4MUc1"
	accountID := "5f41cwEcQNDguhLK5LKRh8bVz42mXwdLGgpuTmqq1uKM"
	address, err := tm.CreateAddress(testApp, walletID, accountID, 5)
	if err != nil {
		log.Error(err)
		return
	}



	log.Info("address:", address)

	tm.CloseDB(testApp)
}

func TestWalletManager_GetAddressList(t *testing.T) {

	tm := testInitWalletManager()

	walletID := "WBup5yKfqFW17qTT3yR4HpdDt8uEj4MUc1"
	//accountID := "3csEgf2TcxwNeoFSTsePXFVmzcyNhHAS49jsTv99n1Nv"
	accountID := "5f41cwEcQNDguhLK5LKRh8bVz42mXwdLGgpuTmqq1uKM"
	list, err := tm.GetAddressList(testApp, walletID, accountID, 0, -1, false)
	if err != nil {
		log.Error("unexpected error:", err)
		return
	}
	for _, w := range list {
		fmt.Printf("%s\n", w.Address)
	}
	log.Info("address count:", len(list))

	tm.CloseDB(testApp)
}

