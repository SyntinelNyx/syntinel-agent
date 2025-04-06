!#/bin/bash

echo "Installing dependencies..."

echo "Installing kopia..."

RELEASE=$(wget -q https://github.com/kopia/kopia/releases/latest -O - | grep "title>Release" | cut -d " " -f 4 | sed 's/^v//')

wget -P ./data/dependencies/ -q https://github.com/kopia/kopia/releases/download/v$RELEASE/kopia-$RELEASE-linux-x64.tar.gz

tar -xzf ./data/dependencies/kopia-$RELEASE-linux-x64.tar.gz -C ./data/dependencies/

rm -rf ./data/dependencies/kopia-$RELEASE-linux-x64*

mv ./data/dependencies/kopia /usr/bin

echo "Kopia installed successfully."


echo "Installing trivy..."

RELEASE=$(wget -q https://github.com/aquasecurity/trivy/releases/latest -O - | grep "title>Release" | cut -d " " -f 4 | sed 's/^v//')
RELEASE_Linux=${RELEASE}_Linux

wget -P ./data/dependencies/ -q https://github.com/aquasecurity/trivy/releases/download/v$RELEASE/trivy_$RELEASE_Linux-64bit.tar.gz

tar -xzf ./data/dependencies/trivy_$RELEASE_Linux-64bit.tar.gz -C ./data/dependencies/

rm -rf ./data/dependencies/trivy_$RELEASE_Linux-64bit.tar.gz
rm -rf ./data/dependencies/README.md
rm -rf ./data/dependencies/LICENSE
rm -rf ./data/dependencies/contrib

mv ./data/dependencies/trivy /usr/bin

echo "Trivy installed successfully."

echo "Caching trivy database..."

trivy fs --download-db-only

echo "Trivy database cached successfully."
