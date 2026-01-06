package xstr

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/sigurn/crc16"
)

// EMVCoQRInfo holds parsed EMVCo QR code information
type EMVCoQRInfo struct {
	Format          string
	MerchantAccount string
	Amount          string
	PhoneNumber     string
	CountryCode     string
	Crc             string
	CurrencyISO4217 string
	BillerID        string
	Ref1            string
	Ref2            string
	Ref3            string
}

var crc16Table *crc16.Table

func init() {
	crc16Table = crc16.MakeTable(crc16.CRC16_CCITT_FALSE)
}

func validateEMVCoQRString(qrString string) error {
	if len(qrString) < 14 {
		return fmt.Errorf("qr string too short")
	}
	data := []byte(qrString[:len(qrString)-4])
	crc := crc16.Checksum(data, crc16Table)
	calculatedCRC := fmt.Sprintf("%04X", int(crc))
	expectedCRC := qrString[len(qrString)-4:]
	if calculatedCRC != expectedCRC {
		return fmt.Errorf("invalid crc: expected %s, got %s", expectedCRC, calculatedCRC)
	}
	return nil
}

func ParseEMVCoQRString(qrString string) (*EMVCoQRInfo, error) {
	if err := validateEMVCoQRString(qrString); err != nil {
		return nil, err
	}
	result := &EMVCoQRInfo{}
	index := 0
	for index < len(qrString) {
		if index+4 > len(qrString) {
			return nil, fmt.Errorf("invalid qr structure")
		}
		id := qrString[index : index+2]
		length, err := strconv.Atoi(qrString[index+2 : index+4])
		if err != nil {
			return nil, fmt.Errorf("invalid qr structure")
		}
		if index+4+length > len(qrString) {
			return nil, fmt.Errorf("invalid specified qr string length")
		}
		value := qrString[index+4 : index+4+length]
		switch id {
		case "01":
			result.Format = value
		case "29":
			prefixPhoneIndex := strings.Index(value, "011300")
			result.MerchantAccount = value
			if prefixPhoneIndex != -1 {
				result.PhoneNumber = value[prefixPhoneIndex+6:]
			} else {
				prefixIDIndex := strings.Index(value, "110213")
				if prefixIDIndex != -1 {
					result.PhoneNumber = value[prefixIDIndex+6:]
				}
			}
		case "30":
			result.MerchantAccount = value
			index2 := 0
			for index2 < len(value) {
				if index2+4 > len(value) {
					return nil, fmt.Errorf("invalid qr structure")
				}
				id2 := value[index2 : index2+2]
				length2, err := strconv.Atoi(value[index2+2 : index2+4])
				if err != nil {
					return nil, fmt.Errorf("invalid qr structure")
				}
				if index2+4+length2 > len(value) {
					return nil, fmt.Errorf("invalid specified qr string length")
				}
				value2 := value[index2+4 : index2+4+length2]
				switch id2 {
				case "01":
					result.BillerID = value2
				case "02":
					result.Ref1 = value2
				case "03":
					result.Ref2 = value2
				}
				index2 += 4 + length2
			}
		case "54":
			result.Amount = value
		case "58":
			result.CountryCode = value
		case "62":
			if len(value) > 4 {
				result.Ref3 = value[4:]
			}
		case "63":
			result.Crc = value
		case "53":
			result.CurrencyISO4217 = value
		}
		index += 4 + length
	}
	return result, nil
}
