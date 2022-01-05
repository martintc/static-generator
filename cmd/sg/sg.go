package main

import (
	"flag"
	"fmt"
	"os"

	lex "github.com/advancebsd/ianus/markdownLexer"
	gemtextRender "github.com/advancebsd/ianus/gemtextRender"
)

func main() {
	helpPtr := flag.Bool("h", false, "help")
	srcPtr := flag.String("i", "", "Source file to convert")
	destPtr := flag.String("o", "", "Destination file for converted file")

	flag.Parse()

	if *helpPtr {
		fmt.Println("Static generator usage")
		fmt.Println("sg -i <input_file> -o <destination file>")
		os.Exit(0)
	}
	if *srcPtr == "" {
		panic("No source file was given")
	}
	if *destPtr == "" {
		panic("No destination file was given")
	}

	source, err := os.ReadFile(*srcPtr)
	if err != nil {
		fmt.Printf("%s is not a valid file!", *srcPtr)
		os.Exit(0)
	}

	fmt.Printf("Tokenizing %s\n", *srcPtr)

	var lexer lex.Lexer
	lexer.InitializeLexer(string(source))
	var token lex.Token
	var tokens []lex.Token
	token = lexer.NextToken()
	for token.Type != lex.EOF {
		tokens = append(tokens, token)
		token = lexer.NextToken()
	}

	fmt.Printf("Rendering %s into %s\n", *srcPtr, *destPtr)

	var g gemtextRender.GemtextRender
	g.InitializeGemtextRender(tokens)
	gemtext, err := g.RenderDocument()
	if err != nil {
		fmt.Println("Could not render the request document to Gemtext")
		os.Exit(1)
	}
	er := os.WriteFile(*destPtr, []byte(gemtext), 0644)
	if er != nil {
		fmt.Println("Could not output the file")
	}

	fmt.Println("Render complete")
}
