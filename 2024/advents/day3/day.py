#!/usr/bin/env python

from collections import Counter
import re

re_mul = re.compile(r'mul\([0-9]{1,3},[0-9]{1,3}\)')

def parse():
    with open('input.txt') as fd:
        data = fd.readlines()

    results = []
    for line in data:
        results.extend(re_mul.findall(line))
        
    return results

def parse_mul(inp):
    inp = inp[4:-1]
    a,b = inp.split(',')
    return int(a),int(b)

def part1():
    muls = parse()
    results = []
    for mul in muls:
        a,b = parse_mul(mul)
        results.append(a*b)

    return sum(results)

def parse_line(line):
    results = ""
    cur = 0
    capture = True
    while cur < len(line)-1:
        if line[cur:cur+7] == "don't()":
            capture = False
            cur += 7
        elif line[cur:cur+4] == 'do()':
            capture = True
            cur += 4

        if capture:
            results += line[cur]

        cur += 1

    print(results)
    import pdb; pdb.set_trace()
     

    return results

def smart_parse():
    with open('input.txt') as fd:
        data = fd.readlines()

    results = []
    all_lines = []
    for line in data:
        all_lines.extend(line)
    all_mul = parse_line(''.join(all_lines))
    results.extend(re_mul.findall(all_mul))
        
    return results

def part2():
    muls = smart_parse()
    results = []
    for mul in muls:
        a,b = parse_mul(mul)
        results.append(a*b)

    return sum(results)


def main():
    #results = part1()
    # 184122457
    results = part2()
    # 108010220 is too high
    # 108609098
    # 107862689

    print(results)



if __name__ == "__main__":
    main()

