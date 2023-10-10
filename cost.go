package ares

import (
	"context"
	"time"
)

// CostByName 计算耗时.
func CostByName(ctx context.Context, name string) func() {
	start := time.Now()
	return func() {
		Info("name:%s func_cost:%s", name, time.Since(start))
	}
}

// Cost 计算耗时.
func Cost(ctx context.Context) func() {
	start := time.Now()
	return func() {
		Info("func_cost:%s", time.Since(start))
	}
}
