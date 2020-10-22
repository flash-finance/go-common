package tools

import (
	"github.com/flash-finance/go-common/tron/utils"
	"github.com/tronprotocol/grpc-gateway/core"
	"time"
)

func GenAccountCreateContract(from, to string) (core.Transaction_Contract_ContractType, *core.AccountCreateContract, error) {
	ctx := new(core.AccountCreateContract)
	ctx.OwnerAddress = utils.Base58DecodeAddr(from)
	ctx.AccountAddress = utils.Base58DecodeAddr(to)
	return core.Transaction_Contract_AccountCreateContract, ctx, nil
}

func GenTransferContract(from, to string, amount int64) (core.Transaction_Contract_ContractType, *core.TransferContract, error) {
	ctx := new(core.TransferContract)
	ctx.OwnerAddress = utils.Base58DecodeAddr(from)
	ctx.ToAddress = utils.Base58DecodeAddr(to)
	ctx.Amount = amount
	return core.Transaction_Contract_TransferContract, ctx, nil
}

func GenTransferAssetContract(from, to, assetName string, amount int64) (core.Transaction_Contract_ContractType, *core.TransferAssetContract, error) {
	ctx := new(core.TransferAssetContract)
	ctx.OwnerAddress = utils.Base58DecodeAddr(from)
	ctx.ToAddress = utils.Base58DecodeAddr(to)
	ctx.AssetName = []byte(assetName)
	ctx.Amount = amount
	return core.Transaction_Contract_TransferAssetContract, ctx, nil
}

func GenTriggerSmartContract(from, to string, callValue int64, data []byte) (core.Transaction_Contract_ContractType, *core.TriggerSmartContract, error) {
	ctx := new(core.TriggerSmartContract)
	ctx.OwnerAddress = utils.Base58DecodeAddr(from)
	ctx.ContractAddress = utils.Base58DecodeAddr(to)
	ctx.CallValue = callValue
	ctx.Data = data
	return core.Transaction_Contract_TriggerSmartContract, ctx, nil
}

func GenFreezeBalanceContract(from string, amount int64, frozenDay int64, resourceType core.ResourceCode) (core.Transaction_Contract_ContractType, *core.FreezeBalanceContract, error) {
	ctx := new(core.FreezeBalanceContract)
	ctx.OwnerAddress = utils.Base58DecodeAddr(from)
	ctx.FrozenBalance = amount
	ctx.FrozenDuration = frozenDay
	ctx.Resource = resourceType
	return core.Transaction_Contract_FreezeBalanceContract, ctx, nil
}

func GenUnfreezeBalanceContract(from string, resourceType core.ResourceCode) (core.Transaction_Contract_ContractType, *core.UnfreezeBalanceContract, error) {
	ctx := new(core.UnfreezeBalanceContract)
	ctx.OwnerAddress = utils.Base58DecodeAddr(from)
	ctx.Resource = resourceType
	return core.Transaction_Contract_UnfreezeBalanceContract, ctx, nil
}

func GenAssetIssueContract(owner, assetName, abbrName, assetDesc, assetURL string, FrozenSupply []*core.AssetIssueContract_FrozenSupply, totalSupply int64, trxNum, num int32, startTime, endTime time.Time) (core.Transaction_Contract_ContractType, *core.AssetIssueContract, error) {
	ctx := new(core.AssetIssueContract)
	ctx.OwnerAddress = utils.Base58DecodeAddr(owner)
	ctx.Name = []byte(assetName)
	ctx.Abbr = []byte(abbrName)
	ctx.TotalSupply = totalSupply
	ctx.FrozenSupply = FrozenSupply
	ctx.TrxNum = trxNum
	ctx.Num = num
	ctx.StartTime = startTime.UnixNano() / 1000000
	ctx.EndTime = endTime.UnixNano() / 1000000
	ctx.Description = []byte(assetDesc)
	ctx.Url = []byte(assetURL)

	return core.Transaction_Contract_AssetIssueContract, ctx, nil
}

func GenExchangeCreateContract(owner, firstToken, secondToken string, firstAmount, secondAmount int64) (core.Transaction_Contract_ContractType, *core.ExchangeCreateContract, error) {
	ctx := new(core.ExchangeCreateContract)
	ctx.OwnerAddress = utils.Base58DecodeAddr(owner)
	ctx.FirstTokenId = []byte(firstToken)
	ctx.FirstTokenBalance = firstAmount
	ctx.SecondTokenId = []byte(secondToken)
	ctx.SecondTokenBalance = secondAmount
	return core.Transaction_Contract_ExchangeCreateContract, ctx, nil
}

func GenExchangeTransactionContract(owner string, exchangeID int64, tokenID string, quant, expected int64) (core.Transaction_Contract_ContractType, *core.ExchangeTransactionContract, error) {

	ctx := new(core.ExchangeTransactionContract)
	ctx.OwnerAddress = utils.Base58DecodeAddr(owner)
	ctx.ExchangeId = exchangeID
	ctx.TokenId = []byte(tokenID)
	ctx.Quant = quant
	ctx.Expected = expected
	return core.Transaction_Contract_ExchangeTransactionContract, ctx, nil
}

func GenExchangeInjectContract(owner string, exchangeID int64, tokenID string, quant int64) (core.Transaction_Contract_ContractType, *core.ExchangeInjectContract, error) {

	ctx := new(core.ExchangeInjectContract)
	ctx.OwnerAddress = utils.Base58DecodeAddr(owner)
	ctx.ExchangeId = exchangeID
	ctx.TokenId = []byte(tokenID)
	ctx.Quant = quant
	return core.Transaction_Contract_ExchangeInjectContract, ctx, nil
}

func GenExchangeWithdrawContract(owner string, exchangeID int64, tokenID string, quant int64) (core.Transaction_Contract_ContractType, *core.ExchangeWithdrawContract, error) {

	ctx := new(core.ExchangeWithdrawContract)
	ctx.OwnerAddress = utils.Base58DecodeAddr(owner)
	ctx.ExchangeId = exchangeID
	ctx.TokenId = []byte(tokenID)
	ctx.Quant = quant
	return core.Transaction_Contract_ExchangeWithdrawContract, ctx, nil
}

func GenVoteWitnessContract(owner string, witnessList []string, voteList []int64) (core.Transaction_Contract_ContractType, *core.VoteWitnessContract, error) {

	if len(witnessList) != len(voteList) || len(witnessList) == 0 {
		return core.Transaction_Contract_VoteWitnessContract, nil, nil
	}

	ctx := new(core.VoteWitnessContract)
	ctx.OwnerAddress = utils.Base58DecodeAddr(owner)
	ctx.Support = true
	for idx := range witnessList {
		ctx.Votes = append(ctx.Votes, &core.VoteWitnessContract_Vote{
			VoteAddress: utils.Base58DecodeAddr(witnessList[idx]),
			VoteCount:   voteList[idx],
		})
	}
	return core.Transaction_Contract_VoteWitnessContract, ctx, nil
}
