import json
import pytest
import mock
import logging

from recommendations.domain.service.rate_service import build_rates

def test_expected_rates(mocker):
  rates = [{'user_uuid': '17f85d7f76e0935e50a78178', 'rates': [{'book_uuid': '621ff675dc39ed5b1bce10b0', 'rate': 5}, {'book_uuid': '621ff675dc39ed5b1bce10b1', 'rate': 5}], 'created_at': '2022-03-14T05:33:57.161Z', 'updated_at': '2022-03-14T05:34:13.502Z'}, {'user_uuid': '17f85d7f76e0935e50a781b9', 'rates': [{'book_uuid': '621ff675dc39ed5b1bce10b0', 'rate': 6}, {'book_uuid': '621ff675dc39ed5b1bce10b1', 'rate': 8}], 'created_at': '2022-03-14T00:33:32.509Z', 'updated_at': '2022-03-14T05:34:32.575Z'}, {'user_uuid': '17f7e9fcad70e012d21e0069', 'rates': [{'book_uuid': '621ff675dc39ed5b1bce10b0', 'rate': 6}, {'book_uuid': '621ff675dc39ed5b1bce10b1', 'rate': 3}], 'created_at': '2022-03-13T19:36:45.341Z', 'updated_at': '2022-03-14T05:39:06.214Z'}]

  expected_rates_dict = {'17f85d7f76e0935e50a78178': {'621ff675dc39ed5b1bce10b0': 5.0, '621ff675dc39ed5b1bce10b1': 5.0}, '17f85d7f76e0935e50a781b9': {'621ff675dc39ed5b1bce10b0': 6.0, '621ff675dc39ed5b1bce10b1': 8.0}, '17f7e9fcad70e012d21e0069': {'621ff675dc39ed5b1bce10b0': 6.0, '621ff675dc39ed5b1bce10b1': 3.0}}
  logging.info(expected_rates_dict)

  rates_dict = build_rates(rates)
  logging.info(rates_dict)

  assert rates_dict == expected_rates_dict
