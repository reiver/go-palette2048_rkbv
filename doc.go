/*
Package palette2048_rkbv provides the RKBV color palette in the form of 256 RGBA color with 8-bits per color channel,
that when used with the https://github.com/reiver/go-palette2048 package works with Go's built-in "image", "image/color", and "image/draw" packages.

Example

Here is an example of how this package can be used:

	import (
		"github.com/reiver/go-palette2048"
		"github.com/reiver/go-palette2048_rkbv"
	)
	
	// ...
	
	palette := palette2048.Slice(palette2048_rkbv.Palette)
	
	// ...
	
	// Get the color at a specific index in the palette.
	color := palette.Color(index)
	
	// ...
	
	// Get the color in the palette that is closet to a reference color.
	closestColor := palette.Convert(referenceColor)

*/
package palette2048_rkbv
