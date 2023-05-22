package cli

import (
	"fmt"
	"os"

	"github.com/alexflint/go-arg"

	"github.com/mirovarga/dmsd/lib"
)

type args struct {
	DataFile string        `arg:"--data-file,-F" default:"dmsd.db" help:"data file" placeholder:"FILE"`
	Tag      *TagCommand   `arg:"subcommand:tag" help:"tag files"`
	List     *ListCommand  `arg:"subcommand:list" help:"list tagged files"`
	Untag    *UntagCommand `arg:"subcommand:untag" help:"untag files"`
}

func (args) Version() string {
	return "DMSd (v0.3.1): A tool for tagging files and querying for files with those tags [github.com/mirovarga/dmsd]\n"
}

type Command interface {
	Run(db lib.DB) error
}

func Run() {
	var args args
	parser := arg.MustParse(&args)

	var cmd Command
	switch {
	case args.Tag != nil:
		if args.Tag.Tags == nil && !args.Tag.AutoTags {
			parser.Fail("--tag, --auto-tags or both required")
		}
		cmd = args.Tag
	case args.Untag != nil:
		cmd = args.Untag
	case args.List != nil:
		cmd = args.List
	}

	if cmd == nil {
		parser.WriteHelp(os.Stdout)
		return
	}

	err := cmd.Run(lib.NewOrDefaultDB(args.DataFile))
	if err != nil {
		fmt.Println("ERROR: ", err)
	}
}

func printFullPaths(items lib.Items) {
	for _, i := range items {
		fmt.Println(i.FullPath)
	}
}
