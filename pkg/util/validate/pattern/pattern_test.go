package pattern_test

import (
	"testing"

	"github.com/4thel00z/templatectl/pkg/util/validate/pattern"
)

func TestUnixPathPattern(t *testing.T) {
	tests := []struct {
		String string
		Valid  bool
	}{
		{"", false},
		{"/", true},
		{"/root", true},
		{"/tmp-dir", true},
		{"/tmp-dir/new_dir", true},
		{"/TMP/dir", true},
		{"rel/dir", true},
	}

	for _, test := range tests {
		if ok := pattern.UnixPath.MatchString(test.String); ok != test.Valid {
			t.Errorf("pattern.UnixPath.MatchString(%q) expected to be %v", test.String, test.Valid)
		}
	}
}

func TestAlphanumericPattern(t *testing.T) {
	tests := []struct {
		String string
		Valid  bool
	}{
		{" ", false},
		{"/", false},
		{"root", true},
		{"tmp-dir", false},
		{"TMPDIR", true},
		{"L33T", true},
		{"L@@T", false},
	}

	for _, test := range tests {
		if ok := pattern.Alphanumeric.MatchString(test.String); ok != test.Valid {
			t.Errorf("pattern.Alphanumeric.MatchString(%q) expected to be %v", test.String, test.Valid)
		}
	}
}

func TestAlphanumericExtPattern(t *testing.T) {
	tests := []struct {
		String string
		Valid  bool
	}{
		{" ", false},
		{"/", false},
		{"root", true},
		{"tmp-dir", true},
		{"tmp-dir_now", true},
		{"TMPDIR", true},
		{"L33T", true},
		{"L@@T", false},
	}

	for _, test := range tests {
		if ok := pattern.AlphanumericExt.MatchString(test.String); ok != test.Valid {
			t.Errorf("pattern.AlphanumericExt.MatchString(%q) expected to be %v", test.String, test.Valid)
		}
	}
}

func TestIntegerPattern(t *testing.T) {
	tests := []struct {
		String string
		Valid  bool
	}{
		{"", false},
		{" ", false},
		{"/", false},
		{"root", false},
		{"L33T", false},
		{"", false},
	}

	for _, test := range tests {
		if ok := pattern.Numeric.MatchString(test.String); ok != test.Valid {
			t.Errorf("pattern.Numeric.MatchString(%q) expected to be %v", test.String, test.Valid)
		}
	}
}

func TestURLPattern(t *testing.T) {
	tests := []struct {
		String string
		Valid  bool
	}{
		{"", false},
		{" ", false},
		{"/", false},
		{"http://", false},
		{"http://github.com/4thel00z/templatectl", true},
		{"https://github.com/4thel00z/templatectl", true},
		{"github.com/4thel00z/templatectl", true},
		{"rawcontent.github.com/4thel00z/templatectl", true},
		{"github.com:80/4thel00z/templatectl", true},
	}

	for _, test := range tests {
		if ok := pattern.URL.MatchString(test.String); ok != test.Valid {
			t.Errorf("pattern.URL.MatchString(%q) expected to be %v", test.String, test.Valid)
		}
	}
}
