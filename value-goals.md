# Backlog CLI

## Goals
- Create a CLI utility to replace the most common Backlog GUI operations
- Practically & pleasantly augment the user experience of interacting with Backlog in clever ways

## Values
- **Convention > configuration**
- **Provide sensible defaults for working IN THE CURRENT directory, not browsing ALL of Backlog.**
  - If a user is in `cacoo-blog` repository, don't provide confusing ways to create a PR for `nulab-website`.
- **Prioritize CREATE/READ/DELETE type operations (for issues, pull requests, etc)** over modification/editing operations (which may be best suited for a GUI)

## Minimum Viability Goals (for v1.0)
- [ ] Quick updating of issue status via `blg issue status [open|inprogress|resolved|closed]`
- Support quick creation/deletion of:
  - [x] pull requests
  - [ ] issues
- Support text-based browsing of:
  - [ ] issues
  - [ ] pull requests
  - [ ] notifications
  - [ ] milestone progress (calculate based on 'closed' issue percentage for current milestone)
- Allow users to easily spit out "quick links" to useful pages on Backlog:
  - [ ] Add `-o` flag to open quick links immediately, e.g. `blg pr all -o`
  - [x] User information & user activity page
  - [ ] 'All PullRequests' page for current space
  - [ ] 'All wiki' page for current space
  - [ ] 'All git repositories' page for current space
- [ ] Add footnotes to indicate where backlog-cli was used to avoid possible confusion
  - > This PR was created using backlog-cli. Notice an error? Tag @kevin in Typetalk

## Cool, non-priority features to try and implement:
- [ ] Support name-guessing for coworkers when assigning tasks or PRs, so that users don't have to memorize exact IDs (@whatever), since they can change
  - For example, if a coworker has a long, or hard-to-spell ID: `ryuzoyamamotosama`
  - Allow for something like: `blg pr --assignee="ryu yama"`
  - [Use string-matching libraries and only accept high-confidence matches](https://godoc.org/github.com/antzucaro/matchr)
- [ ] Add a `maintainer.md` document with a username/ID in each repository to allow for `blg pr assign-maintainer` to create a Pull Request and assign it to the maintainer automatically and easily send them a notification.

## Lots of work, but cool features:
- Support modification of wiki content in local editor, a la Github's Gist sync system
  - Something like: `blg wiki [pull|push]`
- Adding an `-i` flag for certain commands that will enable an interactive, ncurses-style interface for:
  - graphical browsing for browsing `README.md`s for all repositories

## What backlog-cli is NOT trying to be (at least not yet)
- Creating a superset of `git` itself (a la Github's `hub`)