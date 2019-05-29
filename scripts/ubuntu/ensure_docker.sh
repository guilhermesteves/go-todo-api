if ! [ -x "$(command -v docker)" ]; then 
    sudo apt-get remove docker docker-engine docker.io -y 
    sudo apt-get install apt-transport-https ca-certificates curl software-properties-common -y 
    curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add - 
    sudo add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable edge test" 
    sudo apt-get update 
    sudo apt-get install docker-ce -y 
else
    echo "Docker already installed"
fi 
if ! [ -x "$(command -v docker-compose)" ]; then 
    sudo curl -L "https://github.com/docker/compose/releases/download/1.23.2/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose 
    sudo chmod +x /usr/local/bin/docker-compose 
    sudo ln -s /usr/local/bin/docker-compose /usr/bin/docker-compose 
else
    echo "Docker Compose already installed"
fi 
