package main

import (
	"flag"

	"github.com/technicallyjosh/protoc-gen-enum-json-string/internal/generator"
	"google.golang.org/protobuf/compiler/protogen"
)

var flags flag.FlagSet

func main() {
	opts := protogen.Options{
		ParamFunc: flags.Set,
	}

	opts.Run(func(plugin *protogen.Plugin) error {
		return generator.New(plugin).Run()
	})
}
