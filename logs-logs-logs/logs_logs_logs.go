package logs

// Application identifies the application emitting the given log.
func Application(log string) string {
	app := map[rune]string{
		'\u2757':     "recommendation",
		'\U0001f50d': "search",
		'\u2600':     "weather",
	}

	for _, char := range log {
		if name, ok := app[char]; ok {
			return name
		}
	}
	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	result := []rune(log)

	for i := 0; i < len(result); i++ {
		if result[i] == oldRune {
			result[i] = newRune
		}
	}

	s := string(result)

	return s

}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	runeLog := []rune(log)
	return len(runeLog) <= limit
}
