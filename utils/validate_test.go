package Utils

import "testing"

func TestIsChinese(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"张", true},
		{"Hello", false},
		{"123", false},
	}

	for _, tt := range tests {
		got := IsChinese(tt.input)
		if got != tt.want {
			t.Errorf("IsChinese(%q) = %v; want %v", tt.input, got, tt.want)
		}
	}
}

func TestIsEnglish(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"你好", false},
		{"Hello", true},
		{"123", false},
	}

	for _, tt := range tests {
		got := IsEnglish(tt.input)
		if got != tt.want {
			t.Errorf("IsEnglish(%q) = %v; want %v", tt.input, got, tt.want)
		}
	}
}

func TestIsPhoneNum(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"13812345678", true},
		{"hello", false},
		{"1234567890", false},
	}

	for _, tt := range tests {
		got := IsPhoneNum(tt.input)
		if got != tt.want {
			t.Errorf("IsPhoneNum(%q) = %v; want %v", tt.input, got, tt.want)
		}
	}
}

func TestIsNum(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"123", true},
		{"hello", false},
		{"123abc", false},
	}

	for _, tt := range tests {
		got := IsNum(tt.input)
		if got != tt.want {
			t.Errorf("IsNum(%q) = %v; want %v", tt.input, got, tt.want)
		}
	}
}

func TestIsIdCard(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"442000200301123291", true},
		{"44200020030112329X", true},
		{"44200020030112329", false},
		{"4420002003011232", false},
		{"123456789012345", true},
		{"12345678901234", false},
	}

	for _, tt := range tests {
		got, _ := IsIdCard(tt.input)
		if got != tt.want {
			t.Errorf("IsIdCard(%q) = %v; want %v", tt.input, got, tt.want)
		}
	}
}

func TestIsIdCard_Error(t *testing.T) {
	tests := []struct {
		input string
	}{
		{"44200020030112329"},
		{"4420002003011232"},
		{"1234567890123"},
	}

	for _, tt := range tests {
		_, err := IsIdCard(tt.input)
		if err == nil {
			t.Errorf("IsIdCard(%q) should return an error", tt.input)
		}
	}
}
