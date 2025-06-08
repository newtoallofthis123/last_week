package git

import (
	"fmt"
	"os/exec"
	"strings"
	"time"

	"github.com/newtoallofthis123/last_week/utils"
)

func getAuthor() string {
	cmd := exec.Command("git", "config", "user.name")

	outputBytes, err := cmd.Output()
	if err != nil {
		return ""
	}
	return strings.TrimSpace(string(outputBytes))
}

func executeCommand(timeFrame time.Time) (string, error) {

	author := getAuthor()
	if author == "" {
		return "", fmt.Errorf("author not found")
	}

	dateFormat := "Mon Jan 2 15:04:05 2006 -0700"
	formattedDate := timeFrame.Format(dateFormat)
	cmd := exec.Command("git", "--no-pager", "log", "--since=\""+formattedDate+"\"", "--author="+author)

	outputBytes, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}

	return string(outputBytes), nil
}

func GetOutput(timeQuery string) (string, error) {
	date, err := utils.ParseDate(timeQuery)
	if err != nil {
		return "", err
	}

	return executeCommand(date)
}
