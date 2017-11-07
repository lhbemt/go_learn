package memoryCacheChan

import "testing"
import "net/http"
import "io/ioutil"

var cache *MemoryCacheChan

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

func TestMemoryCacheChan_Get(t *testing.T) {
	cache.Get("http://www.baodu.com")
}

func TestMemoryCacheChan_Close(t *testing.T) {
	cache.Close()
}
