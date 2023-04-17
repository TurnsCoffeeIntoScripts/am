# Init command that will create the initial .am file in the $HOME directory and add the 'source ~/.am' line in the .bashrc/.profile/etc...
#   If it already exists, it will print a warning message and do nothing
am init

# Prompt the user to overwrite the .am file with a default/blank version
am reinit

# Following modification to the .am file (either manually of via am command)
# this will re-source the .bashrc/.profile/etc...
am source

# This will check for inconsistencies between the .bashrc/.profile/etc. and the am.yaml file
# If one or more are found, user should be prompted for action(s)
am verify

# Base 'new' command. Will somehow ask user for name/value/param?/etc...
# TODO (prompt, editor, etc... ???)
am new

# Show the details of the command named "command" (name, value, params, etc...)
am show "command"



# TODO vvvvv
am new "name" "value"
am new "name" "value" "category"

am edit "command"

am list
am list category "category"
am list command "command"

am delete "name"