package extend

type skipBaseExecutor struct {
}

func (s skipBaseExecutor) Prepare() error {
	panic("implement me")
}

func (s skipBaseExecutor) VerifyInputs() error {
	panic("implement me")
}

func (s skipBaseExecutor) ApplyTransaction() error {
	panic("implement me")
}

func (s skipBaseExecutor) GeneratePubData() error {
	panic("implement me")
}
