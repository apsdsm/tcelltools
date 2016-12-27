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

package tcelltools_test

import (
	"github.com/apsdsm/binder/fakes"
	. "github.com/apsdsm/tcelltools"
	"github.com/gdamore/tcell"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("drawing wrapped text", func() {

	var (
		screen *fakes.ScreenBridge
		style  = tcell.StyleDefault
	)

	BeforeEach(func() {
		screen = fakes.NewScreenBridge(100, 100)
	})

	It("paints over a section of the screen with a single rune", func() {
		Paint(screen, 0, 0, 2, 2, '!', style)

		for i := 0; i < 3; i++ {
			Expect(screen.GetLine(i, 0, 2)).To(Equal("!!!"))
		}
	})

	It("paints as many double width characters as possible and pads remaining cells with space", func() {
		Paint(screen, 0, 0, 2, 2, '萌', style)

		for i := 0; i < 3; i++ {
			Expect(screen.GetLine(i, 0, 2)).To(Equal("萌 "))
		}
	})
})