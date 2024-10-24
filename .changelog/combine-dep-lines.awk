# This awk script will process the list of dependency changelog entries.
# It will combine multiple bumps for the same library into a single line.
#
# The lines should have been generated by get-dep-changes.sh and relies on its formatting.
# It's assumed that these entries are sorted by library (alphabetically), then by action ("added",
# "bumped to", or "removed"), then by new version number (in version order).
#
# Example input:
# * `lib1` bumped to v1.2.3 (from v1.2.2) <link1>.
# * `lib1` bumped to v1.2.4 (from v1.2.3) <link2>.
# * `lib2` bumped to v0.4.5 (from v0.3.1) <link3>.
# Gets turned into this:
# * `lib1` bumped to v1.2.4 (from v1.2.2) (<link1>, <link2>).
# * `lib2` bumped to v0.4.5 (from v0.3.1) <link3>.
#
# If there were three (or more) lib1 lines, then, they'd all be combined similarly.
#
# This also accounts for a possible "but is still replaced by <other lib>" warning in the line, e.g.:
# * `lib1` bumped to v1.2.3 (from v1.2.2) but is still replaced by <fork1> <link1>.
# If the last line for a lib has that warning, the final line will too, but it's ignored on the other entries.
#
# It allows for the versions to include a replacement library too, e.g.
# * `<lib>` bumped to v0.50.10-pio-1 of `<fork>` (from v0.50.7-pio-1 of `<fork>`) <link>.
#
# Here's the expeced bump line format:
# * `<lib>` bumped to <version_new>[ of `<fork_new>`] (from <version_old>[ of `<fork_old>`])[ <warning>] <link>.
# Where [] denotes an optional portion.
# And <name> represents a string that we will call "name" (but doesn't necessarily reflect a variable).
# And <link> has the format "[<text>](<address>)" (the [] and () in that are the actual expected characters).

# printPrevLib prints the previous library's info and resets things.
function printPrevLib() {
    if (LibCount == 1) {
        # If there was only one line for this library, just print it without any changes.
        # The only real difference between this output and the combined one is the parenthases around the links.
        # But I just felt like it'd be best to keep the original when there's nothing being combined.
        print LibFirstLine
    } else if (LibCount > 1) {
        # If there were multiple lines for this library, create a new line for it with all the needed info.
        print "* " PrevLib " bumped to " NewVer " (from " OrigVer ")" Warning " (" Links ")."
    }
    LibCount=0
}

{
    # Expected bump line format: "* `<lib>` bumped to <new_version>[ of `<fork>`] (from <old_version>[ of `<fork>`])[ <warning>] <link>."
    if (/^\* `.*` bumped to .* \(from .*\).* \[.*[[:digit:]]+\]\(.*\)/) {
        # The library is the second thing in the line (first is the "*"). Get it including the ticks.
        lib=$2

        # The link is everything after "(from ...) " up to the last ")".
        # This also accounts for a possible warning right before the link (the [^[]* part in this Rx).
        link=$0
        sub(/^.*\(from [^)]*\)[^[]* \[/,"[",link)
        sub(/\)[^)]*$/,")",link)

        if (lib != PrevLib) {
            # This line's library is different from the previous. Print out the previous (if we have one).
            printPrevLib()

            # This lib is now the "previous" one and we restart the links list and counter.
            PrevLib=lib
            Links=link
            LibFirstLine=$0
            LibCount=1

            # The original version is everything between "(from " and the next ")" of the first line for a library.
            OrigVer=$0
            sub(/^.*\(from /,"",OrigVer)
            sub(/\).*$/,"",OrigVer)
        } else {
            # This line's library has the same library as the previous, add the link to the list and count it.
            Links=Links ", " link
            LibCount++
        }

        # The warning will be between "(from ...)" and the " [" of the link.
        # It should end up with a leading space character.
        # This resets every line so that we only get it from the last line for a library.
        Warning=""
        if (/\(from [^)]*\) .* \[/) {
            Warning=$0
            sub(/^.*\(from [^)]*\)/,"",Warning)
            sub(/ \[.*$/,"",Warning)
        }

        # The version from this line is now the most recent (so far) for the current library.
        # It's everything between "bumped to " and "(from ".
        NewVer=$0
        sub(/^.* bumped to /,"",NewVer)
        sub(/ \(from .*$/,"",NewVer)
    } else {
        # Unknown line, print out the previous entry (if there is one), then print out this line.
        printPrevLib()
        print $0
    }
}

# Make sure there isn't one left over.
END {
    printPrevLib()
}