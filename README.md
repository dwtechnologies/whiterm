# whiterm

The whiterm package will remove all trailing and ending whitespaces (such as space, tabs)from a String.

It honors quoted texts and supports "\" -escaped characters.

Everything after "#" will be treated as a comment and will be ignored.

## Functions

### Remove

Remove removes all spaces and tabs from String s.
It will also remove everything regarded as a comment from s (#).
It supports  "\"-escaping and doesn't remove anything inside of quotes.
It returns the formated String.