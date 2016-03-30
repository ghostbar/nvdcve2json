nvdcve2json
===========

Pipe-able parser from XML to JSON of the nvdcve list.

HOW DO I USE IT?
----------------
First, install it with:

    go get github.com/ghostbar/nvdcve2json

Or just go to the [releases](https://github.com/ghostbar/nvdcve2json/releases)
page and download the binary for your system.

Then, just run it like:


    $GOPATH/bin/nvdcve2json < nvdcve-2.0-2016.xml

    curl https://nvd.nist.gov/feeds/xml/cve/nvdcve-2.0-Recent.xml.gz | \
      gunzip - | $GOPATH/bin/nvdcve2json

    $GOPATH/bin/nvdcve2json --input nvdcve-2.0-2016.xml

    $GOPATH/bin/nvdcve2json --input nvdcve-2.0-2016.xml > nvdcve-2.0-2016.json

More help can be found on `$GOPATH/bin/nvdcve2json --help`.

WHAT ABOUT FILTERING STUFF OUT?
-------------------------------

You can use the flag `--filter` since `v1.0.0` to just get the CVEs you want,
like: `"cpe:/o:apple:mac_os_x"`, then `nvdcve2json` will use the logical tests
on the `vulnerable-configuration` field to determine if that `cpe` string
matches any of the CVEs and will print out just that.

Protip: you can send multiple `--filter`, like:

    curl https://nvd.nist.gov/feeds/xml/cve/nvdcve-2.0-Recent.xml.gz | \
      gunzip - | $GOPATH/bin/nvdcve2json --filter "cpe:/o:apple:mac_os_x"
      --filter "cpe:/o:microsoft:windows" > cves-for-mac-n-windows.json

AUTHOR AND LICENSE
------------------
© Jose-Luis Rivas `<me@ghostbar.co>`.

This software is licensed under the MIT terms, a copy of the license can be
found in the `LICENSE` file in this repository.
