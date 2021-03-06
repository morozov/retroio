package basic

// ZX Spectrum character set, a variant of ASCII-1967
var CharacterSet = map[byte]string{
	// Everything below ASCII 32
	0x00: "",
	0x01: "",
	0x02: "",
	0x03: "",
	0x04: "",   // True Video
	0x05: "",   // Inv Video
	0x06: "",   // Caps Lock / Comma
	0x07: "",   // Edit
	0x08: "",   // Left
	0x09: "",   // Right
	0x0A: "",   // Down
	0x0B: "",   // Up
	0x0C: "",   // Delete
	0x0D: "\n", // Enter
	0x0E: "",   // Extend / Number
	0x0F: "",   // Graphics
	0x10: "",   // INK
	0x11: "",   // PAPER
	0x12: "",   // FLASH
	0x13: "",   // BRIGHT
	0x14: "",   // INVERSE
	0x15: "",   // OVER
	0x16: "",   // AT
	0x17: "",   // TAB
	0x18: "",
	0x19: "",
	0x1A: "",
	0x1B: "",
	0x1C: "",
	0x1D: "",
	0x1E: "",
	0x1F: "",
	// Normal ASCII set
	0x20: " ",
	0x21: "!",
	0x22: "\"",
	0x23: "#",
	0x24: "$",
	0x25: "%",
	0x26: "&",
	0x27: "'",
	0x28: "(",
	0x29: ")",
	0x2A: "*",
	0x2B: "+",
	0x2C: ",",
	0x2D: "-",
	0x2E: ".",
	0x2F: "/",
	0x30: "0",
	0x31: "1",
	0x32: "2",
	0x33: "3",
	0x34: "4",
	0x35: "5",
	0x36: "6",
	0x37: "7",
	0x38: "8",
	0x39: "9",
	0x3A: ":",
	0x3B: ";",
	0x3C: "<",
	0x3D: "=",
	0x3E: ">",
	0x3F: "?",
	0x40: "@",
	0x41: "A",
	0x42: "B",
	0x43: "C",
	0x44: "D",
	0x45: "E",
	0x46: "F",
	0x47: "G",
	0x48: "H",
	0x49: "I",
	0x4A: "J",
	0x4B: "K",
	0x4C: "L",
	0x4D: "M",
	0x4E: "N",
	0x4F: "O",
	0x50: "P",
	0x51: "Q",
	0x52: "R",
	0x53: "S",
	0x54: "T",
	0x55: "U",
	0x56: "V",
	0x57: "W",
	0x58: "X",
	0x59: "Y",
	0x5A: "Z",
	0x5B: "[",
	0x5C: "\\",
	0x5D: "]",
	0x5E: "↑",
	0x5F: "_",
	0x60: "£",
	0x61: "a",
	0x62: "b",
	0x63: "c",
	0x64: "d",
	0x65: "e",
	0x66: "f",
	0x67: "g",
	0x68: "h",
	0x69: "i",
	0x6A: "j",
	0x6B: "k",
	0x6C: "l",
	0x6D: "m",
	0x6E: "n",
	0x6F: "o",
	0x70: "p",
	0x71: "q",
	0x72: "r",
	0x73: "s",
	0x74: "t",
	0x75: "u",
	0x76: "v",
	0x77: "w",
	0x78: "x",
	0x79: "y",
	0x7A: "z",
	0x7B: "{",
	0x7C: "|",
	0x7D: "}",
	0x7E: "~",
	0x7F: "©",
	// Block graphics without shift
	0x80: " ",
	0x81: "▝",
	0x82: "▘",
	0x83: "▀",
	0x84: "▗",
	0x85: "▐",
	0x86: "▚",
	0x87: "▜",
	// Block graphics with shift
	0x88: "▖",
	0x89: "▞",
	0x8A: "▌",
	0x8B: "▛",
	0x8C: "▄",
	0x8D: "▟",
	0x8E: "▙",
	0x8F: "█",
	// UDGs
	0x90: "Ⓐ",
	0x91: "Ⓑ",
	0x92: "Ⓒ",
	0x93: "Ⓓ",
	0x94: "Ⓔ",
	0x95: "Ⓕ",
	0x96: "Ⓖ",
	0x97: "Ⓗ",
	0x98: "Ⓘ",
	0x99: "Ⓙ",
	0x9A: "Ⓚ",
	0x9B: "Ⓛ",
	0x9C: "Ⓜ",
	0x9D: "Ⓝ",
	0x9E: "Ⓞ",
	0x9F: "Ⓟ",
	0xA0: "Ⓠ",
	0xA1: "Ⓡ",
	0xA2: "Ⓢ",
	0xA3: "Ⓣ", // SPECTRUM
	0xA4: "Ⓤ", // PLAY
	// BASIC tokens - expression
	0xA5: "RND",
	0xA6: "INKEY$",
	0xA7: "PI",

	0xA8: "FN",
	0xA9: "POINT",
	0xAA: "SCREEN$",
	0xAB: "ATTR",
	0xAC: "AT",
	0xAD: "TAB",
	0xAE: "VAL$",
	0xAF: "CODE",
	0xB0: "VAL",
	0xB1: "LEN",
	0xB2: "SIN",
	0xB3: "COS",
	0xB4: "TAN",
	0xB5: "ASN",
	0xB6: "ACS",
	0xB7: "ATN",
	0xB8: "LN",
	0xB9: "EXP",
	0xBA: "INT",
	0xBB: "SQR",
	0xBC: "SGN",
	0xBD: "ABS",
	0xBE: "PEEK",
	0xBF: "IN",
	0xC0: "USR",
	0xC1: "STR$",
	0xC2: "CHR$",
	0xC3: "NOT",
	0xC4: "BIN",
	0xC5: "OR",
	0xC6: "AND",
	0xC7: "<=",
	0xC8: ">=",
	0xC9: "<>",
	0xCA: "LINE",
	0xCB: "THEN",
	0xCC: "TO",
	0xCD: "STEP",
	// BASIC tokens - keywords
	0xCE: "DEF FN",
	0xCF: "CAT",
	0xD0: "FORMAT",
	0xD1: "MOVE",
	0xD2: "ERASE",
	0xD3: "OPEN #",
	0xD4: "CLOSE #",
	0xD5: "MERGE",
	0xD6: "VERIFY",
	0xD7: "BEEP",
	0xD8: "CIRCLE",
	0xD9: "INK",
	0xDA: "PAPER",
	0xDB: "FLASH",
	0xDC: "BRIGHT",
	0xDD: "INVERSE",
	0xDE: "OVER",
	0xDF: "OUT",
	0xE0: "LPRINT",
	0xE1: "LLIST",
	0xE2: "STOP",
	0xE3: "READ",
	0xE4: "DATA",
	0xE5: "RESTORE",
	0xE6: "NEW",
	0xE7: "BORDER",
	0xE8: "CONTINUE",
	0xE9: "DIM",
	0xEA: "REM",
	0xEB: "FOR",
	0xEC: "GO TO",
	0xED: "GO SUB",
	0xEE: "INPUT",
	0xEF: "LOAD",
	0xF0: "LIST",
	0xF1: "LET",
	0xF2: "PAUSE",
	0xF3: "NEXT",
	0xF4: "POKE",
	0xF5: "PRINT",
	0xF6: "PLOT",
	0xF7: "RUN",
	0xF8: "SAVE",
	0xF9: "RANDOMIZE",
	0xFA: "IF",
	0xFB: "CLS",
	0xFC: "DRAW",
	0xFD: "CLEAR",
	0xFE: "RETURN",
	0xFF: "COPY",
}
