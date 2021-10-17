# Enumerations for golang
Check in [test_data](./test_data) `*.out.go` files for a generated output

## Options
Use as CLI arguments
- `-p,--priority` - Generate priority methods (Priority,Compare,GreaterThan,LessThan) based on definition order

```go
package example

//go:generate goenum
type Severity string

const (
	SeverityHigh   Severity = "HIGH"
	SeverityLow    Severity = "LOW"
	SeverityMedium Severity = "MEDIUM"
)
```