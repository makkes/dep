package dep

import (
	"fmt"

	"github.com/makkes/dep2"
)

func Version() string {
	return fmt.Sprintf("v0.0.5 (dep2 %s)", dep2.Version())
}
