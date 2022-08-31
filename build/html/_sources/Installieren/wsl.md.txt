# Wsl Installieren

```
wsl --install
wsl --list --online
wsl --install -d <DistroName>

```

## Ubuntu Installieren

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
### Dotnet

* [Dotnet](https://docs.microsoft.com/de-de/dotnet/core/install/linux-ubuntu)

### PHP

```
sudo apt-get install php  openssl php-imagick php-common php-curl php-gd php-imap php-intl php-json php-ldap php-mbstring php-ssh2 php-xml php-zip unzip -y
sudo apt-get install php-fpm php-cli  composer  -y




