package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/briandowns/spinner"
	"github.com/newtoallofthis123/last_week/git"
	"github.com/newtoallofthis123/last_week/model"
	"github.com/newtoallofthis123/last_week/utils"
)

var (
	timeQuery  = flag.String("time", "last week", "The time frame to query for")
	localStore = flag.Bool("local", false, "Use local store")
	submodules = flag.Bool("submodules", false, "Include submodules")
	pipe       = flag.Bool("pipe", false, "Pipe the output to a file")
)

func main() {
	flag.Parse()
	s := spinner.New(spinner.CharSets[11], 100*time.Millisecond)
	s.Suffix = " Generating report..."
	if !*pipe {
		s.Start()
	}
	env := utils.ReadEnv(*localStore)
	m, err := model.NewGeminiModel(context.Background(), &env)
	if err != nil {
		log.Fatal(err)
	}

	gitLogFormat, err := git.GetGitLogFormat(*timeQuery, *submodules)
	if err != nil {
		log.Fatal(err)
	}

	json, err := json.Marshal(gitLogFormat)
	if err != nil {
		log.Fatal(err)
	}

	response, err := m.GenerateResponse(string(json), model.SlidesSystemPrompt())
	if err != nil {
		log.Fatal(err)
	}

	if !*pipe {
		s.Stop()
	}

	fmt.Println(response)
}
