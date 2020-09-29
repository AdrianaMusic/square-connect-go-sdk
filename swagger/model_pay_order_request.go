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

// Defines the fields that are included in requests to the [PayOrder](#endpoint-payorder) endpoint.
type PayOrderRequest struct {
	// A value you specify that uniquely identifies this request among requests you've sent. If you're unsure whether a particular payment request was completed successfully, you can reattempt it with the same idempotency key without worrying about duplicate payments.  See [Idempotency](https://developer.squareup.com/docs/working-with-apis/idempotency) for more information.
	IdempotencyKey string `json:"idempotency_key"`
	// The version of the order being paid. If not supplied, the latest version will be paid.
	OrderVersion int32 `json:"order_version,omitempty"`
	// The IDs of the [payments](#type-payment) to collect. The payment total must match the order total.
	PaymentIds []string `json:"payment_ids,omitempty"`
}