sudo dpkg --configure -a
sudo curl -fsSL https://get.docker.com -o get-docker.sh
sudo sh get-docker.sh

docker run -i --name hello --network=host deepakputhraya/python-hello:v1.0
docker run -i --name world --network=host deepakputhraya/python-world:v1.0
docker run --name lb -v ~/nginx.conf:/etc/nginx/nginx.conf:ro --network=host -d nginx
