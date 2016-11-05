package main

import (
        "fmt"
        "time"

        "github.com/fatih/color"
        "github.com/sparrc/go-ping"
        "gopkg.in/urfave/cli.v1"
)

func IcmpPingCli() cli.Command {
        command := cli.Command{
                Name:      "ping",
                ShortName: "p",
                Usage:     "Ping Google.com which is the most reliable service in the Internet so far",
                Action:    pingGoogle,
        }
        return command
}

func pingGoogle(ctx *cli.Context) {
        //yellow := color.New(color.FgYellow).SprintFunc()
        red := color.New(color.FgRed).SprintFunc()
        pinger, err := ping.NewPinger("10.10.10.10")
        if err != nil {
                fmt.Printf("%s", err.Error())
        }
        pinger.Count = 3
        pinger.Timeout = time.Second * 5
        pinger.Run()
        stats := pinger.Statistics()
        fmt.Printf("%+#v\n", stats)
        if stats.PacketsRecv == 0 && stats.PacketLoss == 100 {
                fmt.Printf("%s\n", "oops!")
        } else {
                fmt.Printf("the address being pinged is %s", red(stats.Addr))
        }
}
