# labelcommit
labelcommit is github actions that merge pull request with commit message including pull request labels.

see example below.
https://github.com/0daryo/labelcommit/pull/2
## usage
1. Write your workflow file.
  ```
  - name: labelcommit
    uses: 0daryo/labelcommit@main
  ```
  https://github.com/0daryo/labelcommit/blob/main/.github/workflows/commitlabel.yaml

2. comment ```/merge``` on github pull request comment.

3. pull request is merged, and commit message includes labels.
```
fix: readme
- documentation
- enhancement
```
