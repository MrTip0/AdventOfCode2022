def compare(l1: list[list|int], l2: list[list|int])->int:
    for i in range(len(l1)):
        if len(l2) <= i:
            return -1
        if type(l1[i]) == type(l2[i]):
            if type(l1[i]) == int:
                if l1[i] > l2[i]:
                    return -1
                elif l2[i] > l1[i]:
                    return 1
            else:
                r = compare(l1[i], l2[i])
                if r != 0:
                    return r
        else:
            if type(l1[i]) == int:
                r = compare([l1[i]], l2[i])
                if r != 0:
                    return r
            elif type(l2[i]) == int:
                r = compare(l1[i], [l2[i]])
                if r != 0:
                    return r
    if len(l1) < len(l2):
        return 1
    return 0


f = open("input.txt", "r")

r = 0

pairs = f.read().strip().split("\n\n")

for i, v in enumerate(pairs):
    v = v.strip()
    if v == "":
        continue
    lines = v.splitlines()
    l1 = eval(lines[0].strip())
    l2 = eval(lines[1].strip())
    if compare(l1, l2) == 1:
        r += i+1
    
print(r)