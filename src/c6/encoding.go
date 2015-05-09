package c6

/*
import "golang.org/x/text/encoding"
import "golang.org/x/text/transform"

rInUTF8 := transform.NewReader(r, e.NewDecoder())
*/
var utf_8_bom []byte = []byte{0xEF, 0xBB, 0xBF}
var utf_16_bom_be []byte = []byte{0xFE, 0xFF}
var utf_16_bom_le []byte = []byte{0xFF, 0xFE}
var utf_32_bom_be []byte = []byte{0x00, 0x00, 0xFE, 0xFF}
var utf_32_bom_le []byte = []byte{0xFF, 0xFE, 0x00, 0x00}
var utf_7_bom_1 []byte = []byte{0x2B, 0x2F, 0x76, 0x38}
var utf_7_bom_2 []byte = []byte{0x2B, 0x2F, 0x76, 0x39}
var utf_7_bom_3 []byte = []byte{0x2B, 0x2F, 0x76, 0x2B}
var utf_7_bom_4 []byte = []byte{0x2B, 0x2F, 0x76, 0x2F}
var utf_7_bom_5 []byte = []byte{0x2B, 0x2F, 0x76, 0x38, 0x2D}
var utf_1_bom []byte = []byte{0xF7, 0x64, 0x4C}
var utf_ebcdic_bom []byte = []byte{0xDD, 0x73, 0x66, 0x73}
var scsu_bom []byte = []byte{0x0E, 0xFE, 0xFF}
var bocu_1_bom []byte = []byte{0xFB, 0xEE, 0x28}
var gb_18030_bom []byte = []byte{0x84, 0x31, 0x95, 0x33}