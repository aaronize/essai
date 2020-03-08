package launcher

import (
    "context"
    "fmt"
    "github.com/aaronize/essai/cmd/essaid/config"
    "log"

    //"github.com/DeanThompson/ginpprof"
    "github.com/aaronize/essai/pkg/middle"
    "github.com/gin-gonic/gin"
    //ginSwagger "github.com/swaggo/gin-swagger"
    //"github.com/swaggo/gin-swagger/swaggerFiles"
    "net/http"
    "os"
)

const (
    RunningEnvProduction = "production"
    RunningEnvDevelopment = "development"
)

var runningEnv = RunningEnvDevelopment

type Launcher struct {
    configPath      string
}

func NewLauncher(configPath string) *Launcher {
    return &Launcher{
        configPath: configPath,
    }
}

func (l *Launcher) Launch(ctx context.Context) error {
    l.setEnv(ctx)
    fmt.Printf("+++ Running Environment: %s +++\n", runningEnv)

    conf := config.NewConfig()
    if l.configPath == "" {
        log.Println("Not specify config path, using default config.")
    }
    if err := conf.Parse(runningEnv, l.configPath); err != nil {
        return err
    }

    return nil
}

func (l *Launcher) setEnv(ctx context.Context) {
    systemSetEnv := os.Getenv("RunningEnv")
    runningSetEnv := ctx.Value("RunningEnv")
    if systemSetEnv == RunningEnvProduction || runningSetEnv == RunningEnvProduction {
        runningEnv = RunningEnvProduction
    }
    return
}

func (l *Launcher) newHttpHandler() http.Handler {
    g := gin.Default()

    g.Use(middle.Cros)

    //ginpprof.Wrap(g)
    //g.GET("api/v1/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))


    return g
}
