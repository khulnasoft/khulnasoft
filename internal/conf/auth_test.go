package conf

import (
	"reflect"
	"testing"

	"github.com/khulnasoft/khulnasoft/internal/dotcom"
)

func TestAuthPublic(t *testing.T) {
	t.Run("Default, self-hosted instance non-public auth", func(t *testing.T) {
		dotcom.MockKhulnasoftDotComMode(t, false)
		got := AuthPublic()
		want := false
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("Khulnasoft.com public auth", func(t *testing.T) {
		dotcom.MockKhulnasoftDotComMode(t, true)
		got := AuthPublic()
		want := true
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}
