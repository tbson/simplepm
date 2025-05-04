package srv

type repo1Provider interface {
	RepoHandler() error
}

type srv struct {
	repo1 repo1Provider
}

func New(repo1 repo1Provider) srv {
	return srv{repo1}
}

func (srv srv) Handler1() error {
	return nil
}
