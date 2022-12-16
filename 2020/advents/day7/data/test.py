#!/usr/bin/env python
import re

with open('input.txt', "r") as file_input:
    lines = file_input.read().splitlines()

bags = {}
bag_count = 0

for line in lines:
    colour = re.match(r"(.+?) bags contain", line)[1]  
    bags[colour] = re.findall(r"(\d+?) (.+?) bags?", line)

def has_shiny_gold(colour):
    if colour == "shiny gold": 
        return True
    else:
        return any(has_shiny_gold(c) for _, c in bags[colour] )

for bag in bags:
    if has_shiny_gold(bag):
        bag_count += 1

print("Part 1: " + str(bag_count - 1))

def count_bags(bag_type):
    return 1 + sum(int(number)*count_bags(colour) for number, colour in bags[bag_type])

#print("Part 2: " + str(count_bags("bright salmon")-1))
#print("Part 2: " + str(count_bags("striped violet")-1))
#print("Part 2: " + str(count_bags("muted aqua")-1))
print("Part 2: " + str(count_bags("shiny gold")-1))
