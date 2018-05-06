/*
 * goBilling
 *
 * goBilling is an open-source billing and payments platform
 *
 * API version: 1.0.0
 *
 */

package model

type BulkBaseSubscriptionAndAddOnsJson struct {

	BaseEntitlementAndAddOns []SubscriptionJson `json:"baseEntitlementAndAddOns"`
}
