#!/usr/bin/expect -f

spawn passwd root
expect "New password:"
send "16313302\n"
expect "Retype new password:"
send "16313302\n"
interact
