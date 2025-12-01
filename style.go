package console

type Alignment uint8

const (
	Left = iota
	Right
	Center
)

type Color uint8

const (
	// Use no color but the default
	None = iota
	Black
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	LightGray
	Gray
	BrightRed
	BrightGreen
	BrightYellow
	BrightBlue
	BrightMagenta
	BrightCyan
	White
)

type Stylesheet struct {
	Alignment Alignment
	Width     int
	Margin    [2]int
	Sequence  [8]Cell
}

type color_key struct {
	Bg    bool
	Color Color
}

var color_codes = map[color_key]uint8{
	{false, None}: 0,
	{true, None}:  0,

	{false, Black}:         30,
	{false, Red}:           31,
	{false, Green}:         32,
	{false, Yellow}:        33,
	{false, Blue}:          34,
	{false, Magenta}:       35,
	{false, Cyan}:          36,
	{false, LightGray}:     37,
	{false, Gray}:          90,
	{false, BrightRed}:     91,
	{false, BrightGreen}:   92,
	{false, BrightYellow}:  93,
	{false, BrightBlue}:    94,
	{false, BrightMagenta}: 95,
	{false, BrightCyan}:    96,
	{false, White}:         97,

	{true, Black}:         40,
	{true, Red}:           41,
	{true, Green}:         42,
	{true, Yellow}:        43,
	{true, Blue}:          44,
	{true, Magenta}:       45,
	{true, Cyan}:          46,
	{true, LightGray}:     47,
	{true, Gray}:          100,
	{true, BrightRed}:     101,
	{true, BrightGreen}:   102,
	{true, BrightYellow}:  103,
	{true, BrightBlue}:    104,
	{true, BrightMagenta}: 105,
	{true, BrightCyan}:    106,
	{true, White}:         107,
}
