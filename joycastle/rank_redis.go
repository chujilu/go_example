package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v9"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"
	"time"
)

const splitNum = 10000000

type RankRedis struct {
	ctx     context.Context
	rdb     *redis.Client
	key     string
	endTime time.Time
	maxNum  int
}

func NewRankRedis() *RankRedis {
	return &RankRedis{}
}
func (r *RankRedis) Init(ctx context.Context, key string, num int, endTime time.Time) error {
	r.rdb = redis.NewClient(&redis.Options{
		Addr:     "j.it603.com:49184",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	r.ctx = ctx
	r.key = key
	r.maxNum = num
	r.endTime = endTime
	return nil
}

func (r RankRedis) encodeScore(score uint16, finishTime time.Time) float64 {
	return float64(int64(score)*splitNum + (r.endTime.Unix()-finishTime.Unix())%splitNum)
}

func (r RankRedis) decodeScore(score float64) (uint16, time.Time) {
	return uint16(int64(score) / splitNum), time.Unix(r.endTime.Unix()-int64(score)%splitNum, 0)
}

func (r *RankRedis) Add(id uint32, score uint16, finishTime time.Time) error {
	cmd := r.rdb.ZAdd(r.ctx, r.key, redis.Z{Score: r.encodeScore(score, finishTime), Member: strconv.Itoa(int(id))})
	if cmd.Err() != nil {
		return cmd.Err()
	}
	return nil
}

func (r RankRedis) GetByID(id uint32) ([]Member, error) {
	cmd := r.rdb.ZRevRank(r.ctx, r.key, strconv.Itoa(int(id)))
	if cmd.Err() != nil {
		return nil, fmt.Errorf("redis zrevrank err:%s", cmd.Err())
	}
	v := cmd.Val()
	fmt.Println(v)
	start := v - 10
	end := v + 10
	if start < 0 {
		start = 0
	}
	zcmd := r.rdb.ZRevRangeWithScores(r.ctx, r.key, start, end)
	if cmd.Err() != nil {
		return nil, cmd.Err()
	}
	result := make([]Member, 0, 20)
	for i, z := range zcmd.Val() {
		id, err := strconv.ParseInt(z.Member.(string), 10, 32)
		if err != nil {
			continue
		}
		score, finish := r.decodeScore(z.Score)
		result = append(result, Member{
			ID:         uint32(id),
			Score:      score,
			Rank:       int32(start + int64(i)),
			FinishTime: finish,
		})
	}
	return result, nil
}

func main() {
	endTime, err := time.Parse(time.RFC3339, "2022-09-01T00:00:00+08:00")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	rank := NewRankRedis()
	err = rank.Init(context.Background(), "MR:202208", math.MaxInt32, endTime)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	http.HandleFunc("/rank", func(writer http.ResponseWriter, request *http.Request) {
		idStr := request.URL.Query().Get("id")
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			writer.Write([]byte(fmt.Sprintf("parse err: %s ", err.Error())))
		}
		r, err := rank.GetByID(uint32(id))
		if err != nil {
			writer.Write([]byte(fmt.Sprintf("get err: %s ", err.Error())))
		}
		res, err := json.Marshal(r)
		if err != nil {
			writer.Write([]byte(fmt.Sprintf("marshl err: %s ", err.Error())))
		}
		writer.Write(res)
	})
	err = http.ListenAndServe(":9110", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	fmt.Println("http://127.0.0.1:9110/rank?id=1")
}
