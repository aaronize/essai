package server

import (
    "context"
    "log"
    "net/http"
)

type Config struct {
    Listen      string      `yaml:"listen"`
}

func NewConfig() *Config {
    return &Config{
        Listen: ":15104",
    }
}

func StartHttpServer(ctx context.Context, handler http.Handler, config *Config) error {
    s := &Server{
        Addr: config.Listen,
        Handler: handler,
    }
    if err := s.ListenAndServe(ctx); err != nil {
        return err
    }

    return nil
}

type Server struct {
    Addr      string
    Handler   http.Handler
}

func (s *Server) ListenAndServe(ctx context.Context) error {
    return s.listenAndServe(ctx)
}

func (s *Server) listenAndServe(ctx context.Context) error {
    sh := &http.Server{
        Addr: s.Addr,
        Handler: s.Handler,
    }

    go func() {
        select {
        case <-ctx.Done():
            log.Println("Shutting down http server...")
            _ = sh.Shutdown(ctx)
            return
        }
    }()

    go func() {
        if err := sh.ListenAndServe(); err != nil {
            log.Println(err)
            return
        }
    }()

    return nil
}
