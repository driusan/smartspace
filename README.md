# Smartspace: Replace spaces with unicode codepoints.

This program takes a filename as a parameter, and outputs to STDOUT a variation of the
file with ASCII space characters replaced with more appropriate unicode codepoints, as interpreted
by historical standards that predate computers.

In particular: if the space is at the start of a sentence (comes after a period), it is replaced it with
an em-quad unicode codepoint. If it comes after a comma, semicolon, colon, exclamation, or
question mark, it is replaced with an en-quad codepoint. If the space is a space character between
words (isn't preceeded by any punctuation), it's replaced by a three-per-em space.

This is roughly the average of the standards of the historical references referenced [here.](http://heracliteanriver.com?p=324)

To see the difference, you can compare README and README.Singlespaced, and README.Doublespaced in this repository.