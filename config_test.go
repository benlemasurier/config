package config

import (
	"fmt"
	"os"
	"testing"
)

const (
	ConfigTestPrefix = "CONFIG_TEST_FOO_BAR"
)

func TestSet(t *testing.T) {
	conf := New(ConfigTestPrefix)

	k := "ZOMG!"
	v := "foo"
	err := conf.Set(k, v)
	if err != nil {
		t.Fatal(err)
	}

	fullKey := fmt.Sprintf("%s_%s", ConfigTestPrefix, k)
	result := os.Getenv(fullKey)
	if result != v {
		t.Errorf("expected value to be '%s', instead got '%s'", v, result)
	}

	err = conf.Unset(k)
	if err != nil {
		t.Error(err)
	}
}

func TestGet(t *testing.T) {
	conf := New(ConfigTestPrefix)

	k := "ZOMG!"
	v := "foo"
	fullKey := fmt.Sprintf("%s_%s", ConfigTestPrefix, k)

	err := os.Setenv(fullKey, v)
	if err != nil {
		t.Error(err)
	}

	result := conf.Get(k)
	if result != v {
		t.Errorf("expected value to be '%s', instead got '%s'", v, result)
	}

	err = conf.Unset(k)
	if err != nil {
		t.Error(err)
	}
}

func TestUnset(t *testing.T) {
	conf := New(ConfigTestPrefix)

	k := "ZOMG!"
	v := "foo"
	fullKey := fmt.Sprintf("%s_%s", ConfigTestPrefix, k)

	err := os.Setenv(fullKey, v)
	if err != nil {
		t.Fatal(err)
	}

	err = conf.Unset(k)
	if err != nil {
		t.Fatal(err)
	}

	if os.Getenv(fullKey) != "" {
		t.Errorf("expected %s environment variable to be unset", fullKey)
	}
}

func TestRequire(t *testing.T) {
	conf := New(ConfigTestPrefix)

	k := "ZOMG!"
	v := "foo"

	var keys = make([]string, 0)
	for i := 0; i < 100; i++ {
		keys = append(keys, fmt.Sprintf("%s-%d", k, i))
	}

	for _, key := range keys {
		if err := conf.Set(key, v); err != nil {
			t.Error(err)
		}
	}

	err := conf.Require(keys...)
	if err != nil {
		t.Error(err)
	}

	for _, key := range keys {
		if err := conf.Unset(key); err != nil {
			t.Error(err)
		}
	}
}

func TestRequireInvalid(t *testing.T) {
	conf := New(ConfigTestPrefix)

	if err := conf.Require("foo-bar-baz"); err == nil {
		t.Error("expected unset key to return error")
	}
}
