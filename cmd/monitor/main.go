package main

import (
	"fmt"
	"os"
	"runtime"
	"text/tabwriter"
)

func main() {
	fmt.Println("This is the simple temperature monitor")
	fmt.Println("==============")
	fmt.Println("Not implemented yet")
	printEnv()
}

func printEnv() {
	fmt.Println("Current environment is:")
	tw := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.Debug)
	fmt.Fprintln(tw, "GOOS\t", runtime.GOOS)
	fmt.Fprintln(tw, "GOARCH\t", runtime.GOARCH)
	tw.Flush()
}
