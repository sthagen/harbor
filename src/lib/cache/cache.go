// Copyright Project Harbor Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cache

import (
	"errors"
	"fmt"
	"net/url"
	"sync"
	"time"

	"github.com/goharbor/harbor/src/lib"
	"github.com/goharbor/harbor/src/lib/log"
)

var (
	// ErrNotFound error returns when the key value not found in the cache
	ErrNotFound = errors.New("key not found")
)

// Cache cache interface
type Cache interface {
	// Contains returns true if key exists
	Contains(key string) bool

	// Delete delete item from cache by key
	Delete(key string) error

	// Fetch retrieve the cached key value
	Fetch(key string, value interface{}) error

	// Ping ping the cache
	Ping() error

	// Save cache the value by key
	Save(key string, value interface{}, expiration ...time.Duration) error
}

var (
	factories   = map[string]func(opts Options) (Cache, error){}
	factoriesMu sync.RWMutex
)

// Register register cache factory for type
func Register(typ string, factory func(opts Options) (Cache, error)) {
	factoriesMu.Lock()
	defer factoriesMu.Unlock()

	factories[typ] = factory
}

// New returns cache from addr
func New(typ string, opt ...Option) (Cache, error) {
	opts := newOptions(opt...)
	opts.Codec = codec // use the default codec for the cache

	factoriesMu.Lock()
	defer factoriesMu.Unlock()

	factory, ok := factories[typ]
	if !ok {
		return nil, fmt.Errorf("cache type %s not support", typ)
	}

	return factory(opts)
}

var (
	cache Cache
)

// Initialize initialize the default cache from the addr
func Initialize(typ, addr string) error {
	c, err := New(typ, Address(addr), Prefix("cache:"))
	if err != nil {
		return err
	}

	redactedAddr := addr
	if u, err := url.Parse(addr); err == nil {
		redactedAddr = redacted(u)
	}

	options := []lib.RetryOption{
		lib.RetryInitialInterval(time.Millisecond * 500),
		lib.RetryMaxInterval(time.Second * 10),
		lib.RetryTimeout(time.Minute),
		lib.RetryCallback(func(err error, sleep time.Duration) {
			log.Errorf("failed to ping %s, retry after %s : %v", redactedAddr, sleep, err)
		}),
	}

	if err := lib.RetryUntil(c.Ping, options...); err != nil {
		return err
	}

	cache = c

	return nil
}

// Default returns the default cache
func Default() Cache {
	return cache
}
