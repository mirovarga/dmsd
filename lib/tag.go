package lib

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/gabriel-vasile/mimetype"
)

const (
	tagNameValueSeparator = ":"
	dmsdTagNamePrefix     = "dmsd."
)

type Tag string

func (t Tag) Name() string {
	return strings.Split(t.String(), tagNameValueSeparator)[0]
}

func (t Tag) Value() string {
	nameValue := strings.Split(t.String(), tagNameValueSeparator)
	return nameValue[len(nameValue)-1]
}

func (t Tag) IsDmsdTag() bool {
	return strings.HasPrefix(t.Name(), dmsdTagNamePrefix)
}

func (t Tag) String() string {
	return string(t)
}

func NewTagWithValue(name, value string) Tag {
	return Tag(name + tagNameValueSeparator + value)
}

func newDmsdTagWithValue(name, value string) Tag {
	return NewTagWithValue(dmsdTagNamePrefix+name, value)
}

type Tags []Tag

func (ts Tags) ContainsWithName(name string) bool {
	for _, t := range ts {
		if t.Name() == name {
			return true
		}
	}
	return false
}

func (ts Tags) ContainsWithValue(value string) bool {
	for _, t := range ts {
		if t.Value() == value {
			return true
		}
	}
	return false
}

func (ts Tags) Contains(tag Tag) bool {
	for _, t := range ts {
		if t.Name() == tag.Name() && t.Value() == tag.Value() {
			return true
		}
	}
	return false
}

func (ts Tags) AutoTags() (tags Tags) {
	for _, t := range ts {
		if t.IsDmsdTag() {
			tags = append(tags, t)
		}
	}
	return
}

func (ts Tags) Remove(tags ...Tag) (tgs Tags) {
	for _, t := range ts {
		if !Tags(tags).Contains(t) {
			tgs = append(tgs, t)
		}
	}
	return
}

func (ts Tags) Strings() []string {
	strs := []string{}
	for _, t := range ts {
		strs = append(strs, t.String())
	}
	return strs
}

func NewTagsFromFile(path string) (Tags, error) {
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

	var tags Tags

	dirs := strings.Split(filepath.Dir(fullPath), string(os.PathSeparator))[1:]
	for _, dir := range dirs {
		tags = append(tags, newDmsdTagWithValue("dir", dir))
	}

	filename := info.Name()
	tags = append(tags, newDmsdTagWithValue("filename", filename))

	ext := filepath.Ext(filename)
	tags = append(tags, newDmsdTagWithValue("ext", ext))

	filenameNoExt := strings.TrimSuffix(filename, ext)
	tags = append(tags, newDmsdTagWithValue("filename-no-ext", filenameNoExt))

	modified := info.ModTime().Format(time.RFC3339)
	tags = append(tags, newDmsdTagWithValue("modified", modified))

	size := info.Size()
	tags = append(tags, newDmsdTagWithValue("size", strconv.FormatInt(size, 10)))

	mime, err := mimetype.DetectFile(fullPath)
	if err != nil {
		return nil, err
	}
	tags = append(tags, newDmsdTagWithValue("mime", mime.String()))

	return tags, nil
}
