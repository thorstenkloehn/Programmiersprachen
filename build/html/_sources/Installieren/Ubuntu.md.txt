# Ubuntu

```
sudo apt-get update
sudo apt-get upgrade
sudo apt-get install git-core
sudo apt-get install  gnupg2

```

## Nginx Installieren

```
sudo apt update
sudo apt install nginx
sudo service nginx start

```

## Programmiersprachen

### Lua Installieren

```

sudo apt install lua5.3

```

### Python

```
sudo apt install python3 python3-dev curl python-is-python3 -y
sudo apt install python3-pip -y

```

### C und C++

```

sudo apt install spawn-fcgi
sudo apt-get install libfcgi-dev
sudo apt-get install cmake gcc clang gdb build-essential

```
### Go

```
sudo mkdir /download
cd /download
sudo wget https://go.dev/dl/go1.19.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.19.linux-amd64.tar.gz

```

### Rust

```
cd /download
sudo curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh


```
