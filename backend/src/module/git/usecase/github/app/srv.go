package app

import (
	"fmt"
	"src/common/ctype"
	"src/module/account/schema"
	"src/module/event"
	"src/module/pm"
	"src/util/dictutil"
	"src/util/numberutil"
	"strings"
)

type Service struct {
	tenantRepo     TenantRepo
	gitAccountRepo GitAccountRepo
	gitRepoRepo    GitRepoRepo
	gitPushRepo    GitPushRepo
	gitCommitRepo  GitCommitRepo
	gitRepo        GitRepo
	messageRepo    MessageRepo
	centrifugoRepo CentrifugoRepo
}

func New(
	tenantRepo TenantRepo,
	gitAccountRepo GitAccountRepo,
	gitRepoRepo GitRepoRepo,
	gitPushRepo GitPushRepo,
	gitCommitRepo GitCommitRepo,
	gitRepo GitRepo,
	messageRepo MessageRepo,
	centrifugoRepo CentrifugoRepo,
) Service {
	return Service{
		tenantRepo,
		gitAccountRepo,
		gitRepoRepo,
		gitPushRepo,
		gitCommitRepo,
		gitRepo,
		messageRepo,
		centrifugoRepo,
	}
}

func (srv Service) HandleInstallCallback(
	uid string,
	tenantUid string,
) (*schema.GitAccount, error) {
	tenant, err := srv.tenantRepo.Retrieve(ctype.QueryOpts{
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

	result, err := srv.gitAccountRepo.UpdateOrCreate(ctype.QueryOpts{
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
	repos []RepoInput,
) (*schema.GitAccount, error) {
	data := ctype.Dict{
		"Uid":    &uid,
		"Title":  title,
		"Avatar": avatar,
	}

	gitAccount, err := srv.gitAccountRepo.UpdateOrCreate(ctype.QueryOpts{
		Filters: ctype.Dict{
			"Uid": &uid,
		},
	}, data)
	if err != nil {
		return nil, err
	}

	gitAccountID := gitAccount.ID

	_, err = srv.gitRepoRepo.DeleteBy(ctype.QueryOpts{
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
	_, err := srv.gitAccountRepo.DeleteBy(ctype.QueryOpts{
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
	result := strings.Replace(ref, "refs/", "", 1)
	result = strings.Replace(result, "heads/", "", 1)
	return result
}

func (srv Service) HandlePushWebhook(
	ref string,
	installationID string,
	gitRepo string,
	commits []CommitInput,
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

	var gitCommits []ctype.Dict
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

		gitCommit := ctype.Dict{
			"id":             numberutil.UintToStr(result.ID),
			"commit_id":      commit.ID,
			"commit_url":     commit.URL,
			"commit_message": commit.Message,
			"created_at":     result.CreatedAt,
		}
		gitCommits = append(gitCommits, gitCommit)
	}

	if taskUser.UserID != nil {
		messageData := ctype.Dict{
			"task_id":     *taskUser.TaskID,
			"project_id":  *taskUser.ProjectID,
			"content":     "",
			"type":        messageType,
			"user_id":     *taskUser.UserID,
			"user_name":   *taskUser.UserName,
			"user_avatar": *taskUser.UserAvatar,
			"user_color":  *taskUser.UserColor,
			"git_push": ctype.Dict{
				"id":          numberutil.UintToStr(gitPush.ID),
				"git_branch":  gitBranch,
				"git_commits": gitCommits,
			},
		}
		message, err := srv.messageRepo.Create(messageData)
		if err != nil {
			return nil, err
		}

		projectID := *taskUser.ProjectID
		taskID := *taskUser.TaskID
		channel := fmt.Sprintf("%d/%d", projectID, taskID)
		socketUser := SocketUser{
			ID:     *taskUser.UserID,
			Name:   *taskUser.UserName,
			Avatar: *taskUser.UserAvatar,
			Color:  *taskUser.UserColor,
		}

		socketMessage := SocketMessage{
			Channel: channel,
			Data: SocketData{
				ID:        message.ID,
				Type:      messageType,
				User:      socketUser,
				TaskID:    taskID,
				ProjectID: projectID,
				Content:   messageType,
				GitData:   dictutil.StructToDict(message.GitPush),
			},
		}
		err = srv.centrifugoRepo.Publish(socketMessage)
		if err != nil {
			return nil, err
		}

	}

	return ctype.Dict{}, nil
}

func (srv Service) HandlePrWebhook(
	installationID string,
	gitRepo string,
	pullRequest PullRequestInput,
	action string,
) (ctype.Dict, error) {
	fmt.Println("HandlePrWebhook............")

	gitBranch := getBranchFromRef(pullRequest.Head.Ref)
	toBranch := getBranchFromRef(pullRequest.Base.Ref)
	mergedAt := pullRequest.MergedAt.TimePtr()

	messageType := event.GIT_PR_CREATED

	if action == GITHUB_WEBHOOK_PR_CLOSED {
		if mergedAt == nil {
			messageType = event.GIT_PR_CLOSED
		} else {
			messageType = event.GIT_PR_MERGED
		}
	}

	taskUser, err := srv.gitRepo.GetTaskUser(gitRepo, gitBranch)
	if err != nil {
		return nil, err
	}
	messageData := ctype.Dict{
		"task_id":     *taskUser.TaskID,
		"project_id":  *taskUser.ProjectID,
		"content":     "",
		"type":        messageType,
		"user_id":     *taskUser.UserID,
		"user_name":   *taskUser.UserName,
		"user_avatar": *taskUser.UserAvatar,
		"user_color":  *taskUser.UserColor,
		"git_pr": ctype.Dict{
			"id":          pullRequest.ID.String(),
			"title":       pullRequest.Title,
			"from_branch": gitBranch,
			"to_branch":   toBranch,
			"url":         pullRequest.URL,
			"merged_at":   mergedAt,
			"state":       pullRequest.State,
		},
	}
	message, err := srv.messageRepo.Create(messageData)
	if err != nil {
		return nil, err
	}

	projectID := *taskUser.ProjectID
	taskID := *taskUser.TaskID
	channel := fmt.Sprintf("%d/%d", projectID, taskID)
	socketUser := SocketUser{
		ID:     *taskUser.UserID,
		Name:   *taskUser.UserName,
		Avatar: *taskUser.UserAvatar,
		Color:  *taskUser.UserColor,
	}

	socketMessage := SocketMessage{
		Channel: channel,
		Data: SocketData{
			ID:        message.ID,
			Type:      messageType,
			User:      socketUser,
			TaskID:    taskID,
			ProjectID: projectID,
			Content:   messageType,
			GitData:   dictutil.StructToDict(message.GitPR),
		},
	}
	err = srv.centrifugoRepo.Publish(socketMessage)
	if err != nil {
		fmt.Println("srv.centrifugoRepo.Publish")
		fmt.Println(err)
		return nil, err
	}

	return ctype.Dict{}, nil
}
