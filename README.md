nvdcve2json
===========

Pipe-able parser from XML to JSON of the nvdcve list.

HOW DO I USE IT?
----------------
First, install it with:

    go get github.com/ghostbar/nvdcve2json

Then, just run it like:


    $GOPATH/bin/nvdcve2json < nvdcve-2.0-2016.xml

    curl https://nvd.nist.gov/feeds/xml/cve/nvdcve-2.0-Recent.xml.gz | \
      gunzip - | $GOPATH/bin/nvdcve2json

    $GOPATH/bin/nvdcve2json --input nvdcve-2.0-2016.xml

    $GOPATH/bin/nvdcve2json --input nvdcve-2.0-2016.xml > nvdcve-2.0-2016.json

More help can be found on `$GOPATH/bin/nvdcve2json --help`.

**NOTE**: filtering is not working yet, that'll be `v1.0.0` and work on it is on
the `devel` branch.

AUTHOR AND LICENSE
------------------
© Jose-Luis Rivas `<me@ghostbar.co>`.

This software is licensed under the MIT terms, a copy of the license can be
found in the `LICENSE` file in this repository.
