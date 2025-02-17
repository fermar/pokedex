package pokeapi

import (
	"fmt"
	"testing"
	"time"

	// "github.com/fermar/pokedex/internal/pokecache"
	"github.com/fermar/pokedex/internal/pokelog"
)

func TestListLoc(t *testing.T) {
	cases := []struct {
		key *string
		val string
	}{
		{
			key: nil,
			val: "canalave-city-area",
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Probando caso %v", i), func(t *testing.T) {
			pokelog.StartPokelogger(true)
			// cache := pokecache.NewCache(5 * time.Second)
			client := NewClient(5*time.Second, 20*time.Second)
			listArea, err := client.ListLoc(c.key)
			if err != nil {
				t.Errorf("listArea returned error")
				return
			}
			found := false
			for _, l := range listArea.Results {
				if l.Name == string(c.val) {
					found = true
					break
				}
			}
			if !found {
				t.Errorf("loction %v encontrada", c.val)
				return
			}

			// 	cache.Add(c.key, c.val)
			// 	val, ok := cache.Get(c.key)
			// 	if !ok {
			// 		t.Errorf("expected to find key")
			// 		return
			// 	}
			// 	if string(val) != string(c.val) {
			// 		t.Errorf("expected to find value")
			// 		return
			// 	}
		})
	}

}
