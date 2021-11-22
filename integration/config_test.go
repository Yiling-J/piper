package config_test

import (
	"embed"
	"os"
	"testing"
	"time"

	"github.com/Yiling-J/piper"
	configtm "github.com/Yiling-J/piper/integration/config_toml_multi"
	"github.com/stretchr/testify/require"
)

//go:embed config_toml_multi/*
var configTomlMulti embed.FS

//go:embed config_yaml_multi/*
var configYamlMulti embed.FS

func assertConfig(t *testing.T) {
	// value from config test
	require.Equal(t, piper.GetString(configtm.Foo), "test")
	// value from config test
	require.Equal(t, piper.GetString(configtm.Foobar.Foo), "test")
	// value from config qux
	require.Equal(t, piper.GetString(configtm.Foofoo), "k")
	// value from config qux
	require.Equal(t, piper.GetString(configtm.Quux.Foo), "j")
	// value from config foo
	require.Equal(t, piper.GetString(configtm.Foobar.Qux.Xyz), "h")
	// value from config bar
	require.Equal(t, piper.GetString(configtm.Barfoo), "m")
	// value from config base
	require.Equal(t, piper.GetStringSlice(configtm.Foobar.Bar), []string{"c", "d", "e"})
}

func assertIConfig(t *testing.T) {
	// value from config test
	require.Equal(t, piper.IGetString(configtm.Foo), "test")
	// value from config test
	require.Equal(t, piper.IGetString(configtm.Foobar.Foo), "test")
	// value from config qux
	require.Equal(t, piper.IGetString(configtm.Foofoo), "k")
	// value from config qux
	require.Equal(t, piper.IGetString(configtm.Quux.Foo), "j")
	// value from config foo
	require.Equal(t, piper.IGetString(configtm.Foobar.Qux.Xyz), "h")
	// value from config bar
	require.Equal(t, piper.IGetString(configtm.Barfoo), "m")
	// value from config base
	require.Equal(t, piper.IGetStringSlice(configtm.Foobar.Bar), []string{"c", "d", "e"})
}

func TestTomlMulti(t *testing.T) {
	piper.Reset()
	err := piper.Load("config_toml_multi/test.toml")
	require.Nil(t, err)
	assertConfig(t)
	assertIConfig(t)
}

func TestTomlMultiEmbed(t *testing.T) {
	piper.Reset()
	piper.SetFS(configTomlMulti)
	err := os.Rename("config_toml_multi", "config_toml_multi_tmp")
	require.Nil(t, err)
	defer func() {
		err = os.Rename("config_toml_multi_tmp", "config_toml_multi")
		require.Nil(t, err)
	}()
	err = piper.Load("config_toml_multi/test.toml")
	require.Nil(t, err)
	assertConfig(t)
	assertIConfig(t)
}

func TestTomlMultiEmbedOverride(t *testing.T) {
	piper.Reset()
	piper.SetFS(configTomlMulti)
	err := os.Rename("config_toml_multi", "config_toml_multi_tmp")
	require.Nil(t, err)
	defer func() {
		err = os.Rename("config_toml_multi_tmp", "config_toml_multi")
		require.Nil(t, err)
	}()
	d1 := []byte("pp_imports = [\"test.toml\"]\nfoo = \"secret\"")
	err = os.Mkdir("config_toml_multi", 0755)
	require.Nil(t, err)
	err = os.WriteFile("config_toml_multi/override.toml", d1, 0755)
	require.Nil(t, err)

	err = piper.Load("config_toml_multi/override.toml")
	require.Nil(t, err)
	require.Equal(t, piper.GetString(configtm.Foo), "secret")
	require.Equal(t, piper.GetString(configtm.Foobar.Foo), "test")
	require.Equal(t, piper.GetString(configtm.Foofoo), "k")
	err = os.RemoveAll("config_toml_multi/")
	require.Nil(t, err)
}

func TestYamlMulti(t *testing.T) {
	piper.Reset()
	err := piper.Load("config_yaml_multi/test.yaml")
	require.Nil(t, err)
	assertConfig(t)
}

func TestYamlMultiEmbed(t *testing.T) {
	piper.Reset()
	piper.SetFS(configYamlMulti)
	err := os.Rename("config_yaml_multi", "config_yaml_multi_tmp")
	require.Nil(t, err)
	defer func() {
		err = os.Rename("config_yaml_multi_tmp", "config_yaml_multi")
		require.Nil(t, err)
	}()
	err = piper.Load("config_yaml_multi/test.yaml")
	require.Nil(t, err)
	assertConfig(t)
}

func TestTomlCycle(t *testing.T) {
	piper.Reset()
	err := piper.Load("config_toml_cycle/d.toml")
	require.NotNil(t, err)
}

func TestMultiPiper(t *testing.T) {
	piper.Reset()
	piperA := piper.New()
	piperB := piper.New()

	err := piper.Load("config_toml_multi/test.toml")
	require.Nil(t, err)
	err = piperA.Load("config_toml_multi/base.toml")
	require.Nil(t, err)
	err = piperB.Load("config_toml_multi/foo.toml")
	require.Nil(t, err)

	// global
	require.Equal(t, piper.GetString(configtm.Foo), "test")
	// a
	require.Equal(t, piperA.GetString(configtm.Foo), "a")
	// b
	require.Equal(t, piperB.GetString(configtm.Foo), "foofoo")

	// global
	require.Equal(t, piper.IGetString(configtm.Foo), "test")
	// a
	require.Equal(t, piperA.IGetString(configtm.Foo), "a")
	// b
	require.Equal(t, piperB.IGetString(configtm.Foo), "foofoo")

}

func TestEnvOverrideI(t *testing.T) {
	piper.Reset()
	err := piper.Load("config_toml_multi/test.toml")
	require.Nil(t, err)

	piper.V().AutomaticEnv()
	err = os.Setenv("FOO", "bar")
	require.Nil(t, err)
	require.Equal(t, piper.GetString(configtm.Foo), "bar")
	require.Equal(t, piper.IGetString(configtm.Foo), "test")
}

func TestEnvOverrideII(t *testing.T) {
	piper.Reset()
	piper.V().AutomaticEnv()
	err := os.Setenv("FOO", "bar")
	require.Nil(t, err)
	err = piper.Load("config_toml_multi/test.toml")
	require.Nil(t, err)

	require.Equal(t, piper.GetString(configtm.Foo), "bar")
	require.Equal(t, piper.IGetString(configtm.Foo), "bar")
}

func TestPiperTypes(t *testing.T) {
	piper.Reset()
	err := piper.Load("config_toml_multi/test.toml")
	require.Nil(t, err)

	require.Equal(t, interface{}(true), piper.Get(configtm.Types.Bool))
	require.Equal(t, true, piper.GetBool(configtm.Types.Bool))
	require.Equal(t, 1, piper.GetInt(configtm.Types.Int))
	require.Equal(t, int32(1), piper.GetInt32(configtm.Types.Int))
	require.Equal(t, int64(1), piper.GetInt64(configtm.Types.Int))
	require.Equal(t, uint(1), piper.GetUint(configtm.Types.Int))
	require.Equal(t, uint32(1), piper.GetUint32(configtm.Types.Int))
	require.Equal(t, uint64(1), piper.GetUint64(configtm.Types.Int))
	require.Equal(t, 3.12, piper.GetFloat64(configtm.Types.Float))
	require.Equal(t, 12*time.Second, piper.GetDuration(configtm.Types.Duration))
	require.Equal(t, []int{1, 2, 3}, piper.GetIntSlice(configtm.Types.Intslice))
	require.Equal(t, []string{"a", "b", "c"}, piper.GetStringSlice(configtm.Types.Stringslice))

	require.Equal(t, interface{}(true), piper.IGet(configtm.Types.Bool))
	require.Equal(t, true, piper.IGetBool(configtm.Types.Bool))
	require.Equal(t, 1, piper.IGetInt(configtm.Types.Int))
	require.Equal(t, int32(1), piper.IGetInt32(configtm.Types.Int))
	require.Equal(t, int64(1), piper.IGetInt64(configtm.Types.Int))
	require.Equal(t, uint(1), piper.IGetUint(configtm.Types.Int))
	require.Equal(t, uint32(1), piper.IGetUint32(configtm.Types.Int))
	require.Equal(t, uint64(1), piper.IGetUint64(configtm.Types.Int))
	require.Equal(t, 3.12, piper.IGetFloat64(configtm.Types.Float))
	require.Equal(t, 12*time.Second, piper.IGetDuration(configtm.Types.Duration))
	require.Equal(t, []int{1, 2, 3}, piper.IGetIntSlice(configtm.Types.Intslice))
	require.Equal(t, []string{"a", "b", "c"}, piper.IGetStringSlice(configtm.Types.Stringslice))

	p := piper.New()
	err = p.Load("config_toml_multi/test.toml")
	require.Nil(t, err)

	require.Equal(t, interface{}(true), p.Get(configtm.Types.Bool))
	require.Equal(t, true, p.GetBool(configtm.Types.Bool))
	require.Equal(t, 1, p.GetInt(configtm.Types.Int))
	require.Equal(t, int32(1), p.GetInt32(configtm.Types.Int))
	require.Equal(t, int64(1), p.GetInt64(configtm.Types.Int))
	require.Equal(t, uint(1), p.GetUint(configtm.Types.Int))
	require.Equal(t, uint32(1), p.GetUint32(configtm.Types.Int))
	require.Equal(t, uint64(1), p.GetUint64(configtm.Types.Int))
	require.Equal(t, 3.12, p.GetFloat64(configtm.Types.Float))
	require.Equal(t, 12*time.Second, p.GetDuration(configtm.Types.Duration))
	require.Equal(t, []int{1, 2, 3}, p.GetIntSlice(configtm.Types.Intslice))
	require.Equal(t, []string{"a", "b", "c"}, p.GetStringSlice(configtm.Types.Stringslice))

	require.Equal(t, interface{}(true), p.IGet(configtm.Types.Bool))
	require.Equal(t, true, p.IGetBool(configtm.Types.Bool))
	require.Equal(t, 1, p.IGetInt(configtm.Types.Int))
	require.Equal(t, int32(1), p.IGetInt32(configtm.Types.Int))
	require.Equal(t, int64(1), p.IGetInt64(configtm.Types.Int))
	require.Equal(t, uint(1), p.IGetUint(configtm.Types.Int))
	require.Equal(t, uint32(1), p.IGetUint32(configtm.Types.Int))
	require.Equal(t, uint64(1), p.IGetUint64(configtm.Types.Int))
	require.Equal(t, 3.12, p.IGetFloat64(configtm.Types.Float))
	require.Equal(t, 12*time.Second, p.IGetDuration(configtm.Types.Duration))
	require.Equal(t, []int{1, 2, 3}, p.IGetIntSlice(configtm.Types.Intslice))
	require.Equal(t, []string{"a", "b", "c"}, p.IGetStringSlice(configtm.Types.Stringslice))
}

func BenchmarkGet(b *testing.B) {
	piper.Reset()
	err := piper.Load("config_yaml_multi/test.yaml")
	if err != nil {
		panic(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r := piper.GetString(configtm.Foobar.Qux.Xyz)
		if r != "h" {
			panic("")
		}
	}
}

func BenchmarkIGet(b *testing.B) {
	piper.Reset()
	err := piper.Load("config_yaml_multi/test.yaml")
	if err != nil {
		panic(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r := piper.IGetString(configtm.Foobar.Qux.Xyz)
		if r != "h" {
			panic("")
		}
	}
}
