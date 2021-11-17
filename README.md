# piper - viper wrapper with config inheritance and key generation

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

And this is how you access settings in Django:

```python
from django.conf import settings

author = settings.Author
```

I want to have similar experience with Viper, so here comes this wrapper. And this is the Piper way:

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
	author = piper.GetString(config.Author)
}
```

Check example folder for more details.

## Installation
## Add Configs
## Config Key Generation
## Use Piper
