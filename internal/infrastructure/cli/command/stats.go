package command

import (
	"coffee/internal/infrastructure"
	"fmt"
)

func Stats() {
	statistics := infrastructure.Statistic
	fmt.Printf("orders=%d revenue=%d\n", statistics.OrderCount, statistics.Revenue)
}
