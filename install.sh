os_linux_bin="https://gitee.com/pangbei/devgo/attach_files/754767/download/devgo_linux_amd64.tar.gz"
os_darwin_bin="https://gitee.com/pangbei/devgo/attach_files/754766/download/devgo_darwin_amd64.tar.gz"

bin=""

if [[ "$OSTYPE" == "linux-gnu"* ]]; then
        bin=$os_linux_bin
elif [[ "$OSTYPE" == "darwin"* ]]; then
        bin=$os_darwin_bin
else
        echo "os type not support: $OSTYPE"
        exit 1
fi

curl -L -o devgo.tar.gz "${bin}"

tar -zxvf devgo.tar.gz
sudo chmod +x dg
sudo mv dg /usr/local/bin
rm devgo.tar.gz