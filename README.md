# JD

A command line tool for jumping to a directory quickly.

```bash
# Add an alias for the current directory
$ jd -a name

# Jump to it from anywhere
$ jd name

# List all aliases
$ jd
```

## Setup

* Build with `go build` or just use make.

  ```bash
  make
  ```

* Create a function to your `.zshrc` or `.bashrc`.

  ```bash
  function jd {
     builtin cd "$(/path/to/the/project/bin/jd $@)"
  }
  ```

## Usage

```bash
$ jd --help

  Usage of you/path/to/bin/jd:
    -a string
          Add an alias
    -d string
          Delete an alias.
    -r string
          Rename an alias.
```

## Notes

This is my first Go program written in 2015.