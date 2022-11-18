package buyer

import (
	"fmt"
	"testing"
)

func TestDefaultAlgo(t *testing.T) {
	p := Parameter{Price: 10, Money: DefaultInitialMoney, SpecialMoney: DefaultSpecialMoney, NormalInterval: DefaultNormalInterval, SpecialInterval: DefaultSpecialInterval}
	r := DefaultAlgo(p)
	fmt.Println(r)
}
