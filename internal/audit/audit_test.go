package audit

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-github/v83/github"
	ghmock "github.com/migueleliasweb/go-github-mock/src/mock"
)

func TestRun(t *testing.T) {
	t.Run("reports on a single repository", func(t *testing.T) {
		mock := ghmock.NewMockedHTTPClient(
			ghmock.WithRequestMatch(
				ghmock.GetReposByOwnerByRepo,
				github.Repository{
					FullName:      new("foo/bar"),
					DefaultBranch: new("main"),
				},
			),
			ghmock.WithRequestMatch(
				ghmock.GetReposBranchesByOwnerByRepoByBranch,
				github.Branch{
					Name:      new("main"),
					Protected: new(true),
				},
			),
		)

		ghclient := github.NewClient(mock)
		audit := NewClient(ghclient, "foo", "bar")

		want := &Report{
			FullName:  "foo/bar",
			Branch:    "main",
			Protected: true,
		}

		got, err := audit.Run(t.Context())
		if err != nil {
			t.Fatalf("failed to run test, got unexpected error: %v", err)
		}
		if !cmp.Equal(want, got) {
			t.Errorf("wanted %+v, got %+v", want, got)
		}
	})
}
