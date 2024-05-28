package techpalace

import (
	"fmt"
	"strings"
)

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	welcome_text := strings.ToUpper(customer)
	return fmt.Sprintf("%s %s","Welcome to the Tech Palace,", welcome_text)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	border := strings.Repeat("*", numStarsPerLine)
	return fmt.Sprintf("%s\n%s\n%s", border, welcomeMsg, border)
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	cleanedMessage := strings.ReplaceAll(oldMsg, "*", "")
	cleanedMessage = strings.TrimSpace(cleanedMessage)
	return cleanedMessage
}
