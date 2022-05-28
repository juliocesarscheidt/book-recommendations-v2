from application.commands.commands import Commands
from infra.handler.event_handler import EventHandler


def test_get_command_valid(mocker):
    commands = Commands()

    handler_cmd = commands.get_command("recommendations.calculate")
    assert handler_cmd != None
    print(handler_cmd)

    handler_fn = commands.validate_command_handler(handler_cmd, EventHandler)
    assert handler_fn != None
    print(handler_fn)


def test_get_command_invalid(mocker):
    commands = Commands()

    handler_cmd = commands.get_command("recommendations.invalid")
    assert handler_cmd == None
    print(handler_cmd)

    handler_fn = commands.validate_command_handler(handler_cmd, EventHandler)
    assert handler_fn == None
    print(handler_fn)
