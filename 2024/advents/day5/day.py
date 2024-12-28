#!/usr/bin/env python

from collections import defaultdict

def parse(test=False):
    d_str = "input.txt"
    if test:
        d_str = "test.txt"
    with open(d_str) as fd:
        data = fd.read()

    rules = defaultdict(list)
    results = []
    for line in data.strip().split('\n'):
        if line:
            if "|" in line:
                a,b = line.split("|")
                rules[a].append(b)
            else:
                results.append(line.split(","))
        
    return rules, results

# each rule for the number must be after the number
def check(rule_nums, row, start):
    for ind, r in enumerate(row):
        if r in rule_nums:
            return True
        
    return False

def verify_rules(row, rules):
    for ind, val in enumerate(row):
        if val in rules and ind >= 1:
            if check(rules[val], row[0:ind+1], ind):
                return False
        
    return True

def part1():
    #rules, results = parse(True)
    rules, results = parse(False)
    print(rules)
    count = 0
    for res in results:
        if verify_rules(res, rules):
            count += int(res[len(res)//2])

    return count


def verify_rules_2(row, rules):
    ind = 0
    swap = False
    while ind < len(row):
        if row[ind] in rules and ind >= 1:
            result = check_2(rules[row[ind]], row[0:ind], ind)
            if result >= 0:
                row[ind],row[result] = row[result],row[ind]
                print(f"swap: {row[result]} => {row[ind]}")
                swap = True
                continue

        ind += 1
        
    if swap:
        return True
    return False

def check_2(rule_nums, row, start):
    for ind, r in enumerate(row):
        if r in rule_nums:
            return ind
        
    return -1


def part2():
    rules, results = parse(False)
    #rules, results = parse(True)
    count = 0
    for res in results:
        if verify_rules_2(res, rules):
            count += int(res[len(res)//2])

    return count


def main():
    #results = part1()
    # 1990 correct
    results = part2()
    # 4884 correct
    print(results)



if __name__ == "__main__":
    main()

