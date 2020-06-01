``testfinder`` outputs a list of Python test cases and test functions
found from the requested directory.

I feed ``testfinder`` output to `fzf`_ to build auto-completion for a
test runner:

.. image:: https://raw.githubusercontent.com/dbaty/testfinder/master/docs/demo.svg
   :width: 100%

The demo above is me typing ``pytest`` followed by ``<Maj>-<Tab>``
(which is my configured key for "advanced" auto-completion) and then
typing characters of the test function I am looking for.

``testfinder`` is fast enough for me: the first list of unfiltered
suggestions appear almost instantaneously. Then ``fzf`` does its
magic, in an even more instantaneous fashion.

On a Python project with 477 test files amongst 995 files in the tests
directory, with almost 5000 test cases and functions, ``testfinder``
takes 10ms. If it's slower for you, you're eligible for a refund.

The latest binary is at `<https://github.com/dbaty/testfinder/releases>`_.


.. _fzf: https://github.com/junegunn/fzf


Example
=======

.. code:: bash

    $ testfinder
    tests/tests.py::TestClassWithMethods::test_method1
    tests/tests.py::TestClassWithMethods::test_method2
    tests/tests.py::test_func

Command-line options:

FIXME: make it configurable: starting directory; filename patterns


Installation
============

FIXME

The latest version for linux/amd64 can be found at `https://github.com/dbaty/testfinder/releases`_.
It has been built with ``make build``.

Alternatively, you may build the sources yourself:

.. code-block:: console

    $ go get https://github.com/dbaty/testfinder
    $ $GOPATH/bin/testfinder -v
    0.1


Usage for auto-completion
=========================

To use with ``fzf`` on ``pytest``, add this in ``.zshrc`` (or adapt
for your shell):

.. code-block:: shell

    _fzf_complete_pytest() {
        _fzf_complete "--multi --reverse" "$@" < <(testfinder)
    }


Status, limitations, future
===========================

It's tailored and works for Python code only for now. The output is
compatible with ``pytest``. File parsing is very simple ("fragile" is
another word that comes to mind), and yet it works surprisingly well
in standard cases.

It's my first program in Go. I skipped "Hello world". Maybe I should
not have. If it looks too much like a Python programmer struggling to
write Go, feel free to educate me. Pull requests are welcome.

Future plans:

- handle other programming languages (not planned yet);
- pivot, disrupt an industry and take over the world (ditto).
