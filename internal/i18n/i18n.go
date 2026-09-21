package i18n

import (
	"fmt"
	"strings"
	"sync"
)

// MessageMap maps message keys to localized string templates.
type MessageMap map[string]string

var (
	mu          sync.RWMutex
	currentLang = "en"
	// messages contains the translations for supported languages.
	// en.go and vi.go populate this map via init() or package-level variables.
	messages = make(map[string]MessageMap)
)

// SupportedLangs returns a slice of supported language codes.
func SupportedLangs() []string {
	return []string{"en", "vi"}
}

// IsSupported checks if a language code is supported.
func IsSupported(lang string) bool {
	normalized := strings.ToLower(strings.TrimSpace(lang))
	for _, l := range SupportedLangs() {
		if l == normalized {
			return true
		}
	}
	return false
}

// SetLang sets the current active language. Returns an error if unsupported.
func SetLang(lang string) error {
	normalized := strings.ToLower(strings.TrimSpace(lang))
	if !IsSupported(normalized) {
		return fmt.Errorf("unsupported language: '%s' (supported: %s)", lang, strings.Join(SupportedLangs(), ", "))
	}
	mu.Lock()
	defer mu.Unlock()
	currentLang = normalized
	return nil
}

// GetLang returns the current active language.
func GetLang() string {
	mu.RLock()
	defer mu.RUnlock()
	return currentLang
}

// RegisterMessages allows registering or extending translations for a specific language.
func RegisterMessages(lang string, msgs MessageMap) {
	mu.Lock()
	defer mu.Unlock()
	normalized := strings.ToLower(strings.TrimSpace(lang))
	if messages[normalized] == nil {
		messages[normalized] = make(MessageMap)
	}
	for k, v := range msgs {
		messages[normalized][k] = v
	}
}

// Tr translates a key and formats it using the provided arguments.
func Tr(key string, a ...any) string {
	res := T(key)
	if len(a) > 0 {
		return fmt.Sprintf(res, a...)
	}
	return res
}

// T translates a message key using the active language.
// Fallback hierarchy: currentLang -> "en" -> key itself.
func T(key string) string {
	mu.RLock()
	lang := currentLang
	msgMap, ok := messages[lang]
	var tmpl string
	found := false

	if ok {
		tmpl, found = msgMap[key]
	}

	// Fallback to "en" if not found in current language
	if !found && lang != "en" {
		if enMap, enOk := messages["en"]; enOk {
			tmpl, found = enMap[key]
		}
	}
	mu.RUnlock()

	// If still not found, return the key as a safe fallback (never panic)
	if !found {
		return key
	}

	return tmpl
}
