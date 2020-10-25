# helpers

helpers is a go module containing some simple helper functions that are not
included with the go standard library.

------

#### Usage

To import this package to your project, run

```bash
go get "github.com/mxdillon/helpers"
```

and add

```go
import(
    "github.com/mxdillon/helpers"
)
```

to your `go.mod`.

Now you can call functions contained in the module, ie:

```go
b, i := helpers.IsInString([]byte("n"), "readme document")
```

