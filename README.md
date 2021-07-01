# boilerplate-golang-cli

Basic framework for a simple golang CLI, combining command-line, config file and environment arguments using [cobra](https://github.com/spf13/cobra) [viper](https://github.com/spf13/viper) and [XDG](https://github.com/adrg/xdg) standard directories configuration. 

Probably thousands of these on github, but this one is mine.

Examples:
`EXAMPLEAPP_FLAGPARAM=42 ./exampleapp example subcommand foo`
`./exampleapp init --loglevel info`
