clear_screen = "\033[2J"

def cursor(x, y):
    return "\033[{:d};{:d}H".format(x, y)
