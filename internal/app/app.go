package app

import (
	"2/internal/a51"
	"fmt"
)

const (
	//key = uint64(11111111111111111111)
	key = uint64(1234567891011121314)
)

func Run() {
	a5 := a51.New()

	err := a5.InitRegs(a51.Method1, key)
	if err != nil {
		panic(err)
	}

	k := a5.GenerateKeyStream(5)
	_ = k
	for i := 0; i < len(k); i++ {
		fmt.Printf("i: %d\n", i)
		fmt.Printf("%0X\n", k[i])
	}
}
