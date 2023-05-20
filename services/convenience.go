package services

import "github.com/emirkosuta/celeritas"

type BaseService struct {
	App *celeritas.Celeritas
}

func (s *BaseService) randomString(n int) string {
	return s.App.RandomString(n)
}

func (s *BaseService) encrypt(text string) (string, error) {
	enc := celeritas.Encryption{Key: []byte(s.App.EncryptionKey)}

	encrypted, err := enc.Encrypt(text)
	if err != nil {
		return "", err
	}
	return encrypted, nil
}

func (s *BaseService) decrypt(crypto string) (string, error) {
	enc := celeritas.Encryption{Key: []byte(s.App.EncryptionKey)}

	decrypted, err := enc.Decrypt(crypto)
	if err != nil {
		return "", err
	}
	return decrypted, nil
}
