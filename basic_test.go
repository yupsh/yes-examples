package yes_test

import (
	gloo "github.com/gloo-foo/framework"
	. "github.com/yupsh/yes"
)

func ExampleYes_basic() {
	// yes | head -n 3
	// Note: yes produces infinite output, so we limit with context or external means
	gloo.MustRun(
		Yes(),
	)
	// Output:
	// y
	// y
	// y
}
