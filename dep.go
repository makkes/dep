package v2

import (
	"fmt"

	"github.com/makkes/dep2"
)

func Version() string {
	return fmt.Sprintf("v2.0.0 (dep2 %s)", dep2.Version())
}
