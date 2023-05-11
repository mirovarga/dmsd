# DMSd

Turn files matching a glob into a DMS.

> **NB**: This project is a (working) prototype to see if it can be useful at
> all.

## Overview

DMSd is a command line tool that allows to tag files matching
a [glob pattern](https://en.wikipedia.org/wiki/Glob_(programming)).

It makes no changes to the files themselves, doesn't copy, move or rename them -
all data about the tagged files is stored in one file.

## Quick Start

1. Download a [release](https://github.com/mirovarga/dmsd/releases) and unpack
   it to a directory.
2. `cd` to the directory
3. Tag the files (with tags derived from the file system) 
   ```
   $ ./dmsd tag -A
   Tagged 3 files
   ```
4. List the tagged files
   ```
   $ ./dmsd list
   ```

## Documentation

> **Work in progress.**

Run `./dmsd -h` to see the integrated help.

### Tagging Files

TODO

### Listing Files

TODO

### Untagging Files

TODO

### Using Several Data Files

TODO
