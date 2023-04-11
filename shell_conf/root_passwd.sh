#!/usr/bin/expect -f

spawn sudo passwd root
expect "password for ubuntu:"
send "ubuntu\n"
expect "New password:"
send "16313302\n"
expect "Retype new password:"
send "16313302\n"
interact
