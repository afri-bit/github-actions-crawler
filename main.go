package main

import (
	"context"
	"fmt"

	"github.com/google/go-github/v60/github"
)

func main() {
	fmt.Print("Hello Github")

	a := 5

	fmt.Print(a)

	client := github.NewClient(nil)

	orgs, _, err := client.Organizations.List(context.Background(), "willnorris", nil)

	fmt.Print(orgs, err)
}
