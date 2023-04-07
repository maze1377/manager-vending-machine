package main

import (
	"fmt"
	_ "net/http/pprof"
	"os"
)

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
