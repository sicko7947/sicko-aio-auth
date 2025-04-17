package redis

import (
	"errors"
	"fmt"
	"time"

	"github.com/garyburd/redigo/redis"
)

// GetCookie2FromRedis : Get Single bm_sz Cookie From Redis
func GetCookieFromRedis() string {
	con := pool.Get()
	defer con.Close()

	cookieList, err := redis.Strings(con.Do("ZRANGE", "cookie2", "0", "0"))
	if len(cookieList) > 0 {
		res, _ := redis.Int64(con.Do("ZREM", "cookie", cookieList[0]))
		if res > 0 {
			return cookieList[0]
		}
		return ""
	}
	if err != nil {
		fmt.Println(err)
	}
	return ""
}

// SaveCookie2ToRedis : Save _abck To Redis
func SaveCookieToRedis(expiry float64, cookie string) error {
	con := pool.Get()
	defer con.Close()

	expiryStr := fmt.Sprintf("%f", expiry)
	_, err := redis.Int64(con.Do("ZADD", "cookie", expiryStr, cookie))
	if err != nil {
		return err
	}
	return nil
}

// CheckExpireCookieInRedis : Check Expire Cookie In Redis
func CheckExpireCookieInRedis() {
	con := pool.Get()
	defer con.Close()

	timestamp := fmt.Sprintf(`%d`, time.Now().Unix())
	redis.Int64(con.Do("ZREMRANGEBYSCORE", "cookie2", "0", timestamp))
}

// CheckKeyExist : Check Key Exist
func CheckKeyExist(key string) error {
	con := pool.Get()
	defer con.Close()

	res, err := redis.Int64(con.Do("EXISTS", key))
	if err != nil {
		return err
	}
	if res == 1 {
		return nil
	}
	return errors.New("Key not found")
}
