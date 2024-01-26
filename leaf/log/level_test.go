package log

import "testing"

func TestLevel_String(t *testing.T) {
	tests := []struct {
		name string
		l    Level
		want string
	}{
		{
			name: "trace",
			l:    TraceLevel,
			want: "trace",
		},
		{
			name: "debug",
			l:    DebugLevel,
			want: "debug",
		},
		{
			name: "info",
			l:    InfoLevel,
			want: "info",
		},
		{
			name: "warn",
			l:    WarnLevel,
			want: "warn",
		},
		{
			name: "error",
			l:    ErrorLevel,
			want: "error",
		},
		{
			name: "fatal",
			l:    FatalLevel,
			want: "fatal",
		},
		{
			name: "other",
			l:    10,
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetLevel(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want Level
	}{
		{
			name: "trace",
			want: TraceLevel,
			s:    "trace",
		},
		{
			name: "debug",
			want: DebugLevel,
			s:    "debug",
		},
		{
			name: "info",
			want: InfoLevel,
			s:    "info",
		},
		{
			name: "warn",
			want: WarnLevel,
			s:    "warn",
		},
		{
			name: "error",
			want: ErrorLevel,
			s:    "error",
		},
		{
			name: "fatal",
			want: FatalLevel,
			s:    "fatal",
		},
		{
			name: "other",
			want: 10,
			s:    "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, err := GetLevel(tt.s); err != nil || got != tt.want {
				t.Errorf("ParseLevel() = %v, want %v", got, tt.want)
			}
		})
	}
}
