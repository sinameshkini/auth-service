package postgres

import (
	"gorm.io/gorm"
	"testing"
)

func TestConnect(t *testing.T) {
	type args struct {
		config *Config
	}
	tests := []struct {
		name    string
		args    args
		wantDb  *gorm.DB
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				config: &Config{
					Host:     "localhost",
					Port:     "7032",
					Username: "auth_user",
					Password: "ajkvxHD3igy842LXNOUS",
					Database: "auth",
					TimeZone: "Asia/Tehran",
				},
			},
		},
		{
			name: "invalid config",
			args: args{
				config: &Config{
					Host: "localhost",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotDb, err := Connect(tt.args.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("Connect() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			_ = gotDb
		})
	}
}
