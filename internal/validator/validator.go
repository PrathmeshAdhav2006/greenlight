package validator 
 
import ( 
    "regexp" 
) 

// Declare a regular expression for sanity checking the format of email addresses (we'll 
// use this later in the book). If you're interested, this regular expression pattern is 
// taken from https://html.spec.whatwg.org/#valid-e-mail-address. Note: if you're  
// reading this in PDF or EPUB format and cannot see the full pattern, please see the 
// note further down the page. 

var ( 
    EmailRX = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA
Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}
[a-zA-Z0-9])?)*$") 
)