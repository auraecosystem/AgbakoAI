package main

import (
    "context"
    "fmt"

    clientv3 "go.etcd.io/etcd/client/v3"
)

func main() {
    cli, err := clientv3.New(clientv3.Config{
        Endpoints: []string{"localhost:2379"},
    })
    if err != nil {
        panic(err)
    }
    defer cli.Close()

    _, err = cli.Put(context.Background(), "web4/project", "qubuhub")
    if err != nil {
        panic(err)
    }

    resp, err := cli.Get(context.Background(), "web4/project")
    if err != nil {
        panic(err)
    }

    for _, kv := range resp.Kvs {
        fmt.Printf("%s=%s\n", kv.Key, kv.Value)
    }
}
