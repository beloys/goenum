# Enumerations for golang
Check in [test_data](./test_data) `*.out.go` files for a generated output

## Options
Use as CLI arguments
- `-p,--priority` - Generate priority methods (Priority,Compare,GreaterThan,LessThan) based on definition order

```go
package example

//go:generate goenum -p
type Severity string

const (
	SeverityHigh   Severity = "HIGH"
	SeverityLow    Severity = "LOW"
	SeverityMedium Severity = "MEDIUM"
)

// Generated below

var prioritySeverity = []Severity{
	SeverityLow,
	SeverityMedium,
	SeverityHigh,
}

func (e Severity) String() string {
	return string(e)
}

func (e Severity) In(values ...Severity) bool {
	for _, v := range values {
		if v == e {
			return true
		}
	}
	return false
}

func (e Severity) Priority() int {
	for k, v := range prioritySeverity {
		if v == e {
			return k
		}
	}
	return -1
}

func (e Severity) Compare(v Severity) int {
	if e.Priority() > v.Priority() {
		return 1
	}
	if e.Priority() < v.Priority() {
		return -1
	}
	return 0
}

func (e Severity) LessThan(v Severity) bool {
	return e.Priority() < v.Priority()
}

func (e Severity) GreaterThan(v Severity) bool {
	return e.Priority() > v.Priority()
}
```

## Todo
- Tests
- Go fmt