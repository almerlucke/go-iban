package iban

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// Test fake ibans
func TestFake(t *testing.T) {
	fakeIBANs := []string{
		"VG96VPVG00000L2345678901",
		"1234567890",
		"12345678901234567890",
		"NL30ABNA0123456789",
		"NL30ABNA0517552264",
		"NL30ABNA05175522AB",
	}

	// Loop through fake ibans, they should all raise an error
	for _, fake := range fakeIBANs {
		iban, err := NewIBAN(fake)
		if err == nil {
			// Fake iban did not raise an error,
			t.Errorf("IBAN fake test error: %v", iban.Code)
		}
	}
}

// Test real ibans
func TestIBANS(t *testing.T) {
	err := filepath.Walk("example-ibans", func(path string, f os.FileInfo, err error) error {
		if f.IsDir() || strings.HasPrefix(f.Name(), ".") {
			return nil
		}

		file, err := os.Open(path)
		if err != nil {
			return err
		}

		defer file.Close()

		scanner := bufio.NewScanner(file)

		t.Logf("Start IBAN file test %v\n", path)

		for scanner.Scan() {
			line := scanner.Text()
			t.Logf("IBAN code input: %v\n", line)
			iban, ibanErr := NewIBAN(line)
			if ibanErr != nil {
				return ibanErr
			}
			t.Logf("IBAN code validated: %v\n", iban.PrintCode)
		}

		if err := scanner.Err(); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		t.Errorf("IBAN test error: %v", err)
	}
}
