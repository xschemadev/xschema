package main

import (
	"fmt"

	"example/xschema"
)

// Mock adapter (normally from github.com/xschema/adapter-gojsonschema)
var gojsonschemaAdapter = xschema.Adapter{
	Name:  "gojsonschema",
	Brand: "xschema-adapter",
}

func main() {
	// From URL
	xschema.FromURL("User", "https://cdn.my/user.json", gojsonschemaAdapter)

	user := xschema.User{
		ID:    "123e4567-e89b-12d3-a456-426614174000",
		Name:  "Alice",
		Email: "alice@example.com",
	}
	if err := user.Validate(); err != nil {
		panic(err)
	}
	fmt.Printf("User: %+v\n", user)

	// From file
	xschema.FromFile("Post", "./schemas/post.json", gojsonschemaAdapter)

	post := xschema.Post{
		ID:        "123e4567-e89b-12d3-a456-426614174000",
		Title:     "Hello World",
		Body:      "This is my first post",
		AuthorID:  "123e4567-e89b-12d3-a456-426614174000",
		Published: true,
	}
	if err := post.Validate(); err != nil {
		panic(err)
	}
	fmt.Printf("Post: %+v\n", post)
}
