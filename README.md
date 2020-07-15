# mw2dbm

Small command line program which takes one or more mw values and converts
them to dBm.  If there's more than one value it prints out all invididual
values and sum of all mw.

This is maybe my third go program ever.  Please be nice.

## Usage:

Multiple values
```
% mw2dbm 1.619 0.028 0.024 0.023
        mw | dbm
     1.619 |  2.092
     0.028 | -15.528
     0.024 | -16.198
     0.023 | -16.383
SUM: 1.694 |  2.289
```

Single values
```
% mw2dbm 1.619
        mw | dbm
     1.619 |  2.092
```
Badly needs testing and all that.

