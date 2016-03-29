nvdcve2json
===========

Pipe-able parser from XML to JSON of the nvdcve list.

HOW DO I USE IT?
----------------

    ./nvdcve2json < nvdcve-2.0-2016.xml

    curl https://nvd.nist.gov/feeds/xml/cve/nvdcve-2.0-Recent.xml.gz | \
      gunzip - | ./nvdcve2json

    ./nvdcve2json --input nvdcve-2.0-2016.xml

    ./nvdcve2json --input nvdcve-2.0-2016.xml > nvdcve-2.0-2016.json

AUTHOR AND LICENSE
------------------
© Jose-Luis Rivas `<me@ghostbar.co>`.

This software is licensed under the MIT terms, a copy of the license can be
found in the `LICENSE` file in this repository.
