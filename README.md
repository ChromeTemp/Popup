<!-- cspell:word Println -->

# Popup
> Create Windows Popup simple UI in Go

## Installation

```bash
go get github.com/ChromeTemp/Popup
```

## Usage

```go
package main

import (
    "github.com/ChromeTemp/Popup"
    "fmt"
)

main() {
    // shows a native Windows alert
    Popup.Alert("Example Title", "Example Content")
    // shows a native Windows dialog (and handle user action)
    res := Popup.Dialog("Dialog", "Want to continue?")
    // Example: user will press Ok
    fmt.Println("Pressed Ok? "+res)
    // logs: Pressed Ok? true
}
```
