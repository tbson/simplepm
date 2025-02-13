package main

import (
	"fmt"
	"src/module/git/repo/github"
)

func main() {
	repo := github.New()
	result, err := repo.GenerateJWT()
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
