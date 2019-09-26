package redis_locker

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

type RedisLocker struct {
	RedisClient *redis.Client
}

func NewRedisLocker(redis_client *redis.Client) *RedisLocker {
	return &RedisLocker{
		RedisClient: redis_client,
	}
}

func LockKey(lock_name string) string {
	return fmt.Sprintf("lock:%s", lock_name)
}

func (locker *RedisLocker) Lock(lock_name string, time time.Duration) bool {
	result, err := locker.RedisClient.SetNX(
		LockKey(lockname), "lcok", locktime).Result()

	if err != nil {
		return false
	}

	return result
}

func (locker *RedisLocker) UnLock(lock_name string) error {
	return cache.CacheClient.Del(global.LockKey(lockname)).Err()
}
