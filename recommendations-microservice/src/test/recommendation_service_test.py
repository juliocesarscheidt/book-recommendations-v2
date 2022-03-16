import json
import pytest
import mock
import logging

from recommendations.application.service.recommendation_service import (
    euclidean_distance,
    get_recommendations,
)


def test_euclidean_distance(mocker):
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

    distance = euclidean_distance(
        "17f906a5d400eac1c4ac9a45", "17f906a7a090f07ef571a723", rates_dict
    )

    assert distance == 4.123105625617661


def test_get_recommendations(mocker):
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

    logging.info(expected_recommendations)

    recommendations = get_recommendations("17f906a5d400eac1c4ac9a45", rates_dict)
    logging.info(recommendations)

    assert recommendations == expected_recommendations
    assert len(recommendations) == 2
