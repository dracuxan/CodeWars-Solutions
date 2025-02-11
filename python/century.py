def century(year):
    sy = str(year)
    l = len(sy)
    if l <= 2:
        return 1
    if int(sy[l // 2 :]) > 0:
        return int(sy[: l // 2]) + 1
    return int(sy[: l // 2])


def newcentury(year):
    return -(-year // 100)


print(century(350))
print(newcentury(356))
