package cmd

import "github.com/WhizZest/gh-pr-review/internal/ghcli"

var apiClientFactory = func(host string) ghcli.API {
	return &ghcli.Client{Host: host}
}
