package main

import "example/xschema"

var zodAdapter = xschema.Adapter{Name: "zod", Brand: "xschema-adapter"}

func init() {
	xschema.FromURL("User", "https://api.example.com/user.json", zodAdapter)
	xschema.FromFile("Post", "./schemas/post.json", zodAdapter)
}
