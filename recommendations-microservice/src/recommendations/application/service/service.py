import json
import logging
import requests

from application.adapter.abstract_publisher import AbstractPublisher
from application.adapter.abstract_cache import AbstractCache
from application.dto.get_recommendation_response_dto import GetRecommendationResponseDto
from application.service.rate_service import build_rates
from application.service.recommendation_service import get_recommendations


def make_request(url, token=None, method="GET"):
    auth = None
    headers = {
        "Content-Type": "application/json;charset=UTF-8",
    }
    if token is not None:
        headers.update({"Authorization": token})
    response = requests.request(method, url, headers=headers, auth=auth,)
    if not (response.status_code >= 200 and response.status_code <= 299):
        raise Exception("[ERROR] Request Error" + str(response.text))
    return response.json()


class Service(object):
    publisher: AbstractPublisher
    cache: AbstractCache
    api_gateway_uri: None

    def __init__(self, publisher, cache, api_gateway_uri) -> None:
        self.publisher = publisher
        self.cache = cache
        self.api_gateway_uri = api_gateway_uri

    def calculate_rates(self, __message, __channel, __properties):
        logging.info("[INFO] message")
        logging.info(__message)

        # list rates from User/Rate service using HTTP
        uri = f"{self.api_gateway_uri}/v1/user/rate?page=0&size=50"

        rates = make_request(uri)

        if rates is None or "data" not in rates:
            return

        rates_dict = build_rates(rates["data"])
        logging.info("rates_dict :: " + str(rates_dict))

        user_uuid = __message["user_uuid"]

        recommendations = get_recommendations(user_uuid, rates_dict)
        logging.info("recommendations :: " + str(recommendations))

        if len(recommendations) > 0:
            cache_key = f"/recommendation/user/{user_uuid}"
            logging.info("cache_key :: " + cache_key)
            self.cache.set_cache(cache_key, json.dumps(recommendations))

        logging.info("[x] Done")

    def get_recommendation(self, __message, __channel, __properties):
        logging.info("[INFO] message")
        logging.info(__message)

        user_uuid = __message["user_uuid"]

        cache_key = f"/recommendation/user/{user_uuid}"
        logging.info("cache_key :: " + cache_key)

        recommendations = self.cache.get_cache(cache_key)
        logging.info("recommendations :: " + str(recommendations))

        if recommendations is None:
            recommendations = json.dumps([])

        get_recommendation_response_dto_json = GetRecommendationResponseDto(
            user_uuid, json.loads(recommendations)
        ).to_json()

        logging.info(
            "get_recommendation_response_dto_json :: "
            + str(get_recommendation_response_dto_json)
        )

        if __properties.reply_to is not None:
            reply_queue_name = __properties.reply_to
            logging.info("reply_queue_name :: " + reply_queue_name)
            self.publisher.publish_temp_queue(
                __channel, reply_queue_name, get_recommendation_response_dto_json
            )
