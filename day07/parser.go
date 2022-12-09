package main

func traceCommands(commands []Command) *Directory {
	dirStack := &Stack{}
	for _, cmd := range commands {
		switch cmd.(type) {
		case *ChangeDirectory:
			handleCdCommand(cmd.(*ChangeDirectory), dirStack)
		case *ListDirectory:
			handleLsCommand(cmd.(*ListDirectory), dirStack)
		}
	}
	return dirStack.data[0]
}

func handleLsCommand(ls *ListDirectory, dirStack *Stack) {
	wd := dirStack.Peek()
	for _, out := range ls.Output {
		wd.Children = append(wd.Children, parse(out))
	}
}

func handleCdCommand(cd *ChangeDirectory, dirStack *Stack) {
	wd := dirStack.Peek()
	if cd.Arg == ".." {
		dirStack.Pop()
	} else {
		d := locateDirectoryToPush(cd, wd)
		if d == nil {
			panic(cd.Arg)
		}
		dirStack.Push(d)
	}
}

func locateDirectoryToPush(cd *ChangeDirectory, wd *Directory) *Directory {
	if wd == nil { // root dir
		return &Directory{DirName: cd.Arg}
	} else {
		return wd.FindDirectory(cd.Arg)
	}
}
