# journaler [![GoDoc](https://godoc.org/github.com/missionMeteora/journaler?status.svg)](https://godoc.org/github.com/missionMeteora/journaler) ![Status](https://img.shields.io/badge/status-beta-yellow.svg)

journaler is an output library which prepends messages with a colored label and a prefix

![Example screenshot](https://raw.githubusercontent.com/missionMeteora/journaler/master/screenshot.png "Example screenshot")

## Usage
``` go
package main

import "github.com/missionMeteora/journaler"

func main() {
        jj := journaler.New("Main service", "Sub service")
        jj.Success("Database entry posted")
        jj.Notification("CPU temperatures are at 40*C")
        jj.Warning("Update to remote server has failed")
        jj.Error("Danger, Will Robinson!")
}
```