package main

// Formatter is a bridge-interface.
type Formatter interface {
	Format(s string) string
}

type BWFormatter struct {
	Fmt BackwardsFormatter
}

func (bf BWFormatter) Format(s string) string {
	return bf.Fmt.Backwards(s)
}

type TBFormatter struct {
	Fmt TabFormatter
}

func (tf TBFormatter) Format(s string) string {
	return tf.Fmt.SpaceToTab(s)
}
