package main

import (
	"fmt"
	"src/module/git/repo/github"
	"src/util/localeutil"
)

func main() {
	localeutil.Init("en")
	installationIDs := []string{"61038391"}
	repo := github.New()
	result, err := repo.GetRepoList(installationIDs)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
