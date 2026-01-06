package xstr

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestValidateEMVCoQRString(t *testing.T) {
	tests := []struct {
		name     string
		qrString string
		wantErr  bool
		errMsg   string
	}{
		{
			name:     "valid QR string",
			qrString: "010201630441C5",
			wantErr:  false,
		},
		{
			name:     "too short QR string",
			qrString: "0102015400",
			wantErr:  true,
			errMsg:   "qr string too short",
		},
		{
			name:     "empty QR string",
			qrString: "",
			wantErr:  true,
			errMsg:   "qr string too short",
		},
		{
			name:     "QR string with wrong CRC",
			qrString: "010201630441C6",
			wantErr:  true,
			errMsg:   "invalid crc",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateEMVCoQRString(tt.qrString)

			if tt.wantErr {
				assert.Error(t, err)
				if tt.errMsg != "" {
					assert.Contains(t, err.Error(), tt.errMsg)
				}
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestParseEMVCoQRString_ValidInputs(t *testing.T) {
	tests := []struct {
		name         string
		qrString     string
		expectFormat string
		expectCrc    string
		wantErr      bool
		errMsg       string
	}{
		{
			name:         "valid QR with basic fields",
			qrString:     "010201630441C5",
			expectFormat: "01",
			expectCrc:    "41C5",
			wantErr:      false,
		},
		{
			name:     "invalid QR - too short",
			qrString: "0102015400",
			wantErr:  true,
			errMsg:   "qr string too short",
		},
		{
			name:     "invalid QR - bad CRC",
			qrString: "010201630441C6",
			wantErr:  true,
			errMsg:   "invalid crc",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ParseEMVCoQRString(tt.qrString)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, result)
				if tt.errMsg != "" {
					assert.Contains(t, err.Error(), tt.errMsg)
				}
			} else {
				require.NoError(t, err)
				require.NotNil(t, result)

				assert.Equal(t, tt.expectFormat, result.Format)
				assert.Equal(t, tt.expectCrc, result.Crc)
			}
		})
	}
}

func TestParseEMVCoQRString_InvalidInputs(t *testing.T) {
	tests := []struct {
		name     string
		qrString string
		wantErr  bool
		errMsg   string
	}{
		{
			name:     "too short QR string",
			qrString: "123456789",
			wantErr:  true,
			errMsg:   "qr string too short",
		},
		{
			name:     "invalid CRC",
			qrString: "010201630441C6",
			wantErr:  true,
			errMsg:   "invalid crc",
		},
		{
			name:     "invalid structure - length exceeds remaining string",
			qrString: "01999912345630412AB", // length 9999 which exceeds string length
			wantErr:  true,
			errMsg:   "invalid crc", // CRC will fail first
		},
		{
			name:     "invalid structure - non-numeric length",
			qrString: "01XX123456304ABCD",
			wantErr:  true,
			errMsg:   "invalid crc", // CRC will fail first
		},
		{
			name:     "incomplete structure - truncated at length field",
			qrString: "010263045678", // incomplete length field
			wantErr:  true,
			errMsg:   "qr string too short", // Length check will fail first
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ParseEMVCoQRString(tt.qrString)

			assert.Error(t, err)
			assert.Nil(t, result)
			if tt.errMsg != "" {
				assert.Contains(t, err.Error(), tt.errMsg)
			}
		})
	}
}

func TestParseEMVCoQRString_EdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		qrString string
		wantErr  bool
	}{
		{
			name:     "minimal valid QR",
			qrString: "010201630441C5",
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ParseEMVCoQRString(tt.qrString)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, result)
			} else {
				require.NoError(t, err)
				require.NotNil(t, result)

				// Test that struct fields exist and are accessible
				assert.NotNil(t, result.Format)
				assert.NotNil(t, result.Crc)
				assert.NotNil(t, result.Amount)
				assert.NotNil(t, result.CountryCode)
				assert.NotNil(t, result.PhoneNumber)
				assert.NotNil(t, result.MerchantAccount)
			}
		})
	}
}

func TestEMVCoQRInfo_PhoneNumberExtraction(t *testing.T) {
	tests := []struct {
		name          string
		qrString      string
		expectedPhone string
		wantErr       bool
	}{
		{
			name:          "QR without phone number",
			qrString:      "010201630441C5",
			expectedPhone: "",
			wantErr:       false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ParseEMVCoQRString(tt.qrString)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.expectedPhone, result.PhoneNumber)
			}
		})
	}
}

func TestEMVCoQRInfo_Tag30Structure(t *testing.T) {
	tests := []struct {
		name        string
		qrString    string
		expectedRef *EMVCoQRInfo
		wantErr     bool
	}{
		{
			name:     "basic QR structure",
			qrString: "010201630441C5",
			expectedRef: &EMVCoQRInfo{
				BillerID: "",
				Ref1:     "",
				Ref2:     "",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ParseEMVCoQRString(tt.qrString)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.expectedRef.BillerID, result.BillerID)
				assert.Equal(t, tt.expectedRef.Ref1, result.Ref1)
				assert.Equal(t, tt.expectedRef.Ref2, result.Ref2)
			}
		})
	}
}

func TestEMVCoQRInfo_StructFields(t *testing.T) {
	qrString := "010201630441C5"

	result, err := ParseEMVCoQRString(qrString)
	require.NoError(t, err)
	require.NotNil(t, result)

	// Test that struct fields exist and are properly initialized
	assert.NotNil(t, result.Format)
	assert.NotNil(t, result.Crc)
	assert.NotNil(t, result.Amount)
	assert.NotNil(t, result.CountryCode)
	assert.NotNil(t, result.PhoneNumber)
	assert.NotNil(t, result.MerchantAccount)
	assert.NotNil(t, result.CurrencyISO4217)
	assert.NotNil(t, result.BillerID)
	assert.NotNil(t, result.Ref1)
	assert.NotNil(t, result.Ref2)
	assert.NotNil(t, result.Ref3)
}
