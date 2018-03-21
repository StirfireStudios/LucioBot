# LucioBot - Sound bot for your Slack Server!

This is a Slack bot. Made to play amusing sounds from a raspberry pi

# Developing

## ANTLR

There is an Antlr language definition, which you'll need to change to add commands. You'll need to regenerate the parsers (which are checked into the repo to make it easier to build)

## Prerequisites

You will need a [properly setup](https://golang.org/doc/install) Go development environment to use this.

## Getting code and all dependencies

Once you have your GOPATH setup as above you should:

  1. `go get -u github.com/nlopes/slack`
  2. `go get -u github.com/antlr/antlr4/runtime/Go/antlr`

once you have all those, you should create a config file from the `config.sample.json` template.
Ensure you have your slack Token and some audio files ready to play!

 1. `cd $GOPATH/src/github.com/StirfireStudios/LucioBot`
 2. `go run *.go`

