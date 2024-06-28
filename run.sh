#/bin/bash
export CGO_CFLAGS="-O -D__BLST_PORTABLE__"
export CGO_CFLAGS_ALLOW="-O -D__BLST_PORTABLE__"
export LD_LIBRARY_PATH=.
rm -rf /home/bot/.tendermint/*
rm -rf contractdb*
rm -rf accdb*
rm -rf condb
rm -rf badg*
rm -rf data
#cp -r /home/userland/tm/* /home/userland/.tendermint/
./tendermint init
#./kvstore init
./kvstore -config ../../../.tendermint/config/config.toml
