package controllers

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"math/rand"
	"strconv"
	"strings"

	"github.com/goravel/framework/contracts/http"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/gobold"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/math/fixed"
)

type ImageController struct {
	// Dependent services
}

func NewImageController() *ImageController {
	return &ImageController{
		// Inject services
	}
}

func (r *ImageController) Index(ctx http.Context) http.Response {
	dimensionsParam := ctx.Request().Route("widthXHeight")

	cleanDimensions := strings.TrimSuffix(dimensionsParam, ".png")

	// Split the dimension string (e.g., "100x100") into width and height parts.
	parts := strings.Split(cleanDimensions, "x")
	if len(parts) != 2 {
		return ctx.Response().String(400, "Invalid dimensions format. Must be in the format 'widthxheight', e.g., '100x100'.")
	}

	// Convert the width string to an integer. Ensure it's a positive value.
	width, err := strconv.Atoi(parts[0])
	if err != nil || width <= 0 || width > 4000 {
		return ctx.Response().String(400, "Invalid width. Must be a positive integer between 1 and 4000.")
	}

	// Convert the height string to an integer. Ensure it's a positive value.
	height, err := strconv.Atoi(parts[1])
	if err != nil || height <= 0 || height > 4000 {
		return ctx.Response().String(400, "Invalid height. Must be a positive integer between 1 and 4000.")
	}

	// Create a new RGBA image with the specified width and height.
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	backgroundColors := []color.RGBA{
		{R: 118, G: 181, B: 197, A: 255}, // Light Blue
		{R: 255, G: 160, B: 122, A: 255}, // Salmon
		{R: 144, G: 238, B: 144, A: 255}, // Light Green
		{R: 255, G: 215, B: 0, A: 255},   // Gold
		{R: 186, G: 85, B: 211, A: 255},  // Medium Purple
	}

	// Random bg color
	bgColor := backgroundColors[rand.Intn(len(backgroundColors))]

	// Fill every pixel of the image with the chosen background color.
	for y := range height {
		for x := range width {
			img.Set(x, y, bgColor)
		}
	}

	texts := []string{
		"Axolotl",
		"Okapi",
		"Quokka",
		"Tarsius",
		"Iriomote",
		"Narwhal",
		"Manul",
		"Zorilla",
		"Kakapo",
		"Saola",
		fmt.Sprintf("%dx%d", width, height),
	}

	randomText := texts[rand.Intn(len(texts))]

	textColors := []color.RGBA{
		{R: 255, G: 255, B: 255, A: 255}, // White
		{R: 0, G: 0, B: 0, A: 255},       // Black
		{R: 50, G: 50, B: 50, A: 255},    // Dark Gray
	}

	// Pick a random text color from the list.
	txtColor := textColors[rand.Intn(len(textColors))]

	// Parse the embedded Go Bold font data.
	ttfFont, err := opentype.Parse(gobold.TTF)
	if err != nil {
		log.Printf("Failed to parse font: %v", err)
		return ctx.Response().String(500, "Failed to parse font.")
	}

	// Determine the font size dynamically, scaled to image dimensions.
	// This heuristic tries to make the text fit well for various sizes.
	fontSize := float64(width) / 5.0
	if fontSize > float64(height)/3.0 {
		fontSize = float64(height) / 3.0
	}
	if fontSize < 10 { // minimum readable font size
		fontSize = 10
	}

	// Create a font face from the parsed font with the calculated size and standard DPI.
	face, err := opentype.NewFace(ttfFont, &opentype.FaceOptions{
		Size:    fontSize,
		DPI:     72,
		Hinting: font.HintingNone, // No hinting for cleaner rendering
	})

	if err != nil {
		log.Printf("Failed to create font face: %v", err)
		return ctx.Response().String(500, "Failed to create font face.")
	}
	defer face.Close() // Release font face

	// Create a Drawer
	dr := &font.Drawer{
		Dst:  img,                        // Destination image for drawing
		Src:  image.NewUniform(txtColor), // Source color for the text
		Face: face,                       // Font face to use
	}

	// Measure the width and height of the text to be drawn.
	textWidth := dr.MeasureString(randomText).Round()
	textHeight := face.Metrics().Ascent.Round() + face.Metrics().Descent.Round()

	// Calculate the X and Y coordinates to center the text on the image.
	// Y-coordinate accounts for font baseline (Ascent).
	x := (width - textWidth) / 2
	y := (height-textHeight)/2 + face.Metrics().Ascent.Round()

	// Set the starting point (baseline) for drawing the text.
	dr.Dot = fixed.Point26_6{
		X: fixed.I(x),
		Y: fixed.I(y),
	}

	// Draw the randomly selected text onto the image.
	dr.DrawString(randomText)

	ctx.Request().Header("Content-Type", "image/png")
	ctx.Request().Header("Cache-Control", "public, max-age=31536000")

	// Encode the generated image as PNG and write it to the HTTP response writer.
	err = png.Encode(ctx.Response().Writer(), img)
	if err != nil {
		return ctx.Response().String(500, "Failed to encode image.")
	}

	return ctx.Response().NoContent()
}
