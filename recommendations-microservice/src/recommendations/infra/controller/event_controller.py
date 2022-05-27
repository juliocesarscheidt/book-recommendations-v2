import json
import logging

from application.adapter.abstract_publisher import AbstractPublisher
from application.adapter.abstract_redis_adapter import AbstractRedisAdapter
from application.adapter.abstract_grpc_adapter import AbastractGrpcAdapter
from application.dto.get_recommendation_response_dto import GetRecommendationResponseDto
from application.service.rate_service import build_rates
from application.service.recommendation_service import get_recommendations


class EventController(object):
    publisher: AbstractPublisher
    redis_adapter: AbstractRedisAdapter
    grpc_adapter: AbastractGrpcAdapter

    def __init__(self, publisher, redis_adapter, grpc_adapter) -> None:
        self.publisher = publisher
        self.redis_adapter = redis_adapter
        self.grpc_adapter = grpc_adapter

    def calculate_recommendation(self, __message, __channel, __properties):
        logging.info("[INFO] message")
        logging.info(__message)

        # TODO: retrieve ratings with pagination
        rates = self.grpc_adapter.list_user_rate(0, 50)
        if rates is None or len(rates) == 0:
            return

        rates_dict = build_rates(rates)
        logging.info("rates_dict :: " + str(rates_dict))

        user_uuid = __message["user_uuid"]

        recommendations = get_recommendations(user_uuid, rates_dict)
        logging.info("recommendations :: " + str(recommendations))

        if recommendations is not None:
            cache_key = f"/recommendation/user/{user_uuid}"
            logging.info("cache_key :: " + cache_key)
            self.redis_adapter.set_cache(cache_key, json.dumps(recommendations))

        logging.info("[x] Done")

    def get_recommendation(self, __message, __channel, __properties):
        logging.info("[INFO] message")
        logging.info(__message)

        user_uuid = __message["user_uuid"]

        cache_key = f"/recommendation/user/{user_uuid}"
        logging.info("cache_key :: " + cache_key)

        recommendations = self.redis_adapter.get_cache(cache_key)
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
