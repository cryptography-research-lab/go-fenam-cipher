package fenam_cipher

import "testing"

func Test_asciiToBinary(t *testing.T) {
	type args struct {
		number int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			args: args{
				number: 'A',
			},
			want: "1000001",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intToBinaryString(tt.args.number); got != tt.want {
				t.Errorf("intToBinaryString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fromBinaryString(t *testing.T) {
	type args struct {
		binaryString string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			args: args{
				binaryString: "1010000",
			},
			want: 'P',
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fromBinaryString(tt.args.binaryString); got != tt.want {
				t.Errorf("fromBinaryString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEncrypt(t *testing.T) {
	type args struct {
		asciiText string
		security  []string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			args: args{
				asciiText: "HELLO",
				security:  []string{"STUDY"},
			},
			want: "00110110010001001100100010000010110",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Encrypt(tt.args.asciiText, tt.args.security...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Encrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Encrypt() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecrypt(t *testing.T) {
	type args struct {
		encryptBinaryText string
		securityKey       []string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			args: args{
				encryptBinaryText: "00110110010001001100100010000010110",
				securityKey:       []string{"STUDY"},
			},
			want: "HELLO",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Decrypt(tt.args.encryptBinaryText, tt.args.securityKey...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Decrypt() got = %v, want %v", got, tt.want)
			}
		})
	}
}
