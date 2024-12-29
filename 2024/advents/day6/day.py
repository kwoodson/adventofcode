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
    start_pos = -1,-1
    for ind, line in enumerate(data.strip().split('\n')):
        if '^' in line:
            start_pos = (ind, line.find('^'))
        results.append(line)
        
    return results, start_pos

@dataclass
class Location():
    x: int
    y: int


# up = 0 (-1, 0)
# right = 1  (0, 1)
# down = 2  (-1, 0)
# left = left (0, -1)
directions = {0: (-1, 0), 1: (0, 1), 2: (1, 0), 3: (0, -1)}

class Puzzle():
    def __init__(self, loc, area, direction=0):
        self._loc = loc
        self._dir = direction
        self._area = area
        self._moves = set()
        self._moves.add((loc.x, loc.y, self._dir))
        #self._moves = []
        #self._moves.append((loc.x, loc.y))
        self._cycles = []

    @property
    def location(self):
        return self._loc

    @property
    def direction(self):
        return self._dir

    @property
    def area(self):
        return self._area

    def rotate(self):
        self._dir = (self._dir + 1) % 4
        #return directions[self._dir]

    @property
    def moves(self):
        return self._moves

    def still_inside(self, x=None, y=None):
        if not x:
            x = self._loc.x
         
        if not y:
            y = self._loc.y
           
        if (0 <= x <= len(self._area) - 1) and (0 <= y <= len(self._area[0])-1):
            return True

        return False

    def move(self):
        dir_x, dir_y = directions[self._dir]

        if self.still_inside(self._loc.x + dir_x, self._loc.y + dir_y):
            # we are still inside, now check the path is #
            if self._area[self._loc.x + dir_x][self._loc.y + dir_y] == "#":
                self.rotate()
                return False

            # still inside and path is '.'
            self._loc.x += dir_x
            self._loc.y += dir_y

            #if (self._loc.x, self._loc.y) not in self._moves:
                #self._moves.append((self._loc.x, self._loc.y))

            if (self._loc.x, self._loc.y) not in self._moves:
                self._moves.add((self._loc.x, self._loc.y))
            return True

        # out of bounds, move and exit
        self._loc.x += dir_x
        self._loc.y += dir_y
        return False

    def move1(self):
        dir_x, dir_y = directions[self._dir]

        if self.still_inside(self._loc.x + dir_x, self._loc.y + dir_y):
            if (self._loc.x + dir_x, self._loc.y + dir_y, self._dir) in self._moves:
                return True

            # we are still inside, now check the path is #
            if self._area[self._loc.x + dir_x][self._loc.y + dir_y] == "#":
                self.rotate()
                return False

            # still inside and path is '.'
            self._loc.x += dir_x
            self._loc.y += dir_y
            self._moves.add((self._loc.x, self._loc.y, self._dir))

            #if (self._loc.x, self._loc.y) not in self._moves:
                #self._moves.append((self._loc.x, self._loc.y))

            return False

        # out of bounds, move and exit
        self._loc.x += dir_x
        self._loc.y += dir_y
        return False

def part1():
    area, cur_pos = parse(False)
    #area, cur_pos = parse(True)

    p = Puzzle(Location(cur_pos[0], cur_pos[1]), area)

    count = 0
    print(p.location)
     
    while p.still_inside():
        p.move()

    print(p.moves)
    print(len(p.moves))
    return len(p.moves)


def part2():
    area, cur_pos = parse(False)
    #area, cur_pos = parse(True)

    count = 0

    for i in range(len(area)):
        for y in range(len(area[0])):
            p = Puzzle(Location(cur_pos[0], cur_pos[1]), area)
            if p.area[i][y] == "#":
                continue
            store = p.area[i][y]
            p.area[i] = p.area[i][:y] + '#' + p.area[i][y+1:]
            while p.still_inside():
                results = p.move1()
                if results == True:
                    count += 1
                    break

            # Exited area or hit loop
            p.area[i] = p.area[i][:y] + store + p.area[i][y+1:]
             

    print(count)
    return count

def main():
    #results = part1()
    # 5443 is too low
    # 6039 is too high
    results = part2()
    # 2154 is too high
    # 1946 is correct
    # 103.56s :( but right!
    print(results)



if __name__ == "__main__":
    main()

