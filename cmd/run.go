package main

import (
	"github.com/megamsys/libgo/cmd"
	"launchpad.net/gnuflag"
	"fmt"
)

type DDStart struct {
	fs *gnuflag.FlagSet
}

func (g *DDStart) Info() *cmd.Info {
	desc := `starts the datadust API Server.`
	return &cmd.Info{
		Name:    "start",
		Usage:   `start`,
		Desc:    desc,
		MinArgs: 0,
	}
}

func (c *DDStart) Run(context *cmd.Context, client *cmd.Client) error {

	serverRun()
	return nil
}

func (c *DDStart) Flags() *gnuflag.FlagSet {
	if c.fs == nil {
		c.fs = gnuflag.NewFlagSet("datadust", gnuflag.ExitOnError)
	}
	return c.fs
}
