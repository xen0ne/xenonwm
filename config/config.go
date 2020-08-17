package config

var (
	// Bar contans all config items related to window title decorations
	Bar struct {
		// Hex Colors of bar background
		Focused, UnFocused uint32
		// Height of the bar in pixels
		Height int
		// Should the title text be centered?
		CenterText bool
		// If the title text is not centered, the text offset from the ends
		TextOffset int
	}
)

// Defaults sets all the config variables to their default values
func Defaults() {
	// set defaults
	Bar.Focused = 0xc1c1c1
	Bar.UnFocused = 0x3f3f3f
	Bar.Height = 20
	Bar.CenterText = true
	Bar.TextOffset = 5
}
