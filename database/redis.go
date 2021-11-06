package database

import (
	"context"
	"strconv"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()
var rdb *redis.Client
var cur_id int

func RedisInitClient() {
	//初始化客户端
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	tmp_id, err := rdb.Get(ctx, "cur_id").Result()
	//读取当前id
	if err == redis.Nil {
		cur_id = 1
		err = rdb.Set(ctx, "cur_id", "1", 0).Err()
		if err != nil {
			panic(err)
		}
	} else if err != nil {
		panic(err)
	} else {
		cur_id, _ = strconv.Atoi(tmp_id)
	}
}

func RedisGetRuleAttr(ruleid string, attrcode string) (string, error) {
	val, err := rdb.HGet(ctx, ruleid, attrcode).Result()
	return val, err

}

func RedisCheckWhiteList(ruleid string, userid string) (bool, error) {
	val, err := rdb.SIsMember(ctx, ruleid, userid).Result()
	return val, err
}

func RedisAddRule(r map[string]string, white_list []string) error {
	err := rdb.HMSet(ctx, strconv.Itoa(cur_id), r).Err()
	if err != nil {
		return err
	}
	err = rdb.SAdd(ctx, strconv.Itoa(cur_id)+"set", white_list).Err()
	if err != nil {
		return err
	}
	cur_id++
	rdb.Incr(ctx, "cur_id")
	return err
}
