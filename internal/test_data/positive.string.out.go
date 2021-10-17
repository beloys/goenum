package test_data

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
