from dataclasses import dataclass

from application.dto.base_dto import BaseDto


@dataclass
class GetRecommendationResponseDto(BaseDto):
    user_uuid: str
    recommendations: list
