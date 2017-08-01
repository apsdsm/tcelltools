//    Copyright 2016 Nick del Pozo
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.

package canvas

import "github.com/gdamore/tcell"

// A Canvas is a drawable area on top of a screen interface.
type Canvas struct {
	screen ScreenBridge
	Layers []*Layer
}

// NewCanvas returns a new, uninitiated canvas object.
func NewCanvas(screen ScreenBridge) *Canvas {
	c := Canvas{}
	c.screen = screen
	c.Layers = make([]*Layer, 0, 5)

	return &c
}

// AddLayer adds a new blank layer to a canvas object.
func (c *Canvas) AddLayer(layers ...*Layer) {

	if len(layers) == 0 {
		width, height := c.screen.Size()
		l := NewLayer(width, height, 0, 0)
		c.Layers = append([]*Layer{l}, c.Layers...)
		return
	}

	for _, l := range layers {
		c.Layers = append([]*Layer{l}, c.Layers...)
	}
}

// Draw will draw the contents of all layers to the screen.
func (c *Canvas) Draw() {

	style := tcell.StyleDefault.Foreground(tcell.ColorWhite)

	width, height := c.screen.Size()

	if width > 0 {
		width--
	}

	if height > 0 {
		height--
	}

	// for each cell in the screen
	for x := 0; x <= width; x++ {
		for y := 0; y < height; y++ {

			// for each layer
			for _, layer := range c.Layers {

				// get the rune at that position if it exists
				offsetX := x - layer.X
				offsetY := y - layer.Y

				// if there is a layer cell at that position which is not empty
				if offsetX >= 0 && offsetY >= 0 && layer.Grid[offsetX][offsetY].Rune != 0 {
					c.screen.SetContent(x, y, layer.Grid[offsetX][offsetY].Rune, nil, style)
					break
				}
			}
		}
	}
}
