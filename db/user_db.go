// db/user_db.go

package db

import (
	"fmt"
	"log"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

var validate *validator.Validate
var translator ut.Translator

type User struct {
	ID       int    `json:"ID"`
	Username string `json:"username" validate:"required,alphanum,min=3,max=25"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=10"`
}

func CreateUser(user User) error {
	_, err := db.Exec("INSERT INTO users (username, email, password) VALUES ($1, $2, $3)", user.Username, user.Email, user.Password)
	if err != nil {
		log.Printf("Error creating new user: %v", err)
		return err
	} else {
		fmt.Println("User added successfully")
	}
	return nil
}

func SetupValidator() {
	validate = validator.New()
	eng := en.New()
	uni := ut.New(eng, eng)

	translator, _ = uni.GetTranslator("en")
	en_translations.RegisterDefaultTranslations(validate, translator)

	// Customizing error messages
	validate.RegisterTranslation("required", translator, func(ut ut.Translator) error {
		return ut.Add("required", "{0} is required", true) // See the placeholder {0}
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required", fe.Field())
		return t
	})

	validate.RegisterTranslation("email", translator, func(ut ut.Translator) error {
		return ut.Add("email", "{0} must be a valid email address", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("email", fe.Field())
		return t
	})

	// Add more custom messages as needed
}

// ValidateUser uses the `validator` package to check if the User struct satisfies the constraints
func ValidateUser(user User) map[string]string {
	err := validate.Struct(user)
	if err != nil {
		errs := make(map[string]string)
		for _, e := range err.(validator.ValidationErrors) {
			errs[e.Field()] = e.Translate(translator)
		}
		return errs
	}
	return nil
}
