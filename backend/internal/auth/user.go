package auth

import (
	"crypto/ed25519"
	"database/sql"
	"encoding/base64"
	"fmt"
	"go-zkp/internal/usb"
	"os"
	"path/filepath"
)

type RegisterRequest struct {
	Username string `json:"username"`
}

type RegisterResponse struct {
	PrivateKey string `json:"private_key"`
}

func CreateUser(db *sql.DB, username string) (RegisterResponse, error) {
	pub, priv, err := ed25519.GenerateKey(nil)
	if err != nil {
		return RegisterResponse{}, err
	}

	_, err = db.Exec("INSERT INTO users (username, public_key) VALUES (?, ?)", username, pub)
	if err != nil {
		return RegisterResponse{}, err
	}

	return RegisterResponse{
		PrivateKey: base64.StdEncoding.EncodeToString(priv),
	}, nil
}

func FindPrivateKeyOnUSB() (string, error) {
	hasUSB, mounts := usb.CheckUSB()
	if !hasUSB {
		return "", fmt.Errorf("usb_not_found")
	}

	for _, mount := range mounts {
		authKeyPath := filepath.Join(mount, "auth_key")
		info, err := os.Stat(authKeyPath)
		if err != nil || !info.IsDir() {
			continue
		}

		fpath := filepath.Join(authKeyPath, "private_key.txt")
		data, err := os.ReadFile(fpath)
		if err != nil {
			return "", fmt.Errorf("private_key_not_found")
		}
		return string(data), nil
	}
	return "", fmt.Errorf("private_key_not_found")
}
