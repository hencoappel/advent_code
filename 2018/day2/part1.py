from itertools import cycle
from collections import Counter

def parse(l):
    c = set(Counter(l.strip()).values())
    return (2 in c, 3 in c)

with open("in.txt") as f:
    vals = (parse(l) for l in f.readlines())
    twos = 0
    threes = 0
    for val in vals:
        twos += val[0]
        threes += val[1]
    print(twos * threes)
