package utils

import (
	"encoding/base64"
	"net/url"
	"reflect"
	"testing"
)

func TestHmac(t *testing.T) {
	type args struct {
		key  string
		data string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "testHmac", args: args{key: "key", data: "hello"}, want: "kwezuRXvtRcf8U2MtV%2B8x5jGwO8UVtZt7RpqpyOli3s%3D"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dig := Hmac(tt.args.key, tt.args.data)
			got := url.QueryEscape(base64.StdEncoding.EncodeToString(dig))
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Hmac() = %v, want %v", got, tt.want)
			}
		})
	}
}
