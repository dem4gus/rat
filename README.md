# Rat: The Repository Audit Tool

## About

So you have your code repositories all set up to meet compliance requirements.
Your infrastructure as code is configured, your branches are protected,
and your pull requests require multiple reviewers before merging.
You *trust* your tools and your people, but you also need to *verify*.

Enter the Rat.
Its job is simple: it examines the settings on your repositories
and ensures reality matches your perfect vision.
Just give it a set of expected rules and a repository,
and Rat will make sure the rules are present.
Additionally, Rat can be set to Snitch.
This will examine the activity in a repository and report any anomalies,
such as commits on a protected branch that aren't associated with a pull request
or pull requests that have been merged without review or approval.

## Roadmap

Here is a list of planned features for Rat:

- Configurable rules
  - [ ] Default branch name
  - [ ] Branch protection for multiple branches
  - [ ] Required PR approvals
  - [ ] GitHub Action (or other CI) configuration
  - Whatever else I think of
- [ ] Configurable severity for rules
- [ ] Scanning all repositories in an organization, with an exclusion list
- [ ] Snitch mode, with report generation in a variety of formats
      (cli table, csv, json, yaml)
- Support for additional repository hosting platforms,
  including privately hosted servers
  - [ ] GitLab
  - [ ] Gitea
  - [ ] Codeberg/Forgejo
  - Others depending on time and interest
- [ ] Web interface (please don't expect much, I'm not a web dev)

## Development

Rat uses [task](https://taskfile.dev/) as a task runner.
You can run the same checks as CI with the following command:

```bash
task ci
```
