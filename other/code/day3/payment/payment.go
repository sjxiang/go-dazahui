package main

import (
	"fmt"
	"sync"
	"time"
)


func main() {
	p := Payment{
		From: "支付宝 Jie",
		To: "KFC",
		Amount: 50.00000,
	}

	p.Process()
	
	p.Process()
}


func (p *Payment) Process() {
	t := time.Now()
	p.once.Do(func() {
		p.process(t)
	})
}

func (p *Payment) process(t time.Time) {
	ts := t.Format(time.RFC3339)
	fmt.Printf("[%s] %s -> $%.2f -> %s\n", ts, p.From, p.Amount, p.To)
}


// 支付
type Payment struct {
	From   string
	To     string
	Amount float64  // 金额

	once sync.Once
}