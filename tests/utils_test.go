package tests

import (
	"testing"
	"sgcodes7471/damsharaz.io-server/internal/pkg"
)

func TestParsePayload_Valid(t *testing.T) {
	payload := "Demo/r/nJOIN/r/nHello World/r/n"

	event, author, msg, err := pkg.Parse_Payload(payload)


	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if event != "JOIN" {
		t.Errorf("expected event JOIN, got %s", event)
	}

	if author != "Demo" {
		t.Errorf("expected author Demo, got %s", author)
	}

	if msg != "Hello World" {
		t.Errorf("expected msg Hello World, got %s", msg)
	}
}

func TestParsePayload_Invalid_NoDelimiters(t *testing.T) {
	payload := "INVALID_PAYLOAD"

	_, _, _, err := pkg.Parse_Payload(payload)

	if err == nil {
		t.Fatalf("expected error but got nil")
	}
}

func TestParsePayload_Invalid_OneDelimiter(t *testing.T) {
	payload := "DEMO/r/nJOIN"

	_, _, _, err := pkg.Parse_Payload(payload)

	if err == nil {
		t.Fatalf("expected error but got nil")
	}
}

func TestParsePayload_EmptyMessage(t *testing.T) {
	payload := "User/r/nMSG/r/n"

	event, author, msg, err := pkg.Parse_Payload(payload)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if event != "MSG" {
		t.Errorf("expected MSG, got %s", event)
	}

	if author != "User" {
		t.Errorf("expected User, got %s", author)
	}

	if msg != "" {
		t.Errorf("expected empty msg, got %s", msg)
	}
}
