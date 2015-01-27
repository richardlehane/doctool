package main

var fieldNames = map[byte]string{
	0x01: "unparseable",
	0x02: "ref - no keyword", // Not Named Specifies that the field represents a REF field where the keyword has been omitted.
	0x03: "ref",              // REF Specified in [ECMA-376] part 4, section 2.16.5.58
	0x05: "ftnref",           // FTNREF This field is identical to NOTEREF specified in [ECMA-376] part 4, section 2.16.5.47.
	0x06: "set",              // SET Specified in [ECMA-376] part 4, section 2.16.5.64.
	0x07: "if",               // IF Specified in [ECMA-376] part 4, section 2.16.5.32.
	0x08: "index",            // INDEX Specified in [ECMA-376] part 4, section 2.16.5.35.
	0x0A: "styleref",         // STYLEREF Specified in [ECMA-376] part 4, section 2.16.5.66.
	0x0C: "seq",              // SEQ Specified in [ECMA-376] part 4, section 2.16.5.63.
	0x0D: "TOC",              // TOC Specified in [ECMA-376] part 4, section 2.16.5.75.
	0x0E: "info",             // INFO Specified in [ECMA-376] part 4, section 2.16.5.36.
	0x0F: "title",            // TITLE Specified in [ECMA-376] part 4, section 2.16.5.73.
	0x10: "subject",          // SUBJECT Specified in [ECMA-376] part 4, section 2.16.5.67.
	0x11: "author",           // AUTHOR Specified in [ECMA-376] part 4, section 2.16.5.4.
	0x12: "keywords",         // KEYWORDS Specified in [ECMA-376] part 4, section 2.16.5.37.
	0x13: "comments",         // COMMENTS Specified in [ECMA-376] part 4, section 2.16.5.14.
	0x14: "last saved by",    // LASTSAVEDBY Specified in [ECMA-376] part 4, section 2.16.5.38.
	0x15: "creation date",    //CREATEDATE Specified in [ECMA-376] part 4, section 2.16.5.16.
	0x16: "save date",        // SAVEDATE Specified in [ECMA-376] part 4, section 2.16.5.60.
	0x17: "print date",       // PRINTDATE Specified in [ECMA-376] part 4, section 2.16.5.54.
	0x18: "revision number",  // REVNUM Specified in [ECMA-376] part 4, section 2.16.5.59.
	0x19: "edit time",        // EDITTIME Specified in [ECMA-376] part 4, section 2.16.5.21.
	0x1A: "number of pages",  // NUMPAGES Specified in [ECMA-376] part 4, section 2.16.5.49.
	0x1B: "number of words",  // NUMWORDS Specified in [ECMA-376] part 4, section 2.16.5.50.
	0x1C: "number of chars",  // NUMCHARS Specified in [ECMA-376] part 4, section 2.16.5.48.
	0x1D: "filename",         // FILENAME Specified in [ECMA-376] part 4, section 2.16.5.23.
	0x1E: "template",         // TEMPLATE Specified in [ECMA-376] part 4, section 2.16.5.71.
	0x1F: "date",             // DATE Specified in [ECMA-376] part 4, section 2.16.5.18.
	0x20: "time",             // TIME Specified in [ECMA-376] part 4, section 2.16.5.72.
	0x21: "page",             // PAGE Specified in [ECMA-376] part 4, section 2.16.5.51.
	0x22: "equals",           // = Specified in [ECMA-376]part 4, section 2.16.3.3.
	0x23: "quote",            // QUOTE Specified in [ECMA-376] part 4, section 2.16.5.56.
	0x24: "include",          // INCLUDE This field is identical to INCLUDETEXT specified in [ECMA-376] part 4, section 2.16.5.34.
	0x25: "pageref",          // PAGEREF Specified in [ECMA-376] part 4, section 2.16.5.52.
	0x26: "ask",              // ASK Specified in [ECMA-376] part 4, section 2.16.5.3.
	0x27: "fill in",          // FILLIN Specified in [ECMA-376] part 4, section 2.16.5.25.
	0x28: "data",             // DATA Usage: DATA datafile [headerfile] Specifies that this field SHOULD<224> redirect the mail merge data and header files to the ones specified.
	0x29: "next",             // NEXT Specified in [ECMA-376] part 4, section 2.16.5.45.
	0x2A: "next if",          // NEXTIF Specified in [ECMA-376] part 4, section 2.16.5.46.
	0x2B: "skip if",          // SKIPIF Specified in [ECMA-376] part 4, section 2.16.5.65.
	0x2C: "merge rec",        // MERGEREC Specified in [ECMA-376] part 4, section 2.16.5.43.
	0x2D: "dde",              // DDE Specified in [MS-OE376] part 2, section 1.3.2.1.
	0x2E: "dde auto",         // DDEAUTO Specified in [MS-OE376] part 2, section 1.3.2.2.
	0x2F: "glossary",         // GLOSSARY This field is identical to AUTOTEXT specified in [ECMA-376] part 4, section 2.16.5.8.
	0x30: "print",            // PRINT Specified in [ECMA-376] part 4, section 2.16.5.53.
	0x31: "eq",               // EQ Specified in [ECMA-376] part 4, section 2.16.5.22.
	0x32: "goto button",      // GOTOBUTTON Specified in [ECMA-376] part 4, section 2.16.5.29.
	0x33: "macro button",     // MACROBUTTON Specified in [ECMA-376] part 4, section 2.16.5.41.
	0x34: "auto num out",     // AUTONUMOUT Specified in [ECMA-376] part 4, section 2.16.5.7.
	0x35: "auto num gl",      // AUTONUMLGL Specified in [ECMA-376] part 4, section 2.16.5.6.
	0x36: "auto num",         // AUTONUM Specified in [ECMA-376] part 4, section 2.16.5.5.
	0x37: "import",           // IMPORT Identical to the INCLUDEPICTURE field specified in [ECMA-376] part 4, section 2.16.5.33.
	0x38: "link",             // LINK Specified in [ECMA-376] part 4, section 2.16.5.39.
	0x39: "symbol",           // SYMBOL Specified in [ECMA-376] part 4, section 2.16.5.68.
	0x3A: "embed",            // EMBED Specifies that the field represents an embedded OLE object.
	0x3B: "merge field",      // MERGEFIELD Specified in [ECMA-376] part 4, section 2.16.5.42.
	0x3C: "user name",        // USERNAME Specified in [ECMA-376] part 4, section 2.16.5.78.
	0x3D: "user initials",    // USERINITIALS Specified in [ECMA-376] part 4, section 2.16.5.77.
	0x3E: "user address",     // USERADDRESS Specified in [ECMA-376] part 4, section 2.16.5.76.
	0x3F: "barcode",          // BARCODE Specified in [ECMA-376] part 4, section 2.16.5.10.
	0x40: "doc variable",     // DOCVARIABLE Specified in [ECMA-376] part 4, section 2.16.5.20.
	0x41: "section",          // SECTION Specified in [ECMA-376] part 4, section 2.16.5.61.
	0x42: "section pages",    // SECTIONPAGES Specified in [ECMA-376] part 4, section 2.16.5.62.
	0x43: "include picture",  // INCLUDEPICTURE Specified in [ECMA-376] part 4, section 2.16.5.33.
	0x44: "include text",     // INCLUDETEXT Specified in [ECMA-376] part 4, section 2.16.5.34.
	0x45: "file size",        // FILESIZE Specified in [ECMA-376] part 4, section 2.16.5.24.
	0x46: "form text",        // FORMTEXT Specified in [ECMA-376] part 4, section 2.16.5.28.
	0x47: "form checkbox",    // FORMCHECKBOX Specified in [ECMA-376] part 4, section 2.16.5.26.
	0x48: "note ref",         // NOTEREF Specified in [ECMA-376] part 4, section 2.16.5.47.
	0x49: "TOA",              // TOA Specified in [ECMA-376] part 4, section 2.16.5.74.
	0x4B: "merge seq",        // MERGESEQ Specified in [ECMA-376] part 4, section 2.16.5.44.
	0x4F: "auto text",        // AUTOTEXT Specified in [ECMA-376] part 4, section 2.16.5.8.
	0x50: "compare",          // COMPARE Specified in [ECMA-376] part 4, section 2.16.5.15.
	0x51: "add in",           // ADDIN Specifies that the field contains data created by an add-in.
	0x53: "form dropdown",    // FORMDROPDOWN Specified in [ECMA-376] part 4, section 2.16.5.27.
	0x54: "advance",          // ADVANCE Specified in [ECMA-376] part 4, section 2.16.5.2.
	0x55: "doc property",     // DOCPROPERTY Specified in [ECMA-376] part 4, section 2.16.5.19.
	0x57: "control",          // CONTROL Specifies that the field represents an OCX control.
	0x58: "hyperlink",        // HYPERLINK Specified in [ECMA-376] part 4, section 2.16.5.31.
	0x59: "auto text list",   // AUTOTEXTLIST Specified in [ECMA-376] part 4, section 2.16.5.9.
	0x5A: "list number",      // LISTNUM Specified in [ECMA-376] part 4, section 2.16.5.40.
	0x5B: "html control",     // HTMLCONTROL Specifies the field represents an HTML control.
	0x5C: "bidi outline",     // BIDIOUTLINE Specified in [ECMA-376] part 4, section 2.16.5.12.
	0x5D: "address block",    // ADDRESSBLOCK Specified in [ECMA-376] part 4, section 2.16.5.1.
	0x5E: "greeting line",    // GREETINGLINE Specified in [ECMA-376] part 4, section 2.16.5.30.
	0x5F: "shape",            // SHAPE This field is identical to QUOTE specified in [ECMA-376] part 4, section 2.16.5.56.
}
