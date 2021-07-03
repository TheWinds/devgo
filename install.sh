version=$1
repo="TheWinds/devgo"

if [ -z "$version" ]; then
  version=$(curl --silent "https://api.github.com/repos/$repo/releases/latest" | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/')
fi

os_linux_bin="https://github.com/TheWinds/devgo/releases/download/$version/dg-$version-linux-amd64.tar.gz"
os_darwin_bin="https://github.com/TheWinds/devgo/releases/download/$version/dg-$version-darwin-amd64.tar.gz"

bin=""

if [[ "$OSTYPE" == "linux-gnu"* ]]; then
        bin=$os_linux_bin
elif [[ "$OSTYPE" == "darwin"* ]]; then
        bin=$os_darwin_bin
else
        echo "os type not support: $OSTYPE"
        exit 1
fi

echo "download $version from $bin"

curl -L -o devgo.tar.gz "${bin}"

tar -zxvf devgo.tar.gz
sudo chmod +x dg
sudo mv dg /usr/local/bin
rm devgo.tar.gz