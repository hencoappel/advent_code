from datetime import datetime
from enum import Enum, auto

class Event(Enum):
    WAKE = auto()
    SLEEP = auto()
    BEGIN = auto()

event_map = {"wakes": Event.WAKE, "Guard": Event.BEGIN, "falls": Event.SLEEP}

def parse(l):
    l = l.strip()
    split_line = l.split()
    date = datetime.strptime(split_line[0][1:], "%Y-%m-%d")
    time = split_line[1][:-1]
    event = event_map[split_line[2]]
    num = None
    if event == Event.BEGIN:
        num = int(split_line[3][1:])
    return (date, time, event, num)

with open("in.txt") as f:
    events = [parse(l) for l in f.readlines()]

