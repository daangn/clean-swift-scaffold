<img src="https://github.com/GeekTree0101/clean-swift-scaffold/blob/develop/logo.png" />

### Clean-Swift source & test code auto generator

[![Go](https://github.com/daangn/clean-swift-scaffold/actions/workflows/go.yml/badge.svg)](https://github.com/daangn/clean-swift-scaffold/actions/workflows/go.yml)


## Overview

#### Run
<img height=300pt src="https://github.com/GeekTree0101/clean-swift-scaffold/blob/develop/res/example.png" />

#### Output
<img height=300pt src="https://github.com/GeekTree0101/clean-swift-scaffold/blob/develop/res/output.png" />

## Basic Usage

#### make config.yaml
```yaml
target_project_name: Miro // target project name
copyright: Geektree0101 // copyright
template_path: ./templates // templates path
source_path: ./Playground/Sources // base source file destination
test_path: ./Playground/Tests // base test file destination
indentation: 2 // indentation
```

#### add clean_swift_scaffold runner command on your command 
```go
var rootCmd = &cobra.Command{
	Use:   "your cmd",
	Short: "your cmd short marty",
	Long:  "your cmd long something",
}

init() {
  rootCmd.AddCommand(clean_swift_scaffold.NewRunnerCommand("**use_name**"))
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
```

#### run 
```sh
your_command **use_name** -n Feed -u Fetch,Delete,Update
```

flag list
```sh
- -n/--name: scene prefix
- -u/--usecase: some model behavior (such as Fetch, Get, Reload, Delete and so on)
- -c/--config: config.yaml path ./some_dir/config.yaml or ./some_dir/some_config.yaml
- -s/--source: custom base source_dir (Default values follow the configuration file.)
- -t/--test: custon base test_dir (Default values follow the configuration file.)
```

- Please set the name and directory of the configuration file freely. Instead, please enter the correct path on -c/--config flag.
- Default values of source & test directoly flag follow the configuration file.
