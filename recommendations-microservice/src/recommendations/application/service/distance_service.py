from math import sqrt


def calculate_euclidean_distance(item1, item2, base):
    # dict of similarities
    si = {}

    for item in base[item1]:
        if item in base[item2]:
            si[item] = 1

    # there are no ratings in common for these users
    if len(si) == 0:
        return 0

    # Euclidean Distance(x, y) = Sqrt( Sum( Pow( (x - y), 2 ) ) )
    result = sum(
        [
            pow(base[item1][item] - base[item2][item], 2)
            for item in base[item1]
            if item in base[item2]
        ]
    )

    dist_real = sqrt(result)
    # dist_uniform = 1 / (1 + dist_real)

    return dist_real
