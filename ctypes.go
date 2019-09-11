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

Some enum fields in structs need to be fixed by hand. For example,
* ColorType	ColorType
* AlphaType	AlphaType

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

type ColorType C.sk_color_type_t

const (
	UNKNOWN_COLORTYPE   ColorType = C.Unknown_SkColorType
	RGBA_8888_COLORTYPE ColorType = C.RGBA_8888_SkColorType
	BGRA_8888_COLORTYPE ColorType = C.BGRA_8888_SkColorType
	ALPHA_8_COLORTYPE   ColorType = C.Alpha_8_SkColorType
)

type AlphaType C.sk_alpha_type_t

const (
	OPAQUE_ALPHATYPE   AlphaType = C.Opaque_SkAlphaType
	PREMUL_ALPHATYPE   AlphaType = C.Premul_SkAlphaType
	UNPREMUL_ALPHATYPE AlphaType = C.Unpremul_SkAlphaType
)

type PixelGeometry C.sk_pixel_geometry_t

const (
	UNKNOWN_SK_PIXELGEOMETRY PixelGeometry = C.Unknown_SkPixelGeometry
	RGB_H_SK_PIXELGEOMETRY   PixelGeometry = C.RGB_H_SkPixelGeometry
	BGR_H_SK_PIXELGEOMETRY   PixelGeometry = C.BGR_H_SkPixelGeometry
	RGB_V_SK_PIXELGEOMETRY   PixelGeometry = C.RGB_V_SkPixelGeometry
	BGR_V_SK_PIXELGEOMETRY   PixelGeometry = C.BGR_V_SkPixelGeometry
)

type ImageInfo C.sk_imageinfo_t

type SurfaceProps C.sk_surfaceprops_t

type Rect C.sk_rect_t
