#!/usr/bin/env python

from collections import namedtuple

Points = {
    # A and X for rock
    # B and Y for paper
    # C and Z for scissors
    'A':  1,
    'B':  2, 
    'C':  3,
    'X':  1,
    'Y':  2, 
    'Z':  3,
}

GameRulesP2 = {
    # ties
    ('A','X'): None,
    ('B','Y'): None,
    ('C','Z'): None,
    # p1 takes rock
    ('A','Y'): True,
    ('A','Z'): False,
    # p1 takes paper
    ('B','X'): False,
    ('B','Z'): True,
    # p1 takes scissors
    ('C','X'): True,
    ('C','Y'): False,
}

LOSE_GAME = {
    'A': 'Z',
    'B': 'X',
    'C': 'Y',
}

WIN_GAME = {
    'A': 'Y',
    'B': 'Z',
    'C': 'X',
}
TIE_GAME = {
    'A': 'X',
    'B': 'Y',
    'C': 'Z',
}

WIN_OR_LOSE = {
    'X': False,
    'Y': None,
    'Z': True,
}


def part_2(p1, p2):
    # if p2 == None, then choose p1
    result = WIN_OR_LOSE[p2]
    if result:
        # We need to win
        return WIN_GAME[p1]
    
    elif result == None:
        # We need to tie
        return TIE_GAME[p1]
    # we need to lose
    return LOSE_GAME[p1]

RPS = namedtuple('RPS', ['p1', 'p2'])

def parse():
    with open('input.txt') as fd:
        data = fd.readlines()

    games = []
    for line in data:
        line = line.strip()
        p1, p2 = line.split()

        games.append(RPS(p1, p2))

    return games
            

def play_rps_2(games):
    points = 0
    for game in games:
        p2 = part_2(game.p1, game.p2)
        out = GameRulesP2[(game.p1, p2)]
        
        if out:
            # we won
            #print(f"Won: {Points[game.p1]} < {Points[game.p2]:}")
            points += 6
        elif out == None:
            # tie
            #print(f"tie: {Points[game.p1]} == {Points[game.p2]:}")
            points += 3
        else:
            # lost
            #print(f"Lost: {Points[game.p1]} > {Points[game.p2]:}")
            pass
            
        points += Points[p2]
    return points

def play_rps(games):
    points = 0
    for game in games:
        
        out = GameRulesP2[(game.p1, game.p2)]
        
        if out:
            # we won
            #print(f"Won: {Points[game.p1]} < {Points[game.p2]:}")
            points += 6
        elif out == None:
            # tie
            #print(f"tie: {Points[game.p1]} == {Points[game.p2]:}")
            points += 3
        else:
            # lost
            #print(f"Lost: {Points[game.p1]} > {Points[game.p2]:}")
            pass
            
        points += Points[game.p2]
    return points

def main():
    games = parse()

    score = play_rps(games)
    print(score)

    score = play_rps_2(games)
    print(score)

if __name__ == "__main__":
    main()

