// Copyright 2012 The go-gl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gl

// #include "gl.h"
import "C"
import "unsafe"

const (
	GL_2_BYTES                    = C.GL_2_BYTES
	GL_2D                         = C.GL_2D
	GL_3_BYTES                    = C.GL_3_BYTES
	GL_3D                         = C.GL_3D
	GL_3D_COLOR                   = C.GL_3D_COLOR
	GL_3D_COLOR_TEXTURE           = C.GL_3D_COLOR_TEXTURE
	GL_4_BYTES                    = C.GL_4_BYTES
	GL_4D_COLOR_TEXTURE           = C.GL_4D_COLOR_TEXTURE
	ACCUM                         = C.GL_ACCUM
	ACCUM_ALPHA_BITS              = C.GL_ACCUM_ALPHA_BITS
	ACCUM_BLUE_BITS               = C.GL_ACCUM_BLUE_BITS
	ACCUM_BUFFER_BIT              = C.GL_ACCUM_BUFFER_BIT
	ACCUM_CLEAR_VALUE             = C.GL_ACCUM_CLEAR_VALUE
	ACCUM_GREEN_BITS              = C.GL_ACCUM_GREEN_BITS
	ACCUM_RED_BITS                = C.GL_ACCUM_RED_BITS
	ADD                           = C.GL_ADD
	ALL_ATTRIB_BITS               = C.GL_ALL_ATTRIB_BITS
	ALPHA                         = C.GL_ALPHA
	ALPHA12                       = C.GL_ALPHA12
	ALPHA16                       = C.GL_ALPHA16
	ALPHA4                        = C.GL_ALPHA4
	ALPHA8                        = C.GL_ALPHA8
	ALPHA_BIAS                    = C.GL_ALPHA_BIAS
	ALPHA_BITS                    = C.GL_ALPHA_BITS
	ALPHA_SCALE                   = C.GL_ALPHA_SCALE
	ALPHA_TEST                    = C.GL_ALPHA_TEST
	ALPHA_TEST_FUNC               = C.GL_ALPHA_TEST_FUNC
	ALPHA_TEST_REF                = C.GL_ALPHA_TEST_REF
	ALWAYS                        = C.GL_ALWAYS
	AMBIENT                       = C.GL_AMBIENT
	AMBIENT_AND_DIFFUSE           = C.GL_AMBIENT_AND_DIFFUSE
	AND                           = C.GL_AND
	AND_INVERTED                  = C.GL_AND_INVERTED
	AND_REVERSE                   = C.GL_AND_REVERSE
	ATTRIB_STACK_DEPTH            = C.GL_ATTRIB_STACK_DEPTH
	AUTO_NORMAL                   = C.GL_AUTO_NORMAL
	AUX0                          = C.GL_AUX0
	AUX1                          = C.GL_AUX1
	AUX2                          = C.GL_AUX2
	AUX3                          = C.GL_AUX3
	AUX_BUFFERS                   = C.GL_AUX_BUFFERS
	BACK                          = C.GL_BACK
	BACK_LEFT                     = C.GL_BACK_LEFT
	BACK_RIGHT                    = C.GL_BACK_RIGHT
	BITMAP                        = C.GL_BITMAP
	BITMAP_TOKEN                  = C.GL_BITMAP_TOKEN
	BLEND                         = C.GL_BLEND
	BLEND_DST                     = C.GL_BLEND_DST
	BLEND_SRC                     = C.GL_BLEND_SRC
	BLUE                          = C.GL_BLUE
	BLUE_BIAS                     = C.GL_BLUE_BIAS
	BLUE_BITS                     = C.GL_BLUE_BITS
	BLUE_SCALE                    = C.GL_BLUE_SCALE
	BYTE                          = C.GL_BYTE
	C3F_V3F                       = C.GL_C3F_V3F
	C4F_N3F_V3F                   = C.GL_C4F_N3F_V3F
	C4UB_V2F                      = C.GL_C4UB_V2F
	C4UB_V3F                      = C.GL_C4UB_V3F
	CCW                           = C.GL_CCW
	CLAMP                         = C.GL_CLAMP
	CLEAR                         = C.GL_CLEAR
	CLIENT_ALL_ATTRIB_BITS        = C.GL_CLIENT_ALL_ATTRIB_BITS
	CLIENT_ATTRIB_STACK_DEPTH     = C.GL_CLIENT_ATTRIB_STACK_DEPTH
	CLIENT_PIXEL_STORE_BIT        = C.GL_CLIENT_PIXEL_STORE_BIT
	CLIENT_VERTEX_ARRAY_BIT       = C.GL_CLIENT_VERTEX_ARRAY_BIT
	CLIP_PLANE0                   = C.GL_CLIP_PLANE0
	CLIP_PLANE1                   = C.GL_CLIP_PLANE1
	CLIP_PLANE2                   = C.GL_CLIP_PLANE2
	CLIP_PLANE3                   = C.GL_CLIP_PLANE3
	CLIP_PLANE4                   = C.GL_CLIP_PLANE4
	CLIP_PLANE5                   = C.GL_CLIP_PLANE5
	COEFF                         = C.GL_COEFF
	COLOR                         = C.GL_COLOR
	COLOR_ARRAY                   = C.GL_COLOR_ARRAY
	COLOR_ARRAY_POINTER           = C.GL_COLOR_ARRAY_POINTER
	COLOR_ARRAY_SIZE              = C.GL_COLOR_ARRAY_SIZE
	COLOR_ARRAY_STRIDE            = C.GL_COLOR_ARRAY_STRIDE
	COLOR_ARRAY_TYPE              = C.GL_COLOR_ARRAY_TYPE
	COLOR_BUFFER_BIT              = C.GL_COLOR_BUFFER_BIT
	COLOR_CLEAR_VALUE             = C.GL_COLOR_CLEAR_VALUE
	COLOR_INDEX                   = C.GL_COLOR_INDEX
	COLOR_INDEX12_EXT             = C.GL_COLOR_INDEX12_EXT
	COLOR_INDEX16_EXT             = C.GL_COLOR_INDEX16_EXT
	COLOR_INDEX1_EXT              = C.GL_COLOR_INDEX1_EXT
	COLOR_INDEX2_EXT              = C.GL_COLOR_INDEX2_EXT
	COLOR_INDEX4_EXT              = C.GL_COLOR_INDEX4_EXT
	COLOR_INDEX8_EXT              = C.GL_COLOR_INDEX8_EXT
	COLOR_INDEXES                 = C.GL_COLOR_INDEXES
	COLOR_LOGIC_OP                = C.GL_COLOR_LOGIC_OP
	COLOR_MATERIAL                = C.GL_COLOR_MATERIAL
	COLOR_MATERIAL_FACE           = C.GL_COLOR_MATERIAL_FACE
	COLOR_MATERIAL_PARAMETER      = C.GL_COLOR_MATERIAL_PARAMETER
	COLOR_WRITEMASK               = C.GL_COLOR_WRITEMASK
	COMPILE                       = C.GL_COMPILE
	COMPILE_AND_EXECUTE           = C.GL_COMPILE_AND_EXECUTE
	CONSTANT_ATTENUATION          = C.GL_CONSTANT_ATTENUATION
	COPY                          = C.GL_COPY
	COPY_INVERTED                 = C.GL_COPY_INVERTED
	COPY_PIXEL_TOKEN              = C.GL_COPY_PIXEL_TOKEN
	CULL_FACE                     = C.GL_CULL_FACE
	CULL_FACE_MODE                = C.GL_CULL_FACE_MODE
	CURRENT_BIT                   = C.GL_CURRENT_BIT
	CURRENT_COLOR                 = C.GL_CURRENT_COLOR
	CURRENT_INDEX                 = C.GL_CURRENT_INDEX
	CURRENT_NORMAL                = C.GL_CURRENT_NORMAL
	CURRENT_RASTER_COLOR          = C.GL_CURRENT_RASTER_COLOR
	CURRENT_RASTER_DISTANCE       = C.GL_CURRENT_RASTER_DISTANCE
	CURRENT_RASTER_INDEX          = C.GL_CURRENT_RASTER_INDEX
	CURRENT_RASTER_POSITION       = C.GL_CURRENT_RASTER_POSITION
	CURRENT_RASTER_POSITION_VALID = C.GL_CURRENT_RASTER_POSITION_VALID
	CURRENT_RASTER_TEXTURE_COORDS = C.GL_CURRENT_RASTER_TEXTURE_COORDS
	CURRENT_TEXTURE_COORDS        = C.GL_CURRENT_TEXTURE_COORDS
	CW                            = C.GL_CW
	DECAL                         = C.GL_DECAL
	DECR                          = C.GL_DECR
	DEPTH                         = C.GL_DEPTH
	DEPTH_BIAS                    = C.GL_DEPTH_BIAS
	DEPTH_BITS                    = C.GL_DEPTH_BITS
	DEPTH_BUFFER_BIT              = C.GL_DEPTH_BUFFER_BIT
	DEPTH_CLEAR_VALUE             = C.GL_DEPTH_CLEAR_VALUE
	DEPTH_COMPONENT               = C.GL_DEPTH_COMPONENT
	DEPTH_FUNC                    = C.GL_DEPTH_FUNC
	DEPTH_RANGE                   = C.GL_DEPTH_RANGE
	DEPTH_SCALE                   = C.GL_DEPTH_SCALE
	DEPTH_TEST                    = C.GL_DEPTH_TEST
	DEPTH_WRITEMASK               = C.GL_DEPTH_WRITEMASK
	DIFFUSE                       = C.GL_DIFFUSE
	DITHER                        = C.GL_DITHER
	DOMAIN                        = C.GL_DOMAIN
	DONT_CARE                     = C.GL_DONT_CARE
	DOUBLE                        = C.GL_DOUBLE
	DOUBLEBUFFER                  = C.GL_DOUBLEBUFFER
	DRAW_BUFFER                   = C.GL_DRAW_BUFFER
	DRAW_PIXEL_TOKEN              = C.GL_DRAW_PIXEL_TOKEN
	DST_ALPHA                     = C.GL_DST_ALPHA
	DST_COLOR                     = C.GL_DST_COLOR
	EDGE_FLAG                     = C.GL_EDGE_FLAG
	EDGE_FLAG_ARRAY               = C.GL_EDGE_FLAG_ARRAY
	EDGE_FLAG_ARRAY_POINTER       = C.GL_EDGE_FLAG_ARRAY_POINTER
	EDGE_FLAG_ARRAY_STRIDE        = C.GL_EDGE_FLAG_ARRAY_STRIDE
	EMISSION                      = C.GL_EMISSION
	ENABLE_BIT                    = C.GL_ENABLE_BIT
	EQUAL                         = C.GL_EQUAL
	EQUIV                         = C.GL_EQUIV
	EVAL_BIT                      = C.GL_EVAL_BIT
	EXP                           = C.GL_EXP
	EXP2                          = C.GL_EXP2
	EXTENSIONS                    = C.GL_EXTENSIONS
	EYE_LINEAR                    = C.GL_EYE_LINEAR
	EYE_PLANE                     = C.GL_EYE_PLANE
	FALSE                         = C.GL_FALSE
	FASTEST                       = C.GL_FASTEST
	FEEDBACK                      = C.GL_FEEDBACK
	FEEDBACK_BUFFER_POINTER       = C.GL_FEEDBACK_BUFFER_POINTER
	FEEDBACK_BUFFER_SIZE          = C.GL_FEEDBACK_BUFFER_SIZE
	FEEDBACK_BUFFER_TYPE          = C.GL_FEEDBACK_BUFFER_TYPE
	FILL                          = C.GL_FILL
	FLAT                          = C.GL_FLAT
	FLOAT                         = C.GL_FLOAT
	FOG                           = C.GL_FOG
	FOG_BIT                       = C.GL_FOG_BIT
	FOG_COLOR                     = C.GL_FOG_COLOR
	FOG_DENSITY                   = C.GL_FOG_DENSITY
	FOG_END                       = C.GL_FOG_END
	FOG_HINT                      = C.GL_FOG_HINT
	FOG_INDEX                     = C.GL_FOG_INDEX
	FOG_MODE                      = C.GL_FOG_MODE
	FOG_START                     = C.GL_FOG_START
	FRONT                         = C.GL_FRONT
	FRONT_AND_BACK                = C.GL_FRONT_AND_BACK
	FRONT_FACE                    = C.GL_FRONT_FACE
	FRONT_LEFT                    = C.GL_FRONT_LEFT
	FRONT_RIGHT                   = C.GL_FRONT_RIGHT
	GEQUAL                        = C.GL_GEQUAL
	GREATER                       = C.GL_GREATER
	GREEN                         = C.GL_GREEN
	GREEN_BIAS                    = C.GL_GREEN_BIAS
	GREEN_BITS                    = C.GL_GREEN_BITS
	GREEN_SCALE                   = C.GL_GREEN_SCALE
	HINT_BIT                      = C.GL_HINT_BIT
	INCR                          = C.GL_INCR
	INDEX_ARRAY                   = C.GL_INDEX_ARRAY
	INDEX_ARRAY_POINTER           = C.GL_INDEX_ARRAY_POINTER
	INDEX_ARRAY_STRIDE            = C.GL_INDEX_ARRAY_STRIDE
	INDEX_ARRAY_TYPE              = C.GL_INDEX_ARRAY_TYPE
	INDEX_BITS                    = C.GL_INDEX_BITS
	INDEX_CLEAR_VALUE             = C.GL_INDEX_CLEAR_VALUE
	INDEX_LOGIC_OP                = C.GL_INDEX_LOGIC_OP
	INDEX_MODE                    = C.GL_INDEX_MODE
	INDEX_OFFSET                  = C.GL_INDEX_OFFSET
	INDEX_SHIFT                   = C.GL_INDEX_SHIFT
	INDEX_WRITEMASK               = C.GL_INDEX_WRITEMASK
	INT                           = C.GL_INT
	INTENSITY                     = C.GL_INTENSITY
	INTENSITY12                   = C.GL_INTENSITY12
	INTENSITY16                   = C.GL_INTENSITY16
	INTENSITY4                    = C.GL_INTENSITY4
	INTENSITY8                    = C.GL_INTENSITY8
	INVALID_ENUM                  = C.GL_INVALID_ENUM
	INVALID_OPERATION             = C.GL_INVALID_OPERATION
	INVALID_VALUE                 = C.GL_INVALID_VALUE
	INVERT                        = C.GL_INVERT
	KEEP                          = C.GL_KEEP
	LEFT                          = C.GL_LEFT
	LEQUAL                        = C.GL_LEQUAL
	LESS                          = C.GL_LESS
	LIGHT0                        = C.GL_LIGHT0
	LIGHT1                        = C.GL_LIGHT1
	LIGHT2                        = C.GL_LIGHT2
	LIGHT3                        = C.GL_LIGHT3
	LIGHT4                        = C.GL_LIGHT4
	LIGHT5                        = C.GL_LIGHT5
	LIGHT6                        = C.GL_LIGHT6
	LIGHT7                        = C.GL_LIGHT7
	LIGHT_MODEL_AMBIENT           = C.GL_LIGHT_MODEL_AMBIENT
	LIGHT_MODEL_LOCAL_VIEWER      = C.GL_LIGHT_MODEL_LOCAL_VIEWER
	LIGHT_MODEL_TWO_SIDE          = C.GL_LIGHT_MODEL_TWO_SIDE
	LIGHTING                      = C.GL_LIGHTING
	LIGHTING_BIT                  = C.GL_LIGHTING_BIT
	LINE                          = C.GL_LINE
	LINE_BIT                      = C.GL_LINE_BIT
	LINE_LOOP                     = C.GL_LINE_LOOP
	LINE_RESET_TOKEN              = C.GL_LINE_RESET_TOKEN
	LINE_SMOOTH                   = C.GL_LINE_SMOOTH
	LINE_SMOOTH_HINT              = C.GL_LINE_SMOOTH_HINT
	LINE_STIPPLE                  = C.GL_LINE_STIPPLE
	LINE_STIPPLE_PATTERN          = C.GL_LINE_STIPPLE_PATTERN
	LINE_STIPPLE_REPEAT           = C.GL_LINE_STIPPLE_REPEAT
	LINE_STRIP                    = C.GL_LINE_STRIP
	LINE_TOKEN                    = C.GL_LINE_TOKEN
	LINE_WIDTH                    = C.GL_LINE_WIDTH
	LINE_WIDTH_GRANULARITY        = C.GL_LINE_WIDTH_GRANULARITY
	LINE_WIDTH_RANGE              = C.GL_LINE_WIDTH_RANGE
	LINEAR                        = C.GL_LINEAR
	LINEAR_ATTENUATION            = C.GL_LINEAR_ATTENUATION
	LINEAR_MIPMAP_LINEAR          = C.GL_LINEAR_MIPMAP_LINEAR
	LINEAR_MIPMAP_NEAREST         = C.GL_LINEAR_MIPMAP_NEAREST
	LINES                         = C.GL_LINES
	LIST_BASE                     = C.GL_LIST_BASE
	LIST_BIT                      = C.GL_LIST_BIT
	LIST_INDEX                    = C.GL_LIST_INDEX
	LIST_MODE                     = C.GL_LIST_MODE
	LOAD                          = C.GL_LOAD
	LOGIC_OP                      = C.GL_LOGIC_OP
	LOGIC_OP_MODE                 = C.GL_LOGIC_OP_MODE
	LUMINANCE                     = C.GL_LUMINANCE
	LUMINANCE12                   = C.GL_LUMINANCE12
	LUMINANCE12_ALPHA12           = C.GL_LUMINANCE12_ALPHA12
	LUMINANCE12_ALPHA4            = C.GL_LUMINANCE12_ALPHA4
	LUMINANCE16                   = C.GL_LUMINANCE16
	LUMINANCE16_ALPHA16           = C.GL_LUMINANCE16_ALPHA16
	LUMINANCE4                    = C.GL_LUMINANCE4
	LUMINANCE4_ALPHA4             = C.GL_LUMINANCE4_ALPHA4
	LUMINANCE6_ALPHA2             = C.GL_LUMINANCE6_ALPHA2
	LUMINANCE8                    = C.GL_LUMINANCE8
	LUMINANCE8_ALPHA8             = C.GL_LUMINANCE8_ALPHA8
	LUMINANCE_ALPHA               = C.GL_LUMINANCE_ALPHA
	MAP1_COLOR_4                  = C.GL_MAP1_COLOR_4
	MAP1_GRID_DOMAIN              = C.GL_MAP1_GRID_DOMAIN
	MAP1_GRID_SEGMENTS            = C.GL_MAP1_GRID_SEGMENTS
	MAP1_INDEX                    = C.GL_MAP1_INDEX
	MAP1_NORMAL                   = C.GL_MAP1_NORMAL
	MAP1_TEXTURE_COORD_1          = C.GL_MAP1_TEXTURE_COORD_1
	MAP1_TEXTURE_COORD_2          = C.GL_MAP1_TEXTURE_COORD_2
	MAP1_TEXTURE_COORD_3          = C.GL_MAP1_TEXTURE_COORD_3
	MAP1_TEXTURE_COORD_4          = C.GL_MAP1_TEXTURE_COORD_4
	MAP1_VERTEX_3                 = C.GL_MAP1_VERTEX_3
	MAP1_VERTEX_4                 = C.GL_MAP1_VERTEX_4
	MAP2_COLOR_4                  = C.GL_MAP2_COLOR_4
	MAP2_GRID_DOMAIN              = C.GL_MAP2_GRID_DOMAIN
	MAP2_GRID_SEGMENTS            = C.GL_MAP2_GRID_SEGMENTS
	MAP2_INDEX                    = C.GL_MAP2_INDEX
	MAP2_NORMAL                   = C.GL_MAP2_NORMAL
	MAP2_TEXTURE_COORD_1          = C.GL_MAP2_TEXTURE_COORD_1
	MAP2_TEXTURE_COORD_2          = C.GL_MAP2_TEXTURE_COORD_2
	MAP2_TEXTURE_COORD_3          = C.GL_MAP2_TEXTURE_COORD_3
	MAP2_TEXTURE_COORD_4          = C.GL_MAP2_TEXTURE_COORD_4
	MAP2_VERTEX_3                 = C.GL_MAP2_VERTEX_3
	MAP2_VERTEX_4                 = C.GL_MAP2_VERTEX_4
	MAP_COLOR                     = C.GL_MAP_COLOR
	MAP_STENCIL                   = C.GL_MAP_STENCIL
	MATRIX_MODE                   = C.GL_MATRIX_MODE
	MAX_ATTRIB_STACK_DEPTH        = C.GL_MAX_ATTRIB_STACK_DEPTH
	MAX_CLIENT_ATTRIB_STACK_DEPTH = C.GL_MAX_CLIENT_ATTRIB_STACK_DEPTH
	MAX_CLIP_PLANES               = C.GL_MAX_CLIP_PLANES
	MAX_EVAL_ORDER                = C.GL_MAX_EVAL_ORDER
	MAX_LIGHTS                    = C.GL_MAX_LIGHTS
	MAX_LIST_NESTING              = C.GL_MAX_LIST_NESTING
	MAX_MODELVIEW_STACK_DEPTH     = C.GL_MAX_MODELVIEW_STACK_DEPTH
	MAX_NAME_STACK_DEPTH          = C.GL_MAX_NAME_STACK_DEPTH
	MAX_PIXEL_MAP_TABLE           = C.GL_MAX_PIXEL_MAP_TABLE
	MAX_PROJECTION_STACK_DEPTH    = C.GL_MAX_PROJECTION_STACK_DEPTH
	MAX_TEXTURE_SIZE              = C.GL_MAX_TEXTURE_SIZE
	MAX_TEXTURE_STACK_DEPTH       = C.GL_MAX_TEXTURE_STACK_DEPTH
	MAX_VIEWPORT_DIMS             = C.GL_MAX_VIEWPORT_DIMS
	MODELVIEW                     = C.GL_MODELVIEW
	MODELVIEW_MATRIX              = C.GL_MODELVIEW_MATRIX
	MODELVIEW_STACK_DEPTH         = C.GL_MODELVIEW_STACK_DEPTH
	MODULATE                      = C.GL_MODULATE
	MULT                          = C.GL_MULT
	N3F_V3F                       = C.GL_N3F_V3F
	NAME_STACK_DEPTH              = C.GL_NAME_STACK_DEPTH
	NAND                          = C.GL_NAND
	NEAREST                       = C.GL_NEAREST
	NEAREST_MIPMAP_LINEAR         = C.GL_NEAREST_MIPMAP_LINEAR
	NEAREST_MIPMAP_NEAREST        = C.GL_NEAREST_MIPMAP_NEAREST
	NEVER                         = C.GL_NEVER
	NICEST                        = C.GL_NICEST
	NO_ERROR                      = C.GL_NO_ERROR
	NONE                          = C.GL_NONE
	NOOP                          = C.GL_NOOP
	NOR                           = C.GL_NOR
	NORMAL_ARRAY                  = C.GL_NORMAL_ARRAY
	NORMAL_ARRAY_POINTER          = C.GL_NORMAL_ARRAY_POINTER
	NORMAL_ARRAY_STRIDE           = C.GL_NORMAL_ARRAY_STRIDE
	NORMAL_ARRAY_TYPE             = C.GL_NORMAL_ARRAY_TYPE
	NORMALIZE                     = C.GL_NORMALIZE
	NOTEQUAL                      = C.GL_NOTEQUAL
	OBJECT_LINEAR                 = C.GL_OBJECT_LINEAR
	OBJECT_PLANE                  = C.GL_OBJECT_PLANE
	ONE                           = C.GL_ONE
	ONE_MINUS_DST_ALPHA           = C.GL_ONE_MINUS_DST_ALPHA
	ONE_MINUS_DST_COLOR           = C.GL_ONE_MINUS_DST_COLOR
	ONE_MINUS_SRC_ALPHA           = C.GL_ONE_MINUS_SRC_ALPHA
	ONE_MINUS_SRC_COLOR           = C.GL_ONE_MINUS_SRC_COLOR
	OR                            = C.GL_OR
	OR_INVERTED                   = C.GL_OR_INVERTED
	OR_REVERSE                    = C.GL_OR_REVERSE
	ORDER                         = C.GL_ORDER
	OUT_OF_MEMORY                 = C.GL_OUT_OF_MEMORY
	PACK_ALIGNMENT                = C.GL_PACK_ALIGNMENT
	PACK_LSB_FIRST                = C.GL_PACK_LSB_FIRST
	PACK_ROW_LENGTH               = C.GL_PACK_ROW_LENGTH
	PACK_SKIP_PIXELS              = C.GL_PACK_SKIP_PIXELS
	PACK_SKIP_ROWS                = C.GL_PACK_SKIP_ROWS
	PACK_SWAP_BYTES               = C.GL_PACK_SWAP_BYTES
	PASS_THROUGH_TOKEN            = C.GL_PASS_THROUGH_TOKEN
	PERSPECTIVE_CORRECTION_HINT   = C.GL_PERSPECTIVE_CORRECTION_HINT
	PIXEL_MAP_A_TO_A              = C.GL_PIXEL_MAP_A_TO_A
	PIXEL_MAP_A_TO_A_SIZE         = C.GL_PIXEL_MAP_A_TO_A_SIZE
	PIXEL_MAP_B_TO_B              = C.GL_PIXEL_MAP_B_TO_B
	PIXEL_MAP_B_TO_B_SIZE         = C.GL_PIXEL_MAP_B_TO_B_SIZE
	PIXEL_MAP_G_TO_G              = C.GL_PIXEL_MAP_G_TO_G
	PIXEL_MAP_G_TO_G_SIZE         = C.GL_PIXEL_MAP_G_TO_G_SIZE
	PIXEL_MAP_I_TO_A              = C.GL_PIXEL_MAP_I_TO_A
	PIXEL_MAP_I_TO_A_SIZE         = C.GL_PIXEL_MAP_I_TO_A_SIZE
	PIXEL_MAP_I_TO_B              = C.GL_PIXEL_MAP_I_TO_B
	PIXEL_MAP_I_TO_B_SIZE         = C.GL_PIXEL_MAP_I_TO_B_SIZE
	PIXEL_MAP_I_TO_G              = C.GL_PIXEL_MAP_I_TO_G
	PIXEL_MAP_I_TO_G_SIZE         = C.GL_PIXEL_MAP_I_TO_G_SIZE
	PIXEL_MAP_I_TO_I              = C.GL_PIXEL_MAP_I_TO_I
	PIXEL_MAP_I_TO_I_SIZE         = C.GL_PIXEL_MAP_I_TO_I_SIZE
	PIXEL_MAP_I_TO_R              = C.GL_PIXEL_MAP_I_TO_R
	PIXEL_MAP_I_TO_R_SIZE         = C.GL_PIXEL_MAP_I_TO_R_SIZE
	PIXEL_MAP_R_TO_R              = C.GL_PIXEL_MAP_R_TO_R
	PIXEL_MAP_R_TO_R_SIZE         = C.GL_PIXEL_MAP_R_TO_R_SIZE
	PIXEL_MAP_S_TO_S              = C.GL_PIXEL_MAP_S_TO_S
	PIXEL_MAP_S_TO_S_SIZE         = C.GL_PIXEL_MAP_S_TO_S_SIZE
	PIXEL_MODE_BIT                = C.GL_PIXEL_MODE_BIT
	POINT                         = C.GL_POINT
	POINT_BIT                     = C.GL_POINT_BIT
	POINT_SIZE                    = C.GL_POINT_SIZE
	POINT_SIZE_GRANULARITY        = C.GL_POINT_SIZE_GRANULARITY
	POINT_SIZE_RANGE              = C.GL_POINT_SIZE_RANGE
	POINT_SMOOTH                  = C.GL_POINT_SMOOTH
	POINT_SMOOTH_HINT             = C.GL_POINT_SMOOTH_HINT
	POINT_TOKEN                   = C.GL_POINT_TOKEN
	POINTS                        = C.GL_POINTS
	POLYGON                       = C.GL_POLYGON
	POLYGON_BIT                   = C.GL_POLYGON_BIT
	POLYGON_MODE                  = C.GL_POLYGON_MODE
	POLYGON_OFFSET_FACTOR         = C.GL_POLYGON_OFFSET_FACTOR
	POLYGON_OFFSET_FILL           = C.GL_POLYGON_OFFSET_FILL
	POLYGON_OFFSET_LINE           = C.GL_POLYGON_OFFSET_LINE
	POLYGON_OFFSET_POINT          = C.GL_POLYGON_OFFSET_POINT
	POLYGON_OFFSET_UNITS          = C.GL_POLYGON_OFFSET_UNITS
	POLYGON_SMOOTH                = C.GL_POLYGON_SMOOTH
	POLYGON_SMOOTH_HINT           = C.GL_POLYGON_SMOOTH_HINT
	POLYGON_STIPPLE               = C.GL_POLYGON_STIPPLE
	POLYGON_STIPPLE_BIT           = C.GL_POLYGON_STIPPLE_BIT
	POLYGON_TOKEN                 = C.GL_POLYGON_TOKEN
	POSITION                      = C.GL_POSITION
	PROJECTION                    = C.GL_PROJECTION
	PROJECTION_MATRIX             = C.GL_PROJECTION_MATRIX
	PROJECTION_STACK_DEPTH        = C.GL_PROJECTION_STACK_DEPTH
	PROXY_TEXTURE_1D              = C.GL_PROXY_TEXTURE_1D
	PROXY_TEXTURE_2D              = C.GL_PROXY_TEXTURE_2D
	Q                             = C.GL_Q
	QUAD_STRIP                    = C.GL_QUAD_STRIP
	QUADRATIC_ATTENUATION         = C.GL_QUADRATIC_ATTENUATION
	QUADS                         = C.GL_QUADS
	R                             = C.GL_R
	R3_G3_B2                      = C.GL_R3_G3_B2
	READ_BUFFER                   = C.GL_READ_BUFFER
	RED                           = C.GL_RED
	RED_BIAS                      = C.GL_RED_BIAS
	RED_BITS                      = C.GL_RED_BITS
	RED_SCALE                     = C.GL_RED_SCALE
	RENDER                        = C.GL_RENDER
	RENDER_MODE                   = C.GL_RENDER_MODE
	RENDERER                      = C.GL_RENDERER
	REPEAT                        = C.GL_REPEAT
	REPLACE                       = C.GL_REPLACE
	RETURN                        = C.GL_RETURN
	RGB                           = C.GL_RGB
	RGB10                         = C.GL_RGB10
	RGB10_A2                      = C.GL_RGB10_A2
	RGB12                         = C.GL_RGB12
	RGB16                         = C.GL_RGB16
	RGB4                          = C.GL_RGB4
	RGB5                          = C.GL_RGB5
	RGB5_A1                       = C.GL_RGB5_A1
	RGB8                          = C.GL_RGB8
	RGBA                          = C.GL_RGBA
	RGBA12                        = C.GL_RGBA12
	RGBA16                        = C.GL_RGBA16
	RGBA2                         = C.GL_RGBA2
	RGBA4                         = C.GL_RGBA4
	RGBA8                         = C.GL_RGBA8
	RGBA_MODE                     = C.GL_RGBA_MODE
	RIGHT                         = C.GL_RIGHT
	S                             = C.GL_S
	SCISSOR_BIT                   = C.GL_SCISSOR_BIT
	SCISSOR_BOX                   = C.GL_SCISSOR_BOX
	SCISSOR_TEST                  = C.GL_SCISSOR_TEST
	SELECT                        = C.GL_SELECT
	SELECTION_BUFFER_POINTER      = C.GL_SELECTION_BUFFER_POINTER
	SELECTION_BUFFER_SIZE         = C.GL_SELECTION_BUFFER_SIZE
	SET                           = C.GL_SET
	SHADE_MODEL                   = C.GL_SHADE_MODEL
	SHININESS                     = C.GL_SHININESS
	SHORT                         = C.GL_SHORT
	SMOOTH                        = C.GL_SMOOTH
	SPECULAR                      = C.GL_SPECULAR
	SPHERE_MAP                    = C.GL_SPHERE_MAP
	SPOT_CUTOFF                   = C.GL_SPOT_CUTOFF
	SPOT_DIRECTION                = C.GL_SPOT_DIRECTION
	SPOT_EXPONENT                 = C.GL_SPOT_EXPONENT
	SRC_ALPHA                     = C.GL_SRC_ALPHA
	SRC_ALPHA_SATURATE            = C.GL_SRC_ALPHA_SATURATE
	SRC_COLOR                     = C.GL_SRC_COLOR
	STACK_OVERFLOW                = C.GL_STACK_OVERFLOW
	STACK_UNDERFLOW               = C.GL_STACK_UNDERFLOW
	STENCIL                       = C.GL_STENCIL
	STENCIL_BITS                  = C.GL_STENCIL_BITS
	STENCIL_BUFFER_BIT            = C.GL_STENCIL_BUFFER_BIT
	STENCIL_CLEAR_VALUE           = C.GL_STENCIL_CLEAR_VALUE
	STENCIL_FAIL                  = C.GL_STENCIL_FAIL
	STENCIL_FUNC                  = C.GL_STENCIL_FUNC
	STENCIL_INDEX                 = C.GL_STENCIL_INDEX
	STENCIL_PASS_DEPTH_FAIL       = C.GL_STENCIL_PASS_DEPTH_FAIL
	STENCIL_PASS_DEPTH_PASS       = C.GL_STENCIL_PASS_DEPTH_PASS
	STENCIL_REF                   = C.GL_STENCIL_REF
	STENCIL_TEST                  = C.GL_STENCIL_TEST
	STENCIL_VALUE_MASK            = C.GL_STENCIL_VALUE_MASK
	STENCIL_WRITEMASK             = C.GL_STENCIL_WRITEMASK
	STEREO                        = C.GL_STEREO
	SUBPIXEL_BITS                 = C.GL_SUBPIXEL_BITS
	T                             = C.GL_T
	T2F_C3F_V3F                   = C.GL_T2F_C3F_V3F
	T2F_C4F_N3F_V3F               = C.GL_T2F_C4F_N3F_V3F
	T2F_C4UB_V3F                  = C.GL_T2F_C4UB_V3F
	T2F_N3F_V3F                   = C.GL_T2F_N3F_V3F
	T2F_V3F                       = C.GL_T2F_V3F
	T4F_C4F_N3F_V4F               = C.GL_T4F_C4F_N3F_V4F
	T4F_V4F                       = C.GL_T4F_V4F
	TEXTURE                       = C.GL_TEXTURE
	TEXTURE_1D                    = C.GL_TEXTURE_1D
	TEXTURE_2D                    = C.GL_TEXTURE_2D
	TEXTURE_ALPHA_SIZE            = C.GL_TEXTURE_ALPHA_SIZE
	TEXTURE_BINDING_1D            = C.GL_TEXTURE_BINDING_1D
	TEXTURE_BINDING_2D            = C.GL_TEXTURE_BINDING_2D
	TEXTURE_BIT                   = C.GL_TEXTURE_BIT
	TEXTURE_BLUE_SIZE             = C.GL_TEXTURE_BLUE_SIZE
	TEXTURE_BORDER                = C.GL_TEXTURE_BORDER
	TEXTURE_BORDER_COLOR          = C.GL_TEXTURE_BORDER_COLOR
	TEXTURE_COMPONENTS            = C.GL_TEXTURE_COMPONENTS
	TEXTURE_COORD_ARRAY           = C.GL_TEXTURE_COORD_ARRAY
	TEXTURE_COORD_ARRAY_POINTER   = C.GL_TEXTURE_COORD_ARRAY_POINTER
	TEXTURE_COORD_ARRAY_SIZE      = C.GL_TEXTURE_COORD_ARRAY_SIZE
	TEXTURE_COORD_ARRAY_STRIDE    = C.GL_TEXTURE_COORD_ARRAY_STRIDE
	TEXTURE_COORD_ARRAY_TYPE      = C.GL_TEXTURE_COORD_ARRAY_TYPE
	TEXTURE_ENV                   = C.GL_TEXTURE_ENV
	TEXTURE_ENV_COLOR             = C.GL_TEXTURE_ENV_COLOR
	TEXTURE_ENV_MODE              = C.GL_TEXTURE_ENV_MODE
	TEXTURE_GEN_MODE              = C.GL_TEXTURE_GEN_MODE
	TEXTURE_GEN_Q                 = C.GL_TEXTURE_GEN_Q
	TEXTURE_GEN_R                 = C.GL_TEXTURE_GEN_R
	TEXTURE_GEN_S                 = C.GL_TEXTURE_GEN_S
	TEXTURE_GEN_T                 = C.GL_TEXTURE_GEN_T
	TEXTURE_GREEN_SIZE            = C.GL_TEXTURE_GREEN_SIZE
	TEXTURE_HEIGHT                = C.GL_TEXTURE_HEIGHT
	TEXTURE_INTENSITY_SIZE        = C.GL_TEXTURE_INTENSITY_SIZE
	TEXTURE_INTERNAL_FORMAT       = C.GL_TEXTURE_INTERNAL_FORMAT
	TEXTURE_LUMINANCE_SIZE        = C.GL_TEXTURE_LUMINANCE_SIZE
	TEXTURE_MAG_FILTER            = C.GL_TEXTURE_MAG_FILTER
	TEXTURE_MATRIX                = C.GL_TEXTURE_MATRIX
	TEXTURE_MIN_FILTER            = C.GL_TEXTURE_MIN_FILTER
	TEXTURE_PRIORITY              = C.GL_TEXTURE_PRIORITY
	TEXTURE_RED_SIZE              = C.GL_TEXTURE_RED_SIZE
	TEXTURE_RESIDENT              = C.GL_TEXTURE_RESIDENT
	TEXTURE_STACK_DEPTH           = C.GL_TEXTURE_STACK_DEPTH
	TEXTURE_WIDTH                 = C.GL_TEXTURE_WIDTH
	TEXTURE_WRAP_S                = C.GL_TEXTURE_WRAP_S
	TEXTURE_WRAP_T                = C.GL_TEXTURE_WRAP_T
	TRANSFORM_BIT                 = C.GL_TRANSFORM_BIT
	TRIANGLE_FAN                  = C.GL_TRIANGLE_FAN
	TRIANGLE_STRIP                = C.GL_TRIANGLE_STRIP
	TRIANGLES                     = C.GL_TRIANGLES
	TRUE                          = C.GL_TRUE
	UNPACK_ALIGNMENT              = C.GL_UNPACK_ALIGNMENT
	UNPACK_LSB_FIRST              = C.GL_UNPACK_LSB_FIRST
	UNPACK_ROW_LENGTH             = C.GL_UNPACK_ROW_LENGTH
	UNPACK_SKIP_PIXELS            = C.GL_UNPACK_SKIP_PIXELS
	UNPACK_SKIP_ROWS              = C.GL_UNPACK_SKIP_ROWS
	UNPACK_SWAP_BYTES             = C.GL_UNPACK_SWAP_BYTES
	UNSIGNED_BYTE                 = C.GL_UNSIGNED_BYTE
	UNSIGNED_INT                  = C.GL_UNSIGNED_INT
	UNSIGNED_SHORT                = C.GL_UNSIGNED_SHORT
	V2F                           = C.GL_V2F
	V3F                           = C.GL_V3F
	VENDOR                        = C.GL_VENDOR
	VERSION                       = C.GL_VERSION
	VERTEX_ARRAY                  = C.GL_VERTEX_ARRAY
	VERTEX_ARRAY_POINTER          = C.GL_VERTEX_ARRAY_POINTER
	VERTEX_ARRAY_SIZE             = C.GL_VERTEX_ARRAY_SIZE
	VERTEX_ARRAY_STRIDE           = C.GL_VERTEX_ARRAY_STRIDE
	VERTEX_ARRAY_TYPE             = C.GL_VERTEX_ARRAY_TYPE
	VIEWPORT                      = C.GL_VIEWPORT
	VIEWPORT_BIT                  = C.GL_VIEWPORT_BIT
	XOR                           = C.GL_XOR
	ZERO                          = C.GL_ZERO
	ZOOM_X                        = C.GL_ZOOM_X
	ZOOM_Y                        = C.GL_ZOOM_Y
)

func Accum(op GLenum, value float32) {
	C.glAccum(C.GLenum(op), C.GLfloat(value))
}

func AlphaFunc(funcion GLenum, ref float32) {
	C.glAlphaFunc(C.GLenum(funcion), C.GLclampf(ref))
}

func AreTexturesResident(textures []Texture, residences []bool) bool {
	sz := len(textures)
	if sz == 0 {
		return false
	}

	if sz != len(residences) {
		panic("Residences slice must be equal in length to textures slice.")
	}

	return goBool(C.glAreTexturesResident(
		C.GLsizei(sz),
		(*C.GLuint)(unsafe.Pointer(&textures[0])),
		(*C.GLboolean)(unsafe.Pointer(&residences[0])),
	))
}

func (texture Texture) IsResident() bool {
	var b C.GLboolean
	return goBool(C.glAreTexturesResident(C.GLsizei(1), (*C.GLuint)(&texture), &b))
}

func ArrayElement(i int) {
	C.glArrayElement(C.GLint(i))
}

func Begin(mode GLenum) {
	C.glBegin(C.GLenum(mode))
}

func (texture Texture) Bind(target GLenum) {
	C.glBindTexture(C.GLenum(target), C.GLuint(texture))
}

func (_ Texture) Unbind(target GLenum) {
	C.glBindTexture(C.GLenum(target), C.GLuint(0))
}

func Bitmap(width int, height int, xorig float32, yorig float32, xmove float32, ymove float32, bitmap []uint8) {
	C.glBitmap(C.GLsizei(width), C.GLsizei(height), C.GLfloat(xorig), C.GLfloat(yorig), C.GLfloat(xmove), C.GLfloat(ymove), (*C.GLubyte)(&bitmap[0]))
}

func BlendFunc(sfactor GLenum, dfactor GLenum) {
	C.glBlendFunc(C.GLenum(sfactor), C.GLenum(dfactor))
}

func (list List) Call() {
	C.glCallList(C.GLuint(list))
}

func CallLists(n int, typ GLenum, lists []List) {
	C.glCallLists(C.GLsizei(n), C.GLenum(typ), glPointer(lists))
}

func Clear(mask uint32) {
	C.glClear(C.GLbitfield(mask))
}

func ClearAccum(red float32, green float32, blue float32, alpha float32) {
	C.glClearAccum(C.GLfloat(red), C.GLfloat(green), C.GLfloat(blue), C.GLfloat(alpha))
}

func ClearColor(red float32, green float32, blue float32, alpha float32) {
	C.glClearColor(C.GLclampf(red), C.GLclampf(green), C.GLclampf(blue), C.GLclampf(alpha))
}

func ClearDepth(depth float64) {
	C.glClearDepth(C.GLclampd(depth))
}

func ClearIndex(c float32) {
	C.glClearIndex(C.GLfloat(c))
}

func ClearStencil(s int) {
	C.glClearStencil(C.GLint(s))
}

func ClipPlane(plane GLenum, equation []float64) {
	C.glClipPlane(C.GLenum(plane), (*C.GLdouble)(&equation[0]))
}

func Color3b(red int8, green int8, blue int8) {
	C.glColor3b(C.GLbyte(red), C.GLbyte(green), C.GLbyte(blue))
}

func Color3bv(v []int8) {
	C.glColor3bv((*C.GLbyte)(&v[0]))
}

func Color3d(red float64, green float64, blue float64) {
	C.glColor3d(C.GLdouble(red), C.GLdouble(green), C.GLdouble(blue))
}

func Color3dv(v []float64) {
	C.glColor3dv((*C.GLdouble)(&v[0]))
}

func Color3f(red float32, green float32, blue float32) {
	C.glColor3f(C.GLfloat(red), C.GLfloat(green), C.GLfloat(blue))
}

func Color3fv(v []float32) {
	C.glColor3fv((*C.GLfloat)(&v[0]))
}

func Color3i(red int, green int, blue int) {
	C.glColor3i(C.GLint(red), C.GLint(green), C.GLint(blue))
}

func Color3iv(v []int32) {
	C.glColor3iv((*C.GLint)(&v[0]))
}

func Color3s(red int16, green int16, blue int16) {
	C.glColor3s(C.GLshort(red), C.GLshort(green), C.GLshort(blue))
}

func Color3sv(v []int16) {
	C.glColor3sv((*C.GLshort)(&v[0]))
}

func Color3ub(red uint8, green uint8, blue uint8) {
	C.glColor3ub(C.GLubyte(red), C.GLubyte(green), C.GLubyte(blue))
}

func Color3ubv(v []uint8) {
	C.glColor3ubv((*C.GLubyte)(&v[0]))
}

func Color3ui(red uint, green uint, blue uint) {
	C.glColor3ui(C.GLuint(red), C.GLuint(green), C.GLuint(blue))
}

func Color3uiv(v []uint32) {
	C.glColor3uiv((*C.GLuint)(&v[0]))
}

func Color3us(red uint16, green uint16, blue uint16) {
	C.glColor3us(C.GLushort(red), C.GLushort(green), C.GLushort(blue))
}

func Color3usv(v []uint16) {
	C.glColor3usv((*C.GLushort)(&v[0]))
}

func Color4b(red int8, green int8, blue int8, alpha int8) {
	C.glColor4b(C.GLbyte(red), C.GLbyte(green), C.GLbyte(blue), C.GLbyte(alpha))
}

func Color4bv(v []int8) {
	C.glColor4bv((*C.GLbyte)(&v[0]))
}

func Color4d(red float64, green float64, blue float64, alpha float64) {
	C.glColor4d(C.GLdouble(red), C.GLdouble(green), C.GLdouble(blue), C.GLdouble(alpha))
}

func Color4dv(v []float64) {
	C.glColor4dv((*C.GLdouble)(&v[0]))
}

func Color4f(red float32, green float32, blue float32, alpha float32) {
	C.glColor4f(C.GLfloat(red), C.GLfloat(green), C.GLfloat(blue), C.GLfloat(alpha))
}

func Color4fv(v []float32) {
	C.glColor4fv((*C.GLfloat)(&v[0]))
}

func Color4i(red int, green int, blue int, alpha int) {
	C.glColor4i(C.GLint(red), C.GLint(green), C.GLint(blue), C.GLint(alpha))
}

func Color4iv(v []int32) {
	C.glColor4iv((*C.GLint)(&v[0]))
}

func Color4s(red int16, green int16, blue int16, alpha int16) {
	C.glColor4s(C.GLshort(red), C.GLshort(green), C.GLshort(blue), C.GLshort(alpha))
}

func Color4sv(v []int16) {
	C.glColor4sv((*C.GLshort)(&v[0]))
}

func Color4ub(red uint8, green uint8, blue uint8, alpha uint8) {
	C.glColor4ub(C.GLubyte(red), C.GLubyte(green), C.GLubyte(blue), C.GLubyte(alpha))
}

func Color4ubv(v []uint8) {
	C.glColor4ubv((*C.GLubyte)(&v[0]))
}

func Color4ui(red uint, green uint, blue uint, alpha uint) {
	C.glColor4ui(C.GLuint(red), C.GLuint(green), C.GLuint(blue), C.GLuint(alpha))
}

func Color4uiv(v []uint32) {
	C.glColor4uiv((*C.GLuint)(&v[0]))
}

func Color4us(red uint16, green uint16, blue uint16, alpha uint16) {
	C.glColor4us(C.GLushort(red), C.GLushort(green), C.GLushort(blue), C.GLushort(alpha))
}

func Color4usv(v []uint16) {
	C.glColor4usv((*C.GLushort)(&v[0]))
}

func ColorMask(red bool, green bool, blue bool, alpha bool) {
	C.glColorMask(glBool(red), glBool(green), glBool(blue), glBool(alpha))
}

func ColorMaterial(face GLenum, mode GLenum) {
	C.glColorMaterial(C.GLenum(face), C.GLenum(mode))
}

func ColorPointer(size int, typ GLenum, stride int, pointer interface{}) {
	C.glColorPointer(C.GLint(size), C.GLenum(typ), C.GLsizei(stride), glPointer(pointer))
}

func CopyPixels(x int, y int, width int, height int, typ GLenum) {
	C.glCopyPixels(C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height), C.GLenum(typ))
}

func CopyTexImage1D(target GLenum, level int, internalFormat GLenum, x int, y int, width int, border int) {
	C.glCopyTexImage1D(C.GLenum(target), C.GLint(level), C.GLenum(internalFormat), C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLint(border))
}

func CopyTexImage2D(target GLenum, level int, internalFormat GLenum, x int, y int, width int, height int, border int) {
	C.glCopyTexImage2D(C.GLenum(target), C.GLint(level), C.GLenum(internalFormat), C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height), C.GLint(border))
}

func CopyTexSubImage1D(target GLenum, level int, xoffset int, x int, y int, width int) {
	C.glCopyTexSubImage1D(C.GLenum(target), C.GLint(level), C.GLint(xoffset), C.GLint(x), C.GLint(y), C.GLsizei(width))
}

func CopyTexSubImage2D(target GLenum, level int, xoffset int, yoffset int, x int, y int, width int, height int) {
	C.glCopyTexSubImage2D(C.GLenum(target), C.GLint(level), C.GLint(xoffset), C.GLint(yoffset), C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height))
}

func CullFace(mode GLenum) {
	C.glCullFace(C.GLenum(mode))
}

func (list List) Delete(rang int) {
	C.glDeleteLists(C.GLuint(list), C.GLsizei(rang))
}

func DeleteTextures(textures []Texture) {
	C.glDeleteTextures(C.GLsizei(len(textures)), (*C.GLuint)(&textures[0]))
}

func (texture Texture) Delete() {
	C.glDeleteTextures(C.GLsizei(1), (*C.GLuint)(&texture))
}

func DepthFunc(function GLenum) {
	C.glDepthFunc(C.GLenum(function))
}

func DepthMask(flag bool) {
	C.glDepthMask(glBool(flag))
}

func DepthRange(zNear float64, zFar float64) {
	C.glDepthRange(C.GLclampd(zNear), C.GLclampd(zFar))
}

func Disable(cap GLenum) {
	C.glDisable(C.GLenum(cap))
}

func DisableClientState(array GLenum) {
	C.glDisableClientState(C.GLenum(array))
}

func DrawArrays(mode GLenum, first int, count int) {
	C.glDrawArrays(C.GLenum(mode), C.GLint(first), C.GLsizei(count))
}

func DrawBuffer(mode GLenum) {
	C.glDrawBuffer(C.GLenum(mode))
}

func DrawElements(mode GLenum, count int, typ GLenum, indices interface{}) {
	C.glDrawElements(C.GLenum(mode), C.GLsizei(count), C.GLenum(typ), glPointer(indices))
}

func DrawPixels(width int, height int, format GLenum, typ GLenum, pixels interface{}) {
	C.glDrawPixels(C.GLsizei(width), C.GLsizei(height), C.GLenum(format), C.GLenum(typ), glPointer(pixels))
}

func EdgeFlag(flag bool) {
	C.glEdgeFlag(glBool(flag))
}

func EdgeFlagPointer(stride int, pointer interface{}) {
	C.glEdgeFlagPointer(C.GLsizei(stride), glPointer(pointer))
}

func EdgeFlagv(flag []bool) {
	C.glEdgeFlagv((*C.GLboolean)(unsafe.Pointer(&flag[0])))
}

func Enable(cap GLenum) {
	C.glEnable(C.GLenum(cap))
}

func EnableClientState(array GLenum) {
	C.glEnableClientState(C.GLenum(array))
}

func End() {
	C.glEnd()
}

func EndList() {
	C.glEndList()
}

func EvalCoord1d(u float64) {
	C.glEvalCoord1d(C.GLdouble(u))
}

func EvalCoord1dv(u []float64) {
	C.glEvalCoord1dv((*C.GLdouble)(&u[0]))
}

func EvalCoord1f(u float32) {
	C.glEvalCoord1f(C.GLfloat(u))
}

func EvalCoord1fv(u []float32) {
	C.glEvalCoord1fv((*C.GLfloat)(&u[0]))
}

func EvalCoord2d(u float64, v float64) {
	C.glEvalCoord2d(C.GLdouble(u), C.GLdouble(v))
}

func EvalCoord2dv(u []float64) {
	C.glEvalCoord2dv((*C.GLdouble)(&u[0]))
}

func EvalCoord2f(u float32, v float32) {
	C.glEvalCoord2f(C.GLfloat(u), C.GLfloat(v))
}

func EvalCoord2fv(u []float32) {
	C.glEvalCoord2fv((*C.GLfloat)(&u[0]))
}

func EvalMesh1(mode GLenum, i1 int, i2 int) {
	C.glEvalMesh1(C.GLenum(mode), C.GLint(i1), C.GLint(i2))
}

func EvalMesh2(mode GLenum, i1 int, i2 int, j1 int, j2 int) {
	C.glEvalMesh2(C.GLenum(mode), C.GLint(i1), C.GLint(i2), C.GLint(j1), C.GLint(j2))
}

func EvalPoint1(i int) {
	C.glEvalPoint1(C.GLint(i))
}

func EvalPoint2(i int, j int) {
	C.glEvalPoint2(C.GLint(i), C.GLint(j))
}

func FeedbackBuffer(size int, typ GLenum, buffer []float32) {
	C.glFeedbackBuffer(C.GLsizei(size), C.GLenum(typ), (*C.GLfloat)(&buffer[0]))
}

func Finish() {
	C.glFinish()
}

func Flush() {
	C.glFlush()
}

func Fogf(pname GLenum, param float32) {
	C.glFogf(C.GLenum(pname), C.GLfloat(param))
}

func Fogfv(pname GLenum, params []float32) {
	C.glFogfv(C.GLenum(pname), (*C.GLfloat)(&params[0]))
}

func Fogi(pname GLenum, param int) {
	C.glFogi(C.GLenum(pname), C.GLint(param))
}

func Fogiv(pname GLenum, params []int32) {
	C.glFogiv(C.GLenum(pname), (*C.GLint)(&params[0]))
}

func FrontFace(mode GLenum) {
	C.glFrontFace(C.GLenum(mode))
}

func Frustum(left float64, right float64, bottom float64, top float64, zNear float64, zFar float64) {
	C.glFrustum(C.GLdouble(left), C.GLdouble(right), C.GLdouble(bottom), C.GLdouble(top), C.GLdouble(zNear), C.GLdouble(zFar))
}

func GenLists(rang int32) List {
	return List(C.glGenLists(C.GLsizei(rang)))
}

func GenTextures(textures []Texture) {
	C.glGenTextures(C.GLsizei(len(textures)), (*C.GLuint)(&textures[0]))
}

func GenTexture() (texture Texture) {
	C.glGenTextures(C.GLsizei(1), (*C.GLuint)(&texture))
	return
}

func GetBooleanv(pname GLenum, params []bool) {
	C.glGetBooleanv(C.GLenum(pname), (*C.GLboolean)(unsafe.Pointer(&params[0])))
}

func GetClipPlane(plane GLenum, equation []float64) {
	C.glGetClipPlane(C.GLenum(plane), (*C.GLdouble)(&equation[0]))
}

func GetDoublev(pname GLenum, params []float64) {
	C.glGetDoublev(C.GLenum(pname), (*C.GLdouble)(&params[0]))
}

func GetError() GLenum {
	return GLenum(C.glGetError())
}

func GetFloatv(pname GLenum, params []float32) {
	C.glGetFloatv(C.GLenum(pname), (*C.GLfloat)(&params[0]))
}

func GetIntegerv(pname GLenum, params []int32) {
	C.glGetIntegerv(C.GLenum(pname), (*C.GLint)(&params[0]))
}

func GetLightfv(light GLenum, pname GLenum, params []float32) {
	C.glGetLightfv(C.GLenum(light), C.GLenum(pname), (*C.GLfloat)(&params[0]))
}

func GetLightiv(light GLenum, pname GLenum, params []int32) {
	C.glGetLightiv(C.GLenum(light), C.GLenum(pname), (*C.GLint)(&params[0]))
}

func GetMapdv(target GLenum, query GLenum, v []float64) {
	C.glGetMapdv(C.GLenum(target), C.GLenum(query), (*C.GLdouble)(&v[0]))
}

func GetMapfv(target GLenum, query GLenum, v []float32) {
	C.glGetMapfv(C.GLenum(target), C.GLenum(query), (*C.GLfloat)(&v[0]))
}

func GetMapiv(target GLenum, query GLenum, v []int32) {
	C.glGetMapiv(C.GLenum(target), C.GLenum(query), (*C.GLint)(&v[0]))
}

func GetMaterialfv(face GLenum, pname GLenum, params []float32) {
	C.glGetMaterialfv(C.GLenum(face), C.GLenum(pname), (*C.GLfloat)(&params[0]))
}

func GetMaterialiv(face GLenum, pname GLenum, params []int32) {
	C.glGetMaterialiv(C.GLenum(face), C.GLenum(pname), (*C.GLint)(&params[0]))
}

func GetPixelMapfv(mp GLenum, values []float32) {
	C.glGetPixelMapfv(C.GLenum(mp), (*C.GLfloat)(&values[0]))
}

func GetPixelMapuiv(mp GLenum, values []uint32) {
	C.glGetPixelMapuiv(C.GLenum(mp), (*C.GLuint)(&values[0]))
}

func GetPixelMapusv(mp GLenum, values []uint16) {
	C.glGetPixelMapusv(C.GLenum(mp), (*C.GLushort)(&values[0]))
}

func GetPointerv(pname GLenum) (ptr unsafe.Pointer) {
	C.glGetPointerv(C.GLenum(pname), &ptr)
	return
}

func GetPolygonStipple(mask []uint8) {
	C.glGetPolygonStipple((*C.GLubyte)(&mask[0]))
}

func GetString(name GLenum) string {
	str := unsafe.Pointer(C.glGetString(C.GLenum(name)))
	return C.GoString((*C.char)(str))
}

func GetTexEnvfv(target GLenum, pname GLenum, params []float32) {
	C.glGetTexEnvfv(C.GLenum(target), C.GLenum(pname), (*C.GLfloat)(&params[0]))
}

func GetTexEnviv(target GLenum, pname GLenum, params []int32) {
	C.glGetTexEnviv(C.GLenum(target), C.GLenum(pname), (*C.GLint)(&params[0]))
}

func GetTexGendv(coord GLenum, pname GLenum, params []float64) {
	C.glGetTexGendv(C.GLenum(coord), C.GLenum(pname), (*C.GLdouble)(&params[0]))
}

func GetTexGenfv(coord GLenum, pname GLenum, params []float32) {
	C.glGetTexGenfv(C.GLenum(coord), C.GLenum(pname), (*C.GLfloat)(&params[0]))
}

func GetTexGeniv(coord GLenum, pname GLenum, params []int32) {
	C.glGetTexGeniv(C.GLenum(coord), C.GLenum(pname), (*C.GLint)(&params[0]))
}

func GetTexImage(target GLenum, level int, format GLenum, typ GLenum, pixels interface{}) {
	C.glGetTexImage(C.GLenum(target), C.GLint(level), C.GLenum(format), C.GLenum(typ), glPointer(pixels))
}

func GetTexLevelParameterfv(target GLenum, level int, pname GLenum, params []float32) {
	C.glGetTexLevelParameterfv(C.GLenum(target), C.GLint(level), C.GLenum(pname), (*C.GLfloat)(&params[0]))
}

func GetTexLevelParameteriv(target GLenum, level int, pname GLenum, params []int32) {
	C.glGetTexLevelParameteriv(C.GLenum(target), C.GLint(level), C.GLenum(pname), (*C.GLint)(&params[0]))
}

func GetTexParameterfv(target GLenum, pname GLenum, params []float32) {
	C.glGetTexParameterfv(C.GLenum(target), C.GLenum(pname), (*C.GLfloat)(&params[0]))
}

func GetTexParameteriv(target GLenum, pname GLenum, params []int32) {
	C.glGetTexParameteriv(C.GLenum(target), C.GLenum(pname), (*C.GLint)(&params[0]))
}

func Hint(target GLenum, mode GLenum) {
	C.glHint(C.GLenum(target), C.GLenum(mode))
}

func IndexMask(mask uint) {
	C.glIndexMask(C.GLuint(mask))
}

func IndexPointer(typ GLenum, stride int, pointer interface{}) {
	C.glIndexPointer(C.GLenum(typ), C.GLsizei(stride), glPointer(pointer))
}

func Indexd(c float64) {
	C.glIndexd(C.GLdouble(c))
}

func Indexdv(c []float64) {
	C.glIndexdv((*C.GLdouble)(&c[0]))
}

func Indexf(c float32) {
	C.glIndexf(C.GLfloat(c))
}

func Indexfv(c []float32) {
	C.glIndexfv((*C.GLfloat)(&c[0]))
}

func Indexi(c int) {
	C.glIndexi(C.GLint(c))
}

func Indexiv(c []int32) {
	C.glIndexiv((*C.GLint)(&c[0]))
}

func Indexs(c int16) {
	C.glIndexs(C.GLshort(c))
}

func Indexsv(c []int16) {
	C.glIndexsv((*C.GLshort)(&c[0]))
}

func Indexub(c uint8) {
	C.glIndexub(C.GLubyte(c))
}

func Indexubv(c []uint8) {
	C.glIndexubv((*C.GLubyte)(&c[0]))
}

func InitNames() {
	C.glInitNames()
}

func InterleavedArrays(format GLenum, stride int, pointer interface{}) {
	C.glInterleavedArrays(C.GLenum(format), C.GLsizei(stride), glPointer(pointer))
}

func IsEnabled(cap GLenum) bool {
	return goBool(C.glIsEnabled(C.GLenum(cap)))
}

func (object Object) IsList() bool {
	return goBool(C.glIsList(C.GLuint(object)))
}

func (object Object) IsTexture() bool {
	return goBool(C.glIsTexture(C.GLuint(object)))
}

func LightModelf(pname GLenum, param float32) {
	C.glLightModelf(C.GLenum(pname), C.GLfloat(param))
}

func LightModelfv(pname GLenum, params []float32) {
	C.glLightModelfv(C.GLenum(pname), (*C.GLfloat)(&params[0]))
}

func LightModeli(pname GLenum, param int) {
	C.glLightModeli(C.GLenum(pname), C.GLint(param))
}

func LightModeliv(pname GLenum, params []int32) {
	C.glLightModeliv(C.GLenum(pname), (*C.GLint)(&params[0]))
}

func Lightf(light GLenum, pname GLenum, param float32) {
	C.glLightf(C.GLenum(light), C.GLenum(pname), C.GLfloat(param))
}

func Lightfv(light GLenum, pname GLenum, params []float32) {
	C.glLightfv(C.GLenum(light), C.GLenum(pname), (*C.GLfloat)(&params[0]))
}

func Lighti(light GLenum, pname GLenum, param int) {
	C.glLighti(C.GLenum(light), C.GLenum(pname), C.GLint(param))
}

func Lightiv(light GLenum, pname GLenum, params []int32) {
	C.glLightiv(C.GLenum(light), C.GLenum(pname), (*C.GLint)(&params[0]))
}

func LineStipple(factor int, pattern uint16) {
	C.glLineStipple(C.GLint(factor), C.GLushort(pattern))
}

func LineWidth(width float32) {
	C.glLineWidth(C.GLfloat(width))
}

func (list List) Base() {
	C.glListBase(C.GLuint(list))
}

func LoadIdentity() {
	C.glLoadIdentity()
}

func LoadMatrixd(m []float64) {
	C.glLoadMatrixd((*C.GLdouble)(&m[0]))
}

func LoadMatrixf(m []float32) {
	C.glLoadMatrixf((*C.GLfloat)(&m[0]))
}

func LoadName(name uint) {
	C.glLoadName(C.GLuint(name))
}

func LogicOp(opcode GLenum) {
	C.glLogicOp(C.GLenum(opcode))
}

func Map1d(target GLenum, u1 float64, u2 float64, stride int, order int, points []float64) {
	C.glMap1d(C.GLenum(target), C.GLdouble(u1), C.GLdouble(u2), C.GLint(stride), C.GLint(order), (*C.GLdouble)(&points[0]))
}

func Map1f(target GLenum, u1 float32, u2 float32, stride int, order int, points []float32) {
	C.glMap1f(C.GLenum(target), C.GLfloat(u1), C.GLfloat(u2), C.GLint(stride), C.GLint(order), (*C.GLfloat)(&points[0]))
}

func Map2d(target GLenum, u1 float64, u2 float64, ustride int, uorder int, v1 float64, v2 float64, vstride int, vorder int, points []float64) {
	C.glMap2d(C.GLenum(target), C.GLdouble(u1), C.GLdouble(u2), C.GLint(ustride), C.GLint(uorder), C.GLdouble(v1), C.GLdouble(v2), C.GLint(vstride), C.GLint(vorder), (*C.GLdouble)(&points[0]))
}

func Map2f(target GLenum, u1 float32, u2 float32, ustride int, uorder int, v1 float32, v2 float32, vstride int, vorder int, points []float32) {
	C.glMap2f(C.GLenum(target), C.GLfloat(u1), C.GLfloat(u2), C.GLint(ustride), C.GLint(uorder), C.GLfloat(v1), C.GLfloat(v2), C.GLint(vstride), C.GLint(vorder), (*C.GLfloat)(&points[0]))
}

func MapGrid1d(un int, u1 float64, u2 float64) {
	C.glMapGrid1d(C.GLint(un), C.GLdouble(u1), C.GLdouble(u2))
}

func MapGrid1f(un int, u1 float32, u2 float32) {
	C.glMapGrid1f(C.GLint(un), C.GLfloat(u1), C.GLfloat(u2))
}

func MapGrid2d(un int, u1 float64, u2 float64, vn int, v1 float64, v2 float64) {
	C.glMapGrid2d(C.GLint(un), C.GLdouble(u1), C.GLdouble(u2), C.GLint(vn), C.GLdouble(v1), C.GLdouble(v2))
}

func MapGrid2f(un int, u1 float32, u2 float32, vn int, v1 float32, v2 float32) {
	C.glMapGrid2f(C.GLint(un), C.GLfloat(u1), C.GLfloat(u2), C.GLint(vn), C.GLfloat(v1), C.GLfloat(v2))
}

func Materialf(face GLenum, pname GLenum, param float32) {
	C.glMaterialf(C.GLenum(face), C.GLenum(pname), C.GLfloat(param))
}

func Materialfv(face GLenum, pname GLenum, params []float32) {
	C.glMaterialfv(C.GLenum(face), C.GLenum(pname), (*C.GLfloat)(&params[0]))
}

func Materiali(face GLenum, pname GLenum, param int) {
	C.glMateriali(C.GLenum(face), C.GLenum(pname), C.GLint(param))
}

func Materialiv(face GLenum, pname GLenum, params []int32) {
	C.glMaterialiv(C.GLenum(face), C.GLenum(pname), (*C.GLint)(&params[0]))
}

func MatrixMode(mode GLenum) {
	C.glMatrixMode(C.GLenum(mode))
}

func MultMatrixd(m []float64) {
	C.glMultMatrixd((*C.GLdouble)(&m[0]))
}

func MultMatrixf(m []float32) {
	C.glMultMatrixf((*C.GLfloat)(&m[0]))
}

func (list List) New(mode GLenum) {
	C.glNewList(C.GLuint(list), C.GLenum(mode))
}

func Normal3b(nx int8, ny int8, nz int8) {
	C.glNormal3b(C.GLbyte(nx), C.GLbyte(ny), C.GLbyte(nz))
}

func Normal3bv(v []int8) {
	C.glNormal3bv((*C.GLbyte)(&v[0]))
}

func Normal3d(nx float64, ny float64, nz float64) {
	C.glNormal3d(C.GLdouble(nx), C.GLdouble(ny), C.GLdouble(nz))
}

func Normal3dv(v []float64) {
	C.glNormal3dv((*C.GLdouble)(&v[0]))
}

func Normal3f(nx float32, ny float32, nz float32) {
	C.glNormal3f(C.GLfloat(nx), C.GLfloat(ny), C.GLfloat(nz))
}

func Normal3fv(v []float32) {
	C.glNormal3fv((*C.GLfloat)(&v[0]))
}

func Normal3i(nx int, ny int, nz int) {
	C.glNormal3i(C.GLint(nx), C.GLint(ny), C.GLint(nz))
}

func Normal3iv(v []int32) {
	C.glNormal3iv((*C.GLint)(&v[0]))
}

func Normal3s(nx int16, ny int16, nz int16) {
	C.glNormal3s(C.GLshort(nx), C.GLshort(ny), C.GLshort(nz))
}

func Normal3sv(v []int16) {
	C.glNormal3sv((*C.GLshort)(&v[0]))
}

func NormalPointer(typ GLenum, stride int, pointer interface{}) {
	C.glNormalPointer(C.GLenum(typ), C.GLsizei(stride), glPointer(pointer))
}

func Ortho(left float64, right float64, bottom float64, top float64, zNear float64, zFar float64) {
	C.glOrtho(C.GLdouble(left), C.GLdouble(right), C.GLdouble(bottom), C.GLdouble(top), C.GLdouble(zNear), C.GLdouble(zFar))
}

func PassThrough(token float32) {
	C.glPassThrough(C.GLfloat(token))
}

func PixelMapfv(mp GLenum, mapsize int, values []float32) {
	C.glPixelMapfv(C.GLenum(mp), C.GLsizei(mapsize), (*C.GLfloat)(&values[0]))
}

func PixelMapuiv(mp GLenum, mapsize int, values []uint32) {
	C.glPixelMapuiv(C.GLenum(mp), C.GLsizei(mapsize), (*C.GLuint)(&values[0]))
}

func PixelMapusv(mp GLenum, mapsize int, values []uint16) {
	C.glPixelMapusv(C.GLenum(mp), C.GLsizei(mapsize), (*C.GLushort)(&values[0]))
}

func PixelStoref(pname GLenum, param float32) {
	C.glPixelStoref(C.GLenum(pname), C.GLfloat(param))
}

func PixelStorei(pname GLenum, param int) {
	C.glPixelStorei(C.GLenum(pname), C.GLint(param))
}

func PixelTransferf(pname GLenum, param float32) {
	C.glPixelTransferf(C.GLenum(pname), C.GLfloat(param))
}

func PixelTransferi(pname GLenum, param int) {
	C.glPixelTransferi(C.GLenum(pname), C.GLint(param))
}

func PixelZoom(xfactor float32, yfactor float32) {
	C.glPixelZoom(C.GLfloat(xfactor), C.GLfloat(yfactor))
}

func PointSize(size float32) {
	C.glPointSize(C.GLfloat(size))
}

func PolygonMode(face GLenum, mode GLenum) {
	C.glPolygonMode(C.GLenum(face), C.GLenum(mode))
}

func PolygonOffset(factor float32, units float32) {
	C.glPolygonOffset(C.GLfloat(factor), C.GLfloat(units))
}

func PolygonStipple(mask []uint8) {
	C.glPolygonStipple((*C.GLubyte)(&mask[0]))
}

func PopAttrib() {
	C.glPopAttrib()
}

func PopClientAttrib() {
	C.glPopClientAttrib()
}

func PopMatrix() {
	C.glPopMatrix()
}

func PopName() {
	C.glPopName()
}

func PrioritizeTextures(textures []Texture, priorities []float32) {
	C.glPrioritizeTextures(C.GLsizei(len(textures)), (*C.GLuint)(&textures[0]), (*C.GLclampf)(&priorities[0]))
}

func PushAttrib(mask uint32) {
	C.glPushAttrib(C.GLbitfield(mask))
}

func PushClientAttrib(mask uint32) {
	C.glPushClientAttrib(C.GLbitfield(mask))
}

func PushMatrix() {
	C.glPushMatrix()
}

func PushName(name uint) {
	C.glPushName(C.GLuint(name))
}

func RasterPos2d(x float64, y float64) {
	C.glRasterPos2d(C.GLdouble(x), C.GLdouble(y))
}

func RasterPos2dv(v []float64) {
	C.glRasterPos2dv((*C.GLdouble)(&v[0]))
}

func RasterPos2f(x float32, y float32) {
	C.glRasterPos2f(C.GLfloat(x), C.GLfloat(y))
}

func RasterPos2fv(v []float32) {
	C.glRasterPos2fv((*C.GLfloat)(&v[0]))
}

func RasterPos2i(x int, y int) {
	C.glRasterPos2i(C.GLint(x), C.GLint(y))
}

func RasterPos2iv(v []int32) {
	C.glRasterPos2iv((*C.GLint)(&v[0]))
}

func RasterPos2s(x int16, y int16) {
	C.glRasterPos2s(C.GLshort(x), C.GLshort(y))
}

func RasterPos2sv(v []int16) {
	C.glRasterPos2sv((*C.GLshort)(&v[0]))
}

func RasterPos3d(x float64, y float64, z float64) {
	C.glRasterPos3d(C.GLdouble(x), C.GLdouble(y), C.GLdouble(z))
}

func RasterPos3dv(v []float64) {
	C.glRasterPos3dv((*C.GLdouble)(&v[0]))
}

func RasterPos3f(x float32, y float32, z float32) {
	C.glRasterPos3f(C.GLfloat(x), C.GLfloat(y), C.GLfloat(z))
}

func RasterPos3fv(v []float32) {
	C.glRasterPos3fv((*C.GLfloat)(&v[0]))
}

func RasterPos3i(x int, y int, z int) {
	C.glRasterPos3i(C.GLint(x), C.GLint(y), C.GLint(z))
}

func RasterPos3iv(v []int32) {
	C.glRasterPos3iv((*C.GLint)(&v[0]))
}

func RasterPos3s(x int16, y int16, z int16) {
	C.glRasterPos3s(C.GLshort(x), C.GLshort(y), C.GLshort(z))
}

func RasterPos3sv(v []int16) {
	C.glRasterPos3sv((*C.GLshort)(&v[0]))
}

func RasterPos4d(x float64, y float64, z float64, w float64) {
	C.glRasterPos4d(C.GLdouble(x), C.GLdouble(y), C.GLdouble(z), C.GLdouble(w))
}

func RasterPos4dv(v []float64) {
	C.glRasterPos4dv((*C.GLdouble)(&v[0]))
}

func RasterPos4f(x float32, y float32, z float32, w float32) {
	C.glRasterPos4f(C.GLfloat(x), C.GLfloat(y), C.GLfloat(z), C.GLfloat(w))
}

func RasterPos4fv(v []float32) {
	C.glRasterPos4fv((*C.GLfloat)(&v[0]))
}

func RasterPos4i(x int, y int, z int, w int) {
	C.glRasterPos4i(C.GLint(x), C.GLint(y), C.GLint(z), C.GLint(w))
}

func RasterPos4iv(v []int32) {
	C.glRasterPos4iv((*C.GLint)(&v[0]))
}

func RasterPos4s(x int16, y int16, z int16, w int16) {
	C.glRasterPos4s(C.GLshort(x), C.GLshort(y), C.GLshort(z), C.GLshort(w))
}

func RasterPos4sv(v []int16) {
	C.glRasterPos4sv((*C.GLshort)(&v[0]))
}

func ReadBuffer(mode GLenum) {
	C.glReadBuffer(C.GLenum(mode))
}

func ReadPixels(x int, y int, width int, height int, format GLenum, typ GLenum, pixels interface{}) {
	C.glReadPixels(C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height), C.GLenum(format), C.GLenum(typ), glPointer(pixels))
}

func Rectd(x1 float64, y1 float64, x2 float64, y2 float64) {
	C.glRectd(C.GLdouble(x1), C.GLdouble(y1), C.GLdouble(x2), C.GLdouble(y2))
}

func Rectdv(v1 []float64, v2 []float64) {
	C.glRectdv((*C.GLdouble)(&v1[0]), (*C.GLdouble)(&v2[0]))
}

func Rectf(x1 float32, y1 float32, x2 float32, y2 float32) {
	C.glRectf(C.GLfloat(x1), C.GLfloat(y1), C.GLfloat(x2), C.GLfloat(y2))
}

func Rectfv(v1 []float32, v2 []float32) {
	C.glRectfv((*C.GLfloat)(&v1[0]), (*C.GLfloat)(&v2[0]))
}

func Recti(x1 int, y1 int, x2 int, y2 int) {
	C.glRecti(C.GLint(x1), C.GLint(y1), C.GLint(x2), C.GLint(y2))
}

func Rectiv(v1 []int32, v2 []int32) {
	C.glRectiv((*C.GLint)(&v1[0]), (*C.GLint)(&v2[0]))
}

func Rects(x1 int16, y1 int16, x2 int16, y2 int16) {
	C.glRects(C.GLshort(x1), C.GLshort(y1), C.GLshort(x2), C.GLshort(y2))
}

func Rectsv(v1 []int16, v2 []int16) {
	C.glRectsv((*C.GLshort)(&v1[0]), (*C.GLshort)(&v2[0]))
}

func RenderMode(mode GLenum) int {
	return int(C.glRenderMode(C.GLenum(mode)))
}

func Rotated(angle float64, x float64, y float64, z float64) {
	C.glRotated(C.GLdouble(angle), C.GLdouble(x), C.GLdouble(y), C.GLdouble(z))
}

func Rotatef(angle float32, x float32, y float32, z float32) {
	C.glRotatef(C.GLfloat(angle), C.GLfloat(x), C.GLfloat(y), C.GLfloat(z))
}

func Scaled(x float64, y float64, z float64) {
	C.glScaled(C.GLdouble(x), C.GLdouble(y), C.GLdouble(z))
}

func Scalef(x float32, y float32, z float32) {
	C.glScalef(C.GLfloat(x), C.GLfloat(y), C.GLfloat(z))
}

func Scissor(x int, y int, width int, height int) {
	C.glScissor(C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height))
}

func SelectBuffer(size int, buffer []uint32) {
	C.glSelectBuffer(C.GLsizei(size), (*C.GLuint)(&buffer[0]))
}

func ShadeModel(mode GLenum) {
	C.glShadeModel(C.GLenum(mode))
}

func StencilFunc(function GLenum, ref int, mask uint) {
	C.glStencilFunc(C.GLenum(function), C.GLint(ref), C.GLuint(mask))
}

func StencilMask(mask uint) {
	C.glStencilMask(C.GLuint(mask))
}

func StencilOp(fail GLenum, zfail GLenum, zpass GLenum) {
	C.glStencilOp(C.GLenum(fail), C.GLenum(zfail), C.GLenum(zpass))
}

func TexCoord1d(s float64) {
	C.glTexCoord1d(C.GLdouble(s))
}

func TexCoord1dv(v []float64) {
	C.glTexCoord1dv((*C.GLdouble)(&v[0]))
}

func TexCoord1f(s float32) {
	C.glTexCoord1f(C.GLfloat(s))
}

func TexCoord1fv(v []float32) {
	C.glTexCoord1fv((*C.GLfloat)(&v[0]))
}

func TexCoord1i(s int) {
	C.glTexCoord1i(C.GLint(s))
}

func TexCoord1iv(v []int32) {
	C.glTexCoord1iv((*C.GLint)(&v[0]))
}

func TexCoord1s(s int16) {
	C.glTexCoord1s(C.GLshort(s))
}

func TexCoord1sv(v []int16) {
	C.glTexCoord1sv((*C.GLshort)(&v[0]))
}

func TexCoord2d(s float64, t float64) {
	C.glTexCoord2d(C.GLdouble(s), C.GLdouble(t))
}

func TexCoord2dv(v []float64) {
	C.glTexCoord2dv((*C.GLdouble)(&v[0]))
}

func TexCoord2f(s float32, t float32) {
	C.glTexCoord2f(C.GLfloat(s), C.GLfloat(t))
}

func TexCoord2fv(v []float32) {
	C.glTexCoord2fv((*C.GLfloat)(&v[0]))
}

func TexCoord2i(s int, t int) {
	C.glTexCoord2i(C.GLint(s), C.GLint(t))
}

func TexCoord2iv(v []int32) {
	C.glTexCoord2iv((*C.GLint)(&v[0]))
}

func TexCoord2s(s int16, t int16) {
	C.glTexCoord2s(C.GLshort(s), C.GLshort(t))
}

func TexCoord2sv(v []int16) {
	C.glTexCoord2sv((*C.GLshort)(&v[0]))
}

func TexCoord3d(s float64, t float64, r float64) {
	C.glTexCoord3d(C.GLdouble(s), C.GLdouble(t), C.GLdouble(r))
}

func TexCoord3dv(v []float64) {
	C.glTexCoord3dv((*C.GLdouble)(&v[0]))
}

func TexCoord3f(s float32, t float32, r float32) {
	C.glTexCoord3f(C.GLfloat(s), C.GLfloat(t), C.GLfloat(r))
}

func TexCoord3fv(v []float32) {
	C.glTexCoord3fv((*C.GLfloat)(&v[0]))
}

func TexCoord3i(s int, t int, r int) {
	C.glTexCoord3i(C.GLint(s), C.GLint(t), C.GLint(r))
}

func TexCoord3iv(v []int32) {
	C.glTexCoord3iv((*C.GLint)(&v[0]))
}

func TexCoord3s(s int16, t int16, r int16) {
	C.glTexCoord3s(C.GLshort(s), C.GLshort(t), C.GLshort(r))
}

func TexCoord3sv(v []int16) {
	C.glTexCoord3sv((*C.GLshort)(&v[0]))
}

func TexCoord4d(s float64, t float64, r float64, q float64) {
	C.glTexCoord4d(C.GLdouble(s), C.GLdouble(t), C.GLdouble(r), C.GLdouble(q))
}

func TexCoord4dv(v []float64) {
	C.glTexCoord4dv((*C.GLdouble)(&v[0]))
}

func TexCoord4f(s float32, t float32, r float32, q float32) {
	C.glTexCoord4f(C.GLfloat(s), C.GLfloat(t), C.GLfloat(r), C.GLfloat(q))
}

func TexCoord4fv(v []float32) {
	C.glTexCoord4fv((*C.GLfloat)(&v[0]))
}

func TexCoord4i(s int, t int, r int, q int) {
	C.glTexCoord4i(C.GLint(s), C.GLint(t), C.GLint(r), C.GLint(q))
}

func TexCoord4iv(v []int32) {
	C.glTexCoord4iv((*C.GLint)(&v[0]))
}

func TexCoord4s(s int16, t int16, r int16, q int16) {
	C.glTexCoord4s(C.GLshort(s), C.GLshort(t), C.GLshort(r), C.GLshort(q))
}

func TexCoord4sv(v []int16) {
	C.glTexCoord4sv((*C.GLshort)(&v[0]))
}

func TexCoordPointer(size int, typ GLenum, stride int, pointer interface{}) {
	C.glTexCoordPointer(C.GLint(size), C.GLenum(typ), C.GLsizei(stride), glPointer(pointer))
}

func TexEnvf(target GLenum, pname GLenum, param float32) {
	C.glTexEnvf(C.GLenum(target), C.GLenum(pname), C.GLfloat(param))
}

func TexEnvfv(target GLenum, pname GLenum, params []float32) {
	C.glTexEnvfv(C.GLenum(target), C.GLenum(pname), (*C.GLfloat)(&params[0]))
}

func TexEnvi(target GLenum, pname GLenum, param int) {
	C.glTexEnvi(C.GLenum(target), C.GLenum(pname), C.GLint(param))
}

func TexEnviv(target GLenum, pname GLenum, params []int32) {
	C.glTexEnviv(C.GLenum(target), C.GLenum(pname), (*C.GLint)(&params[0]))
}

func TexGend(coord GLenum, pname GLenum, param float64) {
	C.glTexGend(C.GLenum(coord), C.GLenum(pname), C.GLdouble(param))
}

func TexGendv(coord GLenum, pname GLenum, params []float64) {
	C.glTexGendv(C.GLenum(coord), C.GLenum(pname), (*C.GLdouble)(&params[0]))
}

func TexGenf(coord GLenum, pname GLenum, param float32) {
	C.glTexGenf(C.GLenum(coord), C.GLenum(pname), C.GLfloat(param))
}

func TexGenfv(coord GLenum, pname GLenum, params []float32) {
	C.glTexGenfv(C.GLenum(coord), C.GLenum(pname), (*C.GLfloat)(&params[0]))
}

func TexGeni(coord GLenum, pname GLenum, param int) {
	C.glTexGeni(C.GLenum(coord), C.GLenum(pname), C.GLint(param))
}

func TexGeniv(coord GLenum, pname GLenum, params []int32) {
	C.glTexGeniv(C.GLenum(coord), C.GLenum(pname), (*C.GLint)(&params[0]))
}

func TexImage1D(target GLenum, level int, internalformat int, width int, border int, format GLenum, typ GLenum, pixels interface{}) {
	C.glTexImage1D(C.GLenum(target), C.GLint(level), C.GLint(internalformat), C.GLsizei(width), C.GLint(border), C.GLenum(format), C.GLenum(typ), glPointer(pixels))
}

func TexImage2D(target GLenum, level int, internalformat int, width int, height int, border int, format GLenum, typ GLenum, pixels interface{}) {
	C.glTexImage2D(C.GLenum(target), C.GLint(level), C.GLint(internalformat), C.GLsizei(width), C.GLsizei(height), C.GLint(border), C.GLenum(format), C.GLenum(typ), glPointer(pixels))
}

func TexParameterf(target GLenum, pname GLenum, param float32) {
	C.glTexParameterf(C.GLenum(target), C.GLenum(pname), C.GLfloat(param))
}

func TexParameterfv(target GLenum, pname GLenum, params []float32) {
	C.glTexParameterfv(C.GLenum(target), C.GLenum(pname), (*C.GLfloat)(&params[0]))
}

func TexParameteri(target GLenum, pname GLenum, param int) {
	C.glTexParameteri(C.GLenum(target), C.GLenum(pname), C.GLint(param))
}

func TexParameteriv(target GLenum, pname GLenum, params []int32) {
	C.glTexParameteriv(C.GLenum(target), C.GLenum(pname), (*C.GLint)(&params[0]))
}

func TexSubImage1D(target GLenum, level int, xoffset int, width int, format GLenum, typ GLenum, pixels interface{}) {
	C.glTexSubImage1D(C.GLenum(target), C.GLint(level), C.GLint(xoffset), C.GLsizei(width), C.GLenum(format), C.GLenum(typ), glPointer(pixels))
}

func TexSubImage2D(target GLenum, level int, xoffset int, yoffset int, width int, height int, format GLenum, typ GLenum, pixels interface{}) {
	C.glTexSubImage2D(C.GLenum(target), C.GLint(level), C.GLint(xoffset), C.GLint(yoffset), C.GLsizei(width), C.GLsizei(height), C.GLenum(format), C.GLenum(typ), glPointer(pixels))
}

func Translated(x float64, y float64, z float64) {
	C.glTranslated(C.GLdouble(x), C.GLdouble(y), C.GLdouble(z))
}

func Translatef(x float32, y float32, z float32) {
	C.glTranslatef(C.GLfloat(x), C.GLfloat(y), C.GLfloat(z))
}

func Vertex2d(x float64, y float64) {
	C.glVertex2d(C.GLdouble(x), C.GLdouble(y))
}

func Vertex2dv(v []float64) {
	C.glVertex2dv((*C.GLdouble)(&v[0]))
}

func Vertex2f(x float32, y float32) {
	C.glVertex2f(C.GLfloat(x), C.GLfloat(y))
}

func Vertex2fv(v []float32) {
	C.glVertex2fv((*C.GLfloat)(&v[0]))
}

func Vertex2i(x int, y int) {
	C.glVertex2i(C.GLint(x), C.GLint(y))
}

func Vertex2iv(v []int32) {
	C.glVertex2iv((*C.GLint)(&v[0]))
}

func Vertex2s(x int16, y int16) {
	C.glVertex2s(C.GLshort(x), C.GLshort(y))
}

func Vertex2sv(v []int16) {
	C.glVertex2sv((*C.GLshort)(&v[0]))
}

func Vertex3d(x float64, y float64, z float64) {
	C.glVertex3d(C.GLdouble(x), C.GLdouble(y), C.GLdouble(z))
}

func Vertex3dv(v []float64) {
	C.glVertex3dv((*C.GLdouble)(&v[0]))
}

func Vertex3f(x float32, y float32, z float32) {
	C.glVertex3f(C.GLfloat(x), C.GLfloat(y), C.GLfloat(z))
}

func Vertex3fv(v []float32) {
	C.glVertex3fv((*C.GLfloat)(&v[0]))
}

func Vertex3i(x int, y int, z int) {
	C.glVertex3i(C.GLint(x), C.GLint(y), C.GLint(z))
}

func Vertex3iv(v []int32) {
	C.glVertex3iv((*C.GLint)(&v[0]))
}

func Vertex3s(x int16, y int16, z int16) {
	C.glVertex3s(C.GLshort(x), C.GLshort(y), C.GLshort(z))
}

func Vertex3sv(v []int16) {
	C.glVertex3sv((*C.GLshort)(&v[0]))
}

func Vertex4d(x float64, y float64, z float64, w float64) {
	C.glVertex4d(C.GLdouble(x), C.GLdouble(y), C.GLdouble(z), C.GLdouble(w))
}

func Vertex4dv(v []float64) {
	C.glVertex4dv((*C.GLdouble)(&v[0]))
}

func Vertex4f(x float32, y float32, z float32, w float32) {
	C.glVertex4f(C.GLfloat(x), C.GLfloat(y), C.GLfloat(z), C.GLfloat(w))
}

func Vertex4fv(v []float32) {
	C.glVertex4fv((*C.GLfloat)(&v[0]))
}

func Vertex4i(x int, y int, z int, w int) {
	C.glVertex4i(C.GLint(x), C.GLint(y), C.GLint(z), C.GLint(w))
}

func Vertex4iv(v []int32) {
	C.glVertex4iv((*C.GLint)(&v[0]))
}

func Vertex4s(x int16, y int16, z int16, w int16) {
	C.glVertex4s(C.GLshort(x), C.GLshort(y), C.GLshort(z), C.GLshort(w))
}

func Vertex4sv(v []int16) {
	C.glVertex4sv((*C.GLshort)(&v[0]))
}

func VertexPointer(size int, typ GLenum, stride int, pointer interface{}) {
	C.glVertexPointer(C.GLint(size), C.GLenum(typ), C.GLsizei(stride), glPointer(pointer))
}

func Viewport(x int, y int, width int, height int) {
	C.glViewport(C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height))
}
