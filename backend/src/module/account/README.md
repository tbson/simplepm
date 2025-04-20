# Account Security Module

This module handles security mechanisms for user accounts, including OTP generation, validation, and policy enforcement.

## Domain Concepts

- **OTP**: A one-time password used for verification purposes
- **OTP Policy**: Rules governing the format and validation of OTPs
- **Security Token**: Authentication tokens used throughout the application

## Architectural Decisions

### OTP as a Value Object

We model OTPs as value objects rather than entities because they:
- Are defined entirely by their value
- Have no separate identity that persists across state changes
- Are immutable once created

### Direct Import for Policy Access

In this module, we use direct imports for policy access rather than dependency injection because:
1. OTP policies are stable application-level configurations
2. Policy parameters are retrieved from local config, not external systems
3. Policy validation is intrinsic to what makes a valid OTP

Example:
```go
// OTP directly imports policies
func NewOTP() OTP {
    otpLength := policy.GetOtpLength()
    // ...
}
