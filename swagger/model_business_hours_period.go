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

// Represents a period of time during which a business location is open.
type BusinessHoursPeriod struct {
	DayOfWeek *DayOfWeek `json:"day_of_week,omitempty"`
	// The start time of a business hours period, specified in local time using partial-time RFC3339 format.
	StartLocalTime string `json:"start_local_time,omitempty"`
	// The end time of a business hours period, specified in local time using partial-time RFC3339 format.
	EndLocalTime string `json:"end_local_time,omitempty"`
}