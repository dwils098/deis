:title: Configure an Application on Deis
:description: Instructions for developers using Deis to configure applications.

.. _config-application:

Configure an Application
========================
A Deis application `stores config in environment variables`_.

Configure the Application
-------------------------
Use ``deis config`` to modify environment variables for a deployed application.

.. code-block:: console

    $ deis help config
    Valid commands for config:

    config:list        list environment variables for an app
    config:set         set environment variables for an app
    config:unset       unset environment variables for an app

    Use `deis help [command]` to learn more

When config is changed, a new release is created and deployed automatically.

Attach to Backing Services
--------------------------
Deis treats backing services like databases, caches and queues as `attached resources`_.
Attachments are performed using environment variables.

For example, use ``deis config`` to set a `DATABASE_URL` that attaches
the application to an external PostgreSQL database.

.. code-block:: console

    $ deis config:set DATABASE_URL=postgres://user:pass@example.com:5432/db
    === peachy-waxworks
    DATABASE_URL: postgres://user:pass@example.com:5432/db

Detachments can be performed with ``deis config:unset``.

Add Custom Domains to the Application
-------------------------------------

You can use ``deis domains`` to add or remove custom domains to your application:

.. code-block:: console

    $ deis domains:add hello.bacongobbler.com.com
    Adding hello.bacongobbler.com to finest-woodshed... done

Once that's done, you can go into your DNS registrar and set up a CNAME from the new
appname to the old one:

.. code-block:: console

    $ dig hello.deisapp.com
    [...]
    ;; ANSWER SECTION:
    hello.bacongobbler.com.         1759    IN    CNAME    finest-woodshed.deisapp.com.
    finest-woodshed.deisapp.com.    270     IN    A        172.17.8.100

.. note::

    Setting a CNAME for your root domain can cause issues. Setting your @ record
    to be a CNAME causes all traffic to go to the other domain, including mail and the SOA
    ("start-of-authority") records. It is highly recommended that you bind a subdomain to
    an application, however you can work around this by pointing the @ record to the
    address of the load balancer (if any).

Track Changes
-------------
Each time a build or config change is made to your application, a new :ref:`release` is created.
Track changes to your application using ``deis releases``.

.. code-block:: console

    $ deis releases
    === peachy-waxworks Releases
    v4      3 minutes ago                     gabrtv deployed d3ccc05
    v3      1 hour 17 minutes ago             gabrtv added DATABASE_URL
    v2      6 hours 2 minutes ago             gabrtv deployed 7cb3321
    v1      6 hours 2 minutes ago             gabrtv deployed deis/helloworld

Rollback the Application
------------------------
Use ``deis rollback`` to revert to a previous release.

.. code-block:: console

    $ deis rollback v2
    Rolled back to v2

    $ deis releases
    === folksy-offshoot Releases
    v5      Just now                          gabrtv rolled back to v2
    v4      4 minutes ago                     gabrtv deployed d3ccc05
    v3      1 hour 18 minutes ago             gabrtv added DATABASE_URL
    v2      6 hours 2 minutes ago             gabrtv deployed 7cb3321
    v1      6 hours 3 minutes ago             gabrtv deployed deis/helloworld

.. note::

    All releases (including rollbacks) append to the release ledger.


.. _`stores config in environment variables`: http://12factor.net/config
.. _`attached resources`: http://12factor.net/backing-services
