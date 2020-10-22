package util

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/rand"
	"reflect"
	"regexp"
	"time"
)

const DATATIMEFORMAT = "2006-01-02 15:04:05"

func GetRandomString(n int) string {
	const symbols = "0123456789abcdefghjkmnopqrstuvwxyzABCDEFGHJKMNOPQRSTUVWXYZ"
	const symbolsIdxBits = 6
	const symbolsIdxMask = 1<<symbolsIdxBits - 1
	const symbolsIdxMax = 63 / symbolsIdxBits

	prng := rand.New(rand.NewSource(time.Now().UnixNano()))

	b := make([]byte, n)
	for i, cache, remain := n-1, prng.Int63(), symbolsIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = prng.Int63(), symbolsIdxMax
		}
		if idx := int(cache & symbolsIdxMask); idx < len(symbols) {
			b[i] = symbols[idx]
			i--
		}
		cache >>= symbolsIdxBits
		remain--
	}
	return string(b)
}

func MD5(text string) string {
	ctx := md5.New()
	ctx.Write([]byte(text))
	return hex.EncodeToString(ctx.Sum(nil))
}

func IsMailFormat(mail string) bool {
	var mailRe = regexp.MustCompile(`\A[\w+\-.]+@[a-z\d\-]+(\.[a-z]+)*\.[a-z]+\z`)
	return mailRe.MatchString(mail)
}

func NowStr() string {
	tm := time.Unix(time.Now().Unix(), 0)
	return tm.Format("2006-01-02 15:04:05")
}

func GenTimeStr(ts int64) string {
	return time.Unix(ts, 0).Format("2006-01-02")
}

func IsEmptyStr(needCheck string) bool {
	if needCheck != "" && len(needCheck) > 0 {
		return false
	}
	return true
}

func ToJSONStr(val interface{}) string {
	if nil == val {
		return ""
	}
	real := reflect.ValueOf(val)
	if real.IsNil() {
		return ""
	}
	if real.Kind() == reflect.Ptr && !real.Elem().IsValid() {
		return ""
	}
	if (real.Kind() == reflect.Slice || real.Kind() == reflect.Array || real.Kind() == reflect.Map) && real.IsNil() {
		fmt.Printf("list:%#v\n", real)
		return ""
	}
	data, err := json.Marshal(val)
	if nil != err {
		return fmt.Sprintf("%#v", val)
	}
	return string(data)
}
