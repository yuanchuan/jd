# JD

A command line tool for jumping to directories quickly using aliases.

## Features

- Add shortcuts to frequently visited directories
- Jump to any directory with a single command
- List all your saved aliases
- Rename and delete aliases as needed
- Simple JSON-based storage in your home directory

## Quick Start

```bash
# Add an alias for the current directory
jd -a work

# Jump to it from anywhere
jd work

# List all aliases
jd
```

## Installation

### Automated Install (Recommended)

```bash
curl -fsSL https://raw.githubusercontent.com/yuanchuan/jd/master/install.sh | bash
```

This script:
- Installs the jd binary
- Detects your shell (zsh/bash)
- Adds the required function to your shell config
- Shows next steps

### Manual Install

```bash
# Install binary
go install github.com/yuanchuan/jd@latest

# Add this function to your ~/.zshrc or ~/.bashrc
function jd {
    local path
    path=$(jd $@) || return
    builtin cd "$path"
}

# Reload your shell
source ~/.zshrc  # or ~/.bashrc
```

### Build from Source

```bash
git clone https://github.com/yuanchuan/jd.git
cd jd
go build
```

## Setup

Add the following function to your `.zshrc` or `.bashrc`:

```bash
function jd {
   builtin cd "$(jd $@)"
}
```

Then reload your shell configuration:

```bash
source ~/.zshrc  # or ~/.bashrc
```

## Usage

### Commands

```bash
# Add an alias for the current directory
jd -a <name>

# Add an alias for a specific path
jd -a <name> <path>

# Jump to a directory
jd <name>

# List all aliases
jd

# Delete an alias
jd -d <name>

# Rename an alias
jd -r <old-name> <new-name>
```

### Examples

```bash
# Save your project directory
cd ~/projects/my-app
jd -a myapp

# Save your config directory
jd -a config ~/.config

# Navigate to any saved directory
cd /somewhere/else
jd myapp        # Now in ~/projects/my-app

# List all your aliases
jd
# Output:
# {
#   "myapp": "/home/user/projects/my-app",
#   "config": "/home/user/.config"
# }

# Delete an alias you no longer need
jd -d config

# Rename an alias
jd -r myapp project
```

## How It Works

JD stores your directory aliases in a JSON file at `~/.jdstorage`. This makes it easy to backup or manually edit your aliases if needed.

## Notes

First written in 2015 as my first Go program. Updated for modern Go versions.
