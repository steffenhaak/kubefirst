package pkg

import (
	"net/http"
)

type HTTPDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

var SupportedPlatforms = []string{"aws-github", "aws-gitlab", "civo-github", "civo-github-adapted", "civo-gitlab", "k3d-github", "k3d-gitlab"}
