# resolver

:zap: Simply DNS resolver with timeouts

Based on [miekg/dns](https://github.com/miekg/dns).

# Status
  Under development

# Installation
```
go get github.com/joy4eg/resolver
```

# Usage

```go
package main

import (
	"fmt"
	"time"

	"github.com/joy4eg/resolver"
)

func main() {
	ip, err := resolver.ResolveIPv4("reddit.com")
	if err != nil {
		panic(err)
	}
	fmt.Println(ip.String()) // 151.101.1.140

	// Or you can use compat API
	ips, err := resolver.LookupIPTimeout("reddit.com", 1 * time.Second)
	if err != nil {
		panic(err)
	}
	fmt.Println(ips) // [151.101.65.140 151.101.1.140 151.101.193.140 151.101.129.140]
}
```

# Author

**joy4eg**

* <http://github.com/joy4eg>
* <me@ys.lc>

# License

Released under the [MIT License](https://github.com/joy4eg/resolver/blob/master/LICENSE).
