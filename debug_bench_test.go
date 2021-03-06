package debug

import (
	"os"
	"path/filepath"
	"testing"
)

func BenchmarkDisabled(b *testing.B) {
	debug := Debug("something")
	for i := 0; i < b.N; i++ {
		debug.Log("stuff")
	}
}

func BenchmarkMatch(b *testing.B) {
	debug := Debug("something")
	Enable("something")
	for i := 0; i < b.N; i++ {
		debug.Log("stuff")
	}
}

func BenchmarkMatchJson(b *testing.B) {
	SetFormatter(&JSONFormatter{})
	SetHasColors(false)
	debug := Debug("something")
	Enable("something")
	for i := 0; i < b.N; i++ {
		debug.Log("stuff")
	}
	SetFormatter(nil)
	SetHasColors(true)
}

func BenchmarkMatchWithFields(b *testing.B) {
	debug := Debug("something")
	Enable("something")
	debug.WithFields(map[string]interface{}{
		"a": "one",
		"b": "two",
	})
	for i := 0; i < b.N; i++ {
		debug.Log("stuff")
	}
}

func BenchmarkMatchError(b *testing.B) {
	debug := Debug("something")
	Enable("error*")
	for i := 0; i < b.N; i++ {
		debug.Error("stuff")
	}
}

func BenchmarkMatchNonStringNonFunc(b *testing.B) {
	debug := Debug("something")
	Enable("something")
	for i := 0; i < b.N; i++ {
		debug.Log(os.Args[:1])
	}
}

func BenchmarkDisabledLazy(b *testing.B) {
	debug := Debug("something")
	for i := 0; i < b.N; i++ {
		debug.Log(func() string { return "lazy" })
	}
}

func BenchmarkNonMatch(b *testing.B) {
	debug := Debug("something")
	Enable("nonmatch")
	for i := 0; i < b.N; i++ {
		debug.Log("stuff")
	}
}

func BenchmarkLargeNonMatch(b *testing.B) {
	debug := Debug("large:not:lazy")

	abs, _ := filepath.Abs("./crashes.json")
	file := GetFileBytes(abs)

	Enable("nonmatch")
	for i := 0; i < b.N; i++ {
		debug.Log(string(file))
	}
}

func BenchmarkLargeLazyNonMatch(b *testing.B) {
	debug := Debug("large:lazy")

	abs, _ := filepath.Abs("./crashes.json")
	file := GetFileBytes(abs)

	Enable("nonmatch")
	for i := 0; i < b.N; i++ {
		debug.Log(func() string {
			return string(file)
		})
	}
}

func BenchmarkLargeMatch(b *testing.B) {
	debug := Debug("large:lazy")

	abs, _ := filepath.Abs("./crashes.json")
	file := GetFileBytes(abs)

	Enable("large:lazy")
	for i := 0; i < b.N; i++ {
		debug.Log(string(file))
	}
}

func BenchmarkLargeLazyMatch(b *testing.B) {
	debug := Debug("large:lazy")

	abs, _ := filepath.Abs("./crashes.json")
	file := GetFileBytes(abs)

	Enable("large:lazy")
	for i := 0; i < b.N; i++ {
		debug.Log(func() string {
			return string(file)
		})
	}
}
