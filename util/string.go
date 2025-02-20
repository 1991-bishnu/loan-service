package util

import (
	"fmt"
	"regexp"

	"github.com/google/uuid"
)

func GeneratePID(prefix string) (pid string) {
	// Generate a new UUID
	u := uuid.New()

	// Remove non-alphanumeric characters
	re := regexp.MustCompile("[^a-zA-Z0-9]+")
	pid = re.ReplaceAllString(u.String(), "")

	// Combine the prefix and the UUID
	pid = fmt.Sprintf("%s_%s", prefix, pid)

	// Ensure the PID is 16 characters long
	if len(pid) > 16 {
		pid = pid[:16]
	}

	return pid
}
