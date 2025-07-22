import argparse
import os
from fontTools.ttLib import TTFont
from fontTools.subset import Subsetter, Options


def main():
    parser = argparse.ArgumentParser(
        description="Subset a TTF font to ASCII to reduce its size.")
    parser.add_argument("src", help="input font file")
    parser.add_argument("dest", help="output font file")
    parser.add_argument(
        "--extra", default="", help="additional characters to keep")
    args = parser.parse_args()

    # Load the font
    font = TTFont(args.src)

    # ASCII range plus any extra characters
    chars = ''.join(chr(i) for i in range(32, 127)) + args.extra

    options = Options()
    subsetter = Subsetter(options=options)
    subsetter.populate(text=chars)
    subsetter.subset(font)

    font.save(args.dest)

    original = os.path.getsize(args.src)
    new_size = os.path.getsize(args.dest)
    print(
        f"Wrote {args.dest} ({new_size/1024:.1f} KB, reduced from {original/1024:.1f} KB)")


if __name__ == "__main__":
    main()
