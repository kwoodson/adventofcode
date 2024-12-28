#!/usr/bin/env python


def parse(test=False):
    d_str = "input.txt"
    if test:
        d_str = "test.txt"
    with open(d_str) as fd:
        data = fd.read()

    results = []
    for line in data.strip().split('\n'):
        results.append(line)
        
    return results

def left(x, y, data, search="XMAS"):
    if y - 3 >= 0:
        if data[x][y-3:y+1][::-1] == search:
            return True
    return False


def right(x, y, data, search="XMAS"):
    if y + 3 < len(data[0]):
        if data[x][y:y+4] == search:
            return True
    return False


def up(x, y, data, search="XMAS"):
    if x - 3 >= 0:
        if data[x][y] + data[x-1][y] + data[x-2][y] + data[x-3][y] == search:
            return True
    return False
    

def down(x, y, data, search="XMAS"):
    if x + 3 < len(data):
        if data[x][y] + data[x+1][y] + data[x+2][y] + data[x+3][y] == search:
            return True
    return False

def diag(x, y, data, search="XMAS"):
    count = 0
    # right - down diag
    if x + 3 < len(data) and y + 3 < len(data):
        if data[x][y] + data[x+1][y+1] + data[x+2][y+2] + data[x+3][y+3] == search:
            print(f"diag (rd) => {x},{y}")
            count += 1
    # right - up diag
    if x - 3 >= 0 and y + 3 < len(data):
        if data[x][y] + data[x-1][y+1] + data[x-2][y+2] + data[x-3][y+3] == search:
            print(f"diag (ru) => {x},{y}")
            count += 1
    # left - up diag
    if x - 3 >= 0 and y - 3 >= 0:
        if data[x][y] + data[x-1][y-1] + data[x-2][y-2] + data[x-3][y-3] == search:
            print(f"diag (lu) => {x},{y}")
            count += 1
    # left - down diag
    if x + 3 < len(data) and y - 3 >= 0:
        if data[x][y] + data[x+1][y-1] + data[x+2][y-2] + data[x+3][y-3] == search:
            print(f"diag (ld) => {x},{y}")
            count += 1
    return count

def part1():
    data = parse(False)
    #data = parse(True)

    count = 0
    x = 0
    while x < len(data):
        y = 0
        while y < len(data[0]):
            #if x == 1 and y == 4:
                #import pdb; pdb.set_trace()
                 
            if data[x][y] == "X":
                if right(x, y, data):
                    print(f"right => {x},{y}")
                    count += 1
                if left(x, y, data):
                    print(f"left => {x},{y}")
                    count += 1
                if up(x, y, data):
                    print(f"up => {x},{y}")
                    count += 1
                if down(x, y, data):
                    print(f"down => {x},{y}")
                    count += 1
                count += diag(x, y, data)
            y += 1
        x += 1

    return count

def right_down(x,y,data):
    if x + 2 < len(data) and y + 2 < len(data):
        if data[x][y] + data[x+1][y+1] + data[x+2][y+2] in ["MAS", "SAM"]:
            #print(f"diag (rd) => {x},{y}")
            return True
    return False

def right_up(x,y,data):
    if x - 2 >= 0 and y + 2 < len(data):
        if data[x][y] + data[x-1][y+1] + data[x-2][y+2] in ["MAS", "SAM"]:
            #print(f"diag (ru) => {x},{y}")
            return True
    return False

def left_up(x,y,data):
    if x - 2 >= 0 and y - 2 >= 0:
        if data[x][y] + data[x-1][y-1] + data[x-2][y-2] in ["MAS", "SAM"]:
            #print(f"diag (lu) => {x},{y}")
            return True
    return False

def left_down(x,y,data):
    if x + 2 < len(data) and y - 2 >= 0:
        if data[x][y] + data[x+1][y-1] + data[x+2][y-2] in ["MAS", "SAM"]:
            #print(f"diag (ld) => {x},{y}")
            return True

    return False

def mas_diag(x, y, data, search="MAS"):
    count = 0
    if right_down(x,y, data) and right_up(x+2,y,data):
        print(f"rd/ru {x},{y}")
        count += 1

    if left_up(x,y, data) and right_up(x,y-2,data):
        print(f"lu/ru {x},{y}")
        count += 1

    return count

def part2():
    data = parse(False)
    #data = parse(True)

    count = 0
    x = 0
    while x < len(data):
        y = 0
        while y < len(data[0]):
            if data[x][y] == "M":
                count += mas_diag(x, y, data)
            y += 1
        x += 1

    return count


def main():
    #results = part1()
    # 2623  too low
    # 2636  too low
    # 2654  correct
    results = part2()
    # 1990 correct
    print(results)



if __name__ == "__main__":
    main()

