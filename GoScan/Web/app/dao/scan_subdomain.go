// ============================================================================
// This is auto-generated by gf cli tool only once. Fill this file as you wish.
// ============================================================================

package dao

import (
	"github.com/CTF-MissFeng/GoScan/Web/app/dao/internal"
)

// scanSubdomainDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type scanSubdomainDao struct {
	*internal.ScanSubdomainDao
}

var (
	// ScanSubdomain is globally public accessible object for table scan_subdomain operations.
	ScanSubdomain = &scanSubdomainDao{
		internal.ScanSubdomain,
	}
)

// Fill with you ideas below.