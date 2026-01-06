// Package xstr provides EMV QR Code decoding functionality.
// This implementation supports EMV QR Code specification with CRC-16 validation,
// TLV parsing, and multi-scheme support (PromptPay, QRIS, DuitNow, etc.).
package xstr

import (
	"fmt"
	"strconv"
)

// QRPaymentType represents the type of QR payment based on AID.
type QRPaymentType string

// POIMethodType represents the Point of Initiation Method type.
type POIMethodType string

// QR Payment Type constants
const (
	QRTypeC2C         QRPaymentType = "C2C"         // Consumer-to-Consumer Transfer
	QRTypeC2B         QRPaymentType = "C2B"         // Consumer-to-Business (Merchant Presented)
	QRTypeBillPayment QRPaymentType = "BillPayment" // Bill Payment
	QRTypeCrossBorder QRPaymentType = "CrossBorder" // Cross-Border Payment
	QRTypeUnknown     QRPaymentType = "Unknown"
)

// POI Method Type constants
const (
	POITypeStatic  POIMethodType = "static"  // Static QR with fixed amount
	POITypeDynamic POIMethodType = "dynamic" // Dynamic QR, amount entered later
	POITypeUnknown POIMethodType = "unknown"
)

// QRPaymentScheme represents the payment scheme/network type.
type QRPaymentScheme string

// QR Payment Scheme constants define supported payment networks globally
const (
	QRSchemePromptPay QRPaymentScheme = "PromptPay" // Thailand national payment system
	QRSchemeQRIS      QRPaymentScheme = "QRIS"      // Indonesia national QR standard
	QRSchemeDuitNow   QRPaymentScheme = "DuitNow"   // Malaysia real-time payment
	QRSchemeUPI       QRPaymentScheme = "UPI"       // India Unified Payments Interface
	QRSchemeNETS      QRPaymentScheme = "NETS"      // Singapore electronic payment
	QRSchemeAlipay    QRPaymentScheme = "Alipay"    // Alipay global payment
	QRSchemeWeChatPay QRPaymentScheme = "WeChatPay" // WeChat Pay global payment
	QRSchemeUnknown   QRPaymentScheme = "Unknown"
)

// MerchantAccount represents detailed merchant account information with parsed sub-fields.
type MerchantAccount struct {
	AID            string            `json:"aid"`             // Tag 00: Application/GUI Identifier
	AIDType        QRPaymentType     `json:"aid_type"`        // Mapped AID type (C2C, C2B, BillPayment, etc.)
	PaymentScheme  QRPaymentScheme   `json:"payment_scheme"`  // Mapped payment scheme (PromptPay, QRIS, etc.)
	MerchantID     string            `json:"merchant_id"`     // Tag 01: Merchant/Biller ID
	Reference1     string            `json:"reference_1"`     // Tag 02: Reference 1
	Reference2     string            `json:"reference_2"`     // Tag 03: Reference 2
	Reference3     string            `json:"reference_3"`     // Tag 04: Reference 3 (if exists)
	RawValue       string            `json:"raw_value"`       // Original raw value
	UnresolvedData map[string]string `json:"unresolved_data"` // Other unresolved sub-fields
}

// EMVData represents decoded EMV QR code data structure.
type EMVData struct {
	PayloadFormatIndicator    string                      `json:"payload_format_indicator"`
	PointOfInitiationMethod   string                      `json:"point_of_initiation_method"`
	POIMethodType             POIMethodType               `json:"poi_method_type"` // Mapped POI method type (static, dynamic)
	MerchantAccountInfo       map[string]*MerchantAccount `json:"merchant_account_info"`
	MerchantCategoryCode      string                      `json:"merchant_category_code"`
	TransactionCurrency       string                      `json:"transaction_currency"`
	TransactionAmount         string                      `json:"transaction_amount"`
	TipOrConvenienceIndicator string                      `json:"tip_or_convenience_indicator"`
	ValueOfConvenienceFee     string                      `json:"value_of_convenience_fee"`
	CountryCode               string                      `json:"country_code"`
	MerchantName              string                      `json:"merchant_name"`
	MerchantCity              string                      `json:"merchant_city"`
	PostalCode                string                      `json:"postal_code"`
	AdditionalData            map[string]string           `json:"additional_data"`
	MerchantInformation       map[string]string           `json:"merchant_information"`
	CRC                       string                      `json:"crc"`
	UnresolvedData            map[string]string           `json:"unresolved_data"`
}

// EMVDataValue represents a single EMV data field with tag, length, and value.
type EMVDataValue struct {
	Tag    string `json:"tag"`
	Length int    `json:"length"`
	Value  string `json:"value"`
}

// DecodeEMVQR decodes EMV QR code string and returns structured data.
// It parses the TLV (Tag-Length-Value) format according to EMV QR Code specification.
func DecodeEMVQR(qrString string) (*EMVData, error) {
	if len(qrString) < 4 {
		return nil, fmt.Errorf("invalid EMV QR code: too short")
	}

	emvData := &EMVData{
		MerchantAccountInfo: make(map[string]*MerchantAccount),
		AdditionalData:      make(map[string]string),
		MerchantInformation: make(map[string]string),
		UnresolvedData:      make(map[string]string),
	}

	// Parse TLV data sequentially from QR string
	position := 0
	for position < len(qrString) {
		if position+4 > len(qrString) {
			break
		}

		// Parse tag (2 digits)
		tag := qrString[position : position+2]
		position += 2

		// Parse length (2 digits)
		lengthStr := qrString[position : position+2]
		position += 2

		length, err := strconv.Atoi(lengthStr)
		if err != nil {
			return nil, fmt.Errorf("invalid length at position %d: %s", position-2, lengthStr)
		}

		if position+length > len(qrString) {
			return nil, fmt.Errorf("invalid data length at tag %s", tag)
		}

		// Parse value
		value := qrString[position : position+length]
		position += length

		// Map to appropriate field
		if err := mapEMVField(emvData, tag, value); err != nil {
			return nil, fmt.Errorf("error mapping field %s: %v", tag, err)
		}
	}

	// Validate CRC checksum to ensure data integrity
	// EMV standard requires CRC-16 validation on complete QR data
	if emvData.CRC != "" {
		// Reconstruct original data for CRC calculation by removing actual CRC value
		// and appending dummy CRC tag "6304" as per EMV specification
		crcTagPosition := len(qrString) - 8 // CRC is always 8 chars: "6304" + 4-digit CRC
		if crcTagPosition >= 0 && qrString[crcTagPosition:crcTagPosition+4] == "6304" {
			dataForCRC := qrString[:crcTagPosition] + "6304"
			calculatedCRC := calculateCRC16(dataForCRC)

			if emvData.CRC != calculatedCRC {
				return nil, fmt.Errorf("invalid CRC: expected %s, got %s", calculatedCRC, emvData.CRC)
			}
		} else {
			return nil, fmt.Errorf("invalid EMV QR format: CRC tag not found at expected position")
		}
	}

	return emvData, nil
}

// ParseEMVTLV parses EMV QR code string into individual TLV structures.
// Returns a slice of EMVDataValue representing each tag-length-value triplet.
func ParseEMVTLV(qrString string) ([]EMVDataValue, error) {
	if len(qrString) < 4 {
		return nil, fmt.Errorf("invalid EMV QR code: too short")
	}

	var tlvData []EMVDataValue
	position := 0

	for position < len(qrString) {
		if position+4 > len(qrString) {
			break
		}

		// Parse tag (2 digits)
		tag := qrString[position : position+2]
		position += 2

		// Parse length (2 digits)
		lengthStr := qrString[position : position+2]
		position += 2

		length, err := strconv.Atoi(lengthStr)
		if err != nil {
			return nil, fmt.Errorf("invalid length at position %d: %s", position-2, lengthStr)
		}

		if position+length > len(qrString) {
			return nil, fmt.Errorf("invalid data length at tag %s", tag)
		}

		// Parse value
		value := qrString[position : position+length]
		position += length

		tlvData = append(tlvData, EMVDataValue{
			Tag:    tag,
			Length: length,
			Value:  value,
		})
	}

	return tlvData, nil
}

// mapEMVField maps EMV tag to appropriate struct field.
func mapEMVField(emvData *EMVData, tag, value string) error {
	switch tag {
	case "00":
		emvData.PayloadFormatIndicator = value
	case "01":
		emvData.PointOfInitiationMethod = value
		emvData.POIMethodType = mapPOIMethodType(value)
	case "52":
		emvData.MerchantCategoryCode = value
	case "53":
		emvData.TransactionCurrency = value
	case "54":
		emvData.TransactionAmount = value
	case "55":
		emvData.TipOrConvenienceIndicator = value
	case "56":
		emvData.ValueOfConvenienceFee = value
	case "58":
		emvData.CountryCode = value
	case "59":
		emvData.MerchantName = value
	case "60":
		emvData.MerchantCity = value
	case "61":
		emvData.PostalCode = value
	case "62":
		// Additional Data Field Template
		subFields, err := parseSubFields(value)
		if err != nil {
			return fmt.Errorf("error parsing additional data: %v", err)
		}
		emvData.AdditionalData = subFields
	case "63":
		emvData.CRC = value
	default:
		// Handle merchant account information (tags 02-51)
		// These tags contain payment provider specific data
		if tag >= "02" && tag <= "51" {
			// Parse merchant account sub-fields
			merchantAccount, err := parseMerchantAccountInfo(value)
			if err != nil {
				return fmt.Errorf("error parsing merchant account info: %v", err)
			}
			merchantAccount.RawValue = value
			emvData.MerchantAccountInfo[tag] = merchantAccount
		} else if tag >= "64" && tag <= "99" {
			// Merchant Information Language Template or RFU
			if tag == "99" {
				// This is actually unresolved/future use
				emvData.UnresolvedData[tag] = value
			} else {
				emvData.MerchantInformation[tag] = value
			}
		} else {
			// Store unresolved data
			emvData.UnresolvedData[tag] = value
		}
	}

	return nil
}

// parseSubFields parses sub-fields within a TLV structure.
func parseSubFields(data string) (map[string]string, error) {
	subFields := make(map[string]string)
	position := 0

	for position < len(data) {
		if position+4 > len(data) {
			break
		}

		// Parse tag (2 digits)
		tag := data[position : position+2]
		position += 2

		// Parse length (2 digits)
		lengthStr := data[position : position+2]
		position += 2

		length, err := strconv.Atoi(lengthStr)
		if err != nil {
			return nil, fmt.Errorf("invalid sub-field length: %s", lengthStr)
		}

		if position+length > len(data) {
			return nil, fmt.Errorf("invalid sub-field data length at tag %s", tag)
		}

		// Parse value
		value := data[position : position+length]
		position += length

		subFields[tag] = value
	}

	return subFields, nil
}

// parseMerchantAccountInfo parses merchant account information sub-fields.
// Returns a MerchantAccount struct with parsed sub-fields according to EMV specification.
func parseMerchantAccountInfo(data string) (*MerchantAccount, error) {
	account := &MerchantAccount{
		UnresolvedData: make(map[string]string),
	}

	subFields, err := parseSubFields(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse merchant account sub-fields: %v", err)
	}

	// Map known sub-fields to struct properties
	// Each payment scheme may use different sub-field combinations
	for tag, value := range subFields {
		switch tag {
		case "00":
			// AID (Application Identifier) determines payment scheme and type
			account.AID = value
			account.AIDType = mapAIDType(value)
			account.PaymentScheme = mapScheme(value)
		case "01":
			account.MerchantID = value
		case "02":
			account.Reference1 = value
		case "03":
			account.Reference2 = value
		case "04":
			account.Reference3 = value
		default:
			// Store unknown sub-fields
			account.UnresolvedData[tag] = value
		}
	}

	return account, nil
}

// calculateCRC16 calculates CRC-16 checksum for EMV QR code validation.
// Uses CRC-16-CCITT polynomial: 0x1021 with EMV QR Code specific parameters:
// Initial value: 0xFFFF, Final XOR: 0x0000, Reflect input/output: false
func calculateCRC16(data string) string {
	crc := uint16(0xFFFF)

	// Process each byte using CRC-16-CCITT polynomial
	for i := 0; i < len(data); i++ {
		crc ^= uint16(data[i]) << 8

		// Process each bit using polynomial 0x1021
		for j := 0; j < 8; j++ {
			if crc&0x8000 != 0 {
				crc = (crc << 1) ^ 0x1021
			} else {
				crc = crc << 1
			}
		}
	}

	// Final XOR with 0x0000 (no change) as per EMV specification
	return fmt.Sprintf("%04X", crc&0xFFFF)
}

// mapScheme determines payment network from AID/GUI identifier.
// This enables proper routing and business logic for different payment systems.
func mapScheme(gui string) QRPaymentScheme {
	switch gui {
	case "A000000677010111",
		"A000000677010112",
		"A000000677010113",
		"A000000677010114":
		return QRSchemePromptPay
	case "ID.CO.QRIS.WWW", "COM.INACASH.WWW":
		return QRSchemeQRIS
	case "COM.MY.DUITNOW":
		return QRSchemeDuitNow
	case "COM.UPI.PAY":
		return QRSchemeUPI
	case "COM.SG.NETS":
		return QRSchemeNETS
	case "COM.ALIPAY.WWW":
		return QRSchemeAlipay
	case "COM.WECHAT.WWW":
		return QRSchemeWeChatPay
	default:
		return QRSchemeUnknown
	}
}

// mapAIDType determines payment type for PromptPay AIDs only.
// Other payment schemes use different classification systems.
func mapAIDType(aid string) QRPaymentType {
	switch aid {
	case "A000000677010111":
		return QRTypeC2C
	case "A000000677010112":
		return QRTypeC2B
	case "A000000677010113":
		return QRTypeBillPayment
	case "A000000677010114":
		return QRTypeCrossBorder
	default:
		return QRTypeUnknown
	}
}

// mapPOIMethodType converts EMV POI method codes to readable types.
// This affects how payment amount is handled (fixed vs dynamic).
func mapPOIMethodType(poiMethod string) POIMethodType {
	switch poiMethod {
	case "11":
		return POITypeStatic
	case "12":
		return POITypeDynamic
	default:
		return POITypeUnknown
	}
}

// QRInfo represents consolidated QR code information from primary merchant account.
// This provides a simplified view of the most important QR data for business logic,
// extracting key information from the primary payment account within the EMV data.
type QRInfo struct {
	AID               string          `json:"aid"`
	AIDType           QRPaymentType   `json:"aid_type"`
	POIMethodType     POIMethodType   `json:"poi_method_type"`
	PaymentScheme     QRPaymentScheme `json:"payment_scheme"`
	TransactionAmount string          `json:"transaction_amount"`
	CountryCode       string          `json:"country_code"`
	MerchantID        string          `json:"merchant_id"`
	Reference1        string          `json:"reference_1"`
	Reference2        string          `json:"reference_2"`
	Reference3        string          `json:"reference_3"`
}

// QRInfo extracts consolidated information from the primary merchant account.
// This method prioritizes merchant accounts and provides unified access to key QR data
// for business logic and payment processing.
func (e *EMVData) QRInfo() QRInfo {
	info := QRInfo{
		POIMethodType:     e.POIMethodType,
		TransactionAmount: e.TransactionAmount,
		CountryCode:       e.CountryCode,
	}

	// Find primary merchant account (prefer lower tag numbers as they're typically primary)
	var primaryAccount *MerchantAccount

	// Look for merchant accounts in order of preference
	preferredTags := []string{"26", "27", "28", "29", "30", "31", "32", "33", "34", "35"}

	// First pass: check preferred tags
	for _, tag := range preferredTags {
		if account, exists := e.MerchantAccountInfo[tag]; exists {
			primaryAccount = account
			break
		}
	}

	// Second pass: if no preferred tag found, use first available
	if primaryAccount == nil {
		for tag, account := range e.MerchantAccountInfo {
			if tag >= "02" && tag <= "51" {
				primaryAccount = account
				break
			}
		}
	}

	// Extract information from primary account
	if primaryAccount != nil {
		info.AID = primaryAccount.AID
		info.AIDType = primaryAccount.AIDType
		info.PaymentScheme = primaryAccount.PaymentScheme
		info.MerchantID = primaryAccount.MerchantID
		info.Reference1 = primaryAccount.Reference1
		info.Reference2 = primaryAccount.Reference2
		info.Reference3 = primaryAccount.Reference3

		// Fill missing references with additional data if available
		fillMissingReferences(&info, e.AdditionalData)
	}

	return info
}

// fillMissingReferences fills empty reference fields with additional data
func fillMissingReferences(info *QRInfo, additionalData map[string]string) {
	references := []*string{&info.Reference1, &info.Reference2, &info.Reference3}

	// Find empty reference slots
	var emptySlots []*string
	for _, ref := range references {
		if *ref == "" {
			emptySlots = append(emptySlots, ref)
		}
	}

	// Fill empty slots with additional data values directly
	slotIndex := 0
	for _, value := range additionalData {
		if slotIndex >= len(emptySlots) {
			break
		}
		// Use additional data value directly without tag prefix
		*emptySlots[slotIndex] = value
		slotIndex++
	}
}
