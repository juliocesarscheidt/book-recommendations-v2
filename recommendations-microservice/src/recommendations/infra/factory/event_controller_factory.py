from infra.adapter.publisher import Publisher
from infra.adapter.redis_adapter import RedisAdapter
from infra.adapter.grpc_adapter import GrpcAdapter
from infra.controller.event_controller import EventController


class EventControllerFactory:
    redis_conn_string: None
    grpc_conn_string: None

    def __init__(self, redis_conn_string, grpc_conn_string) -> None:
        self.redis_conn_string = redis_conn_string
        self.grpc_conn_string = grpc_conn_string

    def create_event_controller(self):
        publisher = Publisher()
        redis_adapter = RedisAdapter(self.redis_conn_string)
        grpc_adapter = GrpcAdapter(self.grpc_conn_string)

        return EventController(publisher, redis_adapter, grpc_adapter)
