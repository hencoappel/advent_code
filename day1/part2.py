from itertools import cycle

def parse(l):
    return int(l.strip())

with open("in.txt") as f:
    seen = set()
    total = 0
    seen.add(total)
    vals = [parse(l) for l in f.readlines()]
    for val in cycle(vals):
        total += val
        if total in seen:
            print("Found")
            print(total)
            break
        seen.add(total)
