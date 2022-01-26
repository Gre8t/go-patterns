package main

import (
	"math/rand"
	"testing"

	"golang.org/x/tools/go/analysis/passes/nilfunc"
)

func Test_UserListProxy(t *testing.T) {
	mockedDatabase := UserList{}

	rand.Seed(2342342)
	for i := 0; i < 1000000; i++ {
		n := rand.Int31()
		mockedDatabase = append(mockedDatabase, User{ID: n})
	}
	proxy := UserListProxy{
		MockedDatabase: &mockedDatabase,
		Stacksize:      2,
		StackCache:     UserList{},
	}
	knownIDs := [3]int32{mockedDatabase[3].ID, mockedDatabase[4].ID, mockedDatabase[5].ID}
	t.Run("FindUser - Empty cache", func(t *testing.T) {
		user, err := proxy.FindUser(knownIDs[0])
		if err != nil {
			t.Fatal(err)
		}
		if user.ID != knownIDs[0] {
			t.Error("Returned user name doesn't match with expected")
			if len(proxy.StackCache) != 1 {
				t.Error("After one successful search in an empty cache, the size of it must be one")
			}
			if proxy.LastSearchUsedCache {
				t.Error("No user can be returned from an empty cache")
			}
		}
	})
	t.Run("FindUser - One user, ask for the same user", func(t *testing.T){
		user, err := proxy.Finduser(knownIDs[0])
		if err != nil {
			t.Fatal(err)
		}
		if user.ID != knownIDs[0]{
			t.Error("Returned user name doesn't match with expected")
		}
		if len(proxy.StackCache) != 1 {
			t.Error("Cache must not grow if we ask for an object that is stored on it")
		}
		if !proxy.LastSearchUsedCache {
			t.Error("The user should have been returned from the cache")
		}
	}
	
	
}
