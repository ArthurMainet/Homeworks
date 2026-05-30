package internal

type LocalRepo struct {
	EmailAndHash map[string]*EmailWithHash
}

func NewLocalRepo() *LocalRepo {
	return &LocalRepo{
		EmailAndHash: make(map[string]*EmailWithHash),
	}
}
