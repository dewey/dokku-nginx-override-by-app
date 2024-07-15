# dokku-nginx-override-by-app

This plugin makes it possible to customize the nginx configuration for an app that is based on an existing
Docker image. This is the case if you create a new app with the Dokku [git:from-image](https://dokku.com/docs/deployment/methods/image/) command.

For regular apps it's already possible to overwrite the nginx configuration by placing a nginx.conf.sigil file in the root of your code base,
this won't work if you are using a third party image where you can't influence the code base. This is where this project comes in.

## Installation

Like every Dokku plugin just run the following command:

```
dokku plugin:install https://github.com/dewey/dokku-nginx-override-by-app.git
```

## Usage

1) With the `dokku` user navigate to `/home/dokku`, create a nginx.conf.sigil file there with the configuration you want to use and make sure it has the right permissions (`chown dokku:dokku nginx.conf.sigil`)
2) Use the app name and the path to the directory where you want to copy the config from to copy the config to the right directory: `dokku nginx-override-by-app:add <app> /home/dokku/nginx.conf.sigil`
3) Run `dokku ps:rebuild <app>` to rebuild your site with the new config
4) Done. Your app should now use your custom `nginx.conf.sigil` file.

If you ever want to tweak a config just navigate to `/var/lib/dokku/data/nginx-override-by-app` and you'll see all your configs there, in directories named after your apps.