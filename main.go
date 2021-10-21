package main

import (
	"context"
	"fmt"
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
	Comment     string `envconfig:"COMMENT"`
	MergeMethod string `envconfig:"MERGEMETHOD" default:"squash"`
}

const (
	mergeComment = "/merge"
)

func main() {
	var e env
	err := envconfig.Process("INPUT", &e)
	if err != nil {
		log.Fatal(err.Error())
	}

	if e.Comment != mergeComment {
		return
	}
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: e.GithubToken},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)
	if err := merge(ctx, client, e.Owner, e.Repo, e.PRNumber, e.MergeMethod); err != nil {
		if err := sendMsg(ctx, client, e.Owner, e.Repo, e.PRNumber, err.Error()); err != nil {
			log.Fatal(err.Error())
		}
		return
	}
	successMsg := "Merged PR #" + fmt.Sprintf("%d", e.PRNumber) + " successfully!"
	if err := sendMsg(ctx, client, e.Owner, e.Repo, e.PRNumber, successMsg); err != nil {
		log.Fatal(err.Error())
		return
	}
	log.Printf(successMsg)
}

func sendMsg(ctx context.Context, client *github.Client, owner, repo string, prNumber int, msg string) error {
	_, ghResp, err := client.Issues.CreateComment(ctx, owner, repo, prNumber, &github.IssueComment{
		Body: &msg,
	})
	if err != nil {
		return fmt.Errorf("failed to send message: %v, githubResponse: %s", err, ghResp.String())
	}
	return nil
}

func merge(ctx context.Context, client *github.Client, owner, repo string, prNumber int, mergeMethod string) error {
	pr, _, err := client.PullRequests.Get(ctx, owner, repo, prNumber)
	if err != nil {
		return fmt.Errorf("failed to get pull request: %v", err)
	}
	labels := make([]string, 0, len(pr.Labels))
	for _, l := range pr.Labels {
		labels = append(labels, l.GetName())
	}

	var commitMessage string
	if len(labels) > 0 {
		commitMessage = "- " + strings.Join(labels, "\n- ")
	}

	_, _, err = client.PullRequests.Merge(ctx, owner, repo, prNumber, commitMessage, &github.PullRequestOptions{
		CommitTitle: pr.GetTitle(),
		MergeMethod: mergeMethod,
	})
	if err != nil {
		return fmt.Errorf("failed to merge pull request: %v", err)
	}
	return nil
}
