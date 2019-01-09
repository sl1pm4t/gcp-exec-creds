package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"golang.org/x/oauth2/google"
)

// ExecCredential is the Kubernetes ExecCredential API object
// Example:
// {
//   "apiVersion": "client.authentication.k8s.io/v1beta1",
//   "kind": "ExecCredential",
//   "status": {
//     "token": "my-bearer-token"
//   }
// }
type ExecCredential struct {
	APIVersion string               `json:"apiVersion"`
	Kind       string               `json:"ExecCredential"`
	Status     ExecCredentialStatus `json:"status"`
}

type ExecCredentialStatus struct {
	Token string `json:"token"`
}

func NewExecCredential(token string) ExecCredential {
	ec := ExecCredential{
		APIVersion: "client.authentication.k8s.io/v1beta1",
		Kind:       "ExecCredential",
		Status: ExecCredentialStatus{
			Token: token,
		},
	}
	return ec
}

func (ec ExecCredential) String() string {
	b, err := json.MarshalIndent(&ec, "", "  ")
	if err != nil {
		fmt.Printf("could not marshal ExecCredentials: %s", err)
		os.Exit(1)
	}
	return string(b)
}

func main() {
	clientScopes := []string{
		"https://www.googleapis.com/auth/cloud-platform",
	}

	tokenSource, err := google.DefaultTokenSource(context.Background(), clientScopes...)
	if err != nil {
		fmt.Printf("could not get tokenSource: %s", err)
		os.Exit(1)
	}

	token, err := tokenSource.Token()
	if err != nil {
		fmt.Printf("could not get token: %s", err)
		os.Exit(1)
	}

	ec := NewExecCredential(token.AccessToken)
	fmt.Println(ec.String())
}
