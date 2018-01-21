def distance(ba1, ba2):
    ba1 = bytearray(ba1)
    ba2 = bytearray(ba2)
    if len(ba1) != len(ba2):
        raise ValueError("Mismatched lengths")
    distance = 0
    for b1, b2 in zip(ba1, ba2):
        for shift in range(8):
            if (b1 & (1 << shift)) != (b2 & (1 << shift)):
                distance += 1
    return distance
