package github

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
	"os"
	"time"

	"src/common/setting"

	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwt"
)

type Repo struct{}

func New() Repo {
	return Repo{}
}

type TokenResult struct {
	Token string `json:"token"`
}

func (r Repo) GetToken() (TokenResult, error) {
	result := TokenResult{
		Token: "github",
	}
	return result, nil
}

func (r Repo) GenerateJWT() (string, error) {
	// Replace with your GitHub App's ID and your private key file path.
	clientID := setting.GITHUB_CLIENT_ID
	privateKeyPath := setting.GITHUB_PRIVATE_KEY_PATH

	// Read your private key file.
	keyBytes, err := os.ReadFile(privateKeyPath)
	if err != nil {
		log.Fatalf("failed to read private key: %v", err)
		return "", err
	}

	// Parse the RSA private key.
	privateKey, err := parseRSAPrivateKeyFromPEM(keyBytes)
	if err != nil {
		log.Fatalf("failed to parse RSA private key: %v", err)
		return "", err
	}

	// Create a new token with claims.
	tok, err := jwt.NewBuilder().
		Claim(jwt.IssuedAtKey, time.Now()).
		Claim(jwt.ExpirationKey, time.Now().Add(10*time.Minute)).
		Claim(jwt.IssuerKey, clientID).
		Build()
	if err != nil {
		log.Fatalf("failed to build token: %v", err)
		return "", err
	}

	// Sign the token using RS256.
	signed, err := jwt.Sign(tok, jwt.WithKey(jwa.RS256, privateKey))
	if err != nil {
		log.Fatalf("failed to sign token: %v", err)
	}

	return string(signed), nil
}

func parseRSAPrivateKeyFromPEM(pemBytes []byte) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode(pemBytes)
	if block == nil {
		return nil, fmt.Errorf("failed to decode PEM block")
	}
	// Try PKCS1 first
	key, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err == nil {
		return key, nil
	}
	// Fallback to PKCS8
	keyInterface, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	rsaKey, ok := keyInterface.(*rsa.PrivateKey)
	if !ok {
		return nil, fmt.Errorf("not an RSA private key")
	}
	return rsaKey, nil
}
