# DMSd

A command line tool for tagging files and querying for files with those tags.

> **NB:** This project is a (working) prototype to see if it can be useful at
> all.

## Table of contents

* [About the name](#about-the-name)
* [Overview](#overview)
* [Quick start](#quick-start)
* [Documentation](#documentation)
  * [Installation](#installation)
  * [Tagging files](#tagging-files)
    * [Examples](#examples)
  * [Listing files](#listing-files)
    * [Examples](#examples-1)
  * [Untagging files](#untagging-files)
    * [Examples](#examples-2)
  * [Using multiple data files](#using-multiple-data-files)

## About the name

`DMS` stands for
[Document Management System](https://en.wikipedia.org/wiki/Document_management_system),
`d` for directory.

The tool was originally meant to create a DMS from files in a directory, but 
now it actually uses [glob patterns](https://en.wikipedia.org/wiki/Glob_(programming))
to match the files to work on.
 
Also, after feedback from
[Reddit](https://www.reddit.com/r/golang/comments/13hqp3f/comment/jk7e65l)
and [HN](https://news.ycombinator.com/item?id=35944339), it seems that it's not
quite a DMS, but rather a file tagging system.

As I have currently no better name for it, it remains DMSd :)

## Overview

DMSd is a command line tool for tagging files and querying for files with those
tags.

It uses [glob patterns](https://en.wikipedia.org/wiki/Glob_(programming))
to match the files to work on.

It makes no changes to the files themselves, doesn't copy, move or rename them -
all data about the tagged files is stored in a data file.

## Quick start

1. Download a [release](https://github.com/mirovarga/dmsd/releases) and unpack
   it to a directory
2. `cd` to the directory
3. Tag all files in the current directory (with tags derived from the file system):
   ```
   ./dmsd tag --auto-tags
   ```
4. List the tagged files:
   ```
   ./dmsd list
   ```

## Documentation

### Installation

Download a [release](https://github.com/mirovarga/dmsd/releases) and unpack
it to a directory. That's all.

> You can add the directory to the `PATH` so you can run `dmsd` from any
> directory. All examples assume you have `dmsd` in your `PATH`.

Alternatively, if you have Go installed, run
`go install github.com/mirovarga/dmsd@latest`.

### Tagging files

> **TL;DR:** To tag files, use the `tag` command (run `dmsd tag -h` for help).

Tagging is attaching a name or a name and value to a file, like `invoice`
(a name) or `due:tomorrow` (a name and value - the `:` separates the tag name
from its value).

Tags can be attached to multiple files and each file can have multiple tags.

The tags can then be used to list files that match certain tags.

> Tagging a file makes no changes to the file itself, doesn't copy, move or
> rename it - all data about the tagged files is stored in the data file.

#### Examples

> Use the `--dry-run` option in all the examples below to avoid making any
> changes to the data file.

Tag all files in the current directory with tags derived from the file system,
like file name, extension, etc.:
```
dmsd tag --auto-tags
```

Tag all files in the current directory with the `a-tag` tag:
```
dmsd tag --tag a-tag
```

Tag all files in the current directory with tags `tag-one` and `tag:two`
(the `:` separates the tag name from its value):
```
dmsd tag --tag tag-one --tag tag:two
```
> You can combine tags derived from the file system with custom tags like this:
>
> `dmsd tag --auto-tags --tag a-tag`

Tag all Markdown files in the current directory with tags derived from the file
system:
```
dmsd tag '*.md' --auto-tags
```
> Note the single quotes - we need them to prevent the shell from interpreting
> the glob pattern and thus matching different files than expected.

Tag all files except Markdown ones in the current directory with tags derived
from the file system:
```
dmsd tag --auto-tags --exclude '*.md'
```

### Listing files

> **TL;DR:** To list tagged files, use the `list` command (run `dmsd list -h`
> for help).

#### Examples

> Use the `--format` option in all examples below to change the listing format.
> Supported formats are `text` (default) and `json`.

List all files:
```
dmsd list
```

List all files with the `a-tag` tag:
```
dmsd list a-tag
```

List all files with tags `tag-one` and `tag:two` (the `:` separates the tag name
from its value):
```
dmsd list tag-one tag:two
```

### Untagging files

> **TL;DR:** To untag files, use the `untag` command (run `dmsd untag -h` for
> help).

Untagging removes tags from already tagged files.

> Like tagging, untagging a file makes no changes to the file itself, doesn't
> copy, move or rename it - the tags are removed only from the data file.

#### Examples

> Use the `--dry-run` option in all the examples below to avoid making any
> changes to the data file.

Remove tags derived from the file system from all files:
```
dmsd untag --auto-tags
```

Remove the `a-tag` tag from all files:
```
dmsd untag --tag a-tag
```

Remove tags `tag-one` and `tag:two` from all files (the `:` separates the tag
name from its value):
```
dmsd untag --tag tag-one --tag tag:two
```
> You can combine tags derived from the file system with custom tags like this:
>
> `dmsd untag --auto-tags --tag a-tag`

Remove tags derived from the file system from all Markdown files:
```
dmsd untag '**/*.md' --auto-tags
```
> Note the double asterisk - we need it because files are indexed by their full
> paths so `*.md` wouldn't work as expected.

Remove tags derived from the file system from all files except Markdown ones:
```
dmsd untag --auto-tags --exclude '**/*.md'
```

### Using multiple data files

By default, information about tagged files is stored in the `dmsd.db` file in 
the current directory. You can override where the information is stored (or read
from) by specifying the `--data-file` option, like this:

```
dmsd --data-file overriden.db tag --auto-tags
```
