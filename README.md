# Backlog CLI

A CLI tool to interact with [Backlog](backlogtool.com), a project management tool by [Nulab, Inc](https://nulab-inc.com).

## Setup
1. Get an api key here: https://nulab.backlog.jp/EditApiSettings.action 
  - For "memo" you can simply put 'backlog-cli' or whatever you'd like.
2. `make init` will copy the sample configuration file to `$HOME/backlog-config.yaml`
3. Add the `API_KEY` to the YAML file.
4. Add the `BASEURL`. This will be the link you visit to use Backlog, e.g: `https://yourbacklogspace.nulab.com`
  - :rotating_light: Do not use a trailing slash! :rotating_light:
5. `make` will build the project and add the `blg` executable to `usr/local/bin` directly.
  - Ensure that `/usr/local/bin` is in your `$PATH`
  - To add elsewhere, simply do: `go build -o /your/preferred/path/blg *.go`

## Usage

```bash
# Test to see executable is found with:
blg --help
# Test Backlog API_KEY is set up properly with:
blg me
```

## Misc Resources

- To check values and goals, look at value-goals.md
- [Issues @ Github](https://github.com/aflashyrhetoric/backlog-cli/issues)
- [Kanban board](https://github.com/aflashyrhetoric/backlog-cli/projects/1)
