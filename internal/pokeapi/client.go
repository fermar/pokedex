package pokeapi

import (
	"net/http"
	"time"

	"github.com/fermar/pokedex/internal/pokecache"
)

type Client struct {
	httpClient http.Client
	cache      pokecache.Cache
}

func NewClient(timeoutHttp, timeoutCache time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeoutHttp,
		},
		cache: pokecache.NewCache(timeoutCache),
	}
}
