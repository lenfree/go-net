package main

import (
        "os"

        "gopkg.in/urfave/cli.v1"
)

func initApp() *cli.App {
        app := cli.NewApp()

        app.Name = "go-net"
        app.Version = "0.0.1"
        app.Usage = "My CLI"
        app.Author = "Lenfree Yeung"
        app.Email = "lenfree.yeung@gmail.com"

        app.Commands = []cli.Command{
                IpNetCli(),
                IcmpPingCli(),
        }
        return app
}

func main() {
        app := initApp()
        app.Run(os.Args)
}
