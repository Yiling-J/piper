package config_test

import (
	"embed"
	"os"
	"testing"

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
