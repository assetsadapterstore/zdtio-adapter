/*
 * Copyright 2018 The openwallet Authors
 * This file is part of the openwallet library.
 *
 * The openwallet library is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The openwallet library is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Lesser General Public License for more details.
 */

package openwtester

import (
	"github.com/astaxie/beego/config"
	"github.com/blocktree/openwallet/openw"
	"path/filepath"
	"testing"

	"github.com/blocktree/openwallet/log"
	"github.com/blocktree/openwallet/openwallet"
)

func TestWalletManager_GetTransactions(t *testing.T) {
	tm := testInitWalletManager()
	list, err := tm.GetTransactions(testApp, 0, -1, "", false)
	if err != nil {
		log.Error("GetTransactions failed, unexpected error:", err)
		return
	}
	for i, tx := range list {
		log.Info("trx[", i, "] :", tx)
	}
	log.Info("trx count:", len(list))
}

func TestWalletManager_GetTransactionByWxID(t *testing.T) {
	tm := testInitWalletManager()
	wxID := openwallet.GenTransactionWxID(&openwallet.Transaction{
		TxID: "4aabaedba12594e869b99916dca8619132a96b7ea00a90f497f57d52c2c2fa68",
		Coin: openwallet.Coin{
			Symbol:     "ZDT",
			IsContract: false,
			ContractID: "",
		},
	})
	log.Info("wxID:", wxID)
	tx, err := tm.GetTransactionByWxID(testApp, wxID)
	if err != nil {
		log.Error("GetTransactionByTxID failed, unexpected error:", err)
		return
	}
	log.Info("tx:", tx)
}

func TestWalletManager_GetAssetsAccountBalance(t *testing.T) {
	tm := testInitWalletManager()
	walletID := "W59hhWFQ4NWt2WB1FPt8XJmi5q88fH4tyR"
	accountID := "6NywxLEwyU7oeaei2UknUNVWyuAwJzDDZMPAgZ4ed2J4"

	balance, err := tm.GetAssetsAccountBalance(testApp, walletID, accountID)
	if err != nil {
		log.Error("GetAssetsAccountBalance failed, unexpected error:", err)
		return
	}
	log.Info("balance:", balance)
}

func TestWalletManager_GetAssetsAccountTokenBalance(t *testing.T) {
	tm := testInitWalletManager()
	walletID := "WD1yNAJEhY61UGo4SdWuGuLyF5S1VV7CWn"
	accountID := "EYu9jxzw869PLo8vRwkte7TXLyrHskp4eVcyc5Vvs5yv"

	contract := openwallet.SmartContract{
		Address:  "tonydchan123:ZING",
		Protocol: "multiple-token",
		Symbol:   "ZING",
		Name:     "ZING",
		Decimals: 8,
	}

	balance, err := tm.GetAssetsAccountTokenBalance(testApp, walletID, accountID, contract)
	if err != nil {
		log.Error("GetAssetsAccountTokenBalance failed, unexpected error:", err)
		return
	}
	log.Info("balance:", balance.Balance)
}

func TestWalletManager_GetEstimateFeeRate(t *testing.T) {
	tm := testInitWalletManager()
	coin := openwallet.Coin{
		Symbol: "ZDT",
	}
	feeRate, unit, err := tm.GetEstimateFeeRate(coin)
	if err != nil {
		log.Error("GetEstimateFeeRate failed, unexpected error:", err)
		return
	}
	log.Std.Info("feeRate: %s %s/%s", feeRate, coin.Symbol, unit)
}


func TestGetAccountTokenBalance(t *testing.T) {
	symbol := "ZDT"
	assetsMgr, err := openw.GetAssetsAdapter(symbol)
	if err != nil {
		log.Error(symbol, "is not support")
		return
	}
	//读取配置
	absFile := filepath.Join(configFilePath, symbol+".ini")

	c, err := config.NewConfig("ini", absFile)
	if err != nil {
		return
	}
	assetsMgr.LoadAssetsConfig(c)
	sm := assetsMgr.GetSmartContractDecoder()

	contract := openwallet.SmartContract{
		Address:  "tonydchan123:ZING",
		Protocol: "",
		Symbol:   "ZDT",
		Name:     "ZING",
		Decimals: 8,
	}

	contractID := openwallet.GenContractID(contract.Symbol, contract.Address)
	log.Infof("contractID = %s", contractID)
	log.Infof("BalanceModelType = %v", assetsMgr.BalanceModelType())

	addrs := []string{
		"zdtjiahua222",
	}

	balances, err := sm.GetTokenBalanceByAddress(contract, addrs...)
	if err != nil {
		log.Errorf(err.Error())
		return
	}
	for _, b := range balances {
		log.Infof("balance[%s] = %s", b.Balance.Address, b.Balance.Balance)
		log.Infof("UnconfirmBalance[%s] = %s", b.Balance.Address, b.Balance.UnconfirmBalance)
		log.Infof("ConfirmBalance[%s] = %s", b.Balance.Address, b.Balance.ConfirmBalance)
	}
}


func TestGetAddressVerify(t *testing.T) {
	symbol := "ZDT"
	assetsMgr, err := openw.GetAssetsAdapter(symbol)
	if err != nil {
		log.Error(symbol, "is not support")
		return
	}
	//读取配置
	absFile := filepath.Join(configFilePath, symbol+".ini")

	c, err := config.NewConfig("ini", absFile)
	if err != nil {
		return
	}
	assetsMgr.LoadAssetsConfig(c)
	addrDec := assetsMgr.GetAddressDecoderV2()

	flag := addrDec.AddressVerify("zdttestercat")
	log.Infof("flag: %v", flag)

}