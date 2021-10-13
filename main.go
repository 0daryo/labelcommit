package main

import (
	"context"
	"log"
	"strings"

	"github.com/google/go-github/github"
	"github.com/kelseyhightower/envconfig"
	"golang.org/x/oauth2"
)

type env struct {
	GithubToken string `envconfig:"GITHUB_TOKEN"`
	Owner       string `envconfig:"OWNER"`
	Repo        string `envconfig:"REPO"`
	PRNumber    int    `envconfig:"PR_NUMBER"`
	MergeMethod string `envconfig:"MERGEMETHOD" default:"squash"`
}

func main() {
	var e env
	err := envconfig.Process("INPUT", &e)
	if err != nil {
		log.Fatal(err.Error())
	}
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: e.GithubToken},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)
	pr, _, err := client.PullRequests.Get(ctx, e.Owner, e.Repo, e.PRNumber)
	if err != nil {
		log.Fatal("failed to get pull request: " + err.Error())
	}
	labels := make([]string, 0, len(pr.Labels))
	for _, l := range pr.Labels {
		labels = append(labels, l.GetName())
	}

	var commitMessage string
	if len(labels) > 0 {
		commitMessage = "- " + strings.Join(labels, "\n- ")
	}

	_, _, err = client.PullRequests.Merge(ctx, e.Owner, e.Repo, e.PRNumber, commitMessage, &github.PullRequestOptions{
		CommitTitle: pr.GetTitle(),
		MergeMethod: e.MergeMethod,
	})
	if err != nil {
		log.Fatal("failed to merge pull request: " + err.Error())
	}
	log.Printf("Merged pull request https://github.com/%s/%s/pull/%d", e.Owner, e.Repo, e.PRNumber)
}
