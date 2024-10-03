package main

// BackwardsFormatter formats string backwards.
type BackwardsFormatter interface {
	Backwards(s string) string
}

// Backwarder implements BackwardsFormatter
type Backwarder struct{}

// Backwards implementation of BackwardsFormatter.
func (b Backwarder) Backwards(s string) string {
	l := len(s)
	bs := make([]byte, 0, l)
	for i := range s {
		bs = append(bs, s[l-1-i])
	}
	return string(bs)
}
