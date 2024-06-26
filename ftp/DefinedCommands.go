package FtpServer

import (
	commands "ftp/ftp/command"
)

type Command404Error struct {
	CommName string
}

func (e Command404Error) Error() string {
	return "Command " + e.CommName + " not implemented!"
}

func ResolveCommand(bc commands.BaseCommand) (commands.Command, error) {
	getCommandFunc, ok := AVAILABLE_COMMANDS[bc.CommandName]
	if ok {
		return getCommandFunc(bc), nil
	}
	return nil, Command404Error{bc.CommandName}
}

var AVAILABLE_COMMANDS = map[string]func(comm commands.BaseCommand) commands.Command{
	"NOOP": func(comm commands.BaseCommand) commands.Command { return commands.NoopCommand(comm) },
	"USER": func(comm commands.BaseCommand) commands.Command { return commands.UserCommand(comm) },
	"TYPE": func(comm commands.BaseCommand) commands.Command { return commands.TypeCommand(comm) },

	"PWD":  func(comm commands.BaseCommand) commands.Command { return commands.PwdCommand(comm) },
	"CWD":  func(comm commands.BaseCommand) commands.Command { return commands.CwdCommand(comm) },
	"MKD":  func(comm commands.BaseCommand) commands.Command { return commands.MkdCommand(comm) },
	"DELE": func(comm commands.BaseCommand) commands.Command { return commands.DeleCommand(comm) },
	"RMD":  func(comm commands.BaseCommand) commands.Command { return commands.RmdCommand(comm) },

	"PASV": func(comm commands.BaseCommand) commands.Command { return commands.PasvCommand(comm) },
	"PORT": func(comm commands.BaseCommand) commands.Command { return commands.PortCommand(comm) },

	"LIST": func(comm commands.BaseCommand) commands.Command { return commands.ListCommand(comm) },
	"STOR": func(comm commands.BaseCommand) commands.Command { return commands.StorCommand(comm) },
	"RETR": func(comm commands.BaseCommand) commands.Command { return commands.RetrCommand(comm) },

	"RNFR": func(comm commands.BaseCommand) commands.Command { return commands.RnfrCommand(comm) },
	"RNTO": func(comm commands.BaseCommand) commands.Command { return commands.RntoCommand(comm) },
}
