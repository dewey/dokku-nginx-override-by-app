Just experimenting!

```
dokku plugin:install https://github.com/dewey/dokku-nginx-override-by-hostname.git
```

1) Copy file into dokku directory (so we have permissions)
2) Use the app name and the path to the directory where you want to copy the config from
3) dokku ps:rebuild <app> to rebuild your site with the new config

```
chown dokku:dokku /home/dokku/nginx.conf.sigil
dokku nginx-override-by-hostname:add beckkla.us /home/dokku/nginx.conf.sigil
```
