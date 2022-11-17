package main

import (
	"fmt"

	"github.com/docker/docker/client"
)

func main() {
	cli, _ := client.NewClientWithOpts(client.FromEnv)

	fmt.Println(cli)
}
