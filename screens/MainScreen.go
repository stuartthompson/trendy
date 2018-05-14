// Copyright 2018 Stuart Thompson.

// This file is part of Trendy.

// Trendy is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// Trendy is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with Trendy.  If not, see <http://www.gnu.org/licenses/>.

package screens

import (
	"github.com/stuartthompson/trendy/io"
)

// MainScreen ...
type MainScreen struct {
}

// Render ...
// Renders the main screen.
func (s *MainScreen) Render() {
	io.ClearScreen(0)
	width, height := io.GetWindowSize()
	io.RenderPaneBorder(0, 0, width-1, height-1, 82, 0)
	io.RenderText("Trendy", 1, 1, 255, 0)

	io.Flush()
}
