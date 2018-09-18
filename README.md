# beezap


## Usage

### Start using it

Download and install it:

```sh
$ go get github.com/GNURub/beezap
```

Import it in your code:

```go
import (
    "github.com/GNURub/beezap"
)
```

## Example

See the [example](example/main.go).

[embedmd]:# (example/main.go go)
```go
package main

import (
	"time"

	"github.com/astaxie/beego"
	"go.uber.org/zap"
	"github.com/GNURub/beezap"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Ctx.WriteString("hello world")
}

func main() {
	logger, _ := zap.NewProduction()
	beezap.InitBeeZapMiddleware(logger, time.RFC3339, true)
	beego.Router("/", &MainController{})
	beego.Run(":8090")
}
```