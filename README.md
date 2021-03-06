# labelcommit
labelcommit is github actions that merges pull request with commit message including pull request labels.

see example below.
https://github.com/0daryo/labelcommit/pull/7
## Usage
1. Write your workflow file.
  ```
  - name: labelcommit
    # uses: 0daryo/labelcommit@main
    # specifying prebuilt docker image is faster.
    uses: docker://0daryo/labelcommit:latest
  ```
  https://github.com/0daryo/labelcommit/blob/main/.github/workflows/commitlabel.yaml

2. comment ```/merge``` on github pull request comment.

3. pull request is merged, and commit message includes labels.
```
fix: readme
* documentation
* enhancement
```

## Parameters
You need to set parameters in workflow.
```
github: token: ${{ secrets.GITHUB_TOKEN }}
owner: repository owner
repo: repository name
pr number: ${{ github.event.issue.number }}
comment: ${{ github.event.comment.body }}
```
