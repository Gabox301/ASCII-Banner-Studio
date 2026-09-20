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

func TestEscapeKotlinDollar(t *testing.T) {
	out := escapeKotlin(`price $5`)
	if out != `price ${'$'}5` {
		t.Fatalf("expected escaped dollar, got %q", out)
	}
}

func TestEscapeRubyInterpolation(t *testing.T) {
	out := escapeRubyDoubleQuoted(`#{name}`)
	if out != `\#{name}` {
		t.Fatalf("expected escaped interpolation, got %q", out)
	}
}

func TestEscapeDartDollar(t *testing.T) {
	out := escapeDartSingleQuoted(`$x`)
	if out != `\$x` {
		t.Fatalf("expected escaped dollar, got %q", out)
	}
}

func TestEscapeShellSingleQuote(t *testing.T) {
	out := escapeShellSingleQuoted(`it's`)
	if out != `it'\''s` {
		t.Fatalf("expected shell-escaped quote, got %q", out)
	}
}

func TestEscapePowerShellSingleQuote(t *testing.T) {
	out := escapePowerShellSingleQuoted(`it's`)
	if out != `it''s` {
		t.Fatalf("expected doubled quote, got %q", out)
	}
}

func TestEscapeSwiftInterpolation(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{name: "interpolation opener is neutralized", input: `\(name)`, want: `\\(name)`},
		{name: "empty interpolation is neutralized", input: `\(`, want: `\\(`},
		{name: "doubled backslash before paren stays paired", input: `\\(x)`, want: `\\\\(x)`},
		{name: "plain parens are untouched", input: `(x)`, want: `(x)`},
		{name: "double quotes still escaped", input: `say "hi"`, want: `say \"hi\"`},
		{name: "lone backslash still escaped", input: `back\slash`, want: `back\\slash`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := escapeSwift(tt.input); got != tt.want {
				t.Fatalf("escapeSwift(%q): got %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}
