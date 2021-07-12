package main

import "encoding/base64"

func main() {
	user := "username"
	userBytes := []byte(user)
	encodedUser := base64.StdEncoding.EncodeToString(userBytes)
	println(encodedUser)

	pwd := "password"
	pwdBytes := []byte(pwd)
	encodedpwd := base64.StdEncoding.EncodeToString(pwdBytes)
	println(encodedpwd)
}
