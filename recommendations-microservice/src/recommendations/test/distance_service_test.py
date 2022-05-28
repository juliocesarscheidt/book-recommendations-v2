from application.service.distance_service import calculate_euclidean_distance


def test_calculate_euclidean_distance(mocker):
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

    distance = calculate_euclidean_distance(
        "17f906a5d400eac1c4ac9a45", "17f906a7a090f07ef571a723", rates_dict
    )

    assert distance == 4.123105625617661
