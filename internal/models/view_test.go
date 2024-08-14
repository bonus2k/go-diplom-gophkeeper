package models

import (
	"reflect"
	"testing"

	"github.com/google/uuid"
)

var (
	baseNote = BaseNote{Id: uuid.Nil, NameRecord: "Test Note", Created: 1723652739, Type: CARD, MetaInfo: []string{"test", "test"}}
	bank     = BankCardNote{Bank: "TEST BANK", Number: "1234-1234-1234-1234", Expiration: "12/24", Cardholder: "Test User", SecurityCode: "456", BaseNote: baseNote}
)

func TestBankCardNote_GetID(t *testing.T) {
	type fields struct {
		Bank         string
		Number       string
		Expiration   string
		Cardholder   string
		SecurityCode string
		BaseNote     BaseNote
	}
	tests := []struct {
		name   string
		fields fields
		want   uuid.UUID
	}{
		{
			name: "bank card",
			fields: fields{
				Bank:         "TEST BANK",
				Number:       "1234-1234-1234-1234",
				Expiration:   "12/24",
				Cardholder:   "Test User",
				SecurityCode: "456",
				BaseNote:     baseNote,
			},
			want: uuid.Nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bnc := BankCardNote{
				Bank:         tt.fields.Bank,
				Number:       tt.fields.Number,
				Expiration:   tt.fields.Expiration,
				Cardholder:   tt.fields.Cardholder,
				SecurityCode: tt.fields.SecurityCode,
				BaseNote:     tt.fields.BaseNote,
			}
			if got := bnc.GetID(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBankCardNote_GetName(t *testing.T) {
	type fields struct {
		note BankCardNote
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "bank card",
			fields: fields{
				note: bank,
			},
			want: "Test Note",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fields.note.GetName(); got != tt.want {
				t.Errorf("GetName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBankCardNote_GetType(t *testing.T) {
	type fields struct {
		note BankCardNote
	}
	tests := []struct {
		name   string
		fields fields
		want   TypeNote
	}{
		{
			name: "bank card",
			fields: fields{
				note: bank,
			},
			want: CARD,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if got := tt.fields.note.GetType(); got != tt.want {
				t.Errorf("GetType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBankCardNote_Print(t *testing.T) {
	type fields struct {
		note BankCardNote
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "bank card",
			fields: fields{
				note: bank,
			},
			want: "Note: Test Note\n" +
				"Bank name: TEST BANK\n" +
				"Card number: 1234-1234-1234-1234\n" +
				"Card expiration: 12/24\n" +
				"Cardholder name: Test User\n" +
				"Security code: 456\n" +
				"Additional information: test; test\n" +
				"Created: 14 Aug 24 19:25 MSK\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fields.note.Print(); got != tt.want {
				t.Errorf("Print() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinaryNote_GetID(t *testing.T) {
	type fields struct {
		Binary   []byte
		BaseNote BaseNote
	}
	tests := []struct {
		name   string
		fields fields
		want   uuid.UUID
	}{
		{
			name: "binary",
			fields: fields{
				Binary:   []byte("TEST BINARY"),
				BaseNote: baseNote,
			},
			want: uuid.Nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bn := BinaryNote{
				Binary:   tt.fields.Binary,
				BaseNote: tt.fields.BaseNote,
			}
			if got := bn.GetID(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinaryNote_GetName(t *testing.T) {
	type fields struct {
		Binary   []byte
		BaseNote BaseNote
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "binary",
			fields: fields{
				Binary:   []byte("TEST BINARY"),
				BaseNote: baseNote,
			},
			want: "Test Note",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bn := BinaryNote{
				Binary:   tt.fields.Binary,
				BaseNote: tt.fields.BaseNote,
			}
			if got := bn.GetName(); got != tt.want {
				t.Errorf("GetName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinaryNote_GetType(t *testing.T) {
	type fields struct {
		Binary   []byte
		BaseNote BaseNote
	}
	tests := []struct {
		name   string
		fields fields
		want   TypeNote
	}{
		{
			name: "binary",
			fields: fields{
				Binary:   []byte("TEST BINARY"),
				BaseNote: baseNote,
			},
			want: BINARY,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bn := BinaryNote{
				Binary:   tt.fields.Binary,
				BaseNote: tt.fields.BaseNote,
			}
			if got := bn.GetType(); got != tt.want {
				t.Errorf("GetType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinaryNote_Print(t *testing.T) {
	type fields struct {
		Binary   []byte
		BaseNote BaseNote
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "binary",
			fields: fields{
				Binary:   []byte("TEST BINARY"),
				BaseNote: baseNote,
			},
			want: "Note: Test Note\n" +
				"Binary: TEST BINARY\n" +
				"Additional information: test; test\n" +
				"Created: 14 Aug 24 19:25 MSK\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bn := BinaryNote{
				Binary:   tt.fields.Binary,
				BaseNote: tt.fields.BaseNote,
			}
			if got := bn.Print(); got != tt.want {
				t.Errorf("Print() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCredentialNote_GetID(t *testing.T) {
	type fields struct {
		Username string
		Password string
		BaseNote BaseNote
	}
	tests := []struct {
		name   string
		fields fields
		want   uuid.UUID
	}{
		{
			name: "username",
			fields: fields{
				Username: "username",
				Password: "password",
				BaseNote: baseNote,
			},
			want: uuid.Nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cn := CredentialNote{
				Username: tt.fields.Username,
				Password: tt.fields.Password,
				BaseNote: tt.fields.BaseNote,
			}
			if got := cn.GetID(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCredentialNote_GetName(t *testing.T) {
	type fields struct {
		Username string
		Password string
		BaseNote BaseNote
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "username",
			fields: fields{
				Username: "username",
				Password: "password",
				BaseNote: baseNote,
			},
			want: "Test Note",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cn := CredentialNote{
				Username: tt.fields.Username,
				Password: tt.fields.Password,
				BaseNote: tt.fields.BaseNote,
			}
			if got := cn.GetName(); got != tt.want {
				t.Errorf("GetName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCredentialNote_GetType(t *testing.T) {
	type fields struct {
		Username string
		Password string
		BaseNote BaseNote
	}
	tests := []struct {
		name   string
		fields fields
		want   TypeNote
	}{
		{
			name: "username",
			fields: fields{
				Username: "username",
				Password: "password",
				BaseNote: baseNote,
			},
			want: CREDENTIAL,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cn := CredentialNote{
				Username: tt.fields.Username,
				Password: tt.fields.Password,
				BaseNote: tt.fields.BaseNote,
			}
			if got := cn.GetType(); got != tt.want {
				t.Errorf("GetType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCredentialNote_Print(t *testing.T) {
	type fields struct {
		Username string
		Password string
		BaseNote BaseNote
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "username",
			fields: fields{
				Username: "username",
				Password: "password",
				BaseNote: baseNote,
			},
			want: "Note: Test Note\n" +
				"Username: username\n" +
				"Password: password\n" +
				"Additional information: test; test\n" +
				"Created: 14 Aug 24 19:25 MSK\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cn := CredentialNote{
				Username: tt.fields.Username,
				Password: tt.fields.Password,
				BaseNote: tt.fields.BaseNote,
			}
			if got := cn.Print(); got != tt.want {
				t.Errorf("Print() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTextNote_GetID(t *testing.T) {
	type fields struct {
		Text     string
		BaseNote BaseNote
	}
	tests := []struct {
		name   string
		fields fields
		want   uuid.UUID
	}{
		{
			name: "text",
			fields: fields{
				Text:     "Test Note",
				BaseNote: baseNote,
			},
			want: uuid.Nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tn := TextNote{
				Text:     tt.fields.Text,
				BaseNote: tt.fields.BaseNote,
			}
			if got := tn.GetID(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTextNote_GetName(t *testing.T) {
	type fields struct {
		Text     string
		BaseNote BaseNote
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "text",
			fields: fields{
				Text:     "Test Note",
				BaseNote: baseNote,
			},
			want: "Test Note",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tn := TextNote{
				Text:     tt.fields.Text,
				BaseNote: tt.fields.BaseNote,
			}
			if got := tn.GetName(); got != tt.want {
				t.Errorf("GetName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTextNote_GetType(t *testing.T) {
	type fields struct {
		Text     string
		BaseNote BaseNote
	}
	tests := []struct {
		name   string
		fields fields
		want   TypeNote
	}{
		{
			name: "text",
			fields: fields{
				Text:     "Test Note",
				BaseNote: baseNote,
			},
			want: TEXT,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tn := TextNote{
				Text:     tt.fields.Text,
				BaseNote: tt.fields.BaseNote,
			}
			if got := tn.GetType(); got != tt.want {
				t.Errorf("GetType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTextNote_Print(t *testing.T) {
	type fields struct {
		Text     string
		BaseNote BaseNote
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "text",
			fields: fields{
				Text:     "Test Note",
				BaseNote: baseNote,
			},
			want: "Note: Test Note\n" +
				"Text: Test Note\n" +
				"Additional information: test; test\n" +
				"Created: 14 Aug 24 19:25 MSK\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tn := TextNote{
				Text:     tt.fields.Text,
				BaseNote: tt.fields.BaseNote,
			}
			if got := tn.Print(); got != tt.want {
				t.Errorf("Print() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTypeNote_String(t *testing.T) {
	tests := []struct {
		name string
		t    TypeNote
		want string
	}{
		{
			name: "text",
			t:    TEXT,
			want: "text",
		},
		{
			name: "binary",
			t:    BINARY,
			want: "binary",
		},
		{
			name: "credential",
			t:    CREDENTIAL,
			want: "credential",
		},
		{
			name: "bank card",
			t:    CARD,
			want: "bank card",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.t.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
