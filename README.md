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
    log := loglevel.New()
    
    // Set loglevel. 3=Debug, 2=Info, 1=Warn, 0=Error. Defaults to Info.
    log.SetLevel(3)
    
    // Log away
    log.Debug("Debug message")
    log.Info("Info message")
    log.Warn("Warn message")
    log.Error("Error message")

}
```

Result
```
2017-01-21T13:56:28+01:00 DEBUG Debug message
2017-01-21T13:56:28+01:00 INFO Info message
2017-01-21T13:56:28+01:00 WARN Warn message
2017-01-21T13:56:28+01:00 ERROR Error message
```

## Formatting 

```go
package main

import "github.com/amimof/loglevel-go"

func main() {
    // Create an instance of loglevel
    log := loglevel.New()
    
    // Set loglevel. 3=Debug, 2=Info, 1=Warn, 0=Error. Defaults to Info.
    log.SetLevel(3)
        
    fname := "Luke"
    lname := "Skywalker"
    log.Infof("The Jedi %s, %s", fname, lname)
}
```

Result:
```
2017-01-21T13:53:54+01:00 INFO The Jedi Luke, Skywalker
```

## Options

```go
package main

import "github.com/amimof/loglevel-go"

func main() {
    // Create an instance of loglevel
    log := loglevel.New()
    log.PrintTime = false
    log.PrintLevel = false
    log.PrintName = false
    log.Warn("Achtung, Achtung!")
}
```

Result
```
Achtung, Achtung!
```