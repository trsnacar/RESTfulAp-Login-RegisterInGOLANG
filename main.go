package main

import (
	"fmt"
	"net/http"

	helper "./helpers"
)

func main() {

	uName, email, pwd, pwdConfirm := "", "", "", ""
	mux := http.NewServeMux()

	//Singup
	mux.HandleFunc("/singup", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		uName = r.FormValue("username")
		email = r.FormValue("email")
		pwd = r.FormValue("password")
		pwdConfirm = r.FormValue("confirm")

		uNameCheck := helper.IsEmpty(uName)
		emailCheck := helper.IsEmpty(email)
		pwdCheck := helper.IsEmpty(pwd)
		pwdConfirmCheck := helper.IsEmpty(pwdConfirm)

		if uNameCheck || emailCheck || pwdCheck || pwdConfirmCheck {
			fmt.Fprintf(w, "There is empty data.")
			return
		}

		if pwd == pwdConfirm {
			// Vertabanına kullanıcıyı kayıt et.
			fmt.Fprintf(w, "Registration succesfull.")
		} else {
			fmt.Fprintf(w, "Password information must be the same.")
		}

	})

	//Login
	mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		email = r.FormValue("email")
		pwd = r.FormValue("password")

		emailCheck := helper.IsEmpty(email)
		pwdCheck := helper.IsEmpty(pwd)

		if emailCheck || pwdCheck {
			fmt.Fprintf(w, "There is empty data.")
			return
		}

		dbPwd := "12345!*."
		dbEmail := "trsn.acr@gmail.com"

		if email == dbEmail && pwd == dbPwd {
			fmt.Fprintf(w, "Login Succesfull.")
		} else {
			fmt.Fprintf(w, "Login failed")
		}
	})

	http.ListenAndServe(":8080", mux)
}

/*
		for key, value := range r.Form {
				fmt.Printf("%s = %s\n", key, value)
	}
*/
