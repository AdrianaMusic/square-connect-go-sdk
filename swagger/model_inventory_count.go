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

// Represents Square's estimated quantity of items in a particular state at a particular location based on the known history of physical counts and inventory adjustments.
type InventoryCount struct {
	// The Square generated ID of the `CatalogObject` being tracked.
	CatalogObjectId string `json:"catalog_object_id,omitempty"`
	// The `CatalogObjectType` of the `CatalogObject` being tracked. Tracking is only supported for the `ITEM_VARIATION` type.
	CatalogObjectType string          `json:"catalog_object_type,omitempty"`
	Status            *InventoryState `json:"status,omitempty"`
	State             *InventoryState `json:"state,omitempty"`
	// The Square ID of the [Location](#type-location) where the related quantity of items are being tracked.
	LocationId string `json:"location_id,omitempty"`
	// The number of items affected by the estimated count as a decimal string. Can support up to 5 digits after the decimal point.
	Quantity string `json:"quantity,omitempty"`
	// A read-only timestamp in RFC 3339 format that indicates when Square received the most recent physical count or adjustment that had an affect on the estimated count.
	CalculatedAt string `json:"calculated_at,omitempty"`
}