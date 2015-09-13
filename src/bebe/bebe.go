/*
 * bebe.go - Bouncer Example Back End.
 *
 * An example of an implementation that can be used by Bouncer to do the actual authentication.
 * In a real implmentation you would write code to interface to your existing legacy system.
 *
 * For Cashbook, the backend will most likely be the host banking system or the online banking
 * front end for it.
 *
 * WARNING:  It should go without saying but while this is useful during development, it 
 *           should not be used for production.
 *
 * License: Public Domain
 */

package main

import (
	"fmt"
	"flag"
	"net/http"
)

func main() {

	port := flag.String("port", "14001", "Port number to listen on")
	users_file := flag.String("users", "users.json", "Json file to load user records from")

    flag.Parse()

	users := load_users(*users_file)

	// Authorization code endpoint
	http.HandleFunc("/authenticate", func(w http.ResponseWriter, r *http.Request) {

		r.ParseForm()

		if r.Method == "POST"  {

			login := r.Form.Get("login")

			if login != "" {

				password := r.Form.Get("password")

				if password != "" {

					user := users.find(login)

					if (user != nil) && (user.Password == password) {						
						w.Write([]byte(fmt.Sprintf("{\"Success\": true, \"Uid\": \"%s\", \"Error\": \"%s\"}", user.Uid, user.Error)))
					} else {
						w.Write([]byte("{\"Error\": \"There was a problem with that login or password.\"}"))
					}
				} else {
					w.Write([]byte("{\"Error\": \"Missing or blank password\"}"))
				}
			} else {
				w.Write([]byte("{\"Error\": \"Missing or blank login\"}"))
			}
		} else {
			w.Write([]byte("{\"Error\": \"POST only please\"}"))
		}
		
	})

	http.ListenAndServe(":" + *port, nil)
}
