package main

import "fmt"

// 动物园要给不同的动物称重，并打印报告。假设你只有小猴子、小鸟、鱼、狗、猫

type Animals struct {
	Name   string
	Weight float64
}
type Cats struct {
	Animals
}
type Dogs struct {
	Animals
}
type Birds struct {
	Animals
}
type Monkeys struct {
	Animals
}

type GetWeight interface {
	getAnimalWeight() float64
}

func (m *Monkeys) getAnimalWeight() float64 { return m.Weight }
func (b *Birds) getAnimalWeight() float64   { return b.Weight }
func (d *Dogs) getAnimalWeight() float64    { return d.Weight }
func (c *Cats) getAnimalWeight() float64    { return c.Weight }

func Reporter(g GetWeight) {
	fmt.Printf("动物 %s 重量: %fkg\n", "no", g.getAnimalWeight())
}
