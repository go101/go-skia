//+build ignore

/*
 * Copyright 2015 Google Inc.
 *
 * Use of this source code is governed by a BSD-style license that can be
 * found in the LICENSE file.
 */

package skia

// This file is used to generate 'types.go'
// from the corresponding type definitions in the C API.
// Any C struct for which we would like to generate a
// Go struct with the same memory layout should defined defined here.
// Any enum that is used in Go should also be listed here, together
// with the enum values that we want to use.

/*

Run the following commands to create the "types.go" file.
The purpose of this file is to have Go structs with the same memory
layout as their C counterparts. For enums we want the underlying primitive
types to match.

      cat credit > types.go
      echo -e "\n// Please see file "ctypes.go" on how this file is created.\n" >> types.go
      go tool cgo -godefs ctypes.go >> types.go

Some enum fields in structs need to be fixed by hand.

TODO(stephan): Add tests that allow to detect failure on platforms other
               than Linux and changes in the underlying C types.
*/

/*
#cgo CFLAGS: -Iskia/include/c
#include "skia/include/c/sk_types.h"
*/
import "C"

// todo: change enum constant names, such as
// * CW_SK_PATH_DIRECTION
// * UNKNOWN_SK_COLORTYPE
// * ...

type Color C.sk_color_t

type ColorType C.sk_colortype_t

const (
	UNKNOWN_COLORTYPE   ColorType = C.UNKNOWN_SK_COLORTYPE
	RGBA_8888_COLORTYPE ColorType = C.RGBA_8888_SK_COLORTYPE
	BGRA_8888_COLORTYPE ColorType = C.BGRA_8888_SK_COLORTYPE
	ALPHA_8_COLORTYPE   ColorType = C.ALPHA_8_SK_COLORTYPE
)

type AlphaType C.sk_alphatype_t

const (
	OPAQUE_ALPHATYPE   AlphaType = C.OPAQUE_SK_ALPHATYPE
	PREMUL_ALPHATYPE   AlphaType = C.PREMUL_SK_ALPHATYPE
	UNPREMUL_ALPHATYPE AlphaType = C.UNPREMUL_SK_ALPHATYPE
)

type PixelGeometry C.sk_pixelgeometry_t

const (
	UNKNOWN_SK_PIXELGEOMETRY PixelGeometry = C.UNKNOWN_SK_PIXELGEOMETRY
	RGB_H_SK_PIXELGEOMETRY   PixelGeometry = C.RGB_H_SK_PIXELGEOMETRY
	BGR_H_SK_PIXELGEOMETRY   PixelGeometry = C.BGR_H_SK_PIXELGEOMETRY
	RGB_V_SK_PIXELGEOMETRY   PixelGeometry = C.RGB_V_SK_PIXELGEOMETRY
	BGR_V_SK_PIXELGEOMETRY   PixelGeometry = C.BGR_V_SK_PIXELGEOMETRY
)

type ImageInfo C.sk_imageinfo_t

type SurfaceProps C.sk_surfaceprops_t

type Rect C.sk_rect_t
