package main

import (
	"github.com/morpheyesh/datadust-server/libs/cmd"
	"launchpad.net/gnuflag"
	"fmt"
)

type DDStart struct {
	fs *gnuflag.FlagSet
}

func (g *DDStart) Info() *cmd.Info {
	desc := `starts the dd API Server.`
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
		c.fs = gnuflag.NewFlagSet("dd", gnuflag.ExitOnError)
	}
	return c.fs
}
