package infra

import (
	"fmt"
	"io"
	"net/http"
	"src/client/scyllaclient"
	"src/common/ctype"
	"src/util/dbutil"
	"src/util/numberutil"
	"src/util/vldtutil"

	"src/module/git/repo/github"

	"src/module/account/repo/gitaccount"
	"src/module/account/repo/gitrepo"
	"src/module/account/repo/tenant"
	"src/module/event/repo/centrifugo"
	"src/module/event/repo/message"
	"src/module/pm/repo/gitcommit"
	"src/module/pm/repo/gitpush"

	"src/module/git/usecase/github/app"

	"github.com/labstack/echo/v4"
)

func GetInstallUrl(c echo.Context) error {
	gitaccountRepo := github.New()
	tenantUid := c.Get("TenantUid").(string)
	url := gitaccountRepo.GetInstallUrl(tenantUid)
	result := ctype.Dict{
		"url": url,
	}
	return c.JSON(http.StatusOK, result)
}

func Callback(c echo.Context) error {
	tenantRepo := tenant.New(dbutil.Db(nil))
	gitaccountRepo := gitaccount.New(dbutil.Db(nil))
	gitRepoRepo := gitrepo.New(dbutil.Db(nil))
	gitPushRepo := gitpush.New(dbutil.Db(nil))
	gitCommitRepo := gitcommit.New(dbutil.Db(nil))
	messageRepo := message.New(scyllaclient.NewClient())
	centrifugoRepo := centrifugo.New()
	gitRepo := New(dbutil.Db(nil))

	srv := app.New(
		tenantRepo,
		gitaccountRepo,
		gitRepoRepo,
		gitPushRepo,
		gitCommitRepo,
		gitRepo,
		messageRepo,
		centrifugoRepo,
	)

	setupAction := c.QueryParam("setup_action")
	installationID := c.QueryParam("installation_id")
	tenantUid := c.QueryParam("state")

	if setupAction == app.GITHUB_CALLBACK_ACTION_INSTALL {
		_, err := srv.HandleInstallCallback(installationID, tenantUid)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		return c.Redirect(http.StatusFound, "/account/tenant/setting")
	}

	return c.JSON(http.StatusOK, ctype.Dict{})
}

func WebhookTest(c echo.Context) error {
	bodyBytes, err := io.ReadAll(c.Request().Body)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	fmt.Println(string(bodyBytes))
	return c.JSON(http.StatusOK, ctype.Dict{})
}

func Webhook(c echo.Context) error {
	tenantRepo := tenant.New(dbutil.Db(nil))
	gitaccountRepo := gitaccount.New(dbutil.Db(nil))
	gitRepoRepo := gitrepo.New(dbutil.Db(nil))
	gitPushRepo := gitpush.New(dbutil.Db(nil))
	gitCommitRepo := gitcommit.New(dbutil.Db(nil))
	messageRepo := message.New(scyllaclient.NewClient())
	centrifugoRepo := centrifugo.New()
	gitRepo := New(dbutil.Db(nil))

	srv := app.New(
		tenantRepo,
		gitaccountRepo,
		gitRepoRepo,
		gitPushRepo,
		gitCommitRepo,
		gitRepo,
		messageRepo,
		centrifugoRepo,
	)

	data, err := vldtutil.ValidatePayload(c, app.GithubWebhook{})
	if err != nil {
		fmt.Println("error")
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, err)
	}

	uid := numberutil.UintToStr(data.Installation.ID)
	title := data.Sender.Login
	avatar := data.Sender.AvatarURL
	repos := data.Repositories

	if data.Action == app.GITHUB_WEBHOOK_ACTION_CREATED {
		_, err = srv.HandleInstallWebhook(uid, title, avatar, repos)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
	}

	if data.Action == app.GITHUB_WEBHOOK_ACTION_DELETED {
		err = srv.HandleUninstallWebhook(uid)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
	}

	if data.Ref != "" {
		ref := data.Ref
		installationID := data.Installation.ID
		repoUid := data.Repository.FullName
		commits := data.Commits
		_, err = srv.HandlePushWebhook(
			ref,
			numberutil.UintToStr(installationID),
			repoUid,
			commits,
		)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
	}

	if data.PullRequest.ID != "" {
		installationID := data.Installation.ID
		pullRequest := data.PullRequest
		repoUid := data.Repository.FullName
		_, err = srv.HandlePrWebhook(
			numberutil.UintToStr(installationID),
			repoUid,
			pullRequest,
			data.Action,
		)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
	}

	return c.JSON(http.StatusOK, ctype.Dict{})
}
