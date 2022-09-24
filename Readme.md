```
sudo tljh-config set https.enabled true
sudo tljh-config set http.port 8089
sudo tljh-config set https.port 8443
sudo tljh-config set https.tls.key /etc/letsencrypt/live/ahrensburg.city/privkey.pem
sudo tljh-config set https.tls.cert /etc/letsencrypt/live/ahrensburg.city/fullchain.pem
sudo tljh-config reload proxy

systemctl start jupyterhub.service

```
