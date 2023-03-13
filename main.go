package main

import (
	"fmt"
	"os"
)

var (
	gopath = os.Getenv("GOPATH")
)

func printHelp() {
	fmt.Println("bAsk - bing chat for the CLI - version v0.2")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("\tbask [args]")
	fmt.Println()
	fmt.Println("Arguments:")
	fmt.Println("\t-q [query]  -> sends your query and gets the response from the Bing chat AI")
	fmt.Println("\t-c [cookie] -> sets your Bing cookie")
	fmt.Println("\t-h          -> displays this message")

	os.Exit(0)
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		printHelp()
	}

	switch args[0] {
	case "-q":
		if len(args) < 2 {
			fmt.Println("bAsk - No query provided!")
		} else {
			if cookieExists() {
				search(args[1])
			} else {
				fmt.Println("bAsk - No cookie set!")
				fmt.Println("Please set your cookie using \"bask -c [cookie]\", then try again.")
			}
		}
	case "-c":
		if len(args) < 2 {
			fmt.Println("bAsk - No cookie provided!")
		} else {
			if gopath != "" {
				setCookie(args[1])
			} else {
				fmt.Println("bAsk - GOPATH environment variable is undefined!")
				fmt.Println("Please verify your GOPATH environment variable is set properly.")
			}
		}
	case "-h":
		printHelp()
	default:
		printHelp()
	}
}
