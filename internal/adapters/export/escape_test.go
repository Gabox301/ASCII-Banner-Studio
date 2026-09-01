package export

import (
	"strings"
	"testing"
)

func TestEscapeSingleQuotedBackslash(t *testing.T) {
	out := escapeSingleQuoted(`a\b`)
	if !strings.Contains(out, `\\`) {
		t.Fatalf("expected escaped backslash, got %q", out)
	}
}

func TestEscapeSingleQuotedQuote(t *testing.T) {
	out := escapeSingleQuoted(`it's`)
	if !strings.Contains(out, `it\'`) {
		t.Fatalf("expected escaped single quote, got %q", out)
	}
}

func TestEscapeSingleQuotedNoSpecialChars(t *testing.T) {
	out := escapeSingleQuoted("abc")
	if out != "abc" {
		t.Fatalf("expected no change, got %q", out)
	}
}

func TestEscapeDoubleQuotedBackslash(t *testing.T) {
	out := escapeDoubleQuoted(`a\b`)
	if !strings.Contains(out, `\\`) {
		t.Fatalf("expected escaped backslash, got %q", out)
	}
}

func TestEscapeDoubleQuotedQuote(t *testing.T) {
	out := escapeDoubleQuoted(`say "hi"`)
	if !strings.Contains(out, `say \"hi\"`) {
		t.Fatalf("expected escaped double quote, got %q", out)
	}
}

func TestEscapeDoubleQuotedNoSpecialChars(t *testing.T) {
	out := escapeDoubleQuoted("abc")
	if out != "abc" {
		t.Fatalf("expected no change, got %q", out)
	}
}

func TestEscapeBacktick(t *testing.T) {
	out := escapeBacktick("`template`")
	if !strings.Contains(out, "\\`template\\`") {
		t.Fatalf("expected escaped backticks, got %q", out)
	}
}

func TestEscapeBacktickDollar(t *testing.T) {
	out := escapeBacktick("${var}")
	if !strings.Contains(out, `\${var}`) {
		t.Fatalf("expected escaped dollar sign, got %q", out)
	}
}

func TestEscapeBacktickNoSpecialChars(t *testing.T) {
	out := escapeBacktick("abc")
	if out != "abc" {
		t.Fatalf("expected no change, got %q", out)
	}
}
