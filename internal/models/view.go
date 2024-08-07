package models

import (
	"strings"
	"time"

	"github.com/google/uuid"
)

type Noteable interface {
	GetName() string
	GetType() TypeNote
	GetId() uuid.UUID
	Print() string
}

type TypeNote string

func (t TypeNote) String() string {
	return string(t)
}

const (
	CARD       TypeNote = "bank card"
	CREDENTIAL TypeNote = "credential"
	TEXT       TypeNote = "text"
	BINARY     TypeNote = "binary"
)

type BaseNote struct {
	Id         uuid.UUID `json:"id"`
	NameRecord string    `json:"name_record"`
	Created    int64     `json:"created"`
	Type       TypeNote  `json:"type"`
	MetaInfo   []string  `json:"meta_info,omitempty"`
}

type CredentialNote struct {
	Username string `json:"username"`
	Password string `json:"password"`
	BaseNote `json:"data"`
}

func (cn CredentialNote) Print() string {
	var str string
	str += "Note: " + cn.NameRecord + "\n"
	str += "Username: " + cn.Username + "\n"
	str += "Password: " + cn.Password + "\n"
	str += "Additional information: " + strings.Join(cn.MetaInfo, "; ") + "\n"
	str += "Created: " + time.Unix(cn.Created, 0).Format(time.RFC822) + "\n"
	return str
}

func (cn CredentialNote) GetName() string {
	return cn.NameRecord
}

func (cn CredentialNote) GetType() TypeNote {
	return CREDENTIAL
}

func (cn CredentialNote) GetId() uuid.UUID {
	return cn.Id
}

type TextNote struct {
	Text     string `json:"text"`
	BaseNote `json:"data"`
}

func (tn TextNote) Print() string {
	var str string
	str += "Note: " + tn.NameRecord + "\n"
	str += "Text: " + tn.Text + "\n"
	str += "Additional information: " + strings.Join(tn.MetaInfo, "; ") + "\n"
	str += "Created: " + time.Unix(tn.Created, 0).Format(time.RFC822) + "\n"
	return str
}

func (tn TextNote) GetName() string {
	return tn.NameRecord
}

func (tn TextNote) GetType() TypeNote {
	return TEXT
}

func (tn TextNote) GetId() uuid.UUID {
	return tn.Id
}

type BinaryNote struct {
	Binary   []byte `json:"binary"`
	BaseNote `json:"data"`
}

func (bn BinaryNote) Print() string {
	var str string
	str += "Note: " + bn.NameRecord + "\n"
	str += "Binary: " + string(bn.Binary) + "\n"
	str += "Additional information: \n" + strings.Join(bn.MetaInfo, "; ") + "\n"
	str += "Created: " + time.Unix(bn.Created, 0).Format(time.RFC822) + "\n"
	return str
}

func (bn BinaryNote) GetName() string {
	return bn.NameRecord
}

func (bn BinaryNote) GetType() TypeNote {
	return BINARY
}

func (bn BinaryNote) GetId() uuid.UUID {
	return bn.Id
}

type BankCardNote struct {
	Bank         string `json:"bank"`
	Number       string `json:"number"`
	Expiration   string `json:"expiration"`
	Cardholder   string `json:"cardholder"`
	SecurityCode string `json:"security_code"`
	BaseNote     `json:"data"`
}

func (bnc BankCardNote) Print() string {
	var str string
	str += "Note: " + bnc.NameRecord + "\n"
	str += "Bank name: " + bnc.Bank + "\n"
	str += "Card number: " + bnc.Number + "\n"
	str += "Card expiration: " + bnc.Expiration + "\n"
	str += "Cardholder name: " + bnc.Cardholder + "\n"
	str += "Security code: " + bnc.SecurityCode + "\n"
	str += "Additional information: " + strings.Join(bnc.MetaInfo, "; ") + "\n"
	str += "Created: " + time.Unix(bnc.Created, 0).Format(time.RFC822) + "\n"
	return str
}

func (bnc BankCardNote) GetName() string {
	return bnc.NameRecord
}

func (bnc BankCardNote) GetType() TypeNote {
	return CARD
}

func (bnc BankCardNote) GetId() uuid.UUID {
	return bnc.Id
}
