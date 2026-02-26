package audit

import (
	"context"

	"github.com/google/go-github/v83/github"
)

// InfoGetter gets information about a single repository from GitHub.
type InfoGetter interface {
	Get(ctx context.Context, owner, name string) (*github.Repository, *github.Response, error)
	GetBranch(ctx context.Context, owner, repo, branch string, maxRedirects int) (*github.Branch, *github.Response, error)
}

// Client contains the information required to run an audit
// and the API client for the hosting platform.
type Client struct {
	owner, name string
	client      InfoGetter
}

// NewClient returns a new auditor for the specified platform and repository.
func NewClient(ghclient *github.Client, owner, name string) *Client {
	client := &Client{
		owner:  owner,
		name:   name,
		client: ghclient.Repositories,
	}
	return client
}

// Run runs an audit for a repository and returns a [*Report].
func (a Client) Run(ctx context.Context) (*Report, error) {
	repo, _, err := a.client.Get(ctx, a.owner, a.name)
	if err != nil {
		return nil, err
	}

	fullName := repo.GetFullName()
	branchName := repo.GetDefaultBranch()

	branch, _, err := a.client.GetBranch(ctx, a.owner, a.name, branchName, 3)
	if err != nil {
		return nil, err
	}

	report := &Report{
		FullName:  fullName,
		Branch:    branchName,
		Protected: branch.GetProtected(),
	}

	return report, nil
}

// Report is the results of the audit of the repository.
type Report struct {
	FullName  string
	Branch    string
	Protected bool
}
