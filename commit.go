package main

import (
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/AlecAivazis/survey"
)

var qs = []*survey.Question{
	{
		Name: "tag",
		Prompt: &survey.Select{
			Message: "Tag",
			Options: []string{"feat", "fix", "style", "refactor", "test", "docs", "clean", "perf", "chore"},
		},
	},
	{
		Name:     "subject",
		Validate: survey.Required,
		Prompt:   &survey.Input{Message: "Subject"},
	},
	{
		Name:   "issue",
		Prompt: &survey.Input{Message: "Issue"},
	},
	{
		Name:   "scope",
		Prompt: &survey.Input{Message: "Scope"},
	},
}

type answersType struct {
	Issue   string
	Tag     string
	Scope   string
	Subject string
}

func main() {
	answers := takeSurvey()
	message := formatCommitMessage(answers)
	commit(message)
}

func takeSurvey() answersType {
	answers := answersType{}

	err := survey.Ask(qs, &answers)
	if err != nil {
		log.Fatal(err)
	}

	return answers
}

func formatCommitMessage(answers answersType) string {
	var message strings.Builder

	if answers.Issue != "" {
		message.WriteString(answers.Issue + " ")
	}

	message.WriteString(answers.Tag)

	if answers.Scope != "" {
		message.WriteString("(" + answers.Scope + ")")
	}

	message.WriteString(": " + answers.Subject)

	return message.String()
}

func commit(message string) {
	cmd := exec.Command("git", "commit", "-m", message)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
	}
}
