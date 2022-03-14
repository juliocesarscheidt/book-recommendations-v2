import json


class BaseDto:
    def __iter__(self):
        yield from {**self}.items()

    def __str__(self):
        return json.dumps(dict(self), ensure_ascii=False)

    def __repr__(self):
        return self.__str__()

    def to_json(self):
        return json.dumps(self.__dict__)
