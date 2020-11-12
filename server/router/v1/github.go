package v1

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	"github.com/google/go-github/v32/github"
	"golang.org/x/oauth2"
)

type (
	// GitHubClientI defines the interface used to retrieve GitHub repository
	// information.
	GitHubClientI interface {
		GetRepository(repoURL string) (Repository, error)
	}

	// Repository defines the relevant information Atlas needs for a GitHub
	// repository in order to publish modules.
	Repository struct {
		Owner        string
		Repo         string
		Contributors map[string]*github.Contributor
	}
)

// GitHubClient implements a wrapper around a GitHub v3 API client.
type GitHubClient struct {
	*github.Client
}

func NewGitHubClient(token string) *GitHubClient {
	if len(token) != 0 {
		return &GitHubClient{
			Client: github.NewClient(
				oauth2.NewClient(context.Background(),
					oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token}),
				),
			),
		}
	}

	return &GitHubClient{Client: github.NewClient(nil)}
}

// GetRepository returns a Repository object which contains information needed
// when publishing a Cosmos SDK module. It returns an error if the repository
// URL is invalid or if any resource fails to be fetched from the GitHub API.
func (gc *GitHubClient) GetRepository(repoURL string) (Repository, error) {
	repo, err := parseGitHubRepo(repoURL)
	if err != nil {
		return Repository{}, err
	}

	ghRepo, _, err := gc.Repositories.Get(context.Background(), repo.Owner, repo.Repo)
	if err != nil {
		return Repository{}, fmt.Errorf("failed to fetch repository: %w", err)
	}

	if repo.Owner != ghRepo.Owner.GetLogin() {
		return Repository{}, fmt.Errorf("unexpected owner; got: %s, want: %s", ghRepo.Owner.GetLogin(), repo.Owner)
	}

	opts := &github.ListContributorsOptions{Anon: "false", ListOptions: github.ListOptions{Page: 1, PerPage: 100}}
	ghContributors, _, err := gc.Repositories.ListContributors(context.Background(), repo.Owner, repo.Repo, opts)
	if err != nil {
		return Repository{}, fmt.Errorf("failed to get repository contributors: %w", err)
	}

	contributors := make(map[string]*github.Contributor)
	for len(ghContributors) > 0 {
		for _, c := range ghContributors {
			contributors[c.GetLogin()] = c
		}

		opts = &github.ListContributorsOptions{Anon: "false", ListOptions: github.ListOptions{Page: opts.Page + 1, PerPage: 100}}
		ghContributors, _, err = gc.Repositories.ListContributors(context.Background(), repo.Owner, repo.Repo, opts)
		if err != nil {
			return Repository{}, fmt.Errorf("failed to get repository contributors: %w", err)
		}
	}

	repo.Contributors = contributors

	return repo, nil
}

func parseGitHubRepo(repoURL string) (Repository, error) {
	u, err := url.Parse(repoURL)
	if err != nil {
		return Repository{}, fmt.Errorf("failed to parse repo URL: %w", err)
	}

	path := u.Path

	// remove query
	path = strings.Split(path, "?")[0]

	// remove .git extension
	path = strings.Split(path, ".git")[0]

	tokens := strings.Split(path, "/")
	if len(tokens) != 3 || len(tokens[1]) == 0 || len(tokens[2]) == 0 {
		return Repository{}, fmt.Errorf("invalid repository: %s", repoURL)
	}

	return Repository{Owner: tokens[1], Repo: tokens[2]}, nil
}
