f = open("input.txt", "r")

r = ""

file = f.read().split("\n\n")

crates = file[0]
movement = file[1]

cr: list[list[str]] = [[], [], [], [], [], [], [], [], []]

crates = crates.splitlines()

for i in range(len(crates)-2, -1, -1):
    crate = crates[i]
    for j in range(1, len(crate), 4):
        if crate[j] != " ":
            cr[int(j/4)].append(crate[j])

movements = movement.splitlines()

for mov in movements:
    elements = mov.split(" ")
    times = int(elements[1])
    fr = int(elements[3])-1
    to = int(elements[5])-1
    for _ in range(times):
        cr[to].append(cr[fr].pop())

for s in cr:
    r += s.pop()

print(r)