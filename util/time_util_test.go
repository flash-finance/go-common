package util

import "testing"

func TestTimeStampBefore24h(t *testing.T) {
	dateTime := TimeStampBefore24h()
	t.Logf("dataTime:%v", dateTime)
}
