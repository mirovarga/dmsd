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
	Tags   lib.Tags `arg:"positional" help:"list only files with the tag in the format name[:value]" placeholder:"TAG"`
	Format string   `arg:"-f" help:"text | json" default:"text"`
}

func (cmd ListCommand) Run(db lib.DB) error {
	items, err := db.All()
	if err != nil {
		return err
	}

	if cmd.Tags != nil {
		items = items.FilterByTags(cmd.Tags...)
	}

	switch cmd.Format {
	case "text":
		if len(items) == 0 {
			fmt.Printf("No matching files found\n")
			return nil
		}

		table := uitable.New()
		table.MaxColWidth = 80
		table.Wrap = true
		table.AddRow("FULL PATH", "TAGS")
		for _, i := range items {
			table.AddRow(i.FullPath, strings.Join(i.Tags.Strings(), ", "))
		}
		fmt.Println(table)
	case "json":
		jsonItems := []*lib.Item{}
		for _, i := range items {
			jsonItems = append(jsonItems, i)
		}
		return json.NewEncoder(os.Stdout).Encode(jsonItems)
	}

	return nil
}
