package logs

import (
	"math"
	"strings"
)

// Application identifies the application emitting the given log.
func Application(log string) string {
	indexReccomendation := strings.IndexRune(log, '❗')
	indexSearch := strings.IndexRune(log, '🔍')
	indexWeather := strings.IndexRune(log, '☀')

	if indexReccomendation == -1 && indexSearch == -1 && indexWeather == -1 {
		return "default"
	}

	if indexSearch == -1 {
		indexSearch = len(log)
	}

	if indexWeather == -1 {
		indexWeather = len(log)
	}

	if indexReccomendation == -1 {
		indexReccomendation = len(log)
	}

	min := math.Min(float64(indexReccomendation), float64(indexSearch))

	min = math.Min(min, float64(indexWeather))

	switch {
	case min == float64(indexReccomendation):
		return "recommendation"
	case min == float64(indexWeather):
		return "weather"
	case min == float64(indexSearch):
		return "search"
	}

	return ""
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
