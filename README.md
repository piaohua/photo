# Photo

    This sample is a API application based on beego. It has two API func:

- /domainurl
- /mmopen

## Installation

```
cd $GOPATH/src/photo
bee run

sh ./service.sh build
```

## Usage:

```
# url
http://wx.qlogo.cn/mmopen/vi_32/EnctdgHtll59eeJIRuh56lA0MZfmMtzwfQUPyJbpX5neicxVEZNq2nj3ibp9dIMUhXIu06N3EnKLbPibmcEiabZycg/0

# url example
http://localhost:8080/domainurl/?url=http://wx.qlogo.cn/mmopen/vi_32/EnctdgHtll59eeJIRuh56lA0MZfmMtzwfQUPyJbpX5neicxVEZNq2nj3ibp9dIMUhXIu06N3EnKLbPibmcEiabZycg/0

# photo example
http://localhost:8080/mmopen/?photo=vi_32/EnctdgHtll59eeJIRuh56lA0MZfmMtzwfQUPyJbpX5neicxVEZNq2nj3ibp9dIMUhXIu06N3EnKLbPibmcEiabZycg/0.jpg

```
