#!/usr/bin/env python

from collections import defaultdict
from dataclasses import dataclass

def parse(test=False):
    d_str = "input.txt"
    if test:
        d_str = "test.txt"
    with open(d_str) as fd:
        data = fd.read()

    results = []
    for line in data.strip().split('\n'):
        a,b = line.split(':')
        results.append((int(a),[int(x) for x in b.split()]))

    return results


def solve(nums, res, test):
    if len(nums) == 0:
        return res

    num = nums.pop(0)
    div_res = solve(nums[:], res * num, test)
    sub_res = solve(nums[:], res + num, test)

    if test in [div_res, sub_res]:
        return test

    return -1

def solve2(nums, res, test):
    if len(nums) == 0:
        return res

    num = nums.pop(0)
    div_res = solve2(nums[:], res * num, test)
    sub_res = solve2(nums[:], res + num, test)
    con_res = solve2(nums[:], int(str(res) + str(num)), test)

    if test in [div_res, sub_res, con_res]:
        return test

    return -1


def part1():
    data = parse(False)
    #data = parse(True)
    count = 0

    for prob in data:
        first = prob[1].pop(0)
        res = solve(prob[1][:], first, prob[0])
         
        if res == prob[0]:
            print(prob)
            count += prob[0]


    #print(data)
    return count

def part2():
    data = parse(False)
    #data = parse(True)
    count = 0

    for prob in data:
        first = prob[1].pop(0)
        res = solve2(prob[1][:], first, prob[0])
         
        if res == prob[0]:
            print(prob)
            count += prob[0]

    #print(data)
    return count

def main():
    #results = part1()
    # 5540634308362 is correct
    results = part2()
    print(results)



if __name__ == "__main__":
    main()

