package git

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type GitLogFormat struct {
	Author string            `json:"author"`
	Repo   map[string]string `json:"repo"`
}

func getSubModules() []string {
	cmd := exec.Command("sh", "-c", "git submodule status | awk '{print $2}'")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil
	}
	return strings.Split(string(output), "\n")
}

func GetGitLogFormat(timeQuery string, submodules bool) (GitLogFormat, error) {
	var globalErr error
	repoMap := make(map[string]string)

	currentDir, err := os.Getwd()
	if err != nil {
		return GitLogFormat{}, err
	}

	repoMap[currentDir], err = GetOutput(timeQuery)
	if err != nil {
		return GitLogFormat{}, err
	}

	if submodules {
		for _, repo := range getSubModules() {
			os.Chdir(repo)
			repoOutput, err := GetOutput(timeQuery)
			if err != nil {
				globalErr = fmt.Errorf("%s \n error getting git log for %s: %w", globalErr, repo, err)
			} else {
				repoMap[repo] = repoOutput
			}
			os.Chdir("..")
		}
	}

	if globalErr != nil {
		return GitLogFormat{}, globalErr
	}

	return GitLogFormat{
		Author: getAuthor(),
		Repo:   repoMap,
	}, nil
}
