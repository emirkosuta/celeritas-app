package services

import "github.com/emirkosuta/celeritas"

func (s *Services) randomString(n int) string {
	return s.App.RandomString(n)
}

func (s *Services) encrypt(text string) (string, error) {
	enc := celeritas.Encryption{Key: []byte(s.App.EncryptionKey)}

	encrypted, err := enc.Encrypt(text)
	if err != nil {
		return "", err
	}
	return encrypted, nil
}

func (s *Services) decrypt(crypto string) (string, error) {
	enc := celeritas.Encryption{Key: []byte(s.App.EncryptionKey)}

	decrypted, err := enc.Decrypt(crypto)
	if err != nil {
		return "", err
	}
	return decrypted, nil
}
