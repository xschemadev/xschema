package language

import (
	"fmt"
	"sort"
	"strings"
	"sync"
)

type Registry struct {
	mu sync.RWMutex

	byName      map[string]*Language
	bySchemaURL map[string]*Language
	bySchemaExt map[string]*Language
	ignoreDirs  map[string]bool
}

func NewRegistry() *Registry {
	return &Registry{
		byName:      make(map[string]*Language),
		bySchemaURL: make(map[string]*Language),
		bySchemaExt: make(map[string]*Language),
		ignoreDirs:  make(map[string]bool),
	}
}

func (r *Registry) Register(lang Language) error {
	if strings.TrimSpace(lang.Name) == "" {
		return fmt.Errorf("language name is required")
	}
	if strings.TrimSpace(lang.SchemaURL) == "" {
		return fmt.Errorf("language %q schema url is required", lang.Name)
	}
	if !IsXSchemaURL(lang.SchemaURL) {
		return fmt.Errorf("language %q schema url must start with %q: %s", lang.Name, XSchemaBaseURL, lang.SchemaURL)
	}

	expectedExt := strings.TrimPrefix(lang.SchemaURL, XSchemaBaseURL)
	if strings.TrimSpace(expectedExt) == "" {
		return fmt.Errorf("language %q schema url must include a filename: %s", lang.Name, lang.SchemaURL)
	}
	if lang.SchemaExt == "" {
		lang.SchemaExt = expectedExt
	}
	if lang.SchemaExt != expectedExt {
		return fmt.Errorf("language %q schema ext must match schema url filename: schemaExt=%q expected=%q", lang.Name, lang.SchemaExt, expectedExt)
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.byName[lang.Name]; exists {
		return fmt.Errorf("duplicate language name: %s", lang.Name)
	}
	if _, exists := r.bySchemaURL[lang.SchemaURL]; exists {
		return fmt.Errorf("duplicate language schema url mapping: %s", lang.SchemaURL)
	}
	if _, exists := r.bySchemaExt[lang.SchemaExt]; exists {
		return fmt.Errorf("duplicate language schema ext mapping: %s", lang.SchemaExt)
	}

	registered := lang
	r.byName[registered.Name] = &registered
	r.bySchemaURL[registered.SchemaURL] = &registered
	r.bySchemaExt[registered.SchemaExt] = &registered

	for _, dir := range registered.IgnoreDirs {
		if strings.TrimSpace(dir) == "" {
			continue
		}
		r.ignoreDirs[dir] = true
	}

	return nil
}

func (r *Registry) ByName(name string) *Language {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.byName[name]
}

// Unregister removes a language from the registry by name.
func (r *Registry) Unregister(name string) {
	r.mu.Lock()
	defer r.mu.Unlock()

	lang, exists := r.byName[name]
	if !exists {
		return
	}

	delete(r.byName, name)
	delete(r.bySchemaURL, lang.SchemaURL)
	delete(r.bySchemaExt, lang.SchemaExt)
}

func (r *Registry) BySchemaURL(url string) *Language {
	if !IsXSchemaURL(url) {
		return nil
	}

	ext := strings.TrimPrefix(url, XSchemaBaseURL)

	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.bySchemaExt[ext]
}

func (r *Registry) AllIgnoreDirs() map[string]bool {
	r.mu.RLock()
	defer r.mu.RUnlock()

	out := make(map[string]bool, len(r.ignoreDirs))
	for dir := range r.ignoreDirs {
		out[dir] = true
	}
	return out
}

func (r *Registry) SupportedLanguages() []string {
	r.mu.RLock()
	defer r.mu.RUnlock()

	langs := make([]string, 0, len(r.byName))
	for name := range r.byName {
		langs = append(langs, name)
	}
	sort.Strings(langs)
	return langs
}

var defaultRegistry = NewRegistry()

// ResetForTests resets the registry state to empty.
// Used to prevent global registry leakage between tests.
func ResetForTests() {
	defaultRegistry = NewRegistry()
}

func Register(lang Language) error {
	return defaultRegistry.Register(lang)
}

// Unregister removes a language from the registry by name.
// Used in tests to cleanup after registering test languages.
func Unregister(name string) {
	defaultRegistry.Unregister(name)
}

func SupportedLanguages() []string {
	return defaultRegistry.SupportedLanguages()
}
