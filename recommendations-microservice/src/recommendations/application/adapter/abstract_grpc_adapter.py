class AbastractGrpcAdapter:
    channel: None
    stub_user_service: None
    stub_user_rate_service: None

    def __init__(self, grpc_conn_string) -> None:
        pass

    def list_user_rate(self, page=0, size=50):
        pass
