#!/usr/bin/env python

def Interval(object):
    def __init__9

def parse():
    with open('input.txt') as fd:
        data = fd.readlines()

    intervals = []
    for line in data:
        line = line.strip()
        sacks.append(Sack(line[:int(len(line)/2)], line[int(len(line)/2):]))
    return sacks

def main():
    sacks = parse()
    score = compare_sacks(sacks)
    print(score)

    sacks = parse2()
    score = compare_sacks_2(sacks)
    print(score)

    #score = play_rps_2(games)
    #print(score)

if __name__ == "__main__":
    main()

