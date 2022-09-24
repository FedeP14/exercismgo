package techpalace

import "strings"

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	tmpString := strings.Repeat("*", numStarsPerLine)

	return tmpString + "\n" + welcomeMsg + "\n" + tmpString
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	oldMsg = strings.Replace(oldMsg, "*", "", -1)
	oldMsg = strings.TrimSpace(oldMsg)
	return oldMsg
}
