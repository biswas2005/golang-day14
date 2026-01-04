package set1

import (
	"fmt"
	"os"
	"strings"
)

type Cache interface {
	Set(key string, value string)
	Get(key string) (string, bool)
}

type InMemoryCache struct {
	store map[string]string
}

func NewInMemoryCache() *InMemoryCache {
	return &InMemoryCache{store: make(map[string]string)}
}
func (c *InMemoryCache) Set(key string, value string) {
	c.store[key] = value
}
func (c *InMemoryCache) Get(key string) (string, bool) {
	val, ok := c.store[key]
	return val, ok
}

type FileCache struct {
	FileName string
}

func NewFileCache(Filename string) *FileCache {
	_, err := os.Stat(Filename)
	if os.IsNotExist(err) {
		os.WriteFile(Filename, []byte(""), 0644)
	}
	return &FileCache{FileName: Filename}
}
func (f *FileCache) Set(key string, value string) {
	entry := fmt.Sprintf("%s=%s\n", key, value)
	file, _ := os.OpenFile(f.FileName, os.O_APPEND|os.O_WRONLY, 0644)
	defer file.Close()
	file.WriteString(entry)
}
func (f *FileCache) Get(key string) (string, bool) {
	data, err := os.ReadFile(f.FileName)
	if err != nil {
		return " ", false
	}
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, key+"=") {
			return strings.TrimPrefix(line, key+"="), true
		}
	}
	return " ", false
}

func CacheSystem() {
	memCache := NewInMemoryCache()
	memCache.Set("foo", "bar")
	val, ok := memCache.Get("foo")
	fmt.Println("In Memory Cache:", val, ok)

	fileCache := NewFileCache("cache.txt")
	fileCache.Set("hello", "world")
	val2, ok2 := fileCache.Get("hello")
	fmt.Println("File cache:", val2, ok2)
}
