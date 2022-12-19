package chapter12

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Status struct {
	Text string
}

type User struct {
	Status Status
}

func Twitter() {
	res, _ := http.Get("http://twitter.com/users/Googland.json")

	user := User{Status{""}}
	temp, _ := ioutil.ReadAll(res.Body)

	body := []byte(temp)
	json.Unmarshal(body, &user)
	fmt.Printf("status %s", user.Status.Text)

}
