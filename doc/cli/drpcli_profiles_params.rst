drpcli profiles params
======================

Gets/sets all parameters for the profile

Synopsis
--------

A helper function to return all or set all the parameters on the profile

::

    drpcli profiles params [id] [json] [flags]

Options
-------

::

      -h, --help   help for params

Options inherited from parent commands
--------------------------------------

::

      -d, --debug             Whether the CLI should run in debug mode
      -E, --endpoint string   The Digital Rebar Provision API endpoint to talk to (default "https://127.0.0.1:8092")
      -f, --force             When needed, attempt to force the operation - used on some update/patch calls
      -F, --format string     The serialzation we expect for output.  Can be "json" or "yaml" (default "json")
      -P, --password string   password of the Digital Rebar Provision user (default "r0cketsk8ts")
      -T, --token string      token of the Digital Rebar Provision access
      -U, --username string   Name of the Digital Rebar Provision user to talk to (default "rocketskates")

SEE ALSO
--------

-  `drpcli profiles <drpcli_profiles.html>`__ - Access CLI commands
   relating to profiles
