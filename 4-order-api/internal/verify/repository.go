package verify

import (
	"encoding/json"
	"os"
)

type LocalRepo struct {
	AdrressAndHash map[string]map[string]*EmailWithHash
}

var storageFile string = "verify.json"

func NewLocalRepo() *LocalRepo {
	repo := LocalRepo{
		AdrressAndHash: make(map[string]map[string]*EmailWithHash),
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
	return json.Unmarshal(data, &l.AdrressAndHash)
}

func (l *LocalRepo) Save() error {
	data, err := json.MarshalIndent(l.AdrressAndHash, "", " ")
	if err != nil {
		return err
	}
	os.WriteFile(storageFile, data, 0644)
	return nil
}

func (l *LocalRepo) Delete(method, hash string) {
	delete(l.AdrressAndHash[method], hash)
}
