def parse(l):
    return int(l.strip())
with open("in.txt") as f:
    print(sum(parse(l) for l in f.readlines()))
