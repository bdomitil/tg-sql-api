package main

import (
	"math/big"
)

func string_to_int64(str string) (bool, int64) {

	big, ok := new(big.Int).SetString(str, 10)
	if ok {
		return ok, big.Int64()
	} else {
		return ok, 0
	}
}
