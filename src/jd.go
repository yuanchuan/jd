package main

import (
	"flag"
	"fmt"
	"os"

	"./storage"
	"github.com/mitchellh/go-homedir"
)

const (
	Version = "0.1.0"
	AppName = "jd"
	AppDesc = "Jump to a directory quickly by alias."
)

var home, err = homedir.Dir()
var db = storage.New(home + "/.jdstorage")

func action(f *flag.Flag) {
	arg := flag.Arg(0)
	if f.Name == "a" {
		curr, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		if len(arg) != 0 {
			curr = arg
		}
		db.Set(f.Value.String(), curr)
	}
	if f.Name == "d" {
		db.Del(f.Value.String())
	}
	if f.Name == "r" {
		db.Rename(f.Value.String(), arg)
	}
}

func main() {
	flag.Parse()
	flag.Visit(action)

	f := flag.NewFlagSet("f", flag.ContinueOnError)
	arg := flag.Arg(0)

	if flag.NFlag() == 0 && len(arg) == 0 {
		fmt.Fprint(f.Output(), db.GetAll(), "\n")
	}
	if flag.NFlag() == 0 && len(arg) != 0 {
		fmt.Println(db.Get(arg))
	}
}

func init() {
	flag.String("a", "", "Add alias")
	flag.String("d", "", "Delete an alias.")
	flag.String("r", "", "Rename an alias.")
}
