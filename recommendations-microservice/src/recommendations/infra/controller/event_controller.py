import json
from infra.common.logger import logger

from application.adapter.abstract_publisher import AbstractPublisher
from application.adapter.abstract_redis_adapter import AbstractRedisAdapter
from application.adapter.abstract_grpc_adapter import AbastractGrpcAdapter
from application.dto.get_recommendation_response_dto import GetRecommendationResponseDto
from application.service.ratings_service import build_ratings
from application.service.recommendations_service import calculate_recommendations


class EventController:
    publisher: AbstractPublisher
    redis_adapter: AbstractRedisAdapter
    grpc_adapter: AbastractGrpcAdapter

    def __init__(self, publisher, redis_adapter, grpc_adapter) -> None:
        self.publisher = publisher
        self.redis_adapter = redis_adapter
        self.grpc_adapter = grpc_adapter

    def calculate_recommendation(self, __message, __channel, __properties):
        logger.info("[INFO] message")
        logger.info(__message)

        # TODO: retrieve ratings with pagination
        rates = self.grpc_adapter.list_user_rate(0, 50)
        if rates is None or len(rates) == 0:
            return

        rates_dict = build_ratings(rates)
        logger.info("rates_dict :: " + str(rates_dict))

        user_uuid = __message["user_uuid"]

        recommendations = calculate_recommendations(user_uuid, rates_dict)
        logger.info("recommendations :: " + str(recommendations))

        if recommendations is not None:
            cache_key = f"/recommendation/user/{user_uuid}"
            logger.info("cache_key :: " + cache_key)
            self.redis_adapter.set_cache(cache_key, json.dumps(recommendations))

        logger.info("[x] Done")

    def get_recommendation(self, __message, __channel, __properties):
        logger.info("[INFO] message")
        logger.info(__message)

        user_uuid = __message["user_uuid"]

        cache_key = f"/recommendation/user/{user_uuid}"
        logger.info("cache_key :: " + cache_key)

        recommendations = self.redis_adapter.get_cache(cache_key)
        logger.info("recommendations :: " + str(recommendations))

        if recommendations is None:
            recommendations = json.dumps([])

        get_recommendation_response_dto_json = GetRecommendationResponseDto(
            user_uuid, json.loads(recommendations)
        ).to_json()

        logger.info(
            "get_recommendation_response_dto_json :: "
            + str(get_recommendation_response_dto_json)
        )

        if __properties.reply_to is not None:
            reply_queue_name = __properties.reply_to
            logger.info("reply_queue_name :: " + reply_queue_name)
            self.publisher.publish_temp_queue(
                __channel, reply_queue_name, get_recommendation_response_dto_json
            )
