package main

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"os"
	"testing"
	"time"
)

const testKey = "MR:202208"

var testEndTime time.Time

func TestMain(t *testing.M) {
	var err error
	testEndTime, err = time.Parse(time.RFC3339, "2022-09-01T00:00:00+08:00")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	fmt.Println(testEndTime)
	t.Run()
}

func TestRankRedis_Init(t *testing.T) {
	r := NewRankRedis()
	err := r.Init(context.Background(), testKey, 1000, testEndTime)
	assert.Nil(t, err)
}

func TestRankRedis_Add(t *testing.T) {
	r := NewRankRedis()
	err := r.Init(context.Background(), testKey, 1000, testEndTime)
	assert.Nil(t, err)
	err = r.Add(1, 18, time.Now())
	assert.Nil(t, err)
}

func TestRankRedis_GetByID(t *testing.T) {
	r := NewRankRedis()
	err := r.Init(context.Background(), testKey, 1000, testEndTime)
	assert.Nil(t, err)
	res, err := r.GetByID(1)
	assert.Nil(t, err)
	assert.NotNil(t, res)
	t.Log(res)
}

func BenchmarkRankRedis_Add(b *testing.B) {
	r := NewRankRedis()
	err := r.Init(context.Background(), testKey, 1000, testEndTime)
	assert.Nil(b, err)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < b.N; i++ {
		err = r.Add(uint32(rand.Int()%50000000), uint16(rand.Int()%10000), time.Now())
		assert.Nil(b, err)
	}
}
