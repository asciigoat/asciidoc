/*
Attributes are dynamic data in the asciidoc templates, turning
a key (a.k.a. name) into a value, or discarding the whole line
if the reference returns undefined.

Attributes can be accessed using the following constructs.
Constructs marked with an * are extensions to the original
asciidoc language.


Simple Attribute References

{name}
	if {name} is defined, it's value is returned.
	if {name} is undefined, undefined is returned.

{name|f1[|f2...]} *
	{fn(...f2(f1({name})))} is returned

Conditional Attribute References

{name=value}
	if {name} is defined, it's value is returned.
	if {name} is undefined, the <value> expression is returned.

{name?value}
	if {name} is defined, the <value> expression is returned.
	if {name} is undefined, an empty string is returned.

{name!value}
	if {name} is defined, an empty string is returned.
	if {name} is undefined, the <value> expression is returned.

{name#value}
	if {name} is defined, the <value> expression is returned.
	if {name} is undefined, undefined is returned.

{name%value}
	if {name} is defined, undefined is returned.
	if {name} is undefined, the <value> expression is returned.

{name@regexp:value1[:value2]}
	if {name} is undefined, undefined is returned.
	if {name} matches the regexp, the <value1> expression is returned.
	if {name} doesn't match the regexp, the <value2> expression is returned, or an empty string if omitted.
	Colons in value1 or value2 need to be escaped with a backslash.

{name$regexp:value1}
	if {name} is undefined or it doesn't match the regexp, the full line is discarded.
	if {name} matches the regexp, the <value1> expression is returned.

{name$regexp::value2}
	if {name} is undefined or it doesn't match the regexp, the <value1> expression is returned.
	if {name} matches the regexp, the full line is discarded.
*/
package attribute
