package main

import (
	"math/rand"
	"time"
)

var users = []User{
	User{
		ID:   1,
		Name: "Konrad Adenauer",
	},
	User{
		ID:   2,
		Name: "Ludwig Erhard",
	},
	User{
		ID:   3,
		Name: "Kurt Georg Kiesinger",
	},
	User{
		ID:   4,
		Name: "Willy Brandt",
	},
	User{
		ID:   5,
		Name: "Helmut Schmidt",
	},
	User{
		ID:   6,
		Name: "Helmut Kohl",
	},
	User{
		ID:   7,
		Name: "Gerhard Schröder",
	},
	User{
		ID:   8,
		Name: "Angela Merkel",
	},
}

func queryDB(id int) (User, bool) {
	time.Sleep(50 * time.Millisecond)
	for _, u := range users {
		if u.ID == id {
			return u, true
		}
	}

	return User{}, false
}

func randomID() int {
	return rand.New(rand.NewSource(time.Now().UnixNano())).Intn(len(users)) + 1
}
