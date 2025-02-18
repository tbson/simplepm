package infra

import (
	"fmt"
	"io"
	"net/http"
	"src/common/ctype"
	"src/util/dbutil"
	"src/util/numberutil"
	"src/util/vldtutil"

	"src/module/git/repo/github"

	"src/module/account/repo/gitaccount"
	"src/module/account/repo/gitrepo"
	"src/module/account/repo/tenant"
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
	tenantRepo := tenant.New(dbutil.Db())
	gitaccountRepo := gitaccount.New(dbutil.Db())
	gitRepoRepo := gitrepo.New(dbutil.Db())
	gitPushRepo := gitpush.New(dbutil.Db())
	gitCommitRepo := gitcommit.New(dbutil.Db())
	gitRepo := New(dbutil.Db())

	srv := app.New(
		tenantRepo,
		gitaccountRepo,
		gitRepoRepo,
		gitPushRepo,
		gitCommitRepo,
		gitRepo,
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

func Webhook1(c echo.Context) error {
	bodyBytes, err := io.ReadAll(c.Request().Body)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	fmt.Println(string(bodyBytes))
	return c.JSON(http.StatusOK, ctype.Dict{})
}

func Webhook(c echo.Context) error {
	tenantRepo := tenant.New(dbutil.Db())
	gitaccountRepo := gitaccount.New(dbutil.Db())
	gitRepoRepo := gitrepo.New(dbutil.Db())
	gitPushRepo := gitpush.New(dbutil.Db())
	gitCommitRepo := gitcommit.New(dbutil.Db())
	gitRepo := New(dbutil.Db())

	srv := app.New(
		tenantRepo,
		gitaccountRepo,
		gitRepoRepo,
		gitPushRepo,
		gitCommitRepo,
		gitRepo,
	)

	structData, err := vldtutil.ValidatePayload(c, app.GithubWebhook{})
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	uid := numberutil.UintToStr(structData.Installation.ID)
	title := structData.Sender.Login
	avatar := structData.Sender.AvatarURL
	repos := structData.Repositories

	if structData.Action == app.GITHUB_WEBHOOK_ACTION_CREATED {
		_, err = srv.HandleInstallWebhook(uid, title, avatar, repos)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
	}

	if structData.Action == app.GITHUB_WEBHOOK_ACTION_DELETED {
		err = srv.HandleUninstallWebhook(uid)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
	}

	fmt.Println("structData.Ref", structData.Ref)
	if structData.Ref != "" {
		ref := structData.Ref
		installationID := structData.Installation.ID
		repoUid := structData.Repository.FullName
		commits := structData.Commits
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

	return c.JSON(http.StatusOK, ctype.Dict{})
}
