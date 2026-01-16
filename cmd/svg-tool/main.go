package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/pauloappbr/svg-tool/pkg/converter"
)

func main() {
	log.SetFlags(0)
	log.SetPrefix("error: ")

	// Flag definitions
	inputPath := flag.String("file", "", "Input SVG file path (required)")
	outputDir := flag.String("dir", ".", "Output directory for generated files")
	sizesList := flag.String("sizes", "", "Comma-separated list of custom sizes (e.g., 16,32,512). If empty, uses web defaults.")
	generateIco := flag.Bool("ico", true, "Generate favicon.ico file")

	// Custom usage message
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of svg-tool:\n")
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\nExample:\n  svg-tool -file logo.svg -dir assets/img\n")
	}

	flag.Parse()

	// Validation
	if *inputPath == "" {
		log.Println("input file is required.")
		flag.Usage()
		os.Exit(1)
	}

	// Define which files to generate
	var specs []converter.OutputSpec
	if *sizesList != "" {
		parts := strings.Split(*sizesList, ",")
		for _, p := range parts {
			s, err := strconv.Atoi(strings.TrimSpace(p))
			if err == nil {
				specs = append(specs, converter.OutputSpec{
					Name: fmt.Sprintf("icon-%d.png", s),
					Size: s,
				})
			}
		}
	} else {
		specs = converter.DefaultWebAssets()
	}

	fmt.Printf("Processing '%s'...\n", *inputPath)

	// Execute conversion
	err := converter.ProcessSVG(*inputPath, *outputDir, specs, *generateIco)
	if err != nil {
		log.Fatalf("failed to process SVG: %v", err)
	}

	fmt.Println("Conversion completed successfully.")
}
