package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

const baseurl = "https://jsonplaceholder.typicode.com"

func GetTodoByUserID(id int) ([]Todo, error) {
	url := baseurl + "/todos?userId=" + strconv.Itoa(id)

	resp, err := http.Get(url)

	if err != nil {
		return []Todo{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		reqErr := errors.New("erro ao obter a todo's - status: " + strconv.Itoa(resp.StatusCode))
		return []Todo{}, reqErr
	}

	var todo []Todo
	err = json.NewDecoder(resp.Body).Decode(&todo)

	if err != nil {
		return []Todo{}, err
	}

	return todo, nil
}

func GetUserByID(id int) (User, error) {

	url := baseurl + "/users/" + strconv.Itoa(id)

	resp, err := http.Get(url)

	if err != nil {
		return User{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		reqErr := errors.New("erro ao obter o usuario - status: " + strconv.Itoa(resp.StatusCode))
		return User{}, reqErr
	}

	var user User
	err = json.NewDecoder(resp.Body).Decode(&user)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func GetUserWithTodos(id int) (*User, error) {
	userChan := make(chan User)
	todosChan := make(chan []Todo)
	errChan := make(chan error)

	go func() {
		user, err := GetUserByID(id)
		if err != nil {
			errChan <- err
			return
		}
		userChan <- user
	}()

	go func() {
		todos, err := GetTodoByUserID(id)
		if err != nil {
			errChan <- err
			return
		}
		todosChan <- todos
	}()

	select {
	case err := <-errChan:
		return nil, err
	default:
	}

	user := <-userChan
	todos := <-todosChan

	user.Todos = todos
	return &user, nil
}
