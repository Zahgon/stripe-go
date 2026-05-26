//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"github.com/stripe/stripe-go/v85/form"
)

type PaymentSourceType string

// List of values that PaymentSourceType can take
const (
	PaymentSourceTypeAccount     PaymentSourceType = "account"
	PaymentSourceTypeBankAccount PaymentSourceType = "bank_account"
	PaymentSourceTypeCard        PaymentSourceType = "card"
	PaymentSourceTypeSource      PaymentSourceType = "source"
)

// List sources for a specified customer.
type PaymentSourceListParams struct {
	ListParams `form:"*"`
	Customer   *string `form:"-"` // Included in URL
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
	// Filter sources according to a particular object type.
	Object *string `form:"object" json:"object,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *PaymentSourceListParams) AddExpand(f string) { _ = "STUB: not implemented"; return }

// PaymentSourceSourceParams is a union struct used to describe an
// arbitrary payment source.
type PaymentSourceSourceParams struct {
	Card  *CardParams `form:"-"`
	Token *string     `form:"source"`
}

// AppendTo implements custom encoding logic for PaymentSourceSourceParams.
func (p *PaymentSourceSourceParams) AppendTo(body *form.Values, keyParts []string) {
	_ = "STUB: not implemented"
	return
}

// SourceParamsFor creates PaymentSourceSourceParams objects around supported
// payment sources, returning errors if not.
//
// Currently supported payment source types are Card (CardParams) and
// Tokens/IDs (string), where Tokens could be single use card
// tokens
func SourceParamsFor(obj interface{}) (*PaymentSourceSourceParams, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// When you create a new credit card, you must specify a customer or recipient on which to create it.
//
// If the card's owner has no default card, then the new card will become the default.
// However, if the owner already has a default, then it will not change.
// To change the default, you should [update the customer](https://docs.stripe.com/api/customers/update) to have a new default_source.
type PaymentSourceParams struct {
	Params   `form:"*"`
	Customer *string `form:"-"` // Included in URL
	// The name of the person or business that owns the bank account.
	AccountHolderName *string `form:"account_holder_name" json:"account_holder_name,omitempty"`
	// The type of entity that holds the account. This can be either `individual` or `company`.
	AccountHolderType *string `form:"account_holder_type" json:"account_holder_type,omitempty"`
	// City/District/Suburb/Town/Village.
	AddressCity *string `form:"address_city" json:"address_city,omitempty"`
	// Billing address country, if provided when creating card.
	AddressCountry *string `form:"address_country" json:"address_country,omitempty"`
	// Address line 1 (Street address/PO Box/Company name).
	AddressLine1 *string `form:"address_line1" json:"address_line1,omitempty"`
	// Address line 2 (Apartment/Suite/Unit/Building).
	AddressLine2 *string `form:"address_line2" json:"address_line2,omitempty"`
	// State/County/Province/Region.
	AddressState *string `form:"address_state" json:"address_state,omitempty"`
	// ZIP or postal code.
	AddressZip *string `form:"address_zip" json:"address_zip,omitempty"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
	// Two digit number representing the card's expiration month.
	ExpMonth *string `form:"exp_month" json:"exp_month,omitempty"`
	// Four digit number representing the card's expiration year.
	ExpYear *string `form:"exp_year" json:"exp_year,omitempty"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// Cardholder name.
	Name  *string                   `form:"name" json:"name,omitempty"`
	Owner *PaymentSourceOwnerParams `form:"owner" json:"owner,omitempty"`
	// Please refer to full [documentation](https://api.stripe.com) instead.
	Source      *PaymentSourceSourceParams      `form:"*"` // PaymentSourceSourceParams has custom encoding so brought to top level with "*"
	Validate    *bool                           `form:"validate" json:"validate,omitempty"`
	UnsetFields []PaymentSourceParamsUnsetField `form:"-" json:"-"`
}

// PaymentSourceParamsUnsetField is the list of fields that can be cleared/unset on PaymentSourceParams.
type PaymentSourceParamsUnsetField string

const (
	PaymentSourceParamsUnsetFieldMetadata PaymentSourceParamsUnsetField = "metadata"
)

// AddUnsetField adds a field to the list of fields to clear/unset on this params object.
func (p *PaymentSourceParams) AddUnsetField(field PaymentSourceParamsUnsetField) {
	_ = "STUB: not implemented"
	return
}

// AddExpand appends a new field to expand.
func (p *PaymentSourceParams) AddExpand(f string) { _ = "STUB: not implemented"; return }

// AddMetadata adds a new key-value pair to the Metadata.
func (p *PaymentSourceParams) AddMetadata(key string, value string) {
	_ = "STUB: not implemented"
	return
}

type PaymentSourceOwnerParams struct {
	// Owner's address.
	Address *AddressParams `form:"address" json:"address,omitempty"`
	// Owner's email address.
	Email *string `form:"email" json:"email,omitempty"`
	// Owner's full name.
	Name *string `form:"name" json:"name,omitempty"`
	// Owner's phone number.
	Phone *string `form:"phone" json:"phone,omitempty"`
}

// Verify a specified bank account for a given customer.
type PaymentSourceVerifyParams struct {
	Params   `form:"*"`
	Customer *string `form:"-"` // Included in URL
	// Two positive integers, in *cents*, equal to the values of the microdeposits sent to the bank account.
	Amounts [2]int64 `form:"amounts" json:"amounts,omitempty"` // Amounts is used when verifying bank accounts
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
	Values []*string `form:"values" json:"values,omitempty"` // Values is used when verifying sources
}

// AddExpand appends a new field to expand.
func (p *PaymentSourceVerifyParams) AddExpand(f string) { _ = "STUB: not implemented"; return }

// When you create a new credit card, you must specify a customer or recipient on which to create it.
//
// If the card's owner has no default card, then the new card will become the default.
// However, if the owner already has a default, then it will not change.
// To change the default, you should [update the customer](https://docs.stripe.com/api/customers/update) to have a new default_source.
type PaymentSourceCreateParams struct {
	Params   `form:"*"`
	Customer *string `form:"-"` // Included in URL
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// Please refer to full [documentation](https://api.stripe.com) instead.
	Source   *PaymentSourceSourceParams `form:"*"` // PaymentSourceSourceParams has custom encoding so brought to top level with "*"
	Validate *bool                      `form:"validate" json:"validate,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *PaymentSourceCreateParams) AddExpand(f string) { _ = "STUB: not implemented"; return }

// AddMetadata adds a new key-value pair to the Metadata.
func (p *PaymentSourceCreateParams) AddMetadata(key string, value string) {
	_ = "STUB: not implemented"
	return
}

// Retrieve a specified source for a given customer.
type PaymentSourceRetrieveParams struct {
	Params   `form:"*"`
	Customer *string `form:"-"` // Included in URL
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *PaymentSourceRetrieveParams) AddExpand(f string) { _ = "STUB: not implemented"; return }

type PaymentSourceUpdateOwnerParams struct {
	// Owner's address.
	Address *AddressParams `form:"address" json:"address,omitempty"`
	// Owner's email address.
	Email *string `form:"email" json:"email,omitempty"`
	// Owner's full name.
	Name *string `form:"name" json:"name,omitempty"`
	// Owner's phone number.
	Phone *string `form:"phone" json:"phone,omitempty"`
}

// Update a specified source for a given customer.
type PaymentSourceUpdateParams struct {
	Params   `form:"*"`
	Customer *string `form:"-"` // Included in URL
	// The name of the person or business that owns the bank account.
	AccountHolderName *string `form:"account_holder_name" json:"account_holder_name,omitempty"`
	// The type of entity that holds the account. This can be either `individual` or `company`.
	AccountHolderType *string `form:"account_holder_type" json:"account_holder_type,omitempty"`
	// City/District/Suburb/Town/Village.
	AddressCity *string `form:"address_city" json:"address_city,omitempty"`
	// Billing address country, if provided when creating card.
	AddressCountry *string `form:"address_country" json:"address_country,omitempty"`
	// Address line 1 (Street address/PO Box/Company name).
	AddressLine1 *string `form:"address_line1" json:"address_line1,omitempty"`
	// Address line 2 (Apartment/Suite/Unit/Building).
	AddressLine2 *string `form:"address_line2" json:"address_line2,omitempty"`
	// State/County/Province/Region.
	AddressState *string `form:"address_state" json:"address_state,omitempty"`
	// ZIP or postal code.
	AddressZip *string `form:"address_zip" json:"address_zip,omitempty"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
	// Two digit number representing the card's expiration month.
	ExpMonth *string `form:"exp_month" json:"exp_month,omitempty"`
	// Four digit number representing the card's expiration year.
	ExpYear *string `form:"exp_year" json:"exp_year,omitempty"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// Cardholder name.
	Name        *string                               `form:"name" json:"name,omitempty"`
	Owner       *PaymentSourceUpdateOwnerParams       `form:"owner" json:"owner,omitempty"`
	UnsetFields []PaymentSourceUpdateParamsUnsetField `form:"-" json:"-"`
}

// PaymentSourceUpdateParamsUnsetField is the list of fields that can be cleared/unset on PaymentSourceUpdateParams.
type PaymentSourceUpdateParamsUnsetField string

const (
	PaymentSourceUpdateParamsUnsetFieldMetadata PaymentSourceUpdateParamsUnsetField = "metadata"
)

// AddUnsetField adds a field to the list of fields to clear/unset on this params object.
func (p *PaymentSourceUpdateParams) AddUnsetField(field PaymentSourceUpdateParamsUnsetField) {
	_ = "STUB: not implemented"
	return
}

// AddExpand appends a new field to expand.
func (p *PaymentSourceUpdateParams) AddExpand(f string) { _ = "STUB: not implemented"; return }

// AddMetadata adds a new key-value pair to the Metadata.
func (p *PaymentSourceUpdateParams) AddMetadata(key string, value string) {
	_ = "STUB: not implemented"
	return
}

// Delete a specified source for a given customer.
type PaymentSourceDeleteParams struct {
	Params   `form:"*"`
	Customer *string `form:"-"` // Included in URL
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *PaymentSourceDeleteParams) AddExpand(f string) { _ = "STUB: not implemented"; return }

type PaymentSource struct {
	APIResource
	BankAccount *BankAccount      `json:"-"`
	Card        *Card             `json:"-"`
	Deleted     bool              `json:"deleted,omitempty"`
	ID          string            `json:"id,omitempty"`
	Source      *Source           `json:"-"`
	Type        PaymentSourceType `json:"object"`
}

// PaymentSourceList is a list of PaymentSources as retrieved from a list endpoint.
type PaymentSourceList struct {
	APIResource
	ListMeta
	Data []*PaymentSource `json:"data"`
}

// UnmarshalJSON handles deserialization of a PaymentSource.
// This custom unmarshaling is needed because the specific
// type of payment instrument it refers to is specified in the JSON
func (s *PaymentSource) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

// MarshalJSON handles serialization of a PaymentSource.
// This custom marshaling is needed because the specific type
// of payment instrument it represents is specified by the Type
func (s *PaymentSource) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
