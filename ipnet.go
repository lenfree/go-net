package main

import (
        "fmt"
        "github.com/fatih/color"
        "net"
        "gopkg.in/urfave/cli.v1"
)

func IpNetCli() cli.Command{
        command := cli.Command{
                Name     : "ipnet",
                ShortName: "ip",
                Usage    : "Return IPv4/IPv6 with interface name",
                Action   : ipnet,
        }
        return command
}

// Returns interface name with IPv4/IPv6 that are present
func ipnet(ctx *cli.Context) {
        yellow := color.New(color.FgYellow).SprintFunc()
        red := color.New(color.FgRed).SprintFunc()

        interfaces, err := net.Interfaces()
        if err != nil {
                fmt.Printf("%s", err.Error())
        }
        for _, i := range interfaces {
                a, err := i.Addrs()
                if err != nil {
                        fmt.Printf("%s", err.Error())
                }
                for _, add := range a {
                        fmt.Printf("%s belongs to %s\n", yellow(add), red(i.Name))
                }
        }
}
