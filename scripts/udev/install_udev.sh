#!/bin/sh
set -e

DIR="$(cd "$(dirname "$0")" && pwd)"
RULE_FILE="99-ffbeast-hid.rules"

install -m 0644 "$DIR/$RULE_FILE" "/etc/udev/rules.d/$RULE_FILE"

udevadm control --reload-rules
udevadm trigger
