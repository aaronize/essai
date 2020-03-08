package config

import (
    "bytes"
    "fmt"
    "github.com/aaronize/essai/internal/api"
	"github.com/aaronize/essai/pkg/db"
	"github.com/aaronize/essai/pkg/queue"
    "github.com/aaronize/essai/pkg/rpc"
    "github.com/aaronize/essai/pkg/server"
    "gopkg.in/yaml.v2"
    "io/ioutil"
)

type Config struct {
	Kind   string         `yaml:"kind"`
	Server *server.Config `yaml:"server"`
	Queue  *queue.Config  `yaml:"queue"`
	DB     *db.Config     `yaml:"database"`
	API    *api.APIs      `yaml:"api"`
	GRPC    *rpc.Config     `yaml:"gprc"`
}

func NewConfig() *Config {
    return &Config{
        Server: server.NewConfig(),
        Queue: queue.NewConfig(),
        DB: db.NewConfig(),
        GRPC: rpc.NewConfig(),
        API: nil,
    }
}

func (c *Config) Parse(kind, path string) error {
    content, err := ioutil.ReadFile(path)
    if err != nil {
        return err
    }

    //var flag bool
    decoder := yaml.NewDecoder(bytes.NewReader(content))
    for decoder.Decode(c) == nil {
        if c.Kind == kind {
            return nil
        }
    }

    return fmt.Errorf("请指定运行环境[production/development]")
}
