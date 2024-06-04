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
	"PWD":  func(comm commands.BaseCommand) commands.Command { return commands.PwdCommand(comm) },
	"CWD":  func(comm commands.BaseCommand) commands.Command { return commands.CwdCommand(comm) },
}
