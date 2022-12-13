import functools

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

sep1 = [[2]]
sep2 = [[6]]
packets = [sep1, sep2]

lines = f.read().strip().splitlines()

for line in lines:
    line = line.strip()
    if line == "":
        continue
    
    packets.append(eval(line))

packets.sort(key=functools.cmp_to_key(compare), reverse=True)

r = packets.index(sep1)+1
r *= packets.index(sep2)+1

print(r)