package main

import (
	"bytes"
	"path/filepath"
	"testing"
)

func TestRunValidFile(t *testing.T) {
	filePath := filepath.Join("tests", "testdata", "valid_single_o.txt")
	var output bytes.Buffer

	err := run([]string{filePath}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	if output.String() != "AA\nAA\n" {
		t.Fatalf("output = %q, want %q", output.String(), "AA\nAA\n")
	}
}

func TestRunRejectsInvalidArgCount(t *testing.T) {
	var output bytes.Buffer

	err := run([]string{}, &output)
	if err == nil {
		t.Fatal("run returned nil error for invalid argument count")
	}
}

func TestRunRejectsInvalidFile(t *testing.T) {
	filePath := filepath.Join("tests", "testdata", "invalid_disconnected.txt")
	var output bytes.Buffer

	err := run([]string{filePath}, &output)
	if err == nil {
		t.Fatal("run returned nil error for invalid tetromino file")
	}
}

func TestRunSubjectSample(t *testing.T) {
	filePath := filepath.Join("tests", "testdata", "sample_subject.txt")
	var output bytes.Buffer

	err := run([]string{filePath}, &output)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	want := "ABBBB.\nACCCEE\nAFFCEE\nA.FFGG\nHHHDDG\n.HDD.G\n"
	if output.String() != want {
		t.Fatalf("output = %q, want %q", output.String(), want)
	}
}
