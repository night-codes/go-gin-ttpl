# go-gin-ttpl
Custom render for gin-gonic by "text/template" instead "html/template"  

## Why Do It?
Sometimes we need to render templates without HTML. Also "html/template" has the [problem](https://github.com/golang/go/issues/12149) with parser.
Recursively parse all template files in all subdirectories.

## How To Use
```go
package main

import (
    "gopkg.in/gin-gonic/gin.v1"
    "gopkg.in/night-codes/go-gin-ttpl.v1"
)

func main() {
    r := gin.New()
    ttpl.Use(r, "templates")
    r.GET("/", func(c *gin.Context) {
        c.HTML(200, "index.html", gin.H{"title": "top"})
    })
    r.Run(":41000")
}
```

When you want to add your own template functions you can pass a third parameter - FuncMap

```go
import (
    // ...
    "text/template"
)

func main() {
    // ...
    funcMap := template.FuncMap{
        "customFunc": customFunc, // A custom template function
    }
    ttpl.Use(r, "templates", funcMap) // <--
    // ...
}
```


## License
DO WHAT THE FUCK YOU WANT TO PUBLIC LICENSE
Version 2, December 2004

Copyright (C) 2015 Oleksiy Chechel <alex.mirrr@gmail.com>

Everyone is permitted to copy and distribute verbatim or modified
copies of this license document, and changing it is allowed as long
as the name is changed.

DO WHAT THE FUCK YOU WANT TO PUBLIC LICENSE
TERMS AND CONDITIONS FOR COPYING, DISTRIBUTION AND MODIFICATION

 0. You just DO WHAT THE FUCK YOU WANT TO.
