// DO NOT MODIFY: this file is autogenerated by makemysqlversions.go

package collations

type collver byte

const (
	collverInvalid    collver = 0
	collverMariaDB100 collver = 1 << 0
	collverMariaDB101 collver = 1 << 1
	collverMariaDB102 collver = 1 << 2
	collverMariaDB103 collver = 1 << 3
	collverMySQL56    collver = 1 << 4
	collverMySQL57    collver = 1 << 5
	collverMySQL80    collver = 1 << 6
)

func (v collver) String() string {
	switch v {
	case collverInvalid:
		return "Invalid"
	case collverMariaDB100:
		return "MariaDB100"
	case collverMariaDB101:
		return "MariaDB101"
	case collverMariaDB102:
		return "MariaDB102"
	case collverMariaDB103:
		return "MariaDB103"
	case collverMySQL56:
		return "MySQL56"
	case collverMySQL57:
		return "MySQL57"
	case collverMySQL80:
		return "MySQL80"
	default:
		panic("invalid version identifier")
	}
}

var globalVersionInfo = map[ID]struct {
	alias     map[collver]string
	isdefault collver
}{
	1:    {alias: map[collver]string{0b01111111: "big5_chinese_ci"}, isdefault: 0b01111111},
	2:    {alias: map[collver]string{0b01111111: "latin2_czech_cs"}, isdefault: 0b00000000},
	3:    {alias: map[collver]string{0b01111111: "dec8_swedish_ci"}, isdefault: 0b01111111},
	4:    {alias: map[collver]string{0b01111111: "cp850_general_ci"}, isdefault: 0b01111111},
	5:    {alias: map[collver]string{0b01111111: "latin1_german1_ci"}, isdefault: 0b00000000},
	6:    {alias: map[collver]string{0b01111111: "hp8_english_ci"}, isdefault: 0b01111111},
	7:    {alias: map[collver]string{0b01111111: "koi8r_general_ci"}, isdefault: 0b01111111},
	8:    {alias: map[collver]string{0b01111111: "latin1_swedish_ci"}, isdefault: 0b01111111},
	9:    {alias: map[collver]string{0b01111111: "latin2_general_ci"}, isdefault: 0b01111111},
	10:   {alias: map[collver]string{0b01111111: "swe7_swedish_ci"}, isdefault: 0b01111111},
	11:   {alias: map[collver]string{0b01111111: "ascii_general_ci"}, isdefault: 0b01111111},
	12:   {alias: map[collver]string{0b01111111: "ujis_japanese_ci"}, isdefault: 0b01111111},
	13:   {alias: map[collver]string{0b01111111: "sjis_japanese_ci"}, isdefault: 0b01111111},
	14:   {alias: map[collver]string{0b01111111: "cp1251_bulgarian_ci"}, isdefault: 0b00000000},
	15:   {alias: map[collver]string{0b01111111: "latin1_danish_ci"}, isdefault: 0b00000000},
	16:   {alias: map[collver]string{0b01111111: "hebrew_general_ci"}, isdefault: 0b01111111},
	18:   {alias: map[collver]string{0b01111111: "tis620_thai_ci"}, isdefault: 0b01111111},
	19:   {alias: map[collver]string{0b01111111: "euckr_korean_ci"}, isdefault: 0b01111111},
	20:   {alias: map[collver]string{0b01111111: "latin7_estonian_cs"}, isdefault: 0b00000000},
	21:   {alias: map[collver]string{0b01111111: "latin2_hungarian_ci"}, isdefault: 0b00000000},
	22:   {alias: map[collver]string{0b01111111: "koi8u_general_ci"}, isdefault: 0b01111111},
	23:   {alias: map[collver]string{0b01111111: "cp1251_ukrainian_ci"}, isdefault: 0b00000000},
	24:   {alias: map[collver]string{0b01111111: "gb2312_chinese_ci"}, isdefault: 0b01111111},
	25:   {alias: map[collver]string{0b01111111: "greek_general_ci"}, isdefault: 0b01111111},
	26:   {alias: map[collver]string{0b01111111: "cp1250_general_ci"}, isdefault: 0b01111111},
	27:   {alias: map[collver]string{0b01111111: "latin2_croatian_ci"}, isdefault: 0b00000000},
	28:   {alias: map[collver]string{0b01111111: "gbk_chinese_ci"}, isdefault: 0b01111111},
	29:   {alias: map[collver]string{0b01111111: "cp1257_lithuanian_ci"}, isdefault: 0b00000000},
	30:   {alias: map[collver]string{0b01111111: "latin5_turkish_ci"}, isdefault: 0b01111111},
	31:   {alias: map[collver]string{0b01111111: "latin1_german2_ci"}, isdefault: 0b00000000},
	32:   {alias: map[collver]string{0b01111111: "armscii8_general_ci"}, isdefault: 0b01111111},
	33:   {alias: map[collver]string{0b01111111: "utf8_general_ci"}, isdefault: 0b01111111},
	34:   {alias: map[collver]string{0b01111111: "cp1250_czech_cs"}, isdefault: 0b00000000},
	35:   {alias: map[collver]string{0b01111111: "ucs2_general_ci"}, isdefault: 0b01111111},
	36:   {alias: map[collver]string{0b01111111: "cp866_general_ci"}, isdefault: 0b01111111},
	37:   {alias: map[collver]string{0b01111111: "keybcs2_general_ci"}, isdefault: 0b01111111},
	38:   {alias: map[collver]string{0b01111111: "macce_general_ci"}, isdefault: 0b01111111},
	39:   {alias: map[collver]string{0b01111111: "macroman_general_ci"}, isdefault: 0b01111111},
	40:   {alias: map[collver]string{0b01111111: "cp852_general_ci"}, isdefault: 0b01111111},
	41:   {alias: map[collver]string{0b01111111: "latin7_general_ci"}, isdefault: 0b01111111},
	42:   {alias: map[collver]string{0b01111111: "latin7_general_cs"}, isdefault: 0b00000000},
	43:   {alias: map[collver]string{0b01111111: "macce_bin"}, isdefault: 0b00000000},
	44:   {alias: map[collver]string{0b01111111: "cp1250_croatian_ci"}, isdefault: 0b00000000},
	45:   {alias: map[collver]string{0b01111111: "utf8mb4_general_ci"}, isdefault: 0b00111111},
	46:   {alias: map[collver]string{0b01111111: "utf8mb4_bin"}, isdefault: 0b00000000},
	47:   {alias: map[collver]string{0b01111111: "latin1_bin"}, isdefault: 0b00000000},
	48:   {alias: map[collver]string{0b01111111: "latin1_general_ci"}, isdefault: 0b00000000},
	49:   {alias: map[collver]string{0b01111111: "latin1_general_cs"}, isdefault: 0b00000000},
	50:   {alias: map[collver]string{0b01111111: "cp1251_bin"}, isdefault: 0b00000000},
	51:   {alias: map[collver]string{0b01111111: "cp1251_general_ci"}, isdefault: 0b01111111},
	52:   {alias: map[collver]string{0b01111111: "cp1251_general_cs"}, isdefault: 0b00000000},
	53:   {alias: map[collver]string{0b01111111: "macroman_bin"}, isdefault: 0b00000000},
	54:   {alias: map[collver]string{0b01111111: "utf16_general_ci"}, isdefault: 0b01111111},
	55:   {alias: map[collver]string{0b01111111: "utf16_bin"}, isdefault: 0b00000000},
	56:   {alias: map[collver]string{0b01111111: "utf16le_general_ci"}, isdefault: 0b01111111},
	57:   {alias: map[collver]string{0b01111111: "cp1256_general_ci"}, isdefault: 0b01111111},
	58:   {alias: map[collver]string{0b01111111: "cp1257_bin"}, isdefault: 0b00000000},
	59:   {alias: map[collver]string{0b01111111: "cp1257_general_ci"}, isdefault: 0b01111111},
	60:   {alias: map[collver]string{0b01111111: "utf32_general_ci"}, isdefault: 0b01111111},
	61:   {alias: map[collver]string{0b01111111: "utf32_bin"}, isdefault: 0b00000000},
	62:   {alias: map[collver]string{0b01111111: "utf16le_bin"}, isdefault: 0b00000000},
	63:   {alias: map[collver]string{0b01111111: "binary"}, isdefault: 0b01111111},
	64:   {alias: map[collver]string{0b01111111: "armscii8_bin"}, isdefault: 0b00000000},
	65:   {alias: map[collver]string{0b01111111: "ascii_bin"}, isdefault: 0b00000000},
	66:   {alias: map[collver]string{0b01111111: "cp1250_bin"}, isdefault: 0b00000000},
	67:   {alias: map[collver]string{0b01111111: "cp1256_bin"}, isdefault: 0b00000000},
	68:   {alias: map[collver]string{0b01111111: "cp866_bin"}, isdefault: 0b00000000},
	69:   {alias: map[collver]string{0b01111111: "dec8_bin"}, isdefault: 0b00000000},
	70:   {alias: map[collver]string{0b01111111: "greek_bin"}, isdefault: 0b00000000},
	71:   {alias: map[collver]string{0b01111111: "hebrew_bin"}, isdefault: 0b00000000},
	72:   {alias: map[collver]string{0b01111111: "hp8_bin"}, isdefault: 0b00000000},
	73:   {alias: map[collver]string{0b01111111: "keybcs2_bin"}, isdefault: 0b00000000},
	74:   {alias: map[collver]string{0b01111111: "koi8r_bin"}, isdefault: 0b00000000},
	75:   {alias: map[collver]string{0b01111111: "koi8u_bin"}, isdefault: 0b00000000},
	76:   {alias: map[collver]string{0b01000000: "utf8_tolower_ci"}, isdefault: 0b00000000},
	77:   {alias: map[collver]string{0b01111111: "latin2_bin"}, isdefault: 0b00000000},
	78:   {alias: map[collver]string{0b01111111: "latin5_bin"}, isdefault: 0b00000000},
	79:   {alias: map[collver]string{0b01111111: "latin7_bin"}, isdefault: 0b00000000},
	80:   {alias: map[collver]string{0b01111111: "cp850_bin"}, isdefault: 0b00000000},
	81:   {alias: map[collver]string{0b01111111: "cp852_bin"}, isdefault: 0b00000000},
	82:   {alias: map[collver]string{0b01111111: "swe7_bin"}, isdefault: 0b00000000},
	83:   {alias: map[collver]string{0b01111111: "utf8_bin"}, isdefault: 0b00000000},
	84:   {alias: map[collver]string{0b01111111: "big5_bin"}, isdefault: 0b00000000},
	85:   {alias: map[collver]string{0b01111111: "euckr_bin"}, isdefault: 0b00000000},
	86:   {alias: map[collver]string{0b01111111: "gb2312_bin"}, isdefault: 0b00000000},
	87:   {alias: map[collver]string{0b01111111: "gbk_bin"}, isdefault: 0b00000000},
	88:   {alias: map[collver]string{0b01111111: "sjis_bin"}, isdefault: 0b00000000},
	89:   {alias: map[collver]string{0b01111111: "tis620_bin"}, isdefault: 0b00000000},
	90:   {alias: map[collver]string{0b01111111: "ucs2_bin"}, isdefault: 0b00000000},
	91:   {alias: map[collver]string{0b01111111: "ujis_bin"}, isdefault: 0b00000000},
	92:   {alias: map[collver]string{0b01111111: "geostd8_general_ci"}, isdefault: 0b01111111},
	93:   {alias: map[collver]string{0b01111111: "geostd8_bin"}, isdefault: 0b00000000},
	94:   {alias: map[collver]string{0b01111111: "latin1_spanish_ci"}, isdefault: 0b00000000},
	95:   {alias: map[collver]string{0b01111111: "cp932_japanese_ci"}, isdefault: 0b01111111},
	96:   {alias: map[collver]string{0b01111111: "cp932_bin"}, isdefault: 0b00000000},
	97:   {alias: map[collver]string{0b01111111: "eucjpms_japanese_ci"}, isdefault: 0b01111111},
	98:   {alias: map[collver]string{0b01111111: "eucjpms_bin"}, isdefault: 0b00000000},
	99:   {alias: map[collver]string{0b01111111: "cp1250_polish_ci"}, isdefault: 0b00000000},
	101:  {alias: map[collver]string{0b01111111: "utf16_unicode_ci"}, isdefault: 0b00000000},
	102:  {alias: map[collver]string{0b01111111: "utf16_icelandic_ci"}, isdefault: 0b00000000},
	103:  {alias: map[collver]string{0b01111111: "utf16_latvian_ci"}, isdefault: 0b00000000},
	104:  {alias: map[collver]string{0b01111111: "utf16_romanian_ci"}, isdefault: 0b00000000},
	105:  {alias: map[collver]string{0b01111111: "utf16_slovenian_ci"}, isdefault: 0b00000000},
	106:  {alias: map[collver]string{0b01111111: "utf16_polish_ci"}, isdefault: 0b00000000},
	107:  {alias: map[collver]string{0b01111111: "utf16_estonian_ci"}, isdefault: 0b00000000},
	108:  {alias: map[collver]string{0b01111111: "utf16_spanish_ci"}, isdefault: 0b00000000},
	109:  {alias: map[collver]string{0b01111111: "utf16_swedish_ci"}, isdefault: 0b00000000},
	110:  {alias: map[collver]string{0b01111111: "utf16_turkish_ci"}, isdefault: 0b00000000},
	111:  {alias: map[collver]string{0b01111111: "utf16_czech_ci"}, isdefault: 0b00000000},
	112:  {alias: map[collver]string{0b01111111: "utf16_danish_ci"}, isdefault: 0b00000000},
	113:  {alias: map[collver]string{0b01111111: "utf16_lithuanian_ci"}, isdefault: 0b00000000},
	114:  {alias: map[collver]string{0b01111111: "utf16_slovak_ci"}, isdefault: 0b00000000},
	115:  {alias: map[collver]string{0b01111111: "utf16_spanish2_ci"}, isdefault: 0b00000000},
	116:  {alias: map[collver]string{0b01111111: "utf16_roman_ci"}, isdefault: 0b00000000},
	117:  {alias: map[collver]string{0b01111111: "utf16_persian_ci"}, isdefault: 0b00000000},
	118:  {alias: map[collver]string{0b01111111: "utf16_esperanto_ci"}, isdefault: 0b00000000},
	119:  {alias: map[collver]string{0b01111111: "utf16_hungarian_ci"}, isdefault: 0b00000000},
	120:  {alias: map[collver]string{0b01111111: "utf16_sinhala_ci"}, isdefault: 0b00000000},
	121:  {alias: map[collver]string{0b01111111: "utf16_german2_ci"}, isdefault: 0b00000000},
	122:  {alias: map[collver]string{0b01110000: "utf16_croatian_ci", 0b00001111: "utf16_croatian_mysql561_ci"}, isdefault: 0b00000000},
	123:  {alias: map[collver]string{0b01111111: "utf16_unicode_520_ci"}, isdefault: 0b00000000},
	124:  {alias: map[collver]string{0b01111111: "utf16_vietnamese_ci"}, isdefault: 0b00000000},
	128:  {alias: map[collver]string{0b01111111: "ucs2_unicode_ci"}, isdefault: 0b00000000},
	129:  {alias: map[collver]string{0b01111111: "ucs2_icelandic_ci"}, isdefault: 0b00000000},
	130:  {alias: map[collver]string{0b01111111: "ucs2_latvian_ci"}, isdefault: 0b00000000},
	131:  {alias: map[collver]string{0b01111111: "ucs2_romanian_ci"}, isdefault: 0b00000000},
	132:  {alias: map[collver]string{0b01111111: "ucs2_slovenian_ci"}, isdefault: 0b00000000},
	133:  {alias: map[collver]string{0b01111111: "ucs2_polish_ci"}, isdefault: 0b00000000},
	134:  {alias: map[collver]string{0b01111111: "ucs2_estonian_ci"}, isdefault: 0b00000000},
	135:  {alias: map[collver]string{0b01111111: "ucs2_spanish_ci"}, isdefault: 0b00000000},
	136:  {alias: map[collver]string{0b01111111: "ucs2_swedish_ci"}, isdefault: 0b00000000},
	137:  {alias: map[collver]string{0b01111111: "ucs2_turkish_ci"}, isdefault: 0b00000000},
	138:  {alias: map[collver]string{0b01111111: "ucs2_czech_ci"}, isdefault: 0b00000000},
	139:  {alias: map[collver]string{0b01111111: "ucs2_danish_ci"}, isdefault: 0b00000000},
	140:  {alias: map[collver]string{0b01111111: "ucs2_lithuanian_ci"}, isdefault: 0b00000000},
	141:  {alias: map[collver]string{0b01111111: "ucs2_slovak_ci"}, isdefault: 0b00000000},
	142:  {alias: map[collver]string{0b01111111: "ucs2_spanish2_ci"}, isdefault: 0b00000000},
	143:  {alias: map[collver]string{0b01111111: "ucs2_roman_ci"}, isdefault: 0b00000000},
	144:  {alias: map[collver]string{0b01111111: "ucs2_persian_ci"}, isdefault: 0b00000000},
	145:  {alias: map[collver]string{0b01111111: "ucs2_esperanto_ci"}, isdefault: 0b00000000},
	146:  {alias: map[collver]string{0b01111111: "ucs2_hungarian_ci"}, isdefault: 0b00000000},
	147:  {alias: map[collver]string{0b01111111: "ucs2_sinhala_ci"}, isdefault: 0b00000000},
	148:  {alias: map[collver]string{0b01111111: "ucs2_german2_ci"}, isdefault: 0b00000000},
	149:  {alias: map[collver]string{0b01110000: "ucs2_croatian_ci", 0b00001111: "ucs2_croatian_mysql561_ci"}, isdefault: 0b00000000},
	150:  {alias: map[collver]string{0b01111111: "ucs2_unicode_520_ci"}, isdefault: 0b00000000},
	151:  {alias: map[collver]string{0b01111111: "ucs2_vietnamese_ci"}, isdefault: 0b00000000},
	159:  {alias: map[collver]string{0b01111111: "ucs2_general_mysql500_ci"}, isdefault: 0b00000000},
	160:  {alias: map[collver]string{0b01111111: "utf32_unicode_ci"}, isdefault: 0b00000000},
	161:  {alias: map[collver]string{0b01111111: "utf32_icelandic_ci"}, isdefault: 0b00000000},
	162:  {alias: map[collver]string{0b01111111: "utf32_latvian_ci"}, isdefault: 0b00000000},
	163:  {alias: map[collver]string{0b01111111: "utf32_romanian_ci"}, isdefault: 0b00000000},
	164:  {alias: map[collver]string{0b01111111: "utf32_slovenian_ci"}, isdefault: 0b00000000},
	165:  {alias: map[collver]string{0b01111111: "utf32_polish_ci"}, isdefault: 0b00000000},
	166:  {alias: map[collver]string{0b01111111: "utf32_estonian_ci"}, isdefault: 0b00000000},
	167:  {alias: map[collver]string{0b01111111: "utf32_spanish_ci"}, isdefault: 0b00000000},
	168:  {alias: map[collver]string{0b01111111: "utf32_swedish_ci"}, isdefault: 0b00000000},
	169:  {alias: map[collver]string{0b01111111: "utf32_turkish_ci"}, isdefault: 0b00000000},
	170:  {alias: map[collver]string{0b01111111: "utf32_czech_ci"}, isdefault: 0b00000000},
	171:  {alias: map[collver]string{0b01111111: "utf32_danish_ci"}, isdefault: 0b00000000},
	172:  {alias: map[collver]string{0b01111111: "utf32_lithuanian_ci"}, isdefault: 0b00000000},
	173:  {alias: map[collver]string{0b01111111: "utf32_slovak_ci"}, isdefault: 0b00000000},
	174:  {alias: map[collver]string{0b01111111: "utf32_spanish2_ci"}, isdefault: 0b00000000},
	175:  {alias: map[collver]string{0b01111111: "utf32_roman_ci"}, isdefault: 0b00000000},
	176:  {alias: map[collver]string{0b01111111: "utf32_persian_ci"}, isdefault: 0b00000000},
	177:  {alias: map[collver]string{0b01111111: "utf32_esperanto_ci"}, isdefault: 0b00000000},
	178:  {alias: map[collver]string{0b01111111: "utf32_hungarian_ci"}, isdefault: 0b00000000},
	179:  {alias: map[collver]string{0b01111111: "utf32_sinhala_ci"}, isdefault: 0b00000000},
	180:  {alias: map[collver]string{0b01111111: "utf32_german2_ci"}, isdefault: 0b00000000},
	181:  {alias: map[collver]string{0b01110000: "utf32_croatian_ci", 0b00001111: "utf32_croatian_mysql561_ci"}, isdefault: 0b00000000},
	182:  {alias: map[collver]string{0b01111111: "utf32_unicode_520_ci"}, isdefault: 0b00000000},
	183:  {alias: map[collver]string{0b01111111: "utf32_vietnamese_ci"}, isdefault: 0b00000000},
	192:  {alias: map[collver]string{0b01111111: "utf8_unicode_ci"}, isdefault: 0b00000000},
	193:  {alias: map[collver]string{0b01111111: "utf8_icelandic_ci"}, isdefault: 0b00000000},
	194:  {alias: map[collver]string{0b01111111: "utf8_latvian_ci"}, isdefault: 0b00000000},
	195:  {alias: map[collver]string{0b01111111: "utf8_romanian_ci"}, isdefault: 0b00000000},
	196:  {alias: map[collver]string{0b01111111: "utf8_slovenian_ci"}, isdefault: 0b00000000},
	197:  {alias: map[collver]string{0b01111111: "utf8_polish_ci"}, isdefault: 0b00000000},
	198:  {alias: map[collver]string{0b01111111: "utf8_estonian_ci"}, isdefault: 0b00000000},
	199:  {alias: map[collver]string{0b01111111: "utf8_spanish_ci"}, isdefault: 0b00000000},
	200:  {alias: map[collver]string{0b01111111: "utf8_swedish_ci"}, isdefault: 0b00000000},
	201:  {alias: map[collver]string{0b01111111: "utf8_turkish_ci"}, isdefault: 0b00000000},
	202:  {alias: map[collver]string{0b01111111: "utf8_czech_ci"}, isdefault: 0b00000000},
	203:  {alias: map[collver]string{0b01111111: "utf8_danish_ci"}, isdefault: 0b00000000},
	204:  {alias: map[collver]string{0b01111111: "utf8_lithuanian_ci"}, isdefault: 0b00000000},
	205:  {alias: map[collver]string{0b01111111: "utf8_slovak_ci"}, isdefault: 0b00000000},
	206:  {alias: map[collver]string{0b01111111: "utf8_spanish2_ci"}, isdefault: 0b00000000},
	207:  {alias: map[collver]string{0b01111111: "utf8_roman_ci"}, isdefault: 0b00000000},
	208:  {alias: map[collver]string{0b01111111: "utf8_persian_ci"}, isdefault: 0b00000000},
	209:  {alias: map[collver]string{0b01111111: "utf8_esperanto_ci"}, isdefault: 0b00000000},
	210:  {alias: map[collver]string{0b01111111: "utf8_hungarian_ci"}, isdefault: 0b00000000},
	211:  {alias: map[collver]string{0b01111111: "utf8_sinhala_ci"}, isdefault: 0b00000000},
	212:  {alias: map[collver]string{0b01111111: "utf8_german2_ci"}, isdefault: 0b00000000},
	213:  {alias: map[collver]string{0b01110000: "utf8_croatian_ci", 0b00001111: "utf8_croatian_mysql561_ci"}, isdefault: 0b00000000},
	214:  {alias: map[collver]string{0b01111111: "utf8_unicode_520_ci"}, isdefault: 0b00000000},
	215:  {alias: map[collver]string{0b01111111: "utf8_vietnamese_ci"}, isdefault: 0b00000000},
	223:  {alias: map[collver]string{0b01111111: "utf8_general_mysql500_ci"}, isdefault: 0b00000000},
	224:  {alias: map[collver]string{0b01111111: "utf8mb4_unicode_ci"}, isdefault: 0b00000000},
	225:  {alias: map[collver]string{0b01111111: "utf8mb4_icelandic_ci"}, isdefault: 0b00000000},
	226:  {alias: map[collver]string{0b01111111: "utf8mb4_latvian_ci"}, isdefault: 0b00000000},
	227:  {alias: map[collver]string{0b01111111: "utf8mb4_romanian_ci"}, isdefault: 0b00000000},
	228:  {alias: map[collver]string{0b01111111: "utf8mb4_slovenian_ci"}, isdefault: 0b00000000},
	229:  {alias: map[collver]string{0b01111111: "utf8mb4_polish_ci"}, isdefault: 0b00000000},
	230:  {alias: map[collver]string{0b01111111: "utf8mb4_estonian_ci"}, isdefault: 0b00000000},
	231:  {alias: map[collver]string{0b01111111: "utf8mb4_spanish_ci"}, isdefault: 0b00000000},
	232:  {alias: map[collver]string{0b01111111: "utf8mb4_swedish_ci"}, isdefault: 0b00000000},
	233:  {alias: map[collver]string{0b01111111: "utf8mb4_turkish_ci"}, isdefault: 0b00000000},
	234:  {alias: map[collver]string{0b01111111: "utf8mb4_czech_ci"}, isdefault: 0b00000000},
	235:  {alias: map[collver]string{0b01111111: "utf8mb4_danish_ci"}, isdefault: 0b00000000},
	236:  {alias: map[collver]string{0b01111111: "utf8mb4_lithuanian_ci"}, isdefault: 0b00000000},
	237:  {alias: map[collver]string{0b01111111: "utf8mb4_slovak_ci"}, isdefault: 0b00000000},
	238:  {alias: map[collver]string{0b01111111: "utf8mb4_spanish2_ci"}, isdefault: 0b00000000},
	239:  {alias: map[collver]string{0b01111111: "utf8mb4_roman_ci"}, isdefault: 0b00000000},
	240:  {alias: map[collver]string{0b01111111: "utf8mb4_persian_ci"}, isdefault: 0b00000000},
	241:  {alias: map[collver]string{0b01111111: "utf8mb4_esperanto_ci"}, isdefault: 0b00000000},
	242:  {alias: map[collver]string{0b01111111: "utf8mb4_hungarian_ci"}, isdefault: 0b00000000},
	243:  {alias: map[collver]string{0b01111111: "utf8mb4_sinhala_ci"}, isdefault: 0b00000000},
	244:  {alias: map[collver]string{0b01111111: "utf8mb4_german2_ci"}, isdefault: 0b00000000},
	245:  {alias: map[collver]string{0b01110000: "utf8mb4_croatian_ci", 0b00001111: "utf8mb4_croatian_mysql561_ci"}, isdefault: 0b00000000},
	246:  {alias: map[collver]string{0b01111111: "utf8mb4_unicode_520_ci"}, isdefault: 0b00000000},
	247:  {alias: map[collver]string{0b01111111: "utf8mb4_vietnamese_ci"}, isdefault: 0b00000000},
	248:  {alias: map[collver]string{0b01100000: "gb18030_chinese_ci"}, isdefault: 0b01100000},
	249:  {alias: map[collver]string{0b01100000: "gb18030_bin"}, isdefault: 0b00000000},
	250:  {alias: map[collver]string{0b01100000: "gb18030_unicode_520_ci"}, isdefault: 0b00000000},
	255:  {alias: map[collver]string{0b01000000: "utf8mb4_0900_ai_ci"}, isdefault: 0b01000000},
	256:  {alias: map[collver]string{0b01000000: "utf8mb4_de_pb_0900_ai_ci"}, isdefault: 0b00000000},
	257:  {alias: map[collver]string{0b01000000: "utf8mb4_is_0900_ai_ci"}, isdefault: 0b00000000},
	258:  {alias: map[collver]string{0b01000000: "utf8mb4_lv_0900_ai_ci"}, isdefault: 0b00000000},
	259:  {alias: map[collver]string{0b01000000: "utf8mb4_ro_0900_ai_ci"}, isdefault: 0b00000000},
	260:  {alias: map[collver]string{0b01000000: "utf8mb4_sl_0900_ai_ci"}, isdefault: 0b00000000},
	261:  {alias: map[collver]string{0b01000000: "utf8mb4_pl_0900_ai_ci"}, isdefault: 0b00000000},
	262:  {alias: map[collver]string{0b01000000: "utf8mb4_et_0900_ai_ci"}, isdefault: 0b00000000},
	263:  {alias: map[collver]string{0b01000000: "utf8mb4_es_0900_ai_ci"}, isdefault: 0b00000000},
	264:  {alias: map[collver]string{0b01000000: "utf8mb4_sv_0900_ai_ci"}, isdefault: 0b00000000},
	265:  {alias: map[collver]string{0b01000000: "utf8mb4_tr_0900_ai_ci"}, isdefault: 0b00000000},
	266:  {alias: map[collver]string{0b01000000: "utf8mb4_cs_0900_ai_ci"}, isdefault: 0b00000000},
	267:  {alias: map[collver]string{0b01000000: "utf8mb4_da_0900_ai_ci"}, isdefault: 0b00000000},
	268:  {alias: map[collver]string{0b01000000: "utf8mb4_lt_0900_ai_ci"}, isdefault: 0b00000000},
	269:  {alias: map[collver]string{0b01000000: "utf8mb4_sk_0900_ai_ci"}, isdefault: 0b00000000},
	270:  {alias: map[collver]string{0b01000000: "utf8mb4_es_trad_0900_ai_ci"}, isdefault: 0b00000000},
	271:  {alias: map[collver]string{0b01000000: "utf8mb4_la_0900_ai_ci"}, isdefault: 0b00000000},
	273:  {alias: map[collver]string{0b01000000: "utf8mb4_eo_0900_ai_ci"}, isdefault: 0b00000000},
	274:  {alias: map[collver]string{0b01000000: "utf8mb4_hu_0900_ai_ci"}, isdefault: 0b00000000},
	275:  {alias: map[collver]string{0b01000000: "utf8mb4_hr_0900_ai_ci"}, isdefault: 0b00000000},
	277:  {alias: map[collver]string{0b01000000: "utf8mb4_vi_0900_ai_ci"}, isdefault: 0b00000000},
	278:  {alias: map[collver]string{0b01000000: "utf8mb4_0900_as_cs"}, isdefault: 0b00000000},
	279:  {alias: map[collver]string{0b01000000: "utf8mb4_de_pb_0900_as_cs"}, isdefault: 0b00000000},
	280:  {alias: map[collver]string{0b01000000: "utf8mb4_is_0900_as_cs"}, isdefault: 0b00000000},
	281:  {alias: map[collver]string{0b01000000: "utf8mb4_lv_0900_as_cs"}, isdefault: 0b00000000},
	282:  {alias: map[collver]string{0b01000000: "utf8mb4_ro_0900_as_cs"}, isdefault: 0b00000000},
	283:  {alias: map[collver]string{0b01000000: "utf8mb4_sl_0900_as_cs"}, isdefault: 0b00000000},
	284:  {alias: map[collver]string{0b01000000: "utf8mb4_pl_0900_as_cs"}, isdefault: 0b00000000},
	285:  {alias: map[collver]string{0b01000000: "utf8mb4_et_0900_as_cs"}, isdefault: 0b00000000},
	286:  {alias: map[collver]string{0b01000000: "utf8mb4_es_0900_as_cs"}, isdefault: 0b00000000},
	287:  {alias: map[collver]string{0b01000000: "utf8mb4_sv_0900_as_cs"}, isdefault: 0b00000000},
	288:  {alias: map[collver]string{0b01000000: "utf8mb4_tr_0900_as_cs"}, isdefault: 0b00000000},
	289:  {alias: map[collver]string{0b01000000: "utf8mb4_cs_0900_as_cs"}, isdefault: 0b00000000},
	290:  {alias: map[collver]string{0b01000000: "utf8mb4_da_0900_as_cs"}, isdefault: 0b00000000},
	291:  {alias: map[collver]string{0b01000000: "utf8mb4_lt_0900_as_cs"}, isdefault: 0b00000000},
	292:  {alias: map[collver]string{0b01000000: "utf8mb4_sk_0900_as_cs"}, isdefault: 0b00000000},
	293:  {alias: map[collver]string{0b01000000: "utf8mb4_es_trad_0900_as_cs"}, isdefault: 0b00000000},
	294:  {alias: map[collver]string{0b01000000: "utf8mb4_la_0900_as_cs"}, isdefault: 0b00000000},
	296:  {alias: map[collver]string{0b01000000: "utf8mb4_eo_0900_as_cs"}, isdefault: 0b00000000},
	297:  {alias: map[collver]string{0b01000000: "utf8mb4_hu_0900_as_cs"}, isdefault: 0b00000000},
	298:  {alias: map[collver]string{0b01000000: "utf8mb4_hr_0900_as_cs"}, isdefault: 0b00000000},
	300:  {alias: map[collver]string{0b01000000: "utf8mb4_vi_0900_as_cs"}, isdefault: 0b00000000},
	303:  {alias: map[collver]string{0b01000000: "utf8mb4_ja_0900_as_cs"}, isdefault: 0b00000000},
	304:  {alias: map[collver]string{0b01000000: "utf8mb4_ja_0900_as_cs_ks"}, isdefault: 0b00000000},
	305:  {alias: map[collver]string{0b01000000: "utf8mb4_0900_as_ci"}, isdefault: 0b00000000},
	306:  {alias: map[collver]string{0b01000000: "utf8mb4_ru_0900_ai_ci"}, isdefault: 0b00000000},
	307:  {alias: map[collver]string{0b01000000: "utf8mb4_ru_0900_as_cs"}, isdefault: 0b00000000},
	308:  {alias: map[collver]string{0b01000000: "utf8mb4_zh_0900_as_cs"}, isdefault: 0b00000000},
	309:  {alias: map[collver]string{0b01000000: "utf8mb4_0900_bin"}, isdefault: 0b00000000},
	576:  {alias: map[collver]string{0b00001111: "utf8_croatian_ci"}, isdefault: 0b00000000},
	577:  {alias: map[collver]string{0b00001111: "utf8_myanmar_ci"}, isdefault: 0b00000000},
	578:  {alias: map[collver]string{0b00001110: "utf8_thai_520_w2"}, isdefault: 0b00000000},
	608:  {alias: map[collver]string{0b00001111: "utf8mb4_croatian_ci"}, isdefault: 0b00000000},
	609:  {alias: map[collver]string{0b00001111: "utf8mb4_myanmar_ci"}, isdefault: 0b00000000},
	610:  {alias: map[collver]string{0b00001110: "utf8mb4_thai_520_w2"}, isdefault: 0b00000000},
	640:  {alias: map[collver]string{0b00001111: "ucs2_croatian_ci"}, isdefault: 0b00000000},
	641:  {alias: map[collver]string{0b00001111: "ucs2_myanmar_ci"}, isdefault: 0b00000000},
	642:  {alias: map[collver]string{0b00001110: "ucs2_thai_520_w2"}, isdefault: 0b00000000},
	672:  {alias: map[collver]string{0b00001111: "utf16_croatian_ci"}, isdefault: 0b00000000},
	673:  {alias: map[collver]string{0b00001111: "utf16_myanmar_ci"}, isdefault: 0b00000000},
	674:  {alias: map[collver]string{0b00001110: "utf16_thai_520_w2"}, isdefault: 0b00000000},
	736:  {alias: map[collver]string{0b00001111: "utf32_croatian_ci"}, isdefault: 0b00000000},
	737:  {alias: map[collver]string{0b00001111: "utf32_myanmar_ci"}, isdefault: 0b00000000},
	738:  {alias: map[collver]string{0b00001110: "utf32_thai_520_w2"}, isdefault: 0b00000000},
	1025: {alias: map[collver]string{0b00001100: "big5_chinese_nopad_ci"}, isdefault: 0b00000000},
	1027: {alias: map[collver]string{0b00001100: "dec8_swedish_nopad_ci"}, isdefault: 0b00000000},
	1028: {alias: map[collver]string{0b00001100: "cp850_general_nopad_ci"}, isdefault: 0b00000000},
	1030: {alias: map[collver]string{0b00001100: "hp8_english_nopad_ci"}, isdefault: 0b00000000},
	1031: {alias: map[collver]string{0b00001100: "koi8r_general_nopad_ci"}, isdefault: 0b00000000},
	1032: {alias: map[collver]string{0b00001100: "latin1_swedish_nopad_ci"}, isdefault: 0b00000000},
	1033: {alias: map[collver]string{0b00001100: "latin2_general_nopad_ci"}, isdefault: 0b00000000},
	1034: {alias: map[collver]string{0b00001100: "swe7_swedish_nopad_ci"}, isdefault: 0b00000000},
	1035: {alias: map[collver]string{0b00001100: "ascii_general_nopad_ci"}, isdefault: 0b00000000},
	1036: {alias: map[collver]string{0b00001100: "ujis_japanese_nopad_ci"}, isdefault: 0b00000000},
	1037: {alias: map[collver]string{0b00001100: "sjis_japanese_nopad_ci"}, isdefault: 0b00000000},
	1040: {alias: map[collver]string{0b00001100: "hebrew_general_nopad_ci"}, isdefault: 0b00000000},
	1042: {alias: map[collver]string{0b00001100: "tis620_thai_nopad_ci"}, isdefault: 0b00000000},
	1043: {alias: map[collver]string{0b00001100: "euckr_korean_nopad_ci"}, isdefault: 0b00000000},
	1046: {alias: map[collver]string{0b00001100: "koi8u_general_nopad_ci"}, isdefault: 0b00000000},
	1048: {alias: map[collver]string{0b00001100: "gb2312_chinese_nopad_ci"}, isdefault: 0b00000000},
	1049: {alias: map[collver]string{0b00001100: "greek_general_nopad_ci"}, isdefault: 0b00000000},
	1050: {alias: map[collver]string{0b00001100: "cp1250_general_nopad_ci"}, isdefault: 0b00000000},
	1052: {alias: map[collver]string{0b00001100: "gbk_chinese_nopad_ci"}, isdefault: 0b00000000},
	1054: {alias: map[collver]string{0b00001100: "latin5_turkish_nopad_ci"}, isdefault: 0b00000000},
	1056: {alias: map[collver]string{0b00001100: "armscii8_general_nopad_ci"}, isdefault: 0b00000000},
	1057: {alias: map[collver]string{0b00001100: "utf8_general_nopad_ci"}, isdefault: 0b00000000},
	1059: {alias: map[collver]string{0b00001100: "ucs2_general_nopad_ci"}, isdefault: 0b00000000},
	1060: {alias: map[collver]string{0b00001100: "cp866_general_nopad_ci"}, isdefault: 0b00000000},
	1061: {alias: map[collver]string{0b00001100: "keybcs2_general_nopad_ci"}, isdefault: 0b00000000},
	1062: {alias: map[collver]string{0b00001100: "macce_general_nopad_ci"}, isdefault: 0b00000000},
	1063: {alias: map[collver]string{0b00001100: "macroman_general_nopad_ci"}, isdefault: 0b00000000},
	1064: {alias: map[collver]string{0b00001100: "cp852_general_nopad_ci"}, isdefault: 0b00000000},
	1065: {alias: map[collver]string{0b00001100: "latin7_general_nopad_ci"}, isdefault: 0b00000000},
	1067: {alias: map[collver]string{0b00001100: "macce_nopad_bin"}, isdefault: 0b00000000},
	1069: {alias: map[collver]string{0b00001100: "utf8mb4_general_nopad_ci"}, isdefault: 0b00000000},
	1070: {alias: map[collver]string{0b00001100: "utf8mb4_nopad_bin"}, isdefault: 0b00000000},
	1071: {alias: map[collver]string{0b00001100: "latin1_nopad_bin"}, isdefault: 0b00000000},
	1074: {alias: map[collver]string{0b00001100: "cp1251_nopad_bin"}, isdefault: 0b00000000},
	1075: {alias: map[collver]string{0b00001100: "cp1251_general_nopad_ci"}, isdefault: 0b00000000},
	1077: {alias: map[collver]string{0b00001100: "macroman_nopad_bin"}, isdefault: 0b00000000},
	1078: {alias: map[collver]string{0b00001100: "utf16_general_nopad_ci"}, isdefault: 0b00000000},
	1079: {alias: map[collver]string{0b00001100: "utf16_nopad_bin"}, isdefault: 0b00000000},
	1080: {alias: map[collver]string{0b00001100: "utf16le_general_nopad_ci"}, isdefault: 0b00000000},
	1081: {alias: map[collver]string{0b00001100: "cp1256_general_nopad_ci"}, isdefault: 0b00000000},
	1082: {alias: map[collver]string{0b00001100: "cp1257_nopad_bin"}, isdefault: 0b00000000},
	1083: {alias: map[collver]string{0b00001100: "cp1257_general_nopad_ci"}, isdefault: 0b00000000},
	1084: {alias: map[collver]string{0b00001100: "utf32_general_nopad_ci"}, isdefault: 0b00000000},
	1085: {alias: map[collver]string{0b00001100: "utf32_nopad_bin"}, isdefault: 0b00000000},
	1086: {alias: map[collver]string{0b00001100: "utf16le_nopad_bin"}, isdefault: 0b00000000},
	1088: {alias: map[collver]string{0b00001100: "armscii8_nopad_bin"}, isdefault: 0b00000000},
	1089: {alias: map[collver]string{0b00001100: "ascii_nopad_bin"}, isdefault: 0b00000000},
	1090: {alias: map[collver]string{0b00001100: "cp1250_nopad_bin"}, isdefault: 0b00000000},
	1091: {alias: map[collver]string{0b00001100: "cp1256_nopad_bin"}, isdefault: 0b00000000},
	1092: {alias: map[collver]string{0b00001100: "cp866_nopad_bin"}, isdefault: 0b00000000},
	1093: {alias: map[collver]string{0b00001100: "dec8_nopad_bin"}, isdefault: 0b00000000},
	1094: {alias: map[collver]string{0b00001100: "greek_nopad_bin"}, isdefault: 0b00000000},
	1095: {alias: map[collver]string{0b00001100: "hebrew_nopad_bin"}, isdefault: 0b00000000},
	1096: {alias: map[collver]string{0b00001100: "hp8_nopad_bin"}, isdefault: 0b00000000},
	1097: {alias: map[collver]string{0b00001100: "keybcs2_nopad_bin"}, isdefault: 0b00000000},
	1098: {alias: map[collver]string{0b00001100: "koi8r_nopad_bin"}, isdefault: 0b00000000},
	1099: {alias: map[collver]string{0b00001100: "koi8u_nopad_bin"}, isdefault: 0b00000000},
	1101: {alias: map[collver]string{0b00001100: "latin2_nopad_bin"}, isdefault: 0b00000000},
	1102: {alias: map[collver]string{0b00001100: "latin5_nopad_bin"}, isdefault: 0b00000000},
	1103: {alias: map[collver]string{0b00001100: "latin7_nopad_bin"}, isdefault: 0b00000000},
	1104: {alias: map[collver]string{0b00001100: "cp850_nopad_bin"}, isdefault: 0b00000000},
	1105: {alias: map[collver]string{0b00001100: "cp852_nopad_bin"}, isdefault: 0b00000000},
	1106: {alias: map[collver]string{0b00001100: "swe7_nopad_bin"}, isdefault: 0b00000000},
	1107: {alias: map[collver]string{0b00001100: "utf8_nopad_bin"}, isdefault: 0b00000000},
	1108: {alias: map[collver]string{0b00001100: "big5_nopad_bin"}, isdefault: 0b00000000},
	1109: {alias: map[collver]string{0b00001100: "euckr_nopad_bin"}, isdefault: 0b00000000},
	1110: {alias: map[collver]string{0b00001100: "gb2312_nopad_bin"}, isdefault: 0b00000000},
	1111: {alias: map[collver]string{0b00001100: "gbk_nopad_bin"}, isdefault: 0b00000000},
	1112: {alias: map[collver]string{0b00001100: "sjis_nopad_bin"}, isdefault: 0b00000000},
	1113: {alias: map[collver]string{0b00001100: "tis620_nopad_bin"}, isdefault: 0b00000000},
	1114: {alias: map[collver]string{0b00001100: "ucs2_nopad_bin"}, isdefault: 0b00000000},
	1115: {alias: map[collver]string{0b00001100: "ujis_nopad_bin"}, isdefault: 0b00000000},
	1116: {alias: map[collver]string{0b00001100: "geostd8_general_nopad_ci"}, isdefault: 0b00000000},
	1117: {alias: map[collver]string{0b00001100: "geostd8_nopad_bin"}, isdefault: 0b00000000},
	1119: {alias: map[collver]string{0b00001100: "cp932_japanese_nopad_ci"}, isdefault: 0b00000000},
	1120: {alias: map[collver]string{0b00001100: "cp932_nopad_bin"}, isdefault: 0b00000000},
	1121: {alias: map[collver]string{0b00001100: "eucjpms_japanese_nopad_ci"}, isdefault: 0b00000000},
	1122: {alias: map[collver]string{0b00001100: "eucjpms_nopad_bin"}, isdefault: 0b00000000},
	1125: {alias: map[collver]string{0b00001100: "utf16_unicode_nopad_ci"}, isdefault: 0b00000000},
	1147: {alias: map[collver]string{0b00001100: "utf16_unicode_520_nopad_ci"}, isdefault: 0b00000000},
	1152: {alias: map[collver]string{0b00001100: "ucs2_unicode_nopad_ci"}, isdefault: 0b00000000},
	1174: {alias: map[collver]string{0b00001100: "ucs2_unicode_520_nopad_ci"}, isdefault: 0b00000000},
	1184: {alias: map[collver]string{0b00001100: "utf32_unicode_nopad_ci"}, isdefault: 0b00000000},
	1206: {alias: map[collver]string{0b00001100: "utf32_unicode_520_nopad_ci"}, isdefault: 0b00000000},
	1216: {alias: map[collver]string{0b00001100: "utf8_unicode_nopad_ci"}, isdefault: 0b00000000},
	1238: {alias: map[collver]string{0b00001100: "utf8_unicode_520_nopad_ci"}, isdefault: 0b00000000},
	1248: {alias: map[collver]string{0b00001100: "utf8mb4_unicode_nopad_ci"}, isdefault: 0b00000000},
	1270: {alias: map[collver]string{0b00001100: "utf8mb4_unicode_520_nopad_ci"}, isdefault: 0b00000000},
}
