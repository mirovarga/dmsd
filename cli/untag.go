package cli

import (
	"fmt"

	"github.com/mirovarga/dmsd/lib"
)

type UntagCommand struct {
	Glob        string   `arg:"positional" help:"include files matching the glob [default: **]" default:"**"`
	ExcludeGlob string   `arg:"--exclude,-E" help:"exclude files matching the glob" placeholder:"GLOB"`
	Tags        lib.Tags `arg:"--tag,-t,separate" help:"remove the tag in the format name[:value]" placeholder:"TAG"`
	AutoTags    bool     `arg:"--auto-tags,-A" help:"remove tags derived from the filesystem"`
	DryRun      bool     `arg:"--dry-run,-D" help:"don't make any changes"`
	Verbose     bool     `arg:"-v" help:"show as much info as available"`
}

func (cmd UntagCommand) Run(db lib.DB) error {
	items, err := db.All()
	if err != nil {
		return err
	}

	items = items.FilterByMatchingGlob(cmd.Glob)
	if cmd.ExcludeGlob != "" {
		items = items.FilterByNotMatchingGlob(cmd.ExcludeGlob)
	}
	if cmd.Tags != nil {
		items.Untag(cmd.Tags...)
	}
	if cmd.AutoTags {
		items.UntagAuto()
	}
	if cmd.DryRun {
		fmt.Printf("Would untag %d files\n", len(items))
		if cmd.Verbose {
			printFullPaths(items)
		} else {
			fmt.Printf("Hint: Use --verbose to list the files\n")
		}
		return nil
	}

	err = db.Store(items, true)
	if err != nil {
		return err
	}

	fmt.Printf("Untagged %d files\n", len(items))
	if cmd.Verbose {
		printFullPaths(items)
	}

	return nil
}
