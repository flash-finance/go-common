package grpcclient

import (
	"github.com/flash-finance/go-common/tron/utils"
	"github.com/tronprotocol/grpc-gateway/api"
	"github.com/tronprotocol/grpc-gateway/core"
)

type WalletExt struct {
	_conn
	client api.WalletExtensionClient
}

func NewWalletExt(serverAddr string) *WalletExt {
	ret := &WalletExt{}
	ret.serverAddr = serverAddr
	return ret
}

func (g *WalletExt) Connect() (err error) {
	err = g._conn.Connect()
	if nil != err {
		return err
	}

	g.client = api.NewWalletExtensionClient(g.c)

	if nil == g.client {
		return utils.ErrorCreateGrpClient
	}

	return nil
}


func (g *WalletExt) GetTransactionsFromThis(addr string, offset, limit int64) ([]*core.Transaction, error) {

	ctx, cancel := getTimeoutContext(defaultTimeout)
	defer cancel()
	callOpt := getDefaultCallOptions()
	msg := &api.AccountPaginated{}
	account := &core.Account{}
	account.Address = utils.Base58DecodeAddr(addr)
	msg.Account = account
	msg.Offset = offset
	msg.Limit = limit

	tranList, err := g.client.GetTransactionsFromThis(ctx, msg, callOpt)

	if nil == tranList {
		return nil, err
	}

	return tranList.Transaction, err

}

// GetTransactionsFromThis2 ...
func (g *WalletExt) GetTransactionsFromThis2(addr string, offset, limit int64) ([]*api.TransactionExtention, error) {

	ctx, cancel := getTimeoutContext(defaultTimeout)
	defer cancel()
	callOpt := getDefaultCallOptions()
	msg := &api.AccountPaginated{}
	account := &core.Account{}
	account.Address = utils.Base58DecodeAddr(addr)
	msg.Account = account
	msg.Offset = offset
	msg.Limit = limit

	tranList, err := g.client.GetTransactionsFromThis2(ctx, msg, callOpt)

	if nil == tranList {
		return nil, err
	}

	return tranList.Transaction, err

}

// GetTransactionsToThis ...
func (g *WalletExt) GetTransactionsToThis(addr string, offset, limit int64) ([]*core.Transaction, error) {

	ctx, cancel := getTimeoutContext(defaultTimeout)
	defer cancel()
	callOpt := getDefaultCallOptions()
	msg := &api.AccountPaginated{}
	account := &core.Account{}
	account.Address = utils.Base58DecodeAddr(addr)
	msg.Account = account
	msg.Offset = offset
	msg.Limit = limit

	tranList, err := g.client.GetTransactionsToThis(ctx, msg, callOpt)

	if nil == tranList {
		return nil, err
	}

	return tranList.Transaction, err

}


func (g *WalletExt) GetTransactionsToThisi2(addr string, offset, limit int64) ([]*api.TransactionExtention, error) {

	ctx, cancel := getTimeoutContext(defaultTimeout)
	defer cancel()
	callOpt := getDefaultCallOptions()
	msg := &api.AccountPaginated{}
	account := &core.Account{}
	account.Address = utils.Base58DecodeAddr(addr)
	msg.Account = account
	msg.Offset = offset
	msg.Limit = limit

	tranList, err := g.client.GetTransactionsToThis2(ctx, msg, callOpt)

	if nil == tranList {
		return nil, err
	}

	return tranList.Transaction, err

}
