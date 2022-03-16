from math import sqrt


def euclidean_distance(item1, item2, base):
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


def get_recommendations(user, base, k=10):
    totals = {}
    sum_similarities = {}

    other_users = [other for other in base if other != user]

    # searching other users different from current user
    for other in other_users:
        similarity = euclidean_distance(user, other, base)

        # if there is any similarity just skip
        if similarity <= 0:
            continue

        for item in base[other]:
            # check if the book is not on the list of books from current user
            if item not in base[user]:
                totals.setdefault(item, 0)
                totals[item] += base[other][item] * similarity

                sum_similarities.setdefault(item, 0)
                sum_similarities[item] += similarity

    recommendations = [
        {"book_uuid": item, "rate": (total / sum_similarities[item])}
        for item, total in totals.items()
    ]
    recommendations_sorted = sorted(
        recommendations, key=lambda o: o["rate"], reverse=True
    )

    return recommendations_sorted[0:k]
