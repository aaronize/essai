package main

import (
    "context"
    "fmt"
    "github.com/aaronize/essai"
    "github.com/aaronize/essai/cmd/essaid/launcher"
    "github.com/spf13/cobra"
    "log"
    "os"
    "os/signal"
    "syscall"
    "time"
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

        ctx, cancel := context.WithCancel(context.WithValue(context.Background(), "RunningMode", runningMode))
        la := launcher.NewLauncher(configFilePath)
        if err := la.Launch(ctx); err != nil {
            log.Fatalf("Service start error, %s\n", err.Error())
        }

        <-exit
        cancel()

        time.Sleep(100)
        log.Println("clean-api service CLOSED!\n Bye!(๑╹◡╹)/\"\"\"")
    },
}

func main() {
    essai.SetBuildInfo(_version, _goVersion, _commit, _date)
    if err := rootCMD.Execute(); err != nil {
        log.Fatal(err)
    }
}

func init() {
    rootCMD.PersistentFlags().StringVarP(&configFilePath, "config", "c", "./conf/server.yaml",
        "specify config file path")
    rootCMD.PersistentFlags().StringVarP(&runningMode, "env", "e", "development",
        "specify running environment [production/development]")
    rootCMD.PersistentFlags().BoolVarP(&versionFlag, "version", "v", false, "print version info")
}
