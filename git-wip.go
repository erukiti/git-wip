package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/github/hub/cmd"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	upstream := "origin"

	app := cli.NewApp()
	app.Name = "git-wip"
	app.Usage = "git-wip"
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name: "T",
		},
	}
	app.Action = func(c *cli.Context) {
		if len(c.Args()) < 1 {
			fmt.Println("git wip [-T] <branch>")
			os.Exit(1)
		}
		template := ""
		branch := c.Args()[0]
		if c.Bool("T") {
			template = exec(os.Getenv("GITWIP_TEMPLATE_CMD"))
		}

		file, err := createComment(template)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer os.Remove(file.Name())

		current := exec("git", "rev-parse", "--abbrev-ref", "HEAD")

		_ = exec("git", "checkout", "-b", branch)
		_ = exec("git", "commit", "--allow-empty", "-m", "[WIP] start")
		_ = exec("git", "push", upstream, branch)

		file.Seek(0, 0)
		_ = exec("hub", "pull-request", "--browse", "-F", file.Name(), "-b", strings.Trim(current, " \n"))
	}
	app.Run(os.Args)
}

func createComment(input string) (*os.File, error) {
	file, err := ioutil.TempFile(os.TempDir(), "prefix")

	file.WriteString(input)

	editor := cmd.New(os.Getenv("EDITOR"))
	editor.WithArg(file.Name())
	_, err = editor.CombinedOutput()

	if err != nil {
		return nil, err
	}

	file.Seek(0, 0)
	return file, err
}

func exec(command string, input ...string) string {
	git := cmd.New(command)

	for _, i := range input {
		git.WithArg(i)
	}

	output, err := git.CombinedOutput()
	if err != nil {
		fmt.Println(output)
		os.Exit(1)
	}

	return string(output)
}
