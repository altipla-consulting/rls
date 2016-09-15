#!/bin/bash

set -eu

go get -d github.com/libgit2/git2go
cd $GOPATH/src/github.com/libgit2/git2go

rm -rf tmp
mkdir tmp
cd tmp

wget -O libgit2-${LIBGIT2_VERSION}.tar.gz https://github.com/libgit2/libgit2/archive/v${LIBGIT2_VERSION}.tar.gz
tar -xzvf libgit2-${LIBGIT2_VERSION}.tar.gz

cd libgit2-${LIBGIT2_VERSION}
mkdir build
cd build

cmake -DTHREADSAFE=ON -DBUILD_CLAR=OFF -DCMAKE_BUILD_TYPE="RelWithDebInfo" ..
make
make install

ldconfig
