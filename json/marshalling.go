package json

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/oapi-codegen/nullable"
)

type User struct {
	FirstName       nullable.Nullable[string] `json:"first_name"`
	LastName        nullable.Nullable[string] `json:"last_name"`
	Email           nullable.Nullable[string] `json:"email"`
	IsEmailVerified nullable.Nullable[bool]   `json:"is_email_verified"`
	Phone           nullable.Nullable[string] `json:"phone"`
}

func UseMarshalling() {

	var userBefore User

	userBefore.FirstName.Set("Dima")
	userBefore.LastName.Set("Dimin")
	userBefore.Email.Set("example@mail.com")
	userBefore.IsEmailVerified.Set(true)
	userBefore.Phone.SetNull()

	res, err := json.Marshal(userBefore)
	if err != nil {
		log.Println(err)

	}

	fmt.Println("\n", string(res))

	var userAfterMapped User

	err = json.Unmarshal(res, &userAfterMapped)

	if err != nil {
		log.Println(err)
	}

	firstName, err := userAfterMapped.FirstName.Get()
	if err != nil {
		log.Println(err)
	}

	fmt.Println(firstName)

}

func UseUnmarshalling() {

	incData := `{
    "first_name":"Dima",
    "last_name":"Dimin",
    "email":"example@mail.com",
    "is_email_verified":true,
    }`

	var user User

	err := json.Unmarshal([]byte(incData), &user)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(user.Phone.IsSpecified())

	fmt.Printf("%+v\n", user)

}
