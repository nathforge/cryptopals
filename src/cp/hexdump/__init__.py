def ascii(bytes, unprintable_char="."):
    return "".join(
        chr(byte) if 0x20 <= byte < 0x7F else unprintable_char
        for byte in bytes
    )
