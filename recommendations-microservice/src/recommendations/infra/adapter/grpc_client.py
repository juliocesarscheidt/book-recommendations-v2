import grpc

import pb.user_pb2 as user_pb2
import pb.user_pb2_grpc as user_pb2_grpc


class GrpcClient:
    channel: None
    stub_user_service: None
    stub_user_rate_service: None

    def __init__(self, grpc_conn_string) -> None:
        self.channel = grpc.insecure_channel(grpc_conn_string)
        try:
            grpc.channel_ready_future(self.channel).result(timeout=30)
            self.stub_user_service = user_pb2_grpc.UserServiceStub(self.channel)
            self.stub_user_rate_service = user_pb2_grpc.UserRateServiceStub(
                self.channel
            )

        except grpc.FutureTimeoutError as e:
            print(e)
        except Exception as e:
            print(e)

    def list_user_rate(self, page=0, size=50):
        try:
            response = self.stub_user_rate_service.ListUserRate(
                user_pb2.ListUserRateRequest(page=page, size=size)
            )
            user_rates = []
            for rate in response.userRates:
                user_rates.append(
                    {
                        "user_uuid": rate.user_uuid,
                        "rates": [
                            {"book_uuid": r.book_uuid, "rate": r.rate}
                            for r in rate.rates
                        ],
                        "created_at": rate.created_at,
                        "updated_at": rate.updated_at,
                    }
                )
            return user_rates

        except Exception as e:
            print(e)
            return None
