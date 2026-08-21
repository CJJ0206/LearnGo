package main

import "slices"

type Config struct {
	item []string
}

func NewConfig(item []string) *Config {
	return &Config{item: slices.Clone(item)}
}
func (c *Config) Items() []string {
	return slices.Clone(c.item)
}
