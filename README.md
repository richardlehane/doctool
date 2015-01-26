Read File Information Block (FIB) in MS Word (97-2003) documents to extract byte size of fields in main, header, footnote and comment sections of the document.

For more information about how this works, see the [MS-DOC Spec](https://msdn.microsoft.com/en-us/library/office/cc313153%28v=office.12%29.aspx). The relevant bit is the FibRgFcLcb97 section. This part of the FIB has offsets and sizes for a bunch of different components of Word Docs.

Examples:

    ./doctool test.doc
 
 Install with `go get` and compile. 
