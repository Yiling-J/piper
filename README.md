# piper - Simple Wrapper For Viper

## Why
If you are familiar with Django, this is how Django settings module looks like:

```console {12-20}
└── settings
    ├── base.py
    ├── dev.py
	├── stage.py
	└── prod.py
```
dev will inherit from base, also stage will inherit from dev.

Start Django with selected settings:
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

```go
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
## Add Config Files
## Config Key Generation
## Use Piper
