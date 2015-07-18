#!/bin/sh

#     This program is free software: you can redistribute it and/or modify
#     it under the terms of the GNU General Public License as published by
#     the Free Software Foundation, either version 2 of the License, or
#     (at your option) any later version.
# 
#     This program is distributed in the hope that it will be useful,
#     but WITHOUT ANY WARRANTY; without even the implied warranty of
#     MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
#     GNU General Public License for more details.
# 
#     You should have received a copy of the GNU General Public License
#     along with this program.  If not, see <http://www.gnu.org/licenses/>.

ROOT="dbus-send --system --print-reply --dest=org.freedesktop.Hal /org/freedesktop/Hal/devices/computer"
HIBERNATE="org.freedesktop.Hal.Device.SystemPowerManagement.Hibernate"
SLEEP="org.freedesktop.Hal.Device.SystemPowerManagement.Suspend int32:0"
REBOOT="org.freedesktop.Hal.Device.SystemPowerManagement.Reboot"
SHUTDOWN="org.freedesktop.Hal.Device.SystemPowerManagement.Shutdown"
LOCK="xscreensaver-command -lock"
 

CAPACITY=`cat /proc/acpi/battery/BAT0/info | grep design | head -n1 |  cut -s -f11 -d" "`

REMAINING=`cat /proc/acpi/battery/BAT0/state | grep remaining | cut -s -f8 -d " "`

REMAINING_PERC=`expr $REMAINING '*' 100 / $CAPACITY`

STATE=`cat /proc/acpi/battery/BAT0/state | grep charging | cut -s -d " " -f 12`

case $STATE in
	charging )    STATE="charging" ;;
	discharging ) STATE="discharging" ;;
    charged )     STATE="charged" ;;
	* )           STATE="unknown state" ;;
esac

ratmenu \
    -style dreary \
    "[$STATE $REMAINING_PERC%]" "echo 1 > /dev/null" \
    SLEEP "${ROOT} ${SLEEP}" \
    HIBERNATE "${ROOT} ${HIBERNATE}" \
    REBOOT "${ROOT} ${REBOOT}" \
    SHUTDOWN "${ROOT} ${SHUTDOWN}" \
    LOCK "$LOCK"

