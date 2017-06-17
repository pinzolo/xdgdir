package xdgdir

import "path/filepath"

// App is application name in XDG Base directories.
type App struct {
	// Name of app
	Name string
}

// NewApp returns new app object that has given name.
func NewApp(name string) App {
	return App{Name: name}
}

// ConfigDir returns base directory path of app's config files.
//
// 1. If XDG_CONFIG_HOME envvar is defiend, returns $XDG_CONFIG_HOME/{{AppName}}.
// 2. IF HOME envvar is defiend, returns $HOME/.config/{{AppName}}
// 3. IF USERPROFILE envvar is defiend, returns $USERPROFILE/.config/{{AppName}} (for Windows)
func (a App) ConfigDir() (string, error) {
	return joinedPath(a.Name, ConfigDir)
}

// ConfigFile returns file path of app's config file that has given file name.
//
// 1. If XDG_CONFIG_HOME envvar is defiend, returns $XDG_CONFIG_HOME/{{AppName}}/{{name}}.
// 2. IF HOME envvar is defiend, returns $HOME/.config/{{AppName}}/{{name}}
// 3. IF USERPROFILE envvar is defiend, returns $USERPROFILE/.config/{{AppName}}/{{name}} (for Windows)
func (a App) ConfigFile(name string) (string, error) {
	return joinedPath(name, a.ConfigDir)
}

// DataDir returns base directory path of app's data files.
//
// 1. If XDG_data_HOME envvar is defiend, returns $XDG_DATA_HOME/{{AppName}}.
// 2. IF HOME envvar is defiend, returns $HOME/.local/share/{{AppName}}
// 3. IF USERPROFILE envvar is defiend, returns $USERPROFILE/.local/share/{{AppName}} (for Windows)
func (a App) DataDir() (string, error) {
	return joinedPath(a.Name, DataDir)
}

// DataFile returns file path of app's data file that has given file name.
//
// 1. If XDG_data_HOME envvar is defiend, returns $XDG_DATA_HOME/{{AppName}}/{{name}}.
// 2. IF HOME envvar is defiend, returns $HOME/.local/share/{{AppName}}/{{name}}
// 3. IF USERPROFILE envvar is defiend, returns $USERPROFILE/.local/share/{{AppName}}/{{name}} (for Windows)
func (a App) DataFile(name string) (string, error) {
	return joinedPath(name, a.DataDir)
}

// CacheDir returns base directory path of app's cache files.
//
// 1. If XDG_cache_HOME envvar is defiend, returns $XDG_CACHE_HOME/{{AppName}}.
// 2. IF HOME envvar is defiend, returns $HOME/.cache/{{AppName}}
// 3. IF USERPROFILE envvar is defiend, returns $USERPROFILE/.cache/{{AppName}} (for Windows)
func (a App) CacheDir() (string, error) {
	return joinedPath(a.Name, CacheDir)
}

// CacheFile returns file path of app's cache file that has given file name.
//
// 1. If XDG_cache_HOME envvar is defiend, returns $XDG_CACHE_HOME/{{AppName}}/{{name}}.
// 2. IF HOME envvar is defiend, returns $HOME/.cache/{{AppName}}/{{name}}
// 3. IF USERPROFILE envvar is defiend, returns $USERPROFILE/.cache/{{AppName}}/{{name}} (for Windows)
func (a App) CacheFile(name string) (string, error) {
	return joinedPath(name, a.CacheDir)
}

// RuntimeDir returns base directory path of app's runtime.
//
// 1. If XDG_RUNTIME_DIR envvar is defiend, returns $XDG_RUNTIME_DIR/{{AppName}}.
// 2. Returns temporary directory path that has subdirectory named AppName.
func (a App) RuntimeDir() string {
	return filepath.Join(RuntimeDir(), a.Name)
}

// RuntimeFile returns file path of app's runtime file that has given file name.
//
// 1. If XDG_RUNTIME_DIR envvar is defiend, returns $XDG_RUNTIME_DIR/{{AppName}}/{{name}}.
// 2. Returns temporary directory path that has subdirectory named AppName.
func (a App) RuntimeFile(name string) string {
	return filepath.Join(a.RuntimeDir(), name)
}

func joinedPath(name string, f func() (string, error)) (string, error) {
	dir, err := f()
	if err != nil {
		return "", err
	}

	return filepath.Join(dir, name), nil
}
