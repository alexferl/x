package main

import (
	"fmt"

	xlog "github.com/alexferl/x/log"
	"github.com/rs/zerolog/log"
	"github.com/spf13/pflag"
)

func main() {
	c := xlog.DefaultConfig // use default settings
	// c := &xlog.Config{LogWriter: "json"} // use custom settings
	c.AddFlags(pflag.CommandLine)
	pflag.Parse()

	err := xlog.New(c)
	if err != nil {
		panic(fmt.Sprintf("Error initializing logger: '%v'", err))
	}

	log.Info().Msg("Hello, world!")
	log.Info().Msgf("Hello, %s!", "world")
	log.Warn().Msg("Hello, warn!")
	log.Error().Msg("Hello, error!")
}
