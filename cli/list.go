package cli

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/gosuri/uitable"

	"github.com/mirovarga/dmsd/lib"
)

type ListCommand struct {
	Glob        string   `arg:"positional" help:"include files matching the glob [default: **]" default:"**"`
	ExcludeGlob string   `arg:"--exclude,-E" help:"exclude files matching the glob" placeholder:"GLOB"`
	Tags        lib.Tags `arg:"--tag,-t,separate" help:"list only files with the tag in the format name[:value]" placeholder:"TAG"`
	Format      string   `arg:"-f" help:"table | card | json" default:"table"`
}

func (cmd ListCommand) Run(db lib.DB) error {
	items, err := db.All()
	if err != nil {
		return err
	}

	if cmd.Glob != "" {
		items = items.FilterByMatchingGlob(cmd.Glob)
	}
	if cmd.ExcludeGlob != "" {
		items = items.FilterByNotMatchingGlob(cmd.ExcludeGlob)
	}
	if cmd.Tags != nil {
		items = items.FilterByTags(cmd.Tags...)
	}

	switch cmd.Format {
	case "table":
		if len(items) == 0 {
			fmt.Printf("No matching files found\n")
			return nil
		}

		t := uitable.New()
		t.Wrap = true
		t.AddRow("FULL PATH", "TAGS")

		for _, i := range items {
			t.AddRow(i.FullPath, strings.Join(i.Tags.Strings(), "\n"))
		}

		fmt.Println(t)
	case "card":
		if len(items) == 0 {
			fmt.Printf("No matching files found\n")
			return nil
		}

		t := uitable.New()
		t.Wrap = true

		for _, i := range items {
			t.AddRow("Full Path: ", i.FullPath)
			t.AddRow("Tags: ", strings.Join(i.Tags.Strings(), "\n"))
			t.AddRow("")
		}

		fmt.Println(t)
	case "json":
		jsonItems := []*lib.Item{}
		for _, i := range items {
			jsonItems = append(jsonItems, i)
		}
		return json.NewEncoder(os.Stdout).Encode(jsonItems)
	}

	return nil
}
