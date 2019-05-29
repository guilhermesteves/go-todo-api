sudo apt-get install jq -y
download_url=$(curl -s https://api.github.com/repos/go-swagger/go-swagger/releases/latest | \
  jq -r '.assets[] | select(.name | contains("'"$(uname | tr '[:upper:]' '[:lower:]')"'_amd64")) | .browser_download_url')
sudo curl -o /usr/local/bin/swagger -L'#' "$download_url"
sudo chmod +x /usr/local/bin/swagger