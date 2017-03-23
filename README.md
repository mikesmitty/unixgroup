# unixgroup
Group membership checking tool, inspired by Jan Wolter's pwauth/unixgroup tools

### Usage

    $ groups michael
    michael users

Set the USER env variable to the user to be checked, and GROUP to a space-delimited list of groups to check against

    $ export USER="michael"
    $ export GROUP="testgroup users wheel"



    $ unixgroup
    $ echo $?
    0

Since the michael user account is in the "users" group it returns status 0

### Exit Statuses

* 0: User exists, and is a member of one of the specified groups
* 1: User exists, but is not a member of any of the groups
* 2: User does not exist (or primary group id is not numeric, e.g. Windows SID)
* 3: USER environment variable is invalid
* 4: GROUP environment variable is invalid
