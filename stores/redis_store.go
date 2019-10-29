//============================================================
// 描述:
// 作者: Simon
// 日期: 2019/10/29 4:18 下午
//
//============================================================

package stores

import (
	"encoding/json"
	"time"

	"github.com/gomodule/redigo/redis"
)

//RedisStorage redis 操作类
type RedisStore struct {
	*redis.Pool
}

//newRedisStorage 实例化redis
func NewRedisStore(addr string, db, idle, active int) *RedisStorage {
	client := &redis.Pool{
		Dial: func() (conn redis.Conn, e error) {
			c, err := redis.Dial("tcp", addr)

			if err != nil {
				panic(err)
			}

			_, err = c.Do("SELECT", db)

			if err != nil {
				panic(err)
			}

			return c, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			if err != nil {
				println(err)
			}
			return err
		},
		MaxIdle:     idle,
		MaxActive:   active,
		IdleTimeout: 300 * time.Second,
	}

	return &RedisStore{client}
}

func (r *RedisStore) getConn() redis.Conn {
	return r.Get()
}

//SetValue 设置值
func (r *RedisStore) SetValue(key string, value interface{}) error {
	c := r.getConn()

	defer c.Close()

	v, _ := json.Marshal(value)

	_, err := c.Do("SET", key, v)

	if err != nil {
		return err
	}

	return nil
}

//GetValue 获取值
func (r *RedisStore) GetValue(key string, value interface{}) error {

	c := r.getConn()

	defer c.Close()

	v, err := redis.Bytes(c.Do("GET", key))

	if err != nil {
		return err
	}

	return json.Unmarshal(v, &value)
}

//Expire 设置过期时间
func (r *RedisStore) Expire(key string, t time.Duration) error {

	c := r.getConn()

	defer c.Close()

	_, err := c.Do("EXPIRE", key, t.Seconds())

	return err
}

//Del 删除Key
func (r *RedisStore) Del(key string) error {

	c := r.getConn()

	defer c.Close()

	_, err := c.Do("DEL", key)

	if err != nil {
		return err
	}

	return nil
}

func (r *RedisStore) SetValueWithTimeout(key string, value interface{}, t time.Duration) error {
	c := r.getConn()
	defer c.Close()
	v, _ := json.Marshal(value)

	_, err := c.Do("SET", key, v)

	if err != nil {
		return err
	}
	_, err = c.Do("EXPIRE", key, t.Seconds())

	return err

}
