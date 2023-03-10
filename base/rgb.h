#pragma once
#ifndef RGB_H
#define RGB_H

#include <stdlib.h> /* For abs() */
#include <math.h>
#include "inline_keywords.h" /* For H_INLINE */
#include <stdint.h>

/* #define MMRGB_IS_BGR (offsetof(MMRGBColor, red) > offsetof(MMRGBColor, blue)) */
#define MMRGB_IS_BGR 1

struct _MMRGBColor {
	uint8_t blue;
	uint8_t green;
	uint8_t red;
};
typedef struct _MMRGBColor MMRGBColor;

/* MMRGBHex is a hexadecimal color value*/
typedef uint32_t MMRGBHex;

#define MMRGBHEX_MIN 0x000000
#define MMRGBHEX_MAX 0xFFFFFF

/* Converts rgb color to hexadecimal value. */
#define RGB_TO_HEX(red, green, blue) (((red) << 16) | ((green) << 8) | (blue))

/* Convenience wrapper for MMRGBColors. */
H_INLINE MMRGBHex hexFromMMRGB(MMRGBColor rgb) {
	return RGB_TO_HEX(rgb.red, rgb.green, rgb.blue);
}

#define RED_FROM_HEX(hex) ((hex >> 16) & 0xFF)
#define GREEN_FROM_HEX(hex) ((hex >> 8) & 0xFF)
#define BLUE_FROM_HEX(hex) (hex & 0xFF)

/* Converts hexadecimal color to MMRGBColor. */
H_INLINE MMRGBColor MMRGBFromHex(MMRGBHex hex) {
	MMRGBColor color;
	color.red = RED_FROM_HEX(hex);
	color.green = GREEN_FROM_HEX(hex);
	color.blue = BLUE_FROM_HEX(hex);
	return color;
}

/* Check absolute equality of two RGB colors. */
#define MMRGBColorEqualToColor(c1, c2) ((c1).red == (c2).red && \
                                        (c1).blue == (c2).blue && \
                                        (c1).green == (c2).green)

/* Returns whether two colors are similar within the given range, |tolerance|.*/
H_INLINE int MMRGBColorSimilarToColor(MMRGBColor c1, MMRGBColor c2, float tolerance) {
	/* Speedy case */
	if (tolerance <= 0.0f) {
		return MMRGBColorEqualToColor(c1, c2);
	} else { /* Otherwise, use a Euclidean space to determine similarity */
		uint8_t d1 = c1.red - c2.red;
		uint8_t d2 = c1.green - c2.green;
		uint8_t d3 = c1.blue - c2.blue;
		return sqrt((double)(d1 * d1) + (d2 * d2) +
		            (d3 * d3)) <= (tolerance * 442.0f);
	}
}

/* Identical to MMRGBColorSimilarToColor, only for hex values. */
H_INLINE int MMRGBHexSimilarToColor(MMRGBHex h1, MMRGBHex h2, float tolerance) {
	if (tolerance <= 0.0f) {
		return h1 == h2;
	} else {
		int d1 = RED_FROM_HEX(h1) - RED_FROM_HEX(h2);
		int d2 = GREEN_FROM_HEX(h1) - GREEN_FROM_HEX(h2);
		int d3 = BLUE_FROM_HEX(h1) - BLUE_FROM_HEX(h2);
		return sqrt((double)(d1 * d1) + (d2 * d2) +
		            (d3 * d3)) <= (tolerance * 442.0f);
	}
}

#endif /* RGB_H */
