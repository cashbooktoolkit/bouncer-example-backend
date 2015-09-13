/*
 * user.go - Bouncer Example Back End User 'model'
 *
 * License: Public Domain
 */

package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"fmt"
	"log"
)

// Our User model
type User struct {
	Login string
	Password string
	Uid string
	Error string  // Allows the system  to simulate errors for specific users
}

type Users []User

func (users Users) find(login string) (found_user *User) {

	for _, user := range users {
		if user.Login == login {
			found_user = &user
			break;
		}
	}

	return
}

/*
 * Our 'Db' of users is loaded from a json file.
 */
func load_users(filename string) (Users) {

	file, e := ioutil.ReadFile(filename)

    if e != nil {
        fmt.Printf("File error: %v\n", e)
        os.Exit(1)
    }

	var users Users
	err := json.Unmarshal(file, &users)
	
	if err != nil {
		log.Printf("Error reading json %v", err)
		os.Exit(1)
	}

	return users
}

