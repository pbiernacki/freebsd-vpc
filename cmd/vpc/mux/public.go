package mux

import (
	"github.com/joyent/freebsd-vpc/cmd/vpc/mux/connect"
	"github.com/joyent/freebsd-vpc/cmd/vpc/mux/create"
	"github.com/joyent/freebsd-vpc/cmd/vpc/mux/destroy"
	"github.com/joyent/freebsd-vpc/cmd/vpc/mux/disconnect"
	"github.com/joyent/freebsd-vpc/cmd/vpc/mux/fte"
	"github.com/joyent/freebsd-vpc/cmd/vpc/mux/listen"
	"github.com/joyent/freebsd-vpc/cmd/vpc/mux/show"
	"github.com/joyent/freebsd-vpc/internal/command"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

const cmdName = "mux"

var Cmd = &command.Command{
	Name: cmdName,

	Cobra: &cobra.Command{
		Use:   cmdName,
		Short: "VPC packet multiplexing configuration",
	},

	Setup: func(self *command.Command) error {
		subCommands := command.Commands{
			create.Cmd,
			connect.Cmd,
			disconnect.Cmd,
			destroy.Cmd,
			fte.Cmd,
			listen.Cmd,
			show.Cmd,
		}

		if err := self.Register(subCommands); err != nil {
			return errors.Wrapf(err, "unable to register sub-commands under %s", cmdName)
		}

		return nil
	},
}