

# Goal

Simple strict go-style library for parsing command-line args, developed using TDD with 100% coverage.


## Philosophy

Things should be clearly defined, strict and unambigous, e.g.:

- only one way to do things, i.e. no alternatives like "-aux" and "-a -u -x" or "-flag" and "--flag"
- no ambiguities like "--some-list opt opt --is-this-opt-or-flag"
- no suppressed args like --name John --name Bob resulting in "Bob" with "John" dropped
- no silent unexpected tails
- support for array/slice args, e.g. --item value1 --item value2 -> {'value1', 'value2'}

### Why?

Because packages with the opposite "expect anything assume nothing" philosophy are already prolific.
 This one will follow the way described above and will never diverge. If it will ever need to,  
 it will be different package, under different name, different version etc.   
 This one is to be relied upon. Period.
