/*
 * Copyright 2019 Tapir Liu.
 *
 * Copyright 2015 Google Inc.
 *
 * Use of this source code is governed by a BSD-style license that can be
 * found in the LICENSE file.
 */
package skia

// ref: https://golang.org/cmd/cgo/

/*
//#cgo LDFLAGS: -L${SRCDIR}/skia/lib/shared/linux/amd64
//#cgo LDFLAGS: -Wl,-rpath,${SRCDIR}/skia/lib/shared/linux/amd64
//#cgo LDFLAGS: -lskia
#cgo linux amd64 LDFLAGS: -L${SRCDIR}/skia/lib/static/linux/amd64
#cgo linux amd64 LDFLAGS: -Wl,-rpath,${SRCDIR}/skia/lib/static/linux/amd64
#cgo linux LDFLAGS:  -lstdc++ -lm
// -lwebp -lwebpmux -lpng -lpng16 -ljpeg -lz -ldl -luuid -lfreetype -lfontconfig -lexpat
#cgo LDFLAGS: -lskia
#cgo CFLAGS: -I./skia/include/c
#include "sk_canvas.h"
#include "sk_data.h"
#include "sk_image.h"
#include "sk_paint.h"
#include "sk_path.h"
#include "sk_surface.h"
*/
import "C"

import (
	"fmt"
	"io"
	"math"
	"runtime"
	"unsafe"
)

const (
	radian2degrees = 180.0 / math.Pi
	degree2radians = math.Pi / 180.0
)

// A Surface holds the destination for drawing to a canvas. For
// raster drawing, the destination is an array of pixels in memory.
// For GPU drawing, the destination is a texture or a framebuffer.
type Surface struct {
	ptr *C.sk_surface_t
}

// NewRasterSurface creates a new Surface, with the properties specified
// by imgInfo, and with the memory for the pixels automatically allocated.
func NewRasterSurface(imgInfo ImageInfo) (*Surface, error) {
	ptr := C.sk_surface_new_raster(imgInfo.cPointer(), (*C.sk_surfaceprops_t)(nil))
	if ptr == nil {
		return nil, fmt.Errorf("Unable to create raster surface.")
	}

	ret := &Surface{ptr: ptr}
	// Todo: use an alternative way to release ref.
	runtime.SetFinalizer(ret, func(s *Surface) {
		C.sk_surface_unref(s.ptr)
	})
	return ret, nil
}

// NewDirectRasterSurface creates a new Surface which will draw into the
// memory specified by pixels, with the number of bytes per row specified
// by rowbytes, and with the properties specified by imgInfo,
// NOTE, it is the caller's duty to ensure the memory is still alive during
// the lifetime of the returned Surface.
func NewDirectRasterSurface(imgInfo ImageInfo, pixels uintptr, rowBytes int) (*Surface, error) {
	ptr := C.sk_surface_new_raster_direct(imgInfo.cPointer(), unsafe.Pointer(pixels), C.size_t(rowBytes), (*C.sk_surfaceprops_t)(nil))
	if ptr == nil {
		return nil, fmt.Errorf("Unable to create direct raster surface.")
	}

	ret := &Surface{ptr: ptr}
	runtime.SetFinalizer(ret, func(s *Surface) {
		C.sk_surface_unref(s.ptr)
	})
	return ret, nil
}

// Canvas returns the canvas associated with the Surface s.
// Note, each Surface owns a Canvas,
// 1. so the returned Canvas is only valid while the owning surface is valid.
// 2. the Canvases returned by two Canvas method calls on the same Surface
//    hold the same internal object.
func (s *Surface) Canvas() *Canvas {
	return &Canvas{
		ptr:             C.sk_surface_get_canvas(s.ptr),
		keepParentAlive: s,
	}
}

// ImageSnapshot returns an Image which represents a snapshot of the Surface s.
func (s *Surface) Snapshot() *Image {
	ret := &Image{
		ptr:             C.sk_surface_new_image_snapshot(s.ptr),
		keepParentAlive: s,
	}
	// Todo: use an alternative way to release ref.
	runtime.SetFinalizer(ret, func(i *Image) {
		C.sk_image_unref(i.ptr)
	})
	return ret
}

// An Image is an abstraction for drawing a rectagle of pixels.
// The content of the image is always immutable, though the actual
// storage may change, if for example that image can be re-created via
// encoded data or other means.
type Image struct {
	ptr             *C.sk_image_t
	keepParentAlive *Surface
}

// WritePNG encodes the image pixels as PNG format and write the PNG data into a Writer.
func (i *Image) WritePNG(w io.Writer) error {
	data := C.sk_image_encode(i.ptr)
	defer C.sk_data_unref(data)

	dataPtr := C.sk_data_get_data(data)
	dataSize := C.sk_data_get_size(data)
	byteSlice := C.GoBytes(dataPtr, C.int(dataSize))
	_, err := w.Write(byteSlice)
	if err != nil {
		return err
	}
	return nil
}

// A Canvas encapsulates all of the state about drawing into a
// destination. This includes a reference to the destination itself,
// and a stack of matrix/clip values.
//
// A Canvas needs a Paint to cooperate in drawing.
type Canvas struct {
	ptr             *C.sk_canvas_t
	keepParentAlive *Surface
}

// Save creates a new restore point from the current matrix and clip on the canvas.
// The new restore point is pushed to a stack maintained by the Canvas c.
// When the balancing call to sk_canvas_restore() is made,
// the previous matrix and clip are restored.
func (c *Canvas) Save() {
	C.sk_canvas_save(c.ptr)
}

// Restore pops the last created restore point as the current one.
// It does nothing if the restore point stack is empty.
func (c *Canvas) Restore() {
	C.sk_canvas_restore(c.ptr)
}

// Translate pre-concats the current coordinate transformation matrix
// with the specified translation.
func (c *Canvas) Translate(dx, dy float64) {
	C.sk_canvas_translate(c.ptr, C.float(dx), C.float(dy))
}

// Scale pre-concats the current coordinate transformation matrix
// with the specified scale.
func (c *Canvas) Scale(sx, sy float64) {
	C.sk_canvas_scale(c.ptr, C.float(sx), C.float(sy))
}

// RotateDegrees pre-concats the current coordinate transformation matrix
// with the specified roration in degrees.
func (c *Canvas) RotateDegrees(degrees float64) {
	C.sk_canvas_rotate_degrees(c.ptr, C.float(degrees))
}

// RotateRadians pre-concats the current coordinate transformation matrix
// with the specified roration in radians.
func (c *Canvas) RotateRadians(radians float64) {
	C.sk_canvas_rotate_radians(c.ptr, C.float(radians))
}

// DrawPaint fills the entire canvas (restricted to the current clip)
// with the specified paint.
func (c *Canvas) DrawPaint(paint *Paint) {
	C.sk_canvas_draw_paint(c.ptr, paint.ptr)
}

// DrawCircle draws a circle centered at (cx, cy) with radius r
// using the specified paint. The circle will be filled or stroked
// based on the style in the specified paint.
func (c *Canvas) DrawCircle(cx, cy, r float32, paint *Paint) {
	C.sk_canvas_draw_circle(c.ptr, C.float(cx), C.float(cy), C.float(r), paint.ptr)
}

// DrawOval draws an oval with the specified paint. The oval, which is bounded
// in rect,  will be filled or stroked based on the style in the specified paint.
func (c *Canvas) DrawOval(rect Rect, paint *Paint) {
	// C.sk_canvas_draw_oval(c.ptr, (*C.sk_rect_t)(unsafe.Pointer(rect)), (*C.sk_paint_t)(paint.ptr))
	C.sk_canvas_draw_oval(c.ptr, rect.cPointer(), paint.ptr)
}

// DrawRect draws a rectangle with the specified paint. The rectangle
// will be filled or stroked based on the style in the specified paint.
func (c *Canvas) DrawRect(rect Rect, paint *Paint) {
	// C.sk_canvas_draw_rect(c.ptr, (*C.sk_rect_t)(unsafe.Pointer(rect)), (*C.sk_paint_t)(paint.ptr))
	C.sk_canvas_draw_rect(c.ptr, rect.cPointer(), paint.ptr)
}

// DrawPath draws a path using the specified paint. The path will be
// filled or stroked based on the style in the paint.
func (c *Canvas) DrawPath(path *Path, paint *Paint) {
	// C.sk_canvas_draw_path(c.ptr, (*C.sk_path_t)(path.ptr), (*C.sk_paint_t)(paint.ptr))
	C.sk_canvas_draw_path(c.ptr, path.ptr, paint.ptr)
}

// A Paint holds the information about how to draw. The information
// includes color, SkTypeface, text size, stroke width, SkShader and so on.
type Paint struct {
	ptr *C.sk_paint_t
}

// A PaintConfig specifies the options of a Paint.
type PaintConfig struct {
	Color       Color
	StrokeWidth float32
	Stroke      bool
	AntiAlias   bool
	// Todo: more properties
}

// NewPaint creates a default Paint with the following option values:
//   antialias : false
//   stroke : false
//   stroke width : 0.0f (hairline)
//   stroke miter : 4.0f
//   stroke cap : BUTT_SK_STROKE_CAP
//   stroke join : MITER_SK_STROKE_JOIN
//   color : opaque black
//   shader : NULL
//   maskfilter : NULL
//   xfermode_mode : SRCOVER_SK_XFERMODE_MODE
//
// A false stroke option means to fill in drawing.
func NewPaint() *Paint {
	ret := &Paint{ptr: C.sk_paint_new()}
	// Todo: use an alternative way to release ref.
	runtime.SetFinalizer(ret, func(p *Paint) {
		C.sk_paint_delete(p.ptr)
	})
	return ret
}

// NewPaint creates a Paint with some common options.
func NewCommonPaint(c Color, stroke bool, strokeWidth float32, antiAlias bool) *Paint {
	ret := NewPaint()
	ret.SetColor(c)
	ret.SetStroke(stroke)
	ret.SetStrokeWidth(strokeWidth)
	ret.SetAntiAlias(antiAlias)
	return ret
}

// NewCustomPaint creates a Paint with full custom options support.
func NewCustomPaint(c *PaintConfig) *Paint {
	ret := NewCommonPaint(c.Color, c.Stroke, c.StrokeWidth, c.AntiAlias)
	return ret
}

// SetColor changes the color value of a Paint.
func (p *Paint) SetColor(color Color) {
	C.sk_paint_set_color(p.ptr, C.sk_color_t(color))
}

// SetAntiAlias changes the antialias flag of a Paint.
// True to enable antialiasing, false to disable it.
func (p *Paint) SetAntiAlias(antiAlias bool) {
	C.sk_paint_set_antialias(p.ptr, C._Bool(antiAlias))
}

// SetStroke changes the stroke flag of a Paint.
// True to enable stroking rather than filling.
func (p *Paint) SetStroke(val bool) {
	C.sk_paint_set_stroke(p.ptr, C._Bool(val))
}

// SetStrokeWidth changes the stroke width value of a Paint.
// A value of 0 strokes in hairline mode (always draw 1-pixel wide,
// regardless of canvas transforms).
func (p *Paint) SetStrokeWidth(width float32) {
	C.sk_paint_set_stroke_width(p.ptr, C.float(width))
}

// A Path encapsulates compound (multiple contour) geometric paths
// consisting of straight line segments, quadratic curves, and cubic curves.
//
// SkPath contain geometry. SkPath may be empty, or contain one or more verbs that
// outline a figure. SkPath always starts with a move verb to a Cartesian coordinate,
// and may be followed by additional verbs that add lines or curves.
// Adding a close verb makes the geometry into a continuous loop, a closed contour.
// SkPath may contain any number of contours, each beginning with a move verb.
//
// SkPath contours may contain only a move verb, or may also contain lines,
// quadratic beziers, conics, and cubic beziers. SkPath contours may be open or closed.
// Whether or not a contour is closed is not important for filling (but for stroking).
//
// When used to draw a filled area, SkPath describes whether the fill is inside or
// outside the geometry. SkPath also describes the winding rule used to fill
// overlapping contours.
//
// Internally, SkPath lazily computes metrics likes bounds and convexity. Call
// SkPath::updateBoundsCache to make SkPath thread safe.
type Path struct {
	ptr *C.sk_path_t
}

// NewPath creates an empty path.
func NewPath() *Path {
	ret := &Path{ptr: C.sk_path_new()}
	// Todo: use an alternative way to release ref.
	runtime.SetFinalizer(ret, func(p *Path) {
		C.sk_path_delete(p.ptr)
	})
	return ret
}

// MoveTo sets the beginning of the next contour to the point (x, y).
func (p *Path) MoveTo(x, y float32) {
	C.sk_path_move_to(p.ptr, C.float(x), C.float(y))
}

// LineTo appends a line from the last point to the specified point (x,y).
// If no sk_path_move_to() call has been made for this contour, the first
// point is automatically set to (0,0).
func (p *Path) LineTo(x, y float32) {
	C.sk_path_line_to(p.ptr, C.float(x), C.float(y))
}

// Todo: support Add***Arc, but not ****ArcTo
// * ArcOfCircleTo
// * ArcOfOvalTo
// * AddCircle
// * AddOval
// * AddRect
// * AddRoundRect
//
// 1. ContourTo
// 2. Contour
// 3. AddContour
// 4. AddContourCCW
// *. .c .h prepares all (stable c++ ones), but some are not exposed.
// a. Path.Action, Path.ActionAs
//
// OvalArcTo appends an oval arc
//func (p *Path) OvalArcTo(rect Rect, radius float32) {
//	C.sk_path_arc_to(p.ptr, C.float(x1), C.float(y1), C.float(x2), C.float(y2), C.float(radius))
//}

// Todo: Need rethink
//func (p *Path) CircleArc(x, y, r, sAngle, eAngle float32, counterclockwise bool) {
//	oval := NewRect(x-r, y-r, x+r, y+r)
//	if counterclockwise {
//		sAngle, eAngle = eAngle, sAngle
//	}
//	sweepAngle := eAngle - sAngle
//	C.sk_path_add_arc(p.ptr, oval.cPointer(), C.float(sAngle)*radian2degrees, C.float(sweepAngle)*radian2degrees)
//}

// QuadTo
func (p *Path) QuadTo(x0, y0, x1, y1 float32) {
	C.sk_path_quad_to(p.ptr, C.float(x0), C.float(y0), C.float(x1), C.float(y1))
}

// ConicTo
func (p *Path) ConicTo(x0, y0, x1, y1, w float32) {
	C.sk_path_conic_to(p.ptr, C.float(x0), C.float(y0), C.float(x1), C.float(y1), C.float(w))
}

// CubicTo
func (p *Path) CubicTo(x0, y0, x1, y1, x2, y2 float32) {
	C.sk_path_cubic_to(p.ptr, C.float(x0), C.float(y0), C.float(x1), C.float(y1), C.float(x2), C.float(y2))
}

// Close closes the current contour in the Path p by connecting
// the first and last point with line. forming a continuous loop.
// Close has no effect if the Path p is empty or the current contour is closed.
func (p *Path) Close() {
	C.sk_path_close(p.ptr)
}

// GetDefaultColortype returns RGBA_8888_COLORTYPE.
func GetDefaultColortype() ColorType {
	return ColorType(C.sk_colortype_get_default_8888())
	//return RGBA_8888_COLORTYPE
}

func (r *Rect) cPointer() *C.sk_rect_t {
	return (*C.sk_rect_t)(unsafe.Pointer(r))
}

func (i *ImageInfo) cPointer() *C.sk_imageinfo_t {
	return (*C.sk_imageinfo_t)(unsafe.Pointer(i))
}









