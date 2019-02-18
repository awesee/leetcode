#!/usr/bin/env bash

# Read from the file words.txt and output the word frequency list to stdout.

awk '{for(i=1;i<=NF;i++) a[$i]++} END {for(k in a) print k,a[k]}' words.txt | sort -k2 -nr
