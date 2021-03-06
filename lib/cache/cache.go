package cache

import (
	"fmt"
	"os"
	"time"

	"github.com/go-redis/redis"
	"timekeeper/tk"
)

var client *redis.Client

// Map accounts to all cache keys they are using.
// When an account is updated, these cached values will be invalidated.
var accKeys map[string][]string

func init() {
	addr := fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"),
		os.Getenv("REDIS_PORT"))
	client = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "",
		DB:       0,
	})

	fmt.Println("redis: sending PING")
	pong, err := client.Ping().Result()
	if err != nil {
		fmt.Fprintf(os.Stderr, "redis: %s\n", err.Error())
		panic("Could not initialize redis connection")
	}

	fmt.Printf("redis: received %s\n", pong)

	accKeys = make(map[string][]string)
}

const DefaultExpiration = 5 * time.Minute

// Write a value to the cache with the default expiration period.
func Set(key string, val string, acc *tk.Account) {
	if acc != nil {
		accKeys[acc.Username] = append(accKeys[acc.Username], key)
	}
	client.Set(key, val, DefaultExpiration)
}

// Write a value to the cache with a specified expiration period.
func SetExpiration(key string, val string, exp time.Duration, acc *tk.Account) {
	if acc != nil {
		accKeys[acc.Username] = append(accKeys[acc.Username], key)
	}
	client.Set(key, val, exp)
}

// Read the value for the specified key from the cache.
func Get(key string) (string, error) {
	return client.Get(key).Result()
}

// Invalidate all cached keys for the specified account.
func InvalidateAccount(acc *tk.Account) {
	if acc == nil {
		return
	}

	keys, prs := accKeys[acc.Username]
	if !prs {
		return
	}

	for _, key := range keys {
		client.Del(key)
	}
	delete(accKeys, acc.Username)
}
