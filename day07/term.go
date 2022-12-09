package main

import "strings"

const CommandPrefix = "$ "

type Command interface {
}

type ChangeDirectory struct {
	Arg string
}

type ListDirectory struct {
	Output []string
}

func ParseCommands(lines []string) []Command {
	result := make([]Command, 0)
	cmd := parseCommand(lines[0])
	for i := 1; i < len(lines); i++ {
		line := lines[i]
		if isCommand(line) {
			result = append(result, cmd)
			cmd = parseCommand(line)
		} else {
			ls := cmd.(*ListDirectory)
			ls.Output = append(ls.Output, line)
		}
	}
	result = append(result, cmd)
	return result
}

func parseCommand(s string) Command {
	cmd := strings.Split(strings.TrimLeft(s, CommandPrefix), " ")
	switch cmd[0] {
	case "cd":
		return &ChangeDirectory{
			Arg: cmd[1],
		}
	case "ls":
		return &ListDirectory{}
	}
	return cmd
}

func isCommand(s string) bool {
	return strings.HasPrefix(s, CommandPrefix)
}
