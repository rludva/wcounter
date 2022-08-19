# Simple Word Counter Reporter
Purpose of this fun and learn project is to practice with fun and try to implement:

```
In your favourite programming language, implement a program that reads from standard input and calculates the number of occurrences of each word read. 
cat After reading all the data, print the number of occurrences to stdout.
# example input:


# foobar
# foo
# bar
# foobar
# bar
# foobar


# example output:


# foobar: 3
# foo: 1
# bar: 2
```

To build it use `./bin/build.sh`
To use it try: `echo "Lorem ipsum, dolor sit ipsum amet." | ./counter_linux_amd64`

```
$ echo "Lorem ipsum, dolor sit ipsum amet." | ./counter_linux_amd64
v[Lorem]:1
v[ipsum]:2
v[dolor]:1
v[sit]:1
v[amet]:1
```

Additional practice:
--------------------
* Try to add failing test with characters that are common in texts like: ?,!,",'.
* Try to add failing test and make it green with ignoring upper/lower case to have "Lorem"|"lorem"|"lOrEm" as the same word.
* Try to implement this in Python/Java/C/Bash or any other language..
