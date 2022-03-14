from dataclasses import dataclass


@dataclass
class Recommendation:
    book_uuid: str
    rate: float
