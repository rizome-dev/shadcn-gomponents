package example

import (
	"errors"
	"testing"
)

func TestNewClient(t *testing.T) {
	tests := []struct {
		name    string
		config  *Config
		wantErr error
	}{
		{
			name:    "nil config",
			config:  nil,
			wantErr: ErrInvalidInput,
		},
		{
			name:    "empty config",
			config:  &Config{},
			wantErr: nil,
		},
		{
			name: "valid config",
			config: &Config{
				Name:    "test",
				Verbose: true,
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewClient(tt.config)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("NewClient() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClient_Process(t *testing.T) {
	client, err := NewClient(&Config{Name: "test"})
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	tests := []struct {
		name    string
		input   string
		want    string
		wantErr bool
	}{
		{
			name:    "empty input",
			input:   "",
			want:    "",
			wantErr: true,
		},
		{
			name:    "valid input",
			input:   "hello",
			want:    "Processed by test: hello",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := client.Process(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.Process() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Client.Process() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkClient_Process(b *testing.B) {
	client, _ := NewClient(DefaultConfig())
	input := "benchmark test"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = client.Process(input)
	}
}