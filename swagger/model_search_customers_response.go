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

// Defines the fields that are included in the response body of a request to the SearchCustomers endpoint.  One of `errors` or `customers` is present in a given response (never both).
type SearchCustomersResponse struct {
	// Any errors that occurred during the request.
	Errors []ModelError `json:"errors,omitempty"`
	// An array of `Customer` objects that match a query.
	Customers []Customer `json:"customers,omitempty"`
	// A pagination cursor that can be used during subsequent calls to SearchCustomers to retrieve the next set of results associated with the original query. Pagination cursors are only present when a request succeeds and additional results are available.  See the [Pagination guide](https://developer.squareup.com/docs/working-with-apis/pagination) for more information.
	Cursor string `json:"cursor,omitempty"`
}