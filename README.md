Identify fields in MS Word (97-2003) documents. Reports names of fields in any of the sections of the document (body, header/footer etc.).

For more information about how this works, see the [MS-DOC Spec](https://msdn.microsoft.com/en-us/library/office/cc313153%28v=office.12%29.aspx). The relevant bit is the FibRgFcLcb97 section of the FIB (in the WordDocument stream). This part of the FIB has offsets and sizes for a bunch of different components of Word Docs. The field data itself is stored in the Table stream.

Examples:

    ./doctool test.doc
 
 Install with `go get` and compile. 
