#!/usr/bin/env python

import re

bags = {}

with open('input.txt') as fd:
    data = fd.read().split('\n')
    for line in data:
        d = line.split(' bags contain')
        bags[d[0]] = re.findall('\s(\d) (\w*\s\w*) bags?', d[1])


count = 0

def solve(color):
    print(f"Processing Color {color}")
    return 1 + sum(int(n) * solve(c) for n, c in bags[color])

#def getChildren(num, color):
#    c = []
#    for n, col in bags[color]:
#        c.append((int(n)*int(num), col))
#    return c
#
#process = [(1, "shiny gold")]
#next_up = []
#while True:
#    for n, color in process:
#        if len(bags[color]) == 0:
#            print(f"Color {color} num: {n}")
#            count += int(n)
#        else:
#            next_up.extend(getChildren(n, color))
#    if len(next_up) == 0:
#        break
#
#    process = next_up
#    next_up = []

print(solve("shiny gold"))
