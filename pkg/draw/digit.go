package draw

import (
	"fmt"
	"image/color"

	"tinygo.org/x/drivers"
)

// Digit renders a LED-segment style digit, bounded by the specified points.
func Digit(display drivers.Displayer, digit byte, x0 int16, y0 int16, x1 int16, y1 int16, color color.RGBA) {
	// Display elements are:
	//  <0>
	// ^   ^
	// 1   2
	// V   V
	//  <3>
	// ^   ^
	// 4   5
	// V   V
	//  <6>
	var segments []bool
	switch digit {
	case '0':
		segments = []bool{
			/**/ true,
			true /**/, true,
			/**/ false,
			true /**/, true,
			/**/ true,
		}
	case '1':
		segments = []bool{
			/**/ false,
			false /**/, true,
			/**/ false,
			false /**/, true,
			/**/ false,
		}
	case '2':
		segments = []bool{
			/**/ true,
			false /**/, true,
			/**/ true,
			true /**/, false,
			/**/ true,
		}
	case '3':
		segments = []bool{
			/**/ true,
			false /**/, true,
			/**/ true,
			false /**/, true,
			/**/ true,
		}
	case '4':
		segments = []bool{
			/**/ false,
			true /**/, true,
			/**/ true,
			false /**/, true,
			/**/ false,
		}
	case '5':
		segments = []bool{
			/**/ true,
			true /**/, false,
			/**/ true,
			false /**/, true,
			/**/ true,
		}
	case '6':
		segments = []bool{
			/**/ true,
			true /**/, false,
			/**/ true,
			true /**/, true,
			/**/ true,
		}
	case '7':
		segments = []bool{
			/**/ true,
			false /**/, true,
			/**/ false,
			false /**/, true,
			/**/ false,
		}
	case '8':
		segments = []bool{
			/**/ true,
			true /**/, true,
			/**/ true,
			true /**/, true,
			/**/ true,
		}
	case '9':
		segments = []bool{
			/**/ true,
			true /**/, true,
			/**/ true,
			false /**/, true,
			/**/ true,
		}
	case '-':
		segments = []bool{
			/**/ false,
			false /**/, false,
			/**/ true,
			false /**/, false,
			/**/ false,
		}
	case ' ':
		segments = []bool{
			/**/ false,
			false /**/, false,
			/**/ false,
			false /**/, false,
			/**/ false,
		}
	default:
		// Show a capital "E" for error.
		segments = []bool{
			/**/ true,
			true /**/, false,
			/**/ true,
			true /**/, false,
			/**/ true,
		}
	}
	fmt.Printf("%v\n", segments)
}
