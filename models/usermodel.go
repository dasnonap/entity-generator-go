package models

import "ivan.mihov/entity-filler/generators"

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewUser() User {
	return User{}
}

func GetUserFieldTypes() map[string]FieldType {
	FieldTypes := make(map[string]FieldType, 3)
	FieldTypes["Username"] = FieldType{"string", generators.USERNAME_REGEX}
	FieldTypes["Email"] = FieldType{"email", generators.EMAIL_REGEX}
	FieldTypes["Password"] = FieldType{"password", generators.PASSWORD_REGEX}

	return FieldTypes
}

func (user *User) SetUserField(fieldName string, value string) {
	switch fieldName {
		case "Username":
			user.Username = value
		case "Email":
			user.Email = value
		case "Password":
			user.Password = value
	}
}