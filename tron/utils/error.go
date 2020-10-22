package utils

import "fmt"

var (
	ErrorDecrypt = fmt.Errorf("decrypt failed")

	ErrorInvalidSign = fmt.Errorf("invalid transaction signature, should be 65 length bytes")

	ErrorCreateGrpClient = fmt.Errorf("create gprc client failed")

	ErrorNotImplement = fmt.Errorf("not implement")
)
