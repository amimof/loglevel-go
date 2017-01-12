# loglevel-go
Standard out formatter that supports log levels and coloured output. 

## Get started

`go get github.com/amimof/loglevel-go`

```go
package main

import "github.com/amimof/loglevel-go"

func main() {
    
}
```


## Example

```go
package main

import "github.com/amimof/loglevel-go"

func main() {
    // Create an instance of loglevel
    log := loglevel.SetupNew("MyLogger")
    
    // Set loglevel. 3=Debug, 2=Info, 1=Warn, 0=Error. Defaults to Info.
    log.Level.SetLevel(3)
    
    // Log away
    log.Debug("Debug message")
    log.Info("Info message")
    log.Warn("Warn message")
    log.Errorf("Error message")
    
    // Also supports formatting
    fname := "Luke"
    lname := "Skywalker"
    log.Infof("The Jedi %s, %s", fname, lname)
}
```
