package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/ortense/consolestyle"
	"github.com/ortense/goroutine-demo/internal/api"
)

func main() {
	startTime := time.Now()

	fmt.Println()

	for i := 1; i <= 10; i++ {
		user, err := api.GetUserWithTodos(i)

		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println(
			consolestyle.Green(strconv.Itoa(i)),
			consolestyle.Magenta(user.Name),
			consolestyle.Italic(user.Email),
			consolestyle.Yellow(strconv.Itoa(len(user.Todos))),
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
