#!/usr/bin/env python3
"""Turn a nanosecond count into a unit a reader feels.

go test reports nanoseconds, which is the right unit for a benchmark and the
wrong one for an article. "2,633,461 ns" is a number nobody converts in their
head; "2.6 ms" is a number with a size. Past a second it matters even more,
because that is where a reader starts recognising the wait.

The rule is one significant unit, picked from the magnitude, alongside the raw
figure rather than replacing it - the raw figure is what makes the benchmark
reproducible.

    python3 tools/humanize.py 2633461        -> 2.6 ms
    python3 tools/humanize.py 864951 12497   -> one per line
"""
import sys

SECOND = 1_000_000_000


def human(ns):
    """One unit, chosen by magnitude. Always rounds to something readable."""
    ns = float(ns)
    if ns < 1_000:
        return "%d ns" % round(ns)
    if ns < 1_000_000:
        return "%.3g µs" % (ns / 1_000)
    if ns < SECOND:
        return "%.3g ms" % (ns / 1_000_000)

    seconds = ns / SECOND
    if seconds < 90:
        return "%.3g s" % seconds
    minutes = seconds / 60
    if minutes < 90:
        # Below ten minutes a reader still thinks in minutes and seconds.
        if minutes < 10:
            return "%dm %ds" % (int(minutes), round(seconds - int(minutes) * 60))
        return "%.3g min" % minutes
    hours = minutes / 60
    if hours < 48:
        if hours < 10:
            return "%dh %dm" % (int(hours), round(minutes - int(hours) * 60))
        return "%.3g hours" % hours
    return "%.3g days" % (hours / 24)


UNITS = [("ns", 1), ("µs", 1_000), ("ms", 1_000_000), ("s", SECOND),
         ("min", 60 * SECOND), ("hours", 3600 * SECOND)]


def human_column(values):
    """One shared unit for a whole table column.

    Per-value units read fine in a sentence and badly in a table: "865 µs"
    beside "2.63 ms" makes the reader do the conversion that the column exists
    to save. So a column picks one unit, from its largest value, and every row
    uses it.
    """
    values = [float(v) for v in values]
    smallest, biggest = min(values), max(values)

    # Pick the unit from the SMALLEST value, not the largest. Choosing by the
    # largest turned a 4,960 ns row into "0.00 ms" - a cell carrying no
    # information at all, in a table whose entire job is comparison.
    # The threshold is a tenth of a unit, not a whole one. Requiring the
    # smallest value to reach 1.0 of its unit is too conservative: a column of
    # 157,175 and 2,633,461 ns then reports microseconds, so the reader compares
    # "157 µs" with "2,633 µs" when "0.16 ms" against "2.63 ms" is the same
    # information in the unit they actually think in.
    unit, scale = UNITS[0]
    for name, size in UNITS:
        # Two conditions. At least one value has to genuinely reach the unit,
        # or [412, 980] ns becomes "0.41 µs" and loses a digit for nothing. And
        # the smallest has to reach a tenth of it, or a row renders as 0.00.
        if biggest >= size and smallest / size >= 0.1:
            unit, scale = name, size

    # A column spanning more than five orders of magnitude cannot share a unit
    # without one end turning unreadable. Fall back to per-value units there.
    if biggest / max(smallest, 1e-9) > 1e5:
        return [human(v) for v in values]

    out = []
    for v in values:
        scaled = v / scale
        if scaled >= 100:
            out.append("%s %s" % ("{:,}".format(round(scaled)), unit))
        elif scaled >= 10:
            out.append("%.1f %s" % (scaled, unit))
        else:
            out.append("%.2f %s" % (scaled, unit))
    return out


def main():
    if len(sys.argv) < 2:
        sys.exit(__doc__)
    for arg in sys.argv[1:]:
        value = arg.replace(",", "").replace("_", "")
        try:
            print("%s ns  ->  %s" % (arg, human(float(value))))
        except ValueError:
            sys.exit("not a number: %r" % arg)


if __name__ == "__main__":
    main()
