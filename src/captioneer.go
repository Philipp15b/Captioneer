package main

import (
	"./ccompiler"
	"./config"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

var language = flag.String("language", "english", "The language of this caption file.")

func main() {
	flag.Parse()

	if len(os.Args) < 2 || os.Args[1] == "" {
		fmt.Println("[ERROR] No input file specified!\nUsage: captioneer filename [--language=\"english\"]")
		return
	}

	input, err := ioutil.ReadFile("./captions.txt")
	if err != nil {
		fmt.Println("[ERROR] " + err.Error())
		return
	}

	captions, err := config.ProcessCaptions(string(input))
	if err != nil {
		fmt.Println("[ERROR] " + err.Error())
		return
	}

  outname := "./closecaption_" + string(*language) + ".txt"
	fo, err := os.OpenFile(outname, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("[ERROR] " + err.Error())
		return
	}
	defer fo.Close()

	err = ccompiler.WriteIntermediate(fo, ccompiler.CompileIntermediate(captions, string(*language)))
	if err != nil {
		fmt.Println("[ERROR] " + err.Error())
		return
	}

	compiler, err := ccompiler.FindCompiler()
	if err != nil {
		fmt.Println("[ERROR] " + err.Error())
		return
	}

	fmt.Println(compiler.Run(outname))

}
