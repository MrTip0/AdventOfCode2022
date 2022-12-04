f = open("input.txt")

lines = f.readlines()
r = 0

for line in lines:
    line = line.strip()
    pair = line.split(",")
    ranges = []
    for range in pair:
        p = range.split("-")
        t = (int(p[0]), int(p[1]))
        ranges.append(t)
    if (ranges[0][0] >= ranges[1][0] and ranges[0][1] <= ranges[1][1]) or (ranges[1][0] >= ranges[0][0] and ranges[1][1] <= ranges[0][1]):
        r += 1

print(r)
