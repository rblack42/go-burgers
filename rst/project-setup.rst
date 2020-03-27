Setting Up a Go Project
#######################

..	_GOPATH:	https://golang.org/doc/gopath_code.html

In the Python world, projects are usually set up inside of a *Virtual
Environment" which serves to isolate the project and all of its dependent
packages from other projects on your system.

The Go ecosystem originally started off using a single directory, defined by a
system environment variable, to store all projects and dependent packages. If
multiple projects needed to use different versions of a single package, the
installed dependencies would be marked with a version number as part of the
directory path. 

While this seemed nice for making it possible for projects to depend on
different versions, the single GOPATH_ directory got to be something of a mess.

Go introduced the  idea of *vendoring* to deal with this issue. A project can
be set  up so that all dependent packages are downloaded into a single
**vendor** directory located inside of the project directory. This is almost
what the Python developers do, except with Python even the version of Python is
defined within the virtual environment. 

..	note::

	After years developing Python projects for my classes, I ended up with
hundreds of copies of Python on my deelopemt systems!


Setting Up the Project Directory
********************************
