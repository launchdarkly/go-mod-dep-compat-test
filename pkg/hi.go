package pkg

import (
	"fmt"

	"gopkg.in/launchdarkly/go-mod-dep-compat-test.v5/internal"
)

func Greet() {
	fmt.Println(internal.Msg)
}
