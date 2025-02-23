package main

import (
	"fmt"
	"src/module/git/usecase/github/infra"
	"src/util/dbutil"
)

func main() {
	dbutil.InitDb()
	repo := infra.New(dbutil.Db(nil))
	gitRepo := "nghiencode/integrate-simplepm"
	gitBranch := "son/test-branch1"
	result, err := repo.GetTaskUser(gitRepo, gitBranch)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result)
}
