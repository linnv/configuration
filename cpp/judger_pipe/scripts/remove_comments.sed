#format the { } in the source code
s/\(.*\)\([{}]\)\(.*\)/\1\n\2\n\3/g

#remove all the blanks in the beginning of each line
s/^[ \t]*//

#remove all c format comments from the source code
s/\/\*.*\*\///
/\/\*/, /\*\// s/.*//

#remove all c++ format comments from source code
s/\([^\/]*\)\/\{2,\}.*/\1/


