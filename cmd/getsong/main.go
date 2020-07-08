package main

import (
	"fmt"
	"os"

	"github.com/prologic/getsong"
)

func main() {
	var (
		fname string
		err   error
	)

	opts := getsong.Options{
		ShowProgress: true,
	}

	if len(os.Args) == 2 {
		fname, err = getsong.GetSong(os.Args[1], opts)
	} else if len(os.Args) == 3 {
		fname, err = getsong.SearchSong(os.Args[1], os.Args[2], opts)
	} else {
		fmt.Println(`Usage:
		
getsong 'title of song' 'artist of song'

or:

getsong 'url'`)
		os.Exit(1)
	}

	if err == nil {
		fmt.Printf("Downloaded '%s'\n", fname)
	} else {
		fmt.Println(err)
	}
}
