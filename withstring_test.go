package yes_test

import (
	gloo "github.com/gloo-foo/framework"
	. "github.com/yupsh/yes"
)

func ExampleYes_withString() {
	// yes "hello" | head -n 2
	gloo.MustRun(
		Yes("hello"),
	)
	// Output:
	// hello
	// hello
}
