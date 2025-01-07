package cliutil

import (
	"fmt"
	"strings"

	"github.com/khulnasoft/khulnasoft/internal/database/migration/schemas"
)

func ConstructLongHelp() string {
	return fmt.Sprintf("Available schemas:\n\n* %s", strings.Join(schemas.SchemaNames, "\n* "))
}
