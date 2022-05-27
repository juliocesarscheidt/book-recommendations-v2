import logging


class Commands:
    commands: dict

    def __init__(self) -> None:
        self.commands = {
            "recommendations.calculate": "calculate_recommendation",
            "recommendations.get": "get_recommendation",
        }

    def get_command(self, cmd):
        return self.commands[cmd] if cmd in self.commands else None

    def validate_command_handler(self, cmd, __instance):
        if cmd is None:
            return None
        try:
            handler_fn = getattr(__instance, cmd)
            if not callable(handler_fn):
                return None
        except Exception as e:
            logging.error(e)
            return None

        return handler_fn
