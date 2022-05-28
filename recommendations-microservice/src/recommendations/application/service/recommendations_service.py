from application.service.distance_service import calculate_euclidean_distance


def calculate_recommendations(user, base, k=50):
    totals = {}
    sum_similarities = {}
    other_users = [other for other in base if other != user]

    # searching other users different from current user
    for other in other_users:
        similarity = calculate_euclidean_distance(user, other, base)

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
