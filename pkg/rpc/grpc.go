package rpc

import (
    "context"
    "google.golang.org/grpc"
    "log"
)

type Config struct {
    Address     string      `yaml:"address"`
}

func NewConfig() *Config {
    return &Config{

    }
}

type Client struct {
    address     string      // ip:port
    connection  *grpc.ClientConn

    //handler     protos.Client
}

func NewClient(address string) (*Client, error) {
    conn, err := grpc.Dial(address, grpc.WithInsecure())
    if err != nil {
        return nil, err
    }

    return &Client{
        address: address,
        connection: conn,
        //handler: ,
    }, nil
}

func (c *Client) Close() {
    if c.connection == nil {
        return
    }
    if err := c.connection.Close(); err != nil {
        log.Println(err.Error())
    }
}

func (c *Client) SendRequest(ctx context.Context, sendData interface{}) error {
    var err error

    switch sendData.(type) {
        // TODO

    }
    if err != nil {
        return err
    }

    return nil
}
