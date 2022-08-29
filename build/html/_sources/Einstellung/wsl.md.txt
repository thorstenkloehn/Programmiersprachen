# Wsl

## C (Fastcgi und Nginx)
### C Beispiel Code
```
#include <fcgi_stdio.h>
#include <stdlib.h>

int main (void) {
    while (FCGI_Accept() >= 0) {
        printf("Status: 200 OK\r\n");
        printf("Content-type: text/html\r\n\r\n");
        printf("<!doctype><html><body>%s</body></html>\n",getenv("SERVER_NAME"));
    }
    return EXIT_SUCCESS;
}

```

### GCC Compiler

```  
gcc main.c -o Testseite -lfcgi



```

### Programm Starten

```

spawn-fcgi -p 8087 Testseite

```

### Nginx Einstellung

```
 sudo service nginx start
sudo nano /etc/nginx/sites-available/default



server {
    listen 80;
    server_name localhost;

    location / {
      fastcgi_pass   127.0.0.1:8087;

      fastcgi_param  GATEWAY_INTERFACE  CGI/1.1;
      fastcgi_param  SERVER_SOFTWARE    nginx;
      fastcgi_param  QUERY_STRING       $query_string;
      fastcgi_param  REQUEST_METHOD     $request_method;
      fastcgi_param  CONTENT_TYPE       $content_type;
      fastcgi_param  CONTENT_LENGTH     $content_length;
      fastcgi_param  SCRIPT_FILENAME    $document_root$fastcgi_script_name;
      fastcgi_param  SCRIPT_NAME        $fastcgi_script_name;
      fastcgi_param  REQUEST_URI        $request_uri;
      fastcgi_param  DOCUMENT_URI       $document_uri;
      fastcgi_param  DOCUMENT_ROOT      $document_root;
      fastcgi_param  SERVER_PROTOCOL    $server_protocol;
      fastcgi_param  REMOTE_ADDR        $remote_addr;
      fastcgi_param  REMOTE_PORT        $remote_port;
      fastcgi_param  SERVER_ADDR        $server_addr;
      fastcgi_param  SERVER_PORT        $server_port;
      fastcgi_param  SERVER_NAME        $server_name;
    }
  }
```

### Clion Einstellung

