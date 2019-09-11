module go101.org/go-skia/skia-go-sdl-example

replace go101.org/skia => ../..

replace github.com/veandco/go-sdl2 => github.com/go-graphics/go-sdl2 v0.3.0

require (
	github.com/veandco/go-sdl2 v0.0.0-00010101000000-000000000000
	go101.org/skia v0.0.0-00010101000000-000000000000
)

go 1.13
