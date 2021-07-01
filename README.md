# boilerplate-golang-cli

The goal of this code is to give a basic framework for a golang CLI that deals with the scaffolding of configuration, logging and state.

Configuration is managed parameters from command-line, config file and environment variables. This is done using [cobra](https://github.com/spf13/cobra) [viper](https://github.com/spf13/viper) and [XDG](https://github.com/adrg/xdg) for standard directories configuration. If the same parameters are provided through different routes then command-line parameters take precedence, then environment variables, then config file values, then finally any flag defaults specified in the code. Config parameters are then accessed through `viper.Get` methods.

Environment variables are automatically bound to Viper config strings with naming system APPNAME_PARAMETERNAME (e.g. EXAMPLEAPP_FLAGPARAM).

This framework will create persistent filesystem objects. An empty config file and directory structure are automatically created in a location specified by the --config flag, which defaults to `$XDG_CONFIG_HOME/appName/config.yaml` with `$XDG_CONFIG_HOME` defaulting to `~/.config`.

Commands and subcommands are managed through [cobra](https://github.com/spf13/cobra). Instead of the default cobra approach of package global variables, this builds a set of command methods. Building the command and subcommand tree is left as a manual exercise but is fairly straightforward, examples are given. Do not try using the cobra CLI!

Logging is provided through [logrus](https://github.com/sirupsen/logrus). This can easily be changed if desired.

Application logic is separated from the CLI, allowing for multiple clients to the same application logic (e.g. HTTP client as well as CLI).

The framework includes a SQLite database for locally persisted application state. It uses [appdb](https://github.com/AndrewMobbs/appdb), a small library that adds some integrity checks to opening and creating application databases. This can easily be removed if the specific application doesn't require one. The assumption is that the user will manually run a one-off "init" command to initialize the database and any other application state (such as config choices), but this could be automated if desired.

Probably thousands of these on github, but this one is mine.

Examples:  
`EXAMPLEAPP_FLAGPARAM=42 ./exampleapp example subcommand foo`  
`./exampleapp init --loglevel info`
