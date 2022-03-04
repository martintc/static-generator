package main

import (
	"flag"
	"fmt"
	"os"

	gemtextRender "github.com/advancebsd/ianus/gemtextRender"
	lex "github.com/advancebsd/ianus/markdownLexer"
	htmlRender "github.com/advancebsd/ianus/htmlRender"
)

func main() {
	helpPtr := flag.Bool("h", false, "help")
	srcPtr := flag.String("i", "", "Source file to convert")
	destPtr := flag.String("o", "", "Destination file for converted file")
	gemPtr := flag.Bool("g", false, "gemtext output")
	htmlPtr := flag.Bool("p", false, "html output")

	flag.Parse()

	if *helpPtr {
		fmt.Println("Static generator usage")
		fmt.Println("sg -<format_specifier> -i <input_file> -o <destination file>")
		fmt.Println("Format specifiers:")
		fmt.Printf("\tp - Html output\n\tg - Gemtext output\n")
		fmt.Println("Both formats can be specified.")
		os.Exit(0)
	}
	if *srcPtr == "" {
		panic("No source file was given")
	}
	if *destPtr == "" {
		panic("No destination file was given")
	}
	if *gemPtr == false && *htmlPtr == false {
		panic("No mode was specified")
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

	if *gemPtr == true {
		var g gemtextRender.GemtextRender
		g.InitializeGemtextRender(tokens)
		gemtext, err := g.RenderDocument()
		if err != nil {
			fmt.Println("Could not render the request document to Gemtext")
			os.Exit(1)
		}
		er := os.WriteFile(*destPtr + ".gmi", []byte(gemtext), 0644)
		if er != nil {
			fmt.Println("Could not output the file to gemtext")
		}
	}

	if *htmlPtr == true {
		file, err := os.Create(*destPtr + ".html")
		defer file.Close()
		if err != nil {
			panic(err)
		}
		var h htmlRender.HtmlRender
		h.InitializeHtmlRender(tokens)
		html_text, err := h.RenderDocument()
		if err != nil {
			fmt.Println("Could not render the requested document to HTML")
			os.Exit(1)
		}
		_, err = file.Write([]byte("<html>\n"))
		if err != nil {
			fmt.Println("Could not output the file to Html")
		}
		_, err = file.Write([]byte(html_text))
		if err != nil {
			fmt.Println("Could not output the file to Html")
		}
		_, err = file.Write([]byte("</html>\n"))
		if err != nil {
			fmt.Println("Could not output the file to Html")
		}
	}

	fmt.Println("Render complete")
}
