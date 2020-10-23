package utils

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

//  是否连接 testNet 默认false
var TestNet bool

//  使用哪个网络 shasta or else
var NetName string

//  随机获取一个full node ip
func GetRandFullNodeAddr() string {
	if TestNet {
		if NetName == "shasta" {
			return TestFullNodeListShasta[rand.Int31n(int32(len(TestFullNodeListShasta)))]
		} else if NetName == "nile" {
			return TestFullNodeListNile[rand.Int31n(int32(len(TestFullNodeListShasta)))]
		}
	}
	return FullNodeList[rand.Int31n(int32(len(FullNodeList)))]
}

// 地址前缀 测试/主网
const (
	AddressPrefixTest = "41" //a0 + address, test net use the same rule now
	AddressPrefixMain = "41" //41 + address

	DefaultGrpPort = 50051
	DefaultP2pPort = 18888
)

var TestFullNodeListShasta = []string{
	"grpc.shasta.trongrid.io:50051",
}

var TestFullNodeListNile = []string{
	"47.252.19.181:50051",
}

// FullNodeList Full节点列表
var FullNodeList = []string{
	"54.236.37.243:50051",  // a // not fully implement
	"52.53.189.99:50051",   // a //  not fully implement
	"18.196.99.16:50051",   // a
	"34.253.187.192:50051", // a
	"52.56.56.149:50051",   // a
	"35.180.51.163:50051",  // a
	"54.252.224.209:50051", // a
	"18.228.15.36:50051",   // a
	"52.15.93.92:50051",    // a
	"34.220.77.106:50051",  // a
	"13.127.47.162:50051",  // a
	"13.124.62.58:50051",   // a
	// "13.229.128.108",
	//"35.182.37.246:50051", // a
	// "34.200.228.125",
	// "18.220.232.201",
	// "13.57.30.186",

	// "35.165.103.105",
	// "18.184.238.21",
	// "34.250.140.143:50051", // b
	// "35.176.192.130:50051", // b
	// "52.47.197.188:50051", // b
	// "52.62.210.100:50051", //b
	// "13.231.4.243:50051",  // b
	// "18.231.76.29",
	// "35.154.90.144:50051",  // b
	// "13.125.210.234:50051", // b
	// "13.250.40.82",
	// "35.183.101.48",
	// "47.104.11.194", // grpc connection failed
}
