# What is this Project?

iTerm has a large repository of [themes and ports to other programs](https://github.com/mbadolato/iTerm2-Color-Schemes).

Some of the most common macOS terminal themes live in the iTerm native format, a `.itermcolors` format.
This format is actuall an Apple `.plist` file (which itself is a wrapped XML format), and the colors within these files
are normalized to be between 0 and 1, despite RGB values being between 0 and 255. It's annoying to extract colors from
an `.itermcolors` file, so I wrote this tool to convert a given theme file into the appropriate hex codes for you.

It's worth noting that the iTerm project has their own python tools, I just wanted to build this using Golang to practice my
data marshalling/unmarshalling, particularly with XML (or a deriavtive, in this case).

I expect this is not a comprehensive solution, nor is it intended to be a replacement of the iTerm tools - this is just a quick
Go package to meet my needs. I may or may not update this in the future as I use it more or encounter issues.

# Dependencies

## plist package

This tool relies heavily on the [plist package](https://github.com/DHowett/go-plist) by DHowett. This library makes parsing `.plist` files as
easy as the encode/json package.

## Go

Thi package was written & built with Go v1.23

# Installation

To install this package, clone the repo:

```bash
git clone https://github.com/taylrfnt/iterm2hex.git

```

Build the package:

```bash
cd iterm2hex && go build
```

Symlink or move the built package into your system's bin (or wherever you want to manage it that is in your `$PATH`).

# Usage

To use this package, simply invoke it by its name and include a file path:

```bash
iterm2hex ${PATH_TO_FILE}
```
