# loglevel
A simple standard out logger and formatter that supports log levels, namespaces and coloured output.

## Get started

Download the package

`go get -t github.com/amimof/loglevel-go`

Import it in your code

```go
package main

import "github.com/amimof/loglevel-go"

func main() {

    // Create an instance of loglevel
    loglevel := loglevel.New()
    
    // Set loglevel. 3=Debug, 2=Info, 1=Warn, 0=Error. Defaults to Info.
    loglevel.SetLevel(3)
    
    // Log away
    loglevel.Debug("Debug message")
    loglevel.Info("Info message")
    loglevel.Warn("Warn message")
    loglevel.Error("Error message")

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
loglevel supports both non-format and format print methods. Formatting is as good as in the native `fmt` package but with the additional log level functionality.

```go
package main

import "github.com/amimof/loglevel-go"

func main() {
    // Create an instance of loglevel
    loglevel := loglevel.New()
    
    // Set loglevel. 3=Debug, 2=Info, 1=Warn, 0=Error. Defaults to Info.
    loglevel.SetLevel(3)
        
    fname := "Luke"
    lname := "Skywalker"
    loglevel.Infof("The Jedi %s, %s", fname, lname)
}
```

Result:
```
2017-01-21T13:53:54+01:00 INFO The Jedi Luke, Skywalker
```

## Namespaces

You can created multiple loglevel instances in your app. Assign different names to them, to make debugging app packages an modules easy.

```go
package main

import "github.com/amimof/loglevel-go"

func main() {
    
    // Create an instance of loglevel
    loglevel := loglevel.New().SetLevel(3)
    
    // Set the namespace of this loglevel instance
    loglevel.Name = "StarWarsNames"    
    loglevel.PrintName =  true

    fname := "Luke"
    lname := "Skywalker"
    loglevel.Infof("The Jedi %s, %s", fname, lname)
}
```

Result:
```
2017-01-21T13:53:54+01:00 INFO StarWarsNames The Jedi Luke, Skywalker
```

## Options

If you like vanilla, `fmt.Println()`, formatting. Disable those fiels.

```go
package main

import "github.com/amimof/loglevel-go"

func main() {
    // Create an instance of loglevel
    loglevel := loglevel.New()
    loglevel.PrintTime = false
    loglevel.PrintLevel = false
    loglevel.PrintName = false
    loglevel.Warn("Achtung, Achtung!")
}
```

Result
```
Achtung, Achtung!
```