.. _config_log_output:

Log Output
==========

Plugin Name: **LogOutput**

Logs messages to stdout using Go's `log` package. Honors hekad's global
``log_flags`` setting.

Config:

<none>

Example:

.. code-block:: ini

    [counter_output]
    type = "LogOutput"
    message_matcher = "Type == 'heka.counter-output'"
    encoder = "PayloadEncoder"
