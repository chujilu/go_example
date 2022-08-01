package main

import (
	"context"
	"time"
)

// Member 玩家
type Member struct {
	ID         uint32
	Score      uint16
	Rank       int32
	FinishTime time.Time
}

type Rank interface {
	Init(ctx context.Context, key string, num int, endTime time.Time) error
	Add(id uint32, score uint16, finishTime time.Time) error
	GetByID(id uint32) ([]Member, error)
}
