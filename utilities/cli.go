package utilities

// SeparateFlags splits a full set of arguments into a two string slices of just args and the flags
func SeparateFlags(args []string) (justArgs []string, flags []string) {
	for _, arg := range args {
		if arg[0] == '-' {
			flags = append(flags, arg)
		} else {
			justArgs = append(justArgs, arg)
		}
	}

	return
}
