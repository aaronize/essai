package main

import (
    "context"
    "fmt"
    "github.com/spf13/cobra"
    "os"
    "os/signal"
    "syscall"
)

var (
    _version    string
    _goVersion  string
    _commit     string
    _date       string

    versionFlag     bool
    runningMode     string
    configFilePath  string
)

var rootCMD = &cobra.Command{
    Use: "essaid",
    Short: "",
    Long: ``,
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Printf("Version: %s; Commit id: %s; Compiler: %s; Date: %s\n",
            _version, _commit, _goVersion, _date)
        if versionFlag {
            return
        }

        exit := make(chan os.Signal)
        signal.Notify(exit, os.Interrupt, os.Kill, syscall.SIGTERM)

        ctx, cancel := context.WithCancel(context.WithValue(context.Background()), "RunningMode", runningMode)
        laun := launcher.
    },
}

func main() {

}
