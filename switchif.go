package switchif


func TryIf(n int) string {
	if n > 20 {
		return "how"
	} else if n > 10 {
		return "are"
	} else if n > 0 {
		return "you"
	}
	return "None"
}

func TrySwitch(n int) string {
	switch {
	case n > 20:
		return "hello"
	case n > 10:
		return "how"
	case n > 0:
		return "are"
	default:
		return "None"
	}
}
