// Code generated from DML.g4 by ANTLR 4.9.1. DO NOT EDIT.

package parser // DML

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// package org.gaterace.dml;

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 77, 433,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4, 39, 9,
	39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44, 9, 44,
	3, 2, 3, 2, 3, 3, 3, 3, 5, 3, 93, 10, 3, 3, 3, 7, 3, 96, 10, 3, 12, 3,
	14, 3, 99, 11, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 7, 4, 106, 10, 4, 12, 4,
	14, 4, 109, 11, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 116, 10, 5, 3, 6,
	7, 6, 119, 10, 6, 12, 6, 14, 6, 122, 11, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3,
	6, 3, 6, 3, 7, 6, 7, 131, 10, 7, 13, 7, 14, 7, 132, 3, 8, 5, 8, 136, 10,
	8, 3, 8, 3, 8, 3, 9, 7, 9, 141, 10, 9, 12, 9, 14, 9, 144, 11, 9, 3, 9,
	3, 9, 3, 9, 5, 9, 149, 10, 9, 3, 9, 3, 9, 3, 10, 7, 10, 154, 10, 10, 12,
	10, 14, 10, 157, 11, 10, 3, 10, 3, 10, 3, 10, 5, 10, 162, 10, 10, 3, 10,
	3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 12, 6, 12, 172, 10, 12, 13,
	12, 14, 12, 173, 3, 13, 7, 13, 177, 10, 13, 12, 13, 14, 13, 180, 11, 13,
	3, 13, 3, 13, 3, 13, 3, 13, 5, 13, 186, 10, 13, 3, 13, 3, 13, 3, 14, 3,
	14, 3, 14, 7, 14, 193, 10, 14, 12, 14, 14, 14, 196, 11, 14, 3, 15, 3, 15,
	3, 15, 7, 15, 201, 10, 15, 12, 15, 14, 15, 204, 11, 15, 3, 16, 3, 16, 3,
	16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 5, 16,
	218, 10, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 5, 16, 225, 10, 16, 3,
	16, 3, 16, 3, 16, 3, 16, 3, 16, 5, 16, 232, 10, 16, 3, 16, 3, 16, 3, 16,
	3, 16, 3, 16, 5, 16, 239, 10, 16, 3, 16, 3, 16, 3, 16, 3, 16, 5, 16, 245,
	10, 16, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 7, 18,
	255, 10, 18, 12, 18, 14, 18, 258, 11, 18, 3, 18, 3, 18, 3, 18, 3, 19, 7,
	19, 264, 10, 19, 12, 19, 14, 19, 267, 11, 19, 3, 19, 3, 19, 3, 19, 3, 19,
	3, 19, 3, 19, 7, 19, 275, 10, 19, 12, 19, 14, 19, 278, 11, 19, 3, 19, 3,
	19, 3, 20, 7, 20, 283, 10, 20, 12, 20, 14, 20, 286, 11, 20, 3, 21, 3, 21,
	5, 21, 290, 10, 21, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 23, 3, 23, 3,
	23, 3, 23, 3, 23, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 25, 6, 25, 308,
	10, 25, 13, 25, 14, 25, 309, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 5, 26,
	317, 10, 26, 3, 27, 3, 27, 3, 27, 5, 27, 322, 10, 27, 3, 27, 3, 27, 3,
	28, 3, 28, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 5, 29, 335,
	10, 29, 3, 30, 7, 30, 338, 10, 30, 12, 30, 14, 30, 341, 11, 30, 3, 30,
	3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 31, 7, 31, 351, 10, 31, 12,
	31, 14, 31, 354, 11, 31, 3, 32, 3, 32, 5, 32, 358, 10, 32, 3, 33, 7, 33,
	361, 10, 33, 12, 33, 14, 33, 364, 11, 33, 3, 33, 3, 33, 3, 33, 5, 33, 369,
	10, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 34, 3, 34, 3, 35, 3, 35, 3, 35,
	3, 35, 3, 35, 3, 35, 3, 36, 6, 36, 384, 10, 36, 13, 36, 14, 36, 385, 3,
	37, 7, 37, 389, 10, 37, 12, 37, 14, 37, 392, 11, 37, 3, 37, 3, 37, 3, 37,
	5, 37, 397, 10, 37, 3, 37, 3, 37, 3, 37, 3, 37, 5, 37, 403, 10, 37, 3,
	38, 3, 38, 3, 39, 6, 39, 408, 10, 39, 13, 39, 14, 39, 409, 3, 40, 3, 40,
	3, 40, 3, 40, 3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 3, 42, 3, 42, 3, 42, 3,
	42, 3, 43, 3, 43, 3, 43, 3, 43, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 2, 2,
	45, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36,
	38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72,
	74, 76, 78, 80, 82, 84, 86, 2, 7, 3, 2, 60, 61, 4, 2, 59, 59, 69, 69, 3,
	2, 46, 48, 3, 2, 38, 40, 3, 2, 50, 53, 2, 450, 2, 88, 3, 2, 2, 2, 4, 90,
	3, 2, 2, 2, 6, 103, 3, 2, 2, 2, 8, 115, 3, 2, 2, 2, 10, 120, 3, 2, 2, 2,
	12, 130, 3, 2, 2, 2, 14, 135, 3, 2, 2, 2, 16, 142, 3, 2, 2, 2, 18, 155,
	3, 2, 2, 2, 20, 167, 3, 2, 2, 2, 22, 171, 3, 2, 2, 2, 24, 178, 3, 2, 2,
	2, 26, 189, 3, 2, 2, 2, 28, 197, 3, 2, 2, 2, 30, 244, 3, 2, 2, 2, 32, 246,
	3, 2, 2, 2, 34, 248, 3, 2, 2, 2, 36, 265, 3, 2, 2, 2, 38, 284, 3, 2, 2,
	2, 40, 289, 3, 2, 2, 2, 42, 291, 3, 2, 2, 2, 44, 296, 3, 2, 2, 2, 46, 301,
	3, 2, 2, 2, 48, 307, 3, 2, 2, 2, 50, 311, 3, 2, 2, 2, 52, 318, 3, 2, 2,
	2, 54, 325, 3, 2, 2, 2, 56, 334, 3, 2, 2, 2, 58, 339, 3, 2, 2, 2, 60, 352,
	3, 2, 2, 2, 62, 357, 3, 2, 2, 2, 64, 362, 3, 2, 2, 2, 66, 374, 3, 2, 2,
	2, 68, 376, 3, 2, 2, 2, 70, 383, 3, 2, 2, 2, 72, 390, 3, 2, 2, 2, 74, 404,
	3, 2, 2, 2, 76, 407, 3, 2, 2, 2, 78, 411, 3, 2, 2, 2, 80, 415, 3, 2, 2,
	2, 82, 420, 3, 2, 2, 2, 84, 424, 3, 2, 2, 2, 86, 428, 3, 2, 2, 2, 88, 89,
	5, 4, 3, 2, 89, 3, 3, 2, 2, 2, 90, 92, 5, 82, 42, 2, 91, 93, 5, 84, 43,
	2, 92, 91, 3, 2, 2, 2, 92, 93, 3, 2, 2, 2, 93, 97, 3, 2, 2, 2, 94, 96,
	5, 86, 44, 2, 95, 94, 3, 2, 2, 2, 96, 99, 3, 2, 2, 2, 97, 95, 3, 2, 2,
	2, 97, 98, 3, 2, 2, 2, 98, 100, 3, 2, 2, 2, 99, 97, 3, 2, 2, 2, 100, 101,
	5, 6, 4, 2, 101, 102, 7, 2, 2, 3, 102, 5, 3, 2, 2, 2, 103, 107, 5, 8, 5,
	2, 104, 106, 5, 8, 5, 2, 105, 104, 3, 2, 2, 2, 106, 109, 3, 2, 2, 2, 107,
	105, 3, 2, 2, 2, 107, 108, 3, 2, 2, 2, 108, 7, 3, 2, 2, 2, 109, 107, 3,
	2, 2, 2, 110, 116, 5, 10, 6, 2, 111, 116, 5, 18, 10, 2, 112, 116, 5, 36,
	19, 2, 113, 116, 5, 58, 30, 2, 114, 116, 5, 64, 33, 2, 115, 110, 3, 2,
	2, 2, 115, 111, 3, 2, 2, 2, 115, 112, 3, 2, 2, 2, 115, 113, 3, 2, 2, 2,
	115, 114, 3, 2, 2, 2, 116, 9, 3, 2, 2, 2, 117, 119, 7, 76, 2, 2, 118, 117,
	3, 2, 2, 2, 119, 122, 3, 2, 2, 2, 120, 118, 3, 2, 2, 2, 120, 121, 3, 2,
	2, 2, 121, 123, 3, 2, 2, 2, 122, 120, 3, 2, 2, 2, 123, 124, 7, 9, 2, 2,
	124, 125, 7, 74, 2, 2, 125, 126, 7, 66, 2, 2, 126, 127, 5, 12, 7, 2, 127,
	128, 7, 67, 2, 2, 128, 11, 3, 2, 2, 2, 129, 131, 5, 16, 9, 2, 130, 129,
	3, 2, 2, 2, 131, 132, 3, 2, 2, 2, 132, 130, 3, 2, 2, 2, 132, 133, 3, 2,
	2, 2, 133, 13, 3, 2, 2, 2, 134, 136, 9, 2, 2, 2, 135, 134, 3, 2, 2, 2,
	135, 136, 3, 2, 2, 2, 136, 137, 3, 2, 2, 2, 137, 138, 7, 55, 2, 2, 138,
	15, 3, 2, 2, 2, 139, 141, 7, 76, 2, 2, 140, 139, 3, 2, 2, 2, 141, 144,
	3, 2, 2, 2, 142, 140, 3, 2, 2, 2, 142, 143, 3, 2, 2, 2, 143, 145, 3, 2,
	2, 2, 144, 142, 3, 2, 2, 2, 145, 148, 7, 74, 2, 2, 146, 147, 7, 68, 2,
	2, 147, 149, 5, 14, 8, 2, 148, 146, 3, 2, 2, 2, 148, 149, 3, 2, 2, 2, 149,
	150, 3, 2, 2, 2, 150, 151, 7, 56, 2, 2, 151, 17, 3, 2, 2, 2, 152, 154,
	7, 76, 2, 2, 153, 152, 3, 2, 2, 2, 154, 157, 3, 2, 2, 2, 155, 153, 3, 2,
	2, 2, 155, 156, 3, 2, 2, 2, 156, 158, 3, 2, 2, 2, 157, 155, 3, 2, 2, 2,
	158, 159, 7, 13, 2, 2, 159, 161, 7, 74, 2, 2, 160, 162, 5, 20, 11, 2, 161,
	160, 3, 2, 2, 2, 161, 162, 3, 2, 2, 2, 162, 163, 3, 2, 2, 2, 163, 164,
	7, 66, 2, 2, 164, 165, 5, 22, 12, 2, 165, 166, 7, 67, 2, 2, 166, 19, 3,
	2, 2, 2, 167, 168, 7, 54, 2, 2, 168, 169, 7, 74, 2, 2, 169, 21, 3, 2, 2,
	2, 170, 172, 5, 24, 13, 2, 171, 170, 3, 2, 2, 2, 172, 173, 3, 2, 2, 2,
	173, 171, 3, 2, 2, 2, 173, 174, 3, 2, 2, 2, 174, 23, 3, 2, 2, 2, 175, 177,
	7, 76, 2, 2, 176, 175, 3, 2, 2, 2, 177, 180, 3, 2, 2, 2, 178, 176, 3, 2,
	2, 2, 178, 179, 3, 2, 2, 2, 179, 181, 3, 2, 2, 2, 180, 178, 3, 2, 2, 2,
	181, 182, 5, 30, 16, 2, 182, 185, 7, 74, 2, 2, 183, 184, 7, 68, 2, 2, 184,
	186, 5, 14, 8, 2, 185, 183, 3, 2, 2, 2, 185, 186, 3, 2, 2, 2, 186, 187,
	3, 2, 2, 2, 187, 188, 7, 56, 2, 2, 188, 25, 3, 2, 2, 2, 189, 194, 7, 74,
	2, 2, 190, 191, 7, 59, 2, 2, 191, 193, 7, 74, 2, 2, 192, 190, 3, 2, 2,
	2, 193, 196, 3, 2, 2, 2, 194, 192, 3, 2, 2, 2, 194, 195, 3, 2, 2, 2, 195,
	27, 3, 2, 2, 2, 196, 194, 3, 2, 2, 2, 197, 202, 7, 74, 2, 2, 198, 199,
	9, 3, 2, 2, 199, 201, 7, 74, 2, 2, 200, 198, 3, 2, 2, 2, 201, 204, 3, 2,
	2, 2, 202, 200, 3, 2, 2, 2, 202, 203, 3, 2, 2, 2, 203, 29, 3, 2, 2, 2,
	204, 202, 3, 2, 2, 2, 205, 245, 7, 21, 2, 2, 206, 245, 7, 22, 2, 2, 207,
	245, 7, 23, 2, 2, 208, 245, 7, 24, 2, 2, 209, 245, 7, 25, 2, 2, 210, 245,
	7, 26, 2, 2, 211, 245, 7, 27, 2, 2, 212, 217, 7, 28, 2, 2, 213, 214, 7,
	63, 2, 2, 214, 215, 5, 14, 8, 2, 215, 216, 7, 62, 2, 2, 216, 218, 3, 2,
	2, 2, 217, 213, 3, 2, 2, 2, 217, 218, 3, 2, 2, 2, 218, 245, 3, 2, 2, 2,
	219, 224, 7, 29, 2, 2, 220, 221, 7, 63, 2, 2, 221, 222, 5, 14, 8, 2, 222,
	223, 7, 62, 2, 2, 223, 225, 3, 2, 2, 2, 224, 220, 3, 2, 2, 2, 224, 225,
	3, 2, 2, 2, 225, 245, 3, 2, 2, 2, 226, 231, 7, 30, 2, 2, 227, 228, 7, 63,
	2, 2, 228, 229, 5, 14, 8, 2, 229, 230, 7, 62, 2, 2, 230, 232, 3, 2, 2,
	2, 231, 227, 3, 2, 2, 2, 231, 232, 3, 2, 2, 2, 232, 245, 3, 2, 2, 2, 233,
	238, 7, 31, 2, 2, 234, 235, 7, 63, 2, 2, 235, 236, 5, 14, 8, 2, 236, 237,
	7, 62, 2, 2, 237, 239, 3, 2, 2, 2, 238, 234, 3, 2, 2, 2, 238, 239, 3, 2,
	2, 2, 239, 245, 3, 2, 2, 2, 240, 245, 7, 32, 2, 2, 241, 245, 7, 33, 2,
	2, 242, 245, 7, 34, 2, 2, 243, 245, 5, 26, 14, 2, 244, 205, 3, 2, 2, 2,
	244, 206, 3, 2, 2, 2, 244, 207, 3, 2, 2, 2, 244, 208, 3, 2, 2, 2, 244,
	209, 3, 2, 2, 2, 244, 210, 3, 2, 2, 2, 244, 211, 3, 2, 2, 2, 244, 212,
	3, 2, 2, 2, 244, 219, 3, 2, 2, 2, 244, 226, 3, 2, 2, 2, 244, 233, 3, 2,
	2, 2, 244, 240, 3, 2, 2, 2, 244, 241, 3, 2, 2, 2, 244, 242, 3, 2, 2, 2,
	244, 243, 3, 2, 2, 2, 245, 31, 3, 2, 2, 2, 246, 247, 9, 4, 2, 2, 247, 33,
	3, 2, 2, 2, 248, 249, 7, 45, 2, 2, 249, 250, 5, 32, 17, 2, 250, 251, 7,
	63, 2, 2, 251, 256, 7, 74, 2, 2, 252, 253, 7, 57, 2, 2, 253, 255, 7, 74,
	2, 2, 254, 252, 3, 2, 2, 2, 255, 258, 3, 2, 2, 2, 256, 254, 3, 2, 2, 2,
	256, 257, 3, 2, 2, 2, 257, 259, 3, 2, 2, 2, 258, 256, 3, 2, 2, 2, 259,
	260, 7, 62, 2, 2, 260, 261, 7, 56, 2, 2, 261, 35, 3, 2, 2, 2, 262, 264,
	7, 76, 2, 2, 263, 262, 3, 2, 2, 2, 264, 267, 3, 2, 2, 2, 265, 263, 3, 2,
	2, 2, 265, 266, 3, 2, 2, 2, 266, 268, 3, 2, 2, 2, 267, 265, 3, 2, 2, 2,
	268, 269, 7, 14, 2, 2, 269, 270, 7, 74, 2, 2, 270, 271, 5, 38, 20, 2, 271,
	272, 7, 66, 2, 2, 272, 276, 5, 48, 25, 2, 273, 275, 5, 34, 18, 2, 274,
	273, 3, 2, 2, 2, 275, 278, 3, 2, 2, 2, 276, 274, 3, 2, 2, 2, 276, 277,
	3, 2, 2, 2, 277, 279, 3, 2, 2, 2, 278, 276, 3, 2, 2, 2, 279, 280, 7, 67,
	2, 2, 280, 37, 3, 2, 2, 2, 281, 283, 5, 40, 21, 2, 282, 281, 3, 2, 2, 2,
	283, 286, 3, 2, 2, 2, 284, 282, 3, 2, 2, 2, 284, 285, 3, 2, 2, 2, 285,
	39, 3, 2, 2, 2, 286, 284, 3, 2, 2, 2, 287, 290, 5, 42, 22, 2, 288, 290,
	5, 44, 23, 2, 289, 287, 3, 2, 2, 2, 289, 288, 3, 2, 2, 2, 290, 41, 3, 2,
	2, 2, 291, 292, 7, 13, 2, 2, 292, 293, 7, 63, 2, 2, 293, 294, 7, 74, 2,
	2, 294, 295, 7, 62, 2, 2, 295, 43, 3, 2, 2, 2, 296, 297, 7, 15, 2, 2, 297,
	298, 7, 63, 2, 2, 298, 299, 7, 74, 2, 2, 299, 300, 7, 62, 2, 2, 300, 45,
	3, 2, 2, 2, 301, 302, 7, 18, 2, 2, 302, 303, 7, 63, 2, 2, 303, 304, 7,
	74, 2, 2, 304, 305, 7, 62, 2, 2, 305, 47, 3, 2, 2, 2, 306, 308, 5, 52,
	27, 2, 307, 306, 3, 2, 2, 2, 308, 309, 3, 2, 2, 2, 309, 307, 3, 2, 2, 2,
	309, 310, 3, 2, 2, 2, 310, 49, 3, 2, 2, 2, 311, 312, 7, 44, 2, 2, 312,
	316, 7, 74, 2, 2, 313, 314, 7, 63, 2, 2, 314, 315, 7, 74, 2, 2, 315, 317,
	7, 62, 2, 2, 316, 313, 3, 2, 2, 2, 316, 317, 3, 2, 2, 2, 317, 51, 3, 2,
	2, 2, 318, 319, 5, 56, 29, 2, 319, 321, 7, 74, 2, 2, 320, 322, 5, 50, 26,
	2, 321, 320, 3, 2, 2, 2, 321, 322, 3, 2, 2, 2, 322, 323, 3, 2, 2, 2, 323,
	324, 7, 56, 2, 2, 324, 53, 3, 2, 2, 2, 325, 326, 9, 5, 2, 2, 326, 55, 3,
	2, 2, 2, 327, 335, 7, 35, 2, 2, 328, 335, 7, 36, 2, 2, 329, 335, 7, 37,
	2, 2, 330, 335, 7, 41, 2, 2, 331, 335, 7, 42, 2, 2, 332, 335, 7, 43, 2,
	2, 333, 335, 5, 54, 28, 2, 334, 327, 3, 2, 2, 2, 334, 328, 3, 2, 2, 2,
	334, 329, 3, 2, 2, 2, 334, 330, 3, 2, 2, 2, 334, 331, 3, 2, 2, 2, 334,
	332, 3, 2, 2, 2, 334, 333, 3, 2, 2, 2, 335, 57, 3, 2, 2, 2, 336, 338, 7,
	76, 2, 2, 337, 336, 3, 2, 2, 2, 338, 341, 3, 2, 2, 2, 339, 337, 3, 2, 2,
	2, 339, 340, 3, 2, 2, 2, 340, 342, 3, 2, 2, 2, 341, 339, 3, 2, 2, 2, 342,
	343, 7, 16, 2, 2, 343, 344, 7, 74, 2, 2, 344, 345, 5, 60, 31, 2, 345, 346,
	7, 66, 2, 2, 346, 347, 5, 48, 25, 2, 347, 348, 7, 67, 2, 2, 348, 59, 3,
	2, 2, 2, 349, 351, 5, 62, 32, 2, 350, 349, 3, 2, 2, 2, 351, 354, 3, 2,
	2, 2, 352, 350, 3, 2, 2, 2, 352, 353, 3, 2, 2, 2, 353, 61, 3, 2, 2, 2,
	354, 352, 3, 2, 2, 2, 355, 358, 5, 42, 22, 2, 356, 358, 5, 44, 23, 2, 357,
	355, 3, 2, 2, 2, 357, 356, 3, 2, 2, 2, 358, 63, 3, 2, 2, 2, 359, 361, 7,
	76, 2, 2, 360, 359, 3, 2, 2, 2, 361, 364, 3, 2, 2, 2, 362, 360, 3, 2, 2,
	2, 362, 363, 3, 2, 2, 2, 363, 365, 3, 2, 2, 2, 364, 362, 3, 2, 2, 2, 365,
	366, 7, 17, 2, 2, 366, 368, 7, 74, 2, 2, 367, 369, 5, 42, 22, 2, 368, 367,
	3, 2, 2, 2, 368, 369, 3, 2, 2, 2, 369, 370, 3, 2, 2, 2, 370, 371, 7, 66,
	2, 2, 371, 372, 5, 70, 36, 2, 372, 373, 7, 67, 2, 2, 373, 65, 3, 2, 2,
	2, 374, 375, 9, 6, 2, 2, 375, 67, 3, 2, 2, 2, 376, 377, 7, 49, 2, 2, 377,
	378, 7, 63, 2, 2, 378, 379, 5, 66, 34, 2, 379, 380, 5, 26, 14, 2, 380,
	381, 7, 62, 2, 2, 381, 69, 3, 2, 2, 2, 382, 384, 5, 72, 37, 2, 383, 382,
	3, 2, 2, 2, 384, 385, 3, 2, 2, 2, 385, 383, 3, 2, 2, 2, 385, 386, 3, 2,
	2, 2, 386, 71, 3, 2, 2, 2, 387, 389, 7, 76, 2, 2, 388, 387, 3, 2, 2, 2,
	389, 392, 3, 2, 2, 2, 390, 388, 3, 2, 2, 2, 390, 391, 3, 2, 2, 2, 391,
	393, 3, 2, 2, 2, 392, 390, 3, 2, 2, 2, 393, 394, 7, 19, 2, 2, 394, 396,
	7, 74, 2, 2, 395, 397, 5, 68, 35, 2, 396, 395, 3, 2, 2, 2, 396, 397, 3,
	2, 2, 2, 397, 398, 3, 2, 2, 2, 398, 399, 7, 66, 2, 2, 399, 400, 5, 74,
	38, 2, 400, 402, 7, 67, 2, 2, 401, 403, 5, 80, 41, 2, 402, 401, 3, 2, 2,
	2, 402, 403, 3, 2, 2, 2, 403, 73, 3, 2, 2, 2, 404, 405, 5, 76, 39, 2, 405,
	75, 3, 2, 2, 2, 406, 408, 5, 78, 40, 2, 407, 406, 3, 2, 2, 2, 408, 409,
	3, 2, 2, 2, 409, 407, 3, 2, 2, 2, 409, 410, 3, 2, 2, 2, 410, 77, 3, 2,
	2, 2, 411, 412, 5, 54, 28, 2, 412, 413, 7, 74, 2, 2, 413, 414, 7, 56, 2,
	2, 414, 79, 3, 2, 2, 2, 415, 416, 7, 20, 2, 2, 416, 417, 7, 66, 2, 2, 417,
	418, 5, 76, 39, 2, 418, 419, 7, 67, 2, 2, 419, 81, 3, 2, 2, 2, 420, 421,
	7, 10, 2, 2, 421, 422, 5, 26, 14, 2, 422, 423, 7, 56, 2, 2, 423, 83, 3,
	2, 2, 2, 424, 425, 7, 11, 2, 2, 425, 426, 5, 28, 15, 2, 426, 427, 7, 56,
	2, 2, 427, 85, 3, 2, 2, 2, 428, 429, 7, 12, 2, 2, 429, 430, 7, 73, 2, 2,
	430, 431, 7, 56, 2, 2, 431, 87, 3, 2, 2, 2, 42, 92, 97, 107, 115, 120,
	132, 135, 142, 148, 155, 161, 173, 178, 185, 194, 202, 217, 224, 231, 238,
	244, 256, 265, 276, 284, 289, 309, 316, 321, 334, 339, 352, 357, 362, 368,
	385, 390, 396, 402, 409,
}
var literalNames = []string{
	"", "'dml'", "'dmlfields'", "'dmlfield'", "'dmlinputs'", "'dmloutputs'",
	"'dmlqualid'", "'enum'", "'package'", "'gopackage'", "'import'", "'fieldset'",
	"'class'", "'table'", "'association'", "'service'", "'mediates'", "'method'",
	"'returns'", "'double'", "'float'", "'int32'", "'int64'", "'uint32'", "'uint64'",
	"'bool'", "'string'", "'bytes'", "'chararray'", "'bytearray'", "'datetime'",
	"'guid'", "'decimal'", "'auto'", "'id'", "'idgen'", "'required'", "'optional'",
	"'repeated'", "'virtual'", "'map'", "'list'", "'references'", "'index'",
	"'primary'", "'unique'", "'nonunique'", "'pattern'", "'insert'", "'update'",
	"'select'", "'delete'", "'extends'", "", "';'", "','", "':'", "'.'", "'+'",
	"'-'", "')'", "'('", "']'", "'['", "'{'", "'}'", "'='", "'/'",
}
var symbolicNames = []string{
	"", "DML", "DMLFIELDS", "DMLFIELD", "DMLINPUTS", "DMLOUTPUTS", "DMLQUALID",
	"ENUM", "PACKAGE", "GOPACKAGE", "IMPORT", "FIELDSET", "CLASS", "TABLE",
	"ASSOCIATION", "SERVICE", "MEDIATES", "METHOD", "RETURNS", "DOUBLE", "FLOAT",
	"INT32", "INT64", "UINT32", "UINT64", "BOOL", "STRING", "BYTES", "CHARARRAY",
	"BYTEARRAY", "DATETIME", "GUID", "DECIMAL", "AUTO", "ID_KEYWORD", "IDGEN",
	"REQUIRED", "OPTIONAL", "REPEATED", "VIRTUAL", "MAP", "LIST", "REFERENCES",
	"INDEX", "PRIMARY", "UNIQUE", "NONUNIQUE", "PATTERN", "INSERT", "UPDATE",
	"SELECT", "DELETE", "EXTENDS", "INTEGER_NUM", "SEMI", "COMMA", "COLON",
	"DOT", "PLUS", "MINUS", "RPAREN", "LPAREN", "RBRACK", "LBRACK", "LBRACE",
	"RBRACE", "EQUALS", "SLASH", "WHITE_SPACE", "HEX_DIGIT", "REAL_NUMBER",
	"TEXT_STRING", "ID", "COMMENT", "KEEP_COMMENT", "LINE_COMMENT",
}

var ruleNames = []string{
	"model", "modelParts", "sourceElements", "sourceElement", "enumElement",
	"enumValues", "integerConst", "enumValue", "fieldsetElement", "extendsQualifier",
	"fieldsetValues", "fieldsetValue", "qualifiedId", "goPath", "fieldType",
	"indexType", "indexElement", "classElement", "classOptions", "classOption",
	"fieldsetOption", "tableOption", "mediatesOption", "classFields", "referencesModifier",
	"classField", "parameterModifier", "fieldModifier", "associationElement",
	"associationOptions", "associationOption", "serviceElement", "patternType",
	"methodPattern", "methods", "method", "inputList", "fieldList", "fieldInstance",
	"methodReturn", "packageElement", "goPackageElement", "importElements",
}

type DMLParser struct {
	*antlr.BaseParser
}

// NewDMLParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *DMLParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewDMLParser(input antlr.TokenStream) *DMLParser {
	this := new(DMLParser)
	deserializer := antlr.NewATNDeserializer(nil)
	deserializedATN := deserializer.DeserializeFromUInt16(parserATN)
	decisionToDFA := make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "DML.g4"

	return this
}

// DMLParser tokens.
const (
	DMLParserEOF          = antlr.TokenEOF
	DMLParserDML          = 1
	DMLParserDMLFIELDS    = 2
	DMLParserDMLFIELD     = 3
	DMLParserDMLINPUTS    = 4
	DMLParserDMLOUTPUTS   = 5
	DMLParserDMLQUALID    = 6
	DMLParserENUM         = 7
	DMLParserPACKAGE      = 8
	DMLParserGOPACKAGE    = 9
	DMLParserIMPORT       = 10
	DMLParserFIELDSET     = 11
	DMLParserCLASS        = 12
	DMLParserTABLE        = 13
	DMLParserASSOCIATION  = 14
	DMLParserSERVICE      = 15
	DMLParserMEDIATES     = 16
	DMLParserMETHOD       = 17
	DMLParserRETURNS      = 18
	DMLParserDOUBLE       = 19
	DMLParserFLOAT        = 20
	DMLParserINT32        = 21
	DMLParserINT64        = 22
	DMLParserUINT32       = 23
	DMLParserUINT64       = 24
	DMLParserBOOL         = 25
	DMLParserSTRING       = 26
	DMLParserBYTES        = 27
	DMLParserCHARARRAY    = 28
	DMLParserBYTEARRAY    = 29
	DMLParserDATETIME     = 30
	DMLParserGUID         = 31
	DMLParserDECIMAL      = 32
	DMLParserAUTO         = 33
	DMLParserID_KEYWORD   = 34
	DMLParserIDGEN        = 35
	DMLParserREQUIRED     = 36
	DMLParserOPTIONAL     = 37
	DMLParserREPEATED     = 38
	DMLParserVIRTUAL      = 39
	DMLParserMAP          = 40
	DMLParserLIST         = 41
	DMLParserREFERENCES   = 42
	DMLParserINDEX        = 43
	DMLParserPRIMARY      = 44
	DMLParserUNIQUE       = 45
	DMLParserNONUNIQUE    = 46
	DMLParserPATTERN      = 47
	DMLParserINSERT       = 48
	DMLParserUPDATE       = 49
	DMLParserSELECT       = 50
	DMLParserDELETE       = 51
	DMLParserEXTENDS      = 52
	DMLParserINTEGER_NUM  = 53
	DMLParserSEMI         = 54
	DMLParserCOMMA        = 55
	DMLParserCOLON        = 56
	DMLParserDOT          = 57
	DMLParserPLUS         = 58
	DMLParserMINUS        = 59
	DMLParserRPAREN       = 60
	DMLParserLPAREN       = 61
	DMLParserRBRACK       = 62
	DMLParserLBRACK       = 63
	DMLParserLBRACE       = 64
	DMLParserRBRACE       = 65
	DMLParserEQUALS       = 66
	DMLParserSLASH        = 67
	DMLParserWHITE_SPACE  = 68
	DMLParserHEX_DIGIT    = 69
	DMLParserREAL_NUMBER  = 70
	DMLParserTEXT_STRING  = 71
	DMLParserID           = 72
	DMLParserCOMMENT      = 73
	DMLParserKEEP_COMMENT = 74
	DMLParserLINE_COMMENT = 75
)

// DMLParser rules.
const (
	DMLParserRULE_model              = 0
	DMLParserRULE_modelParts         = 1
	DMLParserRULE_sourceElements     = 2
	DMLParserRULE_sourceElement      = 3
	DMLParserRULE_enumElement        = 4
	DMLParserRULE_enumValues         = 5
	DMLParserRULE_integerConst       = 6
	DMLParserRULE_enumValue          = 7
	DMLParserRULE_fieldsetElement    = 8
	DMLParserRULE_extendsQualifier   = 9
	DMLParserRULE_fieldsetValues     = 10
	DMLParserRULE_fieldsetValue      = 11
	DMLParserRULE_qualifiedId        = 12
	DMLParserRULE_goPath             = 13
	DMLParserRULE_fieldType          = 14
	DMLParserRULE_indexType          = 15
	DMLParserRULE_indexElement       = 16
	DMLParserRULE_classElement       = 17
	DMLParserRULE_classOptions       = 18
	DMLParserRULE_classOption        = 19
	DMLParserRULE_fieldsetOption     = 20
	DMLParserRULE_tableOption        = 21
	DMLParserRULE_mediatesOption     = 22
	DMLParserRULE_classFields        = 23
	DMLParserRULE_referencesModifier = 24
	DMLParserRULE_classField         = 25
	DMLParserRULE_parameterModifier  = 26
	DMLParserRULE_fieldModifier      = 27
	DMLParserRULE_associationElement = 28
	DMLParserRULE_associationOptions = 29
	DMLParserRULE_associationOption  = 30
	DMLParserRULE_serviceElement     = 31
	DMLParserRULE_patternType        = 32
	DMLParserRULE_methodPattern      = 33
	DMLParserRULE_methods            = 34
	DMLParserRULE_method             = 35
	DMLParserRULE_inputList          = 36
	DMLParserRULE_fieldList          = 37
	DMLParserRULE_fieldInstance      = 38
	DMLParserRULE_methodReturn       = 39
	DMLParserRULE_packageElement     = 40
	DMLParserRULE_goPackageElement   = 41
	DMLParserRULE_importElements     = 42
)

// IModelContext is an interface to support dynamic dispatch.
type IModelContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsModelContext differentiates from other interfaces.
	IsModelContext()
}

type ModelContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyModelContext() *ModelContext {
	var p = new(ModelContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DMLParserRULE_model
	return p
}

func (*ModelContext) IsModelContext() {}

func NewModelContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ModelContext {
	var p = new(ModelContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DMLParserRULE_model

	return p
}

func (s *ModelContext) GetParser() antlr.Parser { return s.parser }

func (s *ModelContext) ModelParts() IModelPartsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IModelPartsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IModelPartsContext)
}

func (s *ModelContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ModelContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ModelContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DMLListener); ok {
		listenerT.EnterModel(s)
	}
}

func (s *ModelContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DMLListener); ok {
		listenerT.ExitModel(s)
	}
}

func (p *DMLParser) Model() (localctx IModelContext) {
	localctx = NewModelContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, DMLParserRULE_model)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(86)
		p.ModelParts()
	}

	return localctx
}

// IModelPartsContext is an interface to support dynamic dispatch.
type IModelPartsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsModelPartsContext differentiates from other interfaces.
	IsModelPartsContext()
}

type ModelPartsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyModelPartsContext() *ModelPartsContext {
	var p = new(ModelPartsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DMLParserRULE_modelParts
	return p
}

func (*ModelPartsContext) IsModelPartsContext() {}

func NewModelPartsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ModelPartsContext {
	var p = new(ModelPartsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DMLParserRULE_modelParts

	return p
}

func (s *ModelPartsContext) GetParser() antlr.Parser { return s.parser }

func (s *ModelPartsContext) PackageElement() IPackageElementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPackageElementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPackageElementContext)
}

func (s *ModelPartsContext) SourceElements() ISourceElementsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISourceElementsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISourceElementsContext)
}

func (s *ModelPartsContext) EOF() antlr.TerminalNode {
	return s.GetToken(DMLParserEOF, 0)
}

func (s *ModelPartsContext) GoPackageElement() IGoPackageElementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGoPackageElementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGoPackageElementContext)
}

func (s *ModelPartsContext) AllImportElements() []IImportElementsContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IImportElementsContext)(nil)).Elem())
	var tst = make([]IImportElementsContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IImportElementsContext)
		}
	}

	return tst
}

func (s *ModelPartsContext) ImportElements(i int) IImportElementsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IImportElementsContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IImportElementsContext)
}

func (s *ModelPartsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ModelPartsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ModelPartsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DMLListener); ok {
		listenerT.EnterModelParts(s)
	}
}

func (s *ModelPartsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DMLListener); ok {
		listenerT.ExitModelParts(s)
	}
}

func (p *DMLParser) ModelParts() (localctx IModelPartsContext) {
	localctx = NewModelPartsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, DMLParserRULE_modelParts)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(88)
		p.PackageElement()
	}
	p.SetState(90)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == DMLParserGOPACKAGE {
		{
			p.SetState(89)
			p.GoPackageElement()
		}

	}
	p.SetState(95)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == DMLParserIMPORT {
		{
			p.SetState(92)
			p.ImportElements()
		}

		p.SetState(97)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(98)
		p.SourceElements()
	}
	{
		p.SetState(99)
		p.Match(DMLParserEOF)
	}

	return localctx
}

// ISourceElementsContext is an interface to support dynamic dispatch.
type ISourceElementsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSourceElementsContext differentiates from other interfaces.
	IsSourceElementsContext()
}

type SourceElementsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySourceElementsContext() *SourceElementsContext {
	var p = new(SourceElementsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DMLParserRULE_sourceElements
	return p
}

func (*SourceElementsContext) IsSourceElementsContext() {}

func NewSourceElementsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SourceElementsContext {
	var p = new(SourceElementsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DMLParserRULE_sourceElements

	return p
}

func (s *SourceElementsContext) GetParser() antlr.Parser { return s.parser }

func (s *SourceElementsContext) AllSourceElement() []ISourceElementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISourceElementContext)(nil)).Elem())
	var tst = make([]ISourceElementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISourceElementContext)
		}
	}

	return tst
}

func (s *SourceElementsContext) SourceElement(i int) ISourceElementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISourceElementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISourceElementContext)
}

func (s *SourceElementsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SourceElementsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SourceElementsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DMLListener); ok {
		listenerT.EnterSourceElements(s)
	}
}

func (s *SourceElementsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DMLListener); ok {
		listenerT.ExitSourceElements(s)
	}
}

func (p *DMLParser) SourceElements() (localctx ISourceElementsContext) {
	localctx = NewSourceElementsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, DMLParserRULE_sourceElements)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(101)
		p.SourceElement()
	}
	p.SetState(105)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<DMLParserENUM)|(1<<DMLParserFIELDSET)|(1<<DMLParserCLASS)|(1<<DMLParserASSOCIATION)|(1<<DMLParserSERVICE))) != 0) || _la == DMLParserKEEP_COMMENT {
		{
			p.SetState(102)
			p.SourceElement()
		}

		p.SetState(107)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ISourceElementContext is an interface to support dynamic dispatch.
type ISourceElementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSourceElementContext differentiates from other interfaces.
	IsSourceElementContext()
}

type SourceElementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySourceElementContext() *SourceElementContext {
	var p = new(SourceElementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DMLParserRULE_sourceElement
	return p
}

func (*SourceElementContext) IsSourceElementContext() {}

func NewSourceElementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SourceElementContext {
	var p = new(SourceElementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DMLParserRULE_sourceElement

	return p
}

func (s *SourceElementContext) GetParser() antlr.Parser { return s.parser }

func (s *SourceElementContext) EnumElement() IEnumElementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnumElementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEnumElementContext)
}

func (s *SourceElementContext) FieldsetElement() IFieldsetElementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldsetElementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldsetElementContext)
}

func (s *SourceElementContext) ClassElement() IClassElementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IClassElementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IClassElementContext)
}

func (s *SourceElementContext) AssociationElement() IAssociationElementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssociationElementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssociationElementContext)
}

func (s *SourceElementContext) ServiceElement() IServiceElementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IServiceElementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IServiceElementContext)
}

func (s *SourceElementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SourceElementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SourceElementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DMLListener); ok {
		listenerT.EnterSourceElement(s)
	}
}

func (s *SourceElementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DMLListener); ok {
		listenerT.ExitSourceElement(s)
	}
}

func (p *DMLParser) SourceElement() (localctx ISourceElementContext) {
	localctx = NewSourceElementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, DMLParserRULE_sourceElement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(113)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(108)
			p.EnumElement()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(109)
			p.FieldsetElement()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(110)
			p.ClassElement()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(111)
			p.AssociationElement()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(112)
			p.ServiceElement()
		}

	}

	return localctx
}

// IEnumElementContext is an interface to support dynamic dispatch.
type IEnumElementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEnumElementContext differentiates from other interfaces.
	IsEnumElementContext()
}

type EnumElementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumElementContext() *EnumElementContext {
	var p = new(EnumElementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DMLParserRULE_enumElement
	return p
}

func (*EnumElementContext) IsEnumElementContext() {}

func NewEnumElementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumElementContext {
	var p = new(EnumElementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DMLParserRULE_enumElement

	return p
}

func (s *EnumElementContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumElementContext) ENUM() antlr.TerminalNode {
	return s.GetToken(DMLParserENUM, 0)
}

func (s *EnumElementContext) ID() antlr.TerminalNode {
	return s.GetToken(DMLParserID, 0)
}

func (s *EnumElementContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(DMLParserLBRACE, 0)
}

func (s *EnumElementContext) EnumValues() IEnumValuesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnumValuesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEnumValuesContext)
}

func (s *EnumElementContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(DMLParserRBRACE, 0)
}

func (s *EnumElementContext) AllKEEP_COMMENT() []antlr.TerminalNode {
	return s.GetTokens(DMLParserKEEP_COMMENT)
}

func (s *EnumElementContext) KEEP_COMMENT(i int) antlr.TerminalNode {
	return s.GetToken(DMLParserKEEP_COMMENT, i)
}

func (s *EnumElementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumElementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumElementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DMLListener); ok {
		listenerT.EnterEnumElement(s)
	}
}

func (s *EnumElementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DMLListener); ok {
		listenerT.ExitEnumElement(s)
	}
}

func (p *DMLParser) EnumElement() (localctx IEnumElementContext) {
	localctx = NewEnumElementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, DMLParserRULE_enumElement)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(118)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == DMLParserKEEP_COMMENT {
		{
			p.SetState(115)
			p.Match(DMLParserKEEP_COMMENT)
		}

		p.SetState(120)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(121)
		p.Match(DMLParserENUM)
	}
	{
		p.SetState(122)
		p.Match(DMLParserID)
	}
	{
		p.SetState(123)
		p.Match(DMLParserLBRACE)
	}
	{
		p.SetState(124)
		p.EnumValues()
	}
	{
		p.SetState(125)
		p.Match(DMLParserRBRACE)
	}

	return localctx
}

// IEnumValuesContext is an interface to support dynamic dispatch.
type IEnumValuesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEnumValuesContext differentiates from other interfaces.
	IsEnumValuesContext()
}

type EnumValuesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumValuesContext() *EnumValuesContext {
	var p = new(EnumValuesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DMLParserRULE_enumValues
	return p
}

func (*EnumValuesContext) IsEnumValuesContext() {}

func NewEnumValuesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumValuesContext {
	var p = new(EnumValuesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DMLParserRULE_enumValues

	return p
}

func (s *EnumValuesContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumValuesContext) AllEnumValue() []IEnumValueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IEnumValueContext)(nil)).Elem())
	var tst = make([]IEnumValueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IEnumValueContext)
		}
	}

	return tst
}

func (s *EnumValuesContext) EnumValue(i int) IEnumValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnumValueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IEnumValueContext)
}

func (s *EnumValuesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumValuesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumValuesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DMLListener); ok {
		listenerT.EnterEnumValues(s)
	}
}

func (s *EnumValuesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DMLListener); ok {
		listenerT.ExitEnumValues(s)
	}
}

func (p *DMLParser) EnumValues() (localctx IEnumValuesContext) {
	localctx = NewEnumValuesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, DMLParserRULE_enumValues)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(128)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == DMLParserID || _la == DMLParserKEEP_COMMENT {
		{
			p.SetState(127)
			p.EnumValue()
		}

		p.SetState(130)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IIntegerConstContext is an interface to support dynamic dispatch.
type IIntegerConstContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIntegerConstContext differentiates from other interfaces.
	IsIntegerConstContext()
}

type IntegerConstContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIntegerConstContext() *IntegerConstContext {
	var p = new(IntegerConstContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DMLParserRULE_integerConst
	return p
}

func (*IntegerConstContext) IsIntegerConstContext() {}

func NewIntegerConstContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntegerConstContext {
	var p = new(IntegerConstContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DMLParserRULE_integerConst

	return p
}

func (s *IntegerConstContext) GetParser() antlr.Parser { return s.parser }

func (s *IntegerConstContext) INTEGER_NUM() antlr.TerminalNode {
	return s.GetToken(DMLParserINTEGER_NUM, 0)
}

func (s *IntegerConstContext) PLUS() antlr.TerminalNode {
	return s.GetToken(DMLParserPLUS, 0)
}

func (s *IntegerConstContext) MINUS() antlr.TerminalNode {
	return s.GetToken(DMLParserMINUS, 0)
}

func (s *IntegerConstContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntegerConstContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IntegerConstContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DMLListener); ok {
		listenerT.EnterIntegerConst(s)
	}
}

func (s *IntegerConstContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DMLListener); ok {
		listenerT.ExitIntegerConst(s)
	}
}

func (p *DMLParser) IntegerConst() (localctx IIntegerConstContext) {
	localctx = NewIntegerConstContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, DMLParserRULE_integerConst)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(133)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == DMLParserPLUS || _la == DMLParserMINUS {
		{
			p.SetState(132)
			_la = p.GetTokenStream().LA(1)

			if !(_la == DMLParserPLUS || _la == DMLParserMINUS) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	}
	{
		p.SetState(135)
		p.Match(DMLParserINTEGER_NUM)
	}

	return localctx
}

// IEnumValueContext is an interface to support dynamic dispatch.
type IEnumValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEnumValueContext differentiates from other interfaces.
	IsEnumValueContext()
}

type EnumValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumValueContext() *EnumValueContext {
	var p = new(EnumValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DMLParserRULE_enumValue
	return p
}

func (*EnumValueContext) IsEnumValueContext() {}

func NewEnumValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumValueContext {
	var p = new(EnumValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DMLParserRULE_enumValue

	return p
}

func (s *EnumValueContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumValueContext) ID() antlr.TerminalNode {
	return s.GetToken(DMLParserID, 0)
}

func (s *EnumValueContext) SEMI() antlr.TerminalNode {
	return s.GetToken(DMLParserSEMI, 0)
}

func (s *EnumValueContext) AllKEEP_COMMENT() []antlr.TerminalNode {
	return s.GetTokens(DMLParserKEEP_COMMENT)
}

func (s *EnumValueContext) KEEP_COMMENT(i int) antlr.TerminalNode {
	return s.GetToken(DMLParserKEEP_COMMENT, i)
}

func (s *EnumValueContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(DMLParserEQUALS, 0)
}

func (s *EnumValueContext) IntegerConst() IIntegerConstContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIntegerConstContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIntegerConstContext)
}

func (s *EnumValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DMLListener); ok {
		listenerT.EnterEnumValue(s)
	}
}

func (s *EnumValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DMLListener); ok {
		listenerT.ExitEnumValue(s)
	}
}

func (p *DMLParser) EnumValue() (localctx IEnumValueContext) {
	localctx = NewEnumValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, DMLParserRULE_enumValue)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(140)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == DMLParserKEEP_COMMENT {
		{
			p.SetState(137)
			p.Match(DMLParserKEEP_COMMENT)
		}

		p.SetState(142)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(143)
		p.Match(DMLParserID)
	}
	p.SetState(146)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == DMLParserEQUALS {
		{
			p.SetState(144)
			p.Match(DMLParserEQUALS)
		}
		{
			p.SetState(145)
			p.IntegerConst()
		}

	}
	{
		p.SetState(148)
		p.Match(DMLParserSEMI)
	}

	return localctx
}

// IFieldsetElementContext is an interface to support dynamic dispatch.
type IFieldsetElementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldsetElementContext differentiates from other interfaces.
	IsFieldsetElementContext()
}

type FieldsetElementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldsetElementContext() *FieldsetElementContext {
	var p = new(FieldsetElementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DMLParserRULE_fieldsetElement
	return p
}

func (*FieldsetElementContext) IsFieldsetElementContext() {}

func NewFieldsetElementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldsetElementContext {
	var p = new(FieldsetElementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DMLParserRULE_fieldsetElement

	return p
}

func (s *FieldsetElementContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldsetElementContext) FIELDSET() antlr.TerminalNode {
	return s.GetToken(DMLParserFIELDSET, 0)
}

func (s *FieldsetElementContext) ID() antlr.TerminalNode {
	return s.GetToken(DMLParserID, 0)
}

func (s *FieldsetElementContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(DMLParserLBRACE, 0)
}

func (s *FieldsetElementContext) FieldsetValues() IFieldsetValuesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldsetValuesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldsetValuesContext)
}

func (s *FieldsetElementContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(DMLParserRBRACE, 0)
}

func (s *FieldsetElementContext) AllKEEP_COMMENT() []antlr.TerminalNode {
	return s.GetTokens(DMLParserKEEP_COMMENT)
}

func (s *FieldsetElementContext) KEEP_COMMENT(i int) antlr.TerminalNode {
	return s.GetToken(DMLParserKEEP_COMMENT, i)
}

func (s *FieldsetElementContext) ExtendsQualifier() IExtendsQualifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExtendsQualifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExtendsQualifierContext)
}

func (s *FieldsetElementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldsetElementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldsetElementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DMLListener); ok {
		listenerT.EnterFieldsetElement(s)
	}
}

func (s *FieldsetElementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DMLListener); ok {
		listenerT.ExitFieldsetElement(s)
	}
}

func (p *DMLParser) FieldsetElement() (localctx IFieldsetElementContext) {
	localctx = NewFieldsetElementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, DMLParserRULE_fieldsetElement)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(153)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == DMLParserKEEP_COMMENT {
		{
			p.SetState(150)
			p.Match(DMLParserKEEP_COMMENT)
		}

		p.SetState(155)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(156)
		p.Match(DMLParserFIELDSET)
	}
	{
		p.SetState(157)
		p.Match(DMLParserID)
	}
	p.SetState(159)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == DMLParserEXTENDS {
		{
			p.SetState(158)
			p.ExtendsQualifier()
		}

	}
	{
		p.SetState(161)
		p.Match(DMLParserLBRACE)
	}
	{
		p.SetState(162)
		p.FieldsetValues()
	}
	{
		p.SetState(163)
		p.Match(DMLParserRBRACE)
	}

	return localctx
}

// IExtendsQualifierContext is an interface to support dynamic dispatch.
type IExtendsQualifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExtendsQualifierContext differentiates from other interfaces.
	IsExtendsQualifierContext()
}

type ExtendsQualifierContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExtendsQualifierContext() *ExtendsQualifierContext {
	var p = new(ExtendsQualifierContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DMLParserRULE_extendsQualifier
	return p
}

func (*ExtendsQualifierContext) IsExtendsQualifierContext() {}

func NewExtendsQualifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExtendsQualifierContext {
	var p = new(ExtendsQualifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DMLParserRULE_extendsQualifier

	return p
}

func (s *ExtendsQualifierContext) GetParser() antlr.Parser { return s.parser }

func (s *ExtendsQualifierContext) EXTENDS() antlr.TerminalNode {
	return s.GetToken(DMLParserEXTENDS, 0)
}

func (s *ExtendsQualifierContext) ID() antlr.TerminalNode {
	return s.GetToken(DMLParserID, 0)
}

func (s *ExtendsQualifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExtendsQualifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExtendsQualifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DMLListener); ok {
		listenerT.EnterExtendsQualifier(s)
	}
}

func (s *ExtendsQualifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DMLListener); ok {
		listenerT.ExitExtendsQualifier(s)
	}
}

func (p *DMLParser) ExtendsQualifier() (localctx IExtendsQualifierContext) {
	localctx = NewExtendsQualifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, DMLParserRULE_extendsQualifier)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(165)
		p.Match(DMLParserEXTENDS)
	}
	{
		p.SetState(166)
		p.Match(DMLParserID)
	}

	return localctx
}

// IFieldsetValuesContext is an interface to support dynamic dispatch.
type IFieldsetValuesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldsetValuesContext differentiates from other interfaces.
	IsFieldsetValuesContext()
}

type FieldsetValuesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldsetValuesContext() *FieldsetValuesContext {
	var p = new(FieldsetValuesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DMLParserRULE_fieldsetValues
	return p
}

func (*FieldsetValuesContext) IsFieldsetValuesContext() {}

func NewFieldsetValuesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldsetValuesContext {
	var p = new(FieldsetValuesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DMLParserRULE_fieldsetValues

	return p
}

func (s *FieldsetValuesContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldsetValuesContext) AllFieldsetValue() []IFieldsetValueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFieldsetValueContext)(nil)).Elem())
	var tst = make([]IFieldsetValueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFieldsetValueContext)
		}
	}

	return tst
}

func (s *FieldsetValuesContext) FieldsetValue(i int) IFieldsetValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldsetValueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFieldsetValueContext)
}

func (s *FieldsetValuesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldsetValuesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldsetValuesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DMLListener); ok {
		listenerT.EnterFieldsetValues(s)
	}
}

func (s *FieldsetValuesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DMLListener); ok {
		listenerT.ExitFieldsetValues(s)
	}
}

func (p *DMLParser) FieldsetValues() (localctx IFieldsetValuesContext) {
	localctx = NewFieldsetValuesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, DMLParserRULE_fieldsetValues)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(169)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la-19)&-(0x1f+1)) == 0 && ((1<<uint((_la-19)))&((1<<(DMLParserDOUBLE-19))|(1<<(DMLParserFLOAT-19))|(1<<(DMLParserINT32-19))|(1<<(DMLParserINT64-19))|(1<<(DMLParserUINT32-19))|(1<<(DMLParserUINT64-19))|(1<<(DMLParserBOOL-19))|(1<<(DMLParserSTRING-19))|(1<<(DMLParserBYTES-19))|(1<<(DMLParserCHARARRAY-19))|(1<<(DMLParserBYTEARRAY-19))|(1<<(DMLParserDATETIME-19))|(1<<(DMLParserGUID-19))|(1<<(DMLParserDECIMAL-19)))) != 0) || _la == DMLParserID || _la == DMLParserKEEP_COMMENT {
		{
			p.SetState(168)
			p.FieldsetValue()
		}

		p.SetState(171)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IFieldsetValueContext is an interface to support dynamic dispatch.
type IFieldsetValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldsetValueContext differentiates from other interfaces.
	IsFieldsetValueContext()
}

type FieldsetValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldsetValueContext() *FieldsetValueContext {
	var p = new(FieldsetValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DMLParserRULE_fieldsetValue
	return p
}

func (*FieldsetValueContext) IsFieldsetValueContext() {}

func NewFieldsetValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldsetValueContext {
	var p = new(FieldsetValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DMLParserRULE_fieldsetValue

	return p
}

func (s *FieldsetValueContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldsetValueContext) FieldType() IFieldTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldTypeContext)
}

func (s *FieldsetValueContext) ID() antlr.TerminalNode {
	return s.GetToken(DMLParserID, 0)
}

func (s *FieldsetValueContext) SEMI() antlr.TerminalNode {
	return s.GetToken(DMLParserSEMI, 0)
}

func (s *FieldsetValueContext) AllKEEP_COMMENT() []antlr.TerminalNode {
	return s.GetTokens(DMLParserKEEP_COMMENT)
}

func (s *FieldsetValueContext) KEEP_COMMENT(i int) antlr.TerminalNode {
	return s.GetToken(DMLParserKEEP_COMMENT, i)
}

func (s *FieldsetValueContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(DMLParserEQUALS, 0)
}

func (s *FieldsetValueContext) IntegerConst() IIntegerConstContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIntegerConstContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIntegerConstContext)
}

func (s *FieldsetValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldsetValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldsetValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DMLListener); ok {
		listenerT.EnterFieldsetValue(s)
	}
}

func (s *FieldsetValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DMLListener); ok {
		listenerT.ExitFieldsetValue(s)
	}
}

func (p *DMLParser) FieldsetValue() (localctx IFieldsetValueContext) {
	localctx = NewFieldsetValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, DMLParserRULE_fieldsetValue)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(176)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == DMLParserKEEP_COMMENT {
		{
			p.SetState(173)
			p.Match(DMLParserKEEP_COMMENT)
		}

		p.SetState(178)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(179)
		p.FieldType()
	}
	{
		p.SetState(180)
		p.Match(DMLParserID)
	}
	p.SetState(183)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == DMLParserEQUALS {
		{
			p.SetState(181)
			p.Match(DMLParserEQUALS)
		}
		{
			p.SetState(182)
			p.IntegerConst()
		}

	}
	{
		p.SetState(185)
		p.Match(DMLParserSEMI)
	}

	return localctx
}

// IQualifiedIdContext is an interface to support dynamic dispatch.
type IQualifiedIdContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsQualifiedIdContext differentiates from other interfaces.
	IsQualifiedIdContext()
}

type QualifiedIdContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQualifiedIdContext() *QualifiedIdContext {
	var p = new(QualifiedIdContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DMLParserRULE_qualifiedId
	return p
}

func (*QualifiedIdContext) IsQualifiedIdContext() {}

func NewQualifiedIdContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QualifiedIdContext {
	var p = new(QualifiedIdContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DMLParserRULE_qualifiedId

	return p
}

func (s *QualifiedIdContext) GetParser() antlr.Parser { return s.parser }

func (s *QualifiedIdContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(DMLParserID)
}

func (s *QualifiedIdContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(DMLParserID, i)
}

func (s *QualifiedIdContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(DMLParserDOT)
}

func (s *QualifiedIdContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(DMLParserDOT, i)
}

func (s *QualifiedIdContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QualifiedIdContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QualifiedIdContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DMLListener); ok {
		listenerT.EnterQualifiedId(s)
	}
}

func (s *QualifiedIdContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DMLListener); ok {
		listenerT.ExitQualifiedId(s)
	}
}

func (p *DMLParser) QualifiedId() (localctx IQualifiedIdContext) {
	localctx = NewQualifiedIdContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, DMLParserRULE_qualifiedId)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(187)
		p.Match(DMLParserID)
	}
	p.SetState(192)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == DMLParserDOT {
		{
			p.SetState(188)
			p.Match(DMLParserDOT)
		}
		{
			p.SetState(189)
			p.Match(DMLParserID)
		}

		p.SetState(194)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IGoPathContext is an interface to support dynamic dispatch.
type IGoPathContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGoPathContext differentiates from other interfaces.
	IsGoPathContext()
}

type GoPathContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGoPathContext() *GoPathContext {
	var p = new(GoPathContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DMLParserRULE_goPath
	return p
}

func (*GoPathContext) IsGoPathContext() {}

func NewGoPathContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GoPathContext {
	var p = new(GoPathContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DMLParserRULE_goPath

	return p
}

func (s *GoPathContext) GetParser() antlr.Parser { return s.parser }

func (s *GoPathContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(DMLParserID)
}

func (s *GoPathContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(DMLParserID, i)
}

func (s *GoPathContext) AllSLASH() []antlr.TerminalNode {
	return s.GetTokens(DMLParserSLASH)
}

func (s *GoPathContext) SLASH(i int) antlr.TerminalNode {
	return s.GetToken(DMLParserSLASH, i)
}

func (s *GoPathContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(DMLParserDOT)
}

func (s *GoPathContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(DMLParserDOT, i)
}

func (s *GoPathContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GoPathContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GoPathContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DMLListener); ok {
		listenerT.EnterGoPath(s)
	}
}

func (s *GoPathContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DMLListener); ok {
		listenerT.ExitGoPath(s)
	}
}

func (p *DMLParser) GoPath() (localctx IGoPathContext) {
	localctx = NewGoPathContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, DMLParserRULE_goPath)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(195)
		p.Match(DMLParserID)
	}
	p.SetState(200)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == DMLParserDOT || _la == DMLParserSLASH {
		{
			p.SetState(196)
			_la = p.GetTokenStream().LA(1)

			if !(_la == DMLParserDOT || _la == DMLParserSLASH) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(197)
			p.Match(DMLParserID)
		}

		p.SetState(202)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IFieldTypeContext is an interface to support dynamic dispatch.
type IFieldTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldTypeContext differentiates from other interfaces.
	IsFieldTypeContext()
}

type FieldTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldTypeContext() *FieldTypeContext {
	var p = new(FieldTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DMLParserRULE_fieldType
	return p
}

func (*FieldTypeContext) IsFieldTypeContext() {}

func NewFieldTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldTypeContext {
	var p = new(FieldTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DMLParserRULE_fieldType

	return p
}

func (s *FieldTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldTypeContext) DOUBLE() antlr.TerminalNode {
	return s.GetToken(DMLParserDOUBLE, 0)
}

func (s *FieldTypeContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(DMLParserFLOAT, 0)
}

func (s *FieldTypeContext) INT32() antlr.TerminalNode {
	return s.GetToken(DMLParserINT32, 0)
}

func (s *FieldTypeContext) INT64() antlr.TerminalNode {
	return s.GetToken(DMLParserINT64, 0)
}

func (s *FieldTypeContext) UINT32() antlr.TerminalNode {
	return s.GetToken(DMLParserUINT32, 0)
}

func (s *FieldTypeContext) UINT64() antlr.TerminalNode {
	return s.GetToken(DMLParserUINT64, 0)
}

func (s *FieldTypeContext) BOOL() antlr.TerminalNode {
	return s.GetToken(DMLParserBOOL, 0)
}

func (s *FieldTypeContext) STRING() antlr.TerminalNode {
	return s.GetToken(DMLParserSTRING, 0)
}

func (s *FieldTypeContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(DMLParserLPAREN, 0)
}

func (s *FieldTypeContext) IntegerConst() IIntegerConstContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIntegerConstContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIntegerConstContext)
}

func (s *FieldTypeContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(DMLParserRPAREN, 0)
}

func (s *FieldTypeContext) BYTES() antlr.TerminalNode {
	return s.GetToken(DMLParserBYTES, 0)
}

func (s *FieldTypeContext) CHARARRAY() antlr.TerminalNode {
	return s.GetToken(DMLParserCHARARRAY, 0)
}

func (s *FieldTypeContext) BYTEARRAY() antlr.TerminalNode {
	return s.GetToken(DMLParserBYTEARRAY, 0)
}

func (s *FieldTypeContext) DATETIME() antlr.TerminalNode {
	return s.GetToken(DMLParserDATETIME, 0)
}

func (s *FieldTypeContext) GUID() antlr.TerminalNode {
	return s.GetToken(DMLParserGUID, 0)
}

func (s *FieldTypeContext) DECIMAL() antlr.TerminalNode {
	return s.GetToken(DMLParserDECIMAL, 0)
}

func (s *FieldTypeContext) QualifiedId() IQualifiedIdContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQualifiedIdContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQualifiedIdContext)
}

func (s *FieldTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DMLListener); ok {
		listenerT.EnterFieldType(s)
	}
}

func (s *FieldTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DMLListener); ok {
		listenerT.ExitFieldType(s)
	}
}

func (p *DMLParser) FieldType() (localctx IFieldTypeContext) {
	localctx = NewFieldTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, DMLParserRULE_fieldType)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(242)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case DMLParserDOUBLE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(203)
			p.Match(DMLParserDOUBLE)
		}

	case DMLParserFLOAT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(204)
			p.Match(DMLParserFLOAT)
		}

	case DMLParserINT32:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(205)
			p.Match(DMLParserINT32)
		}

	case DMLParserINT64:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(206)
			p.Match(DMLParserINT64)
		}

	case DMLParserUINT32:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(207)
			p.Match(DMLParserUINT32)
		}

	case DMLParserUINT64:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(208)
			p.Match(DMLParserUINT64)
		}

	case DMLParserBOOL:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(209)
			p.Match(DMLParserBOOL)
		}

	case DMLParserSTRING:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(210)
			p.Match(DMLParserSTRING)
		}
		p.SetState(215)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == DMLParserLPAREN {
			{
				p.SetState(211)
				p.Match(DMLParserLPAREN)
			}
			{
				p.SetState(212)
				p.IntegerConst()
			}
			{
				p.SetState(213)
				p.Match(DMLParserRPAREN)
			}

		}

	case DMLParserBYTES:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(217)
			p.Match(DMLParserBYTES)
		}
		p.SetState(222)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == DMLParserLPAREN {
			{
				p.SetState(218)
				p.Match(DMLParserLPAREN)
			}
			{
				p.SetState(219)
				p.IntegerConst()
			}
			{
				p.SetState(220)
				p.Match(DMLParserRPAREN)
			}

		}

	case DMLParserCHARARRAY:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(224)
			p.Match(DMLParserCHARARRAY)
		}
		p.SetState(229)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == DMLParserLPAREN {
			{
				p.SetState(225)
				p.Match(DMLParserLPAREN)
			}
			{
				p.SetState(226)
				p.IntegerConst()
			}
			{
				p.SetState(227)
				p.Match(DMLParserRPAREN)
			}

		}

	case DMLParserBYTEARRAY:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(231)
			p.Match(DMLParserBYTEARRAY)
		}
		p.SetState(236)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == DMLParserLPAREN {
			{
				p.SetState(232)
				p.Match(DMLParserLPAREN)
			}
			{
				p.SetState(233)
				p.IntegerConst()
			}
			{
				p.SetState(234)
				p.Match(DMLParserRPAREN)
			}

		}

	case DMLParserDATETIME:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(238)
			p.Match(DMLParserDATETIME)
		}

	case DMLParserGUID:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(239)
			p.Match(DMLParserGUID)
		}

	case DMLParserDECIMAL:
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(240)
			p.Match(DMLParserDECIMAL)
		}

	case DMLParserID:
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(241)
			p.QualifiedId()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IIndexTypeContext is an interface to support dynamic dispatch.
type IIndexTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIndexTypeContext differentiates from other interfaces.
	IsIndexTypeContext()
}

type IndexTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIndexTypeContext() *IndexTypeContext {
	var p = new(IndexTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DMLParserRULE_indexType
	return p
}

func (*IndexTypeContext) IsIndexTypeContext() {}

func NewIndexTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IndexTypeContext {
	var p = new(IndexTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DMLParserRULE_indexType

	return p
}

func (s *IndexTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *IndexTypeContext) PRIMARY() antlr.TerminalNode {
	return s.GetToken(DMLParserPRIMARY, 0)
}

func (s *IndexTypeContext) UNIQUE() antlr.TerminalNode {
	return s.GetToken(DMLParserUNIQUE, 0)
}

func (s *IndexTypeContext) NONUNIQUE() antlr.TerminalNode {
	return s.GetToken(DMLParserNONUNIQUE, 0)
}

func (s *IndexTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IndexTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IndexTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DMLListener); ok {
		listenerT.EnterIndexType(s)
	}
}

func (s *IndexTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DMLListener); ok {
		listenerT.ExitIndexType(s)
	}
}

func (p *DMLParser) IndexType() (localctx IIndexTypeContext) {
	localctx = NewIndexTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, DMLParserRULE_indexType)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(244)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-44)&-(0x1f+1)) == 0 && ((1<<uint((_la-44)))&((1<<(DMLParserPRIMARY-44))|(1<<(DMLParserUNIQUE-44))|(1<<(DMLParserNONUNIQUE-44)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IIndexElementContext is an interface to support dynamic dispatch.
type IIndexElementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIndexElementContext differentiates from other interfaces.
	IsIndexElementContext()
}

type IndexElementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIndexElementContext() *IndexElementContext {
	var p = new(IndexElementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DMLParserRULE_indexElement
	return p
}

func (*IndexElementContext) IsIndexElementContext() {}

func NewIndexElementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IndexElementContext {
	var p = new(IndexElementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DMLParserRULE_indexElement

	return p
}

func (s *IndexElementContext) GetParser() antlr.Parser { return s.parser }

func (s *IndexElementContext) INDEX() antlr.TerminalNode {
	return s.GetToken(DMLParserINDEX, 0)
}

func (s *IndexElementContext) IndexType() IIndexTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIndexTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIndexTypeContext)
}

func (s *IndexElementContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(DMLParserLPAREN, 0)
}

func (s *IndexElementContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(DMLParserID)
}

func (s *IndexElementContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(DMLParserID, i)
}

func (s *IndexElementContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(DMLParserRPAREN, 0)
}

func (s *IndexElementContext) SEMI() antlr.TerminalNode {
	return s.GetToken(DMLParserSEMI, 0)
}

func (s *IndexElementContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(DMLParserCOMMA)
}

func (s *IndexElementContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(DMLParserCOMMA, i)
}

func (s *IndexElementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IndexElementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IndexElementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DMLListener); ok {
		listenerT.EnterIndexElement(s)
	}
}

func (s *IndexElementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DMLListener); ok {
		listenerT.ExitIndexElement(s)
	}
}

func (p *DMLParser) IndexElement() (localctx IIndexElementContext) {
	localctx = NewIndexElementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, DMLParserRULE_indexElement)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(246)
		p.Match(DMLParserINDEX)
	}
	{
		p.SetState(247)
		p.IndexType()
	}
	{
		p.SetState(248)
		p.Match(DMLParserLPAREN)
	}
	{
		p.SetState(249)
		p.Match(DMLParserID)
	}
	p.SetState(254)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == DMLParserCOMMA {
		{
			p.SetState(250)
			p.Match(DMLParserCOMMA)
		}
		{
			p.SetState(251)
			p.Match(DMLParserID)
		}

		p.SetState(256)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(257)
		p.Match(DMLParserRPAREN)
	}
	{
		p.SetState(258)
		p.Match(DMLParserSEMI)
	}

	return localctx
}

// IClassElementContext is an interface to support dynamic dispatch.
type IClassElementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsClassElementContext differentiates from other interfaces.
	IsClassElementContext()
}

type ClassElementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClassElementContext() *ClassElementContext {
	var p = new(ClassElementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DMLParserRULE_classElement
	return p
}

func (*ClassElementContext) IsClassElementContext() {}

func NewClassElementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClassElementContext {
	var p = new(ClassElementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DMLParserRULE_classElement

	return p
}

func (s *ClassElementContext) GetParser() antlr.Parser { return s.parser }

func (s *ClassElementContext) CLASS() antlr.TerminalNode {
	return s.GetToken(DMLParserCLASS, 0)
}

func (s *ClassElementContext) ID() antlr.TerminalNode {
	return s.GetToken(DMLParserID, 0)
}

func (s *ClassElementContext) ClassOptions() IClassOptionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IClassOptionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IClassOptionsContext)
}

func (s *ClassElementContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(DMLParserLBRACE, 0)
}

func (s *ClassElementContext) ClassFields() IClassFieldsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IClassFieldsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IClassFieldsContext)
}

func (s *ClassElementContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(DMLParserRBRACE, 0)
}

func (s *ClassElementContext) AllKEEP_COMMENT() []antlr.TerminalNode {
	return s.GetTokens(DMLParserKEEP_COMMENT)
}

func (s *ClassElementContext) KEEP_COMMENT(i int) antlr.TerminalNode {
	return s.GetToken(DMLParserKEEP_COMMENT, i)
}

func (s *ClassElementContext) AllIndexElement() []IIndexElementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IIndexElementContext)(nil)).Elem())
	var tst = make([]IIndexElementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IIndexElementContext)
		}
	}

	return tst
}

func (s *ClassElementContext) IndexElement(i int) IIndexElementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIndexElementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IIndexElementContext)
}

func (s *ClassElementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClassElementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ClassElementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DMLListener); ok {
		listenerT.EnterClassElement(s)
	}
}

func (s *ClassElementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DMLListener); ok {
		listenerT.ExitClassElement(s)
	}
}

func (p *DMLParser) ClassElement() (localctx IClassElementContext) {
	localctx = NewClassElementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, DMLParserRULE_classElement)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(263)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == DMLParserKEEP_COMMENT {
		{
			p.SetState(260)
			p.Match(DMLParserKEEP_COMMENT)
		}

		p.SetState(265)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(266)
		p.Match(DMLParserCLASS)
	}
	{
		p.SetState(267)
		p.Match(DMLParserID)
	}
	{
		p.SetState(268)
		p.ClassOptions()
	}
	{
		p.SetState(269)
		p.Match(DMLParserLBRACE)
	}
	{
		p.SetState(270)
		p.ClassFields()
	}
	p.SetState(274)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == DMLParserINDEX {
		{
			p.SetState(271)
			p.IndexElement()
		}

		p.SetState(276)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(277)
		p.Match(DMLParserRBRACE)
	}

	return localctx
}

// IClassOptionsContext is an interface to support dynamic dispatch.
type IClassOptionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsClassOptionsContext differentiates from other interfaces.
	IsClassOptionsContext()
}

type ClassOptionsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClassOptionsContext() *ClassOptionsContext {
	var p = new(ClassOptionsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DMLParserRULE_classOptions
	return p
}

func (*ClassOptionsContext) IsClassOptionsContext() {}

func NewClassOptionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClassOptionsContext {
	var p = new(ClassOptionsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DMLParserRULE_classOptions

	return p
}

func (s *ClassOptionsContext) GetParser() antlr.Parser { return s.parser }

func (s *ClassOptionsContext) AllClassOption() []IClassOptionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IClassOptionContext)(nil)).Elem())
	var tst = make([]IClassOptionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IClassOptionContext)
		}
	}

	return tst
}

func (s *ClassOptionsContext) ClassOption(i int) IClassOptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IClassOptionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IClassOptionContext)
}

func (s *ClassOptionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClassOptionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ClassOptionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DMLListener); ok {
		listenerT.EnterClassOptions(s)
	}
}

func (s *ClassOptionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DMLListener); ok {
		listenerT.ExitClassOptions(s)
	}
}

func (p *DMLParser) ClassOptions() (localctx IClassOptionsContext) {
	localctx = NewClassOptionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, DMLParserRULE_classOptions)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(282)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == DMLParserFIELDSET || _la == DMLParserTABLE {
		{
			p.SetState(279)
			p.ClassOption()
		}

		p.SetState(284)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IClassOptionContext is an interface to support dynamic dispatch.
type IClassOptionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsClassOptionContext differentiates from other interfaces.
	IsClassOptionContext()
}

type ClassOptionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClassOptionContext() *ClassOptionContext {
	var p = new(ClassOptionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DMLParserRULE_classOption
	return p
}

func (*ClassOptionContext) IsClassOptionContext() {}

func NewClassOptionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClassOptionContext {
	var p = new(ClassOptionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DMLParserRULE_classOption

	return p
}

func (s *ClassOptionContext) GetParser() antlr.Parser { return s.parser }

func (s *ClassOptionContext) FieldsetOption() IFieldsetOptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldsetOptionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldsetOptionContext)
}

func (s *ClassOptionContext) TableOption() ITableOptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITableOptionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITableOptionContext)
}

func (s *ClassOptionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClassOptionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ClassOptionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DMLListener); ok {
		listenerT.EnterClassOption(s)
	}
}

func (s *ClassOptionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DMLListener); ok {
		listenerT.ExitClassOption(s)
	}
}

func (p *DMLParser) ClassOption() (localctx IClassOptionContext) {
	localctx = NewClassOptionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, DMLParserRULE_classOption)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(287)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case DMLParserFIELDSET:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(285)
			p.FieldsetOption()
		}

	case DMLParserTABLE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(286)
			p.TableOption()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IFieldsetOptionContext is an interface to support dynamic dispatch.
type IFieldsetOptionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldsetOptionContext differentiates from other interfaces.
	IsFieldsetOptionContext()
}

type FieldsetOptionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldsetOptionContext() *FieldsetOptionContext {
	var p = new(FieldsetOptionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DMLParserRULE_fieldsetOption
	return p
}

func (*FieldsetOptionContext) IsFieldsetOptionContext() {}

func NewFieldsetOptionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldsetOptionContext {
	var p = new(FieldsetOptionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DMLParserRULE_fieldsetOption

	return p
}

func (s *FieldsetOptionContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldsetOptionContext) FIELDSET() antlr.TerminalNode {
	return s.GetToken(DMLParserFIELDSET, 0)
}

func (s *FieldsetOptionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(DMLParserLPAREN, 0)
}

func (s *FieldsetOptionContext) ID() antlr.TerminalNode {
	return s.GetToken(DMLParserID, 0)
}

func (s *FieldsetOptionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(DMLParserRPAREN, 0)
}

func (s *FieldsetOptionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldsetOptionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldsetOptionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DMLListener); ok {
		listenerT.EnterFieldsetOption(s)
	}
}

func (s *FieldsetOptionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DMLListener); ok {
		listenerT.ExitFieldsetOption(s)
	}
}

func (p *DMLParser) FieldsetOption() (localctx IFieldsetOptionContext) {
	localctx = NewFieldsetOptionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, DMLParserRULE_fieldsetOption)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(289)
		p.Match(DMLParserFIELDSET)
	}
	{
		p.SetState(290)
		p.Match(DMLParserLPAREN)
	}
	{
		p.SetState(291)
		p.Match(DMLParserID)
	}
	{
		p.SetState(292)
		p.Match(DMLParserRPAREN)
	}

	return localctx
}

// ITableOptionContext is an interface to support dynamic dispatch.
type ITableOptionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTableOptionContext differentiates from other interfaces.
	IsTableOptionContext()
}

type TableOptionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTableOptionContext() *TableOptionContext {
	var p = new(TableOptionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DMLParserRULE_tableOption
	return p
}

func (*TableOptionContext) IsTableOptionContext() {}

func NewTableOptionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TableOptionContext {
	var p = new(TableOptionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DMLParserRULE_tableOption

	return p
}

func (s *TableOptionContext) GetParser() antlr.Parser { return s.parser }

func (s *TableOptionContext) TABLE() antlr.TerminalNode {
	return s.GetToken(DMLParserTABLE, 0)
}

func (s *TableOptionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(DMLParserLPAREN, 0)
}

func (s *TableOptionContext) ID() antlr.TerminalNode {
	return s.GetToken(DMLParserID, 0)
}

func (s *TableOptionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(DMLParserRPAREN, 0)
}

func (s *TableOptionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TableOptionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TableOptionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DMLListener); ok {
		listenerT.EnterTableOption(s)
	}
}

func (s *TableOptionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DMLListener); ok {
		listenerT.ExitTableOption(s)
	}
}

func (p *DMLParser) TableOption() (localctx ITableOptionContext) {
	localctx = NewTableOptionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, DMLParserRULE_tableOption)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(294)
		p.Match(DMLParserTABLE)
	}
	{
		p.SetState(295)
		p.Match(DMLParserLPAREN)
	}
	{
		p.SetState(296)
		p.Match(DMLParserID)
	}
	{
		p.SetState(297)
		p.Match(DMLParserRPAREN)
	}

	return localctx
}

// IMediatesOptionContext is an interface to support dynamic dispatch.
type IMediatesOptionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMediatesOptionContext differentiates from other interfaces.
	IsMediatesOptionContext()
}

type MediatesOptionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMediatesOptionContext() *MediatesOptionContext {
	var p = new(MediatesOptionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DMLParserRULE_mediatesOption
	return p
}

func (*MediatesOptionContext) IsMediatesOptionContext() {}

func NewMediatesOptionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MediatesOptionContext {
	var p = new(MediatesOptionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DMLParserRULE_mediatesOption

	return p
}

func (s *MediatesOptionContext) GetParser() antlr.Parser { return s.parser }

func (s *MediatesOptionContext) MEDIATES() antlr.TerminalNode {
	return s.GetToken(DMLParserMEDIATES, 0)
}

func (s *MediatesOptionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(DMLParserLPAREN, 0)
}

func (s *MediatesOptionContext) ID() antlr.TerminalNode {
	return s.GetToken(DMLParserID, 0)
}

func (s *MediatesOptionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(DMLParserRPAREN, 0)
}

func (s *MediatesOptionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MediatesOptionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MediatesOptionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DMLListener); ok {
		listenerT.EnterMediatesOption(s)
	}
}

func (s *MediatesOptionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DMLListener); ok {
		listenerT.ExitMediatesOption(s)
	}
}

func (p *DMLParser) MediatesOption() (localctx IMediatesOptionContext) {
	localctx = NewMediatesOptionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, DMLParserRULE_mediatesOption)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(299)
		p.Match(DMLParserMEDIATES)
	}
	{
		p.SetState(300)
		p.Match(DMLParserLPAREN)
	}
	{
		p.SetState(301)
		p.Match(DMLParserID)
	}
	{
		p.SetState(302)
		p.Match(DMLParserRPAREN)
	}

	return localctx
}

// IClassFieldsContext is an interface to support dynamic dispatch.
type IClassFieldsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsClassFieldsContext differentiates from other interfaces.
	IsClassFieldsContext()
}

type ClassFieldsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClassFieldsContext() *ClassFieldsContext {
	var p = new(ClassFieldsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DMLParserRULE_classFields
	return p
}

func (*ClassFieldsContext) IsClassFieldsContext() {}

func NewClassFieldsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClassFieldsContext {
	var p = new(ClassFieldsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DMLParserRULE_classFields

	return p
}

func (s *ClassFieldsContext) GetParser() antlr.Parser { return s.parser }

func (s *ClassFieldsContext) AllClassField() []IClassFieldContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IClassFieldContext)(nil)).Elem())
	var tst = make([]IClassFieldContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IClassFieldContext)
		}
	}

	return tst
}

func (s *ClassFieldsContext) ClassField(i int) IClassFieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IClassFieldContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IClassFieldContext)
}

func (s *ClassFieldsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClassFieldsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ClassFieldsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DMLListener); ok {
		listenerT.EnterClassFields(s)
	}
}

func (s *ClassFieldsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DMLListener); ok {
		listenerT.ExitClassFields(s)
	}
}

func (p *DMLParser) ClassFields() (localctx IClassFieldsContext) {
	localctx = NewClassFieldsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, DMLParserRULE_classFields)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(305)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la-33)&-(0x1f+1)) == 0 && ((1<<uint((_la-33)))&((1<<(DMLParserAUTO-33))|(1<<(DMLParserID_KEYWORD-33))|(1<<(DMLParserIDGEN-33))|(1<<(DMLParserREQUIRED-33))|(1<<(DMLParserOPTIONAL-33))|(1<<(DMLParserREPEATED-33))|(1<<(DMLParserVIRTUAL-33))|(1<<(DMLParserMAP-33))|(1<<(DMLParserLIST-33)))) != 0) {
		{
			p.SetState(304)
			p.ClassField()
		}

		p.SetState(307)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IReferencesModifierContext is an interface to support dynamic dispatch.
type IReferencesModifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsReferencesModifierContext differentiates from other interfaces.
	IsReferencesModifierContext()
}

type ReferencesModifierContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReferencesModifierContext() *ReferencesModifierContext {
	var p = new(ReferencesModifierContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DMLParserRULE_referencesModifier
	return p
}

func (*ReferencesModifierContext) IsReferencesModifierContext() {}

func NewReferencesModifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReferencesModifierContext {
	var p = new(ReferencesModifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DMLParserRULE_referencesModifier

	return p
}

func (s *ReferencesModifierContext) GetParser() antlr.Parser { return s.parser }

func (s *ReferencesModifierContext) REFERENCES() antlr.TerminalNode {
	return s.GetToken(DMLParserREFERENCES, 0)
}

func (s *ReferencesModifierContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(DMLParserID)
}

func (s *ReferencesModifierContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(DMLParserID, i)
}

func (s *ReferencesModifierContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(DMLParserLPAREN, 0)
}

func (s *ReferencesModifierContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(DMLParserRPAREN, 0)
}

func (s *ReferencesModifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReferencesModifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReferencesModifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DMLListener); ok {
		listenerT.EnterReferencesModifier(s)
	}
}

func (s *ReferencesModifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DMLListener); ok {
		listenerT.ExitReferencesModifier(s)
	}
}

func (p *DMLParser) ReferencesModifier() (localctx IReferencesModifierContext) {
	localctx = NewReferencesModifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, DMLParserRULE_referencesModifier)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(309)
		p.Match(DMLParserREFERENCES)
	}
	{
		p.SetState(310)
		p.Match(DMLParserID)
	}
	p.SetState(314)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == DMLParserLPAREN {
		{
			p.SetState(311)
			p.Match(DMLParserLPAREN)
		}
		{
			p.SetState(312)
			p.Match(DMLParserID)
		}
		{
			p.SetState(313)
			p.Match(DMLParserRPAREN)
		}

	}

	return localctx
}

// IClassFieldContext is an interface to support dynamic dispatch.
type IClassFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsClassFieldContext differentiates from other interfaces.
	IsClassFieldContext()
}

type ClassFieldContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClassFieldContext() *ClassFieldContext {
	var p = new(ClassFieldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DMLParserRULE_classField
	return p
}

func (*ClassFieldContext) IsClassFieldContext() {}

func NewClassFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClassFieldContext {
	var p = new(ClassFieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DMLParserRULE_classField

	return p
}

func (s *ClassFieldContext) GetParser() antlr.Parser { return s.parser }

func (s *ClassFieldContext) FieldModifier() IFieldModifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldModifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldModifierContext)
}

func (s *ClassFieldContext) ID() antlr.TerminalNode {
	return s.GetToken(DMLParserID, 0)
}

func (s *ClassFieldContext) SEMI() antlr.TerminalNode {
	return s.GetToken(DMLParserSEMI, 0)
}

func (s *ClassFieldContext) ReferencesModifier() IReferencesModifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReferencesModifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IReferencesModifierContext)
}

func (s *ClassFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClassFieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ClassFieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DMLListener); ok {
		listenerT.EnterClassField(s)
	}
}

func (s *ClassFieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DMLListener); ok {
		listenerT.ExitClassField(s)
	}
}

func (p *DMLParser) ClassField() (localctx IClassFieldContext) {
	localctx = NewClassFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, DMLParserRULE_classField)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(316)
		p.FieldModifier()
	}
	{
		p.SetState(317)
		p.Match(DMLParserID)
	}
	p.SetState(319)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == DMLParserREFERENCES {
		{
			p.SetState(318)
			p.ReferencesModifier()
		}

	}
	{
		p.SetState(321)
		p.Match(DMLParserSEMI)
	}

	return localctx
}

// IParameterModifierContext is an interface to support dynamic dispatch.
type IParameterModifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParameterModifierContext differentiates from other interfaces.
	IsParameterModifierContext()
}

type ParameterModifierContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParameterModifierContext() *ParameterModifierContext {
	var p = new(ParameterModifierContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DMLParserRULE_parameterModifier
	return p
}

func (*ParameterModifierContext) IsParameterModifierContext() {}

func NewParameterModifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParameterModifierContext {
	var p = new(ParameterModifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DMLParserRULE_parameterModifier

	return p
}

func (s *ParameterModifierContext) GetParser() antlr.Parser { return s.parser }

func (s *ParameterModifierContext) REQUIRED() antlr.TerminalNode {
	return s.GetToken(DMLParserREQUIRED, 0)
}

func (s *ParameterModifierContext) OPTIONAL() antlr.TerminalNode {
	return s.GetToken(DMLParserOPTIONAL, 0)
}

func (s *ParameterModifierContext) REPEATED() antlr.TerminalNode {
	return s.GetToken(DMLParserREPEATED, 0)
}

func (s *ParameterModifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParameterModifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParameterModifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DMLListener); ok {
		listenerT.EnterParameterModifier(s)
	}
}

func (s *ParameterModifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DMLListener); ok {
		listenerT.ExitParameterModifier(s)
	}
}

func (p *DMLParser) ParameterModifier() (localctx IParameterModifierContext) {
	localctx = NewParameterModifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, DMLParserRULE_parameterModifier)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(323)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-36)&-(0x1f+1)) == 0 && ((1<<uint((_la-36)))&((1<<(DMLParserREQUIRED-36))|(1<<(DMLParserOPTIONAL-36))|(1<<(DMLParserREPEATED-36)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IFieldModifierContext is an interface to support dynamic dispatch.
type IFieldModifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldModifierContext differentiates from other interfaces.
	IsFieldModifierContext()
}

type FieldModifierContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldModifierContext() *FieldModifierContext {
	var p = new(FieldModifierContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DMLParserRULE_fieldModifier
	return p
}

func (*FieldModifierContext) IsFieldModifierContext() {}

func NewFieldModifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldModifierContext {
	var p = new(FieldModifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DMLParserRULE_fieldModifier

	return p
}

func (s *FieldModifierContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldModifierContext) AUTO() antlr.TerminalNode {
	return s.GetToken(DMLParserAUTO, 0)
}

func (s *FieldModifierContext) ID_KEYWORD() antlr.TerminalNode {
	return s.GetToken(DMLParserID_KEYWORD, 0)
}

func (s *FieldModifierContext) IDGEN() antlr.TerminalNode {
	return s.GetToken(DMLParserIDGEN, 0)
}

func (s *FieldModifierContext) VIRTUAL() antlr.TerminalNode {
	return s.GetToken(DMLParserVIRTUAL, 0)
}

func (s *FieldModifierContext) MAP() antlr.TerminalNode {
	return s.GetToken(DMLParserMAP, 0)
}

func (s *FieldModifierContext) LIST() antlr.TerminalNode {
	return s.GetToken(DMLParserLIST, 0)
}

func (s *FieldModifierContext) ParameterModifier() IParameterModifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParameterModifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParameterModifierContext)
}

func (s *FieldModifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldModifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldModifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DMLListener); ok {
		listenerT.EnterFieldModifier(s)
	}
}

func (s *FieldModifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DMLListener); ok {
		listenerT.ExitFieldModifier(s)
	}
}

func (p *DMLParser) FieldModifier() (localctx IFieldModifierContext) {
	localctx = NewFieldModifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, DMLParserRULE_fieldModifier)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(332)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case DMLParserAUTO:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(325)
			p.Match(DMLParserAUTO)
		}

	case DMLParserID_KEYWORD:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(326)
			p.Match(DMLParserID_KEYWORD)
		}

	case DMLParserIDGEN:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(327)
			p.Match(DMLParserIDGEN)
		}

	case DMLParserVIRTUAL:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(328)
			p.Match(DMLParserVIRTUAL)
		}

	case DMLParserMAP:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(329)
			p.Match(DMLParserMAP)
		}

	case DMLParserLIST:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(330)
			p.Match(DMLParserLIST)
		}

	case DMLParserREQUIRED, DMLParserOPTIONAL, DMLParserREPEATED:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(331)
			p.ParameterModifier()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IAssociationElementContext is an interface to support dynamic dispatch.
type IAssociationElementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssociationElementContext differentiates from other interfaces.
	IsAssociationElementContext()
}

type AssociationElementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssociationElementContext() *AssociationElementContext {
	var p = new(AssociationElementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DMLParserRULE_associationElement
	return p
}

func (*AssociationElementContext) IsAssociationElementContext() {}

func NewAssociationElementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssociationElementContext {
	var p = new(AssociationElementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DMLParserRULE_associationElement

	return p
}

func (s *AssociationElementContext) GetParser() antlr.Parser { return s.parser }

func (s *AssociationElementContext) ASSOCIATION() antlr.TerminalNode {
	return s.GetToken(DMLParserASSOCIATION, 0)
}

func (s *AssociationElementContext) ID() antlr.TerminalNode {
	return s.GetToken(DMLParserID, 0)
}

func (s *AssociationElementContext) AssociationOptions() IAssociationOptionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssociationOptionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssociationOptionsContext)
}

func (s *AssociationElementContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(DMLParserLBRACE, 0)
}

func (s *AssociationElementContext) ClassFields() IClassFieldsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IClassFieldsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IClassFieldsContext)
}

func (s *AssociationElementContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(DMLParserRBRACE, 0)
}

func (s *AssociationElementContext) AllKEEP_COMMENT() []antlr.TerminalNode {
	return s.GetTokens(DMLParserKEEP_COMMENT)
}

func (s *AssociationElementContext) KEEP_COMMENT(i int) antlr.TerminalNode {
	return s.GetToken(DMLParserKEEP_COMMENT, i)
}

func (s *AssociationElementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssociationElementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssociationElementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DMLListener); ok {
		listenerT.EnterAssociationElement(s)
	}
}

func (s *AssociationElementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DMLListener); ok {
		listenerT.ExitAssociationElement(s)
	}
}

func (p *DMLParser) AssociationElement() (localctx IAssociationElementContext) {
	localctx = NewAssociationElementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, DMLParserRULE_associationElement)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(337)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == DMLParserKEEP_COMMENT {
		{
			p.SetState(334)
			p.Match(DMLParserKEEP_COMMENT)
		}

		p.SetState(339)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(340)
		p.Match(DMLParserASSOCIATION)
	}
	{
		p.SetState(341)
		p.Match(DMLParserID)
	}
	{
		p.SetState(342)
		p.AssociationOptions()
	}
	{
		p.SetState(343)
		p.Match(DMLParserLBRACE)
	}
	{
		p.SetState(344)
		p.ClassFields()
	}
	{
		p.SetState(345)
		p.Match(DMLParserRBRACE)
	}

	return localctx
}

// IAssociationOptionsContext is an interface to support dynamic dispatch.
type IAssociationOptionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssociationOptionsContext differentiates from other interfaces.
	IsAssociationOptionsContext()
}

type AssociationOptionsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssociationOptionsContext() *AssociationOptionsContext {
	var p = new(AssociationOptionsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DMLParserRULE_associationOptions
	return p
}

func (*AssociationOptionsContext) IsAssociationOptionsContext() {}

func NewAssociationOptionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssociationOptionsContext {
	var p = new(AssociationOptionsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DMLParserRULE_associationOptions

	return p
}

func (s *AssociationOptionsContext) GetParser() antlr.Parser { return s.parser }

func (s *AssociationOptionsContext) AllAssociationOption() []IAssociationOptionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAssociationOptionContext)(nil)).Elem())
	var tst = make([]IAssociationOptionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAssociationOptionContext)
		}
	}

	return tst
}

func (s *AssociationOptionsContext) AssociationOption(i int) IAssociationOptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssociationOptionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAssociationOptionContext)
}

func (s *AssociationOptionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssociationOptionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssociationOptionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DMLListener); ok {
		listenerT.EnterAssociationOptions(s)
	}
}

func (s *AssociationOptionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DMLListener); ok {
		listenerT.ExitAssociationOptions(s)
	}
}

func (p *DMLParser) AssociationOptions() (localctx IAssociationOptionsContext) {
	localctx = NewAssociationOptionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, DMLParserRULE_associationOptions)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(350)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == DMLParserFIELDSET || _la == DMLParserTABLE {
		{
			p.SetState(347)
			p.AssociationOption()
		}

		p.SetState(352)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IAssociationOptionContext is an interface to support dynamic dispatch.
type IAssociationOptionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssociationOptionContext differentiates from other interfaces.
	IsAssociationOptionContext()
}

type AssociationOptionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssociationOptionContext() *AssociationOptionContext {
	var p = new(AssociationOptionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DMLParserRULE_associationOption
	return p
}

func (*AssociationOptionContext) IsAssociationOptionContext() {}

func NewAssociationOptionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssociationOptionContext {
	var p = new(AssociationOptionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DMLParserRULE_associationOption

	return p
}

func (s *AssociationOptionContext) GetParser() antlr.Parser { return s.parser }

func (s *AssociationOptionContext) FieldsetOption() IFieldsetOptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldsetOptionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldsetOptionContext)
}

func (s *AssociationOptionContext) TableOption() ITableOptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITableOptionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITableOptionContext)
}

func (s *AssociationOptionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssociationOptionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssociationOptionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DMLListener); ok {
		listenerT.EnterAssociationOption(s)
	}
}

func (s *AssociationOptionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DMLListener); ok {
		listenerT.ExitAssociationOption(s)
	}
}

func (p *DMLParser) AssociationOption() (localctx IAssociationOptionContext) {
	localctx = NewAssociationOptionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, DMLParserRULE_associationOption)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(355)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case DMLParserFIELDSET:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(353)
			p.FieldsetOption()
		}

	case DMLParserTABLE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(354)
			p.TableOption()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IServiceElementContext is an interface to support dynamic dispatch.
type IServiceElementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsServiceElementContext differentiates from other interfaces.
	IsServiceElementContext()
}

type ServiceElementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyServiceElementContext() *ServiceElementContext {
	var p = new(ServiceElementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DMLParserRULE_serviceElement
	return p
}

func (*ServiceElementContext) IsServiceElementContext() {}

func NewServiceElementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ServiceElementContext {
	var p = new(ServiceElementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DMLParserRULE_serviceElement

	return p
}

func (s *ServiceElementContext) GetParser() antlr.Parser { return s.parser }

func (s *ServiceElementContext) SERVICE() antlr.TerminalNode {
	return s.GetToken(DMLParserSERVICE, 0)
}

func (s *ServiceElementContext) ID() antlr.TerminalNode {
	return s.GetToken(DMLParserID, 0)
}

func (s *ServiceElementContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(DMLParserLBRACE, 0)
}

func (s *ServiceElementContext) Methods() IMethodsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMethodsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMethodsContext)
}

func (s *ServiceElementContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(DMLParserRBRACE, 0)
}

func (s *ServiceElementContext) AllKEEP_COMMENT() []antlr.TerminalNode {
	return s.GetTokens(DMLParserKEEP_COMMENT)
}

func (s *ServiceElementContext) KEEP_COMMENT(i int) antlr.TerminalNode {
	return s.GetToken(DMLParserKEEP_COMMENT, i)
}

func (s *ServiceElementContext) FieldsetOption() IFieldsetOptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldsetOptionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldsetOptionContext)
}

func (s *ServiceElementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ServiceElementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ServiceElementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DMLListener); ok {
		listenerT.EnterServiceElement(s)
	}
}

func (s *ServiceElementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DMLListener); ok {
		listenerT.ExitServiceElement(s)
	}
}

func (p *DMLParser) ServiceElement() (localctx IServiceElementContext) {
	localctx = NewServiceElementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, DMLParserRULE_serviceElement)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(360)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == DMLParserKEEP_COMMENT {
		{
			p.SetState(357)
			p.Match(DMLParserKEEP_COMMENT)
		}

		p.SetState(362)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(363)
		p.Match(DMLParserSERVICE)
	}
	{
		p.SetState(364)
		p.Match(DMLParserID)
	}
	p.SetState(366)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == DMLParserFIELDSET {
		{
			p.SetState(365)
			p.FieldsetOption()
		}

	}
	{
		p.SetState(368)
		p.Match(DMLParserLBRACE)
	}
	{
		p.SetState(369)
		p.Methods()
	}
	{
		p.SetState(370)
		p.Match(DMLParserRBRACE)
	}

	return localctx
}

// IPatternTypeContext is an interface to support dynamic dispatch.
type IPatternTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPatternTypeContext differentiates from other interfaces.
	IsPatternTypeContext()
}

type PatternTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPatternTypeContext() *PatternTypeContext {
	var p = new(PatternTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DMLParserRULE_patternType
	return p
}

func (*PatternTypeContext) IsPatternTypeContext() {}

func NewPatternTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PatternTypeContext {
	var p = new(PatternTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DMLParserRULE_patternType

	return p
}

func (s *PatternTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *PatternTypeContext) INSERT() antlr.TerminalNode {
	return s.GetToken(DMLParserINSERT, 0)
}

func (s *PatternTypeContext) UPDATE() antlr.TerminalNode {
	return s.GetToken(DMLParserUPDATE, 0)
}

func (s *PatternTypeContext) SELECT() antlr.TerminalNode {
	return s.GetToken(DMLParserSELECT, 0)
}

func (s *PatternTypeContext) DELETE() antlr.TerminalNode {
	return s.GetToken(DMLParserDELETE, 0)
}

func (s *PatternTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PatternTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PatternTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DMLListener); ok {
		listenerT.EnterPatternType(s)
	}
}

func (s *PatternTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DMLListener); ok {
		listenerT.ExitPatternType(s)
	}
}

func (p *DMLParser) PatternType() (localctx IPatternTypeContext) {
	localctx = NewPatternTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, DMLParserRULE_patternType)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(372)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-48)&-(0x1f+1)) == 0 && ((1<<uint((_la-48)))&((1<<(DMLParserINSERT-48))|(1<<(DMLParserUPDATE-48))|(1<<(DMLParserSELECT-48))|(1<<(DMLParserDELETE-48)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IMethodPatternContext is an interface to support dynamic dispatch.
type IMethodPatternContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMethodPatternContext differentiates from other interfaces.
	IsMethodPatternContext()
}

type MethodPatternContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMethodPatternContext() *MethodPatternContext {
	var p = new(MethodPatternContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DMLParserRULE_methodPattern
	return p
}

func (*MethodPatternContext) IsMethodPatternContext() {}

func NewMethodPatternContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MethodPatternContext {
	var p = new(MethodPatternContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DMLParserRULE_methodPattern

	return p
}

func (s *MethodPatternContext) GetParser() antlr.Parser { return s.parser }

func (s *MethodPatternContext) PATTERN() antlr.TerminalNode {
	return s.GetToken(DMLParserPATTERN, 0)
}

func (s *MethodPatternContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(DMLParserLPAREN, 0)
}

func (s *MethodPatternContext) PatternType() IPatternTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPatternTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPatternTypeContext)
}

func (s *MethodPatternContext) QualifiedId() IQualifiedIdContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQualifiedIdContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQualifiedIdContext)
}

func (s *MethodPatternContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(DMLParserRPAREN, 0)
}

func (s *MethodPatternContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MethodPatternContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MethodPatternContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DMLListener); ok {
		listenerT.EnterMethodPattern(s)
	}
}

func (s *MethodPatternContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DMLListener); ok {
		listenerT.ExitMethodPattern(s)
	}
}

func (p *DMLParser) MethodPattern() (localctx IMethodPatternContext) {
	localctx = NewMethodPatternContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, DMLParserRULE_methodPattern)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(374)
		p.Match(DMLParserPATTERN)
	}
	{
		p.SetState(375)
		p.Match(DMLParserLPAREN)
	}
	{
		p.SetState(376)
		p.PatternType()
	}
	{
		p.SetState(377)
		p.QualifiedId()
	}
	{
		p.SetState(378)
		p.Match(DMLParserRPAREN)
	}

	return localctx
}

// IMethodsContext is an interface to support dynamic dispatch.
type IMethodsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMethodsContext differentiates from other interfaces.
	IsMethodsContext()
}

type MethodsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMethodsContext() *MethodsContext {
	var p = new(MethodsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DMLParserRULE_methods
	return p
}

func (*MethodsContext) IsMethodsContext() {}

func NewMethodsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MethodsContext {
	var p = new(MethodsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DMLParserRULE_methods

	return p
}

func (s *MethodsContext) GetParser() antlr.Parser { return s.parser }

func (s *MethodsContext) AllMethod() []IMethodContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMethodContext)(nil)).Elem())
	var tst = make([]IMethodContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMethodContext)
		}
	}

	return tst
}

func (s *MethodsContext) Method(i int) IMethodContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMethodContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMethodContext)
}

func (s *MethodsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MethodsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MethodsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DMLListener); ok {
		listenerT.EnterMethods(s)
	}
}

func (s *MethodsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DMLListener); ok {
		listenerT.ExitMethods(s)
	}
}

func (p *DMLParser) Methods() (localctx IMethodsContext) {
	localctx = NewMethodsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, DMLParserRULE_methods)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(381)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == DMLParserMETHOD || _la == DMLParserKEEP_COMMENT {
		{
			p.SetState(380)
			p.Method()
		}

		p.SetState(383)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IMethodContext is an interface to support dynamic dispatch.
type IMethodContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMethodContext differentiates from other interfaces.
	IsMethodContext()
}

type MethodContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMethodContext() *MethodContext {
	var p = new(MethodContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DMLParserRULE_method
	return p
}

func (*MethodContext) IsMethodContext() {}

func NewMethodContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MethodContext {
	var p = new(MethodContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DMLParserRULE_method

	return p
}

func (s *MethodContext) GetParser() antlr.Parser { return s.parser }

func (s *MethodContext) METHOD() antlr.TerminalNode {
	return s.GetToken(DMLParserMETHOD, 0)
}

func (s *MethodContext) ID() antlr.TerminalNode {
	return s.GetToken(DMLParserID, 0)
}

func (s *MethodContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(DMLParserLBRACE, 0)
}

func (s *MethodContext) InputList() IInputListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInputListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInputListContext)
}

func (s *MethodContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(DMLParserRBRACE, 0)
}

func (s *MethodContext) AllKEEP_COMMENT() []antlr.TerminalNode {
	return s.GetTokens(DMLParserKEEP_COMMENT)
}

func (s *MethodContext) KEEP_COMMENT(i int) antlr.TerminalNode {
	return s.GetToken(DMLParserKEEP_COMMENT, i)
}

func (s *MethodContext) MethodPattern() IMethodPatternContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMethodPatternContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMethodPatternContext)
}

func (s *MethodContext) MethodReturn() IMethodReturnContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMethodReturnContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMethodReturnContext)
}

func (s *MethodContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MethodContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MethodContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DMLListener); ok {
		listenerT.EnterMethod(s)
	}
}

func (s *MethodContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DMLListener); ok {
		listenerT.ExitMethod(s)
	}
}

func (p *DMLParser) Method() (localctx IMethodContext) {
	localctx = NewMethodContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, DMLParserRULE_method)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(388)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == DMLParserKEEP_COMMENT {
		{
			p.SetState(385)
			p.Match(DMLParserKEEP_COMMENT)
		}

		p.SetState(390)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(391)
		p.Match(DMLParserMETHOD)
	}
	{
		p.SetState(392)
		p.Match(DMLParserID)
	}
	p.SetState(394)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == DMLParserPATTERN {
		{
			p.SetState(393)
			p.MethodPattern()
		}

	}
	{
		p.SetState(396)
		p.Match(DMLParserLBRACE)
	}
	{
		p.SetState(397)
		p.InputList()
	}
	{
		p.SetState(398)
		p.Match(DMLParserRBRACE)
	}
	p.SetState(400)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == DMLParserRETURNS {
		{
			p.SetState(399)
			p.MethodReturn()
		}

	}

	return localctx
}

// IInputListContext is an interface to support dynamic dispatch.
type IInputListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInputListContext differentiates from other interfaces.
	IsInputListContext()
}

type InputListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInputListContext() *InputListContext {
	var p = new(InputListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DMLParserRULE_inputList
	return p
}

func (*InputListContext) IsInputListContext() {}

func NewInputListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InputListContext {
	var p = new(InputListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DMLParserRULE_inputList

	return p
}

func (s *InputListContext) GetParser() antlr.Parser { return s.parser }

func (s *InputListContext) FieldList() IFieldListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldListContext)
}

func (s *InputListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InputListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InputListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DMLListener); ok {
		listenerT.EnterInputList(s)
	}
}

func (s *InputListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DMLListener); ok {
		listenerT.ExitInputList(s)
	}
}

func (p *DMLParser) InputList() (localctx IInputListContext) {
	localctx = NewInputListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, DMLParserRULE_inputList)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(402)
		p.FieldList()
	}

	return localctx
}

// IFieldListContext is an interface to support dynamic dispatch.
type IFieldListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldListContext differentiates from other interfaces.
	IsFieldListContext()
}

type FieldListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldListContext() *FieldListContext {
	var p = new(FieldListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DMLParserRULE_fieldList
	return p
}

func (*FieldListContext) IsFieldListContext() {}

func NewFieldListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldListContext {
	var p = new(FieldListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DMLParserRULE_fieldList

	return p
}

func (s *FieldListContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldListContext) AllFieldInstance() []IFieldInstanceContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFieldInstanceContext)(nil)).Elem())
	var tst = make([]IFieldInstanceContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFieldInstanceContext)
		}
	}

	return tst
}

func (s *FieldListContext) FieldInstance(i int) IFieldInstanceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldInstanceContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFieldInstanceContext)
}

func (s *FieldListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DMLListener); ok {
		listenerT.EnterFieldList(s)
	}
}

func (s *FieldListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DMLListener); ok {
		listenerT.ExitFieldList(s)
	}
}

func (p *DMLParser) FieldList() (localctx IFieldListContext) {
	localctx = NewFieldListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, DMLParserRULE_fieldList)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(405)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la-36)&-(0x1f+1)) == 0 && ((1<<uint((_la-36)))&((1<<(DMLParserREQUIRED-36))|(1<<(DMLParserOPTIONAL-36))|(1<<(DMLParserREPEATED-36)))) != 0) {
		{
			p.SetState(404)
			p.FieldInstance()
		}

		p.SetState(407)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IFieldInstanceContext is an interface to support dynamic dispatch.
type IFieldInstanceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldInstanceContext differentiates from other interfaces.
	IsFieldInstanceContext()
}

type FieldInstanceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldInstanceContext() *FieldInstanceContext {
	var p = new(FieldInstanceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DMLParserRULE_fieldInstance
	return p
}

func (*FieldInstanceContext) IsFieldInstanceContext() {}

func NewFieldInstanceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldInstanceContext {
	var p = new(FieldInstanceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DMLParserRULE_fieldInstance

	return p
}

func (s *FieldInstanceContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldInstanceContext) ParameterModifier() IParameterModifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParameterModifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParameterModifierContext)
}

func (s *FieldInstanceContext) ID() antlr.TerminalNode {
	return s.GetToken(DMLParserID, 0)
}

func (s *FieldInstanceContext) SEMI() antlr.TerminalNode {
	return s.GetToken(DMLParserSEMI, 0)
}

func (s *FieldInstanceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldInstanceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldInstanceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DMLListener); ok {
		listenerT.EnterFieldInstance(s)
	}
}

func (s *FieldInstanceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DMLListener); ok {
		listenerT.ExitFieldInstance(s)
	}
}

func (p *DMLParser) FieldInstance() (localctx IFieldInstanceContext) {
	localctx = NewFieldInstanceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, DMLParserRULE_fieldInstance)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(409)
		p.ParameterModifier()
	}
	{
		p.SetState(410)
		p.Match(DMLParserID)
	}
	{
		p.SetState(411)
		p.Match(DMLParserSEMI)
	}

	return localctx
}

// IMethodReturnContext is an interface to support dynamic dispatch.
type IMethodReturnContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMethodReturnContext differentiates from other interfaces.
	IsMethodReturnContext()
}

type MethodReturnContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMethodReturnContext() *MethodReturnContext {
	var p = new(MethodReturnContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DMLParserRULE_methodReturn
	return p
}

func (*MethodReturnContext) IsMethodReturnContext() {}

func NewMethodReturnContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MethodReturnContext {
	var p = new(MethodReturnContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DMLParserRULE_methodReturn

	return p
}

func (s *MethodReturnContext) GetParser() antlr.Parser { return s.parser }

func (s *MethodReturnContext) RETURNS() antlr.TerminalNode {
	return s.GetToken(DMLParserRETURNS, 0)
}

func (s *MethodReturnContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(DMLParserLBRACE, 0)
}

func (s *MethodReturnContext) FieldList() IFieldListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldListContext)
}

func (s *MethodReturnContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(DMLParserRBRACE, 0)
}

func (s *MethodReturnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MethodReturnContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MethodReturnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DMLListener); ok {
		listenerT.EnterMethodReturn(s)
	}
}

func (s *MethodReturnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DMLListener); ok {
		listenerT.ExitMethodReturn(s)
	}
}

func (p *DMLParser) MethodReturn() (localctx IMethodReturnContext) {
	localctx = NewMethodReturnContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, DMLParserRULE_methodReturn)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(413)
		p.Match(DMLParserRETURNS)
	}
	{
		p.SetState(414)
		p.Match(DMLParserLBRACE)
	}
	{
		p.SetState(415)
		p.FieldList()
	}
	{
		p.SetState(416)
		p.Match(DMLParserRBRACE)
	}

	return localctx
}

// IPackageElementContext is an interface to support dynamic dispatch.
type IPackageElementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPackageElementContext differentiates from other interfaces.
	IsPackageElementContext()
}

type PackageElementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPackageElementContext() *PackageElementContext {
	var p = new(PackageElementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DMLParserRULE_packageElement
	return p
}

func (*PackageElementContext) IsPackageElementContext() {}

func NewPackageElementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PackageElementContext {
	var p = new(PackageElementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DMLParserRULE_packageElement

	return p
}

func (s *PackageElementContext) GetParser() antlr.Parser { return s.parser }

func (s *PackageElementContext) PACKAGE() antlr.TerminalNode {
	return s.GetToken(DMLParserPACKAGE, 0)
}

func (s *PackageElementContext) QualifiedId() IQualifiedIdContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQualifiedIdContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQualifiedIdContext)
}

func (s *PackageElementContext) SEMI() antlr.TerminalNode {
	return s.GetToken(DMLParserSEMI, 0)
}

func (s *PackageElementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PackageElementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PackageElementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DMLListener); ok {
		listenerT.EnterPackageElement(s)
	}
}

func (s *PackageElementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DMLListener); ok {
		listenerT.ExitPackageElement(s)
	}
}

func (p *DMLParser) PackageElement() (localctx IPackageElementContext) {
	localctx = NewPackageElementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, DMLParserRULE_packageElement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(418)
		p.Match(DMLParserPACKAGE)
	}
	{
		p.SetState(419)
		p.QualifiedId()
	}
	{
		p.SetState(420)
		p.Match(DMLParserSEMI)
	}

	return localctx
}

// IGoPackageElementContext is an interface to support dynamic dispatch.
type IGoPackageElementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGoPackageElementContext differentiates from other interfaces.
	IsGoPackageElementContext()
}

type GoPackageElementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGoPackageElementContext() *GoPackageElementContext {
	var p = new(GoPackageElementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DMLParserRULE_goPackageElement
	return p
}

func (*GoPackageElementContext) IsGoPackageElementContext() {}

func NewGoPackageElementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GoPackageElementContext {
	var p = new(GoPackageElementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DMLParserRULE_goPackageElement

	return p
}

func (s *GoPackageElementContext) GetParser() antlr.Parser { return s.parser }

func (s *GoPackageElementContext) GOPACKAGE() antlr.TerminalNode {
	return s.GetToken(DMLParserGOPACKAGE, 0)
}

func (s *GoPackageElementContext) GoPath() IGoPathContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGoPathContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGoPathContext)
}

func (s *GoPackageElementContext) SEMI() antlr.TerminalNode {
	return s.GetToken(DMLParserSEMI, 0)
}

func (s *GoPackageElementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GoPackageElementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GoPackageElementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DMLListener); ok {
		listenerT.EnterGoPackageElement(s)
	}
}

func (s *GoPackageElementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DMLListener); ok {
		listenerT.ExitGoPackageElement(s)
	}
}

func (p *DMLParser) GoPackageElement() (localctx IGoPackageElementContext) {
	localctx = NewGoPackageElementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, DMLParserRULE_goPackageElement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(422)
		p.Match(DMLParserGOPACKAGE)
	}
	{
		p.SetState(423)
		p.GoPath()
	}
	{
		p.SetState(424)
		p.Match(DMLParserSEMI)
	}

	return localctx
}

// IImportElementsContext is an interface to support dynamic dispatch.
type IImportElementsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsImportElementsContext differentiates from other interfaces.
	IsImportElementsContext()
}

type ImportElementsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImportElementsContext() *ImportElementsContext {
	var p = new(ImportElementsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DMLParserRULE_importElements
	return p
}

func (*ImportElementsContext) IsImportElementsContext() {}

func NewImportElementsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImportElementsContext {
	var p = new(ImportElementsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DMLParserRULE_importElements

	return p
}

func (s *ImportElementsContext) GetParser() antlr.Parser { return s.parser }

func (s *ImportElementsContext) IMPORT() antlr.TerminalNode {
	return s.GetToken(DMLParserIMPORT, 0)
}

func (s *ImportElementsContext) TEXT_STRING() antlr.TerminalNode {
	return s.GetToken(DMLParserTEXT_STRING, 0)
}

func (s *ImportElementsContext) SEMI() antlr.TerminalNode {
	return s.GetToken(DMLParserSEMI, 0)
}

func (s *ImportElementsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImportElementsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ImportElementsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DMLListener); ok {
		listenerT.EnterImportElements(s)
	}
}

func (s *ImportElementsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DMLListener); ok {
		listenerT.ExitImportElements(s)
	}
}

func (p *DMLParser) ImportElements() (localctx IImportElementsContext) {
	localctx = NewImportElementsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, DMLParserRULE_importElements)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(426)
		p.Match(DMLParserIMPORT)
	}
	{
		p.SetState(427)
		p.Match(DMLParserTEXT_STRING)
	}
	{
		p.SetState(428)
		p.Match(DMLParserSEMI)
	}

	return localctx
}
