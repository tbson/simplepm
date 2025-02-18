package profiletype

// For platform owner
const ADMIN = "ADMIN"
const STAFF = "STAFF"

// For customer
const OWNER = "OWNER"
const MANAGER = "MANAGER"
const USER = "USER"

var PlatformProfileTypes = []string{ADMIN, STAFF}
var TenantProfileTypes = []string{OWNER, MANAGER, USER}
