Go bindings for libudis86
=========================

Installation
------------
go get github.com/jroimartin/udis86

Documentation
-------------
go doc github.com/jroimartin/udis86

libudis86 Installation
----------------------
	git clone git://udis86.git.sourceforge.net/gitroot/udis86/udis86
	cd udis86
	./configure --enable-shared --with-python=/usr/bin/python2 && make && make install

LAST SYNC
---------
Commit de4a93b816b2a98c5942e2d44d30e8228e6ec8e1 ("Add some more SSE4.1 insructions") from the official udis86 git repository at sourceforge

TODO
----
* Set custom syntax translators from go
