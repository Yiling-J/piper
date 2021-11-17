package piper

import (
	"embed"
	"errors"
	"path/filepath"
	"time"

	"github.com/spf13/afero"
	"github.com/spf13/viper"
)

var (
	p *Piper
)

func init() {
	p = &Piper{V: viper.GetViper()}
}

// sliceEqual tells whether a and b contain the same elements.
// A nil argument is equivalent to an empty slice.
func sliceEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func sliceContains(a []string, b string) bool {
	for _, v := range a {
		if v == b {
			return true
		}
	}
	return false
}

func copyMap(original map[string]bool) map[string]bool {
	new := make(map[string]bool)
	for key, value := range original {
		new[key] = value
	}
	return new
}

type Piper struct {
	V        *viper.Viper
	fs       embed.FS
	dir      string
	imported []string
}

func Reset() {
	viper.Reset()
	p = &Piper{V: viper.GetViper()}
}

func V() *viper.Viper {
	return viper.GetViper()
}

func New() *Piper {
	return &Piper{V: viper.New()}
}

func SetFS(fs embed.FS) {
	p.SetFS(fs)
}

func (p *Piper) SetFS(fs embed.FS) {
	p.fs = fs
}

func ReadInConfig() error {
	return p.ReadInConfig()
}

func (p *Piper) ReadInConfig() error {
	p.V.SetFs(afero.NewOsFs())
	err := p.V.ReadInConfig()
	if err == nil {
		return err
	}
	p.V.SetFs(afero.FromIOFS{FS: p.fs})
	err = p.V.ReadInConfig()
	return err
}

func (p *Piper) MergeInConfig() error {
	p.V.SetFs(afero.NewOsFs())
	err := p.V.MergeInConfig()
	if err == nil {
		return err
	}
	p.V.SetFs(afero.FromIOFS{FS: p.fs})
	err = p.V.MergeInConfig()
	return err
}

func (p *Piper) Get(key string) interface{} {
	return p.V.Get(key)
}

func Get(key string) interface{} {
	return p.V.Get(key)
}

func (p *Piper) GetBool(key string) bool {
	return p.V.GetBool(key)
}

func GetBool(key string) bool {
	return p.V.GetBool(key)
}

func (p *Piper) GetDuration(key string) time.Duration {
	return p.V.GetDuration(key)
}

func GetDuration(key string) time.Duration {
	return p.V.GetDuration(key)
}

func (p *Piper) GetFloat64(key string) float64 {
	return p.V.GetFloat64(key)
}

func GetFloat64(key string) float64 {
	return p.V.GetFloat64(key)
}

func (p *Piper) GetInt(key string) int {
	return p.V.GetInt(key)
}

func GetInt(key string) int {
	return p.V.GetInt(key)
}

func (p *Piper) GetInt32(key string) int32 {
	return p.V.GetInt32(key)
}

func GetInt32(key string) int32 {
	return p.V.GetInt32(key)
}

func (p *Piper) GetInt64(key string) int64 {
	return p.V.GetInt64(key)
}

func GetInt64(key string) int64 {
	return p.V.GetInt64(key)
}

func (p *Piper) GetUint(key string) uint {
	return p.V.GetUint(key)
}

func GetUint(key string) uint {
	return p.V.GetUint(key)
}

func (p *Piper) GetUint32(key string) uint32 {
	return p.V.GetUint32(key)
}

func GetUint32(key string) uint32 {
	return p.V.GetUint32(key)
}

func (p *Piper) GetUint64(key string) uint64 {
	return p.V.GetUint64(key)
}

func GetUint64(key string) uint64 {
	return p.V.GetUint64(key)
}

func (p *Piper) GetIntSlice(key string) []int {
	return p.V.GetIntSlice(key)
}

func GetIntSlice(key string) []int {
	return p.V.GetIntSlice(key)
}

func (p *Piper) GetString(key string) string {
	return p.V.GetString(key)
}

func GetString(key string) string {
	return p.V.GetString(key)
}

func (p *Piper) GetStringMap(key string) map[string]interface{} {
	return p.V.GetStringMap(key)
}

func GetStringMap(key string) map[string]interface{} {
	return p.V.GetStringMap(key)
}

func (p *Piper) GetStringSlice(key string) []string {
	return p.V.GetStringSlice(key)
}

func GetStringSlice(key string) []string {
	return p.V.GetStringSlice(key)
}

func (p *Piper) GetStringMapString(key string) map[string]string {
	return p.V.GetStringMapString(key)
}

func GetStringMapString(key string) map[string]string {
	return p.V.GetStringMapString(key)
}

func (p *Piper) GetStringMapStringSlice(key string) map[string][]string {
	return p.V.GetStringMapStringSlice(key)
}

func GetStringMapStringSlice(key string) map[string][]string {
	return p.V.GetStringMapStringSlice(key)
}

func Load(name string) error {
	return p.Load(name)
}

func (p *Piper) loadChild(name string, parentImports []string, m map[string]bool) error {
	path := p.dir + "/" + name
	p.V.SetConfigFile(path)
	err := p.MergeInConfig()

	if err != nil {
		return err
	}

	imports := p.V.GetStringSlice("pp_imports")
	if !sliceEqual(imports, parentImports) {
		for _, i := range imports {
			if _, ok := m[i]; ok {
				return errors.New("cycle import found")
			}
			if sliceContains(p.imported, i) {
				continue
			}
			mn := copyMap(m)
			mn[i] = true
			err = p.loadChild(i, imports, mn)
			if err != nil {
				return err
			}
			p.imported = append(p.imported, i)
		}
	}
	// merge original file back after all imports resolved
	p.V.SetConfigFile(path)
	err = p.MergeInConfig()
	if err != nil {
		return err
	}

	return nil

}

func (p *Piper) Load(name string) error {
	path := name
	p.dir = filepath.Dir(name)
	p.V.SetConfigFile(name)
	err := p.ReadInConfig()

	if err != nil {
		return err
	}

	imports := p.V.GetStringSlice("pp_imports")
	for _, i := range imports {
		m := map[string]bool{name: true}
		err = p.loadChild(i, imports, m)
		if err != nil {
			return err
		}
	}
	// merge original file back after all imports resolved
	p.V.SetConfigFile(path)
	err = p.MergeInConfig()
	if err != nil {
		return err
	}

	return nil

}
