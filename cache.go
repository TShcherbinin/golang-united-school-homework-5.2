package cache

import "time"

type Cache struct {
	values map[string]cacheValue
}

type cacheValue struct {
	value   string
	expired time.Time
}

func NewCache() Cache {
	return Cache{map[string]cacheValue{}}
}

func (this Cache) Get(key string) (string, bool) {
	if v, ok := this.values[key]; ok {
		if v.expired.IsZero() || v.expired.After(time.Now()) {
			return v.value, true
		}
	}
	return "", false
}

func (this Cache) Put(key, value string) {
	this.values[key] = cacheValue{value, time.Time{}}
}

func (this Cache) Keys() []string {
	result := []string{}

	now := time.Now()
	for k, v := range this.values {
		if v.expired.IsZero() || v.expired.After(now) {
			result = append(result, k)
		}
	}
	return result
}

func (this Cache) PutTill(key, value string, deadline time.Time) {
	this.values[key] = cacheValue{value, deadline}
}
