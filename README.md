# piper - Simple Wrapper For Viper

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
```console {12-20}
project
└── config
    ├── base.toml
    ├── dev.toml
    ├── stage.toml
    └── prod.toml

```

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

### Strategy I
embed your config folder to your code, single executable when you deploy.
```go
import (
	"github.com/Yiling-J/piper"
	"your_project/example/config"
)

//go:embed config/*
var configFS embed.FS

piper.Load("config/stage.toml")
author = piper.GetString(config.Author)
```

### Strategy II
copy config folder when you building docker image, so the true config folder exists with your executable.
```go
import (
	"github.com/Yiling-J/piper"
	"your_project/example/config"
)

piper.Load("config/stage.toml")
author = piper.GetString(config.Author)
```

### Strategy III
Override when you building docker image or deploy, eg: k8s ConfigMap replacement
```go
import (
	"github.com/Yiling-J/piper"
	"your_project/example/config"
)

//go:embed config/*
var configFS embed.FS

piper.Load("config/stage.toml")
author = piper.GetString(config.Author)
```
True folder will take precedence over embeded folder.
