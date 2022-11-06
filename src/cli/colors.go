package cli

const (
	ColorDefault = "\x1b[39m"

	ColorRed   = "\x1b[91m"
	ColorGreen = "\x1b[32m"
	ColorBlue  = "\x1b[94m"
	ColorGray  = "\x1b[90m"
)

type Color interface {
	Red(string) string
	Green(string) string
	Blue(string) string
	Gray(string) string
}

type ColorImpl struct{}

func NewColor() *ColorImpl {
	return &ColorImpl{}
}

func (c *ColorImpl) Red(s string) string {
	return ColorRed + s + ColorDefault
}

func (c *ColorImpl) Green(s string) string {
	return ColorGreen + s + ColorDefault
}

func (c *ColorImpl) Blue(s string) string {
	return ColorBlue + s + ColorDefault
}

func (c *ColorImpl) Gray(s string) string {
	return ColorGray + s + ColorDefault
}