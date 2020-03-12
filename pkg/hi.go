package pkg

import (
	"fmt"

	"gopkg.in/launchdarkly/go-mod-dep-compat-test.v2/internal"
)

func Greet() {
	fmt.Println(internal.Msg)
}
