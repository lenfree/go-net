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
        noIPFlag := 0
        for _, i := range interfaces {
                a, err := i.Addrs()
                if err != nil {
                        fmt.Printf("%s", err.Error())
                }
                for _, add := range a {
                        ip, _, _ := net.ParseCIDR(add.String())

                        // only return non IPv4 and not loopback interface
                        if ip.IsLoopback() == false && ip.To4() != nil {
                                noIPFlag += 1
                                fmt.Printf("%s belongs to %s\n", yellow(add), red(i.Name))
                        }
                }
        }
        if noIPFlag == 0 {
                fmt.Printf("%s\n", "Please verify that you have properly connected to your network.")
        }
}
