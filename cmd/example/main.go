// Package main implements the example CLI application.
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var (
	// Version is the application version.
	Version = "dev"
	// GitCommit is the git commit hash.
	GitCommit = "unknown"
	// BuildDate is the build date.
	BuildDate = "unknown"
)

func main() {
	var (
		versionFlag = flag.Bool("version", false, "Show version information")
		helpFlag    = flag.Bool("help", false, "Show help message")
	)

	flag.Parse()

	if *helpFlag {
		printUsage()
		os.Exit(0)
	}

	if *versionFlag {
		printVersion()
		os.Exit(0)
	}

	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	fmt.Println("Hello from tmpl!")
	return nil
}

func printUsage() {
	fmt.Fprintf(os.Stderr, "Usage: %s [options]\n", os.Args[0])
	fmt.Fprintln(os.Stderr, "\nOptions:")
	flag.PrintDefaults()
}

func printVersion() {
	fmt.Printf("Version:    %s\n", Version)
	fmt.Printf("Git Commit: %s\n", GitCommit)
	fmt.Printf("Build Date: %s\n", BuildDate)
}
