```
docker-compose up -d
docker-compose exec mysql bash -c 'ulimit -Hn && ulimit -Sn'
mysql> -uroot -p sqlinjection 
mysql> create table products (id varchar(255), name varchar(80), price decimal(10,2), primary key (id));
docker-compose exec mysql bash
```

```
sudo dnf install mariadb-server
sudo systemctl enable mariadb
sudo systemctl start mariadb
sudo mysql_secure_installation
mysql -uroot -p 
create database sqlinjection
```

