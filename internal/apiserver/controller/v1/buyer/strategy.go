package buyer

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/marmotedu/iam/pkg/log"

	"github.com/ggchangan/potato/internal/pkg/util/core"
)

// Strategy get a strategy to buy a new stock.
func (r *BuyerController) Strategy(c *gin.Context) {
	log.L(c).Info("get buyer strategy function called.")
	//price := c.GetFloat64("price")

	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	user, err := r.srv.Stocks().Get(c, uint64(id))
	if err != nil {
		core.WriteResponse(c, err, nil)

		return
	}

	core.WriteResponse(c, nil, user)
}

type buyerAlgo interface {
	Calculate(p Parameter) Result
}

const DefaultInitialMoney = 100000.00
const DefaultNormalInterval = 0.1
const DefaultSpecialInterval = 0.15
const DefaultSpecialMoney = 50000.00

type Parameter struct {
	//初始买入价格
	Price float32
	//初始投入资金
	Money float32
	//前3次买入之后补仓金额
	SpecialMoney    float32
	NormalInterval  float32
	SpecialInterval float32
}

func NewParameter(money, normalInterval, specialInterval float32) Parameter {
	return Parameter{
		Money:           money,
		NormalInterval:  normalInterval,
		SpecialInterval: specialInterval,
	}
}

type Result struct {
	Price          []float32
	Money          []float32
	Loss           []float32
	LossPercentage []float32
}

// 343
type Default struct {
	Percentage []float32
}

// 343
var DefaultAlgo = func(p Parameter) Result {

	price := make([]float32, 4, 4)
	price[0] = p.Price
	for i := 1; i < 3; i++ {
		price[i] = price[i-1] * (1 - p.NormalInterval)
	}
	price[3] = price[2] * (1 - p.SpecialInterval)

	percentage := []float32{0.3, 0.4, 0.3}
	increment := make([]float32, 4, 4)
	for i := 0; i < 3; i++ {
		increment[i] = percentage[i] * p.Money
	}
	increment[3] = p.SpecialMoney

	money := make([]float32, 4, 4)
	money[0] = increment[0]
	for i := 1; i < 4; i++ {
		money[i] = money[i-1] + increment[i]
	}

	interval := []float32{p.NormalInterval, p.NormalInterval, p.SpecialInterval}
	//亏损额
	loss := make([]float32, 4, 4)
	//新一次的亏损 = 之前的亏损 + (已经投入的钱-之前的亏损) * interval
	loss[0] = 0
	for i := 1; i < 4; i++ {
		loss[i] = loss[i-1] + (money[i-1]-loss[i-1])*interval[i-1]
	}

	//亏损比率
	lossPercentage := make([]float32, 4, 4)
	//亏损的钱，投入的钱
	for i := 0; i < 4; i++ {
		lossPercentage[i] = loss[i] / money[i]
	}

	return Result{
		Price:          price,
		Money:          money,
		Loss:           loss,
		LossPercentage: lossPercentage,
	}
}

type Pyramid struct {
	Percentage []float32
}
