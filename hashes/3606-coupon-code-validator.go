package hashes

import (
	"sort"
	"unicode"
)

// https://leetcode.com/problems/coupon-code-validator/

// Approach: Filtering + Sorting with Custom Comparator
// Time: O(n log n) - sorting valid coupons
// Space: O(n) - storing valid coupons
func validateCoupons(code []string, businessLine []string, isActive []bool) []string {
	// Define business line priority order
	priority := map[string]int{
		"electronics": 0,
		"grocery":     1,
		"pharmacy":    2,
		"restaurant":  3,
	}

	isValidCode := func(c string) bool {
		// Check if code is non-empty and contains only alphanumeric chars and underscores
		if len(c) == 0 {
			return false
		}

		for _, ch := range c {
			if !unicode.IsLetter(ch) && !unicode.IsDigit(ch) && ch != '_' {
				return false
			}
		}

		return true
	}

	// Collect valid coupons as structs
	type Coupon struct {
		priority int
		code     string
	}

	var validCoupons []Coupon

	for i := range code {
		_, validBusiness := priority[businessLine[i]]
		if isActive[i] && isValidCode(code[i]) && validBusiness {
			validCoupons = append(validCoupons, Coupon{
				priority: priority[businessLine[i]],
				code:     code[i],
			})
		}
	}

	// Sort by priority first, then by code lexicographically
	sort.Slice(validCoupons, func(i, j int) bool {
		if validCoupons[i].priority != validCoupons[j].priority {
			return validCoupons[i].priority < validCoupons[j].priority
		}

		return validCoupons[i].code < validCoupons[j].code
	})

	// Extract and return just the codes
	result := make([]string, len(validCoupons))
	for i, coupon := range validCoupons {
		result[i] = coupon.code
	}

	return result
}
