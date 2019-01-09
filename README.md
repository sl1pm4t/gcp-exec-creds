# Kubernetes client-go credential helper for Google Cloud

`client-go` credential helper that generates Kubernetes `ExecCredentials` objects from the GCloud SDK account.

## Overview

Kubernetes provides a pluggable mechanism for getting user credentials for authenticating with the API server.
`k8s.io/client-go` client libraries and tools using it such as `kubectl` and `kubelet` are able to execute an external command to receive user credentials.
This tool reads the client credentials from the pre-authenticated GCloud SDK account on the local system and generates the appropriate `ExecCredential` API object that can be used by the above tools to authenticate.

See (https://kubernetes.io/docs/reference/access-authn-authz/authentication/#client-go-credential-plugins)

## Usage

Build and run `gcp-exec-creds`.
The command does not take any flags or arguments.

**Example**

```
gcp-exec-creds
{
  "apiVersion": "client.authentication.k8s.io/v1beta1",
  "ExecCredential": "ExecCredential",
  "status": {
    "token": "ya29.Gl2MBsw3YWqsD5......OqgeJE5LciQ"
  }
}
```

## Building

**Prerequisites**

- Golang 1.11

`go get -v github.com/sl1pm4t/gcp-exec-creds`

### How is this different from 'gcloud config config-helper' command?

This tool outputs the credentials in the `ExecCredential` API document format which may be necessary for some tools.
The `gcloud config config-helper` tool outputs in a non-standard json document.
