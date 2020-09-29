/*
 * Square
 *
 * Use Square APIs to manage and run business including payment, customer, product, inventory, and employee management.
 *
 * API version: 2.0
 * Contact: developers@squareup.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// V1Discount
type V1Discount struct {
	// The discount's unique ID.
	Id string `json:"id,omitempty"`
	// The discount's name.
	Name string `json:"name,omitempty"`
	// The rate of the discount, as a string representation of a decimal number. A value of 0.07 corresponds to a rate of 7%. This rate is 0 if discount_type is VARIABLE_PERCENTAGE.
	Rate         string                  `json:"rate,omitempty"`
	AmountMoney  *V1Money                `json:"amount_money,omitempty"`
	DiscountType *V1DiscountDiscountType `json:"discount_type,omitempty"`
	// Indicates whether a mobile staff member needs to enter their PIN to apply the discount to a payment.
	PinRequired bool             `json:"pin_required,omitempty"`
	Color       *V1DiscountColor `json:"color,omitempty"`
	// The ID of the CatalogObject in the Connect v2 API. Objects that are shared across multiple locations share the same v2 ID.
	V2Id string `json:"v2_id,omitempty"`
}