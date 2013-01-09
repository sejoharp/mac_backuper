# Introduction
This script saves macos x data onto an remote time machine. 

# Destination
* check whether the macos is at home. The script continues only in this case.
* wakes the remote time machine up if it is offline
* saves the data if the time machine is available
* shuts down the remote time machine if it was offline

# Dependencies
* golang
* homebrew formula wol

# My setup
My "time machine" is a freebsd 9.1 with netatalk 2.2.4.