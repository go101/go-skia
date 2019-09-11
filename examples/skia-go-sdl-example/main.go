/*
 * Copyright 2015 Google Inc.
 *
 * Use of this source code is governed by a BSD-style license that can be
 * found in the LICENSE file.
 */
package main

/*
	This is a translation of the example code in
	"experimental/c-api-example/skia-c-example.c" to test the
	go-skia package.
*/

import (
	"log"
	"os"
	"reflect"
	"unsafe"

	"go101.org/skia"
	"github.com/veandco/go-sdl2/sdl"
)


func run() int {

	var winTitle string = "Go-SDL2 Render"
	var winWidth, winHeight int32 = 800, 600

	window, err := sdl.CreateWindow(winTitle, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		winWidth, winHeight, sdl.WINDOW_SHOWN)
	if err != nil {
		log.Printf("Failed to create window: %s\n", err)
		return 1
	}
	defer window.Destroy()

	sdlSurface, err := window.GetSurface()
	if err != nil {
		log.Printf("Failed to get window surface: %s\n", err)
		return 1
	}
	defer sdlSurface.Free()

	// Create skia surface and canvas.
	imgInfo := &skia.ImageInfo{
		Width:     640,
		Height:    480,
		ColorType: skia.BGRA_8888_COLORTYPE,
		AlphaType: skia.PREMUL_ALPHATYPE,
	}
	pixels := sdlSurface.Pixels()
	pixelsAddr := ((*reflect.SliceHeader)(unsafe.Pointer(&pixels))).Data
	rowBytes := int(sdlSurface.Pitch)

	surface, err := skia.NewDirectRasterSurface(imgInfo, pixelsAddr, rowBytes)
	if err != nil {
		log.Printf("Failed to create direct raster surface: %s\n", err)
		return 1
	}

	canvas := surface.Canvas()

	// ...
	var radius float32 = 100

	var event sdl.Event
	var running = true
	for running {
		for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.QuitEvent:
				running = false
				_ = t
			//case *sdl.MouseMotionEvent:
			//	fmt.Printf("[%d ms] MouseMotion\tid:%d\tx:%d\ty:%d\txrel:%d\tyrel:%d\n", t.Timestamp, t.Which, t.X, t.Y, t.XRel, t.YRel)
			}
		}

		func() {
			// needed?
			sdlSurface.Lock()
			defer sdlSurface.Unlock()

			fillPaint := skia.NewPaint()
			fillPaint.SetColor(0xFF0000FF)
			canvas.DrawPaint(fillPaint)
			fillPaint.SetColor(0xFF00FFFF)
			fillPaint.SetAntiAlias(true)
			//fillPaint.SetStroke(true)

			rect := skia.Rect{100, 100, 540, 380}
			canvas.DrawRect(&rect, fillPaint)

			strokePaint := skia.NewPaint()
			strokePaint.SetColor(0xFFFF0000)
			strokePaint.SetAntiAlias(true)
			strokePaint.SetStroke(true)
			strokePaint.SetStrokeWidth(5.0)

			path := skia.NewPath()
			path.MoveTo(50, 50)
			path.LineTo(590, 50)
			path.CubicTo(-490, 50, 1130, 430, 50, 430)
			path.LineTo(590, 430)
			canvas.DrawPath(path, strokePaint)

			fillPaint.SetColor(0x8000FF00)
			canvas.DrawOval(&skia.Rect{120, 120, 520, 360}, fillPaint)

			fillPaint.SetColor(0x80FF00FF)
			radius += 0.1
			if radius > 200 {
				radius = 1
			}
			canvas.DrawCircle(320, 240, radius, fillPaint)

			window.UpdateSurface()
		}()
	}

	return 0
}

func main() {
	os.Exit(run())
}
