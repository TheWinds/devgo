cd build
rm -rf ./*

gox -osarch="linux/amd64" ..
mv devgo_linux_amd64 dg
tar -czvf devgo_linux_amd64.tar.gz dg
rm dg

gox -osarch="darwin/amd64" ..
mv devgo_darwin_amd64 dg
tar -czvf devgo_darwin_amd64.tar.gz dg
rm dg