#!/usr/bin/env python

from collections import Counter

def parse():
    with open('input.txt') as fd:
        data = fd.readlines()

    counts = []
    cur = []
    num1, num2 = [], []
    for line in data:
        one, two = line.strip().split()
        num1.append(int(one))
        num2.append(int(two))
        

    return sorted(num1), sorted(num2)
            


def part2():
    num1, num2 = parse()

    n2 = Counter(num2)

    results = []
    for a in num1:
        results.append(a * n2[a])


    #print(results)
    print(sum(results))

def part1():
    num1, num2 = parse()

    results = []
    for a,b in zip(num1, num2):
        results.append(abs(a-b))
        print(a,b,results[-1])


    #print(results)
    print(abs(sum(results)))


def main():
    part2()



if __name__ == "__main__":
    main()

