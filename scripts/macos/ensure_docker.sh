if ! [ -x "$(command -v docker)" ]; then 
    brew install docker
else
    echo "Docker already installed"
fi 
if ! [ -x "$(command -v docker-compose)" ]; then 
    brew install docker-compose
else
    echo "Docker Compose already installed"
fi
