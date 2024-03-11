package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/ortense/consolestyle"
	"github.com/ortense/goroutine-demo/internal/api"
)

func main() {
	startTime := time.Now()

	users := make(chan *api.User)
	var wg sync.WaitGroup

	getUserAsync := func(id int) {
		defer wg.Done()
		user, err := api.GetUserWithTodos(id)
		if err == nil {
			users <- user
		}
	}

	waitComplete := func() {
		wg.Wait()
		close(users)
	}

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go getUserAsync(i)
	}

	go waitComplete()

	fmt.Println()

	for u := range users {
		fmt.Println(
			consolestyle.Green(strconv.Itoa(u.ID)),
			consolestyle.Magenta(u.Name),
			consolestyle.Italic(u.Email),
			consolestyle.Yellow(strconv.Itoa(len(u.Todos))),
		)
	}

	endTime := time.Now()
	totalTime := endTime.Sub(startTime)

	fmt.Println()
	fmt.Println(
		consolestyle.Cyan("Total time:"),
		consolestyle.Bold(consolestyle.Underline(totalTime.String())),
	)
}
