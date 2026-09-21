package i18n

import (
	"testing"
)

func TestI18nBasic(t *testing.T) {
	// Register sample messages for testing
	RegisterMessages("en", MessageMap{
		"test.hello":   "Hello %s",
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

func TestDictionariesIntegrity(t *testing.T) {
	// Verify that en and vi dictionaries have matching keys
	if len(enMessages) == 0 {
		t.Fatal("enMessages is empty")
	}
	if len(viMessages) == 0 {
		t.Fatal("viMessages is empty")
	}

	for k := range enMessages {
		if _, exists := viMessages[k]; !exists {
			t.Errorf("key '%s' present in enMessages but missing in viMessages", k)
		}
	}

	for k := range viMessages {
		if _, exists := enMessages[k]; !exists {
			t.Errorf("key '%s' present in viMessages but missing in enMessages", k)
		}
	}

	// Verify key translations in both languages
	_ = SetLang("en")
	if got := T("check.short"); got != "Check the process listening on a specified port" {
		t.Errorf("unexpected EN check.short: %s", got)
	}

	_ = SetLang("vi")
	if got := T("check.short"); got != "Kiểm tra process đang chạy trên một port" {
		t.Errorf("unexpected VI check.short: %s", got)
	}

	// Reset to "en"
	_ = SetLang("en")
}
