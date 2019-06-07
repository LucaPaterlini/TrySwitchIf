package switchif

func TryIf(n int) string {
	if n > 0 {
		return "how"
	} else if n > 10 {
		return "are"
	} else if n > 20 {
		return "you"
	}
	return "None"
}

func TrySwitch(n int) string {
	switch {
	case n > 0:
		return "hello"
	case n > 10:
		return "how"
	case n > 20:
		return "are"
	default:
		return "None"
	}
}
