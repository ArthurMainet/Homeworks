package verify

import (
	"encoding/json"
	"os"
)

type LocalRepo struct {
	EmailAndHash map[string]*EmailWithHash
	PhoneAndCode map[int]*SessionWithCode
}

var emailFile string = "verifyEmail.json"
var sessionFile string = "verifySession.json"

func NewLocalRepo() *LocalRepo {
	repo := LocalRepo{
		EmailAndHash: make(map[string]*EmailWithHash),
		PhoneAndCode: make(map[int]*SessionWithCode),
	}
	repo.LoadEmailData()
	return &repo
}

//////////////////////////////////
////////// Phone REPO ///////////

func (l *LocalRepo) LoadPhoneData() error {
	data, err := os.ReadFile(sessionFile)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	return json.Unmarshal(data, &l.PhoneAndCode)
}

func (l *LocalRepo) SaveSessionCode() error {
	data, err := json.MarshalIndent(l.PhoneAndCode, "", " ")
	if err != nil {
		return err
	}
	os.WriteFile(sessionFile, data, 0644)
	return nil
}

func (l *LocalRepo) DeleteSession(code int) {
	delete(l.PhoneAndCode, code)
}

//////////////////////////////////
////////// EMAIL REPO ///////////

func (l *LocalRepo) LoadEmailData() error {
	data, err := os.ReadFile(emailFile)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	return json.Unmarshal(data, &l.EmailAndHash)
}

func (l *LocalRepo) SaveEmailHash() error {
	data, err := json.MarshalIndent(l.EmailAndHash, "", " ")
	if err != nil {
		return err
	}
	os.WriteFile(emailFile, data, 0644)
	return nil
}

func (l *LocalRepo) DeleteEmail(hash string) {
	delete(l.EmailAndHash, hash)
}
