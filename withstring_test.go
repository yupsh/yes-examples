package yes_test

import (
	yup "github.com/gloo-foo/framework"
	. "github.com/yupsh/yes"
)

func ExampleYes_withString() {
	// yes "hello" | head -n 2
	yup.MustRun(
		Yes("hello"),
	)
	// Output:
	// hello
	// hello
}

