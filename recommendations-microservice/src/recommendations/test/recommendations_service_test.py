from application.service.recommendations_service import calculate_recommendations


def test_calculate_recommendations(mocker):
    rates_dict = {
        "17f906a5d400eac1c4ac9a45": {
            "621ff675dc39ed5b1bce10b0": 9.0,
            "621ff675dc39ed5b1bce10b1": 2.0,
        },
        "17f906a7a090f07ef571a723": {
            "621ff675dc39ed5b1bce10b0": 8.0,
            "621ff675dc39ed5b1bce10b1": 6.0,
            "621ff675dc39ed5b1bce10b2": 5.0,
        },
        "17f9055dc9007b0d54d83270": {
            "621ff675dc39ed5b1bce10b0": 8.0,
            "621ff675dc39ed5b1bce10b1": 6.0,
            "621ff675dc39ed5b1bce10b3": 5.0,
        },
    }

    expected_recommendations = [
        {"book_uuid": "621ff675dc39ed5b1bce10b2", "rate": 5.0},
        {"book_uuid": "621ff675dc39ed5b1bce10b3", "rate": 5.0},
    ]

    print(expected_recommendations)

    recommendations = calculate_recommendations("17f906a5d400eac1c4ac9a45", rates_dict)
    print(recommendations)

    assert recommendations == expected_recommendations
    assert len(recommendations) == 2
