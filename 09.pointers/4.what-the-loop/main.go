package main

import "fmt"

type User struct {
	ID   int
	Name string
}

func main() {
	users := []User{{1, "Andy"}, {2, "John"}, {3, "Jane"}}
	// Fix the get user id function.
	userIDs := getUserIDs(users)
	for _, id := range userIDs {
		fmt.Println(id) // check the memory address!
		fmt.Println(*id)
	}
}

// Run the test to make sure you've fixed the problem.
func getUserIDs(users []User) []*int {
	ids := []*int{}
	// Go bug: loop variables are per loop and not per iteration: https://github.com/golang/go/discussions/56010
	for _, user := range users {
		ids = append(ids, &user.ID)
	}
	return ids
}
