package pkg

import (
	"fmt"

	"gopkg.in/launchdarkly/go-mod-dep-compat-test.v3/internal"
)

func Greet() {
	fmt.Println(internal.Msg)
}
