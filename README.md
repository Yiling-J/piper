# piper - Simple Wrapper For Viper
![example workflow](https://github.com/Yiling-J/piper/actions/workflows/go.yml/badge.svg)
![Go Report Card](https://goreportcard.com/badge/github.com/yiling-j/piper?style=flat-square)

- **Single Source of Truth**
- **Code Generation, No Typo**
- **Config Inheritance**
- **Multiple Config Strategies Support**

## Why Piper
If you are familiar with Django, this is how Django settings module looks like:

```console {12-20}
└── settings
    ├── base.py
    ├── dev.py
    ├── stage.py
    └── prod.py
```
`dev.py` will inherit from `base.py`, also `stage.py` will inherit from `dev.py`.

Start Django with selected setting:
```shell
export DJANGO_SETTINGS_MODULE=mysite.settings.prod
django-admin runserver
```

And this is how you access settings in Django:

```python
from django.conf import settings

author = settings.Author
```

I want to have similar experience with Viper, so here comes this wrapper:

```console {12-20}
└── config
    ├── base.toml
    ├── dev.toml
    ├── stage.toml
    └── prod.toml
```
This is how you access config using piper:

```go
import "your_project/config"

func main() {
	piper.Load("config/stage.toml")
	author = piper.GetString(config.Author)
}
```

Check example folder for more details.

## Installation
```shell
go get github.com/Yiling-J/piper/cmd
```
## Add Config Files
Add your config files to your config folder, usually your config folder should be under project root folder.

toml example, you can also use yaml or json.
```console {12-20}
project
└── config
    ├── base.toml
    ├── dev.toml
    ├── stage.toml
    └── prod.toml

```
To support inheritance, you need to add a special key to your config file called `pp_imports`
```
pp_imports = ["base.toml", "dev.toml"]
```
Piper will resolve that automatically.

## Config Key Generation
Run code generation from the root directory of the project as follows:
```shell
go run github.com/Yiling-J/piper/cmd your_config_folder
```
In this step piper will load all files in your config folder and merge them together.
Then piper will generate `config.go` under your config folder, include all your config keys.
Also you will see the config structure when pipe generating code.

After code genertation, your config folder should look like:
```console {12-20}
└── config
    ├── base.toml
    ├── dev.toml
    ├── stage.toml
    ├── prod.toml
    └── config.go
```
## Use Piper

### Strategy I - Embed
embed your config folder into your code, single executable when you deploy.
```go
import (
	"github.com/Yiling-J/piper"
	"your_project/example/config"
)

//go:embed config/*
var configFS embed.FS

piper.SetFS(configFS)
piper.Load("config/stage.toml")
author := piper.GetString(config.Author)
```

### Strategy II - Embed with Env
embed your config folder into your code, single executable when you deploy, and replace secret with env.
```go
import (
	"github.com/Yiling-J/piper"
	"your_project/example/config"
)

//go:embed config/*
var configFS embed.FS

os.Setenv("SECRET", "qux")
piper.SetFS(configFS)
piper.Load("config/stage.toml")
piper.V().AutomaticEnv()
secret := piper.GetString(config.Secret)
```

### Strategy III -  Copy config directory
copy config folder when building docker image, so the true config folder exists.
```go
import (
	"github.com/Yiling-J/piper"
	"your_project/example/config"
)

piper.Load("config/stage.toml")
author := piper.GetString(config.Author)
```

### Strategy IV - Mix embed and copy directory
Embed your config folder, but keep some secret keys in a real config file.
```go
import (
	"github.com/Yiling-J/piper"
	"your_project/example/config"
)

//go:embed config/*
var configFS embed.FS

// "config/stage_with_secret.toml" is not in source code,
// may come from docker build or k8s ConfigMap
piper.SetFS(configFS)
piper.Load("config/stage_with_secret.toml")
author := piper.GetString(config.Author)
```
True directory will take precedence over embeded directory.

## Access Viper
Piper is just a wrapper, so you can always get the wrapped viper instance:
```go
v := piper.V()
```
Be careful when using viper directly to load config, piper may not work properly.

## Piper or Pipers?

Piper comes ready to use out of the box. There is no configuration or
initialization needed to begin using Piper. Since most applications will want
to use a single central repository for their configuration, the piper package
provides this. It is similar to a singleton.

In all of the examples above, they demonstrate using piper in its singleton
style approach.

### Working with multiple pipers

You can also create many different pipers for use in your application. Each will
have its own unique set of configurations and values. Each can read from a
different config file. All of the functions that piper
package supports are mirrored as methods on a piper.

Example:

```go
x := piper.New()
y := piper.New()

x.Load("config/x.toml")
y.Load("config/y.toml")

// access viper
// p := x.V
```

When working with multiple pipers, it is up to the user to keep track of the
different pipers.

## Performance
Sometimes you may find viper a little slow, because viper need to check in the following order on `Get`: flag, env, config file, key/value store.
If you have confidence some of your configs won't change, you can use piper's `IGet*` methods. Piper will build a configs cache on `Load`,
and those `IGet*` methods will get config from cache directly.
```shell
goos: darwin
goarch: amd64
pkg: github.com/Yiling-J/piper/integration
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz

BenchmarkGet-12           895533              1278 ns/op
BenchmarkIGet-12        19141876                61.35 ns/op
```
