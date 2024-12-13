const std = @import("std");
// const input = @embedFile("example.txt");
const input = @embedFile("input.txt");

var gpa = std.heap.GeneralPurposeAllocator(.{}){};
const allocator = gpa.allocator();

const Grid = struct {
    g: std.ArrayList([]const u8),

    fn inBounds(self: Grid, p: Point) bool {
        return p.x >= 0 and p.y >= 0 and p.x < self.g.items[0].len and p.y < self.g.items.len;
    }

    fn Get(self: Grid, p: Point) u8 {
        return self.g.items[@intCast(p.y)][@intCast(p.x)];
    }
};

const Point = struct {
    x: i64,
    y: i64,

    fn move(self: Point, dir: Direction) Point {
        var x = self.x;
        var y = self.y;
        switch (dir) {
            Direction.up => y -= 1,
            Direction.down => y += 1,
            Direction.right => x += 1,
            Direction.left => x -= 1,
        }
        return Point{ .x = x, .y = y };
    }
};

const PointDir = struct {
    p: Point,
    d: Direction,
};

const Direction = enum(u2) {
    up,
    right,
    down,
    left,
    fn turn90(self: Direction) Direction {
        const res = @addWithOverflow(@intFromEnum(self), 1);
        return @enumFromInt(res[0]);
    }
};

const None = struct {};

var start = Point{ .x = 0, .y = 0 };

fn readGrid() !Grid {
    var it = std.mem.tokenizeScalar(u8, input, '\n');
    var g = Grid{
        .g = std.ArrayList([]const u8).init(allocator),
    };
    var y: i64 = 0;
    while (it.next()) |token| {
        try g.g.append(token);
        for (token, 0..) |b, x| {
            if (b == '^') {
                start = Point{ .x = @intCast(x), .y = y };
            }
        }
        y += 1;
    }
    return g;
}

fn findStart(g: Grid) ?Point {
    for (g.g.items, 0..) |row, y| {
        for (row, 0..) |b, x| {
            if (b == '^') {
                return Point{ .x = @intCast(x), .y = @intCast(y) };
            }
        }
    }
    return null;
}

fn fullPath(g: Grid) ![]PointDir {
    var seenDir = std.ArrayList(PointDir).init(allocator);
    var p = start;
    var currentDir: Direction = Direction.up;
    while (true) {
        try seenDir.append(PointDir{ .p = p, .d = currentDir });
        var nextp = p.move(currentDir);
        if (!g.inBounds(nextp)) {
            break;
        }
        while (g.Get(nextp) == '#') { // doesn't handle nextp being out of bounds technically
            currentDir = currentDir.turn90();
            nextp = p.move(currentDir);
        }
        p = nextp;
    }
    return seenDir.items;
}

fn causeLoop(grid: Grid, startP: PointDir, block: Point) !bool {
    var seenDir = std.AutoHashMap(PointDir, void).init(allocator);
    var p = startP.p;
    var currentDir = startP.d;
    while (true) {
        const curLocDir = PointDir{ .p = p, .d = currentDir };
        if (seenDir.contains(curLocDir)) { // been here in current direction, i.e. loop
            return true;
        }
        try seenDir.put(curLocDir, {});
        var nextp = p.move(currentDir);
        if (!grid.inBounds(nextp)) {
            return false;
        }

        while (grid.Get(nextp) == '#' or (nextp.x == block.x and nextp.y == block.y)) {
            currentDir = currentDir.turn90();
            nextp = p.move(currentDir);
        }
        p = nextp;
    }
}

fn solve1(grid: Grid) !void {
    const path = try fullPath(grid);
    var map = std.AutoHashMap(Point, void).init(allocator);
    for (path) |p| {
        try map.put(p.p, {});
    }
    std.debug.print("{d}\n", .{map.count()});
}

fn solve2(grid: Grid) !void {
    const seenDir = try fullPath(grid);
    var count: u32 = 0;
    var seen = std.AutoHashMap(Point, void).init(allocator);
    try seen.put(start, {});
    for (seenDir[0 .. seenDir.len - 1], seenDir[1..seenDir.len]) |prevp, p| {
        if (seen.contains(p.p)) { // already checked, handles start
            continue;
        }
        try seen.put(p.p, {});
        if (try causeLoop(grid, prevp, p.p)) {
            count = count + 1;
        }
    }
    std.debug.print("{d}\n", .{count});
}

pub fn main() !void {
    const grid = try readGrid();
    // try solve1(grid);
    try solve2(grid);
}
