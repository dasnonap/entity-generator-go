package main

import (
	"encoding/json"
	"fmt"
	"regexp/syntax"

	"ivan.mihov/entity-filler/generators"
	"ivan.mihov/entity-filler/models"
)

func main() {
	usersCount := 20
	userFieldTypes := models.GetUserFieldTypes()
	var users []models.User

	for i := 0; i < usersCount; i++ {
		user := models.NewUser()
		for key, t := range userFieldTypes {
			regex, error := syntax.Parse(t.AllowedSymbols, syntax.Perl)
			
			if error != nil {
				panic(error)
			}
			value := generators.Generate(regex)
			
			user.SetUserField(key, value)
		}

		users = append(users, user)
	}

	jsonData, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)
		return
	}

	fmt.Println(string(jsonData))

	// response, error := http.Get("http://127.0.0.1:8000/api/posts")
	// if error != nil {
	// 	panic(error)
	// }

	// defer response.Body.Close()

	// var data map[string]interface {}

	// error = json.NewDecoder(response.Body).Decode(&data)
	
	// if error != nil {
	// 	panic(error)
	// }

	// fmt.Println(data["items"])

	// client := client.NewClient(client.GET, "http://127.0.0.1:8000/api/posts")

	// response, error := client.SendRequest()

	// if (error != nil) {
	// 	panic(error)
	// }

	// response.PrettyPrint()
}
