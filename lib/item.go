package lib

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/bmatcuk/doublestar/v4"
)

type Item struct {
	FullPath string `json:"full-path"`
	Tags     Tags   `json:"tags,omitempty"`
}

func (i *Item) FullPathMatches(glob string) bool {
	ok, _ := doublestar.PathMatch(glob, i.FullPath)
	return ok
}

func (i *Item) HasTagWithName(name string) bool {
	return i.Tags.ContainsWithName(name)
}

func (i *Item) HasTagWithValue(value string) bool {
	return i.Tags.ContainsWithValue(value)
}

func (i *Item) HasTag(tag Tag) bool {
	return i.Tags.Contains(tag)
}

func (i *Item) Tag(tags ...Tag) {
	for _, t := range tags {
		if i.Tags.Contains(t) {
			continue
		}
		i.Tags = append(i.Tags, t)
	}
}

func (i *Item) AutoTag() error {
	tags, err := NewTagsFromFile(i.FullPath)
	if err != nil {
		return err
	}

	i.Tag(tags...)
	return nil
}

func (i *Item) Untag(tags ...Tag) {
	i.Tags = i.Tags.Remove(tags...)
}

func (i *Item) UntagAuto() {
	i.Untag(i.Tags.AutoTags()...)
}

func NewItemFromFile(path string) (*Item, error) {
	fullPath, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}

	info, err := os.Stat(fullPath)
	if err != nil {
		return nil, err
	}
	if info.IsDir() {
		return nil, fmt.Errorf("not a file: %s", fullPath)
	}

	return &Item{
		FullPath: fullPath,
		Tags:     Tags{},
	}, nil
}

type Items map[string]*Item

func (is Items) FilterByMatchingGlob(glob string) Items {
	return is.filterBy(func(i *Item) bool {
		return i.FullPathMatches(glob)
	})
}

func (is Items) FilterByNotMatchingGlob(glob string) Items {
	return is.filterBy(func(i *Item) bool {
		return !i.FullPathMatches(glob)
	})
}

func (is Items) FilterByTagWithName(name string) Items {
	return is.filterBy(func(i *Item) bool {
		return i.HasTagWithName(name)
	})
}

func (is Items) FilterByTagWithValue(value string) Items {
	return is.filterBy(func(i *Item) bool {
		return i.HasTagWithValue(value)
	})
}

func (is Items) FilterByTags(tags ...Tag) Items {
	return is.filterBy(func(i *Item) bool {
		for _, t := range tags {
			if !i.HasTag(t) {
				return false
			}
		}
		return true
	})
}

func (is Items) filterBy(fn func(i *Item) bool) Items {
	items := Items{}
	for _, i := range is {
		if fn(i) {
			items[i.FullPath] = i
		}
	}
	return items
}

func (is Items) Tag(tags ...Tag) {
	for _, i := range is {
		i.Tag(tags...)
	}
}

func (is Items) AutoTag() error {
	for _, i := range is {
		err := i.AutoTag()
		if err != nil {
			return err
		}
	}
	return nil
}

func (is Items) Untag(tags ...Tag) {
	for _, i := range is {
		i.Untag(tags...)
	}
}

func (is Items) UntagAuto() {
	for _, i := range is {
		i.UntagAuto()
	}
}

func NewItemsFromGlob(glob string) (Items, error) {
	files, err := doublestar.FilepathGlob(glob, doublestar.WithFilesOnly())
	if err != nil {
		return nil, err
	}
	if files == nil {
		return Items{}, nil
	}

	items := Items{}
	for _, file := range files {
		item, err := NewItemFromFile(file)
		if err != nil {
			return nil, err
		}

		items[item.FullPath] = item
	}
	return items, nil
}
