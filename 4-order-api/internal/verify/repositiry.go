package verify

import (
	"encoding/json"
	"os"
)

type LocalRepo struct {
	EmailAndHash map[string]*EmailWithHash
}

var storageFile string = "verify.json"

func NewLocalRepo() *LocalRepo {
	repo := LocalRepo{
		EmailAndHash: make(map[string]*EmailWithHash),
	}
	repo.Load()
	return &repo
}

func (l *LocalRepo) Load() error {
	data, err := os.ReadFile(storageFile)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	return json.Unmarshal(data, &l.EmailAndHash)
}

func (l *LocalRepo) Save() error {
	data, err := json.MarshalIndent(l.EmailAndHash, "", " ")
	if err != nil {
		return err
	}
	os.WriteFile(storageFile, data, 0644)
	return nil
}

func (l *LocalRepo) Delete(hash string) {
	delete(l.EmailAndHash, hash)
}
