/*
 * Copyright 2019 Tapir Liu.
 *
 * Copyright 2014 Google Inc.
 *
 * Use of this source code is governed by a BSD-style license that can be
 * found in the LICENSE file.
 */

// EXPERIMENTAL EXPERIMENTAL EXPERIMENTAL EXPERIMENTAL
// DO NOT USE -- FOR INTERNAL TESTING ONLY

#ifndef sk_path_DEFINED
#define sk_path_DEFINED

#include "sk_types.h"

SK_C_PLUS_PLUS_BEGIN_GUARD


/**
    Create a new, empty path.
*/
SK_API sk_path_t* sk_path_new(void);

/**
    Release the memory used by a sk_path_t.
*/
SK_API void sk_path_delete(sk_path_t*);

/**
    Clone a new path.
*/
SK_API sk_path_t* sk_path_clone(const sk_path_t*);

/**
    Copy a new path.
*/
SK_API void sk_path_copy(sk_path_t* dest, const sk_path_t* src);

/*
    Check whether or not two paths are equal.
*/
SK_API bool sk_path_equal(const sk_path_t* a, const sk_path_t* b);

/*
    Check whether or not a path is interpolatable.
*/
SK_API bool sk_path_is_interpolatable(const sk_path_t* cpath, const sk_path_t* compare);

/*
    Interpolate two paths.
*/
SK_API bool sk_path_interpolate(const sk_path_t* cpath, const sk_path_t* ending, float weight, sk_path_t* out);

/*
    Get the fill type of a path.
*/
SK_API sk_path_fill_type_t sk_path_get_fill_type(const sk_path_t* cpath);

/*
    Set the fill type of a path
*/
SK_API void sk_path_set_fill_type(sk_path_t* cpath, sk_path_fill_type_t ft);

/*
    Returns if FillType describes area outside SkPath geometry.
    The inverse fill area extends indefinitely.
*/
// use the static function instead.
//SK_API bool sk_path_is_inverse_fill_type(const sk_path_t* cpath);

/*
    Replaces FillType with its inverse. The inverse of FillType
    describes the area unmodified by the original FillType.
*/
SK_API void sk_path_toggle_inverse_fill_type(sk_path_t* cpath);

/*
    Computes convexity if required, and returns stored value.
    Convexity is computed if stored value is Unknown_Convexity,
    or if path has been altered since convexity was computed or set.
*/
SK_API sk_path_convexity_t sk_path_get_convexity(const sk_path_t* cpath);

/*
    Returns last computed convexity, or Unknown_Convexity if
    path has been altered since convexity was computed or set.
*/
SK_API sk_path_convexity_t sk_path_get_convexity_or_unknown(const sk_path_t* cpath);

/*
    Stores convexity so that it is later returned by
    get_convexity() or get_convexity_or_unknown().
    convexity may differ from get_convexity(),
    although setting an incorrect value may
    cause incorrect or inefficient drawing.
*/
SK_API void sk_path_set_convexity(sk_path_t* cpath, sk_path_convexity_t convexity);

/*
    Computes convexity if required, and returns true
    if value is Convex_Convexity.
    If set_convexity() was called with Convex_Convexity or Concave_Convexity,
    and the path has not been altered, convexity is not recomputed.
*/
SK_API bool sk_path_is_convex(const sk_path_t* cpath);

/*
    Returns true if this path is recognized as an oval or circle.
    bounds receives bounds of oval.
    bounds is unmodified if oval is not found.
*/
SK_API bool sk_path_is_oval(const sk_path_t* cpath, sk_rect_t* bounds);

/*
    Returns true if path is representable as sk_rrect_t.
    Returns false if path is representable as oval, circle, or rect.
    rrect receives bounds of sk_rrect_t.
    rrect is unmodified if sk_rrect_t is not found.
*/
SK_API bool sk_path_is_rrect(const sk_path_t* cpath, sk_rrect_t* rrect);

/*
    Sets SkPath to its initial state.
    Removes verb array, SkPoint array, and weights, and sets FillType to Winding_FillType.
    Internal storage associated with SkPath is released.
*/
SK_API sk_path_t* sk_path_reset(sk_path_t* cpath);

/*
    Sets SkPath to its initial state, preserving internal storage.
    Removes verb array, SkPoint array, and weights, and sets FillType to Winding_FillType.
    Internal storage associated with SkPath is retained.
*/
SK_API sk_path_t* sk_path_rewind(sk_path_t* cpath);



SK_API bool sk_path_is_empty(const sk_path_t* cpath);
SK_API bool sk_path_is_last_contour_closed(const sk_path_t* cpath);
SK_API bool sk_path_is_finite(const sk_path_t* cpath);
SK_API bool sk_path_is_volatile(const sk_path_t* cpath);
SK_API void sk_path_set_volatile(sk_path_t* cpath, bool isVolatile);
SK_API bool sk_path_is_line_degenerate(const sk_point_t* p1, const sk_point_t* p2, bool exact);
SK_API bool sk_path_is_quad_degenerate(const sk_point_t* p1, const sk_point_t* p2, const sk_point_t* p3, bool exact);
SK_API bool sk_path_is_cubic_degenerate(const sk_point_t* p1, const sk_point_t* p2, const sk_point_t* p3, const sk_point_t* p4, bool exact);
SK_API bool sk_path_is_line(const sk_path_t* cpath, sk_path_t* lines);



SK_API int sk_path_count_points(const sk_path_t* cpath);
SK_API sk_point_t sk_path_get_point(const sk_path_t* cpath, int index);
SK_API int sk_path_get_points(const sk_path_t* cpath, sk_point_t cpoints[], int max);
SK_API int sk_path_count_verbs(const sk_path_t* cpath);
SK_API int sk_path_get_verbs(const sk_path_t* cpath, uint8_t verbs[], int max);
SK_API void sk_path_swap(sk_path_t* cpath, sk_path_t* cother);

/**

*/
SK_API sk_rect_t sk_path_get_bounds(const sk_path_t*);
SK_API void sk_path_update_bounds_cache(const sk_path_t* cpath);
SK_API sk_rect_t sk_path_compute_tight_bounds(const sk_path_t* cpath);
SK_API bool sk_path_conservatively_contains_rect(const sk_path_t* cpath, const sk_rect_t* crect);

SK_API void sk_path_inc_reserve(sk_path_t* cpath, int extraPtCount);
//SK_API void shrink_to_fit(sk_path_t* cpath);

/**
    Set the beginning of the next contour to the point (x,y).
*/
SK_API sk_path_t* sk_path_move_to(sk_path_t*, float x, float y);
SK_API sk_path_t* sk_path_move_to_1(sk_path_t*, const sk_point_t* p);
SK_API sk_path_t* sk_path_r_move_to(sk_path_t*, float dx, float dy);

/**
    Add a line from the last point to the specified point (x,y). If no
    sk_path_move_to() call has been made for this contour, the first
    point is automatically set to (0,0).
*/
SK_API sk_path_t* sk_path_line_to(sk_path_t*, float x, float y);
SK_API sk_path_t* sk_path_line_to_1(sk_path_t*, const sk_point_t* p);
SK_API sk_path_t* sk_path_r_line_to(sk_path_t*, float dx, float dy);

/**
    Add a quadratic bezier from the last point, approaching control
    point (x0,y0), and ending at (x1,y1). If no sk_path_move_to() call
    has been made for this contour, the first point is automatically
    set to (0,0).
*/
SK_API sk_path_t* sk_path_quad_to(sk_path_t*, float x0, float y0, float x1, float y1);
SK_API sk_path_t* sk_path_quad_to_1(sk_path_t*, const sk_point_t* p1, const sk_point_t* p2);
SK_API sk_path_t* sk_path_r_quad_to(sk_path_t*, float dx1, float dy1, float dx2, float dy2);

/**
    Add a conic curve from the last point, approaching control point
    (x0,y01), and ending at (x1,y1) with weight w.  If no
    sk_path_move_to() call has been made for this contour, the first
    point is automatically set to (0,0).
*/
SK_API sk_path_t* sk_path_conic_to(sk_path_t*, float x0, float y0, float x1, float y1, float w);
SK_API sk_path_t* sk_path_conic_to_1(sk_path_t*, const sk_point_t* p1, const sk_point_t* p2, float w);
SK_API sk_path_t* sk_path_r_conic_to(sk_path_t*, float dx1, float dy1, float dx2, float dy2, float w);

/**
    Add a cubic bezier from the last point, approaching control points
    (x0,y0) and (x1,y1), and ending at (x2,y2). If no
    sk_path_move_to() call has been made for this contour, the first
    point is automatically set to (0,0).
*/
SK_API sk_path_t* sk_path_cubic_to(sk_path_t*, float x0, float y0, float x1, float y1, float x2, float y2);
SK_API sk_path_t* sk_path_cubic_to_1(sk_path_t*, const sk_point_t* p1, const sk_point_t* p2, const sk_point_t* p3);
SK_API sk_path_t* sk_path_r_cubic_to(sk_path_t*, float dx1, float dy1, float dx2, float dy2, float dx3, float dy3);

/**
    Appends arc to SkPath. Arc added is part of ellipse
    bounded by oval, from startAngle through sweepAngle. Both startAngle and
    sweepAngle are measured in degrees, where zero degrees is aligned with the
    positive x-axis, and positive sweeps extends arc clockwise.
*/
SK_API sk_path_t* sk_path_arc_to_a(sk_path_t* cpath, const sk_rect_t* oval, float startAngle, float sweepAngle, bool forceMoveTo);

/**
    Appends arc to SkPath, after appending line if needed. Arc is implemented by conic
    weighted to describe part of circle. Arc is contained by tangent from
    last SkPath point to (x1, y1), and tangent from (x1, y1) to (x2, y2). Arc
    is part of circle sized to radius, positioned so it touches both tangent lines.
*/
SK_API sk_path_t* sk_path_arc_to_b(sk_path_t* cpath, float x1, float y1, float x2, float y2, float radius);
SK_API sk_path_t* sk_path_arc_to_b_1(sk_path_t* cpath, const sk_point_t p1, const sk_point_t p2, float radius);
SK_API sk_path_t* sk_path_arc_to_c(sk_path_t* cpath, float rx, float ry, float xAxisRotate, sk_path_arc_size_t largeArc, sk_path_direction_t sweep, float x, float y);
SK_API sk_path_t* sk_path_arc_to_c_1(sk_path_t* cpath, const sk_point_t r, float xAxisRotate, sk_path_arc_size_t largeArc, sk_path_direction_t sweep, const sk_point_t xy);
SK_API sk_path_t* sk_path_r_arc_to(sk_path_t* cpath, float rx, float ry, float xAxisRotate, sk_path_arc_size_t largeArc, sk_path_direction_t sweep, float dx, float dy);

// todo: why not "const sk_point_t*" for above 4 "const sk_point_t".


/**
    Close the current contour. If the current point is not equal to the
    first point of the contour, a line segment is automatically added.
*/
SK_API sk_path_t* sk_path_close(sk_path_t*);

SK_API bool sk_path_is_inverse_fill_type(sk_path_fill_type_t fill);
SK_API sk_path_fill_type_t sk_path_convert_to_non_inverse_fill_type(sk_path_fill_type_t fill);
SK_API int sk_path_convert_conic_to_quads(const sk_point_t* p0, const sk_point_t* p1, const sk_point_t* p2, float w, sk_point_t pts[], int pow2);

SK_API bool sk_path_is_rect(const sk_path_t* cpath, sk_rect_t* rect, bool* isClosed, sk_path_direction_t* direction);
SK_API bool sk_path_is_nested_fill_rects(const sk_path_t* cpath, sk_rect_t rect[2], sk_path_direction_t dirs[2]);



/**
    Add a closed rectangle contour to the path.
*/
SK_API sk_path_t* sk_path_add_rect(sk_path_t* cpath, const sk_rect_t* rect, sk_path_direction_t dir, unsigned start);
SK_API sk_path_t* sk_path_add_rect_1(sk_path_t* cpath, const sk_rect_t* rect, sk_path_direction_t dir);
SK_API sk_path_t* sk_path_add_rect_2(sk_path_t* cpath, float left, float top, float right, float bottom, sk_path_direction_t dir);



/**
    Add a closed oval contour to the path
*/
SK_API sk_path_t* sk_path_add_oval(sk_path_t* cpath, const sk_rect_t* oval, sk_path_direction_t dir, unsigned start);
SK_API sk_path_t* sk_path_add_oval_1(sk_path_t* cpath, const sk_rect_t*oval, sk_path_direction_t dir);

SK_API sk_path_t* sk_path_add_circle(sk_path_t* cpath, float x, float y, float radius, sk_path_direction_t dir);

/**
    Add an oval arc to the path.
*/
SK_API sk_path_t* sk_path_add_arc(sk_path_t* cpath, const sk_rect_t* oval, float startAngle, float sweepAngle);
SK_API sk_path_t* sk_path_add_round_rect(sk_path_t* cpath, const sk_rect_t* rect, const float radii[], sk_path_direction_t dir);
SK_API sk_path_t* sk_path_add_round_rect_1(sk_path_t* cpath, const sk_rect_t* rect, float rx, float ry, sk_path_direction_t dir);
SK_API sk_path_t* sk_path_add_rrect(sk_path_t* cpath, const sk_rrect_t* rrect, sk_path_direction_t dir, unsigned start);
SK_API sk_path_t* sk_path_add_rrect_1(sk_path_t* cpath, const sk_rrect_t* rrect, sk_path_direction_t dir);
SK_API sk_path_t* sk_path_add_poly(sk_path_t* cpath, const sk_point_t pts[], int count, bool close);
// SkPath& addPoly(const std::initializer_list<SkPoint>& list, bool close) // todo
SK_API sk_path_t* sk_path_add_path(sk_path_t* cpath, const sk_path_t* src, const sk_matrix_t* matrix, sk_path_add_path_mode_t mode);
SK_API sk_path_t* sk_path_add_path_1(sk_path_t* cpath, const sk_path_t* src, float dx, float dy, sk_path_add_path_mode_t mode);
SK_API sk_path_t* sk_path_add_path_2(sk_path_t* cpath, const sk_path_t* src, sk_path_add_path_mode_t mode);
SK_API sk_path_t* sk_path_reverse_add_path(sk_path_t* cpath, const sk_path_t* src);

SK_API void sk_path_offset_a(const sk_path_t* cpath, float dx, float dy, sk_path_t* dst);
SK_API void sk_path_offset_b(sk_path_t* cpath, float dx, float dy);
SK_API void sk_path_transform_a(const sk_path_t* cpath, const sk_matrix_t* matrix, sk_path_t* dst);
SK_API void sk_path_transform_b(sk_path_t* cpath, const sk_matrix_t* matrix);
SK_API bool sk_path_get_last_pt(const sk_path_t* cpath, sk_point_t* lastPt);
SK_API void sk_path_set_last_pt(sk_path_t* cpath, float x, float y);
SK_API void sk_path_set_last_pt_1(sk_path_t* cpath, const sk_point_t* p);
SK_API uint32_t sk_path_get_segment_masks(const sk_path_t* cpath);
SK_API bool sk_path_contains(const sk_path_t* cpath, float x, float y);
// SK_API void dump(SkWStream* stream, bool forceClose, bool dumpAsHex) // todo
SK_API void sk_path_dump(const sk_path_t* cpath);
SK_API uint32_t sk_path_get_generation_id(const sk_path_t* cpath);
SK_API bool sk_path_is_valid(const sk_path_t* cpath);

// todo: return sk_path_t* -> bool; if (!from_c_path_direction(cdir, &dir)) {...}

// todo: optimize matrix conversions

// modify to_c_path_fill_type, ..., functions

// todo: remove SkXXX from comments.

// todo: matrix

SK_C_PLUS_PLUS_END_GUARD

#endif
