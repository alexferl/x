package main

import (
	"fmt"

	xconfig "github.com/alexferl/x/config/v2"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// Config holds all configuration for our program
type Config struct {
	*xconfig.Config
	MyKey string
}

// NewConfig creates a Config instance
func NewConfig() *Config {
	c := &Config{
		Config: xconfig.New(),
		MyKey:  "value",
	}
	return c
}

// addFlags adds all the flags from the command line
func (c *Config) addFlags(fs *pflag.FlagSet) {
	fs.StringVar(&c.MyKey, "mykey", c.MyKey, "My key.")
}

// BindFlags normalizes and parses the command line flags
func (c *Config) BindFlags() {
	c.addFlags(pflag.CommandLine)
	c.Config.BindFlags() // Bind the default flags from x/config
}

func main() {
	c := NewConfig()
	c.BindFlags()
	fmt.Println(viper.GetString("app-name")) // from xconfig, overloaded in configs/config.dev.toml
	fmt.Println(viper.GetString("mykey"))
}
