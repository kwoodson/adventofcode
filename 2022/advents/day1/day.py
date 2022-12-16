#!/usr/bin/env python


def parse():
    with open('input.txt') as fd:
        data = fd.readlines()

    counts = []
    cur = []
    for line in data:
        line = line.strip()
        if line == '':
            counts.append(cur[:])
            cur = []
            continue
        cur.append(int(line))

    return counts
            

def main():
    counts = parse()

    h = 0
    res = []
    for c in counts:
        res.append(sum(c))
        h = max(res[-1], h)

    print(h)
    print(sum(sorted(res)[-3:]))


if __name__ == "__main__":
    main()

