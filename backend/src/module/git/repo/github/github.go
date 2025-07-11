package github

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"src/common/setting"
	"src/util/errutil"
	"src/util/i18nmsg"

	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwt"
)

type Repo struct{}

func New() Repo {
	return Repo{}
}

type GitRepo struct {
	ID       uint   `json:"id"`
	FullName string `json:"full_name"`
}

type TokenResult struct {
	Token string `json:"token"`
}

type GitbRepoResult struct {
	Repositories []GitRepo `json:"repositories"`
}

func (r Repo) GetInstallUrl(tenantUid string) string {
	publicLink := setting.GITHUB_APP_PUBLIC_LINK()
	return fmt.Sprintf("%s?state=%s", publicLink, tenantUid)
}

func (r Repo) generateJWT() (string, error) {
	clientID := setting.GITHUB_CLIENT_ID()
	privateKeyPath := setting.GITHUB_PRIVATE_KEY_PATH()

	// Read your private key file.
	keyBytes, err := os.ReadFile(privateKeyPath)
	if err != nil {
		return "", errutil.New(i18nmsg.FailedToReadPrivateKey)
	}

	// Parse the RSA private key.
	privateKey, err := parseRSAPrivateKeyFromPEM(keyBytes)
	if err != nil {
		return "", errutil.New(i18nmsg.FailedToReadPrivateKey)
	}

	// Create a new token with claims.
	tok, err := jwt.NewBuilder().
		Claim(jwt.IssuedAtKey, time.Now()).
		Claim(jwt.ExpirationKey, time.Now().Add(10*time.Minute)).
		Claim(jwt.IssuerKey, clientID).
		Build()
	if err != nil {
		return "", errutil.New(i18nmsg.FailedToBuildToken)
	}

	// Sign the token using RS256.
	signed, err := jwt.Sign(tok, jwt.WithKey(jwa.RS256, privateKey))
	if err != nil {
		return "", errutil.New(i18nmsg.FailedToSignToken)
	}

	return string(signed), nil
}

func (r Repo) retInstallationToken(installationID string) (string, error) {
	emptyResult := ""
	jwt, err := r.generateJWT()
	if err != nil {
		return emptyResult, err
	}
	url := fmt.Sprintf(
		"https://api.github.com/app/installations/%s/access_tokens", installationID,
	)
	authHeader := fmt.Sprintf("Bearer %s", jwt)

	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return emptyResult, errutil.New(i18nmsg.CanNotCreateRequest)
	}

	// Set required headers.
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("Authorization", authHeader)
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")

	// Create a HTTP client and send the request.
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return emptyResult, errutil.New(i18nmsg.CanNotSendRequest)
	}
	defer resp.Body.Close()

	// Read and print the response body.
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return emptyResult, errutil.New(i18nmsg.CanNotReadResponse)
	}
	// body is a JSON object, so we can unmarshal it into a struct.
	result := TokenResult{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return emptyResult, errutil.New(i18nmsg.CanNotReadResponse)
	}

	return result.Token, nil
}

func (r Repo) getRepoListOfInstallation(installationID string) (GitbRepoResult, error) {
	emptyResult := GitbRepoResult{}
	jwt, err := r.retInstallationToken(installationID)
	if err != nil {
		return emptyResult, err
	}
	url := "https://api.github.com/installation/repositories"
	authHeader := fmt.Sprintf("Bearer %s", jwt)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return emptyResult, errutil.New(i18nmsg.CanNotCreateRequest)
	}

	// Set required headers.
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("Authorization", authHeader)
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")

	// Create a HTTP client and send the request.
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return emptyResult, errutil.New(i18nmsg.CanNotSendRequest)
	}
	defer resp.Body.Close()

	// Read and print the response body.
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return emptyResult, errutil.New(i18nmsg.CanNotReadResponse)
	}
	// body is a JSON object, so we can unmarshal it into a struct.
	result := GitbRepoResult{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return emptyResult, errutil.New(i18nmsg.CanNotReadResponse)
	}

	return result, nil
}

func (r Repo) GetRepoList(installationIDs []string) (GitbRepoResult, error) {
	// execute getRepoListOfInstallation in parallel
	ch := make(chan GitbRepoResult)
	for _, installationID := range installationIDs {
		go func(installationID string) {
			result, err := r.getRepoListOfInstallation(installationID)
			if err != nil {
				ch <- GitbRepoResult{}
				return
			}
			ch <- result
		}(installationID)
	}

	// collect results
	result := GitbRepoResult{}
	for range installationIDs {
		res := <-ch
		result.Repositories = append(result.Repositories, res.Repositories...)
	}

	return result, nil
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
