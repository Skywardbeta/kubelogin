# kubelogin [![go](https://github.com/int128/kubelogin/actions/workflows/go.yaml/badge.svg)](https://github.com/int128/kubelogin/actions/workflows/go.yaml) [![Go Report Card](https://goreportcard.com/badge/github.com/int128/kubelogin)](https://goreportcard.com/report/github.com/int128/kubelogin)

This is a kubectl plugin for [Kubernetes OpenID Connect (OIDC) authentication](https://kubernetes.io/docs/reference/access-authn-authz/authentication/#openid-connect-tokens), also known as `kubectl oidc-login`.

Here is an example of Kubernetes authentication with the Google Identity Platform:

<img alt="screencast" src="https://user-images.githubusercontent.com/321266/85427290-86e43700-b5b6-11ea-9e97-ffefd736c9b7.gif" width="572" height="391">

Kubelogin is designed to run as a [client-go credential plugin](https://kubernetes.io/docs/reference/access-authn-authz/authentication/#client-go-credential-plugins).
When you run kubectl, kubelogin opens the browser and you can log in to the provider.
Then kubelogin gets a token from the provider and kubectl access Kubernetes APIs with the token.
Take a look at the diagram:

