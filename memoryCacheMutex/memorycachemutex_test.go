package memoryCacheMutex

import "testing"
import "net/http"
import "io/ioutil"

var cache *Memorycache

func GetUrl(key string)(interface{}, error) {
	resp, err := http.Get(key)
	if err != nil {
		return nil, err
	} else {
		body, err := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()
		return body, err
	}
}

func TestNew(t *testing.T) {
	cache = New(GetUrl)
}

func TestMemorycache_Get(t *testing.T) {
	cache.Get("http://www.baidu.com")
}

func TestSpped(t *testing.T) {
	cache.Get("http://www.baidu.com")
}
