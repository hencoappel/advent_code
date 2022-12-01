from itertools import combinations
from collections import Counter

def parse(l):
    return l.strip()

with open("in.txt") as f:
    combos = combinations([parse(l) for l in f.readlines()], 2)
    for s1, s2 in combos:
        reduce(lambda c1, c2, , zip(s1, s2))
        diff = list((c1 - c2).values())
        if diff == [1, 1]:
            print(c1 & c2)
