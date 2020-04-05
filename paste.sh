#! /usr/bin/env bash
# http://redsymbol.net/articles/unofficial-bash-strict-mode/

xattr paste | grep 'com.apple.quarantine' && {
    osascript -e 'display dialog "Alfred: Do you want to trust this version of paste ?" buttons {"No", "Yes"}' | grep 'returned:Yes' && xattr -d com.apple.quarantine paste
}

./paste $*

exit 0
