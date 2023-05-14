package cli

import (
	"fmt"

	"github.com/mirovarga/dmsd/lib"
)

type TagCommand struct {
	Glob        string   `arg:"positional" help:"include files matching the glob [default: *]" default:"*"`
	ExcludeGlob string   `arg:"--exclude,-E" help:"exclude files matching the glob" placeholder:"GLOB"`
	Tags        lib.Tags `arg:"--tag,-t,separate" help:"tag in the format name[:value]" placeholder:"TAG"`
	AutoTags    bool     `arg:"--auto-tags,-A" help:"add tags derived from the file system"`
	DryRun      bool     `arg:"--dry-run,-D" help:"don't make any changes"`
	Verbose     bool     `arg:"-v" help:"show as much info as available"`
}

func (cmd TagCommand) Run(db lib.DB) error {
	items, err := lib.NewItemsFromGlob(cmd.Glob)
	if err != nil {
		return err
	}

	if cmd.ExcludeGlob != "" {
		items = items.FilterByNotMatchingGlob(cmd.ExcludeGlob)
	}
	if cmd.Tags != nil {
		items.Tag(cmd.Tags...)
	}
	if cmd.AutoTags {
		err := items.AutoTag()
		if err != nil {
			return err
		}
	}
	if cmd.DryRun {
		fmt.Printf("Would tag %d files\n", len(items))
		if cmd.Verbose {
			printFullPaths(items)
		} else {
			fmt.Println("Hint: Use --verbose to list the files")
		}
		return nil
	}

	err = db.Store(items, false)
	if err != nil {
		return err
	}

	fmt.Printf("Tagged %d files\n", len(items))
	if cmd.Verbose {
		printFullPaths(items)
	}

	return nil
}
