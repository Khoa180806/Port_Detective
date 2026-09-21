package i18n

import (
	"testing"
)

func TestI18nBasic(t *testing.T) {
	// Register sample messages for testing
	RegisterMessages("en", MessageMap{
		"test.hello": "Hello %s",
		"test.only_en": "English only",
	})
	RegisterMessages("vi", MessageMap{
		"test.hello": "Xin chào %s",
	})

	// Default should be "en"
	if err := SetLang("en"); err != nil {
		t.Fatalf("unexpected error setting lang en: %v", err)
	}

	if got := Tr("test.hello", "Alice"); got != "Hello Alice" {
		t.Errorf("expected 'Hello Alice', got '%s'", got)
	}

	// Switch to "vi"
	if err := SetLang("vi"); err != nil {
		t.Fatalf("unexpected error setting lang vi: %v", err)
	}
	if GetLang() != "vi" {
		t.Errorf("expected currentLang 'vi', got '%s'", GetLang())
	}

	if got := Tr("test.hello", "Alice"); got != "Xin chào Alice" {
		t.Errorf("expected 'Xin chào Alice', got '%s'", got)
	}

	// Fallback to "en" when key is missing in "vi"
	if got := T("test.only_en"); got != "English only" {
		t.Errorf("expected fallback 'English only', got '%s'", got)
	}

	// Fallback to key when key is not registered anywhere
	if got := T("unknown.key"); got != "unknown.key" {
		t.Errorf("expected fallback to key name 'unknown.key', got '%s'", got)
	}

	// Test unsupported language
	if err := SetLang("fr"); err == nil {
		t.Errorf("expected error setting unsupported lang 'fr', got nil")
	}

	// Reset to "en"
	_ = SetLang("en")
}
