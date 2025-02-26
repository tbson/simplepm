package app

import (
	"fmt"
	"src/common/ctype"
	"src/module/account/schema"
	"src/module/event"
	eventSchema "src/module/event/schema"
	"src/module/pm"
	"src/util/numberutil"
)

type Service struct {
	tenantRepo     TenantRepo
	gitAccountRepo GitAccountRepo
	gitRepoRepo    GitRepoRepo
	gitPushRepo    GitPushRepo
	gitCommitRepo  GitCommitRepo
	gitRepo        GitRepo
	messageRepo    MessageRepo
}

func New(
	tenantRepo TenantRepo,
	gitAccountRepo GitAccountRepo,
	gitRepoRepo GitRepoRepo,
	gitPushRepo GitPushRepo,
	gitCommitRepo GitCommitRepo,
	gitRepo GitRepo,
	messageRepo MessageRepo,
) Service {
	return Service{
		tenantRepo,
		gitAccountRepo,
		gitRepoRepo,
		gitPushRepo,
		gitCommitRepo,
		gitRepo,
		messageRepo,
	}
}

func (srv Service) HandleInstallCallback(
	uid string,
	tenantUid string,
) (*schema.GitAccount, error) {
	tenant, err := srv.tenantRepo.Retrieve(ctype.QueryOptions{
		Filters: ctype.Dict{
			"uid": tenantUid,
		},
	})
	if err != nil {
		return nil, err
	}
	tenantID := tenant.ID

	data := ctype.Dict{
		"Uid":      &uid,
		"TenantID": &tenantID,
	}

	result, err := srv.gitAccountRepo.UpdateOrCreate(ctype.QueryOptions{
		Filters: ctype.Dict{
			"Uid": &uid,
		},
	}, data)

	if err != nil {
		return nil, err
	}
	return result, nil
}

func (srv Service) HandleInstallWebhook(
	uid string,
	title string,
	avatar string,
	repos []GithubRepo,
) (*schema.GitAccount, error) {
	data := ctype.Dict{
		"Uid":    &uid,
		"Title":  title,
		"Avatar": avatar,
	}

	gitAccount, err := srv.gitAccountRepo.UpdateOrCreate(ctype.QueryOptions{
		Filters: ctype.Dict{
			"Uid": &uid,
		},
	}, data)
	if err != nil {
		return nil, err
	}

	gitAccountID := gitAccount.ID

	_, err = srv.gitRepoRepo.DeleteBy(ctype.QueryOptions{
		Filters: ctype.Dict{
			"GitAccountID": &gitAccountID,
		},
	})

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	for _, repo := range repos {
		data := ctype.Dict{
			"GitAccountID": gitAccountID,
			"RepoID":       numberutil.UintToStr(repo.ID),
			"Uid":          repo.FullName,
			"Private":      repo.Private,
		}

		_, err = srv.gitRepoRepo.Create(data)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
	}

	return gitAccount, nil
}

func (srv Service) HandleUninstallWebhook(
	uid string,
) error {
	_, err := srv.gitAccountRepo.DeleteBy(ctype.QueryOptions{
		Filters: ctype.Dict{
			"Uid": &uid,
		},
	})
	if err != nil {
		return err
	}

	return nil
}

func getBranchFromRef(ref string) string {
	return ref[11:]
}

func (srv Service) HandlePushWebhook(
	ref string,
	installationID string,
	gitRepo string,
	commits []GithubCommit,
) (ctype.Dict, error) {
	fmt.Println("HandlePushWebhook............")
	messageType := event.GIT_PUSHED
	gitBranch := getBranchFromRef(ref)

	taskUser, err := srv.gitRepo.GetTaskUser(gitRepo, gitBranch)
	if err != nil {
		fmt.Println("srv.gitRepo.GetTaskUser")
		fmt.Println(err)
		// return nil, err
	}

	gitPushData := ctype.Dict{
		"TaskID":        taskUser.TaskID,
		"UserID":        taskUser.UserID,
		"GitAccountUid": installationID,
		"GitRepoUid":    gitRepo,
		"GitHost":       pm.PROJECT_REPO_TYPE_GITHUB,
		"GitBranch":     gitBranch,
	}

	gitPush, err := srv.gitPushRepo.Create(gitPushData)
	if err != nil {
		fmt.Println("srv.gitPushRepo.Create")
		fmt.Println(err)
		return nil, err
	}

	var gitCommits []map[string]interface{}
	for _, commit := range commits {
		gitCommitData := ctype.Dict{
			"GitPushID":     gitPush.ID,
			"CommitID":      commit.ID,
			"CommitURL":     commit.URL,
			"CommitMessage": commit.Message,
		}

		result, err := srv.gitCommitRepo.Create(gitCommitData)
		if err != nil {
			fmt.Println("srv.gitCommitRepo.Create")
			fmt.Println(err)
			return nil, err
		}

		gitCommit := map[string]interface{}{
			"id":             result.ID,
			"commit_id":      commit.ID,
			"commit_url":     commit.URL,
			"commit_message": commit.Message,
			"created_at":     result.CreatedAt,
		}
		gitCommits = append(gitCommits, gitCommit)
	}

	if taskUser.UserID != nil {
		messageData := eventSchema.Message{
			TaskID:     *taskUser.TaskID,
			ProjectID:  *taskUser.ProjectID,
			Content:    "",
			Type:       messageType,
			UserID:     *taskUser.UserID,
			UserName:   *taskUser.UserName,
			UserAvatar: *taskUser.UserAvatar,
			UserColor:  *taskUser.UserColor,
			GitPush: map[string]interface{}{
				"id":          gitPush.ID,
				"git_branch":  gitBranch,
				"git_commits": gitCommits,
			},
		}
		_, err = srv.messageRepo.Create(messageData)
		if err != nil {
			fmt.Println("srv.messageRepo.Create")
			fmt.Println(err)
			return nil, err
		}
	}

	return ctype.Dict{}, nil
}
