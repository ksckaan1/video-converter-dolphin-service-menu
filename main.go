package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	//fmt.Println("Dolphin Video Converter Service Started!")
	file := flag.String("f", "", "input file name")
	process := flag.String("p", "convert", "process type: convert | compress")
	outputFormat := flag.String("o", "", "output format: eg. mkv|mp4|mov|avi...")
	flag.Parse()
	parseDestination := strings.Split(*file, "/")
	dirArray := parseDestination[:len(parseDestination)-1]
	fileName := strings.Split(parseDestination[len(parseDestination)-1:][0], ".")[0]
	fileFormat := strings.Split(parseDestination[len(parseDestination)-1:][0], ".")[1]
	if *outputFormat == "" {
		*outputFormat = fileFormat
	}
	if *process == "convert" {
		convertVideo(fileName, fileFormat, *outputFormat, getDir(dirArray))
		fmt.Println("Converted Successfully!")
	}
	os.Exit(0)
}

func getDir(dirArray []string) string {
	dir := "/"
	for _, v := range dirArray {
		if v != "" {
			dir = dir + v + "/"
		}

	}

	return dir
}

func convertVideo(name, iFormat, oFormat, dir string) error {
	convertCommand := fmt.Sprintf("ffmpeg -i %v%v.%v %v%v.%v", dir, name, iFormat, dir, name, oFormat)
	cmd := exec.Command("konsole", "-e", convertCommand)
	cmd.Stdout = os.Stdout
	cmd.Run()

	return nil
}
