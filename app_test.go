package xdgdir

import (
	"io/ioutil"
	"os"
	"strings"
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
		{"x", "y", "z", path("x", "test"), false},
		{"", "y", "z", path("y", ".config", "test"), false},
		{"", "", "z", path("z", ".config", "test"), false},
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
		{"x", "y", "z", path("x", "test"), false},
		{"", "y", "z", path("y", ".local", "share", "test"), false},
		{"", "", "z", path("z", ".local", "share", "test"), false},
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
		{"x", "y", "z", path("x", "test"), false},
		{"", "y", "z", path("y", ".cache", "test"), false},
		{"", "", "z", path("z", ".cache", "test"), false},
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
		{"x", "y", "z", path("x", "test", "config.json"), false},
		{"", "y", "z", path("y", ".config", "test", "config.json"), false},
		{"", "", "z", path("z", ".config", "test", "config.json"), false},
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
		{"x", "y", "z", path("x", "test", "data.json"), false},
		{"", "y", "z", path("y", ".local", "share", "test", "data.json"), false},
		{"", "", "z", path("z", ".local", "share", "test", "data.json"), false},
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
		{"x", "y", "z", path("x", "test", "cache.json"), false},
		{"", "y", "z", path("y", ".cache", "test", "cache.json"), false},
		{"", "", "z", path("z", ".cache", "test", "cache.json"), false},
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

	os.Setenv("XDG_RUNTIME_DIR", "x")
	if dir := app.RuntimeFile(name); dir != path("x", "test", "runtime.pid") {
		t.Errorf("expected /x/test, but got %s", dir)
	}
}

func TestAppFindConfigFile(t *testing.T) {
	app := NewApp("test")
	os.Setenv("XDG_CONFIG_HOME", path("testdata", "x"))
	os.Setenv("XDG_CONFIG_DIRS", join(path("testdata", "y"), path("testdata", "z")))
	table := []struct {
		name    string
		content string
		err     bool
	}{
		{"aaa.txt", path("testdata", "a", "test", "aaa.txt"), true},
		{"xxx.txt", path("testdata", "x", "test", "xxx.txt"), false},
		{"yyy.txt", path("testdata", "y", "test", "yyy.txt"), false},
		{"zzz.txt", path("testdata", "z", "test", "zzz.txt"), false},
	}
	for _, test := range table {
		f, err := app.FindConfigFile(test.name)
		if test.err {
			if err == nil {
				t.Error("should raise error, but not raised")
			}
			continue
		}

		if err != nil {
			t.Error(err)
			continue
		}
		s, err := openFile(f)
		if err != nil {
			t.Error(err)
			continue
		}
		if s != test.content {
			t.Error("find invalid file")
		}
	}
}

func TestAppFindDataFile(t *testing.T) {
	app := NewApp("test")
	os.Setenv("XDG_DATA_HOME", path("testdata", "x"))
	os.Setenv("XDG_DATA_DIRS", join(path("testdata", "y"), path("testdata", "z")))
	table := []struct {
		name    string
		content string
		err     bool
	}{
		{"aaa.txt", path("testdata", "a", "test", "aaa.txt"), true},
		{"xxx.txt", path("testdata", "x", "test", "xxx.txt"), false},
		{"yyy.txt", path("testdata", "y", "test", "yyy.txt"), false},
		{"zzz.txt", path("testdata", "z", "test", "zzz.txt"), false},
	}
	for _, test := range table {
		f, err := app.FindDataFile(test.name)
		if test.err {
			if err == nil {
				t.Error("should raise error, but not raised")
			}
			continue
		}

		if err != nil {
			t.Error(err)
			continue
		}
		s, err := openFile(f)
		if err != nil {
			t.Error(err)
			continue
		}
		if s != test.content {
			t.Error("find invalid file")
		}
	}
}

func openFile(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	p, err := ioutil.ReadAll(f)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(p)), nil
}

func join(elm ...string) string {
	return strings.Join(elm, string(os.PathListSeparator))
}
