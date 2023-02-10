/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"log"
	"os"

	"github.com/rrd1986/go-cli-app/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
