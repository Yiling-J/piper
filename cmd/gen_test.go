package main

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"
)

func TestGenerate(t *testing.T) {
	const folder = "test_config"
	const file = "test_config/test.toml"
	const expectFile = "//nolint\npackage config\n\ntype bar struct {\n\tFoo string\n}\n\nvar Bar = bar{\n\n\tFoo: \"bar.foo\",\n}\n\nconst Foo = \"foo\"\n"

	path, err := viperLoad(folder)
	require.Nil(t, err)
	require.Equal(t, "test_config", path)
	err = generate(path)
	require.Nil(t, err)
	b, err := ioutil.ReadFile("test_config/config.go")
	require.Nil(t, err)
	err = os.Remove("test_config/config.go")
	require.Nil(t, err)
	require.Equal(t, expectFile, string(b))

	viper.Reset()
	path, err = viperLoad(file)
	require.Nil(t, err)
	require.Equal(t, "test_config", path)
	err = generate(path)
	require.Nil(t, err)
	b, err = ioutil.ReadFile("test_config/config.go")
	require.Nil(t, err)
	err = os.Remove("test_config/config.go")
	require.Nil(t, err)
	require.Equal(t, expectFile, string(b))
}
