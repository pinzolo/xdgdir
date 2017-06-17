package xdgdir

import (
	"os"
	"testing"
)

func TestNewApp(t *testing.T) {
	app := NewApp("test")
	if app.Name != "test" {
		t.Errorf("expected test, but got %s", app.Name)
	}
}

func TestAppConfigDir(t *testing.T) {
	app := NewApp("test")
	table := []struct {
		xdgHome     string
		home        string
		userProfile string
		expected    string
		err         bool
	}{
		{"/x", "/y", "/z", "/x/test", false},
		{"", "/y", "/z", "/y/.config/test", false},
		{"", "", "/z", "/z/.config/test", false},
		{"", "", "", "", true},
	}

	for _, tbl := range table {
		os.Setenv("XDG_CONFIG_HOME", tbl.xdgHome)
		os.Setenv("HOME", tbl.home)
		os.Setenv("USERPROFILE", tbl.userProfile)
		dir, err := app.ConfigDir()
		if tbl.err {
			if err == nil {
				t.Error("should raise error, but not raised")
			}
		} else {
			if err != nil {
				t.Error(err)
			}
			if dir != tbl.expected {
				t.Errorf("expected %s, but got %s", tbl.expected, dir)
			}
		}
	}
}

func TestAppDataDir(t *testing.T) {
	app := NewApp("test")
	table := []struct {
		xdgHome     string
		home        string
		userProfile string
		expected    string
		err         bool
	}{
		{"/x", "/y", "/z", "/x/test", false},
		{"", "/y", "/z", "/y/.local/share/test", false},
		{"", "", "/z", "/z/.local/share/test", false},
		{"", "", "", "", true},
	}

	for _, tbl := range table {
		os.Setenv("XDG_DATA_HOME", tbl.xdgHome)
		os.Setenv("HOME", tbl.home)
		os.Setenv("USERPROFILE", tbl.userProfile)
		dir, err := app.DataDir()
		if tbl.err {
			if err == nil {
				t.Error("should raise error, but not raised")
			}
		} else {
			if err != nil {
				t.Error(err)
			}
			if dir != tbl.expected {
				t.Errorf("expected %s, but got %s", tbl.expected, dir)
			}
		}
	}
}

func TestAppCacheDir(t *testing.T) {
	app := NewApp("test")
	table := []struct {
		xdgHome     string
		home        string
		userProfile string
		expected    string
		err         bool
	}{
		{"/x", "/y", "/z", "/x/test", false},
		{"", "/y", "/z", "/y/.cache/test", false},
		{"", "", "/z", "/z/.cache/test", false},
		{"", "", "", "", true},
	}

	for _, tbl := range table {
		os.Setenv("XDG_CACHE_HOME", tbl.xdgHome)
		os.Setenv("HOME", tbl.home)
		os.Setenv("USERPROFILE", tbl.userProfile)
		dir, err := app.CacheDir()
		if tbl.err {
			if err == nil {
				t.Error("should raise error, but not raised")
			}
		} else {
			if err != nil {
				t.Error(err)
			}
			if dir != tbl.expected {
				t.Errorf("expected %s, but got %s", tbl.expected, dir)
			}
		}
	}
}

func TestAppRuntimeDir(t *testing.T) {
	app := NewApp("test")
	if app.RuntimeDir() == "" {
		t.Error("runtime dir should be not empty")
	}

	os.Setenv("XDG_RUNTIME_DIR", "/x")
	if dir := app.RuntimeDir(); dir != "/x/test" {
		t.Errorf("expected /x/test, but got %s", dir)
	}
}

func TestAppConfigFile(t *testing.T) {
	app := NewApp("test")
	name := "config.json"
	table := []struct {
		xdgHome     string
		home        string
		userProfile string
		expected    string
		err         bool
	}{
		{"/x", "/y", "/z", "/x/test/config.json", false},
		{"", "/y", "/z", "/y/.config/test/config.json", false},
		{"", "", "/z", "/z/.config/test/config.json", false},
		{"", "", "", "", true},
	}

	for _, tbl := range table {
		os.Setenv("XDG_CONFIG_HOME", tbl.xdgHome)
		os.Setenv("HOME", tbl.home)
		os.Setenv("USERPROFILE", tbl.userProfile)
		dir, err := app.ConfigFile(name)
		if tbl.err {
			if err == nil {
				t.Error("should raise error, but not raised")
			}
		} else {
			if err != nil {
				t.Error(err)
			}
			if dir != tbl.expected {
				t.Errorf("expected %s, but got %s", tbl.expected, dir)
			}
		}
	}
}

func TestAppDataFile(t *testing.T) {
	app := NewApp("test")
	name := "data.json"
	table := []struct {
		xdgHome     string
		home        string
		userProfile string
		expected    string
		err         bool
	}{
		{"/x", "/y", "/z", "/x/test/data.json", false},
		{"", "/y", "/z", "/y/.local/share/test/data.json", false},
		{"", "", "/z", "/z/.local/share/test/data.json", false},
		{"", "", "", "", true},
	}

	for _, tbl := range table {
		os.Setenv("XDG_DATA_HOME", tbl.xdgHome)
		os.Setenv("HOME", tbl.home)
		os.Setenv("USERPROFILE", tbl.userProfile)
		dir, err := app.DataFile(name)
		if tbl.err {
			if err == nil {
				t.Error("should raise error, but not raised")
			}
		} else {
			if err != nil {
				t.Error(err)
			}
			if dir != tbl.expected {
				t.Errorf("expected %s, but got %s", tbl.expected, dir)
			}
		}
	}
}

func TestAppCacheFile(t *testing.T) {
	app := NewApp("test")
	name := "cache.json"
	table := []struct {
		xdgHome     string
		home        string
		userProfile string
		expected    string
		err         bool
	}{
		{"/x", "/y", "/z", "/x/test/cache.json", false},
		{"", "/y", "/z", "/y/.cache/test/cache.json", false},
		{"", "", "/z", "/z/.cache/test/cache.json", false},
		{"", "", "", "", true},
	}

	for _, tbl := range table {
		os.Setenv("XDG_CACHE_HOME", tbl.xdgHome)
		os.Setenv("HOME", tbl.home)
		os.Setenv("USERPROFILE", tbl.userProfile)
		dir, err := app.CacheFile(name)
		if tbl.err {
			if err == nil {
				t.Error("should raise error, but not raised")
			}
		} else {
			if err != nil {
				t.Error(err)
			}
			if dir != tbl.expected {
				t.Errorf("expected %s, but got %s", tbl.expected, dir)
			}
		}
	}
}

func TestAppRuntimeFile(t *testing.T) {
	app := NewApp("test")
	name := "runtime.pid"
	if app.RuntimeFile(name) == "" {
		t.Error("runtime dir should be not empty")
	}

	os.Setenv("XDG_RUNTIME_DIR", "/x")
	if dir := app.RuntimeFile(name); dir != "/x/test/runtime.pid" {
		t.Errorf("expected /x/test, but got %s", dir)
	}
}
