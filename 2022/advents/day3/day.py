#!/usr/bin/env python

points = {}
for c in range(97, 97+26):
    points[chr(c)] = c - 96
    points[chr(c).upper()] = c - 96 + 26

class Sack(object):
    def __init__(self, s1, s2, s3=None):
        self.s1 = s1
        self.s2 = s2
        self.s3 = s3

def parse():
    with open('input.txt') as fd:
        data = fd.readlines()

    sacks = []
    for line in data:
        line = line.strip()
        sacks.append(Sack(line[:int(len(line)/2)], line[int(len(line)/2):]))
    return sacks

def parse2():

    with open('input.txt') as fd:
        data = fd.readlines()

    sacks = []
    for i, s in enumerate(data[::3]):
        sacks.append(Sack(data[i*3].strip(),
                          data[i*3+1].strip(),
                          data[i*3+2].strip()))
    return sacks

def compare_sacks(sacks):
    tot = 0
    for i, sack in enumerate(sacks):
        s1, s2 = set(sack.s1), set(sack.s2)
        r = s1 & s2
        tot += points[list(r)[0]]
        
    return tot

def compare_sacks_2(sacks):
    tot = 0
    for i, sack in enumerate(sacks):
        s1, s2, s3 = set(sack.s1), set(sack.s2), set(sack.s3)
        r = s1 & s2 & s3
        tot += points[list(r)[0]]
         
        
    return tot

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

