#!/usr/bin/env python

from collections import Counter

def parse():
    with open('input.txt') as fd:
        data = fd.readlines()

    results = []
    for line in data:
        arr = line.strip().split()
        results.append([int(a) for a in arr])
        
    return results

def part2():
    num1, num2 = parse()

    n2 = Counter(num2)

    results = []
    for a in num1:
        results.append(a * n2[a])


    #print(results)
    print(sum(results))

# report=[7, 8, 5, 8, 10] => True
# report=[20, 21, 24, 25, 27, 29, 27]
def check_asc_recur(report, errors):
    if errors > 1:
        return False

    if len(report)-1 == 0:
        return True


    if report[0] < report[1]:
        return check_asc_recur(report[1:], errors)

    else:
        skip0 = check_asc_recur(report[1:], errors + 1)
        skip1 = check_asc_recur(report[0:1] + report[2:], errors + 1)

        if skip0 or skip1:
            return True

    return False

# [82, 83, 84, 81, 86]
# [82, 85, 83, 84, 86, 90]
# [13, 14, 13, 15, 19]
# [16, 19, 21, 25, 26]
# [20, 21, 24, 25, 27, 29, 27]
def check_order_recur(report, errors):
    if errors > 1:
        return False

    if len(report) == 1 or (len(report) == 2 and errors == 0):
        return True

    if 1 <= (report[1] - report[0] ) <= 3:
        return check_order_recur(report[1:], errors)

    # if next is larger then skip current
    # 0->1 doesn't work, skip 0


    elif len(report) >= 3 and 1 <= (report[2] - report[0] ) <= 3:
            return check_order_recur(report[0:1] + report[2:], errors + 1)

    return False
     

def check_order_2(report):
    if (report[0] > report[-1]):
        report = report[::-1]

    order = None
    #[16, 19, 21, 25, 26]
         
    order = check_asc_recur(report, 0)
    #print(f"report={report} => {order}")
    if not order:
        return False, None, report

    if report[0] == 20 and report[-1] == 27:
        import pdb; pdb.set_trace()
    distance = check_order_recur(report, 0)
    #print(f"report={report} => {distance}")
    return order, distance, report

    # if there are more than 2 duplicates, then return False
    if len(report)-len(set(report)) >= 2:
        return False

    cur = 0
    errors = 0
    while cur < len(report)-1:
        if errors > 1:
            return False

        # can we go 0 -> 1?
        if not (1 <= report[cur+1]-report[cur] <= 3):
            # we can't go from current -> current+1

            # can we go from 0 -> 2 ?
            # if cur + 2 is an option?
            if cur + 2 >= len(report):
                return False

            if not (1 <= report[cur+2]-report[cur] <= 3):
                # we can't go from 0 -> 2 so move cur to +1
                cur += 1

            # we can go from 0 -> 2
            else:
                cur += 2
                
            errors += 1
            continue

        cur += 1 

    if errors > 1:
        return False

    return True

def check_order(report):
    if not (report[0] < report[1] < report[2]):
        report = report[::-1]
    ind = 0
    for ind, val in enumerate(report):
        if ind == len(report)-1:
            break

        _next = report[ind+1]

        #print(f"{val} > {report[ind+1]}")
        if val > _next:
            return False

        result = _next-val
        #print(f"1 <= {result} <= 3 IS {1 <= result <= 3}")
        if not(1 <= result <= 3):
            return False

        ind += 1 
    return True

def is_safe(report):
    if not (report[0] < report[-1]): # incrementing
        report = report[::-1]

    for i in range(1, len(report)):
        diff = report[i] - report[i-1]
        if not 1 <= diff <= 3:
            return False

    return True

def check_is_safe(report):
    if is_safe(report):
        return True

    # skip and try again
    for i in range(len(report)):
        if is_safe(report[:i] + report[i+1:]):
            return True

    return False


def part2():
    reports = parse()
    results = 0
    #for report in reports[0:1]:
    for report in reports:
        results += check_is_safe(report)

    print(results)

def part1():
    reports = parse()
    results = []
    #for report in reports[0:1]:
    for report in reports:
        if is_safe(report):
            results.append(report)
    #print(results)
    return len(results)

def main():
    results = part1()
    print(results)
    results = part2()
    # 555 nope!
    # 560 is too low
    # 569 is right but not for us
    # 578 is incorrect
    print(results)



if __name__ == "__main__":
    main()

