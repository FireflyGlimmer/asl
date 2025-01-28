package utils

func CmdArgs(args []string) {
	if len(args) == 1 {
		ShowUsage()
		return
	}

	switch args[1] {
	case "-t":
		Extractor(args[2], args[3])
	case "-p":
		GetImage(args[2], args[3])
	case "-v", "--version":
		ShowVersion()
	default:
		ShowUsage()
	}

}
