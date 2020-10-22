package common

import (
	"github.com/flash-finance/go-common/tron/grpcclient"
	"github.com/tronprotocol/grpc-gateway/api"
	"github.com/tronprotocol/grpc-gateway/core"
	"sync"
)

var _wallet *WalletClient
var _walletOnce sync.Once

type WalletClient struct {
	*grpcclient.Wallet
	sync.Mutex
}

func GetWalletClient() *WalletClient {
	_walletOnce.Do(func() {
		_wallet = new(WalletClient)
		_wallet.Wallet = grpcclient.GetRandomWallet()
	})
	return _wallet
}

func (wc *WalletClient) Refresh() {
	wc.Lock()
	if nil != wc.Wallet {
		wc.Wallet.Close()
	}
	wc.Wallet = grpcclient.GetRandomWallet()
	wc.Unlock()
}

func (wc *WalletClient) GetExchangeByID(exchangeID int64) (*core.Exchange, error) {
	var ret *core.Exchange
	var err error
	tryCnt := 3
	for tryCnt > 0 {
		tryCnt--
		ret, err = wc.Wallet.GetExchangeByID(exchangeID)
		if nil != err {
			wc.Refresh()
			continue
		}
		return ret, err
	}
	return ret, err
}

func (wc *WalletClient) ExchangeInject(addr string, exchangeID int64, tokenID string, quant int64) (*api.TransactionExtention, error) {
	var ret *api.TransactionExtention
	var err error
	tryCnt := 3
	for tryCnt > 0 {
		tryCnt--
		ret, err = wc.Wallet.ExchangeInject(addr, exchangeID, tokenID, quant)
		if nil != err {
			wc.Refresh()
			continue
		}
		return ret, err
	}
	return ret, err
}

func (wc *WalletClient) ExchangeWithdraw(addr string, exchangeID int64, tokenID string, quant int64) (*api.TransactionExtention, error) {
	var ret *api.TransactionExtention
	var err error
	tryCnt := 3
	for tryCnt > 0 {
		tryCnt--
		ret, err = wc.Wallet.ExchangeWithdraw(addr, exchangeID, tokenID, quant)
		if nil != err {
			wc.Refresh()
			continue
		}
		return ret, err
	}
	return ret, err
}

func (wc *WalletClient) ExchangeCreate(addr string, fromTokenID, toTokenID string, fromTokenAmount, toTokenAmount int64) (*api.TransactionExtention, error) {
	var ret *api.TransactionExtention
	var err error
	tryCnt := 3
	for tryCnt > 0 {
		tryCnt--
		ret, err = wc.Wallet.ExchangeCreate(addr, fromTokenID, toTokenID, fromTokenAmount, toTokenAmount)
		if nil != err {
			wc.Refresh()
			continue
		}
		return ret, err
	}
	return ret, err
}

func (wc *WalletClient) ExchangeList() ([]*core.Exchange, error) {
	var ret []*core.Exchange
	var err error
	tryCnt := 3
	for tryCnt > 0 {
		tryCnt--
		ret, err = wc.Wallet.ListExchanges()
		if nil != err {
			wc.Refresh()
			continue
		}
		return ret, err
	}
	return ret, err
}

func (wc *WalletClient) BroadcastTransaction(trx *core.Transaction) (*api.Return, error) {
	var ret *api.Return
	var err error
	tryCnt := 3
	for tryCnt > 0 {
		tryCnt--
		ret, err = wc.Wallet.BroadcastTransaction(trx)
		if nil != err {
			wc.Refresh()
			continue
		}
		return ret, err
	}
	return ret, err
}

// GetNowBlock ...
func (wc *WalletClient) GetNowBlock() (*core.Block, error) {
	tryCnt := 3
	for tryCnt > 0 {
		tryCnt--
		ret, err := wc.Wallet.GetNowBlock()
		if nil != err || nil == ret || nil == ret.BlockHeader || nil == ret.BlockHeader.RawData {
			wc.Refresh()
			continue
		}
		return ret, err
	}
	return nil, nil
}

func (wc *WalletClient) GetSmartContract(address string) (*core.SmartContract, error) {
	tryCnt := 3
	for tryCnt > 0 {
		tryCnt--
		ret, err := wc.Wallet.GetContract(address)
		if nil != err || nil == ret {
			wc.Refresh()
			continue
		}
		return ret, err
	}
	return nil, nil
}

func (wc *WalletClient) GetMaxBlockID() (*core.Block, error) {
	var ret *core.Block
	var err error
	tryCnt := 3
	for tryCnt > 0 {
		tryCnt--
		ret, err = wc.Wallet.GetNowBlock()
		if nil != err || nil == ret {
			wc.Refresh()
			continue
		}
		return ret, err
	}
	return ret, err
}
