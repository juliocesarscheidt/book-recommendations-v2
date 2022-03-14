class Commands:
    CALCULATE: str
    GET: str

    def __init__(self) -> None:
        self.CALCULATE = "recommendations.calculate"
        self.GET = "recommendations.get"
