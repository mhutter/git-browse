# git-browse

Open the current repository in a browser.

## Usage

    # cd into any git repository
    git browse
    # opens a browser

## Limitations

At the moment:
- `git-browse` uses the URL of the `origin` remote
- only SSH remotes are supported

## Installation

    go get -u github.com/mhutter/git-browse
