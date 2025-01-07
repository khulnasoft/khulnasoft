package app

import "github.com/khulnasoft/khulnasoft/internal/txemail"

func init() {
	txemail.DisableSilently()
}
